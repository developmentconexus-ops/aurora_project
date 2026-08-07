package store

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"testing"
	"time"
)

const helperEnv = "AURORA_SPK001_HELPER"

func TestSPK001HelperProcess(t *testing.T) {
	mode := os.Getenv(helperEnv)
	if mode == "" {
		return
	}
	path := os.Getenv("AURORA_SPK001_DB")
	point := os.Getenv("AURORA_SPK001_FAULT")
	marker := os.Getenv("AURORA_SPK001_MARKER")
	if path == "" || point == "" || marker == "" {
		fmt.Fprintln(os.Stderr, "missing helper environment")
		os.Exit(90)
	}
	hook := blockingFaultHook(point, marker)

	var err error
	switch mode {
	case "transition":
		err = ApplyTransition(path, TransitionInput{
			AttemptID:          "ATTEMPT-CRASH",
			ProjectID:          "PROJECT-SPIKE-001",
			ExpectedRevision:   1,
			NewRevision:        2,
			AuthorityRevision:  "AUTH-1",
			StateKind:          "ACTIVE",
			StateSummary:       "revision two after crash boundary",
			AuditID:            "AUDIT-CRASH",
			EvidenceID:         "EVIDENCE-CRASH",
			EvidenceRef:        "sha256:crash-fixture",
		}, hook)
	case "checkpoint":
		err = Checkpoint(path, hook)
	case "backup":
		destination := os.Getenv("AURORA_SPK001_DEST")
		if destination == "" {
			fmt.Fprintln(os.Stderr, "missing backup destination")
			os.Exit(94)
		}
		err = SupportedBackup(path, destination, hook)
	default:
		fmt.Fprintf(os.Stderr, "unknown helper mode %q\n", mode)
		os.Exit(95)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s operation: %v\n", mode, err)
		os.Exit(92)
	}
	fmt.Fprintln(os.Stderr, "fault point was not reached")
	os.Exit(93)
}

func blockingFaultHook(point, marker string) FaultHook {
	return func(got string) {
		if got != point {
			return
		}
		if err := os.WriteFile(marker, []byte(got+"\n"), 0o600); err != nil {
			fmt.Fprintf(os.Stderr, "write marker: %v\n", err)
			os.Exit(91)
		}
		// Timer-backed sleep prevents Go's deadlock detector from terminating
		// the helper before the parent applies the OS process-kill primitive.
		for {
			time.Sleep(time.Hour)
		}
	}
}

func TestProcessKillAtTransitionBoundaries(t *testing.T) {
	points := []string{
		"before_tx",
		"after_validation",
		"after_revision_insert",
		"after_pointer_update",
		"before_commit",
		"after_commit",
	}
	for _, point := range points {
		t.Run(point, func(t *testing.T) {
			dir := t.TempDir()
			path := filepath.Join(dir, "aurora.db")
			initial := Snapshot{
				SchemaVersion:     1,
				AuroraID:          "AURORA-SPIKE-001",
				ProjectID:         "PROJECT-SPIKE-001",
				CurrentRevision:   1,
				AuthorityRevision: "AUTH-1",
				StateKind:         "ACTIVE",
				StateSummary:      "revision one",
			}
			if err := Bootstrap(path, initial); err != nil {
				t.Fatalf("bootstrap: %v", err)
			}

			cmd, _ := startHelperAtPoint(t, "transition", path, "", point)
			killHelper(t, cmd)

			got, err := Inspect(path)
			if err != nil {
				t.Fatalf("inspect after kill at %s: %v", point, err)
			}
			counts, err := InspectCounts(path)
			if err != nil {
				t.Fatalf("counts after kill at %s: %v", point, err)
			}

			if point == "after_commit" {
				want := initial
				want.CurrentRevision = 2
				want.StateSummary = "revision two after crash boundary"
				if !reflect.DeepEqual(got, want) {
					t.Fatalf("post-commit kill recovered wrong state\nwant: %#v\n got: %#v", want, got)
				}
				wantCounts := Counts{StateRevisions: 2, AcceptedTransitions: 1, AuditRecords: 1, EvidenceRefs: 1}
				if !reflect.DeepEqual(counts, wantCounts) {
					t.Fatalf("post-commit rows incoherent: %#v", counts)
				}
				return
			}

			if !reflect.DeepEqual(got, initial) {
				t.Fatalf("pre-commit kill changed governing state at %s\nwant: %#v\n got: %#v", point, initial, got)
			}
			wantCounts := Counts{StateRevisions: 1}
			if !reflect.DeepEqual(counts, wantCounts) {
				t.Fatalf("pre-commit kill left partial rows at %s: %#v", point, counts)
			}
		})
	}
}

func startHelperAtPoint(t *testing.T, mode, path, destination, point string) (*exec.Cmd, string) {
	t.Helper()
	marker := filepath.Join(filepath.Dir(path), mode+"-"+point+".marker")
	cmd := exec.Command(os.Args[0], "-test.run=^TestSPK001HelperProcess$", "-test.v=false")
	cmd.Env = append(os.Environ(),
		helperEnv+"="+mode,
		"AURORA_SPK001_DB="+path,
		"AURORA_SPK001_DEST="+destination,
		"AURORA_SPK001_FAULT="+point,
		"AURORA_SPK001_MARKER="+marker,
	)
	if err := cmd.Start(); err != nil {
		t.Fatalf("start %s helper: %v", mode, err)
	}
	waitForMarker(t, marker, cmd, 15*time.Second)
	return cmd, marker
}

func killHelper(t *testing.T, cmd *exec.Cmd) {
	t.Helper()
	if err := cmd.Process.Kill(); err != nil {
		t.Fatalf("kill child: %v", err)
	}
	_ = cmd.Wait()
}

func waitForMarker(t *testing.T, marker string, cmd *exec.Cmd, timeout time.Duration) {
	t.Helper()
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if _, err := os.Stat(marker); err == nil {
			return
		}
		time.Sleep(20 * time.Millisecond)
	}
	_ = cmd.Process.Kill()
	_ = cmd.Wait()
	t.Fatalf("timeout waiting for fault marker %s", marker)
}
