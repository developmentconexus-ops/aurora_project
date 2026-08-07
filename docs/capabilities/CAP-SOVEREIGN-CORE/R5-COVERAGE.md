---
id: DOC-AURORA-CAP-SOVEREIGN-CORE-R5-COVERAGE
title: CAP-SOVEREIGN-CORE R5 Mission Requirement Allocation
document_type: requirement_coverage
form: reference
authority: reference
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - current R5 allocation of every M0 CAP-SOVEREIGN-CORE requirement to MIS-M0-SOVEREIGN-CORE-001 criteria
related:
  - DOC-AURORA-MIS-M0-SOVEREIGN-CORE-001
  - DOC-AURORA-CAP-SOVEREIGN-CORE-REQUIREMENTS
  - DOC-AURORA-CAP-SOVEREIGN-CORE-R3-COVERAGE
  - DOC-AURORA-CAP-SOVEREIGN-CORE-TEST-PLAN
  - DOC-AURORA-M0-R5-OPERATOR-AUTHORIZATION
source_revision: 74167bd1404d9076423ffdbae20f97958283527c
last_reviewed: 2026-08-07
---

# CAP-SOVEREIGN-CORE — R5 Mission Requirement Allocation

## 1. Purpose

R5 must prove that the proposed Mission does not hide or orphan any accepted M0 requirement.

```text
Expected Capability requirements: 122
R5 allocation rows:              122
Unallocated:                       0
Mission criteria:                 12
```

Each requirement receives one **primary** Mission criterion. Cross-cutting requirements may also support the integrated Golden Proof, but primary ownership stays unique for discoverability.

Detailed requirement→Spec mechanism→planned test allocation remains owned by `R3-COVERAGE.md`; this file adds the R5 requirement→Mission criterion edge rather than duplicating the test catalog.

## 2. Allocation matrix

| Capability requirement | Primary Mission criterion | Verification source | Status |
|---|---|---|---|
"+table+"

## 3. Criterion summary

| Criterion | Primary requirements | Purpose |
|---|---|---|
| `CRIT-001` | `001..009` | stable identity and bounded M0 scope |
| `CRIT-002` | `010..020` | canonical state ownership |
| `CRIT-003` | `021..030` | revision-bound transition lifecycle |
| `CRIT-004` | `032..045` | authority and next-safe-action correctness |
| `CRIT-005` | `046..055` | fresh-process continuity and recovery |
| `CRIT-006` | `056..066` | export, restore and migration |
| `CRIT-007` | `067..076` | security, sovereignty and secret hygiene |
| `CRIT-008` | `077..088` | audit, evidence and telemetry separation |
| `CRIT-009` | `089..095` | reliability and fault containment |
| `CRIT-010` | `096..107` | accepted architecture guards |
| `CRIT-011` | `108..122` | documentation, traceability and readiness integrity |
| `CRIT-012` | `031` + composition | complete M0 Golden Proof |

## 4. Allocation rule

R6 may decompose criteria into implementation tasks and test groups but MUST preserve this traceability:

```text
Capability Requirement
→ primary Mission Criterion
→ R3 planned verification
→ R6 implementation/test unit
→ R7 receipt/evidence
```

A requirement cannot disappear because no convenient implementation task exists. Any material mismatch triggers Contract replan.
