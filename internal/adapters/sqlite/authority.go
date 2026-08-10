package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

func (s *Store) CommitAuthorityRevision(ctx context.Context, in ports.AuthorityMutation) (ports.AuthorityMutationResult, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil { return ports.AuthorityMutationResult{}, err }
	defer func(){ _ = tx.Rollback() }()
	var currentAuthority, generation uint64
	if err := tx.QueryRowContext(ctx, `SELECT current_authority_revision, governing_generation FROM core_state WHERE singleton_key='core'`).Scan(&currentAuthority, &generation); err != nil {
		if errors.Is(err, sql.ErrNoRows) { return ports.AuthorityMutationResult{}, ports.ErrNotInitialized }
		return ports.AuthorityMutationResult{}, err
	}
	if currentAuthority != in.ExpectedAuthorityRevision { return ports.AuthorityMutationResult{}, ports.ErrAuthorityRevisionConflict }
	if generation != in.ExpectedGeneration { return ports.AuthorityMutationResult{}, ports.ErrGenerationConflict }
	when := in.ChangedAt.UTC().Format(time.RFC3339Nano)
	if _, err := tx.ExecContext(ctx, `INSERT INTO authority_revisions(authority_revision,predecessor_revision,authority_state_json,changed_by,changed_at) VALUES(?,?,?,?,?)`, in.NewAuthorityRevision, in.ExpectedAuthorityRevision, string(in.AuthorityJSON), in.ChangedBy, when); err != nil { return ports.AuthorityMutationResult{}, err }
	if _, err := tx.ExecContext(ctx, `UPDATE core_state SET current_authority_revision=?,governing_generation=?,governing_descriptor_hmac=?,updated_at=? WHERE singleton_key='core' AND current_authority_revision=? AND governing_generation=?`, in.NewAuthorityRevision, in.NewGeneration, in.GoverningHMAC, when, in.ExpectedAuthorityRevision, in.ExpectedGeneration); err != nil { return ports.AuthorityMutationResult{}, err }
	if _, err := tx.ExecContext(ctx, `INSERT INTO records(record_id,kind,operation_id,project_id,state_revision,authority_revision,outcome,reason,details_json,created_at) VALUES(?, 'AUDIT', ?, NULL, NULL, ?, 'ACCEPTED','AUTHORITY_CHANGED','{}',?)`, in.AuditRecordID, in.OperationID, in.NewAuthorityRevision, when); err != nil { return ports.AuthorityMutationResult{}, err }
	if err := tx.Commit(); err != nil { return ports.AuthorityMutationResult{}, err }
	return ports.AuthorityMutationResult{AuthorityRevision: in.NewAuthorityRevision, GoverningGeneration: in.NewGeneration}, nil
}
