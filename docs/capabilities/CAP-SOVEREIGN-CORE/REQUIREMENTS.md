---
id: DOC-AURORA-CAP-SOVEREIGN-CORE-REQUIREMENTS
title: CAP-SOVEREIGN-CORE Atomic Requirements
document_type: capability_requirements
form: reference
authority: specification
status: accepted
accepted_at: 2026-08-09
acceptance_evidence: DOC-AURORA-M0-R5-A2-CONTRACT-OPERATOR-ACCEPTANCE
accepted_from_blob: de234e4a57c04d1d0b68cd017597e06a618fd68b
version: 0.1.1
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - M0 atomic requirements for CAP-SOVEREIGN-CORE
related:
  - DOC-AURORA-STATUS
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DOC-AURORA-REQUIREMENTS-TRACEABILITY
  - DOC-AURORA-CAP-SOVEREIGN-CORE-APPLICABILITY
  - DOC-AURORA-M0-R2-OPERATOR-AUTHORIZATION
  - DOC-AURORA-CAP-SOVEREIGN-CORE-R5-COVERAGE
  - DOC-AURORA-MIS-M0-SOVEREIGN-CORE-001
  - DOC-AURORA-M0-R5-A2-CONTRACT-OPERATOR-ACCEPTANCE
source_revision: 495b712142d7c3d722da2298f7a0b060707f9f5e
review_triggers:
  - applicable constitutional requirement changes
  - M0 roadmap meaning changes
  - accepted ADR changes affecting M0
  - R2 finding or R3 contradiction
last_reviewed: 2026-08-07
---

# CAP-SOVEREIGN-CORE — Atomic Requirements

## 1. R2 scope and authority

This is the R2 requirements package for `M0 — Sovereign Core Walking Skeleton` and `CAP-SOVEREIGN-CORE`.

Fixed source baseline:

```text
495b712142d7c3d722da2298f7a0b060707f9f5e
```

R1 selected 127 active constitutional source rows (`78 APPLIES + 49 PARTIALLY_APPLIES`). R2 derives atomic, verifiable, implementation-neutral Capability requirements from those sources.

This document does **not** define the complete Capability/System design or threat model (R3), select architecture/stack/spike winners (R4), create a Mission Contract (R5), create Microdesign/implementation allocation (R6), or authorize implementation (R7).

All 122 requirements were proposed at R2. This v0.1.1 preserves their normative statements unchanged while making the package review-ready for R5 operator acceptance. R3 owns mechanism/test allocation and `R5-COVERAGE.md` owns requirement-to-Mission allocation; this document was explicitly accepted by the operator through `DOC-AURORA-M0-R5-A2-CONTRACT-OPERATOR-ACCEPTANCE`.

## 2. Interpretation

- **M0 canonical state** — minimum Aurora/Project operational state, current authority state/snapshot inputs, and derived next-safe-action state required by the M0 Golden Proof.
- **fresh process** — process started after all prior Aurora processes are terminated, with no prior-process in-memory state available.
- **current accepted state** — Project state revision the canonical owner currently recognizes as governing.
- **next safe action** — next operator-visible action consistent with current accepted Project state and currently valid authority; it is not itself an Authority Grant.
- **material M0 operation** — initialize/create, state-transition attempt, authority change/validation, restart/recovery, export, restore or migration when it can affect or prove canonical M0 state.
- **successful operation** — operation whose required resulting state and evidence have been observed and verified, not merely dispatched.

Risk levels: `critical` can corrupt identity, authority, canonical state, sovereignty or the Golden Proof; `high` can materially impair recovery, security, evidence, migration or readiness; `medium` can create avoidable coupling, ambiguity or operational/documentation debt.

Verification methods use the accepted ACRM vocabulary and are directions, not the future R3 test plan.

## 3. Summary

| Category | Requirements |
|---|---:|
| Identity and scope | 9 |
| Canonical ownership and state | 11 |
| State transition lifecycle | 11 |
| Authority | 14 |
| Durability and recovery | 10 |
| Export, restore and migration | 11 |
| Security and sovereignty | 10 |
| Event, audit, evidence and telemetry | 12 |
| Reliability and verification | 7 |
| Open decision guards | 12 |
| Documentation and readiness | 15 |
| **TOTAL** | **122** |

## 4. Identity and scope

| ID | Requirement | Risk | Verification | Constitutional sources |
|---|---|---|---|---|
| `CAP-SOVEREIGN-CORE-REQ-001` | Aurora MUST preserve one stable system identity across process termination and restart, independent of any LLM, provider, session, interface, runtime, storage product or device. | critical | CONTRACT_TEST, INTEGRATION, USER_JOURNEY | `AUR-REQ-VIS-001`, `AUR-REQ-VIS-007`, `AUR-REQ-REL-005`, `AUR-REQ-ARC-017`, `AUR-REQ-PRS-001` |
| `CAP-SOVEREIGN-CORE-REQ-002` | Every M0 Project MUST have a stable durable identity that survives restart and restore and MUST NOT rely on a mutable display name as its identity. | critical | SCHEMA_VALIDATION, INTEGRATION, USER_JOURNEY | `AUR-REQ-DOM-001`, `AUR-REQ-DOM-004`, `AUR-REQ-COG-017` |
| `CAP-SOVEREIGN-CORE-REQ-003` | A successful restart or restore MUST recover the same Aurora and Project identities represented by the recovered state; replacement identities MUST require an explicit initialize/create operation rather than happen implicitly. | critical | INTEGRATION, FAULT_INJECTION, USER_JOURNEY | `AUR-REQ-VIS-001`, `AUR-REQ-RDM-007`, `AUR-REQ-ARC-004`, `AUR-REQ-COG-017` |
| `CAP-SOVEREIGN-CORE-REQ-004` | CAP-SOVEREIGN-CORE MUST remain Leandro-first and single-user for M0 and MUST NOT introduce multi-tenant identity, isolation or authority semantics as a hidden current commitment. | medium | DOCUMENT_REVIEW, CONTRACT_TEST | `AUR-REQ-VIS-003`, `AUR-REQ-VIS-009`, `AUR-REQ-RDM-002`, `AUR-REQ-RDM-005` |
| `CAP-SOVEREIGN-CORE-REQ-005` | CAP-SOVEREIGN-CORE MUST have a stable capability identity and version whose meaning is independent of programming language, framework, protocol, database and binding. | high | DOCUMENT_REVIEW, CONTRACT_TEST | `AUR-REQ-CAP-001`, `AUR-REQ-DOM-006`, `AUR-REQ-ARC-007`, `AUR-REQ-ORCH-020` |
| `CAP-SOVEREIGN-CORE-REQ-006` | No model, Harness, provider, framework, interface, workflow runtime, database or device MAY be treated as Aurora or as the canonical owner of Aurora identity. | critical | STATIC_ANALYSIS, DOCUMENT_REVIEW, INTEGRATION | `AUR-REQ-VIS-007`, `AUR-REQ-VIS-011`, `AUR-REQ-DOM-003`, `AUR-REQ-DOM-013`, `AUR-REQ-ORCH-021`, `AUR-REQ-ARC-002` |
| `CAP-SOVEREIGN-CORE-REQ-007` | M0 continuity and recovery MUST NOT require M1 conversational memory, M2 Capability Registry/AHDK, MNFS integration, cloud deployment, multi-Presence operation or physical-device capability. | high | DOCUMENT_REVIEW, INTEGRATION | `AUR-REQ-VIS-008`, `AUR-REQ-VIS-009`, `AUR-REQ-ORCH-022`, `AUR-REQ-MEM-023`, `AUR-REQ-RDM-002`, `AUR-REQ-RDM-005` |
| `CAP-SOVEREIGN-CORE-REQ-008` | The M0 operator interface MUST act only as a proof/control surface; it MUST NOT own canonical identity, Project state or authority, and M0 MUST NOT assume a specific CLI, web, voice, glasses or other UI technology. | high | DOCUMENT_REVIEW, INTEGRATION | `AUR-REQ-PRS-001`, `AUR-REQ-PRS-015`, `AUR-REQ-ARC-016`, `AUR-REQ-VIS-007` |
| `CAP-SOVEREIGN-CORE-REQ-009` | R2 requirements MUST remain limited to the accepted M0 outcome, Golden Proof, dependencies and non-goals and MUST NOT silently promote deferred roadmap behavior into the current executable horizon. | high | DOCUMENT_REVIEW | `AUR-REQ-VIS-008`, `AUR-REQ-VIS-009`, `AUR-REQ-RDM-001`, `AUR-REQ-RDM-002`, `AUR-REQ-RDM-004`, `AUR-REQ-RDM-005` |

## 5. Canonical ownership and state

| ID | Requirement | Risk | Verification | Constitutional sources |
|---|---|---|---|---|
| `CAP-SOVEREIGN-CORE-REQ-010` | Aurora Core MUST be the canonical owner for the M0 subset of Aurora identity, Project operational state, authority state needed by the slice, and the accepted next-action state. | critical | DOCUMENT_REVIEW, CONTRACT_TEST, INTEGRATION | `AUR-REQ-VIS-002`, `AUR-REQ-ORCH-002`, `AUR-REQ-ARC-001`, `AUR-REQ-ARC-006` |
| `CAP-SOVEREIGN-CORE-REQ-011` | Canonical M0 operational state MUST remain distinct from conversation/transcript history, governed memory, model recall, Harness/provider-local state, workflow-engine history, Git documentation and logs/telemetry. | critical | DOCUMENT_REVIEW, INTEGRATION, FAULT_INJECTION | `AUR-REQ-MEM-001`, `AUR-REQ-MEM-003`, `AUR-REQ-MEM-004`, `AUR-REQ-MEM-006`, `AUR-REQ-ARC-009`, `AUR-REQ-ORCH-016` |
| `CAP-SOVEREIGN-CORE-REQ-012` | Current accepted operational state and current valid authority MUST govern M0 decisions over stale narrative, observational memory, cached projection or historical record. | critical | CONTRACT_TEST, SECURITY_TEST, FAULT_INJECTION | `AUR-REQ-MEM-005`, `AUR-REQ-MEM-014`, `AUR-REQ-ARC-012` |
| `CAP-SOVEREIGN-CORE-REQ-013` | A Project MUST preserve, at minimum for M0, its stable identity, operator-recognizable objective/context label, current accepted operational state, state revision, and current next action or explicit absence of a permitted next action. | critical | SCHEMA_VALIDATION, CONTRACT_TEST, USER_JOURNEY | `AUR-REQ-DOM-004`, `AUR-REQ-COG-005`, `AUR-REQ-RDM-007` |
| `CAP-SOVEREIGN-CORE-REQ-014` | When M0 state refers to artifacts or evidence, the Project record MUST preserve stable references and provenance and MUST NOT require duplicating the full referenced content inline. | medium | SCHEMA_VALIDATION, DOCUMENT_REVIEW | `AUR-REQ-DOM-004`, `AUR-REQ-DOM-014`, `AUR-REQ-ARC-013` |
| `CAP-SOVEREIGN-CORE-REQ-015` | Material M0 relationships among Aurora, Project, state revisions, authority state, events, audit records and evidence MUST preserve stable references, provenance, scope and temporal validity sufficient to determine which record governs. | high | SCHEMA_VALIDATION, INTEGRATION | `AUR-REQ-DOM-012`, `AUR-REQ-COG-002`, `AUR-REQ-ORCH-011`, `AUR-REQ-ARC-022` |
| `CAP-SOVEREIGN-CORE-REQ-016` | A Project snapshot, authority snapshot, rendered status or other M0 projection MUST reference its authoritative inputs and MUST NOT become an independent unaudited source of truth. | critical | CONTRACT_TEST, FAULT_INJECTION | `AUR-REQ-DOM-014`, `AUR-REQ-ARC-012`, `AUR-REQ-RELIA-005` |
| `CAP-SOVEREIGN-CORE-REQ-017` | Canonical M0 domain and state semantics MUST NOT embed provider-, protocol-, database-, workflow-engine-, telemetry-backend- or UI-framework-specific types as their product meaning. | high | STATIC_ANALYSIS, DOCUMENT_REVIEW | `AUR-REQ-ARC-002`, `AUR-REQ-ARC-007`, `AUR-REQ-ORCH-020`, `AUR-REQ-CAP-020` |
| `CAP-SOVEREIGN-CORE-REQ-018` | R3 MUST assign exactly one logical owner for every durable M0 concept and distinguish canonical state owners from adapters, projections, audit, telemetry and evidence responsibilities. | high | DOCUMENT_REVIEW | `AUR-REQ-ARC-006`, `AUR-REQ-DOC-001`, `AUR-REQ-ARC-003` |
| `CAP-SOVEREIGN-CORE-REQ-019` | Current accepted state, state-transition attempt, Domain Event, Audit record, Telemetry signal, Artifact, Claim, Evidence, Verdict and Outcome MUST remain logically distinct even if a later implementation co-locates some records physically. | critical | SCHEMA_VALIDATION, CONTRACT_TEST, DOCUMENT_REVIEW | `AUR-REQ-DOM-008`, `AUR-REQ-COG-011`, `AUR-REQ-ORCH-014`, `AUR-REQ-ARC-010`, `AUR-REQ-ARC-013`, `AUR-REQ-RELIA-002`, `AUR-REQ-VIS-006`, `AUR-REQ-DOM-002` |
| `CAP-SOVEREIGN-CORE-REQ-020` | Harness/provider/runtime-local state MAY exist in future integrations but MUST NOT be sufficient to establish or recover M0 canonical Project state or authority. | critical | INTEGRATION, FAULT_INJECTION | `AUR-REQ-ORCH-003`, `AUR-REQ-ORCH-016`, `AUR-REQ-ORCH-021`, `AUR-REQ-DOM-013`, `AUR-REQ-COG-007` |

## 6. State transition lifecycle

| ID | Requirement | Risk | Verification | Constitutional sources |
|---|---|---|---|---|
| `CAP-SOVEREIGN-CORE-REQ-021` | Every material M0 state-transition attempt MUST be attributable to a stable attempt identifier, Project identity, actor/source and recorded time. | high | SCHEMA_VALIDATION, CONTRACT_TEST | `AUR-REQ-COG-002`, `AUR-REQ-COG-009`, `AUR-REQ-ORCH-011`, `AUR-REQ-SEC-010` |
| `CAP-SOVEREIGN-CORE-REQ-022` | A state-transition request MUST identify or unambiguously bind the current state revision/precondition against which it is evaluated. | high | CONTRACT_TEST, FAULT_INJECTION | `AUR-REQ-DOM-012`, `AUR-REQ-COG-009`, `AUR-REQ-ARC-020` |
| `CAP-SOVEREIGN-CORE-REQ-023` | Operator-accepted intent/state input MUST remain distinguishable from any derived plan or next action recorded by M0. | high | SCHEMA_VALIDATION, CONTRACT_TEST | `AUR-REQ-COG-005`, `AUR-REQ-REL-011`, `AUR-REQ-DOM-014` |
| `CAP-SOVEREIGN-CORE-REQ-024` | A transition MAY become current accepted state only after its M0 preconditions, state-revision expectation and applicable current authority checks succeed. | critical | CONTRACT_TEST, SECURITY_TEST | `AUR-REQ-COG-009`, `AUR-REQ-AUT-001`, `AUR-REQ-AUT-003`, `AUR-REQ-ARC-023` |
| `CAP-SOVEREIGN-CORE-REQ-025` | A transition that violates the M0 state model, preconditions or authority MUST be rejected and MUST NOT alter current accepted state, authority state or derived next safe action. | critical | CONTRACT_TEST, FAULT_INJECTION, USER_JOURNEY | `AUR-REQ-RDM-007`, `AUR-REQ-COG-018`, `AUR-REQ-REL-013`, `AUR-REQ-RELIA-007` |
| `CAP-SOVEREIGN-CORE-REQ-026` | When a transition is accepted, the new current state revision and durable audit/evidence references required to explain that acceptance MUST become mutually consistent; a partial update MUST NOT be reported as successful completion. | critical | INTEGRATION, FAULT_INJECTION | `AUR-REQ-COG-010`, `AUR-REQ-COG-011`, `AUR-REQ-ARC-020`, `AUR-REQ-RELIA-001` |
| `CAP-SOVEREIGN-CORE-REQ-027` | A stale, conflicting or corrective transition MUST NOT silently overwrite the currently accepted state; it MUST either be rejected or recorded as an explicit new revision with attribution and preserved material history. | high | CONTRACT_TEST, FAULT_INJECTION | `AUR-REQ-REL-013`, `AUR-REQ-DOM-012`, `AUR-REQ-MEM-014`, `AUR-REQ-SEC-020` |
| `CAP-SOVEREIGN-CORE-REQ-028` | M0 MUST NOT report a transition, restart, recovery, export or restore operation as complete solely because it was dispatched; required resulting state/evidence MUST be observed first. | high | INTEGRATION, FAULT_INJECTION | `AUR-REQ-COG-010`, `AUR-REQ-VIS-010`, `AUR-REQ-RELIA-002` |
| `CAP-SOVEREIGN-CORE-REQ-029` | Material M0 records MUST be routed to their canonical state, audit or evidence owner and MUST NOT be persisted only as conversational narrative. | critical | DOCUMENT_REVIEW, INTEGRATION | `AUR-REQ-COG-012`, `AUR-REQ-MEM-001`, `AUR-REQ-ARC-009` |
| `CAP-SOVEREIGN-CORE-REQ-030` | The operator-facing M0 status surface MUST report material current state, authority validity and recovery/result state rather than raw low-level activity as a substitute for progress. | medium | USER_JOURNEY, DOCUMENT_REVIEW | `AUR-REQ-ORCH-012`, `AUR-REQ-VIS-010`, `AUR-REQ-SEC-019` |
| `CAP-SOVEREIGN-CORE-REQ-031` | CAP-SOVEREIGN-CORE MUST support the M0 Golden Proof sequence end to end: initialize Aurora, create Project, record accepted state and next action, terminate all Aurora processes, start fresh, recover same identities/state, reject an invalid transition, and export/restore required state. | critical | USER_JOURNEY, INTEGRATION, OPERATOR_VERDICT | `AUR-REQ-VIS-005`, `AUR-REQ-COG-001`, `AUR-REQ-RDM-001`, `AUR-REQ-RDM-007` |

## 7. Authority

| ID | Requirement | Risk | Verification | Constitutional sources |
|---|---|---|---|---|
| `CAP-SOVEREIGN-CORE-REQ-032` | Leandro MUST remain the final authority over M0 purpose, material trade-offs and authority grants or revocations. | critical | DOCUMENT_REVIEW, SECURITY_TEST | `AUR-REQ-REL-002`, `AUR-REQ-AUT-001` |
| `CAP-SOVEREIGN-CORE-REQ-033` | Possession of an interface, tool, credential, API access, database access, runtime capability or other technical ability MUST NOT by itself constitute authority to change M0 state or perform an effect. | critical | SECURITY_TEST, CONTRACT_TEST | `AUR-REQ-DOM-009`, `AUR-REQ-AUT-002`, `AUR-REQ-REL-011` |
| `CAP-SOVEREIGN-CORE-REQ-034` | M0 authority state MUST preserve a stable authority identity/revision and the minimum subject/actor, scope, permitted action class, relevant resource/Project scope, conditions and validity needed to determine current permitted next action. | critical | SCHEMA_VALIDATION, CONTRACT_TEST | `AUR-REQ-AUT-007`, `AUR-REQ-DOM-012`, `AUR-REQ-ARC-022` |
| `CAP-SOVEREIGN-CORE-REQ-035` | M0 authority MUST be explicitly scoped; when a requested transition or next action falls outside matching scope, it MUST be treated as not permitted. | critical | CONTRACT_TEST, SECURITY_TEST | `AUR-REQ-AUT-003`, `AUR-REQ-AUT-007`, `AUR-REQ-REL-002` |
| `CAP-SOVEREIGN-CORE-REQ-036` | M0 authority with a validity window MUST become non-permitting when validity expires; recovery MUST evaluate current validity rather than trust a stale persisted allowed result. | critical | SECURITY_TEST, FAULT_INJECTION | `AUR-REQ-AUT-003`, `AUR-REQ-AUT-007`, `AUR-REQ-SEC-022` |
| `CAP-SOVEREIGN-CORE-REQ-037` | A revoked M0 authority MUST remain revoked across restart, export, restore and migration unless a new explicitly authorized grant/change supersedes it. | critical | SECURITY_TEST, INTEGRATION, FAULT_INJECTION | `AUR-REQ-AUT-003`, `AUR-REQ-AUT-014`, `AUR-REQ-SEC-022` |
| `CAP-SOVEREIGN-CORE-REQ-038` | The M0 authority snapshot MUST be a traceable projection of current authoritative authority state and MUST include enough provenance/revision/validity information to determine whether it is current. | critical | SCHEMA_VALIDATION, CONTRACT_TEST | `AUR-REQ-DOM-014`, `AUR-REQ-AUT-001`, `AUR-REQ-SEC-019` |
| `CAP-SOVEREIGN-CORE-REQ-039` | The recorded or presented next safe action MUST be derived from current accepted Project state together with currently valid authority and MUST NOT itself constitute an Authority Grant. | critical | CONTRACT_TEST, SECURITY_TEST, USER_JOURNEY | `AUR-REQ-DOM-009`, `AUR-REQ-COG-017`, `AUR-REQ-AUT-002`, `AUR-REQ-AUT-003` |
| `CAP-SOVEREIGN-CORE-REQ-040` | Computing, recording or presenting a possible next action MUST NOT be treated as permission to execute an external effect; M0 does not authorize an effect plane. | high | DOCUMENT_REVIEW, SECURITY_TEST | `AUR-REQ-REL-011`, `AUR-REQ-AUT-002`, `AUR-REQ-RDM-005` |
| `CAP-SOVEREIGN-CORE-REQ-041` | A fresh process MUST reconstruct and validate current authority state before it presents a next action as permitted. | critical | INTEGRATION, SECURITY_TEST, FAULT_INJECTION | `AUR-REQ-COG-017`, `AUR-REQ-AUT-014`, `AUR-REQ-SEC-022` |
| `CAP-SOVEREIGN-CORE-REQ-042` | Restore MUST NOT silently reactivate expired or revoked authority and MUST NOT treat compromised or explicitly invalid trust/authority state as current permission. | critical | SECURITY_TEST, FAULT_INJECTION, USER_JOURNEY | `AUR-REQ-SEC-022`, `AUR-REQ-AUT-014`, `AUR-REQ-SEC-001` |
| `CAP-SOVEREIGN-CORE-REQ-043` | Untrusted or ordinary Project content MUST be treated as data and MUST NOT redefine M0 authority, policy, canonical ownership or accepted state except through an explicitly governed transition. | critical | SECURITY_TEST, CONTRACT_TEST | `AUR-REQ-SEC-013`, `AUR-REQ-DOM-009`, `AUR-REQ-ARC-023` |
| `CAP-SOVEREIGN-CORE-REQ-044` | If required M0 authority state is missing, corrupt, expired, revoked, incompatible or otherwise unverifiable, the Core MUST fail closed for permission/next-action decisions and preserve a diagnosable failure result. | critical | SECURITY_TEST, FAULT_INJECTION | `AUR-REQ-SEC-001`, `AUR-REQ-SEC-024`, `AUR-REQ-AUT-020`, `AUR-REQ-ARC-020` |
| `CAP-SOVEREIGN-CORE-REQ-045` | A detected M0 authority-integrity incident MUST prevent unsafe state mutation, preserve relevant evidence and surface a recoverable/blocked status rather than continue as though authority were valid. | critical | SECURITY_TEST, FAULT_INJECTION | `AUR-REQ-AUT-020`, `AUR-REQ-SEC-024`, `AUR-REQ-RELIA-024` |

## 8. Durability and recovery

| ID | Requirement | Risk | Verification | Constitutional sources |
|---|---|---|---|---|
| `CAP-SOVEREIGN-CORE-REQ-046` | All canonical M0 state required to satisfy the Golden Proof MUST outlive the process that created or last modified it. | critical | INTEGRATION, FAULT_INJECTION, USER_JOURNEY | `AUR-REQ-ARC-004`, `AUR-REQ-RDM-007`, `AUR-REQ-VIS-001` |
| `CAP-SOVEREIGN-CORE-REQ-047` | Restart recovery MUST NOT require in-memory objects, thread state, model context, transcript context or Harness/runtime-local memory from the terminated process. | critical | FAULT_INJECTION, INTEGRATION | `AUR-REQ-COG-017`, `AUR-REQ-MEM-006`, `AUR-REQ-ORCH-016`, `AUR-REQ-ARC-009` |
| `CAP-SOVEREIGN-CORE-REQ-048` | A successful fresh-process recovery MUST reconstruct the same Aurora identity, Project identity, current accepted Project state/revision, current authority state/snapshot and next safe action required by M0. | critical | INTEGRATION, USER_JOURNEY | `AUR-REQ-COG-017`, `AUR-REQ-RDM-007`, `AUR-REQ-RELIA-001` |
| `CAP-SOVEREIGN-CORE-REQ-049` | When required canonical state is absent or cannot be validated, recovery MUST report an explicit failure/degraded result and MUST NOT fabricate replacement accepted state, authority or next action. | critical | FAULT_INJECTION, SECURITY_TEST | `AUR-REQ-COG-018`, `AUR-REQ-SEC-001`, `AUR-REQ-ARC-020` |
| `CAP-SOVEREIGN-CORE-REQ-050` | M0 recovery failures MUST be classified at least sufficiently to distinguish process/restart failure, durable-state availability/integrity failure, authority-validation failure, export/restore failure and version/migration incompatibility. | high | FAULT_INJECTION, CONTRACT_TEST | `AUR-REQ-COG-018`, `AUR-REQ-RELIA-007`, `AUR-REQ-ARC-020` |
| `CAP-SOVEREIGN-CORE-REQ-051` | For each M0 failure class that can leave current state ambiguous or unsafe, the Capability MUST define a containment outcome that prevents unverifiable state from being presented as current. | critical | FAULT_INJECTION, SECURITY_TEST | `AUR-REQ-ARC-020`, `AUR-REQ-SEC-024`, `AUR-REQ-RELIA-001` |
| `CAP-SOVEREIGN-CORE-REQ-052` | A failed M0 operation MUST be retried only when classified as retryable and the retry is known safe/idempotent for affected state; blind retry of state mutation, migration or restore MUST be prohibited. | critical | FAULT_INJECTION, CONTRACT_TEST | `AUR-REQ-COG-018`, `AUR-REQ-RELIA-008`, `AUR-REQ-ARC-020` |
| `CAP-SOVEREIGN-CORE-REQ-053` | Retry behavior MUST NOT conceal a systematic failure or convert an unknown/ambiguous result into a success claim. | high | FAULT_INJECTION, INTEGRATION | `AUR-REQ-RELIA-008`, `AUR-REQ-COG-010` |
| `CAP-SOVEREIGN-CORE-REQ-054` | Restart/recovery MUST produce a structured result sufficient to identify recovered Project/state revision, authority-validation outcome and any limitation or failure classification. | high | INTEGRATION, USER_JOURNEY | `AUR-REQ-ORCH-012`, `AUR-REQ-RELIA-005`, `AUR-REQ-COG-010` |
| `CAP-SOVEREIGN-CORE-REQ-055` | Any M0 cache or projection used for operator display or convenience MUST either be reconstructable from canonical state or be explicitly non-authoritative after restart. | high | FAULT_INJECTION, DOCUMENT_REVIEW | `AUR-REQ-DOM-014`, `AUR-REQ-ARC-009`, `AUR-REQ-RELIA-005` |

## 9. Export, restore and migration

| ID | Requirement | Risk | Verification | Constitutional sources |
|---|---|---|---|---|
| `CAP-SOVEREIGN-CORE-REQ-056` | An M0 export MUST contain canonical state and metadata necessary for a fresh restore to reproduce M0 continuity obligations for Aurora identity, Project identity/state, authority state and next safe action. | critical | CONTRACT_TEST, INTEGRATION, USER_JOURNEY | `AUR-REQ-SEC-021`, `AUR-REQ-RDM-007`, `AUR-REQ-ARC-021` |
| `CAP-SOVEREIGN-CORE-REQ-057` | Every M0 export MUST identify the logical export/schema/contract version needed to determine restore compatibility without binding that identity to a particular storage product. | high | SCHEMA_VALIDATION, CONTRACT_TEST | `AUR-REQ-ARC-007`, `AUR-REQ-ARC-021`, `AUR-REQ-CAP-020` |
| `CAP-SOVEREIGN-CORE-REQ-058` | M0 export data MUST carry or reference integrity information sufficient for restore to detect material corruption before accepting restored state. | critical | SECURITY_TEST, FAULT_INJECTION | `AUR-REQ-SEC-001`, `AUR-REQ-SEC-021`, `AUR-REQ-RELIA-003` |
| `CAP-SOVEREIGN-CORE-REQ-059` | M0 export/restore correctness MUST NOT depend on transient process state, model context, transcript state, Harness local state or telemetry/logs; such data, if exported for evidence, MUST remain non-authoritative. | critical | CONTRACT_TEST, FAULT_INJECTION | `AUR-REQ-ARC-009`, `AUR-REQ-MEM-001`, `AUR-REQ-ORCH-016` |
| `CAP-SOVEREIGN-CORE-REQ-060` | Restore MUST validate required integrity, version compatibility and authority-safety conditions before restored data may become current canonical M0 state. | critical | SECURITY_TEST, FAULT_INJECTION, INTEGRATION | `AUR-REQ-SEC-022`, `AUR-REQ-ARC-021`, `AUR-REQ-COG-010` |
| `CAP-SOVEREIGN-CORE-REQ-061` | A successful M0 restore into a valid fresh restore context MUST reproduce stable Aurora/Project identities, accepted Project state/revision, authority validity/revocation state and next safe action. | critical | INTEGRATION, USER_JOURNEY | `AUR-REQ-SEC-021`, `AUR-REQ-RDM-007`, `AUR-REQ-RELIA-001` |
| `CAP-SOVEREIGN-CORE-REQ-062` | Restore MUST NOT silently overwrite or merge a different existing Aurora/Project identity or conflicting current state; such a collision MUST require explicit governed resolution or fail without changing current state. | high | FAULT_INJECTION, SECURITY_TEST | `AUR-REQ-DOM-001`, `AUR-REQ-DOM-012`, `AUR-REQ-SEC-001`, `AUR-REQ-ARC-020` |
| `CAP-SOVEREIGN-CORE-REQ-063` | Export, backup and restore state required by M0 MUST remain under Leandro-controlled governance and MUST expose a material result Leandro can inspect. | high | USER_JOURNEY, DOCUMENT_REVIEW | `AUR-REQ-SEC-007`, `AUR-REQ-SEC-019`, `AUR-REQ-SEC-021` |
| `CAP-SOVEREIGN-CORE-REQ-064` | Any M0 migration MUST preserve stable identities, accepted-state meaning/revision, authority validity/revocation semantics, provenance and evidence references needed by the M0 proof. | critical | INTEGRATION, FAULT_INJECTION | `AUR-REQ-ARC-019`, `AUR-REQ-ARC-021`, `AUR-REQ-SEC-021` |
| `CAP-SOVEREIGN-CORE-REQ-065` | When persisted/exported state is incompatible with current accepted M0 schema/contract, Core MUST require an explicit supported migration path or fail clearly; it MUST NOT silently coerce data in a way that can change governing semantics. | critical | FAULT_INJECTION, CONTRACT_TEST | `AUR-REQ-ARC-021`, `AUR-REQ-ARC-023`, `AUR-REQ-SEC-001` |
| `CAP-SOVEREIGN-CORE-REQ-066` | A material migration or restore MUST produce evidence identifying source version, target version, operation result and known limitations sufficient to support a later verdict. | high | INTEGRATION, DOCUMENT_REVIEW | `AUR-REQ-RELIA-003`, `AUR-REQ-SEC-024`, `AUR-REQ-ARC-021` |

## 10. Security and sovereignty

| ID | Requirement | Risk | Verification | Constitutional sources |
|---|---|---|---|---|
| `CAP-SOVEREIGN-CORE-REQ-067` | Canonical M0 identity, Project operational state, authority state, audit records and backup/export material MUST remain under infrastructure and administrative control governed by Leandro; an external model/provider MUST NOT be required as authority for them. | critical | DOCUMENT_REVIEW, SECURITY_TEST, INTEGRATION | `AUR-REQ-SEC-001`, `AUR-REQ-SEC-007`, `AUR-REQ-VIS-002`, `AUR-REQ-ORCH-002` |
| `CAP-SOVEREIGN-CORE-REQ-068` | R3 MUST assign an explicit accepted data classification, or explicitly equivalent accepted classification, to each material M0 data family including identity, Project operational state, authority, audit/evidence references and export/backup material. | high | DOCUMENT_REVIEW, SCHEMA_VALIDATION | `AUR-REQ-SEC-005`, `AUR-REQ-SEC-002`, `AUR-REQ-SEC-004` |
| `CAP-SOVEREIGN-CORE-REQ-069` | M0 MUST NOT require secret values to enter prompts or manifests, and secret values MUST NOT appear in general telemetry or general logs through the normal M0 path merely to persist, inspect or recover state. | high | SECURITY_TEST, STATIC_ANALYSIS | `AUR-REQ-SEC-009`, `AUR-REQ-ARC-022` |
| `CAP-SOVEREIGN-CORE-REQ-070` | Material operator/Core state transitions, recovery operations, export/restore operations and authority changes MUST preserve actor/source attribution appropriate to M0. | high | SCHEMA_VALIDATION, INTEGRATION | `AUR-REQ-SEC-010`, `AUR-REQ-DOM-012`, `AUR-REQ-RELIA-003` |
| `CAP-SOVEREIGN-CORE-REQ-071` | Ordinary or untrusted content stored in or referenced by a Project MUST NOT be interpreted as policy, authority or system configuration solely because it is persisted. | critical | SECURITY_TEST, CONTRACT_TEST | `AUR-REQ-SEC-013`, `AUR-REQ-ARC-023`, `AUR-REQ-DOM-014` |
| `CAP-SOVEREIGN-CORE-REQ-072` | Leandro MUST be able to inspect material M0 Project state, current authority status/snapshot provenance, material audit events, and export/restore result needed to understand what Core considers current. | high | USER_JOURNEY, DOCUMENT_REVIEW | `AUR-REQ-SEC-019`, `AUR-REQ-ORCH-012`, `AUR-REQ-VIS-010` |
| `CAP-SOVEREIGN-CORE-REQ-073` | M0 durable records MUST support explicit current-versus-superseded semantics and retention handling sufficient to avoid ambiguous current state; archive/deletion and the broader M1 privacy lifecycle remain deferred. | high | SCHEMA_VALIDATION, CONTRACT_TEST | `AUR-REQ-SEC-020`, `AUR-REQ-MEM-014`, `AUR-REQ-DOM-012` |
| `CAP-SOVEREIGN-CORE-REQ-074` | For security/integrity incidents relevant to M0, Capability MUST support detection of materially invalid or corrupt state, containment that prevents unsafe current-state or authority use, evidence preservation, and recovery/review hooks; the broader incident-response program remains deferred. | high | DOCUMENT_REVIEW, FAULT_INJECTION, SECURITY_TEST | `AUR-REQ-SEC-024`, `AUR-REQ-AUT-020`, `AUR-REQ-RELIA-024` |
| `CAP-SOVEREIGN-CORE-REQ-075` | CAP-SOVEREIGN-CORE MUST NOT pass R3 without a threat model covering at least canonical-state tampering/corruption, stale or rolled-back authority, unsafe restore, export/backup exposure, identity collision and untrusted-content authority injection. | critical | DOCUMENT_REVIEW | `AUR-REQ-SEC-003`, `AUR-REQ-SEC-001`, `AUR-REQ-SEC-004`, `AUR-REQ-ARC-023` |
| `CAP-SOVEREIGN-CORE-REQ-076` | M0 cross-component observability MUST propagate stable identifiers without sensitive payloads. Sensitive content needed as evidence MUST be governed as separate evidence or artifact data rather than carried in correlation or general telemetry payloads. | high | SECURITY_TEST, DOCUMENT_REVIEW | `AUR-REQ-ARC-022`, `AUR-REQ-SEC-005`, `AUR-REQ-RELIA-005` |

## 11. Event, audit, evidence and telemetry

| ID | Requirement | Risk | Verification | Constitutional sources |
|---|---|---|---|---|
| `CAP-SOVEREIGN-CORE-REQ-077` | Every material M0 transition/recovery/export/restore event MUST carry stable event identity, source, recorded time or sequence, correlation identity and classification sufficient for audit and proof. | high | SCHEMA_VALIDATION, INTEGRATION | `AUR-REQ-ORCH-011`, `AUR-REQ-COG-002`, `AUR-REQ-DOM-012`, `AUR-REQ-ARC-022` |
| `CAP-SOVEREIGN-CORE-REQ-078` | M0 audit MUST distinguish attempted, accepted and rejected state transitions and preserve material reason/classification for rejection. | high | CONTRACT_TEST, USER_JOURNEY | `AUR-REQ-RDM-007`, `AUR-REQ-ORCH-011`, `AUR-REQ-RELIA-007` |
| `CAP-SOVEREIGN-CORE-REQ-079` | M0 MUST record restart/recovery boundary and material recovery result so they can be correlated to recovered Aurora/Project/state identities. | high | INTEGRATION, USER_JOURNEY | `AUR-REQ-RDM-007`, `AUR-REQ-RELIA-003`, `AUR-REQ-ARC-022` |
| `CAP-SOVEREIGN-CORE-REQ-080` | M0 MUST record each material export and restore attempt/result, including logical source/target version information and integrity/authority validation outcome required for proof. | high | INTEGRATION, SECURITY_TEST | `AUR-REQ-SEC-021`, `AUR-REQ-RELIA-003`, `AUR-REQ-ARC-021` |
| `CAP-SOVEREIGN-CORE-REQ-081` | Domain Events, audit records, logs and telemetry MUST NOT be the sole canonical source of M0 current state or authority. | critical | DOCUMENT_REVIEW, FAULT_INJECTION | `AUR-REQ-ORCH-011`, `AUR-REQ-ARC-009`, `AUR-REQ-RELIA-005` |
| `CAP-SOVEREIGN-CORE-REQ-082` | Claim, Receipt, Evidence, Verdict and Outcome MUST remain distinct in M0 verification, and an implementation claim or command exit result MUST NOT by itself establish M0 success. | high | DOCUMENT_REVIEW, CONTRACT_TEST | `AUR-REQ-COG-011`, `AUR-REQ-ORCH-014`, `AUR-REQ-RELIA-002`, `AUR-REQ-RELIA-004` |
| `CAP-SOVEREIGN-CORE-REQ-083` | Material M0 Evidence MUST preserve criterion, producer, verifier, method, environment, relevant versions/revisions, artifact references, known uncertainty and limitations. | high | SCHEMA_VALIDATION, DOCUMENT_REVIEW | `AUR-REQ-RELIA-003`, `AUR-REQ-VIS-010`, `AUR-REQ-DOM-008` |
| `CAP-SOVEREIGN-CORE-REQ-084` | Every M0 Golden Proof run MUST have stable proof-run correlation identity linked to Aurora identity, Project identity and material state/recovery/export/restore evidence. | medium | SCHEMA_VALIDATION, USER_JOURNEY | `AUR-REQ-ARC-022`, `AUR-REQ-RELIA-003`, `AUR-REQ-RDM-001` |
| `CAP-SOVEREIGN-CORE-REQ-085` | Every retained M0 telemetry signal or metric MUST state the proof criterion, operational decision, threshold or recovery action it informs and SHOULD state what it cannot prove. | medium | DOCUMENT_REVIEW, SCHEMA_VALIDATION | `AUR-REQ-RELIA-005`, `AUR-REQ-RELIA-006`, `AUR-REQ-VIS-010` |
| `CAP-SOVEREIGN-CORE-REQ-086` | M0 traces, metrics and logs MUST NOT become canonical owner of Project state, authority or proof verdicts. | high | DOCUMENT_REVIEW, FAULT_INJECTION | `AUR-REQ-RELIA-005`, `AUR-REQ-ARC-009`, `AUR-REQ-ARC-010` |
| `CAP-SOVEREIGN-CORE-REQ-087` | Where M0 Golden Proof relies on state/export integrity, corresponding evidence MUST retain the integrity reference used for the check and link it to relevant state/export revision. | high | SCHEMA_VALIDATION, INTEGRATION | `AUR-REQ-RELIA-003`, `AUR-REQ-RDM-007`, `AUR-REQ-SEC-001` |
| `CAP-SOVEREIGN-CORE-REQ-088` | Passing a component test, completing a task or producing a local green result MUST NOT by itself close M0; M0 acceptance requires the end-to-end Golden Proof and later R8 verdict. | high | DOCUMENT_REVIEW, OPERATOR_VERDICT | `AUR-REQ-RELIA-004`, `AUR-REQ-RDM-001`, `AUR-REQ-RDM-007`, `AUR-REQ-VIS-010` |

## 12. Reliability and verification

| ID | Requirement | Risk | Verification | Constitutional sources |
|---|---|---|---|---|
| `CAP-SOVEREIGN-CORE-REQ-089` | M0 verification and later closeout MUST evaluate, as applicable to this slice, continuity, correctness, authority integrity, operational recovery, evidence quality, security/sovereignty and efficiency. | high | DOCUMENT_REVIEW, BENCHMARK/EVAL, USER_JOURNEY | `AUR-REQ-RELIA-001`, `AUR-REQ-RELIA-009`, `AUR-REQ-VIS-010` |
| `CAP-SOVEREIGN-CORE-REQ-090` | Capability MUST define M0-relevant error classes for invalid state transition, state availability/integrity, authority validation, version/migration, export/restore and internal operational failure before R3 can pass. | high | DOCUMENT_REVIEW, CONTRACT_TEST | `AUR-REQ-RELIA-007`, `AUR-REQ-COG-018`, `AUR-REQ-ARC-020` |
| `CAP-SOVEREIGN-CORE-REQ-091` | Accepted M0 requirements MUST be verifiable with negative cases for invalid transitions, invalid/expired/revoked authority, corrupt or incompatible persisted/exported state, and unsafe retry/restore conditions. | critical | FAULT_INJECTION, SECURITY_TEST, CONTRACT_TEST | `AUR-REQ-RDM-007`, `AUR-REQ-SEC-022`, `AUR-REQ-RELIA-008`, `AUR-REQ-ARC-020` |
| `CAP-SOVEREIGN-CORE-REQ-092` | An M0 proof or review supporting a completion verdict MUST identify exact accepted target revision and relevant state/schema/capability revisions being evaluated. | high | DOCUMENT_REVIEW | `AUR-REQ-RELIA-003`, `AUR-REQ-RDM-001`, `AUR-REQ-DOC-013` |
| `CAP-SOVEREIGN-CORE-REQ-093` | Every material M0 verification result MUST state known limitations, residual risks and unmet/deferred conditions constraining what evidence proves. | medium | DOCUMENT_REVIEW, OPERATOR_VERDICT | `AUR-REQ-RELIA-003`, `AUR-REQ-VIS-010`, `AUR-REQ-RELIA-001` |
| `CAP-SOVEREIGN-CORE-REQ-094` | A material M0 incident or requirement failure MUST update applicable requirements, documentation, evaluation scope and trust/assumptions as appropriate rather than be handled only as an implementation patch. | high | DOCUMENT_REVIEW | `AUR-REQ-RELIA-024`, `AUR-REQ-DOC-021`, `AUR-REQ-SEC-024` |
| `CAP-SOVEREIGN-CORE-REQ-095` | Before any M0 architecture decision is accepted, it MUST be reviewed for duplicated authority, provider/framework leakage into domain semantics, restart/recovery behavior, enforcement boundaries, context/state ownership and ability to produce the Golden Proof. | high | DOCUMENT_REVIEW | `AUR-REQ-ARC-023`, `AUR-REQ-VIS-011`, `AUR-REQ-RDM-005` |

## 13. Open decision guards

| ID | Requirement | Risk | Verification | Constitutional sources |
|---|---|---|---|---|
| `CAP-SOVEREIGN-CORE-REQ-096` | M0 architecture MUST preserve clear logical module boundaries before physical service distribution; physical separation MUST require evidence of need. | medium | DOCUMENT_REVIEW | `AUR-REQ-ARC-003`, `AUR-REQ-ARC-006`, `AUR-REQ-RDM-005` |
| `CAP-SOVEREIGN-CORE-REQ-097` | R4 SHOULD prefer the smallest local recoverable Core topology satisfying R2/R3 requirements; any distributed or microservice topology MUST be justified by evidence rather than assumed. | medium | DOCUMENT_REVIEW | `AUR-REQ-ARC-018`, `AUR-REQ-RDM-005`, `AUR-REQ-SEC-007` |
| `CAP-SOVEREIGN-CORE-REQ-098` | R2 MUST NOT select Aurora Core implementation language/runtime; choice remains an R4 decision informed by authorized M0 uncertainty reduction. | medium | DOCUMENT_REVIEW | `AUR-REQ-VIS-008`, `AUR-REQ-ARC-002`, `AUR-REQ-RDM-003`, `AUR-REQ-DOC-003` |
| `CAP-SOVEREIGN-CORE-REQ-099` | R2 MUST NOT select operational-state storage product/engine; R4 must choose a mechanism only after R2/R3 durability, ownership, recovery and security requirements are fixed. | high | DOCUMENT_REVIEW | `AUR-REQ-ARC-002`, `AUR-REQ-ARC-004`, `AUR-REQ-SEC-003`, `AUR-REQ-RDM-003` |
| `CAP-SOVEREIGN-CORE-REQ-100` | R2 MUST NOT require event sourcing, snapshot-only persistence or another state-versus-event pattern; R4 may select a mechanism only if canonical-state and event/audit distinctions remain satisfied. | high | DOCUMENT_REVIEW | `AUR-REQ-ARC-009`, `AUR-REQ-ARC-010`, `AUR-REQ-ORCH-011`, `AUR-REQ-DOC-003` |
| `CAP-SOVEREIGN-CORE-REQ-101` | R2 MUST define semantic fields/invariants without selecting source schema language, serialization format or code-generation stack. | medium | DOCUMENT_REVIEW | `AUR-REQ-ARC-007`, `AUR-REQ-CAP-020`, `AUR-REQ-DOC-003` |
| `CAP-SOVEREIGN-CORE-REQ-102` | R2 MUST define required event/audit semantics without selecting event transport, broker, append-only technology or audit backend. | medium | DOCUMENT_REVIEW | `AUR-REQ-ARC-010`, `AUR-REQ-ORCH-011`, `AUR-REQ-DOC-003` |
| `CAP-SOVEREIGN-CORE-REQ-103` | R2 MUST define proof/decision-oriented telemetry obligations without selecting telemetry backend, collector, transport or observability framework. | medium | DOCUMENT_REVIEW | `AUR-REQ-RELIA-005`, `AUR-REQ-RELIA-006`, `AUR-REQ-ARC-022` |
| `CAP-SOVEREIGN-CORE-REQ-104` | R2 MUST define backup/export/restore/migration behavior and evidence without selecting concrete backup topology, archive format, migration tool or storage-specific mechanism. | high | DOCUMENT_REVIEW | `AUR-REQ-SEC-021`, `AUR-REQ-ARC-021`, `AUR-REQ-DOC-003` |
| `CAP-SOVEREIGN-CORE-REQ-105` | M0 MUST NOT introduce a durable workflow/execution engine merely because restartable state exists; R4 MAY consider one only if accepted M0 requirements and evidence demonstrate that it is necessary and proportionate to this slice. | medium | DOCUMENT_REVIEW | `AUR-REQ-VIS-009`, `AUR-REQ-RDM-005`, `AUR-REQ-ARC-003` |
| `CAP-SOVEREIGN-CORE-REQ-106` | Persisting M0 operational state MUST NOT select or instantiate the later M1 technical memory architecture, vector-store ownership model or Context Builder. | high | DOCUMENT_REVIEW, STATIC_ANALYSIS | `AUR-REQ-MEM-023`, `AUR-REQ-MEM-001`, `AUR-REQ-RDM-005` |
| `CAP-SOVEREIGN-CORE-REQ-107` | R3/R4 MUST preserve replaceable adapter boundaries sufficient for later evolution from the initial local Core without requiring Aurora/Project/authority domain meaning to be rewritten for a new process topology or binding; later distributed stages need not be implemented in M0. | medium | DOCUMENT_REVIEW | `AUR-REQ-ARC-019`, `AUR-REQ-CAP-020`, `AUR-REQ-ORCH-020` |

## 14. Documentation and readiness

| ID | Requirement | Risk | Verification | Constitutional sources |
|---|---|---|---|---|
| `CAP-SOVEREIGN-CORE-REQ-108` | Every durable M0 concept/readiness artifact MUST have one declared canonical owner; material conflict MUST raise `DOCUMENTATION_DIVERGENCE` or an equivalent recorded Finding rather than be resolved silently. | high | DOCUMENT_REVIEW | `AUR-REQ-DOC-001`, `AUR-REQ-DOC-021`, `AUR-REQ-ARC-006` |
| `CAP-SOVEREIGN-CORE-REQ-109` | R2 requirements/coverage artifacts MUST declare stable document identity, authority, lifecycle/status, owners, source-of-truth scope, related artifacts and fixed source baseline. | medium | DOCUMENT_REVIEW | `AUR-REQ-DOC-004`, `AUR-REQ-DOC-012`, `AUR-REQ-DOC-013` |
| `CAP-SOVEREIGN-CORE-REQ-110` | Every CAP-SOVEREIGN-CORE requirement MUST trace to one or more R1-active constitutional sources or an explicitly accepted later decision; R2 MUST NOT invent unowned product intent. | critical | DOCUMENT_REVIEW | `AUR-REQ-DOC-001`, `AUR-REQ-DOC-002`, `AUR-REQ-RDM-020` |
| `CAP-SOVEREIGN-CORE-REQ-111` | Every one of 127 R1-active constitutional source rows MUST map to at least one R2 requirement or an explicit R2 rationale explaining why no additional atomic requirement is needed. | high | DOCUMENT_REVIEW | `AUR-REQ-DOC-017`, `AUR-REQ-DOC-018`, `AUR-REQ-RDM-003` |
| `CAP-SOVEREIGN-CORE-REQ-112` | Research reports, examples and candidate technologies MAY inform later R4 work but MUST NOT alter R2 normative requirements or become technical decisions without accepted decision process. | high | DOCUMENT_REVIEW | `AUR-REQ-DOC-003`, `AUR-REQ-DOC-007`, `AUR-REQ-DOC-008`, `AUR-REQ-DOC-009`, `AUR-REQ-DOC-010` |
| `CAP-SOVEREIGN-CORE-REQ-113` | Future R3 Capability Spec MUST own reusable CAP-SOVEREIGN-CORE behavior, while any R5 Mission Contract MUST own only an exact scoped implementation commitment; neither may silently substitute the other. | high | DOCUMENT_REVIEW | `AUR-REQ-DOC-011`, `AUR-REQ-RDM-003` |
| `CAP-SOVEREIGN-CORE-REQ-114` | STATUS MUST record current R2 verdict, authorizations/prohibitions, blockers, fixed verification targets and exact next action after R2 closeout. | high | DOCUMENT_REVIEW | `AUR-REQ-DOC-012`, `AUR-REQ-RDM-002` |
| `CAP-SOVEREIGN-CORE-REQ-115` | Every material R2/R3/R4 change MUST declare documentation impact and update/supersede correct canonical artifacts when product meaning or decision status changes. | medium | DOCUMENT_REVIEW | `AUR-REQ-DOC-015`, `AUR-REQ-DOC-013`, `AUR-REQ-RDM-020` |
| `CAP-SOVEREIGN-CORE-REQ-116` | R2 artifacts and closeout/tracking updates MUST pass repository documentation validation applicable to structure, authority, relations, traceability and generated-projection freshness. | medium | DOCUMENT_REVIEW | `AUR-REQ-DOC-016`, `AUR-REQ-DOC-006` |
| `CAP-SOVEREIGN-CORE-REQ-117` | R2 package MUST receive adversarial review checking atomicity, ambiguity, duplicate/conflicting requirements, verification direction, active-source coverage, false inclusion of deferred scope and hidden technical commitments. | high | DOCUMENT_REVIEW | `AUR-REQ-DOC-017`, `AUR-REQ-RDM-003`, `AUR-REQ-VIS-008` |
| `CAP-SOVEREIGN-CORE-REQ-118` | R2/R3 read paths MUST identify smallest correct authoritative source set needed to interpret CAP-SOVEREIGN-CORE while preserving links to accepted constitution and R1 applicability baseline. | medium | DOCUMENT_REVIEW | `AUR-REQ-DOC-018`, `AUR-REQ-DOC-001` |
| `CAP-SOVEREIGN-CORE-REQ-119` | File creation, branch updates, commits, merges or green CI MUST NOT by themselves constitute operator approval, gate verdict or authorization for next gate. | critical | DOCUMENT_REVIEW | `AUR-REQ-DOC-020`, `AUR-REQ-RDM-002`, `AUR-REQ-RDM-003` |
| `CAP-SOVEREIGN-CORE-REQ-120` | R2 MUST NOT silently change M0 outcome, named risk, Golden Proof direction, non-goals, dependencies or promotion/authority boundary; material roadmap change MUST return to canonical owner for evidence-based versioned approval. | critical | DOCUMENT_REVIEW, OPERATOR_VERDICT | `AUR-REQ-RDM-004`, `AUR-REQ-RDM-020`, `AUR-REQ-DOC-001` |
| `CAP-SOVEREIGN-CORE-REQ-121` | Completion of R2 MUST NOT authorize R3, Architecture Spike execution, stack selection, Mission Contract, Microdesign or implementation; those actions require their own later gate authorization. | critical | DOCUMENT_REVIEW | `AUR-REQ-RDM-002`, `AUR-REQ-RDM-003`, `AUR-REQ-DOC-020`, `AUR-REQ-VIS-008` |
| `CAP-SOVEREIGN-CORE-REQ-122` | R2 MUST NOT allocate these requirements to concrete files, packages, database schemas, runtime components or implementation tasks; allocation belongs to later accepted Spec/Contract/Plan artifacts. | high | DOCUMENT_REVIEW | `AUR-REQ-DOC-011`, `AUR-REQ-RDM-003`, `AUR-REQ-ARC-002` |

## 15. Requirement rationale by group

- **Identity and scope:** retire session/component identity risk while preventing future roadmap scope from leaking into M0.
- **Canonical ownership and state:** ensure one sovereign operational truth exists and remains distinct from narrative, telemetry, providers and projections.
- **State transition lifecycle:** make current-state mutation explicit, attributable, rejectable and verifiable.
- **Authority:** ensure current next action reflects valid human-governed authority and restore cannot revive stale permission.
- **Durability and recovery:** prove current state survives process death without transcript/model/Harness dependence and fails explicitly when it cannot be trusted.
- **Export, restore and migration:** make sovereignty portable while preserving identity, authority, integrity and version semantics.
- **Security and sovereignty:** constrain the first canonical durable store before R3/R4 select mechanisms.
- **Event, audit, evidence and telemetry:** make proof explainable without letting logs/events become operational truth.
- **Reliability and verification:** require negative-path and fixed-revision proof, not success-by-claim.
- **Open decision guards:** prevent R2 from smuggling stack choices into requirements.
- **Documentation and readiness:** keep artifact ownership, traceability, review and gate authorization explicit across sessions.

## 16. Open mechanisms intentionally not selected

R2 fixes required behavior while leaving open for later applicable gates: Core language/runtime; operational-state store; state-versus-event mechanism; schema/serialization/codegen representation; event/audit mechanism; telemetry backend; backup/export/restore mechanism/topology; migration tooling; process/deployment topology; interface technology; durable execution engine unless proven necessary; and exact Architecture Spike IDs/candidates/procedures/winners.

## 17. Stop boundary

```text
R2 requirement derivation
→ coverage and adversarial review
→ R2 PASS | FAIL | BLOCKED
→ STOP
→ R3 only after separate explicit operator authorization
```

No requirement in this document is implementation authorization or a substitute for the future Capability Spec, Mission Contract, Microdesign or Evidence Bundle.
