package application

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/adapters/sqlite"
	"github.com/developmentconexus-ops/aurora_project/internal/adapters/trustfs"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
)

func TestTransitionCreatesRevisionThenRejectsStaleWithoutMutation(t *testing.T) {
	ctx := context.Background()
	root := t.TempDir()
	clock := fixedClock{now: time.Unix(200, 0).UTC()}
	store, err := sqlite.Open(root)
	if err != nil { t.Fatal(err) }
	defer store.Close()
	svc := &Service{State: store, Trust: trustfs.New(root), Clock: clock}
	pass := []byte("fixture-passphrase")
	if _, err := svc.Initialize(ctx, pass); err != nil { t.Fatal(err) }
	p, err := svc.CreateProject(ctx, pass, CreateProjectInput{DisplayLabel:"P", ObjectiveSummary:"O"})
	if err != nil { t.Fatal(err) }

	r1, err := svc.TransitionProject(ctx, pass, TransitionProjectInput{AttemptID:"ATT-1", ProjectID:p.ProjectID, ExpectedRevision:nil, State:project.StateEnvelope{SchemaVersion:"1", Kind:"note", Summary:"R1", Payload:json.RawMessage(`{"value":1}`)}})
	if err != nil { t.Fatal(err) }
	if r1.Revision != 1 || r1.PredecessorRevision != nil { t.Fatalf("r1=%+v",r1) }
	one := project.StateRevision(1)
	r2, err := svc.TransitionProject(ctx, pass, TransitionProjectInput{AttemptID:"ATT-2", ProjectID:p.ProjectID, ExpectedRevision:&one, State:project.StateEnvelope{SchemaVersion:"1", Kind:"note", Summary:"R2", Payload:json.RawMessage(`{"value":2}`)}})
	if err != nil { t.Fatal(err) }
	if r2.Revision != 2 || r2.PredecessorRevision == nil || *r2.PredecessorRevision != 1 { t.Fatalf("r2=%+v",r2) }
	before, err := svc.Inspect(ctx, pass); if err != nil { t.Fatal(err) }
	_, err = svc.TransitionProject(ctx, pass, TransitionProjectInput{AttemptID:"ATT-stale", ProjectID:p.ProjectID, ExpectedRevision:&one, State:project.StateEnvelope{SchemaVersion:"1", Kind:"note", Summary:"stale"}})
	if !errors.Is(err, ErrStaleStateRevision) { t.Fatalf("err=%v want ErrStaleStateRevision",err) }
	after, err := svc.Inspect(ctx, pass); if err != nil { t.Fatal(err) }
	view, err := svc.InspectProject(ctx, pass, p.ProjectID); if err != nil { t.Fatal(err) }
	if before.GoverningGeneration != after.GoverningGeneration { t.Fatalf("stale attempt advanced generation: %d -> %d",before.GoverningGeneration,after.GoverningGeneration) }
	if view.CurrentState == nil || view.CurrentState.Revision != 2 { t.Fatalf("current state changed after stale attempt: %+v",view.CurrentState) }
}

func TestTransitionRejectsMalformedAndInterpretedIdentityMismatch(t *testing.T) {
	ctx := context.Background(); root:=t.TempDir(); store,err:=sqlite.Open(root); if err!=nil{t.Fatal(err)}; defer store.Close()
	svc:=&Service{State:store,Trust:trustfs.New(root),Clock:fixedClock{now:time.Unix(300,0).UTC()}}; pass:=[]byte("fixture-passphrase")
	if _,err:=svc.Initialize(ctx,pass);err!=nil{t.Fatal(err)}
	p,err:=svc.CreateProject(ctx,pass,CreateProjectInput{DisplayLabel:"P",ObjectiveSummary:"O"});if err!=nil{t.Fatal(err)}
	before,err:=svc.Inspect(ctx,pass);if err!=nil{t.Fatal(err)}
	_,err=svc.TransitionProject(ctx,pass,TransitionProjectInput{AttemptID:"ATT-mal",ProjectID:p.ProjectID,State:project.StateEnvelope{SchemaVersion:"9",Kind:"note",Summary:"bad"}})
	if err==nil{t.Fatal("malformed envelope accepted")}
	_,err=svc.TransitionProject(ctx,pass,TransitionProjectInput{AttemptID:"ATT-id",ProjectID:p.ProjectID,State:project.StateEnvelope{SchemaVersion:"1",Kind:"note",Summary:"candidate",Payload:json.RawMessage(`{"aurora_id":"AUR-content-only"}`)},ProposedNextAction:&project.ActionDescriptor{ActionClass:"review",Summary:"mismatch",ProjectID:"PRJ-other",RequiredAuthorityAction:"project.state.transition"}})
	if !errors.Is(err,ErrIdentityMutationAttempt){t.Fatalf("err=%v want ErrIdentityMutationAttempt",err)}
	after,err:=svc.Inspect(ctx,pass);if err!=nil{t.Fatal(err)}
	if before.GoverningGeneration!=after.GoverningGeneration{t.Fatalf("rejected attempts advanced generation")}
}

func TestTransitionWrongOwnerPassphraseCannotMutate(t *testing.T) {
	ctx:=context.Background();root:=t.TempDir();store,err:=sqlite.Open(root);if err!=nil{t.Fatal(err)};defer store.Close();svc:=&Service{State:store,Trust:trustfs.New(root),Clock:fixedClock{now:time.Unix(400,0).UTC()}};pass:=[]byte("fixture-passphrase")
	if _,err:=svc.Initialize(ctx,pass);err!=nil{t.Fatal(err)};p,err:=svc.CreateProject(ctx,pass,CreateProjectInput{DisplayLabel:"P",ObjectiveSummary:"O"});if err!=nil{t.Fatal(err)}
	_,err=svc.TransitionProject(ctx,[]byte("wrong"),TransitionProjectInput{AttemptID:"ATT-noauth",ProjectID:p.ProjectID,State:project.StateEnvelope{SchemaVersion:"1",Kind:"note",Summary:"R1"}})
	if err==nil{t.Fatal("unauthenticated transition accepted")}
	view,err:=svc.InspectProject(ctx,pass,p.ProjectID);if err!=nil{t.Fatal(err)};if view.CurrentState!=nil{t.Fatalf("unauthenticated transition mutated state: %+v",view.CurrentState)}
}
