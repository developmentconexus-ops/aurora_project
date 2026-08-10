package sqlite

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
)

var ErrBrokenCurrentStatePointer=errors.New("current Project state pointer does not resolve")

func(s *Store)GetCurrentProjectState(ctx context.Context,id project.ProjectID)(*project.ProjectStateRevision,error){var current sql.NullInt64;if err:=s.db.QueryRowContext(ctx,`SELECT current_state_revision FROM projects WHERE project_id=?`,id).Scan(&current);err!=nil{return nil,err};if !current.Valid{return nil,nil};return s.loadProjectStateRevision(ctx,id,project.StateRevision(current.Int64))}
func(s *Store)loadProjectStateRevision(ctx context.Context,id project.ProjectID,rev project.StateRevision)(*project.ProjectStateRevision,error){var out project.ProjectStateRevision;var pred sql.NullInt64;var payload,intent,next sql.NullString;var acceptedAt string;err:=s.db.QueryRowContext(ctx,`SELECT project_id,state_revision,predecessor_revision,state_schema_version,state_kind,state_summary,state_payload_json,accepted_intent_ref,proposed_next_action_json,accepted_by_actor,accepted_at,transition_attempt_id FROM project_state_revisions WHERE project_id=? AND state_revision=?`,id,rev).Scan(&out.ProjectID,&out.Revision,&pred,&out.State.SchemaVersion,&out.State.Kind,&out.State.Summary,&payload,&intent,&next,&out.AcceptedByActor,&acceptedAt,&out.TransitionAttemptID);if errors.Is(err,sql.ErrNoRows){return nil,ErrBrokenCurrentStatePointer};if err!=nil{return nil,err};if pred.Valid{p:=project.StateRevision(pred.Int64);out.PredecessorRevision=&p};if payload.Valid{out.State.Payload=json.RawMessage(payload.String)};if intent.Valid{out.AcceptedIntentRef=intent.String};if next.Valid{var action project.ActionDescriptor;if err:=json.Unmarshal([]byte(next.String),&action);err!=nil{return nil,err};out.ProposedNextAction=&action};out.AcceptedAt,err=time.Parse(time.RFC3339Nano,acceptedAt);if err!=nil{return nil,err};if err:=project.ValidateStateEnvelope(out.State);err!=nil{return nil,err};return &out,nil}
