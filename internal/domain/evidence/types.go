package evidence

import (
	"context"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/identity"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
)

type Correlation struct {
	ProofRunID       string                `json:"proof_run_id,omitempty"`
	OperationID      string                `json:"operation_id"`
	AuroraID         identity.AuroraID     `json:"aurora_id,omitempty"`
	ProjectID        project.ProjectID     `json:"project_id,omitempty"`
	StateRevision    *project.StateRevision `json:"state_revision,omitempty"`
	AuthorityRevision *uint64              `json:"authority_revision,omitempty"`
}

type DomainEvent struct {
	EventID        string      `json:"event_id"`
	Source         string      `json:"source"`
	RecordedAt     time.Time   `json:"recorded_at"`
	Sequence       uint64      `json:"sequence,omitempty"`
	Classification string      `json:"classification"`
	Correlation    Correlation `json:"correlation"`
}

type AuditRecord struct {
	AuditID     string      `json:"audit_id"`
	Correlation Correlation `json:"correlation"`
	Outcome     string      `json:"outcome"`
	Reason      string      `json:"reason"`
	CreatedAt   time.Time   `json:"created_at"`
}

type Claim struct {
	ClaimID     string    `json:"claim_id"`
	Criterion   string    `json:"criterion"`
	Producer    string    `json:"producer"`
	Statement   string    `json:"statement"`
	CreatedAt   time.Time `json:"created_at"`
}

type Receipt struct {
	ReceiptID    string    `json:"receipt_id"`
	ClaimID      string    `json:"claim_id,omitempty"`
	Producer     string    `json:"producer"`
	Method       string    `json:"method"`
	ArtifactRefs []string  `json:"artifact_refs,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}

type EvidenceRecord struct {
	EvidenceID       string            `json:"evidence_id"`
	Criterion        string            `json:"criterion"`
	Producer         string            `json:"producer"`
	Verifier         string            `json:"verifier"`
	Method           string            `json:"method"`
	Environment      string            `json:"environment"`
	Versions         map[string]string `json:"versions"`
	ArtifactRefs     []string          `json:"artifact_refs,omitempty"`
	IntegrityRef     string            `json:"integrity_ref,omitempty"`
	KnownUncertainty []string          `json:"known_uncertainty,omitempty"`
	Limitations      []string          `json:"limitations"`
	Correlation      Correlation       `json:"correlation"`
	CreatedAt        time.Time         `json:"created_at"`
}

type Verdict struct {
	VerdictID    string    `json:"verdict_id"`
	Criterion    string    `json:"criterion"`
	Decision     string    `json:"decision"`
	DecidedBy    string    `json:"decided_by"`
	EvidenceRefs []string  `json:"evidence_refs,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}

type Outcome struct {
	OutcomeID   string    `json:"outcome_id"`
	OperationID string    `json:"operation_id"`
	Result      string    `json:"result"`
	Reason      string    `json:"reason"`
	CreatedAt   time.Time `json:"created_at"`
}

type StoredRecord struct {
	RecordID          string                 `json:"record_id"`
	Kind              string                 `json:"kind"`
	OperationID       string                 `json:"operation_id"`
	ProjectID         project.ProjectID      `json:"project_id,omitempty"`
	StateRevision     *project.StateRevision `json:"state_revision,omitempty"`
	AuthorityRevision *uint64                `json:"authority_revision,omitempty"`
	Outcome           string                 `json:"outcome"`
	Reason            string                 `json:"reason"`
	Details           any                    `json:"details,omitempty"`
	CreatedAt         time.Time              `json:"created_at"`
}

type proofRunKey struct{}

func WithProofRunID(ctx context.Context, proofRunID string) context.Context {
	return context.WithValue(ctx, proofRunKey{}, proofRunID)
}

func ProofRunIDFromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	value, _ := ctx.Value(proofRunKey{}).(string)
	return value
}
