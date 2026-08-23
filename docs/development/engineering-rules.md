---
id: DOC-AURORA-ENGINEERING-RULES
title: Aurora Repository Engineering Rules
document_type: repository_engineering_rules
form: reference
authority: standard
status: accepted
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - Aurora repository-local specialization of DevelopmentConexus Method and Repository Standard
  - Aurora-specific stop conditions and verification posture
related:
  - DOC-AURORA-BLUEPRINT-15
  - DOC-AURORA-METHODOLOGY-REPOSITORY-REBASELINE
  - DOC-AURORA-PLANNING-READINESS
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
review_triggers:
  - local repository workflow changes
  - verification/CI contract changes
  - Aurora-specific safety or execution boundary changes
  - organizational Method or Repository Standard amendment
last_reviewed: 2026-08-23
accepted_at: 2026-08-23
acceptance_evidence: DOC-AURORA-MR-01-OPERATOR-RATIFICATION
---

# Aurora Repository Engineering Rules

> **ACCEPTED / OPERATOR-RATIFIED.** This document is repository-local specialization only. It becomes active only if MR-01 is finally ratified and integrated.

Cross-repository authorities remain external and canonical:

- `developmentconexus-ops/conexus-methodology/METHOD.md` — DevelopmentConexus Engineering Method v1.0.0;
- `developmentconexus-ops/conexus-methodology/REPOSITORY-STANDARD.md` — DevelopmentConexus Repository Standard v1.0.0.

This file does not copy or redefine those authorities.

## 1. Aurora-specific hard stops

Stop and return to the smallest owning decision when work would:

- change Aurora Product meaning or a constitutional invariant without the owning Product amendment;
- create, move or duplicate a semantic owner;
- change a trust, authority, effect-enforcement or credential boundary;
- invent a cross-system operation, contract or persistent data meaning during implementation;
- generalize an M0-scoped technology decision to Aurora globally without current authority;
- make a model/framework/provider own sovereign Aurora identity, state, authority, governed memory or global Outcome;
- make Conexus OS or another Harness a required sovereign runtime dependency by convenience;
- create an external/physical consequential effect without the exact authorized proof/execution scope;
- weaken a material guard merely to make CI green;
- begin Product implementation while `docs/roadmap.md` blocks it.

## 2. Fresh-actor target route

After the repository migration is ratified and executed:

```text
AGENTS.md
→ docs/index.md
→ docs/roadmap.md
→ 1–2 task-specific owning documents
```

Normal work targets five files or fewer. Exceed that only for a named material reason.

Do not recursively read Product Blueprint, phase history, Evidence, research, Git history, old review chronology, frozen M0 implementation or qualification harnesses before a concrete question requires them.

## 3. Current authority versus history

Current accepted Aurora authority outranks:

- conversation recollection;
- old branch/PR names;
- historical implementation;
- framework docs;
- reviewer findings;
- research reports;
- generated projections.

Implementation/code/tests are Evidence of current mechanism behavior. They do not become target architecture merely because they exist or once passed.

## 4. M0 frozen candidate rule

The preserved M0 R7 branch remains Evidence only until explicitly reopened under current authority.

Do not:

- resume it by implication;
- copy its package/schema/storage shape into new architecture because it is available;
- infer R7 acceptance or R8 closure from green CI;
- delete it while its provenance remains a named current consumer.

Reuse of an implementation mechanism requires a current consumer and proof that the mechanism still fits accepted target ownership/contract direction with less total complexity than rederivation.

## 5. Conexus OS relationship

Conexus OS is the current name of the software-development Harness historically referred to as MNFS.

The Conexus OS sibling repository and the delegated software-development Harness are the same external system in different contexts. The repository is comparison/reference Evidence for Aurora and never becomes Aurora authority by existence; an explicit Delegation may authorize that system to act as a Harness provider within Aurora-owned contracts and limits.

Conexus OS may, when explicitly delegated:

- research;
- plan;
- implement;
- test;
- review;
- package artifacts/Evidence.

It does not own Aurora Product semantics, sovereign state, authority, budgets, provider trust, governed memory or global Outcome. Its own Method/Paved Road/runtime choices do not automatically become Aurora runtime choices.

## 6. Material decision discipline

Apply the organizational Method proportionally:

```text
Evidence
→ Known / Inferred / Unknown / Deferred
→ root cause
→ target invariant
→ constraints
→ credible alternatives
→ Local vs Global Maximum
→ essential vs accidental complexity
→ YAGNI / future cost
→ authority / boundary
→ enforcement
→ proof strategy
→ adversarial challenge
→ decision
→ reopen triggers
```

A material mechanism/technology decision additionally uses the Aurora realization disposition:

```text
ADOPT | ADAPT | BUILD | DEFER | STOP
```

## 7. Research and external facts

Use current external research only when an upcoming material decision depends on unstable standards, exact versions, framework/provider behavior, security status or operational properties.

Prefer normative/official primary sources. Exact version-sensitive claims require exact version/source Evidence before becoming deciding.

Research does not add a dependency or Product capability by itself.

## 8. Temporary work

Temporary planning/review material belongs only on a branch under:

```text
docs/work/current/
```

It is non-authoritative and must be absorbed or deleted before a merge candidate can be promoted.

Permanent `docs/superpowers/**`, session handoffs, dialogue logs, review-round trees and parallel status dashboards are not target live surfaces.

## 9. Independent review

Use the standard isolated Fable review workflow for material architecture/planning/governance decisions when the organizational Method floor is triggered.

Reviewer output is Evidence, never Product/architecture authority.

A material finding is first classified against current authority:

- defect against existing authority → bounded correction;
- proposal creating new authority → return to decision/operator;
- preference/ceremony-only suggestion → no forced change.

The repository verification must be capable of proving the review-isolation invariant mechanically:

```text
review branch - exact candidate branch
= docs/work/current/ai-dialog.md only
```

That guard requires a deterministic negative control; a manually inspected diff is Evidence for one review, not sufficient proof that the repository guard exists.

A second review round is justified only when material corrections change the reviewed property enough that prior challenge no longer covers it.

## 10. Git / PR lifecycle target

After repository migration:

```text
main
→ one branch / one Draft PR per coherent gate by default
→ analysis + candidate
→ independent review when required
→ bounded corrections
→ verification
→ explicit operator ratification where required
→ explicit merge authorization when required
→ squash merge
→ delete temporary review/work surfaces
→ next gate starts from revalidated main
```

Do not:

- commit directly to `main`;
- force-push/rewrite shared history;
- stack a later gate on an unmerged earlier gate by default;
- merge on the basis of file existence or CI alone.

## 11. Verification contract — target properties

Repository migration must establish mechanical checks for at least:

### Bootstrap / authority

- `AGENTS.md + docs/index.md + docs/roadmap.md <= 20 KiB`;
- `docs/roadmap.md` is the sole mutable stage/status/allowed-work/next-action authority;
- `README.md` is landing-only;
- default fresh-actor pack is five files or fewer;
- durable current documents are reachable from `docs/index.md` or a routed child index;
- current relative links resolve;
- no durable authority depends on `docs/work/**`.

### Temporary/bloat controls

Merge candidates and `main` contain no:

- `docs/work/**`;
- `docs/superpowers/**`;
- permanent AI-dialog/session-handoff/review-round artifacts;
- duplicate mutable roadmaps/status surfaces;
- active archive/old trees used as a second current authority.

### Decision/provenance controls

- current decision dispositions are valid/discoverable;
- important unique unmerged provenance remains reachable while a current consumer requires it;
- retired controls have proof that their subject population is zero or replacement coverage is complete.

### Guard quality

- every material behavioral guard has a deterministic negative control or equivalent falsifier;
- PR diff/whitespace checks compare the intended base to candidate;
- review-branch isolation is mechanically checked against the exact candidate;
- while implementation is blocked, allowed live top-level/source surfaces are explicitly allowlisted rather than relying only on old-name denylists, unless a documented local exception proves a smaller equivalent control.

### Required aggregate check

The migrated repository must expose at least one required aggregate verification status protecting `main`. The exact check name/tooling is a migration decision; no current green docs workflow is relabeled by implication.

## 12. Documentation migration safety law

No current Aurora file is removed solely because its path violates the target layout.

Before retiring a live surface:

1. identify its surviving current semantics;
2. assign those semantics to one target owner;
3. update inbound routes/references;
4. decide whether byte-level provenance still has a current consumer;
5. preserve a durable Git ref when required by the Repository Standard reachability law;
6. prove the new fresh-actor route reaches the current meaning;
7. only then delete/rehome the old live file.

Git history is archive only while the history needed by a current claim remains reachable.

## 13. Verification before completion

Never claim a gate, migration or implementation step complete without fresh Evidence against the exact target revision.

A green structural documentation check proves only the properties it actually tests. It does not prove Product correctness, architecture coherence, framework qualification, external integration or runtime behavior.
