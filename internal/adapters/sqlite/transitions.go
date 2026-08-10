package sqlite

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

func (s *Store) CommitProjectTransition(ctx context.Context, in ports.ProjectTransitionMutation) (ports.ProjectTransitionResult, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil { return ports.ProjectTransitionResult{}, err }
	defer tx.Rollback()

	var current sql.NullInt64
	var generation uint64
	if err := tx.QueryRowContext(ctx, `SELECT p.current_state_revision,c.governing_generation FROM projects p JOIN core_state c ON c.singleton_key='core' WHERE p.project_id=?`, in.ProjectID).Scan(&current, &generation); err != nil {
		return ports.ProjectTransitionResult{}, err
	}
	if generation != in.ExpectedGeneration { return ports.ProjectTransitionResult{}, ErrConcurrentMutation }
	if !dbRevisionMatches(current, in.ExpectedRevision) {
		if err := recordRejectedTransitionTx(ctx, tx, in, "STALE_REVISION"); err != nil { return ports.ProjectTransitionResult{}, err }
		if err := tx.Commit(); err != nil { return ports.ProjectTransitionResult{}, err }
		return ports.ProjectTransitionResult{}, ports.ErrStaleStateRevision
	}

	statePayload := nullableRaw(in.Revision.State.Payload)
	nextAction, err := nullableActionJSON(in.Revision.ProposedNextAction)
	if err != nil { return ports.ProjectTransitionResult{}, err }
	pred := nullableRevision(in.Revision.PredecessorRevision)
	expected := nullableRevision(in.ExpectedRevision)
	requestedState, err := json.Marshal(in.Revision.State)
	if err != nil { return ports.ProjectTransitionResult{}, err }
	requestedAction, err := nullableActionJSON(in.Revision.ProposedNextAction)
	if err != nil { return ports.ProjectTransitionResult{}, err }
	at := in.RequestedAt.UTC().Format(time.RFC3339Nano)

	if _, err := tx.ExecContext(ctx, `INSERT INTO transition_attempts(attempt_id,project_id,actor_id,requested_at,expected_state_revision,requested_state_json,proposed_next_action_json,authority_evaluation_ref,result,reason,accepted_state_revision) VALUES(?,?,?,?,?,?,?,?,?,?,?)`, in.AttemptID, in.ProjectID, in.ActorID, at, expected, string(requestedState), requestedAction, in.AuthorityEvaluationRef, "ACCEPTED", "ACCEPTED", in.Revision.Revision); err != nil {
		return ports.ProjectTransitionResult{}, err
	}
	if _, err := tx.ExecContext(ctx, `INSERT INTO project_state_revisions(project_id,state_revision,predecessor_revision,state_schema_version,state_kind,state_summary,state_payload_json,accepted_intent_ref,proposed_next_action_json,accepted_by_actor,accepted_at,transition_attempt_id) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)`, in.ProjectID, in.Revision.Revision, pred, in.Revision.State.SchemaVersion, in.Revision.State.Kind, in.Revision.State.Summary, statePayload, nullableString(in.Revision.AcceptedIntentRef), nextAction, in.Revision.AcceptedByActor, in.Revision.AcceptedAt.UTC().Format(time.RFC3339Nano), in.AttemptID); err != nil {
		return ports.ProjectTransitionResult{}, err
	}
	res, err := tx.ExecContext(ctx, `UPDATE projects SET current_state_revision=?,updated_at=? WHERE project_id=? AND ((current_state_revision IS NULL AND ? IS NULL) OR current_state_revision=?)`, in.Revision.Revision, in.Revision.AcceptedAt.UTC().Format(time.RFC3339Nano), in.ProjectID, expected, expected)
	if err != nil { return ports.ProjectTransitionResult{}, err }
	n, err := res.RowsAffected()
	if err != nil { return ports.ProjectTransitionResult{}, err }
	if n != 1 { return ports.ProjectTransitionResult{}, ports.ErrStaleStateRevision }
	res, err = tx.ExecContext(ctx, `UPDATE core_state SET governing_generation=?,governing_descriptor_hmac=?,updated_at=? WHERE singleton_key='core' AND governing_generation=?`, in.NewGeneration, in.GoverningMAC, in.Revision.AcceptedAt.UTC().Format(time.RFC3339Nano), in.ExpectedGeneration)
	if err != nil { return ports.ProjectTransitionResult{}, err }
	n, err = res.RowsAffected()
	if err != nil { return ports.ProjectTransitionResult{}, err }
	if n != 1 { return ports.ProjectTransitionResult{}, ErrConcurrentMutation }
	if _, err := tx.ExecContext(ctx, `INSERT INTO records(record_id,kind,operation_id,project_id,state_revision,outcome,reason,details_json,created_at) VALUES(?,?,?,?,?,?,?,?,?)`, "REC-"+in.AttemptID, "AUDIT", in.AttemptID, in.ProjectID, in.Revision.Revision, "ACCEPTED", "STATE_TRANSITION", `{}`, in.Revision.AcceptedAt.UTC().Format(time.RFC3339Nano)); err != nil {
		return ports.ProjectTransitionResult{}, err
	}
	if err := tx.Commit(); err != nil { return ports.ProjectTransitionResult{}, err }
	return ports.ProjectTransitionResult{Revision: in.Revision}, nil
}

func (s *Store) RecordNonGoverning(ctx context.Context, in ports.RecordMutation) error {
	requestedState, err := json.Marshal(in.RequestedState)
	if err != nil { return err }
	action, err := nullableActionJSON(in.ProposedNextAction)
	if err != nil { return err }
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil { return err }
	defer tx.Rollback()
	_, err = tx.ExecContext(ctx, `INSERT INTO transition_attempts(attempt_id,project_id,actor_id,requested_at,expected_state_revision,requested_state_json,proposed_next_action_json,authority_evaluation_ref,result,reason,accepted_state_revision) VALUES(?,?,?,?,?,?,?,?,?,?,NULL)`, in.AttemptID, in.ProjectID, in.ActorID, in.RequestedAt.UTC().Format(time.RFC3339Nano), nullableRevision(in.ExpectedRevision), string(requestedState), action, in.AuthorityEvaluationRef, in.Result, in.Reason)
	if err != nil { return err }
	_, err = tx.ExecContext(ctx, `INSERT INTO records(record_id,kind,operation_id,project_id,outcome,reason,details_json,created_at) VALUES(?,?,?,?,?,?,?,?)`, "REC-"+in.AttemptID, "AUDIT", in.AttemptID, in.ProjectID, in.Result, in.Reason, `{}`, in.RequestedAt.UTC().Format(time.RFC3339Nano))
	if err != nil { return err }
	return tx.Commit()
}

func recordRejectedTransitionTx(ctx context.Context, tx *sql.Tx, in ports.ProjectTransitionMutation, reason string) error {
	state, err := json.Marshal(in.Revision.State)
	if err != nil { return err }
	action, err := nullableActionJSON(in.Revision.ProposedNextAction)
	if err != nil { return err }
	at := in.RequestedAt.UTC().Format(time.RFC3339Nano)
	if _, err := tx.ExecContext(ctx, `INSERT INTO transition_attempts(attempt_id,project_id,actor_id,requested_at,expected_state_revision,requested_state_json,proposed_next_action_json,authority_evaluation_ref,result,reason,accepted_state_revision) VALUES(?,?,?,?,?,?,?,?,?,?,NULL)`, in.AttemptID, in.ProjectID, in.ActorID, at, nullableRevision(in.ExpectedRevision), string(state), action, in.AuthorityEvaluationRef, "REJECTED", reason); err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, `INSERT INTO records(record_id,kind,operation_id,project_id,outcome,reason,details_json,created_at) VALUES(?,?,?,?,?,?,?,?)`, "REC-"+in.AttemptID, "AUDIT", in.AttemptID, in.ProjectID, "REJECTED", reason, `{}`, at)
	return err
}

func dbRevisionMatches(v sql.NullInt64, expected *project.StateRevision) bool {
	if expected == nil { return !v.Valid }
	return v.Valid && project.StateRevision(v.Int64) == *expected
}
func nullableRevision(v *project.StateRevision) any { if v == nil { return nil }; return int64(*v) }
func nullableRaw(v json.RawMessage) any { if len(v) == 0 { return nil }; return string(v) }
func nullableString(v string) any { if v == "" { return nil }; return v }
func nullableActionJSON(v *project.ActionDescriptor) (any, error) {
	if v == nil { return nil, nil }
	raw, err := json.Marshal(v)
	if err != nil { return nil, err }
	return string(raw), nil
}
