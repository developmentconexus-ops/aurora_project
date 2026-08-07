package store

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/developmentconexus-ops/aurora_project/spikes/m0-sovereign-store-001/internal/sqlitedriver"
)

const supportedSchemaVersion = 1

var (
	ErrAlreadyExists       = errors.New("store already exists")
	ErrMissingStore        = errors.New("store does not exist")
	ErrIncompatibleSchema  = errors.New("unsupported logical schema version")
	ErrIntegrityCheckFailed = errors.New("sqlite integrity check failed")
)

func openDB(path string) (*sql.DB, error) {
	db, err := sql.Open(sqlitedriver.Name, path)
	if err != nil {
		return nil, fmt.Errorf("sql open: %w", err)
	}
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	pragmas := []string{
		"PRAGMA journal_mode=WAL",
		"PRAGMA synchronous=FULL",
		"PRAGMA foreign_keys=ON",
		"PRAGMA busy_timeout=5000",
		"PRAGMA wal_autocheckpoint=0",
	}
	for _, pragma := range pragmas {
		if _, err := db.Exec(pragma); err != nil {
			db.Close()
			return nil, fmt.Errorf("%s: %w", pragma, err)
		}
	}
	return db, nil
}

func Bootstrap(path string, initial Snapshot) error {
	if initial.SchemaVersion != supportedSchemaVersion {
		return fmt.Errorf("%w: got %d want %d", ErrIncompatibleSchema, initial.SchemaVersion, supportedSchemaVersion)
	}
	if _, err := os.Stat(path); err == nil {
		return ErrAlreadyExists
	} else if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("stat store: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("mkdir store dir: %w", err)
	}

	db, err := openDB(path)
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("begin bootstrap: %w", err)
	}
	defer tx.Rollback()

	for _, stmt := range []string{
		`CREATE TABLE meta (key TEXT PRIMARY KEY, value TEXT NOT NULL)`,
		`CREATE TABLE projects (
			project_id TEXT PRIMARY KEY,
			current_revision INTEGER NOT NULL,
			authority_revision TEXT NOT NULL
		)`,
		`CREATE TABLE state_revisions (
			project_id TEXT NOT NULL,
			revision INTEGER NOT NULL,
			state_kind TEXT NOT NULL,
			state_summary TEXT NOT NULL,
			state_payload TEXT,
			PRIMARY KEY (project_id, revision),
			FOREIGN KEY (project_id) REFERENCES projects(project_id)
		)`,
		`CREATE TABLE transition_attempts (
			attempt_id TEXT PRIMARY KEY,
			project_id TEXT NOT NULL,
			expected_revision INTEGER NOT NULL,
			proposed_revision INTEGER NOT NULL,
			result TEXT NOT NULL,
			reason TEXT NOT NULL
		)`,
		`CREATE TABLE audit_records (
			audit_id TEXT PRIMARY KEY,
			attempt_id TEXT NOT NULL,
			project_id TEXT NOT NULL,
			revision INTEGER NOT NULL,
			event_type TEXT NOT NULL
		)`,
		`CREATE TABLE evidence_refs (
			evidence_id TEXT PRIMARY KEY,
			attempt_id TEXT NOT NULL,
			ref TEXT NOT NULL
		)`,
		`CREATE TABLE backup_attempts (
			backup_id TEXT PRIMARY KEY,
			result TEXT NOT NULL,
			path TEXT NOT NULL
		)`,
	} {
		if _, err := tx.Exec(stmt); err != nil {
			return fmt.Errorf("bootstrap schema: %w", err)
		}
	}

	if _, err := tx.Exec(`INSERT INTO meta(key, value) VALUES ('schema_version', ?), ('aurora_id', ?)`, strconv.Itoa(initial.SchemaVersion), initial.AuroraID); err != nil {
		return fmt.Errorf("insert meta: %w", err)
	}
	if _, err := tx.Exec(`INSERT INTO projects(project_id, current_revision, authority_revision) VALUES (?, ?, ?)`, initial.ProjectID, initial.CurrentRevision, initial.AuthorityRevision); err != nil {
		return fmt.Errorf("insert project: %w", err)
	}
	if _, err := tx.Exec(`INSERT INTO state_revisions(project_id, revision, state_kind, state_summary) VALUES (?, ?, ?, ?)`, initial.ProjectID, initial.CurrentRevision, initial.StateKind, initial.StateSummary); err != nil {
		return fmt.Errorf("insert initial revision: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit bootstrap: %w", err)
	}
	return nil
}

func Inspect(path string) (Snapshot, error) {
	if _, err := os.Stat(path); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return Snapshot{}, ErrMissingStore
		}
		return Snapshot{}, fmt.Errorf("stat store: %w", err)
	}

	db, err := openDB(path)
	if err != nil {
		return Snapshot{}, err
	}
	defer db.Close()

	var integrity string
	if err := db.QueryRow(`PRAGMA integrity_check`).Scan(&integrity); err != nil {
		return Snapshot{}, fmt.Errorf("integrity check: %w", err)
	}
	if integrity != "ok" {
		return Snapshot{}, fmt.Errorf("%w: %s", ErrIntegrityCheckFailed, integrity)
	}

	var schemaRaw, auroraID string
	if err := db.QueryRow(`SELECT value FROM meta WHERE key = 'schema_version'`).Scan(&schemaRaw); err != nil {
		return Snapshot{}, fmt.Errorf("read schema version: %w", err)
	}
	schemaVersion, err := strconv.Atoi(schemaRaw)
	if err != nil {
		return Snapshot{}, fmt.Errorf("parse schema version: %w", err)
	}
	if schemaVersion != supportedSchemaVersion {
		return Snapshot{}, fmt.Errorf("%w: got %d want %d", ErrIncompatibleSchema, schemaVersion, supportedSchemaVersion)
	}
	if err := db.QueryRow(`SELECT value FROM meta WHERE key = 'aurora_id'`).Scan(&auroraID); err != nil {
		return Snapshot{}, fmt.Errorf("read aurora id: %w", err)
	}

	var out Snapshot
	out.SchemaVersion = schemaVersion
	out.AuroraID = auroraID
	if err := db.QueryRow(`
		SELECT p.project_id, p.current_revision, p.authority_revision, r.state_kind, r.state_summary
		FROM projects p
		JOIN state_revisions r
		  ON r.project_id = p.project_id AND r.revision = p.current_revision
		LIMIT 1
	`).Scan(&out.ProjectID, &out.CurrentRevision, &out.AuthorityRevision, &out.StateKind, &out.StateSummary); err != nil {
		return Snapshot{}, fmt.Errorf("read governing snapshot: %w", err)
	}
	return out, nil
}

func Candidate() string { return sqlitedriver.Candidate }
