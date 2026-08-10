package application

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/evidence"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/portability"
)

type ExportResult struct {
	ExportID      string                  `json:"export_id"`
	CreatedAt     time.Time               `json:"created_at"`
	PayloadSHA256 string                  `json:"payload_sha256"`
	Ciphertext    []byte                  `json:"-"`
	OperationID   string                  `json:"operation_id"`
	Evidence      evidence.EvidenceRecord `json:"evidence"`
}

func (s *Service) Export(ctx context.Context, ownerPassphrase, exportSecret []byte) (ExportResult, error) {
	op := s.startOperation(ctx, "portability.export", "", evidence.Correlation{})
	if s.ExportProtection == nil {
		err := errors.New("export protection adapter is not configured")
		op.finish("FAILED", err.Error())
		return ExportResult{}, err
	}
	trusted, err := s.loadTrustedCurrent(op.ctx, ownerPassphrase)
	if err != nil {
		op.finish("FAILED", err.Error())
		return ExportResult{}, err
	}
	defer zero(trusted.ORK)
	authorityRevision := trusted.Snapshot.CurrentAuthorityRevision
	op.updateCorrelation(evidence.Correlation{AuroraID: trusted.Snapshot.Identity.AuroraID, AuthorityRevision: &authorityRevision})
	state, err := s.State.ExportLogicalState(op.ctx)
	if err != nil {
		op.finish("FAILED", err.Error())
		return ExportResult{}, err
	}
	doc := portability.Document{Format: portability.FormatV1, Version: 1, CreatedAt: s.operationNow(), GoverningGeneration: state.GoverningGeneration, Aurora: state.Aurora, Projects: state.Projects, Authority: state.Authority, Attempts: state.Attempts, Records: state.Records, OwnerRecovery: rootToPortable(trusted.Root)}
	if err := portability.Finalize(&doc); err != nil {
		op.finish("FAILED", err.Error())
		return ExportResult{}, err
	}
	plain, err := json.Marshal(doc)
	if err != nil {
		op.finish("FAILED", err.Error())
		return ExportResult{}, err
	}
	ciphertext, err := s.ExportProtection.Protect(op.ctx, plain, exportSecret)
	if err != nil {
		op.finish("FAILED", err.Error())
		return ExportResult{}, err
	}
	record := s.operationEvidence(op, "CRIT-008", "EXPORT_INTEGRITY_AND_PROTECTION", []string{"export:" + doc.ExportID}, "sha256:"+doc.Integrity.PayloadSHA256)
	if err := s.appendOperationalRecord(op.ctx, "EXPORT", "ACCEPTED", "EXPORT_CREATED", record); err != nil {
		op.finish("FAILED", "record export evidence: "+err.Error())
		return ExportResult{}, fmt.Errorf("record export evidence: %w", err)
	}
	op.finish("ACCEPTED", "EXPORT_CREATED")
	return ExportResult{ExportID: doc.ExportID, CreatedAt: doc.CreatedAt, PayloadSHA256: doc.Integrity.PayloadSHA256, Ciphertext: ciphertext, OperationID: op.correlation.OperationID, Evidence: record}, nil
}
