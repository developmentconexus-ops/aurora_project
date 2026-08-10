package application

import (
	"context"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/identity"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
)

type RecoveryStatus string
type RecoveryFailureClass string

const (
	RecoveryRecovered RecoveryStatus = "RECOVERED"
	RecoveryBlocked   RecoveryStatus = "BLOCKED"
	RecoveryDegraded  RecoveryStatus = "DEGRADED"
	RecoveryFailed    RecoveryStatus = "FAILED"

	RecoveryNone                  RecoveryFailureClass = "NONE"
	RecoveryStoreFailure          RecoveryFailureClass = "DURABLE_STATE_FAILURE"
	RecoveryAuthorityTrustFailure RecoveryFailureClass = "AUTHORITY_TRUST_FAILURE"
	RecoveryStateIntegrityFailure RecoveryFailureClass = "STATE_INTEGRITY_FAILURE"
	RecoveryVersionFailure        RecoveryFailureClass = "VERSION_MIGRATION_INCOMPATIBILITY"
)

type RecoveryResult struct {
	Status                 RecoveryStatus       `json:"status"`
	FailureClass           RecoveryFailureClass `json:"failure_class"`
	Reason                 string               `json:"reason"`
	AuroraID               identity.AuroraID    `json:"aurora_id,omitempty"`
	Projects               []project.Project    `json:"projects,omitempty"`
	CurrentStates          []project.ProjectStateRevision `json:"current_states,omitempty"`
	AuthorityStateRevision uint64               `json:"authority_state_revision,omitempty"`
	TrustStatus            TrustStatus          `json:"trust_status,omitempty"`
	GoverningGeneration    uint64               `json:"governing_generation,omitempty"`
}

func (s *Service) RecoverCurrentState(ctx context.Context, ownerPassphrase []byte) RecoveryResult {
	snap, err := s.State.LoadCurrent(ctx)
	if err != nil {
		return RecoveryResult{Status:RecoveryFailed, FailureClass:RecoveryStoreFailure, Reason:err.Error()}
	}
	result := RecoveryResult{AuroraID:snap.Identity.AuroraID, Projects:append([]project.Project(nil),snap.Projects...), CurrentStates:append([]project.ProjectStateRevision(nil),snap.CurrentStates...), AuthorityStateRevision:snap.CurrentAuthorityRevision, GoverningGeneration:snap.GoverningGeneration}
	trust, err := s.ClassifyTrust(ctx, ownerPassphrase)
	if err != nil {
		result.Status=RecoveryFailed;result.FailureClass=RecoveryAuthorityTrustFailure;result.Reason=err.Error();return result
	}
	result.TrustStatus=trust.Status
	if trust.Status != TrustNormal {
		result.Status=RecoveryBlocked;result.FailureClass=RecoveryAuthorityTrustFailure;result.Reason=trust.Reason;return result
	}
	state, err := decodeAuthorityState(snap.AuthorityStateJSON)
	if err != nil {
		result.Status=RecoveryFailed;result.FailureClass=RecoveryStateIntegrityFailure;result.Reason=err.Error();return result
	}
	if uint64(state.Revision) != snap.CurrentAuthorityRevision {
		result.Status=RecoveryFailed;result.FailureClass=RecoveryStateIntegrityFailure;result.Reason="authority state revision does not match current pointer";return result
	}
	result.Status=RecoveryRecovered;result.FailureClass=RecoveryNone;result.Reason="canonical state recovered and current trust validated";return result
}
