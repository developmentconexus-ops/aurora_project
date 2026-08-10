package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

func (s *Store) CommitProjectTransition(ctx context.Context, in ports.ProjectTransitionMutation) (ports.ProjectTransitionResult, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return ports.ProjectTransitionResult{}, err
	}
	defer func() { _ = tx.Rollback() }()

	var generation uint64
	if err := tx.QueryRowContext(ctx, `SELECT governing_generation FROM core_state WHERE singleton_key='core'`).Scan(&generation); err != nil {
		return ports.ProjectTransitionResult{}, err
	}
	if generation != in.ExpectedGeneration {
		return ports.ProjectTransitionResult{}, ports.ErrGenerationConflict
	}
	var current sql.NullInt64
	if err := tx.QueryRowContext(ctx, `SELECT current_state_revision FROM projects WHERE project_id=?`, in.ProjectID).Scan(&current); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ports.ProjectTransitionResult{}, ports.ErrProjectNotFound
		}
		return ports.ProjectTransitionResult{}, err
	}
	if !revisionMatches(current, in.ExpectedRevision) {
		return ports.ProjectTransitionResult{}, ports.ErrStateRevisionConflict
	}
	wantRevision := uint64(1)
	if current.Valid {
		wantRevision = uint64(current.Int64) + 1
	}
	if in.State.Revision != wantRevision {
		return ports.ProjectTransitionResult{}, ports.ErrStateRevisionConflict
	}
	when := in.RequestedAt.UTC().Format(time.RFC3339Nano)
	var predecessor any
	if in.State.PredecessorRevision != nil {
		predecessor = *in.State.PredecessorRevision
	}
	var payload any
	if len(in.State.State.Payload) != 0 {
		payload = string(in.State.State.Payload)
	}
	var intent any
	if in.State.AcceptedIntentRef != "" {
		intent = in.State.AcceptedIntentRef
	}
	var action any
	if len(in.ProposedNextActionJSON) != 0 {
		action = string(in.ProposedNextActionJSON)
	}
	if _, err := tx.ExecContext(ctx, `INSERT INTO transition_attempts(attempt_id,project_id,actor_id,requested_at,expected_state_revision,requested_state_json,proposed_next_action_json,authority_evaluation_ref,result,reason,accepted_state_revision) VALUES(?,?,?,?,?,?,?,?, 'ACCEPTED','ACCEPTED',?)`, in.AttemptID, in.ProjectID, in.ActorID, when, nullableRevision(in.ExpectedRevision), in.RequestedStateJSON, action, in.AuthorityEvaluationRef, in.State.Revision); err != nil {
		return ports.ProjectTransitionResult{}, err
	}
	if _, err := tx.ExecContext(ctx, `INSERT INTO project_state_revisions(project_id,state_revision,predecessor_revision,state_schema_version,state_kind,state_summary,state_payload_json,accepted_intent_ref,proposed_next_action_json,accepted_by_actor,accepted_at,transition_attempt_id) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)`, in.ProjectID, in.State.Revision, predecessor, in.State.State.SchemaVersion, in.State.State.Kind, in.State.State.Summary, payload, intent, action, in.State.AcceptedByActor, when, in.AttemptID); err != nil {
		return ports.ProjectTransitionResult{}, err
	}
	if _, err := tx.ExecContext(ctx, `UPDATE projects SET current_state_revision=?, updated_at=? WHERE project_id=?`, in.State.Revision, when, in.ProjectID); err != nil {
		return ports.ProjectTransitionResult{}, err
	}
	if _, err := tx.ExecContext(ctx, `UPDATE core_state SET governing_generation=?, governing_descriptor_hmac=?, updated_at=? WHERE singleton_key='core' AND governing_generation=?`, in.NewGeneration, in.GoverningHMAC, when, in.ExpectedGeneration); err != nil {
		return ports.ProjectTransitionResult{}, err
	}
	if _, err := tx.ExecContext(ctx, `INSERT INTO records(record_id,kind,operation_id,project_id,state_revision,authority_revision,outcome,reason,details_json,created_at) VALUES(?, 'AUDIT', ?, ?, ?, NULL, 'ACCEPTED','PROJECT_STATE_TRANSITION','{}',?)`, in.AuditRecordID, in.AttemptID, in.ProjectID, in.State.Revision, when); err != nil {
		return ports.ProjectTransitionResult{}, err
	}
	if _, err := tx.ExecContext(ctx, `INSERT INTO records(record_id,kind,operation_id,project_id,state_revision,authority_revision,outcome,reason,details_json,created_at) VALUES(?, 'EVIDENCE', ?, ?, ?, NULL, 'OBSERVED','PROJECT_STATE_TRANSITION','{"criterion":"CRIT-003"}',?)`, in.EvidenceRecordID, in.AttemptID, in.ProjectID, in.State.Revision, when); err != nil {
		return ports.ProjectTransitionResult{}, err
	}
	if err := tx.Commit(); err != nil {
		return ports.ProjectTransitionResult{}, err
	}
	return ports.ProjectTransitionResult{ProjectID: in.ProjectID, StateRevision: in.State.Revision, GoverningGeneration: in.NewGeneration}, nil
}

func (s *Store) RecordTransitionRejection(ctx context.Context, in ports.TransitionRejection) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()
	when := in.RequestedAt.UTC().Format(time.RFC3339Nano)
	var action any
	if len(in.ProposedNextActionJSON) != 0 {
		action = string(in.ProposedNextActionJSON)
	}
	if _, err := tx.ExecContext(ctx, `INSERT INTO transition_attempts(attempt_id,project_id,actor_id,requested_at,expected_state_revision,requested_state_json,proposed_next_action_json,authority_evaluation_ref,result,reason,accepted_state_revision) VALUES(?,?,?,?,?,?,?,?, 'REJECTED',?,NULL)`, in.AttemptID, in.ProjectID, in.ActorID, when, nullableRevision(in.ExpectedRevision), in.RequestedStateJSON, action, in.AuthorityEvaluationRef, in.Reason); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, `INSERT INTO records(record_id,kind,operation_id,project_id,state_revision,authority_revision,outcome,reason,details_json,created_at) VALUES(?, 'AUDIT', ?, ?, NULL, NULL, 'REJECTED',?,'{}',?)`, in.AuditRecordID, in.AttemptID, in.ProjectID, in.Reason, when); err != nil {
		return err
	}
	return tx.Commit()
}

func revisionMatches(current sql.NullInt64, expected *uint64) bool {
	if expected == nil {
		return !current.Valid
	}
	return current.Valid && uint64(current.Int64) == *expected
}

func nullableRevision(rev *uint64) any {
	if rev == nil {
		return nil
	}
	return *rev
}
