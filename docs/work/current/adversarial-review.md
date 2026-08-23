---
id: REVIEW-AURORA-METHODOLOGY-REPOSITORY-REBASELINE-LEAD
title: Aurora Methodology and Repository Rebaseline Lead Adversarial Review
document_type: temporary_adversarial_review
form: reference
authority: evidence
status: current
version: 0.3.1
owners:
  - developmentconexus-ops
related:
  - DOC-AURORA-METHODOLOGY-REPOSITORY-REBASELINE
  - DOC-AURORA-PLANNING-READINESS
  - DOC-AURORA-ENGINEERING-RULES
last_reviewed: 2026-08-23
---

# Aurora Methodology + Repository Rebaseline — Lead Adversarial Review

> **BRANCH-ONLY / NON-INDEPENDENT.** This Lead self-challenge does not satisfy the independent-challenger floor required before final ratification.

## 1. Fixed semantic review target

```text
repository: developmentconexus-ops/aurora_project
base main: 35614c581cea32e04305c1ad63522fee151eb283
candidate branch: docs/methodology-repository-rebaseline-20260823
semantic review target: eb412408edc420d69c38e22ba4070773b0c26284
Lead review commit: 2956be55308f17634d53e24e42458f26c62f8166
Draft PR: #6
```

The post-review commits only align temporary status/provenance wording with the already-reviewed durable semantics; they do not add Product/stack/runtime decisions. Independent review must nevertheless pin its own exact then-current candidate head.

Primary semantic artifacts reviewed at `eb412408...`:

```text
docs/decisions/methodology-repository-rebaseline.md
docs/development/planning-readiness.md
docs/development/engineering-rules.md
docs/work/current/blueprint-15-amendment.md
docs/work/current/plan.md v0.2.0
```

## 2. Finding register

| ID | Finding | Severity | Disposition |
|---|---|---:|---|
| MR-F01 | Blueprint 15 owns the old documentation/status/layout model; Repository Standard cannot silently override it | material | **CORRECTED** — exact bounded constitutional amendment required |
| MR-F02 | Generated repository roadmap collides with target mutable repository-roadmap owner | material | **CORRECTED** — Product aggregate remains generated; repo roadmap changes only in atomic control-plane cutover |
| MR-F03 | TA-03 could duplicate Capability Specs by cataloguing internal methods | material | **CORRECTED** — only material cross-owner/actor/runtime/effect operations admitted |
| MR-F04 | TA-01→TA-13 could become ceremony/waterfall | material | **CORRECTED** — dependency order binds; depth/artifact/PR count are materiality-tailored; empty stages forbidden |
| MR-F05 | No transition/cutover stage would ignore future installed-state continuity | moderate | **CORRECTED** — conditional TA-TX, deferred until real consumer |
| MR-F06 | Frontend v2.1 is reusable Evidence but not current organizational authority | material | **CORRECTED** — specialized method only by real Aurora consumer/adoption |
| MR-F07 | Current authoritative name MNFS is stale | material terminology | **CORRECTED** — current name Conexus OS; historical MNFS provenance retained |
| MR-F08 | Current validator/CI hardcodes old tree/generated roadmap | material control | **CORRECTED AS MIGRATION OBLIGATION** — generator/validator/CI redesign + negative controls |
| MR-F09 | `main` lacks target platform enforcement | material control | **CORRECTED AS MIGRATION OBLIGATION** — branch protection/required aggregate proof before closure |
| MR-F10 | `CONTRIBUTING.md` carries old route/status model | moderate | **CORRECTED** — update/retain only with real consumer in cutover |
| MR-F11 | Mechanical tracking/review cleanup could destroy unique semantics/Evidence | blocking if mechanical | **CORRECTED** — RM-01 semantic/provenance census precedes deletion |
| MR-F12 | Moving production source/build decisions to TA-09 may be too late | moderate | **ACCEPTED TRADE-OFF** — Repository Standard governs planning layout now; production Paved Road waits for actual consumers/contracts |
| MR-F13 | TA-05 might pull full UI/Voice design too early | moderate | **CORRECTED** — current-horizon interaction only; visual/audio specialization by consumer |
| MR-F14 | Paved Road could become premature internal platform | material | **CORRECTED** — only repeated protected properties/current consumers enter AURORA-FOUNDATION |
| MR-F15 | `AURORA-CONTRACT` source class collides with G01 Contract Model | moderate | **CORRECTED** — `AURORA-FOUNDATION` |
| MR-F16 | New TA numbering could falsify prior accepted history | material traceability | **CORRECTED** — TA-01/02 preserved; old TA-03+ explicitly refined/superseded, never retro-rejected |
| MR-F17 | Global readiness could compete with Product Milestone roadmap | material | **CORRECTED** — Blueprint 14 owns Product sequence; repo roadmap owns current program progression |
| MR-F18 | ACRM could become redundant/parallel lifecycle | material | **CORRECTED** — ACRM owns capability/milestone/mission realization only |
| MR-F19 | Repository migration could be mistaken for TA-03/Product execution | material | **CORRECTED** — RM program separate; TA-03 stays blocked without new authorization |
| MR-F20 | Self-review cannot satisfy independent floor | material process | **OPEN BY DESIGN** — independent review required |
| MR-F21 | Initial plan changed roadmap before generator/validator ownership | material transition | **CORRECTED IN PLAN v0.2** — RM-05 atomic cutover |
| MR-F22 | Initial plan refined ACRM but did not explicitly rehome repository-engineering ownership | moderate | **CORRECTED IN PLAN v0.2** — stable ID rehome to `docs/development/capability-realization.md` without dual current copies |
| MR-F23 | Earlier review covered temporary proposal, not compiled durable owners | material review coverage | **CORRECTED** — fixed durable target reviewed at `eb412408...` |

## 3. Strongest counterarguments

### “Aurora is single-user; this is too heavy.”

Single-user removes SaaS/tenant/commercial complexity but not sovereign state, governed memory, provider/Harness substitution, authority/effects, recovery and future physical boundaries. The candidate scales ceremony down while preserving material correctness obligations.

### “Repository/source/build belongs earlier.”

Repository Standard already gives planning files a stable home. TA-09 defers only production source/build/Paved Road decisions whose correct shape depends on operations, contracts, data/security and cognitive-runtime consumers.

### “TA-03 is just premature API design.”

TA-03 owns semantic cross-boundary operations, not HTTP. TA-04 owns executable contract/binding. Capability-internal methods do not become global operations.

### “Keep STATUS/WORKLOG because they worked.”

They were a successful local maximum. After adopting Repository Standard, retaining parallel mutable status/worklog surfaces would reintroduce context ambiguity. RM-01/RM-06 preserve current semantics/provenance before retirement.

### “Why reopen constitutional Blueprint 15 for repository mechanics?”

Because Blueprint 15 explicitly owns those mechanics today. A bounded amendment avoids two competing authorities.

### “TA-13 means every detail must be decided before code.”

No. It blocks material Product/owner/contract/security/data/runtime/user-semantics ambiguity. Reversible local algorithms/functions remain legitimate implementation judgment.

## 4. Cross-document coherence verdict

```text
engineering reasoning                  → external METHOD.md
repository operating envelope          → external REPOSITORY-STANDARD.md
Aurora Product meaning                 → Product Blueprint / current Product owners
repository-local specialization        → docs/development/engineering-rules.md
cross-system readiness                 → docs/development/planning-readiness.md
capability/slice realization           → ACRM (rehome after migration)
specific technical decisions           → exact ADR/Spec/Contract owners
current repo progression after cutover → docs/roadmap.md
review output                           → Evidence only
```

No deliberate duplicate owner remains.

Plan v0.2 correctly sequences:

```text
semantic/provenance census
→ constitutional target
→ target owner preparation
→ decision/architecture/capability reconciliation
→ atomic README/AGENTS/index/roadmap/generator/validator/CI cutover
→ old live-surface retirement
→ negative-control closure
→ Git protection
→ fresh-actor/GCR
→ independent final review/promotion
```

The candidate selects no Product runtime, database, IAM, protocol, model/provider, Voice stack, repository strategy or Mastra version.

## 5. Residual obligations

Completed:

```text
operator design-direction approval             COMPLETE — 2026-08-23
durable decision candidate                     COMPLETE
durable planning/readiness candidate            COMPLETE
durable engineering-rules candidate             COMPLETE
bounded Blueprint 15 amendment                 COMPLETE
repository migration plan v0.2                 COMPLETE
Lead durable-candidate review                  COMPLETE
```

Still required:

1. current documentation CI/validation on the exact independent-review base candidate;
2. independent Fable challenge against that exact candidate;
3. Lead adjudication of every material independent finding;
4. second independent round only if material corrections invalidate prior coverage;
5. explicit final operator ratification;
6. separate repository-migration execution authorization;
7. separate merge authorization.

## 6. Lead verdict

```text
BLOCKING PRODUCT/ARCHITECTURE CONTRADICTIONS: 0
MATERIAL FINDINGS CLOSED:                      22
OPEN MATERIAL PROCESS OBLIGATION:              independent review chain
TA-01/TA-02 REOPEN REQUIRED:                   NO
REPOSITORY MIGRATION AUTHORIZED:               NO
TA-03 AUTHORIZED:                              NO
PRODUCT IMPLEMENTATION AUTHORIZED:             NO
```

**Recommendation:** ready for independent review; not ready for ratification, migration execution or merge.
