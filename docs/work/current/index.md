---
id: DOC-AURORA-WORK-METHODOLOGY-REBASELINE-INDEX
title: Aurora Methodology and Repository Rebaseline Current Work
document_type: temporary_work_index
form: reference
authority: tracking
status: current
version: 0.5.0
owners:
  - developmentconexus-ops
related:
  - DOC-AURORA-METHODOLOGY-REPOSITORY-REBASELINE
  - DOC-AURORA-PLANNING-READINESS
  - DOC-AURORA-ENGINEERING-RULES
last_reviewed: 2026-08-23
---

# Aurora methodology + repository rebaseline — current work

> **NON-AUTHORITATIVE / BRANCH-ONLY.** This directory must be absorbed or deleted before a merge candidate is promoted.

## Current gate

```text
MR-01 — Methodology, Planning-Readiness & Repository Rebaseline
operator design-direction decision: APPROVED — 2026-08-23
durable candidate: COMPILED / PROPOSED
Lead adversarial challenge: COMPLETE
independent challenge: READY TO START / REQUIRED BEFORE FINAL RATIFICATION
repository migration execution: NOT AUTHORIZED
TA-03+: HOLD
Aurora implementation: BLOCKED
M0 R7: FROZEN / NON-CANONICAL
merge: NOT AUTHORIZED
```

## Candidate identity

```text
repository: developmentconexus-ops/aurora_project
base main: 35614c581cea32e04305c1ad63522fee151eb283
candidate branch: docs/methodology-repository-rebaseline-20260823
Draft PR: #6
Lead semantic review target: eb412408edc420d69c38e22ba4070773b0c26284
Lead review commit: 2956be55308f17634d53e24e42458f26c62f8166
```

This commit freezes only the **review routing state**. The exact independent-review base is this branch head after its Documentation checks succeed. Once `review/mr-01-fable` is cut, the candidate branch remains frozen until reviewer Evidence is returned or the review is explicitly restarted.

## Durable candidate owners

```text
docs/decisions/methodology-repository-rebaseline.md
docs/development/planning-readiness.md
docs/development/engineering-rules.md
```

All remain `PROPOSED`; design-direction approval is not final constitutional ratification.

## Temporary review / migration artifacts

```text
docs/work/current/proposal.md v0.3.1 — compiled provenance only
docs/work/current/adversarial-review.md v0.3.1
docs/work/current/blueprint-15-amendment.md
docs/work/current/plan.md v0.2.0
```

They MUST NOT enter the final merge candidate or `main`.

## Independent-review authority pack

MR-01 materially changes both repository operation and global implementation-readiness, so independent review has a named reason to exceed the ordinary five-file work pack. Keep it bounded to:

1. `AGENTS.md` — current Aurora authority/stop model;
2. this file — exact review target/state;
3. `docs/decisions/methodology-repository-rebaseline.md` — durable decision candidate;
4. `docs/development/planning-readiness.md` — proposed global readiness standard;
5. `docs/work/current/blueprint-15-amendment.md` — exact bounded constitutional reopen;
6. `docs/work/current/plan.md` — authority-preserving migration/proof plan.

Add `docs/development/engineering-rules.md` only for repository-local enforcement questions and `docs/work/current/adversarial-review.md` only to inspect Lead findings/adjudication.

Mandatory external authorities:

- `developmentconexus-ops/conexus-methodology/METHOD.md` v1.0.0;
- `developmentconexus-ops/conexus-methodology/REPOSITORY-STANDARD.md` v1.0.0.

MetalDocs, Marketplace Central and Conexus OS remain comparison Evidence/reference only.

## Independent challenge focus

Attack at least:

```text
duplicate/missing authority among Method / Repository Standard / Planning Readiness / ACRM
whether TA-03→TA-13 ordering is dependency-driven or ceremonial
whether TA-03 duplicates Capability Specs
whether TA-09 moves production source/build/Paved Road too late
whether Blueprint 15 reopen is genuinely bounded
whether retirement of STATUS/WORKLOG/docs/superpowers can lose semantics/provenance
whether RM-05 atomic control-plane cutover avoids dual/no status authority
whether target fresh-actor route is adequate for Aurora complexity
whether GENERATED / AURORA-FOUNDATION / MODULE-OWNED creates hidden framework authority
whether Conexus OS terminology refinement changes Product meaning
whether negative controls can falsify the repository properties claimed
whether any Product/runtime/stack decision is smuggled into MR-01
```

Reviewer output is Evidence, never authority. Material corrections return to the candidate. A second round is justified only if corrections materially invalidate first-round coverage.

## Hard boundaries

This gate does **not** authorize:

- Aurora Product implementation;
- repository migration execution;
- continuation, merge, Verdict or R8 closeout of the frozen M0 R7 candidate;
- TA-03 or later technical stages by implication;
- Architecture Spike execution;
- production repository/source restructuring;
- framework/database/authentication/provider selection by convenience;
- frontend Product design work;
- merge.

TA-01 and TA-02 remain accepted/canonical inputs unless independent Evidence proves a concrete material contradiction.
