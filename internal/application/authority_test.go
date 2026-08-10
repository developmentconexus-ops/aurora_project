package application

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/adapters/sqlite"
	"github.com/developmentconexus-ops/aurora_project/internal/adapters/trustfs"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/authority"
)

func TestGrantRevokeAndRestartPreserveAuthorityRevision(t *testing.T) {
	ctx:=context.Background();root:=t.TempDir();clock:=fixedClock{now:time.Date(2026,8,10,12,0,0,0,time.UTC)};pass:=[]byte("fixture-passphrase")
	store,err:=sqlite.Open(root);if err!=nil{t.Fatal(err)};svc:=&Service{State:store,Trust:trustfs.New(root),Clock:clock}
	if _,err:=svc.Initialize(ctx,pass);err!=nil{t.Fatal(err)}
	p,err:=svc.CreateProject(ctx,pass,CreateProjectInput{DisplayLabel:"P",ObjectiveSummary:"O"});if err!=nil{t.Fatal(err)}
	before,err:=svc.ShowAuthority(ctx,pass);if err!=nil{t.Fatal(err)}
	granted,err:=svc.GrantAuthority(ctx,pass,GrantAuthorityInput{ExpectedRevision:before.State.Revision,ProjectID:p.ProjectID,ActionClass:"review"});if err!=nil{t.Fatal(err)}
	if granted.State.Revision != before.State.Revision+1 { t.Fatalf("grant state rev=%d before=%d",granted.State.Revision,before.State.Revision) }
	if len(granted.State.Grants)<2 { t.Fatalf("grant missing from state: %+v",granted.State) }
	var id authority.AuthorityID
	for _,g:=range granted.State.Grants{if g.PermitsActionClass("review"){id=g.AuthorityID}}
	if id==""{t.Fatal("did not find newly granted review authority")}
	revoked,err:=svc.RevokeAuthority(ctx,pass,RevokeAuthorityInput{ExpectedRevision:granted.State.Revision,AuthorityID:id});if err!=nil{t.Fatal(err)}
	if revoked.State.Revision != granted.State.Revision+1 { t.Fatalf("revoke state rev=%d",revoked.State.Revision) }
	if err:=store.Close();err!=nil{t.Fatal(err)}
	store,err=sqlite.Open(root);if err!=nil{t.Fatal(err)};defer store.Close();fresh:=&Service{State:store,Trust:trustfs.New(root),Clock:clock}
	shown,err:=fresh.ShowAuthority(ctx,pass);if err!=nil{t.Fatal(err)}
	if shown.State.Revision!=revoked.State.Revision{t.Fatalf("restart revision=%d want %d",shown.State.Revision,revoked.State.Revision)}
	for _,g:=range shown.State.Grants{if g.AuthorityID==id && g.LifecycleStatus!=authority.Revoked{t.Fatalf("revoked authority revived: %+v",g)}}
}

func TestAuthorityAdminRequiresOwnerAndExpectedRevision(t *testing.T){
	ctx:=context.Background();root:=t.TempDir();clock:=fixedClock{now:time.Date(2026,8,10,12,0,0,0,time.UTC)};pass:=[]byte("fixture-passphrase");store,err:=sqlite.Open(root);if err!=nil{t.Fatal(err)};defer store.Close();svc:=&Service{State:store,Trust:trustfs.New(root),Clock:clock};if _,err:=svc.Initialize(ctx,pass);err!=nil{t.Fatal(err)}
	shown,err:=svc.ShowAuthority(ctx,pass);if err!=nil{t.Fatal(err)}
	if _,err:=svc.GrantAuthority(ctx,[]byte("wrong"),GrantAuthorityInput{ExpectedRevision:shown.State.Revision,ActionClass:"review"});err==nil{t.Fatal("technical access/wrong owner secret changed authority")}
	if _,err:=svc.GrantAuthority(ctx,pass,GrantAuthorityInput{ExpectedRevision:shown.State.Revision+99,ActionClass:"review"});!errors.Is(err,ErrStaleAuthorityRevision){t.Fatalf("err=%v want ErrStaleAuthorityRevision",err)}
}
