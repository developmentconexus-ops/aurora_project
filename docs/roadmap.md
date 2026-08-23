---
id: DOC-AURORA-REPOSITORY-ROADMAP
title: Aurora Repository Roadmap
document_type: repository_program_roadmap
form: reference
authority: tracking
status: current
program_status_authority: true
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - current repository stage and gate
  - implementation permission and blocked work
  - exact next action
  - progression and reopen triggers
related:
  - DOC-AURORA-MR-01-OPERATOR-RATIFICATION
  - DOC-AURORA-MR-01-MIGRATION-AUTHORIZATION
  - DOC-AURORA-PLANNING-READINESS
last_reviewed: 2026-08-23
---

# Aurora Repository Roadmap

> This is the sole mutable repository-program status/permission/next-action authority. Product capability sequence remains in Blueprint 14.

## Current canonical baseline

```text
A0 Product constitution: ACCEPTED / MERGED
System Architecture Rebaseline: ACCEPTED / MERGED
TA-01 Logical Modules & Canonical Ownership: ACCEPTED / CANONICAL
TA-02 Process/Runtime/Evolutionary Topology: ACCEPTED / CANONICAL
M0 R0–R6: historical PASS within M0 scope
M0 R7 candidate: FROZEN / PRESERVED / NON-CANONICAL
M0 R7 Verdict: NOT ISSUED
M0 R8: NOT AUTHORIZED
MR-01 Methodology/Repository/Readiness target: OPERATOR-RATIFIED
```

## Current program

```text
program: MR-01 Repository Migration
migration PR: #8
RM-01: PASS — semantic/provenance census
RM-02: APPLIED — bounded constitutional/method reconciliation
RM-03: APPLIED — target authority preparation
RM-04: APPLIED — decision/architecture/capability reconciliation
RM-05: APPLIED — atomic repository control-plane cutover
RM-06: APPLIED — legacy live-surface retirement/rehome
RM-07: NEXT — verification + deterministic negative controls
RM-08: PENDING — Git/branch-protection enforcement
RM-09: PENDING — fresh-actor + Global Coherence proof
RM-10: PENDING — independent final review / promotion candidate
```

## Authorization boundary

```text
repository migration execution: AUTHORIZED
TA-03+: NOT AUTHORIZED
Architecture Spike execution: NOT AUTHORIZED
Aurora Product/runtime implementation: BLOCKED
M0 R7 continuation/Verdict/R8: NOT AUTHORIZED
framework/database/IAM/model/provider selection: NOT AUTHORIZED BY MIGRATION
merge: NOT AUTHORIZED
```

## Exact next action

```text
run RM-07 positive validation + guard negative controls on the migrated tree
→ correct only migration defects
→ RM-08 platform enforcement
→ RM-09 fresh-actor/global coherence proof
→ prepare isolated independent RM-10 review
→ STOP before merge authorization
```

## Reopen triggers

Reopen the MR-01 target only if migration Evidence shows lost current semantics/provenance, duplicate/missing authority, an unworkable fresh-actor route, or a downstream material decision forced before its owner exists.
