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

	for reopen := 1; reopen <= 5; reopen++ {
		got, err := Inspect(path)
		if err != nil {
			t.Fatalf("inspect after reopen %d: %v", reopen, err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("recovered snapshot mismatch after reopen %d\nwant: %#v\n got: %#v", reopen, want, got)
		}
	}
}
