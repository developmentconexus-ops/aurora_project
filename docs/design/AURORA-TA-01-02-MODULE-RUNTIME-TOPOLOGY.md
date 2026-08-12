---
id: DESIGN-AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY
title: Aurora TA-01/TA-02 Module Ownership and Runtime Topology
document_type: system_architecture_design
form: reference
authority: design
status: proposed
version: 0.3.0
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
8. Operational state, memory, knowledge, history, evidence, audit and telemetry remain distinct.
9. Material effects are deterministically enforced outside model judgment.
10. Logical modularity precedes physical distribution.
11. Processes are disposable; canonical state survives them.
12. Stage A is one Leandro-controlled workstation with a persistent minimum and on-demand cognition.
13. Mastra is preferred-first to evaluate for agentic/cognitive runtimes, never as sovereign state or authority owner.
14. Detailed Presence/session policy is deferred unless it changes a structural or near implementation decision.

---

## 3. Architectural vocabulary

| Kind | Meaning | Canonical Aurora state ownership |
|---|---|---:|
| **Governed architecture owner** | owns versioned cross-system semantics and compatibility policy | owns semantic authority, not operational business state |
| **Domain owner** | owns one durable Aurora concept and its invariants | yes, only for named scope |
| **Application coordinator** | orchestrates use cases across owners/mechanisms | no duplicate truth |
| **Deployable/composition boundary** | packages modules, adapters and lifecycle | none merely because it hosts them |
| **Enforcement boundary** | deterministically allows, denies or executes an effect | execution state only; canonical receipts recorded by Aurora owner |
| **Provider runtime** | performs replaceable cognitive/specialized execution | provider-local state only |
| **Mechanism adapter** | storage, protocol, model, OS, telemetry or UI mechanism | no product meaning |

### 3.1 G01 — Contract Model Governance

**G01 is the named canonical owner of Aurora cross-system contract semantics.** It is a governed, non-deployable architecture/specification authority in Stage A.

G01 owns:

- semantic contract-family definitions;
- semantic contract identity and version lifecycle;
- compatibility, deprecation and removal policy;
- canonical field/behavior meaning independent of representation;
- which schema projections and generated bindings are authoritative derivatives;
- binding-generation source authority and reproducibility requirements;
- conformance-profile definitions and expected evidence criteria;
- mapping rules between Aurora contract versions and provider manifest declarations;
- rules separating semantic, schema, binding, SDK and provider versions.

G01 does **not** own:

- provider identity, trust or approval—C05 owns those;
- evidence records—C07 records evidence against G01 criteria;
- AHDK implementation—AHDK projects G01 semantics;
- runtime dispatch—A03 performs it;
- protocol transport or generated-code implementation;
- provider-local workflow state.

Lifecycle:

```text
proposed semantic revision
→ specification/adversarial review
→ accepted semantic version
→ schema/binding projections
→ conformance profiles and evidence
→ C05 provider compatibility/approval assessment
→ deprecation/removal under accepted policy
```

G01 is currently realized through accepted Aurora Specs/ADRs and later Contract Model artifacts. It must not be turned into a network service merely because it is a canonical owner.

### 3.2 Aurora Core and Sovereign Host

```text
Aurora Core
→ governed set of Aurora-owned domain modules and application coordinators

Aurora Sovereign Host
→ Stage A deployable/composition boundary hosting the current Core subset
```

The Sovereign Host is not one God domain module. Hosting does not merge ownership.

### 3.3 Provider runtime

A provider runtime is independently restartable, replaceable and limited to provider-local execution. A first-party provider can be trusted for a scope without becoming part of the sovereign-state boundary.

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
│ C03 Mission & Delegation Control                              │
│ C04 Authority & Policy                                        │
│ C05 Capability & Provider Registry                            │
│ C06 Memory, Knowledge & Context                               │
│ C07 Artifact, Observation & Evidence                          │
│ C08 Presence & Interaction                                    │
│ C09 Attention & Proactivity                 [future/deferred]  │
│ C10 Failure Intelligence & Evaluation       [future/deferred]  │
│ C11 Environment & Device Registry           [future/deferred]  │
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

Owns root `AuroraIdentity`, `Person`/owner identity, relationship/profile references, protected identity settings and cross-domain actor-reference vocabulary.

Entity-specific identity owners remain:

```text
Provider / Harness / ProviderInstance → C05
Presence / InteractionSession         → C08
Device / Environment                  → C11
Mission / Delegation actor linkage    → C03 references ActorChain
```

Does not own OS login, biometrics, OAuth/OIDC sessions, device authentication mechanism, model personality state or Authority Grants.

**Stage A:** active Core owner.

---

### C02 — Project, World and Experiment State

Owns `Project`, Project revisions/current pointer, Project-owned Goal/Objective, `Hypothesis`, general engineering `Experiment` lifecycle, globally meaningful Experiment Run references/status, bounded world relationships and references to decisions, repositories, devices, knowledge and active work.

Does not own accepted ADR/source contents, Mission/Delegation lifecycle, governed Memory Items, immutable Observation/Measurement records, live telemetry, provider-local execution steps or Authority Grants.

**Stage A:** active Core owner, extending M0 without prebuilding a universal graph.

---

### C03 — Mission and Delegation Control

Owns Mission, Delegation/parent-child relationships, global Run/Attempt references, lifecycle/dependencies/stop/escalation, pending decisions, global budget allocation/reconciled consumption and final Mission/Delegation Outcome after authorized evidence/verdict composition.

Does not own Harness plan/worker graph/retries, Context Pack construction, provider trust, effect permission, artifact contents or durable-engine history.

**Stage A:** owner fixed now; implementation deferred until a current consumer. M1 may use C08 interaction state without prebuilding the full Mission engine.

---

### C04 — Authority and Policy

Owns Authority Grants/current state, actor/delegation/resource/action authority relationships, policy inputs, canonical Policy Decisions, guardrail/revocation/revalidation semantics, Effect Request identity/permission state and authority projections.

Does not own authentication sessions, raw secrets, external execution, target truth, Effect Receipt custody or framework-local permission state.

**Stage A:** M0 minimum active; broader effect semantics grow with first effectful consumer.

---

### C05 — Capability and Provider Registry

Owns Capability Definition identity/version, Provider/Harness/ProviderInstance identity/lifecycle, manifests, provider compatibility assessments against G01, conformance/verification references, multidimensional trust/approval and provider suspension/revocation/retirement.

Does not own contract meaning, active Delegation lifecycle, provider internals, Authority Grants, credentials or self-declared trust as accepted truth.

**Stage A:** owner fixed; implementation deferred to M2.

---

### C06 — Memory, Knowledge and Context

Owns governed Memory Item identity/scope/provenance/epistemic status/lifecycle, candidate/promotion/supersession/retention/deletion propagation, Knowledge Source index/reference metadata, lower-authority synthesis, Context Builder policy, Context Pack identity/provenance and derived-index rebuild/delete relationships.

Does not own Project state, accepted decisions, active grants, exact provider thread/workflow state, live device/provider truth or retrieval index as canonical truth.

**Stage A:** first substantive consumer M1. Governance is Aurora-owned; heavy extraction/indexing may be on-demand workers/providers.

---

### C07 — Artifact, Observation and Evidence

Owns Artifact identity/metadata/integrity/content reference, immutable Observation and Measurement records with provenance/quality/units, Claims, canonical Receipts, Evidence relationships, authorized Verdict records/limitations and retention/provenance metadata.

Does not own Git/source authority, Experiment lifecycle, Mission lifecycle/final Outcome, external target truth, raw telemetry by default or authority to accept its own producer claim.

**Stage A:** minimum M0 records exist; broad capability grows with M2/M3 and experiments.

---

### C08 — Presence and Interaction

Owns Presence identity/registration/revocation/status/channel capabilities, Interaction Session identity/lifecycle, Activation Request semantics, channel/privacy/environment/handoff context and delivery references.

Does not own Aurora/Person identity, Project/Mission state, Authority Grants, governed memory, provider conversation state or OS/device authentication mechanism.

**Stage A:** thin local Presence/activation adapter required; full Fabric deferred.

---

### C09 — Attention and Proactivity

Owns proactive candidates, urgency/relevance/confidence, attention budgets, deduplication and delivery decisions. Does not own Project truth, arbitrary Mission creation, user authority or Presence transport.

**Stage A:** `DEFER`; first consumer M5.

---

### C10 — Failure Intelligence and Evaluation

Owns Incidents, Findings/causal relationships, evaluation definitions/dataset references/result relationships, evaluation-specific experiment correlations, Improvement Opportunities/Candidates and regression/promotion/rollback proposals.

Does not own production promotion authority, arbitrary holdout changes, raw telemetry, C02 Experiment state or C07 evidence records.

**Stage A:** `DEFER`; first consumers M6/M7.

---

### C11 — Environment and Device Registry

Owns Environment/Device identity/lifecycle, inventory/relationships, registration/pairing/revocation, controller/firmware/calibration/manifest references and expected capability/protocol/trust metadata.

Does not own live telemetry, current instrument readings, target truth, actuation, deterministic interlocks, firmware source or controller-local state.

**Stage A:** `DEFER`; Project may reference devices before implementation. First substantial consumers M8/M9/M10.

---

## 4.3 Application and integration components

### A01 — Core Application Coordinator

Coordinates commands/queries, domain-owner validation, current unit-of-work boundaries, required audit/evidence references and post-commit notification. Owns sequencing, not duplicate state.

### A02 — Cognitive Coordination

```text
interaction / Mission need
→ obtain C02/C04 inputs
→ C06 builds Context Pack
→ C05 resolves approved cognitive capability/provider against G01
→ A03 invokes provider through B01
→ receive proposal/observation/artifact/capability request
→ validate and route to canonical owner
```

Owns no provider thread as Aurora truth and cannot commit provider output directly.

### A03 — Capability Fabric / Harness Integration

Owns dispatch/translation/reconciliation mechanics: provider connection, Delegation dispatch, status ingestion, cancellation, health, child capability requests, artifact references and binding translation. Does not own G01 semantics, Registry approval, Mission lifecycle or provider-local state.

### A04 — Durable Execution Port

Domain-oriented port for timers, waits, checkpoints and durable execution. It is not a current service or canonical state owner; engine history remains mechanism state reconciled with C03.

---

## 4.4 Enforcement components

### E01 — Effect Gateway family

Per-effect deterministic enforcement for filesystem, network, repository, communication, deployment, financial and device/lab boundaries.

```text
Effect Request + Policy Decision reference
→ minimum credential when needed
→ execute / deny / preserve ambiguity
→ produce receipt
→ C07 records canonical receipt
```

Gateway does not create authority.

### E02 — Credential Broker boundary

Resolves secret references into minimum short-lived credentials without exposing raw secrets to models, broad environments or general logs. Product/process choice belongs to TA-06 and first credential-mediated consumer.

---

## 4.5 Provider runtime classes

### P01 — Cognitive Runtime Provider

Owns provider-local agent/model loops, thread/tool history, model-routing mechanics, workflow snapshots, local skills/workspaces and provider-local traces/evals.

Must not own Aurora identity, Project/Mission state, Authority, global budgets, provider approval, governed memory authority, Artifact/Evidence identity, global Verdict/Outcome or effect permission.

Mastra/TypeScript remains preferred-first to evaluate for first real consumer.

### P02 — Specialized Harness Provider

Owns domain methodology, local plans, workers, attempts, tools and internal recovery inside an Aurora Delegation. MNFS is one future provider class, not a Core component.

### P03 — Model / External Service Adapter

Wraps local/external model or service behind provider policy, minimization and attribution. Vendor message/session formats stay outside Aurora domain.

### P04 — Device / Laboratory Provider

Future controller/Harness behind C11 registration, C04 authority, E01 gateways and independent interlocks.

---

## 4.6 Mechanism adapters

Stores, indexes, blob systems, protocol bindings, UI/Voice adapters, supervisors, OTel exporters, model SDKs, sandboxes and backup/migration mechanisms remain adapters. Physical co-location never grants ownership.

---

# 5. Canonical ownership matrix

| Entity/data family | Canonical owner | May propose/produce | Cannot commit canonical change |
|---|---|---|---|
| Contract semantics/version/deprecation/conformance criteria | G01 | Specs/ADRs/review process | C05, A03, AHDK, provider, binding generator |
| Aurora/Person/relationship identity | C01 | operator-approved identity process | model, Presence, provider, store |
| Provider/Harness/ProviderInstance identity | C05 | provider manifest | provider claim alone |
| Presence/InteractionSession identity | C08 | Presence adapter/device | Presence cannot broaden actor identity/authority |
| Device/Environment inventory identity | C11 | controller/adapter registration | live telemetry/provider alone |
| Project/current Project state | C02 | Presence, Mission, provider result through A01 | Memory, model, Harness, UI, DB client |
| Hypothesis/Experiment global lifecycle | C02 | operator/cognitive/Harness proposal | provider-local workflow alone |
| Observation/Measurement | C07 | instrument/provider/Harness | producer cannot rewrite provenance after commit |
| Mission/Delegation/global Attempt | C03 | operator/cognitive/Harness child request | provider/Harness directly |
| Global budget allocation/consumption | C03 | provider/gateway receipts | provider counter alone |
| Authority Grant/revocation | C04 | operator/approved policy process | model, Harness, Presence, gateway |
| Effect Request/Policy Decision | C04 | Mission/provider/Presence request | gateway/provider cannot self-authorize |
| Capability/Manifest/Provider Approval | C05 | provider + G01 conformance evidence | provider self-approval |
| Governed Memory Item | C06 | observer/model/provider/user candidate | model/provider/index directly |
| Context Pack | C06 | C02/C03/C04/C05 inputs | provider cannot broaden it |
| Provider thread/workflow state | P01/P02 local | provider runtime | never canonical by convenience |
| Artifact/Observation/Measurement/Evidence metadata | C07 | authorized producers/verifiers | producer cannot self-promote claim |
| Effect Receipt | C07 canonical record | E01 produces | model/provider cannot fabricate target result |
| Verdict record | C07 | authorized verifier/decider | producer alone unless policy permits |
| Mission Outcome | C03 | C07 evidence/verdict package | child completion alone |
| Proactive candidate/attention budget | C09 | any module nominates | provider cannot self-interrupt outside policy |
| Incident/ImprovementCandidate | C10 | modules/evidence processors | candidate cannot self-promote |
| Raw telemetry | source/telemetry mechanism | all components | not product state/verdict automatically |
| Source code/accepted ADR/Spec/Contract | Git/document owner | governed development process | Aurora runtime store is not source owner |
| Live device state | device/telemetry source | P04/controller/sensor | registry/memory cannot override live verification |

### 5.1 Ownership verbs

```text
READ     consume through query/read contract
PROPOSE  submit candidate/requested transition
VALIDATE owner checks invariants
COMMIT   owner creates canonical revision/record
PROJECT  non-authoritative view across sources
```

`READ` or `PROPOSE` never implies `COMMIT`.

---

# 6. Dependency and mutation rules

## 6.1 Allowed direction

```text
Presence / external adapters
          ↓
application coordinators
          ↓
domain-owner public interfaces + ports
          ↓
mechanism adapters implement ports
```

Provider side:

```text
Harness/cognitive domain
→ AHDK or direct public provider boundary
→ G01 Contract Model binding
→ A03 Capability Fabric adapter
```

Cross-owner use cases are coordinated by A01; one owner does not mutate another owner's store directly.

## 6.2 Minimal shared kernel

Only typed IDs/references, revision/time values, data classification, actor/correlation references and truly cross-domain result categories may be shared. No module entities, framework types or infrastructure helpers belong there.

## 6.3 Forbidden dependencies

```text
domain → DB driver / ORM / filesystem mechanism
domain → Mastra/model/provider type
domain → MCP/A2A/gRPC/HTTP type
domain → UI framework or OTel backend
domain → provider-local store
provider/Harness/Presence → canonical Aurora database
memory index → Project/Authority writer
telemetry → canonical mutation authority
Effect Gateway → Authority Grant creation
binding/schema generator → G01 semantic authority
```

## 6.4 Canonical mutation path

```text
request/proposal
→ A01 coordination
→ current owner validation
→ C04 authority/policy evaluation where required
→ owner COMMIT
→ required audit/evidence reference
→ post-commit notification/projection
```

Provider output is a proposal, observation, artifact, claim, capability request, provider-local status or controlled gateway receipt—never an implicit canonical write.

## 6.5 Cross-owner consistency

One local unit of work may coordinate owners when a current invariant requires atomicity: Project transition plus audit; Mission/Delegation plus budget/authority refs; provider approval plus verification refs.

External effects never assume one ACID transaction. They use explicit request/decision/execution/receipt/reconciliation states.

## 6.6 Event rule

```text
Domain Event       owner-committed fact
Transport Message fallible delivery representation
Telemetry          diagnostic signal
Audit/Receipt/Evidence accountability/proof record
```

Missing delivery cannot erase canonical state or required evidence.

---

# 7. TA-02 process-boundary criteria

A logical component crosses a process boundary only when current evidence exists for different runtime, independent lifecycle, fault containment, privilege zone, heavy resources, device locality, remote node, measured scaling, independent upgrade cadence or a current consumer/Spike.

Otherwise it stays logically modular inside the smallest safe deployable.

---

# 8. Three coherent Stage A approaches

## 8.1 A — Core-centric single application

One persistent process contains active Core modules, local Presence/UI, direct model adapters, memory mechanisms and local effects.

**Strength:** simplest install/debug/transaction model.

**Weakness:** cognition/framework/resource/privilege leakage; poor polyglot fit; coupled upgrades; weak fault containment.

**Verdict:** not recommended beyond deterministic M0-like slices.

## 8.2 B — Early service decomposition

Separate Core, Presence, Cognitive, memory/index, effect/credential and provider services.

**Strength:** explicit isolation.

**Weakness:** forces RPC, service identity, discovery, supervision and distributed consistency before Stage A consumers justify them; creates empty services.

**Verdict:** rejected for Stage A.

## 8.3 C — Evolutionary Sovereign Host with one early provider-runtime seam

```text
Stage A workstation
│
├── Aurora Sovereign Host — persistent
│   ├── current Aurora-owned domain modules
│   ├── A01 coordination
│   ├── canonical state/recovery ports
│   ├── local interaction/session coordination
│   ├── thin Presence/activation adapter initially co-packaged
│   └── B01 provider boundary endpoint
│
├── Cognitive Runtime Provider — separate/on demand at first consumer
│   ├── preferred-first Mastra/TypeScript evaluation
│   ├── model/agent/workflow execution
│   └── provider-local state only
│
├── Specialized Harness/provider processes
└── local/external model, store and effect adapters
```

**Reason for early seam:** polyglot runtime, rapid framework cadence, variable resources, untrusted/provider-generated content/tool requests, independent restart, provider replaceability and accepted local/global state separation.

**Why other boundaries stay logical:** Presence, Registry, Memory governance, Artifact/Evidence, Authority and Device inventory do not yet justify one service each.

**Strengths:** small Stage A, isolation of highest-risk framework boundary, deterministic state/authority, Mastra reuse without capture, no service-per-module pattern, Stage B path.

**Costs:** local process crossing when cognition enters; restart/reconciliation/compatibility become real work; Host needs dependency fitness rules.

**Recommendation:** Approach C.

```text
one small persistent sovereign deployable
+
one independently replaceable cognitive/provider runtime seam when consumed
+
all other physical splits require a named current trigger
```

---

# 9. B01 — Transport-neutral Provider Runtime Boundary Profile

B01 defines the minimum semantic/lifecycle behavior for any out-of-process Cognitive Runtime or Harness boundary without selecting HTTP, gRPC, local IPC, event transport or schema technology.

## 9.1 Identity and compatibility envelope

Every provider interaction identifies:

- `provider_id`;
- `provider_instance_id` tied to exact build/environment;
- capability identity/version;
- G01 semantic contract family/version;
- schema projection version;
- binding/SDK version when applicable;
- Aurora/Project/Mission/Delegation references permitted for the scope;
- provider trust/approval snapshot reference from C05.

Dispatch is rejected **before execution** when semantic or declared compatibility is unsupported. A transport connection or successful deserialization is not compatibility.

## 9.2 Correlation and idempotency identities

Every invocation uses distinct stable identities:

```text
request_id
→ one logical requested operation

attempt_id
→ one concrete provider execution attempt

correlation_id / trace_ref
→ cross-boundary diagnostic linkage

idempotency_key
→ duplicate-submission protection for the logical request

parent_delegation_ref
→ global purpose/lifecycle linkage when applicable
```

A retry keeps the same logical `request_id`/idempotency semantics and creates a new `attempt_id`. Providers must reject or return the existing result for duplicate accepted submission according to the contract; they must not start silent duplicate execution.

## 9.3 Authority and context attenuation

The request carries an **authority context reference/envelope**, never an implicit caller privilege.

It includes only the minimum needed:

- actor/delegation chain references;
- allowed action/resource classes;
- budget/deadline limits;
- data classification/provider-policy constraints;
- Context Pack reference or minimized inline data;
- credential references only when the provider is allowed to request a gateway-mediated effect.

A provider cannot widen, forward or mint authority. A child capability request returns to Aurora for a new narrower Delegation/authority decision.

## 9.4 Invocation lifecycle

```text
CREATED
→ SUBMITTED
→ ACCEPTED | REJECTED

ACCEPTED
→ RUNNING
→ SUCCEEDED | FAILED | CANCELED | TIMED_OUT |
  RECONCILIATION_REQUIRED
```

Terminal meanings:

- `SUCCEEDED` means provider execution completed and returned the declared outputs; it does not commit global Aurora state or prove an external effect.
- `FAILED` means a known terminal provider failure.
- `CANCELED` requires provider acknowledgement/snapshot; sending cancellation alone is not terminal.
- `TIMED_OUT` is Aurora's deadline classification; provider state may still need reconciliation.
- `RECONCILIATION_REQUIRED` means terminal truth is not safely known.

Progress/stream messages are informative and fallible. Canonical lifecycle decisions use provider snapshots/terminal responses plus Aurora-owned state.

## 9.5 Deadline, cancellation and retry rules

- each request carries a deadline or explicit no-deadline policy;
- Aurora may request cancellation at any time permitted by the contract;
- cancellation is idempotent;
- provider must expose whether execution stopped, completed, or remains unknown;
- no automatic retry follows unknown completion when duplicate work/effects are possible;
- known retryable failures may create a new Attempt under current authority/budget;
- stale authority, expired deadline or incompatible provider version blocks resume/retry;
- external effects follow E01 receipt/reconciliation rules independently from provider terminal status.

## 9.6 Provider restart and snapshot recovery

A provider supporting recoverable work exposes a bounded snapshot/query operation by `request_id`/`attempt_id` returning:

- provider instance/build identity that created the state;
- current lifecycle state;
- last monotonic provider sequence/checkpoint reference when applicable;
- produced artifact/claim/receipt references;
- pending child/effect requests;
- known error/reason classification;
- whether state is resumable, terminal, lost or requires reconciliation.

After provider restart:

```text
Aurora queries snapshot
→ validates provider/build/contract compatibility
→ reconcile with C03/C07 state
→ resume, new Attempt, fail, or require operator decision
```

Missing provider state never becomes inferred success. Provider snapshot is evidence about provider-local work, not global truth by itself.

## 9.7 Required response categories

Minimum responses/errors are transport-neutral:

```text
ACCEPTED
PROGRESS
OUTPUT_PROPOSAL
ARTIFACT_REFERENCE
CLAIM_REFERENCE
CAPABILITY_REQUEST
EFFECT_REQUEST
TERMINAL_RESULT

INVALID_REQUEST
INCOMPATIBLE_CONTRACT
UNAPPROVED_PROVIDER_INSTANCE
UNAUTHORIZED_OR_STALE_CONTEXT
DEADLINE_EXCEEDED
CANCELED
PROVIDER_UNAVAILABLE
EXECUTION_FAILED
STATE_LOST
RECONCILIATION_REQUIRED
```

Each material response includes request/attempt/provider/contract identities and reason classification.

## 9.8 TA-04 handoff

TA-04 must choose bindings/profiles that implement B01 and define:

- encoding/schema projections;
- local/remote transport;
- streaming mechanics;
- authentication/channel security;
- error wire representation;
- compatibility negotiation;
- backpressure/message limits;
- operational health endpoints.

TA-04 cannot weaken B01 lifecycle, idempotency, authority attenuation or reconciliation semantics for transport convenience.

---

# 10. Proposed Stage A deployables

## 10.1 D01 — Aurora Sovereign Host

The current M0 sovereign executable is Go and is the practical seed of D01.

```text
accepted current fact
→ M0 Sovereign Core runtime is Go

current topology hypothesis
→ existing Go process can seed Stage A Host

open cross-horizon language decision
→ implementing additional M1+ canonical modules in Go requires
   consuming architecture/ADR revalidation when material
```

The topology decision is the sovereign/provider process seam, not universal Go.

Persistent responsibilities: bootstrap/identity, canonical state/recovery, current Project/authority operations, A01 coordination, minimal interaction endpoint, thin Presence adapter, B01 endpoint, fail-closed behavior and current audit/evidence references.

Not continuously active: LLM sessions, Mastra, full STT/TTS, heavy indexing, Harness workers, evaluation campaigns or device control.

Providers/Harnesses never receive canonical DB credentials by convenience. Exact stores belong to TA-05.

## 10.2 D02 — Cognitive Runtime Provider

Starts when current interaction/Mission needs model-mediated work and conforms to B01. Core owners validate all outputs.

## 10.3 D03 — Specialized Harness Provider

Independent lifecycle/execution environment; integrates through G01/B01 capability/Delegation contracts.

## 10.4 D04 — Local model/media/index workers

Replaceable mechanisms/provider workers. First consumer decides embedding versus separate process.

## 10.5 Conditional enforcement deployables

E01/E02 split physically when holding credentials, exposing material effects, requiring narrower OS privilege or containing compromise. No generic privileged `execute anything` process.

---

# 11. Always-active versus on-demand

| Responsibility | Stage A disposition |
|---|---|
| identity/bootstrap/recovery | always available in D01 |
| canonical Project/authority state | always available through D01/stores |
| controlled application/B01 endpoint | always active |
| local activation adapter | available when configured |
| full UI | may start on demand |
| Cognitive Runtime/Mastra | on demand at first consumer |
| external/local models | on demand |
| Context Builder | invoked on demand; governance Aurora-owned |
| memory extraction/indexing | on demand/background after M1 design |
| Harness runtimes | independent/on demand |
| STT/TTS/Voice | on demand |
| artifact/evidence metadata | always queryable; payload mechanism as needed |
| Failure Intelligence/eval workers | deferred/on demand |
| durable engine | deferred |
| effect gateways | only per accepted effect family |
| telemetry backend/exporter | replaceable/degradable; never truth owner |

---

# 12. Runtime/language scope

| Scope | Current posture | Boundary |
|---|---|---|
| G01 Contract Model | language-neutral semantics | schema/binding/SDK versions separate |
| M0 Sovereign Core | Go accepted for M0 | no provider/framework types |
| Stage A Sovereign Host | existing Go Core is seed; M1+ Go expansion requires revalidation | topology does not mandate universal Go |
| Cognitive Runtime | separate provider; Mastra/TypeScript preferred-first | B01; provider-local state only |
| AHDK | first language open | projects G01 semantics |
| Specialized Harness | capability-specific runtime | same G01/B01/authority boundary |
| Presence/UI | runtime open | thin adapter; no canonical DB access |
| model/media services | mechanism-specific | governed, attributed, replaceable |
| device/embedded | device-specific | gateways/interlocks; outside model process |

---

# 13. Failure-domain and restart ownership

| Failure | Response/invariant |
|---|---|
| Sovereign Host crash | recover canonical state; reconcile providers; never infer truth from provider narrative |
| Cognitive Runtime crash | B01 snapshot/reconciliation; Core remains available |
| provider-local store loss | STATE_LOST/RECONCILIATION_REQUIRED; cannot erase sovereign truth |
| Presence adapter crash | session expires/revokes; Core continues |
| Harness crash | B01 resume/new Attempt/block; no blind duplicate effect |
| canonical store failure | block governing mutation; trusted recovery; no stale authority resurrection |
| artifact unavailable | evidence-dependent closeout blocked |
| gateway unknown receipt | preserve ambiguity; reconcile target; no blind retry |
| Credential Broker unavailable | deny credential effects |
| model/provider outage | approved fallback or explicit unavailable; identity/state remain |
| telemetry outage | diagnostics degrade; audit/evidence remain separate |
| Stage B partition | no inferred terminal state; restrict/reconcile remote sessions |

---

# 14. Stage B evolution

```text
persistent personal/home-lab node
├── Aurora Sovereign Host
├── canonical operational/governed-memory data
├── artifact/evidence coordination
├── approved cognitive/provider runtimes
└── risk-specific effect/credential services

workstation/mobile/future device
└── separate Presence client
```

Ownership does not move for G01 semantics, Aurora identity, Project/Mission, authority, governed memory, provider trust, Artifact/Evidence or global Outcome.

Stage B hands service identity, Presence authentication/revocation, transport/discovery, egress, offline behavior, remote supervision, minimization and backup/failover to TA-04/05/06/08.

---

# 15. Split triggers

| Component | Stage A posture | Split trigger |
|---|---|---|
| G01 | non-deployable governance owner | never a service merely for ownership |
| C01/C02/C04 | Sovereign Host | proven security/scale/topology reason |
| C03 | Host when implemented | durable/multi-node needs; engine remains adapter |
| C05 | Host when implemented | remote consumers/independent trust lifecycle |
| C06 governance | Host hypothesis | M1 runtime/resource/sovereignty evidence |
| memory index worker | on-demand candidate | M1 eval proves need |
| C07 metadata | Host | payload scale may split blob store only |
| C08 Presence | co-packaged thin adapter | Stage B, second device, OS/sensor lifecycle |
| C11 Registry | Host when implemented | device node/remote inventory consumer |
| P01 | separate at first consumer | merge not recommended; B01 is intentional |
| E01/E02 | per-risk adapter | credentials/material external/physical effects |
| C10 workers | deferred | sustained evaluation workload |
| observability backend | adapter/external | never canonical owner |

---

# 16. Adversarial flows

## 16.1 Project continuation through model

```text
C08 interaction
→ A02 obtains C02/C04 inputs
→ C06 builds Context Pack
→ C05 validates provider against G01
→ A03 invokes P01 through B01
→ P01 returns proposal/artifact/claim
→ A01 routes to owner
→ owner commits/rejects
→ C07 records evidence
```

P01 cannot write Project state or promote its thread state.

## 16.2 Harness child request

P02 requests child capability through B01; C03 validates relationship; C05 resolves compatible approved provider against G01; C04 creates narrower authority; C06 minimizes context; C03 commits child Delegation. Parent cannot pass full token/context transitively.

## 16.3 Memory conflicts with ADR

C06 records candidate/provenance; accepted ADR/source remains authoritative; no Project/Decision mutation.

## 16.4 Engineering experiment

C02 owns Hypothesis/Experiment; P02/P04 executes local method; C07 records Observation/Measurement/Artifact/Evidence; C02 transitions Experiment; C10 may correlate failures later.

## 16.5 Material effect

Provider proposes request through B01; C03 links purpose/budget; C04 evaluates; E01/E02 enforce; C07 records receipt; C02/C03 update only through reconciled transition.

## 16.6 Cognitive runtime death

B01 snapshot/reconciliation determines resume/new Attempt/failure/unknown. Core remains; sovereign truth survives.

## 16.7 Stage A to Stage B

Export/restore/migrate state, start D01 on personal node, register workstation Presence, add transport/auth boundary, preserve all identities/owners.

---

# 17. Approach comparison

| Criterion | A single application | B early services | C evolutionary Host |
|---|---:|---:|---:|
| Stage A simplicity | high | low | high |
| ownership clarity | medium | high | high |
| cognitive isolation | low | high | high |
| polyglot fit | low | high | high |
| local transaction simplicity | high | low | high inside Host |
| fault containment | low-medium | high | high at B01 seam |
| premature distribution | low initially | very high | low |
| Stage B evolution | extraction-heavy | direct | direct at accepted seams |
| YAGNI | medium | low | high |
| God-process risk | high | low | medium with fitness rules |
| service-sprawl risk | low | high | low |
| M1/M2 fit | medium | high but expensive | high |

---

# 18. Decision disposition register

| ID | Question | Hypothesis | Disposition/owner |
|---|---|---|---|
| TA12-D01 | canonical owner set | G01 + C01–C11 | `DECIDE` through design/review |
| TA12-D02 | Core as God module | prohibited | `DECIDE` |
| TA12-D03 | Stage A topology | Approach C | `DECIDE` |
| TA12-D04 | provider runtime seam | separate at first consumer | `DECIDE` structural |
| TA12-D05 | provider lifecycle profile | B01 | `DECIDE` structural; TA-04 implements |
| TA12-D06 | exact Go↔TypeScript binding | open | `RESEARCH` TA-04; Spike if consequential |
| TA12-D07 | cross-horizon Go scope | M0 seed; M1+ placement open | `RESEARCH/ADR` per consumer |
| TA12-D08 | Presence process | co-packaged initially | `DEFER` split |
| TA12-D09 | Memory placement | Aurora-owned; workers may split | ownership `DECIDE`; mechanisms M1 research |
| TA12-D10 | Device/Environment owner | C11 deferred | `DECIDE` ownership; implementation M8/M9 |
| TA12-D11 | Effect Gateway placement | risk/family specific | `DEFER` TA-06 |
| TA12-D12 | Credential Broker | open | `DEFER` TA-06 |
| TA12-D13 | durable engine | none current | `DEFER` M4 |
| TA12-D14 | Registry service split | Host initially | `DEFER` M2 evidence |
| TA12-D15 | physical stores | open | `DEFER` TA-05 |
| TA12-D16 | repository strategy | blocked | `DEFER` TA-03 |
| TA12-D17 | schema/codegen | open | `RESEARCH` TA-03/04 under G01 |
| TA12-D18 | supervisor/install | open | `DEFER` TA-08 |
| TA12-D19 | immediate Spike | none required for logical topology | `DEFER` until unprovable runtime claim |

---

# 19. Inputs to TA-03 and TA-04

## TA-03

1. domain-owner boundaries must be source-enforceable;
2. D01 contains multiple owners, not one domain package;
3. P01/P02 are independently versioned provider units;
4. Presence is logical-separate but may share Stage A packaging;
5. G01 source is language-neutral and separate from generated Go/TypeScript projections;
6. providers cannot import canonical store/internal packages;
7. builds/releases distinguish Core, contracts, cognitive provider, AHDK, Harness and adapters;
8. monorepo/polyrepo/staged extraction must be compared against these boundaries;
9. Development Harness is build/review tooling, not runtime dependency.

## TA-04

1. implement B01 identities/lifecycle/idempotency/cancellation/reconciliation;
2. preserve G01 semantic/schema/binding/SDK version separation;
3. choose bindings per boundary rather than universally;
4. define channel auth, encoding, errors, streaming/backpressure and compatibility negotiation;
5. never weaken authority attenuation or ambiguous-effect rules.

Both tranches remain separately gated.

---

# 20. Review remediation history

Internal review v0.1→v0.2 resolved:

- missing Device/Environment owner;
- overly broad Identity owner;
- missing Hypothesis/Experiment/Observation/Measurement allocation;
- ambiguous cross-horizon Go scope;
- implicit Contract Model asset.

External review v0.2→v0.3 resolved:

- `Contract Model semantics have no named canonical owner` → G01 now owns semantic/version/compatibility/deprecation/generation/conformance policy;
- `Cognitive Runtime seam lacks a binding-independent boundary profile` → B01 now defines identities, lifecycle, deadlines, cancellation, idempotency, retry, snapshot/recovery, authority attenuation and response categories.

Worklog/fixed-review continuity remains a closeout task after validation.

---

# 21. Explicit non-decisions

No selection of monorepo/polyrepo/package layout; universal Go; concrete Mastra version; HTTP/REST/gRPC/Connect/MCP/A2A/events/IPC; schema generator; new stores; Keycloak/Zitadel/Authentik/Ory/SPIFFE/OPA/Cedar/Vault; supervisor/containers/Kubernetes; model/Voice/sandbox/durable-engine/observability products; device protocol; first AHDK language; implementation roadmap or code.

---

# 22. Acceptance criteria and operator gate

Reviewer must confirm:

1. representative global concepts have one owner;
2. G01 contract semantics cannot be captured by C05/A03/AHDK/provider;
3. provider/model/Harness cannot write canonical state;
4. B01 defines every mandatory process-crossing lifecycle property without selecting transport;
5. Sovereign Host is not a God module;
6. Stage A remains small;
7. cognitive failure cannot erase sovereign truth;
8. no service-per-module pattern;
9. Stage B changes mechanisms, not identity/ownership;
10. M0 decisions are not globalized;
11. Mastra remains preferred-first/replaceable;
12. unresolved mechanisms have treatment/consumer/owner;
13. TA-03/04 receive structure without being pre-decided.

```text
ACCEPT
→ accept Approach C, G01/C01–C11 ownership and B01 profile
→ no implementation or later-tranche finalization follows automatically

REVISE
→ return exact semantic/ownership/topology/profile changes

REJECT
→ preserve accepted Technical Architecture Map
→ replace this proposal with another reviewed approach
```

Until acceptance:

```text
TA-01/TA-02 proposal: NON-CANONICAL
TA-03/TA-04 finalization: BLOCKED
Architecture Spike execution: NOT AUTHORIZED
Aurora implementation: PAUSED
```
