---
id: DOC-AURORA-WORK-METHODOLOGY-REBASELINE-INDEX
title: Aurora Methodology and Repository Rebaseline Current Work
document_type: temporary_work_index
form: reference
authority: tracking
status: current
version: 0.3.0
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
durable candidate: PREPARED / PROPOSED
Lead adversarial challenge: COMPLETE / FIXED DURABLE CANDIDATE REVIEWED
independent challenge: NEXT / REQUIRED BEFORE FINAL RATIFICATION
repository migration execution: NOT AUTHORIZED
TA-03+: HOLD
Aurora implementation: BLOCKED
M0 R7: FROZEN / NON-CANONICAL
merge: NOT AUTHORIZED
```

Base and candidate:

```text
repository: developmentconexus-ops/aurora_project
base main: 35614c581cea32e04305c1ad63522fee151eb283
branch: docs/methodology-repository-rebaseline-20260823
Draft PR: #6
Lead semantic review target: eb412408edc420d69c38e22ba4070773b0c26284
Lead review commit: 2956be55308f17634d53e24e42458f26c62f8166
```

The independent-review branch MUST be cut from the then-current candidate head after this index/status pin and must differ only by `docs/work/current/ai-dialog.md`.

## Durable candidate owners

```text
docs/decisions/methodology-repository-rebaseline.md
docs/development/planning-readiness.md
docs/development/engineering-rules.md
```

All three remain `PROPOSED`; design-direction approval is not final constitutional ratification.

## Temporary review / migration artifacts

```text
docs/work/current/proposal.md
docs/work/current/adversarial-review.md v0.3.0
docs/work/current/blueprint-15-amendment.md
docs/work/current/plan.md v0.2.0
```

They MUST NOT enter the final merge candidate or `main`.

## Independent-review authority pack

MR-01 materially changes both the repository operating envelope and the global implementation-readiness program, so the independent reviewer has a named reason to exceed the ordinary five-file work pack. Keep review bounded to:

1. `AGENTS.md` — current Aurora authority/stop model;
2. this file — exact review target/state;
3. `docs/decisions/methodology-repository-rebaseline.md` — durable decision candidate;
4. `docs/development/planning-readiness.md` — proposed global readiness standard;
5. `docs/work/current/blueprint-15-amendment.md` — exact bounded constitutional reopen;
6. `docs/work/current/plan.md` — authority-preserving migration/proof plan.

Add `docs/development/engineering-rules.md` only for repository-local enforcement questions and `docs/work/current/adversarial-review.md` only to inspect prior Lead findings/adjudication.

External organizational authorities are mandatory comparison inputs, not copied local files:

- `developmentconexus-ops/conexus-methodology/METHOD.md` v1.0.0;
- `developmentconexus-ops/conexus-methodology/REPOSITORY-STANDARD.md` v1.0.0.

MetalDocs, Marketplace Central and Conexus OS remain Evidence/reference only.

## Independent challenge focus

Attack at least:

```text
duplicate/missing authority between Method / Repository Standard / Planning Readiness / ACRM
whether TA-03→TA-13 ordering reflects real dependencies or ceremony
whether TA-03 duplicates Capability Specs
whether TA-09 moves source/build/Paved Road too late
whether Blueprint 15 reopen is genuinely bounded
whether retirement of STATUS/WORKLOG/docs/superpowers can lose semantics/provenance
whether RM-05 atomic control-plane cutover is sufficient
whether target fresh-actor route is adequate for Aurora complexity
whether GENERATED / AURORA-FOUNDATION / MODULE-OWNED creates hidden framework authority
whether Conexus OS terminology refinement changes Product meaning
whether proposed negative controls can falsify claimed repository guards
whether any Product/runtime/stack decision is smuggled into MR-01
```

Reviewer output is Evidence, never authority. Material corrections return to the candidate; a second round is justified only if corrections materially invalidate first-round coverage.

## Hard boundaries

This gate does **not** authorize:

- Aurora Product implementation;
- repository migration execution;
- continuation, merge, Verdict or R8 closeout of the frozen M0 R7 candidate;
- TA-03 or any later technical stage by implication;
- Architecture Spike execution;
- production repository/source restructuring;
- framework/database/authentication/provider selection by convenience;
- frontend Product design work;
- merge.

TA-01 and TA-02 remain accepted/canonical inputs unless independent Evidence proves a concrete material contradiction.
