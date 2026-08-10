package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

func (s *Store) Bootstrap(ctx context.Context, in ports.BootstrapMutation) (ports.BootstrapResult, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return ports.BootstrapResult{}, err
	}
	defer func() { _ = tx.Rollback() }()
	var count int
	if err := tx.QueryRowContext(ctx, "SELECT COUNT(*) FROM core_state").Scan(&count); err != nil {
		return ports.BootstrapResult{}, err
	}
	if count != 0 {
		return ports.BootstrapResult{}, ports.ErrAlreadyInitialized
	}
	when := in.CreatedAt.UTC().Format(time.RFC3339Nano)
	if _, err := tx.ExecContext(ctx, `INSERT INTO authority_revisions(authority_revision, predecessor_revision, authority_state_json, changed_by, changed_at) VALUES(?, NULL, ?, ?, ?)`, in.AuthorityRevision, string(in.AuthorityJSON), in.OwnerOperatorID, when); err != nil {
		return ports.BootstrapResult{}, err
	}
	if _, err := tx.ExecContext(ctx, `INSERT INTO core_state(singleton_key, aurora_id, owner_operator_id, logical_schema_version, current_authority_revision, governing_generation, governing_descriptor_hmac, created_at, updated_at) VALUES('core', ?, ?, 1, ?, ?, ?, ?, ?)`, in.AuroraID, in.OwnerOperatorID, in.AuthorityRevision, in.GoverningGeneration, in.GoverningDescriptorHMAC, when, when); err != nil {
		return ports.BootstrapResult{}, err
	}
	if _, err := tx.ExecContext(ctx, `INSERT INTO records(record_id, kind, operation_id, project_id, state_revision, authority_revision, outcome, reason, details_json, created_at) VALUES(?, 'AUDIT', ?, NULL, NULL, ?, 'ACCEPTED', 'INITIALIZED', '{}', ?)`, in.OperationID, in.OperationID, in.AuthorityRevision, when); err != nil {
		return ports.BootstrapResult{}, err
	}
	if err := tx.Commit(); err != nil {
		return ports.BootstrapResult{}, err
	}
	return ports.BootstrapResult{AuroraID: in.AuroraID, OwnerOperatorID: in.OwnerOperatorID, CreatedAt: in.CreatedAt.UTC(), AuthorityRevision: in.AuthorityRevision, AuthorityJSON: append([]byte(nil), in.AuthorityJSON...), GoverningGeneration: in.GoverningGeneration, GoverningHMAC: append([]byte(nil), in.GoverningDescriptorHMAC...)}, nil
}

func (s *Store) LoadCurrent(ctx context.Context) (ports.CurrentSnapshot, error) {
	var result ports.CurrentSnapshot
	var created string
	if err := s.db.QueryRowContext(ctx, `SELECT aurora_id, owner_operator_id, current_authority_revision, governing_generation, governing_descriptor_hmac, created_at FROM core_state WHERE singleton_key='core'`).Scan(&result.AuroraID, &result.OwnerOperatorID, &result.AuthorityRevision, &result.GoverningGeneration, &result.GoverningHMAC, &created); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ports.CurrentSnapshot{}, ports.ErrNotInitialized
		}
		return ports.CurrentSnapshot{}, err
	}
	parsed, err := time.Parse(time.RFC3339Nano, created)
	if err != nil {
		return ports.CurrentSnapshot{}, err
	}
	result.CreatedAt = parsed
	if err := s.db.QueryRowContext(ctx, `SELECT authority_state_json FROM authority_revisions WHERE authority_revision=?`, result.AuthorityRevision).Scan(&result.AuthorityJSON); err != nil {
		return ports.CurrentSnapshot{}, err
	}
	return result, nil
}
