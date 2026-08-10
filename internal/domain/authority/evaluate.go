package authority

import (
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
)

func Evaluate(state State, projectID project.ProjectID, action string, now time.Time) Snapshot {
	result := Snapshot{SourceRevision: state.Revision, Decision: DecisionBlocked, Status: EffectiveInvalid, Reasons: []string{"NO_MATCHING_AUTHORITY"}, EvaluatedAt: now.UTC()}
	if state.RevalidationRequired {
		result.Decision = DecisionRevalidationRequired
		result.Status = EffectiveRevalidationRequired
		result.Reasons = []string{"REVALIDATION_REQUIRED"}
		return result
	}

	best := EffectiveInvalid
	for _, grant := range state.Grants {
		if !scopeMatches(grant.ProjectScope, projectID) || !contains(grant.PermittedActionClasses, action) {
			continue
		}
		switch grant.LifecycleStatus {
		case StatusRevoked:
			if best == EffectiveInvalid { best = EffectiveRevoked }
			continue
		case StatusSuperseded:
			if best == EffectiveInvalid { best = EffectiveSuperseded }
			continue
		case StatusActive:
		default:
			continue
		}
		if grant.ValidFrom != nil && now.Before(grant.ValidFrom.UTC()) {
			continue
		}
		if grant.ValidUntil != nil && !now.Before(grant.ValidUntil.UTC()) {
			best = EffectiveExpired
			continue
		}
		return Snapshot{SourceRevision: state.Revision, Decision: DecisionPermitted, Status: EffectiveValid, Reasons: []string{"MATCHING_ACTIVE_AUTHORITY"}, EvaluatedAt: now.UTC()}
	}
	result.Status = best
	if best == EffectiveInvalid {
		result.Reasons = []string{"NO_MATCHING_AUTHORITY"}
	} else {
		result.Reasons = []string{string(best)}
	}
	return result
}

func NextSafeAction(state State, stateRevision project.StateRevision, action *project.ActionDescriptor, now time.Time) NextSafeActionProjection {
	projection := NextSafeActionProjection{StateRevision: stateRevision, AuthorityStateRevision: state.Revision, Action: action, EvaluatedAt: now.UTC()}
	if action == nil {
		projection.Decision = DecisionNone
		projection.Reasons = []string{"NO_CANDIDATE_ACTION"}
		return projection
	}
	projection.ProjectID = action.ProjectID
	snapshot := Evaluate(state, action.ProjectID, action.RequiredAuthorityAction, now)
	projection.Decision = snapshot.Decision
	projection.Reasons = append([]string(nil), snapshot.Reasons...)
	return projection
}

func scopeMatches(scope []project.ProjectID, projectID project.ProjectID) bool {
	if len(scope) == 0 { return true }
	for _, item := range scope { if item == projectID { return true } }
	return false
}

func contains(items []string, wanted string) bool {
	for _, item := range items { if item == wanted { return true } }
	return false
}
