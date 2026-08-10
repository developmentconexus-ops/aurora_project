package application

import (
	"bytes"
	"testing"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

func TestClassifyTrustTable(t *testing.T) {
	now := time.Date(2026, 8, 9, 23, 30, 0, 0, time.UTC)
	cases := []struct {
		name string
		in   trustSignals
		want TrustState
	}{
		{"normal", trustSignals{dbMACValid: true, anchorPresent: true, anchorMACValid: true, dbGeneration: 4, anchorGeneration: 4, now: now, highWater: now}, TrustNormal},
		{"anchor lag", trustSignals{dbMACValid: true, anchorPresent: true, anchorMACValid: true, dbGeneration: 5, anchorGeneration: 4, now: now, highWater: now}, TrustAnchorLag},
		{"rollback", trustSignals{dbMACValid: true, anchorPresent: true, anchorMACValid: true, dbGeneration: 3, anchorGeneration: 4, now: now, highWater: now}, TrustStateRollback},
		{"db mac", trustSignals{dbMACValid: false, anchorPresent: true, anchorMACValid: true, dbGeneration: 4, anchorGeneration: 4, now: now, highWater: now}, TrustInvalidDBMAC},
		{"anchor mac", trustSignals{dbMACValid: true, anchorPresent: true, anchorMACValid: false, dbGeneration: 4, anchorGeneration: 4, now: now, highWater: now}, TrustInvalidAnchorMAC},
		{"missing anchor", trustSignals{dbMACValid: true, anchorPresent: false, dbGeneration: 4, now: now}, TrustMissingAnchor},
		{"time rollback", trustSignals{dbMACValid: true, anchorPresent: true, anchorMACValid: true, dbGeneration: 4, anchorGeneration: 4, now: now.Add(-time.Hour), highWater: now}, TrustTimeUntrusted},
		{"revalidation", trustSignals{dbMACValid: true, anchorPresent: true, anchorMACValid: true, dbGeneration: 4, anchorGeneration: 4, now: now, highWater: now, revalidationRequired: true}, TrustRevalidationRequired},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := classifyTrust(tc.in); got != tc.want {
				t.Fatalf("got %s want %s", got, tc.want)
			}
		})
	}
}

func TestGoverningMACBindsCurrentStateContentNotOnlyRevision(t *testing.T) {
	key := bytes.Repeat([]byte{7}, 32)
	rev := uint64(1)
	base := ports.CurrentSnapshot{AuroraID: "AUR-F", OwnerOperatorID: "OWNER-LOCAL", AuthorityRevision: 1, AuthorityJSON: []byte(`{"revision":1,"grants":[]}`), GoverningGeneration: 3, Projects: []ports.ProjectRecord{{ProjectID: "PRJ-F", DisplayLabel: "A", ObjectiveSummary: "A", CurrentStateRevision: &rev, CurrentState: &ports.ProjectStateRecord{ProjectID: "PRJ-F", Revision: 1, State: ports.StateEnvelopeRecord{SchemaVersion: "1", Kind: "NOTE", Summary: "one", Payload: []byte(`{"x":1}`)}, AcceptedByActor: "OWNER-LOCAL", AcceptedAt: time.Date(2026, 8, 9, 23, 0, 0, 0, time.UTC)}}}}
	mac1, err := governingMAC(key, governingSnapshot(base))
	if err != nil {
		t.Fatal(err)
	}
	base.Projects[0].CurrentState.State.Payload = []byte(`{"x":2}`)
	mac2, err := governingMAC(key, governingSnapshot(base))
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Equal(mac1, mac2) {
		t.Fatal("governing MAC did not change when current state content changed at same revision")
	}
}
