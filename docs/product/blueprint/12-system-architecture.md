---
id: DOC-AURORA-BLUEPRINT-12
title: Arquitetura do Sistema e Fronteiras dos Componentes
document_type: product_blueprint_section
form: explanation
authority: constitutional
status: accepted
accepted_at: 2026-08-06
acceptance_evidence: DOC-AURORA-A0-OPERATOR-ACCEPTANCE
version: 0.2.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - logical system architecture
  - component ownership and dependency boundaries
  - control-plane and data-plane model
  - local/cloud evolution principles
related:
  - DOC-AURORA-BLUEPRINT-03
  - DOC-AURORA-BLUEPRINT-05
  - DOC-AURORA-BLUEPRINT-06
  - DOC-AURORA-BLUEPRINT-07
  - DOC-AURORA-BLUEPRINT-08
  - DOC-AURORA-BLUEPRINT-09
  - DOC-AURORA-BLUEPRINT-10
  - DOC-AURORA-BLUEPRINT-11
  - DOC-AURORA-BLUEPRINT-13
  - ADR-AURORA-0001
  - ADR-AURORA-0002
review_triggers:
  - Core component ownership changes
  - authority or state ownership changes
  - new mandatory distributed boundary
  - local-first topology changes
last_reviewed: 2026-08-06
---

# 12. Arquitetura do Sistema e Fronteiras dos Componentes

## 12.1 Propósito

Aurora will combine long-lived personal state, real-time interaction, agent reasoning, durable workflows, external models, specialized harnesses, device telemetry and controlled physical effects.

Without explicit boundaries, the system can collapse into one of two failures:

### Monolithic superagent

One prompt or agent tries to remember, plan, call every tool, manage every device and judge its own work.

### Accidental distributed platform

Many services and frameworks own overlapping state, authority and semantics, making recovery and trust impossible.

The architecture must establish:

- which concepts belong to Aurora Core;
- which concerns are delegated to harnesses;
- how identity, memory, authority and state remain canonical;
- how protocols and frameworks remain replaceable;
- how local operation remains practical;
- how cloud and multiple devices can be added without rewriting the domain;
- where security and safety are enforced;
- how observations, artifacts and evidence flow;
- how failures are contained and reconciled.

This section defines **logical architecture**, not a final process topology or stack.

---

## 12.2 Architectural thesis

> Aurora is a sovereign cognitive control plane with modular domain components, replaceable inference and execution runtimes, contract-governed capability providers and independently enforced effect boundaries.

The initial implementation should prefer the smallest topology that proves the required boundaries.

The target architecture must not require a microservice deployment from the beginning.

---

## 12.3 Global view

```text
┌───────────────────────────────────────────────────────────────┐
│                         LEANDRO                               │
└──────────────────────────────┬────────────────────────────────┘
                               │
                 text • voice • visual • events
                               │
┌──────────────────────────────▼────────────────────────────────┐
│                      PRESENCE FABRIC                          │
│ auth • device context • channel • privacy • handoff           │
└──────────────────────────────┬────────────────────────────────┘
                               │
┌──────────────────────────────▼────────────────────────────────┐
│                         AURORA CORE                           │
│                                                               │
│ Identity & Relationship    Project / World Model              │
│ Memory & Context Builder   Intent / Mission Control           │
│ Capability Registry        Authority & Policy Coordination    │
│ Durable Execution Port     Artifact / Evidence Coordination   │
│ Proactivity & Attention    Failure Intelligence               │
│ Observability & Evaluation Operator Interaction               │
└──────────────────────────────┬────────────────────────────────┘
                               │
                    Aurora Contract Model
                               │
┌──────────────────────────────▼────────────────────────────────┐
│                     CAPABILITY FABRIC                         │
│ Native AHDK • Local RPC • A2A • MCP • HTTP/gRPC • Events      │
└──────────────┬───────────────┬───────────────┬────────────────┘
               │               │               │
       ┌───────▼───────┐ ┌─────▼──────┐ ┌──────▼─────────┐
       │ Research      │ │ Software   │ │ Hardware/FW/Lab│
       │ Harness       │ │ MNFS       │ │ Harnesses      │
       └───────────────┘ └────────────┘ └────────────────┘
                               │
┌──────────────────────────────▼────────────────────────────────┐
│                 EFFECT AND DEVICE PLANE                       │
│ filesystem • network • credentials • repos • deploy • devices │
│ PDP • gateways • broker • sandbox • interlocks • receipts     │
└───────────────────────────────────────────────────────────────┘
```

Cross-cutting stores and services:

```text
Operational State
Memory/Knowledge Stores
Artifact Store
Evidence Store
Event/Audit Store
Telemetry Backend
Backup/Recovery
Build Provenance
```

---

## 12.4 Architectural principles

### 12.4.1 Domain inward, adapters outward

The canonical domain must not import:

- provider SDK semantics;
- MCP/A2A internals;
- database-specific models;
- workflow-engine state;
- device protocol details;
- UI framework types;
- a model vendor's message format.

Adapters translate external mechanisms into Aurora contracts.

### 12.4.2 One authority per concept

Examples:

```text
constitutional identity  → Product Blueprint / accepted changes
current mission state    → Aurora operational state
harness internal plan    → harness local state
source code              → Git
provider approval        → Capability Registry
external effect result   → Effect Receipt + target system
live instrument reading  → device/telemetry source
memory synthesis         → Memory subsystem, below source authority
```

### 12.4.3 Logical modularity before physical distribution

Components should have clear ports and ownership even if they run in one process initially.

A component becomes physically separate only with evidence such as:

- independent lifecycle;
- isolation requirement;
- device boundary;
- remote consumer;
- scale/latency need;
- language/runtime constraint;
- fault containment;
- security zone.

### 12.4.4 State is durable, processes are disposable

Aurora must survive:

- UI close;
- model change;
- process crash;
- provider restart;
- network partition;
- presence handoff.

### 12.4.5 Large data moves by reference or governed channel

Prompts and control messages should not carry full traces, videos, repository trees or high-rate telemetry.

### 12.4.6 Security enforcement stays outside probabilistic control

Model, SDK and agent framework cannot be the only barrier to an effect.

### 12.4.7 Local-first without local-only coupling

Aurora should function meaningfully on infrastructure controlled by Leandro while permitting external models and remote providers through policy.

### 12.4.8 Protocols transport; contracts define

A2A, MCP, RPC and events are mappings, not product ontology.

---

## 12.5 Aurora Core modules

### 12.5.1 Identity and Relationship

Owns:

- Aurora constitutional identity reference;
- Leandro identity/profile references;
- relationship and personality mode;
- authenticated actor context;
- protected identity settings.

Does not own:

- raw biometric engine;
- device OS identity;
- model personality defaults.

### 12.5.2 Project and World Model

Owns global relationships and project-scoped views across:

- projects;
- goals;
- decisions;
- missions;
- devices;
- knowledge;
- experiments;
- incidents.

It references authoritative sources rather than duplicating every content item.

### 12.5.3 Memory and Context Builder

Owns:

- memory write/manage/read lifecycle;
- memory scopes and metadata;
- promotion/supersession/retention;
- retrieval and context compilation;
- authority/freshness resolution;
- token/data minimization;
- context provenance.

Does not own project decisions or live device state.

### 12.5.4 Intent and Mission Control

Owns:

- interpreted user intent;
- mission lifecycle;
- global decomposition;
- delegation relationships;
- dependencies;
- global acceptance;
- decision/escalation routing;
- consolidated outcome.

Harnesses own local execution plans.

### 12.5.5 Capability Registry

Owns:

- capability/provider/instance identity;
- manifests;
- compatibility;
- verification results;
- approvals;
- trust dimensions;
- suspension/revocation;
- build provenance;
- operational availability.

### 12.5.6 Authority and Policy Coordination

Owns Aurora-level grants, requested effects and policy context.

May integrate a policy engine, but domain concepts remain Aurora-owned.

### 12.5.7 Durable Execution Port

Exposes domain-oriented operations such as:

```text
start durable mission/delegation
schedule timer
persist checkpoint reference
await signal/decision
retry eligible step
cancel
resume
query status
```

The adapter may use an engine or a minimal local implementation.

Durable engine history does not become the Product Domain automatically.

### 12.5.8 Artifact and Evidence Coordination

Owns:

- artifact identities and metadata;
- evidence-to-criterion links;
- integrity references;
- promotion/retention;
- outcome composition.

Blob storage and databases are adapters.

### 12.5.9 Proactivity and Attention

Owns:

- candidate notification;
- urgency/relevance/confidence;
- attention budget;
- delivery presence;
- suppression/deduplication;
- explanation of why now.

### 12.5.10 Failure Intelligence

Owns:

- incident/finding correlation;
- causal hypotheses;
- improvement opportunities;
- candidate lifecycle;
- evaluation links;
- promotion proposals.

### 12.5.11 Operator Interaction

Converts domain decisions/status to user-facing communication. UI, voice and device surfaces remain thin adapters.

---

## 12.6 Aurora Contract Model

The Contract Model is a language-neutral canonical definition of boundaries.

Initial entity families:

### Identity

- PersonRef;
- AuroraRef;
- PresenceRef;
- DeviceRef;
- ProviderRef;
- ActorChain.

### Capability

- CapabilityDefinition;
- ProviderManifest;
- CompatibilityProfile;
- Verification;
- Approval.

### Mission

- Mission;
- Delegation;
- ContextPack;
- AuthorityGrant;
- Budget;
- Guardrail;
- DecisionRequest.

### Execution

- Run/Attempt;
- Checkpoint;
- Heartbeat;
- Cancellation;
- Error;
- RecoverySnapshot.

### Effects

- EffectRequest;
- PolicyDecision;
- CredentialReference;
- EffectReceipt.

### Results

- Artifact;
- Claim;
- Observation;
- Evidence;
- Verdict;
- Outcome.

### Events

- material lifecycle events;
- audit correlation;
- privacy classification;
- trace context.

The Contract Model requires versioning and compatibility independent of AHDK implementation.

---

## 12.7 Capability Fabric

The Capability Fabric coordinates communication with providers.

Responsibilities:

- provider discovery/connection;
- manifest retrieval;
- binding negotiation;
- delegation dispatch;
- event/status ingestion;
- artifact/evidence references;
- cancellation/signals;
- child capability requests;
- context and authority delivery;
- health and reconciliation.

Bindings:

```text
Native in-process AHDK
Local RPC
A2A
MCP
HTTP/gRPC
Event transport
Device-specific adapter
```

A provider can support more than one binding. Selection depends on topology, maturity, performance and security.

---

## 12.8 Aurora Harness Development Kit

AHDK is the paved path for first-party harnesses.

Logical modules:

```text
contracts-generated
manifest-builder
provider-runtime
lifecycle-client
context-reader
authority-effect-client
artifact-evidence-client
checkpoint-cancellation
telemetry
simulator-testkit
conformance-runner
scaffolder
provenance-hooks
```

### Dependency direction

```text
Harness domain code
        ↓
AHDK public API
        ↓
Contract types and binding adapter
```

Harness code should not depend on Aurora Core internals.

### Waiver

Direct implementation remains technically possible for external/legacy systems. First-party waiver requires explicit rationale and complete conformance.

---

## 12.9 Conformance architecture

The Conformance Kit is external to provider self-assessment.

Tests:

- schema and manifest;
- lifecycle/transitions;
- version negotiation;
- errors;
- cancellation;
- deadline;
- resume/checkpoint;
- idempotency;
- event ordering/duplicates;
- artifact integrity;
- authority denial;
- budget thresholds;
- effect receipts;
- restart/reconciliation;
- telemetry propagation;
- privacy/redaction.

A provider using AHDK can still fail conformance due to misuse or SDK defect.

---

## 12.10 Operational State

Aurora needs a canonical operational state that is distinct from:

- conversation history;
- durable engine event history;
- harness local database;
- Git documents;
- logs.

It contains current projections for:

- projects;
- missions;
- delegations;
- grants;
- budgets;
- provider approvals/status;
- presence/device status references;
- checkpoints;
- incidents;
- pending decisions;
- artifact/evidence metadata.

The storage technology remains open. Initial topology should favor simplicity and recoverability.

---

## 12.11 Event and audit architecture

Material events support:

- state transitions;
- audit;
- notification;
- reconciliation;
- observability;
- historical analysis.

The system distinguishes:

### Domain Event

A meaningful state change.

### Transport Message

Delivery unit that may duplicate or fail.

### Telemetry Event

Operational measurement for monitoring.

### Audit Record

Security/authority/data/effect accountability.

### Artifact

Potentially large durable output referenced by event.

CloudEvents is a candidate envelope; exact event store/broker remains open.

Critical information must be recoverable from state/artifact even if a stream subscriber misses an event.

---

## 12.12 Control plane and data plane

### Control plane

Governed by Aurora:

- identity;
- mission/delegation;
- contracts;
- authority;
- budget;
- channels;
- decision;
- lifecycle;
- outcome.

### Data plane

May transfer:

- telemetry;
- files;
- traces;
- datasets;
- images/video;
- repository content;
- model artifacts.

Direct channels can bypass Core as a data path but not as an authority path.

Channel contract includes:

- producer/consumer;
- schema;
- scope;
- data classification;
- rate/size;
- encryption;
- duration;
- storage/retention;
- budget;
- revocation;
- audit.

---

## 12.13 Memory and knowledge topology

Logical strata:

```text
L0 authoritative sources/state
L1 current authority/project snapshots
L2 governed memory stores
L3 observational/session synthesis
L4 exact conversation/tool history
L5 ephemeral context and transport
```

These strata may use different physical stores.

Context Builder reads through ports:

- project source;
- decision source;
- memory source;
- knowledge retrieval;
- live state source;
- conversation history;
- provider policy.

No vector database becomes the owner of project truth merely because it serves retrieval.

---

## 12.14 Artifact and Evidence stores

### Artifact Store

Optimized for content:

- documents;
- code patches;
- binaries;
- telemetry files;
- waveforms;
- images/video;
- model outputs;
- reports.

### Evidence Store

Optimized for relationships:

- evidence ID;
- criterion/hypothesis;
- artifact references;
- producer;
- verifier;
- environment;
- method;
- result;
- limitations;
- verdict.

A single physical storage may implement both initially while logical distinction remains.

---

## 12.15 Authority and Effect architecture

```text
Aurora creates Authority Grant
        ↓
Provider receives scoped delegation/token/reference
        ↓
Provider requests effect
        ↓
Policy Decision Point evaluates
        ↓
Effect Gateway enforces
        ↓
Credential Broker resolves minimum credential
        ↓
Target system/device
        ↓
Effect Receipt + observation + audit
```

Gateway families may include:

- filesystem;
- network;
- repository;
- external communication;
- deployment;
- financial;
- credential;
- device/laboratory.

No generic “execute anything” gateway should become the default.

---

## 12.16 Device plane

Device Plane components:

- Device Registry;
- controllers/adapters;
- telemetry ingestion;
- command gateways;
- calibration records;
- leases;
- interlock status;
- firmware/artifact identity;
- laboratory protocol executor.

It remains separated from general model context and external providers by explicit contracts.

---

## 12.17 Presence architecture

Presence Fabric adapters may run on:

- workstation;
- mobile;
- wearable;
- laboratory display;
- custom embedded device.

They communicate with Core through a presence contract carrying:

- authenticated actor;
- device trust;
- channel capabilities;
- environment;
- privacy mode;
- active activity;
- minimal context;
- effective authority.

A presence may host local low-latency functions while canonical state remains in Core.

---

## 12.18 Model and inference architecture

Models are selected capabilities/internal runtimes, not identity.

Possible roles:

- conversation;
- planning;
- coding;
- vision;
- speech;
- embeddings/retrieval;
- review;
- summarization;
- anomaly analysis.

Model Router considerations:

- task fit;
- data sensitivity;
- provider policy;
- cost;
- latency;
- context size;
- tool support;
- reliability/eval history;
- local availability;
- fallback.

Model output remains probabilistic and attributable to provider/version.

---

## 12.19 Initial topology hypothesis

Without selecting language or store, the initial deployable shape should likely be a **modular local Core** with separable adapters.

```text
One local Aurora Core process/service
+ one local operational store
+ local artifact directory/store
+ provider adapters
+ one presentation interface
+ isolated reference harness processes
```

This is a hypothesis guided by YAGNI, not an accepted implementation design.

Do not begin with:

- mandatory Kubernetes;
- many microservices;
- broker cluster;
- remote multi-tenant control plane;
- full device mesh;
- global high-availability system.

Architecture spikes decide the minimum practical topology.

---

## 12.20 Evolution topologies

### Stage A — Single machine

```text
workstation/server controlled by Leandro
Core + store + UI + local harness adapters
external models through governed calls
```

### Stage B — Home/lab node

```text
persistent local server
workstation/mobile presences
laboratory device node
remote model providers
```

### Stage C — Distributed personal system

```text
sovereign Core cluster or primary/backup
remote durable workers
multiple presences
specialized compute nodes
cloud-assisted capabilities
```

### Stage D — Reusable platform components

Selected capabilities/AHDK may be reused outside Leandro's instance without introducing current multi-tenancy into the Core.

Each stage requires explicit evidence; the domain should remain stable while adapters/topology evolve.

---

## 12.21 Failure domains

### Core process failure

Recover operational state and reconcile active providers.

### Operational store failure

Backup/restore and integrity verification; prevent stale grants from reactivating.

### Provider failure

Resume, substitute or block according to contract and evidence.

### Durable engine failure

Recover event history/checkpoints; Aurora domain state remains reconcilable.

### Network partition

Do not infer terminal state; prevent duplicate effects; reconcile later.

### Presence failure

Mission continues; local context is revoked/expired.

### Artifact store failure

Outcome cannot close when required evidence is unavailable.

### Policy/gateway failure

Fail closed for material effects; allow explicitly safe read-only/degraded functions where designed.

### Device controller failure

Enter safe state, preserve ambiguity, use local interlock.

### Model/provider outage

Use approved fallback or declare capability unavailable; identity/state remain.

---

## 12.22 Versioning and migration

Versioned surfaces:

- Product Blueprint;
- Contract Model;
- AHDK;
- bindings;
- provider manifest;
- provider build;
- schemas;
- policy;
- operational store migrations;
- memory representation;
- device manifest;
- protocol;
- artifacts.

Running Delegation remains bound to exact versions unless migration is explicit.

Migration needs:

- compatibility analysis;
- transform;
- backup;
- validation;
- rollback;
- evidence;
- authority.

---

## 12.23 Observability integration

Every component should propagate correlation for:

```text
project
mission
delegation
provider instance
run/attempt
effect
artifact/evidence
decision request
trace
```

OpenTelemetry is the current baseline hypothesis. Domain events and evidence remain separate from telemetry signals.

Sensitive context must be redacted from general telemetry.

---

## 12.24 Testing architecture

### Unit/domain

Invariants, state machines, authority composition, context rules.

### Contract

Schema and compatibility.

### Conformance

Provider black-box behavior.

### Integration

Core/adapters/stores/gateways.

### Fault/recovery

Kill, restart, duplicates, partition, stale state, ambiguous effect.

### Security

Injection, SDK bypass, revoked grant, provider substitution, context leakage.

### Journey

North Star laboratory, campaign, handoff and multi-harness composition.

### Physical

Simulation, HIL, interlock and controlled bench drills.

---

## 12.25 Architecture fitness questions

Before accepting material change:

1. Does it create a second authority for a concept?
2. Does Domain import a provider/framework type?
3. Can the component restart without losing authoritative state?
4. Is authority enforced at the effect?
5. Can a provider/harness be substituted?
6. Is context minimized and attributable?
7. Does the change require distribution now?
8. Does it introduce a bypass around audit or interlock?
9. Can a fresh session understand the boundary?
10. Is there a Golden Proof demonstrating the need?

---

## 12.26 Open technical decisions

- Core language/runtime;
- modular monolith packaging;
- operational database;
- event/audit store;
- memory storage architecture;
- artifact storage;
- durable engine;
- policy engine;
- identity/credential mechanism;
- local RPC;
- schema formats by boundary;
- message/event transport;
- telemetry backend;
- deployment/service manager;
- backup topology;
- device protocols;
- presence transport;
- model routing implementation.

Each requires focused research or architecture spike.

---

## 12.27 Non-goals

This section does not:

- mandate microservices;
- select a framework;
- make all components network services;
- make Aurora a generic workflow platform;
- require every harness to use the same language;
- define a cloud SaaS;
- assume high availability from the start;
- move physical safety into the Core model loop;
- make event history the only state store;
- make a vector store the world model;
- authorize implementation before A0 acceptance.
