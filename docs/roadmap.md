---
id: DOC-AURORA-REPOSITORY-ROADMAP
title: Aurora Repository Roadmap
document_type: repository_program_roadmap
form: reference
authority: tracking
status: current
program_status_authority: true
version: 1.2.0
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
  - DOC-AURORA-MR-01-RM08-PLATFORM-ENFORCEMENT
  - DOC-AURORA-MR-01-RM09-COHERENCE
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
RM-01: PASS — semantic/provenance census; durable Evidence recorded
RM-02: PASS — bounded constitutional/method reconciliation
RM-03: PASS — target authority preparation
RM-04: PASS — decision/architecture/capability reconciliation
RM-05: PASS — atomic repository control-plane cutover
RM-06: PASS — legacy live-surface retirement/rehome
RM-07: PASS — positive validation + deterministic negative controls
RM-08: PASS — GitHub main protection + squash-only merge policy applied and revalidated
RM-09: PASS — fresh-actor + Global Coherence proof
RM-10: NEXT — isolated independent final review / promotion candidate
```

Evidence:

```text
RM-07 Documentation: 32651194229 — SUCCESS
RM-09 Coherence Audit: 32651194244 — SUCCESS
RM-08 main before operator enforcement: protected=false
RM-08 prior API credential probes: 32652270369 / 32652317828 — HTTP 403, preserved as historical Evidence
RM-08 post-operator verification:
  main SHA: 35614c581cea32e04305c1ad63522fee151eb283
  protected: true
  squash: true
  merge commits: false
  rebase: false
  auto-delete merged head branches: true
  PR #8: OPEN / DRAFT / NOT MERGED
  exact head before RM-08 closeout bookkeeping: fed6d962f24c745dee2167507b1cc061eae935f9
  Documentation / validate: 32652516829 — SUCCESS
RM-08 exact Evidence: docs/evidence/mr-01-platform-enforcement.md
```

The effective GitHub Ruleset internals are not individually enumerable through the connected read surface. The operator applied the exact prescribed PR/check/force-push/deletion rule configuration; independent machine-readable repository state confirms `main` is protected and the repository merge policy is now squash-only with automatic head-branch cleanup.

## Authorization boundary

```text
repository migration execution: AUTHORIZED
RM-10 independent review: AUTHORIZED BY COMPLETION OF RM-01..RM-09
TA-03+: NOT AUTHORIZED
Architecture Spike execution: NOT AUTHORIZED
Aurora Product/runtime implementation: BLOCKED
M0 R7 continuation/Verdict/R8: NOT AUTHORIZED
framework/database/IAM/model/provider selection: NOT AUTHORIZED BY MIGRATION
merge: NOT AUTHORIZED
```

## Exact next action

```text
run final Documentation validation on the exact post-RM-08 closeout head
→ freeze the migration candidate revision
→ create isolated review/mr-01-migration-fable from that exact revision
→ allow review branch to differ only by docs/work/current/ai-dialog.md
→ obtain independent RM-10 Evidence
→ Lead adjudicates every finding
→ if a material correction invalidates coverage, isolated Round 2
→ if converged, present separate operator merge/promotion gate
→ STOP before merge
```

## Reopen triggers

Reopen the MR-01 target only if migration Evidence shows lost current semantics/provenance, duplicate/missing authority, an unworkable fresh-actor route, platform enforcement regression, or a downstream material decision forced before its owner exists.
