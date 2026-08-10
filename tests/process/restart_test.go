package process_test

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

func TestFreshProcessesRecoverSameAuroraProjectAndState(t *testing.T){
	root,err:=filepath.Abs("../..");if err!=nil{t.Fatal(err)}
	bin:=filepath.Join(t.TempDir(),"aurora");if runtime.GOOS=="windows"{bin += ".exe"}
	build:=exec.Command("go","build","-o",bin,"./cmd/aurora");build.Dir=root;if out,err:=build.CombinedOutput();err!=nil{t.Fatalf("build: %v\n%s",err,out)}
	data:=t.TempDir();pass:="process-fixture-passphrase\n"
	init:=runJSON(t,bin,data,pass,"init")
	project:=runJSON(t,bin,data,pass,"project","create","--label","Process Project","--objective","Fresh process continuity")
	pid:=stringField(t,project,"project_id")
	state:=runJSON(t,bin,data,pass,"project","set-state","--project",pid,"--expected","none","--kind","note","--summary","R1")
	if numberField(t,state,"state_revision")!=1{t.Fatalf("state=%v",state)}
	status:=runJSON(t,bin,data,pass,"status")
	shown:=runJSON(t,bin,data,pass,"project","show","--project",pid)
	if stringField(t,status,"aurora_id")!=stringField(t,init,"aurora_id"){t.Fatalf("Aurora identity changed: init=%v status=%v",init,status)}
	projectObj:=mapField(t,shown,"project");if stringField(t,projectObj,"project_id")!=pid{t.Fatalf("Project identity changed: %v",shown)}
	current:=mapField(t,shown,"current_state");if numberField(t,current,"state_revision")!=1{t.Fatalf("current state not recovered: %v",shown)}
}

func runJSON(t *testing.T,bin,data,stdin string,args ...string)map[string]any{t.Helper();cmdArgs:=append([]string{"--data-dir",data,"--json"},args...);cmd:=exec.Command(bin,cmdArgs...);cmd.Stdin=bytes.NewBufferString(stdin);var stdout,stderr bytes.Buffer;cmd.Stdout=&stdout;cmd.Stderr=&stderr;if err:=cmd.Run();err!=nil{t.Fatalf("run %v: %v\nstdout=%s\nstderr=%s",args,err,stdout.String(),stderr.String())};var out map[string]any;if err:=json.Unmarshal(stdout.Bytes(),&out);err!=nil{t.Fatalf("decode %v JSON: %v\nstdout=%s\nstderr=%s",args,err,stdout.String(),stderr.String())};return out}
func stringField(t *testing.T,m map[string]any,k string)string{t.Helper();v,ok:=m[k].(string);if !ok{t.Fatalf("%s missing/string in %v",k,m)};return v}
func numberField(t *testing.T,m map[string]any,k string)float64{t.Helper();v,ok:=m[k].(float64);if !ok{t.Fatalf("%s missing/number in %v",k,m)};return v}
func mapField(t *testing.T,m map[string]any,k string)map[string]any{t.Helper();v,ok:=m[k].(map[string]any);if !ok{t.Fatalf("%s missing/object in %v",k,m)};return v}
