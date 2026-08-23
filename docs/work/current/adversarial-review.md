---
id: REVIEW-AURORA-METHODOLOGY-REPOSITORY-REBASELINE-LEAD
title: Aurora Methodology and Repository Rebaseline Lead Adversarial Review
document_type: temporary_adversarial_review
form: reference
authority: evidence
status: current
version: 0.3.0
owners:
  - developmentconexus-ops
related:
  - DOC-AURORA-METHODOLOGY-REPOSITORY-REBASELINE
  - DOC-AURORA-PLANNING-READINESS
  - DOC-AURORA-ENGINEERING-RULES
last_reviewed: 2026-08-23
---

# Aurora Methodology + Repository Rebaseline — Lead Adversarial Review

> **BRANCH-ONLY / NON-INDEPENDENT.** This is Lead self-challenge Evidence. It does not satisfy the independent-challenger floor required before final ratification.

## 1. Fixed review target

The durable semantic candidate reviewed here is the exact branch revision immediately before this review-only update:

```text
repository: developmentconexus-ops/aurora_project
base main: 35614c581cea32e04305c1ad63522fee151eb283
candidate branch: docs/methodology-repository-rebaseline-20260823
semantic review target: eb412408edc420d69c38e22ba4070773b0c26284
Draft PR: #6
```

Primary reviewed artifacts:

```text
docs/decisions/methodology-repository-rebaseline.md
docs/development/planning-readiness.md
docs/development/engineering-rules.md
docs/work/current/blueprint-15-amendment.md
docs/work/current/plan.md v0.2.0
```

Supporting earlier candidate/review history:

```text
docs/work/current/proposal.md v0.2.0
docs/work/current/index.md v0.2.0
```

Review question:

> Does the fixed durable candidate improve Aurora implementation readiness and repository continuity without duplicating authorities, creating a ceremony-heavy second process framework, discarding current semantics/provenance, or smuggling Product/stack decisions into a repository/methodology rebaseline?

## 2. Finding register

| ID | Finding | Severity | Disposition |
|---|---|---:|---|
| MR-F01 | Blueprint 15 owns the old documentation/status/layout model; Repository Standard cannot silently override it | material | **CORRECTED** — exact bounded constitutional amendment required |
| MR-F02 | Current D-050/generated-doc model collides Product roadmap and repository roadmap | material | **CORRECTED** — Product Blueprint aggregate remains generated; repository roadmap becomes independent current-program owner only at atomic cutover |
| MR-F03 | New TA-03 could duplicate Capability Specs if it censuses every internal method | material | **CORRECTED** — only material cross-owner/actor/runtime/effect operations are admitted |
| MR-F04 | TA-01→TA-13 could become architecture waterfall/ceremony | material | **CORRECTED** — dependency order is binding; artifact/PR count and depth are materiality-tailored; empty stages forbidden |
| MR-F05 | Removing transition/cutover entirely would ignore future installed-state migration | moderate | **CORRECTED** — TA-TX is conditional and safely deferred until a real continuity consumer exists |
| MR-F06 | Frontend Product Experience Method v2.1 is reusable Evidence but not organizational authority | material | **CORRECTED** — specialized method activates only when a real Aurora visual consumer deliberately adopts it |
| MR-F07 | Current Aurora text still uses MNFS while current product is Conexus OS | material terminology | **CORRECTED** — current authority uses Conexus OS; historical MNFS wording is preserved when provenance would be falsified by replacement |
| MR-F08 | Current validator/CI hardcodes old tree and generated roadmap | material control | **CORRECTED** — migration includes generator/validator/CI redesign plus negative controls |
| MR-F09 | GitHub `main` currently lacks target branch-protection enforcement | material control | **CORRECTED AS MIGRATION OBLIGATION** — protection/required aggregate proof is explicit before closure |
| MR-F10 | `CONTRIBUTING.md` carries old read/status model | moderate | **CORRECTED** — retained only if real consumer exists and updated in atomic control-plane cutover |
| MR-F11 | Mechanical removal of STATUS/WORKLOG/acceptance/reviews could destroy unique current semantics/Evidence | blocking if mechanical | **CORRECTED** — semantic+provenance census is RM-01, before any deletion |
| MR-F12 | Moving source/repository/build decisions later could leave implementation structure too late | moderate | **ACCEPTED TRADE-OFF** — Repository Standard supplies planning layout now; production Paved Road/source/build waits for consumers and contracts |
| MR-F13 | TA-05 might pull full frontend/Voice design into a non-visual horizon | moderate | **CORRECTED** — only current-horizon human/Presence realization is global; specialized visual/audio methods require real consumers |
| MR-F14 | Paved Road could become an internal platform project before repetition exists | material | **CORRECTED** — only repeated protected properties/current consumers may enter AURORA-FOUNDATION |
| MR-F15 | `AURORA-CONTRACT` source class would collide with G01 Contract Model vocabulary | moderate | **CORRECTED** — class is `AURORA-FOUNDATION` |
| MR-F16 | Renumbering TA stages could rewrite accepted history | material traceability | **CORRECTED** — TA-01/02 preserved; old TA-03+ is explicitly refined/superseded, never retroactively rejected |
| MR-F17 | System readiness graph could compete with Product Milestone roadmap | material | **CORRECTED** — Blueprint 14 owns Product capability sequence; repository roadmap owns current work progression |
| MR-F18 | ACRM could become redundant or a second repository lifecycle | material | **CORRECTED** — ACRM owns capability/milestone/mission realization; global readiness and repository operation remain separate owners |
| MR-F19 | Repository migration could be mistaken for TA-03 or Product execution authority | material | **CORRECTED** — RM program is separate and TA-03 remains blocked after migration unless separately authorized |
| MR-F20 | Initial review was self-authored | material process | **OPEN BY DESIGN** — independent challenge remains mandatory |
| MR-F21 | Initial migration plan switched `docs/roadmap.md` before removing its generator/validator ownership | material transition correctness | **CORRECTED IN PLAN v0.2** — RM-05 is one atomic router/roadmap/generator/validator/CI cutover |
| MR-F22 | Initial plan refined ACRM scope but did not explicitly rehome it from `docs/product/` to repository engineering authority | moderate ownership/routing | **CORRECTED IN PLAN v0.2** — preserve stable doc ID and rehome to `docs/development/capability-realization.md` without dual-current copies |
| MR-F23 | Earlier Lead review targeted the temporary proposal rather than the compiled durable candidate | material review coverage | **CORRECTED BY THIS v0.3 REVIEW** — exact durable semantic target is `eb412408...` |

## 3. Strongest counterarguments

### 3.1 “Aurora is personal/single-user; this is too heavy.”

Single-user reduces tenancy/commercial complexity, not the inherent cost of sovereign state, governed memory, model/Harness substitution, external effects, recovery and future physical-device boundaries. The correct optimization is proportional ceremony, not missing authority.

The candidate therefore binds **decision dependencies and correctness obligations**, while allowing stage bundling, compact artifacts and `DEFER SAFELY` when no current consumer exists.

### 3.2 “Repository/source/build should remain TA-03 because developers need a place to put files.”

Planning repository structure is already solved by Repository Standard. The decision being deferred to TA-09 is **production source/build/Paved Road architecture**, whose correct shape depends on admitted operations, executable contracts, data/security and cognitive-runtime boundaries. Choosing it earlier risks allowing folders/workspaces to decide system architecture by convenience.

### 3.3 “Cross-System Operation Surface is just API design with another name.”

It is not. TA-03 owns semantic operation admission across material boundaries. TA-04 owns executable schema/binding. An operation can be in-process, local IPC, provider call, effect request or later HTTP/RPC binding. Pure internal methods do not enter the global census.

### 3.4 “Why remove STATUS/WORKLOG if they worked?”

They were a successful earlier local maximum. The organizational Repository Standard now provides a smaller global operating model: one mutable roadmap owner, task router, bounded context and Git as archive. Keeping old parallel surfaces after adopting the standard would preserve context ambiguity. Their unique semantics/provenance must be consolidated before retirement, which RM-01/RM-06 explicitly require.

### 3.5 “Why is Blueprint 15 constitutional reopen necessary for repository mechanics?”

Because Blueprint 15 explicitly owns Aurora documentation governance and names the current status owner, layout, tracking files and generated roadmap. Ignoring it would create two authorities. The bounded amendment is therefore smaller and safer than an out-of-band override.

### 3.6 “Will TA-13 make implementation impossible because every detail must be decided?”

No. TA-13 blocks **material owner/contract/security/data/runtime/user-semantics decisions** from leaking into coding. Local reversible algorithm/function composition remains legitimate implementation judgment. The target is constrained freedom, not zero engineering judgment.

## 4. Cross-document coherence

### 4.1 Authority uniqueness

The fixed candidate has one intended owner per level:

```text
engineering reasoning                  conexus-methodology/METHOD.md
repository operating envelope          conexus-methodology/REPOSITORY-STANDARD.md
Aurora Product meaning                 Product Blueprint / current Product owners
repository-local specialization        docs/development/engineering-rules.md
cross-system implementation readiness  docs/development/planning-readiness.md
capability/slice realization           ACRM, rehomed under docs/development after migration
specific material decision             exact ADR/Spec/Contract owner
current repository progression         docs/roadmap.md after atomic RM-05 cutover
review outcome                         Evidence only
```

No deliberate duplicate authority remains.

### 4.2 Product vs repository roadmap

```text
Blueprint 14
= long-horizon Product capability sequence + Golden Proof direction

docs/roadmap.md
= one mutable current repository gate/status/permission/next action
```

Generator ownership changes in the same atomic cutover that activates the new repository roadmap.

### 4.3 ACRM vs global readiness

```text
Planning Readiness
= close recurring cross-system decisions needed by many later slices

ACRM
= realize one selected milestone/capability/mission against current global authority
```

A later capability consumes current TA outputs; it does not rerun TA-01→TA-13.

### 4.4 TA-03 vs Capability Specs

TA-03 only admits material cross-boundary operations. Capability Specs remain free to define reusable internal behavior/lifecycle. Therefore the same method need not be represented twice.

### 4.5 Repository migration correctness

Plan v0.2 corrects the prior ordering hazard:

```text
census first
→ constitutional target
→ prepare target owners
→ reconcile decision/architecture/capability routes
→ ATOMIC control-plane cutover
   README + AGENTS + index + roadmap + generator + validator + CI
→ retire old live surfaces
→ prove negative controls
→ platform protection
→ fresh-actor/GCR
→ independent final challenge
```

This avoids a period where generated roadmap semantics and mutable-roadmap semantics both claim current authority.

### 4.6 Product/stack non-selection

The candidate names technologies only as future decision classes/examples inherited from prior accepted context. It selects none of:

```text
monorepo/polyrepo
universal language/runtime
Mastra version/integration
RPC/API/event binding
data/memory store
IAM/policy/secrets product
model/provider
Voice stack
service supervisor/container/orchestrator
first AHDK language
```

## 5. Anti-overengineering check

Explicit protections remain:

- stage with no real consumer may `DEFER SAFELY`;
- no empty artifact for numbering;
- adjacent stages may share a coherent gate when dependencies/owners stay explicit;
- no exhaustive future operation inventory;
- no universal protocol/store/IAM decision;
- no Paved Road mechanism without repeated protected property/current consumer;
- no frontend/Voice design without real visual/audio consumer;
- no transition/cutover capability without installed-state continuity need;
- no new package manager/framework solely to make repository controls look uniform.

## 6. Residual obligations before final ratification

Completed since v0.2 Lead review:

```text
operator design-direction approval             COMPLETE — 2026-08-23
durable rebaseline decision candidate          COMPLETE
durable planning/readiness standard candidate  COMPLETE
durable local engineering-rules candidate      COMPLETE
bounded Blueprint 15 amendment                 COMPLETE
repository migration plan v0.2                 COMPLETE
Lead review of fixed durable candidate          COMPLETE — this review
```

Still required:

1. current documentation CI/validation must pass on the post-review exact branch head;
2. independent Fable challenge must review the fixed durable candidate/review pack;
3. Lead must adjudicate every material independent finding;
4. if material corrections change review coverage, run a bounded second independent round;
5. operator final ratification remains separate;
6. repository migration execution remains separately authorized after ratification;
7. merge remains separately authorized.

## 7. Lead verdict

```text
BLOCKING PRODUCT/ARCHITECTURE CONTRADICTIONS: 0
MATERIAL DESIGN/TRANSITION FINDINGS CLOSED:   22
OPEN MATERIAL PROCESS OBLIGATION:              1 independent challenge chain
STACK/PRODUCT DECISIONS SMUGGLED IN:           0 found
TA-01/TA-02 REOPEN REQUIRED:                    NO
REPOSITORY MIGRATION AUTHORIZED:                NO
TA-03 AUTHORIZED:                               NO
PRODUCT IMPLEMENTATION AUTHORIZED:              NO
```

**Lead recommendation:** the fixed durable MR-01 candidate is ready for independent review. It is not final authority and must not be merged or executed as repository migration yet.
