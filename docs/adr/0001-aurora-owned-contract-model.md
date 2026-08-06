---
id: ADR-AURORA-0001
title: Aurora-owned Contract Model and Replaceable Bindings
document_type: adr
form: explanation
authority: decision
status: proposed
version: 0.2.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - proposed ownership of cross-Harness semantics
  - proposed protocol binding policy
related:
  - DOC-AURORA-BLUEPRINT-03
  - DOC-AURORA-BLUEPRINT-05
  - DOC-AURORA-BLUEPRINT-07
  - DOC-AURORA-BLUEPRINT-10
  - DOC-AURORA-BLUEPRINT-12
  - RESEARCH-AURORA-HARNESS-ARCHITECTURE-V1
  - RESEARCH-AURORA-HARNESS-INTEROPERABILITY-V1
  - DESIGN-AURORA-CAPABILITY-HARNESS-ARCHITECTURE
supersedes: []
superseded_by: null
review_triggers:
  - an open protocol covers the complete Aurora delegation domain
  - contract-model spike evidence contradicts the proposal
  - adapter/conformance cost becomes disproportionate
last_reviewed: 2026-08-06
---

# ADR-0001 — Aurora-owned Contract Model and Replaceable Bindings

## 1. Status

```text
PROPOSED
```

This ADR records a proposed architecture decision for operator review. It does not authorize implementation or a protocol selection.

## 2. Decision scope

This decision concerns the semantic boundary between Aurora and specialized Harnesses.

It does **not** decide:

- the Aurora Core language;
- the wire format of every contract;
- the first local RPC technology;
- whether A2A or MCP is adopted in a particular milestone;
- the internal framework/runtime of a Harness;
- the durable execution engine;
- the operational database;
- the first real Harness integration.

Those choices remain under future Capability Specs, spikes and ADRs.

## 3. Context

Aurora must eventually coordinate heterogeneous systems such as:

- software engineering Harnesses, including a future MNFS adapter;
- research and evaluation Harnesses;
- hardware and firmware systems;
- laboratory telemetry and device-control systems;
- deterministic services;
- local processes and remote agent applications;
- future personal-domain capabilities.

These providers may use different:

- languages;
- frameworks;
- process boundaries;
- transports;
- security profiles;
- state models;
- durations;
- data volumes;
- evidence types;
- failure and recovery behavior.

Examples discussed during discovery include:

```text
Research Harness
→ may use a TypeScript/Python agent workflow

MNFS
→ may use Pi and its own deterministic development control plane

Laboratory Harness
→ may combine deterministic device services, firmware and analysis agents

Evaluation Harness
→ may consume large telemetry/data artifacts without receiving global authority
```

Aurora must understand a common external meaning without requiring all systems to share one internal runtime.

## 4. Problem

If the model of one framework or protocol becomes the Aurora domain model, several failures become likely:

### Framework lock-in

Aurora concepts become constrained by how one runtime represents agents, graphs, sessions, tools or workflows.

Changing the framework becomes a product-domain migration rather than an adapter change.

### Semantic gaps hidden in prompts

Concepts not supported by the chosen framework—such as delegated authority, data classification, evidence linked to criteria, build-bound trust, physical interlocks or global budget—may be improvised in text rather than represented as contracts.

### Heterogeneous systems forced into an agent abstraction

A deterministic hardware controller or telemetry service may be modeled as an “agent” even when it does not reason, creating unclear state and authority.

### Global state leakage into Harnesses

A Harness may start owning project state, memory, authority or cross-domain acceptance because the protocol lacks an explicit Aurora owner.

### Protocol capability mistaken for permission

A tool/agent advertises an operation and Aurora treats technical availability as authority to execute it.

### Recovery and evidence divergence

Each provider reports completion differently, making global reconciliation, audit and acceptance dependent on free-form narratives.

## 5. Decision drivers

The decision must support:

1. stable product semantics across model/framework/provider replacement;
2. first-party Harness ergonomics without universal runtime lock-in;
3. local-first operation and future remote providers;
4. agentic and deterministic providers;
5. explicit identity, authority, budget, data and effect boundaries;
6. durable/observable work lasting seconds to days;
7. cancellation, restart, reconciliation and idempotency;
8. artifacts and evidence, not only chat messages;
9. physical and digital safety progression;
10. machine-verifiable versioning and conformance;
11. YAGNI—only concepts required by proved slices should be implemented;
12. ability to adopt open standards where they fit.

## 6. Alternatives

### Option A — One framework owns the domain

Examples:

```text
Aurora = one large Mastra application
Aurora = one LangGraph graph
Aurora = one OpenAI Agents SDK manager
```

#### Advantages

- fastest initial prototyping;
- fewer custom contracts;
- framework-provided tools, memory, streaming and tracing;
- one programming model.

#### Disadvantages

- framework lifecycle becomes product lifecycle;
- physical/deterministic systems are awkwardly modeled;
- global memory/authority/evidence inherit framework limitations;
- future migration is broad and risky;
- Harness independence becomes superficial;
- provider-specific session state can be mistaken for authoritative Aurora state.

#### Assessment

Rejected as the constitutional architecture. Frameworks remain valid Harness-internal candidates.

---

### Option B — MCP owns the domain

Aurora models all external work as MCP tools/resources/prompts and optional Tasks.

#### Advantages

- open ecosystem;
- strong tool/resource discovery;
- standardized local/remote integrations;
- official SDKs and conformance work.

#### Disadvantages

MCP does not by itself own the complete Aurora semantics for:

- global Mission decomposition;
- delegated authority and effect policy;
- cross-Harness child Delegations;
- project/world memory governance;
- global budget and acceptance criteria;
- provider/build trust lifecycle;
- evidence and verdict distinction;
- physical-safety envelopes.

#### Assessment

Rejected as the complete domain. Retained as a candidate binding for tools/resources and bounded asynchronous operations.

---

### Option C — A2A owns the domain

Aurora models Harnesses directly as A2A agents, tasks, messages and artifacts.

#### Advantages

- designed for opaque agent applications;
- task lifecycle, artifacts, streaming/polling/push;
- Agent Cards and multi-turn collaboration;
- official conformance/TCK direction.

#### Disadvantages

A2A does not fully define Aurora-specific:

- personal/project memory governance;
- fine-grained Authority Grants and effect enforcement;
- global budget/guardrail/stop conditions;
- provider trust and data-class approvals;
- causal evidence and Product Milestone acceptance;
- physical device safety;
- distinction between Aurora constitutional state and provider-local state.

Language SDK maturity and protocol evolution may also differ at implementation time.

#### Assessment

Rejected as the complete domain. Retained as a leading candidate binding for remote, opaque and long-running Harness providers.

---

### Option D — Every integration is bespoke

Aurora has no canonical external model; each adapter translates directly to Core internals.

#### Advantages

- no up-front semantic design;
- maximum freedom per integration;
- quick first adapter.

#### Disadvantages

- semantic drift;
- duplicate lifecycle and error handling;
- no universal conformance;
- testing and trust become adapter-specific;
- difficult provider substitution;
- Aurora cannot safely scaffold new Harnesses;
- cross-Harness composition becomes special-case orchestration.

#### Assessment

Rejected.

---

### Option E — Aurora-owned semantics with replaceable bindings

Aurora defines a language/framework-neutral Contract Model. Bindings map it to native SDK calls, local RPC, A2A, MCP, HTTP/gRPC or event transports.

#### Advantages

- product meaning survives infrastructure replacement;
- agentic and deterministic providers fit;
- authority/evidence/recovery are explicit;
- first-party Golden Paths can be rich without closing external interoperability;
- conformance can test semantics independently of SDK/runtime;
- protocols are adopted where strong rather than stretched beyond their domain.

#### Disadvantages

- Aurora must maintain schemas, versioning and conformance;
- mapping to external protocols requires adapters;
- risk of overengineering unused concepts;
- semantic duplication is possible when Aurora concepts overlap standards.

#### Assessment

Recommended, subject to SPK-001/002/003/008 evidence.

## 7. Proposed decision

Aurora will own the canonical meaning of cross-domain concepts including:

```text
Capability
Provider
Provider Instance
Capability Manifest
Verification and Trust Assessment
Approval for Scope
Mission
Delegation and Child Delegation
Context Pack
Authority Grant / Delegation Token
Budget and Guardrails
Decision Request
Event and Checkpoint
Effect Request / Decision / Receipt
Artifact
Observation
Claim
Receipt
Evidence
Verdict
Outcome
Cancellation and Recovery
```

The canonical model will be:

- independent of programming language;
- independent of provider framework;
- versioned;
- schema-addressable;
- compatible with code generation;
- testable through black-box conformance;
- minimal and expanded only by a proved Product Milestone.

## 8. Binding policy

Bindings transport or adapt Aurora semantics; they do not own them.

### Native AHDK / in-process

Preferred candidate for first-party local providers where low overhead and typed integration matter.

### Local RPC

Candidate when process isolation, independent restart or cross-language support is required.

### A2A

Candidate for remote/opaque agent applications with stateful tasks and artifacts.

### MCP

Candidate for tools, resources and bounded calls, including provider-internal tool access.

### HTTP/gRPC/Event transport

Candidates where explicit APIs, high-volume data or performance requirements justify them.

A provider may expose multiple bindings. The Capability Registry records exact contract/binding versions and conformance results.

## 9. Mapping rule

For every binding:

```text
Aurora canonical concept
→ adapter mapping
→ protocol concept
→ provider behavior
```

The adapter must document:

- lossless mappings;
- approximations;
- unsupported behavior;
- version constraints;
- retry/idempotency semantics;
- cancellation/recovery semantics;
- security assumptions;
- evidence/conformance coverage.

If a protocol cannot represent a required invariant, Aurora must either:

1. wrap it with an Aurora envelope;
2. add a narrow extension after gap analysis;
3. select another binding;
4. reject the provider for that capability scope.

Aurora must not silently weaken the invariant.

## 10. State ownership consequences

### Aurora owns global state

- identity and user authority;
- Projects and global Missions;
- Delegation relationships;
- grants, approvals and global budgets;
- Capability Registry trust/approval;
- cross-Harness composition;
- global criteria and outcomes;
- operator interaction;
- canonical project/world memory.

### Harness owns local execution state

- internal plan and agent/workflow graph;
- local workers/attempts;
- domain-specific intermediate state;
- local checkpoints;
- intermediate artifacts;
- specialized methodology.

### Shared boundary

- observable lifecycle snapshot;
- context and authority contract;
- significant events;
- decision/escalation requests;
- budget consumption;
- artifacts/evidence;
- cancellation and recovery.

Connection loss does not imply completion or failure. Aurora reconciles against durable provider state/checkpoints.

## 11. Security consequences

The Contract Model represents authority; it does not enforce effects by itself.

```text
Contract / Authority Grant
→ policy decision
→ Effect Gateway / Credential Broker
→ sandbox or device boundary
→ Effect Receipt
```

A protocol connection or advertised operation never grants permission.

Context transfer is minimized and classified. Child Delegations receive independent grants rather than inheriting parent credentials transitively.

## 12. Reliability and observability consequences

- lifecycle transitions are structured and versioned;
- significant state is recoverable, not only streamed;
- retries depend on error class/idempotency;
- ambiguous effects block until receipt reconciliation;
- traces correlate Mission, Delegation, provider, tool/effect and artifact/evidence;
- provider “success” is a Claim, not global acceptance;
- contract/version changes during a run require explicit migration or binding to the original version.

## 13. Implementation constraints if accepted

1. No universal schema catalog before SPK-001 proves the minimal slice.
2. The first Contract Model must include only concepts used by the reference journey.
3. Each concept must declare a canonical owner and invariant.
4. Source schemas and generated SDK types must remain distinguishable.
5. A direct implementation and an AHDK implementation must pass the same conformance suite.
6. Protocol adapters must not write Core state directly.
7. The first real provider cannot define missing semantics ad hoc.
8. No custom wire protocol unless a documented gap survives standard mapping/spike.

## 14. Validation plan

### SPK-001 — Contract, AHDK and conformance

Implement one minimal capability twice:

- official first-party AHDK path;
- direct protocol path.

Both must produce identical observable lifecycle, artifacts, errors and conformance results.

### SPK-002 — MCP mapping

Prove tool/resource mapping, auth denial, error, cancellation and async behavior against the current official conformance suite.

### SPK-003 — A2A mapping

Prove Agent Card/task/artifact/stream/cancel/restart mapping and pass the official TCK plus Aurora conformance.

### SPK-008 — Runtime neutrality

Implement the same capability in two materially different internal runtimes without changing Aurora Mission/Delegation semantics.

## 15. Acceptance evidence

This ADR should not be accepted only from documentary reasoning. Required evidence before use in production planning:

- approved Product Blueprint sections 03, 05, 07, 10 and 12;
- reviewed interoperability research current for chosen protocol versions;
- SPK-001 result showing independent semantics/conformance;
- at least one protocol mapping spike;
- explicit cost/complexity assessment;
- documented removal criteria for abstractions not used by the first real slice.

A0 may accept the architectural principle while exact schemas and bindings remain future decisions, provided STATUS and roadmap preserve that boundary.

## 16. Consequences

### Positive

- stable Aurora identity and architecture across runtime changes;
- explicit boundaries for memory, authority, evidence and recovery;
- provider substitution and multi-language support;
- reusable first-party AHDK and external adapters;
- safety model applicable to digital and physical systems;
- conformance and provider trust tied to exact versions/builds;
- avoids making MNFS or any current framework the center.

### Negative

- more product-owned design and maintenance;
- need for versioned schemas and compatibility policy;
- adapter and conformance workload;
- duplicated terminology if mappings are poorly managed;
- first milestone can be delayed by over-modeling.

### Neutral/trade-off

Aurora intentionally owns semantics that no current standard fully covers, while reusing external standards for transport/tool/task behavior.

## 17. Risks and mitigations

| Risk | Mitigation |
|---|---|
| speculative platform architecture | only add concepts consumed by a Product Milestone/Golden Proof |
| adapter explosion | standardize official bindings; reject low-value providers; use AHDK first-party |
| schema drift | canonical source, code generation, compatibility checks and conformance |
| duplicate protocol semantics | mapping tables and explicit ownership; do not rename external concepts unnecessarily |
| false sense of safety | enforcement remains in policy/gateways/sandbox, not schema |
| central Core bottleneck | control plane remains governed; authorized data plane may be direct |
| abstraction blocks domain-specific work | Harness retains internal autonomy and domain artifacts |

## 18. Reconsideration triggers

Re-open or supersede this ADR if:

- one stable open standard covers all required Aurora semantics without weakening invariants;
- SPK-001 shows the canonical model cannot remain independent of implementation;
- adapter/conformance maintenance exceeds demonstrated value;
- most providers prove deterministic/simple enough for a smaller model;
- protocol evolution makes the proposed wrapping approach incompatible;
- security review finds authority/effect semantics cannot be safely carried;
- a real Product Milestone requires a fundamentally different state owner.

## 19. Decision record on acceptance

If accepted, record:

- exact accepted version/hash;
- which principles are accepted now;
- which schema/binding details remain open;
- which spikes are authorized separately;
- affected requirements and future Capability Specs;
- explicit statement that no runtime implementation is authorized by the ADR alone.
