package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunHelp(t *testing.T) {
	var out, errOut bytes.Buffer
	code := Run([]string{"--help"}, &out, &errOut)
	if code != 0 {
		t.Fatalf("Run(--help) code = %d, want 0; stderr=%q", code, errOut.String())
	}
	if !strings.Contains(out.String(), "Aurora Sovereign Core") {
		t.Fatalf("help output %q does not contain product name", out.String())
	}
}
