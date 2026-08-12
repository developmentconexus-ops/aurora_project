---
id: DESIGN-AURORA-TECHNICAL-ARCHITECTURE-BASELINE-MAP
title: Aurora Technical Architecture Baseline Map
document_type: system_design
form: reference
authority: design
status: accepted
accepted_at: 2026-08-12
acceptance_evidence: DOC-AURORA-TECHNICAL-ARCHITECTURE-MAP-OPERATOR-ACCEPTANCE
version: 0.1.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - ordered Aurora technical architecture baseline work map
  - technical-architecture materiality and decision sequence
  - required outputs and stop conditions for each architecture area
  - treatment of accepted Presence details that are not current architecture priorities
  - current first technical architecture tranche
related:
  - DOC-AURORA-STATUS
  - DOC-AURORA-DECISIONS
  - DOC-AURORA-WORKLOG
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DESIGN-AURORA-SYSTEM-ARCHITECTURE-REBASELINE
  - DESIGN-AURORA-SYSTEM-ARCHITECTURE-DECISION-LANDSCAPE
  - DESIGN-AURORA-SAR-A1-STAGE-A-AVAILABILITY-ACTIVATION
  - DOC-AURORA-TECHNICAL-ARCHITECTURE-MAP-OPERATOR-ACCEPTANCE
  - PLAN-AURORA-TECHNICAL-ARCHITECTURE-BASELINE
source_revision: a520df862ad39d8e17e8bd17a80da8b8b2f1a900
review_triggers:
  - a work package discovers a material cross-system contradiction
  - a new Product Milestone becomes the executable horizon
  - an accepted technical decision changes module ownership or runtime topology
  - implementation cannot proceed without inventing an unowned cross-system boundary
  - the map creates architecture work that has no near consumer or stop condition
last_reviewed: 2026-08-12
---

# Aurora Technical Architecture Baseline Map

## 1. Purpose

Aurora already has a strong accepted product constitution, logical architecture, domain vocabulary, authority model, memory principles, Harness boundary, AHDK direction, roadmap and ACRM realization method.

The current program need is different:

> Convert that accepted product architecture into a coherent technical map that tells future implementers what components exist, who owns each responsibility and datum, how components communicate, which runtime and repository boundaries apply, and where technology choices belong.

This map prevents two opposite failures:

```text
code-first fragmentation
→ each capability invents its own modules, APIs, identity, storage and runtime assumptions

architecture without delivery
→ the program researches distant mechanisms indefinitely and never reaches an executable baseline
```

The map therefore organizes technical architecture work in a strict dependency order and applies a materiality rule to every question.

It does not authorize Aurora implementation.

---

## 2. Direction correction

SAR-A1 initially entered useful but increasingly specific Presence/session-policy questions, including activation while a workstation is locked. Those decisions remain valid within their stated Stage A Presence scope.

The operator then identified a priority error:

```text
useful future Presence detail
≠ current cross-system technical architecture blocker
```

Questions such as who may use an already-unlocked workstation can matter later to a Presence, Voice or authentication threat model. They do not currently decide:

- global module ownership;
- process/runtime topology;
- repository boundaries;
- inter-component contracts;
- data architecture;
- identity/service trust architecture;
- Brain/model/Harness integration;
- deployment and observability foundations.

Further Presence micro-policy exploration is therefore `DEFER` until a consuming Capability reaches the appropriate ACRM gate.

The accepted Stage A constraints are preserved as downstream inputs, not expanded as the present workstream.

---

## 3. Existing accepted foundation

The technical baseline must preserve the following accepted product meaning and scoped decisions.

### 3.1 Product and ownership constraints

- Aurora is one persistent personal intelligence and the sovereign cognitive/operational control plane.
- Leandro retains final authority over goals, values, material decisions and grants.
- Aurora owns global Project/Mission meaning, authority, budgets, governed context, provider trust and global Outcome.
- Harnesses own specialized local methodology, workers, attempts and provider-local state.
- Aurora owns Contract Model semantics; protocols, SDKs and runtimes are bindings.
- AHDK is the first-party Golden Path by policy but is not the specification or security boundary.
- memory, knowledge, history, operational state, active context, telemetry and evidence are distinct.
- material effects require deterministic enforcement outside probabilistic model judgment.
- local-first/cloud-assisted sovereignty is required.
- one durable concept has one canonical owner.

### 3.2 Scoped technical decisions already accepted

- Go is selected for the M0 Sovereign Core, not every Aurora component.
- M0 uses one local modular Core and explicit current state/revisions.
- SQLite is the M0 operational-state baseline, not Aurora's universal database.
- JSON Schema/JSON/JCS and protected logical export are M0 portability decisions.
- OTel/slog is the M0 observability boundary; no universal backend is selected.
- Mastra is preferred-first to evaluate for first-party agentic Harness/cognitive runtimes, never as global state or authority owner.

No scoped M0 mechanism is generalized by implication.

---

## 4. Technical architecture materiality rule

A question belongs in the current technical architecture work only when its answer changes at least one of:

1. a component or module boundary;
2. canonical responsibility or data ownership;
3. an allowed or forbidden dependency;
4. a process/runtime/deployment boundary;
5. a contract, API, event or compatibility rule;
6. a security, authority, credential or effect-enforcement boundary;
7. storage role, portability, recovery or deletion ownership;
8. the next evidence-supported implementation decision.

Otherwise:

```text
record consumer and trigger
→ classify DEFER
→ continue the current architecture work
```

Examples of valid current questions:

- Is the Cognitive Runtime part of the Sovereign Core process or a replaceable provider process?
- Which module owns Project current state?
- Can a Harness write governed memory directly?
- Which contract crosses the Go ↔ TypeScript boundary?
- Which data classes are canonical versus rebuildable indexes?

Examples to defer until a real consumer:

- exact wake-word model;
- which family member may ask a public question while a workstation is unlocked;
- exact wearable handoff behavior;
- laboratory voice confirmation wording;
- TTS voice selection;
- Kubernetes cluster layout.

---

## 5. Decision workflow for every architecture area

Technical choices must follow this sequence:

```text
1. boundary and responsibility
2. technical requirements and invariants
3. first real consumer
4. viable approaches
5. current primary-source research when material
6. trade-offs, migration and operational cost
7. DECIDE | RESEARCH | SPIKE | DEFER disposition
8. owning ADR / Specification / Standard / Contract
9. resulting contracts and implementation impact
```

Technology names may appear as candidates during comparison. They do not become decisions until promoted by the correct owner.

Examples:

```text
identity model and topology
→ authentication requirements
→ compare Keycloak / Zitadel / Ory / Authentik / SPIFFE classes
→ decide per actor class if required
```

```text
data categories and access patterns
→ store roles
→ compare SQLite / PostgreSQL / object storage / vector or graph mechanisms
→ decide per consuming boundary
```

```text
communication boundary
→ latency, streaming, recovery and compatibility requirements
→ compare in-process / IPC / HTTP / Connect / gRPC / events / MCP / A2A
→ select only the binding needed by that boundary
```

---

# 6. Ordered technical architecture work map

The eight areas below are dependency-ordered. Later areas may be discussed provisionally, but they must not finalize assumptions owned by an earlier area.

## TA-01 — Logical modules and canonical ownership

### Question

What are Aurora's principal technical modules, what does each own, and which dependencies are allowed or forbidden?

### Required outputs

- component/module catalog;
- responsibility matrix;
- canonical entity/data owner matrix;
- allowed dependency direction;
- forbidden cross-module writes;
- module-owned ports and events;
- separation between domain modules and mechanism adapters;
- mapping to accepted Blueprint concepts and M0 evidence.

### Minimum candidate modules to test—not pre-accept

```text
Sovereign Core foundation
Identity and Relationship
Project / World State
Mission / Delegation Control
Authority and Effect Coordination
Memory and Context
Cognitive Runtime Coordination
Capability / Provider Registry
Harness Integration
Artifact and Evidence
Presence Coordination
Proactivity / Attention
Failure Intelligence
Observability Coordination
```

The work must merge, split or defer these based on responsibility cohesion rather than preserve the list mechanically.

### Stop condition

A reviewer can answer for every current global concept:

```text
who owns it
who may read it
who may propose change
who may commit change
what cannot own it
```

No repository or framework decision is final before this output exists.

---

## TA-02 — Process, runtime and evolutionary topology

### Question

Which logical modules share a process initially, which require independent failure/security/runtime boundaries, and how does Stage A evolve to Stage B?

### Required outputs

- Stage A process/runtime diagram;
- always-active versus on-demand responsibilities;
- same-process versus separate-process criteria;
- Go, TypeScript/Mastra and future runtime scope constraints;
- startup, shutdown, restart and supervision ownership;
- local failure domains;
- migration path from workstation-sovereign Stage A to persistent-node Stage B;
- explicit list of boundaries that stay logical until evidence justifies physical separation.

### Accepted Stage A input

```text
one Leandro-controlled workstation
→ initial sovereign host + first Presence

hybrid availability
→ minimum Core + activation responsibility available
→ heavy cognition starts on demand
```

This input constrains topology but does not select one-process versus multi-process packaging, IPC or a service manager.

### Stop condition

Every TA-01 module has a Stage A runtime placement and Stage B evolution hypothesis, with a reason for each physical boundary.

---

## TA-03 — Repository, source and build architecture

### Question

Which source units version, test, release and deploy together, and how are contracts shared without coupling all components?

### Required outputs

- monorepo versus polyrepo decision or staged strategy;
- repository/component ownership map;
- package/module layout;
- dependency direction enforcement;
- contract/schema source location;
- generated-type policy;
- language/workspace boundaries;
- CI test layers;
- release/versioning model;
- development-Harness relationship;
- branch/PR and supply-chain baseline.

### Technology candidates considered only here or later

- Go workspace/module structure;
- TypeScript workspace tooling;
- schema/code generation;
- build orchestration;
- dependency graph enforcement;
- release automation.

### Stop condition

A new implementer can place a module, contract, test and deployment artifact without inventing repository structure.

---

## TA-04 — Contracts, APIs, events and communication

### Question

How do modules, processes, Presences, models and Harnesses communicate while preserving Aurora-owned semantics and compatibility?

### Required outputs

- contract family catalog;
- in-process versus inter-process versus external boundary profiles;
- synchronous, streaming and asynchronous interaction rules;
- error taxonomy;
- deadlines, cancellation and retry semantics;
- idempotency and ambiguous-result rules;
- correlation and causation identity;
- schema/binding/SDK/provider version separation;
- event versus message versus audit versus telemetry taxonomy;
- compatibility and deprecation policy.

### Candidate mechanisms after requirements

```text
in-process ports
local IPC
HTTP/REST
Connect/gRPC
WebSocket/SSE
message/event transport
MCP
A2A
ACP
```

There is no requirement for one universal protocol.

### Stop condition

Every TA-02 process crossing has a contract class and failure/compatibility semantics, even when the concrete binding is deferred.

---

## TA-05 — Data, storage, portability and lifecycle architecture

### Question

What data classes exist, who owns them, what guarantees do they need and which physical store roles are justified?

### Required data families

```text
canonical operational state
governed durable memory
knowledge sources and documents
conversation / interaction history
artifacts
evidence and receipts
audit records
telemetry
vector/full-text/graph derived indexes
ephemeral active context
provider-local runtime state
secrets and credential references
```

### Required outputs

- canonical versus derived versus ephemeral classification;
- owner and writer matrix;
- consistency and freshness requirements;
- query/access patterns;
- retention and deletion propagation;
- confidentiality/integrity requirements;
- backup/export/restore classes;
- migration ownership;
- rebuildability rules;
- physical co-location versus separation criteria;
- expected volume and performance assumptions where needed.

### Candidate mechanisms after classification

- SQLite;
- PostgreSQL;
- file/object storage;
- full-text search;
- vector extensions or databases;
- graph mechanisms;
- specialized telemetry stores.

### Stop condition

No database is selected merely because a framework supports it. Each physical store has a named logical role and an exit/migration path.

---

## TA-06 — Identity, authentication, authorization, policy and secrets

### Question

How do human, Aurora, service, Presence, Harness, provider and device actors prove identity, receive authority, obtain credentials and reach effect-enforcement points?

### Required outputs

- actor/identity class catalog;
- stable identity versus runtime instance distinction;
- authentication proof by actor class and topology;
- session/token/credential lifecycle;
- actor → delegate → executor → Presence chain;
- authority and capability model;
- policy-decision contract;
- Effect Gateway enforcement topology;
- secret classes and credential-reference flow;
- step-up, revocation, recovery and audit requirements;
- Stage A versus Stage B differences.

### Candidate products only after the model exists

```text
Keycloak
Zitadel
Authentik
Ory
SPIFFE/SPIRE
OAuth/OIDC mechanisms
OPA
Cedar
Vault or equivalent credential brokers
minimal Aurora-owned mechanisms where justified
```

No identity product may define Aurora's domain identity or authority semantics.

### Stop condition

Every actor and material effect has an authentication/authority/enforcement path, even if a product choice remains deferred.

---

## TA-07 — Brain, models, memory and Harness integration

### Question

How do deterministic Core responsibilities, cognitive runtimes, model providers, Context Builder, memory and specialized Harnesses compose without duplicating global truth or authority?

### Required outputs

- cognitive responsibility matrix;
- deterministic versus model-mediated boundaries;
- Context Builder ownership and inputs;
- model role taxonomy and provider boundary;
- model routing/fallback contract;
- memory write/read/promotion path;
- tool call and effect authorization path;
- Harness Delegation request/result lifecycle;
- provider-local versus Aurora-global state reconciliation;
- Mastra fit and isolation boundary for the first consumer;
- artifact/evidence return path;
- restart, cancellation and ambiguous completion behavior.

### Reference flow

```text
Presence
→ Sovereign Core / session coordination
→ Context Builder
→ Cognitive Runtime / model
→ capability or Harness request
→ authority/effect enforcement
→ Artifact / Evidence
→ global Outcome and governed state update
```

### Stop condition

Loss or replacement of a model, Mastra runtime or Harness cannot redefine Aurora identity, authority, canonical Project state or governed memory semantics.

---

## TA-08 — Configuration, observability, deployment and operation

### Question

How is the complete system configured, started, observed, updated, recovered and operated on Stage A and Stage B?

### Required outputs

- configuration source/precedence model;
- environment and feature-profile model;
- secret-reference handling;
- startup, readiness, liveness and dependency health;
- structured logging, traces, metrics, audit and evidence separation;
- semantic conventions and redaction;
- migrations and compatibility checks;
- packaging and installation;
- Windows/Linux posture;
- updates, rollback and recovery;
- backup/drill/runbook ownership;
- CI/CD and artifact provenance;
- supply-chain controls;
- resource and latency budgets where material.

### Candidate mechanisms after topology

- service supervisors;
- container packaging;
- OTel collectors/backends;
- deployment automation;
- update channels;
- secret-management integrations.

### Stop condition

The architecture can explain how a clean machine becomes a recoverable Aurora installation and how an operator diagnoses failure without treating telemetry as product truth.

---

## 7. Current first architecture tranche

The current active technical work is the coupled foundation:

```text
TA-01 — Logical modules and canonical ownership
+
TA-02 — Process, runtime and evolutionary topology
```

They are developed together because process placement depends on coherent module ownership, while proposed modules must be tested against real runtime/failure boundaries.

The first reviewed package must contain:

1. a technical component catalog;
2. a module responsibility and canonical ownership matrix;
3. allowed and forbidden dependency directions;
4. Stage A process/runtime topology;
5. always-active and on-demand placement;
6. Stage B evolution diagram;
7. runtime/language scope map, including Go and Mastra constraints;
8. failure-domain and restart ownership;
9. questions classified `DECIDE`, `RESEARCH`, `SPIKE` or `DEFER`;
10. exact inputs to TA-03 repository architecture.

The first tranche must not choose repository strategy, universal API protocol, new database, authentication product, policy engine, Voice provider or observability backend.

---

## 8. Treatment of accepted Stage A Presence constraints

The following accepted decisions are preserved:

- one workstation is the Stage A sovereign host and first Presence;
- minimum Core responsibilities remain available;
- cognition and heavy capabilities start on demand;
- activation belongs to Presence semantics;
- button/UI/hotkey are baseline triggers;
- local wake word is an optional adapter;
- activation is not authentication or authority;
- while locked, Aurora may acknowledge availability but requires unlock before private interaction.

Their current disposition is:

```text
accepted downstream constraints
→ consumed by TA-02 and later Presence/Voice Capability work
→ no further session-policy decomposition now
```

Reopen only when a consuming architecture decision changes because of them.

---

## 9. Research and Architecture Spike policy

Research is required only when an upcoming decision depends on current external facts, standards, versions or products.

Use primary sources:

- official specifications;
- official documentation;
- official repositories;
- standards;
- research papers where appropriate.

A Spike is required only when documentary analysis cannot prove a material runtime property, such as:

- process isolation or recovery behavior;
- streaming latency/interruption;
- compatibility/code generation;
- crash consistency;
- credential delivery/revocation;
- sandbox containment;
- provider state reconciliation.

Spike specification is architecture work. Spike execution still requires explicit authorization.

---

## 10. Anti-overengineering rules

The baseline must not:

- create one service per conceptual module;
- choose microservices because Aurora is large;
- choose a broker before durable asynchronous requirements exist;
- choose Kubernetes before a topology requires orchestration;
- select a universal database for unrelated data classes;
- create proprietary protocols when standard bindings fit;
- research distant M8–M10 mechanisms as current blockers;
- produce diagrams without owners, contracts, consumers and stop conditions;
- make every reversible library choice an ADR;
- wait for all eight areas to be perfect before a bounded executable horizon can be planned.

Preferred starting posture:

```text
modular ownership
+ smallest operational topology
+ explicit ports/contracts
+ replaceable mechanisms
+ evidence before irreversible commitment
```

---

## 11. Integrated Technical Architecture Baseline deliverables

When the ordered work is mature enough for the next executable horizon, the integrated baseline should expose:

```text
01 System/technical context
02 Component and module catalog
03 Canonical ownership matrix
04 Stage A and Stage B runtime topology
05 Repository and build map
06 Contract/API/event catalog
07 Data/store/lifecycle map
08 Identity/authority/credential/effect map
09 Brain/model/memory/Harness integration map
10 Operations/deployment/observability map
11 Accepted ADR/Standard index
12 Deferred-decision and reconsideration-trigger register
13 Evidence/spike requirements
14 First implementation-horizon impact statement
```

The baseline is sufficiently complete when implementers can design a bounded Capability without inventing cross-system product meaning or structural architecture.

---

## 12. Fresh-session read order

A new session working on technical architecture must read:

1. `AGENTS.md`;
2. `docs/tracking/STATUS.md`;
3. this map;
4. `docs/design/AURORA-SYSTEM-ARCHITECTURE-DECISION-LANDSCAPE.md`;
5. Product Blueprint 03, 05, 06, 07, 10, 11, 12 and 13 as relevant;
6. accepted ADRs within their exact scope;
7. current technical work-package artifacts;
8. relevant current research only after the architecture question is known.

A fresh session must be able to state:

- which TA area is active;
- what it must produce;
- what is deferred;
- which technical choices remain open;
- what implementation remains prohibited;
- the exact next question.

---

## 13. Exact next question

The next dialogue begins TA-01/TA-02 with one bounded question:

> What is the minimum set of technical components required to preserve Aurora's accepted ownership boundaries, and which of those components should share or cross a process boundary in Stage A?

The answer must compare 2–3 coherent topology approaches rather than start by selecting repositories, databases or frameworks.
