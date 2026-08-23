---
id: PLAN-AURORA-MR-01-REPOSITORY-MIGRATION
title: Aurora Methodology and Repository Rebaseline Migration Plan
document_type: temporary_execution_plan
form: reference
authority: design
status: proposed
version: 0.2.0
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

**Architecture:** Treat migration as an authority-preserving compiler, not cleanup. First census every current semantic/provenance obligation; prepare target owners as non-current candidates; then perform one atomic repository-control-plane cutover in which `AGENTS`, `docs/index`, hand-maintained `docs/roadmap`, generator, validator and CI change together. Retire old live surfaces only after replacement coverage and reachability are proven.

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
- The target fresh-actor pack is `AGENTS.md → docs/index.md → docs/roadmap.md → 1–2 owners`, normally five files or fewer total.
- `docs/roadmap.md` does not become mutable authority until the same atomic candidate also removes its generated-projection ownership from generator/validator/CI.
- Historical closed-stage wording may retain `MNFS`; current authoritative terminology is `Conexus OS`.
- Final migration requires independent challenge and explicit operator merge authorization.
- A temporary CI-red intermediate commit is not target Evidence; final candidate must prove all new guards on one exact revision.

---

## Corrected gate decomposition

```text
RM-01 Current Semantic + Provenance Census
RM-02 Constitutional / Method Reconciliation
RM-03 Target Authority Preparation
RM-04 Decision / Architecture / Capability Reconciliation
RM-05 Atomic Repository Control-Plane Cutover
RM-06 Legacy Live-Surface Retirement
RM-07 Verification / Negative-Control Closure
RM-08 Git / Branch Protection
RM-09 Fresh-Actor + Global Coherence Proof
RM-10 Independent Review + Final Promotion
```

Adjacent gates may share a PR only when their owner/proof boundaries remain explicit and the operator authorizes that combined scope. RM-05 is intentionally atomic because splitting it would temporarily create either two status authorities or a roadmap that the old generator still overwrites.

---

### RM-01 — Current Semantic + Provenance Census

**Outcome:** Classify every current live Aurora documentation surface before any destructive move.

**Inputs:**

```text
README.md
AGENTS.md
CONTRIBUTING.md
docs/DOCUMENTATION-MAP.md
docs/roadmap.md
docs/product/**
docs/adr/**
docs/capabilities/**
docs/design/**
docs/research/**
docs/reviews/**
docs/acceptance/**
docs/history/**
docs/tracking/**
docs/superpowers/**
scripts/generate_docs.py
scripts/validate_docs.py
.github/workflows/docs.yml
frozen M0 R7 branch/ref and any named unmerged provenance
```

For every file/class classify:

```text
CURRENT AUTHORITY
CURRENT ROUTER / PROGRAM STATE
DURABLE EVIDENCE WITH CURRENT CONSUMER
CURRENT FORWARD OBLIGATION
HISTORICAL / GIT SUFFICIENT
UNMERGED UNIQUE PROVENANCE STILL REQUIRED
TEMPORARY / DUPLICATE
```

Each non-historical item receives exactly one target owner/path.

**Required special proofs:**

- `docs/tracking/STATUS.md` current facts map completely into target roadmap/decision owners;
- `WORKLOG.md` contains no unique current obligation that would disappear into Git history;
- acceptance/review files still referenced by accepted frontmatter are identified before deletion;
- `docs/superpowers/**` current semantics are mapped before retirement;
- frozen M0 R7 exact ref remains reachable while named as Evidence.

**Exit:** zero current live file/class remains unclassified; no deletion is authorized by this census itself.

---

### RM-02 — Constitutional / Method Reconciliation

**Outcome:** Apply the bounded Blueprint 15 amendment and refine ACRM scope while keeping the repository control plane on the old model until RM-05.

**Files:**
- Modify: `docs/product/blueprint/15-documentation-research-governance.md`
- Modify: `docs/product/CAPABILITY-REALIZATION-METHOD.md` as source before rehome
- Modify: `docs/product/blueprint/07-harness-orchestration.md` for current `Conexus OS (historically MNFS)` terminology
- Modify only where forward-looking semantics require it: `docs/product/blueprint/14-capability-roadmap.md`
- Regenerate under the still-current generator: `docs/product/PRODUCT-BLUEPRINT.md`
- Keep the current generated `docs/roadmap.md` until RM-05 atomic cutover

**Constitutional deltas:** exactly those enumerated in `docs/work/current/blueprint-15-amendment.md`.

**ACRM scope delta:**

```text
ACRM = selected milestone/capability/mission realization lifecycle
repository operating envelope = external Repository Standard + Aurora engineering rules
global cross-system readiness = Aurora Planning and Implementation-Readiness Standard
```

R0–R8 names and already-recorded M0 Evidence remain historically valid.

**Current-path note:** Blueprint/ACRM may describe the *target* roadmap owner before cutover, but the branch work index must explicitly state that `main` authority remains the old model until RM-05 lands atomically.

**Exit:** accepted constitutional target no longer requires permanent STATUS/WORKLOG/docs-superpowers or generated repository-roadmap semantics.

---

### RM-03 — Target Authority Preparation

**Outcome:** Prepare durable target owners without yet making them the repository's current program/status route.

**Files to create/finalize as `PROPOSED` until final promotion:**

```text
docs/development/planning-readiness.md
docs/development/engineering-rules.md
docs/development/capability-realization.md   # rehome of current ACRM, preserving stable doc identity
docs/decisions/methodology-repository-rebaseline.md
docs/decisions/index.md
docs/architecture/index.md
```

**ACRM rehome law:**

- preserve `DOC-AURORA-CAPABILITY-REALIZATION-METHOD` stable identity;
- move current semantic ownership from `docs/product/CAPABILITY-REALIZATION-METHOD.md` to `docs/development/capability-realization.md` only after all current links/routers are updated in the same candidate;
- do not keep two current copies after cutover.

**Planning graph mapping required:**

```text
old TA-01 → PRESERVE / new TA-01
old TA-02 → PRESERVE / new TA-02
old TA-03 repository/source/build → SUPERSEDED by new TA-09
old TA-04 contracts → REFINED/SPLIT into new TA-03/TA-04
old TA-05 data → new TA-06
old TA-06 identity/security → new TA-07
old TA-07 cognition/Harnesses → new TA-08
old TA-08 operations → new TA-10
new TA-05, TA-11, TA-12, TA-13 → new readiness owners justified by cross-project Evidence
TA-TX → conditional only
```

**Exit:** target owners exist and can be reviewed without yet creating parallel current program authority.

---

### RM-04 — Decision / Architecture / Capability Reconciliation

**Outcome:** Prepare semantic routing from current legacy paths to target decision/architecture/capability homes before the repository-control-plane cutover.

#### Decisions / ADRs

Prepare:

```text
docs/decisions/index.md
docs/decisions/adr/0001-*.md ... 0009-*.md
```

Preserve exact ADR IDs/status/scope; M0-local decisions remain M0-local.

Decision dispositions use a compact current vocabulary:

```text
CURRENT
PRESERVE
REFINED
REOPEN
DEFERRED
SUPERSEDED
REJECTED
```

Current software-Harness entries use Conexus OS; historical provenance may retain MNFS.

#### Architecture

Prepare `docs/architecture/index.md` and route current structural authorities.

Classify current `docs/design/**` individually:

```text
current structural architecture → docs/architecture/**
current material decision → docs/decisions/**
closed-stage implementation design / M0 microdesign → docs/phases/** or Evidence/history disposition
spike specification/result → phase/evidence owner
historical-only design → Git/history when no current live consumer
```

Do not move files merely for aesthetics when a stable path has a real consumer and a justified local deviation is smaller.

#### Capabilities

Preserve `docs/capabilities/CAP-*/` as Aurora-specific specialization. Global TA documents must not duplicate internal Capability Spec behavior.

**Exit:** every current decision/architecture/capability meaning has one target route ready for `docs/index.md`.

---

### RM-05 — Atomic Repository Control-Plane Cutover

**Outcome:** Switch the repository from the old bootstrap/status/generated-roadmap model to the Repository Standard target on one coherent exact revision.

This gate MUST change the following together; do not split the status switch from generator/validator ownership.

**Files:**

```text
README.md
AGENTS.md
CONTRIBUTING.md when retained
docs/index.md
docs/roadmap.md
docs/development/engineering-rules.md
docs/development/capability-realization.md
docs/decisions/index.md
docs/architecture/index.md
scripts/generate_docs.py
scripts/validate_docs.py
.github/workflows/docs.yml
all current inbound links needed for the new router/roadmap owner
```

#### `README.md`

Landing-only:

```text
Aurora one-paragraph orientation
links to AGENTS.md + docs/index.md
stable verification/setup entrypoint when useful
no mutable stage/status/architecture authority
```

#### `AGENTS.md`

Compact bootstrap only:

```text
AGENTS → docs/index → docs/roadmap → 1–2 owners
Method/Repository Standard refs
Aurora-specific hard stops
verification command/gate
Git/review rules
```

#### `docs/index.md`

Task/intention router, no mutable status. Minimum routes:

```text
current stage/permission → docs/roadmap.md
Product/North Star → docs/product/README.md / exact Blueprint owner
module/runtime ownership → docs/architecture/index.md → exact owner
planning/readiness → docs/development/planning-readiness.md
current decisions → docs/decisions/index.md
capability work → exact CAP owner
research → exact question-specific research
frozen M0 Evidence → exact phase/evidence route, never default
```

#### `docs/roadmap.md`

Hand-maintained and sole mutable repository-program authority:

```text
MR-01 / repository migration exact state
TA-01/TA-02 canonical
next TA stage blocked until migration closeout and authorization
M0 R7 frozen/non-canonical
implementation blocked
Architecture Spike execution blocked unless separately authorized
exact next action
```

Blueprint 14 remains Product/capability roadmap authority; repository roadmap does not duplicate long-horizon Product content.

#### Generator separation

Modify `scripts/generate_docs.py` so it generates only the Product aggregate(s) with real consumers.

Required result:

```text
GENERATE:
docs/product/PRODUCT-BLUEPRINT.md

NEVER GENERATE/OVERWRITE:
docs/roadmap.md
```

Do not create a new generated Product-roadmap projection without a real consumer; Blueprint 14 is already reachable through Product routing.

#### Validator / CI switch

In the same cutover, remove checks that assume roadmap is generated and add target repository-control properties. The exact final validator still retains useful Aurora-specific controls for Product Blueprint, requirements and research manifests.

**Atomic negative controls:**

1. generator cannot modify `docs/roadmap.md`;
2. stale Product Blueprint projection fails;
3. missing `docs/index.md` fails;
4. bootstrap size >20 KiB fails;
5. duplicate mutable status authority fixture fails;
6. durable authority depending on `docs/work/**` fails;
7. base→candidate whitespace/diff control operates on intended range.

**Transition status law:** On this revision, old `docs/tracking/STATUS.md` is either removed (if RM-01 census proves safe) or unmistakably marked superseded/non-current with no router pointing to it. There must never be two current status authorities.

**Exit:** new fresh-actor route is mechanically current; generator cannot overwrite roadmap; validation recognizes only the new program/status owner.

---

### RM-06 — Legacy Live-Surface Retirement

**Outcome:** Remove old live surfaces whose semantics/provenance were covered by RM-01–RM-05.

Candidate retirements, subject to census proof:

```text
docs/DOCUMENTATION-MAP.md
docs/tracking/STATUS.md
docs/tracking/WORKLOG.md
docs/tracking/BACKLOG.md
docs/tracking/DOCUMENTATION-COVERAGE.md
docs/superpowers/**
old docs/adr/** after decision rehome
old docs/design/** items after architecture/phase/evidence disposition
obsolete permanent review/acceptance chronology whose current semantics are consolidated
```

`docs/history/**` stays live only when a current historical consumer justifies it; otherwise Git history may be sufficient.

**Before each retirement class:**

```text
surviving semantic obligation = mapped
current inbound refs = updated
required acceptance/provenance ref = reachable
unique unmerged provenance consumer = protected by durable ref when needed
replacement route = tested
```

**Exit:** no parallel mutable current-state surface, permanent temporary-work tree or superseded duplicate authority remains.

---

### RM-07 — Verification / Negative-Control Closure

**Outcome:** Prove the migrated documentation/repository guard actually enforces claimed properties.

**Retain when current consumers remain:**

```text
Product Blueprint 01..15 source/order/freshness
stable material document IDs where useful
research source-manifest integrity
constitutional requirement identity/coverage
current local links/router reachability
```

**Required repository-standard controls:**

```text
AGENTS + docs/index + docs/roadmap <= 20 KiB
docs/roadmap sole mutable stage/status/allowed-work/next-action authority
README landing-only
default routed task pack <= 5 files unless named reason
no durable current authority depends on docs/work
no docs/work in merge candidate/main
no docs/superpowers in merge candidate/main
no permanent session/handoff/review-round trees
no duplicate mutable roadmap/status surface
current decision dispositions valid/discoverable
required unique unmerged provenance not deleted
PR diff checks compare intended base→candidate
```

**Guard falsification:** every material guard must have a deterministic negative control or equivalent test demonstrating it fires. Presence-only checks are insufficient for behavioral properties.

**Verification entrypoint:** expose one stable local command/script used by CI. Keep Python if it remains the smallest mechanism; do not add npm/package-management infrastructure solely for uniformity.

**Exit:** positive candidate is green and each material repository guard has demonstrated failure on its negative fixture/control.

---

### RM-08 — Git / Branch Protection

**Outcome:** Align GitHub enforcement to Repository Standard after the aggregate migrated gate is stable.

Target settings:

```text
main deletion forbidden
main force-push forbidden
PR-based change integration
at least one required aggregate status check
normal merge method = squash
automatic head deletion when appropriate
```

Do not rename a functioning required check solely for aesthetics. Choose/retain the exact aggregate check only after RM-07 proves it stable.

**Proof:** query GitHub repository/branch-protection settings and record exact current Evidence.

**Exit:** unchecked direct integration can no longer bypass the intended `main` contract.

---

### RM-09 — Fresh-Actor + Global Coherence Proof

**Outcome:** Prove session resilience and absence of duplicate/missing authority after migration.

A fresh reviewer, using repository only, must correctly state:

```text
what Aurora is
current gate / blocked work / exact next action
TA-01/TA-02 accepted state
where Product roadmap lives vs repository roadmap
where current decisions live
where architecture and Capability owners live
where temporary work may exist
what Conexus OS is relative to Aurora
where M0 frozen Evidence/provenance is routed
that TA-03 and Product implementation are not authorized by migration completion
```

The answer must be reachable through the default task pack rather than recursive repository archaeology.

**Global Coherence Review attacks:**

```text
duplicate/missing authority
circular routing
Planning Readiness vs ACRM duplication
stale STATUS/old-roadmap references
current-name MNFS leakage
lost semantic obligation from cleanup
generated projection becoming authority
old TA ordering surviving as current
permanent temporary-work surface
organizational standards copied locally as second authority
one current path requiring >5 files without a named reason
```

**Exit:** zero unresolved material coherence finding.

---

### RM-10 — Independent Review + Final Promotion

**Outcome:** Independent challenger reviews the exact merge candidate, Lead adjudicates findings, operator ratifies final repository/methodology target, and merge remains separately authorized.

Use canonical isolated review:

```text
exact candidate branch/head
→ review/<gate>-fable
→ only docs/work/current/ai-dialog.md differs from candidate
→ reviewer reconstructs authority first
→ reviewer output = Evidence
→ Lead adjudicates
→ accepted corrections land on candidate
→ review branch never merges
```

A second round occurs only if material corrections invalidate prior review coverage.

Before promotion:

```text
docs/work/** absent from candidate
docs/superpowers/** absent if ratified retirement applies
all current routes valid
all required aggregate CI green on exact candidate
negative controls green as tests (i.e. each expected failure demonstrated)
independent material findings = 0 unresolved
operator final ratification = explicit
merge authorization = explicit and separate
```

**Exit:** migrated MR-01 is integrated; `docs/roadmap.md` names the next allowed action. TA-03 remains `NOT STARTED / NOT AUTHORIZED` unless the operator separately opens it.

---

## Plan self-review

### Dependency correction

The plan intentionally does **not** switch roadmap authority before generator/validator ownership. RM-05 is the atomic control-plane cutover that changes all of those surfaces together.

### Coverage

The plan covers:

- Method/Repository Standard adoption boundary;
- Blueprint 15 bounded amendment;
- ACRM scope refinement and target rehome;
- new TA planning graph;
- repository bootstrap/status/router model;
- decisions/ADR routing;
- architecture/Capability routing;
- semantic/provenance census before cleanup;
- Product aggregate vs repository roadmap separation;
- validator/CI negative controls;
- Git platform enforcement;
- fresh-session/global coherence proof;
- independent review/operator/merge gates;
- Conexus OS current terminology;
- frozen M0 Evidence preservation.

### Explicitly deferred

This migration does not select or execute any Aurora Product stack, TA-03 operation semantics, API binding, data store, IAM product, Mastra integration, model/provider or runtime deployment.

### Completion definition

Repository rebaseline is complete only when the target operating model is integrated on `main`, all temporary work is absent, repository-standard guards are proven to fire, a fresh actor routes correctly, required provenance remains reachable and `docs/roadmap.md` still blocks Product implementation until later readiness gates close.
