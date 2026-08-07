---
id: DOC-AURORA-CAP-SOVEREIGN-CORE-R4-DECISION-COVERAGE
title: CAP-SOVEREIGN-CORE R4 Decision Coverage
document_type: architecture_decision_coverage
form: reference
authority: reference
status: current
version: 0.3.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - current R4 disposition of all M0 implementation-blocking architecture questions
related:
  - DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC
  - DESIGN-AURORA-M0-R4-DECISION-LANDSCAPE
  - DOC-AURORA-ADR-INDEX
  - DOC-AURORA-RESEARCH-MAP
  - DOC-AURORA-M0-R4-OPERATOR-AUTHORIZATION
  - DOC-AURORA-M0-R4-ADR-ACCEPTANCE-SPK001-AUTHORIZATION
  - DOC-AURORA-M0-R4-SPK001-EVIDENCE-RECEIPT
  - REVIEW-AURORA-M0-R4-SPK001-SOVEREIGN-STORE-2026-08-07
  - RESEARCH-AURORA-MASTRA-COGNITIVE-HARNESS-R4-V1
  - ADR-AURORA-0009
source_revision: d00cc1abfc2a41ac7e81e1f3478e188b3c5e9e52
last_reviewed: 2026-08-07
---

# CAP-SOVEREIGN-CORE — R4 Decision Coverage

## 1. Purpose

R3 identified exactly 15 implementation-blocking M0 R4 questions. R4 must make every question discoverably resolved, intentionally deferred by a named trigger, or blocked by explicit missing evidence.

```text
R3 M0 R4 questions expected: 15
R4 disposition rows:          15
Missing:                        0
```

A later focused Mastra study found an important cross-horizon runtime direction. It changes interpretation of existing rows but does **not** add a sixteenth M0 implementation blocker because M0 does not consume Mastra.

After operator acceptance and SPK-001 execution:

- ADR-0003, ADR-0004, ADR-0005, ADR-0006 and ADR-0009 are `accepted`;
- SPK-001 is `PASS / REVIEWED / DECISION_INFORMED / CLOSED`;
- ADR-0007 is `proposed / evidence-ready` and awaits operator decision;
- ADR-0008 remains `proposed / spike-blocked`;
- SPK-002 remains `NOT AUTHORIZED`.

## 2. Coverage matrix

| R3 question | Class | Research / decision owner | Executable evidence | Current disposition |
|---|---|---|---|---|
| `R4-Q-CORE-001` Sovereign Core runtime | HIGH_LOCK_IN | ADR-0003 → accepted Go Sovereign Core; ADR-0009 separates Harness runtime | language-specific M0 spike not required | `DECIDED / ACCEPTED` |
| `R4-Q-STORE-001` operational store | HIGH_LOCK_IN | ADR-0007 v0.2.0 → SQLite + modernc proposed | SPK-001 PASS/CLOSED | `EVIDENCE_READY / OPERATOR_DECISION_REQUIRED` |
| `R4-Q-STATE-001` state-versus-event | NO_REGRET | ADR-0004 → accepted current state + immutable revisions + separate audit/events | SPK-001 additionally proved transactional physical fit | `DECIDED / ACCEPTED` |
| `R4-Q-SCHEMA-001` schema/serialization | REVERSIBLE_MECHANISM | ADR-0005 → accepted JSON Schema 2020-12 + JSON/JCS | implementation conformance later | `DECIDED / ACCEPTED` |
| `R4-Q-ATOMIC-001` crash-consistent commit | HIGH_LOCK_IN | ADR-0007 proposed transaction posture | SPK-001 real process-kill matrix PASS | `EVIDENCE_READY / OPERATOR_DECISION_REQUIRED` |
| `R4-Q-INTEGRITY-001` integrity | HIGH_LOCK_IN_SECURITY | ADR-0005 accepted SHA-256/content boundary; ADR-0008 owns authenticated governing anchor | SPK-001 proved SQLite structural/integrity handling; SPK-002 still required for authenticated trust anchor | `PARTIALLY_DECIDED / SPK002_REQUIRED` |
| `R4-Q-TIME-001` time/rollback | HIGH_LOCK_IN_SECURITY | ADR-0008 → proposed fail-closed time/high-water model | SPK-002 required | `SPIKE_REQUIRED / EXECUTION_NOT_AUTHORIZED` |
| `R4-Q-AUTHN-001` owner auth/bootstrap | HIGH_LOCK_IN_SECURITY | ADR-0008 → proposed random wrapped ORK + Argon2id owner unlock | SPK-002 required | `SPIKE_REQUIRED / EXECUTION_NOT_AUTHORIZED` |
| `R4-Q-EXPORT-001` export/backup | REVERSIBLE_SECURITY | ADR-0005 accepted logical export; ADR-0007 proposed operational SQLite backup | SPK-001 PASS for operational backup/restore; SPK-002 remains for owner-root recovery | `PARTIALLY_DECIDED / SPK002_DEPENDENCY_REMAINS` |
| `R4-Q-MIGRATE-001` migration | REVERSIBLE_MECHANISM | ADR-0005 accepted application-owned explicit version-pair migration | SPK-001 v1→v2 fixture PASS | `DECIDED / ACCEPTED + EVIDENCED` |
| `R4-Q-AUDIT-001` event/audit physical mechanism | REVERSIBLE_MECHANISM | ADR-0004 accepted logical distinction + M0 co-located transaction | SPK-001 atomic state/audit/evidence cases PASS | `DECIDED / ACCEPTED + EVIDENCED` |
| `R4-Q-TELEM-001` telemetry | REVERSIBLE_MECHANISM | ADR-0006 accepted OTel traces/metrics + slog; backend optional | no architecture spike required | `DECIDED / ACCEPTED` |
| `R4-Q-TOPOLOGY-001` M0 Core process topology | NO_REGRET_M0 | ADR-0004 accepted one local modular Sovereign Core | SPK-001 proved embedded local store fit on Linux/Windows | `DECIDED / ACCEPTED` |
| `R4-Q-ENGINE-001` M0 durable engine | NOT_YET_A_DECISION | ADR-0004 accepted non-selection for M0; future Mastra/durable mechanisms remain provider/port candidates | no M0 engine spike required | `DECIDED / INTENTIONAL_NON_SELECTION` |
| `R4-Q-RESTORE-001` authority freshness after restore | HIGH_LOCK_IN_SECURITY | ADR-0008 → proposed `REVALIDATION_REQUIRED` + owner-only new authority revision | SPK-002 required | `SPIKE_REQUIRED / EXECUTION_NOT_AUTHORIZED` |

## 3. Accepted M0 decisions

```text
ADR-0003 — Go as initial Sovereign Core runtime
ADR-0004 — local modular Sovereign Core + explicit current state/revisions + no full event sourcing/durable engine in M0
ADR-0005 — JSON Schema/JSON/JCS portable logical state + protected export + explicit migrations
ADR-0006 — OpenTelemetry traces/metrics + slog, optional exporter/backend
```

These were explicitly accepted by the operator on 2026-08-07.

## 4. Accepted cross-horizon direction

```text
ADR-0009 — Mastra as preferred-first first-party cognitive/Harness runtime substrate
```

Mastra remains outside M0's sovereignty boundary and adds no M0 implementation blocker.

## 5. SPK-001 result and ADR-0007

```text
SPK-AURORA-M0-SOVEREIGN-STORE-001
→ PASS
→ EVIDENCE_COMPLETE
→ REVIEWED
→ DECISION_INFORMED
→ CLOSED
```

The final 4-case Ubuntu/Windows × modernc/mattn matrix passed all correctness gates and required evidence receipts.

The reviewed recommendation is:

```text
SQLite
+ database/sql
+ modernc.org/sqlite preferred initial binding
+ WAL
+ synchronous=FULL
+ exact compatible module lock
```

ADR-0007 revision `0.2.0` remains non-governing until operator acceptance.

## 6. Remaining architecture blocker

ADR-0008 remains blocked by:

```text
SPK-AURORA-M0-OWNER-TRUST-002
```

SPK-002 is sequenced after the now-reviewed SPK-001 result but is **not authorized** merely because that dependency is satisfied.

It must still prove:

- Owner Root unlock/custody behavior;
- rollback detection/trust-anchor protocol;
- crash gap between database and trust anchor;
- backward-time handling;
- historical restore freshness;
- owner-only revalidation and self-revalidation denial.

## 7. R4 gate implication

Current state:

```text
15/15 M0 questions accounted for
ADR-0003 / 0004 / 0005 / 0006 accepted
ADR-0009 accepted cross-horizon
SPK-001 closed with PASS
ADR-0007 evidence-ready, operator decision pending
ADR-0008 still spike-blocked
SPK-002 not authorized
```

Therefore:

```text
R4 BLOCKED
```

The blockers are now narrow:

1. operator accepts/rejects/revises evidence-informed ADR-0007;
2. SPK-002 requires separate explicit execution authorization;
3. SPK-002 evidence must be reviewed and ADR-0008 decided.

No R5 work may begin by implication.
