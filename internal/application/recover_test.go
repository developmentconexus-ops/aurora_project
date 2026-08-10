package application

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/adapters/sqlite"
	"github.com/developmentconexus-ops/aurora_project/internal/adapters/trustfs"
)

func TestRecoverCurrentStateReturnsStructuredRecoveredResult(t *testing.T){
	ctx:=context.Background();root:=t.TempDir();store,err:=sqlite.Open(root);if err!=nil{t.Fatal(err)};svc:=&Service{State:store,Trust:trustfs.New(root),Clock:fixedClock{now:time.Date(2026,8,10,12,0,0,0,time.UTC)}};pass:=[]byte("fixture-passphrase")
	init,err:=svc.Initialize(ctx,pass);if err!=nil{t.Fatal(err)};p,err:=svc.CreateProject(ctx,pass,CreateProjectInput{DisplayLabel:"P",ObjectiveSummary:"O"});if err!=nil{t.Fatal(err)};if err:=store.Close();err!=nil{t.Fatal(err)}
	store,err=sqlite.Open(root);if err!=nil{t.Fatal(err)};defer store.Close();fresh:=&Service{State:store,Trust:trustfs.New(root),Clock:fixedClock{now:time.Date(2026,8,10,12,0,0,0,time.UTC)}}
	result:=fresh.RecoverCurrentState(ctx,pass)
	if result.Status!=RecoveryRecovered || result.AuroraID!=init.AuroraID || len(result.Projects)!=1 || result.Projects[0].ProjectID!=p.ProjectID { t.Fatalf("recovery=%+v",result) }
	if result.TrustStatus!=TrustNormal || result.AuthorityStateRevision==0 { t.Fatalf("recovery trust/authority=%+v",result) }
}

func TestRecoverClassifiesMissingStoreAndMissingAnchorWithoutFabrication(t *testing.T){
	ctx:=context.Background();empty:=t.TempDir();store,err:=sqlite.Open(empty);if err!=nil{t.Fatal(err)};svc:=&Service{State:store,Trust:trustfs.New(empty),Clock:fixedClock{now:time.Now().UTC()}}
	result:=svc.RecoverCurrentState(ctx,[]byte("anything"));_ = store.Close()
	if result.Status!=RecoveryFailed || result.FailureClass!=RecoveryStoreFailure || result.AuroraID!="" { t.Fatalf("empty recovery=%+v",result) }

	root:=t.TempDir();store,err=sqlite.Open(root);if err!=nil{t.Fatal(err)};svc=&Service{State:store,Trust:trustfs.New(root),Clock:fixedClock{now:time.Date(2026,8,10,12,0,0,0,time.UTC)}};pass:=[]byte("fixture-passphrase");if _,err:=svc.Initialize(ctx,pass);err!=nil{t.Fatal(err)};if err:=os.Remove(filepath.Join(root,"trust","owner-anchor.json"));err!=nil{t.Fatal(err)}
	result=svc.RecoverCurrentState(ctx,pass);_ = store.Close()
	if result.Status!=RecoveryBlocked || result.FailureClass!=RecoveryAuthorityTrustFailure || result.TrustStatus!=TrustMissingAnchor { t.Fatalf("missing-anchor recovery=%+v",result) }
}
