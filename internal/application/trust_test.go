package application

import (
	"context"
	"crypto/hmac"
	"database/sql"
	"encoding/base64"
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/adapters/sqlite"
	"github.com/developmentconexus-ops/aurora_project/internal/adapters/trustfs"
	"github.com/developmentconexus-ops/aurora_project/internal/ports"
	_ "modernc.org/sqlite"
)

func TestTrustClassificationMatrixAndAnchorReconciliation(t *testing.T) {
	cases := []struct {
		name string
		mutate func(t *testing.T, f trustFixture)
		want TrustStatus
	}{
		{name:"normal", want:TrustNormal},
		{name:"anchor lag", want:TrustAnchorLag, mutate:func(t *testing.T,f trustFixture){
			snap, err := f.store.LoadCurrent(f.ctx); if err != nil { t.Fatal(err) }
			snap.GoverningGeneration++
			mac, err := governingMAC(f.ork, snap); if err != nil { t.Fatal(err) }
			setDBGenerationAndMAC(t,f.root,snap.GoverningGeneration,mac)
		}},
		{name:"state rollback", want:TrustStateRollback, mutate:func(t *testing.T,f trustFixture){
			a, err := f.trust.LoadAnchor(f.ctx); if err != nil { t.Fatal(err) }
			a.GoverningGeneration++
			mac, err := anchorMAC(f.ork,a); if err != nil { t.Fatal(err) }
			a.HMAC=encodeMAC(mac)
			if err:=f.trust.PublishAnchor(f.ctx,a);err!=nil{t.Fatal(err)}
		}},
		{name:"invalid db mac", want:TrustInvalidDBMAC, mutate:func(t *testing.T,f trustFixture){
			setDBMAC(t,f.root,[]byte("tampered"))
		}},
		{name:"invalid anchor mac", want:TrustInvalidAnchorMAC, mutate:func(t *testing.T,f trustFixture){
			a,err:=f.trust.LoadAnchor(f.ctx);if err!=nil{t.Fatal(err)};a.HMAC=base64.RawURLEncoding.EncodeToString([]byte("tampered"));if err:=f.trust.PublishAnchor(f.ctx,a);err!=nil{t.Fatal(err)}
		}},
		{name:"missing anchor", want:TrustMissingAnchor, mutate:func(t *testing.T,f trustFixture){
			if err:=os.Remove(filepath.Join(f.root,"trust","owner-anchor.json"));err!=nil{t.Fatal(err)}
		}},
		{name:"time untrusted", want:TrustTimeUntrusted, mutate:func(t *testing.T,f trustFixture){
			a,err:=f.trust.LoadAnchor(f.ctx);if err!=nil{t.Fatal(err)};a.ObservedWallTimeHighWater=f.clock.now.Add(time.Hour);mac,err:=anchorMAC(f.ork,a);if err!=nil{t.Fatal(err)};a.HMAC=encodeMAC(mac);if err:=f.trust.PublishAnchor(f.ctx,a);err!=nil{t.Fatal(err)}
		}},
	}
	for _,tc:=range cases{t.Run(tc.name,func(t *testing.T){f:=newTrustFixture(t);defer f.close();if tc.mutate!=nil{tc.mutate(t,f)};result,err:=f.svc.ClassifyTrust(f.ctx,f.pass);if err!=nil{t.Fatal(err)};if result.Status!=tc.want{t.Fatalf("status=%s want=%s result=%+v",result.Status,tc.want,result)};if tc.want!=TrustNormal{if _,err:=f.svc.Inspect(f.ctx,f.pass);err==nil{t.Fatalf("ordinary governing access allowed under %s",tc.want)}}})}

	f:=newTrustFixture(t);defer f.close();snap,err:=f.store.LoadCurrent(f.ctx);if err!=nil{t.Fatal(err)};snap.GoverningGeneration++;mac,err:=governingMAC(f.ork,snap);if err!=nil{t.Fatal(err)};setDBGenerationAndMAC(t,f.root,snap.GoverningGeneration,mac)
	result,err:=f.svc.ClassifyTrust(f.ctx,f.pass);if err!=nil{t.Fatal(err)};if result.Status!=TrustAnchorLag{t.Fatalf("before reconcile=%+v",result)}
	if err:=f.svc.ReconcileAnchor(f.ctx,f.pass);err!=nil{t.Fatal(err)}
	after,err:=f.svc.ClassifyTrust(f.ctx,f.pass);if err!=nil{t.Fatal(err)};if after.Status!=TrustNormal{t.Fatalf("after reconcile=%+v",after)}
}

func TestRotateOwnerPassphrasePreservesORKAndOldPassphraseStopsWorking(t *testing.T){
	f:=newTrustFixture(t);defer f.close();before:=append([]byte(nil),f.ork...);newPass:=[]byte("new-fixture-passphrase")
	if err:=f.svc.RotateOwnerPassphrase(f.ctx,f.pass,newPass);err!=nil{t.Fatal(err)}
	if _,err:=f.svc.Inspect(f.ctx,f.pass);err==nil{t.Fatal("old passphrase still unlocks owner root")}
	if _,err:=f.svc.Inspect(f.ctx,newPass);err!=nil{t.Fatalf("new passphrase cannot recover current state: %v",err)}
	env,err:=f.trust.LoadRootEnvelope(f.ctx);if err!=nil{t.Fatal(err)};after,err:=unlockORK(newPass,env);if err!=nil{t.Fatal(err)};defer zero(after);if !hmac.Equal(before,after){t.Fatal("passphrase rotation changed ORK lineage")}
}

type mutableClock struct{now time.Time};func(c *mutableClock)Now()time.Time{return c.now}
type trustFixture struct{ctx context.Context;root string;store *sqlite.Store;trust *trustfs.Store;clock *mutableClock;svc *Service;pass []byte;ork []byte}
func newTrustFixture(t *testing.T)trustFixture{t.Helper();ctx:=context.Background();root:=t.TempDir();store,err:=sqlite.Open(root);if err!=nil{t.Fatal(err)};trust:=trustfs.New(root);clock:=&mutableClock{now:time.Date(2026,8,10,12,0,0,0,time.UTC)};svc:=&Service{State:store,Trust:trust,Clock:clock};pass:=[]byte("fixture-passphrase");if _,err:=svc.Initialize(ctx,pass);err!=nil{t.Fatal(err)};env,err:=trust.LoadRootEnvelope(ctx);if err!=nil{t.Fatal(err)};ork,err:=unlockORK(pass,env);if err!=nil{t.Fatal(err)};return trustFixture{ctx:ctx,root:root,store:store,trust:trust,clock:clock,svc:svc,pass:pass,ork:ork}}
func(f trustFixture)close(){zero(f.ork);_ = f.store.Close()}
func setDBGenerationAndMAC(t *testing.T,root string,generation uint64,mac []byte){t.Helper();db:=openRawDB(t,root);defer db.Close();if _,err:=db.Exec(`UPDATE core_state SET governing_generation=?, governing_descriptor_hmac=? WHERE singleton_key='core'`,generation,mac);err!=nil{t.Fatal(err)}}
func setDBMAC(t *testing.T,root string,mac []byte){t.Helper();db:=openRawDB(t,root);defer db.Close();if _,err:=db.Exec(`UPDATE core_state SET governing_descriptor_hmac=? WHERE singleton_key='core'`,mac);err!=nil{t.Fatal(err)}}
func openRawDB(t *testing.T,root string)*sql.DB{t.Helper();db,err:=sql.Open("sqlite",filepath.Join(root,"state","aurora.db"));if err!=nil{t.Fatal(err)};return db}

func TestReconcileRejectsNonAnchorLag(t *testing.T){f:=newTrustFixture(t);defer f.close();if err:=f.svc.ReconcileAnchor(f.ctx,f.pass);!errors.Is(err,ErrReconcileNotAllowed){t.Fatalf("err=%v want ErrReconcileNotAllowed",err)}}

var _ = ports.Anchor{}
