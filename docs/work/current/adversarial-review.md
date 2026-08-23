---
id: REVIEW-AURORA-METHODOLOGY-REPOSITORY-REBASELINE-LEAD
title: Aurora Methodology and Repository Rebaseline Lead Adversarial Review
document_type: temporary_adversarial_review
form: reference
authority: evidence
status: current
version: 0.2.0
owners:
  - developmentconexus-ops
related:
  - DESIGN-AURORA-METHODOLOGY-REPOSITORY-REBASELINE-CANDIDATE
last_reviewed: 2026-08-23
---

# Aurora Methodology + Repository Rebaseline — Lead Adversarial Review

> **BRANCH-ONLY / NON-INDEPENDENT REVIEW.** This is Lead self-challenge Evidence. It does not satisfy the independent-challenger floor for final ratification of a material repository/architecture-governance change.

## 1. Review target

```text
branch: docs/methodology-repository-rebaseline-20260823
candidate: docs/work/current/proposal.md v0.2.0
base main: 35614c581cea32e04305c1ad63522fee151eb283
```

Review question:

> Does candidate v0.2 improve Aurora implementation readiness and repository continuity without duplicating authorities, overengineering the planning lifecycle, discarding valid provenance, or silently changing Product architecture?

## 2. Finding register

| ID | Finding | Severity | Disposition |
|---|---|---:|---|
| MR-F01 | Blueprint 15 itself owns the old documentation/status/layout model; repository migration cannot bypass it | material | **CORRECTED IN v0.2** — bounded constitutional reopen required |
| MR-F02 | Existing D-050 says both Product Blueprint and `docs/roadmap.md` are generated projections | material | **CORRECTED IN v0.2** — preserve Blueprint aggregate, supersede generated repo roadmap |
| MR-F03 | New TA-03 could duplicate Capability Specs by becoming a census of every internal operation | material | **CORRECTED IN v0.2** — only current-horizon cross-system/actor/trust/owner operations belong there |
| MR-F04 | A 13-stage graph could become ceremony / architecture waterfall | material | **CORRECTED IN v0.2** — dependency order + materiality tailoring + explicit allowed bundling |
| MR-F05 | Removing transition/cutover entirely ignores a future real installed-state migration problem | moderate | **CORRECTED IN v0.2** — conditional TA-TX, DEFER SAFELY now |
| MR-F06 | Frontend v2.1 is reusable but not yet canonical in `conexus-methodology` | material | **CORRECTED IN v0.2** — reference only; exact adoption/promotion at first consumer |
| MR-F07 | Current Aurora authority still names MNFS, but the current software-development product is Conexus OS | material terminology/identity | **CORRECTED IN v0.2** — current references refined, historical MNFS provenance retained |
| MR-F08 | Current validator/CI hardcodes the old tree and generated roadmap | material repository control | **CORRECTED IN v0.2** — explicit RM migration and negative-control guard target |
| MR-F09 | GitHub shows `main` currently unprotected | material repository control | **CORRECTED IN v0.2** — branch protection/required aggregate is an explicit migration proof |
| MR-F10 | `CONTRIBUTING.md` also carries the old route/status owner | moderate | **CORRECTED IN v0.2** — keep only with a real contributor consumer and no parallel authority |
| MR-F11 | Replacing STATUS/WORKLOG/acceptance trees could destroy unique durable decisions/Evidence | blocking if done mechanically | **CORRECTED IN v0.2** — semantic/provenance census before deletion; Git reachability law applies |
| MR-F12 | Moving repository/source/build later could leave source organization unknown too long | moderate | **ACCEPTED TRADE-OFF** — only production/source architecture selection is deferred; temporary docs layout remains governed by Repository Standard |
| MR-F13 | TA-05 interaction planning might drag full UI/Voice design into a non-visual Core horizon | moderate | **CORRECTED IN v0.2** — current-horizon human/Presence realization only; visual/voice methods trigger only by real consumer |
| MR-F14 | Paved Road could become an internal framework/platform project before Aurora has enough implementation repetition | material | **CORRECTED IN v0.2** — TA-09 materializes only evidenced repeated protected properties and current consumers; YAGNI remains binding |
| MR-F15 | `AURORA-CONTRACT` implementation ownership class could be confused with G01 Contract Model | moderate | **CORRECTED** — use `AURORA-FOUNDATION` for protected source/scaffold surfaces |
| MR-F16 | Renumbering TA stages could make older accepted documents look wrong or retroactively rejected | material traceability | **CORRECTED IN v0.2** — preserve TA-01/02; decision register records old TA-03+ as refined/superseded ordering, not false history |
| MR-F17 | System-wide readiness program could compete with Product Milestone roadmap | material | **CORRECTED IN v0.2** — Blueprint 14 owns long-horizon Product sequence; `docs/roadmap.md` owns current repository program progression only |
| MR-F18 | ACRM could become redundant after global readiness stages | material | **CORRECTED IN v0.2** — ACRM remains capability/slice realization/evidence lifecycle, global stages own cross-system readiness |
| MR-F19 | Repo migration itself could be mistaken for TA-03 or implementation authority | material | **CORRECTED IN v0.2** — separate RM migration gate; TA-03 and Product implementation remain blocked |
| MR-F20 | Review is self-authored and insufficient for final ratification | material process | **OPEN BY DESIGN** — independent challenge required against fixed durable candidate before final ratification |

## 3. Strongest counterarguments

### 3.1 “Aurora is personal/single-user; this process is too heavy.”

Valid concern. The answer is not to remove structural correctness but to make the graph materiality-driven. Aurora's distributed cognition, governed memory, provider delegation, authority/effects and future physical integration create high-cost boundaries even for one user. What must be avoided is empty ceremony, not planning itself.

Candidate v0.2 therefore makes stage **order and protected-property closure binding**, while document length, separate PR count and depth remain proportional.

### 3.2 “TA-09 repository architecture should still come early because developers need a place for files.”

A temporary repository can host planning without selecting the final source architecture. The failure class is allowing an early folder/workspace decision to decide service, language, contract or dependency boundaries before consumers are known. Repository Standard solves planning-file placement now; TA-09 solves production Paved Road/source/build after upstream semantics are exact enough.

### 3.3 “Operation surface sounds like premature API design.”

Only if operation == HTTP endpoint. Candidate v0.2 explicitly separates semantic operation from protocol. TA-03 records owner/caller/consumer/meaning/failure; TA-04 decides executable contract/binding. Pure capability-internal methods never enter the global operation census.

### 3.4 “Why not keep STATUS and WORKLOG because they worked?”

They worked as a local early solution. The current organizational Repository Standard establishes a smaller invariant: one mutable roadmap/status owner and Git as archive. Keeping parallel status/worklog surfaces would now increase context ambiguity across projects. Their unique semantic/provenance value must be consolidated first; then the live duplication can be removed.

### 3.5 “Why change Blueprint 15 if Product meaning is not changing?”

Blueprint 15 explicitly owns repository/documentation governance and hardcodes the old model. That is constitutional meaning inside its scope. Silently overriding it with an external Repository Standard would create two authorities. A bounded refinement is therefore the smallest correct reopen.

## 4. Global coherence checks

### Authority uniqueness

Candidate v0.2 produces one intended owner per concern:

```text
engineering reasoning            conexus-methodology/METHOD.md
repository operating envelope    conexus-methodology/REPOSITORY-STANDARD.md
Aurora Product meaning           Product Blueprint / accepted Product authorities
current repository progression   docs/roadmap.md after migration
cross-system readiness order     durable Aurora planning/readiness program
capability/slice realization     ACRM
specific technical choice        ADR/Spec/Contract owner
review result                     Evidence
```

No intentional duplicate owner remains.

### Product-vs-program separation

```text
Blueprint 14
= what Product capability sequence Aurora pursues over time

docs/roadmap.md
= what repository gate is active now and what work is allowed next
```

This resolves the current path collision.

### Architecture-vs-implementation separation

TA-03→TA-10 reduce material implementation degrees of freedom. TA-12 then compiles bounded work. TA-13 attacks readiness. Product code remains behind a separate execution grant.

### Anti-overengineering

The candidate explicitly forbids:

- empty stage artifacts;
- exhaustive future operation census;
- universal protocol/store/IAM choice;
- Paved Road capability with no current consumer;
- full frontend/Voice design without a current visual/audio consumer;
- transition machinery without installed-state continuity need.

## 5. Residual obligations before final ratification

1. Perform exact current-path/provenance census during repository-migration design before deleting/re-homing live files.
2. Run current Aurora documentation validation on the branch candidate.
3. Obtain operator approval of candidate v0.2 direction.
4. After approval, write the durable rebaseline + bounded Blueprint 15 refinement + repository migration plan.
5. Run independent Fable challenge against the fixed durable candidate before final ratification.

## 6. Lead verdict

```text
BLOCKING PRODUCT/ARCHITECTURE CONTRADICTIONS: 0
MATERIAL REMEDIATED FINDINGS:              19
OPEN MATERIAL PROCESS OBLIGATION:           1 (independent challenge before final ratification)
OPEN NON-BLOCKING NAMING FINDINGS:          0
```

**Lead recommendation:** candidate v0.2 is coherent enough for operator design review. It is not yet final authority and not yet independently reviewed.