package portability

import (
	"encoding/json"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/authority"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/identity"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
)

const FormatV1 = "aurora-sovereign-export"

type RootEnvelope struct{Version int `json:"version"`;RootID string `json:"root_id"`;KDF string `json:"kdf"`;MemoryKiB uint32 `json:"memory_kib"`;Iterations uint32 `json:"iterations"`;Parallelism uint8 `json:"parallelism"`;Salt string `json:"salt"`;Nonce string `json:"nonce"`;WrappedORK string `json:"wrapped_ork"`}
type Integrity struct{PayloadSHA256 string `json:"payload_sha256"`}
type ProjectBundle struct{Project project.Project `json:"project"`;Revisions []project.ProjectStateRevision `json:"revisions"`}
type AuthorityBundle struct{CurrentRevision authority.Revision `json:"current_revision"`;Revisions []authority.State `json:"revisions"`}
type TransitionAttempt struct{AttemptID string `json:"attempt_id"`;ProjectID project.ProjectID `json:"project_id"`;ActorID string `json:"actor_id"`;RequestedAt time.Time `json:"requested_at"`;ExpectedRevision *project.StateRevision `json:"expected_state_revision,omitempty"`;RequestedState json.RawMessage `json:"requested_state"`;ProposedNextAction json.RawMessage `json:"proposed_next_action,omitempty"`;AuthorityEvaluationRef string `json:"authority_evaluation_ref,omitempty"`;Result string `json:"result"`;Reason string `json:"reason"`;AcceptedStateRevision *project.StateRevision `json:"accepted_state_revision,omitempty"`}
type Record struct{RecordID string `json:"record_id"`;Kind string `json:"kind"`;OperationID string `json:"operation_id"`;ProjectID project.ProjectID `json:"project_id,omitempty"`;StateRevision *project.StateRevision `json:"state_revision,omitempty"`;AuthorityRevision *authority.Revision `json:"authority_revision,omitempty"`;Outcome string `json:"outcome"`;Reason string `json:"reason"`;Details json.RawMessage `json:"details"`;CreatedAt time.Time `json:"created_at"`}
type StoreState struct{GoverningGeneration uint64 `json:"governing_generation"`;Aurora identity.AuroraIdentity `json:"aurora"`;Projects []ProjectBundle `json:"projects"`;Authority AuthorityBundle `json:"authority"`;Attempts []TransitionAttempt `json:"transition_attempts"`;Records []Record `json:"records"`}
type Document struct{Format string `json:"format"`;Version int `json:"version"`;ExportID string `json:"export_id"`;CreatedAt time.Time `json:"created_at"`;GoverningGeneration uint64 `json:"governing_generation"`;Aurora identity.AuroraIdentity `json:"aurora"`;Projects []ProjectBundle `json:"projects"`;Authority AuthorityBundle `json:"authority"`;Attempts []TransitionAttempt `json:"transition_attempts"`;Records []Record `json:"records"`;OwnerRecovery RootEnvelope `json:"owner_recovery"`;Integrity Integrity `json:"integrity"`}
