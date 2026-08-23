---
id: DOC-AURORA-WORK-METHODOLOGY-REBASELINE-INDEX
title: Aurora Methodology and Repository Rebaseline Current Work
document_type: temporary_work_index
form: reference
authority: tracking
status: current
version: 0.7.0
owners:
  - developmentconexus-ops
related:
  - DOC-AURORA-METHODOLOGY-REPOSITORY-REBASELINE
  - DOC-AURORA-PLANNING-READINESS
  - DOC-AURORA-ENGINEERING-RULES
  - DOC-AURORA-MR-01-OPERATOR-RATIFICATION
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
bounded corrections: APPLIED / VALIDATED
final MR-01 operator ratification: COMPLETE — 2026-08-23
MR-01 semantic target: OPERATOR-RATIFIED / ACCEPTED
canonical main integration: NOT PERFORMED
repository migration execution: NOT AUTHORIZED
TA-03+: HOLD / NOT AUTHORIZED
Aurora implementation: BLOCKED
M0 R7: FROZEN / NON-CANONICAL
merge: NOT AUTHORIZED
```

## Ratification identity

```text
base main: 35614c581cea32e04305c1ad63522fee151eb283
candidate branch: docs/methodology-repository-rebaseline-20260823
Draft PR: #6
independent candidate reviewed: d3d453401a9b4244b14e6e833d0aeb9da2caea41
review output HEAD: 5ab89f8cc21c120afa439a89f35983cf072843b1
post-adjudication candidate presented for final ratification: f4007eca5406ff71a944f5f980cdf3007bfc7504
post-adjudication Documentation: 32648955774 — SUCCESS
operator ratification Evidence: docs/work/current/operator-ratification.md
```

The operator explicitly ratified the final MR-01 semantic target after independent review, Lead adjudication and post-adjudication validation.

Ratification does **not** make this branch canonical and does not authorize migration or later technical work by implication.

## Durable target owners

```text
docs/decisions/methodology-repository-rebaseline.md
docs/development/planning-readiness.md
docs/development/engineering-rules.md
```

These files remain `PROPOSED` on this non-canonical pre-migration branch by design. Their semantics are operator-ratified as the target, but current `main` authority does not switch until a separately authorized repository migration applies the ratified target, removes temporary work and integrates a coherent final candidate.

Lifecycle distinction:

```text
operator-ratified semantic target
≠ canonical main authority
≠ repository migration authorization
≠ TA-03 authorization
≠ Product implementation authorization
≠ merge authorization
```

## Temporary Evidence / migration material

```text
docs/work/current/proposal.md
docs/work/current/adversarial-review.md
docs/work/current/blueprint-15-amendment.md
docs/work/current/plan.md
docs/work/current/independent-review-adjudication.md
docs/work/current/operator-ratification.md
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
STOP
→ await explicit operator authorization to execute the MR-01 repository migration
```

If migration is authorized, it follows the ratified RM plan on one migration branch and one final merge candidate. It must not partially merge intermediate RM states.

The repository migration must complete before the ratified target becomes canonical and before the next technical stage is considered.

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
