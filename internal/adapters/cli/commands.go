package cli

import (
	"context"
	"flag"
	"fmt"
	"io"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/adapters/sqlite"
	"github.com/developmentconexus-ops/aurora_project/internal/adapters/trustfs"
	"github.com/developmentconexus-ops/aurora_project/internal/application"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
)

type wallClock struct{};func(wallClock)Now()time.Time{return time.Now()}
func runCommand(command []string,opts globalOptions,out,errOut io.Writer)int{if len(command)==1&&command[0]=="help"{fmt.Fprint(out,helpText);return 0};store,err:=sqlite.Open(opts.dataDir);if err!=nil{fmt.Fprintln(errOut,"open state:",err);return 1};defer store.Close();svc:=&application.Service{State:store,Trust:trustfs.New(opts.dataDir),Clock:wallClock{}};ctx:=context.Background();switch command[0]{case"init":return runInit(ctx,svc,opts,out,errOut);case"status":return runStatus(ctx,svc,opts,out,errOut);case"project":return runProject(ctx,svc,command[1:],opts,out,errOut);default:fmt.Fprintf(errOut,"unknown command %q; use --help\n",command[0]);return 2}}
func ownerSecret(errOut io.Writer)([]byte,error){return promptSecret("Owner passphrase: ",errOut)}
func runInit(ctx context.Context,svc *application.Service,opts globalOptions,out,errOut io.Writer)int{secret,err:=ownerSecret(errOut);if err!=nil{fmt.Fprintln(errOut,"read owner passphrase:",err);return 1};defer wipe(secret);result,err:=svc.Initialize(ctx,secret);if err!=nil{fmt.Fprintln(errOut,"initialize:",err);return 1};if err:=renderResult(out,opts.json,result);err!=nil{fmt.Fprintln(errOut,err);return 1};return 0}
func runStatus(ctx context.Context,svc *application.Service,opts globalOptions,out,errOut io.Writer)int{secret,err:=ownerSecret(errOut);if err!=nil{fmt.Fprintln(errOut,"read owner passphrase:",err);return 1};defer wipe(secret);result,err:=svc.Inspect(ctx,secret);if err!=nil{fmt.Fprintln(errOut,"status:",err);return 1};if err:=renderResult(out,opts.json,result);err!=nil{fmt.Fprintln(errOut,err);return 1};return 0}
func runProject(ctx context.Context,svc *application.Service,args []string,opts globalOptions,out,errOut io.Writer)int{if len(args)==0{fmt.Fprintln(errOut,"project requires create or show");return 2};switch args[0]{case"create":fs:=flag.NewFlagSet("project create",flag.ContinueOnError);fs.SetOutput(errOut);label:=fs.String("label","","Project display label");objective:=fs.String("objective","","Project objective summary");if err:=fs.Parse(args[1:]);err!=nil{return 2};secret,err:=ownerSecret(errOut);if err!=nil{fmt.Fprintln(errOut,err);return 1};defer wipe(secret);p,err:=svc.CreateProject(ctx,secret,application.CreateProjectInput{DisplayLabel:*label,ObjectiveSummary:*objective});if err!=nil{fmt.Fprintln(errOut,"project create:",err);return 1};if err:=renderResult(out,opts.json,p);err!=nil{fmt.Fprintln(errOut,err);return 1};return 0;case"show":fs:=flag.NewFlagSet("project show",flag.ContinueOnError);fs.SetOutput(errOut);id:=fs.String("project","","Project ID");if err:=fs.Parse(args[1:]);err!=nil{return 2};secret,err:=ownerSecret(errOut);if err!=nil{fmt.Fprintln(errOut,err);return 1};defer wipe(secret);view,err:=svc.ShowProject(ctx,secret,project.ProjectID(*id));if err!=nil{fmt.Fprintln(errOut,"project show:",err);return 1};if err:=renderResult(out,opts.json,view);err!=nil{fmt.Fprintln(errOut,err);return 1};return 0;default:fmt.Fprintf(errOut,"unknown project command %q\n",args[0]);return 2}}
func wipe(v []byte){for i:=range v{v[i]=0}}
