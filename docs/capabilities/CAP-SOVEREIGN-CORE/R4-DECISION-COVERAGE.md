---
id: DOC-AURORA-CAP-SOVEREIGN-CORE-R4-DECISION-COVERAGE
title: CAP-SOVEREIGN-CORE R4 Decision Coverage
document_type: architecture_decision_coverage
form: reference
authority: reference
status: current
version: 0.2.0
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

`PROPOSED_ADR_READY` does not mean accepted. `SPIKE_REQUIRED` cannot become decided merely from documentary research.

A later focused Mastra study found an important cross-horizon runtime direction. It changes interpretation of existing rows but does **not** add a sixteenth M0 implementation blocker because M0 does not consume Mastra.

## 2. Coverage matrix

| R3 question | Class | Research | Proposed decision owner | Executable evidence | Current disposition |
|---|---|---|---|---|---|
| `R4-Q-CORE-001` Sovereign Core runtime | HIGH_LOCK_IN | RUNTIME-PERSISTENCE-R4 + MASTRA-COGNITIVE-HARNESS-R4 | ADR-0003 → Go **Sovereign Core**; ADR-0009 separates future Harness runtime | language-specific M0 spike not required | `PROPOSED_ADR_READY / OPERATOR_REVIEW_REQUIRED` |
| `R4-Q-STORE-001` operational store | HIGH_LOCK_IN | RUNTIME-PERSISTENCE-R4 | ADR-0007 → SQLite candidate | SPK-AURORA-M0-SOVEREIGN-STORE-001 | `SPIKE_REQUIRED / EXECUTION_NOT_AUTHORIZED` |
| `R4-Q-STATE-001` state-versus-event | NO_REGRET | RUNTIME-PERSISTENCE-R4 | ADR-0004 → current state + immutable revisions + separate audit/events; no full event sourcing | covered by store spike atomic boundary only; no pattern-comparison spike required | `PROPOSED_ADR_READY / OPERATOR_REVIEW_REQUIRED` |
| `R4-Q-SCHEMA-001` schema/serialization | REVERSIBLE_MECHANISM | PORTABILITY-INTEGRITY-R4 | ADR-0005 → JSON Schema 2020-12 + JSON/JCS | implementation conformance later | `PROPOSED_ADR_READY / OPERATOR_REVIEW_REQUIRED` |
| `R4-Q-ATOMIC-001` crash-consistent commit | HIGH_LOCK_IN | RUNTIME-PERSISTENCE-R4 | ADR-0007 conditional transaction boundary | SPK-AURORA-M0-SOVEREIGN-STORE-001 | `SPIKE_REQUIRED / EXECUTION_NOT_AUTHORIZED` |
| `R4-Q-INTEGRITY-001` integrity | HIGH_LOCK_IN_SECURITY | PORTABILITY-INTEGRITY-R4 + OWNER-AUTHORITY-RECOVERY-R4 | ADR-0005 → SHA-256 content; ADR-0008 → HMAC authenticated governing anchor | SPK-AURORA-M0-OWNER-TRUST-002 | `PARTIAL_DOCUMENTARY_DECISION / SPIKE_REQUIRED` |
| `R4-Q-TIME-001` time/rollback | HIGH_LOCK_IN_SECURITY | OWNER-AUTHORITY-RECOVERY-R4 | ADR-0008 → monotonic in-process + authenticated wall-time high-water, fail closed | SPK-AURORA-M0-OWNER-TRUST-002 | `SPIKE_REQUIRED / EXECUTION_NOT_AUTHORIZED` |
| `R4-Q-AUTHN-001` owner auth/bootstrap | HIGH_LOCK_IN_SECURITY | OWNER-AUTHORITY-RECOVERY-R4 | ADR-0008 → random wrapped ORK + Argon2id owner unlock | SPK-AURORA-M0-OWNER-TRUST-002 | `SPIKE_REQUIRED / EXECUTION_NOT_AUTHORIZED` |
| `R4-Q-EXPORT-001` export/backup | REVERSIBLE_SECURITY | PORTABILITY-INTEGRITY-R4 + RUNTIME-PERSISTENCE-R4 | ADR-0005 → logical export + age; operational backup separate | SPK-001 proves operational backup; SPK-002 proves owner-root recovery behavior | `PROPOSED_DIRECTION / SPIKE_DEPENDENCIES_REMAIN` |
| `R4-Q-MIGRATE-001` migration | REVERSIBLE_MECHANISM | PORTABILITY-INTEGRITY-R4 + RUNTIME-PERSISTENCE-R4 | ADR-0005 → application-owned explicit version-pair migration | SPK-001 migration fixture | `PROPOSED_DIRECTION / SPIKE_DEPENDENCY_REMAINS` |
| `R4-Q-AUDIT-001` event/audit physical mechanism | REVERSIBLE_MECHANISM | RUNTIME-PERSISTENCE-R4 | ADR-0004 → logically distinct, physically co-located transactionally in M0 | SPK-001 atomic audit/state cases | `PROPOSED_DIRECTION / SPIKE_DEPENDENCY_REMAINS` |
| `R4-Q-TELEM-001` telemetry | REVERSIBLE_MECHANISM | OBSERVABILITY-R4 | ADR-0006 → OTel traces/metrics + slog; backend optional | no architecture spike required | `PROPOSED_ADR_READY / OPERATOR_REVIEW_REQUIRED` |
| `R4-Q-TOPOLOGY-001` M0 Core process topology | NO_REGRET_M0 | RUNTIME-PERSISTENCE-R4 + MASTRA-COGNITIVE-HARNESS-R4 | ADR-0004 → one local modular Sovereign Core; future Mastra process is a provider boundary, not M0 Core topology | no topology spike required for M0 | `PROPOSED_ADR_READY / OPERATOR_REVIEW_REQUIRED` |
| `R4-Q-ENGINE-001` M0 durable engine | NOT_YET_A_DECISION | RUNTIME-PERSISTENCE-R4 + A0 DURABLE-EXECUTION + MASTRA-COGNITIVE-HARNESS-R4 | ADR-0004 → do not introduce in M0; Mastra durable/workflow mechanisms remain future provider/port candidates | no M0 engine spike | `INTENTIONAL_NON_SELECTION / OPERATOR_REVIEW_REQUIRED` |
| `R4-Q-RESTORE-001` authority freshness after restore | HIGH_LOCK_IN_SECURITY | OWNER-AUTHORITY-RECOVERY-R4 | ADR-0008 → default REVALIDATION_REQUIRED + owner-only new authority revision | SPK-AURORA-M0-OWNER-TRUST-002 | `SPIKE_REQUIRED / EXECUTION_NOT_AUTHORIZED` |

## 3. Mastra cross-horizon finding

The current Mastra surface is broad enough to replace substantial generic future Harness infrastructure:

```text
AgentController
Observational Memory / extraction
workflows / HITL / durable agents
RAG
Workspaces / Skills
signals / schedules / background tasks
ACP / A2A / SDK subagents
tracing / evals
```

The resulting proposed decision is ADR-0009:

```text
Mastra
→ preferred substrate to evaluate first for first-party agentic Harnesses

Aurora Core
→ remains owner of identity / state / authority / governance / global verdict
```

This is **not required to close M0 R4**. Its first execution/conformance proof is deferred to the first Mastra-backed Capability because M0 explicitly must survive without a Harness as authority.

## 4. Proposed M0 decisions ready for operator review without further architecture spike

```text
ADR-0003 — Go as initial Sovereign Core runtime
ADR-0004 — local modular Sovereign Core + explicit current state/revisions + no event sourcing/durable engine in M0
ADR-0005 — JSON Schema/JSON/JCS portable logical state + age outer export + explicit migrations
ADR-0006 — OpenTelemetry traces/metrics + slog, optional exporter/backend
```

ADR-0003 and ADR-0004 were refined after the Mastra assessment but remain the same M0-level decisions.

These are non-governing until operator acceptance.

## 5. Cross-horizon proposal ready for separate operator review

```text
ADR-0009 — Mastra as preferred first-party cognitive/Harness runtime substrate
```

ADR-0009 may be reviewed independently, but accepting or rejecting it does not remove the M0 storage/recovery blockers and does not authorize implementation.

## 6. Decisions that cannot be accepted yet

```text
ADR-0007 — SQLite operational store
  BLOCKED BY SPK-AURORA-M0-SOVEREIGN-STORE-001

ADR-0008 — Owner Root / trust anchor / time / restore freshness
  BLOCKED BY SPK-AURORA-M0-OWNER-TRUST-002
```

SPK-002 is sequenced after a viable SPK-001 store/binding result.

## 7. R4 gate implication

At this documentary checkpoint:

```text
all 15 M0 questions are accounted for
Mastra cross-horizon finding is accounted for without widening M0
BUT
2 material M0 Architecture Spikes require execution/review
AND
M0 material ADRs remain unaccepted
```

Therefore R4 cannot return PASS yet.

Unless additional authority/evidence changes the state, the correct readiness verdict is:

```text
R4 BLOCKED
```

Blockers remain narrow:

1. Architecture Spike execution is still prohibited/not authorized;
2. SPK-001 evidence is required for store/atomicity/backup/migration physical behavior;
3. SPK-002 evidence is required for owner-root/trust-anchor/time/restore behavior;
4. operator review/acceptance is required for material M0 ADRs after evidence is sufficient.

Mastra is not a fifth M0 blocker.

No R5 work may begin by implication.
