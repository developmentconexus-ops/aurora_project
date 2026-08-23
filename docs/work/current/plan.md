---
id: PLAN-AURORA-MR-01-REPOSITORY-MIGRATION
title: Aurora Methodology and Repository Rebaseline Migration Plan
document_type: temporary_execution_plan
form: reference
authority: design
status: proposed
version: 0.1.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - branch-only proposed execution plan for MR-01 repository migration after final ratification
related:
  - DOC-AURORA-METHODOLOGY-REPOSITORY-REBASELINE
  - DOC-AURORA-PLANNING-READINESS
  - DOC-AURORA-ENGINEERING-RULES
  - DESIGN-AURORA-BLUEPRINT-15-REPOSITORY-AMENDMENT
last_reviewed: 2026-08-23
---

# Aurora Methodology and Repository Rebaseline Migration Plan

> **For agentic workers:** execute this plan only after MR-01 final ratification and an explicit migration/execution authorization. This file is branch-only temporary work and MUST be deleted before a merge candidate or `main` promotion.

**Goal:** Migrate Projeto Aurora from its pre-Repository-Standard documentation/status model to the ratified DevelopmentConexus repository operating envelope while preserving current Product/architecture authority and installing the new cross-system planning/readiness graph.

**Architecture:** Treat the migration as an authority-preserving compiler, not a cleanup. Introduce target routers/owners first, prove reachability and equivalence, then retire old live surfaces only after their current semantic obligations and required provenance have a destination. Keep Product/runtime implementation blocked throughout.

**Tech Stack:** Markdown, existing Python documentation generator/validator, GitHub Actions, Git/PR controls. No Aurora Product runtime/library dependency is introduced.

**Spec:** `docs/decisions/methodology-repository-rebaseline.md` + `docs/development/planning-readiness.md` + `docs/work/current/blueprint-15-amendment.md`.

## Global constraints

- Revalidate exact `main`, branch, PR and CI before every migration gate.
- No Aurora Product/runtime code, schema, dependency or deployment implementation.
- No M0 R7 continuation/Verdict/R8.
- No TA-03+ execution by implication.
- TA-01 and TA-02 remain canonical unless a concrete contradiction is proven.
- One current semantic meaning has one owner at every transition point.
- No current live file is deleted before surviving semantics and required provenance have a proven destination.
- `docs/work/**` and all temporary review artifacts are removed before final merge candidate.
- Permanent `docs/superpowers/**` is eliminated only after current semantics are absorbed or proven historical.
- The target fresh-actor pack remains `AGENTS.md → docs/index.md → docs/roadmap.md → 1–2 owners`, normally five files or fewer total.
- `docs/roadmap.md` becomes sole mutable stage/status/allowed-work/next-action authority only when its replacement coverage and validators are ready in the same coherent candidate.
- Historical closed-stage wording may retain `MNFS`; current authoritative terminology is `Conexus OS`.
- Final migration requires independent challenge and explicit operator merge authorization.

---

## Gate decomposition

The migration is one program, but not one all-at-once destructive edit. Each `RM-*` is a coherent reviewable gate/work unit; adjacent gates may share one PR only when their authority/proof boundaries remain explicit and the operator authorizes the combined scope.

```text
RM-01 Bootstrap + Status Authority
RM-02 Constitutional/Method Reconciliation
RM-03 Planning/Readiness Authority
RM-04 Decision/ADR Reconciliation
RM-05 Architecture/Capability Routing
RM-06 Evidence/History/Tracking Census
RM-07 Legacy Live-Surface Retirement
RM-08 Generator/Projection Separation
RM-09 Validator/CI Conformance
RM-10 Git/Branch Protection
RM-11 Fresh-Actor + Global Coherence Proof
RM-12 Independent Review + Final Promotion
```

---

### RM-01 — Bootstrap + Status Authority

**Outcome:** Introduce the target fresh-actor route and a hand-maintained repository roadmap without yet deleting old status surfaces.

**Files:**
- Modify: `README.md`
- Modify: `AGENTS.md`
- Create: `docs/index.md`
- Replace semantics: `docs/roadmap.md` from generated Product projection to hand-maintained program authority
- Preserve temporarily: `docs/tracking/STATUS.md` as explicitly superseded transition source until RM-06/RM-07
- Modify later in same gate if needed: `CONTRIBUTING.md`

**Required content:**

`README.md`:
```text
landing-only Product introduction
→ links to AGENTS.md and docs/index.md
→ no mutable stage/status/next action
```

`AGENTS.md`:
```text
fresh route
organizational Method + Repository Standard refs
Aurora hard stops
current verification entrypoint
Git/review rules
no Product/roadmap prose duplication
```

`docs/index.md` task routes must include at minimum:

```text
current stage/permission → docs/roadmap.md
Product/North Star → docs/product/README.md or exact Blueprint owner
module/runtime ownership → accepted architecture owner
planning/readiness → docs/development/planning-readiness.md
current decisions → docs/decisions/index.md once RM-04 lands
capability-specific work → exact CAP owner
research → exact routed research only
frozen M0 candidate → exact phase/evidence route, never default
```

`docs/roadmap.md` must identify:

```text
MR-01 current state
TA-01/TA-02 canonical
TA-03 blocked until migration closure
M0 R7 frozen/non-canonical
implementation blocked
exact next authorized action
```

**Negative controls:**
- remove/rename `docs/index.md` in a controlled validation fixture → repository gate must fail;
- inject a second mutable `CURRENT STAGE` owner into a forbidden status fixture → status-authority guard must fail;
- exceed the bootstrap 20 KiB budget in a fixture → guard must fail.

**Exit:** A fresh actor can find exact current state and task owner without opening `docs/tracking/STATUS.md` directly; old STATUS remains only as transition source and is explicitly non-current.

---

### RM-02 — Constitutional / Method Reconciliation

**Outcome:** Apply the bounded Blueprint 15 amendment and refine ACRM scope without changing unrelated Aurora Product meaning.

**Files:**
- Modify: `docs/product/blueprint/15-documentation-research-governance.md`
- Modify: `docs/product/CAPABILITY-REALIZATION-METHOD.md`
- Modify: `docs/product/blueprint/07-harness-orchestration.md` for current `Conexus OS (historically MNFS)` terminology
- Modify only where forward-looking semantics require it: `docs/product/blueprint/14-capability-roadmap.md`
- Regenerate: `docs/product/PRODUCT-BLUEPRINT.md`
- Do NOT regenerate repository `docs/roadmap.md` from Blueprint 14 after RM-08 separation

**Constitutional deltas:** exactly those enumerated in `docs/work/current/blueprint-15-amendment.md`.

**ACRM scope delta:**

```text
ACRM = capability/milestone/slice realization lifecycle
repository operation = external Repository Standard + local engineering rules
global cross-system readiness = docs/development/planning-readiness.md
```

R0–R8 names and already-recorded M0 Evidence remain historically valid.

**Negative controls:**
- a validator fixture claiming STATUS is current after roadmap migration must fail;
- a fixture that makes ACRM override current Product/global architecture must fail a routing/authority check.

**Exit:** No accepted constitutional text requires the retired repository model.

---

### RM-03 — Planning / Readiness Authority

**Outcome:** Promote the accepted MR-01 planning graph into one durable standard and supersede the old TA-03+ ordering without rewriting TA-01/TA-02.

**Files:**
- Finalize: `docs/development/planning-readiness.md`
- Finalize/rehome as needed: `docs/decisions/methodology-repository-rebaseline.md`
- Refine or supersede: `docs/design/AURORA-TECHNICAL-ARCHITECTURE-BASELINE-MAP.md`
- Preserve: `docs/design/AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY.md` semantics
- Create/update architecture index in RM-05 for final routing

**Required decision mapping:**

```text
old TA-01 → PRESERVE / new TA-01
old TA-02 → PRESERVE / new TA-02
old TA-03 repository/source/build → SUPERSEDED by new TA-09
old TA-04 contracts → REFINED/SPLIT into new TA-03/TA-04
old TA-05 data → new TA-06
old TA-06 identity/security → new TA-07
old TA-07 cognition/Harnesses → new TA-08
old TA-08 operations → new TA-10
new TA-05, TA-11, TA-12, TA-13 → added because current cross-project Evidence proves missing implementation-readiness owners
TA-TX → conditional only
```

No stage status is inferred from numbering alone.

**Exit:** A new planner can identify every remaining major technical decision class and its prerequisite owner without reading the superseded baseline map as current.

---

### RM-04 — Decision / ADR Reconciliation

**Outcome:** Create one current decision-disposition route and rehome ADR authority without losing current/open/reopen obligations.

**Files:**
- Create/finalize: `docs/decisions/index.md`
- Rehome: `docs/adr/README.md` semantics into the decision index/ADR child router
- Rehome current ADR files: `docs/adr/0001-*.md` … `0009-*.md` → `docs/decisions/adr/`
- Update inbound references mechanically
- Retire after coverage proof: `docs/tracking/DECISIONS.md`

**Decision vocabulary target:**

```text
CURRENT
PRESERVE
REFINED
REOPEN
DEFERRED
SUPERSEDED
REJECTED
```

The index must preserve exact scope for M0-local ADRs and prevent global inheritance by implication.

**Historical-name correction:** D-033/current software-Harness entries use Conexus OS; historical provenance may mention MNFS.

**Negative control:** a deliberately unresolved current decision ID/path in a fixture must fail routing validation.

**Exit:** Fresh actors discover current decision disposition without reconstructing chronology.

---

### RM-05 — Architecture / Capability Routing

**Outcome:** Route current structural architecture and Capability authorities through semantic owners, without mechanically flattening everything into one directory.

**Files:**
- Create: `docs/architecture/index.md`
- Rehome current structural authorities from `docs/design/` only when their current meaning is architecture authority
- Keep accepted `docs/capabilities/CAP-*` as Aurora-specific specialization and route it from `docs/index.md`
- Classify old M0 microdesign/implementation-plan artifacts as phase/history/evidence rather than current global architecture
- Update current links/IDs after each move

**Rules:**

```text
current architecture → docs/architecture/**
material current decision → docs/decisions/**
capability reusable behavior → docs/capabilities/**
closed-stage result/closure → docs/phases/** when useful
proof/receipt → docs/evidence/**
historical origin only → Git/history route, not default current surface
```

Do not rename every file for aesthetic consistency if the current semantic identity/path is already clear and a move has no real navigation benefit.

**Exit:** `docs/index.md` can route Product, architecture, decision and Capability questions without `DOCUMENTATION-MAP.md` or recursive reading.

---

### RM-06 — Evidence / History / Tracking Census

**Outcome:** Build a semantic census of permanent tracking/review/acceptance/history surfaces before retirement.

**Inputs:**

```text
docs/tracking/STATUS.md
docs/tracking/WORKLOG.md
docs/tracking/BACKLOG.md
docs/tracking/DOCUMENTATION-COVERAGE.md
docs/acceptance/**
docs/reviews/**
docs/history/**
docs/superpowers/**
```

For every live file/class classify:

```text
CURRENT SEMANTIC OBLIGATION → target owner
DURABLE EVIDENCE WITH CURRENT CONSUMER → docs/evidence or exact phase owner
HISTORICAL / GIT SUFFICIENT → remove from live tree after reachability check
UNMERGED UNIQUE PROVENANCE STILL REQUIRED → durable ref required before deletion
TEMPORARY / DUPLICATE → delete from merge candidate
```

**Special requirement:** preserve exact M0 frozen R7 branch/ref while it remains a named Evidence consumer.

**Exit:** zero old live file is unclassified before RM-07 deletion/rehome.

---

### RM-07 — Legacy Live-Surface Retirement

**Outcome:** Remove duplicate/obsolete live routing/status/planning surfaces after RM-01–RM-06 prove replacements.

**Candidate removals after proof:**

```text
docs/DOCUMENTATION-MAP.md
docs/tracking/STATUS.md
docs/tracking/WORKLOG.md
docs/tracking/BACKLOG.md
docs/tracking/DOCUMENTATION-COVERAGE.md
docs/superpowers/**
obsolete permanent review/acceptance chronology whose current semantics are consolidated
```

`docs/history/**` is retained only when a current historical consumer justifies a live file; otherwise Git history may suffice.

**Guard:** before each deletion, search/update inbound current routes and confirm any required unique unmerged provenance remains reachable.

**Exit:** no parallel mutable current-state or permanent session/work artifact remains.

---

### RM-08 — Generator / Projection Separation

**Outcome:** Stop generating repository roadmap from Product Blueprint while retaining generated Product Blueprint publication.

**Files:**
- Modify: `scripts/generate_docs.py`
- Modify: `scripts/validate_docs.py`
- Modify: `.github/workflows/docs.yml`
- Regenerate: `docs/product/PRODUCT-BLUEPRINT.md`
- Preserve hand-maintained: `docs/roadmap.md`

**Generator target:**

```text
GENERATE:
docs/product/PRODUCT-BLUEPRINT.md

DO NOT GENERATE:
docs/roadmap.md
```

**Test / negative controls:**

1. change Blueprint source without regenerating Product aggregate → fail;
2. change `docs/roadmap.md` without changing Blueprint 14 → allowed when roadmap structure/rules remain valid;
3. generator output must never overwrite roadmap;
4. stale Product aggregate still fails.

**Exit:** Product capability roadmap and mutable repository roadmap are mechanically separate authorities.

---

### RM-09 — Validator / CI Repository-Standard Conformance

**Outcome:** Replace old documentation-presence assumptions with target repository-standard controls while retaining useful Aurora-specific validation.

**Files:**
- Modify: `scripts/validate_docs.py`
- Modify: `.github/workflows/docs.yml`
- Add small fixture/negative-control mechanism only if required to prove guards fire

**Retain where current consumer exists:**

```text
Product Blueprint source count/order/freshness
unique material document IDs where still used
research source-manifest integrity
constitutional requirement identity/coverage where still current
local links/current router reachability
```

**Add/enforce:**

```text
bootstrap <= 20 KiB
roadmap sole mutable status/next action
README landing-only
default task route <= 5 files
no docs/work in main/merge candidate
no docs/superpowers in main/merge candidate
no permanent session/review handoff artifacts
no duplicate roadmap/status surface
current decision dispositions valid
required provenance refs preserved when declared
base...candidate diff check
material repository guards demonstrably fire
```

**Verification command target:** expose one stable local command/script that CI uses identically. The exact command may remain Python-based; do not add npm or another package manager solely for uniformity.

**Exit:** aggregate documentation/repository gate passes positive candidate and each material negative control proves failure when its protected property is violated.

---

### RM-10 — Git / Branch Protection

**Outcome:** Align platform enforcement to Repository Standard instead of relying only on behavioral convention.

**Target GitHub settings:**

```text
main deletion forbidden
main force-push forbidden
PR-based changes required
at least one required aggregate status check
normal merge method = squash
head branch deletion after merge when appropriate
```

Do not rename a functioning required check solely for aesthetics. Choose/retain the exact aggregate name only after the migrated CI status is stable.

**Proof:** query GitHub protection/settings after configuration and record exact current Evidence.

**Exit:** a direct/unchecked path cannot silently bypass the protected `main` integration contract.

---

### RM-11 — Fresh-Actor + Global Coherence Proof

**Outcome:** Prove the repository can be resumed without conversation archaeology and the migration did not create duplicate/missing authority.

**Fresh-actor proof:**

Starting from repository only, a fresh reviewer must correctly state:

```text
what Aurora is
current gate and implementation prohibition
TA-01/TA-02 accepted status
current next stage after MR-01
where current decisions live
where Product roadmap lives versus repository roadmap
where temporary work may exist
what Conexus OS is relative to Aurora
where M0 frozen Evidence lives
exact next action
```

The reviewer should reach the answer through the target default pack, not by reading the repository recursively.

**Global Coherence Review attacks:**

```text
duplicate authority
missing owner
circular routing
stage duplication with ACRM
stale STATUS/roadmap references
legacy MNFS current-name leakage
current semantic obligation lost in cleanup
generated projection treated as authority
old TA ordering surviving as current
permanent temporary-work surface
repository standard copied locally as second authority
```

**Exit:** no unresolved material finding.

---

### RM-12 — Independent Review + Final Promotion

**Outcome:** Independent challenger reviews the exact merge candidate, Lead adjudicates findings, operator ratifies final repository/methodology target, and merge remains a separate explicit authorization.

**Review isolation:** use the canonical Repository Standard/Fable flow:

```text
exact candidate branch/head
→ review/<gate>-fable
→ only docs/work/current/ai-dialog.md differs
→ reviewer output = Evidence
→ Lead adjudication/corrections return to candidate
→ review branch never merges
```

A second round occurs only if material corrections invalidate prior coverage.

Before promotion:

```text
docs/work/** = absent
docs/superpowers/** = absent if ratified migration requires retirement
all current routes valid
all required CI green on exact candidate
independent material findings = 0 unresolved
operator ratification = explicit
merge authorization = explicit and separate
```

**Exit:** MR-01/RM migration is integrated and `docs/roadmap.md` points to the first newly authorized planning gate; TA-03 remains not started unless separately authorized.

---

## Plan self-review

### Coverage

The plan covers:

- Method/Repository Standard adoption boundary;
- Blueprint 15 constitutional amendment;
- ACRM scope refinement;
- new TA planning graph;
- repository bootstrap/status/router model;
- decisions/ADR routing;
- architecture/Capability routing;
- safe evidence/history/tracking cleanup;
- Product aggregate vs repository roadmap generation split;
- validator/CI negative controls;
- Git platform enforcement;
- fresh-session/global coherence proof;
- independent review/operator/merge gates;
- Conexus OS current terminology;
- frozen M0 evidence preservation.

### Explicitly deferred

This migration does not select or execute any Aurora Product stack, TA-03 operation semantics, API binding, data store, IAM product, Mastra integration, model/provider or runtime deployment.

### Completion definition

The repository rebaseline is complete only when the target operating model is integrated on `main`, all temporary work is absent, the repository-standard guards are proven to fire, a fresh actor can route correctly, and the roadmap still blocks Product implementation until later readiness gates close.
