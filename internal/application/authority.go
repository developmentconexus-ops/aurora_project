package application

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/authority"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/identity"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

type GrantAuthorityInput struct {
	ProjectID        project.ProjectID
	Action           string
	ExpectedRevision *authority.Revision
	ValidUntil       *time.Time
}
type GrantAuthorityResult struct {
	AuthorityID            string             `json:"authority_id"`
	AuthorityRevision      uint64             `json:"authority_revision"`
	AuthorityStateRevision authority.Revision `json:"authority_state_revision"`
	GoverningGeneration    uint64             `json:"governing_generation"`
}
type RevokeAuthorityInput struct {
	AuthorityID      string
	ExpectedRevision *authority.Revision
}
type RevokeAuthorityResult struct {
	AuthorityID            string             `json:"authority_id"`
	AuthorityRevision      uint64             `json:"authority_revision"`
	AuthorityStateRevision authority.Revision `json:"authority_state_revision"`
	GoverningGeneration    uint64             `json:"governing_generation"`
}

func decodeAuthorityState(raw []byte) (authority.State, error) {
	var state authority.State
	if err := json.Unmarshal(raw, &state); err != nil {
		return authority.State{}, err
	}
	return state, nil
}

func (s *Service) ShowAuthority(ctx context.Context, pass []byte, projectID project.ProjectID, action string) (authority.Snapshot, error) {
	if _, err := s.Inspect(ctx, pass); err != nil {
		return authority.Snapshot{}, err
	}
	current, err := s.State.LoadCurrent(ctx)
	if err != nil {
		return authority.Snapshot{}, err
	}
	state, err := decodeAuthorityState(current.AuthorityJSON)
	if err != nil {
		return authority.Snapshot{}, err
	}
	return authority.Evaluate(state, projectID, action, s.Clock.Now().UTC()), nil
}

func (s *Service) GrantAuthority(ctx context.Context, pass []byte, in GrantAuthorityInput) (GrantAuthorityResult, error) {
	if in.ProjectID == "" || in.Action == "" {
		return GrantAuthorityResult{}, errors.New("project and action are required")
	}
	if _, err := s.Inspect(ctx, pass); err != nil {
		return GrantAuthorityResult{}, err
	}
	current, err := s.State.LoadCurrent(ctx)
	if err != nil {
		return GrantAuthorityResult{}, err
	}
	state, err := decodeAuthorityState(current.AuthorityJSON)
	if err != nil {
		return GrantAuthorityResult{}, err
	}
	if in.ExpectedRevision != nil && *in.ExpectedRevision != state.Revision {
		return GrantAuthorityResult{}, ports.ErrAuthorityRevisionConflict
	}
	now := s.Clock.Now().UTC()
	id, err := randomIdentifier("AUTH-")
	if err != nil {
		return GrantAuthorityResult{}, err
	}
	grant := authority.Grant{AuthorityID: id, AuthorityRevision: 1, SubjectOperatorID: identity.OperatorID(ownerOperatorID), ActorID: "AURORA_CORE", ProjectScope: []project.ProjectID{in.ProjectID}, PermittedActionClasses: []string{in.Action}, Conditions: map[string]string{}, ValidUntil: in.ValidUntil, LifecycleStatus: authority.StatusActive, IssuedAt: now, Provenance: "OWNER_OPERATOR"}
	previous := state.Revision
	next := state
	next.Revision = previous + 1
	next.PredecessorRevision = &previous
	next.Grants = append(append([]authority.Grant(nil), state.Grants...), grant)
	next.ChangedBy = ownerOperatorID
	next.ChangedAt = now
	stored, err := s.commitAuthorityState(ctx, pass, current, next)
	if err != nil {
		return GrantAuthorityResult{}, err
	}
	return GrantAuthorityResult{AuthorityID: id, AuthorityRevision: 1, AuthorityStateRevision: next.Revision, GoverningGeneration: stored.GoverningGeneration}, nil
}

func (s *Service) RevokeAuthority(ctx context.Context, pass []byte, in RevokeAuthorityInput) (RevokeAuthorityResult, error) {
	if in.AuthorityID == "" {
		return RevokeAuthorityResult{}, errors.New("authority id is required")
	}
	if _, err := s.Inspect(ctx, pass); err != nil {
		return RevokeAuthorityResult{}, err
	}
	current, err := s.State.LoadCurrent(ctx)
	if err != nil {
		return RevokeAuthorityResult{}, err
	}
	state, err := decodeAuthorityState(current.AuthorityJSON)
	if err != nil {
		return RevokeAuthorityResult{}, err
	}
	if in.ExpectedRevision != nil && *in.ExpectedRevision != state.Revision {
		return RevokeAuthorityResult{}, ports.ErrAuthorityRevisionConflict
	}
	now := s.Clock.Now().UTC()
	grants := append([]authority.Grant(nil), state.Grants...)
	found := -1
	for i := range grants {
		if grants[i].AuthorityID == in.AuthorityID {
			found = i
			break
		}
	}
	if found < 0 {
		return RevokeAuthorityResult{}, ports.ErrAuthorityNotFound
	}
	grants[found].AuthorityRevision++
	grants[found].LifecycleStatus = authority.StatusRevoked
	grants[found].RevokedAt = &now
	previous := state.Revision
	next := state
	next.Revision = previous + 1
	next.PredecessorRevision = &previous
	next.Grants = grants
	next.ChangedBy = ownerOperatorID
	next.ChangedAt = now
	stored, err := s.commitAuthorityState(ctx, pass, current, next)
	if err != nil {
		return RevokeAuthorityResult{}, err
	}
	return RevokeAuthorityResult{AuthorityID: in.AuthorityID, AuthorityRevision: grants[found].AuthorityRevision, AuthorityStateRevision: next.Revision, GoverningGeneration: stored.GoverningGeneration}, nil
}

func (s *Service) commitAuthorityState(ctx context.Context, pass []byte, current ports.CurrentSnapshot, next authority.State) (ports.AuthorityMutationResult, error) {
	root, err := s.Trust.LoadRootEnvelope(ctx)
	if err != nil {
		return ports.AuthorityMutationResult{}, err
	}
	ork, err := unlockORK(root, pass)
	if err != nil {
		return ports.AuthorityMutationResult{}, err
	}
	defer clear(ork)
	encoded, err := json.Marshal(next)
	if err != nil {
		return ports.AuthorityMutationResult{}, err
	}
	nextSnapshot := current
	nextSnapshot.AuthorityRevision = uint64(next.Revision)
	nextSnapshot.AuthorityJSON = encoded
	nextSnapshot.GoverningGeneration = current.GoverningGeneration + 1
	mac, err := governingMAC(ork, governingSnapshot(nextSnapshot))
	if err != nil {
		return ports.AuthorityMutationResult{}, err
	}
	opID, err := randomIdentifier("OP-")
	if err != nil {
		return ports.AuthorityMutationResult{}, err
	}
	auditID, err := randomIdentifier("AUD-")
	if err != nil {
		return ports.AuthorityMutationResult{}, err
	}
	stored, err := s.State.CommitAuthorityRevision(ctx, ports.AuthorityMutation{OperationID: opID, AuditRecordID: auditID, ExpectedAuthorityRevision: current.AuthorityRevision, NewAuthorityRevision: uint64(next.Revision), AuthorityJSON: encoded, ChangedBy: ownerOperatorID, ChangedAt: next.ChangedAt, ExpectedGeneration: current.GoverningGeneration, NewGeneration: nextSnapshot.GoverningGeneration, GoverningHMAC: mac})
	if err != nil {
		return ports.AuthorityMutationResult{}, err
	}
	anchor, err := s.Trust.LoadAnchor(ctx)
	if err != nil {
		return ports.AuthorityMutationResult{}, err
	}
	anchor.GoverningGeneration = stored.GoverningGeneration
	if next.ChangedAt.After(anchor.ObservedWallTimeHighWater) {
		anchor.ObservedWallTimeHighWater = next.ChangedAt
	}
	anchorBytes, err := anchorMAC(ork, anchor)
	if err != nil {
		return ports.AuthorityMutationResult{}, err
	}
	anchor.HMAC = base64.RawURLEncoding.EncodeToString(anchorBytes)
	if err := s.Trust.PublishAnchor(ctx, anchor); err != nil {
		return ports.AuthorityMutationResult{}, err
	}
	return stored, nil
}
