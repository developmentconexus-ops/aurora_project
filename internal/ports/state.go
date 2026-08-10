package ports

import (
	"context"
	"errors"
	"time"
)

var (
	ErrAlreadyInitialized = errors.New("Aurora is already initialized")
	ErrNotInitialized     = errors.New("Aurora is not initialized")
	ErrProjectExists      = errors.New("project already exists")
	ErrProjectNotFound    = errors.New("project not found")
	ErrCurrentStateMissing = errors.New("current project state revision is missing")
	ErrGenerationConflict = errors.New("governing generation conflict")
)

type StateEnvelopeRecord struct {
	SchemaVersion string
	Kind          string
	Summary       string
	Payload       []byte
}

type ProjectStateRecord struct {
	ProjectID              string
	Revision               uint64
	PredecessorRevision    *uint64
	State                   StateEnvelopeRecord
	AcceptedIntentRef       string
	ProposedNextActionJSON  []byte
	AcceptedByActor         string
	AcceptedAt              time.Time
	TransitionAttemptID     string
}

type ProjectRecord struct {
	ProjectID            string
	DisplayLabel         string
	ObjectiveSummary     string
	CurrentStateRevision *uint64
	CurrentState         *ProjectStateRecord
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

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
	Projects            []ProjectRecord
}

type CurrentSnapshot = BootstrapResult

type CreateProjectMutation struct {
	OperationID        string
	Project            ProjectRecord
	ExpectedGeneration uint64
	NewGeneration      uint64
	GoverningHMAC      []byte
}

type StateStore interface {
	Bootstrap(context.Context, BootstrapMutation) (BootstrapResult, error)
	LoadCurrent(context.Context) (CurrentSnapshot, error)
	CreateProject(context.Context, CreateProjectMutation) (ProjectRecord, error)
	LoadProject(context.Context, string) (ProjectRecord, error)
	LoadProjectCurrent(context.Context, string) (ProjectRecord, *ProjectStateRecord, error)
}
