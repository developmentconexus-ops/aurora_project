package sqlite

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

func (s *Store) LoadProjectCurrent(ctx context.Context, id string) (ports.ProjectRecord, *ports.ProjectStateRecord, error) {
	projectRow, err := s.loadProjectBase(ctx, id)
	if err != nil {
		return ports.ProjectRecord{}, nil, err
	}
	if projectRow.CurrentStateRevision == nil {
		return projectRow, nil, nil
	}
	state, err := s.loadStateRevision(ctx, id, *projectRow.CurrentStateRevision)
	if err != nil {
		return ports.ProjectRecord{}, nil, err
	}
	projectRow.CurrentState = state
	return projectRow, state, nil
}

func (s *Store) loadStateRevision(ctx context.Context, projectID string, revision uint64) (*ports.ProjectStateRecord, error) {
	var r ports.ProjectStateRecord
	var predecessor sql.NullInt64
	var payload, action, intent sql.NullString
	var acceptedAt string
	err := s.db.QueryRowContext(ctx, `SELECT project_id,state_revision,predecessor_revision,state_schema_version,state_kind,state_summary,state_payload_json,accepted_intent_ref,proposed_next_action_json,accepted_by_actor,accepted_at,transition_attempt_id FROM project_state_revisions WHERE project_id=? AND state_revision=?`, projectID, revision).Scan(
		&r.ProjectID, &r.Revision, &predecessor, &r.State.SchemaVersion, &r.State.Kind, &r.State.Summary, &payload, &intent, &action, &r.AcceptedByActor, &acceptedAt, &r.TransitionAttemptID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ports.ErrCurrentStateMissing
		}
		return nil, err
	}
	if predecessor.Valid {
		v := uint64(predecessor.Int64)
		r.PredecessorRevision = &v
	}
	if payload.Valid {
		if !json.Valid([]byte(payload.String)) {
			return nil, ports.ErrCurrentStateMissing
		}
		r.State.Payload = []byte(payload.String)
	}
	if intent.Valid {
		r.AcceptedIntentRef = intent.String
	}
	if action.Valid {
		if !json.Valid([]byte(action.String)) {
			return nil, ports.ErrCurrentStateMissing
		}
		r.ProposedNextActionJSON = []byte(action.String)
	}
	parsed, err := time.Parse(time.RFC3339Nano, acceptedAt)
	if err != nil {
		return nil, ports.ErrCurrentStateMissing
	}
	r.AcceptedAt = parsed
	return &r, nil
}
