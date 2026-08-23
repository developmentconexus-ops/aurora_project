---
id: work-current-methodology-repository-rebaseline
kind: temporary-work
owner: architecture
status: candidate
summary: Branch-only router for the Aurora methodology, planning-readiness and repository-operating rebaseline.
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
4. external `developmentconexus-ops/conexus-methodology/METHOD.md` v1.0.0
5. external `developmentconexus-ops/conexus-methodology/REPOSITORY-STANDARD.md` v1.0.0

Task-specific source projects are Evidence/reference only:

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