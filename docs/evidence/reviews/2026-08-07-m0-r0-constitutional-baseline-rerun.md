---
id: REVIEW-AURORA-M0-R0-CONSTITUTIONAL-BASELINE-RERUN-2026-08-07
title: M0 ACRM R0 Constitutional Baseline Re-run
document_type: adversarial_review
form: explanation
authority: evidence
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - M0 ACRM R0 re-run observations and verdict against the fixed canonical target
related:
  - DOC-AURORA-STATUS
  - DOC-AURORA-M0-OPERATOR-SELECTION
  - REVIEW-AURORA-M0-R0-CONSTITUTIONAL-BASELINE-2026-08-06
  - DOC-AURORA-M0-R0-REMEDIATION-OPERATOR-ACCEPTANCE
  - DOC-AURORA-M0-R0-REMEDIATION-MERGE-CLOSEOUT
  - DOC-AURORA-M0-R0-RERUN-TARGET-FINDING
target_revision: 6054f84d007347c0aa9eef9e71317134b1047d3c
reviewed_at: 2026-08-07
last_reviewed: 2026-08-07
---

# M0 ACRM R0 Constitutional Baseline Re-run

## 1. Executive verdict

```text
R0 PASS
```

Fixed target:

```text
6054f84d007347c0aa9eef9e71317134b1047d3c
```

The accepted constitutional intent required to begin M0 applicability analysis is coherent, discoverable, sufficiently owned and explicitly authorized. The initial gate-failing documentation divergences were repaired, explicitly accepted by the operator, canonically integrated and revalidated before this re-run.

This verdict does **not** authorize R1 or any technical/implementation work.

## 2. Sources read

Mandatory/current entrypoints:

- `AGENTS.md`;
- `docs/tracking/STATUS.md`;
- `docs/DOCUMENTATION-MAP.md`;
- `docs/product/README.md`;
- `docs/roadmap.md`;
- `docs/product/CAPABILITY-REALIZATION-METHOD.md`.

Constitutional owners material to M0:

- Blueprint 01 — Product Vision;
- Blueprint 03 — Domain and World Model;
- Blueprint 04 — Cognitive Lifecycle and Journeys;
- Blueprint 06 — Memory, Knowledge and Context, specifically the separation of operational state from memory;
- Blueprint 07 — Harness boundary, as a negative sovereignty constraint;
- Blueprint 10 — Autonomy, Authority and Safety;
- Blueprint 11 — Security, Privacy and Sovereignty;
- Blueprint 12 — System Architecture;
- Blueprint 13 — Reliability, Observability and Evaluation;
- Blueprint 14 — Capability Roadmap / accepted M0 definition;
- Blueprint 15 — Documentation and Research Governance.

Related accepted/traceability sources:

- `docs/product/REQUIREMENTS-TRACEABILITY.md`;
- `docs/adr/README.md`;
- ADR-0001;
- ADR-0002;
- `docs/tracking/DECISIONS.md`;
- `docs/tracking/DOCUMENTATION-COVERAGE.md`;
- `docs/research/RESEARCH-MAP.md`;
- M0 operator-selection evidence;
- M0 R0 remediation operator-acceptance and merge-closeout evidence.

## 3. Constitutional intent map

### Sovereign identity

Aurora identity is stable across model, process, session, interface and provider replacement. Process/session/transcript are explicitly not Aurora identity.

### Project continuity

`Project` is the main durable work-continuity boundary and includes current state and next actions. A snapshot is a projection rather than a competing owner.

### Operational state

Operational state is structured and authoritative, distinct from conversation history, model memory, Harness local state, workflow-engine history, Git documents and logs.

### Authority

Access/capability is not authority. Current authority must be independently resolvable, and an authority snapshot is a projection over governing authority rather than permission invented from narrative.

### Restart/recovery

State is durable while processes are disposable. A fresh process must preserve/reconstruct stable IDs, authority, state and the next safe action without transcript dependency.

### Sovereignty and restore

Canonical state remains under Leandro-controlled governance. Export/backup/restore/migration are part of sovereignty. Restore must preserve integrity and must not silently reactivate expired authority or compromised trust.

### Reliability/evidence

Restart/restore success is not a self-declared claim. M0 closes only from criterion-linked evidence and its end-to-end Golden Proof.

### Harness/provider boundary

Harnesses and models remain external/specialized mechanisms. They do not own Aurora global identity, project state or authority. M0 explicitly does not depend on M2 Registry/AHDK or MNFS.

### Technical neutrality

Language, storage, runtime, topology, schema, transport, telemetry backend, audit/event mechanism and backup mechanism remain deliberately open.

## 4. Accepted M0 anatomy

### Outcome

A minimal Aurora Core persists identity, project state, authority snapshot and one interaction lifecycle across restart without depending on an external model or Harness as authority.

### Operator-visible value

Leandro can initialize Aurora, record project state, restart it, and recover the exact current state and permitted next action.

### Risk retired

```text
Aurora is merely a running session; restart destroys identity and state.
```

### Capabilities involved

- sovereign identity;
- project registry;
- operational state;
- authority snapshot;
- event/audit minimum;
- CLI or simple interface.

### Architecture-spike intent

The roadmap now names only the required uncertainty classes: local state/recovery and language/runtime fit. It explicitly does not select spike IDs, technologies, procedures or winners and preserves separate authorization for execution.

### Golden Proof

```text
initialize Aurora instance
→ create project
→ record accepted state and next action
→ terminate all Aurora processes
→ start fresh process
→ recover same identities/state
→ reject an invalid transition
→ export and restore state
```

### Evidence requirements

- state hashes/IDs;
- restart receipt;
- invalid-transition test;
- backup/restore result;
- no transcript dependency.

### Exit criteria

M0 now explicitly requires the complete end-to-end Golden Proof, persistence of relevant identities/state/authority/next action, rejection of the invalid transition, successful export/restore, proof of sovereignty from transcript/model/Harness state, explicit limitations/residual risks and R8 closeout rather than component-level inference.

### Telemetry baseline

Only signals necessary to explain the walking skeleton are constitutionally required. No telemetry backend, event transport or schema technology is selected.

### Dependencies

M0 depends on its own entry prerequisites and accepted constitutional owners. It explicitly does not depend on M1 conversational memory, M2 Registry/AHDK, MNFS, cloud or physical-device capability.

### Non-goals

- conversational memory;
- model routing;
- Harness Registry;
- voice;
- multi-device;
- cloud;
- physical devices.

### Replan triggers

- storage cannot preserve the required state simply;
- domain model is too broad for the slice;
- operational burden exceeds the single-user baseline.

### Promotion/authority boundary

R0–R6 may refine applicability, verifiable requirements, reusable Capability design, technical decisions and exact implementation commitment without silently changing outcome/risk/Golden Proof/non-goals. Every next gate and every spike/implementation action remains separately authorized.

## 5. M0 entry criteria analysis

| Entry criterion | R0 assessment | Owning later readiness work |
|---|---|---|
| A0 accepted | SATISFIED | none |
| Core boundaries approved | SATISFIED at constitutional/logical level | exact Capability mechanisms refine at R3 |
| minimal domain/entity spec | NOT YET an executable M0 spec; constitutional domain is sufficient to derive it | R1 applicability → R2 requirements → R3 Capability Spec |
| storage and language spikes complete enough | NOT SATISFIED and intentionally not executed | R4, after applicable requirements/Capability readiness and separate spike authorization |
| backup/restore and migration strategy for slice | behavior is constitutionally owned; exact strategy remains open | R2/R3 behavior/specification and R4 technical decision |

Unsatisfied later-stage entry prerequisites do not fail R0. R0 asks whether accepted intent is ready for applicability analysis, not whether M0 is ready for implementation.

## 6. Initial R0 findings resolution

### R0-F01 — milestone anatomy divergence

**RESOLVED.** Blueprint 14 now distinguishes directional future milestones from a milestone promoted into the executable horizon and provides the complete required M0 anatomy.

### R0-F02 — ADR status divergence

**RESOLVED.** ADR index and accepted ADR files agree that ADR-0001 and ADR-0002 are accepted. The index also makes clear that the post-A0 ADR readiness rule does not retroactively invalidate the A0 decisions.

### R0-F03 — mutable-state duplication/drift

**RESOLVED.** `STATUS.md` is the single owner of mutable milestone/gate/authorization/next-action coordination. Stable bootstrap/index/constitutional documents point to it instead of duplicating current state.

### R0-F04 — re-run target continuity

A tracking-only issue discovered during closeout hardcoded a pre-closeout merge SHA as the future review target. **RESOLVED** before this review by changing the rule to resolve and pin current canonical `main` HEAD at review start. No constitutional meaning changed.

## 7. Remaining open items by correct gate

### R1 — Applicability

When separately authorized, classify which of the 294 accepted constitutional requirements apply, partially apply, are deferred by roadmap or are not applicable to `CAP-SOVEREIGN-CORE`, with rationale and dependency ownership.

### R2 — Requirements

Examples that need atomic/verifiable M0 definition rather than constitutional invention:

- exact minimum semantics of sovereign identity;
- minimum Project state;
- authority-snapshot correctness/freshness behavior;
- `permitted next action` semantics;
- the M0 interaction-lifecycle obligations;
- valid/invalid transition behavior;
- restart/recovery invariants;
- export/restore security behavior;
- minimum event/audit evidence obligations.

### R3 — Capability Readiness

- `CAP-SOVEREIGN-CORE` reusable Capability/System Spec;
- minimal domain model for the slice;
- lifecycle/state model;
- security/threat model;
- failure/recovery behavior;
- observability/evidence model;
- Capability test plan and requirements coverage.

### R4 — Architecture/Decision Readiness

Still open and not selected by R0:

- Core language/runtime;
- process/deployment topology;
- operational state store;
- state-versus-event implementation;
- schema representation;
- audit/event mechanism;
- backup/restore mechanism/topology;
- migration mechanism;
- any durable engine only if M0 requirements justify it;
- exact Architecture Spike scopes/IDs/candidates and their separately authorized execution.

### R5 — Mission Contract

Only after R0–R4 and separate authorization: exact scoped M0 implementation commitment, criteria, baseline, versions, authority, environment, non-goals and evidence allocation.

### R6 — Microdesign / Implementation Plan

Only after an approved Contract: exact files/modules, algorithms, schemas, migrations, commands, tests, rollback and implementation sequence.

## 8. ADR applicability at R0

### ADR-0001

Applicable as a negative architecture/ownership guard: no external framework/protocol/runtime may become the owner of Aurora semantics. It does not require M0 to select or implement a protocol/binding.

### ADR-0002

Accepted policy remains valid but does not make AHDK part of M0. M0 explicitly excludes the M2 Registry/AHDK dependency and current STATUS continues to prohibit AHDK implementation.

## 9. Premature-scope risks

R0 must not be used to:

- turn M0 into M1 conversational memory;
- pull Capability Registry/AHDK/MNFS from M2/M6 into the sovereign Core;
- implement the full future Authority/Effect plane merely because M0 persists an authority snapshot;
- choose a durable workflow engine before an M0 requirement demonstrates need;
- turn `event/audit minimum` into a distributed event platform;
- turn telemetry baseline into an observability-stack selection;
- treat `CLI or simple interface` as UI/runtime selection;
- generalize schemas, protocols or multi-service topology before the walking skeleton requires them.

## 10. Non-blocking documentation note

The Documentation Map canonical-entrypoint table still describes `REQUIREMENTS-TRACEABILITY.md` using the historical phrase “proposed constitutional requirements and proof intent,” while the requirement document itself is now explicitly `accepted`. This is a minor index-description hygiene issue, not a competing lifecycle owner and not material to M0 constitutional intent/discoverability. The canonical requirement document and current authorization remain unambiguous. It should be cleaned in a future authorized documentation-maintenance change rather than smuggled into this R0 verdict.

## 11. R0 checks

| R0 check | Result |
|---|---|
| relevant Blueprint sections accepted/current | PASS |
| no unresolved material constitutional divergence | PASS |
| Product Milestone exists | PASS |
| Golden Proof exists | PASS |
| scope/non-goals clear | PASS |
| current STATUS authorizes R0 readiness work | PASS |
| fresh repository-only reader can discover owners/current gate | PASS |

## 12. Verdict and next action

```text
R0 PASS
```

Exact next permitted action:

```text
record R0 PASS and stop
→ await explicit operator authorization for M0 ACRM R1 — Applicability
```

R1 is **not** authorized by this verdict. No applicability package, Capability requirement derivation, Architecture Spike execution, stack choice, Mission Contract, Microdesign or implementation may begin by implication.
