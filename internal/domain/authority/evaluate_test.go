package authority

import (
	"testing"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
)

func TestEvaluateScopeExpiryRevocationAndSupersession(t *testing.T) {
	now := time.Date(2026, 8, 9, 23, 0, 0, 0, time.UTC)
	validUntil := now.Add(time.Hour)
	state := State{Revision: 4, Grants: []Grant{
		{AuthorityID: "AUTH-OLD", AuthorityRevision: 1, SubjectOperatorID: "OWNER-LOCAL", ActorID: "AURORA_CORE", ProjectScope: []project.ProjectID{"PRJ-A"}, PermittedActionClasses: []string{"BUILD"}, LifecycleStatus: StatusSuperseded, IssuedAt: now.Add(-2 * time.Hour)},
		{AuthorityID: "AUTH-A", AuthorityRevision: 2, SubjectOperatorID: "OWNER-LOCAL", ActorID: "AURORA_CORE", ProjectScope: []project.ProjectID{"PRJ-A"}, PermittedActionClasses: []string{"BUILD"}, LifecycleStatus: StatusActive, IssuedAt: now.Add(-time.Hour), ValidUntil: &validUntil, Supersedes: "AUTH-OLD"},
		{AuthorityID: "AUTH-B", AuthorityRevision: 1, SubjectOperatorID: "OWNER-LOCAL", ActorID: "AURORA_CORE", ProjectScope: []project.ProjectID{"PRJ-B"}, PermittedActionClasses: []string{"BUILD"}, LifecycleStatus: StatusRevoked, IssuedAt: now.Add(-time.Hour), RevokedAt: &now},
	}}
	got := Evaluate(state, "PRJ-A", "BUILD", now)
	if got.Decision != DecisionPermitted || got.Status != EffectiveValid || got.SourceRevision != 4 { t.Fatalf("valid evaluation=%+v", got) }
	if got := Evaluate(state, "PRJ-B", "BUILD", now); got.Decision != DecisionBlocked { t.Fatalf("revoked grant permitted: %+v", got) }
	if got := Evaluate(state, "PRJ-C", "BUILD", now); got.Decision != DecisionBlocked { t.Fatalf("wrong scope permitted: %+v", got) }
	late := now.Add(2*time.Hour)
	if got := Evaluate(state, "PRJ-A", "BUILD", late); got.Status != EffectiveExpired || got.Decision != DecisionBlocked { t.Fatalf("expired evaluation=%+v", got) }
}

func TestNextSafeActionIsProjectionNotGrant(t *testing.T) {
	now := time.Date(2026, 8, 9, 23, 0, 0, 0, time.UTC)
	action := project.ActionDescriptor{ActionClass: "BUILD", Summary: "Build", ProjectID: "PRJ-A", RequiredAuthorityAction: "BUILD"}
	state := State{Revision: 1, Grants: []Grant{}}
	projection := NextSafeAction(state, 1, &action, now)
	if projection.Decision != DecisionBlocked || projection.AuthorityStateRevision != 1 { t.Fatalf("projection=%+v", projection) }
}
