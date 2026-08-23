---
id: DOC-AURORA-MR-01-MIGRATION-CENSUS
title: MR-01 Semantic and Provenance Census Evidence
document_type: repository_migration_census_evidence
form: reference
authority: evidence
status: accepted
version: 1.0.0
owners:
  - developmentconexus-ops
related:
  - DOC-AURORA-MR-01-OPERATOR-RATIFICATION
  - DOC-AURORA-MR-01-MIGRATION-AUTHORIZATION
  - DOC-AURORA-DOCUMENTATION-MAP
  - DOC-AURORA-REPOSITORY-ROADMAP
last_reviewed: 2026-08-23
---

# MR-01 — Semantic and Provenance Census Evidence

RM-01 classified every live pre-migration documentation/control-plane surface before destructive rehoming. The full branch-working census was recorded before the cutover; this durable summary preserves its accepted outcome while Git/PR #8 retains the exact working artifact history.

## Classification model

```text
CURRENT AUTHORITY
CURRENT ROUTER / PROGRAM STATE
DURABLE EVIDENCE WITH CURRENT CONSUMER
CURRENT FORWARD OBLIGATION
HISTORICAL / GIT SUFFICIENT
UNMERGED UNIQUE PROVENANCE STILL REQUIRED
TEMPORARY / DUPLICATE
```

No source was retired solely because its old path violated the target layout. Surviving meaning, inbound routing and provenance were assigned first.

## Durable dispositions

| Pre-MR-01 surface | Durable disposition |
|---|---|
| `README.md` | landing-only root |
| `AGENTS.md` | compact bootstrap |
| `docs/DOCUMENTATION-MAP.md` | stable identity preserved as `docs/index.md` task/intention router |
| generated `docs/roadmap.md` Product projection | Product sequence remains Blueprint 14; old projection identity preserved as phase snapshot; path becomes hand-maintained repository roadmap |
| `docs/product/blueprint/01..14` | preserved constitutional Product owners |
| Blueprint 15 | bounded repository-governance refinement under MR-01 |
| Product aggregate | preserved generated read-only projection |
| Product ACRM file | stable ID rehomed to `docs/development/capability-realization.md` |
| `docs/adr/**` | exact decision docs/IDs rehomed under `docs/decisions/adr/**` |
| tracking `DECISIONS` | current dispositions/forward obligations consolidated into `docs/decisions/index.md` |
| tracking `STATUS` | archived pre-MR-01 snapshot; current authority moved atomically to `docs/roadmap.md` |
| `WORKLOG` | archived Evidence; Git/PR history owns chronology |
| A0 documentation coverage/backlog | Evidence snapshots where useful, not current program owners |
| current TA-01/TA-02 / system architecture / Stage-A designs | rehomed under `docs/architecture/**` |
| old Technical Architecture ordering | superseded phase snapshot; TA-01/TA-02 results preserved |
| M0 R4/R6/spike designs | `docs/phases/m0/**` Evidence/history, not global target architecture |
| `docs/acceptance/**`, `docs/reviews/**` | exact IDs/blobs rehomed under `docs/evidence/**` where current accepted/frozen claims consume them |
| `docs/superpowers/**` | current semantic specs absorbed into architecture; still-referenced plan IDs preserved as concise phase snapshots; remaining chronology retired to Git |
| `docs/research/**` | preserved opt-in Evidence with source manifests |
| frozen M0 R7 branch | exact non-canonical ref preserved; migration does not normalize it into target source architecture |

## Provenance exception discovered during RM-07

Validation demonstrated that two historical plan IDs and the old generated Product-roadmap ID still had current Evidence consumers. Rather than weaken ID validation, MR-01 retained semantic snapshots under `docs/phases/**` while the original full bytes remain reachable in the pre-migration Git lineage.

This is evidence that the census/reachability rule was actually enforced rather than treated as a cleanup checklist.

## Result

```text
RM-01: PASS
current semantic class without target owner: 0
required current provenance knowingly discarded: 0
M0 R7 exact-ref preservation: retained
Product/runtime implementation authority created: no
TA-03 authority created: no
```
