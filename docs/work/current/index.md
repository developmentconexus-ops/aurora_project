---
id: DOC-AURORA-WORK-METHODOLOGY-REBASELINE-INDEX
title: Aurora Methodology and Repository Rebaseline Current Work
document_type: temporary_work_index
form: reference
authority: tracking
status: current
version: 0.1.0
owners:
  - developmentconexus-ops
related: []
last_reviewed: 2026-08-23
---

# Aurora methodology + repository rebaseline — current work

> **NON-AUTHORITATIVE / BRANCH-ONLY.** This directory must be absorbed or deleted before a merge candidate is promoted.

## Current gate

```text
MR-01 — Methodology, Planning-Readiness & Repository Rebaseline
state: DISCOVERY / DESIGN
implementation: BLOCKED
TA-03+: HOLD
M0 R7: FROZEN / NON-CANONICAL
```

Base revalidated before opening this work:

```text
repository: developmentconexus-ops/aurora_project
main: 35614c581cea32e04305c1ad63522fee151eb283
branch: docs/methodology-repository-rebaseline-20260823
```

## Read pack for this gate

1. `AGENTS.md`
2. this file
3. `proposal.md`
4. `adversarial-review.md`
5. external `developmentconexus-ops/conexus-methodology/METHOD.md` v1.0.0 and `REPOSITORY-STANDARD.md` v1.0.0 as the named organizational authorities

The three product repositories are comparison Evidence/reference only:

- `developmentconexus-ops/MetalDocs`
- `developmentconexus-ops/marketplace-central`
- `developmentconexus-ops/conexus-os`

## Gate target

Reconcile Aurora with the current DevelopmentConexus engineering and repository standards and define a complete pre-implementation planning/readiness graph that makes future coding a constrained realization of accepted authority rather than a place where material architecture is invented.

## Hard boundaries

This gate does **not** authorize:

- Aurora Product implementation;
- continuation, merge, Verdict or R8 closeout of the frozen M0 R7 candidate;
- TA-03 or any later technical stage by implication;
- Architecture Spike execution;
- production repository/source restructuring before the rebaseline is accepted;
- framework/database/authentication/provider selection by convenience;
- frontend Product design work.

TA-01 and TA-02 remain accepted/canonical inputs unless this rebaseline finds a concrete material contradiction.