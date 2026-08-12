---
id: DESIGN-AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY
title: Aurora TA-01/TA-02 Module Ownership and Runtime Topology
document_type: system_architecture_design
form: reference
authority: design
status: proposed
version: 0.1.0
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

## 1. Purpose and decision boundary

This design proposes the first cross-system technical answer to the coupled questions:

```text
TA-01
What is the minimum coherent set of Aurora technical components,
and who owns each canonical responsibility and datum?

TA-02
Which of those responsibilities share a process in Stage A,
which cross an independent runtime boundary,
and how does the topology evolve to Stage B?
```

The proposal converts the accepted product architecture into a technical structure without selecting repository strategy, wire protocol, universal database, authentication product, policy engine, service supervisor, model provider, Voice provider or observability backend.

It is deliberately a **proposed design for operator review**. It does not authorize implementation, Architecture Spike execution, TA-03 finalization, M0 R7 continuation or M0 R8.

---

## 2. Fixed source baseline

### 2.1 Canonical baseline

```text
repository: developmentconexus-ops/aurora_project
canonical branch: main
source revision: 9cbce1efe4f742f90623b894c0c1ba2eaa3cebcc
```

### 2.2 Frozen implementation evidence

```text
branch: feat/m0-r7-sovereign-core-20260810
observed revision: 7ec999b093205a9d82eef2802eca60330d96e14d
classification: FROZEN / PRESERVED / NON-CANONICAL
```

The frozen candidate demonstrates that the M0 pattern below is executable:

```text
adapters / composition
        ↓
application use cases
        ↓
domain owners + ports
        ↓
mechanism adapters
```

That evidence is useful, but the M0 source tree is not generalized into the complete Aurora architecture by implication.

### 2.3 Governing constraints

This proposal preserves these accepted constraints:

1. Aurora is one persistent sovereign cognitive/operational control plane.
2. One durable concept has one canonical owner.
3. Identity, Project/Mission meaning, authority, governed context, provider trust and global Outcome remain Aurora-owned.
4. Harnesses own specialized local methodology, attempts, workers and provider-local state.
5. Models, runtimes, protocols, SDKs and stores are mechanisms, not product ontology.
6. Aurora owns canonical Contract Model semantics; bindings remain replaceable.
7. AHDK is a first-party Golden Path, not the specification or security boundary.
8. Current operational state, memory, knowledge, history, evidence, audit and telemetry remain distinct.
9. Material effects are deterministically enforced outside model judgment.
10. Logical modularity precedes physical distribution.
11. Processes are disposable; canonical state survives them.
12. Stage A remains one Leandro-controlled workstation with a persistent minimum and on-demand cognition.
13. M0 Go, SQLite, JSON/JCS and OTel decisions remain scoped unless separately promoted.
14. Mastra is preferred-first to evaluate for agentic/cognitive runtimes, never as sovereign state or authority owner.

### 2.4 Deferred Presence detail

Already accepted Stage A activation and locked-workstation constraints remain inputs. Further questions about individual users, speaker behavior or public-mode interaction are `DEFER` unless they change a module, process, security boundary or near implementation decision.

---

## 3. Architectural vocabulary

The proposal distinguishes six kinds of technical element.

| Kind | Meaning | May own canonical Aurora state? |
|---|---|---:|
| **Domain owner** | owns one durable Aurora concept and its invariants | yes, only for its named scope |
| **Application coordinator** | orchestrates a use case across owners and mechanisms | no duplicate truth |
| **Deployable / composition boundary** | packages modules, adapters and lifecycle into a running unit | no merely because it hosts them |
| **Enforcement boundary** | deterministically allows, denies or executes an effect | owns execution state only; canonical receipts are recorded through Aurora owners |
| **Provider runtime** | performs replaceable cognitive or specialized execution | provider-local state only |
| **Mechanism adapter** | storage, protocol, model, OS, telemetry or UI implementation | no product meaning |

This distinction prevents the term `Core` from becoming either:

- one giant domain object that owns everything; or
- a synonym for every process installed with Aurora.

### 3.1 Aurora Core

In this design, **Aurora Core** is the governed set of Aurora-owned domain modules and application coordinators.

### 3.2 Aurora Sovereign Host

The **Aurora Sovereign Host** is the Stage A deployable/composition boundary that runs the Core modules selected for the current executable horizon and owns their process lifecycle. Hosting a module does not merge its ownership with other modules.

### 3.3 Provider runtime

A **provider runtime** is independently restartable, replaceable and limited to provider-local execution. A Cognitive Runtime or Harness runtime may be first-party without becoming part of the sovereign state boundary.

---

# 4. TA-01 — Minimum coherent component model

## 4.1 Component model overview

```text
┌──────────────────────────────────────────────────────────────┐
│                    AURORA-OWNED DOMAIN                       │
│                                                              │
│ Identity & Relationship                                      │
│ Project & World State                                        │
│ Mission & Delegation Control                                 │
│ Authority & Policy                                           │
│ Capability & Provider Registry                               │
│ Memory, Knowledge & Context                                  │
│ Artifact & Evidence                                          │
│ Presence & Interaction                                       │
│ Attention & Proactivity                 [future/deferred]     │
│ Failure Intelligence & Evaluation       [future/deferred]     │
└───────────────────────────────┬──────────────────────────────┘
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

Mechanisms such as stores, protocols, event transport, model SDKs, OTel and operating-system services remain adapters around these responsibilities.

---

## 4.2 Canonical domain-owner catalog

### C01 — Identity and Relationship

**Purpose:** preserve stable Aurora, Leandro/person and relationship identities independently from process, model, Presence or authentication product.

**Owns:**

- `AuroraIdentity`;
- `Person` and owner/operator references;
- relationship/profile references and protected identity settings;
- stable actor identity references and identity lifecycle semantics;
- constitutional identity references, not the constitutional documents themselves.

**Does not own:**

- operating-system login;
- biometric recognition;
- OAuth/OIDC sessions;
- device identity mechanism;
- model personality state;
- Authority Grants.

**Stage A:** active Core module.

**First consumers:** M0 continuity, M1 interaction, every later actor chain.

---

### C02 — Project and World State

**Purpose:** preserve Project identity, current Project operational meaning and governed relationships/views over external authoritative sources.

**Owns:**

- `Project`;
- Project current-state revisions and current pointer;
- `Goal`/`Objective` when represented as Project-owned intent;
- bounded Project/world relationships and temporal validity;
- Project snapshots and current-next-action source references;
- references to decisions, repositories, devices, knowledge and active work.

**Does not own:**

- accepted ADR or source-code contents;
- Mission/Delegation lifecycle;
- governed Memory Items;
- live device telemetry;
- provider-local execution state;
- Authority Grants.

**Stage A:** active Core module, initially extending the proven M0 owner rather than becoming a universal graph.

**First consumers:** M0, M1 Project context.

---

### C03 — Mission and Delegation Control

**Purpose:** own global work intent, decomposition, cross-provider relationships, lifecycle, budgets and consolidated global Outcome.

**Owns:**

- `Mission`;
- `Delegation` and parent/child relationships;
- global `Attempt`/`Run` references, not provider internals;
- Mission/Delegation lifecycle and pending decisions;
- global budget allocation and reconciled consumption state;
- dependency/stop/escalation state;
- final Mission/Delegation `Outcome` after authorized evidence/verdict composition.

**Does not own:**

- Harness local plan, worker graph or retry implementation;
- Context Pack construction;
- provider approval/trust;
- permission to execute an effect;
- artifacts/evidence contents;
- durable-engine history.

**Stage A:** logical owner fixed now; implementation deferred until a current Mission/Delegation consumer.

**First consumers:** M2/M3, with M1 possibly using a smaller interaction/activity boundary without prebuilding the full mission engine.

---

### C04 — Authority and Policy

**Purpose:** own Aurora authority semantics and deterministic decisions about what may be disclosed or requested as an effect.

**Owns:**

- `AuthorityGrant` and current authority state;
- actor/delegation/resource/action authority relationships;
- policy inputs and canonical `PolicyDecision` records;
- guardrail and revocation semantics;
- canonical `EffectRequest` identity and permission state;
- authority projections used by Context Builder, Mission and gateways.

**Does not own:**

- authentication product/session;
- raw secret values;
- external-system execution;
- target-system truth;
- `EffectReceipt` content custody;
- model/tool permission state as canonical authority.

**Stage A:** M0 minimum is active; broader policy/effect semantics stay logical until first effectful consumer.

**First consumers:** M0 internal actions, M1 disclosure boundary, M3 effects.

---

### C05 — Capability and Provider Registry

**Purpose:** own what capabilities exist, who claims to offer them, exact provider/build identity, compatibility, verification, trust and approval scope.

**Owns:**

- `CapabilityDefinition` identity/version;
- `Provider`, `Harness` and `ProviderInstance` identity;
- Provider Manifest records;
- compatibility assessments;
- conformance/verification references;
- multidimensional trust and approval lifecycle;
- suspension, revocation and retirement;
- provider availability as a current projection over observed health.

**Does not own:**

- active Delegation lifecycle;
- provider internal state;
- Authority Grants;
- provider credentials;
- provider self-declared trust as accepted truth;
- contract semantics owned by Aurora specifications.

**Stage A:** logical owner fixed now; implementation deferred to M2.

**First consumer:** M2 Capability Registry/AHDK.

---

### C06 — Memory, Knowledge and Context

**Purpose:** govern memory write/manage/read lifecycles and compile the smallest correct, attributable Context Pack for a decision or Delegation.

**Owns:**

- governed `MemoryItem` identity, scope, provenance, epistemic status and lifecycle;
- memory candidates, promotion, supersession, retention and deletion propagation;
- Knowledge Source index/reference metadata, not source truth itself;
- observational/session synthesis as explicitly lower-authority data;
- Context Builder policy and compiled `ContextPack` identity/provenance;
- derived retrieval-index registration and rebuild/delete relationships.

**Does not own:**

- Project current state;
- accepted decisions/source documents;
- active Authority Grants;
- exact provider-local thread/workflow state;
- live device/provider truth;
- vector/full-text/graph index as canonical truth.

**Stage A:** first substantive architecture consumer is M1. Canonical governance may live in the Sovereign Host; heavy extraction/indexing mechanisms may be on-demand providers/workers.

**First consumer:** M1 governed conversation/project context/memory.

---

### C07 — Artifact and Evidence

**Purpose:** preserve materialized outputs, claims, receipts and evidence relationships independently from executors and telemetry.

**Owns:**

- `Artifact` identity, metadata, integrity and content reference;
- `Claim`;
- immutable canonical `Receipt` records, including Effect Receipts supplied by gateways;
- `Evidence` and criterion/hypothesis relationships;
- authorized `Verdict` records and limitations;
- artifact/evidence retention and provenance metadata.

**Does not own:**

- raw source-code authority, which remains Git;
- Mission lifecycle or final Mission Outcome;
- external target-system state;
- telemetry as evidence by default;
- authority to accept its own producer claim.

**Stage A:** minimum M0 records exist; broad artifact/evidence capability grows with M2/M3.

**First consumers:** M0 evidence metadata, M2 provider evidence, M3 composition.

---

### C08 — Presence and Interaction

**Purpose:** own contextual interaction sessions, channel/device/environment context and translation between a Presence and Aurora application operations.

**Owns:**

- `Presence` registration/status and Presence-scoped capability declarations;
- `InteractionSession` identity and lifecycle;
- `ActivationRequest` semantics;
- channel, privacy/environment and handoff context;
- operator-facing interaction state and delivery references.

**Does not own:**

- Aurora or Person identity;
- Project/Mission state;
- Authority Grants;
- governed memory;
- model conversation state as product truth;
- device operating-system authentication mechanism.

**Stage A:** a thin local Presence/activation adapter is required; full Presence Fabric remains deferred.

**First consumers:** Stage A local interaction, M8 for multi-Presence.

---

### C09 — Attention and Proactivity

**Purpose:** govern proactive candidates, attention budget, urgency, suppression and delivery selection.

**Owns:**

- proactive notification candidates;
- urgency/relevance/confidence classifications;
- attention budgets and deduplication state;
- delivery decisions and explanations.

**Does not own:**

- Project truth;
- arbitrary background Mission creation;
- user authority;
- Presence transport.

**Stage A:** `DEFER`; do not create an empty module before a proactive consumer.

**First consumer:** M5.

---

### C10 — Failure Intelligence and Evaluation

**Purpose:** turn incidents, findings and evaluations into causal hypotheses, improvement candidates and governed promotion proposals.

**Owns:**

- `Incident`, `Finding` and causal relationships;
- evaluation definitions/dataset references and result relationships;
- `ImprovementOpportunity` and `ImprovementCandidate` lifecycle;
- regression/promotion/rollback proposals;
- failure-correlation views.

**Does not own:**

- production promotion authority;
- arbitrary changes to evaluation/holdout rules;
- raw telemetry;
- the canonical state of affected modules.

**Stage A:** `DEFER`; preserve data/contract hooks only when a current capability produces incidents or evidence.

**First consumers:** M6/M7 and later self-improvement.

---

## 4.3 Application and integration components

These components coordinate domain owners but cannot become duplicate state owners.

### A01 — Core Application Coordinator

Coordinates commands/queries across current domain owners, transaction boundaries and ports. It owns use-case sequencing, not a second Project/Mission/Authority model.

Representative responsibilities:

- validate command context;
- invoke one or more domain owners;
- coordinate atomic canonical mutations where supported;
- record required audit/evidence references;
- return structured application results;
- publish post-commit domain notifications through adapters.

### A02 — Cognitive Coordination

Coordinates deterministic context construction and replaceable cognitive execution.

```text
interaction / Mission need
→ request Context Pack from C06
→ select an approved cognitive capability/provider through C05
→ invoke provider through A03
→ receive proposal/observation/artifact/capability request
→ validate and route to canonical owner
```

It owns no durable model thread as Aurora truth and cannot commit provider output directly.

### A03 — Capability Fabric / Harness Integration

Maps Aurora Delegations and provider contracts to concrete providers/bindings.

Responsibilities:

- dispatch;
- lifecycle/status ingestion;
- cancellation and health;
- provider snapshot/reconciliation;
- child capability request routing;
- artifact/evidence references;
- binding translation.

It does not own Capability approval, Mission lifecycle or provider-local state.

### A04 — Durable Execution Port

A domain-oriented port for timers, waits, checkpoints and durable execution. It is not a current service or state owner. Any engine history remains mechanism state and reconciles with C03.

**Stage A disposition:** interface remains conceptual until M4 or an earlier accepted consumer proves need.

---

## 4.4 Enforcement components

### E01 — Effect Gateway family

Deterministically enforces an allowed effect against a specific external boundary.

Potential families:

- filesystem;
- network;
- repository;
- external communication;
- deployment;
- financial;
- device/laboratory.

A gateway:

1. receives a canonical Effect Request and Policy Decision reference;
2. obtains minimum credentials through E02 when necessary;
3. executes, denies or reports ambiguity;
4. emits a receipt;
5. never grants itself broader authority.

Physical process isolation is risk-driven and deferred per effect family.

### E02 — Credential Broker boundary

Resolves secret references into minimum short-lived credentials. It must not expose raw secrets to models, broad process environments or general logs.

**Stage A disposition:** logical boundary now; physical product/process choice belongs to TA-06 and the first credential-mediated effect.

---

## 4.5 External/provider runtime classes

### P01 — Cognitive Runtime Provider

Executes agent/model reasoning, planning, local workflow and provider-local memory. Mastra/TypeScript is the accepted preferred-first substrate to evaluate for the first real consumer.

May own locally:

- agent loop/thread/tool history;
- model routing and invocation mechanics;
- provider-local workflow snapshots;
- observational synthesis candidates;
- local skills/workspaces;
- provider-local traces/evals.

Must not own:

- Aurora identity;
- Project/Mission global state;
- Authority Grants;
- global budgets;
- provider approval/trust;
- governed memory authority;
- Artifact/Evidence identity;
- global Verdict/Outcome;
- effect permission.

### P02 — Specialized Harness Provider

Owns domain methodology, workers, local plans, attempts, tools and recovery inside an Aurora Delegation. MNFS is one future instance of this class, not a Core component.

### P03 — Model / External Service Adapter

Wraps an external or local model/service endpoint behind provider policy, data minimization and attribution. The vendor's message/session format remains outside Aurora domain contracts.

### P04 — Device / Laboratory Provider

Future controller/Harness class for device and laboratory work. It remains outside the general model process and behind gateways/interlocks.

---

## 4.6 Mechanism adapter classes

These remain technical mechanisms rather than canonical owners:

- operational-state adapters;
- memory/knowledge/index adapters;
- artifact-blob adapters;
- protocol/binding adapters;
- UI/Voice/Presence adapters;
- operating-system supervisor adapters;
- observability exporters/backends;
- model SDK adapters;
- sandbox/execution-environment adapters;
- backup/export/migration adapters.

A mechanism may be physically co-located with a domain owner without acquiring its meaning.

---

# 5. Canonical ownership matrix

## 5.1 Representative entity ownership

| Entity / data family | Canonical owner | May propose / produce | Cannot commit canonical change |
|---|---|---|---|
| Aurora identity | C01 Identity & Relationship | operator-approved constitutional process | model, Presence, provider, store adapter |
| Person/relationship profile | C01 | operator, governed memory candidate | model/provider directly |
| Project and current Project state | C02 Project & World State | Presence, Mission, provider result through application | Memory, model, Harness, UI, DB client |
| Project snapshot | C02 projection | C06 may request/compile reference | provider-local summary |
| Mission / Delegation / global Attempt | C03 Mission & Delegation | operator, cognitive proposal, Harness child request | provider/Harness directly |
| Global budget allocation/consumption | C03 | provider/gateway receipts | provider-local counter alone |
| Authority Grant / revocation | C04 Authority & Policy | operator/approved policy workflow | model, Harness, Presence, gateway |
| Effect Request / Policy Decision | C04 | Mission, provider or Presence may request | gateway/provider cannot self-authorize |
| Capability / Provider / Manifest / Approval | C05 Registry | provider manifest, conformance/evidence | provider self-declaration alone |
| Governed Memory Item | C06 Memory & Context | observer/model/provider/user produces candidate | model/provider/index directly |
| Context Pack | C06 | C02/C03/C04/C05 inputs | provider cannot broaden it |
| Exact model/provider thread state | P01 provider-local | cognitive runtime | never canonical Aurora state by convenience |
| Artifact metadata/content reference | C07 Artifact & Evidence | any authorized producer | producer cannot rewrite identity/provenance after commit |
| Claim | C07 | executor/provider | claim cannot become Verdict automatically |
| Effect Receipt | C07 canonical receipt record | E01 gateway produces | model/provider cannot fabricate target result |
| Evidence / Verdict record | C07 | authorized verifier/decider | producer alone unless policy explicitly permits |
| Mission Outcome | C03 | C07 supplies evidence/verdict package | child completion cannot close parent Mission |
| Presence / InteractionSession / ActivationRequest | C08 | device/UI/Voice adapter | Presence cannot alter identity/authority/Project directly |
| Proactive candidate / attention budget | C09 | any module may nominate | provider cannot self-interrupt outside policy |
| Incident / ImprovementCandidate | C10 | modules, telemetry/evidence processors | candidate cannot self-promote |
| Raw telemetry | telemetry mechanism/source | all components | cannot become product state/verdict automatically |
| Source code | Git repository | development workflow | Aurora state store is not source owner |
| Live device state | device/telemetry source | P04 controller/sensor | memory snapshot cannot override live verification |

## 5.2 Ownership verbs

For every canonical concept, the architecture distinguishes:

```text
READ
→ consume through a query/read contract

PROPOSE
→ submit a candidate or requested transition

VALIDATE
→ determine whether the proposal satisfies the owner's invariants

COMMIT
→ create the canonical revision/record

PROJECT
→ build a non-authoritative view over one or more sources
```

A component allowed to `READ` or `PROPOSE` does not inherit `COMMIT` authority.

---

# 6. Dependency and mutation rules

## 6.1 Allowed dependency direction

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
Harness / cognitive domain code
          ↓
AHDK or direct provider boundary
          ↓
Aurora Contract Model binding
          ↓
Capability Fabric adapter
```

Cross-module use case:

```text
Domain Owner A
      ↑
application coordinator
      ↓
Domain Owner B
```

The coordinator invokes both owners; one domain module does not mutate another module's store directly.

## 6.2 Minimal shared kernel

Domain modules may share only a deliberately small, Aurora-owned kernel of cross-cutting value concepts such as:

- stable typed references/IDs;
- time and revision value types;
- data classification;
- actor/correlation references;
- result/error categories that are truly cross-domain.

The shared kernel must not become a dumping ground for module entities, framework types or infrastructure helpers.

## 6.3 Forbidden dependencies

```text
domain → database driver / ORM
domain → model or Mastra type
domain → MCP/A2A/gRPC/HTTP type
domain → UI framework
domain → OTel exporter/backend
domain → operating-system service API
domain → provider-local store
provider/Harness → canonical Aurora database
Presence → canonical Aurora database
memory index → Project/Authority writer
telemetry → state mutation authority
Effect Gateway → authority grant creation
```

## 6.4 Canonical mutation rule

Every canonical mutation follows:

```text
request/proposal
→ application coordination
→ current owner validation
→ current authority/policy evaluation where required
→ canonical owner commit
→ required audit/evidence reference
→ post-commit notification/projection
```

Provider output is always one of:

- proposal;
- observation;
- artifact;
- claim;
- capability request;
- provider-local status;
- receipt from a controlled gateway.

It is never an implicit canonical write.

## 6.5 Cross-owner consistency

A cross-owner operation may use one application-level unit of work when all affected owners are locally co-located and a current invariant requires atomicity.

Examples likely eligible:

- accepted Project transition plus required audit reference;
- Mission/Delegation creation plus initial budget and authority references;
- provider approval transition plus associated verification references.

The architecture must not assume one ACID transaction for an external effect. External effects use explicit request/decision/execution/receipt/reconciliation states.

## 6.6 Event rule

A domain owner may emit a meaningful post-commit domain event about its own state. Application/adapters may emit transport messages and telemetry.

```text
Domain Event
→ owner-committed fact

Transport Message
→ fallible delivery representation

Telemetry
→ diagnostic signal

Audit/Receipt/Evidence
→ accountability/proof record
```

A missed transport message cannot erase the canonical state or required artifact/evidence.

---

# 7. TA-02 — Process-boundary criteria

A logical component crosses a physical process boundary only when one or more of the following are material now:

1. **language/runtime constraint** — e.g. TypeScript/Mastra versus Go;
2. **independent lifecycle** — it must start, stop, update or recover independently;
3. **fault containment** — a crash or leak must not terminate/corrupt the sovereign host;
4. **security/privilege zone** — it receives untrusted code, network, credentials or device access;
5. **resource profile** — model, indexing or media work is materially heavy;
6. **latency/device locality** — work must remain close to a microphone, GPU, instrument or remote Presence;
7. **remote consumer/location** — it runs on another node;
8. **independent scale/concurrency** — measured demand requires separate capacity;
9. **upgrade cadence/supply chain** — rapid provider/framework changes should not force Core upgrade;
10. **evidence** — a current consumer or Spike proves the boundary necessary.

Absent a material trigger, the component remains a logical module inside the smallest safe deployable.

---

# 8. Three coherent Stage A topology approaches

## 8.1 Approach A — Core-centric single application

### Shape

```text
Stage A workstation
└── one persistent Aurora application process
    ├── all active Core domain modules
    ├── local Presence/UI adapter
    ├── direct model-provider adapters
    ├── initial memory/context mechanisms
    └── local effect adapters

external Harnesses only when explicitly integrated
```

### Advantages

- minimum process/IPC/supervision burden;
- simple local debugging and transactions;
- fastest path for a small deterministic application;
- easiest installation and restart model.

### Weaknesses

- encourages cognition/model orchestration to leak into the Core;
- poor fit for Mastra/TypeScript without child-process special cases;
- one process accumulates network, model, UI and storage privileges;
- heavy indexing/model work can degrade the sovereign state process;
- framework upgrades couple to Core release;
- failure containment is weak.

### Stage B migration

Requires extracting Presence and cognitive responsibilities after they have already accumulated in one process. Domain seams can reduce the cost, but runtime coupling is likely.

### Assessment

**Not recommended as the cross-system baseline.** It is acceptable only for M0-like deterministic slices with no cognitive runtime consumer.

---

## 8.2 Approach B — Early service decomposition

### Shape

```text
Stage A workstation
├── aurora-core service
├── aurora-presence service
├── aurora-cognitive service
├── aurora-memory/index service
├── aurora-effect/credential service
└── Harness/provider services
```

### Advantages

- explicit isolation and independent lifecycle;
- clear polyglot boundaries;
- natural Stage B distribution;
- lower risk that provider/framework state becomes Core internals;
- narrower privileges per service in principle.

### Weaknesses

- forces local RPC, service identity, discovery, configuration and supervision before consumers require them;
- multiplies failure/recovery states and compatibility surfaces;
- makes a one-user workstation operate like a distributed platform;
- creates pressure for brokers, shared infrastructure and deployment tooling;
- cross-service consistency becomes harder before domain rules are stable;
- high risk of empty services around future-only modules.

### Stage B migration

Physical migration is straightforward, but the cost is paid too early and continuously.

### Assessment

**Rejected for Stage A.** It violates the accepted preference for logical modularity and smallest operational topology.

---

## 8.3 Approach C — Evolutionary Sovereign Host with one early provider-runtime seam

### Shape

```text
Stage A workstation
│
├── Aurora Sovereign Host — Go — persistent
│   ├── active Aurora-owned domain modules
│   ├── application coordination
│   ├── canonical state/recovery ports
│   ├── local interaction/session coordination
│   ├── thin local Presence/activation adapter initially co-packaged
│   ├── capability/provider contract endpoint
│   └── logical authority/effect interfaces
│
├── Cognitive Runtime Provider — on demand — separate process
│   ├── preferred-first Mastra/TypeScript evaluation at first consumer
│   ├── model/agent/workflow execution
│   ├── provider-local memory/thread/snapshot state
│   └── no direct canonical-store access
│
├── Specialized Harness/provider processes — on demand or independently managed
│
├── local/external model services — adapters/providers
│
└── stores and effect adapters
    ├── canonical stores writable only through Aurora owners
    ├── artifact/blob mechanisms
    └── risk-specific gateways separated when introduced
```

### Why one early physical seam is justified

The boundary between sovereign deterministic state and cognitive/provider-local runtime already has current evidence:

- different expected languages/runtimes;
- rapid framework/model dependency cadence;
- heavy and variable resource use;
- untrusted/provider-generated content and tool requests;
- independent restart requirement;
- provider replaceability;
- explicit accepted rule that provider-local state cannot become global truth.

This seam should exist before a first cognitive implementation, even though the exact transport remains TA-04 work.

### Why other boundaries remain logical initially

Presence, Registry, Memory governance, Artifact/Evidence and internal Authority modules do not yet have enough Stage A evidence to justify one service each. They remain separate domain owners inside the Sovereign Host until a split trigger is reached.

Heavy memory extraction/indexing may later run as an on-demand worker/provider while memory governance and Context Builder policy remain Aurora-owned.

### Advantages

- preserves a small Stage A operating model;
- isolates the highest-risk polyglot/framework boundary early;
- keeps canonical state and authority deterministic;
- supports Mastra reuse without making Mastra the Core;
- lets a thin local Presence remain simple;
- avoids one service per conceptual module;
- gives Stage B a clean extraction path;
- allows Harnesses to remain truly independent providers.

### Weaknesses

- introduces local inter-process communication when cognition enters;
- requires explicit provider restart/reconciliation semantics;
- the Sovereign Host still needs architectural fitness checks to avoid becoming a God process;
- co-packaged Presence and logical gateways must not be mistaken for permanent physical placement;
- Go/TypeScript contract compatibility becomes a real maintenance concern.

### Stage B migration

```text
Stage A
workstation
├── Sovereign Host
├── local Presence adapter
└── Cognitive/Harness providers

        ↓ explicit migration/export/restore

Stage B
persistent personal node
├── Sovereign Host + canonical data
├── cognitive/providers as approved local/remote runtimes
└── effect/credential services by risk

workstation/mobile
└── separate Presence clients
```

The migration adds transport, service identity, network policy and Presence authentication. It does not move canonical ownership to the Presence or cognitive runtime.

### Assessment

**Recommended working architecture for operator review.**

---

# 9. Proposed Stage A runtime topology

## 9.1 Persistent minimum

### Deployable D01 — Aurora Sovereign Host

**Runtime:** Go, consistent with the accepted initial Sovereign Core direction; exact future global scope remains governed by later decisions.

**Always available responsibilities:**

- bootstrap and stable Aurora identity;
- canonical state access and integrity/recovery;
- active Project and current authority operations;
- application command/query coordination;
- minimal interaction/session coordination;
- local Presence activation endpoint/adapter;
- provider contract endpoint and runtime start/stop coordination;
- deterministic denial/fail-closed behavior;
- audit/evidence references required by current operations.

**Not continuously active by implication:**

- LLM/model sessions;
- Mastra;
- full STT/TTS;
- heavy retrieval/indexing;
- Harness workers;
- evaluation campaigns;
- device control.

### Store posture

- canonical stores are local to infrastructure controlled by Leandro;
- only Aurora-owned application/domain paths commit canonical state;
- provider/Harness processes receive contracts and references, never database credentials by convenience;
- exact physical stores and co-location belong to TA-05.

## 9.2 On-demand deployables

### D02 — Cognitive Runtime Provider

Starts when a current interaction/Mission needs model-mediated work.

Minimum boundary:

```text
input
→ versioned provider request / Context Pack reference / authority context

output
→ proposal / observation / artifact reference / claim /
   capability request / provider-local status
```

The Core validates and routes outputs. Provider-local completion does not commit Project/Mission/Memory/Authority state.

### D03 — Specialized Harness Provider

Each Harness retains its own lifecycle and execution environment. It integrates through Capability/Delegation contracts and may be long-lived or started on demand.

### D04 — Local model/media/index workers

Local model servers, embeddings, STT/TTS, media or indexing workers remain replaceable mechanisms. They may run as provider-internal processes or separate providers depending on the first consumer.

## 9.3 Conditional enforcement deployables

Effect Gateways and Credential Broker components may be in-process adapters only for explicitly low-risk local boundaries. They cross a process/security boundary when they:

- hold credentials;
- expose network/repository/deployment effects;
- run untrusted provider requests;
- control a device;
- require a narrower OS privilege set;
- must survive/contain a Core or provider compromise independently.

No generic privileged `execute anything` process is proposed.

---

# 10. Always-active versus on-demand matrix

| Responsibility | Stage A disposition | Reason |
|---|---|---|
| identity/bootstrap/recovery | always active in Sovereign Host | continuity and availability |
| canonical Project/authority state | always available in Sovereign Host/store | sovereign truth |
| minimal application endpoint | always active | interfaces/providers need controlled entry |
| local activation adapter | available when configured | accepted Stage A trigger experience |
| full interaction UI | may start on demand | not required for Core continuity |
| cognitive runtime/Mastra | on demand | resource/framework/fault isolation |
| external/cloud models | on demand | provider policy, cost and network |
| governed Context Builder | invoked on demand; governance in Host | context is decision-specific |
| memory extraction/consolidation/indexing | on demand/background only when M1 defines it | heavy/probabilistic/rebuildable mechanisms |
| Harness runtimes | independent/on demand | specialized provider lifecycle |
| STT/TTS/Voice pipeline | on demand | privacy/resource boundary |
| artifact/evidence metadata | always queryable; payload mechanism as needed | acceptance/recovery dependency |
| evaluation/Failure Intelligence workers | deferred/on demand | no current consumer |
| durable scheduler/engine | deferred | M4 or earlier proven consumer |
| effect gateways | present only per accepted effect family | risk-specific enforcement |
| observability exporter/backend | replaceable and degradable by profile | telemetry is not product truth |

---

# 11. Runtime and language scope map

| Scope | Current technical posture | Explicit boundary |
|---|---|---|
| Sovereign Host / initial Core | Go is the accepted initial runtime direction; M0 evidence informs the pattern | no provider/framework types in domain |
| Cognitive Runtime | separate provider runtime; Mastra/TypeScript preferred-first to evaluate | no canonical state/authority ownership |
| AHDK | first language not selected | generated/public contracts remain language-neutral |
| Specialized Harness | language/runtime chosen by capability | same Aurora conformance/authority boundary |
| Presence/UI | runtime not selected | thin adapter; no canonical DB access |
| local model/media services | provider/mechanism specific | attributed, governed, replaceable |
| device/embedded | capability/device specific | outside general cognitive runtime; gateways/interlocks |
| Contract Model | language-neutral | semantic, schema, binding and SDK versions separated |

No decision here makes TypeScript mandatory for all cognition or Go mandatory for all Aurora components.

---

# 12. Failure-domain and restart ownership

| Failure | Canonical owner response | Required invariant |
|---|---|---|
| Sovereign Host crash | recover canonical state; inspect/reconcile active providers | no identity/state invention from provider narrative |
| cognitive runtime crash | C03/A02 mark interaction/delegation interrupted or ambiguous; query provider snapshot when supported | Core remains available; no auto-success |
| provider-local store loss | provider run degrades/fails or restarts under contract | cannot erase Project/Mission/Authority truth |
| Presence adapter crash | C08 expires/revokes session context; Core continues | interaction loss is not Core identity loss |
| Harness crash | C03/A03 resume, substitute, block or reconcile | stable Delegation identity; no blind duplicate effect |
| canonical store failure | block governing mutation; recover/restore under trusted procedure | stale grants/state cannot reactivate silently |
| artifact payload unavailable | C07 prevents evidence-dependent closeout | missing artifact cannot be replaced by claim |
| gateway crash before known receipt | C04/C03 preserve ambiguity and reconcile target | no blind retry |
| Credential Broker unavailable | deny credential-mediated effects; safe local reads may continue by profile | fail closed for material effects |
| model/provider outage | approved fallback or explicit capability unavailable | Aurora identity/state remain |
| telemetry backend outage | diagnostic degradation reported | audit/evidence required for governing work remain separate |
| network partition in Stage B | no inferred terminal state; expire/restrict remote sessions and reconcile | no transitive authority or duplicate effects |

---

# 13. Stage B evolutionary topology

## 13.1 Bounded Stage B hypothesis

Stage B moves sovereign hosting from the workstation to a persistent personal/home-lab node.

```text
Leandro-controlled persistent node
├── Aurora Sovereign Host
├── canonical operational and governed-memory data
├── artifact/evidence coordination
├── approved cognitive/provider runtimes
└── risk-specific effect/credential services

Workstation / mobile / future device
└── Presence adapter/client

External providers/models/Harnesses
└── governed provider boundaries
```

## 13.2 Ownership invariants across migration

The following do not change owner:

- Aurora identity;
- Project/Mission meaning;
- authority and grants;
- governed memory semantics;
- provider trust/approval;
- artifact/evidence identity;
- global Outcome.

## 13.3 New Stage B questions handed forward

Stage B introduces, but this design does not solve:

- service/workload identity;
- Presence authentication and revocation;
- local network transport and discovery;
- egress/network zones;
- synchronization and offline behavior;
- remote update/supervision;
- data minimization before Presence-to-Core transfer;
- node backup/failover posture.

These become inputs to TA-04, TA-05, TA-06 and TA-08.

---

# 14. Process split and merge triggers

## 14.1 Split a logical component when

- a second node/Presence consumes it;
- it requires a different runtime or upgrade cadence;
- it handles untrusted code or broad network access;
- it must hold privileged credentials or device access;
- its resource profile threatens Core availability;
- it needs independent restart/scaling;
- an accepted threat model requires privilege separation;
- a current Spike proves a fault/latency property;
- a real consumer needs independent release/versioning.

## 14.2 Keep or merge in the Sovereign Host when

- the module shares the same sovereign lifecycle and trust zone;
- cross-owner atomicity is currently important;
- no independent scaling/runtime/security need exists;
- process separation would introduce unsupported distributed consistency;
- the component has no current consumer;
- the only rationale is conceptual neatness.

## 14.3 Component-specific hypotheses

| Component | Stage A physical posture | Split trigger |
|---|---|---|
| Identity, Project, Authority | Sovereign Host | only a proven security/scale/topology reason; ownership stays Core |
| Mission/Delegation | Sovereign Host when implemented | durable engine remains adapter; multi-node writers require re-evaluation |
| Registry | Sovereign Host when implemented | remote registry consumers or independent trust service need |
| Memory governance/Context Builder | Sovereign Host | remote/shared memory service only with sovereignty and deletion proof |
| memory extraction/index worker | on-demand provider/worker candidate | M1 evaluation proves separate resource/runtime need |
| Artifact/Evidence metadata | Sovereign Host | payload scale may split content store, not metadata ownership |
| Presence | co-packaged thin adapter initially | second device, independent OS lifecycle, sensors/latency or Stage B |
| Cognitive Coordination | application logic in Host | cognitive execution remains P01 separate |
| Cognitive Runtime | separate on first consumer | merge is not recommended due runtime/ownership boundary |
| Effect Gateway | per-risk adapter; may split | credentials, network, repository, deployment or device privilege |
| Failure Intelligence | deferred logical owner | sustained independent evaluation workloads |
| Observability backend | external adapter/backend | no canonical ownership regardless of placement |

---

# 15. Adversarial flow tests

## 15.1 Continue a Project through a model

```text
Presence submits interaction
→ C08 opens interaction context
→ A02 requests Project snapshot from C02
→ C06 builds attributable Context Pack using C04 policy
→ P01 reasons and returns proposal
→ application routes proposal to C02/C03 as appropriate
→ owner validates and commits or rejects
→ C07 records artifacts/evidence when required
```

**Adversarial assertion:** P01 cannot write the Project store or convert its thread state into current Project state.

## 15.2 Harness requests a child capability

```text
P02 Harness emits child capability request
→ A03 receives through provider contract
→ C03 evaluates relationship to Mission
→ C05 resolves approved candidates
→ C04 establishes narrower authority
→ C06 builds minimized child Context Pack
→ new child Delegation is committed by C03
```

**Adversarial assertion:** parent Harness cannot pass its token, complete context or budget directly to the child.

## 15.3 Memory candidate conflicts with an ADR

```text
P01 or observer emits MemoryCandidate
→ C06 records candidate with provenance
→ accepted ADR remains authoritative source
→ Context Builder preserves conflict/low authority
→ no Project/Decision state change occurs
```

**Adversarial assertion:** repeated or confident synthesis cannot overwrite accepted authority.

## 15.4 Material effect

```text
provider proposes Effect Request
→ C03 links purpose/Delegation
→ C04 evaluates grant/policy/budget context
→ E01 enforces using E02 reference when needed
→ gateway emits Effect Receipt
→ C07 commits receipt/integrity metadata
→ C03/C02 update their state only through explicit reconciled transition
```

**Adversarial assertion:** a provider tool call, model approval or transport success cannot substitute the gateway receipt or owner transition.

## 15.5 Cognitive runtime dies

```text
P01 process dies
→ Sovereign Host remains available
→ current interaction/delegation becomes interrupted/ambiguous
→ provider snapshot is queried when contract supports it
→ resume/new Attempt/failure is explicit
→ no canonical identity, Project, Authority or governed memory is lost
```

## 15.6 Stage A to Stage B migration

```text
export/restore/migrate sovereign state
→ start Sovereign Host on personal node
→ register workstation as Presence
→ establish new transport/auth boundary
→ preserve all domain identities and ownership
```

**Adversarial assertion:** host migration does not create a second Aurora or turn workstation cache into canonical truth.

---

# 16. Approaches comparison

| Criterion | A — Core-centric application | B — early services | C — evolutionary Sovereign Host |
|---|---:|---:|---:|
| Stage A operational simplicity | high | low | high |
| canonical ownership clarity | medium; leakage risk | high if well designed | high |
| cognitive/framework isolation | low | high | high |
| polyglot fit | low/awkward | high | high |
| cross-module transaction simplicity | high | low | high inside Host |
| fault containment | low-medium | high | high at cognitive/Harness boundary |
| premature distributed complexity | low initially | very high | low |
| Stage B evolution | extraction-heavy | direct | direct at accepted seams |
| YAGNI compliance | medium | low | high |
| risk of God process | high | low | medium, controlled by module fitness rules |
| risk of service sprawl | low | high | low |
| first M1/M2 fit | medium | high but expensive | high |

## Recommendation

Adopt **Approach C — Evolutionary Sovereign Host with one early provider-runtime seam** as the TA-01/TA-02 working architecture, subject to operator acceptance.

The recommendation is not “two services because two services are modern.” It is:

```text
one small persistent sovereign deployable
+
one independently replaceable cognitive/provider runtime boundary when consumed
+
all other components remain logically modular until a named split trigger exists
```

---

# 17. Decision disposition register

| ID | Question | Current hypothesis | Disposition | Evidence / owner |
|---|---|---|---|---|
| TA12-D01 | minimum canonical domain-owner set | C01–C10, with C09/C10 future-deferred | `DECIDE` in this design | duplicate-owner review + operator acceptance |
| TA12-D02 | is `Core` one God module? | no; Core is a governed set of owners/coordinators | `DECIDE` | this design + architecture fitness tests |
| TA12-D03 | Stage A base topology | Approach C | `DECIDE` | operator review; no runtime Spike required for conceptual selection |
| TA12-D04 | cognitive execution process boundary | separate from Sovereign Host at first consumer | `DECIDE` structural boundary | ADR-0009 + first-consumer contract/recovery proof |
| TA12-D05 | exact Go ↔ TypeScript binding | unknown | `RESEARCH` in TA-04; `SPIKE` only if recovery/streaming property is consequential | TA-04 ADR/profile |
| TA12-D06 | Presence physical process | co-packaged thin adapter initially | `DEFER` split | reconsider at Stage B/second Presence/OS lifecycle need |
| TA12-D07 | Memory governance placement | Sovereign Host; heavy workers may separate | `DECIDE` ownership, `RESEARCH` M1 mechanisms | TA-05/TA-07 + M1 Spec/evals |
| TA12-D08 | Effect Gateway physical placement | risk/effect-family specific | `DEFER` | TA-06 and first effectful consumer |
| TA12-D09 | Credential Broker product/process | unknown | `DEFER` | TA-06 threat/secret-flow model |
| TA12-D10 | durable engine/service | none for current tranche | `DEFER` | M4 Capability Spec and comparative proof |
| TA12-D11 | Registry physical service | logical module in Host initially | `DEFER` split | M2 consumer and topology evidence |
| TA12-D12 | one physical store versus several | unknown | `DEFER` mechanism | TA-05 logical data/access/failure analysis |
| TA12-D13 | repository strategy | blocked on accepted TA-01/TA-02 | `DEFER` to TA-03 | compare monorepo/polyrepo/staged after approval |
| TA12-D14 | contract schema/codegen | unknown | `RESEARCH` TA-04/TA-03 | canonical contract inventory + current standards |
| TA12-D15 | service supervisor/install model | unknown | `DEFER` TA-08 | first multi-process packaging consumer |
| TA12-D16 | immediate Architecture Spike | none required to choose logical ownership/topology | `DEFER` execution | specify only when a binding/isolation/recovery property cannot be decided documentarily |

---

# 18. Inputs handed to TA-03

If this TA-01/TA-02 proposal is accepted, TA-03 receives these fixed inputs:

1. Aurora has distinct canonical domain modules; source layout must make dependency direction enforceable.
2. `Aurora Sovereign Host` is one Stage A deployable containing multiple owners, not one domain package.
3. Cognitive Runtime is an independently versioned provider-runtime class from its first consumer.
4. Specialized Harnesses are independent provider source/release units, even when first-party.
5. Presence is logically separate but may share the Stage A install/deployable initially.
6. Contract Model source must remain language-neutral and separate from Go/TypeScript generated projections.
7. Canonical stores are accessible only through owner/application code; providers do not share internal database packages.
8. Build/release/versioning must distinguish Core, contract/schema, cognitive provider, AHDK, Harness and adapter versions.
9. Repository strategy must compare monorepo, polyrepo and staged extraction against these deployable/version boundaries.
10. The Development Harness builds/tests/reviews Aurora but does not become an Aurora runtime dependency.

TA-03 remains separately gated and cannot be finalized by accepting this document alone.

---

# 19. Explicit non-decisions

This proposal does not select:

- monorepo, polyrepo or workspace tooling;
- exact package/folder layout beyond ownership direction;
- Go as the universal Aurora language;
- exact Mastra version or implementation;
- HTTP, REST, gRPC, Connect, MCP, A2A, events or local IPC;
- schema language or code generator;
- SQLite, PostgreSQL, object storage, vector or graph stores beyond already scoped M0 decisions;
- Keycloak, Zitadel, Authentik, Ory, SPIFFE, OPA, Cedar, Vault or equivalent;
- service supervisor, containerization or Kubernetes;
- Voice/STT/TTS or model provider;
- durable workflow engine;
- observability backend;
- sandbox mechanism;
- device/laboratory protocol;
- first AHDK language;
- implementation roadmap or production code.

---

# 20. Review criteria

The proposal is ready for acceptance only if an adversarial reviewer can confirm:

1. every representative global concept has one canonical owner;
2. a provider/model/Harness cannot write canonical state directly;
3. the Sovereign Host is not defined as one God module;
4. the recommended Stage A topology is operationally small;
5. the cognitive runtime can fail without erasing sovereign truth;
6. the design does not create one service per module;
7. Stage B changes transport/authentication/topology, not domain identity;
8. scoped M0 decisions are not generalized silently;
9. Mastra remains preferred-first and replaceable rather than sovereign;
10. no hidden repository, protocol, database or authentication product selection exists;
11. all unresolved mechanisms have a treatment, consumer and owner;
12. TA-03 receives sufficient structural inputs without being pre-decided.

---

# 21. Operator decision gate

The proposed decision is:

```text
ACCEPT
→ accept Approach C as the TA-01/TA-02 module/runtime baseline
→ promote the ownership/dependency matrices
→ authorize documentary TA-03 discovery only when separately recorded
→ no implementation follows automatically

REVISE
→ preserve current main and return exact ownership/topology changes

REJECT
→ preserve the accepted Technical Architecture Map
→ replace this proposed TA-01/TA-02 approach with another reviewed design
```

Until operator acceptance:

```text
TA-01/TA-02 proposal: NON-CANONICAL
TA-03 finalization: BLOCKED
Architecture Spike execution: NOT AUTHORIZED
Aurora implementation: PAUSED
```
