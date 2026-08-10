package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

var ErrUnsupportedPhysicalSchema = errors.New("unsupported SQLite physical schema version")

type physicalMigration struct {
	Version int
	Name    string
	SQL     string
}

var physicalMigrations = []physicalMigration{
	{Version: 1, Name: "0001_initial.sql", SQL: initialSchema},
}

func applyMigrations(ctx context.Context, db *sql.DB) error {
	var current int
	if err := db.QueryRowContext(ctx, "PRAGMA user_version").Scan(&current); err != nil {
		return fmt.Errorf("read SQLite user_version: %w", err)
	}
	latest := physicalMigrations[len(physicalMigrations)-1].Version
	if current < 0 || current > latest {
		return fmt.Errorf("%w: current=%d latest=%d", ErrUnsupportedPhysicalSchema, current, latest)
	}
	for _, migration := range physicalMigrations {
		if migration.Version <= current {
			continue
		}
		if migration.Version != current+1 {
			return fmt.Errorf("%w: migration gap after version %d", ErrUnsupportedPhysicalSchema, current)
		}
		if err := applyPhysicalMigration(ctx, db, migration); err != nil {
			return err
		}
		current = migration.Version
	}
	return nil
}

func applyPhysicalMigration(ctx context.Context, db *sql.DB, migration physicalMigration) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin %s: %w", migration.Name, err)
	}
	defer tx.Rollback()
	if _, err := tx.ExecContext(ctx, migration.SQL); err != nil {
		return fmt.Errorf("apply %s: %w", migration.Name, err)
	}
	if _, err := tx.ExecContext(ctx, fmt.Sprintf("PRAGMA user_version=%d", migration.Version)); err != nil {
		return fmt.Errorf("mark %s: %w", migration.Name, err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit %s: %w", migration.Name, err)
	}
	return nil
}

const initialSchema = `
CREATE TABLE IF NOT EXISTS authority_revisions (
  authority_revision INTEGER PRIMARY KEY,
  predecessor_revision INTEGER NULL,
  authority_state_json TEXT NOT NULL,
  changed_by TEXT NOT NULL,
  changed_at TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS core_state (
  singleton_key TEXT PRIMARY KEY CHECK(singleton_key = 'core'),
  aurora_id TEXT NOT NULL UNIQUE,
  owner_operator_id TEXT NOT NULL,
  logical_schema_version INTEGER NOT NULL,
  current_authority_revision INTEGER NOT NULL,
  governing_generation INTEGER NOT NULL,
  governing_descriptor_hmac BLOB NOT NULL,
  created_at TEXT NOT NULL,
  updated_at TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS projects (
  project_id TEXT PRIMARY KEY,
  display_label TEXT NOT NULL,
  objective_summary TEXT NOT NULL,
  current_state_revision INTEGER NULL,
  created_at TEXT NOT NULL,
  updated_at TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS project_state_revisions (
  project_id TEXT NOT NULL,
  state_revision INTEGER NOT NULL,
  predecessor_revision INTEGER NULL,
  state_schema_version TEXT NOT NULL,
  state_kind TEXT NOT NULL,
  state_summary TEXT NOT NULL,
  state_payload_json TEXT NULL,
  accepted_intent_ref TEXT NULL,
  proposed_next_action_json TEXT NULL,
  accepted_by_actor TEXT NOT NULL,
  accepted_at TEXT NOT NULL,
  transition_attempt_id TEXT NOT NULL UNIQUE,
  PRIMARY KEY(project_id, state_revision),
  FOREIGN KEY(project_id) REFERENCES projects(project_id)
);
CREATE TABLE IF NOT EXISTS transition_attempts (
  attempt_id TEXT PRIMARY KEY,
  project_id TEXT NOT NULL,
  actor_id TEXT NOT NULL,
  requested_at TEXT NOT NULL,
  expected_state_revision INTEGER NULL,
  requested_state_json TEXT NOT NULL,
  proposed_next_action_json TEXT NULL,
  authority_evaluation_ref TEXT NULL,
  result TEXT NOT NULL,
  reason TEXT NOT NULL,
  accepted_state_revision INTEGER NULL,
  FOREIGN KEY(project_id) REFERENCES projects(project_id)
);
CREATE TABLE IF NOT EXISTS records (
  record_id TEXT PRIMARY KEY,
  kind TEXT NOT NULL,
  operation_id TEXT NOT NULL,
  project_id TEXT NULL,
  state_revision INTEGER NULL,
  authority_revision INTEGER NULL,
  outcome TEXT NOT NULL,
  reason TEXT NOT NULL,
  details_json TEXT NOT NULL,
  created_at TEXT NOT NULL
);
`
