package cli

import (
	"context"
	"fmt"
	"io"
)

const helpText = `Aurora Sovereign Core

Usage:
  aurora [--data-dir PATH] [--json] init
  aurora [--data-dir PATH] [--json] status
  aurora --help
`

func Run(args []string, out, errOut io.Writer) int {
	return runWithSecretReader(args, out, errOut, newTerminalSecretReader(errOut))
}

func runWithSecretReader(args []string, out, errOut io.Writer, secrets SecretReader) int {
	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		_, _ = io.WriteString(out, helpText)
		return 0
	}
	opts, rest, err := parseGlobal(args)
	if err != nil {
		_, _ = fmt.Fprintln(errOut, err)
		return 2
	}
	if len(rest) != 1 {
		_, _ = fmt.Fprintln(errOut, "expected exactly one command")
		return 2
	}
	service, state, err := newService(opts.dataDir)
	if err != nil {
		_, _ = fmt.Fprintln(errOut, err)
		return 1
	}
	defer state.Close()
	passphrase, err := secrets.ReadSecret("Owner passphrase: ")
	if err != nil {
		_, _ = fmt.Fprintln(errOut, err)
		return 1
	}
	defer clear(passphrase)
	switch rest[0] {
	case "init":
		result, err := service.Initialize(context.Background(), passphrase)
		if err != nil {
			_, _ = fmt.Fprintln(errOut, err)
			return 1
		}
		if err := renderInitialize(out, result, opts.json); err != nil {
			_, _ = fmt.Fprintln(errOut, err)
			return 1
		}
		return 0
	case "status":
		result, err := service.Inspect(context.Background(), passphrase)
		if err != nil {
			_, _ = fmt.Fprintln(errOut, err)
			return 1
		}
		if err := renderInspect(out, result, opts.json); err != nil {
			_, _ = fmt.Fprintln(errOut, err)
			return 1
		}
		return 0
	default:
		_, _ = fmt.Fprintf(errOut, "unknown command %q\n", rest[0])
		return 2
	}
}
