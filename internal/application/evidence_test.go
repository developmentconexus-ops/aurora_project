package application

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/adapters/exportage"
	"github.com/developmentconexus-ops/aurora_project/internal/adapters/sqlite"
	"github.com/developmentconexus-ops/aurora_project/internal/adapters/trustfs"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/evidence"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/identity"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
)

func TestTransitionAuditDistinguishesAcceptedAndRejectedAttempts(t *testing.T) {
	ctx := evidence.WithProofRunID(context.Background(), "PROOF-AUDIT-1")
	clock := fixedClock{now: time.Date(2026, 8, 10, 12, 0, 0, 0, time.UTC)}
	owner := []byte("owner-passphrase")
	root := t.TempDir()
	store, err := sqlite.Open(root)
	if err != nil {
		t.Fatal(err)
	}
	defer store.Close()
	svc := &Service{State: store, Trust: trustfs.New(root), Clock: clock, ExportProtection: exportage.Protection{}}
	if _, err := svc.Initialize(ctx, owner); err != nil {
		t.Fatal(err)
	}
	p, err := svc.CreateProject(ctx, owner, CreateProjectInput{DisplayLabel: "Audit", ObjectiveSummary: "Evidence"})
	if err != nil {
		t.Fatal(err)
	}
	accepted, err := svc.TransitionProject(ctx, owner, TransitionProjectInput{
		AttemptID: "ATT-EVID-ACCEPT",
		ProjectID: p.ProjectID,
		State: project.StateEnvelope{SchemaVersion: "1", Kind: "accepted", Summary: "R1"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := svc.TransitionProject(ctx, owner, TransitionProjectInput{
		AttemptID:       "ATT-EVID-REJECT",
		ProjectID:       p.ProjectID,
		ExpectedRevision: nil,
		State:           project.StateEnvelope{SchemaVersion: "1", Kind: "stale", Summary: "must reject"},
	}); err == nil {
		t.Fatal("stale transition unexpectedly accepted")
	}

	state, err := store.ExportLogicalState(ctx)
	if err != nil {
		t.Fatal(err)
	}
	var acceptedRecord, rejectedRecord bool
	for _, record := range state.Records {
		switch record.OperationID {
		case "ATT-EVID-ACCEPT":
			acceptedRecord = record.Outcome == "ACCEPTED" && record.Reason == "STATE_TRANSITION" && record.StateRevision != nil && *record.StateRevision == accepted.Revision
		case "ATT-EVID-REJECT":
			rejectedRecord = record.Outcome == "REJECTED" && record.Reason != "" && record.StateRevision == nil
		}
	}
	if !acceptedRecord || !rejectedRecord {
		t.Fatalf("accepted/rejected audit missing: records=%+v", state.Records)
	}
	var acceptedAttempt, rejectedAttempt bool
	for _, attempt := range state.Attempts {
		switch attempt.AttemptID {
		case "ATT-EVID-ACCEPT":
			acceptedAttempt = attempt.Result == "ACCEPTED" && attempt.ActorID != "" && !attempt.RequestedAt.IsZero()
		case "ATT-EVID-REJECT":
			rejectedAttempt = attempt.Result == "REJECTED" && attempt.Reason != "" && attempt.ActorID != "" && !attempt.RequestedAt.IsZero()
		}
	}
	if !acceptedAttempt || !rejectedAttempt {
		t.Fatalf("attempt attribution missing: attempts=%+v", state.Attempts)
	}
}

func TestRecoveryExportRestoreAndMigrationReturnAttributableEvidence(t *testing.T) {
	ctx := evidence.WithProofRunID(context.Background(), "PROOF-CORR-1")
	clock := fixedClock{now: time.Date(2026, 8, 10, 12, 0, 0, 0, time.UTC)}
	owner := []byte("owner-passphrase")
	exportSecret := []byte("separate-export-secret")

	sourceRoot := t.TempDir()
	sourceStore, err := sqlite.Open(sourceRoot)
	if err != nil {
		t.Fatal(err)
	}
	defer sourceStore.Close()
	source := &Service{State: sourceStore, Trust: trustfs.New(sourceRoot), Clock: clock, ExportProtection: exportage.Protection{}}
	initialized, err := source.Initialize(ctx, owner)
	if err != nil {
		t.Fatal(err)
	}
	p, err := source.CreateProject(ctx, owner, CreateProjectInput{DisplayLabel: "Proof", ObjectiveSummary: "Correlation"})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := source.TransitionProject(ctx, owner, TransitionProjectInput{AttemptID: "ATT-CORR-1", ProjectID: p.ProjectID, State: project.StateEnvelope{SchemaVersion: "1", Kind: "note", Summary: "R1"}}); err != nil {
		t.Fatal(err)
	}

	recovery := source.RecoverCurrentState(ctx, owner)
	assertEvidenceCorrelation(t, recovery.OperationID, recovery.Evidence, "PROOF-CORR-1", initialized.AuroraID)
	if len(recovery.Evidence.ArtifactRefs) == 0 {
		t.Fatalf("recovery evidence has no refs: %+v", recovery.Evidence)
	}

	exported, err := source.Export(ctx, owner, exportSecret)
	if err != nil {
		t.Fatal(err)
	}
	assertEvidenceCorrelation(t, exported.OperationID, exported.Evidence, "PROOF-CORR-1", initialized.AuroraID)
	if exported.Evidence.IntegrityRef == "" {
		t.Fatalf("export evidence lacks integrity reference: %+v", exported.Evidence)
	}

	targetRoot := t.TempDir()
	targetStore, err := sqlite.Open(targetRoot)
	if err != nil {
		t.Fatal(err)
	}
	defer targetStore.Close()
	target := &Service{State: targetStore, Trust: trustfs.New(targetRoot), Clock: clock, ExportProtection: exportage.Protection{}}
	restored, err := target.Restore(ctx, exported.Ciphertext, exportSecret, owner)
	if err != nil {
		t.Fatal(err)
	}
	assertEvidenceCorrelation(t, restored.OperationID, restored.Evidence, "PROOF-CORR-1", initialized.AuroraID)
	if restored.Evidence.IntegrityRef == "" {
		t.Fatalf("restore evidence lacks integrity reference: %+v", restored.Evidence)
	}

	fixture, err := os.ReadFile("../../testdata/migration/v0-valid.json")
	if err != nil {
		t.Fatal(err)
	}
	protected, err := (exportage.Protection{}).Protect(ctx, fixture, exportSecret)
	if err != nil {
		t.Fatal(err)
	}
	migrator := &Service{Clock: clock, ExportProtection: exportage.Protection{}}
	migrated, err := migrator.MigratePackage(ctx, protected, exportSecret)
	if err != nil {
		t.Fatal(err)
	}
	if migrated.OperationID == "" || migrated.Evidence.Criterion == "" || migrated.Evidence.Method == "" || migrated.Evidence.Environment == "" {
		t.Fatalf("migration evidence incomplete: %+v", migrated)
	}
	if migrated.Evidence.Correlation.ProofRunID != "PROOF-CORR-1" || migrated.Evidence.Correlation.OperationID != migrated.OperationID {
		t.Fatalf("migration correlation=%+v", migrated.Evidence.Correlation)
	}
}

func TestTelemetryFailureCannotChangeCoreResultAndSecretsAreRedacted(t *testing.T) {
	ctx := context.Background()
	clock := fixedClock{now: time.Date(2026, 8, 10, 12, 0, 0, 0, time.UTC)}
	owner := []byte("owner-secret-do-not-log")
	exportSecret := []byte("export-secret-do-not-log")
	root := t.TempDir()
	store, err := sqlite.Open(root)
	if err != nil {
		t.Fatal(err)
	}
	defer store.Close()

	var logs bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&logs, nil))
	svc := &Service{State: store, Trust: trustfs.New(root), Clock: clock, ExportProtection: exportage.Protection{}, Logger: logger}
	if _, err := svc.Initialize(ctx, owner); err != nil {
		t.Fatal(err)
	}
	p, err := svc.CreateProject(ctx, owner, CreateProjectInput{DisplayLabel: "Sensitive", ObjectiveSummary: "No payload in telemetry"})
	if err != nil {
		t.Fatal(err)
	}
	payload := "sensitive-payload-do-not-log"
	if _, err := svc.TransitionProject(ctx, owner, TransitionProjectInput{AttemptID: "ATT-SECRET-1", ProjectID: p.ProjectID, State: project.StateEnvelope{SchemaVersion: "1", Kind: "secret", Summary: "redacted", Payload: json.RawMessage(`{"secret":"` + payload + `"}`)}}); err != nil {
		t.Fatal(err)
	}
	if _, err := svc.Export(ctx, owner, exportSecret); err != nil {
		t.Fatal(err)
	}
	text := logs.String()
	for _, forbidden := range []string{string(owner), string(exportSecret), payload} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("telemetry leaked secret %q: %s", forbidden, text)
		}
	}
	if !strings.Contains(text, "operation_id") || !strings.Contains(text, "operation_name") {
		t.Fatalf("stable correlation fields missing from logs: %s", text)
	}

	failing := &Service{State: store, Trust: trustfs.New(root), Clock: clock, ExportProtection: exportage.Protection{}, Logger: slog.New(alwaysFailHandler{})}
	result := failing.RecoverCurrentState(ctx, owner)
	if result.Status != RecoveryRecovered {
		t.Fatalf("telemetry failure changed recovery result: %+v", result)
	}
}

func assertEvidenceCorrelation(t *testing.T, operationID string, record evidence.EvidenceRecord, proofRunID string, auroraID identity.AuroraID) {
	t.Helper()
	if operationID == "" || record.EvidenceID == "" || record.Criterion == "" || record.Producer == "" || record.Verifier == "" || record.Method == "" || record.Environment == "" {
		t.Fatalf("evidence incomplete: operation=%q evidence=%+v", operationID, record)
	}
	if record.Correlation.OperationID != operationID || record.Correlation.ProofRunID != proofRunID || record.Correlation.AuroraID != auroraID {
		t.Fatalf("correlation mismatch: %+v", record.Correlation)
	}
	if len(record.Versions) == 0 || len(record.Limitations) == 0 {
		t.Fatalf("version/limitation metadata missing: %+v", record)
	}
}

type alwaysFailHandler struct{}

func (alwaysFailHandler) Enabled(context.Context, slog.Level) bool { return true }
func (alwaysFailHandler) Handle(context.Context, slog.Record) error { return errors.New("telemetry sink unavailable") }
func (h alwaysFailHandler) WithAttrs([]slog.Attr) slog.Handler { return h }
func (h alwaysFailHandler) WithGroup(string) slog.Handler { return h }
