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
	if err != nil {
		t.Fatal(err)
	}
	defer store.Close()

	created := time.Date(2026, 8, 9, 21, 0, 0, 0, time.UTC)
	_, err = store.Bootstrap(ctx, ports.BootstrapMutation{
		OperationID:             "OP-BOOTSTRAP",
		AuroraID:                "AUR-FIXTURE",
		OwnerOperatorID:         "OWNER-LOCAL",
		CreatedAt:               created,
		AuthorityRevision:       1,
		AuthorityJSON:           []byte(`{"revision":1,"grants":[]}`),
		GoverningGeneration:     1,
		GoverningDescriptorHMAC: []byte("mac-generation-1"),
	})
	if err != nil {
		t.Fatal(err)
	}

	p := ports.ProjectRecord{
		ProjectID:        "PRJ-FIXTURE",
		DisplayLabel:     "Fonte",
		ObjectiveSummary: "Construir fonte",
		CreatedAt:        created,
		UpdatedAt:        created,
	}
	stored, err := store.CreateProject(ctx, ports.CreateProjectMutation{
		OperationID:        "OP-PROJECT-1",
		Project:            p,
		ExpectedGeneration: 1,
		NewGeneration:      2,
		GoverningHMAC:      []byte("mac-generation-2"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if stored != p {
		t.Fatalf("stored project = %#v want %#v", stored, p)
	}

	got, err := store.LoadProject(ctx, p.ProjectID)
	if err != nil {
		t.Fatal(err)
	}
	if got != p {
		t.Fatalf("loaded project = %#v want %#v", got, p)
	}
	current, err := store.LoadCurrent(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if current.GoverningGeneration != 2 || len(current.Projects) != 1 {
		t.Fatalf("current snapshot = %+v, want generation 2 and one Project", current)
	}

	replacement := p
	replacement.DisplayLabel = "Replacement"
	_, err = store.CreateProject(ctx, ports.CreateProjectMutation{
		OperationID:        "OP-PROJECT-2",
		Project:            replacement,
		ExpectedGeneration: 2,
		NewGeneration:      3,
		GoverningHMAC:      []byte("mac-generation-3"),
	})
	if !errors.Is(err, ports.ErrProjectExists) {
		t.Fatalf("duplicate create err=%v, want ErrProjectExists", err)
	}
	got, err = store.LoadProject(ctx, p.ProjectID)
	if err != nil {
		t.Fatal(err)
	}
	if got.DisplayLabel != "Fonte" {
		t.Fatalf("duplicate create replaced project: %#v", got)
	}
	current, err = store.LoadCurrent(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if current.GoverningGeneration != 2 {
		t.Fatalf("duplicate create advanced generation to %d, want 2", current.GoverningGeneration)
	}
}
