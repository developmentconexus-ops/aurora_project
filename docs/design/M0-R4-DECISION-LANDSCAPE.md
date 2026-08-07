---
id: DESIGN-AURORA-M0-R4-DECISION-LANDSCAPE
title: M0 R4 Architecture Decision Landscape
document_type: architecture_decision_landscape
form: reference
authority: design
status: proposed
version: 0.1.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - proposed dependency and lock-in map for M0 R4 architecture decisions
related:
  - DOC-AURORA-STATUS
  - DOC-AURORA-M0-R4-OPERATOR-AUTHORIZATION
  - DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC
  - DOC-AURORA-CAP-SOVEREIGN-CORE-THREAT-MODEL
  - DOC-AURORA-CAP-SOVEREIGN-CORE-TEST-PLAN
source_revision: d00cc1abfc2a41ac7e81e1f3478e188b3c5e9e52
last_reviewed: 2026-08-07
---

# M0 R4 — Architecture Decision Landscape

## 1. Purpose

R4 must not optimize fifteen isolated technology choices. It must resolve the smallest coherent set of implementation-blocking decisions for M0 while testing those choices against:

```text
H0 — M0 Golden Proof and R2/R3 requirements
H1 — foreseeable M1–M4 evolution
H2 — constitutional Aurora direction
```

Approved philosophy:

```text
Explore long horizon.
Commit only to the next evidence-supported irreversible step.
```

This landscape does not authorize Architecture Spike execution or implementation.

## 2. Decision classes

| Class | Meaning | R4 treatment |
|---|---|---|
| `NO_REGRET` | semantic/structural choice already strongly constrained by R3 and cheap to preserve | decide strongly if compatible with accepted sources |
| `REVERSIBLE_MECHANISM` | implementation choice intentionally isolated behind a boundary | select current mechanism with explicit exit path |
| `HIGH_LOCK_IN` | choice can shape data, runtime, recovery or security for years | deep research plus executable evidence where documentary proof is insufficient |
| `NOT_YET_A_DECISION` | current M0 does not justify introducing the mechanism | record explicit non-selection and reconsideration trigger |

## 3. Dependency graph

```text
RUNTIME ──────────────┬──────────────→ STORAGE DRIVER / LIBRARY
                     │
                     ├──────────────→ SCHEMA TOOLING
                     │
                     └──────────────→ TELEMETRY / CRYPTO / MIGRATION LIBRARIES

STORAGE ──────────────┬──────────────→ ATOMICITY / CRASH CONSISTENCY
                     ├──────────────→ AUDIT PHYSICAL PLACEMENT
                     ├──────────────→ BACKUP / EXPORT SOURCE
                     └──────────────→ MIGRATION MECHANISM

STATE MODEL ──────────┬──────────────→ STORAGE SCHEMA
                     ├──────────────→ AUDIT / EVENT SEMANTICS
                     └──────────────→ RESTORE / MIGRATION VALIDATION

SCHEMA / SERIALIZATION ──────────────→ EXPORT FORMAT ─→ INTEGRITY / CONFIDENTIALITY

AUTHN / OWNER ROOT ───┬──────────────→ INTEGRITY KEY CUSTODY
                     └──────────────→ RESTORE REVALIDATION

TIME / ROLLBACK ─────────────────────→ AUTHORITY EXPIRY / RESTORE SAFETY

TOPOLOGY ─────────────┬──────────────→ STORAGE ACCESS MODEL
                     └──────────────→ DURABLE ENGINE APPLICABILITY
```

The practical implication is that runtime/storage/atomicity/migration and authn/time/restore cannot be finalized as independent popularity contests.

## 4. Fifteen R4 decisions

| R3 ID | Decision | Class | Main dependency | Initial R4 posture |
|---|---|---|---|---|
| `R4-Q-CORE-001` | Core language/runtime | `HIGH_LOCK_IN` | storage/schema/ops ecosystem | compare current Go/Rust and only add a third candidate if evidence justifies it |
| `R4-Q-STORE-001` | operational-state store | `HIGH_LOCK_IN` | atomicity/backup/migration/topology | compare embedded relational, local client/server relational and embedded KV classes; spike final candidate(s) |
| `R4-Q-STATE-001` | state-versus-event persistence | `NO_REGRET` | R3 canonical-state semantics | prefer explicit current state + immutable revisions + separate audit/events; full event sourcing requires independent justification |
| `R4-Q-SCHEMA-001` | schema/serialization | `REVERSIBLE_MECHANISM` | export/integrity/contracts | distinguish internal DB schema from portable logical export schema; do not force one format to own both |
| `R4-Q-ATOMIC-001` | crash-consistent commit | `HIGH_LOCK_IN` | storage | prove accepted-state/current-pointer/audit consistency under kill/crash injection |
| `R4-Q-INTEGRITY-001` | integrity mechanism | `HIGH_LOCK_IN` security | schema/authn/export/store | distinguish accidental-corruption detection from authenticated tamper detection |
| `R4-Q-TIME-001` | expiry/time rollback semantics | `HIGH_LOCK_IN` security | authority | fail closed on suspicious backward wall-clock movement; prove behavior |
| `R4-Q-AUTHN-001` | local owner authentication/bootstrap | `HIGH_LOCK_IN` security | key custody/revalidation | explicit owner credential/trust root outside canonical Project state; technical DB access alone is insufficient |
| `R4-Q-EXPORT-001` | export/backup format/topology | `REVERSIBLE_MECHANISM` with security impact | schema/integrity/store | portable logical package plus explicit encryption/integrity; raw DB copy alone cannot be sovereignty contract |
| `R4-Q-MIGRATE-001` | migration mechanism/tooling | `REVERSIBLE_MECHANISM` | schema/store/runtime | application-owned ordered/version-pair migrations with invariant verification; exact tooling follows stack |
| `R4-Q-AUDIT-001` | event/audit physical mechanism | `REVERSIBLE_MECHANISM` | store/atomicity | co-locate logically distinct audit/event records with canonical transactional boundary unless evidence requires a separate system |
| `R4-Q-TELEM-001` | telemetry backend/transport | `REVERSIBLE_MECHANISM` | runtime | vendor-neutral instrumentation semantics; no backend should be required for proof correctness |
| `R4-Q-TOPOLOGY-001` | process/deployment topology | `NO_REGRET` for M0 | store/runtime | one local modular Core process unless evidence proves physical distribution necessary |
| `R4-Q-ENGINE-001` | durable workflow engine applicability | `NOT_YET_A_DECISION` | topology/future mission lifecycle | do not introduce for M0 unless an M0 test cannot be satisfied safely without it |
| `R4-Q-RESTORE-001` | authority freshness after restore | `HIGH_LOCK_IN` security | authn/time/integrity | default `REVALIDATION_REQUIRED`; any automatic freshness proof must be stronger than restored state itself |

## 5. Decision order

R4 should resolve in this order because later choices depend on earlier ones:

```text
A. structural/no-regret decisions
   state model
   topology
   durable-engine applicability

B. runtime + persistence cluster
   runtime
   store
   atomicity
   audit placement
   migrations

C. portable representation cluster
   schema/serialization
   export package
   integrity/confidentiality

D. authority security cluster
   owner authentication/bootstrap
   time rollback semantics
   restore freshness/revalidation

E. observability
   telemetry API/transport/backend posture
```

## 6. Evidence policy

A documentary decision is permitted when official specifications/documentation are enough to establish the required property and migration/rollback implications are understood.

A spike is required when the claim depends on observed runtime behavior such as:

- crash consistency under kill/power-loss-like fault injection;
- driver/library behavior rather than engine specification alone;
- backup/restore behavior across process death;
- ambiguous commit reconciliation;
- clock rollback handling through actual runtime/storage integration;
- local owner credential/key-custody behavior across restart/restore;
- operational burden that materially differs between candidates.

## 7. Known R4 proof blocker from accepted research governance

`docs/research/RESEARCH-MAP.md` explicitly requires a `crash/restart/restore spike` before technical commitment for Sovereign Core storage and recovery.

Therefore R4 cannot legitimately PASS merely from documentary comparison if the selected persistence architecture has not received the required executable evidence.

Architecture Spike **specification** is within R4. Architecture Spike **execution** remains separately prohibited until explicit operator authorization.

## 8. Stop condition

R4 may close only when:

```text
all implementation-blocking decisions
→ DECIDED or explicitly NOT_YET_A_DECISION
→ required ADRs accepted at the appropriate authority level
→ required spike evidence completed/reviewed
→ migration/rollback implications recorded
→ no current material architecture choice unresolved
```

Until then the valid R4 verdict is `BLOCKED`, not an optimistic PASS.
