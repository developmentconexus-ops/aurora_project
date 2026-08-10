package cli

import("fmt";"io";"os")
const helpText=`Aurora Sovereign Core

Usage:
  aurora [--data-dir <path>] [--json] init
  aurora [--data-dir <path>] [--json] status
  aurora [--data-dir <path>] [--json] project create --label <label> --objective <summary>
  aurora [--data-dir <path>] [--json] project show --project <id>
  aurora [--data-dir <path>] [--json] project set-state --project <id> --expected <n|none> --kind <kind> --summary <summary> [--payload-json <json>]
  aurora [--data-dir <path>] [--json] authority show|grant|revoke|revalidate ...
  aurora [--data-dir <path>] [--json] export --out <file.aurora.age>
  aurora [--json] restore --in <file.aurora.age> --target-data-dir <fresh-path>
  aurora --help
`
type globalOptions struct{dataDir string;json bool}
func Run(args []string,out,errOut io.Writer)int{if len(args)==1&&(args[0]=="--help"||args[0]=="-h"){fmt.Fprint(out,helpText);return 0};opts,rest,err:=parseGlobals(args);if err!=nil{fmt.Fprintln(errOut,err);return 2};if len(rest)==0{fmt.Fprintln(errOut,"missing command; use --help");return 2};if rest[0]!="restore"&&opts.dataDir==""{opts.dataDir,err=defaultDataDir();if err!=nil{fmt.Fprintln(errOut,err);return 1}};return runCommand(rest,opts,out,errOut)}
func parseGlobals(args []string)(globalOptions,[]string,error){var opts globalOptions;var rest []string;for i:=0;i<len(args);i++{switch args[i]{case"--data-dir":if i+1>=len(args){return opts,nil,fmt.Errorf("--data-dir requires a path")};i++;opts.dataDir=args[i];case"--json":opts.json=true;case"--help","-h":return opts,[]string{"help"},nil;default:rest=append(rest,args[i:]...);i=len(args)}};return opts,rest,nil}
func defaultDataDir()(string,error){if v:=os.Getenv("AURORA_DATA_DIR");v!=""{return v,nil};home,err:=os.UserHomeDir();if err!=nil{return "",fmt.Errorf("resolve home directory: %w",err)};return home+string(os.PathSeparator)+".aurora",nil}
