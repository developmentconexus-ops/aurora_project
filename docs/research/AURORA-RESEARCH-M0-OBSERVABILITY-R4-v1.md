---
id: RESEARCH-AURORA-M0-OBSERVABILITY-R4-V1
title: Aurora M0 R4 Research — Observability Mechanism
document_type: research_report
form: explanation
authority: research
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - M0 R4 observability mechanism research through 2026-08-07
related:
  - DOC-AURORA-M0-R4-OPERATOR-AUTHORIZATION
  - DESIGN-AURORA-M0-R4-DECISION-LANDSCAPE
  - DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC
source_manifest: AURORA-RESEARCH-M0-OBSERVABILITY-R4-v1.sources.json
review_triggers:
  - OpenTelemetry Go signal maturity change
  - M0 telemetry requirement change
  - remote/distributed topology promotion
last_reviewed: 2026-08-07
---

# Aurora M0 R4 — Observability Mechanism

## 1. Question

Which observability mechanism gives M0 stable correlation and later ecosystem compatibility without creating a Collector/backend dependency or letting telemetry become canonical state/evidence?

## 2. Current maturity

OpenTelemetry Go currently reports:

```text
traces  = Stable
metrics = Stable
logs    = Beta
```

[S01]. OTLP 1.11.0 is stable for trace, metric and log signals and supports HTTP/gRPC transport [S02].

Go's standard `log/slog` provides structured logging with level/message/key-value attributes [S03]. OpenTelemetry documents bridges from common Go loggers including `slog` and supports either forwarding local logs or exporting them through its log SDK [S04].

## 3. M0 needs

R3 requires correlation identifiers and proof/decision-oriented telemetry, not an observability platform.

Minimum M0 signal needs:

```text
traces
→ transition / recovery / export / restore operation correlation

metrics
→ latency, failure counts, recovery outcomes, proof-related thresholds where useful

structured logs
→ local diagnostics with stable IDs and redaction
```

Audit and Evidence remain separate domain records.

## 4. Proposed mechanism

```text
OpenTelemetry API/SDK for traces + metrics
Go log/slog for application structured logs
optional slog→OpenTelemetry bridge later
optional OTLP exporter/backend
```

No Collector or remote backend is required for M0 correctness.

Default local operation can use no-op/in-memory/test exporters for tests and a local diagnostic exporter only when configured. A future backend can be attached without changing domain code.

## 5. Why not make OTel Logs mandatory now?

The OpenTelemetry Go logs implementation is still Beta [S01], while `slog` is standard-library structured logging [S03]. M0 can preserve structured diagnostic semantics in `slog` today and add an OTel bridge/exporter without changing business/domain state.

That is a cleaner evidence-bounded commitment than making beta signal implementation a foundational dependency.

## 6. Privacy rule

Telemetry attributes may include stable IDs/revisions/classifications but not unrestricted state/authority payloads.

Sensitive material remains in governed state/evidence stores. Trace/log loss or exporter failure cannot affect state transition success, authority evaluation or Golden Proof verdict.

## 7. Proposed R4 decision

```text
R4-Q-TELEM-001
→ OpenTelemetry traces and metrics APIs/SDKs
→ Go slog for structured local logs
→ OTLP exporter/backend optional and replaceable
→ OTel Go logs SDK deferred until maturity/need justifies it
```

Classification: `REVERSIBLE_MECHANISM`.

No Architecture Spike is required merely to choose this boundary. R7 tests should verify redaction, no-op/exporter-failure independence and correlation propagation.

## 8. Exit path

If OpenTelemetry later becomes unsuitable, Aurora-specific correlation attributes remain domain-independent conventions and can be mapped to another observability adapter. `slog` handlers are already replaceable.

No telemetry data format is allowed to become a migration dependency for canonical state.
