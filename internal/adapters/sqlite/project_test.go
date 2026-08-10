package sqlite

import (
	"context"
	"testing"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/identity"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

func TestCreateProjectRejectsDuplicateIdentityAndSurvivesReopen(t *testing.T) {
	root := t.TempDir()
	store, err := Open(root)
	if err != nil { t.Fatal(err) }
	p := project.Project{ProjectID:"PRJ-fixed", DisplayLabel:"Power", ObjectiveSummary:"Build supply", CreatedAt:time.Unix(10,0).UTC(), UpdatedAt:time.Unix(10,0).UTC()}
	seedStoreForProjectTest(t, store)
	_, err = store.CreateProject(context.Background(), ports.CreateProjectMutation{Project:p, ExpectedGeneration:1, NewGeneration:2, GoverningMAC:[]byte("m2"), OperationID:"OP-PROJ-1"})
	if err != nil { t.Fatal(err) }
	if _, err := store.CreateProject(context.Background(), ports.CreateProjectMutation{Project:p, ExpectedGeneration:2, NewGeneration:3, GoverningMAC:[]byte("m3"), OperationID:"OP-PROJ-2"}); err == nil { t.Fatal("duplicate Project identity replaced existing row") }
	if err := store.Close(); err != nil { t.Fatal(err) }
	store, err = Open(root)
	if err != nil { t.Fatal(err) }
	defer store.Close()
	got, err := store.GetProject(context.Background(), p.ProjectID)
	if err != nil { t.Fatal(err) }
	if got.ProjectID != p.ProjectID || got.DisplayLabel != p.DisplayLabel || got.ObjectiveSummary != p.ObjectiveSummary { t.Fatalf("got=%+v want=%+v", got,p) }
}

func seedStoreForProjectTest(t *testing.T, store *Store) {
	t.Helper()
	id := testIdentity(time.Unix(1,0).UTC())
	if _, err := store.Bootstrap(context.Background(), ports.BootstrapMutation{Identity:id, AuthorityStateJSON:[]byte(`{"revision":1,"grants":[]}`), CurrentAuthorityRevision:1, GoverningGeneration:1, GoverningMAC:[]byte("m1")}); err != nil { t.Fatal(err) }
}

func testIdentity(at time.Time) identity.AuroraIdentity {
	return identity.AuroraIdentity{AuroraID:"AUR-test", OwnerOperatorID:"OWNER-LEANDRO", CreatedAt:at, IdentityRevision:1, CapabilityContractVersion:"0.2.0"}
}
