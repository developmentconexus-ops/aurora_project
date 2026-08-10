package project

import (
	"crypto/rand"
	"encoding/hex"
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

func NewProjectID() (ProjectID, error) {
	var raw [16]byte
	if _, err := rand.Read(raw[:]); err != nil {
		return "", err
	}
	return ProjectID("PRJ-" + hex.EncodeToString(raw[:])), nil
}
