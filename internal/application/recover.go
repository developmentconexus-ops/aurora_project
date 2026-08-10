package application

import (
	"context"
	"errors"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/authority"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/portability"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

func (s *Service) RecoverCurrentState(ctx context.Context, pass []byte) (portability.RecoveryResult, error) {
	opID, err := randomIdentifier("REC-")
	if err != nil {
		return portability.RecoveryResult{}, err
	}
	failed := func(classification string) portability.RecoveryResult {
		return portability.RecoveryResult{OperationID: opID, Status: portability.RecoveryFailed, Classification: classification, Projects: []portability.RecoveredProject{}, Limitations: []string{classification}}
	}

	current, err := s.State.LoadCurrent(ctx)
	if err != nil {
		if errors.Is(err, ports.ErrNotInitialized) {
			return failed("DURABLE_STATE_UNAVAILABLE"), nil
		}
		return failed("DURABLE_STATE_INTEGRITY_FAILURE"), nil
	}
	assessment, err := s.AssessTrust(ctx, pass)
	if err != nil {
		return portability.RecoveryResult{}, err
	}
	authorityState, err := decodeAuthorityState(current.AuthorityJSON)
	if err != nil {
		result := failed("AUTHORITY_VALIDATION_FAILURE")
		result.AuroraID = current.AuroraID
		return result, nil
	}
	result := portability.RecoveryResult{
		OperationID:            opID,
		AuroraID:               current.AuroraID,
		TrustState:             string(assessment.State),
		AuthorityStateRevision: authorityState.Revision,
		Projects:               []portability.RecoveredProject{},
		Limitations:            []string{},
	}
	if assessment.State != TrustNormal {
		result.Status = portability.RecoveryBlocked
		result.Classification = string(assessment.State)
		result.Limitations = []string{string(assessment.State)}
		return result, nil
	}

	for _, record := range current.Projects {
		recovered := portability.RecoveredProject{ProjectID: project.ProjectID(record.ProjectID)}
		if record.CurrentStateRevision != nil {
			rev := project.StateRevision(*record.CurrentStateRevision)
			recovered.CurrentStateRevision = &rev
		}
		if record.CurrentState != nil {
			state, err := projectStateFromRecord(*record.CurrentState)
			if err != nil {
				return failed("DURABLE_STATE_INTEGRITY_FAILURE"), nil
			}
			projection := authority.NextSafeAction(authorityState, state.Revision, state.ProposedNextAction, s.Clock.Now().UTC())
			recovered.NextSafeAction = &projection
		}
		result.Projects = append(result.Projects, recovered)
	}
	result.Status = portability.RecoveryRecovered
	result.Classification = "RECOVERED_CURRENT_STATE"
	return result, nil
}
