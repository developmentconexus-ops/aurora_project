package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestMigrateRequiresFilesWithoutOpeningCurrentState(t *testing.T) {
	dataDir := filepath.Join(t.TempDir(), "must-not-be-created")
	t.Setenv("AURORA_DATA_DIR", dataDir)

	var out, errOut bytes.Buffer
	code := Run([]string{"migrate"}, &out, &errOut)
	if code != 2 {
		t.Fatalf("code=%d want 2; stderr=%q", code, errOut.String())
	}
	if _, err := os.Stat(dataDir); !os.IsNotExist(err) {
		t.Fatalf("migrate touched current data dir: stat err=%v", err)
	}
	if got := errOut.String(); got == "" {
		t.Fatal("expected migrate usage error")
	}
}
