---
id: RESEARCH-AURORA-EVENTS-OBSERVABILITY-SCHEMAS-V1
title: Aurora Research — Events, Observability, Schemas and Compatibility
document_type: research_report
form: explanation
authority: research
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - events, telemetry and schema research through 2026-08-05
related:
  - DOC-AURORA-BLUEPRINT-05
  - DOC-AURORA-BLUEPRINT-07
  - DOC-AURORA-BLUEPRINT-12
  - DOC-AURORA-BLUEPRINT-13
source_manifest: AURORA-RESEARCH-EVENTS-OBSERVABILITY-SCHEMAS-v1.sources.json
review_triggers:
  - Contract Model schema selection
  - SPK-006 result
  - event transport selection
last_reviewed: 2026-08-05
---

# Aurora Research — Events, Observability, Schemas and Compatibility

## 1. Research question

Which standards and patterns can support Aurora's need for:

- language-neutral contracts;
- domain events;
- distributed trace correlation;
- artifact/evidence references;
- compatibility and code generation;
- streaming/provider events;
- sensitive-data redaction;
- human-readable manifests;
- optional high-performance RPC?

Candidates:

- OpenTelemetry;
- CloudEvents;
- AsyncAPI;
- JSON Schema 2020-12;
- Protocol Buffers/gRPC;
- Buf compatibility tooling;
- W3C Trace Context.

---

## 2. Executive finding

No single schema/event standard should own every Aurora boundary.

A plausible division to prove is:

```text
JSON Schema
→ human-readable manifests/config/contracts/artifact metadata

CloudEvents
→ common envelope for domain/provider events

AsyncAPI
→ documentation/generation for event channels

Protocol Buffers/gRPC
→ typed high-performance RPC where justified

OpenTelemetry + W3C Trace Context
→ traces, metrics, logs and correlation

Buf/schema checks
→ mechanical compatibility detection

Aurora semantic conventions
→ domain meaning above all formats
```

Observability signals are not domain state or evidence. Events are not necessarily durable state. Schema compatibility does not prove semantic compatibility.

---

## 3. Signal taxonomy

### Domain Event

A durable meaningful change:

- Delegation authorized;
- provider suspended;
- artifact published;
- grant revoked;
- experiment concluded.

### Transport Message

Delivery unit that may be retried, duplicated or reordered.

### Telemetry

Operational signal for diagnosis/measurement:

- trace span;
- metric;
- log.

### Audit Record

Security/authority/data/effect accountability.

### Receipt

Controlled record of an effect/verification.

### Artifact

Material output referenced by events/evidence.

Conflating these creates state and retention errors.

---

## 4. OpenTelemetry

OpenTelemetry defines vendor-neutral APIs/SDKs and data models for traces, metrics and logs [S01][S02].

### Aurora fit

- trace Presence → Core → Harness → Tool/Gateway;
- provider/model/tool latency;
- cost and retry metrics;
- standardized instrumentation in AHDK;
- export to local or external backends;
- correlation through trace/span context;
- framework-specific adapter integration.

### Required Aurora semantic conventions

Examples:

```text
aurora.project.id
aurora.mission.id
aurora.delegation.id
aurora.provider.id
aurora.provider.build_digest
aurora.capability.id
aurora.effect.type
aurora.artifact.id
aurora.evidence.id
aurora.data.classification
aurora.authority.grant_ref
aurora.budget.cost
```

### Boundaries

OpenTelemetry does not own:

- Mission state;
- Authority Grant;
- Artifact content;
- Evidence verdict;
- audit retention;
- provider trust.

Telemetry may be sampled or lost; critical domain state cannot depend on it.

---

## 5. W3C Trace Context

W3C Trace Context defines `traceparent` and `tracestate` for distributed trace propagation [S03].

Aurora can propagate trace context through:

- native RPC;
- HTTP/gRPC;
- A2A metadata/headers;
- MCP transport;
- effect gateways;
- data-channel setup.

Trace context is not authority. A malicious provider can copy/spoof correlation fields, so actor identity and grants remain independently verified.

---

## 6. CloudEvents

CloudEvents defines a common event envelope with attributes such as:

- `specversion`;
- `id`;
- `source`;
- `type`;
- `subject`;
- `time`;
- `datacontenttype`;
- `dataschema`;
- `data` [S04].

### Aurora fit

A common envelope can carry provider/domain events across bindings.

Example:

```json
{
  "specversion": "1.0",
  "id": "EVT-008822",
  "source": "harness://research/local-01",
  "type": "com.aurora.delegation.progress.v1",
  "subject": "delegation/DEL-0042",
  "time": "2026-08-05T23:00:00Z",
  "dataschema": "schema://delegation-progress/1",
  "data": {
    "phase": "source-comparison",
    "budget_remaining_percent": 62
  }
}
```

### Limits

CloudEvents standardizes envelope, not Aurora lifecycle/evidence semantics, durability or ordering.

---

## 7. AsyncAPI

AsyncAPI describes event-driven APIs, channels, messages, operations, servers and bindings, with tooling/generation support [S05].

Potential Aurora use:

- document provider event streams;
- data-channel contracts;
- generated client/server stubs;
- schema and binding references;
- event catalog.

Risks:

- documentation drift;
- over-modeling early topology;
- generated clients do not enforce authority;
- channel semantics differ across brokers/transports.

AsyncAPI is a documentation/codegen candidate after transport/channel design, not the Product Domain.

---

## 8. JSON Schema

JSON Schema 2020-12 supports validation, references, vocabularies and machine-readable contracts [S06].

Strong fit candidates:

- Provider Manifest;
- Context Pack metadata;
- Authority Grant metadata;
- artifact/evidence metadata;
- configuration;
- research source manifests;
- human-reviewable contracts.

Benefits:

- human-readable JSON/YAML ecosystem;
- broad validation/tooling;
- language-neutral;
- flexible composition.

Limits:

- semantic versioning policy not inherent;
- code generation quality varies;
- large binary/high-rate data unsuitable;
- compatibility rules need separate tooling;
- numeric/unit/domain constraints may require additional validation.

---

## 9. Protocol Buffers and gRPC

Protocol Buffers define typed schemas and multi-language code generation; gRPC uses them for RPC/streaming [S07].

Potential fit:

- local/remote high-performance RPC;
- device/controller communication where ecosystem supports;
- high-rate typed streams;
- stable generated SDKs.

Benefits:

- compact binary;
- mature code generation;
- streaming/RPC ecosystem;
- field evolution rules.

Limits:

- less human-readable;
- browser/tooling complexity;
- semantic compatibility beyond fields;
- may be unnecessary for early local calls;
- not ideal as editable Product Blueprint/manifest source.

Aurora may use Protobuf for selected bindings while canonical semantics remain format-neutral.

---

## 10. Buf breaking-change detection

Buf provides mechanical Protobuf breaking-change detection against configurable rules [S08].

General lesson:

- contract compatibility should be automated;
- compare against accepted baseline;
- fail CI on breaking schema changes;
- allow explicit version bump/migration rather than silent break.

Equivalent checks are needed for JSON Schema/Aurora semantics even if tooling differs.

---

## 11. Compatibility dimensions

### Structural

Can payload parse/validate?

### Behavioral

Do lifecycle and errors behave the same?

### Semantic

Does a field/state still mean the same?

### Authority/security

Did effects/data requirements change?

### Operational

Did cancellation/recovery/ordering change?

### Evidence

Did artifact/evidence requirements change?

A structurally additive field can be semantically breaking if it changes authority or acceptance.

---

## 12. Event naming/versioning

Recommended properties:

- stable namespaced type;
- major semantic version in type/schema;
- source/subject identities;
- explicit schema reference;
- event ID;
- occurrence time and receive time where needed;
- trace/audit correlation;
- data classification;
- artifact reference for large payload;
- no secret values.

Example type:

```text
com.aurora.delegation.completed.v1
```

Versioning policy requires a dedicated Contract Spec.

---

## 13. Ordering, duplication and delivery

Distributed messages can be:

- duplicated;
- delayed;
- reordered;
- lost;
- replayed.

Aurora should use:

- unique event ID;
- source sequence where available;
- idempotent projection updates;
- terminal-state validation;
- state/snapshot query;
- artifact/receipt reconciliation;
- no assumption of global total order;
- explicit causal relationships.

A broker's delivery guarantee does not remove application reconciliation.

---

## 14. Large artifacts and telemetry

Events should reference rather than embed:

- repository archives;
- videos;
- waveforms;
- high-rate telemetry;
- model checkpoints;
- large reports.

Reference metadata:

- artifact ID;
- URI/location abstraction;
- digest;
- media type/schema;
- size;
- classification;
- encryption/access;
- retention;
- producer/version.

High-rate telemetry can use dedicated data-plane protocol/storage while control events remain in Aurora's event model.

---

## 15. Observability redaction

Default telemetry must exclude:

- prompts with sensitive content;
- secrets/tokens;
- personal memory payloads;
- full project files;
- raw audio/video;
- private tool outputs.

Instead record:

- IDs;
- sizes;
- classifications;
- hashes/references;
- model/provider;
- timing;
- outcome/error class;
- redaction count/policy.

Debug capture requires explicit temporary authority and retention.

---

## 16. Trace versus evidence

Trace answers:

> “Which operations happened and how long did they take?”

Evidence answers:

> “Why should this criterion/hypothesis be accepted?”

A trace may support evidence but cannot automatically prove product correctness.

Example:

- trace shows test command ran;
- receipt shows exit 0;
- test report/artifacts show assertions;
- evidence links report to acceptance criterion;
- verdict accepts coverage.

---

## 17. SPK-006

Golden flow:

```text
Presence request
→ Aurora Mission/Delegation
→ A2A provider
→ provider calls MCP tool
→ tool requests Effect Gateway
→ artifact published
→ evidence recorded
```

Prove:

- one trace/correlation across boundaries;
- domain IDs available without sensitive payload;
- logs/metrics/spans exported locally;
- effect/audit receipt linked but stored separately;
- sampling loss does not lose state/evidence;
- provider without native OTel can be adapted;
- direct data channel references correlate.

---

## 18. Schema spike questions

- Can JSON Schema express Manifest/Context/Grant cleanly?
- Can types be generated for first AHDK language?
- Which compatibility changes are detectable?
- Where is Protobuf/gRPC materially better?
- How do A2A/MCP schemas map without duplicate ontology?
- How are units, money and time represented?
- How are artifact references shared?
- What data classifications are machine-enforced?
- How are schema versions bound to in-flight Delegations?

---

## 19. Decision implications

### Supported

- OpenTelemetry as current observability baseline hypothesis;
- W3C trace propagation;
- separate domain event, telemetry, audit, receipt and evidence;
- CloudEvents as candidate event envelope;
- JSON Schema as candidate for human-readable contracts;
- Protobuf/gRPC for selected typed/performance boundaries;
- mechanical compatibility checks;
- reference large artifacts;
- redact sensitive telemetry by default.

### Not decided

- exact event transport/broker;
- storage backend;
- JSON Schema versus Protobuf per boundary;
- AsyncAPI adoption timing;
- semantic convention names;
- telemetry backend;
- sampling/retention;
- canonical event format.

---

## 20. Limitations

- Standards solve syntax/transport more than Aurora semantics.
- Tooling maturity differs by language.
- Schema generation can introduce awkward APIs.
- Physical telemetry has domain timing/unit requirements beyond generic events.
- Observability itself creates privacy/cost risk.

---

## 21. Conclusion

Aurora should build a semantic contract layer and then use fit-for-purpose standards:

```text
OpenTelemetry for observability
CloudEvents as event-envelope candidate
JSON Schema for human-readable contracts/manifests
Protobuf/gRPC where type/performance justify
AsyncAPI for channel documentation after topology is known
mechanical compatibility plus semantic conformance
```

No format or telemetry backend should become a second source of product truth.
