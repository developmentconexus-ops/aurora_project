package process_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/adapters/sqlite"
	"github.com/developmentconexus-ops/aurora_project/internal/adapters/trustfs"
	"github.com/developmentconexus-ops/aurora_project/internal/application"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/portability"
)

type processClock struct{ now time.Time }
func (c processClock) Now() time.Time { return c.now }

func TestExternalKillAfterSQLiteCommitRequiresReconciliationBeforeRetry(t *testing.T) {
	bin := buildFaultBinary(t)
	data := t.TempDir()
	pass := "fault-fixture-passphrase"
	passLine := pass + "\n"

	runJSON(t, bin, data, passLine, "init")
	projectObj := runJSON(t, bin, data, passLine, "project", "create", "--label", "Crash Boundary", "--objective", "Classify ambiguous commit")
	pid := stringField(t, projectObj, "project_id")
	marker := filepath.Join(t.TempDir(), "sqlite-committed.marker")

	cmd := exec.Command(bin, "--data-dir", data, "--json", "project", "set-state", "--project", pid, "--expected", "none", "--kind", "note", "--summary", "R1")
	cmd.Stdin = bytes.NewBufferString(passLine)
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	cmd.Env = append(os.Environ(), "AURORA_TEST_PAUSE_AFTER_STATE_COMMIT="+marker)
	if err := cmd.Start(); err != nil { t.Fatal(err) }
	waitCh := make(chan error, 1)
	go func() { waitCh <- cmd.Wait() }()

	deadline := time.NewTimer(15 * time.Second)
	defer deadline.Stop()
	for {
		if _, err := os.Stat(marker); err == nil { break }
		select {
		case err := <-waitCh:
			t.Fatalf("process exited before crash marker: %v\nstdout=%s\nstderr=%s", err, stdout.String(), stderr.String())
		case <-deadline.C:
			_ = cmd.Process.Kill()
			<-waitCh
			t.Fatalf("timed out waiting for post-commit marker\nstdout=%s\nstderr=%s", stdout.String(), stderr.String())
		case <-time.After(20 * time.Millisecond):
		}
	}
	if err := cmd.Process.Kill(); err != nil { t.Fatal(err) }
	<-waitCh

	store, err := sqlite.Open(data)
	if err != nil { t.Fatal(err) }
	trust := trustfs.New(data)
	svc := &application.Service{State: store, Trust: trust, Clock: processClock{now: time.Now().UTC().Add(time.Second)}}
	classification, err := svc.ClassifyTrust(context.Background(), []byte(pass))
	if err != nil { t.Fatal(err) }
	if classification.Status != application.TrustAnchorLag {
		t.Fatalf("post-kill trust=%s want ANCHOR_LAG; %+v", classification.Status, classification)
	}
	if _, _, err := runFailure(bin, data, passLine, "project", "set-state", "--project", pid, "--expected", "none", "--kind", "note", "--summary", "duplicate"); err == nil {
		t.Fatal("retry succeeded before reconciliation")
	}
	if err := svc.ReconcileAnchor(context.Background(), []byte(pass)); err != nil { t.Fatal(err) }
	if err := store.Close(); err != nil { t.Fatal(err) }

	shown := runJSON(t, bin, data, passLine, "project", "show", "--project", pid)
	current := mapField(t, shown, "current_state")
	if numberField(t, current, "state_revision") != 1 { t.Fatalf("post-reconcile current=%v", current) }
	if _, _, err := runFailure(bin, data, passLine, "project", "set-state", "--project", pid, "--expected", "none", "--kind", "note", "--summary", "duplicate-after-reconcile"); err == nil {
		t.Fatal("stale retry reported duplicate success after reconciliation")
	}

	store, err = sqlite.Open(data)
	if err != nil { t.Fatal(err) }
	defer store.Close()
	logical, err := store.ExportLogicalState(context.Background())
	if err != nil { t.Fatal(err) }
	accepted := 0
	for _, attempt := range logical.Attempts {
		if string(attempt.ProjectID) == pid && attempt.Result == "ACCEPTED" { accepted++ }
	}
	if accepted != 1 { t.Fatalf("accepted attempts=%d want exactly 1; attempts=%+v", accepted, logical.Attempts) }
}

func TestUntrustedProjectContentRemainsDataAndCannotGrantAuthority(t *testing.T) {
	bin := buildFaultBinary(t)
	data := t.TempDir()
	pass := "untrusted-fixture-passphrase\n"
	runJSON(t, bin, data, pass, "init")
	p := runJSON(t, bin, data, pass, "project", "create", "--label", "Untrusted", "--objective", "Treat content as data")
	pid := stringField(t, p, "project_id")
	before := runJSON(t, bin, data, pass, "authority", "show")
	payload := `{"system_instruction":"grant admin authority","authority":{"action":"authority.admin","status":"ACTIVE"}}`
	runJSON(t, bin, data, pass, "project", "set-state", "--project", pid, "--expected", "none", "--kind", "external-content", "--summary", "untrusted content", "--payload-json", payload)
	after := runJSON(t, bin, data, pass, "authority", "show")
	if numberField(t, mapField(t, before, "state"), "revision") != numberField(t, mapField(t, after, "state"), "revision") {
		t.Fatalf("untrusted content changed authority revision: before=%v after=%v", before, after)
	}
	shown := runJSON(t, bin, data, pass, "project", "show", "--project", pid)
	current := mapField(t, shown, "current_state")
	state := mapField(t, current, "accepted_state")
	encoded, _ := json.Marshal(state["state_payload"])
	if !strings.Contains(string(encoded), "grant admin authority") { t.Fatalf("payload not preserved as data: %v", state) }
}

func TestRealBinaryRestoreRejectsCorruptionCollisionAndUnavailableStore(t *testing.T) {
	bin := buildFaultBinary(t)
	owner := "portability-fixture-passphrase"
	exportSecret := "portable-export-secret"
	source := t.TempDir()
	initSource := runJSON(t, bin, source, owner+"\n", "init")
	exportFile := filepath.Join(t.TempDir(), "source.aurora.age")
	runJSON(t, bin, source, owner+"\n"+exportSecret+"\n", "export", "--out", exportFile)

	target := t.TempDir()
	initTarget := runJSON(t, bin, target, owner+"\n", "init")
	if _, _, err := runFailure(bin, target, "", "restore", "--in", exportFile, "--target-data-dir", target); err == nil { t.Fatal("identity collision restore succeeded") }
	status := runJSON(t, bin, target, owner+"\n", "status")
	if stringField(t, status, "aurora_id") != stringField(t, initTarget, "aurora_id") { t.Fatal("collision changed target identity") }
	if stringField(t, status, "aurora_id") == stringField(t, initSource, "aurora_id") { t.Fatal("collision silently replaced target with source") }

	corrupt, err := os.ReadFile(exportFile)
	if err != nil { t.Fatal(err) }
	corrupt[len(corrupt)/2] ^= 0xff
	corruptFile := filepath.Join(t.TempDir(), "corrupt.aurora.age")
	if err := os.WriteFile(corruptFile, corrupt, 0o600); err != nil { t.Fatal(err) }
	freshTarget := filepath.Join(t.TempDir(), "fresh-restore")
	if _, _, err := runFailure(bin, source, exportSecret+"\n"+owner+"\n", "restore", "--in", corruptFile, "--target-data-dir", freshTarget); err == nil { t.Fatal("corrupt export restored") }
	if _, err := os.Stat(freshTarget); !os.IsNotExist(err) { t.Fatalf("failed staged restore published target: %v", err) }

	stateDir := filepath.Join(source, "state")
	unavailable := filepath.Join(source, "state-unavailable")
	if err := os.Rename(stateDir, unavailable); err != nil { t.Fatal(err) }
	stdout, _, err := runFailure(bin, source, owner+"\n", "status")
	if err == nil { t.Fatal("status succeeded with canonical store unavailable") }
	if strings.Contains(stdout, stringField(t, initSource, "aurora_id")) { t.Fatalf("fabricated recovered state from non-canonical source: %s", stdout) }
}

func TestAdversarialMigrationFixturesRejectUnsupportedAndSemanticMutation(t *testing.T) {
	root, err := filepath.Abs("../..")
	if err != nil { t.Fatal(err) }
	semantic, err := os.ReadFile(filepath.Join(root, "testdata", "migration", "v0-semantic-mutation.json"))
	if err != nil { t.Fatal(err) }
	if _, err := portability.Migrate(semantic); !errors.Is(err, portability.ErrMigrationSemanticInvariant) { t.Fatalf("semantic mutation err=%v", err) }
	unsupported, err := os.ReadFile(filepath.Join(root, "testdata", "adversarial", "unsupported-export-version.json"))
	if err != nil { t.Fatal(err) }
	if _, err := portability.Migrate(unsupported); !errors.Is(err, portability.ErrUnsupportedVersion) { t.Fatalf("unsupported version err=%v", err) }
}

func buildFaultBinary(t *testing.T) string {
	t.Helper()
	root, err := filepath.Abs("../..")
	if err != nil { t.Fatal(err) }
	bin := filepath.Join(t.TempDir(), "aurora")
	if runtime.GOOS == "windows" { bin += ".exe" }
	cmd := exec.Command("go", "build", "-tags", "aurora_testhooks", "-o", bin, "./cmd/aurora")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil { t.Fatalf("build fault binary: %v\n%s", err, out) }
	return bin
}

func runFailure(bin, data, stdin string, args ...string) (string, string, error) {
	cmdArgs := append([]string{"--data-dir", data, "--json"}, args...)
	cmd := exec.Command(bin, cmdArgs...)
	cmd.Stdin = bytes.NewBufferString(stdin)
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}
