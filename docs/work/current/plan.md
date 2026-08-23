---
id: PLAN-AURORA-MR-01-REPOSITORY-MIGRATION
title: Aurora Methodology and Repository Rebaseline Migration Plan
document_type: temporary_execution_plan
form: reference
authority: design
status: proposed
version: 0.3.0
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

> **Execution boundary:** this plan becomes executable only after MR-01 final ratification and a separate explicit repository-migration authorization. It does not authorize Aurora Product/runtime implementation, TA-03+, M0 R7/R8, Architecture Spikes or merge.

## 1. Goal

Migrate Projeto Aurora from the pre-Repository-Standard documentation/status model to the DevelopmentConexus repository operating envelope while preserving all current Product/architecture authority, installing the new Planning/Implementation-Readiness program, and proving that no current semantic obligation or required provenance is lost.

This is an **authority-preserving compiler**, not a cleanup project.

## 2. Integration topology — binding correction from independent review

The migration MUST use **one migration branch and one final merge candidate**.

```text
revalidated main
→ one migration branch / Draft PR
→ RM-01 … RM-09 checkpoints on that same branch
→ no partial RM merge to main
→ final independent review of exact consolidated candidate
→ final operator ratification
→ separate merge authorization
→ one squash merge
```

RM numbers are review/checkpoint boundaries, **not independent merge gates**. This prevents a main-branch window where Blueprint 15 names `docs/roadmap.md` as current authority while the old generator/validator/status model still controls the repository.

RM-08 is GitHub platform enforcement performed after the migrated aggregate check is stable on the open candidate and before final merge authorization. If a GitHub limitation makes one protection setting impossible before merge, MR-01 closeout remains open until that setting is applied and revalidated immediately after merge; it never authorizes TA-03 in the interim.

## 3. Global constraints

- Revalidate exact `main`, candidate branch, PR and CI before each RM checkpoint.
- TA-01 and TA-02 remain canonical unless concrete contradictory Evidence appears.
- No Product/runtime code, Product schema/dependency or production deployment change is part of this migration.
- No live surface is retired until surviving semantics and required provenance have one proven destination.
- `docs/work/**` remains branch-only and is removed before final merge candidate.
- Permanent `docs/superpowers/**` is retired only after current semantic/provenance coverage is proven.
- Current authoritative name is `Conexus OS`; historical snapshots may retain `MNFS` when that is the historically correct name.
- `docs/roadmap.md` becomes sole mutable repository-program authority only in RM-05, atomically with router/generator/validator/CI changes.
- The final exact candidate must be green; temporary intermediate CI failures are not completion Evidence.

## 4. Program

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

---

## RM-01 — Current Semantic + Provenance Census

**Outcome:** zero current live file/class remains unclassified; this checkpoint authorizes no deletion by itself.

**Census inputs:**

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

Classify every item as exactly one primary disposition:

```text
CURRENT AUTHORITY
CURRENT ROUTER / PROGRAM STATE
DURABLE EVIDENCE WITH CURRENT CONSUMER
CURRENT FORWARD OBLIGATION
HISTORICAL / GIT SUFFICIENT
UNMERGED UNIQUE PROVENANCE STILL REQUIRED
TEMPORARY / DUPLICATE
```

**Mandatory specific proofs:**

- every fact in `docs/tracking/STATUS.md` maps to future roadmap/decision/owner homes;
- every current disposition and forward obligation in `docs/tracking/DECISIONS.md` maps to `docs/decisions/index.md` or an exact decision owner;
- `docs/tracking/WORKLOG.md` contains no unique current obligation that would disappear into chronology-only Git history;
- A0 operator-acceptance/fresh-session acceptance records and any other acceptance record still cited by current accepted authority are classified explicitly as `DURABLE EVIDENCE WITH CURRENT CONSUMER` until their target `docs/phases/**` or `docs/evidence/**` home is proven;
- `docs/superpowers/**` current semantics are mapped before retirement;
- frozen M0 R7 exact ref remains reachable while named as Evidence;
- every unique unmerged branch/PR provenance consumer gets a durable ref if Repository Standard §10 requires one.

---

## RM-02 — Constitutional / Method Reconciliation

**Outcome:** canonical target constitution no longer requires the superseded repository/status model, while `main` remains on the old model until the eventual single migration merge.

**Modify on the migration branch:**

```text
docs/product/blueprint/15-documentation-research-governance.md
docs/product/CAPABILITY-REALIZATION-METHOD.md
docs/product/blueprint/07-harness-orchestration.md
forward-looking references in docs/product/blueprint/14-capability-roadmap.md only when required
```

Apply exactly the bounded deltas in `docs/work/current/blueprint-15-amendment.md`, including sections 15.26, 15.29 and 15.31 and the explicit retirement/rehome of `docs/tracking/DECISIONS.md`.

Preserve R0–R8 identities and historical M0 Evidence. Refine ACRM scope only:

```text
ACRM → milestone/capability/mission realization
Planning Readiness → cross-system implementation-readiness
Repository Standard + local engineering rules → repository operation
```

Regenerate the Product Blueprint aggregate while the current generator still owns it. Do not treat the branch target wording as `main` current-state authority before final migration merge.

---

## RM-03 — Target Authority Preparation

**Outcome:** target durable owners exist and are reviewable without becoming a second current program/status authority.

Create/finalize as `PROPOSED`:

```text
docs/development/planning-readiness.md
docs/development/engineering-rules.md
docs/development/capability-realization.md
docs/decisions/methodology-repository-rebaseline.md
docs/decisions/index.md
docs/architecture/index.md
```

**ACRM rehome law:** preserve stable document identity `DOC-AURORA-CAPABILITY-REALIZATION-METHOD`; remove the old current copy only in the same consolidated candidate that updates all current links/routers.

Planning-stage reconciliation must explicitly map old TA-03…TA-08 to new TA-03…TA-13 dispositions while preserving TA-01/TA-02 unchanged.

---

## RM-04 — Decision / Architecture / Capability Reconciliation

**Outcome:** every current decision/architecture/capability meaning has one target route before the control-plane cutover.

### Decisions

Prepare:

```text
docs/decisions/index.md
docs/decisions/adr/0001-*.md … 0009-*.md
```

Preserve exact ADR IDs/status/scope. Carry all forward obligations from `docs/tracking/DECISIONS.md` using controlled dispositions:

```text
CURRENT | PRESERVE | REFINED | REOPEN | DEFERRED | SUPERSEDED | REJECTED
```

### Architecture

Prepare `docs/architecture/index.md`. Classify current `docs/design/**` individually:

```text
current structural architecture → docs/architecture/**
current decision → docs/decisions/**
closed-stage design/microdesign → docs/phases/** or Evidence/history disposition
spike result/spec → phase/evidence owner
historical-only design → Git when no current live consumer remains
```

### Capabilities

Keep `docs/capabilities/CAP-*/` as an Aurora-specific durable surface. Global TA documents must not duplicate internal Capability Spec behavior.

---

## RM-05 — Atomic Repository Control-Plane Cutover

**Outcome:** on one exact candidate revision, switch repository routing/status ownership and its mechanical guards together.

Change together:

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
all inbound links required by the new current route
```

Target route:

```text
AGENTS.md
→ docs/index.md
→ docs/roadmap.md
→ 1–2 exact owners
```

Target roles:

```text
README.md       landing only
docs/index.md   task/intention routing only
docs/roadmap.md sole mutable stage/status/allowed-work/next-action authority
Blueprint 14    Product capability-roadmap authority
```

**Generator rule:** `scripts/generate_docs.py` continues to generate `docs/product/PRODUCT-BLUEPRINT.md` but MUST NOT generate or overwrite `docs/roadmap.md`.

**Current-authority rule:** `docs/tracking/STATUS.md` is removed if RM-01 proves safe, otherwise unmistakably marked superseded/non-current and unreachable from the current router until RM-06 retirement. Never two current status authorities.

Initial cutover negative controls must prove at least:

1. generator cannot modify `docs/roadmap.md`;
2. stale Product Blueprint projection fails;
3. missing `docs/index.md` fails;
4. bootstrap budget >20 KiB fails;
5. duplicate mutable status owner fixture fails;
6. durable authority depending on `docs/work/**` fails;
7. base→candidate diff/whitespace guard uses the intended range.

---

## RM-06 — Legacy Live-Surface Retirement

**Outcome:** no superseded duplicate authority or permanent temporary-work surface remains in the merge candidate.

Candidate retirements after RM-01 coverage proof:

```text
docs/DOCUMENTATION-MAP.md
docs/tracking/STATUS.md
docs/tracking/WORKLOG.md
docs/tracking/DECISIONS.md
docs/tracking/BACKLOG.md
docs/tracking/DOCUMENTATION-COVERAGE.md
docs/superpowers/**
old docs/adr/** after decision rehome
old docs/design/** after individual disposition
obsolete permanent review/acceptance chronology after current Evidence is rehomed
```

For every retirement class prove:

```text
surviving semantic obligation mapped
current inbound refs updated
required acceptance/provenance reachable
unique unmerged provenance protected when required
replacement route tested
```

`docs/history/**` stays live only for a named current historical consumer; otherwise reachable Git history is sufficient.

---

## RM-07 — Verification / Negative-Control Closure

**Outcome:** the repository guard proves the target properties and each material behavioral guard is shown capable of firing.

Required controls:

```text
AGENTS + docs/index + docs/roadmap <= 20 KiB
docs/roadmap sole mutable stage/status/allowed-work/next-action owner
README landing-only
default routed task pack <= 5 files unless named reason
durable current docs reachable from router
no durable current authority depends on docs/work
no docs/work in merge candidate/main
no docs/superpowers in merge candidate/main
no permanent session/handoff/review-round trees
no duplicate mutable roadmap/status surface
current decision dispositions valid/discoverable
required unique unmerged provenance preserved
PR diff checks compare intended base→candidate
review branch - exact candidate branch = docs/work/current/ai-dialog.md only
implementation-blocked repository surfaces satisfy an explicit top-level/source allowlist
```

The review-isolation guard and the blocked-implementation allowlist each require a deterministic negative fixture/equivalent falsifier. Presence-only checks do not count.

Retain useful Aurora-specific controls for Product Blueprint generation, stable requirement identities, research source manifests and current link/routing integrity.

Use one stable local verification entrypoint consumed by CI. Keep Python if it remains the smallest mechanism; do not add npm/tooling only for uniformity.

---

## RM-08 — Git / Branch Protection

**Outcome:** GitHub enforcement prevents bypass of the migrated repository contract.

After RM-07 makes the aggregate candidate check stable, configure and verify on `main` protection before merge authorization:

```text
force-push forbidden
deletion forbidden
PR-based integration
at least one required aggregate status check
normal merge method = squash
automatic head deletion when appropriate
```

Do not rename a functioning check solely for aesthetics. Query GitHub settings and record exact Evidence.

---

## RM-09 — Fresh-Actor + Global Coherence Proof

**Outcome:** a repository-only fresh actor can recover current authority without archaeology and no duplicate/missing owner remains.

The actor must correctly state:

```text
what Aurora is
current gate / blocked work / exact next action
TA-01/TA-02 accepted state
Product roadmap vs repository roadmap ownership
current decision route
architecture/capability routes
temporary-work rule
Conexus OS relationship to Aurora
frozen M0 Evidence route
TA-03 and Product implementation remain unauthorized
```

Global Coherence Review attacks at least:

```text
duplicate/missing authority
circular routing
Planning Readiness vs ACRM duplication
stale STATUS/Documentation-Map/old-roadmap references
current-name MNFS leakage
lost semantic obligation
generated projection becoming authority
old TA order surviving as current
permanent temporary-work surface
organizational standards copied locally as second authority
unjustified task path requiring >5 files
```

Exit only with zero unresolved material coherence finding.

---

## RM-10 — Independent Review + Final Promotion

**Outcome:** independent challenger reviews the exact consolidated migration candidate; Lead adjudicates; operator ratifies; merge stays separately authorized.

Canonical isolation:

```text
exact candidate
→ review/<gate>-fable
→ review branch adds only docs/work/current/ai-dialog.md
→ reviewer output = Evidence
→ Lead adjudicates on candidate
→ review branch closes unmerged
```

A second round occurs only if material corrections change the reviewed property enough that first-round coverage no longer applies.

Before asking for merge authorization, prove on the exact candidate:

```text
docs/work/** absent
docs/superpowers/** absent when ratified retirement applies
current routes valid
required aggregate CI green
all material negative controls demonstrated
independent material findings = 0 unresolved
operator final ratification = explicit
```

Merge authorization is separate. After integration, revalidate `main`; `docs/roadmap.md` must still state `TA-03 NOT STARTED / NOT AUTHORIZED` unless separately opened.

## 5. Completion definition

MR-01 migration is complete only when the target operating model is integrated, all temporary surfaces are absent from `main`, repository-standard guards are proved to fire, required provenance remains reachable, fresh-actor/global-coherence proof passes, GitHub protection is verified and no downstream Product/TA authorization has been inferred from the migration.
