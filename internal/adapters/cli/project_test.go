package cli

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestProjectCreateThenFreshShowPreservesIdentity(t *testing.T) {
	dataDir := t.TempDir()
	secrets := fixedSecretReader{secret: []byte("fixture-owner-passphrase")}
	var out, errOut bytes.Buffer
	if code := runWithSecretReader([]string{"--json", "--data-dir", dataDir, "init"}, &out, &errOut, secrets); code != 0 {
		t.Fatalf("init code=%d stderr=%q", code, errOut.String())
	}

	out.Reset()
	errOut.Reset()
	if code := runWithSecretReader([]string{"--json", "--data-dir", dataDir, "project", "create", "--label", "Projeto Fonte", "--objective", "Construir fonte"}, &out, &errOut, secrets); code != 0 {
		t.Fatalf("create code=%d stderr=%q", code, errOut.String())
	}
	var created map[string]any
	if err := json.Unmarshal(out.Bytes(), &created); err != nil {
		t.Fatal(err)
	}
	id, _ := created["project_id"].(string)
	if id == "" {
		t.Fatalf("missing project_id: %s", out.String())
	}

	out.Reset()
	errOut.Reset()
	if code := runWithSecretReader([]string{"--json", "--data-dir", dataDir, "project", "show", "--project", id}, &out, &errOut, secrets); code != 0 {
		t.Fatalf("show code=%d stderr=%q", code, errOut.String())
	}
	var shown map[string]any
	if err := json.Unmarshal(out.Bytes(), &shown); err != nil {
		t.Fatal(err)
	}
	if shown["project_id"] != id || shown["display_label"] != "Projeto Fonte" || shown["objective_summary"] != "Construir fonte" {
		t.Fatalf("shown=%v created=%v", shown, created)
	}
}
