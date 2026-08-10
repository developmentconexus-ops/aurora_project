package sqlite

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

func TestOpenConfiguresSQLiteAndBootstrapIsAtomic(t *testing.T) {
	ctx := context.Background()
	store, err := Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	defer store.Close()

	var journal string
	if err := store.db.QueryRowContext(ctx, "PRAGMA journal_mode").Scan(&journal); err != nil {
		t.Fatal(err)
	}
	if journal != "wal" {
		t.Fatalf("journal_mode = %q, want wal", journal)
	}
	var synchronous int
	if err := store.db.QueryRowContext(ctx, "PRAGMA synchronous").Scan(&synchronous); err != nil {
		t.Fatal(err)
	}
	if synchronous != 2 {
		t.Fatalf("synchronous = %d, want 2 (FULL)", synchronous)
	}
	var foreignKeys int
	if err := store.db.QueryRowContext(ctx, "PRAGMA foreign_keys").Scan(&foreignKeys); err != nil {
		t.Fatal(err)
	}
	if foreignKeys != 1 {
		t.Fatalf("foreign_keys = %d, want 1", foreignKeys)
	}

	mutation := ports.BootstrapMutation{
		OperationID:             "OP-BOOTSTRAP-1",
		AuroraID:                "AUR-FIXTURE",
		OwnerOperatorID:         "OWNER-LOCAL",
		CreatedAt:               time.Date(2026, 8, 9, 20, 0, 0, 0, time.UTC),
		AuthorityRevision:       1,
		AuthorityJSON:           []byte(`{"revision":1,"grants":[]}`),
		GoverningGeneration:     1,
		GoverningDescriptorHMAC: []byte("fixture-mac"),
	}
	result, err := store.Bootstrap(ctx, mutation)
	if err != nil {
		t.Fatal(err)
	}
	if result.AuroraID != mutation.AuroraID || result.GoverningGeneration != 1 {
		t.Fatalf("bootstrap result = %+v", result)
	}
	if _, err := store.Bootstrap(ctx, mutation); !errors.Is(err, ports.ErrAlreadyInitialized) {
		t.Fatalf("second bootstrap error = %v, want ErrAlreadyInitialized", err)
	}

	for table, want := range map[string]int{"core_state": 1, "authority_revisions": 1, "records": 1} {
		var got int
		if err := store.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM "+table).Scan(&got); err != nil {
			t.Fatal(err)
		}
		if got != want {
			t.Fatalf("%s rows = %d, want %d", table, got, want)
		}
	}
}
