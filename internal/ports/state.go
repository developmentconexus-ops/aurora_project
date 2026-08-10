package ports

import (
	"context"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/identity"
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

type CurrentSnapshot struct {
	Identity                 identity.AuroraIdentity
	AuthorityStateJSON       []byte
	CurrentAuthorityRevision uint64
	GoverningGeneration      uint64
	GoverningMAC             []byte
}

type StateStore interface {
	Bootstrap(context.Context, BootstrapMutation) (BootstrapResult, error)
	LoadCurrent(context.Context) (CurrentSnapshot, error)
}
