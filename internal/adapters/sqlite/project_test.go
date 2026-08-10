package sqlite

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

func TestCreateAndLoadProjectWithoutSilentReplacement(t *testing.T) {
	ctx := context.Background()
	store, err := Open(t.TempDir())
	if err != nil { t.Fatal(err) }
	defer store.Close()

	p := ports.ProjectRecord{ProjectID:"PRJ-FIXTURE", DisplayLabel:"Fonte", ObjectiveSummary:"Construir fonte", CreatedAt:time.Date(2026,8,9,21,0,0,0,time.UTC), UpdatedAt:time.Date(2026,8,9,21,0,0,0,time.UTC)}
	if err := store.CreateProject(ctx, p); err != nil { t.Fatal(err) }
	got, err := store.LoadProject(ctx, p.ProjectID)
	if err != nil { t.Fatal(err) }
	if got != p { t.Fatalf("loaded project = %#v want %#v", got, p) }
	p.DisplayLabel = "Replacement"
	if err := store.CreateProject(ctx, p); !errors.Is(err, ports.ErrProjectExists) { t.Fatalf("duplicate create err=%v, want ErrProjectExists", err) }
	got, err = store.LoadProject(ctx, p.ProjectID)
	if err != nil { t.Fatal(err) }
	if got.DisplayLabel != "Fonte" { t.Fatalf("duplicate create replaced project: %#v", got) }
}
