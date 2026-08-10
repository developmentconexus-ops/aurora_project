package ports

import (
	"context"
	"errors"
	"time"
)

var (
	ErrAlreadyInitialized        = errors.New("Aurora is already initialized")
	ErrNotInitialized            = errors.New("Aurora is not initialized")
	ErrProjectExists             = errors.New("project already exists")
	ErrProjectNotFound           = errors.New("project not found")
	ErrCurrentStateMissing       = errors.New("current project state revision is missing")
	ErrGenerationConflict        = errors.New("governing generation conflict")
	ErrStateRevisionConflict     = errors.New("project state revision conflict")
	ErrAuthorityRevisionConflict = errors.New("authority state revision conflict")
	ErrAuthorityNotFound         = errors.New("authority grant not found")
	ErrUnauthorized              = errors.New("operation is not authorized")
	ErrInvalidTransition         = errors.New("invalid project state transition")
)

type StateEnvelopeRecord struct {
	SchemaVersion, Kind, Summary string
	Payload                      []byte
}
type ProjectStateRecord struct {
	ProjectID              string
	Revision               uint64
	PredecessorRevision    *uint64
	State                  StateEnvelopeRecord
	AcceptedIntentRef      string
	ProposedNextActionJSON []byte
	AcceptedByActor        string
	AcceptedAt             time.Time
	TransitionAttemptID    string
}
type ProjectRecord struct {
	ProjectID, DisplayLabel, ObjectiveSummary string
	CurrentStateRevision                      *uint64
	CurrentState                              *ProjectStateRecord
	CreatedAt, UpdatedAt                      time.Time
}

type BootstrapMutation struct {
	OperationID, AuroraID, OwnerOperatorID string
	CreatedAt                              time.Time
	AuthorityRevision                      uint64
	AuthorityJSON                          []byte
	GoverningGeneration                    uint64
	GoverningDescriptorHMAC                []byte
}
type BootstrapResult struct {
	AuroraID, OwnerOperatorID string
	CreatedAt                 time.Time
	AuthorityRevision         uint64
	AuthorityJSON             []byte
	GoverningGeneration       uint64
	GoverningHMAC             []byte
	Projects                  []ProjectRecord
}
type CurrentSnapshot = BootstrapResult

type CreateProjectMutation struct {
	OperationID                       string
	Project                           ProjectRecord
	ExpectedGeneration, NewGeneration uint64
	GoverningHMAC                     []byte
}
type ProjectTransitionMutation struct {
	AttemptID, AuditRecordID, EvidenceRecordID, ProjectID, ActorID string
	RequestedAt                                                    time.Time
	ExpectedRevision                                               *uint64
	State                                                          ProjectStateRecord
	RequestedStateJSON                                             string
	ProposedNextActionJSON                                         []byte
	AuthorityEvaluationRef                                         string
	ExpectedGeneration, NewGeneration                              uint64
	GoverningHMAC                                                  []byte
}
type ProjectTransitionResult struct {
	ProjectID                          string
	StateRevision, GoverningGeneration uint64
}
type TransitionRejection struct {
	AttemptID, AuditRecordID, ProjectID, ActorID string
	RequestedAt                                  time.Time
	ExpectedRevision                             *uint64
	RequestedStateJSON                           string
	ProposedNextActionJSON                       []byte
	AuthorityEvaluationRef, Reason               string
}

type AuthorityMutation struct {
	OperationID, AuditRecordID                      string
	ExpectedAuthorityRevision, NewAuthorityRevision uint64
	AuthorityJSON                                   []byte
	ChangedBy                                       string
	ChangedAt                                       time.Time
	ExpectedGeneration, NewGeneration               uint64
	GoverningHMAC                                   []byte
}
type AuthorityMutationResult struct{ AuthorityRevision, GoverningGeneration uint64 }

type StateStore interface {
	Bootstrap(context.Context, BootstrapMutation) (BootstrapResult, error)
	LoadCurrent(context.Context) (CurrentSnapshot, error)
	CreateProject(context.Context, CreateProjectMutation) (ProjectRecord, error)
	LoadProject(context.Context, string) (ProjectRecord, error)
	LoadProjectCurrent(context.Context, string) (ProjectRecord, *ProjectStateRecord, error)
	CommitProjectTransition(context.Context, ProjectTransitionMutation) (ProjectTransitionResult, error)
	RecordTransitionRejection(context.Context, TransitionRejection) error
	CommitAuthorityRevision(context.Context, AuthorityMutation) (AuthorityMutationResult, error)
}
