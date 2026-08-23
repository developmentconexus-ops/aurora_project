---
id: DESIGN-AURORA-METHODOLOGY-REPOSITORY-REBASELINE-CANDIDATE
title: Aurora Methodology, Planning-Readiness and Repository Rebaseline Candidate
document_type: temporary_design_candidate
form: explanation
authority: tracking
status: proposed
version: 0.3.1
owners:
  - developmentconexus-ops
related:
  - DOC-AURORA-METHODOLOGY-REPOSITORY-REBASELINE
  - DOC-AURORA-PLANNING-READINESS
  - DOC-AURORA-ENGINEERING-RULES
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DESIGN-AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY
  - DOC-AURORA-BLUEPRINT-15
last_reviewed: 2026-08-23
---

# Aurora Methodology, Planning-Readiness & Repository Rebaseline — Compiled Design Record

> **NON-AUTHORITATIVE / BRANCH-ONLY.** The operator approved the design direction represented by the earlier v0.2 candidate on 2026-08-23. Its semantic result has been compiled into the durable `PROPOSED` owners below. This temporary file is review provenance only and MUST be deleted before merge.

## Current durable candidate

```text
decision:
docs/decisions/methodology-repository-rebaseline.md

cross-system planning/readiness standard:
docs/development/planning-readiness.md

repository-local engineering specialization:
docs/development/engineering-rules.md

bounded constitutional amendment proposal:
docs/work/current/blueprint-15-amendment.md

repository migration/proof plan:
docs/work/current/plan.md v0.2.0

Lead adversarial review:
docs/work/current/adversarial-review.md v0.3.0
```

## Approved design direction

```text
DevelopmentConexus Engineering Method v1
+ DevelopmentConexus Repository Standard v1
+ preserve Aurora Product meaning and canonical TA-01 / TA-02
+ bounded Blueprint 15 repository-governance reopen
+ preserve ACRM, narrow it to capability/slice realization
+ dependency-ordered Aurora implementation-readiness graph
+ target repository route: README → AGENTS → docs/index → docs/roadmap → task owners
+ docs/roadmap sole mutable repository-program owner after atomic migration
+ semantic/provenance census before retiring old live surfaces
+ current software-Harness name Conexus OS; historical MNFS provenance retained
+ Aurora Product implementation remains blocked
```

The earlier exploratory prose is superseded by the durable candidate owners and the fixed Lead review. Do not use this file as parallel authority.

## Fixed Lead review provenance

```text
durable semantic target reviewed: eb412408edc420d69c38e22ba4070773b0c26284
Lead review commit:                2956be55308f17634d53e24e42458f26c62f8166
post-review edits:                 temporary status/provenance alignment only
```

No Product/stack/runtime semantic decision was added after the fixed Lead target.

## Hard boundary

Design-direction approval does not authorize:

- final constitutional ratification;
- repository migration execution;
- TA-03 or later TA work;
- M0 R7 continuation/Verdict/R8;
- Architecture Spike execution;
- framework/database/IAM/model/provider selection;
- Product implementation;
- merge.

The next required step is independent challenge against the fixed compiled candidate. Reviewer output is Evidence and must be Lead-adjudicated before final operator ratification.
