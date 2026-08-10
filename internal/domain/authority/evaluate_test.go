package authority

import (
	"testing"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/identity"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
)

func TestEvaluateMatchingScopeLifecycleAndTime(t *testing.T) {
	now := time.Date(2026,8,10,12,0,0,0,time.UTC)
	projectID := project.ProjectID("PRJ-1")
	action := &project.ActionDescriptor{ActionClass:"review", Summary:"Review result", ProjectID:projectID, RequiredAuthorityAction:"project.state.transition"}
	base := Grant{AuthorityID:"AUTH-1", AuthorityRevision:1, SubjectOperatorID:identity.OperatorID("OWNER-LEANDRO"), ActorID:"AURORA_CORE", ProjectScope:[]project.ProjectID{projectID}, PermittedActionClasses:[]string{"project.state.transition"}, LifecycleStatus:Active, IssuedAt:now.Add(-time.Hour)}

	state := State{Revision:1, Grants:[]Grant{base}}
	if got := Evaluate(state, projectID, action, now); got.Status != Valid { t.Fatalf("matching active grant status=%s want VALID", got.Status) }

	wrong := Evaluate(state, project.ProjectID("PRJ-other"), action, now)
	if wrong.Status == Valid { t.Fatalf("wrong scope unexpectedly valid: %+v", wrong) }

	expiredGrant := base; until := now.Add(-time.Minute); expiredGrant.ValidUntil=&until
	if got := Evaluate(State{Revision:2,Grants:[]Grant{expiredGrant}},projectID,action,now); got.Status != Expired { t.Fatalf("expired status=%s",got.Status) }

	revoked := base; revoked.LifecycleStatus=Revoked
	if got := Evaluate(State{Revision:3,Grants:[]Grant{revoked}},projectID,action,now); got.Status != RevokedStatus { t.Fatalf("revoked status=%s",got.Status) }

	superseded := base; superseded.LifecycleStatus=Superseded
	if got := Evaluate(State{Revision:4,Grants:[]Grant{superseded}},projectID,action,now); got.Status != SupersededStatus { t.Fatalf("superseded status=%s",got.Status) }

	if got := Evaluate(State{Revision:5,RevalidationRequired:true,Grants:[]Grant{base}},projectID,action,now); got.Status != RevalidationRequiredStatus { t.Fatalf("revalidation status=%s",got.Status) }
}

func TestNextSafeActionIsProjectionNotGrant(t *testing.T) {
	now:=time.Date(2026,8,10,12,0,0,0,time.UTC); pid:=project.ProjectID("PRJ-1")
	action:=&project.ActionDescriptor{ActionClass:"review",Summary:"Review",ProjectID:pid,RequiredAuthorityAction:"review"}
	rev:=project.ProjectStateRevision{ProjectID:pid,Revision:3,ProposedNextAction:action}
	grant:=Grant{AuthorityID:"AUTH-1",AuthorityRevision:1,SubjectOperatorID:"OWNER-LEANDRO",ActorID:"AURORA_CORE",ProjectScope:[]project.ProjectID{pid},PermittedActionClasses:[]string{"review"},LifecycleStatus:Active,IssuedAt:now}
	projection:=NextSafeAction(State{Revision:7,Grants:[]Grant{grant}},rev,now)
	if projection.Decision != Permitted || projection.AuthorityStateRevision != 7 || projection.StateRevision != 3 { t.Fatalf("projection=%+v",projection) }
	if projection.Action == nil || projection.Action.RequiredAuthorityAction!="review" { t.Fatalf("projection action=%+v",projection.Action) }
	// Mutating the projection cannot add an authority record: it contains no Grant.
}
