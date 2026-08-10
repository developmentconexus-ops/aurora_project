package cli

import (
	"fmt"
	"io"
	"os"
)

const helpText = `Aurora Sovereign Core

Usage:
  aurora [--data-dir <path>] [--json] init
  aurora [--data-dir <path>] [--json] status
  aurora --help
`

type globalOptions struct {
	dataDir string
	json    bool
}

// Run executes the M0 operator adapter. Domain/application code never renders terminal text.
func Run(args []string, out, errOut io.Writer) int {
	if len(args) == 1 && (args[0] == "--help" || args[0] == "-h") {
		fmt.Fprint(out, helpText)
		return 0
	}
	opts, command, err := parseGlobals(args)
	if err != nil {
		fmt.Fprintln(errOut, err)
		return 2
	}
	if command == "" {
		fmt.Fprintln(errOut, "missing command; use --help")
		return 2
	}
	if opts.dataDir == "" {
		opts.dataDir, err = defaultDataDir()
		if err != nil {
			fmt.Fprintln(errOut, err)
			return 1
		}
	}
	return runCommand(command, opts, out, errOut)
}

func parseGlobals(args []string) (globalOptions, string, error) {
	var opts globalOptions
	var command string
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--data-dir":
			if i+1 >= len(args) { return opts, "", fmt.Errorf("--data-dir requires a path") }
			i++
			opts.dataDir = args[i]
		case "--json":
			opts.json = true
		case "--help", "-h":
			return opts, "help", nil
		default:
			if command != "" { return opts, "", fmt.Errorf("unexpected argument %q", args[i]) }
			command = args[i]
		}
	}
	return opts, command, nil
}

func defaultDataDir() (string, error) {
	if v := os.Getenv("AURORA_DATA_DIR"); v != "" { return v, nil }
	home, err := os.UserHomeDir()
	if err != nil { return "", fmt.Errorf("resolve home directory: %w", err) }
	return home + string(os.PathSeparator) + ".aurora", nil
}
