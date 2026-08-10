---
id: DESIGN-AURORA-M0-R6-SOVEREIGN-CORE-MICRODESIGN
title: M0 Sovereign Core Implementation Microdesign
document_type: implementation_microdesign
form: reference
authority: design
status: accepted
accepted_at: 2026-08-09
acceptance_evidence: DOC-AURORA-M0-R6-MICRODESIGN-OPERATOR-ACCEPTANCE
accepted_from_blob: d76cf237211b7fe35c33d1a32f14905e769702a7
version: 0.1.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - accepted M0 R6 implementation design for MIS-M0-SOVEREIGN-CORE-001 v0.1.0
related:
  - DOC-AURORA-STATUS
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DOC-AURORA-M0-R6-OPERATOR-AUTHORIZATION
  - DOC-AURORA-MIS-M0-SOVEREIGN-CORE-001
  - DOC-AURORA-CAP-SOVEREIGN-CORE-REQUIREMENTS
  - DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC
  - DOC-AURORA-CAP-SOVEREIGN-CORE-THREAT-MODEL
  - DOC-AURORA-CAP-SOVEREIGN-CORE-TEST-PLAN
  - DOC-AURORA-CAP-SOVEREIGN-CORE-R5-COVERAGE
  - ADR-AURORA-0003
  - ADR-AURORA-0004
  - ADR-AURORA-0005
  - ADR-AURORA-0006
  - ADR-AURORA-0007
  - ADR-AURORA-0008
  - ADR-AURORA-0009
source_revision: a6769fe8e28dc2dd693f12ad8b9f2460e95b8bc5
mission_contract_revision: 0.1.0
review_triggers:
  - accepted A2 or Mission Contract change
  - accepted ADR change affecting M0
  - implementation finding contradicts a design invariant
  - dependency revalidation changes mechanism semantics
  - new current product requirement changes M0 scope
last_reviewed: 2026-08-09
---

# M0 Sovereign Core — R6 Microdesign

## 1. Purpose and gate boundary

This document turns the accepted `CAP-SOVEREIGN-CORE` A2 package, accepted ADR-0003..0008 and approved `MIS-M0-SOVEREIGN-CORE-001 v0.1.0` into a concrete implementation design.

R6 question:

> Can an implementer execute the approved M0 Contract without inventing material architecture or product behavior while coding?

This Microdesign was explicitly accepted by the operator through `DOC-AURORA-M0-R6-MICRODESIGN-OPERATOR-ACCEPTANCE`. Acceptance fixes the R6 design but does **not** authorize R7 or production/source implementation.

```text
R6 design only
→ no production Go source
→ no implementation execution
→ no R7 authority
```

## 2. Design philosophy

The implementation follows four rules.

### D1 — Stable seams, not speculative systems

Structure early what is cheap now and expensive to remodel later:

- domain ownership;
- application/use-case boundary;
- persistence/trust ports;
- replaceable adapters;
- versioned logical schemas.

Do not prebuild future systems behind those seams.

### D2 — No abstraction without justification

Create an abstraction only when at least one is true:

1. there is more than one current implementation/consumer;
2. an accepted requirement/ADR requires the boundary;
3. the boundary materially improves deterministic testing of a critical property.

Otherwise use the concrete implementation directly.

### D3 — Evolution follows product requirements and evidence

New structure is introduced by current Blueprint/requirements/ADRs/Contract changes or by material implementation/test findings, not by imagined future convenience.

### D4 — Working vertical slices before foundation expansion

No long foundation sequence may delay the first real `aurora init`, restart continuity or Project-state proof. Each implementation slice must end in observable behavior or a contract test that immediately unlocks observable behavior.

## 3. Cross-horizon Presence decision

Concept Bytes JARVIS/HoloMat is a non-governing product/reference input. It reinforces that future voice, desktop, projector/HoloMat and wearable experiences should be replaceable presences rather than Core owners.

M0 does not implement Presence Fabric, voice, HoloMat, camera, desktop UI or Mastra. It only preserves an application boundary that future adapters can call.

```text
M0 CLI ─────────────┐
future Desktop ─────┤
future Voice ───────┤
future HoloMat ─────┼→ application → Sovereign Core
future Glasses ─────┘
```

There is no `PresentationPort` framework in M0. Application operations return structured Go results; the CLI renders them as text or JSON. A later Presence may call the same application boundary or a future transport adapter without changing domain meaning.

## 4. Runtime topology

M0 is one Go module, one binary and one local process.

```text
cmd/aurora
    ↓
adapters/cli
    ↓
application
    ↓
domain + ports
    ↓
adapters/sqlite + adapters/trustfs
```

No microservices, broker, durable workflow engine, DI container, ORM, event bus, plugin framework, Mastra runtime or model call is required.

## 5. Proposed production source tree

The first production tree is:

```text
go.mod
go.sum

cmd/aurora/
  main.go

internal/domain/identity/
  types.go
  rules.go

internal/domain/project/
  types.go
  transition.go

internal/domain/authority/
  types.go
  evaluate.go

internal/domain/portability/
  types.go
  invariants.go

internal/domain/evidence/
  types.go

internal/application/
  service.go
  initialize.go
  create_project.go
  transition_project.go
  authority.go
  inspect.go
  recover.go
  export.go
  restore.go
  migrate.go

internal/ports/
  state.go
  owner_trust.go
  clock.go
  export_protection.go

internal/adapters/sqlite/
  store.go
  schema.go
  mutations.go
  queries.go

internal/adapters/trustfs/
  store.go
  publish.go
  publish_unix.go
  publish_windows.go

internal/adapters/exportage/
  protection.go

internal/adapters/cli/
  cli.go
  commands.go
  render.go

internal/adapters/observability/
  observability.go

schemas/
  project-state-v1.schema.json
  sovereign-export-v1.schema.json

migrations/
  sqlite/
    0001_initial.sql

testdata/
  fixtures/
  adversarial/
  migration/
  golden/
```

File splitting is by responsibility. R7 may combine files that remain very small, but it may not collapse the approved domain/application/ports/adapters dependency direction.

## 6. Dependency direction

Allowed:

```text
cmd/adapters → application → domain + ports
adapters → ports/domain types where required
```

`application` may receive the accepted non-authoritative standard observability APIs (`*slog.Logger`, OTel `trace.Tracer`, OTel `metric.Meter`) from the composition root. Exporter/provider construction remains in `adapters/observability`; domain packages never depend on observability.

Prohibited:

```text
domain → SQLite
domain → filesystem
domain → CLI
domain → OpenTelemetry
domain → age
domain → future Presence/Mastra
```

Application coordinates use cases and owns no second copy of canonical state.

## 7. Domain types

### 7.1 Identity

`internal/domain/identity/types.go` defines strong logical types and immutable identity values:

```go
type AuroraID string
type OperatorID string

type AuroraIdentity struct {
    AuroraID                  AuroraID
    OwnerOperatorID           OperatorID
    CreatedAt                 time.Time
    IdentityRevision          uint64
    CapabilityContractVersion string
}
```

IDs are generated independently from process/session/DB/UI identity.

### 7.2 Project

`internal/domain/project/types.go` defines:

```go
type ProjectID string
type StateRevision uint64

type Project struct {
    ProjectID            ProjectID
    DisplayLabel         string
    ObjectiveSummary     string
    CurrentStateRevision *StateRevision
    CreatedAt            time.Time
    UpdatedAt            time.Time
}

type StateEnvelope struct {
    SchemaVersion string
    Kind          string
    Summary       string
    Payload       json.RawMessage
}

type ActionDescriptor struct {
    ActionClass     string
    Summary         string
    ProjectID       ProjectID
    RequiredAuthorityAction string
    PreconditionRef string
}

type ProjectStateRevision struct {
    ProjectID           ProjectID
    Revision            StateRevision
    PredecessorRevision *StateRevision
    State               StateEnvelope
    AcceptedIntentRef   string
    ProposedNextAction  *ActionDescriptor
    AcceptedByActor     string
    AcceptedAt          time.Time
    TransitionAttemptID string
}
```

`Payload` is opaque Project data. It cannot redefine identity, authority, canonical ownership or transition protocol.

### 7.3 Authority

`internal/domain/authority/types.go` defines the complete current logical authority state; the SQLite adapter may serialize each immutable authority revision as one JSON value rather than normalizing a future authorization platform.

```go
type Revision uint64

type Grant struct {
    AuthorityID            string
    AuthorityRevision      uint64
    SubjectOperatorID      identity.OperatorID
    ActorID                 string
    ProjectScope            []project.ProjectID
    PermittedActionClasses []string
    Conditions              map[string]string
    ValidFrom               *time.Time
    ValidUntil              *time.Time
    LifecycleStatus         string
    IssuedAt                time.Time
    RevokedAt               *time.Time
    Supersedes              string
    Provenance              string
}

type State struct {
    Revision              Revision
    PredecessorRevision   *Revision
    Grants                []Grant
    RevalidationRequired  bool
    ChangedBy             string
    ChangedAt             time.Time
}
```

Pure evaluation functions derive `AuthoritySnapshot` and `NextSafeActionProjection`. A projection never becomes a grant.

### 7.4 Portability/evidence

Portable/export and evidence structs mirror accepted Spec fields but remain independent from SQLite rows and CLI rendering. Claim, Receipt, Evidence and Verdict remain distinct domain concepts even when M0 co-locates audit/evidence records physically.

## 8. Ports

Only four ports are introduced because each has a current architectural justification.

### 8.1 `StateStore`

`internal/ports/state.go` owns the consistency boundary, not one repository per entity.

```go
type StateStore interface {
    Bootstrap(ctx context.Context, in BootstrapMutation) (BootstrapResult, error)
    LoadCurrent(ctx context.Context) (CurrentSnapshot, error)
    CreateProject(ctx context.Context, in CreateProjectMutation) (project.Project, error)
    CommitProjectTransition(ctx context.Context, in ProjectTransitionMutation) (ProjectTransitionResult, error)
    CommitAuthorityRevision(ctx context.Context, in AuthorityMutation) (authority.State, error)
    RecordNonGoverning(ctx context.Context, in RecordMutation) error
    BuildFreshFromExport(ctx context.Context, in RestoreSnapshot) error
}
```

The mutation/result structs named above are defined in `ports/state.go` and contain domain values plus operation IDs/preconditions only; they contain no SQL rows, statements or driver-specific types. Accepted governing mutations include required audit/evidence references in the same SQLite transaction. Rejected/non-governing attempts may be recorded without advancing governing state.

### 8.2 `OwnerTrustStore`

Required because ADR-0008 mandates a physically independent trust boundary. `RootEnvelope` and `Anchor` are port contract structs declared in `ports/owner_trust.go`; neither exposes the plaintext ORK.

```go
type OwnerTrustStore interface {
    LoadRootEnvelope(ctx context.Context) (RootEnvelope, error)
    StoreRootEnvelope(ctx context.Context, RootEnvelope) error
    LoadAnchor(ctx context.Context) (Anchor, error)
    PublishAnchor(ctx context.Context, Anchor) error
}
```

### 8.3 `Clock`

Required for deterministic authority/time tests.

```go
type Clock interface {
    Now() time.Time
}
```

Production uses wall time; tests use a deterministic fake. No generic time framework is introduced.

### 8.4 `ExportProtection`

Required by ADR-0005 because age protection is replaceable independently from logical export meaning.

```go
type ExportProtection interface {
    Protect(ctx context.Context, plaintext []byte, secret []byte) ([]byte, error)
    Unprotect(ctx context.Context, ciphertext []byte, secret []byte) ([]byte, error)
}
```

M0 `exportage` uses age passphrase/scrypt protection with an export secret separate from owner authority credentials.

## 9. Physical state design

### 9.1 Data directory

Resolution order:

```text
--data-dir
→ AURORA_DATA_DIR
→ ~/.aurora
```

Physical layout:

```text
<data-dir>/
  state/
    aurora.db
  trust/
    owner-root.json
    owner-anchor.json
```

The trust files are not a second operational database.

### 9.2 SQLite posture

```text
journal_mode = WAL
synchronous = FULL
foreign_keys = ON
MaxOpenConns = 1
MaxIdleConns = 1
```

SQL is explicit through `database/sql`; no ORM or generic repository framework.

### 9.3 Initial schema

M0 begins with six tables.

#### `core_state`

One governing root row:

```text
singleton_key TEXT PRIMARY KEY CHECK(singleton_key = 'core')
aurora_id TEXT NOT NULL UNIQUE
owner_operator_id TEXT NOT NULL
logical_schema_version INTEGER NOT NULL
current_authority_revision INTEGER NOT NULL
governing_generation INTEGER NOT NULL
governing_descriptor_hmac BLOB NOT NULL
created_at TEXT NOT NULL
updated_at TEXT NOT NULL
```

#### `projects`

```text
project_id TEXT PRIMARY KEY
display_label TEXT NOT NULL
objective_summary TEXT NOT NULL
current_state_revision INTEGER NULL
created_at TEXT NOT NULL
updated_at TEXT NOT NULL
```

#### `project_state_revisions`

```text
project_id TEXT NOT NULL
state_revision INTEGER NOT NULL
predecessor_revision INTEGER NULL
state_schema_version TEXT NOT NULL
state_kind TEXT NOT NULL
state_summary TEXT NOT NULL
state_payload_json TEXT NULL
accepted_intent_ref TEXT NULL
proposed_next_action_json TEXT NULL
accepted_by_actor TEXT NOT NULL
accepted_at TEXT NOT NULL
transition_attempt_id TEXT NOT NULL UNIQUE
PRIMARY KEY(project_id, state_revision)
FOREIGN KEY(project_id) REFERENCES projects(project_id)
```

#### `transition_attempts`

```text
attempt_id TEXT PRIMARY KEY
project_id TEXT NOT NULL
actor_id TEXT NOT NULL
requested_at TEXT NOT NULL
expected_state_revision INTEGER NULL
requested_state_json TEXT NOT NULL
proposed_next_action_json TEXT NULL
authority_evaluation_ref TEXT NULL
result TEXT NOT NULL
reason TEXT NOT NULL
accepted_state_revision INTEGER NULL
FOREIGN KEY(project_id) REFERENCES projects(project_id)
```

#### `authority_revisions`

M0 stores each complete immutable authority state revision as canonical JSON instead of prebuilding normalized grant/state/member tables.

```text
authority_revision INTEGER PRIMARY KEY
predecessor_revision INTEGER NULL
authority_state_json TEXT NOT NULL
changed_by TEXT NOT NULL
changed_at TEXT NOT NULL
```

#### `records`

Physical co-location for attributable audit/evidence/recovery/export/restore/migration metadata:

```text
record_id TEXT PRIMARY KEY
kind TEXT NOT NULL
operation_id TEXT NOT NULL
project_id TEXT NULL
state_revision INTEGER NULL
authority_revision INTEGER NULL
outcome TEXT NOT NULL
reason TEXT NOT NULL
details_json TEXT NOT NULL
created_at TEXT NOT NULL
```

Domain semantics still distinguish Audit/Evidence/Recovery/etc.; `kind` only co-locates their M0 storage.

No additional table is introduced without a current requirement/test/query need.

## 10. Governing integrity design

SQLite provides transaction atomicity/durability; Aurora does not rebuild a transaction journal above SQLite.

ADR-0008 requires one authenticated governing descriptor so raw DB modification alone cannot silently redefine governing truth.

The descriptor is one compact **current governing logical snapshot**, not a list of physical SQLite bytes and not merely revision pointers. It binds:

- immutable Aurora/owner identity binding required by M0;
- `governing_generation`;
- every Project's current canonical metadata, current revision number and current `StateEnvelope`/proposed-next-action/acceptance attribution;
- the complete current logical `authority.State`, including revalidation status.

Conceptually:

```json
{
  "version": 1,
  "aurora": {"aurora_id":"AUR-...","owner_operator_id":"OWNER-..."},
  "governing_generation": 14,
  "projects": [
    {
      "project_id":"PRJ-...",
      "display_label":"...",
      "current_state_revision":3,
      "current_state": {"schema_version":"1","kind":"...","summary":"...","payload":{}},
      "proposed_next_action": null,
      "accepted_by_actor":"...",
      "accepted_at":"..."
    }
  ],
  "authority": {"revision":5,"revalidation_required":false,"grants":[]}
}
```

Project entries and any set-like fields use stable documented ordering before JCS. The logical object is JCS-canonicalized and authenticated **once** with HMAC-SHA-256 using the ORK-derived governing-state key. Therefore changing the contents of the current Project or authority revision without the ORK invalidates governing integrity even if revision numbers are left unchanged.

M0 does **not** add per-table/per-row HMACs, Merkle trees, hash chains, event-sourced integrity or a custom transaction protocol.

## 11. Owner root/trust files

### 11.1 `owner-root.json`

Contains only:

```text
version
root_id
kdf = argon2id
memory = 64 MiB
iterations = 3
parallelism = 4
salt
nonce
wrapped_ork
```

Decoding validates/allowlists version and KDF parameters before Argon2 allocation. Passphrase rotation rewraps the same ORK.

### 11.2 `owner-anchor.json`

Contains only:

```text
version
root_id
aurora_id
governing_generation
observed_wall_time_high_water
hmac
```

It is a small authenticated high-water artifact, not a database/history log.

### 11.3 Trust classification

Production implements the accepted ADR-0008 vocabulary and no stronger speculative protocol:

```text
NORMAL
ANCHOR_LAG
STATE_ROLLBACK
INVALID_DB_MAC
INVALID_ANCHOR_MAC
MISSING_ANCHOR
TIME_UNTRUSTED
REVALIDATION_REQUIRED
```

Normal governing mutation preflights trust before writing. A SQLite commit followed by process death before anchor publication is classified by the next process instead of guessed as success.

No additional `anchor + 1` restriction or stronger time-high-water write policy is introduced unless a test/finding proves it necessary.

## 12. Trust-file publication

`trustfs` uses a small explicit atomic publication helper:

1. serialize the complete next root/anchor;
2. create a temp file in the target directory;
3. write all bytes;
4. call `File.Sync()`;
5. close;
6. atomically rename/replace target;
7. on Unix, open and sync the parent directory;
8. on Windows, after temp-file `Sync` + close, replace the target with `MoveFileExW` using `MOVEFILE_REPLACE_EXISTING | MOVEFILE_WRITE_THROUGH` through `golang.org/x/sys/windows`; do not claim Unix-style parent-directory fsync semantics on Windows beyond the R7 evidence.

R7 fault tests verify the supported target classes. No generalized filesystem durability framework is created.

## 13. Governing mutation protocol

The implementation remains intentionally small:

```text
trust preflight
→ BEGIN SQLite transaction
→ validate expected state/authority revision again inside transaction
→ write governing mutation + required records
→ governing_generation++
→ recompute one governing descriptor HMAC
→ COMMIT
→ publish owner-anchor
→ report success only after required result is observed
```

SQLite owns intra-DB atomicity. The only cross-medium asymmetry to classify is DB commit versus trust-anchor publication, as already accepted in ADR-0008.

Rejected transitions do not advance governing generation or anchor. They may append a non-governing transition/audit record.

## 14. Portability design

### 14.1 One logical document

M0 export is one logical JSON document; no ZIP/TAR/member framework.

```json
{
  "format": "aurora-sovereign-export",
  "version": 1,
  "created_at": "...",
  "aurora": {},
  "projects": [],
  "authority": {},
  "records": [],
  "owner_recovery": {"wrapped_root_envelope": {}},
  "integrity": {"payload_sha256": "..."}
}
```

Logical content is validated by `schemas/sovereign-export-v1.schema.json`. `integrity.payload_sha256` is SHA-256 over the JCS-canonical JSON of the export document with the top-level `integrity` member omitted; verification reconstructs that same canonical input before accepting the package. The complete document is then protected externally with age.

Normal sensitive file:

```text
*.aurora.age
```

### 14.2 Export secret

M0 uses an operator-provided export passphrase/secret independent from the owner authority passphrase. The CLI reads it from a non-echoing terminal prompt; tests inject deterministic secrets through the application/adapter boundary, not hardcoded production defaults.

### 14.3 Restore

Restore is staged into a fresh target data directory:

```text
age decrypt
→ parse JSON
→ schema validate
→ verify digest
→ validate logical version
→ validate identity collision
→ recover wrapped ORK lineage
→ build fresh SQLite store
→ authority REVALIDATION_REQUIRED
→ publish recovered root envelope
→ do NOT import historical owner-anchor as current freshness
→ owner explicit revalidation creates new authority revision/generation/anchor
```

A failed staged restore never replaces a currently governing target directory.

## 15. Migration design

No migration framework/registry/graph.

Logical migration uses explicit source→target functions:

```text
v1 → v2 only when v2 exists
v2 → v3 only when v3 exists
```

Unsupported versions fail clearly. The M0 test corpus includes a deterministic prior-version compatibility fixture and a semantic-mutation fixture proving stable identity/authority meaning cannot be silently changed.

Physical SQLite DDL migrations are explicit numbered SQL files. `0001_initial.sql` is the only initial migration.

## 16. CLI/proof surface

M0 uses a minimal standard-library CLI, not Cobra/TUI/web/interactive-shell infrastructure.

Global options:

```text
--data-dir <path>
--json
```

Commands:

```text
aurora init
aurora status

aurora project create
aurora project show
aurora project set-state

aurora authority show
aurora authority grant
aurora authority revoke
aurora authority revalidate

aurora export
aurora restore
aurora migrate
```

CLI parsing/rendering stays inside `adapters/cli`. Application returns structured results used by text and JSON renderers. No domain rule prints to stdout/stderr.

## 17. Observability design

M0 uses:

```text
structured logs: log/slog
traces/metrics: OpenTelemetry Go API/SDK
backend/exporter: optional
```

No Collector, Grafana or telemetry backend is required for correctness.

Initial stable correlation fields:

```text
operation_id
aurora_id
project_id when applicable
state_revision when applicable
authority_revision when applicable
operation_name
result
reason
```

No owner passphrase, ORK, raw sensitive Project payload or export secret is logged/traced.

Initial metrics remain few and decision-oriented, e.g. operation total/failure total. New metrics require a stated proof/decision purpose.

## 18. Vertical implementation slices

R7, if separately authorized, executes the following order. Each slice must finish green and usable before the next.

### Slice 0 — production skeleton

Real Go module/binary, approved package boundaries, `aurora --help`, build/vet/test baseline.

### Slice 1 — sovereign bootstrap

Real SQLite + real owner root/anchor + `aurora init` + restart `aurora status` showing the same Aurora identity.

### Slice 2 — Project continuity

Create/show one Project, terminate process, recover same Project identity.

### Slice 3 — revision-bound Project state

Create state R1/R2, preserve history, reject stale/malformed/identity-mutating transitions without current-state mutation.

### Slice 4 — minimum authority

Grant/revoke/expire/scope/evaluate + next-safe-action behavior; restart preserves revocation/current evaluation.

### Slice 5 — trust anomalies

Reproduce accepted ADR-0008 anomaly classes from SPK-002 without expanding the protocol.

### Slice 6 — fresh-process integrated continuity

Automated real-process bootstrap→Project→state→authority→kill→fresh recovery.

### Slice 7 — export/restore

Real `.aurora.age`, fresh staged restore, `REVALIDATION_REQUIRED`, non-owner deny, owner revalidation/new authority revision.

### Slice 8 — migration

Explicit valid logical migration and semantic-mutation rejection fixture.

### Slice 9 — audit/evidence/observability completion

Complete required metadata, redaction and optional telemetry failure behavior.

### Slice 10 — fault/security matrix

Systematic negative/fault coverage from the accepted Test Plan with minimal purpose-specific fault hooks; no chaos framework.

### Slice 11 — full M0 Golden Proof

Run the real binary, real filesystem, real SQLite and real process termination on a fixed revision. No in-memory-only substitute.

## 19. Test strategy

TDD is slice-local, not an 84-test upfront project.

```text
select slice
→ write relevant failing domain/contract/integration test
→ verify RED
→ minimal implementation
→ verify GREEN
→ real vertical journey
→ commit/review
```

Test layers:

```text
domain unit tests
→ pure invariants/evaluation

port contract tests
→ reusable adapter expectations where replacement is real

adapter integration tests
→ SQLite/trustfs/age behavior

process/fault tests
→ kill/restart/corrupt/rollback/restore

Golden Proof
→ complete operator-visible product journey
```

The accepted 84-test catalog is a final coverage obligation, not an instruction to create 84 test files before the first executable behavior.

## 20. Fault hooks

Fault injection is purpose-specific and test-only. Allowed seams include deterministic callbacks or test helpers around:

- before/after SQLite commit;
- before anchor publication;
- time source;
- corrupted/removed trust file fixture;
- corrupted/export input fixture;
- process kill via external test runner.

No production chaos/fault framework is introduced.

## 21. Dependency policy

R7 must resolve and pin current supported versions before the first production source commit, starting from the evidence-qualified R4 baseline and rerunning material conformance when versions change.

Allowed dependency classes are only those required by accepted decisions:

- Go standard library;
- `modernc.org/sqlite` compatible exact module set;
- `golang.org/x/crypto` where required by ADR-0008;
- `golang.org/x/sys/windows` for the Windows trust-file replacement primitive;
- maintained age implementation;
- maintained JSON Schema 2020-12 validator;
- maintained RFC 8785/JCS implementation;
- OpenTelemetry Go API/SDK.

No framework is added merely to decorate architecture. A new dependency requires a current requirement/ADR/testability justification and license/maintenance review.

## 22. Change/replan rules

R7 implementers may simplify internals when all accepted semantics/tests remain unchanged.

Replan is mandatory when implementation would require any of:

- change to accepted domain behavior;
- weakening an authority/integrity/restore invariant;
- replacing an accepted mechanism class;
- new cloud/network/process topology;
- M1+ Presence/memory/agent/Harness scope;
- material new dependency/framework that changes ownership/topology;
- inability to satisfy a requirement without redesign.

Do not make product/architecture changes silently to make tests pass.

## 23. Rollback strategy

Implementation is delivered in independently green vertical-slice commits. A failed later slice can be reverted without invalidating prior accepted slice behavior.

Persistent schema changes are forward-migrated explicitly and tested against fixtures. A migration that cannot preserve accepted semantics blocks rather than silently coercing data.

Architecture spike code remains disposable evidence and is not promoted by copy/merge as production implementation.

## 24. Mission-criterion design coverage

| Mission criterion | Primary design realization | Vertical proof slice |
|---|---|---|
| `CRIT-001` Stable identity/scope | §§7.1, 9, 11 | Slices 1–2 |
| `CRIT-002` Canonical state ownership | §§7.2, 9–10 | Slices 2–3 |
| `CRIT-003` Revision-bound transitions | §§7.2, 8.1, 13 | Slice 3 |
| `CRIT-004` Authority/next-safe-action | §§7.3, 11, 13 | Slices 4–5 |
| `CRIT-005` Fresh-process recovery | §§9–13 | Slice 6 |
| `CRIT-006` Export/restore/migration | §§14–15 | Slices 7–8 |
| `CRIT-007` Security/sovereignty/secrets | §§10–14, 17 | Slices 1, 4–5, 7, 10 |
| `CRIT-008` Audit/evidence/telemetry | §§9.3, 17, 19 | Slice 9 |
| `CRIT-009` Reliability/fault containment | §§12–13, 20, 23 | Slices 5–6, 10 |
| `CRIT-010` Architecture guards | §§2–6, 21–22 | all slices / static review |
| `CRIT-011` Documentation/traceability | §§1–2, 21–25 | R6/R7 reviews |
| `CRIT-012` Complete M0 Golden Proof | §§18–20 | Slice 11 |

The task-by-task Implementation Plan created after written Microdesign approval must expand this to explicit requirement/test allocations; this table proves no Mission criterion is structurally orphaned at the design level.

## 25. R6 self-review checklist

The written design must pass these checks before an R6 verdict can be proposed:

- exact M0 production source tree identified;
- domain/application/ports/adapters direction explicit;
- physical persistence remains minimal and consistent with ADR-0007/0008;
- no custom transaction/event/authorization framework;
- bounded Argon2 envelope parsing specified;
- target trust-file publication semantics specified without overclaim;
- portable logical format and restore staging exact enough to implement;
- migration path explicit without speculative framework;
- CLI/proof surface exact enough for Golden Proof;
- observability non-authoritative and optional-backend safe;
- implementation slices make Aurora work early;
- accepted fault/security cases have concrete implementation seams;
- future Presence/Mastra evolution remains possible without prebuilding it;
- no R7 implementation authority implied.

## 26. Written-review gate

This synthesis incorporates the operator-approved conversational direction from the R6 design discussion, including the explicit anti-overengineering correction. The written form was explicitly accepted by the operator. The remaining R6 work is exact task allocation plus adversarial plan review.

After written approval:

```text
Microdesign approved for R6 review
→ create exact task-by-task Implementation Plan
→ map tasks to Mission criteria/requirements/tests
→ adversarial R6 review
→ PASS | FAIL | BLOCKED
→ STOP before R7
```
