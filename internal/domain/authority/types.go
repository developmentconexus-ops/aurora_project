package authority

import (
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/identity"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
)

type Revision uint64

type LifecycleStatus string

const (
	StatusActive     LifecycleStatus = "ACTIVE"
	StatusRevoked    LifecycleStatus = "REVOKED"
	StatusSuperseded LifecycleStatus = "SUPERSEDED"
)

type EffectiveStatus string

const (
	EffectiveValid                EffectiveStatus = "VALID"
	EffectiveExpired              EffectiveStatus = "EXPIRED"
	EffectiveRevoked              EffectiveStatus = "REVOKED"
	EffectiveSuperseded           EffectiveStatus = "SUPERSEDED"
	EffectiveRevalidationRequired EffectiveStatus = "REVALIDATION_REQUIRED"
	EffectiveInvalid              EffectiveStatus = "INVALID"
)

type Decision string

const (
	DecisionPermitted            Decision = "PERMITTED"
	DecisionBlocked              Decision = "BLOCKED"
	DecisionNone                 Decision = "NONE"
	DecisionRevalidationRequired Decision = "REVALIDATION_REQUIRED"
)

type Grant struct {
	AuthorityID            string                `json:"authority_id"`
	AuthorityRevision      uint64                `json:"authority_revision"`
	SubjectOperatorID      identity.OperatorID   `json:"subject_operator_id"`
	ActorID                string                `json:"actor_id"`
	ProjectScope           []project.ProjectID   `json:"project_scope"`
	PermittedActionClasses []string              `json:"permitted_action_classes"`
	Conditions             map[string]string     `json:"conditions"`
	ValidFrom              *time.Time            `json:"valid_from"`
	ValidUntil             *time.Time            `json:"valid_until"`
	LifecycleStatus        LifecycleStatus       `json:"lifecycle_status"`
	IssuedAt               time.Time             `json:"issued_at"`
	RevokedAt              *time.Time            `json:"revoked_at"`
	Supersedes             string                `json:"supersedes"`
	Provenance             string                `json:"provenance"`
}

type State struct {
	Revision             Revision  `json:"revision"`
	PredecessorRevision  *Revision `json:"predecessor_revision"`
	Grants               []Grant   `json:"grants"`
	RevalidationRequired bool      `json:"revalidation_required"`
	ChangedBy            string    `json:"changed_by"`
	ChangedAt            time.Time `json:"changed_at"`
}

type Snapshot struct {
	SourceRevision Revision        `json:"authority_state_revision"`
	Decision       Decision        `json:"decision"`
	Status         EffectiveStatus `json:"status"`
	Reasons        []string        `json:"reasons"`
	EvaluatedAt    time.Time       `json:"evaluated_at"`
}

type NextSafeActionProjection struct {
	ProjectID              project.ProjectID         `json:"project_id"`
	StateRevision          project.StateRevision     `json:"state_revision"`
	AuthorityStateRevision Revision                  `json:"authority_state_revision"`
	Action                 *project.ActionDescriptor `json:"action,omitempty"`
	Decision               Decision                  `json:"decision"`
	Reasons                []string                  `json:"reasons"`
	EvaluatedAt            time.Time                 `json:"evaluated_at"`
}
