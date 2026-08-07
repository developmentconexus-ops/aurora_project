package store

import (
	"errors"
	"path/filepath"
	"reflect"
	"testing"
)

func TestAcceptedTransitionIsAtomicAndStaleTransitionMutatesNothing(t *testing.T) {
	path := filepath.Join(t.TempDir(), "aurora.db")
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

	accepted := TransitionInput{
		AttemptID:          "ATTEMPT-2",
		ProjectID:          initial.ProjectID,
		ExpectedRevision:   1,
		NewRevision:        2,
		AuthorityRevision:  "AUTH-1",
		StateKind:          "ACTIVE",
		StateSummary:       "revision two",
		AuditID:            "AUDIT-2",
		EvidenceID:         "EVIDENCE-2",
		EvidenceRef:        "sha256:fixture-2",
	}
	if err := ApplyTransition(path, accepted, nil); err != nil {
		t.Fatalf("accepted transition: %v", err)
	}

	wantAfterAccepted := initial
	wantAfterAccepted.CurrentRevision = 2
	wantAfterAccepted.StateSummary = "revision two"
	got, err := Inspect(path)
	if err != nil {
		t.Fatalf("inspect accepted state: %v", err)
	}
	if !reflect.DeepEqual(got, wantAfterAccepted) {
		t.Fatalf("accepted state mismatch\nwant: %#v\n got: %#v", wantAfterAccepted, got)
	}
	countsBeforeStale, err := InspectCounts(path)
	if err != nil {
		t.Fatalf("counts before stale: %v", err)
	}
	if countsBeforeStale.StateRevisions != 2 || countsBeforeStale.AcceptedTransitions != 1 || countsBeforeStale.AuditRecords != 1 || countsBeforeStale.EvidenceRefs != 1 {
		t.Fatalf("accepted transition rows are incoherent: %#v", countsBeforeStale)
	}

	stale := accepted
	stale.AttemptID = "ATTEMPT-STALE"
	stale.NewRevision = 3
	stale.StateSummary = "must not become governing"
	stale.AuditID = "AUDIT-STALE"
	stale.EvidenceID = "EVIDENCE-STALE"
	stale.EvidenceRef = "sha256:must-not-exist"
	if err := ApplyTransition(path, stale, nil); !errors.Is(err, ErrStaleRevision) {
		t.Fatalf("stale transition error = %v, want ErrStaleRevision", err)
	}

	gotAfterStale, err := Inspect(path)
	if err != nil {
		t.Fatalf("inspect after stale: %v", err)
	}
	if !reflect.DeepEqual(gotAfterStale, wantAfterAccepted) {
		t.Fatalf("stale transition changed governing state\nwant: %#v\n got: %#v", wantAfterAccepted, gotAfterStale)
	}
	countsAfterStale, err := InspectCounts(path)
	if err != nil {
		t.Fatalf("counts after stale: %v", err)
	}
	if !reflect.DeepEqual(countsAfterStale, countsBeforeStale) {
		t.Fatalf("stale transition mutated durable rows\nbefore: %#v\n after: %#v", countsBeforeStale, countsAfterStale)
	}
}
