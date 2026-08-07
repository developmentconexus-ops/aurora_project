---
id: DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC
title: CAP-SOVEREIGN-CORE Capability Specification
document_type: capability_spec
form: reference
authority: specification
status: proposed
version: 0.1.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - proposed reusable behavior and logical design of CAP-SOVEREIGN-CORE
  - proposed M0 sovereign Core domain and lifecycle semantics
related:
  - DOC-AURORA-STATUS
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DOC-AURORA-CAP-SOVEREIGN-CORE-APPLICABILITY
  - DOC-AURORA-CAP-SOVEREIGN-CORE-REQUIREMENTS
  - DOC-AURORA-CAP-SOVEREIGN-CORE-THREAT-MODEL
  - DOC-AURORA-CAP-SOVEREIGN-CORE-TEST-PLAN
  - DOC-AURORA-CAP-SOVEREIGN-CORE-R3-COVERAGE
  - DOC-AURORA-M0-R3-OPERATOR-AUTHORIZATION
  - ADR-AURORA-0001
  - ADR-AURORA-0002
source_revision: 9ea8adf5c115f54071d7e36e312695d19420d8b0
review_triggers:
  - R2 requirement change
  - accepted ADR affecting M0
  - R3 finding or R4 decision contradicting an assumption
  - M0 roadmap meaning change
last_reviewed: 2026-08-07
---

# CAP-SOVEREIGN-CORE — Capability Specification

## 1. Status, purpose and R3 boundary

This document is the proposed reusable Capability/System Specification for:

```text
CAP-SOVEREIGN-CORE
M0 — Sovereign Core Walking Skeleton
```

Fixed R3 source baseline:

```text
9ea8adf5c115f54071d7e36e312695d19420d8b0
```

Purpose:

> Define the complete logical behavior and ownership needed for Aurora to preserve sovereign identity, Project operational state, current authority semantics and a safe next action across process restart, and to export/restore/migrate that state without making a model, Harness, interface, telemetry stream or storage product the authority.

This R3 specification fixes **logical semantics and boundaries**. It does not select implementation language, process topology, database/storage engine, state-versus-event persistence pattern, serialization format, event transport, telemetry backend, backup technology, migration tooling, durable workflow engine or UI technology. Those remain R4 decisions where applicable.

R3 does not authorize Architecture Spike execution, Mission Contract creation, Microdesign or implementation.

---

## 2. Use cases

### UC-01 — Initialize sovereign Aurora

A valid operator bootstrap initializes one Aurora identity and establishes the initial owner authority context without depending on an LLM, Harness or cloud service.

### UC-02 — Create and inspect a Project

The operator creates one Project with stable identity and can inspect its current accepted operational state, state revision, authority status and next-safe-action projection.

### UC-03 — Accept a valid state transition

A transition request that names the current revision, is structurally valid and is permitted by current authority becomes a new accepted immutable state revision with attributable audit/evidence references.

### UC-04 — Reject an invalid transition

A stale-revision, malformed, identity-mutating or unauthorized transition is rejected without changing current accepted state, authority state or the next-safe-action projection.

### UC-05 — Recover after process death

After all Aurora processes terminate, a fresh process reconstructs the same Aurora identity, Project identity, current accepted state/revision, current valid authority state and next safe action solely from canonical durable state plus required local trusted inputs.

### UC-06 — Export and restore

The operator exports the M0 continuity state with logical versions and integrity metadata. Restore validates compatibility, integrity, identity collision and authority safety before any restored state can become current.

### UC-07 — Migrate state explicitly

A supported source-to-target logical version migration preserves stable identities, state meaning, authority semantics, provenance and evidence references and produces a migration result suitable for verification.

---

## 3. Goals

1. Preserve stable Aurora and Project identities across restart and restore.
2. Establish one canonical owner for M0 operational state and authority semantics.
3. Make every state mutation revision-bound, attributable, rejectable and observable.
4. Keep current state, authority, events, audit, telemetry and evidence logically distinct.
5. Derive next safe action from current accepted state plus current effective authority.
6. Recover after process termination without transcript/model/Harness/runtime-local memory.
7. Make export/restore/migration portable at the semantic level and safe against stale authority.
8. Define security, failure and evidence behavior before R4 chooses mechanisms.
9. Remain local-first, single-user and framework/storage neutral.
10. Support the complete M0 Golden Proof without importing M1+ capabilities.

---

## 4. Non-goals

M0 does not provide:

- governed conversational/project memory or Context Builder;
- model routing or model-dependent reasoning;
- Capability Registry, Provider trust lifecycle, AHDK or conformance;
- Mission/Delegation orchestration;
- external Effect Gateway, PDP or Credential Broker;
- cloud topology or multi-node synchronization;
- multi-Presence handoff, voice, vision or sensors;
- physical devices or laboratory control;
- long-running campaigns, budgets or durable timers;
- self-improvement;
- multi-tenancy;
- a generic workflow platform;
- a generalized event platform;
- a generalized authorization platform.

A future capability may reuse or extend the M0 semantics, but M0 does not pre-build those systems.

---

## 5. Applicability and requirement baseline

R1 classified all 294 accepted constitutional requirements and selected:

```text
78 APPLIES
49 PARTIALLY_APPLIES
127 active constitutional source rows
```

R2 derived 122 proposed atomic requirements from that active set. R3 allocates all 122 requirements to the logical mechanisms in this Spec and to the separate Capability Test Plan.

Canonical R3 package:

```text
APPLICABILITY.md
REQUIREMENTS.md
SPEC.md
THREAT-MODEL.md
TEST-PLAN.md
R3-COVERAGE.md
```

`SPEC.md` owns reusable behavior. The threat model and test plan are specialized parts of the R3 package and MUST NOT redefine this Spec.

---

## 6. Design alternatives considered

### 6.1 Documentation/package shape

**A — One monolithic R3 document**

Rejected because behavior, threat analysis, test design and coverage would become difficult to review independently.

**B — One Capability Spec plus specialized threat/test/coverage artifacts**

Selected. One reusable Capability remains the semantic owner while specialized artifacts stay independently reviewable.

**C — Split M0 into identity/state/authority/recovery/evidence sub-capabilities**

Rejected for M0 because it would generalize platform boundaries before multiple independent consumers exist.

### 6.2 State semantics

**A — One mutable opaque state blob**

Rejected because revision history, concurrency/precondition checks, authority provenance and evidence would become ambiguous.

**B — Logical immutable accepted-state revisions with an explicit current pointer**

Selected as a semantic model. This does **not** select event sourcing or snapshot persistence; R4 may realize the semantics through any mechanism that preserves the invariants.

**C — Full event-sourced product model**

Not selected. Events remain distinct from canonical current state, and R4 may consider an event-oriented mechanism only if evidence proves it useful without making events the sole product authority.

---

## 7. Domain model

### 7.1 `AuroraIdentity`

Canonical owner: **Identity Module**.

Required semantic fields:

| Field | Meaning |
|---|---|
| `aurora_id` | stable immutable Aurora identity |
| `created_at` | recorded initialization time |
| `identity_revision` | logical identity metadata revision; does not replace `aurora_id` |
| `owner_operator_id` | stable reference to Leandro/operator identity |
| `capability_contract_version` | logical CAP-SOVEREIGN-CORE semantic version |

`aurora_id` MUST NOT be derived from process ID, hostname, database identifier, model session or UI session.

### 7.2 `OperatorIdentityRef`

Canonical owner: **Identity Module**.

M0 requires one stable logical operator identity reference for Leandro. R3 defines the semantic identity; R4 must decide the concrete authentication mechanism. A caller's technical access does not establish this identity by itself.

Required fields:

- `operator_id`;
- `identity_kind = OWNER_OPERATOR` for M0;
- provenance/reference to the authentication boundary that established the current caller context.

### 7.3 `Project`

Canonical owner: **Project State Module**.

Required semantic fields:

| Field | Meaning |
|---|---|
| `project_id` | stable immutable Project identity |
| `display_label` | operator-readable mutable label |
| `objective_summary` | bounded operator-readable project objective/context for M0 |
| `current_state_revision` | reference to the governing `ProjectStateRevision` |
| `created_at` | Project creation time |
| `updated_at` | last accepted state change time |

M0 does not model full tasks, hypotheses, memory or artifact contents inline. References may be preserved where needed for proof.

### 7.4 `ProjectStateRevision`

Canonical owner: **Project State Module**.

An accepted operational state is immutable once accepted. A later change creates a new revision.

Required semantic fields:

| Field | Meaning |
|---|---|
| `project_id` | owning Project |
| `state_revision` | stable logical revision identity/order |
| `predecessor_revision` | prior accepted revision, absent only for the first revision |
| `accepted_state` | bounded structured M0 state envelope |
| `accepted_intent_ref` | optional reference distinguishing accepted operator intent/input from derived action |
| `proposed_next_action` | optional `ActionDescriptor`; it is a candidate, not permission |
| `accepted_by_actor` | actor/source attribution |
| `accepted_at` | recorded acceptance time |
| `transition_attempt_id` | transition attempt that produced the revision |

The `accepted_state` envelope MUST be semantically versioned but R3 does not select a serialization language. Core treats domain-specific content as data; it validates the envelope, revision/identity invariants and M0 transition rules rather than inventing domain workflow semantics.

### 7.5 `TransitionAttempt`

Canonical owner: **Project State Module**.

Required fields:

- `attempt_id`;
- `project_id`;
- `actor_id`;
- `requested_at`;
- `expected_state_revision`;
- requested `accepted_state` envelope;
- optional proposed `ActionDescriptor`;
- authority evaluation reference;
- lifecycle result and reason classification.

Lifecycle:

```text
RECEIVED
→ VALIDATING
→ ACCEPTED | REJECTED | FAILED
```

`ACCEPTED` creates exactly one new accepted state revision. `REJECTED` and `FAILED` create none.

A stale `expected_state_revision` is a canonical M0 invalid-transition case and MUST be rejected.

### 7.6 `ActionDescriptor`

Canonical owner: **Project State Module** for the candidate; **Authority Module** owns its permission evaluation.

Required semantic fields:

- stable action identifier or action class;
- operator-readable summary;
- Project/resource scope;
- required authority action class;
- optional precondition reference.

An `ActionDescriptor` describes what may happen next. It is not an Authority Grant and cannot execute an external effect in M0.

### 7.7 `AuthorityGrantRecord`

Canonical owner: **Authority Module**.

M0 defines the minimum authority profile necessary to decide internal Project/state operations and next-safe-action permission. It is not the later full Mission/Delegation/effect grant model.

Required semantic fields:

| Field | Meaning |
|---|---|
| `authority_id` | stable grant identity |
| `authority_revision` | grant/state revision |
| `subject_operator_id` | Leandro/operator subject for M0 |
| `actor_id` | actor receiving/using the permission, normally Aurora Core in M0 |
| `project_scope` | Project scope or explicit Aurora-instance scope |
| `permitted_action_classes` | bounded internal M0 action classes |
| `resource_scope` | resource/project scope where applicable |
| `conditions` | semantic conditions that must hold |
| `valid_from` | start of validity if applicable |
| `valid_until` | expiry boundary if applicable |
| `lifecycle_status` | `ACTIVE`, `REVOKED` or `SUPERSEDED` |
| `issued_at` | issue time |
| `revoked_at` | revocation time when applicable |
| `supersedes` | replaced authority reference when applicable |
| `provenance` | source/operator attribution |

M0 internal authority action classes are limited to operations needed for this slice, such as Project-state mutation, authority administration, export, restore and migration. The exact transport/API encoding remains open.

### 7.8 `AuthorityState`

Canonical owner: **Authority Module**.

`AuthorityState` is the current canonical authority revision for the M0 scope and references the currently governing grant/revocation/supersession records.

Required fields:

- `authority_state_revision`;
- governing authority record references;
- last material change attribution/time;
- restoration/revalidation marker when applicable.

### 7.9 `AuthoritySnapshot`

Canonical owner: **Authority Module** as a projection.

A snapshot is derived from `AuthorityState` at an evaluation time. It MUST record:

- source `authority_state_revision`;
- evaluation time;
- effective statuses of relevant authorities;
- provenance/validity information sufficient to explain the result.

Effective authority status vocabulary:

```text
VALID
EXPIRED
REVOKED
SUPERSEDED
REVALIDATION_REQUIRED
INVALID
```

`EXPIRED` is derived from validity/time; it need not mutate the canonical grant record merely because time passed.

### 7.10 `NextSafeActionProjection`

Canonical owner: **Authority Module**.

It is derived from:

```text
current ProjectStateRevision
+ current AuthorityState
+ evaluation time
→ NextSafeActionProjection
```

Required fields:

- `project_id`;
- source `state_revision`;
- source `authority_state_revision`;
- evaluated `ActionDescriptor` or explicit absence;
- decision: `PERMITTED`, `BLOCKED`, `NONE` or `REVALIDATION_REQUIRED`;
- reason classifications;
- evaluation time.

It MUST NOT be used as an Authority Grant. A cached projection is non-authoritative when either source revision or the relevant validity time changes.

### 7.11 `AuditRecord`

Canonical owner: **Audit/Evidence Module**.

Audit records explain material operations but do not own current state.

Required fields:

- stable audit record identity;
- operation/attempt identity;
- Aurora/Project correlation references;
- actor/source attribution;
- recorded time/sequence;
- classification;
- outcome/reason classification;
- state/authority revision references where material.

### 7.12 `EvidenceRecord`

Canonical owner: **Audit/Evidence Module**.

Required fields follow R2 evidence semantics:

- evidence identity;
- criterion/reference supported;
- producer;
- verifier;
- method;
- environment;
- relevant versions/revisions;
- artifact/integrity references;
- known uncertainty and limitations.

Claim, Receipt, Evidence, Verdict and Outcome remain distinct concepts even if M0 stores them physically together later.

### 7.13 `ExportManifest`

Canonical owner: **Portability/Recovery Module**.

Required semantic fields:

- export identity;
- export/package contract version;
- CAP-SOVEREIGN-CORE semantic version;
- Aurora identity;
- Project identities/current state revisions included;
- authority-state revision and authority records required for continuity;
- required audit/evidence references;
- source schema/logical versions;
- creation time/source environment reference;
- integrity descriptor/reference;
- known limitations.

The manifest format and cryptographic/integrity mechanism remain R4 decisions.

### 7.14 `RecoveryResult`, `RestoreResult` and `MigrationResult`

Canonical owner: **Portability/Recovery Module**.

Each structured result identifies:

- operation identity;
- source/target logical versions where applicable;
- affected Aurora/Project identities;
- resulting state revision or no-change result;
- authority validation/revalidation status;
- integrity/compatibility result;
- failure classification when applicable;
- evidence references and limitations.

---

## 8. Logical architecture and ownership boundaries

```text
Operator Adapter
      │ commands / inspection
      ▼
Sovereign Core Coordinator
      ├── Identity Module
      ├── Project State Module
      ├── Authority Module
      ├── Portability / Recovery Module
      └── Audit / Evidence Module
              │
              ├── Durable State Port
              ├── Evidence/Artifact Port
              ├── Time Source Port
              └── Integrity Port
```

### 8.1 Sovereign Core Coordinator

Owns cross-module orchestration and operation completion semantics. It does not own a second copy of module state.

Responsibilities:

- bind one material operation to stable correlation/attempt identity;
- coordinate validation before canonical mutation;
- require all state-affecting invariants to succeed before reporting completion;
- coordinate audit/evidence references;
- route inspection responses through projections without letting the adapter become authority.

### 8.2 Identity Module

Owns:

- `AuroraIdentity`;
- `OperatorIdentityRef`;
- stable identity invariants and collision checks.

Does not own Project state, authority policy, authentication technology or persistence technology.

### 8.3 Project State Module

Owns:

- `Project`;
- `ProjectStateRevision`;
- `TransitionAttempt`;
- candidate `ActionDescriptor`;
- current accepted state pointer/revision rules.

Does not own authority decisions or audit truth.

### 8.4 Authority Module

Owns:

- `AuthorityGrantRecord`;
- `AuthorityState`;
- `AuthoritySnapshot`;
- `NextSafeActionProjection`;
- M0 internal permission evaluation semantics.

Does not execute external effects and does not broker credentials.

### 8.5 Portability / Recovery Module

Owns:

- recovery orchestration;
- `ExportManifest` semantics;
- restore validation/application sequence;
- migration compatibility sequence;
- structured recovery/restore/migration results.

It consumes canonical state from owners; it does not become the owner of a second canonical copy.

### 8.6 Audit / Evidence Module

Owns:

- `AuditRecord`;
- `EvidenceRecord` metadata;
- required proof linkage and integrity references.

Audit, logs and telemetry MUST NOT become the sole canonical Project/authority source.

### 8.7 Operator Adapter

A replaceable input/output boundary for commands and inspection. It may be CLI, web or another later mechanism. It MUST NOT own canonical state, identity or authority.

### 8.8 Durable State Port

Logical persistence boundary for canonical M0 state. The port contract requires:

- persistence across process termination;
- revision/precondition-safe writes;
- recovery of a mutually consistent accepted view;
- explicit failure when required state cannot be validated;
- support for logical export/restore/migration semantics.

R3 does not select storage product, transaction technology, event sourcing or snapshot strategy.

### 8.9 Evidence/Artifact Port

Logical boundary for evidence/artifact payloads or references. M0 may physically co-locate these later, but the logical roles remain distinct.

### 8.10 Time Source Port

Provides the time context used for authority validity and evidence timestamps. R3 requires explicit time-source semantics because expired authority cannot be evaluated safely without time. R4 must determine the concrete local time/rollback-resistance approach appropriate to M0.

### 8.11 Integrity Port

Provides creation and validation of integrity descriptors for state/export evidence. The concrete checksum/signature/key mechanism remains R4 scope.

---

## 9. Semantic command contracts

The following are language-neutral command families. Names are semantic, not wire/API commitments.

### 9.1 `InitializeAurora`

Inputs:

- authenticated `OperatorIdentityRef` context;
- initialization request identity;
- semantic contract version.

Preconditions:

- no Aurora identity already governs the target sovereign store;
- operator identity satisfies the owner-bootstrap boundary.

On success:

- create stable `AuroraIdentity`;
- create initial owner authority state for the M0 scope;
- emit audit/evidence references;
- return identity and authority-state revision.

Repeated initialization against an already initialized identity MUST NOT silently replace the identity.

### 9.2 `CreateProject`

Inputs:

- request identity;
- `aurora_id`;
- actor/operator context;
- display label and objective summary.

On success, creates one stable `Project` with no accepted operational-state revision until the first state transition is accepted.

### 9.3 `TransitionProjectState`

Inputs:

- request/attempt identity;
- `project_id`;
- actor identity;
- `expected_state_revision` or explicit `NO_CURRENT_REVISION` for the first accepted state;
- new state envelope;
- optional proposed `ActionDescriptor`.

Validation order is semantically:

1. Project/identity validity;
2. state-envelope/contract validity;
3. expected current revision/precondition;
4. current authority evaluation for state mutation;
5. immutable-identity and ownership invariants;
6. commitability/containment preconditions.

Success creates one new immutable revision and current pointer. Any failed check rejects without changing current accepted state.

### 9.4 `ChangeAuthority`

Inputs:

- request identity;
- authenticated owner operator context;
- expected current authority-state revision;
- grant/revoke/supersede operation;
- target `AuthorityGrantRecord` semantics.

Only owner-authorized M0 authority administration may create/revoke/supersede authority. Stale revision or invalid scope fails closed.

### 9.5 `InspectCurrentProject`

Returns:

- Aurora/Project identities;
- current accepted Project state/revision or explicit absence;
- current authority snapshot provenance/status;
- `NextSafeActionProjection`;
- material limitations/degraded state;
- references to relevant audit/evidence.

Inspection cannot mutate state.

### 9.6 `RecoverCurrentState`

Used on fresh-process startup.

Semantics:

1. load canonical identity/Project/state/authority records;
2. validate required structural/integrity/version invariants;
3. evaluate current authority using current time;
4. reconstruct projections;
5. return `RECOVERED`, `DEGRADED`, `BLOCKED` or `FAILED` with classification.

Recovery MUST NOT fabricate missing accepted state or authority.

### 9.7 `ExportState`

Produces a logically consistent export package and `ExportManifest` for the requested M0 scope plus proof result. It MUST NOT report success until required contents and integrity descriptor are produced and validated.

### 9.8 `RestoreState`

Restore phases:

```text
RECEIVED
→ STRUCTURAL_VALIDATION
→ VERSION_COMPATIBILITY
→ INTEGRITY_VALIDATION
→ IDENTITY_COLLISION_CHECK
→ AUTHORITY_SAFETY_EVALUATION
→ SAFE_TO_APPLY
→ APPLIED | REJECTED | FAILED
```

Canonical state MUST NOT change before `SAFE_TO_APPLY` succeeds.

#### Restore authority safety rule

An export package may be older than later revocations. Therefore a restored authority record MUST NOT automatically become effective merely because the package was valid when created.

After restore:

- `REVOKED`, `SUPERSEDED`, `EXPIRED` or invalid authority remains non-permitting;
- an apparently active authority whose freshness relative to later revocation cannot be proven becomes `REVALIDATION_REQUIRED`;
- `NextSafeActionProjection` MUST be `REVALIDATION_REQUIRED` or `BLOCKED` for actions requiring that authority;
- only a new explicit owner revalidation/grant operation, or another R4-approved freshness proof satisfying this Spec, may return it to `VALID`.

This rule applies to restore, not ordinary restart from the current canonical store.

### 9.9 `MigrateState`

Migration is an explicit source-version to target-version operation. It MUST preserve identity, current-state meaning, authority semantics, provenance and evidence references or fail without silently changing governing semantics.

---

## 10. Lifecycle and state model

### 10.1 Aurora lifecycle

```text
ABSENT
→ INITIALIZED
```

M0 defines no autonomous replacement/deletion lifecycle.

### 10.2 Project lifecycle

```text
ABSENT
→ CREATED
```

M0 does not add archive/delete project lifecycle. Operational evolution occurs through immutable `ProjectStateRevision` records.

### 10.3 Project state revision lifecycle

```text
no current revision
→ revision 1 accepted
→ revision N accepted
→ revision N+1 accepted
```

Only one revision is current for a Project. Superseded revisions remain historical and non-governing.

### 10.4 Transition attempt lifecycle

```text
RECEIVED
→ VALIDATING
→ ACCEPTED | REJECTED | FAILED
```

A rejected/failed attempt cannot become current state.

### 10.5 Authority lifecycle

Canonical grant lifecycle:

```text
ACTIVE
→ REVOKED | SUPERSEDED
```

Effective evaluation can additionally yield:

```text
VALID | EXPIRED | REVALIDATION_REQUIRED | INVALID
```

No effective status is permission without matching action/scope/conditions.

### 10.6 Recovery lifecycle

```text
STARTING
→ VALIDATING
→ RECOVERED | DEGRADED | BLOCKED | FAILED
```

`DEGRADED` may allow safe inspection while denying mutation. `BLOCKED` denies operations whose authority/state cannot be trusted. `FAILED` means required M0 state could not be reconstructed.

### 10.7 Export/restore/migration lifecycle

Each operation has an attempt identity and ends in an explicit success/failure result. Unknown/ambiguous outcomes MUST NOT be silently retried as mutations.

---

## 11. Context and memory boundary

M0 canonical operational state is **not memory**.

The Core MUST NOT require:

- transcript replay;
- LLM context reconstruction;
- vector retrieval;
- observational memory;
- Context Builder;
- Harness local history.

Future M1 memory may reference Project/state owners but MUST NOT become their owner. Narrative or memory content cannot override current accepted state or authority.

---

## 12. Authority and effects boundary

M0 authority exists only to govern the M0 internal state/portability operations and safe-next-action projection needed by the walking skeleton.

M0 explicitly does **not** implement external effects. Therefore:

- `ActionDescriptor` is descriptive;
- `NextSafeActionProjection` is a permission projection for the next internal/operator-visible action;
- no external credential/effect execution follows from it;
- no AHDK/Provider may mint authority;
- technical access remains distinct from authority.

Future effect-plane capabilities must extend authority semantics through later gates without weakening M0's owner/revision/expiry/revocation rules.

---

## 13. Security, privacy and data classification

Detailed threat analysis is owned by `THREAT-MODEL.md`.

### 13.1 M0 data-family classification

The accepted Aurora classification vocabulary is used as follows for R3:

| Data family | R3 classification | Rationale |
|---|---|---|
| Aurora identity metadata | `INTERNAL` | stable product identity, not intended for unrestricted disclosure |
| operator identity references/authentication provenance | `SENSITIVE` | misuse can affect authority attribution |
| Project operational state/objective label | `CONFIDENTIAL` | may contain private project/business context |
| authority grants/state/snapshots | `SENSITIVE` | disclosure/tampering can affect permitted operations |
| audit records | `CONFIDENTIAL` | may reveal project state, actions and security outcomes |
| evidence metadata/references | `CONFIDENTIAL` minimum | actual evidence payload may require a higher inherited class |
| telemetry/correlation identifiers | `INTERNAL` | must exclude sensitive payloads |
| export/backup package | `SENSITIVE` minimum | aggregate contains operational and authority state; inherits any higher contained classification |
| integrity descriptor | `INTERNAL` minimum | may inherit package classification depending on mechanism |

M0 does not require secret values in canonical state, prompts, manifests or general telemetry. If a future mechanism introduces secret material, that is a new security/architecture input and must not be smuggled into M0 state.

### 13.2 Sovereignty

Canonical identity, Project state, authority, audit and export/backup material remain under Leandro-governed infrastructure/administration. External model/provider availability cannot be required for current-state authority or recovery.

---

## 14. Failure and recovery model

### 14.1 Failure classes

R3 defines these M0 failure classes:

| Class | Examples | Required containment |
|---|---|---|
| `INVALID_TRANSITION` | stale revision, malformed envelope, identity mutation | reject; no canonical mutation |
| `AUTHORITY_INVALID` | missing, expired, revoked, scope mismatch | fail closed; block mutation/next action |
| `STATE_UNAVAILABLE` | canonical store unavailable | no fabricated current state; inspection may be degraded only if non-authoritative |
| `STATE_CORRUPT` | failed structural/integrity validation | block governing use; preserve evidence |
| `VERSION_INCOMPATIBLE` | unsupported persisted/export version | explicit migration required or fail |
| `RESTORE_UNSAFE` | identity collision, stale authority, failed integrity | do not apply restored state |
| `MIGRATION_FAILED` | semantic/identity preservation cannot be proven | retain prior governing state or explicit no-change failure |
| `OPERATION_AMBIGUOUS` | process failure around material mutation with unproven result | reconcile before retry; never claim success |
| `INTERNAL_OPERATIONAL` | unexpected internal failure | classify, preserve evidence, no unsafe success claim |

### 14.2 Retry rule

Retry is permitted only when failure is classified retryable and the operation is safe/idempotent for the affected state. State mutation, restore and migration are never blindly retried after an ambiguous result.

### 14.3 Cache/projection rule

Caches and projections are reconstructable/non-authoritative. After restart, source revision mismatch or authority-validity change, they must be recomputed or treated stale.

---

## 15. Observability, audit and evidence

### 15.1 Required correlation identities

M0 observability must carry, where applicable:

- `aurora_id`;
- `project_id`;
- operation/attempt ID;
- state revision;
- authority-state revision;
- proof-run ID;
- export/restore/migration attempt ID.

Sensitive content must not be placed in general telemetry payloads merely for correlation.

### 15.2 Audit minimum

Audit must distinguish:

- attempted/accepted/rejected transitions;
- authority changes/evaluation result classes where material;
- restart/recovery boundary/result;
- export attempt/result;
- restore attempt/result;
- migration attempt/result;
- material failure classification.

### 15.3 Evidence minimum

Evidence supports a criterion but is not itself the verdict. M0 evidence must preserve method, producer/verifier, environment/revisions, integrity/artifact references and limitations.

### 15.4 Telemetry purpose rule

Every retained M0 metric/signal must identify the proof criterion, threshold, operational decision or recovery action it informs. Telemetry, logs and audit cannot become the sole canonical state or verdict source.

---

## 16. Evaluation and evidence model

R3 evaluation focuses on deterministic M0 behavior rather than model quality.

Evaluation dimensions:

- continuity;
- state correctness;
- authority integrity;
- negative-path rejection;
- recovery behavior;
- export/restore/migration correctness;
- sovereignty/security;
- evidence completeness;
- implementation/operational efficiency once R4/R7 provide measurable mechanisms.

The Capability Test Plan defines representative journeys, adversarial/fault cases, thresholds and requirement mappings.

---

## 17. Compatibility and migration

### 17.1 Version identities

At minimum, M0 must distinguish:

- CAP-SOVEREIGN-CORE semantic/contract version;
- persisted-state logical schema version;
- export package version;
- migration source/target version.

R4 chooses concrete representation.

### 17.2 Compatibility rule

A reader/restorer may accept a version only when compatibility is explicit. Unknown or incompatible versions cannot be silently coerced.

### 17.3 Migration rule

Migration must be explicit, version-pair scoped and evidence-producing. Stable identities and authority semantics are protected invariants.

### 17.4 Adapter evolution

Later process/binding changes must adapt to the domain rather than rewrite Aurora/Project/authority meaning. M0 does not implement future distributed topology.

---

## 18. Rollout and graduation

R3 defines logical graduation levels, not deployment technology.

### G0 — Spec complete

- all 122 R2 requirements allocated to Spec mechanisms and test plan;
- threat model complete;
- R4 decisions clearly enumerated;
- no unresolved R3 semantic contradiction.

### G1 — Architecture ready

Future R4 condition: implementation-blocking technical choices/spikes resolved sufficiently for one scoped Mission Contract.

### G2 — Contract/design ready

Future R5/R6 condition: exact Mission Contract and Microdesign approved.

### G3 — Implemented/evidenced

Future R7 condition: implementation evidence exists for the approved contract.

### G4 — Product milestone accepted

Future R8 condition: complete M0 Golden Proof passes against a fixed accepted revision and receives required operator verdict.

R3 does not promote directly beyond G0.

---

## 19. Open R4 decisions and uncertainty classes

These questions are **not R3 architecture placeholders**; R3 fixes the semantics they must satisfy and deliberately leaves their mechanisms to R4.

| ID | R4 decision/uncertainty | R3 constraints |
|---|---|---|
| `R4-Q-CORE-001` | Core implementation language/runtime | must not own domain semantics; smallest local fit preferred |
| `R4-Q-STORE-001` | operational-state storage mechanism | process-independent durability, revision/precondition safety, restore/migration support |
| `R4-Q-STATE-001` | state-versus-event persistence pattern | current state remains canonical; events/audit distinct and not sole authority |
| `R4-Q-SCHEMA-001` | schema/serialization representation | semantic versions/invariants remain implementation-neutral |
| `R4-Q-ATOMIC-001` | crash-consistent commit mechanism | accepted state/audit/evidence references cannot be reported partially successful |
| `R4-Q-INTEGRITY-001` | integrity mechanism | detect material corruption for state/export before governing use |
| `R4-Q-TIME-001` | time/rollback semantics for expiry | expired authority must not become valid because time is stale/rolled back |
| `R4-Q-AUTHN-001` | local operator authentication/bootstrap mechanism | technical access cannot equal owner authority |
| `R4-Q-EXPORT-001` | export/backup format/topology | portable, classified, integrity-verifiable, Leandro-controlled |
| `R4-Q-MIGRATE-001` | migration mechanism/tooling | explicit version-pair migration preserving protected semantics |
| `R4-Q-AUDIT-001` | event/audit physical mechanism | audit explainability without becoming current-state owner |
| `R4-Q-TELEM-001` | telemetry backend/transport | IDs without sensitive payload; evidence independent of backend |
| `R4-Q-TOPOLOGY-001` | process/deployment topology | logical modularity first; no distribution without evidence |
| `R4-Q-ENGINE-001` | durable execution engine applicability | do not introduce unless M0 requirements prove proportionate need |
| `R4-Q-RESTORE-001` | mechanism for authority freshness/revalidation after restore | default fail-closed `REVALIDATION_REQUIRED` unless freshness can be proven safely |

No candidate technology or winner is selected here. Architecture Spike execution remains prohibited until separately authorized.

---

## 20. Requirement coverage

Detailed requirement-to-Spec/test allocation is owned by:

```text
docs/capabilities/CAP-SOVEREIGN-CORE/R3-COVERAGE.md
```

R3 cannot pass unless all 122 R2 requirement IDs are allocated to at least one current Spec mechanism/section and at least one test-plan case or explicit documentation-review verification case.

---

## 21. Owner and reviewer

Capability owner:

```text
developmentconexus-ops
```

Required acceptance authority for normative promotion:

```text
operator
```

R3 gate verdict is evidence that the proposed package satisfies readiness conditions; it does not convert `status: proposed` to operator-accepted content merely because files are merged or CI is green.

---

## 22. Stop boundary

```text
R3 Capability Spec + Threat Model + Test Plan + Coverage
→ adversarial review
→ R3 PASS | FAIL | BLOCKED
→ STOP
→ R4 only after separate explicit operator authorization
```
