---
id: DOC-AURORA-DECISIONS
title: Aurora Decision Index
document_type: decision_index
form: reference
authority: tracking
status: current
version: 0.9.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - current decision discovery and status index
related:
  - DESIGN-AURORA-SYSTEM-ARCHITECTURE-REBASELINE
  - DOC-AURORA-SYSTEM-ARCHITECTURE-REBASELINE-OPERATOR-DIRECTION
  - DESIGN-AURORA-SYSTEM-ARCHITECTURE-DECISION-LANDSCAPE
  - DESIGN-AURORA-TECHNICAL-ARCHITECTURE-BASELINE-MAP
  - DOC-AURORA-TECHNICAL-ARCHITECTURE-MAP-OPERATOR-ACCEPTANCE
  - PLAN-AURORA-TECHNICAL-ARCHITECTURE-BASELINE
  - DESIGN-AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY
  - DOC-AURORA-TA-01-02-OPERATOR-ACCEPTANCE
last_reviewed: 2026-08-13
---

# Aurora Decision Index

## 1. Authority notice

This file indexes decisions and open choices. It does not replace the Product Blueprint, accepted ADRs, accepted System Architecture Rebaseline, accepted Technical Architecture Baseline Map or operator decision evidence.

```text
Index entry
→ points to canonical owner/evidence
→ reports lifecycle/status
→ never redefines the decision
```

A0, the current M0 ADR set and the R4-aligned CAP-SOVEREIGN-CORE A2 package are operator-accepted; `MIS-M0-SOVEREIGN-CORE-001` v0.1.0 is the approved first M0 Mission Contract. M0 R7 produced a frozen non-canonical candidate, but no independent R7 Verdict or R8 closeout exists. Current implementation remains paused while the accepted Technical Architecture Baseline begins with TA-01/TA-02.

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
| D-059 | CAP-SOVEREIGN-CORE Requirements v0.1.1, Spec v0.2.0, Threat Model v0.2.0 and Test Plan v0.2.0 are the accepted R4-aligned M0 A2 package | CAP-SOVEREIGN-CORE A2 documents + R5 operator acceptance | accepted |
| D-060 | MIS-M0-SOVEREIGN-CORE-001 v0.1.0 is the approved first scoped M0 Mission Contract | MIS-M0-SOVEREIGN-CORE-001 + R5 operator acceptance | approved |

## 7. Methodology, documentation and architecture-program decisions

| ID | Decision/direction | Canonical owner | Current state |
|---|---|---|---|
| D-045 | documentation is product memory/governance and must precede implementation | Blueprint 15, Documentation Map | accepted |
| D-046 | one durable concept has one canonical owner | Blueprint 15, Documentation Map | accepted |
| D-047 | research informs, ADR decides, Spec defines, Contract commits, Plan implements and Evidence proves | ACRM | accepted |
| D-048 | material capabilities follow readiness gates R0–R8 | ACRM | accepted |
| D-049 | Product Milestones close through end-to-end Golden Proofs, not task completion | Blueprint 14, ACRM | accepted |
| D-050 | aggregate Product Blueprint and roadmap are generated from modular canonical sources | Blueprint 15, Product Index | accepted |
| D-051 | `M0 — Sovereign Core Walking Skeleton` is the first Product Milestone after A0 | Blueprint 14 + M0 operator selection evidence | accepted |
| D-061 | before further multi-subsystem implementation, Aurora performs a System Architecture Rebaseline inside ACRM rather than creating a parallel lifecycle | accepted Rebaseline design + operator direction + ACRM | accepted |
| D-062 | global architecture questions use `DECIDE`, `RESEARCH`, `SPIKE` or `DEFER`, with consumer, evidence, owner and reconsideration trigger | accepted Rebaseline design + ACRM | accepted |
| D-063 | the software-development Harness may build, test, review and package evidence for Aurora but is not a sovereign Aurora runtime dependency | accepted Rebaseline design | accepted |
| D-064 | the M0 R7 implementation candidate remains frozen, preserved and non-canonical; code/CI existence is not acceptance | operator direction + STATUS | accepted current coordination |
| D-065 | current program priority is the cross-system Technical Architecture Baseline, not additional broad product-definition dialogue or Presence micro-policy decomposition | Technical Architecture Map + operator acceptance | accepted |
| D-066 | a technical question is current only when it changes ownership, structural/runtime/contract/security/data boundaries or the next implementation decision; otherwise it is deferred | Technical Architecture Map | accepted |
| D-067 | technical architecture proceeds in dependency order TA-01 modules, TA-02 runtimes, TA-03 repositories, TA-04 contracts, TA-05 data, TA-06 identity/security, TA-07 cognition/Harnesses and TA-08 operations | Technical Architecture Map | accepted |
| D-068 | accepted Stage A Presence/activation/locked-workstation rules remain downstream constraints, while further session-policy detail is deferred until a consuming Capability | Stage A design + Technical Architecture Map | accepted |
| D-069 | TA-01 + TA-02 is the accepted first Technical Architecture tranche, fixing module ownership and Stage A/B runtime-topology direction before repository or stack finalization | TA-01/TA-02 design + operator acceptance | accepted |
| D-070 | canonical ownership is G01 Contract Model Governance plus C01–C12 domain owners, including C03 Intent ownership and C12 Audit/Exact History | TA-01/TA-02 design v0.5.0 + operator acceptance | accepted |
| D-071 | Stage A uses Approach C: one small persistent Evolutionary Sovereign Host plus a separate on-demand provider-runtime seam at the first consumer; other process splits require evidence | TA-01/TA-02 design v0.5.0 + operator acceptance | accepted |
| D-072 | A05 owns Aurora-side runtime lifecycle policy and B01 owns transport-neutral provider identity/lifecycle/idempotency/cancellation/reconciliation semantics before TA-04 selects a binding | TA-01/TA-02 design v0.5.0 + operator acceptance | accepted |

## 8. Deliberately open decisions

These are not yet decisions and must not be inferred from examples, research candidates, the selected M0 milestone or the frozen R7 candidate. The complete dependency map is `DESIGN-AURORA-TECHNICAL-ARCHITECTURE-BASELINE-MAP` and `DESIGN-AURORA-SYSTEM-ARCHITECTURE-DECISION-LANDSCAPE`.

| Open ID | Decision required | Expected owner/path |
|---|---|---|
| O-002 | first AHDK language and source-code generation stack | CAP-AHDK research/spike + ADR |
| O-003 | schema representation per boundary | Contract Model Spec + consuming R4 evidence |
| O-004 | local RPC binding | TA-04 interoperability research/spike + ADR |
| O-005 | exact MCP/A2A/ACP adoption/mapping | consuming capability spike + ADR/profile |
| O-006 | durable execution engine | M4 Capability Spec + comparative spike + ADR |
| O-007 | policy decision implementation | TA-06 Effect/Authority Spec + spike + ADR |
| O-008 | workload/device identity implementation | TA-06 security Spec + actor-specific spike |
| O-010 | Artifact/Evidence Store | TA-05/TA-07 evidence capability research/spike |
| O-011 | event transport and telemetry backend | TA-04/TA-08 research/spike |
| O-012 | memory storage/retrieval/consolidation mix | TA-05/TA-07 + CAP-MEMORY-CONTEXT eval spikes |
| O-013 | first reference Harness runtime | TA-07/M2 after current architecture review |
| O-014 | first real engineering Harness | roadmap readiness; MNFS remains one candidate |
| O-016 | Stage A/B actor authentication mechanisms by class | TA-06 research + capability-specific ADRs |
| O-017 | API profiles, error taxonomy, idempotency and streaming conventions per boundary | TA-04 + first consuming capability R4 |
| O-018 | configuration, environment and secret-reference precedence model | TA-08 Standard/ADR |
| O-019 | model/inference and Brain/Core ownership boundary | TA-07 + M1 research/Capability Spec |
| O-020 | physical storage beyond M0 for memory, artifacts, telemetry and derived indexes | TA-05 consumer-specific research/spike |
| O-023 | monorepo, polyrepo or staged source strategy | TA-03 after TA-01/TA-02 acceptance |
| O-024 | exact Go ↔ TypeScript/Mastra process and contract boundary | TA-02/TA-04/TA-07, first consumer evidence |

## 9. Status-change rule

An index entry changes only when its canonical owner/evidence changes lifecycle/status.

```text
conversation approval
→ permits faithful documentation or records an explicit operator decision when directly asked

accepted Product Blueprint/design
→ governs only its stated authority scope

accepted ADR
→ promotes a specific technical decision

accepted Technical Architecture Map
→ orders questions and work; does not select mechanisms by itself

approved Contract
→ creates scoped implementation commitment

implementation branch/green CI
→ produces Claim/Evidence only; never acceptance by itself
```

No entry is promoted merely because code exists, CI is green, a framework is popular or a PR is merged.
