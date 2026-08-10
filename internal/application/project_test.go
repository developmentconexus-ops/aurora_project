package application

import (
	"context"
	"testing"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/adapters/sqlite"
	"github.com/developmentconexus-ops/aurora_project/internal/adapters/trustfs"
)

type fixedClock struct{ now time.Time }
func (c fixedClock) Now() time.Time { return c.now }

func TestCreateAndShowProjectAcrossFreshStoreOpen(t *testing.T) {
	ctx := context.Background()
	root := t.TempDir()
	clock := fixedClock{now:time.Unix(100,0).UTC()}
	store, err := sqlite.Open(root)
	if err != nil { t.Fatal(err) }
	svc := &Service{State:store, Trust:trustfs.New(root), Clock:clock}
	pass := []byte("fixture-passphrase")
	if _, err := svc.Initialize(ctx, pass); err != nil { t.Fatal(err) }
	created, err := svc.CreateProject(ctx, pass, CreateProjectInput{DisplayLabel:"Power Supply", ObjectiveSummary:"Prototype revision A"})
	if err != nil { t.Fatal(err) }
	if err := store.Close(); err != nil { t.Fatal(err) }

	store, err = sqlite.Open(root)
	if err != nil { t.Fatal(err) }
	defer store.Close()
	fresh := &Service{State:store, Trust:trustfs.New(root), Clock:clock}
	view, err := fresh.ShowProject(ctx, pass, created.ProjectID)
	if err != nil { t.Fatal(err) }
	if view.Project.ProjectID != created.ProjectID || view.Project.DisplayLabel != "Power Supply" { t.Fatalf("view=%+v created=%+v",view,created) }
}
