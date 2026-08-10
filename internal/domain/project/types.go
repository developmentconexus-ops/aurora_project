package project

import "time"

type ProjectID string
type StateRevision uint64

type Project struct {
	ProjectID            ProjectID       `json:"project_id"`
	DisplayLabel         string          `json:"display_label"`
	ObjectiveSummary     string          `json:"objective_summary"`
	CurrentStateRevision *StateRevision  `json:"current_state_revision,omitempty"`
	CreatedAt            time.Time       `json:"created_at"`
	UpdatedAt            time.Time       `json:"updated_at"`
}
