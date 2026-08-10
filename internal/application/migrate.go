package application

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/portability"
)

type MigratePackageResult struct {
	SourceVersion int    `json:"source_version"`
	TargetVersion int    `json:"target_version"`
	PayloadSHA256 string `json:"payload_sha256"`
	Ciphertext    []byte `json:"-"`
}

func (s *Service) MigratePackage(ctx context.Context, ciphertext, secret []byte) (MigratePackageResult, error) {
	if s.ExportProtection == nil {
		return MigratePackageResult{}, errors.New("export protection adapter is not configured")
	}
	plain, err := s.ExportProtection.Unprotect(ctx, ciphertext, secret)
	if err != nil {
		return MigratePackageResult{}, err
	}
	migrated, err := portability.Migrate(plain)
	if err != nil {
		return MigratePackageResult{}, err
	}
	encoded, err := json.Marshal(migrated.Document)
	if err != nil {
		return MigratePackageResult{}, err
	}
	protected, err := s.ExportProtection.Protect(ctx, encoded, secret)
	if err != nil {
		return MigratePackageResult{}, err
	}
	return MigratePackageResult{
		SourceVersion: migrated.SourceVersion,
		TargetVersion: migrated.TargetVersion,
		PayloadSHA256: migrated.Document.Integrity.PayloadSHA256,
		Ciphertext:    protected,
	}, nil
}
