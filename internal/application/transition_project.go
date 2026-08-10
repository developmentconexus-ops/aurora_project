package application

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

type TransitionProjectInput struct {
	AttemptID          string
	ProjectID          project.ProjectID
	ExpectedRevision   *project.StateRevision
	State              project.StateEnvelope
	ProposedNextAction *project.ActionDescriptor
}

type TransitionProjectResult struct {
	ProjectID           project.ProjectID     `json:"project_id"`
	StateRevision       project.StateRevision `json:"state_revision"`
	GoverningGeneration uint64                `json:"governing_generation"`
}

type authorityCheckState struct {
	Revision uint64                `json:"revision"`
	Grants   []authorityCheckGrant `json:"grants"`
}
type authorityCheckGrant struct {
	ProjectScope           []string   `json:"project_scope"`
	PermittedActionClasses []string   `json:"permitted_action_classes"`
	ValidFrom              *time.Time `json:"valid_from"`
	ValidUntil             *time.Time `json:"valid_until"`
	LifecycleStatus        string     `json:"lifecycle_status"`
}

func (s *Service) TransitionProjectState(ctx context.Context, ownerPassphrase []byte, in TransitionProjectInput) (TransitionProjectResult, error) {
	if in.ProjectID == "" {
		return TransitionProjectResult{}, ports.ErrInvalidTransition
	}
	if in.AttemptID == "" {
		id, err := randomIdentifier("ATT-")
		if err != nil {
			return TransitionProjectResult{}, err
		}
		in.AttemptID = id
	}
	now := s.Clock.Now().UTC()
	requestedJSON := requestedStateJSON(in.State)
	actionJSON, err := json.Marshal(in.ProposedNextAction)
	if err != nil {
		return TransitionProjectResult{}, err
	}
	if in.ProposedNextAction == nil {
		actionJSON = nil
	}
	reject := func(reason string, cause error) (TransitionProjectResult, error) {
		auditID, _ := randomIdentifier("AUD-")
		var expected *uint64
		if in.ExpectedRevision != nil {
			v := uint64(*in.ExpectedRevision)
			expected = &v
		}
		_ = s.State.RecordTransitionRejection(ctx, ports.TransitionRejection{AttemptID: in.AttemptID, AuditRecordID: auditID, ProjectID: string(in.ProjectID), ActorID: ownerOperatorID, RequestedAt: now, ExpectedRevision: expected, RequestedStateJSON: requestedJSON, ProposedNextActionJSON: actionJSON, Reason: reason})
		return TransitionProjectResult{}, cause
	}
	if _, err := s.Inspect(ctx, ownerPassphrase); err != nil {
		return TransitionProjectResult{}, err
	}
	if err := validateStateEnvelope(in.State); err != nil {
		return reject("MALFORMED_STATE", fmt.Errorf("%w: %v", ports.ErrInvalidTransition, err))
	}
	if in.ProposedNextAction != nil && in.ProposedNextAction.ProjectID != in.ProjectID {
		return reject("PROJECT_SCOPE_MISMATCH", ports.ErrInvalidTransition)
	}
	current, err := s.State.LoadCurrent(ctx)
	if err != nil {
		return TransitionProjectResult{}, err
	}
	var currentProject *ports.ProjectRecord
	for i := range current.Projects {
		if current.Projects[i].ProjectID == string(in.ProjectID) {
			currentProject = &current.Projects[i]
			break
		}
	}
	if currentProject == nil {
		return TransitionProjectResult{}, ports.ErrProjectNotFound
	}
	if !expectedMatches(currentProject.CurrentStateRevision, in.ExpectedRevision) {
		return reject("STALE_REVISION", ports.ErrStateRevisionConflict)
	}
	if !authorityAllowsProjectMutation(current.AuthorityJSON, in.ProjectID, now) {
		return reject("UNAUTHORIZED", ports.ErrUnauthorized)
	}
	root, err := s.Trust.LoadRootEnvelope(ctx)
	if err != nil {
		return TransitionProjectResult{}, err
	}
	ork, err := unlockORK(root, ownerPassphrase)
	if err != nil {
		return TransitionProjectResult{}, err
	}
	defer clear(ork)
	nextRevision := uint64(1)
	var predecessor *uint64
	if currentProject.CurrentStateRevision != nil {
		nextRevision = *currentProject.CurrentStateRevision + 1
		p := *currentProject.CurrentStateRevision
		predecessor = &p
	}
	stateRecord := ports.ProjectStateRecord{ProjectID: string(in.ProjectID), Revision: nextRevision, PredecessorRevision: predecessor, State: ports.StateEnvelopeRecord{SchemaVersion: in.State.SchemaVersion, Kind: in.State.Kind, Summary: in.State.Summary, Payload: append([]byte(nil), in.State.Payload...)}, ProposedNextActionJSON: append([]byte(nil), actionJSON...), AcceptedByActor: ownerOperatorID, AcceptedAt: now, TransitionAttemptID: in.AttemptID}
	next := current
	next.GoverningGeneration = current.GoverningGeneration + 1
	next.Projects = append([]ports.ProjectRecord(nil), current.Projects...)
	for i := range next.Projects {
		if next.Projects[i].ProjectID == string(in.ProjectID) {
			v := nextRevision
			next.Projects[i].CurrentStateRevision = &v
			next.Projects[i].CurrentState = &stateRecord
			next.Projects[i].UpdatedAt = now
		}
	}
	mac, err := governingMAC(ork, governingSnapshot(next))
	if err != nil {
		return TransitionProjectResult{}, err
	}
	auditID, err := randomIdentifier("AUD-")
	if err != nil {
		return TransitionProjectResult{}, err
	}
	evidenceID, err := randomIdentifier("EVID-")
	if err != nil {
		return TransitionProjectResult{}, err
	}
	var expected *uint64
	if in.ExpectedRevision != nil {
		v := uint64(*in.ExpectedRevision)
		expected = &v
	}
	stored, err := s.State.CommitProjectTransition(ctx, ports.ProjectTransitionMutation{AttemptID: in.AttemptID, AuditRecordID: auditID, EvidenceRecordID: evidenceID, ProjectID: string(in.ProjectID), ActorID: ownerOperatorID, RequestedAt: now, ExpectedRevision: expected, State: stateRecord, RequestedStateJSON: requestedJSON, ProposedNextActionJSON: actionJSON, AuthorityEvaluationRef: fmt.Sprintf("AUTHORITY-REV-%d", current.AuthorityRevision), ExpectedGeneration: current.GoverningGeneration, NewGeneration: next.GoverningGeneration, GoverningHMAC: mac})
	if err != nil {
		if errors.Is(err, ports.ErrStateRevisionConflict) || errors.Is(err, ports.ErrGenerationConflict) {
			return reject("STALE_REVISION", err)
		}
		return TransitionProjectResult{}, err
	}
	anchor, err := s.Trust.LoadAnchor(ctx)
	if err != nil {
		return TransitionProjectResult{}, err
	}
	anchor.GoverningGeneration = stored.GoverningGeneration
	if now.After(anchor.ObservedWallTimeHighWater) {
		anchor.ObservedWallTimeHighWater = now
	}
	anchorBytes, err := anchorMAC(ork, anchor)
	if err != nil {
		return TransitionProjectResult{}, err
	}
	anchor.HMAC = base64.RawURLEncoding.EncodeToString(anchorBytes)
	if err := s.Trust.PublishAnchor(ctx, anchor); err != nil {
		return TransitionProjectResult{}, err
	}
	return TransitionProjectResult{ProjectID: in.ProjectID, StateRevision: project.StateRevision(stored.StateRevision), GoverningGeneration: stored.GoverningGeneration}, nil
}

func requestedStateJSON(env project.StateEnvelope) string {
	encoded, err := json.Marshal(env)
	if err == nil {
		return string(encoded)
	}
	fallback, _ := json.Marshal(map[string]any{"schema_version": env.SchemaVersion, "kind": env.Kind, "summary": env.Summary, "payload_raw": string(env.Payload)})
	return string(fallback)
}
func expectedMatches(current *uint64, expected *project.StateRevision) bool {
	if expected == nil {
		return current == nil
	}
	return current != nil && *current == uint64(*expected)
}
func authorityAllowsProjectMutation(raw []byte, projectID project.ProjectID, now time.Time) bool {
	var state authorityCheckState
	if json.Unmarshal(raw, &state) != nil {
		return false
	}
	for _, g := range state.Grants {
		if g.LifecycleStatus != "ACTIVE" {
			continue
		}
		if g.ValidFrom != nil && now.Before(*g.ValidFrom) {
			continue
		}
		if g.ValidUntil != nil && !now.Before(*g.ValidUntil) {
			continue
		}
		allowed := false
		for _, a := range g.PermittedActionClasses {
			if a == "PROJECT_STATE_MUTATE" {
				allowed = true
				break
			}
		}
		if !allowed {
			continue
		}
		if len(g.ProjectScope) == 0 {
			return true
		}
		for _, p := range g.ProjectScope {
			if p == string(projectID) {
				return true
			}
		}
	}
	return false
}
