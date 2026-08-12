---
id: DESIGN-AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY
title: Aurora TA-01/TA-02 Module Ownership and Runtime Topology
document_type: system_architecture_design
form: reference
authority: design
status: proposed
version: 0.2.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - proposed minimum coherent Aurora technical component set
  - proposed canonical module and entity ownership for TA-01
  - proposed Stage A process and runtime topology for TA-02
  - proposed Stage B evolutionary topology and process split triggers
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
| **Governed architecture asset** | semantic/specification source projected into implementations | owns semantics, not runtime state |
| **Domain owner** | owns one durable Aurora concept and its invariants | yes, only for named scope |
| **Application coordinator** | orchestrates use cases across owners/mechanisms | no duplicate truth |
| **Deployable/composition boundary** | packages modules, adapters and lifecycle | none merely because it hosts them |
| **Enforcement boundary** | deterministically allows, denies or executes an effect | execution state only; canonical receipts recorded by Aurora owner |
| **Provider runtime** | performs replaceable cognitive/specialized execution | provider-local state only |
| **Mechanism adapter** | storage, protocol, model, OS, telemetry or UI mechanism | no product meaning |

### 3.1 G01 — Canonical Contract Model

The Canonical Contract Model is a **governed architecture/specification asset**, not a runtime service.

It is owned by Aurora Specifications and accepted decisions, then projected into:

- schemas;
- generated language types;
- bindings;
- AHDK APIs;
- conformance profiles.

It is not owned by AHDK, Capability Fabric, Mastra, an RPC framework or a provider runtime.

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
│                    AURORA-OWNED DOMAIN                        │
│                                                               │
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

Not every logical owner becomes a current code module or service. Future/deferred owners exist to prevent missing or duplicate ownership when their consumer arrives.

---

## 4.2 Canonical domain-owner catalog

### C01 — Identity and Relationship

**Purpose:** preserve root Aurora, Leandro/person and relationship identity independently from process, model, Presence or authentication mechanism.

**Owns:**

- `AuroraIdentity`;
- `Person` and owner/operator identity;
- relationship/profile references and protected identity settings;
- cross-domain actor-reference vocabulary and root actor relationships;
- constitutional identity references, not the constitutional documents themselves.

**Entity-specific identity owners remain:**

```text
Provider / Harness / ProviderInstance → C05
Presence / InteractionSession         → C08
Device / Environment                  → C11
Mission / Delegation actor linkage    → C03 references ActorChain
```

**Does not own:** OS login, biometrics, OAuth/OIDC sessions, device authentication mechanism, model personality state or Authority Grants.

**Stage A:** active Core owner.

---

### C02 — Project, World and Experiment State

**Purpose:** preserve Project identity/current meaning and governed relationships over world, engineering and authoritative-source references.

**Owns:**

- `Project` and Project state revisions/current pointer;
- Project-owned `Goal`/`Objective`;
- `Hypothesis` and `Experiment` global identity/lifecycle/relationships;
- global `ExperimentRun` references and status when materially meaningful to the Project;
- bounded world relationships and temporal validity;
- Project snapshots and next-action source references;
- references to decisions, repositories, devices, knowledge and active work.

**Does not own:** accepted ADR/source-code contents, Mission/Delegation lifecycle, governed Memory Items, immutable Observation/Measurement records, live telemetry, provider-local execution steps or Authority Grants.

**Stage A:** active Core owner, extending the M0 Project owner without prebuilding a universal graph.

---

### C03 — Mission and Delegation Control

**Purpose:** own global work intent, cross-provider decomposition, lifecycle, budgets and consolidated Outcome.

**Owns:**

- `Mission`;
- `Delegation` and parent/child relationships;
- global `Run`/`Attempt` references, not provider internals;
- lifecycle, dependencies, stop/escalation and pending decisions;
- global budget allocation and reconciled consumption state;
- final Mission/Delegation `Outcome` after authorized evidence/verdict composition.

**Does not own:** Harness plan/worker graph/retry implementation, Context Pack construction, provider approval/trust, effect permission, artifact/evidence contents or durable-engine history.

**Stage A:** owner fixed now; implementation deferred until a current consumer. M1 may use C08 interaction state without prematurely implementing the full Mission engine.

---

### C04 — Authority and Policy

**Purpose:** own Aurora authority semantics and deterministic decisions about disclosure and requested effects.

**Owns:**

- `AuthorityGrant` and current authority state;
- actor/delegation/resource/action authority relationships;
- policy inputs and canonical `PolicyDecision` records;
- guardrail, revocation and revalidation semantics;
- canonical `EffectRequest` identity and permission state;
- authority projections used by Context Builder, Mission and gateways.

**Does not own:** authentication product/session, raw secrets, external execution, target truth, Effect Receipt custody or framework-local tool permission.

**Stage A:** M0 minimum active; broader effect semantics grow with first effectful consumer.

---

### C05 — Capability and Provider Registry

**Purpose:** own what capabilities exist and the exact identity, compatibility, verification, trust and approval of providers.

**Owns:**

- `CapabilityDefinition` identity/version;
- `Provider`, `Harness` and `ProviderInstance` identity/lifecycle;
- Provider Manifest records;
- compatibility and conformance/verification references;
- multidimensional trust/approval, suspension, revocation and retirement;
- provider availability as a projection over observed health.

**Does not own:** active Delegation lifecycle, provider internals, Authority Grants, provider credentials, self-declared trust or Contract Model semantics.

**Stage A:** owner fixed; implementation deferred to M2.

---

### C06 — Memory, Knowledge and Context

**Purpose:** govern memory lifecycle and compile the smallest correct, attributable Context Pack for a decision or Delegation.

**Owns:**

- governed `MemoryItem` identity, scope, provenance, epistemic status and lifecycle;
- memory candidate/promotion/supersession/retention/deletion propagation;
- Knowledge Source index/reference metadata, not source truth;
- explicitly lower-authority observational/session synthesis;
- Context Builder policy and compiled `ContextPack` identity/provenance;
- derived retrieval-index registration and rebuild/delete relationships.

**Does not own:** Project state, accepted decisions, active grants, exact provider thread/workflow state, live device/provider truth or retrieval indexes as canonical truth.

**Stage A:** first substantive consumer M1. Governance/policy is Aurora-owned; heavy extraction/indexing may be on-demand workers/providers.

---

### C07 — Artifact, Observation and Evidence

**Purpose:** preserve immutable materialized outputs, observations, measurements, claims, receipts and evidence relationships independently from executors and telemetry.

**Owns:**

- `Artifact` identity, metadata, integrity and content reference;
- immutable `Observation` and `Measurement` records with provenance/quality/units;
- `Claim`;
- canonical `Receipt` records, including gateway-produced Effect Receipts;
- `Evidence` and criterion/hypothesis relationships;
- authorized `Verdict` records and limitations;
- artifact/evidence retention and provenance metadata.

**Does not own:** Git/source-code authority, Experiment lifecycle, Mission lifecycle/final Outcome, external target truth, raw telemetry by default or authority to accept its own producer claim.

**Stage A:** minimum M0 records exist; broad capability grows with M2/M3 and engineering experiments.

---

### C08 — Presence and Interaction

**Purpose:** own Presence-specific identity/lifecycle, interaction sessions, channel/environment context and translation into Aurora application operations.

**Owns:**

- `Presence` identity, registration, revocation/status and declared channel capabilities;
- `InteractionSession` identity/lifecycle;
- `ActivationRequest` semantics;
- channel, privacy/environment and handoff context;
- operator-facing interaction/delivery references.

**Does not own:** Aurora/Person identity, Project/Mission state, Authority Grants, governed memory, provider conversation state or OS/device authentication mechanisms.

**Stage A:** thin local Presence/activation adapter required; full Presence Fabric deferred.

---

### C09 — Attention and Proactivity

**Purpose:** govern proactive candidates, attention budget, urgency, suppression and delivery selection.

**Owns:** proactive candidates, urgency/relevance/confidence, attention budgets, deduplication and delivery decisions.

**Does not own:** Project truth, arbitrary Mission creation, user authority or Presence transport.

**Stage A:** `DEFER`; first consumer M5.

---

### C10 — Failure Intelligence and Evaluation

**Purpose:** correlate incidents/evaluations into causal hypotheses, improvement candidates and governed promotion proposals.

**Owns:**

- `Incident`, `Finding` and causal relationships;
- evaluation definitions/dataset references/result relationships;
- evaluation-specific experiment correlations, without taking C02's general Experiment lifecycle;
- `ImprovementOpportunity` and `ImprovementCandidate` lifecycle;
- regression/promotion/rollback proposals and failure-correlation views.

**Does not own:** production promotion authority, arbitrary holdout changes, raw telemetry, C02 Project/Experiment state or C07 immutable evidence records.

**Stage A:** `DEFER`; first consumers M6/M7.

---

### C11 — Environment and Device Registry

**Purpose:** own Aurora's canonical identity/inventory and governed metadata relationships for environments and devices while leaving live truth and actuation outside the registry.

**Owns:**

- `Environment` and `Device` identity/lifecycle;
- device/environment inventory and relationships;
- registration, pairing, revocation and ownership metadata;
- controller/firmware/calibration/manifest references;
- expected capability/protocol and trust metadata references.

**Does not own:** live telemetry, current instrument reading, target-system truth, device actuation, deterministic interlocks, firmware source or controller-local state.

**Stage A:** `DEFER`; Project may reference devices before this owner is implemented. First substantial consumers M8/M9/M10.

---

## 4.3 Application and integration components

### A01 — Core Application Coordinator

Coordinates commands/queries, domain-owner validation, current transactions/units of work, required audit/evidence references and post-commit notification. It owns sequencing, not duplicate state.

### A02 — Cognitive Coordination

```text
interaction / Mission need
→ obtain Project/authority/provider inputs
→ C06 builds Context Pack
→ C05 resolves approved cognitive capability/provider
→ A03 invokes provider
→ receive proposal/observation/artifact/capability request
→ validate and route to canonical owner
```

It owns no provider thread as Aurora truth and cannot commit provider output directly.

### A03 — Capability Fabric / Harness Integration

Owns dispatch/translation/reconciliation mechanics across provider boundaries: manifest connection, Delegation dispatch, status ingestion, cancellation, health, child capability requests, artifact references and binding translation. It does not own Registry approval, Mission lifecycle or provider-local state.

### A04 — Durable Execution Port

A domain-oriented port for timers, waits, checkpoints and durable execution. It is not a current service or canonical state owner. Engine history remains mechanism state reconciled with C03.

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

Resolves secret references into minimum short-lived credentials without exposing raw secrets to models, broad environments or general logs. Product/process choice belongs to TA-06 and the first credential-mediated consumer.

---

## 4.5 Provider runtime classes

### P01 — Cognitive Runtime Provider

Owns provider-local agent/model loops, thread/tool history, model routing mechanics, workflow snapshots, local skills/workspaces and provider-local traces/evals.

It must not own Aurora identity, Project/Mission state, Authority, global budgets, provider approval, governed memory authority, Artifact/Evidence identity, global Verdict/Outcome or effect permission.

Mastra/TypeScript remains the preferred-first substrate to evaluate for the first real consumer.

### P02 — Specialized Harness Provider

Owns domain methodology, local plans, workers, attempts, tools and internal recovery inside an Aurora Delegation. MNFS is one future provider of this class, not a Core component.

### P03 — Model / External Service Adapter

Wraps an external or local model/service behind provider policy, data minimization and attribution. Vendor message/session formats remain outside Aurora domain.

### P04 — Device / Laboratory Provider

Future controller/Harness for device/laboratory work behind C11 registration, C04 authority, E01 gateways and independent interlocks.

---

## 4.6 Mechanism adapters

Stores, indexes, blob systems, protocol bindings, UI/Voice adapters, service supervisors, OTel exporters, model SDKs, sandboxes and backup/migration mechanisms are adapters. Physical co-location never grants product ownership.

---

# 5. Canonical ownership matrix

| Entity/data family | Canonical owner | May propose/produce | Cannot commit canonical change |
|---|---|---|---|
| Aurora/Person/relationship identity | C01 | operator-approved identity process | model, Presence, provider, store |
| Provider/Harness/ProviderInstance identity | C05 | provider manifest | provider self-declaration alone |
| Presence/InteractionSession identity | C08 | Presence adapter/device | Presence cannot broaden actor identity/authority |
| Device/Environment inventory identity | C11 | controller/adapter registration | live telemetry/provider alone |
| Project/current Project state | C02 | Presence, Mission, provider result through A01 | Memory, model, Harness, UI, DB client |
| Hypothesis/Experiment global lifecycle | C02 | operator/cognitive/Harness proposal | provider-local workflow alone |
| Observation/Measurement | C07 | instrument/provider/Harness | producer cannot rewrite provenance after commit |
| Mission/Delegation/global Attempt | C03 | operator/cognitive proposal/Harness child request | provider/Harness directly |
| Global budget allocation/consumption | C03 | provider/gateway receipts | provider counter alone |
| Authority Grant/revocation | C04 | operator/approved policy process | model, Harness, Presence, gateway |
| Effect Request/Policy Decision | C04 | Mission/provider/Presence request | gateway/provider cannot self-authorize |
| Capability/Manifest/Approval | C05 | provider + conformance evidence | provider claim alone |
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
| Source code | Git | development workflow | Aurora store is not source owner |
| Live device state | device/telemetry source | P04/controller/sensor | registry/memory cannot override live verification |

### 5.1 Ownership verbs

```text
READ     consume through a query/read contract
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
→ Contract Model binding
→ Capability Fabric adapter
```

Cross-owner use cases are coordinated by A01; one domain owner does not mutate another owner's store directly.

## 6.2 Minimal shared kernel

Only genuinely cross-domain values may be shared:

- typed IDs/references;
- revision/time values;
- data classification;
- actor/correlation references;
- truly cross-domain error/result categories.

No module entities, framework types or infrastructure helpers belong in the shared kernel.

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
```

## 6.4 Canonical mutation path

```text
request/proposal
→ A01 use-case coordination
→ current owner validation
→ C04 authority/policy evaluation where required
→ owner COMMIT
→ required audit/evidence reference
→ post-commit notification/projection
```

Provider output is a proposal, observation, artifact, claim, capability request, provider-local status or controlled gateway receipt—never an implicit canonical write.

## 6.5 Cross-owner consistency

One application-level unit of work may coordinate locally co-located owners when a current invariant requires atomicity, such as:

- Project transition plus required audit reference;
- Mission/Delegation creation plus budget/authority references;
- provider approval plus verification references.

External effects never assume one ACID transaction. They require explicit request/decision/execution/receipt/reconciliation states.

## 6.6 Event rule

```text
Domain Event      owner-committed fact
Transport Message fallible delivery representation
Telemetry         diagnostic signal
Audit/Receipt/Evidence accountability/proof record
```

Missing transport delivery cannot erase canonical state or required evidence.

---

# 7. TA-02 process-boundary criteria

A logical component crosses a process boundary only when material current evidence exists for one or more:

1. different language/runtime;
2. independent lifecycle/restart;
3. fault containment;
4. security/privilege zone;
5. heavy resource profile;
6. latency/device locality;
7. remote node/consumer;
8. measured scale/concurrency;
9. independent supply-chain/upgrade cadence;
10. current consumer or Spike evidence.

Otherwise it stays logically modular inside the smallest safe deployable.

---

# 8. Three coherent Stage A approaches

## 8.1 A — Core-centric single application

```text
one persistent local application
├── active Core modules
├── local Presence/UI
├── direct model adapters
├── memory/context mechanisms
└── local effect adapters
```

**Strengths:** simplest install/debug/transaction model.

**Weaknesses:** cognition/framework/resource/privilege leakage into sovereign process; poor Mastra/polyglot fit; framework upgrades couple to Core; weak fault containment.

**Verdict:** not recommended beyond deterministic M0-like slices.

---

## 8.2 B — Early service decomposition

```text
core service
presence service
cognitive service
memory/index service
effect/credential service
Harness/provider services
```

**Strengths:** explicit isolation and independent lifecycle.

**Weaknesses:** forces RPC, service identity, discovery, supervision, distributed consistency and operational burden before Stage A consumers justify them; creates empty services around future modules.

**Verdict:** rejected for Stage A.

---

## 8.3 C — Evolutionary Sovereign Host with one early provider-runtime seam

```text
Stage A workstation
│
├── Aurora Sovereign Host — persistent
│   ├── current Aurora-owned domain modules
│   ├── A01 application coordination
│   ├── canonical state/recovery ports
│   ├── local interaction/session coordination
│   ├── thin local Presence/activation adapter initially co-packaged
│   └── provider contract endpoint
│
├── Cognitive Runtime Provider — separate/on demand at first consumer
│   ├── preferred-first Mastra/TypeScript evaluation
│   ├── model/agent/workflow execution
│   └── provider-local state only
│
├── Specialized Harness/provider processes
└── local/external model, store and effect adapters
```

### Structural reason for the early cognitive seam

- intentional polyglot possibility;
- rapid framework/model dependency cadence;
- heavy/variable resources;
- untrusted/provider-generated content and tool requests;
- independent restart;
- provider replaceability;
- accepted prohibition on provider-local state becoming global truth.

### Why other boundaries remain logical

Presence, Registry, Memory governance, Artifact/Evidence, Authority and Device inventory do not yet justify one service each. They split only on a named runtime/security/fault/resource/remote-consumer trigger.

### Strengths

- small Stage A operation;
- early isolation of highest-risk framework boundary;
- deterministic canonical state/authority;
- Mastra reuse without sovereignty loss;
- no service-per-module pattern;
- clean Stage B extraction path.

### Costs/risks

- local inter-process boundary when cognition enters;
- provider restart/reconciliation and contract compatibility become real work;
- the Sovereign Host needs dependency/fitness enforcement to avoid becoming a God process;
- co-packaging must not become permanent-placement mythology.

### Recommendation

**Approach C is the recommended TA-01/TA-02 working architecture.**

```text
one small persistent sovereign deployable
+
one independently replaceable cognitive/provider runtime seam when consumed
+
all other physical splits require a named current trigger
```

---

# 9. Proposed Stage A runtime topology

## 9.1 D01 — Aurora Sovereign Host

The current M0 sovereign executable is Go and is the practical seed of D01.

This statement does **not** extend ADR-0003 silently:

```text
accepted current fact
→ M0 Sovereign Core implementation/runtime is Go

current topology hypothesis
→ the existing Go process can seed the Stage A Sovereign Host

open cross-horizon language decision
→ implementing additional M1+ canonical modules in Go requires
   consuming architecture/ADR revalidation when material
```

The topology decision is the sovereign/provider process boundary, not a universal language mandate.

**Persistent responsibilities:** bootstrap/identity, canonical state/recovery, current Project/authority operations, A01 coordination, minimal interaction endpoint, thin Presence activation adapter, provider contract endpoint, deterministic fail-closed behavior and current audit/evidence references.

**Not continuously active by implication:** LLM sessions, Mastra, full STT/TTS, heavy retrieval/indexing, Harness workers, evaluation campaigns or device control.

**Store rule:** providers/Harnesses never receive canonical DB credentials by convenience; only owner/application paths commit canonical data. Exact stores belong to TA-05.

## 9.2 D02 — Cognitive Runtime Provider

Starts when a current interaction/Mission needs model-mediated work.

```text
input
→ versioned provider request
→ Context Pack/reference
→ authority/provider context

output
→ proposal / observation / artifact / claim /
   capability request / provider-local status
```

Core owners validate and route outputs. Provider completion cannot commit Project/Mission/Memory/Authority state.

## 9.3 D03 — Specialized Harness Provider

Independent lifecycle/execution environment; integrates through Capability/Delegation contracts. It may be long-lived or started on demand.

## 9.4 D04 — Local model/media/index workers

Replaceable mechanisms or provider-internal workers. Their first consuming capability decides whether they are embedded in P01/P02 or separate.

## 9.5 Conditional enforcement deployables

Effect Gateways/Credential Broker split physically when they hold credentials, expose material network/repository/deployment/device effects, need narrower OS privilege or must contain a provider/Core compromise. No generic privileged `execute anything` service is proposed.

---

# 10. Always-active versus on-demand

| Responsibility | Stage A disposition |
|---|---|
| identity/bootstrap/recovery | always available in D01 |
| canonical Project/authority state | always available through D01/stores |
| minimal controlled application endpoint | always active |
| local activation adapter | available when configured |
| full UI | may start on demand |
| Cognitive Runtime/Mastra | on demand at first consumer |
| external/local models | on demand |
| Context Builder | invoked on demand; governance Aurora-owned |
| memory extraction/consolidation/indexing | on demand/background after M1 design |
| Harness runtimes | independent/on demand |
| STT/TTS/Voice | on demand |
| artifact/evidence metadata | always queryable; payload mechanism as needed |
| Failure Intelligence/eval workers | deferred/on demand |
| durable engine | deferred |
| effect gateways | only per accepted effect family |
| telemetry backend/exporter | replaceable/degradable; never truth owner |

---

# 11. Runtime/language scope map

| Scope | Current posture | Boundary |
|---|---|---|
| M0 Sovereign Core | Go accepted for M0 | no provider/framework types |
| Stage A Sovereign Host | existing Go Core is seed; M1+ Go expansion requires revalidation | topology does not mandate universal Go |
| Cognitive Runtime | separate provider runtime; Mastra/TypeScript preferred-first | provider-local state only |
| AHDK | first language open | contracts remain language-neutral |
| Specialized Harness | capability-specific language/runtime | same conformance/authority boundary |
| Presence/UI | runtime open | thin adapter; no DB access |
| model/media services | mechanism/provider-specific | governed, attributed, replaceable |
| device/embedded | device-specific | gateways/interlocks; outside general model process |
| Contract Model | language-neutral | semantic/schema/binding/SDK versions separated |

---

# 12. Failure-domain and restart ownership

| Failure | Response/invariant |
|---|---|
| Sovereign Host crash | recover canonical state; reconcile providers; never infer truth from provider narrative |
| Cognitive Runtime crash | interaction/Delegation interrupted or ambiguous; Core remains available |
| provider-local store loss | provider run degrades/fails; cannot erase sovereign truth |
| Presence adapter crash | session expires/revokes; Core continues |
| Harness crash | resume/substitute/block/reconcile; no blind duplicate effect |
| canonical store failure | block governing mutation; trusted recovery; no stale authority resurrection |
| artifact payload unavailable | evidence-dependent closeout blocked |
| gateway crash/unknown receipt | preserve ambiguity and reconcile target; no blind retry |
| Credential Broker unavailable | deny credential effects; safe profile-specific reads may continue |
| model/provider outage | approved fallback or explicit unavailable; identity/state remain |
| telemetry backend outage | diagnostics degrade; audit/evidence remain separate |
| Stage B network partition | no inferred terminal state; restrict remote sessions and reconcile |

---

# 13. Stage B evolution

```text
persistent personal/home-lab node
├── Aurora Sovereign Host
├── canonical operational/governed-memory data
├── artifact/evidence coordination
├── approved cognitive/provider runtimes
└── risk-specific effect/credential services

workstation/mobile/future device
└── separate Presence client

external models/Harnesses
└── governed provider boundaries
```

Ownership does not move for Aurora identity, Project/Mission, authority, governed memory, provider trust, Artifact/Evidence or global Outcome.

Stage B adds questions for TA-04/05/06/08: service identity, Presence authentication/revocation, transport/discovery, egress zones, offline behavior, remote update/supervision, data minimization and backup/failover.

---

# 14. Split triggers

| Component | Stage A physical posture | Split trigger |
|---|---|---|
| C01/C02/C04 | Sovereign Host | proven security/scale/topology reason; ownership remains Core |
| C03 | Host when implemented | durable/multi-node needs; engine remains adapter |
| C05 | Host when implemented | remote consumers/independent trust lifecycle |
| C06 governance/Context Builder | Host | remote/shared service only with sovereignty/deletion proof |
| memory extraction/index worker | on-demand candidate | M1 eval proves resource/runtime need |
| C07 metadata | Host | payload scale may split blob store, not metadata ownership |
| C08 Presence | co-packaged thin adapter | Stage B, second device, independent OS/sensor lifecycle |
| C11 Registry | Host when implemented | device node/remote inventory consumer |
| P01 Cognitive Runtime | separate at first consumer | merge not recommended; provider boundary is intentional |
| E01/E02 | per-risk adapter | credentials/material external or physical effects |
| C10 workers | deferred | sustained independent evaluation workload |
| observability backend | adapter/external | never canonical ownership |

---

# 15. Adversarial flow tests

## 15.1 Project continuation through a model

```text
C08 interaction
→ A02 obtains C02/C04 inputs
→ C06 builds Context Pack
→ P01 reasons and returns proposal
→ A01 routes to C02/C03
→ owner commits/rejects
→ C07 records required artifacts/evidence
```

P01 cannot write Project state or promote its thread state.

## 15.2 Harness child capability request

```text
P02 requests child capability
→ A03 receives
→ C03 validates Mission relationship
→ C05 resolves approved candidates
→ C04 creates narrower authority
→ C06 builds minimized Context Pack
→ C03 commits child Delegation
```

Parent Harness cannot pass full context/authority/token transitively.

## 15.3 Memory conflicts with an ADR

C06 records a candidate/provenance; accepted ADR remains authoritative; no Project/Decision state changes.

## 15.4 General engineering experiment

```text
C02 owns Hypothesis/Experiment lifecycle
→ P02/P04 executes local method/run
→ C07 records Observation/Measurement/Artifact/Evidence
→ C02 updates Experiment state through explicit transition
→ C10 may later correlate evaluation/failure patterns
```

Provider-local run state never replaces globally meaningful Experiment state or immutable observations.

## 15.5 Material effect

```text
provider proposes Effect Request
→ C03 links purpose/budget
→ C04 evaluates authority/policy
→ E01/E02 enforce
→ C07 records receipt
→ C02/C03 update only through reconciled owner transition
```

Tool-call success cannot replace target receipt/owner transition.

## 15.6 Cognitive runtime death

Core remains; status becomes interrupted/ambiguous; resume/new Attempt/failure is explicit; sovereign truth survives.

## 15.7 Stage A to Stage B

Export/restore/migrate state, start D01 on personal node, register workstation as C08 Presence, add transport/auth boundary, preserve domain identities.

---

# 16. Approach comparison

| Criterion | A single application | B early services | C evolutionary Host |
|---|---:|---:|---:|
| Stage A simplicity | high | low | high |
| ownership clarity | medium | high | high |
| cognitive isolation | low | high | high |
| polyglot fit | low | high | high |
| local transaction simplicity | high | low | high inside Host |
| fault containment | low-medium | high | high at provider seam |
| premature distribution | low initially | very high | low |
| Stage B evolution | extraction-heavy | direct | direct at accepted seams |
| YAGNI | medium | low | high |
| God-process risk | high | low | medium with fitness rules |
| service-sprawl risk | low | high | low |
| M1/M2 fit | medium | high but expensive | high |

---

# 17. Decision disposition register

| ID | Question | Current hypothesis | Disposition / owner |
|---|---|---|---|
| TA12-D01 | minimum domain owners | C01–C11, future owners deferred | `DECIDE` through this design/review |
| TA12-D02 | Core as God module | prohibited; Core is owner set/coordinators | `DECIDE` |
| TA12-D03 | Stage A topology | Approach C | `DECIDE` |
| TA12-D04 | cognitive process seam | separate at first consumer | `DECIDE` structural; consumer proof required |
| TA12-D05 | exact Go↔TypeScript binding | open | `RESEARCH` TA-04; Spike only for consequential runtime property |
| TA12-D06 | cross-horizon Go scope | M0 Go process seeds D01; M1+ placement undecided | `RESEARCH/ADR` per consumer; not decided here |
| TA12-D07 | Presence physical process | co-packaged thin adapter initially | `DEFER` split until Stage B/second Presence/OS need |
| TA12-D08 | Memory governance placement | Aurora-owned/Host; heavy workers may separate | ownership `DECIDE`; mechanisms M1 research |
| TA12-D09 | Device/Environment owner | C11 future/deferred | `DECIDE` ownership; implementation M8/M9 |
| TA12-D10 | Effect Gateway placement | risk/family specific | `DEFER` TA-06/first effect |
| TA12-D11 | Credential Broker product/process | open | `DEFER` TA-06 |
| TA12-D12 | durable engine | none current | `DEFER` M4 |
| TA12-D13 | Registry service split | Host initially | `DEFER` M2 evidence |
| TA12-D14 | physical stores | open | `DEFER` TA-05 |
| TA12-D15 | repository strategy | blocked on TA-01/02 | `DEFER` TA-03 |
| TA12-D16 | schema/codegen | open | `RESEARCH` TA-03/04 |
| TA12-D17 | supervisor/install model | open | `DEFER` TA-08 |
| TA12-D18 | immediate Spike | none required for logical topology | `DEFER`; specify at first unprovable property |

---

# 18. Inputs to TA-03

Upon acceptance, TA-03 receives:

1. distinct domain-owner boundaries must be enforceable in source dependencies;
2. D01 is one deployable containing multiple owners, not one domain package;
3. P01 is independently versioned from its first consumer;
4. Harnesses are independent provider source/release units;
5. Presence is logical-separate but may share Stage A packaging;
6. G01 Contract Model source is language-neutral and separate from generated Go/TypeScript projections;
7. providers cannot import canonical store/internal packages;
8. build/release/versioning distinguishes Core, contracts, cognitive provider, AHDK, Harness and adapters;
9. monorepo/polyrepo/staged extraction must be compared against these boundaries;
10. Development Harness builds/tests/reviews Aurora but is not a runtime dependency.

TA-03 remains separately gated.

---

# 19. Remediation from v0.1.0

This revision resolves the initial review findings:

- `TA12-F01`: added C11 Environment and Device Registry;
- `TA12-F02`: narrowed C01 and assigned entity-specific identities to C05/C08/C11;
- `TA12-F03`: allocated Hypothesis/Experiment to C02 and Observation/Measurement to C07;
- `TA12-F04`: separated M0 Go fact, Stage A host hypothesis and cross-horizon Go decision;
- `TA12-F05`: added G01 Canonical Contract Model as a non-runtime governed asset.

Worklog/fixed-review continuity remains a closeout task after validation.

---

# 20. Explicit non-decisions

This proposal does not select:

- monorepo/polyrepo/package layout;
- Go as universal Aurora language;
- concrete Mastra implementation/version;
- HTTP/REST/gRPC/Connect/MCP/A2A/events/local IPC;
- schema language/code generator;
- new relational/vector/graph/object/telemetry stores;
- Keycloak/Zitadel/Authentik/Ory/SPIFFE/OPA/Cedar/Vault;
- supervisor/containers/Kubernetes;
- model/Voice/sandbox/durable-engine/observability products;
- device/lab protocol;
- first AHDK language;
- implementation roadmap or production code.

---

# 21. Acceptance criteria and operator gate

A reviewer must confirm:

1. representative global concepts have one owner;
2. provider/model/Harness cannot write canonical state directly;
3. Sovereign Host is not one God module;
4. Stage A remains operationally small;
5. cognitive/runtime failure cannot erase sovereign truth;
6. design does not create one service per module;
7. Stage B changes mechanisms, not identity/ownership;
8. M0 decisions are not globalized silently;
9. Mastra remains preferred-first/replaceable;
10. no hidden repository/protocol/database/auth product selection exists;
11. unresolved mechanisms have treatment/consumer/owner;
12. TA-03 receives structure without being pre-decided.

```text
ACCEPT
→ accept Approach C and ownership/dependency matrices
→ no implementation or TA-03 finalization follows automatically

REVISE
→ return exact ownership/topology changes

REJECT
→ preserve accepted Technical Architecture Map
→ replace this proposal with another reviewed approach
```

Until acceptance:

```text
TA-01/TA-02 proposal: NON-CANONICAL
TA-03 finalization: BLOCKED
Architecture Spike execution: NOT AUTHORIZED
Aurora implementation: PAUSED
```
