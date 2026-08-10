package cli

import (
	"fmt"
	"io"
)

const helpText = `Aurora Sovereign Core

Usage:
  aurora --help
`

// Run executes the minimal M0 operator adapter.
func Run(args []string, out, errOut io.Writer) int {
	if len(args) == 1 && (args[0] == "--help" || args[0] == "-h") {
		fmt.Fprint(out, helpText)
		return 0
	}
	fmt.Fprintln(errOut, "unknown command; use --help")
	return 2
}
