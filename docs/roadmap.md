---
id: DOC-AURORA-REPOSITORY-ROADMAP
title: Aurora Repository Roadmap
document_type: repository_program_roadmap
form: reference
authority: tracking
status: current
program_status_authority: true
version: 1.3.0
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
  - DOC-AURORA-MR-01-RM10-INDEPENDENT-REVIEW
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
RM-10: PASS — isolated Fable review CONVERGED / PASS_WITH_FINDINGS; Lead adjudication complete; Round 2 not required
```

Evidence:

```text
RM-07 Documentation: 32651194229 — SUCCESS
RM-09 Coherence Audit: 32651194244 — SUCCESS
RM-08 post-operator verification:
  main SHA: 35614c581cea32e04305c1ad63522fee151eb283
  protected: true
  squash: true
  merge commits: false
  rebase: false
  auto-delete merged head branches: true
  PR #8: OPEN / DRAFT / NOT MERGED
RM-08 exact Evidence: docs/evidence/mr-01-platform-enforcement.md

RM-10 reviewed candidate: eb402577c3cda638102408543c412b70c72b18d4
RM-10 review PR: #9 — NEVER MERGE
RM-10 review Evidence head: abf035d7cdefe7c7339fd1dcdf4dcaef2e20c025
RM-10 review Documentation: 32665723826 — SUCCESS
RM-10 verdict: CONVERGED / PASS_WITH_FINDINGS
RM-10 findings: 0 blocking / 1 material / 2 moderate / 4 minor
RM-10 RED regression run: 32666339230 — expected FAILURE reproducing RM-I01/RM-I02
RM-10 corrected candidate validation: 32666630531 — SUCCESS
RM-10 regression suite: 7/7 PASS
RM-10 adjudication: docs/evidence/mr-01-rm10-independent-review.md
```

RM-I01, RM-I02, RM-I03, RM-I05 and RM-I07 were corrected without semantic-owner change. RM-I04 requires no candidate change; mechanical review-isolation proof already exists. RM-I06 is `DEFER_SAFELY`: before either frozen M0 R7 branch is deleted/renamed/cleaned, create and verify a durable archival tag/ref at the exact frozen commit(s).

## Authorization boundary

```text
repository migration execution: COMPLETE THROUGH RM-10
MR-01 merge/promotion: PENDING EXPLICIT OPERATOR DECISION
TA-03+: NOT AUTHORIZED
Architecture Spike execution: NOT AUTHORIZED
Aurora Product/runtime implementation: BLOCKED
M0 R7 continuation/Verdict/R8: NOT AUTHORIZED
framework/database/IAM/model/provider selection: NOT AUTHORIZED BY MIGRATION
merge: NOT AUTHORIZED UNTIL EXPLICIT OPERATOR GRANT
```

## Exact next action

```text
run final Documentation validation on the exact post-adjudication candidate head
→ confirm PR #8 remains OPEN / DRAFT / NOT MERGED and main has not drifted
→ close RM-10 review PR #9 unmerged after Evidence preservation
→ present the exact MR-01 repository migration candidate to the operator
→ request separate merge/promotion decision
→ STOP before merge
```

## Reopen triggers

Reopen the MR-01 target only if migration Evidence shows lost current semantics/provenance, duplicate/missing authority, an unworkable fresh-actor route, platform enforcement regression, a material RM-10 correction invalidating review coverage, or a downstream material decision forced before its owner exists.
