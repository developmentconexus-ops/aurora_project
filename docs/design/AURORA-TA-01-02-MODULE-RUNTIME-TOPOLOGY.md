---
id: DESIGN-AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY
title: Aurora TA-01/TA-02 Module Ownership and Runtime Topology
document_type: system_architecture_design
form: reference
authority: design
status: proposed
version: 0.5.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - proposed minimum coherent Aurora technical component set
  - proposed canonical module and entity ownership for TA-01
  - proposed Stage A process and runtime topology for TA-02
  - proposed Stage B evolutionary topology and process split triggers
  - proposed transport-neutral provider runtime boundary profile
  - proposed TA-01 and TA-02 decision disposition register
related:
  - DOC-AURORA-STATUS
  - DESIGN-AURORA-TECHNICAL-ARCHITECTURE-BASELINE-MAP
  - PLAN-AURORA-TECHNICAL-ARCHITECTURE-BASELINE
  - DESIGN-AURORA-SYSTEM-ARCHITECTURE-DECISION-LANDSCAPE
  - REVIEW-AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY-2026-08-12
  - DOC-AURORA-BLUEPRINT-03
  - DOC-AURORA-BLUEPRINT-05
  - DOC-AURORA-BLUEPRINT-06
  - DOC-AURORA-BLUEPRINT-07
  - DOC-AURORA-BLUEPRINT-10
  - DOC-AURORA-BLUEPRINT-11
  - DOC-AURORA-BLUEPRINT-12
  - DOC-AURORA-BLUEPRINT-13
  - ADR-AURORA-0001
  - ADR-AURORA-0002
  - ADR-AURORA-0003
  - ADR-AURORA-0004
  - ADR-AURORA-0009
  - DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC
  - DESIGN-AURORA-M0-R6-SOVEREIGN-CORE-MICRODESIGN
source_revision: 9cbce1efe4f742f90623b894c0c1ba2eaa3cebcc
observed_r7_candidate_revision: 7ec999b093205a9d82eef2802eca60330d96e14d
review_triggers:
  - an accepted TA-01 or TA-02 finding changes canonical ownership
  - a first M1 or M2 consumer proves the proposed runtime seam invalid
  - a process boundary cannot satisfy recovery, security or compatibility requirements
  - a component acquires a second canonical owner
  - a provider or framework type leaks into Aurora-owned domain contracts
  - Stage A cannot remain operationally small under the proposed topology
last_reviewed: 2026-08-12
---

# Aurora TA-01/TA-02 — Module Ownership and Runtime Topology

## 1. Purpose and authority boundary

This document proposes the first cross-system technical answer to:

```text
TA-01
What is the minimum coherent set of Aurora technical components,
and who owns each canonical responsibility and datum?

TA-02
Which responsibilities share a process in Stage A,
which cross an independent runtime boundary,
and how does the topology evolve to Stage B?
```

It does not select repository strategy, wire protocol, universal database, authentication product, policy engine, service supervisor, model provider, Voice provider or observability backend.

The design remains **proposed** until operator acceptance. It does not authorize implementation, TA-03 finalization, Architecture Spike execution, M0 R7 continuation or M0 R8.

---

## 2. Fixed source baseline

```text
repository: developmentconexus-ops/aurora_project
canonical branch: main
source revision: 9cbce1efe4f742f90623b894c0c1ba2eaa3cebcc

frozen R7 evidence branch: feat/m0-r7-sovereign-core-20260810
observed revision: 7ec999b093205a9d82eef2802eca60330d96e14d
classification: FROZEN / PRESERVED / NON-CANONICAL
```

The frozen candidate demonstrates that this M0 pattern is executable:

```text
composition/adapters
→ application use cases
→ domain owners + ports
→ mechanism adapters
```

That pattern is evidence. The M0 source tree, Go runtime, SQLite store and process shape are not universalized by implication.

### 2.1 Accepted cross-system constraints

1. Aurora is one persistent sovereign cognitive/operational control plane.
2. One durable concept has one canonical owner.
3. Identity, Project/Mission meaning, authority, governed context, provider trust and global Outcome remain Aurora-owned.
4. Harnesses own specialized local methodology, attempts, workers and provider-local state.
5. Models, runtimes, protocols, SDKs and stores are mechanisms, not product ontology.
6. Aurora owns canonical Contract Model semantics; bindings remain replaceable.
7. AHDK is a first-party Golden Path, not the specification or security boundary.
8. Operational state, memory, knowledge, exact history, evidence, audit and telemetry remain distinct.
9. Material effects are deterministically enforced outside model judgment.
10. Logical modularity precedes physical distribution.
11. Processes are disposable; canonical state survives them.
12. Stage A is one Leandro-controlled workstation with a persistent minimum and on-demand cognition.
13. Mastra is preferred-first to evaluate for agentic/cognitive runtimes, never as sovereign state or authority owner.
14. Detailed Presence/session policy is deferred unless it changes a structural or near-term implementation decision.

---

## 3. Architectural vocabulary

| Kind | Meaning | Canonical Aurora state ownership |
|---|---|---:|
| **Governed architecture owner** | owns versioned cross-system semantics and compatibility policy | semantic authority, not operational business state |
| **Domain owner** | owns one durable Aurora concept and its invariants | yes, only for named scope |
| **Application coordinator** | orchestrates use cases across owners/mechanisms | no duplicate truth |
| **Deployable/composition boundary** | packages modules, adapters and lifecycle | none merely because it hosts them |
| **Enforcement boundary** | deterministically allows, denies or executes an effect | execution state only; canonical receipts recorded by Aurora owner |
| **Provider runtime** | performs replaceable cognitive/specialized execution | provider-local state only |
| **Mechanism adapter** | storage, protocol, model, OS, telemetry or UI mechanism | no product meaning |

### 3.1 G01 — Contract Model Governance

G01 is the named, non-deployable canonical owner of Aurora cross-system contract semantics.

It owns semantic contract families/versions, compatibility/deprecation/removal policy, canonical field/behavior meaning, projection/generation authority, conformance-profile criteria, provider-manifest mapping and semantic/schema/binding/SDK/provider version separation.

It does not own provider trust/approval, evidence records, AHDK implementation, dispatch, transport or provider-local state.

```text
proposed semantic revision
→ review
→ accepted semantic version
→ schema/binding projections
→ conformance criteria/evidence
→ C05 compatibility/approval
→ governed deprecation/removal
```

G01 is realized through accepted Specs/ADRs and later Contract Model artifacts, not a network service.

### 3.2 Aurora Core and Sovereign Host

```text
Aurora Core
→ governed set of Aurora-owned domain modules and application coordinators

Aurora Sovereign Host
→ Stage A deployable/composition boundary hosting the current Core subset
```

The Sovereign Host is not one God domain module. Hosting does not merge ownership.

### 3.3 Provider runtime

A provider runtime is independently restartable, replaceable and limited to provider-local execution. First-party trust does not move it into the sovereign-state boundary.

---

# 4. TA-01 — Minimum coherent component model

## 4.1 Overview

```text
┌───────────────────────────────────────────────────────────────┐
│                GOVERNED CROSS-SYSTEM SEMANTICS                │
│ G01 Contract Model Governance                                 │
└───────────────────────────────┬───────────────────────────────┘
                                │ projections / conformance
┌───────────────────────────────▼───────────────────────────────┐
│                    AURORA-OWNED DOMAIN                        │
│ C01 Identity & Relationship                                   │
│ C02 Project, World & Experiment State                         │
│ C03 Intent, Mission & Delegation Control                      │
│ C04 Authority & Policy                                        │
│ C05 Capability & Provider Registry                            │
│ C06 Memory, Knowledge & Context                               │
│ C07 Artifact, Observation & Evidence                          │
│ C08 Presence & Interaction                                    │
│ C09 Attention & Proactivity                 [future/deferred]  │
│ C10 Failure Intelligence & Evaluation       [future/deferred]  │
│ C11 Environment & Device Registry           [future/deferred]  │
│ C12 Audit & Exact History                                     │
└───────────────────────────────┬───────────────────────────────┘
                                │
                    application coordination
                                │
        ┌───────────────────────┼────────────────────────┐
        │                       │                        │
 Cognitive Coordination   Capability Fabric     Effect Coordination
        │                       │                        │
        ▼                       ▼                        ▼
 Cognitive Runtime       Harness / Provider       PDP / Gateway /
 model providers         runtimes                 Credential boundary
```

Not every logical owner becomes a current code module or service. Future/deferred owners prevent missing or duplicate ownership when their consumer arrives.

---

## 4.2 Canonical domain-owner catalog

### C01 — Identity and Relationship

Owns root AuroraIdentity, Person/owner identity, relationship/profile references, protected identity settings and cross-domain actor-reference vocabulary.

```text
Provider / Harness / ProviderInstance → C05
Presence / InteractionSession         → C08
Device / Environment                  → C11
Mission / Delegation actor linkage    → C03 references ActorChain
```

Does not own OS login, biometrics, OAuth/OIDC, device authentication mechanism, model personality state or Authority Grants.

**Stage A:** active Core owner.

---

### C02 — Project, World and Experiment State

Owns Project/current revisions, Project-owned Goal/Objective, Hypothesis, general engineering Experiment lifecycle, globally meaningful Experiment Run references/status, bounded world relationships and references to decisions/repositories/devices/knowledge/work.

Does not own accepted ADR/source contents, Mission lifecycle, governed Memory, immutable Observation/Measurement, live telemetry, provider-local steps or Authority Grants.

**Stage A:** active Core owner, extending M0 without a universal graph.

---

### C03 — Intent, Mission and Delegation Control

**Purpose:** own interpreted intent once it becomes decision-relevant, then govern explicit promotion into Mission/Delegation work without treating raw utterance or model inference as committed intent.

Owns:

- canonical `Intent` identity, interpreted meaning, provenance and confidence once validated;
- the relationship from an Interaction Session to one or more committed Intents;
- Intent lifecycle and explicit promotion/rejection/expiry decisions;
- Mission, Delegation and parent/child relationships;
- global Run/Attempt references, lifecycle, dependencies, stop/escalation and pending decisions;
- global budgets/reconciled consumption;
- final Mission/Delegation Outcome after evidence/verdict composition.

Intent lifecycle:

```text
CANDIDATE
→ VALIDATED
→ COMMITTED_TO_INTERACTION | PROMOTED_TO_MISSION |
  REJECTED | EXPIRED
```

Raw audio/text and exact interaction history remain C08/C12 records. A02 or P01 may propose an Intent candidate; only C03 validates and commits the interpreted Intent. Creating a Mission requires an explicit C03 transition and does not follow automatically from model classification.

Does not own Harness plan/workers/retries, Context Pack, provider trust, effect permission, artifact contents or durable-engine history.

**Stage A:** Intent ownership is required by M1 interaction; full Mission/Delegation implementation grows with the first consuming milestone.

---

### C04 — Authority and Policy

Owns Authority Grants/current state, actor/delegation/resource/action authority, policy inputs, canonical Policy Decisions, guardrail/revocation/revalidation, Effect Request identity/permission and authority projections.

Does not own authentication sessions, raw secrets, external execution, target truth, Effect Receipt custody or framework-local permissions.

**Stage A:** M0 minimum active; broader effects grow with consumer.

---

### C05 — Capability and Provider Registry

Owns Capability Definition identity/version, Provider/Harness/ProviderInstance identity/lifecycle, manifests, compatibility assessments against G01, conformance/verification references, trust/approval and suspension/revocation/retirement.

Does not own G01 semantics, Delegation lifecycle, provider internals, Authority Grants, credentials or self-declared trust.

**Stage A:** owner fixed; implementation deferred to M2.

---

### C06 — Memory, Knowledge and Context

Owns governed Memory Item scope/provenance/epistemic/lifecycle, candidate/promotion/supersession/retention/deletion, Knowledge Source references, lower-authority synthesis, Context Builder policy, Context Pack provenance and index rebuild/delete relationships.

Does not own Project state, accepted decisions, active grants, exact provider state, live truth, retrieval indexes as truth or L4 exact interaction/tool/event history.

**Stage A:** first substantive consumer M1; heavy mechanisms may be workers/providers.

---

### C07 — Artifact, Observation and Evidence

Owns Artifact identity/integrity/content reference, immutable Observation/Measurement with provenance/quality/units, Claims, Receipts, Evidence relationships, authorized Verdict records and evidence retention/provenance.

Does not own source-code authority, Experiment lifecycle, Mission Outcome, target truth, raw telemetry, general Audit Records or exact interaction/tool/event history.

**Stage A:** minimum M0 records exist; broad capability grows with M2/M3/experiments.

---

### C08 — Presence and Interaction

Owns Presence identity/registration/revocation/status/capabilities, Interaction Session lifecycle, Activation Request semantics, channel/privacy/environment/handoff context and delivery references.

Does not own Person identity, Project/Mission, Authority, governed memory, provider state, OS/device authentication or durable exact-history custody.

**Stage A:** thin local adapter required; full Fabric deferred.

---

### C09 — Attention and Proactivity

Owns proactive candidates, urgency/relevance/confidence, attention budgets, deduplication and delivery decisions. Does not own Project truth, arbitrary Mission creation, authority or Presence transport.

**Stage A:** DEFER; first consumer M5.

---

### C10 — Failure Intelligence and Evaluation

Owns Incidents, Findings/causal relationships, evaluation definitions/datasets/result relationships, evaluation-specific correlations, Improvement Opportunities/Candidates and regression/promotion/rollback proposals.

Does not own production promotion, holdout changes, raw telemetry, C02 Experiment lifecycle, C07 evidence or C12 audit/history.

**Stage A:** DEFER; first consumers M6/M7.

---

### C11 — Environment and Device Registry

Owns Environment/Device identity/lifecycle, inventory/relationships, registration/pairing/revocation, controller/firmware/calibration/manifest references and expected capability/protocol/trust metadata.

Does not own live telemetry, readings, target truth, actuation, interlocks, firmware source or controller-local state.

**Stage A:** DEFER; first substantial consumers M8/M9/M10.

---

### C12 — Audit and Exact History

**Purpose:** preserve the attributable, append-oriented historical record required for accountability, reconstruction and L4 exact history without conflating it with telemetry, memory, artifacts or current domain state.

Owns:

- canonical Audit Record identity and append semantics;
- actor/subject/executor/Presence/Mission/Delegation/correlation references for audited actions;
- exact L4 interaction, tool invocation, provider-boundary and material domain-event history;
- sequence/time/source/integrity metadata needed for reconstruction;
- audit-history retention, access, export and redaction-policy references;
- commit/append path for governance-required audit facts produced by owners, B01, E01 and C08.

Does not own:

- current Project/Mission/Authority state;
- governed Memory Items or synthesis;
- Artifact/Receipt/Evidence/Verdict semantics;
- raw telemetry or provider traces;
- transport delivery state;
- external target truth;
- domain-event meaning, which remains with the emitting owner.

Relationship:

```text
owner commits domain fact
→ C12 appends attributable exact-history/audit record when profile requires
→ C06 may later derive governed memory candidate
→ C07 may cite the audit record as evidence only through explicit promotion
```

M0 may physically store Audit/Evidence together. TA-01 keeps C07 and C12 as distinct logical owners because their authority, retention and query semantics differ.

**Stage A:** active foundational owner; may share process/store adapters with current Core while retaining separate interfaces and data classification.

---

## 4.3 Application and integration components

### A01 — Core Application Coordinator

Coordinates commands/queries, owner validation, unit-of-work boundaries, required C12/C07 records and post-commit notification. Owns sequencing, not truth.

### A02 — Cognitive Coordination

```text
interaction / Mission need
→ obtain C02/C04 inputs
→ C06 builds Context Pack
→ C05 resolves provider against G01
→ A05 ensures an approved runtime incarnation is READY
→ A03 invokes through B01
→ receive Intent/proposal/observation/artifact/capability request
→ route Intent candidate to C03 and other outputs to their owners
```

A02 owns reasoning-use-case sequencing only. It neither commits interpreted Intent nor turns a provider thread into Aurora truth.

### A03 — Capability Fabric / Harness Integration

Owns dispatch/translation/reconciliation mechanics: connection, Delegation dispatch, status ingestion, cancellation, health, child requests, artifact refs and binding translation. It dispatches only to an A05-READY runtime and does not own G01, C05 approval, C03 lifecycle, A05 desired runtime state or provider-local execution state.

### A04 — Durable Execution Port

Port for timers, waits, checkpoints and durable execution. Engine history remains mechanism state reconciled with C03/C12.

### A05 — Runtime Lifecycle Coordination

**Purpose:** own Aurora-side desired lifecycle and recovery decisions for separately running providers without selecting or becoming an operating-system supervisor product.

Owns:

- desired runtime state for each approved provider role;
- start, attach, readiness, drain, stop, restart and reconcile decisions;
- one new `runtime_incarnation_id` for every concrete process/runtime start;
- coordination of D01/D02/D03 shutdown ordering;
- restart policy decisions constrained by current C03/C04/C05 state;
- runtime health interpretation and transition to READY, DEGRADED, FAILED or UNKNOWN;
- C12 lifecycle/audit append requests for material runtime transitions.

Lifecycle:

```text
ABSENT
→ STARTING
→ READY
→ DRAINING
→ STOPPED

STARTING | READY | DRAINING
→ DEGRADED | FAILED | UNKNOWN
```

Startup sequence:

```text
current demand + C05-approved provider instance
→ A05 selects approved runtime configuration
→ supervisor/process adapter starts or attaches
→ new runtime_incarnation_id
→ G01/B01 compatibility and readiness handshake
→ READY or explicit failure
```

Shutdown/restart sequence:

```text
A05 blocks new dispatch
→ requests B01 cancel/drain
→ reconciles C03/C07/C12 state
→ supervisor adapter stops the process
→ restart, replacement or terminal failure decision
→ every restart creates a new runtime_incarnation_id
```

A05 does not own provider identity/trust/approval (C05), interpreted Intent/Mission/Attempt state (C03), provider checkpoints, G01 semantics, evidence, telemetry or the OS supervisor mechanism. TA-08 chooses supervisor adapters and operational products; it cannot displace A05 as lifecycle-policy owner.

---

## 4.4 Enforcement components

### E01 — Effect Gateway family

```text
Effect Request + Policy Decision
→ minimum credential
→ execute / deny / preserve ambiguity
→ produce receipt
→ C07 records receipt
→ C12 records attributable exact history/audit
```

Gateway does not create authority.

### E02 — Credential Broker boundary

Resolves secret references into minimum short-lived credentials without exposing raw secrets to models, broad environments or logs. Product/process belongs to TA-06.

---

## 4.5 Provider runtime classes

### P01 — Cognitive Runtime Provider

Owns provider-local model/agent loops, thread/tool history, routing mechanics, workflow snapshots, skills/workspaces and local traces/evals. It does not own global identity/state/authority/budgets/trust/memory/evidence/outcome/effect permission.

Mastra/TypeScript remains preferred-first to evaluate.

### P02 — Specialized Harness Provider

Owns domain methodology, plans, workers, attempts, tools and internal recovery within Delegation. MNFS is one future provider, not Core.

### P03 — Model / External Service Adapter

Wraps local/external model/service behind policy, minimization and attribution. Vendor session format stays external.

### P04 — Device / Laboratory Provider

Future controller/Harness behind C11, C04, E01 and independent interlocks.

---

## 4.6 Mechanism adapters

Stores, indexes, blobs, protocol bindings, UI/Voice adapters, supervisors, OTel exporters, model SDKs, sandboxes and backup/migration mechanisms remain adapters. Co-location never grants ownership.

---

# 5. Canonical ownership matrix

| Entity/data family | Canonical owner | May propose/produce | Cannot commit canonical change |
|---|---|---|---|
| Contract semantics/version/deprecation/conformance | G01 | Specs/ADRs/review | C05/A03/AHDK/provider/generator |
| Aurora/Person/relationship identity | C01 | approved identity process | model/Presence/provider/store |
| Provider/Harness/ProviderInstance | C05 | provider manifest | provider claim alone |
| Presence/InteractionSession | C08 | Presence adapter | cannot broaden actor/authority |
| Device/Environment inventory | C11 | registration adapter | live telemetry/provider alone |
| Project/current Project state | C02 | proposals through A01 | Memory/model/Harness/UI/DB client |
| Interpreted Intent and promotion decision | C03 | C08/A02/P01 candidate | model/Presence/provider cannot commit or auto-create Mission |
| Hypothesis/Experiment lifecycle | C02 | operator/cognitive/Harness | provider workflow alone |
| Observation/Measurement | C07 | instrument/provider/Harness | producer cannot rewrite provenance |
| Mission/Delegation/global Attempt | C03 | operator/cognitive/child request | provider/Harness directly |
| Global budget | C03 | provider/gateway receipts | provider counter alone |
| Authority Grant/revocation | C04 | approved policy process | model/Harness/Presence/gateway |
| Effect Request/Policy Decision | C04 | Mission/provider/Presence | gateway cannot self-authorize |
| Capability/Manifest/Provider Approval | C05 | provider + G01/C07 evidence | provider self-approval |
| Governed Memory Item | C06 | candidate producers | model/provider/index directly |
| Context Pack | C06 | owner inputs | provider cannot broaden |
| Provider local workflow/thread | P01/P02 | provider | never canonical by convenience |
| Artifact/Observation/Evidence | C07 | authorized producers/verifiers | producer cannot self-promote |
| Effect Receipt | C07 | E01 produces | model/provider cannot fabricate |
| Verdict record | C07 | authorized verifier | producer alone unless policy |
| Mission Outcome | C03 | C07 verdict package | child completion alone |
| Audit Record | C12 | owners/B01/E01/C08 facts | producer cannot rewrite committed history |
| L4 exact interaction/tool/event history | C12 | C08/B01/domain owners | C06/provider/telemetry cannot replace |
| Proactive candidate/attention budget | C09 | module nomination | provider cannot self-interrupt |
| Incident/ImprovementCandidate | C10 | modules/evidence processors | candidate cannot self-promote |
| Raw telemetry | source/telemetry mechanism | components | not state/verdict/history automatically |
| Source-code bytes and commit history | Git repository | governed development workflow | Aurora runtime/provider cannot redefine source |
| Accepted ADR/Spec/Contract document custody and commit history | Git/document repository | governed authoring and promotion workflow | custody does not own G01 Contract semantics/version/deprecation/conformance |
| Live device state | device/telemetry source | controller/sensor | registry/memory cannot override |

## 5.1 Ownership verbs

```text
READ     query/consume
PROPOSE  candidate/request
VALIDATE owner checks invariants
COMMIT   owner creates canonical record
PROJECT  non-authoritative view
APPEND   C12 creates immutable attributable history
```

READ/PROPOSE does not imply COMMIT/APPEND.

---

# 6. Dependency and mutation rules

## 6.1 Allowed direction

```text
Presence/external adapters
→ application coordinators
→ domain-owner public interfaces + ports
→ mechanism adapters
```

Provider side:

```text
Harness/cognitive domain
→ AHDK/public boundary
→ G01 binding
→ A03
```

Cross-owner use cases are coordinated by A01; no owner mutates another store directly.

## 6.2 Minimal shared kernel

Only typed refs, revision/time values, data classification, actor/correlation refs and true cross-domain result categories. No entities/framework/infrastructure helpers.

## 6.3 Forbidden dependencies

```text
domain → DB/ORM/filesystem mechanism
domain → Mastra/model/provider type
domain → transport/UI/OTel type
domain → provider-local store
provider/Harness/Presence → canonical DB
memory index → Project/Authority writer
telemetry → canonical mutation or exact-history authority
Effect Gateway → Authority Grant creation
generator → G01 semantic authority
```

## 6.4 Canonical mutation/audit path

```text
request/proposal
→ A01
→ owner validation
→ C04 policy when required
→ owner COMMIT
→ C12 APPEND when audit/history profile requires
→ C07 evidence/receipt record when proof is required
→ post-commit notification/projection
```

Provider output is an Intent candidate, proposal, observation, artifact, claim, capability request, provider-local status or reference to an E01-controlled receipt—never an implicit canonical mutation.

## 6.5 Cross-owner consistency

Local unit of work may coordinate Project+audit, Mission+budget+authority refs or provider approval+verification refs. External effects use explicit request/decision/execution/receipt/reconciliation, never one assumed transaction.

If a governance-required C12 append cannot be guaranteed, the corresponding material mutation/effect must fail closed or remain uncommitted according to the owning future profile. TA-05/TA-08 will define physical atomicity/outbox/recovery mechanisms.

## 6.6 Event/history rule

```text
Domain Event       owner-committed fact
Transport Message fallible delivery
Telemetry          diagnostic signal
C12 Exact History attributable durable record
C07 Evidence       promoted proof relationship
```

A transport event or telemetry sample is not automatically C12 history or C07 evidence. Missing delivery cannot erase canonical state; missing required audit history blocks/flags the governing operation per profile.

---

# 7. TA-02 process-boundary criteria

Split only for current evidence: different runtime, independent lifecycle, fault containment, privilege, heavy resources, device locality, remote node, measured scaling, upgrade cadence or consumer/Spike. Otherwise remain logical in smallest deployable.

---

# 8. Three Stage A approaches

## A — Core-centric single application

Simplest install/transactions; weak framework/resource/privilege/fault isolation. Not recommended beyond deterministic M0-like slices.

## B — Early services

Strong isolation; premature RPC, identity, discovery, supervision and distributed consistency. Rejected for Stage A.

## C — Evolutionary Sovereign Host with one early provider seam

```text
Stage A workstation
├── Aurora Sovereign Host — persistent
│   ├── current domain owners including C03 Intent and C12 history
│   ├── A01 application coordination
│   ├── A05 runtime lifecycle coordination
│   ├── canonical state/recovery
│   ├── thin Presence adapter
│   └── B01 provider endpoint
├── Cognitive Runtime — separate/on demand at first consumer
├── Specialized Harness providers
└── model/store/effect adapters
```

Early seam is justified by polyglot runtime, framework cadence, resources, untrusted content/tool requests, independent restart and provider-local state. Other modules stay co-located until named trigger.

**Recommendation:** Approach C.

---

# 9. B01 — Transport-neutral Provider Runtime Boundary Profile

B01 defines process-crossing semantics without selecting transport/schema.

## 9.1 Identity/compatibility envelope

Provider ID/instance/build/environment, `runtime_incarnation_id`, capability version, G01 semantic version, schema/binding version, permitted Aurora refs and C05 approval snapshot. Reject incompatibility before execution. `provider_instance_id` identifies an approved exact build/environment registration; `runtime_incarnation_id` identifies one concrete process start under A05.

## 9.2 Correlation/idempotency

```text
request_id      logical operation
attempt_id      concrete execution attempt
correlation_id  diagnostics
idempotency_key duplicate protection
parent_delegation_ref purpose/lifecycle
```

Retry keeps logical request/idempotency identity and creates new attempt. Duplicate accepted submission must not silently start duplicate execution.

## 9.3 Authority/context attenuation

Minimum actor/delegation refs, action/resource classes, budget/deadline, data/provider policy, Context Pack and credential refs. Provider cannot widen/forward/mint authority; child request returns to Aurora.

## 9.4 Lifecycle

```text
CREATED → SUBMITTED → ACCEPTED | REJECTED
ACCEPTED → RUNNING → SUCCEEDED | FAILED | CANCELED |
                     TIMED_OUT | RECONCILIATION_REQUIRED
```

SUCCEEDED is provider execution completion, not global commit/effect proof. Cancellation requires acknowledgement/snapshot. Progress is fallible. Material transitions append C12 history according to profile.

## 9.5 Deadline/cancel/retry

Explicit deadline/no-deadline; idempotent cancellation; known stopped/completed/unknown; no blind retry after unknown completion; new Attempt for retry; stale authority/deadline/incompatibility blocks resume; effects reconcile independently.

## 9.6 Restart/snapshot

Snapshot by request/attempt returns provider build, lifecycle, last sequence/checkpoint, artifact/claim/receipt refs, pending child/effect requests, error class and resumable/terminal/lost/unknown state.

Aurora validates and reconciles C03/C07/C12. Missing provider state never becomes inferred success.

## 9.7 Response/error categories

```text
ACCEPTED PROGRESS OUTPUT_PROPOSAL ARTIFACT_REFERENCE CLAIM_REFERENCE
CAPABILITY_REQUEST EFFECT_REQUEST TERMINAL_RESULT

INVALID_REQUEST INCOMPATIBLE_CONTRACT UNAPPROVED_PROVIDER_INSTANCE
UNAUTHORIZED_OR_STALE_CONTEXT DEADLINE_EXCEEDED CANCELED
PROVIDER_UNAVAILABLE EXECUTION_FAILED STATE_LOST RECONCILIATION_REQUIRED
```

Material responses include request/attempt/provider/contract identities and are eligible for C12 append under profile.

## 9.8 TA-04 handoff

TA-04 defines encoding, transport, streaming, channel auth, wire errors, negotiation, backpressure and health without weakening B01 lifecycle/idempotency/authority/reconciliation.

---

# 10. Stage A deployables

## D01 — Aurora Sovereign Host

M0 Go executable is practical seed, not universal Go:

```text
M0 Go fact
→ Stage A Go Host seed hypothesis
→ M1+ placement/language requires consuming revalidation
```

Persistent: bootstrap/identity, state/recovery, current Project/authority, C03 Intent ownership, A01 coordination, A05 runtime lifecycle coordination, C12 append path, minimal interaction, thin Presence, B01 endpoint and fail-closed behavior.

Not always active: LLM/Mastra/STT/TTS/heavy indexing/Harness/eval/device control.

## D02 — Cognitive Runtime

A05 starts or attaches D02 only after current demand and C05 approval, creates a new runtime incarnation, requires the G01/B01 readiness handshake and exposes it to A03 only in READY state. Core owners validate every output.

## D03 — Harness Provider

Independent lifecycle through G01/B01 Capability/Delegation contracts.

## D04 — Model/media/index workers

Replaceable; first consumer decides embedding/separation.

## Conditional enforcement

E01/E02 split when credentials/material effects/privilege/containment demand. No generic privileged executor.

---

# 11. Always-active versus on-demand

| Responsibility | Stage A disposition |
|---|---|
| identity/bootstrap/recovery | always D01 |
| Project/authority state | always available |
| C12 audit/exact-history append path | always available for governed operations |
| controlled application/B01 endpoint | always |
| A05 runtime lifecycle coordination | always available; provider processes remain on demand |
| local activation adapter | configured |
| full UI | possible on demand |
| Cognitive Runtime/Mastra | on demand |
| models | on demand |
| Context Builder | on demand; governance owned |
| memory extraction/indexing | M1 background/on demand |
| Harnesses | independent/on demand |
| STT/TTS | on demand |
| evidence metadata/history query | available; payload/query mechanism as needed |
| eval workers | deferred/on demand |
| durable engine | deferred |
| effect gateways | per accepted family |
| telemetry backend | replaceable; never truth/history owner |

---

# 12. Runtime/language scope

| Scope | Posture | Boundary |
|---|---|---|
| G01 | language-neutral semantics | projections/version separation |
| M0 Core | Go accepted M0 | no provider types |
| Stage A Host | Go seed; M1+ expansion revalidated | no universal Go mandate |
| Cognitive Runtime | separate; Mastra/TS preferred-first | B01/provider-local |
| AHDK | language open | projects G01 |
| Harness | capability-specific | G01/B01/authority |
| Presence/UI | runtime open | no canonical DB |
| model/media | mechanism-specific | governed/replaceable |
| device/embedded | device-specific | gateways/interlocks |

---

# 13. Failure/restart ownership

| Failure | Response/invariant |
|---|---|
| Host crash | recover state/C12 continuity and A05 desired runtime state; reattach or reconcile provider incarnations |
| Cognitive crash | A05 marks the incarnation FAILED/UNKNOWN, blocks dispatch and drives B01 snapshot/reconciliation; Core remains |
| provider state loss | STATE_LOST/RECONCILIATION_REQUIRED |
| Presence crash | session expires; Core continues |
| Harness crash | A05/C03 reconcile, then resume/new Attempt/block; no duplicate effect |
| canonical store failure | block mutation; trusted recovery |
| audit/history store/path unavailable | governed material operation blocks or records explicit integrity failure per future profile |
| artifact unavailable | evidence closeout blocked |
| gateway unknown receipt | ambiguity + target reconciliation |
| credential unavailable | deny credential effects |
| model outage | fallback/unavailable; identity/state remain |
| telemetry outage | diagnostics degrade; C12/C07 separate |
| Stage B partition | no inferred terminal state; restrict/reconcile |

---

# 14. Stage B evolution

A persistent personal node hosts the Sovereign Host, canonical data, C12 history/audit, evidence coordination, approved providers and risk gateways. Workstation/mobile become Presence clients.

Stable Aurora domain identities and G01/C01–C12 ownership do not move. Environment-bound runtime identities do change:

```text
provider_id
→ may remain stable for the same provider product/version lineage

provider_instance_id
→ MUST be newly registered for the new build/environment placement
→ requires a fresh C05 compatibility/trust/approval snapshot

runtime_incarnation_id
→ MUST be new for every process start, including post-migration start
```

Provider-local thread/checkpoint/workflow state is reconciled through B01 and may be resumed only when compatible and authorized. It is never imported into C02/C03/C04/C06/C12 as canonical truth merely because files were copied.

Stage B adds service identity, Presence authentication, transport, egress, offline behavior, supervision, minimization, history/evidence retention and backup questions to TA-04/05/06/08. A05 remains the Aurora-side lifecycle-policy owner while TA-08 supplies node-specific supervisor adapters.

---

# 15. Split triggers

| Component | Stage A posture | Trigger |
|---|---|---|
| G01 | non-deployable | never service for ownership alone |
| C01/C02/C03/C04/C12 | Host | proven security/scale/topology/storage lifecycle reason |
| A05 lifecycle coordination | Host | remains Aurora-side owner; TA-08 may replace only the supervisor/process adapter |
| C05 | Host when implemented | remote/independent trust lifecycle |
| C06 governance | Host hypothesis | M1 runtime/resource/sovereignty evidence |
| memory index worker | on-demand | M1 eval |
| C07 metadata | Host | blob scale only |
| C08 | co-packaged thin | Stage B/second device/OS lifecycle |
| C11 | Host when implemented | remote device inventory need |
| P01 | separate | intentional B01 seam |
| E01/E02 | per-risk | credentials/material effects |
| C10 workers | deferred | sustained eval workload |
| telemetry backend | external adapter | never canonical owner |

---

# 16. Adversarial flows

## Model-assisted Project continuation

C08 captures interaction → A02 obtains C02/C04 inputs → C06 builds context → C05/G01 validate provider → A05 ensures READY runtime → A03/B01/P01 returns an Intent candidate/proposal → C03 validates and commits/rejects Intent, explicitly promoting to Mission only when warranted → other owners commit their state → C12 exact history → C07 evidence when required. P01 never writes Intent, Project or Mission state.

## Harness child request

P02/B01 → C03 → C05/G01 → C04 narrower authority → C06 minimized context → C03 child Delegation → C12 history. No transitive token/context.

## Memory conflicts with ADR

C06 candidate/provenance; accepted source remains authority; C12 preserves exact interaction/history; no Decision/Project mutation.

## Engineering experiment

C02 Hypothesis/Experiment → provider run → C07 Observation/Measurement/Evidence → C12 exact action/history → C02 explicit transition → C10 later correlation.

## Material effect

B01 proposal → C03 purpose/budget → C04 decision → E01/E02 → C07 receipt → C12 audit/history → reconciled C02/C03 transition.

## Cognitive death

B01 snapshot/reconcile; C12 last known events support reconstruction; sovereign truth survives.

## Stage migration

Export/restore/migrate canonical state/history under policy, start D01 on the new node, create new environment-bound provider_instance_id registrations and C05 approval snapshots, create new runtime_incarnation_id values through A05, register workstation Presence, reconcile rather than canonize provider-local state, and preserve stable Aurora domain identities/owners.

---

# 17. Approach comparison

| Criterion | A | B | C |
|---|---:|---:|---:|
| Stage A simplicity | high | low | high |
| ownership clarity | medium | high | high |
| cognitive isolation | low | high | high |
| polyglot fit | low | high | high |
| local transaction/audit coordination | high | low | high inside Host |
| fault containment | low-medium | high | high at B01 |
| premature distribution | low initially | very high | low |
| Stage B evolution | extraction-heavy | direct | direct |
| YAGNI | medium | low | high |
| God-process risk | high | low | medium with fitness rules |
| service-sprawl risk | low | high | low |

---

# 18. Decision disposition register

| ID | Question | Hypothesis | Disposition |
|---|---|---|---|
| TA12-D01 | owner set | G01 + C01–C12 | DECIDE |
| TA12-D02 | interpreted Intent owner | C03 before explicit Mission promotion | DECIDE |
| TA12-D03 | God module | prohibited | DECIDE |
| TA12-D04 | Stage A topology | Approach C | DECIDE |
| TA12-D05 | provider seam | separate at first consumer | DECIDE |
| TA12-D06 | B01 lifecycle profile | as defined | DECIDE; TA-04 implements |
| TA12-D07 | runtime lifecycle owner | A05 in Sovereign Host | DECIDE; TA-08 supplies adapters/products |
| TA12-D08 | exact binding | open | RESEARCH TA-04/SPIKE if needed |
| TA12-D09 | Go scope | M0 seed; M1+ open | RESEARCH/ADR |
| TA12-D10 | Presence process | co-packaged | DEFER split |
| TA12-D11 | Memory placement | owned/Host hypothesis | DECIDE ownership; M1 mechanisms |
| TA12-D12 | Device owner | C11 | DECIDE ownership |
| TA12-D13 | Audit/exact-history owner | C12 | DECIDE ownership; TA-05/08 mechanisms |
| TA12-D14 | Effect Gateway | risk-specific | DEFER TA-06 |
| TA12-D15 | Credential Broker | open | DEFER TA-06 |
| TA12-D16 | durable engine | none | DEFER M4 |
| TA12-D17 | Registry service | Host initially | DEFER M2 |
| TA12-D18 | stores | open | DEFER TA-05 |
| TA12-D19 | repository | blocked | DEFER TA-03 |
| TA12-D20 | schema/codegen | open | RESEARCH TA-03/04 under G01 |
| TA12-D21 | supervisor/install product | open; cannot own lifecycle policy | DEFER TA-08 |
| TA12-D22 | immediate Spike | none | DEFER until unprovable claim |

---

# 19. Inputs to later tranches

## TA-03

Enforce G01/C01–C12 and A01–A05 dependency boundaries; D01 contains multiple owners; P01/P02 are separately versioned; Presence is logically separate; G01 source stays separate from generated code; providers cannot import canonical internals; builds distinguish Core/contracts/providers/AHDK/Harness/adapters; compare repository strategies; Development Harness remains non-runtime tooling.

## TA-04

Implement B01 including `runtime_incarnation_id`, G01 version separation and A05 readiness/cancellation/reconciliation integration; choose bindings per boundary; define channel auth, encoding, errors, streaming, backpressure and negotiation; preserve attenuation; define which provider/domain events become C12 history.

## TA-05

Define physical roles, retention, export and recovery for canonical state, C12 audit/history, C07 evidence/artifacts, C06 memory and derived indexes without merging logical ownership. Define physical atomicity/outbox/recovery between owner COMMIT and required C12 APPEND.

## TA-08

Provide supervisor/process adapters used by A05; define startup/shutdown hooks, OS integration, update/rollback, health probes, C12 availability/integrity profile, telemetry separation and diagnostics. TA-08 may choose mechanisms but cannot move runtime lifecycle-policy ownership out of A05.

---

# 20. Review remediation history

Internal v0.1→v0.2 resolved Device owner, identity overlap, experiment/observation allocation, Go scope and implicit Contract Model.

External v0.2→v0.3 resolved named G01 governance and B01 lifecycle/reconciliation profile.

Codex v0.3→v0.4 resolved missing Audit Record and L4 exact-history ownership through C12, mutation/audit paths, failure behavior and later-tranche handoff.

Final reviewer remediation v0.4→v0.5 assigned interpreted Intent to C03, runtime lifecycle policy to A05, separated G01 semantics from Git custody, restored E01-only receipt production, required provider re-registration across Stage B migration and corrected heading/wording findings.

Worklog/final tracking continuity remains before operator-promotion readiness.

---

# 21. Explicit non-decisions

No choice of monorepo/polyrepo/package layout; universal Go; concrete Mastra; transport; schema generator; stores; auth/policy/secrets products; supervisor implementation/containers/Kubernetes; model/Voice/sandbox/durable engine/observability backend; device protocol; AHDK language; implementation plan/code. A05 owns lifecycle policy, not the supervisor product.

---

# 22. Acceptance criteria and operator gate

Confirm:

1. representative concepts have one owner;
2. G01 semantics are distinct from Git/document custody and cannot be captured by C05/A03/AHDK/provider;
3. C03 owns validated interpreted Intent before explicit Mission promotion;
4. C12 distinctly owns audit/L4 exact history;
5. providers cannot write state/audit/evidence or produce Effect Receipts outside owner/gateway paths;
6. B01 defines process lifecycle without transport selection;
7. A05 owns start/attach/readiness/drain/stop/restart/reconciliation policy without selecting a supervisor product;
8. Sovereign Host is not a God module;
9. Stage A remains small;
10. cognitive failure cannot erase truth/history;
11. no service-per-module pattern;
12. Stage B preserves stable domain identity while re-registering environment-bound provider instances/incarnations;
13. M0 choices are not globalized;
14. Mastra remains replaceable;
15. unresolved mechanisms have owner/treatment;
16. TA-03/04/05/08 receive constraints without preselection.

```text
ACCEPT
→ accept Approach C, G01/C01–C12, A05 and B01
→ no implementation or later-tranche finalization automatically

REVISE
→ exact changes

REJECT
→ replace proposal under accepted Technical Architecture Map
```

Until acceptance:

Until acceptance:

```text
proposal NON-CANONICAL
TA-03/04/05/08 finalization BLOCKED
Spikes NOT AUTHORIZED
implementation PAUSED
```
