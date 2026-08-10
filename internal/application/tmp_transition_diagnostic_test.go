package application

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/adapters/sqlite"
	"github.com/developmentconexus-ops/aurora_project/internal/adapters/trustfs"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
)

func TestDiagnosticTransitionSnapshotRoundTrip(t *testing.T) {
	ctx := context.Background()
	root := t.TempDir()
	store, err := sqlite.Open(root)
	if err != nil { t.Fatal(err) }
	defer store.Close()
	svc := &Service{State: store, Trust: trustfs.New(root), Clock: fixedClock{now: time.Unix(200,0).UTC()}}
	pass := []byte("fixture-passphrase")
	if _, err := svc.Initialize(ctx, pass); err != nil { t.Fatal(err) }
	p, err := svc.CreateProject(ctx, pass, CreateProjectInput{DisplayLabel:"P", ObjectiveSummary:"O"})
	if err != nil { t.Fatal(err) }
	before, err := svc.loadTrustedCurrent(ctx, pass)
	if err != nil { t.Fatal(err) }
	defer zero(before.ORK)
	r1, err := svc.TransitionProject(ctx, pass, TransitionProjectInput{AttemptID:"ATT-DIAG", ProjectID:p.ProjectID, State:project.StateEnvelope{SchemaVersion:"1", Kind:"note", Summary:"R1", Payload:json.RawMessage(`{"value":1}`)}})
	if err != nil { t.Fatal(err) }

	expected := before.Snapshot
	expected.GoverningGeneration++
	for i := range expected.Projects {
		if expected.Projects[i].ProjectID == p.ProjectID {
			r := project.StateRevision(1)
			expected.Projects[i].CurrentStateRevision = &r
			expected.Projects[i].UpdatedAt = r1.AcceptedAt
		}
	}
	expected.CurrentStates = append(expected.CurrentStates, r1)
	actual, err := store.LoadCurrent(ctx)
	if err != nil { t.Fatal(err) }
	expected.GoverningMAC = nil
	actual.GoverningMAC = nil
	eb, _ := json.MarshalIndent(expected, "", "  ")
	ab, _ := json.MarshalIndent(actual, "", "  ")
	if string(eb) != string(ab) {
		t.Fatalf("snapshot mismatch\nEXPECTED:\n%s\nACTUAL:\n%s", eb, ab)
	}
}
