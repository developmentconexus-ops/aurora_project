---
id: DOC-AURORA-REPOSITORY-ROADMAP
title: Aurora Repository Roadmap
document_type: repository_program_roadmap
form: reference
authority: tracking
status: current
program_status_authority: true
version: 1.6.0
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
  - DOC-AURORA-TA-03-DESIGN-REVIEW
  - DESIGN-AURORA-TA-03-CROSS-SYSTEM-OPERATION-SURFACE
last_reviewed: 2026-09-05
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
TA-03 Cross-System Operation Surface: OPERATOR-AUTHORIZED / IN PROGRESS; scoped admission STOP below
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
Section 1 admission/mutation law: OPERATOR-APPROVED FOR DESIGN; recorded in proposed architecture
Section 2 / complete operation catalogue: NOT ACCEPTED
TA03-F01 first M1 cognitive-provider admission: STOP / SPLIT PREREQUISITE
upstream C05 disposition amendment: NOT AUTHORIZED / NOT APPLIED
independent TA-03 review: NOT COMPLETED
TA-03 ratification / merge: NOT AUTHORIZED
TA-04+: NOT AUTHORIZED
Architecture Spike execution: NOT AUTHORIZED
Aurora Product/runtime implementation: BLOCKED
M0 R7 continuation/Verdict/R8: NOT AUTHORIZED
framework/database/IAM/model/provider selection: NOT AUTHORIZED BY TA-03
```

TA-03 execution authorization remains recorded by `DOC-AURORA-TA-03-EXECUTION-AUTHORIZATION`. The Section-1 approval and bounded TA03-F01 finding are recorded in [design-review Evidence](evidence/ta-03-design-review.md); the [proposed operation design](architecture/cross-system-operation-surface.md) does not become canonical through section approval.

TA03-F01 suspends only admission of the disputed M1/first-cognitive-provider path and claims of complete catalogue readiness. Unaffected TA-03 analysis remains authorized. The accepted C05 ownership and A05 readiness/approval requirements are preserved; the upstream owner must resolve how the proposed first consumer is compatible with C05's explicit M2 implementation deferral. MR-01 and unrelated TA-01/TA-02 decisions are not reopened.

## Exact next action

```text
present TA03-F01 and its bounded alternatives to the operator
→ request a scoped resolution at module-runtime-topology.md: C05 Stage-A disposition, read with A02/A05/D02
→ do not change accepted upstream semantics by implication
→ after an accepted owning resolution, update only affected TA-03 consumer/precondition records
→ complete operation catalogue and cross-operation review
→ independent challenge, adjudication and exact-revision verification
→ separate operator TA-03 ratification and merge decision
```

The recommended resolution is a minimum C05-owned prerequisite before the first consumer that requires it, while leaving the complete M2 Registry/AHDK/reference-provider capability at M2. This recommendation is not an accepted amendment. An existing accepted path that satisfies the disputed prerequisite can instead falsify the finding and permit correction of TA-03 interpretation only.

No frozen M0 R7 execution, TA-04 wire/schema/transport work, technology selection or runtime implementation follows from this checkpoint.

## Reopen triggers

Reopen MR-01 only if evidence shows lost current semantics/provenance, duplicate/missing authority, an unworkable fresh-actor route, platform-enforcement regression, or a downstream material decision that cannot be owned coherently by the accepted readiness graph.

Reopen only the implicated TA-01/TA-02 clause when material evidence shows missing/duplicated ownership or an accepted prerequisite/runtime boundary cannot support a required operation without violating its invariant. TA03-F01 currently requests that bounded owning resolution; it does not revoke the accepted owner/topology baseline. A naming/numbering artifact from the superseded Technical Architecture ordering is not itself a reopen trigger.
