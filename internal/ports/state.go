package ports

import (
	"context"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/identity"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
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
	AuthorityStateJSON       []byte
	CurrentAuthorityRevision uint64
	GoverningGeneration      uint64
	GoverningMAC             []byte
}

type StateStore interface {
	Bootstrap(context.Context, BootstrapMutation) (BootstrapResult, error)
	LoadCurrent(context.Context) (CurrentSnapshot, error)
	CreateProject(context.Context, CreateProjectMutation) (project.Project, error)
	GetProject(context.Context, project.ProjectID) (project.Project, error)
}
