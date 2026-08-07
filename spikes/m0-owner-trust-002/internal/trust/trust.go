package trust

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hkdf"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/crypto/argon2"
	_ "modernc.org/sqlite"
)

const (
	rootEnvelopeVersion = 1
	anchorVersion       = 1
	bundleVersion       = 1
	kdfMemoryKiB        = 64 * 1024
	kdfIterations       = 3
	kdfParallelism      = 4
	kdfKeyBytes         = 32
	orkBytes            = 32
	saltBytes           = 16
)

var (
	ErrUnlockFailed                = errors.New("owner root unlock failed")
	ErrMissingRoot                 = errors.New("wrapped owner root is missing")
	ErrMissingAnchor               = errors.New("owner trust anchor is missing")
	ErrInvalidDBMAC                = errors.New("governing state MAC is invalid")
	ErrInvalidAnchorMAC            = errors.New("owner trust anchor MAC is invalid")
	ErrStateRollback               = errors.New("governing state rollback detected")
	ErrAnchorLag                   = errors.New("owner trust anchor lags authenticated governing state")
	ErrTimeUntrusted               = errors.New("wall time is behind authenticated high-water")
	ErrRevalidationRequired        = errors.New("restored historical state requires owner revalidation")
	ErrStaleGeneration             = errors.New("stale expected generation")
	ErrOwnerAuthenticationRequired = errors.New("authenticated owner required")
	ErrRevalidationNotRequired     = errors.New("restore revalidation is not pending")
)

type Classification string

const (
	ClassNormal               Classification = "NORMAL"
	ClassStateRollback        Classification = "STATE_ROLLBACK"
	ClassAnchorLag            Classification = "ANCHOR_LAG"
	ClassInvalidDBMAC         Classification = "INVALID_DB_MAC"
	ClassInvalidAnchorMAC     Classification = "INVALID_ANCHOR_MAC"
	ClassMissingAnchor        Classification = "MISSING_ANCHOR"
	ClassMissingRoot          Classification = "MISSING_WRAPPED_ROOT"
	ClassUnlockFailed         Classification = "OWNER_UNLOCK_FAILED"
	ClassTimeUntrusted        Classification = "TIME_UNTRUSTED"
	ClassRevalidationRequired Classification = "REVALIDATION_REQUIRED"
)

type AuthorityStatus string

const (
	StatusActive  AuthorityStatus = "ACTIVE"
	StatusRevoked AuthorityStatus = "REVOKED"
)

type Actor string

const (
	ActorOwner         Actor = "AUTHENTICATED_OWNER"
	ActorRestoredGrant Actor = "RESTORED_HISTORICAL_GRANT"
)

type Layout struct {
	BaseDir           string
	DBPath            string
	RootPath          string
	AnchorPath        string
	RestoreMarkerPath string
}

func NewLayout(baseDir string) Layout {
	return Layout{
		BaseDir:           baseDir,
		DBPath:            filepath.Join(baseDir, "state", "aurora.db"),
		RootPath:          filepath.Join(baseDir, "trust", "owner-root.json"),
		AnchorPath:        filepath.Join(baseDir, "trust", "owner-trust.json"),
		RestoreMarkerPath: filepath.Join(baseDir, "trust", "restore-pending.json"),
	}
}

type OwnerSession struct {
	OwnerID string
	ork     []byte
}

type Evaluation struct {
	Classification    Classification  `json:"classification"`
	Permitting        bool            `json:"permitting"`
	Generation        int64           `json:"generation"`
	OwnerID           string          `json:"owner_id,omitempty"`
	AuthorityRevision string          `json:"authority_revision,omitempty"`
	AuthorityStatus   AuthorityStatus `json:"authority_status,omitempty"`
	ExpiresAtUnix     int64           `json:"expires_at_unix,omitempty"`
	NextSafeAction    string          `json:"next_safe_action,omitempty"`
	Diagnostic        string          `json:"diagnostic"`
}

type Mutation struct {
	ExpectedGeneration int64
	AuthorityRevision  string
	AuthorityStatus    AuthorityStatus
	ExpiresAt          time.Time
	StateSummary       string
}

type Revalidation struct {
	AuthorityRevision string
	ExpiresAt         time.Time
}

type FaultHook func(point string)

type AdvanceMetrics struct {
	DBCommitNS    int64 `json:"db_commit_ns"`
	AnchorWriteNS int64 `json:"anchor_write_ns"`
}

type rootEnvelope struct {
	Version      int    `json:"version"`
	OwnerID      string `json:"owner_id"`
	KDF          string `json:"kdf"`
	ArgonVersion int    `json:"argon_version"`
	MemoryKiB    uint32 `json:"memory_kib"`
	Iterations   uint32 `json:"iterations"`
	Parallelism  uint8  `json:"parallelism"`
	Salt         string `json:"salt_b64"`
	Cipher       string `json:"cipher"`
	Nonce        string `json:"nonce_b64"`
	WrappedORK   string `json:"wrapped_ork_b64"`
}

type anchorRecord struct {
	Version               int    `json:"version"`
	OwnerID               string `json:"owner_id"`
	Generation            int64  `json:"generation"`
	WallTimeHighWaterUnix int64  `json:"wall_time_high_water_unix"`
	MAC                   string `json:"mac_b64"`
}

type dbState struct {
	Generation        int64
	OwnerID           string
	AuthorityRevision string
	AuthorityStatus   AuthorityStatus
	ExpiresAtUnix     int64
	StateSummary      string
	MAC               []byte
}

type restoreMarker struct {
	Version          int   `json:"version"`
	SourceGeneration int64 `json:"source_generation"`
}

type bundleManifest struct {
	Version          int    `json:"version"`
	OwnerID          string `json:"owner_id"`
	SourceGeneration int64  `json:"source_generation"`
}

func Bootstrap(layout Layout, ownerID string, passphrase []byte, now, expiresAt time.Time) (*OwnerSession, error) {
	if ownerID == "" || len(passphrase) == 0 {
		return nil, errors.New("owner id and passphrase are required")
	}
	if err := os.MkdirAll(filepath.Dir(layout.RootPath), 0o700); err != nil {
		return nil, fmt.Errorf("create trust dir: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(layout.DBPath), 0o700); err != nil {
		return nil, fmt.Errorf("create state dir: %w", err)
	}
	ork := make([]byte, orkBytes)
	if _, err := io.ReadFull(rand.Reader, ork); err != nil {
		return nil, fmt.Errorf("generate owner root: %w", err)
	}
	session := &OwnerSession{OwnerID: ownerID, ork: append([]byte(nil), ork...)}
	env, err := wrapRoot(ownerID, passphrase, ork)
	clear(ork)
	if err != nil {
		return nil, err
	}
	if err := writeJSONAtomic(layout.RootPath, env, 0o600); err != nil {
		return nil, fmt.Errorf("write wrapped owner root: %w", err)
	}

	db, err := openDB(layout.DBPath, true)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	if err := initSchema(db); err != nil {
		return nil, err
	}
	state := dbState{
		Generation:        1,
		OwnerID:           ownerID,
		AuthorityRevision: "AUTH-1",
		AuthorityStatus:   StatusActive,
		ExpiresAtUnix:     expiresAt.Unix(),
		StateSummary:      "initial active authority",
	}
	if err := signState(session, &state); err != nil {
		return nil, err
	}
	tx, err := db.Begin()
	if err != nil {
		return nil, fmt.Errorf("begin bootstrap state: %w", err)
	}
	if err := insertInitialState(tx, state, now); err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("commit bootstrap state: %w", err)
	}
	anchor, err := newAnchor(session, 1, now.Unix())
	if err != nil {
		return nil, err
	}
	if _, err := writeAnchorAtomic(layout.AnchorPath, anchor, nil); err != nil {
		return nil, err
	}
	return session, nil
}

func Unlock(rootPath string, passphrase []byte) (*OwnerSession, error) {
	payload, err := os.ReadFile(rootPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrMissingRoot
		}
		return nil, fmt.Errorf("read wrapped owner root: %w", err)
	}
	var env rootEnvelope
	if err := json.Unmarshal(payload, &env); err != nil {
		return nil, ErrUnlockFailed
	}
	if env.Version != rootEnvelopeVersion || env.KDF != "Argon2id" || env.Cipher != "AES-256-GCM" || env.OwnerID == "" {
		return nil, ErrUnlockFailed
	}
	salt, err := base64.StdEncoding.DecodeString(env.Salt)
	if err != nil || len(salt) < 8 {
		return nil, ErrUnlockFailed
	}
	nonce, err := base64.StdEncoding.DecodeString(env.Nonce)
	if err != nil {
		return nil, ErrUnlockFailed
	}
	wrapped, err := base64.StdEncoding.DecodeString(env.WrappedORK)
	if err != nil {
		return nil, ErrUnlockFailed
	}
	kek := argon2.IDKey(passphrase, salt, env.Iterations, env.MemoryKiB, env.Parallelism, kdfKeyBytes)
	defer clear(kek)
	block, err := aes.NewCipher(kek)
	if err != nil {
		return nil, ErrUnlockFailed
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil || len(nonce) != gcm.NonceSize() {
		return nil, ErrUnlockFailed
	}
	ork, err := gcm.Open(nil, nonce, wrapped, rootAAD(env))
	if err != nil || len(ork) != orkBytes {
		clear(ork)
		return nil, ErrUnlockFailed
	}
	return &OwnerSession{OwnerID: env.OwnerID, ork: ork}, nil
}

func RotatePassphrase(rootPath string, oldPassphrase, newPassphrase []byte) error {
	if len(newPassphrase) == 0 {
		return errors.New("new passphrase is required")
	}
	session, err := Unlock(rootPath, oldPassphrase)
	if err != nil {
		return err
	}
	defer clear(session.ork)
	env, err := wrapRoot(session.OwnerID, newPassphrase, session.ork)
	if err != nil {
		return err
	}
	return writeJSONAtomic(rootPath, env, 0o600)
}

func rootFingerprint(session *OwnerSession) []byte {
	h := sha256.New()
	_, _ = h.Write([]byte("aurora/spk002/ork-fingerprint/v1"))
	_, _ = h.Write(session.ork)
	return h.Sum(nil)
}

func wrapRoot(ownerID string, passphrase, ork []byte) (rootEnvelope, error) {
	salt := make([]byte, saltBytes)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return rootEnvelope{}, fmt.Errorf("generate argon2 salt: %w", err)
	}
	env := rootEnvelope{
		Version:      rootEnvelopeVersion,
		OwnerID:      ownerID,
		KDF:          "Argon2id",
		ArgonVersion: argon2.Version,
		MemoryKiB:    kdfMemoryKiB,
		Iterations:   kdfIterations,
		Parallelism:  kdfParallelism,
		Salt:         base64.StdEncoding.EncodeToString(salt),
		Cipher:       "AES-256-GCM",
	}
	kek := argon2.IDKey(passphrase, salt, env.Iterations, env.MemoryKiB, env.Parallelism, kdfKeyBytes)
	defer clear(kek)
	block, err := aes.NewCipher(kek)
	if err != nil {
		return rootEnvelope{}, fmt.Errorf("create AES: %w", err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return rootEnvelope{}, fmt.Errorf("create AES-GCM: %w", err)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return rootEnvelope{}, fmt.Errorf("generate GCM nonce: %w", err)
	}
	wrapped := gcm.Seal(nil, nonce, ork, rootAAD(env))
	env.Nonce = base64.StdEncoding.EncodeToString(nonce)
	env.WrappedORK = base64.StdEncoding.EncodeToString(wrapped)
	return env, nil
}

func rootAAD(env rootEnvelope) []byte {
	return []byte(fmt.Sprintf("aurora-owner-root-v1|%s|%s|%d|%d|%d|%d|%s", env.OwnerID, env.KDF, env.ArgonVersion, env.MemoryKiB, env.Iterations, env.Parallelism, env.Cipher))
}

func deriveSubkey(session *OwnerSession, label string) ([]byte, error) {
	if session == nil || len(session.ork) != orkBytes {
		return nil, ErrOwnerAuthenticationRequired
	}
	return hkdf.Key(sha256.New, session.ork, nil, "aurora/spk002/"+label, 32)
}

func openDB(path string, create bool) (*sql.DB, error) {
	if create {
		if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
			return nil, fmt.Errorf("create DB dir: %w", err)
		}
	} else if _, err := os.Stat(path); err != nil {
		return nil, fmt.Errorf("state DB unavailable: %w", err)
	}
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("open sqlite: %w", err)
	}
	db.SetMaxOpenConns(1)
	for _, pragma := range []string{
		"PRAGMA journal_mode=WAL",
		"PRAGMA synchronous=FULL",
		"PRAGMA foreign_keys=ON",
		"PRAGMA busy_timeout=5000",
	} {
		if _, err := db.Exec(pragma); err != nil {
			_ = db.Close()
			return nil, fmt.Errorf("apply %s: %w", pragma, err)
		}
	}
	return db, nil
}

func initSchema(db *sql.DB) error {
	statements := []string{
		`CREATE TABLE IF NOT EXISTS governing_state (
			singleton INTEGER PRIMARY KEY CHECK(singleton = 1),
			generation INTEGER NOT NULL,
			owner_id TEXT NOT NULL,
			authority_revision TEXT NOT NULL,
			authority_status TEXT NOT NULL,
			expires_at_unix INTEGER NOT NULL,
			state_summary TEXT NOT NULL,
			mac BLOB NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS authority_revisions (
			authority_revision TEXT PRIMARY KEY,
			generation INTEGER NOT NULL UNIQUE,
			authority_status TEXT NOT NULL,
			expires_at_unix INTEGER NOT NULL,
			created_at_unix INTEGER NOT NULL
		)`,
	}
	for _, stmt := range statements {
		if _, err := db.Exec(stmt); err != nil {
			return fmt.Errorf("create schema: %w", err)
		}
	}
	return nil
}

func insertInitialState(tx *sql.Tx, state dbState, now time.Time) error {
	if _, err := tx.Exec(`INSERT INTO governing_state(singleton, generation, owner_id, authority_revision, authority_status, expires_at_unix, state_summary, mac) VALUES (1, ?, ?, ?, ?, ?, ?, ?)`, state.Generation, state.OwnerID, state.AuthorityRevision, state.AuthorityStatus, state.ExpiresAtUnix, state.StateSummary, state.MAC); err != nil {
		return fmt.Errorf("insert governing state: %w", err)
	}
	if _, err := tx.Exec(`INSERT INTO authority_revisions(authority_revision, generation, authority_status, expires_at_unix, created_at_unix) VALUES (?, ?, ?, ?, ?)`, state.AuthorityRevision, state.Generation, state.AuthorityStatus, state.ExpiresAtUnix, now.Unix()); err != nil {
		return fmt.Errorf("insert authority revision: %w", err)
	}
	return nil
}

func readState(q interface{ QueryRow(string, ...any) *sql.Row }) (dbState, error) {
	var s dbState
	err := q.QueryRow(`SELECT generation, owner_id, authority_revision, authority_status, expires_at_unix, state_summary, mac FROM governing_state WHERE singleton = 1`).Scan(&s.Generation, &s.OwnerID, &s.AuthorityRevision, &s.AuthorityStatus, &s.ExpiresAtUnix, &s.StateSummary, &s.MAC)
	if err != nil {
		return dbState{}, fmt.Errorf("read governing state: %w", err)
	}
	return s, nil
}

func signState(session *OwnerSession, s *dbState) error {
	key, err := deriveSubkey(session, "state-mac")
	if err != nil {
		return err
	}
	defer clear(key)
	mac := hmac.New(sha256.New, key)
	_, _ = mac.Write(stateMACPayload(*s))
	s.MAC = mac.Sum(nil)
	return nil
}

func verifyState(session *OwnerSession, s dbState) bool {
	key, err := deriveSubkey(session, "state-mac")
	if err != nil {
		return false
	}
	defer clear(key)
	mac := hmac.New(sha256.New, key)
	_, _ = mac.Write(stateMACPayload(s))
	return hmac.Equal(s.MAC, mac.Sum(nil))
}

func stateMACPayload(s dbState) []byte {
	var b bytes.Buffer
	writeI64(&b, s.Generation)
	writeString(&b, s.OwnerID)
	writeString(&b, s.AuthorityRevision)
	writeString(&b, string(s.AuthorityStatus))
	writeI64(&b, s.ExpiresAtUnix)
	writeString(&b, s.StateSummary)
	return b.Bytes()
}

func newAnchor(session *OwnerSession, generation int64, highWater int64) (anchorRecord, error) {
	a := anchorRecord{Version: anchorVersion, OwnerID: session.OwnerID, Generation: generation, WallTimeHighWaterUnix: highWater}
	key, err := deriveSubkey(session, "trust-anchor-mac")
	if err != nil {
		return anchorRecord{}, err
	}
	defer clear(key)
	mac := hmac.New(sha256.New, key)
	_, _ = mac.Write(anchorMACPayload(a))
	a.MAC = base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return a, nil
}

func verifyAnchor(session *OwnerSession, a anchorRecord) bool {
	if a.Version != anchorVersion || a.OwnerID != session.OwnerID {
		return false
	}
	provided, err := base64.StdEncoding.DecodeString(a.MAC)
	if err != nil {
		return false
	}
	key, err := deriveSubkey(session, "trust-anchor-mac")
	if err != nil {
		return false
	}
	defer clear(key)
	mac := hmac.New(sha256.New, key)
	_, _ = mac.Write(anchorMACPayload(a))
	return hmac.Equal(provided, mac.Sum(nil))
}

func anchorMACPayload(a anchorRecord) []byte {
	var b bytes.Buffer
	writeI64(&b, int64(a.Version))
	writeString(&b, a.OwnerID)
	writeI64(&b, a.Generation)
	writeI64(&b, a.WallTimeHighWaterUnix)
	return b.Bytes()
}

func loadAnchor(path string) (anchorRecord, error) {
	payload, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return anchorRecord{}, ErrMissingAnchor
		}
		return anchorRecord{}, fmt.Errorf("read trust anchor: %w", err)
	}
	var a anchorRecord
	if err := json.Unmarshal(payload, &a); err != nil {
		return anchorRecord{}, ErrInvalidAnchorMAC
	}
	return a, nil
}

func Evaluate(layout Layout, session *OwnerSession, now time.Time) (Evaluation, error) {
	db, err := openDB(layout.DBPath, false)
	if err != nil {
		return Evaluation{}, err
	}
	state, err := readState(db)
	_ = db.Close()
	if err != nil {
		return Evaluation{}, err
	}
	base := Evaluation{Generation: state.Generation, OwnerID: state.OwnerID, AuthorityRevision: state.AuthorityRevision, AuthorityStatus: state.AuthorityStatus, ExpiresAtUnix: state.ExpiresAtUnix}
	if !verifyState(session, state) {
		base.Classification = ClassInvalidDBMAC
		base.Diagnostic = diagnostic(ClassInvalidDBMAC)
		return base, nil
	}
	if _, err := os.Stat(layout.RestoreMarkerPath); err == nil {
		base.Classification = ClassRevalidationRequired
		base.Diagnostic = diagnostic(ClassRevalidationRequired)
		return base, nil
	} else if !os.IsNotExist(err) {
		return Evaluation{}, fmt.Errorf("inspect restore marker: %w", err)
	}

	anchor, err := loadAnchor(layout.AnchorPath)
	if err != nil {
		if errors.Is(err, ErrMissingAnchor) {
			base.Classification = ClassMissingAnchor
			base.Diagnostic = diagnostic(ClassMissingAnchor)
			return base, nil
		}
		base.Classification = ClassInvalidAnchorMAC
		base.Diagnostic = diagnostic(ClassInvalidAnchorMAC)
		return base, nil
	}
	if !verifyAnchor(session, anchor) {
		base.Classification = ClassInvalidAnchorMAC
		base.Diagnostic = diagnostic(ClassInvalidAnchorMAC)
		return base, nil
	}
	if state.Generation < anchor.Generation {
		base.Classification = ClassStateRollback
		base.Diagnostic = diagnostic(ClassStateRollback)
		return base, nil
	}
	if state.Generation > anchor.Generation {
		base.Classification = ClassAnchorLag
		base.Diagnostic = diagnostic(ClassAnchorLag)
		return base, nil
	}
	if now.Unix() < anchor.WallTimeHighWaterUnix {
		base.Classification = ClassTimeUntrusted
		base.Diagnostic = diagnostic(ClassTimeUntrusted)
		return base, nil
	}
	if now.Unix() > anchor.WallTimeHighWaterUnix {
		advanced, err := newAnchor(session, anchor.Generation, now.Unix())
		if err != nil {
			return Evaluation{}, err
		}
		if _, err := writeAnchorAtomic(layout.AnchorPath, advanced, nil); err != nil {
			return Evaluation{}, err
		}
	}
	base.Classification = ClassNormal
	base.Permitting = state.AuthorityStatus == StatusActive && now.Unix() < state.ExpiresAtUnix
	if base.Permitting {
		base.NextSafeAction = "CONTINUE"
	}
	base.Diagnostic = diagnostic(ClassNormal)
	return base, nil
}

func Recover(layout Layout, passphrase []byte, now time.Time) (Evaluation, error) {
	session, err := Unlock(layout.RootPath, passphrase)
	if err != nil {
		if errors.Is(err, ErrMissingRoot) {
			return Evaluation{Classification: ClassMissingRoot, Diagnostic: diagnostic(ClassMissingRoot)}, nil
		}
		if errors.Is(err, ErrUnlockFailed) {
			return Evaluation{Classification: ClassUnlockFailed, Diagnostic: diagnostic(ClassUnlockFailed)}, nil
		}
		return Evaluation{}, err
	}
	defer clear(session.ork)
	return Evaluate(layout, session, now)
}

func preflightOperationalMutation(layout Layout, session *OwnerSession, now time.Time) (dbState, anchorRecord, error) {
	if _, err := os.Stat(layout.RestoreMarkerPath); err == nil {
		return dbState{}, anchorRecord{}, ErrRevalidationRequired
	} else if !os.IsNotExist(err) {
		return dbState{}, anchorRecord{}, fmt.Errorf("inspect restore marker: %w", err)
	}

	db, err := openDB(layout.DBPath, false)
	if err != nil {
		return dbState{}, anchorRecord{}, err
	}
	state, err := readState(db)
	_ = db.Close()
	if err != nil {
		return dbState{}, anchorRecord{}, err
	}
	if !verifyState(session, state) {
		return dbState{}, anchorRecord{}, ErrInvalidDBMAC
	}

	anchor, err := loadAnchor(layout.AnchorPath)
	if err != nil {
		return dbState{}, anchorRecord{}, err
	}
	if !verifyAnchor(session, anchor) {
		return dbState{}, anchorRecord{}, ErrInvalidAnchorMAC
	}
	if state.Generation < anchor.Generation {
		return dbState{}, anchorRecord{}, ErrStateRollback
	}
	if state.Generation > anchor.Generation {
		return dbState{}, anchorRecord{}, ErrAnchorLag
	}
	if now.Unix() < anchor.WallTimeHighWaterUnix {
		return dbState{}, anchorRecord{}, ErrTimeUntrusted
	}
	return state, anchor, nil
}

func Advance(layout Layout, session *OwnerSession, mutation Mutation, now time.Time, hook FaultHook) (AdvanceMetrics, error) {
	_, anchor, err := preflightOperationalMutation(layout, session, now)
	if err != nil {
		return AdvanceMetrics{}, err
	}

	startDB := time.Now()
	state, err := advanceDBTransaction(layout.DBPath, session, mutation, now, hook)
	if err != nil {
		return AdvanceMetrics{}, err
	}
	dbNS := time.Since(startDB).Nanoseconds()
	if hook != nil {
		hook("after_db_commit")
	}
	highWater := anchor.WallTimeHighWaterUnix
	if now.Unix() > highWater {
		highWater = now.Unix()
	}
	updated, err := newAnchor(session, state.Generation, highWater)
	if err != nil {
		return AdvanceMetrics{}, err
	}
	startAnchor := time.Now()
	anchorNS, err := writeAnchorAtomic(layout.AnchorPath, updated, hook)
	if err != nil {
		return AdvanceMetrics{}, err
	}
	if anchorNS == 0 {
		anchorNS = time.Since(startAnchor).Nanoseconds()
	}
	return AdvanceMetrics{DBCommitNS: dbNS, AnchorWriteNS: anchorNS}, nil
}

func advanceDBOnly(layout Layout, session *OwnerSession, mutation Mutation, now time.Time) error {
	_, err := advanceDBTransaction(layout.DBPath, session, mutation, now, nil)
	return err
}

func advanceDBTransaction(path string, session *OwnerSession, mutation Mutation, now time.Time, hook FaultHook) (dbState, error) {
	db, err := openDB(path, false)
	if err != nil {
		return dbState{}, err
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		return dbState{}, fmt.Errorf("begin governing mutation: %w", err)
	}
	defer tx.Rollback()
	current, err := readState(tx)
	if err != nil {
		return dbState{}, err
	}
	if !verifyState(session, current) {
		return dbState{}, ErrInvalidDBMAC
	}
	if current.Generation != mutation.ExpectedGeneration {
		return dbState{}, fmt.Errorf("%w: current=%d expected=%d", ErrStaleGeneration, current.Generation, mutation.ExpectedGeneration)
	}
	if mutation.AuthorityRevision == "" || (mutation.AuthorityStatus != StatusActive && mutation.AuthorityStatus != StatusRevoked) {
		return dbState{}, errors.New("invalid authority mutation")
	}
	next := dbState{
		Generation:        current.Generation + 1,
		OwnerID:           current.OwnerID,
		AuthorityRevision: mutation.AuthorityRevision,
		AuthorityStatus:   mutation.AuthorityStatus,
		ExpiresAtUnix:     mutation.ExpiresAt.Unix(),
		StateSummary:      mutation.StateSummary,
	}
	if err := signState(session, &next); err != nil {
		return dbState{}, err
	}
	if _, err := tx.Exec(`INSERT INTO authority_revisions(authority_revision, generation, authority_status, expires_at_unix, created_at_unix) VALUES (?, ?, ?, ?, ?)`, next.AuthorityRevision, next.Generation, next.AuthorityStatus, next.ExpiresAtUnix, now.Unix()); err != nil {
		return dbState{}, fmt.Errorf("insert authority revision: %w", err)
	}
	if _, err := tx.Exec(`UPDATE governing_state SET generation=?, owner_id=?, authority_revision=?, authority_status=?, expires_at_unix=?, state_summary=?, mac=? WHERE singleton=1`, next.Generation, next.OwnerID, next.AuthorityRevision, next.AuthorityStatus, next.ExpiresAtUnix, next.StateSummary, next.MAC); err != nil {
		return dbState{}, fmt.Errorf("update governing state: %w", err)
	}
	if hook != nil {
		hook("before_db_commit")
	}
	if err := tx.Commit(); err != nil {
		return dbState{}, fmt.Errorf("commit governing mutation: %w", err)
	}
	return next, nil
}

func ReconcileAnchor(layout Layout, session *OwnerSession, now time.Time) error {
	db, err := openDB(layout.DBPath, false)
	if err != nil {
		return err
	}
	state, err := readState(db)
	_ = db.Close()
	if err != nil {
		return err
	}
	if !verifyState(session, state) {
		return ErrInvalidDBMAC
	}
	anchor, err := loadAnchor(layout.AnchorPath)
	if err != nil {
		return err
	}
	if !verifyAnchor(session, anchor) {
		return ErrInvalidAnchorMAC
	}
	if state.Generation < anchor.Generation {
		return ErrStateRollback
	}
	if state.Generation == anchor.Generation {
		return nil
	}
	highWater := anchor.WallTimeHighWaterUnix
	if now.Unix() > highWater {
		highWater = now.Unix()
	}
	advanced, err := newAnchor(session, state.Generation, highWater)
	if err != nil {
		return err
	}
	_, err = writeAnchorAtomic(layout.AnchorPath, advanced, nil)
	return err
}

func Revalidate(layout Layout, session *OwnerSession, actor Actor, input Revalidation, now time.Time) error {
	if actor != ActorOwner || session == nil {
		return ErrOwnerAuthenticationRequired
	}
	if _, err := os.Stat(layout.RestoreMarkerPath); err != nil {
		if os.IsNotExist(err) {
			return ErrRevalidationNotRequired
		}
		return err
	}
	db, err := openDB(layout.DBPath, false)
	if err != nil {
		return err
	}
	current, err := readState(db)
	_ = db.Close()
	if err != nil {
		return err
	}
	if !verifyState(session, current) {
		return ErrInvalidDBMAC
	}
	mutation := Mutation{
		ExpectedGeneration: current.Generation,
		AuthorityRevision:  input.AuthorityRevision,
		AuthorityStatus:    StatusActive,
		ExpiresAt:          input.ExpiresAt,
		StateSummary:       "owner revalidated restored authority",
	}
	next, err := advanceDBTransaction(layout.DBPath, session, mutation, now, nil)
	if err != nil {
		return err
	}
	anchor, err := newAnchor(session, next.Generation, now.Unix())
	if err != nil {
		return err
	}
	if _, err := writeAnchorAtomic(layout.AnchorPath, anchor, nil); err != nil {
		return err
	}
	if err := os.Remove(layout.RestoreMarkerPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("remove restore marker: %w", err)
	}
	return nil
}

func writeAnchorAtomic(path string, anchor anchorRecord, hook FaultHook) (int64, error) {
	payload, err := json.MarshalIndent(anchor, "", "  ")
	if err != nil {
		return 0, err
	}
	payload = append(payload, '\n')
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return 0, err
	}
	tmp := path + ".next"
	_ = os.Remove(tmp)
	start := time.Now()
	f, err := os.OpenFile(tmp, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o600)
	if err != nil {
		return 0, fmt.Errorf("create anchor temp: %w", err)
	}
	if _, err := f.Write(payload); err != nil {
		_ = f.Close()
		return 0, fmt.Errorf("write anchor temp: %w", err)
	}
	if err := f.Sync(); err != nil {
		_ = f.Close()
		return 0, fmt.Errorf("sync anchor temp: %w", err)
	}
	if err := f.Close(); err != nil {
		return 0, fmt.Errorf("close anchor temp: %w", err)
	}
	if hook != nil {
		hook("after_anchor_temp_sync")
	}
	if err := os.Rename(tmp, path); err != nil {
		return 0, fmt.Errorf("publish anchor: %w", err)
	}
	if hook != nil {
		hook("after_anchor_publish")
	}
	return time.Since(start).Nanoseconds(), nil
}

func SnapshotDB(dbPath, destination string) error {
	db, err := openDB(dbPath, false)
	if err != nil {
		return err
	}
	if _, err := db.Exec(`PRAGMA wal_checkpoint(TRUNCATE)`); err != nil {
		_ = db.Close()
		return fmt.Errorf("checkpoint snapshot DB: %w", err)
	}
	if err := db.Close(); err != nil {
		return fmt.Errorf("close snapshot DB: %w", err)
	}
	return copyFile(dbPath, destination, 0o600)
}

func ReplaceDB(dbPath, snapshot string) error {
	for _, suffix := range []string{"", "-wal", "-shm"} {
		_ = os.Remove(dbPath + suffix)
	}
	return copyFile(snapshot, dbPath, 0o600)
}

func CreateRecoveryBundle(layout Layout, bundleDir string) error {
	if _, err := os.Stat(layout.RootPath); err != nil {
		if os.IsNotExist(err) {
			return ErrMissingRoot
		}
		return err
	}
	if err := os.MkdirAll(bundleDir, 0o700); err != nil {
		return err
	}
	if err := SnapshotDB(layout.DBPath, filepath.Join(bundleDir, "state.db")); err != nil {
		return err
	}
	if err := copyFile(layout.RootPath, filepath.Join(bundleDir, "owner-root.json"), 0o600); err != nil {
		return err
	}
	if _, err := os.Stat(layout.AnchorPath); err == nil {
		if err := copyFile(layout.AnchorPath, filepath.Join(bundleDir, "historical-anchor.json"), 0o600); err != nil {
			return err
		}
	}
	db, err := openDB(layout.DBPath, false)
	if err != nil {
		return err
	}
	state, err := readState(db)
	_ = db.Close()
	if err != nil {
		return err
	}
	manifest := bundleManifest{Version: bundleVersion, OwnerID: state.OwnerID, SourceGeneration: state.Generation}
	return writeJSONAtomic(filepath.Join(bundleDir, "manifest.json"), manifest, 0o600)
}

func RestoreRecoveryBundle(bundleDir string, layout Layout) error {
	rootSource := filepath.Join(bundleDir, "owner-root.json")
	dbSource := filepath.Join(bundleDir, "state.db")
	if _, err := os.Stat(rootSource); err != nil {
		if os.IsNotExist(err) {
			return ErrMissingRoot
		}
		return err
	}
	if _, err := os.Stat(dbSource); err != nil {
		return fmt.Errorf("recovery state missing: %w", err)
	}
	manifestPayload, err := os.ReadFile(filepath.Join(bundleDir, "manifest.json"))
	if err != nil {
		return fmt.Errorf("read recovery manifest: %w", err)
	}
	var manifest bundleManifest
	if err := json.Unmarshal(manifestPayload, &manifest); err != nil || manifest.Version != bundleVersion {
		return errors.New("invalid recovery manifest")
	}
	if err := os.MkdirAll(filepath.Dir(layout.DBPath), 0o700); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(layout.RootPath), 0o700); err != nil {
		return err
	}
	if err := copyFile(dbSource, layout.DBPath, 0o600); err != nil {
		return err
	}
	if err := copyFile(rootSource, layout.RootPath, 0o600); err != nil {
		return err
	}
	_ = os.Remove(layout.AnchorPath)
	marker := restoreMarker{Version: 1, SourceGeneration: manifest.SourceGeneration}
	if err := writeJSONAtomic(layout.RestoreMarkerPath, marker, 0o600); err != nil {
		return err
	}
	return nil
}

func rawMutateStateSummary(dbPath, summary string) error {
	db, err := openDB(dbPath, false)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec(`UPDATE governing_state SET state_summary=? WHERE singleton=1`, summary)
	return err
}

func rawTamperAnchor(path string) error {
	payload, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var anchor anchorRecord
	if err := json.Unmarshal(payload, &anchor); err != nil {
		return err
	}
	anchor.Generation++
	payload, err = json.MarshalIndent(anchor, "", "  ")
	if err != nil {
		return err
	}
	payload = append(payload, '\n')
	return os.WriteFile(path, payload, 0o600)
}

func forceValidAnchorGeneration(path string, session *OwnerSession, generation int64, now time.Time) error {
	anchor, err := newAnchor(session, generation, now.Unix())
	if err != nil {
		return err
	}
	_, err = writeAnchorAtomic(path, anchor, nil)
	return err
}

func diagnostic(c Classification) string {
	switch c {
	case ClassNormal:
		return "state and trust anchor are coherent"
	case ClassStateRollback:
		return "DB generation is behind authenticated trust high-water; governing use denied"
	case ClassAnchorLag:
		return "DB commit is authenticated and ahead of trust anchor; owner reconciliation required"
	case ClassInvalidDBMAC:
		return "governing state authentication failed"
	case ClassInvalidAnchorMAC:
		return "owner trust anchor authentication failed"
	case ClassMissingAnchor:
		return "current trust anchor is missing; governing use denied"
	case ClassMissingRoot:
		return "wrapped owner root recovery material is missing"
	case ClassUnlockFailed:
		return "owner root could not be authenticated"
	case ClassTimeUntrusted:
		return "wall clock is behind authenticated time high-water; expiry-dependent authority denied"
	case ClassRevalidationRequired:
		return "historically authentic restored state is non-current until authenticated owner revalidation"
	default:
		return "unknown recovery state; governing use denied"
	}
}

func writeJSONAtomic(path string, value any, mode os.FileMode) error {
	payload, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	payload = append(payload, '\n')
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, payload, mode); err != nil {
		return err
	}
	f, err := os.OpenFile(tmp, os.O_RDWR, mode)
	if err != nil {
		return err
	}
	if err := f.Sync(); err != nil {
		_ = f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}

func copyFile(src, dst string, mode os.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(dst), 0o700); err != nil {
		return err
	}
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, mode)
	if err != nil {
		return err
	}
	_, copyErr := io.Copy(out, in)
	syncErr := out.Sync()
	closeErr := out.Close()
	if copyErr != nil {
		return copyErr
	}
	if syncErr != nil {
		return syncErr
	}
	return closeErr
}

func writeI64(w io.Writer, v int64) {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], uint64(v))
	_, _ = w.Write(buf[:])
}

func writeString(w io.Writer, s string) {
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], uint32(len(s)))
	_, _ = w.Write(buf[:])
	_, _ = io.WriteString(w, s)
}
