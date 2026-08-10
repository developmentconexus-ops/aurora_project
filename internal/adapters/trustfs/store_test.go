package trustfs

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

func TestRootAndAnchorRoundTrip(t *testing.T) {
	root := t.TempDir()
	store := New(root)
	env := ports.RootEnvelope{Version:1, RootID:"ROOT-1", KDF:"argon2id", MemoryKiB:65536, Iterations:3, Parallelism:4, Salt:"AA", Nonce:"BB", WrappedORK:"CC"}
	if err := store.StoreRootEnvelope(context.Background(), env); err != nil { t.Fatal(err) }
	got, err := store.LoadRootEnvelope(context.Background())
	if err != nil { t.Fatal(err) }
	if got.RootID != env.RootID { t.Fatalf("root=%+v", got) }
	anchor := ports.Anchor{Version:1, RootID:"ROOT-1", AuroraID:"AUR-1", GoverningGeneration:1, ObservedWallTimeHighWater:time.Unix(5,0).UTC(), HMAC:"DD"}
	if err := store.PublishAnchor(context.Background(), anchor); err != nil { t.Fatal(err) }
	gotA, err := store.LoadAnchor(context.Background())
	if err != nil { t.Fatal(err) }
	if gotA.GoverningGeneration != 1 { t.Fatalf("anchor=%+v", gotA) }
}

func TestPartialRootJSONIsRejected(t *testing.T) {
	root := t.TempDir()
	path := filepath.Join(root, "trust", "owner-root.json")
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil { t.Fatal(err) }
	if err := os.WriteFile(path, []byte(`{"version":1`), 0o600); err != nil { t.Fatal(err) }
	_, err := New(root).LoadRootEnvelope(context.Background())
	if err == nil { t.Fatal("partial JSON unexpectedly accepted") }
}
