package application

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/adapters/exportage"
	"github.com/developmentconexus-ops/aurora_project/internal/adapters/sqlite"
	"github.com/developmentconexus-ops/aurora_project/internal/adapters/trustfs"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/authority"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
)

func TestExportRestoreFreshContextRequiresOwnerRevalidation(t *testing.T){
	ctx:=context.Background();clock:=fixedClock{now:time.Date(2026,8,10,12,0,0,0,time.UTC)};owner:=[]byte("owner-passphrase");exportSecret:=[]byte("separate-export-secret")
	sourceRoot:=t.TempDir();sourceStore,err:=sqlite.Open(sourceRoot);if err!=nil{t.Fatal(err)};source:=&Service{State:sourceStore,Trust:trustfs.New(sourceRoot),Clock:clock,ExportProtection:exportage.Protection{}}
	init,err:=source.Initialize(ctx,owner);if err!=nil{t.Fatal(err)};p,err:=source.CreateProject(ctx,owner,CreateProjectInput{DisplayLabel:"P",ObjectiveSummary:"O"});if err!=nil{t.Fatal(err)}
	rev,err:=source.TransitionProject(ctx,owner,TransitionProjectInput{AttemptID:"ATT-EXPORT",ProjectID:p.ProjectID,State:project.StateEnvelope{SchemaVersion:"1",Kind:"note",Summary:"R1",Payload:json.RawMessage(`{"value":1}`)},ProposedNextAction:&project.ActionDescriptor{ActionClass:"review",Summary:"Review",ProjectID:p.ProjectID,RequiredAuthorityAction:"review"}});if err!=nil{t.Fatal(err)}
	shown,err:=source.ShowAuthority(ctx,owner);if err!=nil{t.Fatal(err)};if _,err:=source.GrantAuthority(ctx,owner,GrantAuthorityInput{ExpectedRevision:shown.State.Revision,ProjectID:p.ProjectID,ActionClass:"review"});err!=nil{t.Fatal(err)}
	pkg,err:=source.Export(ctx,owner,exportSecret);if err!=nil{t.Fatal(err)};if pkg.ExportID==""||len(pkg.Ciphertext)==0{t.Fatalf("export=%+v",pkg)};_ = sourceStore.Close()

	targetRoot:=t.TempDir();targetStore,err:=sqlite.Open(targetRoot);if err!=nil{t.Fatal(err)};defer targetStore.Close();target:=&Service{State:targetStore,Trust:trustfs.New(targetRoot),Clock:clock,ExportProtection:exportage.Protection{}}
	restored,err:=target.Restore(ctx,pkg.Ciphertext,exportSecret,owner);if err!=nil{t.Fatal(err)}
	if restored.AuroraID!=init.AuroraID || restored.TrustStatus!=TrustRevalidationRequired{t.Fatalf("restore=%+v",restored)}
	if _,err:=target.ShowProject(ctx,owner,p.ProjectID);err==nil{t.Fatal("ordinary project inspection permitted before restore revalidation")}
	if err:=target.RevalidateRestoredAuthority(ctx,[]byte("wrong"));err==nil{t.Fatal("non-owner/wrong owner secret revalidated authority")}
	before:=restored.AuthorityStateRevision;if err:=target.RevalidateRestoredAuthority(ctx,owner);err!=nil{t.Fatal(err)}
	after,err:=target.ShowAuthority(ctx,owner);if err!=nil{t.Fatal(err)};if uint64(after.State.Revision)<=before||after.State.RevalidationRequired{t.Fatalf("authority after revalidation=%+v",after.State)}
	view,err:=target.ShowProject(ctx,owner,p.ProjectID);if err!=nil{t.Fatal(err)};if view.CurrentState==nil||view.CurrentState.Revision!=rev.Revision{t.Fatalf("restored Project=%+v",view)};if view.NextSafeAction==nil||view.NextSafeAction.Decision!=authority.Permitted{t.Fatalf("next safe action=%+v",view.NextSafeAction)}
}

func TestRestoreRejectsCollisionAndCorruptPackageWithoutChangingTarget(t *testing.T){
	ctx:=context.Background();clock:=fixedClock{now:time.Date(2026,8,10,12,0,0,0,time.UTC)};owner:=[]byte("owner-passphrase");secret:=[]byte("export-secret")
	sourceRoot:=t.TempDir();sourceStore,_:=sqlite.Open(sourceRoot);source:=&Service{State:sourceStore,Trust:trustfs.New(sourceRoot),Clock:clock,ExportProtection:exportage.Protection{}};if _,err:=source.Initialize(ctx,owner);err!=nil{t.Fatal(err)};pkg,err:=source.Export(ctx,owner,secret);if err!=nil{t.Fatal(err)};_ = sourceStore.Close()
	targetRoot:=t.TempDir();targetStore,_:=sqlite.Open(targetRoot);defer targetStore.Close();target:=&Service{State:targetStore,Trust:trustfs.New(targetRoot),Clock:clock,ExportProtection:exportage.Protection{}};targetInit,err:=target.Initialize(ctx,owner);if err!=nil{t.Fatal(err)}
	if _,err:=target.Restore(ctx,pkg.Ciphertext,secret,owner);!errors.Is(err,ErrRestoreTargetNotFresh){t.Fatalf("collision err=%v",err)}
	status,err:=target.Inspect(ctx,owner);if err!=nil{t.Fatal(err)};if status.AuroraID!=targetInit.AuroraID{t.Fatal("collision changed target identity")}
	corrupt:=append([]byte(nil),pkg.Ciphertext...);corrupt[len(corrupt)-1]^=0xff
	freshRoot:=t.TempDir();freshStore,_:=sqlite.Open(freshRoot);defer freshStore.Close();fresh:=&Service{State:freshStore,Trust:trustfs.New(freshRoot),Clock:clock,ExportProtection:exportage.Protection{}}
	if _,err:=fresh.Restore(ctx,corrupt,secret,owner);err==nil{t.Fatal("corrupt package restored")}
	if result:=fresh.RecoverCurrentState(ctx,owner);result.AuroraID!=""{t.Fatalf("failed restore fabricated state: %+v",result)}
}
