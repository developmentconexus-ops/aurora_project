---
id: DOC-AURORA-ADR-INDEX
title: Aurora ADR Index
document_type: adr_index
form: reference
authority: decision
status: proposed
version: 0.2.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - Aurora ADR lifecycle
  - architecture decision discovery
  - ADR numbering and current status
related:
  - DOC-AURORA-DOCUMENTATION-MAP
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DOC-AURORA-RESEARCH-MAP
last_reviewed: 2026-08-06
---

# Architecture Decision Records

## 1. Purpose

An ADR records one material decision after its context, alternatives and evidence are understood.

An ADR is not:

- a research report;
- a brainstorm;
- a feature request;
- an implementation plan;
- a retroactive justification for code already written;
- permission to begin the next readiness gate.

## 2. Authority boundary

- Product Blueprint owns constitutional meaning.
- ADR owns a specific accepted/rejected choice compatible with that constitution.
- Capability Spec applies the choice to reusable product behavior.
- Mission Contract binds exact scope/version for implementation.
- Code and Evidence prove the choice in a real system.

An ordinary ADR cannot override a constitutional invariant. A constitutional change requires explicit Blueprint revision and operator approval.

## 3. Lifecycle

```text
DRAFT
→ PROPOSED
→ ACCEPTED | REJECTED | WITHDRAWN
→ SUPERSEDED
```

### DRAFT

Incomplete working material; not ready for decision.

### PROPOSED

Context, alternatives, evidence, trade-offs and validation path are sufficient for review. Not governing implementation.

### ACCEPTED

Approved by the required authority for its stated scope. Becomes an A1 decision source.

### REJECTED

Reviewed and intentionally not chosen. Preserved to prevent rediscovery.

### WITHDRAWN

Proposal removed before a decision because scope or premise changed.

### SUPERSEDED

Previously accepted decision replaced by a newer accepted ADR. History remains discoverable.

## 4. Required content

A material ADR contains:

1. decision status and scope;
2. context/problem;
3. decision drivers and constraints;
4. related Blueprint requirements;
5. research and spike evidence;
6. alternatives, including “do nothing” when meaningful;
7. decision and explicit non-decision;
8. positive/negative consequences;
9. security, privacy, reliability and operational implications;
10. compatibility/migration/rollback where applicable;
11. verification and graduation evidence;
12. reconsideration triggers;
13. supersession metadata.

## 5. Acceptance gate

Before acceptance:

- relevant ACRM R0–R3 gates are current;
- material uncertainty has evidence or an authorized spike;
- the decision does not silently expand current Product Milestone scope;
- affected requirements and documents are identified;
- operational burden and exit conditions are explicit;
- Leandro reviews material architecture decisions.

Acceptance of an ADR does not automatically authorize a Capability Spec, Mission Contract, Architecture Spike or implementation.

## 6. Index

| ADR | Title | Status | Primary decision |
|---|---|---|---|
| [ADR-0001](0001-aurora-owned-contract-model.md) | Aurora-owned Contract Model and Replaceable Bindings | proposed | Aurora owns cross-Harness semantics; protocols remain bindings |
| [ADR-0002](0002-first-party-harness-development-kit.md) | First-party Harness Development Kit and Universal Conformance | proposed | first-party Harnesses use AHDK by policy; conformance remains universal |

No ADR is accepted in A0 yet.

## 7. Numbering and filenames

```text
0001-short-kebab-title.md
0002-short-kebab-title.md
```

IDs are stable:

```text
ADR-AURORA-0001
```

A rejected or withdrawn number is never reused.

## 8. Review and supersession

When new evidence conflicts with an accepted ADR:

```text
open Finding
→ assess affected requirements/contracts
→ propose new ADR or amendment path
→ preserve old decision
→ accept replacement
→ mark old ADR SUPERSEDED
→ re-run affected readiness gates
```

Do not rewrite accepted history to make the old decision appear obvious or never chosen.

## 9. Research relationship

```text
Research
→ describes evidence and alternatives

Architecture Spike
→ produces executable evidence

ADR
→ records the choice

Capability Spec / Contract
→ applies the choice

Implementation / Evidence
→ tests the choice in operation
```

Research detail may be summarized in the ADR, but primary findings and limitations remain in the research report.
