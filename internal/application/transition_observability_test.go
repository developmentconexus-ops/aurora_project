package application

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"strings"
	"testing"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/adapters/exportage"
	"github.com/developmentconexus-ops/aurora_project/internal/adapters/sqlite"
	"github.com/developmentconexus-ops/aurora_project/internal/adapters/trustfs"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/evidence"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
)

func TestProjectTransitionTelemetryIsCorrelatedAndRedacted(t *testing.T) {
	ctx := evidence.WithProofRunID(context.Background(), "PROOF-TRANSITION-1")
	clock := fixedClock{now: time.Date(2026, 8, 10, 12, 0, 0, 0, time.UTC)}
	owner := []byte("owner-secret-never-log")
	root := t.TempDir()
	store, err := sqlite.Open(root)
	if err != nil {
		t.Fatal(err)
	}
	defer store.Close()

	var logs bytes.Buffer
	svc := &Service{
		State:            store,
		Trust:            trustfs.New(root),
		Clock:            clock,
		ExportProtection: exportage.Protection{},
		Logger:           slog.New(slog.NewJSONHandler(&logs, nil)),
	}
	if _, err := svc.Initialize(ctx, owner); err != nil {
		t.Fatal(err)
	}
	p, err := svc.CreateProject(ctx, owner, CreateProjectInput{DisplayLabel: "Telemetry", ObjectiveSummary: "Transition audit"})
	if err != nil {
		t.Fatal(err)
	}
	logs.Reset()
	payloadSecret := "payload-secret-never-log"
	accepted, err := svc.TransitionProject(ctx, owner, TransitionProjectInput{
		AttemptID: "ATT-TELEMETRY-ACCEPT",
		ProjectID: p.ProjectID,
		State: project.StateEnvelope{
			SchemaVersion: "1",
			Kind:          "accepted",
			Summary:       "safe summary",
			Payload:       json.RawMessage(`{"secret":"` + payloadSecret + `"}`),
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	text := logs.String()
	for _, want := range []string{"ATT-TELEMETRY-ACCEPT", "project.state.transition", string(p.ProjectID), "PROOF-TRANSITION-1", "ACCEPTED"} {
		if !strings.Contains(text, want) {
			t.Fatalf("transition telemetry missing %q: %s", want, text)
		}
	}
	for _, forbidden := range []string{string(owner), payloadSecret} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("transition telemetry leaked %q: %s", forbidden, text)
		}
	}

	logs.Reset()
	if _, err := svc.TransitionProject(ctx, owner, TransitionProjectInput{
		AttemptID:       "ATT-TELEMETRY-REJECT",
		ProjectID:       p.ProjectID,
		ExpectedRevision: nil,
		State: project.StateEnvelope{SchemaVersion: "1", Kind: "stale", Summary: "reject"},
	}); err == nil {
		t.Fatal("stale transition unexpectedly accepted")
	}
	text = logs.String()
	for _, want := range []string{"ATT-TELEMETRY-REJECT", "project.state.transition", "REJECTED"} {
		if !strings.Contains(text, want) {
			t.Fatalf("rejected transition telemetry missing %q: %s", want, text)
		}
	}

	state, err := store.ExportLogicalState(ctx)
	if err != nil {
		t.Fatal(err)
	}
	var foundAccepted, foundRejected bool
	for _, record := range state.Records {
		if record.OperationID != "ATT-TELEMETRY-ACCEPT" && record.OperationID != "ATT-TELEMETRY-REJECT" {
			continue
		}
		var details struct {
			Source         string `json:"source"`
			Classification string `json:"classification"`
			ActorID        string `json:"actor_id"`
		}
		if err := json.Unmarshal(record.Details, &details); err != nil {
			t.Fatalf("decode audit details: %v", err)
		}
		if record.RecordID == "" || record.CreatedAt.IsZero() || details.Source != "aurora-sovereign-core" || details.Classification != "PROJECT_TRANSITION" || details.ActorID == "" {
			t.Fatalf("incomplete transition event metadata: record=%+v details=%+v", record, details)
		}
		if record.OperationID == "ATT-TELEMETRY-ACCEPT" {
			foundAccepted = record.StateRevision != nil && *record.StateRevision == accepted.Revision
		} else {
			foundRejected = record.StateRevision == nil
		}
	}
	if !foundAccepted || !foundRejected {
		t.Fatalf("transition metadata missing accepted=%t rejected=%t", foundAccepted, foundRejected)
	}
}
