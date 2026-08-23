---
id: DESIGN-AURORA-METHODOLOGY-REPOSITORY-REBASELINE-CANDIDATE
title: Aurora Methodology, Planning-Readiness and Repository Rebaseline Candidate
document_type: temporary_design_candidate
form: explanation
authority: tracking
status: proposed
version: 0.3.0
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

> **NON-AUTHORITATIVE / BRANCH-ONLY.** The operator approved the design direction represented by the earlier v0.2 candidate on 2026-08-23. Its current semantic result has now been compiled into the durable `PROPOSED` owners below. This temporary file is retained only as review provenance and MUST be deleted before merge.

## Current durable candidate

```text
decision:
docs/decisions/methodology-repository-rebaseline.md

cross-system planning/readiness standard:
docs/development/planning-readiness.md

repository-local engineering specialization:
docs/development/engineering-rules.md

exact bounded constitutional amendment proposal:
docs/work/current/blueprint-15-amendment.md

repository migration/proof plan:
docs/work/current/plan.md v0.2.0

Lead adversarial review:
docs/work/current/adversarial-review.md v0.3.0
```

## Approved design direction

The operator-approved direction is:

```text
DevelopmentConexus Engineering Method v1
+ DevelopmentConexus Repository Standard v1
+ preserved Aurora Product / TA-01 / TA-02
+ bounded Blueprint 15 repository-governance reopen
+ ACRM preserved and narrowed to capability/slice realization
+ dependency-ordered Aurora implementation-readiness graph
+ repository route: README → AGENTS → docs/index → docs/roadmap → task owners
+ docs/roadmap as sole mutable repository-program owner after atomic migration
+ safe semantic/provenance census before retiring old live surfaces
+ current software-Harness name Conexus OS; historical MNFS provenance preserved
+ implementation remains blocked
```

The detailed earlier exploratory reasoning is superseded by the durable candidate owners and Lead review. Do not use this file as a parallel authority.

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

The next required step is an independent challenge against the fixed durable candidate. Reviewer output is Evidence and must be Lead-adjudicated before final operator ratification.
