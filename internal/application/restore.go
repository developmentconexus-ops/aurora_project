package application

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/authority"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/evidence"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/identity"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/portability"
	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

var ErrRestoreTargetNotFresh = ports.ErrRestoreTargetNotFresh
var ErrRestoreNotRevalidatable = errors.New("restored authority is not in REVALIDATION_REQUIRED state")

type RestoreResult struct {
	ExportID               string                  `json:"export_id"`
	AuroraID               identity.AuroraID       `json:"aurora_id"`
	AuthorityStateRevision uint64                  `json:"authority_state_revision"`
	GoverningGeneration    uint64                  `json:"governing_generation"`
	TrustStatus            TrustStatus             `json:"trust_status"`
	OperationID            string                  `json:"operation_id"`
	Evidence               evidence.EvidenceRecord `json:"evidence"`
}

func (s *Service) Restore(ctx context.Context, ciphertext, exportSecret, ownerPassphrase []byte) (RestoreResult, error) {
	op := s.startOperation(ctx, "portability.restore", "", evidence.Correlation{})
	fail := func(err error) (RestoreResult, error) { op.finish("FAILED", err.Error()); return RestoreResult{}, err }
	if s.ExportProtection == nil {
		return fail(errors.New("export protection adapter is not configured"))
	}
	plain, err := s.ExportProtection.Unprotect(op.ctx, ciphertext, exportSecret)
	if err != nil { return fail(fmt.Errorf("unprotect export: %w", err)) }
	var doc portability.Document
	if err := json.Unmarshal(plain, &doc); err != nil { return fail(fmt.Errorf("decode export JSON: %w", err)) }
	if err := portability.Verify(doc); err != nil { return fail(err) }
	root := portableToRoot(doc.OwnerRecovery)
	ork, err := unlockORK(ownerPassphrase, root)
	if err != nil { return fail(fmt.Errorf("recover owner root: %w", err)) }
	defer zero(ork)
	state := storeStateFromDocument(doc)
	found := false
	for i := range state.Authority.Revisions {
		if state.Authority.Revisions[i].Revision == state.Authority.CurrentRevision {
			state.Authority.Revisions[i].RevalidationRequired = true
			found = true
			break
		}
	}
	if !found { return fail(errors.New("current authority revision absent from export")) }
	authorityRevision := uint64(state.Authority.CurrentRevision)
	op.updateCorrelation(evidence.Correlation{AuroraID: state.Aurora.AuroraID, AuthorityRevision: &authorityRevision})
	snap, err := storeStateToCurrentSnapshot(state)
	if err != nil { return fail(err) }
	mac, err := governingMAC(ork, snap)
	if err != nil { return fail(err) }
	if err := s.State.BuildFreshFromExport(op.ctx, ports.RestoreSnapshot{State: state, GoverningMAC: mac}); err != nil {
		if errors.Is(err, ports.ErrRestoreTargetNotFresh) { return fail(ErrRestoreTargetNotFresh) }
		return fail(err)
	}
	if err := s.Trust.StoreRootEnvelope(op.ctx, root); err != nil { return fail(fmt.Errorf("publish recovered owner root: %w", err)) }
	now := s.operationNow()
	anchor := ports.Anchor{Version: 1, RootID: root.RootID, AuroraID: state.Aurora.AuroraID, GoverningGeneration: state.GoverningGeneration, ObservedWallTimeHighWater: now}
	am, err := anchorMAC(ork, anchor)
	if err != nil { return fail(err) }
	anchor.HMAC = encodeMAC(am)
	if err := s.Trust.PublishAnchor(op.ctx, anchor); err != nil { return fail(fmt.Errorf("publish restore anchor: %w", err)) }
	classification, err := s.ClassifyTrust(op.ctx, ownerPassphrase)
	if err != nil { return fail(err) }
	if classification.Status != TrustRevalidationRequired { return fail(fmt.Errorf("restored trust status=%s want REVALIDATION_REQUIRED", classification.Status)) }
	record := s.operationEvidence(op, "CRIT-008", "RESTORE_INTEGRITY_AND_AUTHORITY_VALIDATION", []string{"export:" + doc.ExportID, "aurora:" + string(state.Aurora.AuroraID)}, "sha256:"+doc.Integrity.PayloadSHA256)
	if err := s.appendOperationalRecord(op.ctx, "RESTORE", "ACCEPTED", "REVALIDATION_REQUIRED", record); err != nil { return fail(fmt.Errorf("record restore evidence: %w", err)) }
	op.finish("ACCEPTED", "REVALIDATION_REQUIRED")
	return RestoreResult{ExportID: doc.ExportID, AuroraID: state.Aurora.AuroraID, AuthorityStateRevision: authorityRevision, GoverningGeneration: state.GoverningGeneration, TrustStatus: classification.Status, OperationID: op.correlation.OperationID, Evidence: record}, nil
}

func (s *Service) RevalidateRestoredAuthority(ctx context.Context, ownerPassphrase []byte) error {
	assessment, err := s.assessTrust(ctx, ownerPassphrase)
	if err != nil { return err }
	defer func() { if assessment.ORK != nil { zero(assessment.ORK) } }()
	if assessment.Result.Status != TrustRevalidationRequired { return ErrRestoreNotRevalidatable }
	current, err := decodeAuthorityState(assessment.Snapshot.AuthorityStateJSON)
	if err != nil { return err }
	if !current.RevalidationRequired { return ErrRestoreNotRevalidatable }
	now := s.Clock.Now().UTC()
	pred := current.Revision
	next := authority.State{Revision: current.Revision + 1, PredecessorRevision: &pred, Grants: append([]authority.Grant(nil), current.Grants...), RevalidationRequired: false, ChangedBy: string(assessment.Snapshot.Identity.OwnerOperatorID), ChangedAt: now}
	_, err = s.commitAuthorityState(ctx, assessment.trustedCurrent, next, current.Revision, "OP-AUTH-RESTORE-REVALIDATE")
	return err
}
