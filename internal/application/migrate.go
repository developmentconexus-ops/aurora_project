package application

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/evidence"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/portability"
)

type MigratePackageResult struct {
	SourceVersion int                     `json:"source_version"`
	TargetVersion int                     `json:"target_version"`
	PayloadSHA256 string                  `json:"payload_sha256"`
	Ciphertext    []byte                  `json:"-"`
	OperationID   string                  `json:"operation_id"`
	Evidence      evidence.EvidenceRecord `json:"evidence"`
}

func (s *Service) MigratePackage(ctx context.Context, ciphertext, secret []byte) (MigratePackageResult, error) {
	op := s.startOperation(ctx, "portability.migrate", "", evidence.Correlation{})
	fail := func(err error) (MigratePackageResult, error) { op.finish("FAILED", err.Error()); return MigratePackageResult{}, err }
	if s.ExportProtection == nil { return fail(errors.New("export protection adapter is not configured")) }
	plain, err := s.ExportProtection.Unprotect(op.ctx, ciphertext, secret)
	if err != nil { return fail(err) }
	migrated, err := portability.Migrate(plain)
	if err != nil { return fail(err) }
	op.updateCorrelation(evidence.Correlation{AuroraID: migrated.Document.Aurora.AuroraID})
	encoded, err := json.Marshal(migrated.Document)
	if err != nil { return fail(err) }
	protected, err := s.ExportProtection.Protect(op.ctx, encoded, secret)
	if err != nil { return fail(err) }
	record := s.operationEvidence(op, "CRIT-008", "LOGICAL_MIGRATION_AND_SEMANTIC_VALIDATION", []string{"export:" + migrated.Document.ExportID}, "sha256:"+migrated.Document.Integrity.PayloadSHA256)
	op.finish("ACCEPTED", "MIGRATED")
	return MigratePackageResult{SourceVersion: migrated.SourceVersion, TargetVersion: migrated.TargetVersion, PayloadSHA256: migrated.Document.Integrity.PayloadSHA256, Ciphertext: protected, OperationID: op.correlation.OperationID, Evidence: record}, nil
}
