---
id: DOC-AURORA-CAP-SOVEREIGN-CORE-R4-DECISION-COVERAGE
title: CAP-SOVEREIGN-CORE R4 Decision Coverage
document_type: architecture_decision_coverage
form: reference
authority: reference
status: current
version: 0.4.0
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
  - DOC-AURORA-M0-R4-ADR0007-ACCEPTANCE-SPK002-AUTHORIZATION
  - DOC-AURORA-M0-R4-SPK002-EVIDENCE-RECEIPT
  - REVIEW-AURORA-M0-R4-SPK001-SOVEREIGN-STORE-2026-08-07
  - REVIEW-AURORA-M0-R4-SPK002-OWNER-TRUST-2026-08-07
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

A later focused Mastra study added a cross-horizon runtime direction but did not create a sixteenth M0 blocker.

Current state after both required M0 architecture spikes:

```text
ADR-0003 / 0004 / 0005 / 0006 / 0007: ACCEPTED
ADR-0009: ACCEPTED cross-horizon
SPK-001: PASS / REVIEWED / CLOSED
SPK-002: PASS / REVIEWED / CLOSED
ADR-0008 v0.2.0: PROPOSED / EVIDENCE-READY
R4: BLOCKED only on operator decision for ADR-0008
```

## 2. Coverage matrix

| R3 question | Class | Decision owner | Executable evidence | Current disposition |
|---|---|---|---|---|
| `R4-Q-CORE-001` Sovereign Core runtime | HIGH_LOCK_IN | ADR-0003 accepted Go Sovereign Core; ADR-0009 separates Harness runtime | language-specific M0 spike not required | `DECIDED / ACCEPTED` |
| `R4-Q-STORE-001` operational store | HIGH_LOCK_IN | ADR-0007 accepted SQLite + `modernc.org/sqlite` | SPK-001 PASS/CLOSED | `DECIDED / ACCEPTED + EVIDENCED` |
| `R4-Q-STATE-001` state-versus-event | NO_REGRET | ADR-0004 accepted current state + immutable revisions + separate audit/events | SPK-001 proved transactional physical fit | `DECIDED / ACCEPTED + EVIDENCED` |
| `R4-Q-SCHEMA-001` schema/serialization | REVERSIBLE_MECHANISM | ADR-0005 accepted JSON Schema 2020-12 + JSON/JCS | implementation conformance later | `DECIDED / ACCEPTED` |
| `R4-Q-ATOMIC-001` crash-consistent commit | HIGH_LOCK_IN | ADR-0007 accepted SQLite transaction/WAL posture | SPK-001 process-kill matrix PASS | `DECIDED / ACCEPTED + EVIDENCED` |
| `R4-Q-INTEGRITY-001` integrity | HIGH_LOCK_IN_SECURITY | ADR-0005 accepted content digest boundary; ADR-0008 v0.2.0 proposes ORK-derived HMAC governing integrity + external trust anchor | SPK-002 PASS/CLOSED | `EVIDENCE_READY / ADR0008_OPERATOR_DECISION_REQUIRED` |
| `R4-Q-TIME-001` time/rollback | HIGH_LOCK_IN_SECURITY | ADR-0008 v0.2.0 proposes monotonic local duration + authenticated observed wall-time high-water / fail closed | SPK-002 S07 + final matrix PASS | `EVIDENCE_READY / ADR0008_OPERATOR_DECISION_REQUIRED` |
| `R4-Q-AUTHN-001` owner auth/bootstrap | HIGH_LOCK_IN_SECURITY | ADR-0008 v0.2.0 proposes random ORK + Argon2id KEK + AES-256-GCM wrapped root | SPK-002 S01/S02/S11/S12 PASS | `EVIDENCE_READY / ADR0008_OPERATOR_DECISION_REQUIRED` |
| `R4-Q-EXPORT-001` export/backup | REVERSIBLE_SECURITY | ADR-0005 accepted logical export; ADR-0007 accepted operational SQLite backup; ADR-0008 v0.2.0 defines encrypted root recovery without historical current high-water | SPK-001 + SPK-002 PASS/CLOSED | `EVIDENCE_READY / ADR0008_OPERATOR_DECISION_REQUIRED` |
| `R4-Q-MIGRATE-001` migration | REVERSIBLE_MECHANISM | ADR-0005 accepted explicit application-owned migration | SPK-001 v1→v2 fixture PASS | `DECIDED / ACCEPTED + EVIDENCED` |
| `R4-Q-AUDIT-001` event/audit physical mechanism | REVERSIBLE_MECHANISM | ADR-0004 accepted logical distinction + M0 co-located transaction | SPK-001 atomic state/audit/evidence cases PASS | `DECIDED / ACCEPTED + EVIDENCED` |
| `R4-Q-TELEM-001` telemetry | REVERSIBLE_MECHANISM | ADR-0006 accepted OTel traces/metrics + slog; backend optional | no architecture spike required | `DECIDED / ACCEPTED` |
| `R4-Q-TOPOLOGY-001` M0 Core process topology | NO_REGRET_M0 | ADR-0004 accepted one local modular Sovereign Core | SPK-001/002 prove local embedded state/trust fit on Linux/Windows | `DECIDED / ACCEPTED + EVIDENCED` |
| `R4-Q-ENGINE-001` M0 durable engine | NOT_YET_A_DECISION | ADR-0004 accepted non-selection for M0 | no M0 engine spike required | `DECIDED / INTENTIONAL_NON_SELECTION` |
| `R4-Q-RESTORE-001` authority freshness after restore | HIGH_LOCK_IN_SECURITY | ADR-0008 v0.2.0 proposes `REVALIDATION_REQUIRED` + authenticated-owner-only new authority revision | SPK-002 S08–S12 PASS | `EVIDENCE_READY / ADR0008_OPERATOR_DECISION_REQUIRED` |

## 3. Accepted M0 decisions

```text
ADR-0003 — Go as initial Sovereign Core runtime
ADR-0004 — local modular Core + explicit current state/revisions + no full event sourcing/durable engine in M0
ADR-0005 — JSON Schema/JSON/JCS portable logical state + protected export + explicit migrations
ADR-0006 — OpenTelemetry traces/metrics + slog, optional exporter/backend
ADR-0007 — SQLite + database/sql + modernc.org/sqlite operational-state baseline
```

These are governing R4 decisions but do not authorize implementation.

## 4. Accepted cross-horizon direction

```text
ADR-0009 — Mastra preferred-first first-party cognitive/Harness runtime substrate
```

Mastra remains outside M0's sovereignty boundary and adds no M0 implementation blocker.

## 5. Executable evidence closure

### SPK-001

```text
SPK-AURORA-M0-SOVEREIGN-STORE-001
→ PASS
→ EVIDENCE_COMPLETE
→ REVIEWED
→ DECISION_INFORMED
→ CLOSED
```

This evidence informed accepted ADR-0007.

### SPK-002

```text
SPK-AURORA-M0-OWNER-TRUST-002
→ PASS
→ EVIDENCE_COMPLETE
→ REVIEWED
→ DECISION_INFORMED
→ CLOSED
```

Final hardened execution:

```text
revision: c76b96fee36878f15c54028b4ba1896f84ebdeca
run: 31219882882
Ubuntu: S01-S12 + classifications + hygiene PASS
Windows: S01-S12 + classifications + hygiene PASS
aggregate: PASS
```

The reviewed evidence recommends ADR-0008 v0.2.0:

```text
random 256-bit ORK
+ Argon2id KEK (64 MiB / t=3 / p=4)
+ AES-256-GCM wrapped ORK
+ HKDF-SHA-256 purpose subkeys
+ HMAC-SHA-256 governing descriptor and external trust anchor
+ authenticated generation + observed wall-time high-water
+ explicit ANCHOR_LAG reconciliation
+ STATE_ROLLBACK/TIME_UNTRUSTED fail closed
+ historical restore → REVALIDATION_REQUIRED
+ authenticated-owner-only new post-restore authority revision
```

## 6. Remaining R4 decision

There is no remaining M0 architecture spike to execute.

The only remaining material R4 decision is:

```text
ADR-0008 v0.2.0
→ operator ACCEPT | REJECT | REVISE
```

ADR-0008 remains non-governing until that explicit operator decision.

## 7. R4 gate implication

Current state:

```text
15/15 M0 architecture questions accounted for
all required M0 architecture spikes PASS / REVIEWED / CLOSED
ADR-0003 / 0004 / 0005 / 0006 / 0007 accepted
ADR-0009 accepted cross-horizon
ADR-0008 evidence-ready / proposed
```

Therefore the correct current verdict is still:

```text
R4 BLOCKED
```

but the blocker is now exactly one operator decision:

```text
ADR-0008 v0.2.0
```

If ADR-0008 is accepted and no new material finding appears, R4 may then be reviewed for PASS. R5 still requires separate explicit operator authorization; it does not begin automatically from ADR acceptance or R4 PASS.
