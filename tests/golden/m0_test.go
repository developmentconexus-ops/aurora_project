package golden_test

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

type goldenFixture struct {
	FixtureID               string `json:"fixture_id"`
	ProjectLabel            string `json:"project_label"`
	ProjectObjective        string `json:"project_objective"`
	StateKind               string `json:"state_kind"`
	StateSummary            string `json:"state_summary"`
	NextAction              string `json:"next_action"`
	NextSummary             string `json:"next_summary"`
	RequiredAuthorityAction string `json:"required_authority_action"`
}

type goldenEvidence struct {
	ProofID                    string   `json:"proof_id"`
	SourceRevision             string   `json:"source_revision"`
	BinarySHA256               string   `json:"binary_sha256"`
	DependencyLockSHA256       string   `json:"dependency_lock_sha256"`
	GoVersion                  string   `json:"go_version"`
	GOOS                       string   `json:"goos"`
	GOARCH                     string   `json:"goarch"`
	FixtureID                  string   `json:"fixture_id"`
	AuroraID                   string   `json:"aurora_id"`
	ProjectID                  string   `json:"project_id"`
	ExportID                   string   `json:"export_id"`
	SourceAuthorityRevision    uint64   `json:"source_authority_revision"`
	RestoredAuthorityRevision  uint64   `json:"restored_authority_revision"`
	Limitations                []string `json:"limitations"`
}

func TestM0SovereignCoreGoldenProof(t *testing.T) {
	root := repoRoot(t)
	fixture := loadFixture(t, filepath.Join(root, "testdata", "golden", "m0-journey-v1.json"))
	bin, binaryHash := buildBinary(t, root)
	ownerPass := "golden-owner-fixture-passphrase"
	exportPass := "golden-export-fixture-passphrase"
	sourceData := t.TempDir()

	initialized := runJSON(t, bin, sourceData, ownerPass+"\n", "init")
	auroraID := stringField(t, initialized, "aurora_id")
	if auroraID == "" { t.Fatal("init returned empty Aurora identity") }

	created := runJSON(t, bin, sourceData, ownerPass+"\n", "project", "create", "--label", fixture.ProjectLabel, "--objective", fixture.ProjectObjective)
	projectID := stringField(t, created, "project_id")
	if projectID == "" { t.Fatal("create returned empty Project identity") }

	grant := runJSON(t, bin, sourceData, ownerPass+"\n", "authority", "grant", "--expected", "1", "--project", projectID, "--action", fixture.RequiredAuthorityAction)
	grantState := mapField(t, grant, "state")
	if got := uint64Field(t, grantState, "revision"); got != 2 { t.Fatalf("authority revision after grant=%d want 2", got) }

	accepted := runJSON(t, bin, sourceData, ownerPass+"\n", "project", "set-state", "--project", projectID, "--expected", "none", "--kind", fixture.StateKind, "--summary", fixture.StateSummary, "--next-action", fixture.NextAction, "--next-summary", fixture.NextSummary, "--next-authority-action", fixture.RequiredAuthorityAction)
	if got := uint64Field(t, accepted, "state_revision"); got != 1 { t.Fatalf("accepted state revision=%d want 1", got) }

	// Every CLI invocation above has exited. These fresh invocations prove process-death recovery.
	statusAfterRestart := runJSON(t, bin, sourceData, ownerPass+"\n", "status")
	if got := stringField(t, statusAfterRestart, "aurora_id"); got != auroraID { t.Fatalf("Aurora ID after restart=%q want %q", got, auroraID) }
	projectAfterRestart := runJSON(t, bin, sourceData, ownerPass+"\n", "project", "show", "--project", projectID)
	assertProjectStateAndAction(t, projectAfterRestart, projectID, fixture, 1, 2)

	beforeStale := runJSON(t, bin, sourceData, ownerPass+"\n", "status")
	if _, _, err := runFailure(bin, sourceData, ownerPass+"\n", "project", "set-state", "--project", projectID, "--expected", "none", "--kind", "stale", "--summary", "must reject"); err == nil { t.Fatal("stale transition unexpectedly succeeded") }
	afterStale := runJSON(t, bin, sourceData, ownerPass+"\n", "status")
	if uint64Field(t, beforeStale, "governing_generation") != uint64Field(t, afterStale, "governing_generation") { t.Fatalf("rejected stale transition changed governing generation: before=%v after=%v", beforeStale, afterStale) }
	assertProjectStateAndAction(t, runJSON(t, bin, sourceData, ownerPass+"\n", "project", "show", "--project", projectID), projectID, fixture, 1, 2)

	exportFile := filepath.Join(t.TempDir(), "m0-golden.aurora.age")
	exported := runJSON(t, bin, sourceData, ownerPass+"\n"+exportPass+"\n", "export", "--out", exportFile)
	exportID := stringField(t, exported, "export_id")
	if exportID == "" { t.Fatal("export returned empty export ID") }

	restoredData := filepath.Join(t.TempDir(), "restored")
	restored := runJSON(t, bin, sourceData, exportPass+"\n"+ownerPass+"\n", "restore", "--in", exportFile, "--target-data-dir", restoredData)
	if got := stringField(t, restored, "aurora_id"); got != auroraID { t.Fatalf("restored Aurora ID=%q want %q", got, auroraID) }
	if got := stringField(t, restored, "trust_status"); got != "REVALIDATION_REQUIRED" { t.Fatalf("restore trust=%q want REVALIDATION_REQUIRED", got) }

	if _, _, err := runFailure(bin, restoredData, "not-the-owner-passphrase\n", "authority", "revalidate"); err == nil { t.Fatal("non-owner credential revalidated restored authority") }
	revalidated := runJSON(t, bin, restoredData, ownerPass+"\n", "authority", "revalidate")
	revalidatedState := mapField(t, revalidated, "state")
	if got := uint64Field(t, revalidatedState, "revision"); got != 3 { t.Fatalf("authority revision after owner revalidation=%d want 3", got) }
	if got, ok := revalidatedState["revalidation_required"].(bool); !ok || got { t.Fatalf("revalidation_required=%v want false", revalidatedState["revalidation_required"]) }

	restoredProject := runJSON(t, bin, restoredData, ownerPass+"\n", "project", "show", "--project", projectID)
	assertProjectStateAndAction(t, restoredProject, projectID, fixture, 1, 3)
	restoredStatus := runJSON(t, bin, restoredData, ownerPass+"\n", "status")
	if got := stringField(t, restoredStatus, "aurora_id"); got != auroraID { t.Fatalf("post-revalidation Aurora ID=%q want %q", got, auroraID) }

	evidence := goldenEvidence{
		ProofID:                   "GP-M0-SOVEREIGN-CORE-001",
		SourceRevision:            gitRevision(t, root),
		BinarySHA256:              binaryHash,
		DependencyLockSHA256:      dependencyLockHash(t, root),
		GoVersion:                 runtime.Version(),
		GOOS:                      runtime.GOOS,
		GOARCH:                    runtime.GOARCH,
		FixtureID:                 fixture.FixtureID,
		AuroraID:                  auroraID,
		ProjectID:                 projectID,
		ExportID:                  exportID,
		SourceAuthorityRevision:   2,
		RestoredAuthorityRevision: 3,
		Limitations: []string{
			"hosted process execution does not prove physical power-loss or storage-controller behavior",
			"R7 self-produced proof is Claim/Evidence and is not an R8 Product Milestone Verdict",
		},
	}
	raw, err := json.Marshal(evidence)
	if err != nil { t.Fatal(err) }
	t.Logf("M0_GOLDEN_EVIDENCE_JSON=%s", raw)
}

func assertProjectStateAndAction(t *testing.T, view map[string]any, projectID string, fixture goldenFixture, stateRevision, authorityRevision uint64) {
	t.Helper()
	p := mapField(t, view, "project")
	if got := stringField(t, p, "project_id"); got != projectID { t.Fatalf("Project ID=%q want %q", got, projectID) }
	current := mapField(t, view, "current_state")
	if got := uint64Field(t, current, "state_revision"); got != stateRevision { t.Fatalf("state revision=%d want %d", got, stateRevision) }
	state := mapField(t, current, "accepted_state")
	if got := stringField(t, state, "state_kind"); got != fixture.StateKind { t.Fatalf("state kind=%q want %q", got, fixture.StateKind) }
	if got := stringField(t, state, "state_summary"); got != fixture.StateSummary { t.Fatalf("state summary=%q want %q", got, fixture.StateSummary) }
	action := mapField(t, current, "proposed_next_action")
	if got := stringField(t, action, "action_class"); got != fixture.NextAction { t.Fatalf("next action=%q want %q", got, fixture.NextAction) }
	projection := mapField(t, view, "next_safe_action")
	if got := stringField(t, projection, "decision"); got != "PERMITTED" { t.Fatalf("next action decision=%q want PERMITTED; projection=%v", got, projection) }
	if got := uint64Field(t, projection, "authority_state_revision"); got != authorityRevision { t.Fatalf("projection authority revision=%d want %d", got, authorityRevision) }
}

func loadFixture(t *testing.T, path string) goldenFixture {
	t.Helper()
	raw, err := os.ReadFile(path)
	if err != nil { t.Fatalf("load Golden Proof fixture: %v", err) }
	var fixture goldenFixture
	if err := json.Unmarshal(raw, &fixture); err != nil { t.Fatalf("decode Golden Proof fixture: %v", err) }
	if fixture.FixtureID == "" || fixture.ProjectLabel == "" || fixture.ProjectObjective == "" || fixture.StateKind == "" || fixture.StateSummary == "" || fixture.NextAction == "" || fixture.RequiredAuthorityAction == "" { t.Fatalf("incomplete Golden Proof fixture: %+v", fixture) }
	return fixture
}

func repoRoot(t *testing.T) string {
	t.Helper()
	root, err := filepath.Abs("../..")
	if err != nil { t.Fatal(err) }
	return root
}

func buildBinary(t *testing.T, root string) (string, string) {
	t.Helper()
	bin := filepath.Join(t.TempDir(), "aurora")
	if runtime.GOOS == "windows" { bin += ".exe" }
	cmd := exec.Command("go", "build", "-o", bin, "./cmd/aurora")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil { t.Fatalf("build Golden Proof binary: %v\n%s", err, out) }
	raw, err := os.ReadFile(bin)
	if err != nil { t.Fatal(err) }
	sum := sha256.Sum256(raw)
	return bin, hex.EncodeToString(sum[:])
}

func gitRevision(t *testing.T, root string) string {
	t.Helper()
	cmd := exec.Command("git", "rev-parse", "HEAD")
	cmd.Dir = root
	raw, err := cmd.Output()
	if err != nil { t.Fatalf("resolve source revision: %v", err) }
	return strings.TrimSpace(string(raw))
}

func dependencyLockHash(t *testing.T, root string) string {
	t.Helper()
	h := sha256.New()
	for _, name := range []string{"go.mod", "go.sum"} {
		raw, err := os.ReadFile(filepath.Join(root, name))
		if err != nil { t.Fatal(err) }
		_, _ = h.Write([]byte(name + "\x00"))
		_, _ = h.Write(raw)
	}
	return hex.EncodeToString(h.Sum(nil))
}

func runJSON(t *testing.T, bin, dataDir, stdin string, args ...string) map[string]any {
	t.Helper()
	stdout, stderr, err := run(bin, dataDir, stdin, args...)
	if err != nil { t.Fatalf("command %v failed: %v\nstdout=%s\nstderr=%s", args, err, stdout, stderr) }
	var out map[string]any
	if err := json.Unmarshal([]byte(stdout), &out); err != nil { t.Fatalf("decode command %v JSON: %v\nstdout=%s\nstderr=%s", args, err, stdout, stderr) }
	return out
}

func runFailure(bin, dataDir, stdin string, args ...string) (string, string, error) { return run(bin, dataDir, stdin, args...) }

func run(bin, dataDir, stdin string, args ...string) (string, string, error) {
	cmdArgs := append([]string{"--data-dir", dataDir, "--json"}, args...)
	cmd := exec.Command(bin, cmdArgs...)
	cmd.Stdin = bytes.NewBufferString(stdin)
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func mapField(t *testing.T, obj map[string]any, name string) map[string]any {
	t.Helper()
	value, ok := obj[name]
	if !ok { t.Fatalf("missing field %q in %v", name, obj) }
	out, ok := value.(map[string]any)
	if !ok { t.Fatalf("field %q is %T, want object: %v", name, value, obj) }
	return out
}

func stringField(t *testing.T, obj map[string]any, name string) string {
	t.Helper()
	value, ok := obj[name]
	if !ok { t.Fatalf("missing field %q in %v", name, obj) }
	out, ok := value.(string)
	if !ok { t.Fatalf("field %q is %T, want string: %v", name, value, obj) }
	return out
}

func uint64Field(t *testing.T, obj map[string]any, name string) uint64 {
	t.Helper()
	value, ok := obj[name]
	if !ok { t.Fatalf("missing field %q in %v", name, obj) }
	n, ok := value.(float64)
	if !ok || n < 0 || n != float64(uint64(n)) { t.Fatalf("field %q=%v (%T), want uint64-compatible number", name, value, value) }
	return uint64(n)
}

func _unusedFormatGuard() { _ = fmt.Sprintf("") }
