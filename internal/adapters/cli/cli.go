package cli

import (
	"context"
	"fmt"
	"io"
)

const helpText=`Aurora Sovereign Core

Usage:
  aurora [--data-dir PATH] [--json] init
  aurora [--data-dir PATH] [--json] status
  aurora [--data-dir PATH] [--json] project create --label LABEL --objective OBJECTIVE
  aurora [--data-dir PATH] [--json] project show --project PROJECT_ID
  aurora [--data-dir PATH] [--json] project set-state --project PROJECT_ID --expected REV|none --kind KIND --summary SUMMARY [--payload JSON]
  aurora --help
`

func Run(args []string,out,errOut io.Writer)int{return runWithSecretReader(args,out,errOut,newTerminalSecretReader(errOut))}
func runWithSecretReader(args []string,out,errOut io.Writer,secrets SecretReader)int{if len(args)==0||args[0]=="--help"||args[0]=="-h"{_,_=io.WriteString(out,helpText);return 0};opts,rest,err:=parseGlobal(args);if err!=nil{_,_=fmt.Fprintln(errOut,err);return 2};if len(rest)==0{_,_=fmt.Fprintln(errOut,"missing command");return 2};service,state,err:=newService(opts.dataDir);if err!=nil{_,_=fmt.Fprintln(errOut,err);return 1};defer state.Close();passphrase,err:=secrets.ReadSecret("Owner passphrase: ");if err!=nil{_,_=fmt.Fprintln(errOut,err);return 1};defer clear(passphrase);switch rest[0]{case "init":if len(rest)!=1{return usageError(errOut,"init takes no arguments")};r,err:=service.Initialize(context.Background(),passphrase);if err!=nil{return runtimeError(errOut,err)};if err:=renderInitialize(out,r,opts.json);err!=nil{return runtimeError(errOut,err)};return 0;case "status":if len(rest)!=1{return usageError(errOut,"status takes no arguments")};r,err:=service.Inspect(context.Background(),passphrase);if err!=nil{return runtimeError(errOut,err)};if err:=renderInspect(out,r,opts.json);err!=nil{return runtimeError(errOut,err)};return 0;case "project":return runProject(context.Background(),service,passphrase,rest[1:],out,errOut,opts.json);default:_,_=fmt.Fprintf(errOut,"unknown command %q\n",rest[0]);return 2}}
func usageError(w io.Writer,err string)int{_,_=fmt.Fprintln(w,err);return 2}
func runtimeError(w io.Writer,err error)int{_,_=fmt.Fprintln(w,err);return 1}
