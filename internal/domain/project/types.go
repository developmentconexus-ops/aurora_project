package project

import (
	"encoding/json"
	"time"
)

type ProjectID string
type StateRevision uint64

type Project struct {
	ProjectID            ProjectID      `json:"project_id"`
	DisplayLabel         string         `json:"display_label"`
	ObjectiveSummary     string         `json:"objective_summary"`
	CurrentStateRevision *StateRevision `json:"current_state_revision,omitempty"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
}

type StateEnvelope struct {
	SchemaVersion string          `json:"state_schema_version"`
	Kind          string          `json:"state_kind"`
	Summary       string          `json:"state_summary"`
	Payload       json.RawMessage `json:"state_payload,omitempty"`
}

type ActionDescriptor struct {
	ActionClass               string    `json:"action_class"`
	Summary                   string    `json:"summary"`
	ProjectID                 ProjectID `json:"project_id"`
	RequiredAuthorityAction   string    `json:"required_authority_action"`
	PreconditionRef           string    `json:"precondition_ref,omitempty"`
}

type ProjectStateRevision struct {
	ProjectID           ProjectID         `json:"project_id"`
	Revision            StateRevision     `json:"state_revision"`
	PredecessorRevision *StateRevision    `json:"predecessor_revision,omitempty"`
	State               StateEnvelope     `json:"accepted_state"`
	AcceptedIntentRef   string            `json:"accepted_intent_ref,omitempty"`
	ProposedNextAction  *ActionDescriptor `json:"proposed_next_action,omitempty"`
	AcceptedByActor     string            `json:"accepted_by_actor"`
	AcceptedAt          time.Time         `json:"accepted_at"`
	TransitionAttemptID string            `json:"transition_attempt_id"`
}
