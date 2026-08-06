---
id: DOC-AURORA-BLUEPRINT-07
title: Orquestração de Harnesses e Delegações
document_type: product_blueprint_section
form: explanation
authority: constitutional
status: proposed
version: 0.2.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - Aurora-harness responsibility boundary
  - hierarchical orchestration
  - delegation and context contracts
  - control-plane/data-plane principles
  - protocol and runtime neutrality
related:
  - DOC-AURORA-BLUEPRINT-03
  - DOC-AURORA-BLUEPRINT-04
  - DOC-AURORA-BLUEPRINT-05
  - DOC-AURORA-BLUEPRINT-10
  - DOC-AURORA-BLUEPRINT-11
  - DOC-AURORA-BLUEPRINT-12
  - DOC-AURORA-BLUEPRINT-13
  - ADR-AURORA-0001
  - ADR-AURORA-0002
  - DESIGN-AURORA-CAPABILITY-HARNESS-ARCHITECTURE
review_triggers:
  - delegation boundary changes
  - peer/child orchestration changes
  - state ownership changes
  - protocol mapping changes
last_reviewed: 2026-08-05
---

# 7. Orquestração de Harnesses e Delegações

## 7.1 Purpose

Aurora will not perform every specialized domain directly.

A real engineering mission may require:

- technical research;
- architecture/design;
- software implementation;
- firmware build and flashing;
- circuit simulation;
- laboratory measurement;
- data analysis;
- independent review;
- documentation;
- safety verification.

Each domain can use different agents, deterministic workflows, tools, storage, environments and evidence methods.

The orchestration problem is not merely:

> “How can one agent call another agent?”

It is:

> **How can Aurora preserve one global objective, context, authority, budget and evidence model while specialized systems retain enough internal autonomy to do expert work?**

The accepted model is **hierarchical, contractual orchestration**.

---

## 7.2 What is a harness?

A Harness is a specialized system that can contain:

- one or more agents;
- deterministic and adaptive workflows;
- tools;
- state machines;
- a runtime;
- local storage;
- execution environments;
- policy/guardrails local to the domain;
- evaluation;
- recovery;
- human review interfaces;
- methodology and Golden Paths.

Examples:

### MNFS

Software-engineering harness that may plan, implement, review, test, integrate and produce evidence.

### Research Harness

Searches primary sources, compares claims, records provenance and produces research artifacts.

### Firmware Harness

Builds versioned firmware, runs static/unit tests, flashes approved targets and collects logs.

### Laboratory Harness

Coordinates instruments, protocols, telemetry, interlocks and experiment evidence.

### Evaluation Harness

Runs controlled evals, holdouts, regression and comparison.

A Harness is not simply a model endpoint or a tool.

---

## 7.3 Core responsibility boundary

### Aurora owns global concerns

- original user intent;
- product/project objective;
- project and personal context;
- global mission;
- cross-domain decomposition;
- capability selection;
- provider selection;
- global authority and data policy;
- global budgets;
- dependencies;
- child Delegations;
- material decisions and escalation;
- global mission state;
- composition of artifacts/evidence;
- acceptance/outcome;
- communication with Leandro;
- longitudinal memory and learning.

### Harness owns specialized local concerns

- domain methodology;
- internal agents/workflows;
- local plan;
- tool choice inside authority;
- local attempts/retries;
- local concurrency;
- intermediate artifacts;
- detailed provider state;
- domain-specific checks;
- local recovery according to contract;
- internal optimization.

### Shared contract boundary

- Delegation Envelope;
- Context Pack;
- Authority Grant;
- budget/deadline;
- lifecycle/status;
- material events;
- decision requests;
- effects/receipts;
- artifacts;
- evidence;
- outcome;
- cancellation/recovery;
- trace/audit.

> Aurora governs why, what, global limits and composition. A Harness governs how to produce its specialized result inside the Delegation.

---

## 7.4 Why not centralize every step?

If Aurora plans every local operation:

- Core needs detailed knowledge of every domain;
- local expertise becomes weak;
- every framework leaks into the global model;
- plans become huge;
- local adaptation requires constant round trips;
- Aurora becomes a monolithic superagent.

Example of bad centralization:

```text
Aurora tells Hardware Harness:
1. choose buck topology
2. calculate inductor
3. select MOSFET
4. run simulator command X
5. inspect waveform Y
```

Better:

```text
Aurora delegates:
Produce and compare up to three supply architectures for 24 V input,
5 A output, specified thermal/cost constraints and required evidence.
```

The Hardware Harness decides which topologies and calculations are appropriate.

---

## 7.5 Why not fully decentralize?

If harnesses freely call one another and propagate authority:

- context leaks;
- credentials spread;
- budgets become unowned;
- loops appear;
- global state diverges;
- failure recovery becomes ambiguous;
- responsibility for a result is unclear;
- one provider can silently involve an unapproved external provider.

Therefore:

> Harnesses may request collaboration; Aurora governs every delegation crossing an authority boundary.

---

## 7.6 Hierarchical orchestration

Flow:

```text
Leandro defines/approves global objective
        ↓
Aurora creates Mission and global plan
        ↓
Aurora creates Delegation A
        ↓
Harness A executes internally
        ↓
Harness A requests Capability B
        ↓
Aurora evaluates need and scope
        ↓
Aurora creates child Delegation B
        ↓
Harness B returns artifact/evidence
        ↓
Aurora exposes only relevant result to Harness A
        ↓
Aurora composes global outcome
```

The child does not inherit the parent's complete Context Pack, authority, credentials or budget.

---

## 7.7 Example — cross-domain engineering mission

Mission:

> “Develop and validate a new control module for the programmable power supply.”

```text
MIS-CONTROL-MODULE
│
├── DEL-RESEARCH
│   Capability: technical_research
│   Outcome: candidate approaches with primary sources
│
├── DEL-HARDWARE
│   Capability: circuit_design_review
│   Input: selected requirements and research artifacts
│   │
│   └── requests thermal_analysis
│       └── Aurora creates DEL-THERMAL
│
├── DEL-FIRMWARE
│   Capability: firmware_variants
│   Input: accepted design and control requirements
│
├── DEL-LAB
│   Capability: laboratory_protocol_execution
│   Input: approved protocol, artifacts and authority
│
└── DEL-EVALUATION
    Capability: multi_objective_evaluation
    Input: telemetry/artifacts from lab
```

Aurora owns global acceptance:

- design satisfies requirements;
- firmware artifact identity is preserved;
- laboratory limits were respected;
- evidence covers objective;
- remaining risk is stated.

A completed child Delegation does not close the global Mission automatically.

---

## 7.8 Delegation Envelope

A Delegation is a versioned scoped contract.

### 7.8.1 Identity

```text
mission_id
delegation_id
parent_delegation_id
capability_id/version
provider instance/build
contract model version
attempt/run identity
```

### 7.8.2 Intent

- objective;
- expected outcome;
- scope;
- non-goals;
- assumptions;
- dependencies;
- rationale;
- priority.

### 7.8.3 Context

- project snapshot;
- authority sources;
- selected memory;
- artifact references;
- data classification;
- freshness requirements;
- known uncertainty;
- excluded context.

### 7.8.4 Control

- Authority Grant;
- allowed/prohibited effects;
- environment;
- budget;
- deadline/window;
- guardrails;
- stop conditions;
- escalation;
- cancellation;
- credential references.

### 7.8.5 Quality

- acceptance criteria;
- evidence requirements;
- output schemas;
- reviewer requirements;
- reproducibility;
- limitations format.

### 7.8.6 Recovery

- checkpoint semantics;
- idempotency;
- heartbeat;
- retry class;
- snapshot query;
- retention;
- migration rules.

---

## 7.9 Complete Delegation example

```yaml
delegation:
  id: DEL-RESEARCH-0042
  mission: MIS-AURORA-MEMORY-01
  parent: null

  target:
    capability: technical-research/1
    provider_instance: research-local-01@sha256:...
    contract_model: aurora/1.0

  intent:
    objective: compare long-term agent memory architectures
    outcome: research report with source manifest and decision implications
    scope:
      - observational memory
      - structured project memory
      - temporal retrieval
      - evaluation
    non_goals:
      - select production database
      - implement prototype

  context:
    project_snapshot: CTX-PRJ-AURORA-019
    authority_sources:
      - DOC-AURORA-BLUEPRINT-06
    memories:
      - MEM-AURORA-DISCOVERY-...
    data_max: INTERNAL
    freshness:
      web_sources: current_at_execution

  authority:
    grant: GRANT-RESEARCH-014
    allow_effects:
      - read_public_web
      - write_artifact
      - use_approved_models
    deny_effects:
      - repository_write
      - external_message
      - credential_read

  budget:
    max_cost_usd: 15
    max_wall_time: 4h
    max_provider_calls: 100

  quality:
    criteria:
      - primary_sources_only_for_normative_claims
      - alternatives_and_limitations
      - source_claim_traceability
    evidence:
      - source_manifest
      - coverage_report

  stop:
    - budget_exhausted
    - material_source_conflict_requires_operator
    - no_primary_evidence
```

Exact schema remains open; concepts are mandatory.

---

## 7.10 Context Pack

A Context Pack is a provider-specific compiled projection, not the full project or conversation.

It may contain:

```text
objective and expected outcome
current project/mission state
authoritative decisions
relevant hypotheses
constraints/non-goals
selected memories with provenance
artifact/knowledge references
allowed effects and authority snapshot
freshness and sensitivity
known omissions/limitations
```

### 7.10.1 Context Pack properties

- immutable or content-addressed for a run/revision;
- attributable to a builder/version;
- scoped to Delegation;
- minimized for provider;
- explicit authority precedence;
- sensitive fields redacted/referenced;
- refresh rules;
- inspectable during review.

### 7.10.2 Context Pack example

```yaml
context_pack:
  id: CTX-DEL-0042-R1
  delegation: DEL-RESEARCH-0042
  built_at: ...
  builder_version: context-builder/0.3
  source_hashes:
    project: ...
    authority: ...
  sections:
    objective: ...
    current_decisions:
      - ref: DOC-AURORA-BLUEPRINT-06
        excerpt_ref: section-6.27
    project_state: ...
    memories:
      - ref: MEM-...
        epistemic: USER_APPROVED
    artifacts:
      - ref: ART-...
  excluded:
    - personal_global_memory:not_required
    - credentials:prohibited
```

Provider cannot expand scope by reading the entire memory store.

---

## 7.11 Authority Grant

Authority Grant answers:

```text
who acts
on behalf of whom
for which Delegation
which action/effect
on which resource
during which time
inside which environment
under which budget/guardrails
```

A parent Delegation cannot simply hand its token to a child provider.

When a Harness requests another Capability, Aurora creates a new grant with no more authority than needed and no more than the parent/global authority permits.

---

## 7.12 Child capability request

A Harness may emit:

```yaml
capability_request:
  requester: DEL-HARDWARE-01
  capability: thermal_analysis/1
  reason: compare junction temperature across candidate topologies
  proposed_input_artifacts:
    - ART-CIRCUIT-CANDIDATES
  requested_data_classes:
    - INTERNAL
  requested_effects:
    - compute
    - write_artifact
  budget_request:
    max_compute_minutes: 30
  blocks:
    - hardware_recommendation
```

Aurora evaluates:

- relevance to global objective;
- whether parent could do it locally;
- compatible approved providers;
- context minimization;
- authority/budget;
- data channel need;
- decision/escalation.

Result can be approved, denied, modified or escalated.

---

## 7.13 Lifecycle

Initial common lifecycle:

```text
DRAFT
→ PROPOSED
→ AUTHORIZED
→ QUEUED
→ RUNNING
→ WAITING_FOR_INPUT | BLOCKED | DEGRADED
→ CANCEL_REQUESTED
→ VERIFYING
→ COMPLETED | FAILED | CANCELED | REJECTED
```

Potential dimensions:

```text
lifecycle
phase
attention
health
```

### Invariants

- no RUNNING without authorized exact contract;
- terminal state explicit;
- completion requires required artifact/evidence, not provider Boolean;
- transition records actor/reason/time/version;
- cancellation may be cooperative but authority revocation is immediate at gateways;
- restart cannot create a second active run silently;
- provider state and Aurora global state reconcile;
- failed verification can return to correction/re-execution without rewriting history.

Exact FSM needs a Capability Spec/spike.

---

## 7.14 Attempts and runs

Delegation identity remains stable across attempts.

```text
Delegation DEL-001
├── Attempt ATT-001
│   └── Run RUN-001 on provider instance A
└── Attempt ATT-002
    └── Run RUN-002 on provider instance B
```

Attempt records:

- hypothesis/reason for retry;
- provider/version;
- Context Pack revision;
- authority/budget revision;
- prior failure class;
- reused artifacts/effects;
- outcome.

No blind retry. A new attempt needs a valid retry reason or changed condition.

---

## 7.15 Event contract

Material events include:

### Lifecycle

- proposed;
- authorized;
- queued;
- started;
- checkpoint;
- waiting;
- blocked;
- verifying;
- completed/failed/canceled.

### Work

- material progress;
- hypothesis changed;
- plan adapted;
- artifact published;
- evidence recorded;
- provider substitution.

### Control

- decision requested/resolved;
- budget threshold;
- authority expiring/revoked;
- cancellation requested;
- stop condition.

### Effect

- requested;
- allowed/denied;
- started;
- receipt;
- ambiguous/reconciled.

### Security/incident

- policy violation;
- provider suspended;
- context anomaly;
- channel revoked;
- recovery failure.

Event fields should include:

- ID/type/version;
- source/provider;
- mission/delegation/run;
- timestamp/sequence;
- actor;
- data classification;
- trace context;
- payload/artifact reference;
- idempotency/deduplication where applicable.

A stream message is not durable truth. Critical state/artifact remains queryable.

---

## 7.16 Progress reporting

Progress should report material state, not token-by-token activity.

Bad:

```text
read file
read another file
thinking
called model
thinking
```

Useful:

```text
- source discovery complete: 18 primary sources
- two architecture families remain
- one conflict requires decision
- 62% budget remains
- report draft expected after comparison phase
```

Aurora aggregates progress according to Leandro's attention preference.

---

## 7.17 Decision requests

A provider asks when the choice changes:

- objective;
- scope;
- material architecture;
- risk;
- authority;
- budget;
- data boundary;
- irreversible effect;
- acceptance interpretation.

It should not ask Leandro to choose ordinary domain details that belong to the Harness.

Decision Request includes:

- exact question;
- why it cannot decide locally;
- alternatives;
- evidence;
- recommendation;
- impact on cost/time/risk;
- default safe behavior;
- what is blocked;
- deadline if any.

---

## 7.18 Artifact and evidence flow

### Artifact publication

Provider publishes content-addressed output/reference.

### Claim

Provider states which criteria it believes are satisfied.

### Verification

Deterministic/reviewer/eval systems produce receipts and findings.

### Evidence

Artifacts and receipts are linked to criteria/hypotheses.

### Verdict

Allowed authority accepts, rejects or requests correction.

### Outcome

Aurora composes the global result and remaining limitations.

Example:

```text
Research Harness publishes report + source manifest
→ Conformance validates artifact schemas
→ Review verifies claim-source coverage
→ Evidence supports research criteria
→ Aurora accepts research Delegation
→ Architecture decision remains separate ADR
```

---

## 7.19 Control plane versus data plane

### Control plane

Always governed by Aurora:

- identities;
- mission/delegation;
- contracts;
- authority;
- budget;
- channels;
- decisions;
- lifecycle;
- acceptance.

### Data plane

May flow directly:

```text
Laboratory Harness
→ high-rate telemetry
→ Evaluation Harness / Artifact Store
```

Aurora authorizes channel with:

- producer/consumer;
- schema;
- data classes;
- rate/volume;
- encryption;
- storage/retention;
- duration;
- budget;
- trace;
- revocation;
- allowed transformations.

A direct channel cannot create a new Delegation, permission or credential.

---

## 7.20 Data-channel lifecycle

```text
PROPOSED
→ AUTHORIZED
→ ESTABLISHED
→ ACTIVE
→ THROTTLED | DEGRADED
→ REVOKED | CLOSED | FAILED
```

Required behaviors:

- close when Delegation ends unless explicitly retained;
- rate/volume budget;
- schema/version validation;
- producer/consumer identity;
- revocation propagation;
- artifact/evidence references;
- no hidden side channel.

---

## 7.21 State ownership and reconciliation

### Aurora global state

- Mission/Delegation lifecycle;
- selected Provider;
- grants/budgets;
- dependencies;
- pending decisions;
- artifact/evidence metadata;
- global acceptance;
- last verified provider snapshot.

### Harness local state

- internal plan;
- worker runs;
- local queue;
- intermediate artifacts;
- domain checkpoints;
- local retry details.

### Reconciliation

Aurora queries a Provider snapshot:

```text
provider run identity
current local state
checkpoint
last sequence/event
active effects
published artifacts
health
```

Cases:

- provider says running, Aurora says canceled → revoke effects, request cancel, reconcile;
- Aurora says running, provider has no run → classify lost/unrecoverable or restore;
- events missing but terminal artifact exists → verify artifact and reconstruct state;
- effect result ambiguous → query Effect Gateway/target receipt.

Missing connection never implies completion.

---

## 7.22 Cancellation

Cancellation has layers:

### User/mission cancellation

Stop global objective/delegations.

### Delegation cancellation

Provider receives cooperative signal.

### Authority revocation

Effect Gateways block new effects immediately.

### Process termination

May be used after grace/containment but does not replace cleanup/reconciliation.

Provider declares:

- cancellable phases;
- safe checkpoints;
- cleanup;
- non-cancelable external operations;
- compensation;
- expected confirmation.

A canceled Delegation may still need a cleanup Outcome and incident if effects remain ambiguous.

---

## 7.23 Recovery

Provider Manifest/contract declares:

- resumable;
- checkpoint model;
- snapshot query;
- heartbeat;
- idempotency;
- effect reconciliation;
- state compatibility;
- retention;
- failure taxonomy.

### Core restart

Reload global state, reconnect providers, compare snapshots/events and resume decisions.

### Harness restart

Resume exact checkpoint/build or report failure.

### Network partition

Preserve last verified state, avoid duplicate effect, continue only if local envelope explicitly allows disconnected operation.

### Provider version changed

Do not resume silently on incompatible build. Use explicit migration or original version.

### Context stale during long run

Provider/Aurora may request Context Pack refresh. Refresh cannot silently alter original objective/authority; revision is recorded.

---

## 7.24 Error taxonomy

### Contract

- schema violation;
- invalid transition;
- incompatible version;
- missing required field/artifact.

### Authority

- effect outside grant;
- expired token;
- data class violation;
- unauthorized child request.

### Provider

- unavailable;
- unhealthy;
- non-conformant behavior;
- internal failure;
- lost checkpoint.

### Dependency

- tool/service unavailable;
- transient network;
- rate limit;
- missing artifact.

### Budget

- warning;
- hard limit;
- accounting divergence.

### Evidence

- invalid artifact;
- incomplete coverage;
- non-reproducible result;
- reviewer rejection.

### Security/safety

- context exfiltration;
- credential misuse;
- provider compromise;
- unsafe device effect.

Error class determines retry, substitute, replan, contain or escalate.

---

## 7.25 Provider substitution example

Research cloud provider becomes unavailable halfway.

Aurora should:

1. preserve current artifacts/checkpoint;
2. determine whether results are provider-specific;
3. select an approved fallback;
4. build a new Context Pack for fallback;
5. issue a new grant;
6. record new Attempt;
7. avoid assuming hidden local state compatibility;
8. compare or reverify outcome;
9. report cost/quality difference if material.

Aurora should not forward the original provider's credential or entire transcript.

---

## 7.26 Protocol bindings

### Native AHDK / in-process

Benefits:

- low overhead;
- generated types;
- direct cancellation/telemetry;
- first-party ergonomics.

Risks:

- shared process failure/permissions;
- language coupling;
- weaker isolation.

### Local RPC

Benefits:

- process isolation;
- independent restart;
- cross-language;
- explicit serialization.

Costs:

- lifecycle/protocol complexity;
- debugging/distribution.

### A2A

Candidate for:

- remote opaque agent application;
- stateful Task;
- messages/artifacts;
- streaming/polling/push;
- Agent Card discovery.

Aurora still wraps/maps authority, budget, context and evidence.

### MCP

Candidate for:

- tools;
- resources;
- bounded calls;
- ecosystem interoperability;
- Tasks extension for certain async operations.

MCP Task is not automatically an Aurora Delegation.

### HTTP/gRPC/events

Used where requirements justify.

> Transport carries Aurora semantics; it does not define them.

---

## 7.27 Protocol mapping table

| Aurora concept | Native/RPC | A2A candidate | MCP candidate |
|---|---|---|---|
| Provider discovery | registry/manifest | Agent Card + registry | server discovery/config |
| Delegation | typed call/task | Task | tool call/Task extension for bounded case |
| Context Pack | typed reference/payload | message parts/artifact refs | resources/arguments/refs |
| Progress | event stream | status/stream | progress notification/task status |
| Artifact | artifact API/ref | Artifact | resource/content/ref |
| Decision Request | signal/event | multi-turn message/input-required | elicitation/user input pattern where supported |
| Authority | Aurora token/grant | auth + Aurora envelope | MCP auth + Aurora envelope |
| Budget | Aurora field/enforcement | extension/metadata | extension/metadata |
| Evidence | Aurora artifact/schema | Artifact/profile | resource/result profile |
| Recovery | provider snapshot/checkpoint | Task state + Aurora state | Task handle + Aurora state |

Mapping remains adapter responsibility and must be tested against current specs.

---

## 7.28 Framework neutrality

Harness internal runtime can be:

- Mastra;
- LangGraph;
- Pi;
- OpenHands;
- OpenAI Agents SDK;
- Temporal-backed workflow;
- deterministic service;
- Go/Python/TypeScript application;
- firmware/RTOS;
- manual workflow adapter.

Aurora cares about externally observable contract and evidence.

### Neutrality test

Implement same reference Capability in two runtimes.

Aurora must not change:

- Mission semantics;
- Delegation Envelope;
- Authority model;
- artifact/evidence criteria;
- global lifecycle.

Provider-specific adapter details can differ.

---

## 7.29 Durable execution boundary

Long campaigns cannot depend only on one agent process.

Aurora exposes a Durable Execution Port covering:

- schedule/start;
- timer;
- wait/signal;
- checkpoint;
- retry eligible step;
- cancellation;
- resume;
- workflow version;
- status query.

A Harness may also have internal durability.

State ownership remains explicit:

- engine history executes;
- provider local state specializes;
- Aurora operational state owns global semantics.

Temporal, DBOS, Restate and alternatives require spike; no engine is accepted here.

---

## 7.30 MNFS integration boundary

MNFS is expected to become a provider of software-engineering capabilities.

Aurora should not:

- depend on MNFS internal SQLite/entities;
- assume MNFS workers are Aurora workers;
- treat MNFS plan as Aurora Mission automatically;
- reuse MNFS authority without mapping;
- block Core development on MNFS readiness.

Integration prerequisites:

- MNFS offers stable external boundary;
- capabilities defined;
- Delegation mapping;
- Context Pack mapping;
- authority/effect mapping;
- artifacts/evidence mapping;
- lifecycle/recovery mapping;
- conformance;
- readiness and risk review.

MNFS remains one provider among future research, firmware, hardware and laboratory harnesses.

---

## 7.31 Anti-patterns

### Tool explosion

Hundreds of tools exposed directly to Aurora without capability/domain composition.

### Prompt contract

Provider behavior depends on prose only, without schemas/state/evidence.

### Shared super-context

All harnesses receive full project/personal memory.

### Transitive credential propagation

Parent harness hands credentials to child.

### Central micromanagement

Aurora plans every local tool call.

### Peer swarm

Harnesses autonomously create work and authority loops.

### Stream as state

Missed event destroys mission truth.

### Provider claim as outcome

`success=true` closes criteria.

### Framework leakage

Aurora domain adopts LangGraph/Mastra/Pi state as canonical.

### Retry until green

Same attempt repeated without new hypothesis/evidence.

---

## 7.32 Evaluation requirements

Future implementation must prove:

1. Harness can plan internally without Core micromanagement;
2. child capability request returns through Aurora;
3. child receives narrower context/authority;
4. provider cannot mint/grant authority;
5. global Mission survives provider restart;
6. duplicate/out-of-order events do not corrupt state;
7. missing stream event does not lose critical artifact/state;
8. direct data channel is authorized, bounded and revocable;
9. provider substitution creates a new Attempt/context/grant;
10. cancellation blocks new effects and reconciles cleanup;
11. ambiguous effect does not trigger blind retry;
12. artifact/claim/evidence/verdict remain distinct;
13. same Capability works in two runtimes;
14. A2A/MCP adapters preserve Aurora semantics where mapped;
15. MNFS integration can be added without changing global domain;
16. Leandro receives a consolidated decision/status rather than raw multi-agent noise.

---

## 7.33 Open decisions

- exact Mission/Delegation FSM;
- child request protocol;
- Context Pack schema;
- Authority Grant/token format;
- event envelope/transport;
- artifact/evidence schemas;
- provider snapshot API;
- direct channel implementation;
- durable engine;
- A2A/MCP adoption scope;
- local RPC;
- provider concurrency and leases;
- global planner architecture;
- first real provider.

---

## 7.34 Non-goals

- unrestricted peer-to-peer federation;
- one framework for all harnesses;
- all bytes routed through Aurora Core;
- one generic agent that performs every domain;
- provider access to global memory;
- automatic authority inheritance;
- state derived only from chat/events;
- a universal workflow language;
- immediate MNFS integration;
- public agent marketplace;
- maximum number of workers/providers as a success metric.
