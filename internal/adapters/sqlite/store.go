package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/identity"
	"github.com/developmentconexus-ops/aurora_project/internal/ports"
	_ "modernc.org/sqlite"
)

var (
	ErrAlreadyInitialized = errors.New("state store already initialized")
	ErrNotInitialized = errors.New("state store is not initialized")
)

type Store struct { db *sql.DB }

func Open(dataDir string) (*Store, error) {
	stateDir := filepath.Join(dataDir, "state")
	if err := os.MkdirAll(stateDir, 0o700); err != nil { return nil, err }
	db, err := sql.Open("sqlite", filepath.Join(stateDir, "aurora.db"))
	if err != nil { return nil, err }
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	cleanup := func(err error) (*Store,error) { _ = db.Close(); return nil, err }
	for _, pragma := range []string{"PRAGMA journal_mode=WAL", "PRAGMA synchronous=FULL", "PRAGMA foreign_keys=ON"} {
		if _, err := db.Exec(pragma); err != nil { return cleanup(fmt.Errorf("%s: %w", pragma, err)) }
	}
	if _, err := db.Exec(initialSchema); err != nil { return cleanup(fmt.Errorf("apply initial schema: %w", err)) }
	return &Store{db:db}, nil
}

func (s *Store) Close() error { return s.db.Close() }

func (s *Store) Bootstrap(ctx context.Context, in ports.BootstrapMutation) (ports.BootstrapResult, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil { return ports.BootstrapResult{}, err }
	defer tx.Rollback()
	var count int
	if err := tx.QueryRowContext(ctx, "SELECT COUNT(*) FROM core_state").Scan(&count); err != nil { return ports.BootstrapResult{}, err }
	if count != 0 { return ports.BootstrapResult{}, ErrAlreadyInitialized }
	rev := in.CurrentAuthorityRevision
	if rev == 0 { rev = 1 }
	now := in.Identity.CreatedAt.UTC().Format(time.RFC3339Nano)
	if _, err := tx.ExecContext(ctx, `INSERT INTO authority_revisions(authority_revision, predecessor_revision, authority_state_json, changed_by, changed_at) VALUES(?,NULL,?,?,?)`, rev, string(in.AuthorityStateJSON), in.Identity.OwnerOperatorID, now); err != nil { return ports.BootstrapResult{}, err }
	if _, err := tx.ExecContext(ctx, `INSERT INTO core_state(singleton_key,aurora_id,owner_operator_id,logical_schema_version,current_authority_revision,governing_generation,governing_descriptor_hmac,created_at,updated_at) VALUES('core',?,?,?,?,?,?,?,?)`, in.Identity.AuroraID, in.Identity.OwnerOperatorID, 1, rev, in.GoverningGeneration, in.GoverningMAC, now, now); err != nil { return ports.BootstrapResult{}, err }
	if _, err := tx.ExecContext(ctx, `INSERT INTO records(record_id,kind,operation_id,outcome,reason,details_json,created_at,authority_revision) VALUES(?,?,?,?,?,?,?,?)`, "REC-BOOTSTRAP-1", "AUDIT", "OP-BOOTSTRAP-1", "ACCEPTED", "INITIALIZED", `{}`, now, rev); err != nil { return ports.BootstrapResult{}, err }
	if err := tx.Commit(); err != nil { return ports.BootstrapResult{}, err }
	return ports.BootstrapResult{Identity:in.Identity, GoverningGeneration:in.GoverningGeneration}, nil
}

func (s *Store) LoadCurrent(ctx context.Context) (ports.CurrentSnapshot, error) {
	var snap ports.CurrentSnapshot
	var auroraID, ownerID, createdAt string
	var identityRev uint64 = 1
	var contract = "0.2.0"
	err := s.db.QueryRowContext(ctx, `SELECT aurora_id,owner_operator_id,current_authority_revision,governing_generation,governing_descriptor_hmac,created_at FROM core_state WHERE singleton_key='core'`).Scan(&auroraID,&ownerID,&snap.CurrentAuthorityRevision,&snap.GoverningGeneration,&snap.GoverningMAC,&createdAt)
	if errors.Is(err, sql.ErrNoRows) { return snap, ErrNotInitialized }
	if err != nil { return snap, err }
	created, err := time.Parse(time.RFC3339Nano, createdAt)
	if err != nil { return snap, err }
	snap.Identity = identity.AuroraIdentity{AuroraID:identity.AuroraID(auroraID), OwnerOperatorID:identity.OperatorID(ownerID), CreatedAt:created, IdentityRevision:identityRev, CapabilityContractVersion:contract}
	if err := s.db.QueryRowContext(ctx, `SELECT authority_state_json FROM authority_revisions WHERE authority_revision=?`, snap.CurrentAuthorityRevision).Scan(&snap.AuthorityStateJSON); err != nil { return snap, err }
	return snap, nil
}

var _ ports.StateStore = (*Store)(nil)
