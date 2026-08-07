---
id: ADR-AURORA-0002
title: First-party Harness Development Kit and Universal Conformance
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
  - first-party AHDK policy
  - universal conformance policy
  - waiver boundary
related:
  - DOC-AURORA-BLUEPRINT-05
  - DOC-AURORA-BLUEPRINT-07
  - DOC-AURORA-BLUEPRINT-10
  - DOC-AURORA-BLUEPRINT-13
  - RESEARCH-AURORA-AHDK-CONFORMANCE-GOLDEN-PATHS-V1
  - ADR-AURORA-0001
supersedes: []
superseded_by: null
review_triggers:
  - first-party integration repeatedly bypasses AHDK
  - AHDK complexity exceeds measured value
  - conformance cannot remain implementation-independent
  - the first AHDK language constrains required providers
last_reviewed: 2026-08-06
---

# ADR-0002 — First-party Harness Development Kit and Universal Conformance

## 1. Status

```text
ACCEPTED
```

This ADR establishes an organizational and technical policy accepted in A0. It does not choose the first SDK language or authorize implementation.

## 2. Decision scope

The decision addresses how Harnesses built inside Projeto Aurora should implement the Aurora Contract Model consistently and how every provider—first-party or external—is verified.

It does not decide:

- source schema format;
- code-generation framework;
- first AHDK language;
- package layout;
- RPC transport;
- runtime framework;
- security/policy engine;
- scaffolder UI;
- public SDK support;
- multi-language SDK roadmap.

## 3. Context

Most early Aurora Harnesses are expected to be first-party and built by Leandro with AI assistance, eventually including work initiated or scaffolded by Aurora itself.

Repeated concerns from discovery:

- manually implementing manifests, lifecycle, telemetry and errors in every Harness invites divergence;
- direct low-level integration makes it easy for an agent to omit authority, cancellation, budgets, evidence or recovery;
- Aurora should be able to create new Harnesses through a high-level, safe Golden Path rather than starting from an empty folder;
- a consistent abstraction improves review, testability and future self-improvement;
- however, making a library the specification would lock semantics to one language and make the SDK impossible to test independently;
- an SDK running inside a provider process is not a security boundary.

The desired shape is:

```text
Canonical Contract Model
→ language-neutral source schemas
→ generated types/API
→ official AHDK implementation
→ black-box Conformance Kit
```

## 4. Problem

Without a first-party development kit:

- every provider invents manifest fields and defaults;
- lifecycle transitions are manually serialized;
- events and trace attributes drift;
- effect calls bypass common receipts and idempotency;
- Context Pack handling exposes too much data;
- cancellation and checkpoint behavior differ;
- artifacts are returned as arbitrary text;
- security-sensitive omissions become easy;
- future Aurora-generated code repeats mistakes;
- review focuses on boilerplate rather than domain logic.

If the SDK is universally mandatory at the protocol level:

- external/legacy providers become difficult to integrate;
- non-supported languages are excluded;
- contract meaning becomes coupled to library behavior;
- an SDK bug can be mistaken for the specification;
- independent conformance testing becomes weaker;
- migrating the SDK may require rewriting providers even when contracts are stable.

## 5. Decision drivers

1. make the correct first-party path the easiest path;
2. prevent semantic drift and missing cross-cutting behavior;
3. support Aurora-assisted Harness creation;
4. preserve language/protocol neutrality of contracts;
5. allow external and legacy providers;
6. test the SDK itself through independent conformance;
7. instrument every provider by default;
8. expose authority/effects safely without pretending the SDK enforces them;
9. allow deterministic and agentic Harnesses;
10. start with one proved SDK, not three incomplete SDKs;
11. preserve YAGNI and remove unused modules;
12. support versioning, waivers and migration.

## 6. Alternatives

### Option A — One mandatory runtime and SDK for all providers

Every Harness must use one Aurora runtime/library.

#### Advantages

- strongest uniformity;
- easiest central tooling;
- one support matrix;
- fewer adapters.

#### Disadvantages

- framework/language lock-in;
- poor fit for firmware, deterministic services and external systems;
- makes internal runtime a constitutional dependency;
- difficult adoption of existing Harnesses;
- conflates contract and implementation.

#### Assessment

Rejected.

---

### Option B — SDK optional for every provider

Contracts exist, but first-party teams may implement directly.

#### Advantages

- maximum technical freedom;
- lower initial SDK investment;
- simple experimentation.

#### Disadvantages

- first-party integrations become handcrafted;
- duplicated bugs and inconsistent semantics;
- weak Golden Path for Aurora-generated Harnesses;
- conformance failures are discovered late;
- common observability/security/error behavior requires repeated review.

#### Assessment

Rejected as organizational policy.

---

### Option C — SDK is the specification

The library's types and behavior define the protocol/domain.

#### Advantages

- no separate schema authoring;
- quickest typed implementation in one language;
- fewer artifacts.

#### Disadvantages

- impossible to distinguish library defect from product rule;
- external languages must reverse-engineer behavior;
- documentation and generated types can diverge silently;
- independent implementation/conformance is weakened;
- migration is tightly coupled.

#### Assessment

Rejected.

---

### Option D — AHDK mandatory by first-party policy; contracts and conformance universal

First-party Harnesses must use the official development kit unless an explicit waiver is approved. External/legacy providers may use adapters or direct bindings. All providers must pass the same relevant conformance suite.

#### Advantages

- strong internal Golden Path;
- contract/implementation separation;
- external interoperability;
- SDK remains testable and replaceable;
- lower boilerplate and safer Aurora-assisted generation;
- consistent telemetry, errors and lifecycle.

#### Disadvantages

- AHDK is a product to maintain;
- conformance suite also requires maintenance;
- first-language choice can influence architecture;
- waiver governance adds process;
- possible hidden coupling between SDK convenience API and contract model.

#### Assessment

Selected as the A0 policy direction. Implementation readiness remains subject to SPK-001 and the applicable ACRM gates.

## 7. Decision

### Universal requirements

Every Harness provider must:

- implement an approved Aurora contract/binding version;
- publish or map to a valid Capability Manifest;
- pass applicable black-box conformance profiles;
- expose required lifecycle, error, cancellation, recovery, artifact and evidence behavior;
- respect authority/effect boundaries;
- propagate required trace/correlation identity;
- bind trust to exact provider/build/environment identity.

### First-party policy

A first-party Harness **MUST** use the official Aurora Harness Development Kit unless an explicit waiver is accepted.

“First-party” includes Harnesses:

- maintained in Projeto Aurora or controlled repositories;
- created specifically for Leandro/Aurora;
- generated or materially modified by Aurora under project ownership;
- treated operationally as trusted internal platform components.

### External/legacy policy

An external or legacy provider **MAY** integrate through:

- an adapter;
- a direct protocol implementation;
- MCP/A2A/native RPC/HTTP or another accepted binding;
- another SDK;
- a deterministic device/service bridge.

It receives no reduced conformance, trust or authority requirements because it does not use AHDK.

## 8. AHDK is a development kit

The AHDK is broader than a client library.

Proposed modules:

```text
Contract-generated types
Manifest builder
Provider bootstrap/runtime adapter
Delegation lifecycle API
Context Pack reader and data-policy guards
Decision/escalation API
Budget/deadline/cancellation API
Checkpoint/resume API
Effect request client
Credential-reference client
Artifact/Observation/Claim/Evidence API
Error taxonomy and retry helpers
OpenTelemetry semantic instrumentation
Local simulator and fake Aurora Core
Mocks and test fixtures
Fault injection
Conformance runner/client
Compatibility and version report
Build metadata/provenance hooks
Scaffolder and Golden Path templates
Documentation/reference generation
```

Not every module is implemented initially. SPK-001 determines the minimal first slice.

## 9. Conceptual provider experience

```ts
const provider = defineHarness({
  id: "research-harness",
  version: "0.1.0",
  capabilities: [technicalResearchV1],
  effects: ["public_web.read", "artifact.write"],
});

provider.handle(technicalResearchV1, async ({ delegation, context, effects }) => {
  delegation.reportProgress({ phase: "source-discovery" });

  const result = await effects.request({
    action: "public_web.read",
    resource: sourceRef,
  });

  const artifact = await delegation.publishArtifact({
    kind: "research-report",
    content: buildReport(result),
  });

  await delegation.recordEvidence({
    criterion: "CRIT-SOURCE-PROVENANCE",
    artifacts: [artifact.ref],
  });

  return delegation.complete({ artifacts: [artifact.ref] });
});
```

This example is illustrative, not an accepted TypeScript API or language choice.

The high-level API should make required behavior natural while still allowing domain-specific internal workflows.

## 10. Contract/SDK independence

The canonical source must exist outside any one generated SDK package.

Required separation:

| Layer | Owns |
|---|---|
| Contract Model/specification | semantic meaning and invariants |
| Source schemas | machine-readable representation and compatibility |
| Generated API/types | language projection |
| AHDK implementation | official first-party ergonomics and defaults |
| Provider adapter | protocol/runtime mapping |
| Conformance Kit | black-box observed behavior |
| Policy/Gateway/Sandbox | security decision and enforcement |

AHDK versions and Contract Model versions may evolve independently within an explicit compatibility matrix.

## 11. Conformance profiles

Conformance should be capability/profile aware rather than one universal all-or-nothing suite.

Potential profiles:

```text
CORE
→ identity, manifest, lifecycle, errors and correlation

SYNC_PROVIDER
→ bounded request/response behavior

ASYNC_PROVIDER
→ progress, wait, cancellation and terminal state

RESUMABLE_PROVIDER
→ checkpoint/restart/reconciliation

EFFECTFUL_PROVIDER
→ authority denial, effect receipts and idempotency

STREAMING_PROVIDER
→ ordering, reconnect and recovery of critical state

ARTIFACT_PROVIDER
→ hashes, provenance, retention and schema

DEVICE_PROVIDER
→ pairing, telemetry, safety and fail-safe behavior
```

A provider passes only profiles it declares/needs. A declaration without passing evidence does not create trust.

## 12. Conformance examples

### Lifecycle

```text
AUTHORIZED
→ QUEUED
→ RUNNING
→ COMPLETED
```

Reject:

- completion before start;
- event after terminal state without a new run identity;
- success with missing required artifacts;
- silent state loss after reconnect.

### Authority

```text
allowed effect
→ executed through gateway
→ receipt linked to Delegation

denied effect
→ no external change
→ structured denial event
```

### Cancellation

```text
cancel requested
→ provider acknowledges
→ stops new effects
→ reports cleanup/reconciliation
→ terminal CANCELED or explicit failure
```

### Recovery

```text
checkpoint
→ provider process killed
→ restart with same Delegation/run
→ no duplicate external effect
→ artifact/evidence continuity
```

### Context isolation

A provider cannot read project/global memories outside the Context Pack/grants delivered for the Delegation.

## 13. Security boundary

AHDK is not a sandbox and cannot prevent a provider process from calling operating-system APIs directly.

```text
AHDK
→ correct API and audit context

Policy Decision Point
→ authorization decision

Effect Gateway / Credential Broker
→ enforcement and secret handling

Sandbox / OS / device boundary
→ containment

Receipt / trace
→ observable proof
```

A Harness that bypasses AHDK must still be blocked by external enforcement when attempting an unauthorized effect.

A provider using AHDK is not automatically trusted. Trust is tied to exact build, environment, conformance and approved scope.

## 14. Golden Path and scaffolder

The intended creation flow is:

```text
select capability archetype
→ answer domain/effect/data questions
→ generate manifest and source schemas
→ generate AHDK bootstrap
→ include telemetry and error conventions
→ include unit/contract/conformance tests
→ include simulator configuration
→ include CI and provenance hooks
→ run dry-run/sandbox
→ register as candidate provider
```

Conceptual command:

```bash
aurora harness create research-harness
```

The command/interface is not accepted yet. The requirement is a governed scaffolding path.

Templates must not create broad permissions or unused modules by default.

## 15. Waiver model

A first-party provider may bypass AHDK only when a waiver records:

```yaml
id: WAIVER-AHDK-...
provider: ...
reason: ...
scope: ...
owner: ...
risk: ...
compensating_controls: []
conformance_profiles: []
contract_version: ...
expires_at: ...
removal_condition: ...
approver: ...
```

Valid reasons may include:

- unsupported language/device environment;
- verified performance or footprint constraint;
- integration with an existing mature system;
- architecture spike requiring independent implementation.

Invalid reasons:

- convenience;
- avoiding tests;
- undocumented speed of delivery;
- bypassing authority/effect APIs.

Waivers are narrow, time/condition-bound and discoverable in provider trust assessment.

## 16. Version and compatibility policy

The ecosystem must represent separately:

- Contract Model version;
- capability contract version;
- manifest schema version;
- AHDK version;
- binding/protocol version;
- provider source/build digest;
- conformance suite/profile version;
- Aurora Core compatibility.

A text version alone does not grant trust.

Breaking changes require:

- machine-detected compatibility report where possible;
- migration path;
- provider revalidation;
- updated generated types/docs;
- trust/approval review for material changes.

## 17. Implementation constraints under the accepted policy

1. Build one AHDK first; no premature TypeScript/Python/Go parity.
2. Generate repeated types from canonical schemas where practical.
3. Keep AHDK core small; domain-specific helpers live outside the universal core.
4. Provide a no-network/local simulator before a real provider.
5. Conformance runs against the public provider boundary, not internal methods.
6. AHDK may request effects but never make authorization decisions.
7. Default templates deny undeclared effects and data classes.
8. Telemetry and errors are enabled by default with privacy/redaction controls.
9. Every helper must map to an explicit Contract Model concept or demonstrated Golden Path need.
10. Remove speculative modules before graduation.

## 18. Validation plan

### SPK-001 — Same capability through two implementations

Build a minimal artifact-transformation capability:

- implementation A: official AHDK;
- implementation B: direct binding without AHDK.

Both must pass identical:

- manifest/schema tests;
- lifecycle transitions;
- errors;
- cancellation;
- artifact/evidence output;
- correlation/telemetry;
- authority denial;
- restart behavior appropriate to profile.

Measure:

- implementation size/complexity;
- defects/omissions;
- conformance clarity;
- runtime overhead;
- debugging experience;
- migration/compatibility effort.

### Scaffolder proof

Create a new provider through the Golden Path and demonstrate that a fresh implementation starts with required contracts, tests, telemetry and least-authority defaults.

### SDK bypass security proof

Attempt direct filesystem/network/credential access outside AHDK and prove external enforcement denies or contains it.

## 19. Implementation-readiness evidence

Before advancing this accepted A0 policy into implementation planning:

- ADR-0001 principle is accepted or compatible;
- AHDK-focused research is current;
- SPK-001 has an approved design and authorization path;
- the Conformance Kit has a minimal independent architecture;
- waiver ownership is clear;
- first Product Milestone identifies an actual first-party consumer;
- removal conditions exist for unused abstractions.

A0 accepts the policy direction without selecting a language or implementing the kit; implementation readiness remains a later gate.

## 20. Consequences

### Positive

- consistent first-party provider experience;
- fewer omitted lifecycle/security/evidence requirements;
- easier Aurora-assisted Harness creation;
- contract-driven code generation and documentation;
- uniform observability and errors;
- provider/runtime independence preserved;
- external systems remain integrable;
- conformance gives an objective trust input.

### Negative

- AHDK and Conformance Kit are ongoing products;
- additional release/version matrix;
- early language choice influences developer experience;
- templates can ossify bad assumptions;
- runtime overhead is possible;
- waiver governance requires discipline.

## 21. Risks and mitigations

| Risk | Mitigation |
|---|---|
| SDK becomes de facto specification | canonical schemas/spec; direct implementation in SPK-001 |
| false security | external policy/gateways/sandbox; bypass tests |
| framework bloat | minimal core; profile-based modules; removal conditions |
| one-language lock-in | language-neutral contracts; adapters; defer second SDK until need |
| scaffolder creates cargo-cult code | archetype-specific templates, dry-run and generated rationale |
| conformance only tests happy path | cancellation, faults, denied effects, restarts and malformed input |
| breaking SDK update invalidates providers | compatibility matrix, generated reports and staged migration |
| Aurora-generated Harness self-approves | independent conformance and required reviewer/verdict |

## 22. Reconsideration triggers

Re-open or supersede if:

- most first-party providers require waivers;
- AHDK adds more code/cognitive load than direct contracts;
- conformance cannot detect meaningful divergence;
- one language/runtime genuinely covers all proven provider domains and a simpler policy is justified;
- device/firmware providers cannot participate without a separate kit/profile;
- SDK release burden blocks Product Milestones;
- an open standard provides a mature equivalent Golden Path without loss of Aurora semantics.

## 23. Decision record on acceptance

If accepted, record:

- policy scope and definition of first-party;
- exact waiver authority;
- which AHDK modules are principles versus future candidates;
- the first consumer Product Milestone;
- separately authorized SPK-001 scope;
- explicit statement that the ADR does not choose language or authorize implementation.
