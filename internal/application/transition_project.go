package application

import (
	"context"
	"errors"
	"fmt"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/authority"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/evidence"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

var (
	ErrStaleStateRevision     = ports.ErrStaleStateRevision
	ErrIdentityMutationAttempt = errors.New("interpreted transition field attempts to mutate Project identity")
	ErrTransitionNotAuthorized = errors.New("current authority does not permit Project state transition")
)

type TransitionProjectInput struct {
	AttemptID          string
	ProjectID          project.ProjectID
	ExpectedRevision   *project.StateRevision
	State              project.StateEnvelope
	ProposedNextAction *project.ActionDescriptor
}

func (s *Service) TransitionProject(ctx context.Context, ownerPassphrase []byte, in TransitionProjectInput) (project.ProjectStateRevision, error) {
	op := s.startOperation(ctx, "project.state.transition", in.AttemptID, evidence.Correlation{ProjectID: in.ProjectID})

	trusted, err := s.loadTrustedCurrent(op.ctx, ownerPassphrase)
	if err != nil {
		op.finish("FAILED", "TRUST_LOAD_FAILED")
		return project.ProjectStateRevision{}, err
	}
	defer zero(trusted.ORK)

	authorityRevision := trusted.Snapshot.CurrentAuthorityRevision
	op.updateCorrelation(evidence.Correlation{
		AuroraID:          trusted.Snapshot.Identity.AuroraID,
		ProjectID:         in.ProjectID,
		AuthorityRevision: &authorityRevision,
	})

	p, err := s.State.GetProject(op.ctx, in.ProjectID)
	if err != nil {
		op.finish("FAILED", "PROJECT_LOAD_FAILED")
		return project.ProjectStateRevision{}, fmt.Errorf("load Project: %w", err)
	}
	now := s.Clock.Now().UTC()

	reject := func(reason string, cause error) (project.ProjectStateRevision, error) {
		_ = s.State.RecordNonGoverning(op.ctx, ports.RecordMutation{
			AttemptID:              in.AttemptID,
			ProjectID:              in.ProjectID,
			ActorID:                string(trusted.Snapshot.Identity.OwnerOperatorID),
			RequestedAt:            now,
			ExpectedRevision:       in.ExpectedRevision,
			RequestedState:         in.State,
			ProposedNextAction:     in.ProposedNextAction,
			AuthorityEvaluationRef: "AUTH-TRANSITION-PREFLIGHT",
			Result:                 "REJECTED",
			Reason:                 reason,
		})
		op.finish("REJECTED", reason)
		return project.ProjectStateRevision{}, cause
	}

	if in.AttemptID == "" {
		return reject("MALFORMED_ATTEMPT", errors.New("attempt ID is required"))
	}
	if err := project.ValidateStateEnvelope(in.State); err != nil {
		return reject("INVALID_STATE_ENVELOPE", err)
	}
	if err := project.ValidateActionDescriptor(in.ProjectID, in.ProposedNextAction); err != nil {
		return reject("IDENTITY_SCOPE_MISMATCH", fmt.Errorf("%w: %v", ErrIdentityMutationAttempt, err))
	}
	if !sameRevision(p.CurrentStateRevision, in.ExpectedRevision) {
		return reject("STALE_REVISION", ErrStaleStateRevision)
	}

	state, err := decodeAuthorityState(trusted.Snapshot.AuthorityStateJSON)
	if err != nil {
		return reject("INVALID_AUTHORITY", err)
	}
	operation := &project.ActionDescriptor{
		ActionClass:             "project.state.transition",
		Summary:                 "Accept Project state transition",
		ProjectID:               in.ProjectID,
		RequiredAuthorityAction: "project.state.transition",
	}
	auth := authority.Evaluate(state, in.ProjectID, operation, now)
	if auth.Status != authority.Valid {
		return reject("UNAUTHORIZED", ErrTransitionNotAuthorized)
	}

	var next project.StateRevision = 1
	if p.CurrentStateRevision != nil {
		next = *p.CurrentStateRevision + 1
	}
	rev := project.ProjectStateRevision{
		ProjectID:           in.ProjectID,
		Revision:            next,
		PredecessorRevision: cloneRevision(p.CurrentStateRevision),
		State:               in.State,
		ProposedNextAction:  in.ProposedNextAction,
		AcceptedByActor:     string(trusted.Snapshot.Identity.OwnerOperatorID),
		AcceptedAt:          now,
		TransitionAttemptID: in.AttemptID,
	}

	newSnap := trusted.Snapshot
	newSnap.GoverningGeneration++
	for i := range newSnap.Projects {
		if newSnap.Projects[i].ProjectID == in.ProjectID {
			r := next
			newSnap.Projects[i].CurrentStateRevision = &r
			newSnap.Projects[i].UpdatedAt = now
		}
	}
	replaced := false
	for i := range newSnap.CurrentStates {
		if newSnap.CurrentStates[i].ProjectID == in.ProjectID {
			newSnap.CurrentStates[i] = rev
			replaced = true
		}
	}
	if !replaced {
		newSnap.CurrentStates = append(newSnap.CurrentStates, rev)
	}

	mac, err := governingMAC(trusted.ORK, newSnap)
	if err != nil {
		op.finish("FAILED", "GOVERNING_MAC_FAILED")
		return project.ProjectStateRevision{}, err
	}
	result, err := s.State.CommitProjectTransition(op.ctx, ports.ProjectTransitionMutation{
		AttemptID:              in.AttemptID,
		ProjectID:              in.ProjectID,
		ActorID:                string(trusted.Snapshot.Identity.OwnerOperatorID),
		RequestedAt:            now,
		ExpectedRevision:       in.ExpectedRevision,
		Revision:               rev,
		ExpectedGeneration:     trusted.Snapshot.GoverningGeneration,
		NewGeneration:          newSnap.GoverningGeneration,
		GoverningMAC:           mac,
		AuthorityEvaluationRef: fmt.Sprintf("AUTH-STATE-%d", state.Revision),
	})
	if err != nil {
		if errors.Is(err, ports.ErrStaleStateRevision) {
			op.finish("REJECTED", "STALE_REVISION")
		} else {
			op.finish("FAILED", "STATE_COMMIT_FAILED")
		}
		return project.ProjectStateRevision{}, err
	}

	stateRevision := result.Revision.Revision
	op.updateCorrelation(evidence.Correlation{
		AuroraID:          trusted.Snapshot.Identity.AuroraID,
		ProjectID:         in.ProjectID,
		StateRevision:     &stateRevision,
		AuthorityRevision: &authorityRevision,
	})

	anchor := trusted.Anchor
	anchor.GoverningGeneration = newSnap.GoverningGeneration
	if now.After(anchor.ObservedWallTimeHighWater) {
		anchor.ObservedWallTimeHighWater = now
	}
	am, err := anchorMAC(trusted.ORK, anchor)
	if err != nil {
		op.finish("FAILED", "ANCHOR_MAC_FAILED")
		return project.ProjectStateRevision{}, err
	}
	anchor.HMAC = encodeMAC(am)
	if err := s.Trust.PublishAnchor(op.ctx, anchor); err != nil {
		op.finish("FAILED", "ANCHOR_PUBLISH_FAILED")
		return project.ProjectStateRevision{}, fmt.Errorf("publish transition anchor: %w", err)
	}

	op.finish("ACCEPTED", "STATE_TRANSITION")
	return result.Revision, nil
}

func sameRevision(a, b *project.StateRevision) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return *a == *b
}

func cloneRevision(v *project.StateRevision) *project.StateRevision {
	if v == nil {
		return nil
	}
	x := *v
	return &x
}
