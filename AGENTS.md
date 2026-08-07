# AGENTS.md

This file is the minimum mandatory bootstrap for any agent or fresh session working on Projeto Aurora.

## 1. Mandatory read order

Always begin with:

1. `docs/tracking/STATUS.md`
2. `docs/DOCUMENTATION-MAP.md`
3. `docs/product/README.md`
4. `docs/roadmap.md`
5. only the Product Blueprint sections relevant to the current scope
6. related requirements, research, ADRs, specs, contracts, designs and evidence

For A0 constitutional review, read the complete generated Product Blueprint and the discovery/coverage records.

Do not infer current authority from an old transcript, branch name, issue title or file existence.

## 2. Authority rules

- Product Blueprint owns constitutional product meaning.
- Accepted ADRs own specific decisions compatible with the constitution.
- Capability Specs own reusable capability behavior.
- Approved Mission Contracts own exact scoped commitments.
- Standards/Policies own methods, Golden Paths, enforcement profiles and waivers.
- Reference documents describe current machinery.
- Evidence proves an observation against a target.
- Tracking reports current coordination only.
- Research and historical records inform; they do not govern.
- Generated projections follow their canonical modular sources.

When sources conflict, do not choose silently. Raise `DOCUMENTATION_DIVERGENCE`, identify the canonical owner and block affected work when material.

## 3. Current phase and authorization

Mutable coordination state has exactly one owner:

```text
docs/tracking/STATUS.md
```

A fresh session MUST read `STATUS.md` for the selected Product Milestone, current ACRM gate, blockers, authorizations, prohibitions and exact next action. This file intentionally does not duplicate those mutable values.

Stable accepted baseline:

```text
A0 constitutional baseline: ACCEPTED / MERGED
ADR-0001 / ADR-0002: ACCEPTED
Stack selection from A0: NOT PERFORMED
Aurora Core implementation: NOT AUTHORIZED BY A0
AHDK implementation: NOT AUTHORIZED BY A0
Architecture Spike execution: NOT AUTHORIZED BY A0
MNFS integration: NOT AUTHORIZED BY A0
```

After A0, every readiness transition remains explicit. Selection of a Product Milestone authorizes only the readiness work recorded in `STATUS`; completing one ACRM gate does not authorize the next gate, and no implementation follows from file existence or absence of a prohibition.

Do not implement runtime, database, SDK, protocol binding, UI, model router, memory engine, device controller or MNFS adapter until:

1. A0 is explicitly accepted by Leandro;
2. the next Product Milestone is selected;
3. relevant ACRM R0–R6 gates pass;
4. exact implementation is separately authorized.

## 4. Aurora Capability Realization Method

Material work follows:

```text
R0 Constitutional Baseline
→ R1 Applicability
→ R2 Requirements
→ R3 Capability Readiness
→ R4 Architecture/Decision Readiness
→ R5 Contract Readiness
→ R6 Implementation Design Readiness
→ R7 Execution and Evidence
→ R8 Product Milestone Closeout
```

Read `docs/product/CAPABILITY-REALIZATION-METHOD.md` before planning a capability.

Rules:

- no material implementation without traceable approved intent;
- no approved requirement without evidence or explicit deferral;
- research, ADR, Spec, Contract, Plan, code and Evidence remain distinct;
- a plan cannot alter approved product intent;
- a Claim is not acceptance;
- Product Milestone closure requires its end-to-end Golden Proof;
- completing one gate does not authorize the next gate;
- material change triggers replan through the owning artifact.

## 5. Research conduct

- Use current primary sources: specifications, official docs, official repositories, standards and papers where appropriate.
- Record source ID, title, publisher, URL, access date, relevant version and supported claims.
- State limitations and disagreements.
- Separate fact, inference, recommendation and decision.
- Treat every framework/protocol/runtime as a candidate until accepted through the appropriate gate.
- Revalidate temporally unstable information when implementation approaches.
- Pair each material report with a `*.sources.json` manifest.

## 6. Working with the Product Blueprint

- Edit modular files under `docs/product/blueprint/`.
- Never edit `docs/product/PRODUCT-BLUEPRINT.md` directly.
- Edit Blueprint 14 rather than `docs/roadmap.md` directly.
- Regenerate projections with `python scripts/generate_docs.py`.
- Run validation through the documentation workflow or validator.
- Update `REQUIREMENTS-TRACEABILITY.md` when constitutional meaning changes.
- Update `DOCUMENTATION-COVERAGE.md` when discovery coverage changes.

## 7. Planning and implementation conduct

Before any future implementation:

- inspect the actual repository/code state;
- identify the exact accepted Contract and baseline;
- allocate requirements to exact criteria;
- produce a reviewed Microdesign/Implementation Plan;
- use test-first RED/GREEN where applicable;
- include cross-layer wiring, failures, recovery, observability, security and docs impact;
- do not introduce unrelated refactoring;
- stop and replan when a material decision appears during implementation.

The implementer must not be the sole authority accepting material work.

## 8. Evidence and status

At the end of material work:

- update `docs/tracking/STATUS.md`;
- append `docs/tracking/WORKLOG.md`;
- update `docs/tracking/DECISIONS.md` when a decision owner/status changes;
- update `docs/tracking/BACKLOG.md` for non-commitment ideas;
- record Findings, limitations and next action;
- declare documentation impact;
- preserve exact target revision and evidence references;
- do not claim completion without fresh verification.

## 9. Documentation checks

Current CI validates:

- expected files;
- frontmatter;
- unique IDs;
- relation targets;
- local links;
- source manifests and citations;
- requirement IDs;
- discovery coverage gaps;
- normative placeholders;
- generated Blueprint and Roadmap freshness.

A green documentation check proves structural consistency, not product correctness. Adversarial review and operator acceptance remain required.

## 10. Fresh-session obligation

A new session must be able to state, from the repository alone:

- what Aurora is;
- current phase and prohibitions;
- canonical owners;
- open versus accepted decisions;
- exact next action.

If this cannot be done, stop and repair documentation continuity before proceeding.
