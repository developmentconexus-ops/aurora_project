package cli

import (
	"fmt"
	"io"
)

const helpText = `Aurora Sovereign Core

Usage:
  aurora --help
`

func Run(args []string, out, errOut io.Writer) int {
	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		_, _ = io.WriteString(out, helpText)
		return 0
	}

	_, _ = fmt.Fprintf(errOut, "unknown command %q\n", args[0])
	return 2
}
