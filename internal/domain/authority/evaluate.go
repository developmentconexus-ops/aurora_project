package authority

import (
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
)

func Evaluate(state State, projectID project.ProjectID, action *project.ActionDescriptor, now time.Time) Snapshot {
	out := Snapshot{SourceRevision:state.Revision, EvaluatedAt:now.UTC(), Status:Invalid}
	if state.RevalidationRequired {
		out.Status = RevalidationRequiredStatus
		out.Reasons = []string{"RESTORE_REVALIDATION_REQUIRED"}
		return out
	}
	if action == nil || action.RequiredAuthorityAction == "" {
		out.Reasons = []string{"NO_REQUIRED_ACTION"}
		return out
	}
	var sawExpired, sawRevoked, sawSuperseded bool
	for _, g := range state.Grants {
		if g.ActorID != "AURORA_CORE" || !scopeMatches(g.ProjectScope, projectID) || !g.PermitsActionClass(action.RequiredAuthorityAction) {
			continue
		}
		switch g.LifecycleStatus {
		case Revoked:
			sawRevoked = true
			continue
		case Superseded:
			sawSuperseded = true
			continue
		case Active:
		default:
			continue
		}
		if g.ValidFrom != nil && now.Before(g.ValidFrom.UTC()) { continue }
		if g.ValidUntil != nil && !now.Before(g.ValidUntil.UTC()) {
			sawExpired = true
			continue
		}
		if len(g.Conditions) != 0 { continue }
		out.Status = Valid
		out.MatchingAuthorityIDs = append(out.MatchingAuthorityIDs, g.AuthorityID)
	}
	if out.Status == Valid {
		out.Reasons = []string{"MATCHING_ACTIVE_AUTHORITY"}
		return out
	}
	switch {
	case sawRevoked:
		out.Status = RevokedStatus; out.Reasons=[]string{"MATCHING_AUTHORITY_REVOKED"}
	case sawSuperseded:
		out.Status = SupersededStatus; out.Reasons=[]string{"MATCHING_AUTHORITY_SUPERSEDED"}
	case sawExpired:
		out.Status = Expired; out.Reasons=[]string{"MATCHING_AUTHORITY_EXPIRED"}
	default:
		out.Reasons=[]string{"NO_MATCHING_CURRENT_AUTHORITY"}
	}
	return out
}

func NextSafeAction(state State, rev project.ProjectStateRevision, now time.Time) NextSafeActionProjection {
	out:=NextSafeActionProjection{ProjectID:rev.ProjectID,StateRevision:rev.Revision,AuthorityStateRevision:state.Revision,Action:rev.ProposedNextAction,Decision:None,EvaluatedAt:now.UTC()}
	if rev.ProposedNextAction==nil { out.Reasons=[]string{"NO_PROPOSED_NEXT_ACTION"}; return out }
	snapshot:=Evaluate(state,rev.ProjectID,rev.ProposedNextAction,now)
	out.Reasons=append([]string(nil),snapshot.Reasons...)
	switch snapshot.Status { case Valid:out.Decision=Permitted;case RevalidationRequiredStatus:out.Decision=ProjectionRevalidate;default:out.Decision=Blocked }
	return out
}

func scopeMatches(scope []project.ProjectID, projectID project.ProjectID) bool {
	for _, candidate := range scope { if candidate == "*" || candidate == projectID { return true } }
	return false
}
