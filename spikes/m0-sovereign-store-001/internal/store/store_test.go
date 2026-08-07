package store

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestBootstrapReopenPreservesIdentity(t *testing.T) {
	path := filepath.Join(t.TempDir(), "aurora.db")
	want := Snapshot{
		SchemaVersion:     1,
		AuroraID:          "AURORA-SPIKE-001",
		ProjectID:         "PROJECT-SPIKE-001",
		CurrentRevision:   1,
		AuthorityRevision: "AUTH-1",
		StateKind:         "ACTIVE",
		StateSummary:      "initial accepted spike state",
	}

	if err := Bootstrap(path, want); err != nil {
		t.Fatalf("bootstrap: %v", err)
	}

	got, err := Inspect(path)
	if err != nil {
		t.Fatalf("inspect after reopen: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("recovered snapshot mismatch\nwant: %#v\n got: %#v", want, got)
	}
}
