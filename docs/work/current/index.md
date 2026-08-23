---
id: DOC-AURORA-WORK-METHODOLOGY-REBASELINE-INDEX
title: Aurora Methodology and Repository Rebaseline Current Work
document_type: temporary_work_index
form: reference
authority: tracking
status: current
version: 0.6.0
owners:
  - developmentconexus-ops
related:
  - DOC-AURORA-METHODOLOGY-REPOSITORY-REBASELINE
  - DOC-AURORA-PLANNING-READINESS
  - DOC-AURORA-ENGINEERING-RULES
last_reviewed: 2026-08-23
---

# Aurora methodology + repository rebaseline — current work

> **NON-AUTHORITATIVE / BRANCH-ONLY.** This directory must be absorbed or deleted before a final merge candidate is promoted.

## Current gate

```text
MR-01 — Methodology, Planning-Readiness & Repository Rebaseline
operator design-direction decision: APPROVED — 2026-08-23
independent Fable review: COMPLETE — CONVERGED / PASS_WITH_FINDINGS
independent review PR #7: CLOSED / NOT MERGED
Lead adjudication: COMPLETE
independent blocking findings: 0
TA-01/TA-02 reopen: NOT REQUIRED
second independent round: NOT REQUIRED for accepted bounded corrections
bounded corrections: APPLIED / POST-ADJUDICATION VALIDATION PENDING
final MR-01 operator ratification: PENDING
repository migration execution: NOT AUTHORIZED
TA-03+: HOLD
Aurora implementation: BLOCKED
M0 R7: FROZEN / NON-CANONICAL
merge: NOT AUTHORIZED
```

## Review identity

```text
base main: 35614c581cea32e04305c1ad63522fee151eb283
candidate branch: docs/methodology-repository-rebaseline-20260823
Draft PR: #6
independent candidate reviewed: d3d453401a9b4244b14e6e833d0aeb9da2caea41
review branch: review/mr-01-fable
review output HEAD: 5ab89f8cc21c120afa439a89f35983cf072843b1
review verdict: CONVERGED / PASS_WITH_FINDINGS
```

The candidate changed after review only through the bounded corrections accepted in `independent-review-adjudication.md`. Fable explicitly stated those correction classes do not invalidate first-round coverage and do not require Round 2.

## Durable PROPOSED owners

```text
docs/decisions/methodology-repository-rebaseline.md
docs/development/planning-readiness.md
docs/development/engineering-rules.md
```

They remain `PROPOSED` until final operator ratification and canonical integration.

## Temporary Evidence / migration material

```text
docs/work/current/proposal.md
docs/work/current/adversarial-review.md
docs/work/current/blueprint-15-amendment.md
docs/work/current/plan.md
docs/work/current/independent-review-adjudication.md
```

The Fable `ai-dialog.md` remains only in closed unmerged PR #7 / review branch and never enters this candidate.

All `docs/work/**` material MUST be absorbed/deleted before a final merge candidate or `main` promotion.

## Adjudicated independent findings

```text
MR-I01 MATERIAL  → CORRECTED — residual Blueprint 15 old-route/status clauses enumerated
MR-I02 MODERATE  → CORRECTED — tracking/DECISIONS explicit census/rehome/retirement
MR-I03 MATERIAL  → CORRECTED — one migration branch + one final merge candidate; no partial RM merges
MR-I04 MODERATE  → CORRECTED — bounded non-production pre-TA-09 qualification/proof surface
MR-I05 MODERATE  → CORRECTED — review-isolation guard + blocked-implementation allowlist
MR-I06 MINOR     → CORRECTED — Conexus OS repository/Harness referent disambiguated
MR-I07 MINOR     → DEFER_SAFELY — exact frontend-method id/version/location pinned only at adoption
```

Additional independent recommendation incorporated: current A0 acceptance records are explicit RM-01 census/provenance subjects.

## Current exact next action

```text
run Documentation validation against exact post-adjudication candidate HEAD
→ confirm PR #6 remains mergeable and no unresolved material finding remains
→ present MR-01 for final operator ratification
```

Final ratification still does **not** authorize repository migration execution, TA-03+, Product implementation or merge by implication.

## Hard boundaries

Not authorized:

- Aurora Product/runtime implementation;
- repository migration execution;
- M0 R7 continuation/Verdict/R8;
- TA-03 or any later technical-stage execution;
- Architecture Spike execution;
- production repository/source restructuring;
- framework/database/authentication/provider selection;
- merge.
