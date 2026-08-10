package authority

import (
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/identity"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
)

type AuthorityID string
type Revision uint64
type LifecycleStatus string
type EffectiveStatus string
type ProjectionDecision string

const (
	Active     LifecycleStatus = "ACTIVE"
	Revoked    LifecycleStatus = "REVOKED"
	Superseded LifecycleStatus = "SUPERSEDED"

	Valid                      EffectiveStatus = "VALID"
	Expired                    EffectiveStatus = "EXPIRED"
	RevokedStatus              EffectiveStatus = "REVOKED"
	SupersededStatus           EffectiveStatus = "SUPERSEDED"
	RevalidationRequiredStatus EffectiveStatus = "REVALIDATION_REQUIRED"
	Invalid                    EffectiveStatus = "INVALID"

	Permitted            ProjectionDecision = "PERMITTED"
	Blocked              ProjectionDecision = "BLOCKED"
	None                 ProjectionDecision = "NONE"
	ProjectionRevalidate ProjectionDecision = "REVALIDATION_REQUIRED"
)

type Grant struct {
	AuthorityID            AuthorityID          `json:"authority_id"`
	AuthorityRevision      uint64               `json:"authority_revision"`
	SubjectOperatorID      identity.OperatorID  `json:"subject_operator_id"`
	ActorID                 string               `json:"actor_id"`
	ProjectScope            []project.ProjectID  `json:"project_scope"`
	PermittedActionClasses []string             `json:"permitted_action_classes"`
	Conditions              map[string]string    `json:"conditions,omitempty"`
	ValidFrom               *time.Time           `json:"valid_from,omitempty"`
	ValidUntil              *time.Time           `json:"valid_until,omitempty"`
	LifecycleStatus         LifecycleStatus      `json:"lifecycle_status"`
	IssuedAt                time.Time            `json:"issued_at"`
	RevokedAt               *time.Time           `json:"revoked_at,omitempty"`
	Supersedes              AuthorityID          `json:"supersedes,omitempty"`
	Provenance              string               `json:"provenance"`
}

func (g Grant) PermitsActionClass(action string) bool {
	for _, candidate := range g.PermittedActionClasses {
		if candidate == action { return true }
	}
	return false
}

type State struct {
	Revision             Revision  `json:"revision"`
	PredecessorRevision  *Revision `json:"predecessor_revision,omitempty"`
	Grants               []Grant   `json:"grants"`
	RevalidationRequired bool      `json:"revalidation_required"`
	ChangedBy            string    `json:"changed_by"`
	ChangedAt            time.Time `json:"changed_at"`
}

type Snapshot struct {
	SourceRevision       Revision        `json:"source_authority_state_revision"`
	EvaluatedAt          time.Time       `json:"evaluated_at"`
	Status               EffectiveStatus `json:"status"`
	MatchingAuthorityIDs []AuthorityID   `json:"matching_authority_ids,omitempty"`
	Reasons              []string        `json:"reasons"`
}

type NextSafeActionProjection struct {
	ProjectID              project.ProjectID         `json:"project_id"`
	StateRevision          project.StateRevision     `json:"state_revision"`
	AuthorityStateRevision Revision                  `json:"authority_state_revision"`
	Action                 *project.ActionDescriptor `json:"action,omitempty"`
	Decision               ProjectionDecision        `json:"decision"`
	Reasons                []string                  `json:"reasons"`
	EvaluatedAt            time.Time                 `json:"evaluated_at"`
}
