package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	_ "modernc.org/sqlite"
)

func TestApplyMigrationsAppliesInitialSchemaAndIsIdempotent(t *testing.T) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	ctx := context.Background()
	if err := applyMigrations(ctx, db); err != nil {
		t.Fatal(err)
	}

	var version int
	if err := db.QueryRowContext(ctx, "PRAGMA user_version").Scan(&version); err != nil {
		t.Fatal(err)
	}
	if version != 1 {
		t.Fatalf("user_version=%d want 1", version)
	}

	var table string
	if err := db.QueryRowContext(ctx, `SELECT name FROM sqlite_master WHERE type='table' AND name='core_state'`).Scan(&table); err != nil {
		t.Fatal(err)
	}
	if table != "core_state" {
		t.Fatalf("table=%q want core_state", table)
	}

	if err := applyMigrations(ctx, db); err != nil {
		t.Fatalf("idempotent apply: %v", err)
	}
}

func TestApplyMigrationsRejectsUnknownFuturePhysicalVersion(t *testing.T) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	ctx := context.Background()
	if _, err := db.ExecContext(ctx, "PRAGMA user_version=2"); err != nil {
		t.Fatal(err)
	}
	if err := applyMigrations(ctx, db); !errors.Is(err, ErrUnsupportedPhysicalSchema) {
		t.Fatalf("err=%v want ErrUnsupportedPhysicalSchema", err)
	}
}
