---
id: DESIGN-AURORA-CAPABILITY-HARNESS-ARCHITECTURE
title: Aurora Capability and Harness Architecture Design
document_type: system_design
form: explanation
authority: design
status: proposed
version: 0.1.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - proposed capability and harness architecture
related:
  - DOC-AURORA-BLUEPRINT-05
  - DOC-AURORA-BLUEPRINT-07
  - DOC-AURORA-BLUEPRINT-10
  - RESEARCH-AURORA-HARNESS-ARCHITECTURE-V1
  - ADR-AURORA-0001
  - ADR-AURORA-0002
last_reviewed: 2026-08-05
---

# Aurora Capability and Harness Architecture Design

## 1. Status e objetivo

Este design transforma decisões aprovadas na conversa e pesquisa primária em uma proposta revisável.

Objetivo:

> Definir uma arquitetura global em que Aurora coordena sistemas especializados por contracts estáveis, enquanto frameworks, protocols e runtimes permanecem substituíveis.

Não autoriza implementação.

## 2. Problema

Aurora precisará integrar:

- coding harnesses;
- research harnesses;
- hardware/firmware systems;
- laboratory controllers;
- evaluation systems;
- personal operations;
- self-improvement campaigns;
- deterministic services and devices.

Esses sistemas terão runtimes, linguagens, latências, risks e lifecycles diferentes.

Sem boundary comum:

- Aurora se torna monolítica;
- authority se propaga;
- state diverge;
- context vaza;
- recovery falha;
- integrations viram adapters artesanais;
- um framework governa o produto.

## 3. Goals

1. Domain model owned by Aurora.
2. Harness internal autonomy.
3. Hierarchical orchestration.
4. Explicit authority and budgets.
5. Protocol neutrality.
6. First-party Golden Path.
7. Universal conformance.
8. Durable state and recovery.
9. Evidence-oriented outcomes.
10. Distributed observability.
11. Build-bound trust.
12. Local-first operation.

## 4. Non-goals

- escolher stack;
- implementar Core;
- criar marketplace;
- definir hardware do laboratório;
- integrar MNFS agora;
- protocol proprietário completo;
- peer-to-peer authority federation;
- multi-tenancy;
- três SDKs;
- general workflow platform.

## 5. Architecture

```text
Leandro / Presence
        ↓
Aurora Core
├── Identity and Project Model
├── Memory and Context Builder
├── Mission/Delegation Control
├── Capability Registry
├── Authority and Policy
├── Durable Execution Port
├── Artifact/Evidence Coordination
└── Observability
        ↓
Aurora Contract Model
        ↓
Bindings
├── Native AHDK
├── Local RPC
├── A2A
├── MCP
└── HTTP/gRPC/Event
        ↓
Harness Providers
```

Cross-cutting:

- Effect Gateways;
- Credential Broker;
- Sandbox/Environment;
- Artifact Store;
- Evidence Store;
- provenance.

## 6. Component boundaries

### 6.1 Aurora Core

Owns global state, intention, authority, composition and operator interaction.

Must not know internal prompt graphs.

### 6.2 Contract Model

Language-neutral definitions.

Initial entities:

- Capability;
- Provider;
- Instance;
- Manifest;
- Verification;
- Approval;
- Mission;
- Delegation;
- Context Pack;
- Authority Grant;
- Budget;
- Effect Request/Receipt;
- Event;
- Decision Request;
- Artifact;
- Evidence;
- Outcome;
- Checkpoint.

### 6.3 Capability Registry

Stores exact provider/build identity, contracts, verification, approval scope, incidents and status.

### 6.4 AHDK

Implements first-party Golden Path. Generates types and hides repetitive transport mechanics, not domain semantics.

### 6.5 Conformance Kit

External black-box authority for protocol behavior.

### 6.6 Durable Execution Port

Abstracts scheduling, checkpoints, timers, resume and retries. Does not own product semantics.

### 6.7 Policy Decision Point

Evaluates principal/actor/action/resource/context.

### 6.8 Effect Gateway

Applies decision at point of effect and issues receipt.

### 6.9 Artifact/Evidence coordination

Persists deliverables and links evidence to criteria.

## 7. Data flow

### 7.1 Discovery

```text
Provider publishes manifest
→ Registry ingests
→ schema and signature checks
→ status DISCOVERED
```

No execution.

### 7.2 Verification

```text
Inspect
→ sandbox
→ conformance
→ security/operational evidence
→ trust assessment
→ operator approval for scope
```

### 7.3 Delegation

```text
Leandro intent
→ Aurora resolves project/context
→ capability selection
→ provider selection
→ Delegation Envelope
→ policy decision
→ durable start
→ provider execution
```

### 7.4 Collaboration

Harness requests another capability. Aurora creates child delegation. Data plane may be direct through an authorized channel.

### 7.5 Completion

```text
provider publishes artifacts/claims
→ verification produces receipts/evidence
→ Aurora evaluates criteria
→ outcome
→ memory/project update
→ operator summary
```

## 8. State ownership

### Global state

Aurora:

- mission;
- delegations;
- grants;
- budgets;
- approvals;
- relationships;
- acceptance;
- consolidated events.

### Local state

Harness:

- internal plan;
- workers;
- local attempts;
- intermediate artifacts;
- domain-specific checkpoints.

### Reconciliation

Provider exposes observable snapshot and checkpoint identity. Aurora never infers terminal state only from missing connection.

## 9. Lifecycle

Proposed common lifecycle:

```text
PROPOSED
AUTHORIZED
QUEUED
RUNNING
WAITING_FOR_INPUT
BLOCKED
CANCEL_REQUESTED
COMPLETED
FAILED
CANCELED
REJECTED
```

Requirements:

- state machine validation;
- transition reason;
- timestamp;
- actor;
- attempt/run identity;
- terminal invariants;
- recoverable/non-recoverable error class.

## 10. Error model

Classes:

- contract violation;
- compatibility failure;
- unauthorized effect;
- budget exceeded;
- provider unavailable;
- transient dependency;
- unrecoverable execution;
- invalid evidence;
- stale context;
- lost heartbeat;
- cancellation failure;
- reconciliation conflict;
- security incident.

No blind retry. Retry policy depends on error class and idempotency.

## 11. Recovery

### Core restart

Global state reloads and reconciles active providers.

### Harness restart

Provider resumes from checkpoint or reports failure class.

### Network partition

Aurora preserves last verified state, avoids duplicate delegation and reconciles later.

### Ambiguous effect

Effect Gateway uses idempotency key and receipt lookup. Ambiguity blocks advancement until resolved.

### Version change

Running delegation stays bound to exact contract/provider version or follows explicit migration.

## 12. Security

- grants are scoped and expiring;
- provider cannot mint new authority;
- child delegation returns to Aurora;
- credentials are brokered;
- effects pass gateways;
- devices have own trust;
- external providers receive minimized context;
- manifests do not grant trust;
- revocation has operational effect;
- sandbox complements policy.

## 13. Protocol mapping

### Native AHDK

Preferred first-party local integration.

### A2A

Candidate for remote opaque tasks.

### MCP

Candidate for tools/resources and bounded async calls.

### Direct channel

Authorized data plane for volume/latency.

Mapping is adapter responsibility. Domain remains canonical.

## 14. Testing strategy

### Contract tests

- schemas;
- compatibility;
- state transitions;
- invariants.

### Conformance tests

- SDK and direct implementation;
- protocol/TCK;
- error behavior.

### Fault tests

- kill/restart;
- duplicate delivery;
- lost messages;
- delayed events;
- stale manifest;
- revoked token;
- partial artifact.

### Security tests

- SDK bypass attempt;
- denied effect;
- context exfiltration;
- privilege propagation;
- compromised provider version.

### Observability tests

- complete trace;
- redaction;
- budget attribution;
- correlation.

### Acceptance proof

Reference harness A produces artifact. Reference harness B consumes authorized artifact. Aurora:

1. discovers;
2. verifies;
3. approves;
4. delegates;
5. survives restart;
6. mediates child work;
7. preserves authority boundaries;
8. collects evidence;
9. completes outcome.

No MNFS dependency.

## 15. Rollout

### Phase 1

Contracts and conformance prototype.

### Phase 2

One reference SDK and two trivial providers.

### Phase 3

Durability and authority gateways.

### Phase 4

A2A/MCP adapters.

### Phase 5

Real domain harness integration.

Exact milestones require implementation plan after approval.

## 16. Open technical decisions

These are deliberately undecided:

- Core language;
- first AHDK language;
- schema formats per boundary;
- RPC binding;
- durable engine;
- PDP;
- identity framework;
- storage;
- broker;
- artifact store;
- reference runtimes.

Each requires spike or evidence.

## 17. Acceptance of this design

Design is ready for planning only when Leandro confirms:

- boundary Aurora/harness;
- first-party AHDK policy;
- universal conformance;
- hierarchical orchestration;
- separate durability;
- separate policy/enforcement;
- protocol mapping direction;
- spike portfolio;
- non-goals.

Until then, status remains `proposed`.
