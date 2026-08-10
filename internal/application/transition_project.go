package application

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

var (
	ErrStaleStateRevision = ports.ErrStaleStateRevision
	ErrIdentityMutationAttempt = errors.New("interpreted transition field attempts to mutate Project identity")
	ErrTransitionNotAuthorized = errors.New("current authority does not permit Project state transition")
)

type TransitionProjectInput struct { AttemptID string; ProjectID project.ProjectID; ExpectedRevision *project.StateRevision; State project.StateEnvelope; ProposedNextAction *project.ActionDescriptor }

func (s *Service) TransitionProject(ctx context.Context, ownerPassphrase []byte, in TransitionProjectInput) (project.ProjectStateRevision,error) {
	trusted,err:=s.loadTrustedCurrent(ctx,ownerPassphrase);if err!=nil{return project.ProjectStateRevision{},err};defer zero(trusted.ORK)
	p,err:=s.State.GetProject(ctx,in.ProjectID);if err!=nil{return project.ProjectStateRevision{},fmt.Errorf("load Project: %w",err)}
	now:=s.Clock.Now().UTC()
	reject:=func(reason string, cause error)(project.ProjectStateRevision,error){_ = s.State.RecordNonGoverning(ctx,ports.RecordMutation{AttemptID:in.AttemptID,ProjectID:in.ProjectID,ActorID:string(trusted.Snapshot.Identity.OwnerOperatorID),RequestedAt:now,ExpectedRevision:in.ExpectedRevision,RequestedState:in.State,ProposedNextAction:in.ProposedNextAction,AuthorityEvaluationRef:"AUTH-TRANSITION-PREFLIGHT",Result:"REJECTED",Reason:reason});return project.ProjectStateRevision{},cause}
	if in.AttemptID==""{return reject("MALFORMED_ATTEMPT",errors.New("attempt ID is required"))}
	if err:=project.ValidateStateEnvelope(in.State);err!=nil{return reject("INVALID_STATE_ENVELOPE",err)}
	if err:=project.ValidateActionDescriptor(in.ProjectID,in.ProposedNextAction);err!=nil{return reject("IDENTITY_SCOPE_MISMATCH",fmt.Errorf("%w: %v",ErrIdentityMutationAttempt,err))}
	if !sameRevision(p.CurrentStateRevision,in.ExpectedRevision){return reject("STALE_REVISION",ErrStaleStateRevision)}
	if !transitionPermittedByCurrentAuthority(trusted.Snapshot.AuthorityStateJSON,in.ProjectID){return reject("UNAUTHORIZED",ErrTransitionNotAuthorized)}
	var next project.StateRevision=1;if p.CurrentStateRevision!=nil{next=*p.CurrentStateRevision+1}
	rev:=project.ProjectStateRevision{ProjectID:in.ProjectID,Revision:next,PredecessorRevision:cloneRevision(p.CurrentStateRevision),State:in.State,ProposedNextAction:in.ProposedNextAction,AcceptedByActor:string(trusted.Snapshot.Identity.OwnerOperatorID),AcceptedAt:now,TransitionAttemptID:in.AttemptID}
	newSnap:=trusted.Snapshot;newSnap.GoverningGeneration++;for i:=range newSnap.Projects{if newSnap.Projects[i].ProjectID==in.ProjectID{r:=next;newSnap.Projects[i].CurrentStateRevision=&r;newSnap.Projects[i].UpdatedAt=now}}
	replaced:=false;for i:=range newSnap.CurrentStates{if newSnap.CurrentStates[i].ProjectID==in.ProjectID{newSnap.CurrentStates[i]=rev;replaced=true}};if !replaced{newSnap.CurrentStates=append(newSnap.CurrentStates,rev)}
	mac,err:=governingMAC(trusted.ORK,newSnap);if err!=nil{return project.ProjectStateRevision{},err}
	result,err:=s.State.CommitProjectTransition(ctx,ports.ProjectTransitionMutation{AttemptID:in.AttemptID,ProjectID:in.ProjectID,ActorID:string(trusted.Snapshot.Identity.OwnerOperatorID),RequestedAt:now,ExpectedRevision:in.ExpectedRevision,Revision:rev,ExpectedGeneration:trusted.Snapshot.GoverningGeneration,NewGeneration:newSnap.GoverningGeneration,GoverningMAC:mac,AuthorityEvaluationRef:"AUTH-TRANSITION-PREFLIGHT"});if err!=nil{return project.ProjectStateRevision{},err}
	anchor:=trusted.Anchor;anchor.GoverningGeneration=newSnap.GoverningGeneration;if now.After(anchor.ObservedWallTimeHighWater){anchor.ObservedWallTimeHighWater=now};am,err:=anchorMAC(trusted.ORK,anchor);if err!=nil{return project.ProjectStateRevision{},err};anchor.HMAC=encodeMAC(am);if err:=s.Trust.PublishAnchor(ctx,anchor);err!=nil{return project.ProjectStateRevision{},fmt.Errorf("publish transition anchor: %w",err)}
	return result.Revision,nil
}

func sameRevision(a,b *project.StateRevision)bool{if a==nil||b==nil{return a==nil&&b==nil};return *a==*b}
func cloneRevision(v *project.StateRevision)*project.StateRevision{if v==nil{return nil};x:=*v;return &x}

func transitionPermittedByCurrentAuthority(raw []byte, projectID project.ProjectID) bool {
	var doc struct{ RevalidationRequired bool `json:"revalidation_required"`; Grants []struct{ ProjectScope []string `json:"project_scope"`; Permitted []string `json:"permitted_action_classes"`; Status string `json:"lifecycle_status"` } `json:"grants"` }
	if json.Unmarshal(raw,&doc)!=nil || doc.RevalidationRequired { return false }
	for _,g:=range doc.Grants{if g.Status!="ACTIVE"{continue};scope:=false;for _,p:=range g.ProjectScope{if p=="*"||p==string(projectID){scope=true}};if !scope{continue};for _,a:=range g.Permitted{if a=="project.state.transition"{return true}}};return false
}
