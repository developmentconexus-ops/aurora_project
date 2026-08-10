package cli

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestAnchorLagBlocksMutationUntilOwnerReconcile(t *testing.T) {
	dataDir := t.TempDir()
	secrets := fixedSecretReader{secret: []byte("fixture-owner-passphrase")}
	run := func(args ...string) (map[string]any, int, string) {
		all := append([]string{"--json", "--data-dir", dataDir}, args...)
		var out, errOut bytes.Buffer
		code := runWithSecretReader(all, &out, &errOut, secrets)
		var doc map[string]any
		if out.Len() != 0 {
			if err := json.Unmarshal(out.Bytes(), &doc); err != nil {
				t.Fatalf("json %v: %v %q", args, err, out.String())
			}
		}
		return doc, code, errOut.String()
	}
	if _, code, e := run("init"); code != 0 {
		t.Fatal(e)
	}
	anchorPath := filepath.Join(dataDir, "trust", "owner-anchor.json")
	oldAnchor, err := os.ReadFile(anchorPath)
	if err != nil {
		t.Fatal(err)
	}
	if _, code, e := run("project", "create", "--label", "A", "--objective", "A"); code != 0 {
		t.Fatal(e)
	}
	if err := os.WriteFile(anchorPath, oldAnchor, 0o600); err != nil {
		t.Fatal(err)
	}
	status, code, e := run("status")
	if code != 0 {
		t.Fatalf("status code=%d err=%q", code, e)
	}
	if status["trust_state"] != "ANCHOR_LAG" {
		t.Fatalf("status=%v", status)
	}
	if _, code, _ := run("project", "create", "--label", "B", "--objective", "B"); code == 0 {
		t.Fatal("mutation succeeded during ANCHOR_LAG")
	}
	if _, code, e := run("authority", "reconcile-anchor"); code != 0 {
		t.Fatalf("reconcile code=%d err=%q", code, e)
	}
	status, code, e = run("status")
	if code != 0 {
		t.Fatal(e)
	}
	if status["trust_state"] != "NORMAL" {
		t.Fatalf("after reconcile=%v", status)
	}
}
