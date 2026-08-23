---
id: DESIGN-AURORA-BLUEPRINT-15-REPOSITORY-AMENDMENT
title: Blueprint 15 Bounded Repository Governance Amendment
document_type: temporary_constitutional_amendment
form: reference
authority: design
status: proposed
version: 0.2.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - branch-only exact MR-01 proposal for bounded Blueprint 15 repository-governance refinement
related:
  - DOC-AURORA-BLUEPRINT-15
  - DOC-AURORA-METHODOLOGY-REPOSITORY-REBASELINE
  - DOC-AURORA-PLANNING-READINESS
  - DOC-AURORA-ENGINEERING-RULES
last_reviewed: 2026-08-23
---

# Blueprint 15 Bounded Repository Governance Amendment

> **BRANCH-ONLY / NON-AUTHORITATIVE.** This file specifies the exact constitutional scope to amend if MR-01 is finally ratified. It must not merge. The accepted amendment is applied directly to the canonical Blueprint 15 source and this temporary file is then deleted.

## 1. Amendment boundary

The current Blueprint 15 remains correct in its constitutional principles:

```text
one durable concept → one canonical owner
conversation = discovery; repository = canonical project memory
research ≠ authority
current authority beats historical mechanism by existence
minimal sufficient context
accepted intent/evidence gates implementation
history/provenance remains discoverable
unknown/stale/rejected states stay explicit
generated projections do not become independent authority
```

MR-01 does **not** reopen those principles.

The bounded reopen affects only clauses that hard-code the pre-Repository-Standard operating model:

- current-status owner;
- canonical entrypoint/read route;
- documentation live-tree layout;
- permanent tracking surfaces;
- permanent `docs/superpowers/**` design/plan surface;
- generated `docs/roadmap.md` model;
- repository bootstrap/CI/temporary-work rules;
- ACRM scope wording where it overlaps repository/global readiness;
- current software-Harness terminology where it says MNFS as a current name.

## 2. Section-by-section amendment

### 15.2 Governing principles — refine P4 only

Preserve the principle that coordination does not own Product doctrine, but change the concrete owner.

**Current meaning to retire:**

```text
tracking coordinates current work
STATUS.md is the current coordination authority
```

**Replacement meaning:**

```text
docs/roadmap.md is the sole mutable repository-program authority for:
- current stage/gate
- implementation permission / blocked work
- exact next action
- current progression state

roadmap coordination does not own Product/architecture semantics
```

Other 15.2 principles remain unchanged except link/path updates.

### 15.4 Authority classes — preserve categories, refine A0 example and A8

Preserve the authority categories but remove old-path examples that would survive as false current routing.

For A0/constitutional examples, replace `Documentation Map for authority/read paths` with the durable constitutional Product/Documentation governance owner plus the Repository Standard route. `docs/index.md` is a router, not constitutional Product authority.

Preserve the existence of a coordination/tracking authority class but make it repository-standard neutral:

```text
A8 — Program Coordination / Tracking
owns current repository-program stage/status/next-action only through docs/roadmap.md;
issues/boards may coordinate work but do not become parallel semantic or status authority.
```

Do not require permanent STATUS/WORKLOG/BACKLOG files as members of this class.

### 15.5 Precedence and conflict — preserve semantics

No precedence change is required except replacing old STATUS/tracking path examples with `docs/roadmap.md` and the current decision/router owners.

`DOCUMENTATION_DIVERGENCE` remains a valid Aurora failure class.

### 15.6 Canonical ownership table — change current coordination owner

Replace:

```text
Current project coordination → STATUS/tracking
```

with:

```text
Current repository program/gate/implementation permission/next action → docs/roadmap.md
Task/intention routing → docs/index.md
Current material decision dispositions → docs/decisions/index.md or exact current decision owner
Repository-local engineering rules → docs/development/engineering-rules.md
```

### 15.7 Metadata — relax uniformity to material utility

Current frontmatter remains useful for Product/decision/specification/contract/evidence identities but is too universal for the Repository Standard target.

Replacement law:

```text
stable IDs/metadata are required when they materially improve authority, machine routing, traceability or generation;
not every Markdown file needs a uniform metadata envelope when path/index routing is sufficient;
generated/machine-readable authorities keep explicit provenance/version identity.
```

Repository migration may keep richer existing metadata when useful. Do not rewrite all durable documents only for metadata aesthetics.

### 15.9 Product Blueprint architecture — preserve

The fifteen modular Blueprint sources and generated aggregate remain current.

`docs/product/PRODUCT-BLUEPRINT.md` stays a generated read-only publication.

### 15.10 Documentation layout — replace target live-tree model

Replace the existing broad target tree with the repository-standard semantic model, creating only paths with a real consumer:

```text
README.md
AGENTS.md
CONTRIBUTING.md when useful

docs/
├── index.md
├── roadmap.md
├── product/
├── architecture/
├── decisions/
├── phases/
├── development/
├── capabilities/
├── reference/
├── research/
├── evidence/
├── diagrams/
└── work/          # branch-only temporary; forbidden in merge candidate/main
```

Aurora-specific additions such as `capabilities/` are valid local specializations because ACRM has a real reusable Capability consumer.

Explicit target exclusions from `main`:

```text
docs/work/**
docs/superpowers/**
permanent session handoffs/dialogues/round trees
parallel mutable status/roadmap surfaces
active archive/old trees used as current authority
```

Do not create empty directories.

### 15.11 Conversation-to-canonical promotion — preserve, update destinations

Keep the promotion law, but route durable content through `docs/index.md` to one current owner. Temporary planning/review material is branch-only under `docs/work/current/` and must be absorbed/deleted before merge.

### 15.12–15.16 Research / ADR lifecycle — preserve principles, update paths

Preserve research evidence discipline and ADR semantic role.

Path target:

```text
docs/research/**
docs/decisions/index.md
docs/decisions/adr/**
```

The exact migration may retain filenames/IDs while rehoming paths. Research remains opt-in and non-authoritative.

### 15.17 Capability Specifications — preserve semantic surface

Keep `docs/capabilities/CAP-*/` as an Aurora-specific durable surface because it has an active ACRM consumer.

Do not force every global architecture stage to duplicate Capability Specs.

### 15.18 Capability Realization Method — refine scope

Replace wording that makes ACRM appear to govern the entire repository/program lifecycle with:

```text
ACRM owns realization of a selected Product Milestone / Capability / Mission:
applicability → requirements → capability readiness → decisions/research/spikes → scoped contract → implementation-design readiness → execution Evidence → Product Milestone closeout.

Cross-system planning/readiness is owned separately by the Aurora Planning and Implementation-Readiness Standard and is consumed by ACRM as current authority.

Repository navigation/Git/review/context operating rules are owned by the DevelopmentConexus Repository Standard plus Aurora local engineering rules.
```

R0–R8 identities and their accepted M0 evidence are not invalidated by this scope refinement.

### 15.19 Tracking documents — replace permanent-tracking requirement

Retire the requirement that these permanent live documents must exist:

```text
docs/tracking/STATUS.md
docs/tracking/WORKLOG.md
docs/tracking/DECISIONS.md
docs/tracking/BACKLOG.md
docs/tracking/DOCUMENTATION-COVERAGE.md
```

Replace with:

```text
docs/roadmap.md
→ sole mutable stage/status/implementation permission/next action

docs/decisions/index.md
→ current material decision disposition and forward obligation discovery

Git + merged PR history
→ chronological change record/provenance

docs/evidence/** / docs/phases/**
→ durable proof/closure only when a current consumer exists

docs/work/**
→ temporary branch-only planning/review, never current main authority
```

Existing tracking content is retired only after semantic census/rehome and reachability proof. Every current disposition and forward obligation in `docs/tracking/DECISIONS.md` must be mapped into `docs/decisions/index.md` or an exact durable decision owner before retirement.

### 15.20 Status and authorization vocabulary — preserve vocabulary, move owner

Keep explicit authorization distinctions such as:

```text
DISCOVERY
RESEARCH
DESIGN
SPIKE
PLAN
IMPLEMENTATION
EXTERNAL EFFECT
MERGE
```

But current permission/next action is read only from `docs/roadmap.md` after migration.

### 15.21 Supersession — preserve and align reachability law

Keep history preservation and add:

```text
Git is archive only when required provenance remains reachable.
Unique unmerged provenance with a current consumer receives a durable ref before last branch deletion.
No archive directory is created solely to preserve superseded files.
```

### 15.22 Generated projections — stop generating repository roadmap

Keep:

```text
docs/product/PRODUCT-BLUEPRINT.md
← generated from Blueprint 01–15
```

Retire:

```text
docs/roadmap.md
← generated from Blueprint 14
```

`docs/roadmap.md` becomes hand-maintained current program authority.

Blueprint 14 remains Product/capability-roadmap authority. The two documents intentionally answer different questions.

Update `scripts/generate_docs.py` and validation accordingly during repository migration.

### 15.23 Documentation impact — simplify into repository workflow

Preserve the law that material changes must account for documentation/authority impact. Do not require one particular YAML block on every PR if the Repository Standard/PR template can enforce the property more compactly.

### 15.24 Documentation checks — replace/extend target controls

The migrated repository verification must cover at least:

```text
AGENTS + docs/index + docs/roadmap <= 20 KiB
docs/roadmap sole mutable status/next-action authority
README landing-only
default task pack <= 5 files
durable current docs reachable from router
relative links resolve
no durable authority depends on docs/work
no docs/work in merge candidate/main
no docs/superpowers in merge candidate/main
no permanent handoff/dialogue/review-round tree
no duplicate mutable roadmap/status surface
current decision dispositions valid/discoverable
unique required unmerged provenance preserved
review branch - exact candidate branch = docs/work/current/ai-dialog.md only
blocked-implementation top-level/source surfaces satisfy an explicit allowlist
material guards have deterministic negative controls
aggregate required CI gate exists/protects main
```

The review-isolation and blocked-implementation allowlist controls require deterministic negative fixtures/equivalent falsifiers; manual inspection or old-name denylists alone are insufficient.

Aurora-specific validators may continue checking Product Blueprint generation, requirement identities, research manifests and other current real consumers when those controls remain useful.

### 15.26 Human read paths — replace old bootstrap routes

Replace current-contributor and implementation-worker routes that require `STATUS` / `Documentation Map` with:

```text
AGENTS
→ docs/index
→ docs/roadmap
→ 1–2 exact owning documents
```

Architecture/Product/research readers enter through `docs/index.md` and the routed semantic owner. Do not make `docs/index.md` or `docs/roadmap.md` a replacement Product/architecture owner.

### 15.29 Ownership / CODEOWNERS examples — update target paths

Keep operator approval distinct from Git write permission. Replace old future CODEOWNERS path examples with current target semantic surfaces, for example:

```text
/docs/product/
/docs/architecture/
/docs/decisions/
/docs/development/
/docs/capabilities/
/docs/research/
/docs/evidence/
AGENTS.md
contract/schema directories when they exist
security/policy paths when they exist
```

Do not create directories solely to satisfy a CODEOWNERS example.

### 15.31 A0 acceptance rule and post-A0 current-state ownership — preserve A0, replace current owner

Preserve the historical A0 acceptance record and its proof requirements.

Replace the post-A0 current-state clause with:

```text
After repository migration, docs/roadmap.md owns the selected current repository-program stage/gate,
blockers, authorization boundary and exact next action.
Product/capability roadmap meaning remains in Blueprint 14.
Current decisions remain in their exact owners / docs/decisions index.
```

A0 acceptance still never authorizes later gates, Architecture Spike execution or implementation by implication.

## 3. Current-name refinement: MNFS → Conexus OS

Do not globally replace historical text.

Apply this law:

```text
current authoritative software-Harness references → Conexus OS
historical evidence/closed-stage snapshots written when old name was current → MNFS may remain
first current semantic mention where needed → "Conexus OS (historically MNFS)"
```

Blueprint 07 should update the current example heading/text to `Conexus OS`, preserving the same Harness responsibility boundary.

Blueprint 14 closed A0/M0 historical snapshots may retain MNFS where changing it would falsify historical wording; current forward-looking references use Conexus OS.

## 4. Product meaning explicitly unchanged

This amendment does not change:

- Aurora definition/North Star;
- Leandro authority;
- memory/product semantics;
- G01/C01–C12 ownership;
- TA-01/TA-02 topology;
- Contract Model/AHDK direction;
- Stage A Presence constraints;
- M0 Golden Proof or frozen R7 disposition;
- any technology selection.

## 5. Acceptance/proof obligations

Before applying this amendment to canonical Blueprint 15:

1. every removed old operating responsibility has a target owner;
2. the migration plan proves no current semantic obligation is lost;
3. generated roadmap separation is mechanically feasible;
4. the target fresh-actor route is testable;
5. independent review challenges the exact amendment;
6. operator final ratification is explicit;
7. a completeness sweep proves no residual current-authority reference to `STATUS`, `Documentation Map`, permanent tracking, `docs/superpowers/**` or superseded path names survives outside explicitly labeled historical snapshots/evidence.
