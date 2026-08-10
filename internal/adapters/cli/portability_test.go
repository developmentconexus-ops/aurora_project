package cli

import(
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

type sequenceSecretReader struct{ values [][]byte; i int }
func(s *sequenceSecretReader)ReadSecret(string)([]byte,error){v:=append([]byte(nil),s.values[s.i]...);s.i++;return v,nil}

func TestExportOldActiveAuthorityRestoresAsRevalidationRequired(t *testing.T){
	source:=t.TempDir();target:=t.TempDir();exportPath:=filepath.Join(t.TempDir(),"state.aurora.age");owner:=[]byte("fixture-owner-passphrase");exportSecret:=[]byte("different-export-passphrase")
	run:=func(dataDir string,secrets SecretReader,args ...string)(map[string]any,int,string){all:=append([]string{"--json","--data-dir",dataDir},args...);var out,errOut bytes.Buffer;code:=runWithSecretReader(all,&out,&errOut,secrets);var doc map[string]any;if out.Len()!=0{if err:=json.Unmarshal(out.Bytes(),&doc);err!=nil{t.Fatalf("%v json=%q err=%v",args,out.String(),err)}};return doc,code,errOut.String()}
	fixed:=fixedSecretReader{secret:owner};if _,c,e:=run(source,fixed,"init");c!=0{t.Fatal(e)};p,c,e:=run(source,fixed,"project","create","--label","Portable","--objective","Restore");if c!=0{t.Fatal(e)};pid:=p["project_id"].(string);if _,c,e:=run(source,fixed,"project","set-state","--project",pid,"--expected","none","--kind","WORK_NOTE","--summary","R1","--payload",`{}`,"--next-action","BUILD");c!=0{t.Fatal(e)};grant,c,e:=run(source,fixed,"authority","grant","--project",pid,"--action","BUILD");if c!=0{t.Fatal(e)};aid:=grant["authority_id"].(string)
	exportReader:=&sequenceSecretReader{values:[][]byte{owner,exportSecret}};if _,c,e:=run(source,exportReader,"export","--output",exportPath);c!=0{t.Fatalf("export code=%d err=%q",c,e)};if _,err:=os.Stat(exportPath);err!=nil{t.Fatal(err)}
	if _,c,e:=run(source,fixed,"authority","revoke","--authority",aid);c!=0{t.Fatal(e)}
	restoreReader:=&sequenceSecretReader{values:[][]byte{owner,exportSecret}};restored,c,e:=run(target,restoreReader,"restore","--input",exportPath);if c!=0{t.Fatalf("restore code=%d err=%q",c,e)};if restored["trust_state"]!="REVALIDATION_REQUIRED"{t.Fatalf("restore=%v",restored)}
	status,c,e:=run(target,fixed,"status");if c!=0{t.Fatal(e)};if status["trust_state"]!="REVALIDATION_REQUIRED"{t.Fatalf("status=%v",status)}
	wrong:=fixedSecretReader{secret:[]byte("wrong-owner-passphrase")};if _,c,_:=run(target,wrong,"authority","revalidate");c==0{t.Fatal("non-owner revalidation succeeded")}
	if _,c,e:=run(target,fixed,"authority","revalidate");c!=0{t.Fatalf("owner revalidation code=%d err=%q",c,e)}
	shown,c,e:=run(target,fixed,"project","show","--project",pid);if c!=0{t.Fatal(e)};if shown["current_state_revision"].(float64)!=1{t.Fatalf("shown=%v",shown)};projection:=shown["next_safe_action"].(map[string]any);if projection["decision"]!="PERMITTED"{t.Fatalf("projection=%v",projection)}
}

func TestRestoreCollisionAndWrongExportSecretLeaveTargetUnchanged(t *testing.T){
	source:=t.TempDir();target:=t.TempDir();path:=filepath.Join(t.TempDir(),"state.aurora.age");owner:=[]byte("fixture-owner-passphrase");exportSecret:=[]byte("export-secret")
	run:=func(dataDir string,secrets SecretReader,args ...string)(map[string]any,int,string){all:=append([]string{"--json","--data-dir",dataDir},args...);var out,errOut bytes.Buffer;code:=runWithSecretReader(all,&out,&errOut,secrets);var doc map[string]any;if out.Len()!=0{_ = json.Unmarshal(out.Bytes(),&doc)};return doc,code,errOut.String()}
	fixed:=fixedSecretReader{secret:owner};srcInit,c,e:=run(source,fixed,"init");if c!=0{t.Fatal(e)};exp:=&sequenceSecretReader{values:[][]byte{owner,exportSecret}};if _,c,e:=run(source,exp,"export","--output",path);c!=0{t.Fatal(e)}
	wrong:=&sequenceSecretReader{values:[][]byte{owner,[]byte("wrong-export-secret")}};if _,c,_:=run(t.TempDir(),wrong,"restore","--input",path);c==0{t.Fatal("wrong export secret restored package")}
	targetInit,c,e:=run(target,fixed,"init");if c!=0{t.Fatal(e)};restore:=&sequenceSecretReader{values:[][]byte{owner,exportSecret}};if _,c,_:=run(target,restore,"restore","--input",path);c==0{t.Fatal("identity collision restore succeeded")};status,c,e:=run(target,fixed,"status");if c!=0{t.Fatal(e)};if status["aurora_id"]!=targetInit["aurora_id"]||status["aurora_id"]==srcInit["aurora_id"]{t.Fatalf("target changed after collision: %v",status)}
}
