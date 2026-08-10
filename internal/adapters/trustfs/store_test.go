package trustfs

import (
	"context"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

func TestRootAndAnchorRoundTripAndCorruptionDetection(t *testing.T) {
	ctx := context.Background()
	dataDir := t.TempDir()
	store := New(dataDir)

	root := ports.RootEnvelope{
		Version: 1,
		RootID:  "ROOT-FIXTURE",
		KDF: ports.KDFParams{Algorithm: "argon2id", MemoryKiB: 64 * 1024, Iterations: 3, Parallelism: 4},
		Salt: "c2FsdA",
		Nonce: "bm9uY2U",
		WrappedORK: "d3JhcHBlZA",
	}
	if err := store.StoreRootEnvelope(ctx, root); err != nil {
		t.Fatal(err)
	}
	gotRoot, err := store.LoadRootEnvelope(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(gotRoot, root) {
		t.Fatalf("root round trip = %#v, want %#v", gotRoot, root)
	}

	anchor := ports.Anchor{
		Version: 1,
		RootID: "ROOT-FIXTURE",
		AuroraID: "AUR-FIXTURE",
		GoverningGeneration: 1,
		ObservedWallTimeHighWater: time.Date(2026, 8, 9, 20, 0, 0, 0, time.UTC),
		HMAC: "bWFj",
	}
	if err := store.PublishAnchor(ctx, anchor); err != nil {
		t.Fatal(err)
	}
	gotAnchor, err := store.LoadAnchor(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(gotAnchor, anchor) {
		t.Fatalf("anchor round trip = %#v, want %#v", gotAnchor, anchor)
	}

	rootPath := filepath.Join(dataDir, "trust", "owner-root.json")
	if err := os.WriteFile(rootPath, []byte("{"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := store.LoadRootEnvelope(ctx); err == nil {
		t.Fatal("corrupt root JSON unexpectedly accepted")
	}
}
