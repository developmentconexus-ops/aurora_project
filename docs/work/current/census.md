---
id: DOC-AURORA-MR-01-MIGRATION-CENSUS
title: MR-01 Current Semantic and Provenance Census
document_type: temporary_migration_census
form: reference
authority: evidence
status: current
version: 0.1.0
owners:
  - developmentconexus-ops
related: []
last_reviewed: 2026-08-23
---

# MR-01 — Current Semantic + Provenance Census

> **BRANCH-ONLY EVIDENCE.** This census is a migration instrument, not Product/architecture authority. It must be removed before final merge.

## 1. Classification law

Every current live surface is classified as one of:

```text
CURRENT AUTHORITY
CURRENT ROUTER / PROGRAM STATE
DURABLE EVIDENCE WITH CURRENT CONSUMER
CURRENT FORWARD OBLIGATION
HISTORICAL / GIT SUFFICIENT
UNMERGED UNIQUE PROVENANCE STILL REQUIRED
TEMPORARY / DUPLICATE
```

No deletion is authorized merely by this census. A source retires only after its surviving semantics, inbound routes and required provenance have a destination.

## 2. Root / repository control plane

| Current surface | Classification | Target |
|---|---|---|
| `README.md` | CURRENT ROUTER with duplicated Product/status prose | rewrite as landing-only README |
| `AGENTS.md` | CURRENT ROUTER, old operating model | rewrite as compact bootstrap to `docs/index.md` + `docs/roadmap.md` |
| `CONTRIBUTING.md` | GUIDANCE with old read/status route | retain only as short contributor guidance; route through new bootstrap |
| `.github/workflows/docs.yml` | CURRENT CONTROL, old generator/validator assumptions | replace during atomic RM-05 cutover |
| `scripts/generate_docs.py` | CURRENT CONTROL, generates Product aggregate + repository roadmap | retain Product aggregate generation; remove roadmap generation in RM-05 |
| `scripts/validate_docs.py` | CURRENT CONTROL, old tree/status assumptions | realign in RM-05/RM-07; preserve useful Product/research/requirement checks |

## 3. Product authority

| Current surface | Classification | Target |
|---|---|---|
| `docs/product/blueprint/01..14` | CURRENT CONSTITUTIONAL AUTHORITY | preserve; only bounded terminology/path updates where required |
| `docs/product/blueprint/15-documentation-research-governance.md` | CURRENT CONSTITUTIONAL AUTHORITY, old repository model in scope | apply ratified bounded MR-01 amendment |
| `docs/product/PRODUCT-BLUEPRINT.md` | GENERATED PROJECTION with real consumer | preserve as generated read-only publication |
| `docs/product/README.md` | CURRENT PRODUCT ROUTER | preserve/update links to target repository model |
| `docs/product/REQUIREMENTS-TRACEABILITY.md` | CURRENT PRODUCT/TRACEABILITY REFERENCE | preserve |
| `docs/product/CAPABILITY-REALIZATION-METHOD.md` | CURRENT METHOD, wrong semantic home | rehome with stable ID to `docs/development/capability-realization.md`; no duplicate current copy |
| `docs/roadmap.md` | GENERATED PRODUCT projection occupying repository-roadmap path | separate: Blueprint 14 remains Product roadmap; `docs/roadmap.md` becomes hand-maintained repository-program authority at RM-05 |

## 4. Decisions

| Current surface | Classification | Target |
|---|---|---|
| `docs/adr/0001..0009` | CURRENT DECISION AUTHORITY | rehome exact documents/IDs under `docs/decisions/adr/` |
| `docs/adr/README.md` | CURRENT ADR ROUTER | absorb into `docs/decisions/index.md` |
| `docs/tracking/DECISIONS.md` | CURRENT FORWARD OBLIGATION + decision discovery | consolidate every accepted/open/reopen/deferred disposition into `docs/decisions/index.md`; then retire |
| `docs/tracking/BACKLOG.md` | historical/non-commitment A0 candidate set | current relevant open decisions/research triggers already route through Blueprint 14, decision landscape and new decision register; retire after explicit open-item comparison |

Preserve current non-global scope: ADR-0003..0008 remain M0-scoped; ADR-0009 remains preferred-first-to-evaluate, not global Mastra ownership.

## 5. Current architecture

### Preserve/rehome as current architecture

```text
docs/design/AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY.md
→ docs/architecture/module-runtime-topology.md

AURORA System Architecture accepted structural meaning / decision landscape
→ docs/architecture/system-architecture.md and/or routed current reference

SAR-A1 Stage-A availability/activation accepted constraints
→ docs/architecture/stage-a-presence-activation.md

ARCHITECTURE-SPIKES.md current portfolio/triggers
→ docs/reference/architecture-spikes.md or exact decision route
```

### Superseded/reconciled

```text
docs/design/AURORA-TECHNICAL-ARCHITECTURE-BASELINE-MAP.md
→ old TA-03+ ordering is superseded by ratified Planning Readiness;
  preserve TA-01/TA-02 history in Git/phase Evidence, not as current router.
```

### M0-specific design/evidence

```text
docs/design/M0-R4-DECISION-LANDSCAPE.md
docs/design/M0-R4-MASTRA-FIT-MATRIX.md
docs/design/M0-R6-SOVEREIGN-CORE-IMPLEMENTATION-PLAN.md
docs/design/M0-R6-SOVEREIGN-CORE-MICRODESIGN.md
→ docs/phases/m0/** or docs/evidence/m0/**
```

These remain evidence for the frozen M0 candidate and must not become current cross-system target architecture by existence.

## 6. Capabilities

`docs/capabilities/**` is a justified Aurora-specific durable specialization with an active ACRM consumer. Preserve it. Update inbound/outbound routes only.

## 7. Acceptance / review Evidence

### Current rule

Acceptance and review outputs are Evidence, not Product/architecture authority. Their stable document IDs and exact bytes remain useful because accepted Blueprint/ADR/Capability documents reference them and because M0 is frozen rather than erased.

### Target

```text
docs/acceptance/** → docs/evidence/acceptance/**
docs/reviews/**    → docs/evidence/reviews/**
```

Move exact blobs/IDs where current consumers exist. Git remains provenance for ordinary chronology, but the live target retains evidence needed for current accepted/frozen claims.

Explicit high-value current consumers include:

- A0 operator acceptance and fresh-session Golden Proof;
- M0 selection, R0–R6 authorizations/acceptances and spike receipts;
- System Architecture Rebaseline acceptance/closeout;
- Stage-A Presence/activation acceptances;
- Technical Architecture Baseline acceptance/closeout;
- TA-01/TA-02 operator acceptance/merge closeout and adversarial review.

## 8. Tracking

| Surface | Classification | Target |
|---|---|---|
| `docs/tracking/STATUS.md` | CURRENT PROGRAM STATE until RM-05 | semantic facts compiled into new `docs/roadmap.md`; retire atomically |
| `docs/tracking/WORKLOG.md` | chronological history with occasional pointers | Git/merged PR history is sufficient after current pointers/obligations are mapped; retire |
| `docs/tracking/DECISIONS.md` | current decision discovery/forward obligations | `docs/decisions/index.md`; retire |
| `docs/tracking/BACKLOG.md` | non-commitment historical candidate ideas | current relevant triggers route to Product/decisions/research; retire |
| `docs/tracking/DOCUMENTATION-COVERAGE.md` | A0 discovery coverage proof | preserve only as A0 Evidence if still cited; otherwise Git sufficient after Blueprint completeness is retained |

## 9. Documentation Map

`docs/DOCUMENTATION-MAP.md` is the old combined router/authority map. Its surviving responsibilities split into:

```text
task/intention routing             → docs/index.md
current stage/status/next action   → docs/roadmap.md
current decision discovery         → docs/decisions/index.md
architecture routing               → docs/architecture/index.md
repository rules                   → docs/development/engineering-rules.md
Product authority                  → docs/product/** unchanged
```

After new routes are proven, the old map retires.

## 10. `docs/superpowers/**`

Current population:

```text
plans/
  2026-08-05-a0-documentation-remediation.md
  2026-08-12-aurora-system-architecture-rebaseline.md
  2026-08-12-aurora-technical-architecture-baseline.md

specs/
  2026-08-05-aurora-capability-harness-architecture-design.md
  2026-08-12-aurora-system-architecture-rebaseline-design.md
```

Disposition:

- accepted/current semantic results already exist in Product Blueprint, architecture owners, decisions and acceptance/review Evidence;
- no `docs/superpowers/**` file remains a required current authority after those target owners are established;
- previously merged content remains reachable in Git history;
- therefore retire the entire live tree in RM-06 after inbound references are updated.

## 11. Research and history

`docs/research/**` remains opt-in Evidence with named decision consumers and source manifests. Preserve.

`docs/history/**` is not default-read. Preserve only files with a current explicit historical consumer (notably Aurora origin/discovery) when that improves traceability; Git is sufficient for ordinary superseded chronology.

## 12. Frozen M0 R7 provenance

The frozen M0 R7 implementation branch remains named Evidence and is not canonical. Its exact ref must stay reachable while current Aurora documentation refers to it. Migration must not delete or normalize it into current source architecture.

## 13. RM-01 exit

The census finds no class that needs deletion before a destination exists. Target owners are defined for every current semantic class.

```text
RM-01 RESULT: PASS FOR RM-02
DELETIONS AUTHORIZED BY RM-01 ALONE: NONE
Product implementation: BLOCKED
TA-03+: NOT AUTHORIZED
merge: NOT AUTHORIZED
```
