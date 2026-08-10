package cli

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestAuthorityGrantRestartRevokeAndProjection(t *testing.T) {
	dataDir := t.TempDir()
	secrets := fixedSecretReader{secret: []byte("fixture-owner-passphrase")}
	run := func(args ...string) (map[string]any, int, string) {
		all := append([]string{"--json", "--data-dir", dataDir}, args...)
		var out, errOut bytes.Buffer
		code := runWithSecretReader(all, &out, &errOut, secrets)
		var doc map[string]any
		if out.Len() != 0 {
			if err := json.Unmarshal(out.Bytes(), &doc); err != nil {
				t.Fatalf("%v json: %v %q", args, err, out.String())
			}
		}
		return doc, code, errOut.String()
	}
	if _, code, e := run("init"); code != 0 {
		t.Fatal(e)
	}
	p, code, e := run("project", "create", "--label", "A", "--objective", "A")
	if code != 0 {
		t.Fatal(e)
	}
	pid := p["project_id"].(string)
	if _, code, e := run("project", "set-state", "--project", pid, "--expected", "none", "--kind", "WORK_NOTE", "--summary", "R1", "--payload", `{}`, "--next-action", "BUILD"); code != 0 {
		t.Fatal(e)
	}
	before, code, e := run("authority", "show", "--project", pid, "--action", "BUILD")
	if code != 0 {
		t.Fatal(e)
	}
	if before["decision"] != "BLOCKED" {
		t.Fatalf("before grant=%v", before)
	}
	grant, code, e := run("authority", "grant", "--project", pid, "--action", "BUILD")
	if code != 0 {
		t.Fatal(e)
	}
	aid := grant["authority_id"].(string)
	after, code, e := run("authority", "show", "--project", pid, "--action", "BUILD")
	if code != 0 {
		t.Fatal(e)
	}
	if after["decision"] != "PERMITTED" {
		t.Fatalf("after grant=%v", after)
	}
	if _, code, e := run("authority", "revoke", "--authority", aid); code != 0 {
		t.Fatal(e)
	}
	revoked, code, e := run("authority", "show", "--project", pid, "--action", "BUILD")
	if code != 0 {
		t.Fatal(e)
	}
	if revoked["decision"] != "BLOCKED" {
		t.Fatalf("after revoke=%v", revoked)
	}
}
