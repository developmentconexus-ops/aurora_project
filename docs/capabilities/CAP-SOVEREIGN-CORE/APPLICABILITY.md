---
id: DOC-AURORA-CAP-SOVEREIGN-CORE-APPLICABILITY
title: CAP-SOVEREIGN-CORE Constitutional Applicability
document_type: capability_applicability
form: reference
authority: reference
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - current M0 R1 applicability classification for CAP-SOVEREIGN-CORE
related:
  - DOC-AURORA-STATUS
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DOC-AURORA-REQUIREMENTS-TRACEABILITY
  - DOC-AURORA-BLUEPRINT-14
  - REVIEW-AURORA-M0-R0-CONSTITUTIONAL-BASELINE-RERUN-2026-08-07
  - DOC-AURORA-M0-R1-OPERATOR-AUTHORIZATION
last_reviewed: 2026-08-07
---

# CAP-SOVEREIGN-CORE — R1 Constitutional Applicability

## 1. Scope and fixed baseline

Capability under analysis:

```text
CAP-SOVEREIGN-CORE
```

Product Milestone:

```text
M0 — Sovereign Core Walking Skeleton
```

Fixed canonical R1 source revision:

```text
735f269025e2cc317424e4931f3a5cd414cd6f2a
```

R0 prerequisite:

```text
M0 ACRM R0 — PASS
```

Operator authorization for this gate is recorded in `docs/acceptance/2026-08-07-m0-r1-operator-authorization.md`.

R1 answers only which accepted constitutional requirements govern this Capability. It does **not** derive atomic Capability requirements (R2), design the Capability (R3), choose architecture/stack/spikes (R4), create a Mission Contract (R5), create Microdesign (R6), or implement anything.

## 2. Applicability semantics

The ACRM classes are used exactly as follows:

- `APPLIES` — the source requirement governs the M0 Capability and must be carried into R2.
- `PARTIALLY_APPLIES` — a bounded M0 subset governs; R2 must derive the applicable slice and explicitly preserve the deferred remainder.
- `DEFERRED_BY_ROADMAP` — constitutionally valid direction, but the accepted roadmap places the behavior outside M0.
- `NOT_APPLICABLE` — the statement is not behavior/obligation of this Capability (for example, a historical A0-only gate or repository-global publication rule); rationale is mandatory.
- `CONFLICT_REQUIRES_DECISION` — applicability cannot be resolved without a higher-owner decision.

A negative invariant can `APPLY` even when M0 does not build the corresponding future subsystem. Example: “Harness state must not become Aurora global state” constrains M0 precisely because M0 must prove sovereignty without a Harness.

## 3. M0 applicability lens

The accepted M0 outcome is the governing lens:

> A minimal Aurora Core persists identity, project state, authority snapshot and one interaction lifecycle across restart without depending on an external model or Harness as authority.

Therefore R1 treats the following as M0-owned concerns:

- stable sovereign Aurora and Project identities;
- minimal Project/current-state persistence;
- current authority snapshot as a governed projection;
- valid/invalid state transitions;
- process-independent restart/recovery;
- export, backup/restore and migration behavior required by the slice;
- minimum event/audit and proof telemetry;
- evidence and integrity required for the Golden Proof;
- simple operator interface only as a proof surface;
- negative boundaries preventing transcript/model/Harness/framework/UI from becoming authority/state owner.

Explicitly outside M0 as behavior:

- governed conversational/project memory and Context Builder;
- Capability Registry, Provider lifecycle, AHDK and conformance;
- cross-Harness Delegation orchestration;
- model routing;
- voice, sensors, handoff and multi-Presence operation;
- cloud deployment;
- physical devices/laboratory;
- adaptive campaigns and broad effect execution;
- self-improvement.

## 4. Classification summary

| Source | Total | APPLIES | PARTIAL | DEFERRED | NOT APPLICABLE | CONFLICT |
|---|---:|---:|---:|---:|---:|---:|
| Blueprint 01 — Product Vision | 12 | 7 | 3 | 1 | 1 | 0 |
| Blueprint 02 — Human–Aurora Relationship | 16 | 1 | 3 | 12 | 0 | 0 |
| Blueprint 03 — Domain and World Model | 14 | 8 | 2 | 4 | 0 | 0 |
| Blueprint 04 — Cognitive Lifecycle and Journeys | 18 | 5 | 5 | 8 | 0 | 0 |
| Blueprint 05 — Capabilities, Registry and AHDK | 20 | 1 | 1 | 18 | 0 | 0 |
| Blueprint 06 — Memory, Knowledge and Context | 23 | 4 | 3 | 16 | 0 | 0 |
| Blueprint 07 — Harness Orchestration | 22 | 5 | 4 | 13 | 0 | 0 |
| Blueprint 08 — Interaction, Multimodality and Presence | 15 | 1 | 1 | 13 | 0 | 0 |
| Blueprint 09 — Tools, Devices and Laboratory | 20 | 0 | 0 | 20 | 0 | 0 |
| Blueprint 10 — Autonomy, Authority and Safety | 20 | 2 | 4 | 14 | 0 | 0 |
| Blueprint 11 — Security, Privacy and Sovereignty | 25 | 5 | 9 | 11 | 0 | 0 |
| Blueprint 12 — System Architecture | 23 | 10 | 8 | 5 | 0 | 0 |
| Blueprint 13 — Reliability, Evaluation and Self-Improvement | 24 | 8 | 2 | 14 | 0 | 0 |
| Blueprint 14 — Capability Roadmap | 20 | 7 | 0 | 12 | 1 | 0 |
| Blueprint 15 — Documentation and Research Governance | 22 | 14 | 4 | 0 | 4 | 0 |
| **TOTAL** | **294** | **78** | **49** | **161** | **6** | **0** |

Active source set for R2, if R2 is later authorized:

```text
APPLIES             78
PARTIALLY_APPLIES   49
-----------------------
ACTIVE SOURCES      127
```

R1 does not assume one R2 requirement per source row. R2 may derive, split or consolidate atomic Capability requirements while preserving source traceability and without weakening any active source.

## 5. Full 294-requirement matrix

### Blueprint 01 — Product Vision

Directly constrains identity, single-user scope, framework independence, evidence and the narrow executable horizon. Broader North Star/experience requirements apply only to the M0 continuity slice.

| Requirement | Classification | Rationale | Downstream owner |
|---|---|---|---|
| `AUR-REQ-VIS-001` | `APPLIES` | Directly required by the M0 outcome/Golden Proof: sovereign identity and restart continuity. | R2/R3 CAP-SOVEREIGN-CORE |
| `AUR-REQ-VIS-002` | `PARTIALLY_APPLIES` | The broader constitutional statement applies only to the M0-owned identity/project/authority/state subset. | R2/R3 CAP-SOVEREIGN-CORE |
| `AUR-REQ-VIS-003` | `APPLIES` | Constrains M0 to the current Leandro-first/single-user horizon. | R2/R3 CAP-SOVEREIGN-CORE |
| `AUR-REQ-VIS-004` | `DEFERRED_BY_ROADMAP` | Future domain expansion is constitutional direction but outside the M0 walking skeleton. | future domain capabilities |
| `AUR-REQ-VIS-005` | `PARTIALLY_APPLIES` | M0 proves the continuity slice of the North Star, not the complete laboratory/capability journey. | R2/R3 CAP-SOVEREIGN-CORE |
| `AUR-REQ-VIS-006` | `PARTIALLY_APPLIES` | M0 owns one minimum interaction/state lifecycle; the complete cognitive/interaction lifecycle remains later. | R2/R3 CAP-SOVEREIGN-CORE |
| `AUR-REQ-VIS-007` | `APPLIES` | Negative invariant applies now to prevent a replaceable component from becoming Aurora/Core authority. | R2/R3 CAP-SOVEREIGN-CORE |
| `AUR-REQ-VIS-008` | `APPLIES` | Constrains M0 readiness to the current executable horizon and prevents distant commitments. | ACRM/STATUS |
| `AUR-REQ-VIS-009` | `APPLIES` | Applies as a no-premature-generalization/roadmap sequencing constraint. | R2/R3 + R4 guard |
| `AUR-REQ-VIS-010` | `APPLIES` | M0 requires criterion-linked evidence and cannot close by activity or claim alone. | R3/R8 evidence |
| `AUR-REQ-VIS-011` | `APPLIES` | Negative invariant applies now to prevent a replaceable component from becoming Aurora/Core authority. | R2/R3 CAP-SOVEREIGN-CORE |
| `AUR-REQ-VIS-012` | `NOT_APPLICABLE` | Historical A0-only gate requirement; already satisfied and not behavior of CAP-SOVEREIGN-CORE. | A0 closeout evidence |

### Blueprint 02 — Human–Aurora Relationship

Only operator authority, stable identity and state-correction boundaries touch M0. Conversational disagreement, personality, proactivity, relationship memory and interaction modes remain later.

| Requirement | Classification | Rationale | Downstream owner |
|---|---|---|---|
| `AUR-REQ-REL-001` | `DEFERRED_BY_ROADMAP` | Conversation/relationship behavior belongs to later interaction/memory work, not M0. | M1+/interaction capability |
| `AUR-REQ-REL-002` | `APPLIES` | Directly constrains the M0 authority snapshot and operator authority semantics. | R2/R3 CAP-SOVEREIGN-CORE |
| `AUR-REQ-REL-003` | `DEFERRED_BY_ROADMAP` | Conversation/relationship behavior belongs to later interaction/memory work, not M0. | M1+/interaction capability |
| `AUR-REQ-REL-004` | `DEFERRED_BY_ROADMAP` | Conversation/relationship behavior belongs to later interaction/memory work, not M0. | M1+/interaction capability |
| `AUR-REQ-REL-005` | `PARTIALLY_APPLIES` | Stable Aurora identity applies; broader relationship/personality behavior is outside M0. | R2/R3 CAP-SOVEREIGN-CORE |
| `AUR-REQ-REL-006` | `DEFERRED_BY_ROADMAP` | Personality behavior is outside M0. | interaction capability |
| `AUR-REQ-REL-007` | `DEFERRED_BY_ROADMAP` | Personality behavior is outside M0. | interaction capability |
| `AUR-REQ-REL-008` | `DEFERRED_BY_ROADMAP` | Personality behavior is outside M0. | interaction capability |
| `AUR-REQ-REL-009` | `DEFERRED_BY_ROADMAP` | Attention/proactivity behavior is outside M0. | attention/proactivity capability |
| `AUR-REQ-REL-010` | `DEFERRED_BY_ROADMAP` | Attention/proactivity behavior is outside M0. | attention/proactivity capability |
| `AUR-REQ-REL-011` | `PARTIALLY_APPLIES` | Only the authority/preparation-vs-execution boundary applies; full interaction behavior is later. | R2/R3 CAP-SOVEREIGN-CORE |
| `AUR-REQ-REL-012` | `DEFERRED_BY_ROADMAP` | Conversation/relationship behavior belongs to later interaction/memory work, not M0. | interaction/evaluation |
| `AUR-REQ-REL-013` | `PARTIALLY_APPLIES` | Only current-state correction/incident integrity applies; conversational trust-repair behavior is later. | R2/R3 CAP-SOVEREIGN-CORE |
| `AUR-REQ-REL-014` | `DEFERRED_BY_ROADMAP` | Governed conversational/project memory is M1 scope, explicitly excluded from M0. | M1 memory capability |
| `AUR-REQ-REL-015` | `DEFERRED_BY_ROADMAP` | Governed conversational/project memory is M1 scope, explicitly excluded from M0. | M1 memory capability |
| `AUR-REQ-REL-016` | `DEFERRED_BY_ROADMAP` | Conversation/relationship behavior belongs to later interaction/memory work, not M0. | interaction capability |

### Blueprint 03 — Domain and World Model

Stable identities, Project state, authority distinction, provenance/time, snapshot ownership and Harness/provider separation are active. Mission/Delegation, provider approval, effects and physical-device entities are later.

| Requirement | Classification | Rationale | Downstream owner |
|---|---|---|---|
| `AUR-REQ-DOM-001` | `APPLIES` | Stable durable identifiers are core to M0 restart/recovery. | R2/R3 |
| `AUR-REQ-DOM-002` | `PARTIALLY_APPLIES` | Only M0 entities must be represented now; future Mission/Provider/Device/etc. concepts remain distinct but unimplemented. | R2/R3 |
| `AUR-REQ-DOM-003` | `APPLIES` | Negative invariant applies now to prevent a replaceable component from becoming Aurora/Core authority. | R2/R3 |
| `AUR-REQ-DOM-004` | `APPLIES` | Project state/next-action continuity is a direct M0 responsibility. | R2/R3 |
| `AUR-REQ-DOM-005` | `DEFERRED_BY_ROADMAP` | Mission/Delegation behavior is later orchestration scope. | M3+ |
| `AUR-REQ-DOM-006` | `APPLIES` | CAP-SOVEREIGN-CORE must describe reusable product outcome independently from any future Provider implementation. | R2/R3 |
| `AUR-REQ-DOM-007` | `DEFERRED_BY_ROADMAP` | Provider/Registry/trust behavior is M2+ scope. | M2+ |
| `AUR-REQ-DOM-008` | `PARTIALLY_APPLIES` | Evidence distinctions apply to M0 proof; hypothesis/experiment/measurement portions are later. | R2/R3 |
| `AUR-REQ-DOM-009` | `APPLIES` | Directly constrains the M0 authority snapshot and operator authority semantics. | R2/R3 |
| `AUR-REQ-DOM-010` | `DEFERRED_BY_ROADMAP` | External-effect execution is outside M0. | M3+/effect capability |
| `AUR-REQ-DOM-011` | `DEFERRED_BY_ROADMAP` | Device/firmware/lab behavior is M9/M10 scope. | M9/M10 |
| `AUR-REQ-DOM-012` | `APPLIES` | M0 state/authority/event relationships need owner, provenance and temporal validity. | R2/R3 |
| `AUR-REQ-DOM-013` | `APPLIES` | M0 must prove global state/authority is not owned by Harness/provider internals. | R2/R3 |
| `AUR-REQ-DOM-014` | `APPLIES` | Project/authority snapshots are projections over owners, not new sources of truth. | R2/R3 |

### Blueprint 04 — Cognitive Lifecycle and Journeys

M0 needs only a minimum interaction/state lifecycle, material records, observation/evidence, restart reconstruction and classified recovery. Reasoning, provider selection, campaigns, interruptions and concurrency remain later.

| Requirement | Classification | Rationale | Downstream owner |
|---|---|---|---|
| `AUR-REQ-COG-001` | `PARTIALLY_APPLIES` | M0 owns one minimum interaction/state lifecycle; the complete cognitive/interaction lifecycle remains later. | R2/R3 |
| `AUR-REQ-COG-002` | `PARTIALLY_APPLIES` | Only metadata needed for the minimal M0 interaction/event proof applies. | R2/R3 |
| `AUR-REQ-COG-003` | `DEFERRED_BY_ROADMAP` | Reasoning/ambiguity behavior is not needed to retire M0's persistence risk. | M1+/interaction |
| `AUR-REQ-COG-004` | `DEFERRED_BY_ROADMAP` | Reasoning/ambiguity behavior is not needed to retire M0's persistence risk. | M1+/interaction |
| `AUR-REQ-COG-005` | `PARTIALLY_APPLIES` | Preserve accepted intent/state separately from derived next action; full mission planning is later. | R2/R3 |
| `AUR-REQ-COG-006` | `DEFERRED_BY_ROADMAP` | Adaptive/long-running campaign behavior is later roadmap scope. | M7 |
| `AUR-REQ-COG-007` | `PARTIALLY_APPLIES` | M0 must prove global state/authority is not owned by Harness/provider internals. | R2/R3 |
| `AUR-REQ-COG-008` | `DEFERRED_BY_ROADMAP` | Provider/Registry/trust behavior is M2+ scope. | M2/M3 |
| `AUR-REQ-COG-009` | `PARTIALLY_APPLIES` | Only the M0 state-transition/action-observation chain applies. | R2/R3 |
| `AUR-REQ-COG-010` | `APPLIES` | Restart/restore/transition operations must be observed before claiming success. | R2/R3 |
| `AUR-REQ-COG-011` | `APPLIES` | M0 requires criterion-linked evidence and cannot close by activity or claim alone. | R2/R3 |
| `AUR-REQ-COG-012` | `APPLIES` | M0 records must route to state/audit/evidence owners rather than transcript narrative. | R2/R3 |
| `AUR-REQ-COG-013` | `DEFERRED_BY_ROADMAP` | Learning/self-improvement is later roadmap scope. | M11 |
| `AUR-REQ-COG-014` | `DEFERRED_BY_ROADMAP` | The listed multi-domain journeys are future milestone proofs, not M0 commitments. | future milestones |
| `AUR-REQ-COG-015` | `DEFERRED_BY_ROADMAP` | Attention/proactivity behavior is outside M0. | interaction/proactivity |
| `AUR-REQ-COG-016` | `DEFERRED_BY_ROADMAP` | Multi-mission/provider concurrency is not required by the M0 single walking skeleton. | M3+/M4 |
| `AUR-REQ-COG-017` | `APPLIES` | Direct M0 Golden Proof requirement for fresh-process reconstruction without transcript dependency. | R2/R3 |
| `AUR-REQ-COG-018` | `APPLIES` | M0 requires classified failure/recovery behavior for restart/restore/state operations. | R2/R3 |

### Blueprint 05 — Capabilities, Registry and AHDK

The Capability itself must be implementation-neutral, but Registry, Provider lifecycle, AHDK and conformance are M2+. Replaceable-binding semantics apply only as a guard against premature coupling.

| Requirement | Classification | Rationale | Downstream owner |
|---|---|---|---|
| `AUR-REQ-CAP-001` | `APPLIES` | CAP-SOVEREIGN-CORE itself must have a stable language/framework-independent identity/version even before Registry exists. | R2/R3 |
| `AUR-REQ-CAP-002` | `DEFERRED_BY_ROADMAP` | Capability Registry, Provider lifecycle, AHDK and conformance belong to M2+ and are explicit M0 non-dependencies. | M2+ |
| `AUR-REQ-CAP-003` | `DEFERRED_BY_ROADMAP` | Capability Registry, Provider lifecycle, AHDK and conformance belong to M2+ and are explicit M0 non-dependencies. | M2+ |
| `AUR-REQ-CAP-004` | `DEFERRED_BY_ROADMAP` | Capability Registry, Provider lifecycle, AHDK and conformance belong to M2+ and are explicit M0 non-dependencies. | M2+ |
| `AUR-REQ-CAP-005` | `DEFERRED_BY_ROADMAP` | Capability Registry, Provider lifecycle, AHDK and conformance belong to M2+ and are explicit M0 non-dependencies. | M2+ |
| `AUR-REQ-CAP-006` | `DEFERRED_BY_ROADMAP` | Capability Registry, Provider lifecycle, AHDK and conformance belong to M2+ and are explicit M0 non-dependencies. | M2+ |
| `AUR-REQ-CAP-007` | `DEFERRED_BY_ROADMAP` | Capability Registry, Provider lifecycle, AHDK and conformance belong to M2+ and are explicit M0 non-dependencies. | M2+ |
| `AUR-REQ-CAP-008` | `DEFERRED_BY_ROADMAP` | Capability Registry, Provider lifecycle, AHDK and conformance belong to M2+ and are explicit M0 non-dependencies. | M2+ |
| `AUR-REQ-CAP-009` | `DEFERRED_BY_ROADMAP` | Capability Registry, Provider lifecycle, AHDK and conformance belong to M2+ and are explicit M0 non-dependencies. | M2+ |
| `AUR-REQ-CAP-010` | `DEFERRED_BY_ROADMAP` | Capability Registry, Provider lifecycle, AHDK and conformance belong to M2+ and are explicit M0 non-dependencies. | M2+ |
| `AUR-REQ-CAP-011` | `DEFERRED_BY_ROADMAP` | Capability Registry, Provider lifecycle, AHDK and conformance belong to M2+ and are explicit M0 non-dependencies. | M2+ |
| `AUR-REQ-CAP-012` | `DEFERRED_BY_ROADMAP` | Capability Registry, Provider lifecycle, AHDK and conformance belong to M2+ and are explicit M0 non-dependencies. | M2+ |
| `AUR-REQ-CAP-013` | `DEFERRED_BY_ROADMAP` | Capability Registry, Provider lifecycle, AHDK and conformance belong to M2+ and are explicit M0 non-dependencies. | M2+ |
| `AUR-REQ-CAP-014` | `DEFERRED_BY_ROADMAP` | Capability Registry, Provider lifecycle, AHDK and conformance belong to M2+ and are explicit M0 non-dependencies. | M2+ |
| `AUR-REQ-CAP-015` | `DEFERRED_BY_ROADMAP` | Capability Registry, Provider lifecycle, AHDK and conformance belong to M2+ and are explicit M0 non-dependencies. | M2+ |
| `AUR-REQ-CAP-016` | `DEFERRED_BY_ROADMAP` | Capability Registry, Provider lifecycle, AHDK and conformance belong to M2+ and are explicit M0 non-dependencies. | M2+ |
| `AUR-REQ-CAP-017` | `DEFERRED_BY_ROADMAP` | Capability Registry, Provider lifecycle, AHDK and conformance belong to M2+ and are explicit M0 non-dependencies. | M2+ |
| `AUR-REQ-CAP-018` | `DEFERRED_BY_ROADMAP` | Capability Registry, Provider lifecycle, AHDK and conformance belong to M2+ and are explicit M0 non-dependencies. | M2+ |
| `AUR-REQ-CAP-019` | `DEFERRED_BY_ROADMAP` | Capability Registry, Provider lifecycle, AHDK and conformance belong to M2+ and are explicit M0 non-dependencies. | M2+ |
| `AUR-REQ-CAP-020` | `PARTIALLY_APPLIES` | Only the replaceability/semantic-neutrality guard applies; no binding is selected in M0. | R2/R3; R4 if boundary needed |

### Blueprint 06 — Memory, Knowledge and Context

Operational state must remain distinct from memory/history and cannot depend on model recall. Governed conversational/project memory, Context Builder and memory lifecycle are M1 and are not pulled into M0.

| Requirement | Classification | Rationale | Downstream owner |
|---|---|---|---|
| `AUR-REQ-MEM-001` | `APPLIES` | M0 operational state must remain distinct from memory/history/context. | R2/R3 |
| `AUR-REQ-MEM-002` | `DEFERRED_BY_ROADMAP` | Governed conversational/project memory is M1 scope, explicitly excluded from M0. | M1 |
| `AUR-REQ-MEM-003` | `PARTIALLY_APPLIES` | Only authoritative state/current snapshot boundaries apply; full memory strata are M1. | R2/R3 |
| `AUR-REQ-MEM-004` | `PARTIALLY_APPLIES` | Only authoritative state/current snapshot boundaries apply; full memory strata are M1. | R2/R3 |
| `AUR-REQ-MEM-005` | `APPLIES` | Narrative/observational memory must not become M0 current authority/state. | R2/R3 |
| `AUR-REQ-MEM-006` | `PARTIALLY_APPLIES` | Structured operational-state rule applies to M0 project/authority state; broader mission/device/budget state is later. | R2/R3 |
| `AUR-REQ-MEM-007` | `DEFERRED_BY_ROADMAP` | Governed conversational/project memory is M1 scope, explicitly excluded from M0. | M1 |
| `AUR-REQ-MEM-008` | `DEFERRED_BY_ROADMAP` | Governed conversational/project memory is M1 scope, explicitly excluded from M0. | M1 |
| `AUR-REQ-MEM-009` | `DEFERRED_BY_ROADMAP` | Governed conversational/project memory is M1 scope, explicitly excluded from M0. | M1 |
| `AUR-REQ-MEM-010` | `DEFERRED_BY_ROADMAP` | Governed conversational/project memory is M1 scope, explicitly excluded from M0. | M1 |
| `AUR-REQ-MEM-011` | `DEFERRED_BY_ROADMAP` | Governed conversational/project memory is M1 scope, explicitly excluded from M0. | M1 |
| `AUR-REQ-MEM-012` | `DEFERRED_BY_ROADMAP` | Governed conversational/project memory is M1 scope, explicitly excluded from M0. | M1 |
| `AUR-REQ-MEM-013` | `DEFERRED_BY_ROADMAP` | Governed conversational/project memory is M1 scope, explicitly excluded from M0. | M1 |
| `AUR-REQ-MEM-014` | `APPLIES` | Current accepted state/authority must override stale narrative and survive restart. | R2/R3 |
| `AUR-REQ-MEM-015` | `DEFERRED_BY_ROADMAP` | Governed conversational/project memory is M1 scope, explicitly excluded from M0. | M1 |
| `AUR-REQ-MEM-016` | `DEFERRED_BY_ROADMAP` | Governed conversational/project memory is M1 scope, explicitly excluded from M0. | M1 |
| `AUR-REQ-MEM-017` | `DEFERRED_BY_ROADMAP` | Governed conversational/project memory is M1 scope, explicitly excluded from M0. | M1 |
| `AUR-REQ-MEM-018` | `DEFERRED_BY_ROADMAP` | Governed conversational/project memory is M1 scope, explicitly excluded from M0. | M1 |
| `AUR-REQ-MEM-019` | `DEFERRED_BY_ROADMAP` | Governed conversational/project memory is M1 scope, explicitly excluded from M0. | M1 |
| `AUR-REQ-MEM-020` | `DEFERRED_BY_ROADMAP` | Governed conversational/project memory is M1 scope, explicitly excluded from M0. | M1 |
| `AUR-REQ-MEM-021` | `DEFERRED_BY_ROADMAP` | Governed conversational/project memory is M1 scope, explicitly excluded from M0. | M1 |
| `AUR-REQ-MEM-022` | `DEFERRED_BY_ROADMAP` | Governed conversational/project memory is M1 scope, explicitly excluded from M0. | M1 |
| `AUR-REQ-MEM-023` | `APPLIES` | M0 must not select/implement the later memory architecture merely to persist operational state. | R4 guard |

### Blueprint 07 — Harness Orchestration

M0 uses Harness/orchestration requirements primarily as negative ownership guards: Core state/authority cannot be provider-local. Cross-Harness Delegation behavior remains later.

| Requirement | Classification | Rationale | Downstream owner |
|---|---|---|---|
| `AUR-REQ-ORCH-001` | `DEFERRED_BY_ROADMAP` | Cross-Harness orchestration is later roadmap scope. | M3/M5 |
| `AUR-REQ-ORCH-002` | `PARTIALLY_APPLIES` | M0 proves the Core owns global state/authority; cross-domain/budget composition is later. | R2/R3 |
| `AUR-REQ-ORCH-003` | `PARTIALLY_APPLIES` | M0 must prove global state/authority is not owned by Harness/provider internals. | R2/R3 |
| `AUR-REQ-ORCH-004` | `DEFERRED_BY_ROADMAP` | Cross-Harness orchestration is later roadmap scope. | M3/M5 |
| `AUR-REQ-ORCH-005` | `DEFERRED_BY_ROADMAP` | Cross-Harness orchestration is later roadmap scope. | M3/M5 |
| `AUR-REQ-ORCH-006` | `DEFERRED_BY_ROADMAP` | Delegation/context-pack behavior belongs to M3+. | M3 |
| `AUR-REQ-ORCH-007` | `DEFERRED_BY_ROADMAP` | Delegation/context-pack behavior belongs to M3+. | M3 |
| `AUR-REQ-ORCH-008` | `DEFERRED_BY_ROADMAP` | Delegation/context-pack behavior belongs to M3+. | M3/M5 |
| `AUR-REQ-ORCH-009` | `DEFERRED_BY_ROADMAP` | Delegation/context-pack behavior belongs to M3+. | M3 |
| `AUR-REQ-ORCH-010` | `DEFERRED_BY_ROADMAP` | Delegation/context-pack behavior belongs to M3+. | M3/M4 |
| `AUR-REQ-ORCH-011` | `APPLIES` | Event/audit minimum applies and events must not become the sole current-state authority. | R2/R3 |
| `AUR-REQ-ORCH-012` | `PARTIALLY_APPLIES` | Only material current status/recovery state should surface through the simple interface. | R2/R3 |
| `AUR-REQ-ORCH-013` | `DEFERRED_BY_ROADMAP` | Structured Decision Request behavior belongs to later mission/delegation orchestration. | M3 |
| `AUR-REQ-ORCH-014` | `PARTIALLY_APPLIES` | Evidence distinctions apply, without a Harness boundary in M0. | R2/R3 |
| `AUR-REQ-ORCH-015` | `DEFERRED_BY_ROADMAP` | Bulk/direct data channels are later cross-provider capability scope. | M3+ |
| `AUR-REQ-ORCH-016` | `APPLIES` | Explicitly prevents Harness local state from satisfying M0 global persistence. | R2/R3 |
| `AUR-REQ-ORCH-017` | `DEFERRED_BY_ROADMAP` | Delegation/context-pack behavior belongs to M3+. | M3/M4 |
| `AUR-REQ-ORCH-018` | `DEFERRED_BY_ROADMAP` | Provider resume/checkpoint/heartbeat semantics are later M2/M4 scope. | M2/M4 |
| `AUR-REQ-ORCH-019` | `DEFERRED_BY_ROADMAP` | Provider retry semantics are later M2/M4 scope. | M2/M4 |
| `AUR-REQ-ORCH-020` | `APPLIES` | ADR-0001 guard: transport/framework semantics may not redefine Aurora's M0 domain. | R2/R3; ADR-0001 guard |
| `AUR-REQ-ORCH-021` | `APPLIES` | Internal framework state must not become Aurora authority/state. | R2/R3 |
| `AUR-REQ-ORCH-022` | `APPLIES` | MNFS remains a future provider and cannot become a dependency of the sovereign Core. | R2/R3 |

### Blueprint 08 — Interaction, Multimodality and Presence

M0 has one simple interface only. Identity must not be owned by that interface and no UI technology may be assumed; multi-presence/sensor/handoff behavior is M8.

| Requirement | Classification | Rationale | Downstream owner |
|---|---|---|---|
| `AUR-REQ-PRS-001` | `PARTIALLY_APPLIES` | M0 has one simple interface, but Aurora identity must not belong to that Presence/interface. | R2/R3 |
| `AUR-REQ-PRS-002` | `DEFERRED_BY_ROADMAP` | Multi-presence, sensors, handoff and channel-specific behavior are M8 scope. | M8 |
| `AUR-REQ-PRS-003` | `DEFERRED_BY_ROADMAP` | Multi-presence, sensors, handoff and channel-specific behavior are M8 scope. | M8 |
| `AUR-REQ-PRS-004` | `DEFERRED_BY_ROADMAP` | Multi-presence, sensors, handoff and channel-specific behavior are M8 scope. | M8 |
| `AUR-REQ-PRS-005` | `DEFERRED_BY_ROADMAP` | Multi-presence, sensors, handoff and channel-specific behavior are M8 scope. | M8 |
| `AUR-REQ-PRS-006` | `DEFERRED_BY_ROADMAP` | Multi-presence, sensors, handoff and channel-specific behavior are M8 scope. | M8 |
| `AUR-REQ-PRS-007` | `DEFERRED_BY_ROADMAP` | Multi-presence, sensors, handoff and channel-specific behavior are M8 scope. | M8 |
| `AUR-REQ-PRS-008` | `DEFERRED_BY_ROADMAP` | Multi-presence, sensors, handoff and channel-specific behavior are M8 scope. | M8 |
| `AUR-REQ-PRS-009` | `DEFERRED_BY_ROADMAP` | Multi-presence, sensors, handoff and channel-specific behavior are M8 scope. | M8 |
| `AUR-REQ-PRS-010` | `DEFERRED_BY_ROADMAP` | Multi-presence, sensors, handoff and channel-specific behavior are M8 scope. | M8 |
| `AUR-REQ-PRS-011` | `DEFERRED_BY_ROADMAP` | Multi-presence, sensors, handoff and channel-specific behavior are M8 scope. | M8 |
| `AUR-REQ-PRS-012` | `DEFERRED_BY_ROADMAP` | Multi-presence, sensors, handoff and channel-specific behavior are M8 scope. | M8 |
| `AUR-REQ-PRS-013` | `DEFERRED_BY_ROADMAP` | Multi-presence, sensors, handoff and channel-specific behavior are M8 scope. | M8 |
| `AUR-REQ-PRS-014` | `DEFERRED_BY_ROADMAP` | Multi-presence, sensors, handoff and channel-specific behavior are M8 scope. | M8 |
| `AUR-REQ-PRS-015` | `APPLIES` | The simple interface must not imply a specific UI/voice/glasses technology. | R4 guard |

### Blueprint 09 — Tools, Devices and Laboratory

All laboratory/device requirements are deferred because M0 explicitly excludes physical devices. They remain constitutional and must not be weakened.

| Requirement | Classification | Rationale | Downstream owner |
|---|---|---|---|
| `AUR-REQ-LAB-001` | `DEFERRED_BY_ROADMAP` | Physical-device/laboratory behavior is explicitly excluded from M0 and deferred to M9/M10. | M9/M10 |
| `AUR-REQ-LAB-002` | `DEFERRED_BY_ROADMAP` | Physical-device/laboratory behavior is explicitly excluded from M0 and deferred to M9/M10. | M9/M10 |
| `AUR-REQ-LAB-003` | `DEFERRED_BY_ROADMAP` | Physical-device/laboratory behavior is explicitly excluded from M0 and deferred to M9/M10. | M9/M10 |
| `AUR-REQ-LAB-004` | `DEFERRED_BY_ROADMAP` | Physical-device/laboratory behavior is explicitly excluded from M0 and deferred to M9/M10. | M9/M10 |
| `AUR-REQ-LAB-005` | `DEFERRED_BY_ROADMAP` | Physical-device/laboratory behavior is explicitly excluded from M0 and deferred to M9/M10. | M9/M10 |
| `AUR-REQ-LAB-006` | `DEFERRED_BY_ROADMAP` | Physical-device/laboratory behavior is explicitly excluded from M0 and deferred to M9/M10. | M9/M10 |
| `AUR-REQ-LAB-007` | `DEFERRED_BY_ROADMAP` | Physical-device/laboratory behavior is explicitly excluded from M0 and deferred to M9/M10. | M9/M10 |
| `AUR-REQ-LAB-008` | `DEFERRED_BY_ROADMAP` | Physical-device/laboratory behavior is explicitly excluded from M0 and deferred to M9/M10. | M9/M10 |
| `AUR-REQ-LAB-009` | `DEFERRED_BY_ROADMAP` | Physical-device/laboratory behavior is explicitly excluded from M0 and deferred to M9/M10. | M9/M10 |
| `AUR-REQ-LAB-010` | `DEFERRED_BY_ROADMAP` | Physical-device/laboratory behavior is explicitly excluded from M0 and deferred to M9/M10. | M9/M10 |
| `AUR-REQ-LAB-011` | `DEFERRED_BY_ROADMAP` | Physical-device/laboratory behavior is explicitly excluded from M0 and deferred to M9/M10. | M9/M10 |
| `AUR-REQ-LAB-012` | `DEFERRED_BY_ROADMAP` | Physical-device/laboratory behavior is explicitly excluded from M0 and deferred to M9/M10. | M9/M10 |
| `AUR-REQ-LAB-013` | `DEFERRED_BY_ROADMAP` | Physical-device/laboratory behavior is explicitly excluded from M0 and deferred to M9/M10. | M9/M10 |
| `AUR-REQ-LAB-014` | `DEFERRED_BY_ROADMAP` | Physical-device/laboratory behavior is explicitly excluded from M0 and deferred to M9/M10. | M9/M10 |
| `AUR-REQ-LAB-015` | `DEFERRED_BY_ROADMAP` | Physical-device/laboratory behavior is explicitly excluded from M0 and deferred to M9/M10. | M9/M10 |
| `AUR-REQ-LAB-016` | `DEFERRED_BY_ROADMAP` | Physical-device/laboratory behavior is explicitly excluded from M0 and deferred to M9/M10. | M9/M10 |
| `AUR-REQ-LAB-017` | `DEFERRED_BY_ROADMAP` | Physical-device/laboratory behavior is explicitly excluded from M0 and deferred to M9/M10. | M9/M10 |
| `AUR-REQ-LAB-018` | `DEFERRED_BY_ROADMAP` | Physical-device/laboratory behavior is explicitly excluded from M0 and deferred to M9/M10. | M9/M10 |
| `AUR-REQ-LAB-019` | `DEFERRED_BY_ROADMAP` | Physical-device/laboratory behavior is explicitly excluded from M0 and deferred to M9/M10. | M9/M10 |
| `AUR-REQ-LAB-020` | `DEFERRED_BY_ROADMAP` | Physical-device/laboratory behavior is explicitly excluded from M0 and deferred to M9/M10. | M9/M10 |

### Blueprint 10 — Autonomy, Authority and Safety

M0 owns a minimal current-authority snapshot sufficient for safe restart/restore. Effects, credentials, budgets, campaigns, interlocks and emergency authority are later.

| Requirement | Classification | Rationale | Downstream owner |
|---|---|---|---|
| `AUR-REQ-AUT-001` | `PARTIALLY_APPLIES` | Only persisted current authority/authority-snapshot semantics apply; delegated autonomy is later. | R2/R3 |
| `AUR-REQ-AUT-002` | `APPLIES` | Directly constrains the M0 authority snapshot and operator authority semantics. | R2/R3 |
| `AUR-REQ-AUT-003` | `APPLIES` | Directly constrains the M0 authority snapshot and operator authority semantics. | R2/R3 |
| `AUR-REQ-AUT-004` | `DEFERRED_BY_ROADMAP` | Adaptive autonomy belongs to later milestones. | M7 |
| `AUR-REQ-AUT-005` | `DEFERRED_BY_ROADMAP` | Adaptive/long-running campaign behavior is later roadmap scope. | M7 |
| `AUR-REQ-AUT-006` | `DEFERRED_BY_ROADMAP` | Adaptive/long-running campaign behavior is later roadmap scope. | M7 |
| `AUR-REQ-AUT-007` | `PARTIALLY_APPLIES` | M0 must preserve enough authority identity/scope/validity to recover safely; full Delegation/effect grant model is later. | R2/R3 |
| `AUR-REQ-AUT-008` | `DEFERRED_BY_ROADMAP` | External-effect execution is outside M0. | M3+ |
| `AUR-REQ-AUT-009` | `DEFERRED_BY_ROADMAP` | External-effect execution is outside M0. | M3+ |
| `AUR-REQ-AUT-010` | `DEFERRED_BY_ROADMAP` | Credential brokering/effect authorization is later effect-plane scope. | M3+ |
| `AUR-REQ-AUT-011` | `DEFERRED_BY_ROADMAP` | Budget enforcement belongs to delegated/durable campaign milestones, not M0. | M3/M4/M7 |
| `AUR-REQ-AUT-012` | `DEFERRED_BY_ROADMAP` | Physical-device/laboratory behavior is explicitly excluded from M0 and deferred to M9/M10. | M10 |
| `AUR-REQ-AUT-013` | `DEFERRED_BY_ROADMAP` | Adaptive/long-running campaign behavior is later roadmap scope. | M7 |
| `AUR-REQ-AUT-014` | `PARTIALLY_APPLIES` | Restore/restart must preserve revoked/expired authority state even though no external effects are executed. | R2/R3 |
| `AUR-REQ-AUT-015` | `DEFERRED_BY_ROADMAP` | External-effect execution is outside M0. | M3+ |
| `AUR-REQ-AUT-016` | `DEFERRED_BY_ROADMAP` | Emergency/physical containment behavior is future physical-autonomy scope. | M10 |
| `AUR-REQ-AUT-017` | `DEFERRED_BY_ROADMAP` | Effect confirmation UX is outside M0. | M3+ |
| `AUR-REQ-AUT-018` | `DEFERRED_BY_ROADMAP` | Self-improvement/evolution is M11 scope. | M11 |
| `AUR-REQ-AUT-019` | `DEFERRED_BY_ROADMAP` | Self-improvement/evolution is M11 scope. | M11 |
| `AUR-REQ-AUT-020` | `PARTIALLY_APPLIES` | Only corruption/invalid-transition authority incidents need M0 threat/recovery consideration; full authority incident program is later. | R3 threat/recovery |

### Blueprint 11 — Security, Privacy and Sovereignty

The first canonical durable store makes sovereignty, threat modeling, data classification, authority integrity, inspectability and restore security material now. Provider/presence/physical/effect-specific controls remain later.

| Requirement | Classification | Rationale | Downstream owner |
|---|---|---|---|
| `AUR-REQ-SEC-001` | `APPLIES` | Security/sovereignty is cross-cutting because M0 creates the first canonical durable store. | R3 threat model |
| `AUR-REQ-SEC-002` | `PARTIALLY_APPLIES` | M0 must keep relevant planes distinct; provider/effect/presence/supply-chain portions are not implemented yet. | R2/R3 |
| `AUR-REQ-SEC-003` | `APPLIES` | A material sovereign-state Capability requires a threat model at R3. | R3 |
| `AUR-REQ-SEC-004` | `PARTIALLY_APPLIES` | The sovereign local Core/environment boundary applies; multi-provider/device trust zones are later. | R3/R4 |
| `AUR-REQ-SEC-005` | `APPLIES` | Persisted project/authority/audit/export data requires explicit classification semantics. | R2/R3 |
| `AUR-REQ-SEC-006` | `DEFERRED_BY_ROADMAP` | Cross-provider data transfer is outside M0. | M2/M3 |
| `AUR-REQ-SEC-007` | `APPLIES` | Identity, authority, audit and operational state must remain under Leandro-controlled governance. | R2/R3/R4 |
| `AUR-REQ-SEC-008` | `DEFERRED_BY_ROADMAP` | Provider/Registry/trust behavior is M2+ scope. | M2/M3 |
| `AUR-REQ-SEC-009` | `PARTIALLY_APPLIES` | General secret/log hygiene applies defensively; M0 does not yet implement a credential system. | R3 |
| `AUR-REQ-SEC-010` | `PARTIALLY_APPLIES` | Actor attribution applies to operator/Core transitions; Harness/worker/effect chains are later. | R2/R3 |
| `AUR-REQ-SEC-011` | `DEFERRED_BY_ROADMAP` | External-effect execution is outside M0. | M3+ |
| `AUR-REQ-SEC-012` | `DEFERRED_BY_ROADMAP` | Provider/Harness sandbox containment is later provider/effect capability scope. | M2/M3 |
| `AUR-REQ-SEC-013` | `PARTIALLY_APPLIES` | Input/content must not silently redefine policy/authority; especially relevant to persisted project state. | R2/R3 |
| `AUR-REQ-SEC-014` | `DEFERRED_BY_ROADMAP` | Governed conversational/project memory is M1 scope, explicitly excluded from M0. | M1 |
| `AUR-REQ-SEC-015` | `DEFERRED_BY_ROADMAP` | Transcript/audio/video/memory retention is part of later M1/M8 capabilities. | M1/M8 |
| `AUR-REQ-SEC-016` | `DEFERRED_BY_ROADMAP` | Provider/Registry/trust behavior is M2+ scope. | M2 |
| `AUR-REQ-SEC-017` | `DEFERRED_BY_ROADMAP` | Multi-presence, sensors, handoff and channel-specific behavior are M8 scope. | M8 |
| `AUR-REQ-SEC-018` | `DEFERRED_BY_ROADMAP` | Physical-device/laboratory behavior is explicitly excluded from M0 and deferred to M9/M10. | M10 |
| `AUR-REQ-SEC-019` | `PARTIALLY_APPLIES` | Leandro must inspect M0 state/authority/audit; provider/effect disclosures are later. | R2/R3 |
| `AUR-REQ-SEC-020` | `PARTIALLY_APPLIES` | Integrity/retention/supersession apply now; the full memory/privacy deletion program is later. | R2/R3 |
| `AUR-REQ-SEC-021` | `PARTIALLY_APPLIES` | Export, backup/restore and migration apply now; provider revocation/credential rotation/full deletion are later. | R2/R3/R4 |
| `AUR-REQ-SEC-022` | `APPLIES` | Direct M0 security invariant: restore may not reactivate expired/revoked authority or compromised trust. | R2/R3 |
| `AUR-REQ-SEC-023` | `DEFERRED_BY_ROADMAP` | External-effect execution is outside M0. | M4+ |
| `AUR-REQ-SEC-024` | `PARTIALLY_APPLIES` | M0 needs recovery/evidence behavior for state/security failures; full operational incident program can grow later. | R3 |
| `AUR-REQ-SEC-025` | `DEFERRED_BY_ROADMAP` | Broad authority usability/step-up behavior belongs to later effect/presence capabilities. | M3+ |

### Blueprint 12 — System Architecture

This is the strongest M0 source: sovereign Core, durable state, operational-state ownership, signal distinctions, simple topology, version/migration and failure/recovery constraints all shape later R2–R4 work without selecting a stack.

| Requirement | Classification | Rationale | Downstream owner |
|---|---|---|---|
| `AUR-REQ-ARC-001` | `APPLIES` | Direct architectural requirement for the M0 Capability. | R2/R3/R4 |
| `AUR-REQ-ARC-002` | `APPLIES` | M0 domain/state types cannot be owned by database/framework/runtime semantics. | R2/R3 |
| `AUR-REQ-ARC-003` | `APPLIES` | Logical ownership must be clear before R4 chooses topology. | R3/R4 |
| `AUR-REQ-ARC-004` | `APPLIES` | Direct M0 requirement: state outlives process. | R2/R3 |
| `AUR-REQ-ARC-005` | `DEFERRED_BY_ROADMAP` | Bulk/high-rate data paths are outside M0. | M3+/M9 |
| `AUR-REQ-ARC-006` | `PARTIALLY_APPLIES` | Explicit owners are required for M0 identity/project/state/authority/audit/interface; future modules need not be built. | R3 |
| `AUR-REQ-ARC-007` | `PARTIALLY_APPLIES` | M0 domain contracts must remain implementation-neutral; full cross-provider Contract Model is later. | R2/R3 |
| `AUR-REQ-ARC-008` | `DEFERRED_BY_ROADMAP` | Provider discovery/dispatch fabric is M2/M3 scope. | M2/M3 |
| `AUR-REQ-ARC-009` | `APPLIES` | Direct requirement for canonical current state distinct from conversation/engine/Harness/Git/logs. | R2/R3 |
| `AUR-REQ-ARC-010` | `APPLIES` | M0 includes event/audit minimum plus telemetry baseline, so signal types cannot collapse. | R2/R3 |
| `AUR-REQ-ARC-011` | `DEFERRED_BY_ROADMAP` | Direct/bulk data-plane channels are later. | M3+ |
| `AUR-REQ-ARC-012` | `PARTIALLY_APPLIES` | Authoritative state/snapshot precedence applies; full memory topology is later. | R2/R3 |
| `AUR-REQ-ARC-013` | `PARTIALLY_APPLIES` | Evidence relationships must remain distinct even if M0 uses a simple physical store. | R2/R3 |
| `AUR-REQ-ARC-014` | `DEFERRED_BY_ROADMAP` | External-effect execution is outside M0. | M3+ |
| `AUR-REQ-ARC-015` | `DEFERRED_BY_ROADMAP` | Device plane architecture is M9/M10 scope. | M9/M10 |
| `AUR-REQ-ARC-016` | `PARTIALLY_APPLIES` | The CLI/simple interface is an adapter and cannot own canonical M0 state. | R2/R3 |
| `AUR-REQ-ARC-017` | `APPLIES` | M0 cannot bind Aurora identity/state to a model, even if a model is absent. | R2/R3 |
| `AUR-REQ-ARC-018` | `APPLIES` | R4 must prefer the smallest recoverable topology and justify any distribution. | R4 |
| `AUR-REQ-ARC-019` | `PARTIALLY_APPLIES` | M0 architecture should preserve clean adapters without implementing later distributed stages. | R3/R4 |
| `AUR-REQ-ARC-020` | `APPLIES` | Restart/store failure recovery is part of the M0 risk being retired. | R2/R3 |
| `AUR-REQ-ARC-021` | `PARTIALLY_APPLIES` | M0 state/schema migration and restore compatibility apply; provider/delegation versioning is later. | R2/R3/R4 |
| `AUR-REQ-ARC-022` | `PARTIALLY_APPLIES` | Project/proof/state/recovery correlation applies; mission/provider/effect IDs are later. | R2/R3 |
| `AUR-REQ-ARC-023` | `APPLIES` | Authority duplication, restart, framework leakage and proof fitness questions directly apply. | R3/R4 |

### Blueprint 13 — Reliability, Evaluation and Self-Improvement

M0 must prove recovery with evidence and decision-oriented telemetry. Self-improvement, formal candidate evaluation and specialized future evaluation suites remain later.

| Requirement | Classification | Rationale | Downstream owner |
|---|---|---|---|
| `AUR-REQ-RELIA-001` | `APPLIES` | M0 closeout must evaluate continuity/correctness/authority/operations/evidence/security/efficiency as applicable. | R3/R8 |
| `AUR-REQ-RELIA-002` | `APPLIES` | M0 requires criterion-linked evidence and cannot close by activity or claim alone. | R2/R3 |
| `AUR-REQ-RELIA-003` | `APPLIES` | M0 requires criterion-linked evidence and cannot close by activity or claim alone. | R2/R3 |
| `AUR-REQ-RELIA-004` | `APPLIES` | M0 cannot close from component/task completion alone. | R8 |
| `AUR-REQ-RELIA-005` | `APPLIES` | M0 telemetry/logs must support decisions and not become domain truth. | R2/R3 |
| `AUR-REQ-RELIA-006` | `APPLIES` | Every retained M0 signal must have an explicit proof/decision purpose. | R2/R3 |
| `AUR-REQ-RELIA-007` | `PARTIALLY_APPLIES` | Only M0-relevant contract/operational/security/state error classes must be realized now. | R2/R3 |
| `AUR-REQ-RELIA-008` | `APPLIES` | Any M0 retry must be classified/safe/idempotent; blind retry could corrupt state/restore. | R2/R3 |
| `AUR-REQ-RELIA-009` | `PARTIALLY_APPLIES` | M0 needs capability/journey/security/recovery evaluation, not the later full longitudinal suite. | R3 |
| `AUR-REQ-RELIA-010` | `DEFERRED_BY_ROADMAP` | Formal eval-dataset role separation is not needed for this deterministic walking skeleton. | later eval capability |
| `AUR-REQ-RELIA-011` | `DEFERRED_BY_ROADMAP` | Candidate/self-improvement evaluation is M11 scope. | M11 |
| `AUR-REQ-RELIA-012` | `DEFERRED_BY_ROADMAP` | Probabilistic confidence modeling is not required by the deterministic M0 continuity proof. | M1+/model use |
| `AUR-REQ-RELIA-013` | `DEFERRED_BY_ROADMAP` | Feedback-learning governance is M11 scope. | M11 |
| `AUR-REQ-RELIA-014` | `DEFERRED_BY_ROADMAP` | Self-improvement/evolution is M11 scope. | M11 |
| `AUR-REQ-RELIA-015` | `DEFERRED_BY_ROADMAP` | Self-improvement/evolution is M11 scope. | M11 |
| `AUR-REQ-RELIA-016` | `DEFERRED_BY_ROADMAP` | Self-improvement/evolution is M11 scope. | M11 |
| `AUR-REQ-RELIA-017` | `DEFERRED_BY_ROADMAP` | Self-improvement/evolution is M11 scope. | M11 |
| `AUR-REQ-RELIA-018` | `DEFERRED_BY_ROADMAP` | Self-improvement/evolution is M11 scope. | M11 |
| `AUR-REQ-RELIA-019` | `DEFERRED_BY_ROADMAP` | Self-improvement/evolution is M11 scope. | M11 |
| `AUR-REQ-RELIA-020` | `DEFERRED_BY_ROADMAP` | Self-improvement/evolution is M11 scope. | M11 |
| `AUR-REQ-RELIA-021` | `DEFERRED_BY_ROADMAP` | Self-improvement/evolution is M11 scope. | M11 |
| `AUR-REQ-RELIA-022` | `DEFERRED_BY_ROADMAP` | Promotion to procedural memory/Golden Paths is later learning scope. | M11 |
| `AUR-REQ-RELIA-023` | `DEFERRED_BY_ROADMAP` | Dedicated memory/personality/orchestration/physical evaluation suites belong to their future capabilities. | M1/M3+/M9 |
| `AUR-REQ-RELIA-024` | `APPLIES` | Material M0 incidents must update evidence/docs/requirements rather than only future code. | R3/R8 |

### Blueprint 14 — Capability Roadmap

M0 roadmap anatomy/Golden Proof/gate separation apply. Requirements that specifically define M1–M12 are deferred; the historical A0 closeout rule is not CAP behavior.

| Requirement | Classification | Rationale | Downstream owner |
|---|---|---|---|
| `AUR-REQ-RDM-001` | `APPLIES` | Directly governs the selected M0 milestone/gate/Golden Proof and readiness boundaries. | ACRM/R2+ |
| `AUR-REQ-RDM-002` | `APPLIES` | Directly governs the selected M0 milestone/gate/Golden Proof and readiness boundaries. | ACRM/R2+ |
| `AUR-REQ-RDM-003` | `APPLIES` | Directly governs the selected M0 milestone/gate/Golden Proof and readiness boundaries. | ACRM/R2+ |
| `AUR-REQ-RDM-004` | `APPLIES` | Directly governs the selected M0 milestone/gate/Golden Proof and readiness boundaries. | ACRM/R2+ |
| `AUR-REQ-RDM-005` | `APPLIES` | Directly governs the selected M0 milestone/gate/Golden Proof and readiness boundaries. | ACRM/R2+ |
| `AUR-REQ-RDM-006` | `NOT_APPLICABLE` | Historical A0-only gate requirement; already satisfied and not behavior of CAP-SOVEREIGN-CORE. | A0 closeout evidence |
| `AUR-REQ-RDM-007` | `APPLIES` | Directly governs the selected M0 milestone/gate/Golden Proof and readiness boundaries. | ACRM/R2+ |
| `AUR-REQ-RDM-008` | `DEFERRED_BY_ROADMAP` | This requirement defines a later Product Milestone and is explicitly deferred by the accepted roadmap. | M1-M12 as identified by roadmap |
| `AUR-REQ-RDM-009` | `DEFERRED_BY_ROADMAP` | This requirement defines a later Product Milestone and is explicitly deferred by the accepted roadmap. | M1-M12 as identified by roadmap |
| `AUR-REQ-RDM-010` | `DEFERRED_BY_ROADMAP` | This requirement defines a later Product Milestone and is explicitly deferred by the accepted roadmap. | M1-M12 as identified by roadmap |
| `AUR-REQ-RDM-011` | `DEFERRED_BY_ROADMAP` | This requirement defines a later Product Milestone and is explicitly deferred by the accepted roadmap. | M1-M12 as identified by roadmap |
| `AUR-REQ-RDM-012` | `DEFERRED_BY_ROADMAP` | This requirement defines a later Product Milestone and is explicitly deferred by the accepted roadmap. | M1-M12 as identified by roadmap |
| `AUR-REQ-RDM-013` | `DEFERRED_BY_ROADMAP` | This requirement defines a later Product Milestone and is explicitly deferred by the accepted roadmap. | M1-M12 as identified by roadmap |
| `AUR-REQ-RDM-014` | `DEFERRED_BY_ROADMAP` | This requirement defines a later Product Milestone and is explicitly deferred by the accepted roadmap. | M1-M12 as identified by roadmap |
| `AUR-REQ-RDM-015` | `DEFERRED_BY_ROADMAP` | This requirement defines a later Product Milestone and is explicitly deferred by the accepted roadmap. | M1-M12 as identified by roadmap |
| `AUR-REQ-RDM-016` | `DEFERRED_BY_ROADMAP` | This requirement defines a later Product Milestone and is explicitly deferred by the accepted roadmap. | M1-M12 as identified by roadmap |
| `AUR-REQ-RDM-017` | `DEFERRED_BY_ROADMAP` | This requirement defines a later Product Milestone and is explicitly deferred by the accepted roadmap. | M1-M12 as identified by roadmap |
| `AUR-REQ-RDM-018` | `DEFERRED_BY_ROADMAP` | This requirement defines a later Product Milestone and is explicitly deferred by the accepted roadmap. | M1-M12 as identified by roadmap |
| `AUR-REQ-RDM-019` | `DEFERRED_BY_ROADMAP` | This requirement defines a later Product Milestone and is explicitly deferred by the accepted roadmap. | M1-M12 as identified by roadmap |
| `AUR-REQ-RDM-020` | `APPLIES` | Directly governs the selected M0 milestone/gate/Golden Proof and readiness boundaries. | ACRM/R2+ |

### Blueprint 15 — Documentation and Research Governance

R1/R2 artifacts must preserve canonical ownership, status, evidence, validation and explicit authorization. A0-only/global publication requirements that CAP-SOVEREIGN-CORE does not implement are marked not applicable.

| Requirement | Classification | Rationale | Downstream owner |
|---|---|---|---|
| `AUR-REQ-DOC-001` | `APPLIES` | Applies to R1+ artifact ownership and conflict handling. | R1+ all artifacts |
| `AUR-REQ-DOC-002` | `APPLIES` | This gate authorization and durable decision must be promoted to repository evidence. | current authorization/evidence |
| `AUR-REQ-DOC-003` | `APPLIES` | Research/open technology candidates may inform R4 but cannot become R1 decisions. | R4 guard |
| `AUR-REQ-DOC-004` | `APPLIES` | The R1 applicability artifact needs identity, lifecycle, owner and source scope. | R1 artifact |
| `AUR-REQ-DOC-005` | `NOT_APPLICABLE` | Repository-wide Blueprint publication requirement was satisfied by A0; CAP-SOVEREIGN-CORE does not implement it. | already satisfied A0 baseline |
| `AUR-REQ-DOC-006` | `APPLIES` | The first capability may create only documentation with a present R1 consumer. | R1 artifact/layout |
| `AUR-REQ-DOC-007` | `PARTIALLY_APPLIES` | Applies conditionally if M0 R4 needs focused research/spikes; R1 performs no new research. | R4 if research |
| `AUR-REQ-DOC-008` | `PARTIALLY_APPLIES` | Applies conditionally if M0 R4 needs focused research/spikes; R1 performs no new research. | R4 if research |
| `AUR-REQ-DOC-009` | `PARTIALLY_APPLIES` | Applies conditionally if M0 R4 needs focused research/spikes; R1 performs no new research. | R4 if research |
| `AUR-REQ-DOC-010` | `PARTIALLY_APPLIES` | Applies conditionally if R4 requires a material decision; R1 selects no technical decision. | R4 if ADR needed |
| `AUR-REQ-DOC-011` | `APPLIES` | Capability Spec and later Mission Contract must remain separate owners. | R3/R5 |
| `AUR-REQ-DOC-012` | `APPLIES` | STATUS must reflect R1 authorization/verdict and the next explicit boundary. | tracking |
| `AUR-REQ-DOC-013` | `APPLIES` | Later revisions must preserve history and update discovery paths. | all readiness artifacts |
| `AUR-REQ-DOC-014` | `NOT_APPLICABLE` | R1 applicability is a canonical source artifact, not a generated projection. | no generated projection owned by R1 |
| `AUR-REQ-DOC-015` | `APPLIES` | R1 repository changes must declare documentation impact. | all material changes |
| `AUR-REQ-DOC-016` | `APPLIES` | R1 artifact/tracking changes must pass documentation validation. | all readiness artifacts |
| `AUR-REQ-DOC-017` | `APPLIES` | R1 classification needs adversarial exclusion/dependency review. | R1/R3 |
| `AUR-REQ-DOC-018` | `APPLIES` | Future agents should load this capability's smallest correct authority set. | all agents/review |
| `AUR-REQ-DOC-019` | `NOT_APPLICABLE` | Historical A0-only gate requirement; already satisfied and not behavior of CAP-SOVEREIGN-CORE. | A0 closeout evidence |
| `AUR-REQ-DOC-020` | `APPLIES` | Git writes/CI do not replace explicit gate authorization. | all gates |
| `AUR-REQ-DOC-021` | `APPLIES` | Material applicability conflicts/gaps must be recorded rather than silently classified. | all gates |
| `AUR-REQ-DOC-022` | `NOT_APPLICABLE` | A0 implementation gate already satisfied; current prohibition comes from ACRM/STATUS, not this historical CAP behavior. | A0 closeout/current ACRM |

## 6. High-risk dependencies carried forward

R1 identifies the following dependencies as material to M0. None is permission to execute the downstream work.

### D1 — Authority snapshot correctness and restore safety

M0 cannot persist an arbitrary “allowed next action” string and call it authority. R2/R3 must define the minimum current-authority semantics needed for the slice, including freshness/validity and the rule that restore cannot silently reactivate expired/revoked authority or compromised trust.

Owners downstream:
- R2 — atomic authority-snapshot/recovery requirements;
- R3 — CAP-SOVEREIGN-CORE domain/lifecycle/threat model;
- R4 — only the mechanism needed to realize those accepted semantics.

### D2 — Operational-state ownership

Project/current state, authority snapshot, event/audit and evidence must remain distinct from:
- transcript/history;
- model memory;
- Harness/provider local state;
- workflow-engine history;
- Git documentation;
- logs/telemetry.

R2/R3 must define the minimum source-of-truth and projection boundaries before R4 selects storage.

### D3 — Stable identity and migration

Aurora/Project/state identities must survive process restart and restore. M0 also has a migration prerequisite. R2/R3 must define identity/version invariants before R4 selects schemas/store/migration mechanism.

### D4 — Backup/restore integrity and portability

The Golden Proof includes export/restore. R2/R3 must define what state is exported/restored, integrity expectations, and authority-safety behavior. R4 may then compare mechanisms/topologies.

### D5 — Event/audit versus telemetry

M0 requires an event/audit minimum and a telemetry baseline. They must not collapse into one source of truth. R2/R3 define semantics/evidence; R4 chooses any backend/transport only if required.

### D6 — Security of the first canonical durable store

Because M0 creates the first sovereign durable state, R3 requires a threat model covering at minimum:
- confidentiality/integrity/availability;
- local sovereign ownership;
- state corruption/tampering;
- unsafe restore;
- authority-state corruption;
- data classification;
- audit/evidence exposure;
- backup sensitivity.

This does not import the future full Provider, Presence or physical-device security planes.

### D7 — Simple proof surface is not Presence Fabric

M0 may expose CLI/simple interaction, but the interface cannot own Aurora identity/state and cannot force a UI framework or multi-Presence design.

## 7. Cross-capability boundary decisions

### Memory / M1

`CAP-SOVEREIGN-CORE` owns operational-state durability and the minimum current snapshots required by M0. It does **not** own conversational/project memory lifecycle, retrieval, consolidation, Context Builder, forgetting or memory evaluation. M1 will consume sovereign state; M0 must not depend on M1 to exist.

### Registry/AHDK / M2

M0 does not depend on Capability Registry, Provider trust lifecycle, AHDK, Conformance Kit or a reference Harness. ADR-0002 remains accepted but is deferred as behavior for M0.

### Delegation/Effects / M3+

M0 persists a minimal authority snapshot; it does not implement the full Delegation/Effect/PDP/Gateway/Credential-Broker system. Later capabilities may own those behaviors. M0's persisted projection must be designed so it cannot become a competing authority owner.

### Presence / M8

M0 uses only a proof interface. Presence Fabric, device trust, handoff, sensors and privacy modes are later.

### Laboratory / M9–M10

All physical-device/laboratory requirements remain deferred. No physical capability is needed to close M0.

### Self-improvement / M11

No learning/candidate/promotion subsystem is part of M0.

## 8. Accepted ADR applicability

### ADR-0001 — applies as a boundary constraint

ADR-0001 applies to M0 only insofar as M0 domain/state semantics must remain language/framework/transport independent. It does not require a network protocol, provider binding or AHDK.

### ADR-0002 — deferred as behavior

ADR-0002 remains an accepted first-party Harness policy. Since M0 explicitly does not depend on Registry/AHDK/Harness integration, its implementation obligations are deferred. Its negative lesson—that an SDK is not the specification/security boundary—remains compatible with M0's framework-neutral design.

## 9. Open research and decisions

No requirement is classified `CONFLICT_REQUIRES_DECISION`.

The following are **open mechanisms**, not R1 applicability conflicts:

- Core language/runtime;
- initial process/deployment topology;
- operational state storage;
- state-versus-event implementation;
- schema representation;
- event/audit mechanism;
- telemetry backend;
- backup/restore topology/mechanism;
- migration mechanism;
- any durable execution engine only if R2/R3 prove that M0 needs it;
- exact architecture-spike scopes/IDs/candidates.

Their expected decision path remains R4, after R2 requirements and R3 Capability readiness, with separate Architecture Spike authorization where required.

## 10. Adversarial exclusion check

The highest-risk false exclusions were checked explicitly:

- identity continuity — active;
- Project/current state — active;
- authority scope/expiry/revocation on restore — active;
- transcript/model/Harness non-authority — active;
- sovereign local ownership — active;
- export/restore/migration — active or partial;
- restore security — active;
- event/audit versus current state — active;
- evidence/telemetry integrity — active;
- failure/recovery — active;
- threat model — active downstream dependency;
- framework/database/topology neutrality — active boundary;
- documentation/authorization governance — active.

The highest-risk false inclusions were also checked:

- M1 conversational memory — deferred;
- M2 Registry/AHDK — deferred;
- cross-Harness Delegations — deferred;
- cloud/multi-Presence — deferred;
- physical devices — deferred;
- adaptive campaigns/effect plane — deferred;
- self-improvement — deferred.

## 11. R1 gate result

Gate conditions:

- every known cross-cutting source considered — **PASS**;
- all 294 accepted constitutional requirements classified — **PASS**;
- every `NOT_APPLICABLE` classification has rationale — **PASS**;
- no unjustified exclusion identified — **PASS**;
- high-risk dependencies identified and assigned — **PASS**;
- conflicts/open research separated from applicability — **PASS**;
- no stack, mechanism or Architecture Spike execution selected — **PASS**.

Verdict:

```text
R1 PASS
```

## 12. Stop boundary

R1 PASS does **not** authorize R2.

Exact next permitted sequence:

```text
record R1 PASS
→ stop at R1
→ await explicit operator authorization for M0 ACRM R2 — Requirements
```

If R2 is later authorized, it must consume the 127 active source rows (`78 APPLIES + 49 PARTIALLY_APPLIES`), derive atomic/verifiable `CAP-SOVEREIGN-CORE` requirements, preserve source traceability, and explicitly keep the deferred remainder of every partial classification out of the M0 implementation commitment.
