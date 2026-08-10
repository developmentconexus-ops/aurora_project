package evidence

import (
	"context"
	"reflect"
	"testing"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/identity"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
)

func TestVerificationConceptsRemainDistinct(t *testing.T) {
	types := []reflect.Type{
		reflect.TypeOf(Claim{}),
		reflect.TypeOf(Receipt{}),
		reflect.TypeOf(EvidenceRecord{}),
		reflect.TypeOf(Verdict{}),
		reflect.TypeOf(Outcome{}),
	}
	for i := range types {
		for j := i + 1; j < len(types); j++ {
			if types[i] == types[j] {
				t.Fatalf("verification concepts collapsed: %v", types[i])
			}
		}
	}

	record := EvidenceRecord{
		EvidenceID:       "EVID-1",
		Criterion:        "CRIT-008",
		Producer:         "aurora-core",
		Verifier:         "go-test",
		Method:           "INTEGRATION",
		Environment:      "ubuntu-24.04",
		Versions:         map[string]string{"source_revision": "abc123", "go": "1.26.5"},
		ArtifactRefs:      []string{"records/REC-1"},
		IntegrityRef:      "sha256:deadbeef",
		KnownUncertainty:  []string{"hosted runner is not physical power-loss evidence"},
		Limitations:       []string{"no external telemetry exporter configured"},
		Correlation: Correlation{
			ProofRunID:  "PROOF-1",
			OperationID: "OP-1",
			AuroraID:    identity.AuroraID("AUR-1"),
			ProjectID:   project.ProjectID("PRJ-1"),
		},
	}
	if record.Criterion == "" || record.Producer == "" || record.Verifier == "" || record.Method == "" || record.Environment == "" {
		t.Fatalf("incomplete evidence metadata: %+v", record)
	}
	if len(record.KnownUncertainty) == 0 || len(record.Limitations) == 0 {
		t.Fatalf("uncertainty/limitations missing: %+v", record)
	}
}

func TestProofRunCorrelationRoundTripsContext(t *testing.T) {
	ctx := WithProofRunID(context.Background(), "PROOF-CORRELATION-1")
	if got := ProofRunIDFromContext(ctx); got != "PROOF-CORRELATION-1" {
		t.Fatalf("proof run id=%q", got)
	}
	if got := ProofRunIDFromContext(context.Background()); got != "" {
		t.Fatalf("unexpected proof run id=%q", got)
	}
}
