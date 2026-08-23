---
id: DOC-AURORA-MR-01-MIGRATION-WORK
title: MR-01 Repository Migration Work Gate
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

# MR-01 Repository Migration — Current Work

> **BRANCH-ONLY / NON-AUTHORITATIVE.** This directory must be removed before the final merge candidate.

## Exact authority

```text
canonical main at migration start: 35614c581cea32e04305c1ad63522fee151eb283
MR-01 semantic target ratified from: f4007eca5406ff71a944f5f980cdf3007bfc7504
ratification Evidence: PR #6 + docs/work/current/operator-ratification.md on ratification branch
migration authorization: EXPLICIT — operator, 2026-08-23
migration branch: docs/mr-01-repository-migration-20260823
```

## Current gate

```text
RM-01 — CURRENT SEMANTIC + PROVENANCE CENSUS
state: IN PROGRESS
Product/runtime implementation: BLOCKED
TA-03+: NOT AUTHORIZED
M0 R7/R8: NOT AUTHORIZED
Architecture Spike execution: NOT AUTHORIZED
merge: NOT AUTHORIZED
```

## Migration law

The migration is executed on this one branch and will produce one final merge candidate. RM intermediate states do not merge to `main`.

```text
RM-01 census
→ RM-02 constitutional/method reconciliation
→ RM-03 target authority preparation
→ RM-04 decision/architecture/capability reconciliation
→ RM-05 atomic repository control-plane cutover
→ RM-06 legacy live-surface retirement
→ RM-07 verification / negative controls
→ RM-08 Git protection proof/configuration
→ RM-09 fresh-actor + global coherence proof
→ RM-10 independent review + final promotion candidate
```

## Required final property

```text
README
→ AGENTS
→ docs/index
→ docs/roadmap
→ 1–2 owning documents
```

Before final promotion:

- `docs/work/**` absent;
- `docs/superpowers/**` absent;
- one mutable repository-program status owner only;
- current semantics/provenance reachable;
- all material guards demonstrated capable of firing;
- independent review complete;
- merge remains separately authorized.
