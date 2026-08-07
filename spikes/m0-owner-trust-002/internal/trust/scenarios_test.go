package trust

import (
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

var (
	fixturePassphrase    = []byte("spk002-owner-passphrase-fixture")
	rotatedPassphrase    = []byte("spk002-owner-passphrase-rotated")
	fixtureNow           = time.Date(2026, 8, 7, 20, 0, 0, 0, time.UTC)
	fixtureExpiry        = fixtureNow.Add(90 * time.Minute)
)

func TestSPK002Scenarios(t *testing.T) {
	t.Run("S01_bootstrap_unlock_and_no_plaintext_secrets", func(t *testing.T) {
		layout := NewLayout(t.TempDir())
		session, err := Bootstrap(layout, "OWNER-SPIKE-001", fixturePassphrase, fixtureNow, fixtureExpiry)
		if err != nil {
			t.Fatalf("bootstrap: %v", err)
		}
		unlocked, err := Unlock(layout.RootPath, fixturePassphrase)
		if err != nil {
			t.Fatalf("unlock: %v", err)
		}
		if !bytes.Equal(rootFingerprint(session), rootFingerprint(unlocked)) {
			t.Fatal("owner root identity changed across unlock")
		}
		if _, err := Unlock(layout.RootPath, []byte("wrong-passphrase")); !errors.Is(err, ErrUnlockFailed) {
			t.Fatalf("wrong-passphrase error = %v, want ErrUnlockFailed", err)
		}
		for _, path := range []string{layout.DBPath, layout.RootPath, layout.AnchorPath} {
			payload, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("read %s: %v", path, err)
			}
			if bytes.Contains(payload, fixturePassphrase) {
				t.Fatalf("plaintext passphrase found on disk in %s", path)
			}
			if bytes.Contains(payload, session.ork) {
				t.Fatalf("plaintext ORK found on disk in %s", path)
			}
		}
	})

	baseDir := t.TempDir()
	base := NewLayout(filepath.Join(baseDir, "base"))
	baseSession, err := Bootstrap(base, "OWNER-SPIKE-001", fixturePassphrase, fixtureNow, fixtureExpiry)
	if err != nil {
		t.Fatalf("base bootstrap: %v", err)
	}

	t.Run("S02_passphrase_rotation_preserves_owner_root", func(t *testing.T) {
		layout := cloneLayout(t, base, filepath.Join(baseDir, "s02"))
		before := rootFingerprint(baseSession)
		if err := RotatePassphrase(layout.RootPath, fixturePassphrase, rotatedPassphrase); err != nil {
			t.Fatalf("rotate: %v", err)
		}
		if _, err := Unlock(layout.RootPath, fixturePassphrase); !errors.Is(err, ErrUnlockFailed) {
			t.Fatalf("old passphrase still unlocks: %v", err)
		}
		afterSession, err := Unlock(layout.RootPath, rotatedPassphrase)
		if err != nil {
			t.Fatalf("new passphrase unlock: %v", err)
		}
		if !bytes.Equal(before, rootFingerprint(afterSession)) {
			t.Fatal("ORK identity changed during passphrase rotation")
		}
		eval, err := Evaluate(layout, afterSession, fixtureNow.Add(time.Minute))
		if err != nil {
			t.Fatalf("evaluate after rotation: %v", err)
		}
		if eval.Classification != ClassNormal || !eval.Permitting {
			t.Fatalf("post-rotation evaluation = %#v", eval)
		}
	})

	t.Run("S03_raw_DB_mutation_fails_MAC", func(t *testing.T) {
		layout := cloneLayout(t, base, filepath.Join(baseDir, "s03"))
		if err := rawMutateStateSummary(layout.DBPath, "tampered without ORK"); err != nil {
			t.Fatalf("raw mutate: %v", err)
		}
		eval, err := Evaluate(layout, baseSession, fixtureNow.Add(time.Minute))
		if err != nil {
			t.Fatalf("evaluate tampered DB: %v", err)
		}
		if eval.Classification != ClassInvalidDBMAC || eval.Permitting {
			t.Fatalf("tampered DB evaluation = %#v", eval)
		}
	})

	t.Run("S04_DB_only_rollback_is_detected", func(t *testing.T) {
		layout := cloneLayout(t, base, filepath.Join(baseDir, "s04"))
		historical := filepath.Join(baseDir, "s04-gen1.db")
		if err := SnapshotDB(layout.DBPath, historical); err != nil {
			t.Fatalf("snapshot generation 1: %v", err)
		}
		if _, err := Advance(layout, baseSession, Mutation{
			ExpectedGeneration: 1,
			AuthorityRevision:  "AUTH-2",
			AuthorityStatus:    StatusActive,
			ExpiresAt:          fixtureExpiry.Add(time.Hour),
			StateSummary:       "generation two",
		}, fixtureNow.Add(time.Minute), nil); err != nil {
			t.Fatalf("advance: %v", err)
		}
		if err := ReplaceDB(layout.DBPath, historical); err != nil {
			t.Fatalf("rollback DB: %v", err)
		}
		eval, err := Evaluate(layout, baseSession, fixtureNow.Add(2*time.Minute))
		if err != nil {
			t.Fatalf("evaluate rollback: %v", err)
		}
		if eval.Classification != ClassStateRollback || eval.Permitting {
			t.Fatalf("rollback evaluation = %#v", eval)
		}
	})

	t.Run("S05_crash_after_DB_commit_is_anchor_lag_and_owner_reconciles", func(t *testing.T) {
		for _, point := range []string{"before_db_commit", "after_db_commit", "after_anchor_temp_sync", "after_anchor_publish"} {
			t.Run(point, func(t *testing.T) {
				layout := cloneLayout(t, base, filepath.Join(baseDir, "s05-"+point))
				cmd := startFaultHelper(t, layout, point)
				killFaultHelper(t, cmd)
				eval, err := Evaluate(layout, baseSession, fixtureNow.Add(5*time.Minute))
				if err != nil {
					t.Fatalf("evaluate after kill: %v", err)
				}
				switch point {
				case "before_db_commit":
					if eval.Classification != ClassNormal || eval.Generation != 1 {
						t.Fatalf("before commit = %#v", eval)
					}
				case "after_db_commit", "after_anchor_temp_sync":
					if eval.Classification != ClassAnchorLag || eval.Permitting {
						t.Fatalf("anchor lag = %#v", eval)
					}
					if err := ReconcileAnchor(layout, baseSession, fixtureNow.Add(6*time.Minute)); err != nil {
						t.Fatalf("owner reconcile: %v", err)
					}
					recovered, err := Evaluate(layout, baseSession, fixtureNow.Add(7*time.Minute))
					if err != nil || recovered.Classification != ClassNormal || recovered.Generation != 2 {
						t.Fatalf("reconciled = %#v err=%v", recovered, err)
					}
				case "after_anchor_publish":
					if eval.Classification != ClassNormal || eval.Generation != 2 {
						t.Fatalf("after publish = %#v", eval)
					}
				}
			})
		}
	})

	t.Run("S06_anchor_ahead_DB_behind_fails_closed", func(t *testing.T) {
		layout := cloneLayout(t, base, filepath.Join(baseDir, "s06"))
		if err := forceValidAnchorGeneration(layout.AnchorPath, baseSession, 2, fixtureNow.Add(time.Minute)); err != nil {
			t.Fatalf("force anchor: %v", err)
		}
		eval, err := Evaluate(layout, baseSession, fixtureNow.Add(2*time.Minute))
		if err != nil {
			t.Fatalf("evaluate: %v", err)
		}
		if eval.Classification != ClassStateRollback || eval.Permitting {
			t.Fatalf("anchor-ahead classification = %#v", eval)
		}
		if err := ReconcileAnchor(layout, baseSession, fixtureNow.Add(3*time.Minute)); !errors.Is(err, ErrStateRollback) {
			t.Fatalf("anchor rollback reconciliation error = %v", err)
		}
	})

	t.Run("S07_backward_wall_clock_cannot_revive_expired_authority", func(t *testing.T) {
		layout := cloneLayout(t, base, filepath.Join(baseDir, "s07"))
		afterExpiry := fixtureNow.Add(2 * time.Hour)
		expired, err := Evaluate(layout, baseSession, afterExpiry)
		if err != nil {
			t.Fatalf("evaluate after expiry: %v", err)
		}
		if expired.Permitting {
			t.Fatalf("expired authority permitted: %#v", expired)
		}
		backward := fixtureNow.Add(time.Hour)
		rolled, err := Evaluate(layout, baseSession, backward)
		if err != nil {
			t.Fatalf("evaluate backward time: %v", err)
		}
		if rolled.Classification != ClassTimeUntrusted || rolled.Permitting {
			t.Fatalf("backward-time evaluation = %#v", rolled)
		}
		start := time.Now()
		time.Sleep(5 * time.Millisecond)
		if time.Since(start) <= 0 {
			t.Fatal("process-local monotonic elapsed time did not advance")
		}
	})

	t.Run("S08_old_authentic_backup_requires_revalidation_after_live_revocation", func(t *testing.T) {
		live := cloneLayout(t, base, filepath.Join(baseDir, "s08-live"))
		bundle := filepath.Join(baseDir, "s08-bundle")
		if err := CreateRecoveryBundle(live, bundle); err != nil {
			t.Fatalf("create old active bundle: %v", err)
		}
		if _, err := Advance(live, baseSession, Mutation{
			ExpectedGeneration: 1,
			AuthorityRevision:  "AUTH-REVOKED-2",
			AuthorityStatus:    StatusRevoked,
			ExpiresAt:          fixtureExpiry,
			StateSummary:       "revoked live authority",
		}, fixtureNow.Add(time.Minute), nil); err != nil {
			t.Fatalf("live revoke: %v", err)
		}
		restored := NewLayout(filepath.Join(baseDir, "s08-restored"))
		if err := RestoreRecoveryBundle(bundle, restored); err != nil {
			t.Fatalf("restore old bundle: %v", err)
		}
		restoredSession, err := Unlock(restored.RootPath, fixturePassphrase)
		if err != nil {
			t.Fatalf("unlock restored root: %v", err)
		}
		eval, err := Evaluate(restored, restoredSession, fixtureNow.Add(2*time.Minute))
		if err != nil {
			t.Fatalf("evaluate restored: %v", err)
		}
		if eval.Classification != ClassRevalidationRequired || eval.Permitting || eval.NextSafeAction != "" {
			t.Fatalf("historical restore = %#v", eval)
		}
		if _, err := os.Stat(restored.AnchorPath); !os.IsNotExist(err) {
			t.Fatalf("historical anchor was imported as current: %v", err)
		}
	})

	t.Run("S09_restored_grant_cannot_self_revalidate", func(t *testing.T) {
		layout, session := restoredFixture(t, base, baseDir, "s09")
		err := Revalidate(layout, session, ActorRestoredGrant, Revalidation{
			AuthorityRevision: "AUTH-SELF-ATTACK",
			ExpiresAt:         fixtureExpiry.Add(time.Hour),
		}, fixtureNow.Add(5*time.Minute))
		if !errors.Is(err, ErrOwnerAuthenticationRequired) {
			t.Fatalf("self-revalidation error = %v", err)
		}
		eval, err := Evaluate(layout, session, fixtureNow.Add(6*time.Minute))
		if err != nil || eval.Classification != ClassRevalidationRequired || eval.Permitting {
			t.Fatalf("self attack changed restore posture: %#v err=%v", eval, err)
		}
	})

	t.Run("S10_authenticated_owner_revalidates_with_new_revision", func(t *testing.T) {
		layout, session := restoredFixture(t, base, baseDir, "s10")
		before, err := Evaluate(layout, session, fixtureNow.Add(5*time.Minute))
		if err != nil || before.NextSafeAction != "" {
			t.Fatalf("before revalidation: %#v err=%v", before, err)
		}
		if err := Revalidate(layout, session, ActorOwner, Revalidation{
			AuthorityRevision: "AUTH-REVALIDATED-2",
			ExpiresAt:         fixtureExpiry.Add(time.Hour),
		}, fixtureNow.Add(6*time.Minute)); err != nil {
			t.Fatalf("owner revalidation: %v", err)
		}
		after, err := Evaluate(layout, session, fixtureNow.Add(7*time.Minute))
		if err != nil {
			t.Fatalf("after revalidation: %v", err)
		}
		if after.Classification != ClassNormal || !after.Permitting || after.AuthorityRevision != "AUTH-REVALIDATED-2" || after.NextSafeAction == "" || after.Generation != 2 {
			t.Fatalf("after owner revalidation = %#v", after)
		}
	})

	t.Run("S11_fresh_machine_root_recovery_stays_revalidation_required", func(t *testing.T) {
		bundle := filepath.Join(baseDir, "s11-bundle")
		if err := CreateRecoveryBundle(base, bundle); err != nil {
			t.Fatalf("bundle: %v", err)
		}
		fresh := NewLayout(filepath.Join(baseDir, "s11-fresh"))
		if err := RestoreRecoveryBundle(bundle, fresh); err != nil {
			t.Fatalf("fresh restore: %v", err)
		}
		session, err := Unlock(fresh.RootPath, fixturePassphrase)
		if err != nil {
			t.Fatalf("owner root recovery: %v", err)
		}
		eval, err := Evaluate(fresh, session, fixtureNow.Add(time.Minute))
		if err != nil || eval.Classification != ClassRevalidationRequired || eval.Permitting {
			t.Fatalf("fresh-machine recovery = %#v err=%v", eval, err)
		}
	})

	t.Run("S12_missing_root_material_fails_without_fabrication", func(t *testing.T) {
		bundle := filepath.Join(baseDir, "s12-bundle")
		if err := CreateRecoveryBundle(base, bundle); err != nil {
			t.Fatalf("bundle: %v", err)
		}
		if err := os.Remove(filepath.Join(bundle, "owner-root.json")); err != nil {
			t.Fatalf("remove root fixture: %v", err)
		}
		target := NewLayout(filepath.Join(baseDir, "s12-target"))
		if err := RestoreRecoveryBundle(bundle, target); !errors.Is(err, ErrMissingRoot) {
			t.Fatalf("restore missing-root error = %v", err)
		}
		if _, err := os.Stat(target.RootPath); !os.IsNotExist(err) {
			t.Fatalf("replacement root was fabricated: %v", err)
		}
	})
}

func TestRecoveryClassificationMatrix(t *testing.T) {
	baseDir := t.TempDir()
	base := NewLayout(filepath.Join(baseDir, "base"))
	session, err := Bootstrap(base, "OWNER-SPIKE-001", fixturePassphrase, fixtureNow, fixtureExpiry)
	if err != nil {
		t.Fatalf("bootstrap: %v", err)
	}

	t.Run("DB_generation_equals_anchor", func(t *testing.T) {
		layout := cloneLayout(t, base, filepath.Join(baseDir, "equal"))
		e, err := Evaluate(layout, session, fixtureNow.Add(time.Minute))
		if err != nil || e.Classification != ClassNormal {
			t.Fatalf("equal = %#v err=%v", e, err)
		}
	})
	t.Run("DB_generation_greater_than_anchor", func(t *testing.T) {
		layout := cloneLayout(t, base, filepath.Join(baseDir, "dbgt"))
		if err := advanceDBOnly(layout, session, Mutation{ExpectedGeneration: 1, AuthorityRevision: "AUTH-2", AuthorityStatus: StatusActive, ExpiresAt: fixtureExpiry, StateSummary: "db ahead"}, fixtureNow.Add(time.Minute)); err != nil {
			t.Fatalf("db-only advance: %v", err)
		}
		e, err := Evaluate(layout, session, fixtureNow.Add(2*time.Minute))
		if err != nil || e.Classification != ClassAnchorLag {
			t.Fatalf("db>anchor = %#v err=%v", e, err)
		}
	})
	t.Run("DB_generation_less_than_anchor", func(t *testing.T) {
		layout := cloneLayout(t, base, filepath.Join(baseDir, "dblt"))
		if err := forceValidAnchorGeneration(layout.AnchorPath, session, 2, fixtureNow.Add(time.Minute)); err != nil {
			t.Fatalf("force anchor: %v", err)
		}
		e, err := Evaluate(layout, session, fixtureNow.Add(2*time.Minute))
		if err != nil || e.Classification != ClassStateRollback {
			t.Fatalf("db<anchor = %#v err=%v", e, err)
		}
	})
	t.Run("invalid_DB_MAC", func(t *testing.T) {
		layout := cloneLayout(t, base, filepath.Join(baseDir, "bad-db-mac"))
		if err := rawMutateStateSummary(layout.DBPath, "corrupt mac"); err != nil {
			t.Fatal(err)
		}
		e, err := Evaluate(layout, session, fixtureNow.Add(time.Minute))
		if err != nil || e.Classification != ClassInvalidDBMAC {
			t.Fatalf("invalid DB MAC = %#v err=%v", e, err)
		}
	})
	t.Run("invalid_anchor_MAC", func(t *testing.T) {
		layout := cloneLayout(t, base, filepath.Join(baseDir, "bad-anchor-mac"))
		if err := rawTamperAnchor(layout.AnchorPath); err != nil {
			t.Fatal(err)
		}
		e, err := Evaluate(layout, session, fixtureNow.Add(time.Minute))
		if err != nil || e.Classification != ClassInvalidAnchorMAC {
			t.Fatalf("invalid anchor MAC = %#v err=%v", e, err)
		}
	})
	t.Run("missing_anchor", func(t *testing.T) {
		layout := cloneLayout(t, base, filepath.Join(baseDir, "missing-anchor"))
		if err := os.Remove(layout.AnchorPath); err != nil {
			t.Fatal(err)
		}
		e, err := Evaluate(layout, session, fixtureNow.Add(time.Minute))
		if err != nil || e.Classification != ClassMissingAnchor {
			t.Fatalf("missing anchor = %#v err=%v", e, err)
		}
	})
	t.Run("missing_wrapped_root", func(t *testing.T) {
		layout := cloneLayout(t, base, filepath.Join(baseDir, "missing-root"))
		if err := os.Remove(layout.RootPath); err != nil {
			t.Fatal(err)
		}
		eval, err := Recover(layout, fixturePassphrase, fixtureNow.Add(time.Minute))
		if err != nil || eval.Classification != ClassMissingRoot || eval.Permitting {
			t.Fatalf("missing root = %#v err=%v", eval, err)
		}
	})
}

func TestSPK002FaultHelper(t *testing.T) {
	if os.Getenv("SPK002_FAULT_HELPER") != "1" {
		return
	}
	layout := NewLayout(os.Getenv("SPK002_LAYOUT"))
	point := os.Getenv("SPK002_POINT")
	marker := os.Getenv("SPK002_MARKER")
	session, err := Unlock(layout.RootPath, fixturePassphrase)
	if err != nil {
		os.Exit(20)
	}
	_, err = Advance(layout, session, Mutation{
		ExpectedGeneration: 1,
		AuthorityRevision:  "AUTH-CRASH-2",
		AuthorityStatus:    StatusActive,
		ExpiresAt:          fixtureExpiry.Add(time.Hour),
		StateSummary:       "crash generation two",
	}, fixtureNow.Add(time.Minute), func(p string) {
		if p != point {
			return
		}
		if err := os.WriteFile(marker, []byte("reached\n"), 0o600); err != nil {
			os.Exit(21)
		}
		for {
			time.Sleep(time.Second)
		}
	})
	if err != nil {
		os.Exit(22)
	}
	os.Exit(0)
}

func startFaultHelper(t *testing.T, layout Layout, point string) *exec.Cmd {
	t.Helper()
	marker := filepath.Join(t.TempDir(), "marker")
	cmd := exec.Command(os.Args[0], "-test.run=^TestSPK002FaultHelper$")
	cmd.Env = append(os.Environ(),
		"SPK002_FAULT_HELPER=1",
		"SPK002_LAYOUT="+layout.BaseDir,
		"SPK002_POINT="+point,
		"SPK002_MARKER="+marker,
	)
	if err := cmd.Start(); err != nil {
		t.Fatalf("start fault helper: %v", err)
	}
	deadline := time.Now().Add(20 * time.Second)
	for {
		if _, err := os.Stat(marker); err == nil {
			return cmd
		}
		if time.Now().After(deadline) {
			_ = cmd.Process.Kill()
			_, _ = cmd.Process.Wait()
			t.Fatalf("fault helper did not reach %s", point)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func killFaultHelper(t *testing.T, cmd *exec.Cmd) {
	t.Helper()
	if err := cmd.Process.Kill(); err != nil {
		t.Fatalf("kill helper: %v", err)
	}
	_, _ = cmd.Process.Wait()
}

func cloneLayout(t *testing.T, src Layout, dstBase string) Layout {
	t.Helper()
	if err := copyTree(src.BaseDir, dstBase); err != nil {
		t.Fatalf("clone layout: %v", err)
	}
	return NewLayout(dstBase)
}

func copyTree(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		if info.IsDir() {
			return os.MkdirAll(target, info.Mode())
		}
		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()
		out, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, info.Mode())
		if err != nil {
			return err
		}
		_, copyErr := io.Copy(out, in)
		closeErr := out.Close()
		if copyErr != nil {
			return copyErr
		}
		return closeErr
	})
}

func restoredFixture(t *testing.T, base Layout, baseDir, name string) (Layout, *OwnerSession) {
	t.Helper()
	bundle := filepath.Join(baseDir, name+"-bundle")
	if err := CreateRecoveryBundle(base, bundle); err != nil {
		t.Fatalf("create recovery bundle: %v", err)
	}
	layout := NewLayout(filepath.Join(baseDir, name+"-restored"))
	if err := RestoreRecoveryBundle(bundle, layout); err != nil {
		t.Fatalf("restore bundle: %v", err)
	}
	session, err := Unlock(layout.RootPath, fixturePassphrase)
	if err != nil {
		t.Fatalf("unlock restored root: %v", err)
	}
	return layout, session
}
