package application

import (
	"encoding/json"
	"sort"

	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

type governingProjectSnapshot struct {
	ProjectID            string  `json:"project_id"`
	DisplayLabel         string  `json:"display_label"`
	ObjectiveSummary     string  `json:"objective_summary"`
	CurrentStateRevision *uint64 `json:"current_state_revision"`
}

func governingSnapshot(current ports.CurrentSnapshot) bootstrapGoverningSnapshot {
	projects := make([]governingProjectSnapshot, 0, len(current.Projects))
	for _, p := range current.Projects {
		projects = append(projects, governingProjectSnapshot{
			ProjectID: p.ProjectID, DisplayLabel: p.DisplayLabel, ObjectiveSummary: p.ObjectiveSummary,
			CurrentStateRevision: p.CurrentStateRevision,
		})
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
