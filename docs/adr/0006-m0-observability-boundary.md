---
id: ADR-AURORA-0006
title: M0 Observability Boundary
document_type: adr
form: explanation
authority: decision
status: proposed
version: 0.1.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - proposed M0 traces, metrics, structured logging and exporter boundary
related:
  - DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC
  - RESEARCH-AURORA-M0-OBSERVABILITY-R4-V1
  - DESIGN-AURORA-M0-R4-DECISION-LANDSCAPE
supersedes: []
superseded_by: null
last_reviewed: 2026-08-07
---

# ADR-0006 — M0 Observability Boundary

## Context

R3 requires stable correlation IDs and decision/proof-oriented telemetry while explicitly prohibiting telemetry/logs from becoming canonical state or evidence authority.

## Decision drivers

- vendor neutrality;
- stable trace/metric APIs;
- structured local diagnostics;
- no Collector/backend dependency for correctness;
- sensitive-payload minimization;
- future distributed correlation without redesign.

Affected requirements include `REQ-076`, `077`, `079`, `080`, `084..086`, `103`.

## Options

- OpenTelemetry for every signal/backend immediately;
- bespoke metrics/tracing;
- OpenTelemetry traces/metrics + standard structured logs;
- no instrumentation until later.

## Decision

**Proposed:**

```text
traces + metrics = OpenTelemetry Go API/SDK
structured logs  = Go log/slog
OTLP exporter/backend = optional adapter
OpenTelemetry Go Logs SDK = deferred while Beta unless a concrete M0 need justifies it
```

M0 correctness, state recovery, audit and evidence MUST remain valid with no Collector/backend configured and when telemetry export fails.

Logs/traces carry stable IDs and redacted metadata, not unrestricted Project/authority payloads.

## Consequences

### Positive

- mature current traces/metrics path;
- easy future OTLP/backend integration;
- stable standard-library logs today;
- no beta logs SDK in the foundational critical path.

### Negative

- logs and traces are initially separate pipelines;
- optional bridges/exporters need later wiring/testing.

### Risks

Instrumentation may leak sensitive payloads or accidentally become relied upon for product state. R3 privacy/state ownership rules and redaction tests are mandatory.

## Compatibility / migration / rollback

OTel exporters and `slog` handlers are replaceable. Aurora correlation field semantics should be documented independently from any backend vendor.

## Validation

Documentary evidence is sufficient. R7 tests must prove exporter absence/failure cannot affect governing operations and sensitive payloads are absent.

## Reconsideration triggers

- OTel Go log SDK reaches Stable and a unified pipeline gives concrete value;
- future distributed topology requires stronger propagation conventions;
- OTel ceases to meet required signal semantics.

## References

- `RESEARCH-AURORA-M0-OBSERVABILITY-R4-V1`
- `CAP-SOVEREIGN-CORE/SPEC.md`
