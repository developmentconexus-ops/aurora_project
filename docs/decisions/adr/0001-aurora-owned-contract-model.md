---
id: ADR-AURORA-0001
title: Aurora-owned Contract Model and Replaceable Bindings
document_type: adr
form: explanation
authority: decision
status: accepted
accepted_at: 2026-08-06
acceptance_evidence: DOC-AURORA-A0-OPERATOR-ACCEPTANCE
version: 0.2.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - ownership of cross-Harness semantics
  - protocol binding policy
related:
  - DOC-AURORA-BLUEPRINT-03
  - DOC-AURORA-BLUEPRINT-05
  - DOC-AURORA-BLUEPRINT-07
  - DOC-AURORA-BLUEPRINT-10
  - DOC-AURORA-BLUEPRINT-12
  - RESEARCH-AURORA-HARNESS-INTEROPERABILITY-V1
  - RESEARCH-AURORA-HARNESS-ARCHITECTURE-V1
  - DESIGN-AURORA-ARCHITECTURE-SPIKES
supersedes: []
superseded_by: null
review_triggers:
  - open standard covers Aurora semantics without invariant loss
  - adapter/conformance cost proves structurally disproportionate
  - SPK-001, SPK-002, SPK-003 or SPK-008 contradict assumptions
last_reviewed: 2026-08-06
---

# ADR-0001 — Aurora-owned Contract Model and Replaceable Bindings

## 1. Context

Aurora must coordinate specialized providers that may be:

- agentic or deterministic;
- local or remote;
- TypeScript, Python, Go, firmware or another language;
- implemented with Mastra, LangGraph, Pi, OpenHands, custom code or a future runtime;
- exposed through local SDK calls, RPC, A2A, MCP or other transports.

The product also owns concepts not completely represented by any one studied protocol/framework:

```text
Identity
Project / Goal / Mission
Capability / Provider / Harness
Delegation
Context Pack
Authority Grant
Budget / Guardrail
Effect / Receipt
Artifact / Claim / Evidence / Verdict
Presence / Device
Incident / Improvement Candidate
```

Using an external protocol or framework as the product model would create semantic coupling between Aurora's constitution and a replaceable implementation mechanism.

Research evidence:

- MCP provides a strong tool/resource protocol with extensions/tasks but does not own Aurora's Mission, authority, evidence and physical-safety model;
- A2A provides agent discovery, tasks/messages/artifacts and remote opaque collaboration but does not own Aurora's global sovereignty/acceptance model;
- agent frameworks intentionally expose framework-specific agent/workflow abstractions;
- conformance ecosystems work better when protocol contracts are testable independently from one SDK implementation.

See:

- `docs/research/AURORA-RESEARCH-HARNESS-INTEROPERABILITY-v1.md`;
- `docs/research/AURORA-RESEARCH-HARNESS-ARCHITECTURE-v1.md`;
- `docs/research/AURORA-RESEARCH-AGENT-FRAMEWORKS-RUNTIMES-v1.md`.

---

## 2. Decision drivers

- preserve Aurora identity and product invariants across implementation replacement;
- support heterogeneous digital and physical providers;
- local-first operation;
- least-authority context exchange;
- independent conformance;
- protocol/framework replaceability;
- understandable versioning/migration;
- avoid inventing a wire protocol without proven need;
- YAGNI: define only semantics required by accepted Product Milestones.

---

## 3. Alternatives

### A — Framework-owned product model

Choose Mastra/LangGraph/OpenHands/Pi or another framework and model Aurora directly in its concepts.

**Advantages**

- fastest initial coding;
- framework-native tooling;
- fewer adapters initially.

**Rejected because**

- identity/state becomes coupled to framework lifecycle;
- non-agentic devices/services fit poorly;
- replacing the framework becomes a product migration;
- security/authority may be reduced to whatever the framework exposes;
- cross-Harness semantics become implementation-specific.

### B — Protocol-owned product model

Treat MCP or A2A objects as Aurora's canonical domain.

**Advantages**

- standards-based immediately;
- interoperability tooling already exists.

**Rejected because**

- studied protocols solve narrower interoperability questions;
- Aurora requires global authority, memory/context, budget, evidence, Presence and physical-effect semantics;
- protocol evolution would become product-constitution evolution;
- one protocol does not cover all provider classes efficiently.

### C — Entirely custom Aurora protocol/runtime

Invent domain plus wire protocol plus runtime immediately.

**Advantages**

- maximum control;
- theoretically exact semantics.

**Rejected because**

- reinvents mature transport/interoperability work;
- increases maintenance and security surface;
- creates premature commitment before spikes prove gaps;
- slows the product before actual integration evidence exists.

### D — Aurora-owned domain contracts with replaceable bindings

Aurora owns canonical semantics. Existing protocols/runtimes are adapted where suitable. Custom wire extensions appear only after demonstrated gaps.

**Selected.**

---

## 4. Decision

Aurora SHALL own a language-, framework- and transport-independent **Canonical Contract Model** for global cross-boundary concepts.

Initial semantic families include, only as required by milestones:

```text
Capability / Provider / Provider Instance
Manifest / compatibility
Mission / Delegation / child Delegation
Context Pack
Authority Grant / Budget / Guardrail
Lifecycle / checkpoint / cancellation / recovery
Event
Effect Request / Effect Receipt
Artifact / Claim / Evidence / Verdict / Outcome
Presence / Device identity references where crossing a boundary
```

The Contract Model is:

- owned by Aurora Product/Capability Specs;
- versioned independently from SDKs/transports;
- implementation-neutral;
- compatible with generated language types where chosen;
- governed through ACRM/ADR change control.

Bindings MAY include:

- in-process AHDK implementation;
- local RPC;
- A2A;
- MCP;
- HTTP/gRPC;
- event/message bindings;
- future device protocols.

A binding MUST NOT silently change semantic meaning.

---

## 5. Protocol adoption policy

Aurora will not create a custom network protocol merely because it owns domain contracts.

For each boundary:

1. define required Aurora semantics;
2. identify open-standard candidates;
3. map semantics to candidate protocol;
4. identify information/authority/lifecycle gaps;
5. run an Architecture Spike when consequential;
6. adopt standard directly when it fits;
7. create an adapter/extension when fit is partial;
8. create a new binding/protocol only when evidence proves need.

This preserves both sovereignty and interoperability.

---

## 6. Candidate role mapping

The following is **research-informed and not an implementation selection**:

| Need | Candidate mechanism | Current stance |
|---|---|---|
| local first-party Harness integration | native AHDK / local RPC | spike required |
| tools/resources/prompts | MCP | strong candidate |
| remote opaque task-oriented provider | A2A | strong candidate |
| high-performance typed local/remote call | gRPC/Protobuf | candidate |
| human-readable web API | HTTP/JSON | candidate |
| event envelope | CloudEvents | candidate |
| event-channel description | AsyncAPI | candidate |

The table guides SPK-001/002/003; it does not select stack.

---

## 7. Semantic ownership invariant

```text
Aurora semantics
        ↓
Canonical schemas/contracts
        ↓
Adapter/binding
        ↓
External protocol/runtime
```

Never invert:

```text
External framework/protocol
        ↓
forces Aurora product semantics
```

If an external protocol cannot represent an Aurora invariant, the adapter must either:

- carry approved extension metadata;
- place a semantic gateway around the protocol;
- reject unsupported operation;
- or cause architecture reconsideration.

It must not drop the invariant silently.

---

## 8. Versioning principles

Contract compatibility must distinguish:

```text
semantic contract version
schema representation version
binding/protocol version
AHDK version
provider implementation/build version
```

Example:

```text
Delegation semantics: 1.1
Delegation JSON schema: 1.1.2
A2A binding profile: 0.3
AHDK TS: 0.7
Provider build: sha256:...
```

One update does not automatically require all layers to version together.

A breaking semantic change requires:

- impacted requirement analysis;
- affected Capability Specs;
- migration plan;
- provider compatibility review;
- conformance update;
- ADR/contract change when material.

---

## 9. Translation loss

Every adapter must make translation loss explicit.

Potential classes:

- unsupported lifecycle state;
- missing authority field;
- artifact representation mismatch;
- absent cancellation semantics;
- streaming mismatch;
- provider discovery mismatch;
- error taxonomy mismatch;
- identity/auth mismatch.

Adapter policy:

```text
lossless
→ use normally

lossy but safe and declared
→ restrict capability/profile

loss violates invariant
→ unsupported / fail closed
```

---

## 10. Conformance consequence

Because semantics are Aurora-owned, conformance cannot test only transport syntax.

A provider must satisfy:

### Semantic conformance

- state transitions;
- required identifiers;
- error behavior;
- authority propagation rules;
- artifact/evidence relationships;
- cancellation/recovery expectations.

### Binding conformance

- MCP/A2A/RPC framing;
- schema serialization;
- protocol lifecycle;
- compatibility negotiation.

A provider may pass its upstream protocol TCK and still fail Aurora semantic conformance.

---

## 11. Security consequences

Aurora-owned contracts carry authority *descriptions*, but do not themselves enforce effects.

Security path remains:

```text
Authority Grant
→ Policy Decision
→ Effect Gateway
→ Credential/OS/device enforcement
→ Receipt
```

Therefore:

- protocol authentication is necessary but not full authorization;
- SDK types are not sandboxing;
- an A2A/MCP provider cannot self-declare its permitted scope;
- translation must preserve identity and authority references;
- unrepresentable security semantics fail closed.

---

## 12. Control and data planes

The Contract Model governs the control plane even when data moves directly.

Example:

```text
Aurora Delegation
→ authorize telemetry channel
→ Lab Harness ─────────────→ Evaluation Harness
             large waveform
```

The byte stream may use a protocol optimized for telemetry while the Aurora contracts record:

- purpose;
- producer/consumer;
- data class;
- schema;
- duration/rate;
- authorization;
- retention;
- evidence references.

---

## 13. Failure and degraded behavior

If a binding is unavailable:

- Capability Registry may choose another compatible provider/binding;
- Aurora may degrade to local capability;
- work may pause awaiting reconnection;
- no broader authority is granted to recover automatically.

If adapter behavior becomes ambiguous:

```text
ambiguity
→ block affected operation
→ emit Finding
→ inspect contract/protocol mapping
→ repair or revise ADR/profile
```

---

## 14. Consequences

### Positive

- product semantics survive framework replacement;
- heterogeneous Harnesses can coexist;
- simpler mock/reference providers;
- contract-first code generation possible;
- standardized conformance;
- independent security enforcement;
- future protocol adoption remains possible;
- MNFS does not define Aurora's architecture.

### Negative

- Aurora must own schemas/versioning;
- adapters add work;
- risk of duplicating concepts already present in protocols;
- semantic migration becomes a project responsibility;
- poorly designed abstractions can become an internal framework lock-in.

### Operational costs

- conformance suite maintenance;
- compatibility matrix;
- schema/code generation tooling;
- adapter test environments;
- migration documentation.

---

## 15. Risks and mitigations

### Overengineering contracts before real providers

Mitigation:

- only formalize semantics needed by accepted milestones;
- SPK-001 uses a deliberately small reference capability;
- no public SDK ecosystem in current scope.

### Adapter explosion

Mitigation:

- support a small set of paved bindings;
- per-provider custom adapters only when justified;
- Capability profiles declare supported features.

### Lowest-common-denominator contracts

Mitigation:

- semantic Core is minimal but extensible;
- provider-specific capabilities remain namespaced;
- optional profiles/extensions explicitly versioned.

### Contract drift

Mitigation:

- generated types;
- compatibility checks;
- conformance suite;
- stable IDs;
- CI freshness;
- migration policy.

---

## 16. Evidence and validation plan

This decision requires evidence before production implementation.

### SPK-001

Proves one Capability is usable:

- via AHDK;
- via direct protocol implementation;
- under the same Aurora conformance suite.

### SPK-002

Tests MCP mapping:

- tool/resource;
- Tasks;
- cancellation;
- auth;
- errors;
- Aurora extension gaps.

### SPK-003

Tests A2A mapping:

- Agent Card;
- Task;
- streaming;
- artifact;
- input-required;
- cancellation/reconnect;
- Aurora Delegation gap.

### SPK-008

Implements same Capability in two internal runtimes and verifies Aurora contracts do not change.

---

## 17. Acceptance criteria

The decision remains valid when:

- two different provider implementations satisfy the same canonical contract;
- Aurora does not branch domain logic based on internal framework;
- binding failures do not corrupt global state;
- semantic version mismatch is detected before unsafe execution;
- adapter cannot silently drop authority/evidence fields;
- protocol-specific fields remain outside constitutional sources unless promoted intentionally.

---

## 18. Reconsideration triggers

Reconsider this ADR when:

- an open standard demonstrably covers all required Aurora semantics without invariant loss;
- maintaining Aurora contracts produces greater complexity than adapters eliminate;
- real providers expose a fundamentally different boundary;
- AHDK/contract conformance proves impractical;
- physical device domains cannot fit the abstraction safely;
- semantic versioning becomes operationally unmanageable.

A reconsideration does not mean automatically adopting a specific framework. It reopens the architecture question with evidence.

---

## 19. Non-decisions

This ADR does **not** select:

- schema language;
- wire protocol;
- serialization;
- AHDK implementation language;
- MCP or A2A adoption;
- durable execution engine;
- event transport;
- Capability Registry storage;
- provider runtime;
- MNFS adapter design.
