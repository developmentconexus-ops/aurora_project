package ports

import (
	"context"
	"errors"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/authority"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/evidence"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/identity"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/portability"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
)

var (
	ErrStaleStateRevision     = errors.New("stale expected Project state revision")
	ErrStaleAuthorityRevision = errors.New("stale expected authority state revision")
	ErrRestoreTargetNotFresh  = errors.New("restore target already contains an Aurora identity")
)

type BootstrapMutation struct {
	Identity                 identity.AuroraIdentity
	AuthorityStateJSON       []byte
	CurrentAuthorityRevision uint64
	GoverningGeneration      uint64
	GoverningMAC             []byte
}
type BootstrapResult struct {
	Identity            identity.AuroraIdentity
	GoverningGeneration uint64
}
type CreateProjectMutation struct {
	Project            project.Project
	ExpectedGeneration uint64
	NewGeneration      uint64
	GoverningMAC       []byte
	OperationID        string
}
type CurrentSnapshot struct {
	Identity                 identity.AuroraIdentity
	Projects                 []project.Project
	CurrentStates            []project.ProjectStateRevision
	AuthorityStateJSON       []byte
	CurrentAuthorityRevision uint64
	GoverningGeneration      uint64
	GoverningMAC             []byte
}
type ProjectTransitionMutation struct {
	AttemptID              string
	ProjectID              project.ProjectID
	ActorID                string
	RequestedAt            time.Time
	ExpectedRevision       *project.StateRevision
	Revision               project.ProjectStateRevision
	ExpectedGeneration     uint64
	NewGeneration          uint64
	GoverningMAC           []byte
	AuthorityEvaluationRef string
}
type ProjectTransitionResult struct{ Revision project.ProjectStateRevision }
type RecordMutation struct {
	AttemptID              string
	ProjectID              project.ProjectID
	ActorID                string
	RequestedAt            time.Time
	ExpectedRevision       *project.StateRevision
	RequestedState         project.StateEnvelope
	ProposedNextAction     *project.ActionDescriptor
	AuthorityEvaluationRef string
	Result                 string
	Reason                 string
}
type AuthorityMutation struct {
	State                     authority.State
	ExpectedAuthorityRevision authority.Revision
	ExpectedGeneration        uint64
	NewGeneration             uint64
	GoverningMAC              []byte
	OperationID               string
}
type RestoreSnapshot struct {
	State        portability.StoreState
	GoverningMAC []byte
}

type StateStore interface {
	Bootstrap(context.Context, BootstrapMutation) (BootstrapResult, error)
	LoadCurrent(context.Context) (CurrentSnapshot, error)
	CreateProject(context.Context, CreateProjectMutation) (project.Project, error)
	GetProject(context.Context, project.ProjectID) (project.Project, error)
	GetCurrentProjectState(context.Context, project.ProjectID) (*project.ProjectStateRevision, error)
	CommitProjectTransition(context.Context, ProjectTransitionMutation) (ProjectTransitionResult, error)
	RecordNonGoverning(context.Context, RecordMutation) error
	AppendRecord(context.Context, evidence.StoredRecord) error
	CommitAuthorityRevision(context.Context, AuthorityMutation) (authority.State, error)
	ExportLogicalState(context.Context) (portability.StoreState, error)
	BuildFreshFromExport(context.Context, RestoreSnapshot) error
}
