package cli

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestRevisionBoundTransitionsRejectStaleWithoutMutation(t *testing.T) {
	dataDir := t.TempDir()
	secrets := fixedSecretReader{secret: []byte("fixture-owner-passphrase")}
	runJSON := func(args ...string) (map[string]any, int, string) {
		all := append([]string{"--json", "--data-dir", dataDir}, args...)
		var out, errOut bytes.Buffer
		code := runWithSecretReader(all, &out, &errOut, secrets)
		var doc map[string]any
		if out.Len() != 0 {
			if err := json.Unmarshal(out.Bytes(), &doc); err != nil {
				t.Fatalf("%v output is not JSON: %q: %v", args, out.String(), err)
			}
		}
		return doc, code, errOut.String()
	}
	if _, code, stderr := runJSON("init"); code != 0 {
		t.Fatalf("init code=%d stderr=%q", code, stderr)
	}
	created, code, stderr := runJSON("project", "create", "--label", "Projeto Fonte", "--objective", "Construir fonte")
	if code != 0 {
		t.Fatalf("create code=%d stderr=%q", code, stderr)
	}
	projectID := created["project_id"].(string)

	r1, code, stderr := runJSON("project", "set-state", "--project", projectID, "--expected", "none", "--kind", "WORK_NOTE", "--summary", "R1", "--payload", `{"aurora_id":"AUR-FAKE","authority":{"grant":"pretend"}}`)
	if code != 0 {
		t.Fatalf("R1 code=%d stderr=%q", code, stderr)
	}
	if r1["state_revision"].(float64) != 1 {
		t.Fatalf("R1=%v", r1)
	}

	r2, code, stderr := runJSON("project", "set-state", "--project", projectID, "--expected", "1", "--kind", "WORK_NOTE", "--summary", "R2", "--payload", `{"value":2}`)
	if code != 0 {
		t.Fatalf("R2 code=%d stderr=%q", code, stderr)
	}
	if r2["state_revision"].(float64) != 2 {
		t.Fatalf("R2=%v", r2)
	}

	_, code, _ = runJSON("project", "set-state", "--project", projectID, "--expected", "1", "--kind", "WORK_NOTE", "--summary", "STALE", "--payload", `{"value":999}`)
	if code == 0 {
		t.Fatal("stale transition unexpectedly accepted")
	}

	shown, code, stderr := runJSON("project", "show", "--project", projectID)
	if code != 0 {
		t.Fatalf("show code=%d stderr=%q", code, stderr)
	}
	if shown["current_state_revision"].(float64) != 2 {
		t.Fatalf("current pointer changed after stale reject: %v", shown)
	}
	current := shown["current_state"].(map[string]any)
	if current["state_revision"].(float64) != 2 {
		t.Fatalf("current state changed after stale reject: %v", shown)
	}

	status, code, stderr := runJSON("status")
	if code != 0 {
		t.Fatalf("status code=%d stderr=%q", code, stderr)
	}
	if status["aurora_id"] == "AUR-FAKE" {
		t.Fatalf("opaque payload mutated Aurora identity: %v", status)
	}
	if status["governing_generation"].(float64) != 4 {
		t.Fatalf("stale rejection advanced generation: %v", status)
	}
}

func TestTransitionRejectsMismatchedNextActionProject(t *testing.T) {
	dataDir := t.TempDir()
	secrets := fixedSecretReader{secret: []byte("fixture-owner-passphrase")}
	var out, errOut bytes.Buffer
	if code := runWithSecretReader([]string{"--json", "--data-dir", dataDir, "init"}, &out, &errOut, secrets); code != 0 {
		t.Fatal(errOut.String())
	}
	out.Reset()
	errOut.Reset()
	if code := runWithSecretReader([]string{"--json", "--data-dir", dataDir, "project", "create", "--label", "A", "--objective", "A"}, &out, &errOut, secrets); code != 0 {
		t.Fatal(errOut.String())
	}
	var created map[string]any
	if err := json.Unmarshal(out.Bytes(), &created); err != nil {
		t.Fatal(err)
	}
	id := created["project_id"].(string)
	out.Reset()
	errOut.Reset()
	code := runWithSecretReader([]string{"--json", "--data-dir", dataDir, "project", "set-state", "--project", id, "--expected", "none", "--kind", "WORK_NOTE", "--summary", "R1", "--payload", `{}`, "--next-action", "BUILD", "--next-project", "PRJ-OTHER"}, &out, &errOut, secrets)
	if code == 0 {
		t.Fatal("mismatched next-action Project unexpectedly accepted")
	}
}
