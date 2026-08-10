package application

import (
	"encoding/json"
	"sort"

	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

type governingStateSnapshot struct {
	Revision              uint64          `json:"state_revision"`
	PredecessorRevision   *uint64         `json:"predecessor_revision,omitempty"`
	SchemaVersion         string          `json:"schema_version"`
	Kind                  string          `json:"kind"`
	Summary               string          `json:"summary"`
	Payload               json.RawMessage `json:"payload,omitempty"`
	ProposedNextAction    json.RawMessage `json:"proposed_next_action,omitempty"`
	AcceptedByActor       string          `json:"accepted_by_actor"`
	AcceptedAt            string          `json:"accepted_at"`
}

type governingProjectSnapshot struct {
	ProjectID            string                  `json:"project_id"`
	DisplayLabel         string                  `json:"display_label"`
	ObjectiveSummary     string                  `json:"objective_summary"`
	CurrentStateRevision *uint64                 `json:"current_state_revision"`
	CurrentState         *governingStateSnapshot `json:"current_state,omitempty"`
}

func governingSnapshot(current ports.CurrentSnapshot) bootstrapGoverningSnapshot {
	projects := make([]governingProjectSnapshot, 0, len(current.Projects))
	for _, p := range current.Projects {
		entry := governingProjectSnapshot{ProjectID: p.ProjectID, DisplayLabel: p.DisplayLabel, ObjectiveSummary: p.ObjectiveSummary, CurrentStateRevision: p.CurrentStateRevision}
		if p.CurrentState != nil {
			entry.CurrentState = &governingStateSnapshot{
				Revision: p.CurrentState.Revision,
				PredecessorRevision: p.CurrentState.PredecessorRevision,
				SchemaVersion: p.CurrentState.State.SchemaVersion,
				Kind: p.CurrentState.State.Kind,
				Summary: p.CurrentState.State.Summary,
				Payload: append(json.RawMessage(nil), p.CurrentState.State.Payload...),
				ProposedNextAction: append(json.RawMessage(nil), p.CurrentState.ProposedNextActionJSON...),
				AcceptedByActor: p.CurrentState.AcceptedByActor,
				AcceptedAt: p.CurrentState.AcceptedAt.UTC().Format("2006-01-02T15:04:05.999999999Z07:00"),
			}
		}
		projects = append(projects, entry)
	}
	sort.Slice(projects, func(i, j int) bool { return projects[i].ProjectID < projects[j].ProjectID })
	return bootstrapGoverningSnapshot{
		Version: 1,
		Aurora: bootstrapAurora{AuroraID: current.AuroraID, OwnerOperatorID: current.OwnerOperatorID},
		GoverningGeneration: current.GoverningGeneration,
		Projects: projects,
		Authority: json.RawMessage(current.AuthorityJSON),
	}
}
