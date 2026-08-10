package application

import (
	"context"
	"fmt"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/evidence"
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
	Status                 RecoveryStatus                 `json:"status"`
	FailureClass           RecoveryFailureClass           `json:"failure_class"`
	Reason                 string                         `json:"reason"`
	AuroraID               identity.AuroraID              `json:"aurora_id,omitempty"`
	Projects               []project.Project              `json:"projects,omitempty"`
	CurrentStates          []project.ProjectStateRevision `json:"current_states,omitempty"`
	AuthorityStateRevision uint64                         `json:"authority_state_revision,omitempty"`
	TrustStatus            TrustStatus                    `json:"trust_status,omitempty"`
	GoverningGeneration    uint64                         `json:"governing_generation,omitempty"`
	OperationID            string                         `json:"operation_id"`
	Evidence               evidence.EvidenceRecord        `json:"evidence"`
}

func (s *Service) RecoverCurrentState(ctx context.Context, ownerPassphrase []byte) RecoveryResult {
	op := s.startOperation(ctx, "recover.current_state", "", evidence.Correlation{})
	snap, err := s.State.LoadCurrent(op.ctx)
	if err != nil {
		result := RecoveryResult{Status: RecoveryFailed, FailureClass: RecoveryStoreFailure, Reason: err.Error()}
		return s.completeRecovery(op.ctx, op, result, false)
	}
	authorityRevision := snap.CurrentAuthorityRevision
	op.updateCorrelation(evidence.Correlation{AuroraID: snap.Identity.AuroraID, AuthorityRevision: &authorityRevision})
	result := RecoveryResult{AuroraID: snap.Identity.AuroraID, Projects: append([]project.Project(nil), snap.Projects...), CurrentStates: append([]project.ProjectStateRevision(nil), snap.CurrentStates...), AuthorityStateRevision: snap.CurrentAuthorityRevision, GoverningGeneration: snap.GoverningGeneration}
	trust, err := s.ClassifyTrust(op.ctx, ownerPassphrase)
	if err != nil {
		result.Status, result.FailureClass, result.Reason = RecoveryFailed, RecoveryAuthorityTrustFailure, err.Error()
		return s.completeRecovery(op.ctx, op, result, true)
	}
	result.TrustStatus = trust.Status
	if trust.Status != TrustNormal {
		result.Status, result.FailureClass, result.Reason = RecoveryBlocked, RecoveryAuthorityTrustFailure, trust.Reason
		return s.completeRecovery(op.ctx, op, result, true)
	}
	state, err := decodeAuthorityState(snap.AuthorityStateJSON)
	if err != nil {
		result.Status, result.FailureClass, result.Reason = RecoveryFailed, RecoveryStateIntegrityFailure, err.Error()
		return s.completeRecovery(op.ctx, op, result, true)
	}
	if uint64(state.Revision) != snap.CurrentAuthorityRevision {
		result.Status, result.FailureClass, result.Reason = RecoveryFailed, RecoveryStateIntegrityFailure, "authority state revision does not match current pointer"
		return s.completeRecovery(op.ctx, op, result, true)
	}
	result.Status, result.FailureClass, result.Reason = RecoveryRecovered, RecoveryNone, "canonical state recovered and current trust validated"
	return s.completeRecovery(op.ctx, op, result, true)
}

func (s *Service) completeRecovery(ctx context.Context, op *operationScope, result RecoveryResult, canRecord bool) RecoveryResult {
	refs := []string{}
	if result.AuroraID != "" {
		refs = append(refs, "aurora:"+string(result.AuroraID))
	}
	for _, state := range result.CurrentStates {
		refs = append(refs, fmt.Sprintf("project:%s@%d", state.ProjectID, state.Revision))
	}
	evidenceRecord := s.operationEvidence(op, "CRIT-008", "RECOVERY_OBSERVATION", refs, "")
	result.OperationID = op.correlation.OperationID
	result.Evidence = evidenceRecord
	if canRecord {
		if err := s.appendOperationalRecord(ctx, "RECOVERY", string(result.Status), result.Reason, evidenceRecord); err != nil && result.Status == RecoveryRecovered {
			result.Status = RecoveryFailed
			result.FailureClass = RecoveryStoreFailure
			result.Reason = "record recovery evidence: " + err.Error()
		}
	}
	op.finish(string(result.Status), result.Reason)
	return result
}
