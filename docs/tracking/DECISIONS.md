---
id: DOC-AURORA-DECISIONS
title: Aurora Decision Index
document_type: decision_index
form: reference
authority: tracking
status: current
version: 0.5.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - current decision discovery and status index
last_reviewed: 2026-08-06
---

# Aurora Decision Index

## 1. Authority notice

This file indexes decisions and open choices. It does not replace the Product Blueprint, accepted ADRs or operator decision evidence.

```text
Index entry
→ points to canonical owner/evidence
→ reports lifecycle/status
→ never redefines the decision
```

A0 and ADR-0001/0002 were explicitly accepted by the operator and merged to `main`. The first Product Milestone was subsequently selected as M0. Technical mechanisms and later ACRM gates remain open unless separately accepted/authorized.

## 2. Constitutional direction

| ID | Decision/direction | Canonical owner | Current state |
|---|---|---|---|
| D-001 | preserve complete long-term vision while limiting technical commitment to the next executable horizon | Blueprint 01, 14 | accepted |
| D-002 | Aurora is Leandro-first and single-user in the current horizon | Blueprint 01 | accepted |
| D-003 | Aurora is a broad personal intelligence with engineering as the first deep domain | Blueprint 01, 03 | accepted |
| D-004 | Aurora is a trusted intellectual copilot, not a passive or blindly obedient assistant | Blueprint 02 | accepted |
| D-005 | loyalty is to Leandro's objective and values, not every momentary premise | Blueprint 02 | accepted |
| D-006 | Leandro retains final authority over goals, values, material decisions and grants | Blueprint 02, 10 | accepted |
| D-007 | Aurora has a stable, expressive and transparent AI identity | Blueprint 02, 08 | accepted |
| D-008 | personality combines J.A.R.V.I.S.-like precision/humor with E.V.I.E.-like proximity/energy without copying either | Blueprint 02, History | accepted |
| D-009 | personality must have presence without performative humanity or emotional manipulation | Blueprint 02, 11 | accepted |
| D-010 | default proactivity is contextual and controlled by an attention budget | Blueprint 02, 13 | accepted |
| D-011 | one Aurora may manifest through multiple presences/devices | Blueprint 08, 12 | accepted |
| D-012 | handoff preserves intention but adapts context to device, environment, identity and sensitivity | Blueprint 08, 11 | accepted |
| D-013 | availability does not imply continuous recording or surveillance | Blueprint 08, 11 | accepted |

## 3. Memory and learning direction

| ID | Decision/direction | Canonical owner | Current state |
|---|---|---|---|
| D-014 | memory is a governed multi-scope, multi-temporal subsystem rather than transcript storage | Blueprint 06 | accepted |
| D-015 | memory, history, knowledge, operational state, active context and source of truth are distinct | Blueprint 03, 06 | accepted |
| D-016 | memory promotion is proportional to risk, scope, sensitivity and authority | Blueprint 06 | accepted |
| D-017 | conflicts use supersession and provenance rather than silent overwrite | Blueprint 06 | accepted |
| D-018 | memory guides reasoning while canonical authority, evidence and live state govern action | Blueprint 06 | accepted |
| D-019 | Aurora may detect self-improvement opportunities continuously but experiment only within authorization | Blueprint 13 | accepted |
| D-020 | self-improvement must investigate causal classes, not patch only the visible symptom | Blueprint 13 | accepted |
| D-021 | material self-improvement promotion requires independent evaluation/review, canary and rollback | Blueprint 10, 13 | accepted |
| D-022 | constitutional identity, Leandro authority and safety boundaries cannot be autonomously promoted | Blueprint 10, 13 | accepted |

## 4. Autonomy, sovereignty and physical direction

| ID | Decision/direction | Canonical owner | Current state |
|---|---|---|---|
| D-023 | authority is progressive, explicit, scoped, expiring and revocable | Blueprint 10, 11 | accepted |
| D-024 | Aurora may receive autonomy by action, workflow or adaptive campaign | Blueprint 04, 10 | accepted |
| D-025 | Aurora is autonomous inside an approved envelope and conservative at its boundary | Blueprint 10 | accepted |
| D-026 | campaign baselines, evaluation criteria, budgets and protected areas cannot be silently changed by the executor | Blueprint 10, 13 | accepted |
| D-027 | physical and critical digital guardrails require deterministic enforcement independent of model judgment | Blueprint 09, 10 | accepted |
| D-028 | emergency authority is narrow, preauthorized, tested and observable | Blueprint 10 | accepted |
| D-029 | Aurora is local-first and cloud-assisted; intelligence may be distributed while sovereignty remains controlled by Leandro | Blueprint 11, 12 | accepted |
| D-030 | external providers receive minimized authorized context rather than unrestricted memory | Blueprint 06, 11 | accepted |

## 5. Capability and Harness architecture

| ID | Decision/direction | Canonical owner | Current state |
|---|---|---|---|
| D-031 | Aurora is the global cognitive/operational control plane | Blueprint 01, 07, 12 | accepted |
| D-032 | Harnesses are specialized capability providers with internal methodology and execution autonomy | Blueprint 05, 07 | accepted |
| D-033 | MNFS is a future software-engineering provider and does not define Aurora architecture | Blueprint 07, 14 | accepted |
| D-034 | Aurora owns why, what, global limits and composition; Harness owns local how | Blueprint 07 | accepted |
| D-035 | cross-Harness collaboration is hierarchical and mediated by Aurora child Delegations | Blueprint 07 | accepted |
| D-036 | control plane is centrally governed; authorized high-volume/low-latency data plane may be direct | Blueprint 07, 12 | accepted |
| D-037 | discovery, compatibility, trust, authority and execution are separate states | Blueprint 05 | accepted |
| D-038 | trust is multidimensional and bound to exact provider version/build/environment/scope | Blueprint 05, 11 | accepted |
| D-039 | manifests declare behavior but do not prove it or grant authority | Blueprint 05 | accepted |
| D-040 | providers are selected by fit, evidence, sensitivity, authority, environment, cost, latency, recovery and availability | Blueprint 05 | accepted |

## 6. ADR decisions

| ID | Decision | Canonical owner | Current state |
|---|---|---|---|
| D-041 | Aurora owns cross-Harness Contract Model semantics; protocols/runtimes remain replaceable bindings | ADR-0001 | accepted |
| D-042 | first-party Harnesses use AHDK by policy unless waived; all providers pass universal/profile conformance | ADR-0002 | accepted |
| D-043 | SDK is neither specification nor security boundary | Blueprint 05, 10, ADR-0002 | accepted |
| D-044 | durable execution, policy decision, effect enforcement, containment, evidence and observability are separate layers | Blueprint 10–13 | accepted |
| D-052 | Go is the accepted initial runtime for the Aurora Sovereign Core; other runtimes remain free behind contracts | ADR-0003 | accepted |
| D-053 | M0 uses one local modular Sovereign Core with explicit current state/revisions and no full event sourcing or durable workflow engine | ADR-0004 | accepted |
| D-054 | M0 portable logical state uses JSON Schema/JSON/JCS boundaries with protected export and application-owned migration semantics | ADR-0005 | accepted |
| D-055 | M0 observability uses OTel traces/metrics plus structured Go logging while exporter/backend remains optional | ADR-0006 | accepted |
| D-056 | Mastra is the accepted preferred-first substrate to evaluate for first-party agentic Harnesses while sovereign truth/authority remain Aurora-owned | ADR-0009 | accepted |
| D-057 | SQLite + `database/sql` + `modernc.org/sqlite` is the accepted M0 operational-state baseline | ADR-0007 | accepted |
| D-058 | M0 owner trust uses a random wrapped ORK + authenticated external generation/time high-water with fail-closed restore/revalidation semantics | ADR-0008 | accepted |

## 7. Methodology, documentation and milestone decisions

| ID | Decision/direction | Canonical owner | Current state |
|---|---|---|---|
| D-045 | documentation is product memory/governance and must precede implementation | Blueprint 15, Documentation Map | accepted |
| D-046 | one durable concept has one canonical owner | Blueprint 15, Documentation Map | accepted |
| D-047 | research informs, ADR decides, Spec defines, Contract commits, Plan implements and Evidence proves | ACRM | accepted |
| D-048 | material capabilities follow readiness gates R0–R8 | ACRM | accepted |
| D-049 | Product Milestones close through end-to-end Golden Proofs, not task completion | Blueprint 14, ACRM | accepted |
| D-050 | aggregate Product Blueprint and roadmap are generated from modular canonical sources | Blueprint 15, Product Index | accepted |
| D-051 | `M0 — Sovereign Core Walking Skeleton` is the first Product Milestone after A0; R0 begins in a fresh session before any technical commitment | Blueprint 14 + M0 operator selection evidence | accepted |

## 8. Deliberately open decisions

These are not yet decisions and must not be inferred from examples, research candidates or the selected M0 milestone:

| Open ID | Decision required | Expected owner/path |
|---|---|---|
| O-002 | first AHDK language and source-code generation stack | CAP-AHDK research/spike + ADR |
| O-003 | schema representation per boundary | Contract Model Spec + SPK-001 |
| O-004 | local RPC binding | interoperability research/spike + ADR |
| O-005 | exact MCP/A2A adoption/mapping | SPK-002/SPK-003 + ADR/profile |
| O-006 | durable execution engine | SPK-004 + ADR |
| O-007 | policy decision implementation | SPK-005 + ADR |
| O-008 | workload/device identity implementation | security Capability Spec + spike |
| O-010 | Artifact/Evidence Store | evidence capability research/spike |
| O-011 | event transport and telemetry backend | observability capability research/spike |
| O-012 | memory storage/retrieval/consolidation mix | CAP-MEMORY-CONTEXT + eval spikes |
| O-013 | first reference Harness runtime | M2/M3 contract after M0/M1 readiness |
| O-014 | first real engineering Harness | roadmap readiness; MNFS is one candidate |
| O-015 | exact first Mission Contract for selected M0 | `MIS-M0-SOVEREIGN-CORE-001` v0.1.0 proposed in R5; operator approval pending |

`O-015` previously combined milestone selection and first Contract. The milestone portion is resolved by `D-051`. R5 now proposes `MIS-M0-SOVEREIGN-CORE-001` v0.1.0; the decision remains open until explicit operator approval.

## 9. Status-change rule

An index entry changes only when its canonical owner/evidence changes lifecycle/status.

```text
conversation approval
→ permits faithful documentation or records an explicit operator decision when directly asked

A0 operator acceptance
→ promotes the approved constitutional package

accepted ADR
→ promotes a specific technical decision

selected Product Milestone
→ identifies the next readiness subject but does not authorize implementation

approved Contract
→ creates scoped implementation commitment
```

No entry is promoted merely because code exists, CI is green or a PR is merged.
