---
id: DOC-AURORA-WORK-METHODOLOGY-REBASELINE-INDEX
title: Aurora Methodology and Repository Rebaseline Current Work
document_type: temporary_work_index
form: reference
authority: tracking
status: current
version: 0.2.0
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
Lead adversarial challenge: COMPLETE / MATERIAL DESIGN FINDINGS REMEDIATED
independent challenge: NEXT / REQUIRED BEFORE FINAL RATIFICATION
repository migration execution: NOT AUTHORIZED
TA-03+: HOLD
Aurora implementation: BLOCKED
M0 R7: FROZEN / NON-CANONICAL
merge: NOT AUTHORIZED
```

Base revalidated before opening this work:

```text
repository: developmentconexus-ops/aurora_project
base main: 35614c581cea32e04305c1ad63522fee151eb283
branch: docs/methodology-repository-rebaseline-20260823
Draft PR: #6
```

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
docs/work/current/adversarial-review.md
docs/work/current/blueprint-15-amendment.md
docs/work/current/plan.md
```

They MUST NOT enter the final merge candidate or `main`.

## Independent-review authority pack

The material reason for exceeding the normal five-file repository pack is that MR-01 simultaneously changes the repository operating envelope and the global implementation-readiness program. Review should still remain bounded to:

1. `AGENTS.md` — current Aurora authority/stop model;
2. this file — exact review target and state;
3. `docs/decisions/methodology-repository-rebaseline.md` — durable decision candidate;
4. `docs/development/planning-readiness.md` — proposed global readiness standard;
5. `docs/work/current/blueprint-15-amendment.md` — exact bounded constitutional reopen;
6. `docs/work/current/plan.md` — authority-preserving repository migration/proof plan.

Add `docs/development/engineering-rules.md` only for repository-local enforcement questions and `docs/work/current/adversarial-review.md` only to inspect prior Lead findings/adjudication.

External organizational authorities are mandatory comparison inputs, not copied local files:

- `developmentconexus-ops/conexus-methodology/METHOD.md` v1.0.0;
- `developmentconexus-ops/conexus-methodology/REPOSITORY-STANDARD.md` v1.0.0.

MetalDocs, Marketplace Central and Conexus OS remain Evidence/reference only.

## Gate target

Reconcile Aurora with the current DevelopmentConexus engineering and repository standards and define a complete pre-implementation planning/readiness graph that makes future coding a constrained realization of accepted authority rather than a place where material architecture is invented.

## Independent challenge focus

Attack at least:

```text
duplicate or missing authority between Method / Repository Standard / Planning Readiness / ACRM
whether TA-03→TA-13 ordering has a real dependency basis or ceremony
whether TA-03 cross-system operations duplicates Capability Specs
whether TA-09 moves repository/source/build too late
whether the Blueprint 15 reopen is truly bounded
whether retiring STATUS/WORKLOG/docs/superpowers can lose current semantics/provenance
whether the target fresh-actor route is sufficient for Aurora complexity
whether GENERATED / AURORA-FOUNDATION / MODULE-OWNED creates hidden framework authority
whether Conexus OS current terminology changes Product meaning
whether migration negative controls actually cover the repository properties claimed
whether any Product/runtime/stack decision is being smuggled into MR-01
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
