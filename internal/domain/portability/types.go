package portability

import (
	"github.com/developmentconexus-ops/aurora_project/internal/domain/authority"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
)

type RecoveryStatus string

const (
	RecoveryRecovered RecoveryStatus = "RECOVERED"
	RecoveryDegraded  RecoveryStatus = "DEGRADED"
	RecoveryBlocked   RecoveryStatus = "BLOCKED"
	RecoveryFailed    RecoveryStatus = "FAILED"
)

type RecoveredProject struct {
	ProjectID            project.ProjectID                    `json:"project_id"`
	CurrentStateRevision *project.StateRevision               `json:"current_state_revision,omitempty"`
	NextSafeAction       *authority.NextSafeActionProjection  `json:"next_safe_action,omitempty"`
}

type RecoveryResult struct {
	OperationID            string              `json:"operation_id"`
	Status                 RecoveryStatus      `json:"status"`
	Classification         string              `json:"classification"`
	AuroraID               string              `json:"aurora_id,omitempty"`
	TrustState             string              `json:"trust_state,omitempty"`
	AuthorityStateRevision authority.Revision  `json:"authority_state_revision,omitempty"`
	Projects               []RecoveredProject  `json:"projects"`
	Limitations            []string            `json:"limitations"`
}
