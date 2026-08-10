package ports

import (
	"context"
	"errors"
	"time"
)

var (
	ErrAlreadyInitialized = errors.New("Aurora is already initialized")
	ErrNotInitialized     = errors.New("Aurora is not initialized")
)

type BootstrapMutation struct {
	OperationID             string
	AuroraID                string
	OwnerOperatorID         string
	CreatedAt               time.Time
	AuthorityRevision       uint64
	AuthorityJSON           []byte
	GoverningGeneration     uint64
	GoverningDescriptorHMAC []byte
}

type BootstrapResult struct {
	AuroraID            string
	OwnerOperatorID     string
	CreatedAt           time.Time
	AuthorityRevision   uint64
	AuthorityJSON       []byte
	GoverningGeneration uint64
	GoverningHMAC       []byte
}

type CurrentSnapshot = BootstrapResult

type StateStore interface {
	Bootstrap(context.Context, BootstrapMutation) (BootstrapResult, error)
	LoadCurrent(context.Context) (CurrentSnapshot, error)
}
