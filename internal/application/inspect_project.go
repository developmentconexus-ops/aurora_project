package application

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

type ProjectInspection struct {
	project.Project
	CurrentState *project.ProjectStateRevision `json:"current_state,omitempty"`
}

func (s *Service) InspectProject(ctx context.Context, ownerPassphrase []byte, id project.ProjectID) (ProjectInspection, error) {
	if _, err := s.Inspect(ctx, ownerPassphrase); err != nil {
		return ProjectInspection{}, err
	}
	record, state, err := s.State.LoadProjectCurrent(ctx, string(id))
	if err != nil {
		return ProjectInspection{}, err
	}
	result := ProjectInspection{Project: projectFromRecord(record)}
	if state != nil {
		translated, err := projectStateFromRecord(*state)
		if err != nil {
			return ProjectInspection{}, err
		}
		result.CurrentState = &translated
	}
	return result, nil
}

func projectStateFromRecord(r ports.ProjectStateRecord) (project.ProjectStateRevision, error) {
	var predecessor *project.StateRevision
	if r.PredecessorRevision != nil {
		v := project.StateRevision(*r.PredecessorRevision)
		predecessor = &v
	}
	var action *project.ActionDescriptor
	if len(r.ProposedNextActionJSON) != 0 {
		action = &project.ActionDescriptor{}
		if err := json.Unmarshal(r.ProposedNextActionJSON, action); err != nil {
			return project.ProjectStateRevision{}, fmt.Errorf("decode proposed next action: %w", err)
		}
	}
	return project.ProjectStateRevision{
		ProjectID:           project.ProjectID(r.ProjectID),
		Revision:            project.StateRevision(r.Revision),
		PredecessorRevision: predecessor,
		State: project.StateEnvelope{
			SchemaVersion: r.State.SchemaVersion,
			Kind:          r.State.Kind,
			Summary:       r.State.Summary,
			Payload:       append(json.RawMessage(nil), r.State.Payload...),
		},
		AcceptedIntentRef:   r.AcceptedIntentRef,
		ProposedNextAction:  action,
		AcceptedByActor:     r.AcceptedByActor,
		AcceptedAt:          r.AcceptedAt,
		TransitionAttemptID: r.TransitionAttemptID,
	}, nil
}
