package cli

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"path/filepath"
	"testing"

	_ "modernc.org/sqlite"
)

func TestRecoverCurrentStateReconstructsProjectAuthorityAndNextAction(t *testing.T) {
	dataDir := t.TempDir()
	secrets := fixedSecretReader{secret: []byte("fixture-owner-passphrase")}
	run := func(args ...string) (map[string]any, int, string) {
		all := append([]string{"--json", "--data-dir", dataDir}, args...)
		var out, errOut bytes.Buffer
		code := runWithSecretReader(all, &out, &errOut, secrets)
		var doc map[string]any
		if out.Len() != 0 {
			if err := json.Unmarshal(out.Bytes(), &doc); err != nil { t.Fatalf("%v json: %v %q", args, err, out.String()) }
		}
		return doc, code, errOut.String()
	}
	if _, code, e := run("init"); code != 0 { t.Fatal(e) }
	p, code, e := run("project", "create", "--label", "Recovery", "--objective", "Recover same state")
	if code != 0 { t.Fatal(e) }
	pid := p["project_id"].(string)
	if _, code, e := run("project", "set-state", "--project", pid, "--expected", "none", "--kind", "WORK_NOTE", "--summary", "R1", "--payload", `{}`, "--next-action", "BUILD"); code != 0 { t.Fatal(e) }
	if _, code, e := run("authority", "grant", "--project", pid, "--action", "BUILD"); code != 0 { t.Fatal(e) }

	recovered, code, e := run("recover")
	if code != 0 { t.Fatalf("recover code=%d err=%q", code, e) }
	if recovered["status"] != "RECOVERED" || recovered["trust_state"] != "NORMAL" { t.Fatalf("recovered=%v", recovered) }
	projects := recovered["projects"].([]any)
	if len(projects) != 1 { t.Fatalf("projects=%v", projects) }
	projectView := projects[0].(map[string]any)
	if projectView["project_id"] != pid || projectView["current_state_revision"].(float64) != 1 { t.Fatalf("project recovery=%v", projectView) }
	projection := projectView["next_safe_action"].(map[string]any)
	if projection["decision"] != "PERMITTED" { t.Fatalf("next safe action=%v", projection) }
}

func TestRecoverCurrentStateDoesNotPromoteHistoryWhenCurrentPointerIsBroken(t *testing.T) {
	dataDir := t.TempDir()
	secrets := fixedSecretReader{secret: []byte("fixture-owner-passphrase")}
	var out, errOut bytes.Buffer
	if code := runWithSecretReader([]string{"--json", "--data-dir", dataDir, "init"}, &out, &errOut, secrets); code != 0 { t.Fatal(errOut.String()) }
	out.Reset(); errOut.Reset()
	if code := runWithSecretReader([]string{"--json", "--data-dir", dataDir, "project", "create", "--label", "Broken", "--objective", "Broken pointer"}, &out, &errOut, secrets); code != 0 { t.Fatal(errOut.String()) }
	var created map[string]any
	if err := json.Unmarshal(out.Bytes(), &created); err != nil { t.Fatal(err) }
	pid := created["project_id"].(string)
	out.Reset(); errOut.Reset()
	if code := runWithSecretReader([]string{"--json", "--data-dir", dataDir, "project", "set-state", "--project", pid, "--expected", "none", "--kind", "WORK_NOTE", "--summary", "R1", "--payload", `{}`}, &out, &errOut, secrets); code != 0 { t.Fatal(errOut.String()) }

	db, err := sql.Open("sqlite", filepath.Join(dataDir, "state", "aurora.db"))
	if err != nil { t.Fatal(err) }
	if _, err := db.Exec(`UPDATE projects SET current_state_revision=99 WHERE project_id=?`, pid); err != nil { t.Fatal(err) }
	if err := db.Close(); err != nil { t.Fatal(err) }

	out.Reset(); errOut.Reset()
	code := runWithSecretReader([]string{"--json", "--data-dir", dataDir, "recover"}, &out, &errOut, secrets)
	if code != 0 { t.Fatalf("recover diagnostic code=%d err=%q", code, errOut.String()) }
	var result map[string]any
	if err := json.Unmarshal(out.Bytes(), &result); err != nil { t.Fatal(err) }
	if result["status"] != "FAILED" || result["classification"] != "DURABLE_STATE_INTEGRITY_FAILURE" { t.Fatalf("result=%v", result) }
	if projects, ok := result["projects"].([]any); ok && len(projects) != 0 { t.Fatalf("fabricated/promoted projects=%v", projects) }
}
