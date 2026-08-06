---
id: RESEARCH-AURORA-AHDK-CONFORMANCE-GOLDEN-PATHS-V1
title: Aurora Research — Harness Development Kit, Conformance and Golden Paths
document_type: research_report
form: explanation
authority: research
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - research findings on SDK, conformance and paved paths through 2026-08-05
related:
  - DOC-AURORA-BLUEPRINT-05
  - ADR-AURORA-0002
source_manifest: AURORA-RESEARCH-AHDK-CONFORMANCE-GOLDEN-PATHS-v1.sources.json
review_triggers:
  - AHDK language selection
  - SPK-001 result
  - contract/code generation design
last_reviewed: 2026-08-05
---

# Aurora Research — Harness Development Kit, Conformance and Golden Paths

## 1. Research question

How can Aurora make first-party Harness creation high-level, consistent and easy—especially for Aurora herself—without making one SDK the canonical specification or a false security boundary?

---

## 2. Executive finding

The strongest pattern separates:

```text
Specification / semantic conventions
→ canonical behavior

API / generated contracts
→ stable programming surface

SDK / implementation kit
→ default behavior and ergonomics

Adapters / exporters / bindings
→ external mechanisms

Conformance / compatibility tests
→ independent observed behavior

Golden Paths / scaffolding
→ organizationally preferred creation path
```

This supports the refined policy:

> First-party harnesses use AHDK by engineering policy. Contracts and black-box conformance remain universally authoritative and independently testable.

AHDK should be a **development kit**, not merely a client library. It should provide generated contracts, lifecycle, context, effects, artifacts/evidence, observability, simulation, fault injection, scaffolding and provenance.

---

## 3. Specification versus SDK

OpenTelemetry explicitly separates specifications, APIs, SDKs and exporters while defining semantic conventions that can be implemented across languages [S01][S02]. This pattern is relevant because Aurora needs:

- language-neutral product semantics;
- one first SDK initially;
- future cross-language implementations;
- adapters to protocols/stores;
- consistent telemetry;
- no silent divergence.

Implication:

```text
Aurora Contract Model
→ source of behavior

AHDK
→ reference/paved implementation

Provider
→ consumer of AHDK

Conformance Kit
→ external behavior check
```

A bug in AHDK must be detectable by conformance. A direct implementation can be valid if it passes the same contract profile.

---

## 4. Generated contracts and compatibility

Protobuf demonstrates schema-driven multi-language code generation [S03]. Buf adds mechanical breaking-change detection for protobuf APIs [S04]. JSON Schema provides machine-readable validation for JSON-based contracts [S05].

Aurora implications:

- canonical schemas should generate types where practical;
- generated code declares schema/generator version;
- compatibility is checked in CI;
- provider manifest and event/artifact contracts should not be hand-copied across languages;
- semantic changes beyond field compatibility need additional conformance/evals.

No schema format is selected yet. Different boundaries may use JSON Schema and Protobuf.

---

## 5. Conformance as independent authority

MCP and A2A both maintain conformance/TCK initiatives [S06][S07]. The key pattern is testing a protocol implementation externally rather than trusting its library or self-description.

Aurora Conformance profiles should test:

- manifest and contract versions;
- lifecycle/transitions;
- error mapping;
- cancel/deadline;
- restart/resume;
- idempotency;
- event duplication/order;
- artifact integrity;
- evidence requirements;
- authority denial;
- effect receipt;
- privacy/redaction;
- trace propagation.

A provider can pass a synchronous profile and fail resumable/effect profiles. Conformance should be capability/profile-specific.

---

## 6. Golden Paths and scaffolding

Backstage Software Templates demonstrate a platform-engineering pattern where teams create components through templates that encode organization standards while preserving choice [S08]. The important lesson is not adoption of Backstage itself, but the concept:

> Autonomy works better when the correct path is the easiest, visible and self-service path.

Aurora Harness scaffolding should generate:

- manifest definition;
- capability contract references;
- AHDK bootstrap;
- handler skeleton;
- effect/data declarations;
- telemetry;
- test kit;
- conformance profile;
- fault tests;
- CI;
- ownership;
- provenance hooks;
- documentation.

It should also ask design questions so the template does not create empty abstraction.

---

## 7. AHDK module responsibilities

### Contracts-generated

- types;
- validators;
- version compatibility;
- generated documentation.

### Manifest Builder

- required identity/build/capability/effect/data fields;
- safe defaults;
- digest/provenance metadata;
- schema references.

### Provider Runtime

- typed handler registration;
- lifecycle state validation;
- correlation;
- cancellation/deadline;
- error mapping;
- cleanup.

### Context Reader

- authorized Context Pack only;
- provenance/sensitivity/freshness;
- artifact references;
- no unrestricted Core query.

### Effect Client

- typed Effect Requests;
- actor/delegation propagation;
- receipt handling;
- no policy decision or raw secret exposure.

### Artifact/Evidence Client

- content-addressed publishing;
- criterion links;
- method/result/limitations;
- integrity metadata.

### Decision/Signal Client

- structured decision request;
- wait/resume;
- no permission creation.

### Budget/Checkpoint

- consumed/remaining;
- warnings/hard limits;
- checkpoint/snapshot;
- heartbeat and resume.

### Telemetry

- default spans/metrics/logs;
- semantic attributes;
- redaction.

### Simulator/Test Kit

- fake Core/gateway/store;
- controlled clock;
- fault injection;
- restart, duplicates and partitions.

### Scaffolder

- Golden Path creation;
- ownership and CI;
- documentation and conformance.

---

## 8. Why first-party policy should be mandatory

Benefits:

- Aurora-created harnesses start from a verified path;
- repeated integration code disappears;
- lifecycle and telemetry stay consistent;
- contract evolution can be migrated mechanically;
- secure defaults and test fixtures are reused;
- onboarding and review focus on domain logic;
- fewer subtle errors in authority/context/evidence.

Risks:

- SDK monoculture defect;
- premature abstraction;
- one language dominates unsuitable domains;
- provider code becomes coupled to SDK internals;
- false confidence that SDK enforces security.

Mitigations:

- canonical contract independent;
- external Conformance Kit;
- versioned AHDK;
- public API/internal separation;
- waiver process;
- second implementation spike;
- adapters for embedded/legacy/external systems;
- no three-language expansion before one SDK is proven.

---

## 9. SDK is not a security boundary

A library inside a provider process cannot prevent direct filesystem, network, process or credential access if the OS environment already grants it.

Therefore:

```text
AHDK
→ ergonomics, validation and standard requests

Policy Decision Point
→ authorization decision

Effect Gateway / Credential Broker
→ enforce action/secret access

Sandbox / OS / Device Controller
→ contain bypass

Audit/Receipt
→ record outcome
```

AHDK may make the safe path easy and produce warnings. It cannot be the only defense against a malicious or compromised provider.

---

## 10. Public API design principles

AHDK should:

- expose domain terms, not raw protocol frames;
- make required concepts unavoidable;
- default to least authority/data;
- separate stable API from implementation;
- return typed receipts/errors;
- make lifecycle invalid states difficult;
- preserve async/cancel semantics;
- avoid hidden network/global state;
- support testing without running full Core;
- expose exact version/build information.

It should not:

- expose Core database;
- allow arbitrary state transition;
- approve its own evidence;
- hide expensive/provider-specific behavior;
- require all harness internals to use one agent framework.

---

## 11. Error and lifecycle ergonomics

Bad SDK:

```ts
return { success: false, message: "failed" }
```

Desired behavior:

```text
ContractValidationError
AuthorityDenied
BudgetExceeded
DeadlineExceeded
CancellationRequested
ProviderDependencyUnavailable
EffectAmbiguous
CheckpointIncompatible
EvidenceIncomplete
```

Typed errors improve recovery and prevent generic retry.

Lifecycle methods should enforce transitions and terminal behavior rather than accept arbitrary strings.

---

## 12. Testing strategy

### Unit

Provider domain logic using fake Context/effects.

### Contract

Generated schemas/types and compatibility.

### Conformance

Black-box process/protocol behavior.

### Fault

- Core/provider kill;
- duplicate event;
- stale grant;
- effect timeout;
- corrupt artifact;
- budget divergence;
- context injection;
- network restriction.

### Golden Path acceptance

Fresh scaffold implements a minimal capability and passes all baseline checks without editing generated plumbing.

---

## 13. AHDK waiver

A waiver is an explicit engineering exception, not normal freedom.

It requires:

- why AHDK is unsuitable;
- target/runtime constraint;
- alternative adapter design;
- conformance profile;
- security and observability implications;
- owner;
- expiry/removal condition;
- migration strategy.

Likely valid cases:

- microcontroller/firmware;
- hard real-time controller;
- unsupported language;
- external/legacy system;
- performance-isolated protocol adapter.

---

## 14. Multi-language strategy

Do not begin with TypeScript, Python and Go SDKs simultaneously.

Selection criteria for first SDK:

- likely first providers;
- existing MNFS/Pi/Mastra ecosystem;
- AI framework ecosystem;
- type/schema tooling;
- runtime isolation;
- contributor fluency;
- long-term Core language decision;
- code-generation support.

A second SDK should be justified by a real provider and use the same conformance suite.

---

## 15. Required spike — SPK-001

Implement one trivial Capability twice:

1. AHDK reference provider;
2. direct protocol/contract provider.

Measure:

- handwritten code surface;
- contract divergence;
- lifecycle/error correctness;
- observability completeness;
- conformance pass;
- restart/fault behavior;
- performance;
- developer/Aurora creation clarity.

Golden Proof:

```text
same Contract Model
→ two implementations
→ same black-box outcome/evidence
→ AHDK path materially simpler
→ direct path remains possible
→ SDK defect can be caught by Conformance Kit
```

---

## 16. Decision implications

### Supported

- first-party AHDK mandatory by policy;
- universal contract/conformance independent of SDK;
- schema/code generation;
- simulator/fault kit;
- Golden Path scaffolder;
- automatic observability;
- waiver for constrained cases;
- one SDK first.

### Not decided

- first language;
- schema format;
- package layout;
- generator;
- CLI/template engine;
- distribution model;
- exact public API;
- RPC binding.

---

## 17. Limitations

- OpenTelemetry and Backstage patterns come from broader platform ecosystems, not personal-agent Harnesses.
- Protocol TCKs do not cover Aurora domain semantics.
- Code generation can preserve syntax while semantic compatibility still breaks.
- AHDK ergonomics need a working reference provider to evaluate.
- Security claims require environment/gateway tests.

---

## 18. Conclusion

AHDK should be treated as an internal platform product:

```text
canonical contracts
+ generated APIs
+ runtime helpers
+ secure effect requests
+ artifacts/evidence
+ observability
+ simulator and faults
+ scaffolding
+ independent conformance
```

This creates leverage for humans and for Aurora herself while keeping the architecture portable and enforceable outside the SDK.
