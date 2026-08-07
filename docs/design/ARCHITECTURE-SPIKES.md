---
id: DESIGN-AURORA-ARCHITECTURE-SPIKES
title: Aurora Architecture Spike Portfolio
document_type: spike_portfolio
form: reference
authority: design
status: proposed
version: 0.2.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - proposed A0 architecture spike questions and proof designs
related:
  - DESIGN-AURORA-CAPABILITY-HARNESS-ARCHITECTURE
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - RESEARCH-AURORA-HARNESS-ARCHITECTURE-V1
  - RESEARCH-AURORA-HARNESS-INTEROPERABILITY-V1
  - RESEARCH-AURORA-AHDK-CONFORMANCE-GOLDEN-PATHS-V1
  - RESEARCH-AURORA-DURABLE-EXECUTION-V1
  - RESEARCH-AURORA-AUTHORITY-IDENTITY-EFFECTS-V1
  - RESEARCH-AURORA-EVENTS-OBSERVABILITY-SCHEMAS-V1
last_reviewed: 2026-08-06
---

# Aurora Architecture Spike Portfolio

## 1. Authorization notice

```text
Portfolio design: AUTHORIZED
Spike planning/specification: NOT YET STARTED
Spike execution: PROHIBITED
Production promotion: PROHIBITED
```

This portfolio identifies material uncertainties and possible proof procedures. It is not a batch implementation plan.

After A0 acceptance, each spike must pass ACRM R0–R4 independently and receive:

- exact scope;
- approved environment;
- budget;
- authority/effects boundary;
- implementation/microdesign;
- evidence format;
- explicit execution authorization.

No spike is automatically required. The selected first Product Milestone determines applicability and order.

## 2. Spike principles

### Reduce one named uncertainty

A spike answers a decision-relevant question. It is not an excuse to build a platform prototype.

### Minimal executable evidence

Implement only enough to distinguish alternatives or falsify a premise.

### Fixed evaluation before execution

Criteria, workload and expected observations are declared before results are known.

### Reproducible environment

Record:

- source revision;
- versions;
- environment;
- configuration;
- commands;
- datasets/fixtures;
- network/credential policy;
- generated artifacts and hashes.

### Experimental code is disposable by default

Every spike declares:

```text
DISCARD
PROMOTE_WITH_REWRITE
PROMOTE_AFTER_REVIEW
REFERENCE_ONLY
```

No experiment becomes production because it works once.

### Negative evidence is valid

A failed mapping or excessive operational burden may close a candidate more usefully than a successful demo.

### No self-certification

The implementing actor produces Claims. A reviewer verifies evidence and decision implications.

## 3. Common spike artifact

Each authorized spike must create a dedicated spec containing:

```yaml
id: SPK-AURORA-...
question: ...
decision_informed: ...
baseline: ...
alternatives: []
assumptions: []
environment: ...
authority: ...
budget: ...
procedure: []
metrics: []
golden_proof: ...
stop_conditions: []
artifacts: []
disposal_rule: ...
reviewers: []
```

Closeout records:

- observations versus expectations;
- raw artifact references;
- limitations;
- reproducibility result;
- operational burden;
- security/privacy findings;
- decision recommendation;
- ADR/Spec/requirements impact;
- disposal/promotion action.

---

# SPK-001 — Contract Model, AHDK and Independent Conformance

## Decision informed

- ADR-0001 semantic independence;
- ADR-0002 AHDK policy;
- initial Contract Model and Conformance Kit boundaries;
- first AHDK language evaluation inputs.

## Question

Can the same Aurora capability be implemented through the official AHDK and through a direct binding while preserving identical external semantics and passing one independent black-box conformance suite?

## Premises to test

1. Contract meaning can exist independently from AHDK.
2. AHDK materially reduces repeated work and omissions.
3. Conformance can detect both SDK and direct-provider divergence.
4. Minimal Contract Model concepts are sufficient for a first vertical slice.
5. AHDK does not need global Core state or authorization decisions.

## Reference capability

Candidate:

```text
artifact_transform.v1
```

Input:

- immutable text artifact reference;
- transformation profile;
- output schema;
- deadline/budget.

Output:

- transformed artifact;
- Claim;
- criterion-linked Evidence;
- trace/cost metadata.

Effects:

- read one authorized artifact;
- write one output artifact;
- no network/credential/device effect.

This keeps the domain trivial so the spike evaluates the platform boundary rather than agent quality.

## Implementations

### A — AHDK provider

Uses generated types, manifest builder, lifecycle, artifact/evidence and telemetry helpers.

### B — direct provider

Implements the accepted binding/contracts without AHDK.

## Conformance cases

- valid manifest and capability discovery;
- schema-valid request/response;
- invalid schema rejection;
- transition order;
- duplicate start/idempotency;
- cancellation before and during work;
- deadline/budget exhaustion;
- missing required artifact;
- corrupt artifact hash;
- provider restart if resumable profile applies;
- trace/correlation propagation;
- forbidden effect attempt;
- version incompatibility;
- breaking schema mutation.

## Measurements

- provider code/boilerplate size;
- implementation time;
- defects/omissions found;
- runtime latency/overhead;
- debugging clarity;
- conformance-suite complexity;
- generated documentation quality;
- migration impact from one compatible change.

## Golden Proof

```text
same canonical capability contract
→ AHDK provider passes
→ direct provider passes
→ intentionally broken AHDK provider fails
→ intentionally broken direct provider fails
→ Aurora observes equivalent lifecycle/artifact/evidence
```

## Stop conditions

- Contract Model cannot be expressed independently;
- direct implementation requires undocumented SDK behavior;
- conformance cannot distinguish known violations;
- abstraction size exceeds the reference capability value without a removal plan.

## Disposal

Reference provider code is `REFERENCE_ONLY` unless a future Contract explicitly promotes/reworks it.

---

# SPK-002 — MCP Binding Boundary

## Decision informed

- whether and how MCP is supported;
- which Aurora concepts map directly, require wrapping or remain unsupported;
- current specification/SDK maturity at implementation time.

## Question

Can Aurora safely use MCP for tools/resources and bounded asynchronous operations without weakening Delegation, authority, evidence and recovery invariants?

## Scope

Test the then-current stable MCP specification and official conformance suite.

Mapping candidates:

| Aurora | MCP candidate |
|---|---|
| capability/tool operation | Tool |
| contextual source | Resource |
| bounded async tool call | Tasks extension where stable/applicable |
| progress | notifications/task status |
| input request | elicitation/input mechanism where supported |
| cancellation | protocol cancellation/task cancellation |
| provider description | server metadata/discovery |

Aurora-specific envelope remains responsible for:

- Mission/Delegation identity;
- Authority Grant;
- data classes;
- global budget/criteria;
- trust/approval;
- Evidence/Verdict;
- cross-Harness relationships.

## Procedure

- implement minimal server and client adapter;
- expose one resource and one effect-free tool;
- add one effectful tool routed through Effect Gateway;
- run official conformance;
- test local and remote transport as applicable;
- test auth allowed/denied/revoked;
- test progress/cancel/error;
- test reconnect and critical-state recovery;
- trace call through Aurora correlation.

## Adversarial cases

- server advertises undeclared capability/effect;
- manifest and actual schema differ;
- tool returns narrative “success” with no receipt;
- duplicate request after timeout;
- cancellation races effect;
- server attempts context exfiltration;
- protocol version mismatch.

## Golden Proof

```text
Aurora-authorized Delegation
→ MCP adapter exposes minimum Context/Grant
→ resource/tool operation executes
→ effect is enforced outside server
→ cancellation/denial/retry remain correct
→ artifact/evidence and trace return under Aurora identities
```

## Exit classifications

- `SUPPORTED_CORE_BINDING`;
- `SUPPORTED_FOR_TOOLS_RESOURCES_ONLY`;
- `SUPPORTED_WITH_AURORA_EXTENSION`;
- `DEFERRED_FOR_MATURITY`;
- `REJECTED_FOR_SCOPE`.

---

# SPK-003 — A2A Binding Boundary

## Decision informed

- whether A2A is the preferred remote opaque-Harness binding;
- mapping between Agent Card/Task/Artifact and Aurora Provider/Delegation/Artifact;
- SDK/TCK maturity by selected language.

## Question

Can a remote opaque Harness participate in Aurora Missions through A2A while Aurora retains global authority, state ownership, evidence and recovery?

## Mapping candidates

| Aurora | A2A candidate |
|---|---|
| Provider/Manifest | Agent Card + Aurora metadata/wrapper |
| Delegation | Task |
| operator/provider communication | Message/Parts |
| deliverable | Artifact |
| progress | task status/stream events |
| waiting for input | input-required state |
| cancellation | task cancellation |

## Procedure

- publish/discover Agent Card;
- validate identity/signature options;
- create Task from an Aurora Delegation;
- stream progress;
- request input through Aurora;
- publish artifact;
- cancel and reconnect;
- kill/restart provider;
- run official TCK;
- run Aurora conformance;
- test version/SDK compatibility.

## Security cases

- Agent Card changes build/version after approval;
- provider requests broader data than Context Pack;
- push/webhook sender cannot be authenticated;
- child capability request attempts direct peer delegation;
- task artifact includes undeclared sensitive data;
- trust from one Agent Card is reused for another instance.

## Golden Proof

```text
approved exact remote provider
→ A2A Task created from Delegation
→ state survives connection loss
→ input/cancel/artifact work
→ provider cannot mint authority or child work
→ official TCK + Aurora conformance pass
```

## Exit classifications

Same as SPK-002, plus a maturity matrix per SDK/binding.

---

# SPK-004 — Durable Execution and Recovery

## Decision informed

- DurableExecutionPort semantics;
- initial engine/baseline choice;
- operational burden acceptable for a local-first single-user system.

## Question

Which minimal durable-execution approach can preserve Mission/Delegation progress, timers, human input and effect idempotency across crashes without overbuilding infrastructure?

## Candidates

At authorization time, revalidate:

- minimal project-owned baseline;
- DBOS;
- Restate;
- Temporal;
- another candidate only if a requirement justifies it.

Inngest or cloud-oriented options may remain research references depending on local-first fit.

## Common workflow

```text
create Delegation
→ durable step A
→ request idempotent external effect
→ record receipt/checkpoint
→ process crash
→ recovery
→ durable timer
→ wait for operator input
→ provider restart
→ resume
→ publish artifact
→ complete
```

## Fault matrix

- Core crash before/after checkpoint;
- worker crash during activity;
- database/runtime restart;
- network partition;
- duplicate event delivery;
- timeout with unknown effect result;
- cancellation while waiting;
- code/workflow version change;
- backup and restore on new environment;
- corrupt/missing checkpoint;
- long idle period.

## Metrics

- setup and operational components;
- local resource use;
- programming constraints/determinism;
- recovery semantics;
- idempotency support;
- human-in-the-loop support;
- timers/scheduling;
- versioning/migration;
- observability;
- backup/restore;
- language/runtime fit;
- debugging experience;
- lock-in and removal cost.

## Golden Proof

```text
start
→ effect receipt
→ checkpoint
→ kill Core and provider
→ restore/restart
→ no duplicate effect
→ wait for input
→ resume
→ complete with one coherent trace/evidence chain
```

## Decision guard

The fastest demo does not win if it hides effect ambiguity, recovery or operational state ownership.

---

# SPK-005 — Delegated Authority and Effect Enforcement

## Decision informed

- Authority Grant/Delegation Token model;
- PDP, Effect Gateway and Credential Broker boundaries;
- revocation semantics;
- candidate policy/identity technologies.

## Question

Can Leandro delegate a narrow authority chain through Aurora to a Harness/worker such that allowed effects succeed, denied effects produce no external change, revocation acts during execution and audit explains every decision?

## Reference effect

Use a reversible local effect and one denied target, for example:

```text
allowed: write under isolated test directory
prohibited: write outside directory / access network secret
```

No production credentials or real devices.

## Identity chain

```text
subject: Leandro
actor: Aurora
executor: Provider Instance / Worker Run
action: effect kind
resource: exact target
context: Delegation, device, environment, risk, expiry
```

## Procedure

- issue short-lived scoped grant/token;
- provider requests effect through gateway;
- PDP allows one exact resource;
- gateway executes and returns receipt;
- provider requests prohibited resource;
- deny with no effect;
- revoke during run;
- reject subsequent request;
- simulate gateway/PDP partition;
- reconcile uncertain result;
- rotate/revoke credential reference;
- query complete audit chain.

## Adversarial cases

- provider bypasses AHDK;
- token replay;
- confused deputy/subject–actor loss;
- child provider attempts inherited grant;
- resource path traversal;
- expiry/clock skew;
- request modified after policy decision;
- receipt spoofing;
- direct secret read from environment;
- revocation during disconnected execution.

## Golden Proof

```text
narrow grant
→ allowed effect + verifiable receipt
→ denied effect + no external mutation
→ AHDK bypass still denied by boundary
→ live revocation stops further effects
→ audit explains subject/actor/executor/rule/result
```

---

# SPK-006 — Distributed Observability, Events and Evidence Correlation

## Decision informed

- Aurora semantic conventions;
- event envelope/schema boundaries;
- trace propagation and privacy;
- relation between telemetry and durable state/evidence.

## Question

Can one trace and evidence graph explain a cross-boundary journey without turning logs into product authority or leaking sensitive context?

## Journey

```text
Aurora Mission
→ Delegation to remote/local reference Harness
→ tool/resource call
→ Effect Gateway request/receipt
→ artifact publication
→ evidence record
→ outcome
```

## Required correlation

- Mission/Delegation/run IDs;
- provider/build/capability versions;
- model/tool where policy permits;
- Context Pack classification/hash, not unrestricted contents;
- authority/policy decision ID;
- budget consumption;
- effect request/receipt;
- artifact/evidence hashes;
- errors/retry/recovery;
- parent/child spans/events.

## Tests

- baggage/context propagation;
- process/protocol boundary;
- restart continues logical correlation without invalid span reuse;
- redaction of secrets/personal data;
- high-cardinality protection;
- sampling versus required audit/evidence;
- duplicate/out-of-order events;
- trace backend unavailable;
- Evidence remains verifiable without telemetry backend;
- cost attribution.

## Golden Proof

A reviewer can reconstruct why Aurora selected a provider, what it received, which effect was allowed, what artifact was produced and which criterion was supported—without reading raw transcripts or exposing secrets.

---

# SPK-007 — Capability Registry, Exact Build Trust and Provenance

## Decision informed

- Registry data model;
- provider/build identity;
- provenance/attestation requirements;
- trust inheritance and revalidation policy.

## Question

Can Aurora bind provider trust and approval to an exact build/environment so a changed artifact cannot silently inherit previous authority?

## Procedure

- register provider source revision/build digest;
- attach manifest, contract/binding versions and conformance result;
- record approval for narrow scope;
- run approved provider;
- rebuild with no semantic change but new digest;
- attempt execution under old approval;
- change manifest/effect declaration;
- test suspension/revocation/retirement;
- test rollback to known build;
- inspect provenance and audit.

## Adversarial cases

- version string unchanged but binary changed;
- signed manifest points to different build;
- provider configuration/environment changes behavior;
- dependency compromise or missing provenance;
- old instance continues after revocation;
- approval copied across instance/publisher.

## Golden Proof

```text
build A verified/approved
→ build A executes in scope
→ build B with same name/version is discovered but blocked
→ revalidation/approval required
→ revocation prevents new and contains active work
```

---

# SPK-008 — Framework and Runtime Neutrality

## Decision informed

- validity of Aurora-owned semantics;
- adapter burden;
- first reference-Harness runtime considerations;
- whether Contract Model accidentally assumes one framework.

## Question

Can two providers implemented with materially different internal approaches expose the same capability and lifecycle without changing Aurora Mission/Delegation semantics?

## Candidate implementations

Choose at authorization time based on current maturity, for example:

- Mastra versus LangGraph;
- Pi/custom deterministic service versus another agent runtime;
- agentic provider versus deterministic implementation.

The goal is difference, not a framework popularity contest.

## Same external contract

- identical capability version;
- same input/output schemas;
- same allowed effects;
- same lifecycle/conformance profile;
- same artifact/evidence requirements;
- same error/recovery expectations.

## Comparison

- adapter/code size;
- lifecycle impedance mismatch;
- cancellation/recovery behavior;
- telemetry mapping;
- provider-local state exposure;
- framework-specific concepts leaking into contract;
- performance/cost;
- developer experience;
- testing and failure diagnosis.

## Golden Proof

Aurora swaps/selects between both providers without changing Mission Contract, Delegation semantics, authority or acceptance criteria.

## Failure condition

If canonical contracts need framework-specific fields or one provider cannot represent required lifecycle without semantic loss, record exactly whether the problem is:

- bad Aurora abstraction;
- inadequate adapter;
- provider incompatibility;
- capability profile mismatch.

---

## 4. Suggested dependency relationships

This is not authorization or fixed sequence:

```text
SPK-001 Contract/AHDK
├── informs SPK-002 MCP
├── informs SPK-003 A2A
└── informs SPK-008 neutrality

SPK-005 Authority/Effects
├── required for effectful protocol tests
└── informs SPK-007 provider trust

SPK-004 Durability
└── required before strong restart/recovery claims

SPK-006 Observability
└── should instrument other spikes after semantic conventions exist
```

A smaller first Product Milestone may authorize only one or two spikes.

## 5. Portfolio closeout rule

A spike closes only when:

- fixed target revision and environment are recorded;
- Golden Proof result is explicit;
- raw evidence is preserved;
- known limitations are stated;
- decision implications are reviewed;
- experimental artifacts are discarded/promoted according to rule;
- affected research/ADR/requirements are updated;
- STATUS records whether any next step is authorized.

No successful spike automatically selects a technology. It informs a separately reviewed decision.
