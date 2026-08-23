---
id: DESIGN-AURORA-M0-R4-DECISION-LANDSCAPE
title: M0 R4 Architecture Decision Landscape
document_type: architecture_decision_landscape
form: reference
authority: design
status: proposed
version: 0.2.0
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
  - RESEARCH-AURORA-MASTRA-COGNITIVE-HARNESS-R4-V1
  - DESIGN-AURORA-M0-R4-MASTRA-FIT-MATRIX
  - ADR-AURORA-0009
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

Implementation-near materiality rule:

```text
If an uncertainty does not change the next architecture/build decision,
record its trigger and keep moving.
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
SOVEREIGN CORE RUNTIME ─┬────────────→ STORAGE DRIVER / LIBRARY
                        ├────────────→ SCHEMA TOOLING
                        └────────────→ TELEMETRY / CRYPTO / MIGRATION LIBRARIES

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

CORE TOPOLOGY ────────┬──────────────→ STORAGE ACCESS MODEL
                     └──────────────→ CORE DURABLE ENGINE APPLICABILITY

FUTURE AGENTIC HARNESS RUNTIME
  └─ Mastra candidate/default substrate
       ├─ memory / AgentController / workflow / RAG / workspaces
       ├─ ACP / A2A / MCP provider bindings
       └─ provider-local durable state
          NEVER → Aurora canonical identity/state/authority ownership
```

The practical implication is that runtime/storage/atomicity/migration and authn/time/restore cannot be finalized as independent popularity contests. It also prevents a second mistake: treating the Sovereign Core runtime decision as the universal runtime decision for future Harnesses.

## 4. Fifteen M0 R4 decisions

| R3 ID | Decision | Class | Main dependency | Current R4 posture |
|---|---|---|---|---|
| `R4-Q-CORE-001` | Sovereign Core language/runtime | `HIGH_LOCK_IN` | storage/schema/ops ecosystem | ADR-0003 narrows Go to Sovereign Core; future agentic Harness runtime is separate |
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
| `R4-Q-TOPOLOGY-001` | M0 Core process/deployment topology | `NO_REGRET` for M0 | store/runtime | one local modular Sovereign Core process unless evidence proves physical distribution necessary |
| `R4-Q-ENGINE-001` | M0 durable workflow engine applicability | `NOT_YET_A_DECISION` | topology/future mission lifecycle | do not introduce for M0; Mastra durable agents/workflows remain future provider/port candidates |
| `R4-Q-RESTORE-001` | authority freshness after restore | `HIGH_LOCK_IN` security | authn/time/integrity | default `REVALIDATION_REQUIRED`; any automatic freshness proof must be stronger than restored state itself |

## 5. Emergent cross-horizon finding — Mastra

After the first M0 R4 documentary checkpoint, current Mastra research produced a material long-horizon finding:

```text
Mastra is no longer well modeled as merely an agent library.
It now covers a large portion of generic Harness/cognitive-runtime infrastructure.
```

Focused evidence is recorded in:

- `RESEARCH-AURORA-MASTRA-COGNITIVE-HARNESS-R4-V1`;
- `DESIGN-AURORA-M0-R4-MASTRA-FIT-MATRIX`;
- proposed `ADR-0009`.

### Material effect on M0 decisions

It changes the **interpretation** of `R4-Q-CORE-001`:

```text
Go
→ proposed Sovereign Core runtime only

Mastra / TypeScript
→ proposed preferred future agentic Harness substrate
```

It reinforces `R4-Q-ENGINE-001`:

```text
Mastra workflows / Durable Agents / Temporal integration
→ valuable future provider-local / DurableExecutionPort mechanisms
→ still not required in M0
```

### Why it is not a sixteenth M0 blocker

M0's Golden Proof explicitly requires the Sovereign Core to survive without an external model or Harness as authority. M0 does not consume Mastra.

Therefore researching a Go↔Mastra integration exhaustively now would not change the next M0 build decision and would violate the materiality rule.

The first Mastra-backed Capability must prove the provider boundary when that Capability enters its own implementation horizon. Until then, Mastra is a cross-horizon architecture direction, not a reason to delay M0 storage/recovery evidence.

## 6. Decision order

M0 R4 should still resolve in this order:

```text
A. structural/no-regret decisions
   state model
   M0 topology
   M0 durable-engine non-selection

B. Sovereign Core runtime + persistence cluster
   Go Core runtime
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

Cross-horizon runtime direction is recorded in parallel:

```text
Mastra preferred Harness substrate proposal
→ does not enter M0 execution path
→ revisit at first consuming Capability
```

## 7. Evidence policy

A documentary decision is permitted when official specifications/documentation are enough to establish the required property and migration/rollback implications are understood.

A spike is required when the claim depends on observed runtime behavior such as:

- crash consistency under kill/power-loss-like fault injection;
- driver/library behavior rather than engine specification alone;
- backup/restore behavior across process death;
- ambiguous commit reconciliation;
- clock rollback handling through actual runtime/storage integration;
- local owner credential/key-custody behavior across restart/restore;
- operational burden that materially differs between candidates.

A future Mastra boundary proof becomes required only when Mastra is in the actual implementation path of a Capability.

## 8. Known M0 R4 proof blocker from accepted research governance

`docs/research/RESEARCH-MAP.md` explicitly requires a `crash/restart/restore spike` before technical commitment for Sovereign Core storage and recovery.

Therefore R4 cannot legitimately PASS merely from documentary comparison if the selected persistence architecture has not received the required executable evidence.

Architecture Spike **specification** is within R4. Architecture Spike **execution** remains separately prohibited until explicit operator authorization.

The Mastra finding does not alter or broaden this blocker.

## 9. Stop condition

M0 R4 may close only when:

```text
all M0 implementation-blocking decisions
→ DECIDED or explicitly NOT_YET_A_DECISION
→ required ADRs accepted at the appropriate authority level
→ required M0 spike evidence completed/reviewed
→ migration/rollback implications recorded
→ no current material M0 architecture choice unresolved
```

Future cross-horizon choices such as the exact Mastra provider binding/version MUST NOT be pulled into M0 merely to make R4 appear globally complete.

Until the existing M0 blockers are resolved, the valid verdict remains `R4 BLOCKED`.
