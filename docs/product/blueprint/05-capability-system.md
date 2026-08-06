---
id: DOC-AURORA-BLUEPRINT-05
title: Sistema de Capabilities, Registry e Development Kit
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
  - capability model
  - provider and registry model
  - discovery, compatibility, trust, authority and execution separation
  - first-party AHDK policy
  - universal conformance principles
related:
  - DOC-AURORA-BLUEPRINT-03
  - DOC-AURORA-BLUEPRINT-07
  - DOC-AURORA-BLUEPRINT-10
  - DOC-AURORA-BLUEPRINT-11
  - DOC-AURORA-BLUEPRINT-12
  - DOC-AURORA-BLUEPRINT-13
  - ADR-AURORA-0001
  - ADR-AURORA-0002
  - RESEARCH-AURORA-HARNESS-ARCHITECTURE-V1
review_triggers:
  - capability, provider or registry semantic changes
  - first-party AHDK policy changes
  - trust lifecycle changes
  - conformance boundary changes
last_reviewed: 2026-08-05
---

# 5. Sistema de Capabilities, Registry e Development Kit

## 5.1 Propósito

Aurora will need to use systems created at different times, in different languages and for very different domains:

- research;
- software engineering;
- hardware analysis;
- firmware build/flash/debug;
- laboratory observation;
- device control;
- evaluation;
- memory consolidation;
- personal operations;
- external services.

Without a capability model, every integration becomes a tool name or bespoke prompt. Aurora would need to remember implementation details, infer permissions from availability and manually adapt context to each provider.

The Capability System creates a stable boundary between:

```text
what outcome is needed
and
which concrete system can produce it now
```

It must answer:

- What can Aurora request?
- Which providers claim to offer it?
- Which contract/version do they implement?
- What effects and data do they require?
- What has actually been verified?
- For which scope are they approved?
- Which instance/build is available?
- Why was one provider selected over another?
- How is trust suspended, revoked or migrated?
- How can Aurora create a new first-party harness without rebuilding integration plumbing incorrectly?

---

## 5.2 Canonical vocabulary

| Concept | Meaning | Not equivalent to |
|---|---|---|
| Tool | bounded operation | harness, authority, mission |
| Agent | probabilistic reasoning unit | Aurora identity, provider approval |
| Workflow | sequence/graph of steps | product mission semantics |
| Runtime | executes agents/workflows | capability definition |
| Harness | complete specialized domain system | one tool or one agent |
| Capability | reusable outcome Aurora can request | provider implementation |
| Provider | entity offering one or more capabilities | active execution |
| Provider Instance | exact running/build/environment instance | provider family |
| Manifest | provider's structured declaration | proof or authority |
| Binding | protocol/transport mapping | canonical contract semantics |
| Adapter | translation from Aurora contracts | authority owner |
| Registry | provider/capability/trust/approval catalog | arbitrary service discovery |
| Verification | evidence from inspection/test | universal trust |
| Approval | allowed scope for an exact provider/build | delegation |
| Delegation | concrete assigned work | capability or provider |
| AHDK | first-party development Golden Path | canonical specification |
| Conformance Kit | black-box behavior verification | self-declaration |

### Canonical invariant

> Capability describes **what** reusable result can be delivered. Provider/Harness describes **who** can deliver it. Delegation describes **which concrete work** was assigned under **which authority**.

---

## 5.3 Capability identity

A capability has stable semantic identity and version.

Examples:

```text
capability://aurora/technical-research/1
capability://aurora/software-delivery/1
capability://aurora/firmware-build-flash/1
capability://aurora/laboratory-measurement/1
capability://aurora/context-evaluation/1
```

Exact naming is not yet selected. Required properties:

- globally unambiguous within the Aurora instance/ecosystem;
- versioned independently from provider implementation;
- associated with input/output/event/evidence semantics;
- capable of compatibility/supersession;
- not tied to one framework or language.

### Capability Definition

A future definition may own:

- purpose;
- use cases;
- input contract;
- result contract;
- allowed effect categories;
- required evidence profile;
- expected lifecycle patterns;
- error taxonomy;
- compatibility;
- evaluation and graduation.

A Provider Manifest references a Capability Definition. It should not silently redefine it.

---

## 5.4 Provider and Instance

### Provider

A logical system offering capabilities.

Examples:

```text
research-harness
mnfs
firmware-lab-harness
public-web-search-service
local-memory-consolidator
```

### Provider Instance

A concrete deployable/running instance with exact:

- provider version;
- source revision;
- build digest;
- configuration/profile;
- runtime environment;
- binding endpoints;
- identity/attestation;
- operational health.

Trust and approval bind to the exact identity needed by risk.

```text
Provider research-harness v1.2.0
├── Instance local-01, build A, local-restricted
└── Instance cloud-01, build B, external-network
```

The two instances can have different data and effect approvals.

---

## 5.5 Five independent stages

```text
DISCOVERY
→ Aurora knows the provider/capability exists

COMPATIBILITY
→ Aurora can understand and map its contracts

TRUST
→ evidence exists about identity and observed behavior

AUTHORITY
→ use is approved for a specific scope

EXECUTION
→ a concrete Delegation is active
```

### Example

```yaml
provider: external-research-harness
capability: technical-research

discovery: known
compatibility: contract-v1-supported
trust:
  identity: verified
  functional: verified_on_public_research_suite
  security: partial
  recovery: unverified
authority:
  allowed_data_max: PUBLIC
  allowed_effects:
    - read_public_web
    - publish_research_artifact
  prohibited:
    - read_personal_memory
    - access_credentials
    - modify_repository
execution: none
```

Knowing how to call the provider does not authorize sending context.

---

## 5.6 Capability Manifest

A Provider publishes a structured declaration.

### 5.6.1 Complete conceptual example

```yaml
manifest:
  schema_version: aurora-provider-manifest/1

  identity:
    provider_id: research-harness
    provider_version: 1.2.0
    publisher: developmentconexus-ops
    source_revision: git:abc123
    build_digest: sha256:...
    provenance_ref: attestation://...
    instance_id: research-local-01
    environment_profile: local-restricted

  capabilities:
    - capability_id: technical-research
      capability_version: "1"
      input_schema: schema://technical-research-request/1
      output_schema: schema://research-report/1
      event_schemas:
        - schema://delegation-progress/1
        - schema://decision-request/1
      evidence_profile:
        - source_manifest
        - claim_source_mapping
        - limitation_report

  interaction:
    modes:
      - synchronous
      - asynchronous
      - streaming
    resumable: true
    cancellable: true
    heartbeat: supported
    idempotency: delegation-and-effect
    maximum_concurrency: 3

  effects:
    declared:
      - read_public_web
      - write_artifact
      - use_approved_model
    prohibited:
      - repository_write
      - external_message
      - device_control

  data:
    required_classes:
      - PUBLIC
      - INTERNAL
    optional_classes:
      - CONFIDENTIAL
    secrets: references_only
    retention_profile: local-project-default

  runtime:
    protocols:
      - native-rpc
      - a2a
    network:
      - approved_web_domains
    compute_profile: standard

  recovery:
    checkpoint: durable
    cancellation: cooperative
    snapshot_query: supported

  observability:
    opentelemetry: supported
    trace_propagation: w3c
    audit_receipts: supported

  compatibility:
    aurora_contract_model: ">=1.0 <2.0"
    ahdk_version: "1.1.0"
```

The exact syntax is not selected. The example protects required concepts.

### 5.6.2 Manifest validation

Validation layers:

1. schema syntax;
2. stable IDs/version;
3. referenced contract availability;
4. build/provenance integrity;
5. declared effect/data consistency;
6. protocol compatibility;
7. publisher/instance identity;
8. no forbidden ambiguity for material fields.

### 5.6.3 Manifest limitations

A manifest is a claim.

It cannot prove:

- provider does not access undeclared network paths;
- cancellation truly works;
- results are correct;
- build matches source;
- context is deleted;
- recovery is safe;
- effect boundaries are respected.

Verification and enforcement remain separate.

---

## 5.7 Capability Registry

The Registry stores the current and historical relationship between:

```text
Capability
→ Provider
→ Provider Instance
→ Manifest
→ Compatibility Assessment
→ Verification
→ Trust Assessment
→ Approval
→ Incident
→ Delegation
```

### 5.7.1 Registry responsibilities

- index capability definitions;
- ingest manifests;
- preserve exact versions/builds;
- track bindings/endpoints;
- store compatibility results;
- link conformance and sandbox evidence;
- record trust dimensions;
- store approvals by scope;
- expose health/availability;
- suspend/revoke/retire;
- provide provider selection candidates;
- preserve incident and provenance history.

### 5.7.2 Registry does not

- grant authority merely because an entry exists;
- guarantee runtime availability;
- replace policy enforcement;
- store every provider internal state;
- become a package marketplace in current scope;
- reduce trust to one score.

---

## 5.8 Trust lifecycle

Initial lifecycle:

```text
DISCOVERED
→ INSPECTED
→ SANDBOX_VALIDATED
→ CONFORMANCE_VERIFIED
→ TRUST_ASSESSED
→ APPROVED_FOR_SCOPE
→ SUSPENDED | REVOKED | RETIRED
```

### DISCOVERED

Manifest/endpoint found. No material execution.

### INSPECTED

Identity, contracts, requested data/effects and environment analyzed.

### SANDBOX_VALIDATED

Provider executed with synthetic/non-sensitive data and constrained effects.

### CONFORMANCE_VERIFIED

Black-box tests demonstrate Aurora contract behavior for a version/profile.

### TRUST_ASSESSED

Evidence summarized by dimension and limitation.

### APPROVED_FOR_SCOPE

Authority may allow concrete Delegations within specific boundaries.

### SUSPENDED

Temporarily blocked because of incident, stale evidence, incompatible environment or availability concern.

### REVOKED

Approval/trust removed; active Delegations are contained/reconciled according to policy.

### RETIRED

Historical provider remains discoverable but is not selected for new work.

No state change occurs only from text description. Transitions need actor, reason and evidence.

---

## 5.9 Multidimensional trust

Trust dimensions include:

### Identity

Is publisher/instance who it claims?

### Build integrity

Does running build match source/provenance?

### Contract compatibility

Does it implement the expected semantic version?

### Functional correctness

Does it produce valid outcomes on representative evaluations?

### Security

Does it respect data/effect/environment boundaries under test?

### Recovery

Can it cancel, resume and reconcile?

### Evidence quality

Are claims supported by artifacts and methods?

### Operational reliability

Latency, availability, cost and error behavior.

### Domain scope

Which problem classes were actually tested?

Example:

```text
Provider can be:
- strong functional trust for public technical research;
- weak privacy trust for confidential project context;
- unverified recovery trust;
- prohibited for repository effects.
```

A universal `trusted=true` is constitutionally invalid.

---

## 5.10 Verification profiles

Verification should be scoped and repeatable.

### Contract

- schema;
- state transitions;
- error behavior;
- version negotiation.

### Functional

- representative task suite;
- expected artifacts;
- quality rubric;
- limitations.

### Security

- undeclared effect attempts;
- context exfiltration;
- credential handling;
- SDK bypass attempt;
- malicious input.

### Operational

- restart;
- timeout;
- duplicate messages;
- cancellation;
- resource limits;
- degraded dependencies.

### Supply chain

- digest;
- source revision;
- provenance;
- signature;
- dependency/build profile.

Verification result must declare what it did not test.

---

## 5.11 Approval scope

An approval may constrain:

- capability;
- provider/build/instance;
- projects/domains;
- data classes;
- effects;
- environments;
- model/providers usable internally;
- network;
- maximum budget;
- time window;
- concurrency;
- evidence requirements;
- human confirmation points.

Example:

```yaml
approval:
  provider_instance: research-local-01@sha256:...
  capability: technical-research/1
  projects:
    - PRJ-AURORA
  data_max: CONFIDENTIAL
  effects_allow:
    - read_approved_web
    - write_research_artifact
  effects_deny:
    - repository_write
    - external_message
    - credential_use
  environment: local-restricted
  valid_until: 2026-11-01
  evidence_profile: research-primary-sources-v1
```

Approval is not a Delegation. Each concrete job still receives objective, context, authority and budget.

---

## 5.12 Provider selection

Aurora first resolves the capability, then candidates.

Factors:

```text
capability fit
+ contract compatibility
+ approval scope
+ data sensitivity
+ effect requirements
+ functional evidence
+ environment
+ availability
+ cost
+ latency
+ recovery
+ current incidents
+ user preference
```

### Selection explanation

Material selection can produce:

```yaml
selection:
  capability: technical-research/1
  selected: research-local-01
  reasons:
    - approved for CONFIDENTIAL project context
    - complete primary-source evidence profile
    - local availability
  rejected_candidates:
    - provider: research-cloud-02
      reason: data approval only PUBLIC
    - provider: generic-agent-01
      reason: no conformance for artifact/evidence contract
  fallback: research-local-02
```

Selection should not expose unnecessary internal ranking noise to Leandro unless relevant.

---

## 5.13 Fallback and substitution

Fallback may occur when:

- provider unavailable;
- budget exhausted;
- capability mismatch discovered;
- incident/suspension;
- environment changes;
- data sensitivity changes.

Substitution rules:

- new provider must satisfy contract and authority;
- context is rebuilt/minimized for new provider;
- prior local state is not assumed compatible;
- exact difference is recorded;
- evidence may need revalidation;
- material quality/cost change is surfaced.

A provider cannot silently delegate to an unapproved external provider if that changes data/effect boundaries.

---

## 5.14 Compatibility and versioning

Separate versions:

- capability semantic version;
- provider version;
- manifest schema;
- Aurora Contract Model;
- binding protocol;
- AHDK;
- event/artifact schemas.

Compatibility rules must distinguish:

- additive optional field;
- new capability version;
- breaking schema;
- changed effect declaration;
- changed evidence semantics;
- changed recovery behavior;
- changed data handling.

Material changes trigger reinspection or reverification.

### Example

Provider v1.2 → v1.3 changes only internal model and passes same eval/conformance. Approval may be promoted under policy.

Provider v1.3 → v2.0 adds external network and credential requirement. Previous approval cannot be inherited.

Breaking-change detection should be mechanical where possible.

---

## 5.15 Incidents and trust downgrade

Incidents can affect:

- one instance;
- one build;
- one capability;
- one environment;
- the provider family;
- a binding/SDK version.

Example:

```text
Incident: provider leaked confidential context to general logs
→ suspend confidential-data approval
→ allow public sandbox only if safe
→ revoke active credentials
→ preserve evidence
→ investigate SDK/logger version
→ decide provider/build scope of impact
```

Registry status must drive actual selection/authority, not merely display a warning.

---

## 5.16 Provider retirement and migration

Retirement plan includes:

- replacement capability/provider;
- active Delegation policy;
- artifact/state export;
- contract adapter;
- trust/incident history preservation;
- final date;
- rollback window;
- documentation update.

Historical provider remains discoverable to explain old artifacts and decisions.

---

# 5.17 Aurora Harness Development Kit (AHDK)

## 5.17.1 Purpose

AHDK turns the canonical Contract Model into a safe and productive first-party engineering path.

Leandro's requirement is explicit:

> Aurora and human developers should create new harnesses through high-level abstractions and approved Golden Paths, not by repeatedly hand-coding protocol, lifecycle, telemetry and authority plumbing.

First-party harnesses **must use AHDK by engineering policy**, unless an explicit waiver is approved.

Technical interoperability remains open through contracts/adapters so the SDK does not become the specification.

---

## 5.18 AHDK architecture

```text
Canonical Contract Model
        ↓
Schema/code generation
        ↓
AHDK public APIs
├── Provider/Manifest
├── Delegation Lifecycle
├── Context Reader
├── Authority/Effect Client
├── Artifact/Evidence
├── Decision Requests
├── Budget/Deadline
├── Checkpoint/Recovery
├── Telemetry
└── Test/Simulator
        ↓
Binding adapters
Native • RPC • A2A • MCP • HTTP/Event
```

### Package families

Exact packages are future design, but responsibilities include:

- contracts-generated;
- manifest-builder;
- provider-runtime;
- lifecycle;
- context;
- effects;
- artifacts-evidence;
- observability;
- testkit;
- scaffolder;
- conformance;
- provenance.

---

## 5.19 Generated contracts

AHDK should generate or expose consistent types from canonical schemas.

Benefits:

- fewer spelling/type divergences;
- multi-language potential;
- shared validation;
- version compatibility;
- documentation generation;
- easier Aurora-created harnesses.

Generated code must be reproducible and declare:

- source schema version;
- generator version;
- output language/version;
- digest.

Developers should not edit generated files directly.

---

## 5.20 Manifest Builder

Instead of handcrafted YAML/JSON:

```ts
const provider = defineHarness({
  id: "research-harness",
  version: "1.2.0",
  publisher: "developmentconexus-ops",
  capabilities: [
    defineCapability({
      id: "technical-research",
      version: "1",
      input: TechnicalResearchRequest,
      output: ResearchReport,
      evidence: [SourceManifest, ClaimSourceCoverage],
    }),
  ],
  effects: allow("read_public_web", "write_artifact"),
});
```

The builder can:

- validate IDs;
- insert contract references;
- require effect/data declarations;
- apply safe defaults;
- generate manifest;
- calculate metadata/digest;
- reject incomplete recovery/evidence definitions.

This code is illustrative, not a selected API/language.

---

## 5.21 Provider Runtime API

A provider handler receives a typed Delegation context.

```ts
provider.handle("technical-research", async (delegation) => {
  const request = delegation.input();
  const context = await delegation.context.getAuthorizedSources();

  await delegation.progress({ phase: "source-discovery" });

  const artifact = await delegation.artifacts.publish(...);
  await delegation.evidence.record(...);
  return delegation.complete({ artifact });
});
```

The runtime should manage:

- lifecycle validation;
- correlation IDs;
- cancellation signal;
- deadlines;
- heartbeats;
- telemetry;
- error mapping;
- cleanup hooks.

Provider domain code should not mutate Aurora state directly.

---

## 5.22 Context API

AHDK should expose only the Context Pack granted to the Delegation.

Possible high-level access:

```ts
const project = delegation.context.projectSnapshot();
const decisions = delegation.context.authoritySources();
const memories = delegation.context.memories({ scope: "delegation" });
const artifact = await delegation.context.openArtifact(ref);
```

Context API must preserve:

- provenance;
- sensitivity;
- freshness;
- read audit where required;
- reference rather than secret exposure;
- no unrestricted global memory query.

A provider cannot request “all Leandro memory” through convenience API.

---

## 5.23 Authority and Effect Client

Illustrative request:

```ts
const receipt = await delegation.effects.request({
  effect: "repository.write_file",
  resource: "repo://project/path/file.ts",
  inputArtifact: patchRef,
  idempotencyKey: "...",
});
```

The client:

- builds a typed Effect Request;
- propagates actor/delegation context;
- sends to the Effect Gateway;
- returns a Receipt;
- does not decide permission;
- does not expose raw credentials.

Direct filesystem/network APIs may still exist in the process; sandbox/environment enforcement must prevent bypass for material scopes.

---

## 5.24 Artifact and Evidence API

Illustrative distinctions:

```ts
const report = await artifacts.publish({
  type: "research-report",
  content: ...,
  classification: "INTERNAL",
});

await evidence.record({
  criterion: "CRIT-SOURCE-COVERAGE",
  artifacts: [report, sourceManifest],
  method: "claim-source-crosscheck",
  result: "PASS",
  limitations: [...],
});
```

AHDK should make it difficult to return only `{ success: true }` when a capability contract requires evidence.

---

## 5.25 Decision Request API

A provider can escalate a material choice:

```ts
await delegation.requestDecision({
  question: "Use external provider for confidential source?",
  alternatives: [...],
  recommendation: "remain local",
  impact: {...},
  blocks: ["next-research-step"],
});
```

Decision Request must not grant the provider a new permission by waiting. Aurora/Leandro responds through a new decision/grant or denies.

---

## 5.26 Budget and deadline API

Provider should query and receive signals:

- consumed;
- remaining;
- soft threshold;
- hard limit;
- extension pending;
- deadline/cancel signal.

Budgets may include money, tokens, runs, compute, storage, energy and device cycles.

AHDK can standardize accounting reports while authoritative enforcement remains in Core/gateway/environment.

---

## 5.27 Checkpoint, resume and cancellation

AHDK should provide abstractions for:

- durable checkpoint reference;
- local provider snapshot;
- resume token/version;
- cooperative cancellation;
- cleanup;
- already-produced effects;
- heartbeat;
- lost-connection reconciliation.

Provider declares semantics; Conformance Kit verifies.

Example:

```text
checkpoint created after variant 12
→ provider process killed
→ new instance resumes exact build/contract
→ completed variants not repeated
→ previous external effects reconciled
```

---

## 5.28 Automatic observability

AHDK should instrument by default:

- provider run;
- capability invocation;
- lifecycle transitions;
- context access metadata;
- decision requests;
- effect requests/receipts;
- artifact/evidence;
- checkpoints;
- errors/recovery;
- budget consumption.

Sensitive payloads remain excluded/redacted by default.

Provider-specific telemetry can extend Aurora semantic conventions.

---

## 5.29 Simulator and Test Kit

A local simulator should provide:

- fake Aurora Core;
- synthetic Context Pack;
- scoped grants;
- fake Effect Gateway;
- artifact/evidence stores;
- decision responses;
- budget/deadline;
- clock control;
- restart/cancel;
- duplicate/out-of-order events;
- network failure;
- malicious input;
- trace collector.

This is essential for Aurora herself to build and verify providers safely.

---

## 5.30 Scaffolder and Golden Paths

Command concept:

```text
aurora harness create research-harness
```

Generated project should include:

```text
manifest definition
capability handlers
contract versions
AHDK bootstrap
observability
policy/effect declarations
tests
conformance configuration
fault tests
README
CI
ownership
provenance hooks
```

Golden Path should ask domain questions, not only choose a template:

- What outcome does the capability own?
- Which data/effects are required?
- What evidence closes it?
- Can it cancel/resume?
- What remains local state?
- What is the safe sandbox profile?
- What is explicitly not supported?

---

## 5.31 AHDK waiver

A first-party waiver requires:

- material reason;
- owner;
- exact scope;
- alternative implementation path;
- security/operational impact;
- complete Conformance Kit pass;
- maintenance responsibility;
- expiry or removal condition;
- accepted ADR/waiver record when material.

Valid future reasons may include:

- embedded/firmware target cannot run SDK;
- language/runtime lacks official AHDK;
- integrating legacy provider;
- performance/real-time boundary;
- security isolation requires a minimal adapter.

“Faster to hand-code” is not sufficient.

---

# 5.32 Universal Conformance Kit

## 5.32.1 Purpose

Conformance proves externally observable behavior independent of implementation.

All providers/adapters—AHDK or direct—must pass the applicable profile.

### Profiles

- manifest/contract;
- synchronous provider;
- asynchronous provider;
- resumable provider;
- effect-producing provider;
- streaming/data-channel provider;
- device provider;
- remote A2A;
- MCP tool/resource adapter.

### Test areas

```text
identity and manifest
contract/schema
version negotiation
lifecycle and terminal state
errors
cancel/deadline
resume/checkpoint
idempotency
event duplicate/order behavior
artifact integrity
evidence requirements
authority denial
effect receipt
budget
restart/reconciliation
trace propagation
privacy/redaction
```

A provider passes only the capabilities/profiles tested.

---

## 5.33 Reference providers

Before integrating a complex real harness, Aurora should implement at least two minimal providers.

### Provider A — Artifact Producer

Receives structured text, produces a transformed artifact and evidence.

### Provider B — Artifact Consumer/Evaluator

Consumes an authorized artifact, verifies a criterion and returns a verdict.

They prove:

- provider discovery;
- AHDK ergonomics;
- direct implementation/conformance comparison;
- child Delegation;
- artifact handoff;
- authority isolation;
- restart;
- trace.

Toy providers are architecture probes, not product value proof. M6 introduces real domain behavior.

---

## 5.34 Protocol mapping principles

### Native AHDK/in-process

Best for first-party low-overhead integration in the same language/process, while preserving domain boundary.

### Local RPC

Best for process isolation, independent restart and cross-language local harnesses.

### A2A

Candidate for remote, opaque and stateful agent applications.

### MCP

Candidate for tools/resources and bounded asynchronous operations.

### Direct HTTP/gRPC/events/device protocol

Used when requirements justify it.

Provider may expose several bindings. Aurora contracts remain canonical.

---

## 5.35 Attack and failure modes

### Manifest deception

Provider declares fewer effects/data needs than actual behavior.

Mitigation: sandbox, environment restrictions, gateway enforcement, audit.

### Trust inheritance

New build receives previous approval silently.

Mitigation: build-bound identity and promotion policy.

### SDK monoculture bug

AHDK defect affects many first-party providers.

Mitigation: Conformance Kit independent from SDK; versioned rollout/canary.

### Registry stale state

Suspended provider remains selectable.

Mitigation: transactional status/selection and freshness requirements.

### Capability semantic drift

Providers interpret same capability differently.

Mitigation: canonical spec, evals, versioning, examples.

### Authority smuggling

Provider embeds secondary call using its own credentials.

Mitigation: environment/egress policy, child Delegation, context/effect audit.

### Provider self-certification

Provider emits its own “verified” status.

Mitigation: Verification identity and independent Conformance authority.

### Selection gaming

Provider optimizes cost/latency metadata but quality degrades.

Mitigation: multi-objective eval and current evidence.

### Orphan provider

Instance disappears with active Delegation.

Mitigation: heartbeat, durable state, reconciliation, substitution/stop.

---

## 5.36 Evaluation requirements

Future system must prove:

1. same capability can be implemented with AHDK and direct protocol;
2. both pass the same black-box conformance;
3. AHDK reduces implementation surface without changing semantics;
4. incompatible capability version fails closed;
5. manifest alone cannot produce approval;
6. provider receives only allowed context/effects;
7. new build does not inherit approval automatically;
8. provider selection explains material factors;
9. suspended provider is removed from selection immediately;
10. provider fallback rebuilds context/authority;
11. provider cannot mint child authority;
12. cancellation/restart behavior matches declaration;
13. artifact/evidence contracts reject bare success;
14. trace crosses AHDK/direct implementations consistently;
15. SDK bypass attempt remains contained by environment/gateway;
16. scaffolder produces a conformant project from Golden Path.

---

## 5.37 Open technical decisions

- capability/manifest schema syntax;
- first AHDK language;
- package/module layout;
- code generator;
- Registry storage/query model;
- signing/attestation;
- compatibility rules;
- conformance runner language;
- sandbox profile implementation;
- provider health protocol;
- local RPC;
- A2A/MCP mapping;
- distribution/versioning of AHDK;
- template engine and CLI.

These require focused research and spikes.

---

## 5.38 Non-goals

- public capability marketplace;
- automatic provider installation from internet;
- multi-tenant registry;
- one universal trust score;
- runtime mandatory for every harness;
- three SDK languages before one is proven;
- signing treated as behavioral proof;
- broad peer-to-peer authority;
- manifest interpreted as permission;
- AHDK replacing domain specs;
- bespoke protocol before standards gaps are demonstrated;
- integrating MNFS before its boundary and readiness are suitable.
