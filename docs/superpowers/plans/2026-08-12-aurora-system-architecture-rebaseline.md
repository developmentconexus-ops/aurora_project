---
id: PLAN-AURORA-SYSTEM-ARCHITECTURE-REBASELINE
title: Aurora System Architecture Rebaseline Implementation Plan
document_type: implementation_plan
form: reference
authority: design
status: active
version: 1.0.0
owners:
  - developmentconexus-ops
related:
  - DESIGN-AURORA-SYSTEM-ARCHITECTURE-REBASELINE
  - DOC-AURORA-SYSTEM-ARCHITECTURE-REBASELINE-OPERATOR-DIRECTION
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DOC-AURORA-STATUS
last_reviewed: 2026-08-12
---

# Aurora System Architecture Rebaseline Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: execute this plan task-by-task with an independent review before any completion claim. This plan changes documentation and coordination only; it must not modify Aurora runtime code or the frozen M0 R7 branch.

**Goal:** Establish the approved program-level System Architecture Rebaseline, freeze the non-canonical M0 R7 candidate, restore truthful repository continuity and create the first global architecture decision landscape without selecting new technology.

**Architecture:** The existing ACRM R0–R8 lifecycle remains authoritative. A program-level rebaseline supplies cross-system constraints and a `DECIDE / RESEARCH / SPIKE / DEFER` landscape to later capability gates, while accepted ADRs retain explicit scope and Blueprint 12 remains the logical constitutional owner.

**Tech Stack:** Markdown, YAML frontmatter, existing Aurora documentation generator/validator, GitHub branch/PR review.

## Global constraints

- Canonical starting revision is `e7ca5ffb652fbbd68b35d4434506c58d26daf0e1` on `main`.
- Work occurs only on `docs/system-architecture-rebaseline-20260812`.
- Preserve `feat/m0-r7-sovereign-core-20260810` at observed head `7ec999b093205a9d82eef2802eca60330d96e14d`; do not modify, merge or promote it.
- Do not continue M0 R7 TASK-13 and do not authorize R8.
- Do not implement Aurora runtime, M1+, AHDK, MNFS, Mastra adapters, voice, memory, models, Presence, devices or laboratory functions.
- Do not select an authentication product, policy engine, database, API style, event broker, observability backend, voice provider or model provider.
- Do not create a new ACRM gate, score, FSM, authority hierarchy or parallel lifecycle.
- Keep accepted ADRs accepted within their stated scopes; correct lifecycle wording only where explicitly planned.
- Do not modify Blueprint 12 unless a material constitutional architecture defect is found. None is admitted by this plan.
- A green validator proves structural consistency only; adversarial review and operator review remain separate.

---

## Task 1 — Integrate System Architecture Rebaseline into ACRM

**Files:**
- Modify: `docs/product/CAPABILITY-REALIZATION-METHOD.md`

**Produces:**
- one program-level architecture activity inside the existing R0–R8 method;
- exact triggers, inputs, outputs, classifications, guardrails and gate relationships;
- no new readiness gate or implementation authority.

- [ ] Add `DESIGN-AURORA-SYSTEM-ARCHITECTURE-REBASELINE` and the operator-direction record to frontmatter relations.
- [ ] Increment the method version and `last_reviewed` date without altering existing R0–R8 semantics.
- [ ] Add a section before the gate definitions named `Program-level System Architecture Rebaseline`.
- [ ] Define triggers: multi-subsystem expansion, architecture drift, material cross-capability finding, implementation evidence contradicting global assumptions or explicit operator direction.
- [ ] Define inputs: Product Blueprint, roadmap, accepted ADRs, current decision/research maps, fixed implementation/spike evidence and current status.
- [ ] Define outputs: system context, logical ownership/boundaries, decision landscape, earliest consumers, evidence obligations, Findings and authorization boundary.
- [ ] Define `DECIDE`, `RESEARCH`, `SPIKE` and `DEFER` exactly as accepted by the design.
- [ ] State that capability R4 consumes and revalidates the landscape but still requires accepted decision owners where material.
- [ ] State that rebaseline completion does not authorize R7, implementation or a Product Milestone closeout.
- [ ] Self-review for accidental new lifecycle, hidden technology selection or conflict with M1–M10 principles.

## Task 2 — Publish the initial global Architecture Decision Landscape

**Files:**
- Create: `docs/design/AURORA-SYSTEM-ARCHITECTURE-DECISION-LANDSCAPE.md`

**Produces:**
- one proposed global map of accepted constraints, existing scoped decisions, open questions, dependencies and treatment;
- an exact next architecture workflow rather than a technology shopping list.

- [ ] Add valid frontmatter with authority `design`, status `proposed`, version `0.1.0`, fixed source revision and relations to Blueprints 03–13, ACRM, accepted ADR index and the approved rebaseline design.
- [ ] Explain the gap between accepted logical architecture and capability-specific technical commitment.
- [ ] Record global invariants that must not be lost: Aurora-owned semantics, one owner per concept, memory below authority, non-transitive grants, framework-neutral Core, logical modularity before distribution and effect enforcement outside models.
- [ ] Add the 28 accepted architecture clusters from the design.
- [ ] For every cluster, record at minimum: current constraint, existing decision/non-decision, concrete open question, earliest consumer and treatment.
- [ ] Classify only near-horizon cross-system mapping work as current; mark distant Voice/Presence/Lab selections `DEFER` with milestone triggers.
- [ ] Explicitly scope M0 Go/SQLite/JSON-JCS/OTel/owner-root decisions and Mastra preferred-first posture.
- [ ] Add dependency clusters for identity→authority→credentials/effects, data classes→storage→memory, contracts→bindings→AHDK, and telemetry→audit/evidence.
- [ ] Add a decision-promotion workflow and a stop rule preventing speculative code.
- [ ] Run a placeholder and ambiguity self-review.

## Task 3 — Restore truthful current coordination

**Files:**
- Replace: `docs/tracking/STATUS.md`

**Produces:**
- a concise current-state owner rather than a historical ledger;
- explicit R7 freeze and rebaseline authorization boundary.

- [ ] Preserve the stable accepted A0/M0 R0–R6 history by reference rather than repeating every historical detail.
- [ ] Record canonical `main` head `e7ca5ffb652fbbd68b35d4434506c58d26daf0e1` as R6 PASS baseline.
- [ ] Record that R7 authorization was received and produced a non-canonical candidate branch, but no independent R7 Verdict or R8 closeout exists.
- [ ] Classify the candidate branch as `FROZEN / PRESERVED / NON-CANONICAL` and record observed head `7ec999b093205a9d82eef2802eca60330d96e14d`.
- [ ] Set current program mode to `SYSTEM ARCHITECTURE REBASELINE`.
- [ ] List exactly what is authorized: documentation, architecture mapping, current primary-source research, decision proposals and separately authorized spikes.
- [ ] List exactly what is prohibited: new runtime implementation, TASK-13 continuation, R7 promotion, R8, M1+, AHDK/MNFS/Mastra adapter implementation and ungoverned stack selection.
- [ ] Define the exact next action as completing, validating and adversarially reviewing the global landscape, then presenting a fixed revision for operator review.
- [ ] State that completion of the rebaseline does not automatically resume implementation.

## Task 4 — Update decision discovery and chronology

**Files:**
- Modify: `docs/tracking/DECISIONS.md`
- Modify: `docs/tracking/WORKLOG.md`

**Produces:**
- discoverable accepted methodological direction;
- open technical questions without false promotion;
- a durable record of the R7 pause and rebaseline selection.

- [ ] Add one accepted decision that Aurora uses a program-level System Architecture Rebaseline before further multi-subsystem implementation expansion.
- [ ] Add one accepted decision that the Development Harness may build/verify Aurora but is not a sovereign runtime dependency.
- [ ] Add one accepted decision that the R7 candidate remains preserved and non-canonical pending rebaseline review.
- [ ] Expand the open-decision index to point to the global landscape rather than prematurely naming products.
- [ ] Append a 2026-08-12 worklog entry recording the operator concern, design approval, R7 freeze, authorized documentary scope, prohibited implementation and exact next action.
- [ ] Preserve all earlier decision/history entries unchanged except version/date/frontmatter relations needed for this update.

## Task 5 — Repair accepted ADR lifecycle wording drift

**Files:**
- Modify: `docs/adr/0003-m0-go-core-runtime.md`
- Modify: `docs/adr/0004-m0-local-state-execution-shape.md`
- Modify: `docs/adr/0009-mastra-cognitive-harness-runtime.md`

**Produces:**
- accepted ADR bodies that no longer describe their governing decisions as merely proposed;
- no change to decision scope, alternatives, consequences or reconsideration triggers.

- [ ] Replace the stale decision lead-in `Proposed:` with accepted lifecycle wording in ADR-0003.
- [ ] Replace the stale decision lead-in `Proposed:` with accepted lifecycle wording in ADR-0004.
- [ ] Replace the stale decision lead-in `Proposed:` with accepted lifecycle wording in ADR-0009.
- [ ] Preserve every substantive decision sentence, scope boundary and consequence.
- [ ] Update only version/last-reviewed metadata necessary to identify the lifecycle wording repair.
- [ ] Compare the final decision text against the pre-change accepted text and record that no semantic change occurred.

## Task 6 — Validate the complete documentary package

**Files:**
- Temporary validation workflow only if required because the canonical Documentation workflow does not run on the rebaseline branch; remove or generalize it before final review so no branch-specific permanent machinery remains.

**Required commands in GitHub Actions:**

```bash
python scripts/generate_docs.py --output-root "$RUNNER_TEMP/aurora-generated"
python scripts/validate_docs.py \
  --generated-root "$RUNNER_TEMP/aurora-generated" \
  --report "$RUNNER_TEMP/aurora-generated/docs-validation-report.json"
git diff --check
```

- [ ] Run generated-projection freshness without modifying Blueprint sources.
- [ ] Run full documentation validation and inspect the report, not only the job conclusion.
- [ ] Verify unique IDs, valid relations, frontmatter, local links, accepted-document placeholders and generated projection freshness.
- [ ] Search changed files for `TODO`, `TBD`, fake success, accidental technology selection, new lifecycle vocabulary and implementation authorization.
- [ ] Confirm no runtime/code/config/schema/dependency files changed.
- [ ] Fix every admitted structural error before review.

## Task 7 — Perform adversarial review

**Files:**
- Create: `docs/reviews/2026-08-12-system-architecture-rebaseline-review.md`

**Produces:**
- fixed-revision findings and verdict for the documentary package;
- explicit limitations and next action;
- no Product Milestone or R7/R8 verdict.

- [ ] Review against the accepted design proposal blob and operator direction.
- [ ] Test for duplicate governance, architecture theater, infinite planning, hidden implementation, framework capture and M0-scope globalization.
- [ ] Verify every new open question has an owner, earliest consumer, treatment and evidence path.
- [ ] Verify deferred items have concrete reconsideration triggers.
- [ ] Verify STATUS can orient a fresh session without the current conversation.
- [ ] Record findings as `BLOCKING`, `MATERIAL_NON_BLOCKING` or `MINOR` with remediation status.
- [ ] Issue only a documentary rebaseline readiness verdict: `PASS`, `FAIL` or `BLOCKED`.
- [ ] State that the review is not an R7 Verdict, R8 closeout or implementation authorization.

## Task 8 — Package for operator review

**Files:**
- Update final frontmatter/version references only where validation/review introduces fixed revision metadata.
- Open a draft pull request from `docs/system-architecture-rebaseline-20260812` to `main`.

- [ ] Compare branch against `main` and confirm the diff is documentation-only.
- [ ] Confirm the frozen R7 branch head is unchanged.
- [ ] Confirm the final validation run targets the final branch revision or document the exact final no-semantic marker relationship.
- [ ] Open a draft PR summarizing accepted direction, files changed, validation evidence, findings, limitations and explicit non-authority.
- [ ] Stop for operator review; do not merge.
