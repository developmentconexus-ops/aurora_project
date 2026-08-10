package application

import (
	"context"
	"encoding/base64"
	"errors"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

type CreateProjectInput struct {
	DisplayLabel     string
	ObjectiveSummary string
}

type ProjectView = project.Project

func (s *Service) CreateProject(ctx context.Context, ownerPassphrase []byte, in CreateProjectInput) (project.Project, error) {
	if in.DisplayLabel == "" || in.ObjectiveSummary == "" {
		return project.Project{}, errors.New("project label and objective are required")
	}
	if _, err := s.Inspect(ctx, ownerPassphrase); err != nil {
		return project.Project{}, err
	}
	current, err := s.State.LoadCurrent(ctx)
	if err != nil { return project.Project{}, err }
	root, err := s.Trust.LoadRootEnvelope(ctx)
	if err != nil { return project.Project{}, err }
	ork, err := unlockORK(root, ownerPassphrase)
	if err != nil { return project.Project{}, err }
	defer clear(ork)
	id, err := project.NewProjectID()
	if err != nil { return project.Project{}, err }
	now := s.Clock.Now().UTC()
	record := ports.ProjectRecord{ProjectID:string(id), DisplayLabel:in.DisplayLabel, ObjectiveSummary:in.ObjectiveSummary, CreatedAt:now, UpdatedAt:now}
	next := current
	next.GoverningGeneration = current.GoverningGeneration + 1
	next.Projects = append(append([]ports.ProjectRecord(nil), current.Projects...), record)
	snap := governingSnapshot(next)
	mac, err := governingMAC(ork, snap)
	if err != nil { return project.Project{}, err }
	opID, err := randomIdentifier("OP-")
	if err != nil { return project.Project{}, err }
	stored, err := s.State.CreateProject(ctx, ports.CreateProjectMutation{OperationID:opID, Project:record, ExpectedGeneration:current.GoverningGeneration, NewGeneration:next.GoverningGeneration, GoverningHMAC:mac})
	if err != nil { return project.Project{}, err }
	anchor, err := s.Trust.LoadAnchor(ctx)
	if err != nil { return project.Project{}, err }
	anchor.GoverningGeneration = next.GoverningGeneration
	if now.After(anchor.ObservedWallTimeHighWater) { anchor.ObservedWallTimeHighWater = now }
	anchorMACBytes, err := anchorMAC(ork, anchor)
	if err != nil { return project.Project{}, err }
	anchor.HMAC = base64.RawURLEncoding.EncodeToString(anchorMACBytes)
	if err := s.Trust.PublishAnchor(ctx, anchor); err != nil { return project.Project{}, err }
	return projectFromRecord(stored), nil
}

func (s *Service) ShowProject(ctx context.Context, ownerPassphrase []byte, id project.ProjectID) (ProjectView, error) {
	if _, err := s.Inspect(ctx, ownerPassphrase); err != nil { return ProjectView{}, err }
	record, err := s.State.LoadProject(ctx, string(id))
	if err != nil { return ProjectView{}, err }
	return projectFromRecord(record), nil
}

func projectFromRecord(r ports.ProjectRecord) project.Project {
	var rev *project.StateRevision
	if r.CurrentStateRevision != nil { v := project.StateRevision(*r.CurrentStateRevision); rev = &v }
	return project.Project{ProjectID:project.ProjectID(r.ProjectID), DisplayLabel:r.DisplayLabel, ObjectiveSummary:r.ObjectiveSummary, CurrentStateRevision:rev, CreatedAt:r.CreatedAt, UpdatedAt:r.UpdatedAt}
}

var _ = time.Time{}
