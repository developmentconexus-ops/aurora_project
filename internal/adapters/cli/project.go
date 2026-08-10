package cli

import (
	"context"
	"errors"
	"io"

	"github.com/developmentconexus-ops/aurora_project/internal/application"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
)

func runProject(ctx context.Context,s *application.Service,pass []byte,args []string,out,errOut io.Writer,asJSON bool)int{
	if len(args)==0{return usageError(errOut,"project requires create or show")}
	switch args[0]{
	case "create":
		vals,err:=parseNamed(args[1:],map[string]bool{"--label":true,"--objective":true});if err!=nil{return usageError(errOut,err.Error())}
		if vals["--label"]==""||vals["--objective"]==""{return usageError(errOut,"--label and --objective are required")}
		p,err:=s.CreateProject(ctx,pass,application.CreateProjectInput{DisplayLabel:vals["--label"],ObjectiveSummary:vals["--objective"]});if err!=nil{return runtimeError(errOut,err)}
		if err:=renderProject(out,p,asJSON);err!=nil{return runtimeError(errOut,err)};return 0
	case "show":
		vals,err:=parseNamed(args[1:],map[string]bool{"--project":true});if err!=nil{return usageError(errOut,err.Error())};id:=vals["--project"];if id==""{return usageError(errOut,"--project is required")}
		p,err:=s.ShowProject(ctx,pass,project.ProjectID(id));if err!=nil{return runtimeError(errOut,err)};if err:=renderProject(out,p,asJSON);err!=nil{return runtimeError(errOut,err)};return 0
	default:return usageError(errOut,"unknown project command")
	}
}

func parseNamed(args []string,allowed map[string]bool)(map[string]string,error){
	vals:=map[string]string{}
	for i:=0;i<len(args);i++{key:=args[i];if !allowed[key]{return nil,errors.New("unknown option "+key)};if i+1>=len(args){return nil,errors.New(key+" requires a value")};i++;vals[key]=args[i]}
	return vals,nil
}
