package process_test

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"
)

func TestRealBinaryExportRestoreRequiresRevalidation(t *testing.T) {
	root, err := filepath.Abs("../..")
	if err != nil { t.Fatal(err) }
	bin := filepath.Join(t.TempDir(), "aurora")
	if runtime.GOOS == "windows" { bin += ".exe" }
	build := exec.Command("go", "build", "-o", bin, "./cmd/aurora")
	build.Dir = root
	if out, err := build.CombinedOutput(); err != nil { t.Fatalf("build: %v\n%s", err, out) }

	source := filepath.Join(t.TempDir(), "source")
	target := filepath.Join(t.TempDir(), "restored")
	exportPath := filepath.Join(t.TempDir(), "m0.aurora.age")
	owner := "owner-process-passphrase"
	exportSecret := "export-process-passphrase"

	init := runJSONInput(t, bin, source, owner+"\n", "init")
	projectObj := runJSONInput(t, bin, source, owner+"\n", "project", "create", "--label", "Portable Project", "--objective", "Restore proof")
	pid := stringField(t, projectObj, "project_id")
	state := runJSONInput(t, bin, source, owner+"\n", "project", "set-state", "--project", pid, "--expected", "none", "--kind", "note", "--summary", "R1", "--next-action", "review", "--next-summary", "Review restored state", "--next-authority-action", "review")
	if numberField(t, state, "state_revision") != 1 { t.Fatalf("state=%v", state) }

	auth := runJSONInput(t, bin, source, owner+"\n", "authority", "show")
	authState := mapField(t, auth, "state")
	expected := numberField(t, authState, "revision")
	runJSONInput(t, bin, source, owner+"\n", "authority", "grant", "--expected", formatInt(expected), "--project", pid, "--action", "review")

	exported := runJSONInput(t, bin, source, owner+"\n"+exportSecret+"\n", "export", "--out", exportPath)
	if stringField(t, exported, "export_id") == "" { t.Fatalf("export=%v", exported) }
	if info, err := os.Stat(exportPath); err != nil || info.Size() == 0 { t.Fatalf("export file: info=%v err=%v", info, err) }

	restored := runRestoreJSON(t, bin, exportPath, target, exportSecret+"\n"+owner+"\n")
	if stringField(t, restored, "aurora_id") != stringField(t, init, "aurora_id") { t.Fatalf("restore identity mismatch init=%v restore=%v", init, restored) }
	if stringField(t, restored, "trust_status") != "REVALIDATION_REQUIRED" { t.Fatalf("restore=%v", restored) }

	if _, stderr, err := runCommandInput(bin, append(globalArgs(target), "project", "show", "--project", pid), owner+"\n"); err == nil {
		t.Fatalf("project show unexpectedly permitted before revalidation; stderr=%s", stderr)
	}

	runJSONInput(t, bin, target, owner+"\n", "authority", "revalidate")
	shown := runJSONInput(t, bin, target, owner+"\n", "project", "show", "--project", pid)
	current := mapField(t, shown, "current_state")
	if numberField(t, current, "state_revision") != 1 { t.Fatalf("restored current state=%v", shown) }
	next := mapField(t, shown, "next_safe_action")
	if stringField(t, next, "decision") != "PERMITTED" { t.Fatalf("restored next action=%v", next) }
}

func globalArgs(data string) []string { return []string{"--data-dir", data, "--json"} }

func runJSONInput(t *testing.T, bin, data, stdin string, args ...string) map[string]any {
	t.Helper()
	stdout, stderr, err := runCommandInput(bin, append(globalArgs(data), args...), stdin)
	if err != nil { t.Fatalf("run %v: %v\nstdout=%s\nstderr=%s", args, err, stdout, stderr) }
	var out map[string]any
	if err := json.Unmarshal([]byte(stdout), &out); err != nil { t.Fatalf("decode %v JSON: %v\nstdout=%s\nstderr=%s", args, err, stdout, stderr) }
	return out
}

func runRestoreJSON(t *testing.T, bin, exportPath, target, stdin string) map[string]any {
	t.Helper()
	stdout, stderr, err := runCommandInput(bin, []string{"--json", "restore", "--in", exportPath, "--target-data-dir", target}, stdin)
	if err != nil { t.Fatalf("restore: %v\nstdout=%s\nstderr=%s", err, stdout, stderr) }
	var out map[string]any
	if err := json.Unmarshal([]byte(stdout), &out); err != nil { t.Fatalf("decode restore JSON: %v\nstdout=%s\nstderr=%s", err, stdout, stderr) }
	return out
}

func runCommandInput(bin string, args []string, stdin string) (string, string, error) {
	cmd := exec.Command(bin, args...)
	cmd.Stdin = bytes.NewBufferString(stdin)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func formatInt(v float64) string { return strconv.FormatUint(uint64(v), 10) }
