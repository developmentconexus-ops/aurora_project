---
id: DOC-AURORA-REPOSITORY-ROADMAP
title: Aurora Repository Roadmap
document_type: repository_program_roadmap
form: reference
authority: tracking
status: current
program_status_authority: true
version: 1.5.0
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
  - DOC-AURORA-MR-01-PROMOTION-CLOSEOUT
  - DOC-AURORA-TA-03-EXECUTION-AUTHORIZATION
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
MR-01 Methodology/Repository/Readiness: OPERATOR-RATIFIED / ACCEPTED / MERGED / CANONICAL
TA-03 Cross-System Operation Surface: OPERATOR-AUTHORIZED / IN PROGRESS
M0 R0–R6: historical PASS within M0 scope
M0 R7 candidate: FROZEN / PRESERVED / NON-CANONICAL
M0 R7 Verdict: NOT ISSUED
M0 R8: NOT AUTHORIZED
```

## MR-01 closeout

```text
RM-01: PASS — semantic/provenance census
RM-02: PASS — bounded constitutional/method reconciliation
RM-03: PASS — target authority preparation
RM-04: PASS — decision/architecture/capability reconciliation
RM-05: PASS — atomic repository control-plane cutover
RM-06: PASS — legacy live-surface retirement/rehome
RM-07: PASS — positive validation + deterministic negative controls
RM-08: PASS — GitHub main protection + squash-only merge policy applied/revalidated
RM-09: PASS — fresh-actor + Global Coherence proof
RM-10: PASS — isolated Fable review CONVERGED / PASS_WITH_FINDINGS; Lead adjudication complete; Round 2 not required
promotion PR #8: SQUASH MERGED / CLOSED
promotion candidate: d8727e8564cd1e91a0bd7335e2b3649280592459
canonical main after promotion: 6cc92067b56aeb0de260d36b764f1d7220944781
MR-01 Repository Migration: ACCEPTED / MERGED / CLOSED
```

Evidence:

```text
RM-07 Documentation: 32651194229 — SUCCESS
RM-09 Coherence Audit: 32651194244 — SUCCESS
RM-08 exact Evidence: docs/evidence/mr-01-platform-enforcement.md
RM-10 review PR: #9 — CLOSED / NOT MERGED
RM-10 review Evidence head: abf035d7cdefe7c7339fd1dcdf4dcaef2e20c025
RM-10 verdict: CONVERGED / PASS_WITH_FINDINGS
RM-10 findings: 0 blocking / 1 material / 2 moderate / 4 minor
RM-10 RED regression run: 32666339230 — expected FAILURE reproducing RM-I01/RM-I02
RM-10 corrected candidate validation: 32666630531 — SUCCESS
RM-10 final candidate validation: 32666770939 — SUCCESS
RM-10 regression suite: 7/7 PASS
RM-10 adjudication: docs/evidence/mr-01-rm10-independent-review.md
promotion closeout: docs/evidence/mr-01-promotion-closeout.md
candidate/main tree identity: b5332ea1f6c1e942ec587e96e7394d752c5b8403
```

RM-I01, RM-I02, RM-I03, RM-I05 and RM-I07 were corrected without Product/architecture-owner change. RM-I04 required no candidate change. RM-I06 remains `DEFER_SAFELY`: before either frozen M0 R7 branch is deleted, renamed or cleaned, create and verify a durable archival tag/ref at the exact frozen commit(s).

## Current program gate

```text
current planning stage: TA-03 — Cross-System Operation Surface
TA-03 execution: OPERATOR-AUTHORIZED / IN PROGRESS
TA-04+: NOT AUTHORIZED
Architecture Spike execution: NOT AUTHORIZED
Aurora Product/runtime implementation: BLOCKED
M0 R7 continuation/Verdict/R8: NOT AUTHORIZED
framework/database/IAM/model/provider selection: NOT AUTHORIZED BY TA-03
```

TA-03 authorization is recorded by `DOC-AURORA-TA-03-EXECUTION-AUTHORIZATION`. It authorizes architectural discovery/design/review only. It does not imply acceptance, merge, TA-04 progression, Spike execution, technology selection or Product implementation.

## Exact next action

```text
execute TA-03 from accepted Planning Readiness + TA-01/TA-02 authority
→ derive real cross-system consumers / protected properties / boundary crossings
→ admit only operations whose cross-system semantics must be owned now
→ compare and resolve material semantic alternatives proportionally
→ STOP and reopen the smallest owning authority if a required decision belongs elsewhere
→ present the resulting TA-03 semantic design to the operator before ratification/promotion
```

During TA-03 do not resume the frozen M0 R7 execution path and do not choose TA-04 wire/schema/transport mechanisms by implication.

## Reopen triggers

Reopen MR-01 only if evidence shows lost current semantics/provenance, duplicate/missing authority, an unworkable fresh-actor route, platform-enforcement regression, or a downstream material decision that cannot be owned coherently by the accepted readiness graph.

Reopen TA-01/TA-02 only if TA-03 finds material evidence that an accepted owner is missing/duplicated or an accepted runtime/provider boundary cannot support a required cross-system operation without violating its invariant. A naming/numbering artifact from the superseded Technical Architecture ordering is not itself a reopen trigger.
