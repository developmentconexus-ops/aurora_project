package sqlite

import (
	"context"
	"testing"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/identity"
	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

func TestOpenAppliesDurabilityPostureAndBootstrapAtomically(t *testing.T) {
	store, err := Open(t.TempDir())
	if err != nil { t.Fatal(err) }
	defer store.Close()
	var journal string
	if err := store.db.QueryRow("PRAGMA journal_mode").Scan(&journal); err != nil { t.Fatal(err) }
	if journal != "wal" { t.Fatalf("journal_mode=%q want wal", journal) }
	var sync int
	if err := store.db.QueryRow("PRAGMA synchronous").Scan(&sync); err != nil { t.Fatal(err) }
	if sync != 2 { t.Fatalf("synchronous=%d want 2(FULL)", sync) }
	var fk int
	if err := store.db.QueryRow("PRAGMA foreign_keys").Scan(&fk); err != nil { t.Fatal(err) }
	if fk != 1 { t.Fatalf("foreign_keys=%d want 1", fk) }

	id := identity.AuroraIdentity{AuroraID:"AUR-test", OwnerOperatorID:"OWNER-LEANDRO", CreatedAt:time.Unix(1,0).UTC(), IdentityRevision:1, CapabilityContractVersion:"0.2.0"}
	_, err = store.Bootstrap(context.Background(), ports.BootstrapMutation{Identity:id, AuthorityStateJSON:[]byte(`{"revision":1,"grants":[]}`), GoverningGeneration:1, GoverningMAC:[]byte("mac")})
	if err != nil { t.Fatal(err) }
	snap, err := store.LoadCurrent(context.Background())
	if err != nil { t.Fatal(err) }
	if snap.Identity.AuroraID != id.AuroraID || snap.GoverningGeneration != 1 { t.Fatalf("snapshot=%+v", snap) }
	if _, err := store.Bootstrap(context.Background(), ports.BootstrapMutation{Identity:id}); err == nil { t.Fatal("second bootstrap unexpectedly replaced identity") }
}
