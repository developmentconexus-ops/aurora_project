package trust

import (
	"errors"
	"path/filepath"
	"testing"
	"time"
)

func TestOperationalMutationRequiresNormalTrustState(t *testing.T) {
	baseDir := t.TempDir()
	base := NewLayout(filepath.Join(baseDir, "base"))
	session, err := Bootstrap(base, "OWNER-MUTATION-BOUNDARY", fixturePassphrase, fixtureNow, fixtureExpiry)
	if err != nil {
		t.Fatalf("bootstrap: %v", err)
	}

	mutation := func(expected int64, revision string) Mutation {
		return Mutation{
			ExpectedGeneration: expected,
			AuthorityRevision:  revision,
			AuthorityStatus:    StatusActive,
			ExpiresAt:          fixtureExpiry.Add(time.Hour),
			StateSummary:       "must require normal preflight",
		}
	}

	t.Run("STATE_ROLLBACK_requires_explicit_recovery", func(t *testing.T) {
		layout := cloneLayout(t, base, filepath.Join(baseDir, "rollback"))
		if err := forceValidAnchorGeneration(layout.AnchorPath, session, 2, fixtureNow.Add(time.Minute)); err != nil {
			t.Fatalf("force anchor ahead: %v", err)
		}
		if _, err := Advance(layout, session, mutation(1, "AUTH-ROLLBACK-BYPASS"), fixtureNow.Add(2*time.Minute), nil); !errors.Is(err, ErrStateRollback) {
			t.Fatalf("advance from rollback error = %v, want ErrStateRollback", err)
		}
		eval, err := Evaluate(layout, session, fixtureNow.Add(2*time.Minute))
		if err != nil || eval.Classification != ClassStateRollback || eval.Generation != 1 {
			t.Fatalf("rollback state mutated: %#v err=%v", eval, err)
		}
	})

	t.Run("ANCHOR_LAG_requires_reconciliation_before_next_mutation", func(t *testing.T) {
		layout := cloneLayout(t, base, filepath.Join(baseDir, "anchor-lag"))
		if err := advanceDBOnly(layout, session, mutation(1, "AUTH-LAG-2"), fixtureNow.Add(time.Minute)); err != nil {
			t.Fatalf("create DB-ahead fixture: %v", err)
		}
		if _, err := Advance(layout, session, mutation(2, "AUTH-LAG-BYPASS"), fixtureNow.Add(2*time.Minute), nil); !errors.Is(err, ErrAnchorLag) {
			t.Fatalf("advance from anchor lag error = %v, want ErrAnchorLag", err)
		}
		eval, err := Evaluate(layout, session, fixtureNow.Add(2*time.Minute))
		if err != nil || eval.Classification != ClassAnchorLag || eval.Generation != 2 {
			t.Fatalf("anchor-lag state mutated: %#v err=%v", eval, err)
		}
		if err := ReconcileAnchor(layout, session, fixtureNow.Add(3*time.Minute)); err != nil {
			t.Fatalf("reconcile: %v", err)
		}
		if _, err := Advance(layout, session, mutation(2, "AUTH-AFTER-RECONCILE-3"), fixtureNow.Add(4*time.Minute), nil); err != nil {
			t.Fatalf("advance after explicit reconciliation: %v", err)
		}
	})

	t.Run("TIME_UNTRUSTED_blocks_operational_mutation", func(t *testing.T) {
		layout := cloneLayout(t, base, filepath.Join(baseDir, "time-untrusted"))
		future := fixtureNow.Add(30 * time.Minute)
		if eval, err := Evaluate(layout, session, future); err != nil || eval.Classification != ClassNormal {
			t.Fatalf("establish high-water: %#v err=%v", eval, err)
		}
		backward := fixtureNow.Add(5 * time.Minute)
		if _, err := Advance(layout, session, mutation(1, "AUTH-TIME-BYPASS"), backward, nil); !errors.Is(err, ErrTimeUntrusted) {
			t.Fatalf("advance under backward time error = %v, want ErrTimeUntrusted", err)
		}
		eval, err := Evaluate(layout, session, backward)
		if err != nil || eval.Classification != ClassTimeUntrusted || eval.Generation != 1 {
			t.Fatalf("time-untrusted state mutated: %#v err=%v", eval, err)
		}
	})

	t.Run("REVALIDATION_REQUIRED_allows_only_explicit_revalidation", func(t *testing.T) {
		bundle := filepath.Join(baseDir, "restore-bundle")
		if err := CreateRecoveryBundle(base, bundle); err != nil {
			t.Fatalf("bundle: %v", err)
		}
		layout := NewLayout(filepath.Join(baseDir, "restored"))
		if err := RestoreRecoveryBundle(bundle, layout); err != nil {
			t.Fatalf("restore: %v", err)
		}
		restoredSession, err := Unlock(layout.RootPath, fixturePassphrase)
		if err != nil {
			t.Fatalf("unlock restored root: %v", err)
		}
		if _, err := Advance(layout, restoredSession, mutation(1, "AUTH-RESTORE-BYPASS"), fixtureNow.Add(time.Minute), nil); !errors.Is(err, ErrRevalidationRequired) {
			t.Fatalf("advance from restore-pending error = %v, want ErrRevalidationRequired", err)
		}
		eval, err := Evaluate(layout, restoredSession, fixtureNow.Add(time.Minute))
		if err != nil || eval.Classification != ClassRevalidationRequired || eval.Generation != 1 {
			t.Fatalf("restore-pending state mutated: %#v err=%v", eval, err)
		}
	})
}
