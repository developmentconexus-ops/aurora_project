---
id: DOC-AURORA-STATUS
title: Aurora Project Status
document_type: project_status
form: reference
authority: tracking
status: current
version: 0.31.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - current Aurora project phase
  - current authorization boundary
  - current blockers and immediate next action
related:
  - DOC-AURORA-DOCUMENTATION-MAP
  - DOC-AURORA-PRODUCT-INDEX
  - DOC-AURORA-ROADMAP
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DESIGN-AURORA-SYSTEM-ARCHITECTURE-REBASELINE
  - DESIGN-AURORA-SYSTEM-ARCHITECTURE-DECISION-LANDSCAPE
  - DESIGN-AURORA-TECHNICAL-ARCHITECTURE-BASELINE-MAP
  - PLAN-AURORA-TECHNICAL-ARCHITECTURE-BASELINE
  - DOC-AURORA-TECHNICAL-ARCHITECTURE-MAP-OPERATOR-ACCEPTANCE
  - REVIEW-AURORA-TECHNICAL-ARCHITECTURE-BASELINE-MAP-2026-08-12
  - DESIGN-AURORA-SAR-A1-STAGE-A-AVAILABILITY-ACTIVATION
  - DOC-AURORA-DECISIONS
  - DOC-AURORA-WORKLOG
last_reviewed: 2026-08-12
---

# Aurora Project Status

## 1. Current summary

- **Project:** Projeto Aurora
- **Canonical branch:** `main`
- **Canonical System Architecture Rebaseline merge:** `59f5819de97208bea88fdd3c2b30e13f417c2963`
- **Current documentary branch:** `docs/technical-architecture-baseline-20260812`
- **Draft PR:** `#4 — docs: establish Aurora Technical Architecture Baseline map`
- **PR target:** `main`
- **PR state:** OPEN / DRAFT / NOT MERGED
- **A0 Product/Discovery/Architecture baseline:** ACCEPTED / MERGED
- **ADR-0001..0009:** ACCEPTED within exact scope
- **M0 ACRM R0–R6:** PASS
- **M0 R7 candidate:** FROZEN / PRESERVED / NON-CANONICAL
- **M0 R7 independent Verdict:** NOT ISSUED
- **M0 R8:** NOT AUTHORIZED / NOT PERFORMED
- **System Architecture Rebaseline:** ACCEPTED / MERGED
- **Current program:** `AURORA TECHNICAL ARCHITECTURE BASELINE`
- **Accepted work map:** `TA-01` through `TA-08`
- **Current active tranche:** `TA-01 + TA-02`
- **Current tranche state:** DISCOVERY / DESIGN AUTHORIZED
- **Map package review:** PASS FOR OPERATOR / PR REVIEW
- **Aurora implementation:** PAUSED

## 2. Current direction

Aurora product meaning is already deeply defined. The current objective is not another broad product-discovery cycle and not further decomposition of narrow Presence behavior.

The current objective is:

> Build the technical architecture map that will govern Aurora components, modules, runtime/process topology, repositories, contracts, data, security, cognition/Harness integration and operation before implementation resumes.

The operator corrected a priority drift during SAR-A1:

```text
useful Stage A Presence/session detail
→ preserve as accepted downstream constraint
→ DEFER further micro-policy exploration

cross-system technical architecture
→ current priority
```

Questions are current only when they change a structural boundary, ownership, contract, process/runtime, security/effect boundary, data role or the next implementation decision.

## 3. Accepted Technical Architecture Baseline map

The accepted dependency order is:

```text
TA-01 Logical modules and canonical ownership
TA-02 Process, runtime and evolutionary topology
TA-03 Repository, source and build architecture
TA-04 Contracts, APIs, events and communication
TA-05 Data, storage, portability and lifecycle architecture
TA-06 Identity, authentication, authorization, policy and secrets
TA-07 Brain, models, memory and Harness integration
TA-08 Configuration, observability, deployment and operation
```

Governing map:

```text
docs/design/AURORA-TECHNICAL-ARCHITECTURE-BASELINE-MAP.md
```

Documentary execution plan:

```text
docs/superpowers/plans/2026-08-12-aurora-technical-architecture-baseline.md
```

The map is not a new ACRM gate, lifecycle or score. It is the ordered program-level architecture work inside the accepted System Architecture Rebaseline.

## 4. Current active tranche — TA-01 + TA-02

The operator authorized discovery/design for:

```text
TA-01 — Logical modules and canonical ownership
+
TA-02 — Process, runtime and evolutionary topology
```

The tranche must produce:

1. technical component catalog;
2. module responsibility matrix;
3. canonical entity/data ownership matrix;
4. allowed and forbidden dependency directions;
5. 2–3 coherent module/runtime approaches;
6. Stage A process/runtime topology;
7. always-active versus on-demand responsibilities;
8. Go, TypeScript/Mastra and future-runtime scope map;
9. failure-domain and restart ownership;
10. Stage B evolution path;
11. `DECIDE / RESEARCH / SPIKE / DEFER` register;
12. exact inputs required before TA-03 repository architecture.

TA-01 and TA-02 are coupled because process placement must follow coherent ownership, while proposed modules must be tested against real runtime/failure/security boundaries.

## 5. Review and validation evidence

Fixed adversarial-review target:

```text
5f22d83ba4bf7e893c4857bf887eabbe4aef1feb
```

Review record:

```text
docs/reviews/2026-08-12-technical-architecture-baseline-map-review.md
review commit: ec8e5582cd374b80102305f360af5fb4304c4ddf
verdict: PASS FOR OPERATOR / PR REVIEW
```

Validation:

```text
Documentation run 31618020814 — SUCCESS
review-commit run 31618361364 — SUCCESS
canonical documents / IDs at reviewed target: 130 / 130
```

The review resolved the priority-drift, framework-first and fresh-session-continuity risks. It is an adversarial authoring-session review, not an independent technical-architecture acceptance verdict.

## 6. Current authorized work

The current authorization permits:

- operator/PR review and requested documentary corrections;
- inspection of accepted Blueprint, ADR, research and M0 artifacts relevant to TA-01/TA-02;
- inspection of the frozen R7 candidate as evidence only;
- deriving and challenging candidate technical components;
- defining canonical responsibility and data ownership;
- defining allowed/forbidden dependencies;
- comparing 2–3 complete Stage A topology approaches;
- mapping Stage B evolution without designing Stage B in full;
- current primary-source research only where it changes a near decision;
- Architecture Spike specifications where runtime evidence would be required;
- preparing a fixed, reviewable TA-01/TA-02 design package;
- documentation validation and adversarial review.

## 7. Explicitly prohibited

The current authorization does **not** permit:

- automatic merge of PR #4 without explicit operator direction;
- new Aurora runtime implementation;
- continuation, modification, merge or promotion of the frozen M0 R7 candidate;
- an M0 R7 acceptance Verdict;
- M0 R8 closeout;
- M1+ implementation;
- Architecture Spike execution without separate exact authorization;
- creating or restructuring production repositories;
- selecting monorepo versus polyrepo before TA-01/TA-02 review;
- selecting a universal API protocol;
- selecting a new database/store before TA-05 requirements;
- selecting Keycloak, Zitadel, Ory, Authentik, SPIFFE, OPA, Cedar, Vault or equivalents before TA-06 modeling;
- implementing AHDK, MNFS or a Mastra adapter;
- implementing Brain, memory, Voice, Presence, model routing or observability systems;
- treating M0 Go/SQLite/JSON-JCS/OTel decisions as universal;
- returning to detailed Presence/session policy unless it materially changes TA-01/TA-02;
- creating another readiness lifecycle, score or authority hierarchy.

## 8. Preserved Stage A constraints

The following are accepted but not the current discussion priority:

- one Leandro-controlled workstation is the Stage A sovereign host and first Presence;
- minimum Core and activation responsibilities may remain available;
- heavy cognition starts on demand;
- activation belongs to Presence semantics;
- button/UI/hotkey are baseline activation classes;
- local wake word is optional;
- activation is not authentication or authority;
- while locked, Aurora may acknowledge availability but requires unlock before private interaction.

These are consumed as TA-02 inputs and future Presence/Voice constraints. Further user/session-policy decomposition is `DEFER`.

## 9. Current blocker

```text
PR #4 CANONICAL PROMOTION: AWAITING OPERATOR DIRECTION
TA-01/TA-02 DESIGN: NOT YET PRODUCED OR REVIEWED
TA-03 FINALIZATION: BLOCKED ON TA-01/TA-02
AURORA IMPLEMENTATION RESUMPTION: BLOCKED
ARCHITECTURE SPIKE EXECUTION: NOT AUTHORIZED
```

This is an intentional architecture-readiness block, not a rejection of Aurora or of the frozen M0 candidate.

## 10. Immediate next action

Two activities may proceed without implying implementation:

```text
A. operator reviews draft PR #4 for canonical promotion

B. TA-01/TA-02 discovery dialogue begins
   → derive minimum coherent technical component set
   → compare 2–3 complete module/runtime approaches
   → recommend one with explicit trade-offs
```

Exact TA-01/TA-02 question:

> What is the minimum set of technical components required to preserve Aurora's accepted ownership boundaries, and which of those components should share or cross a process boundary in Stage A?

No repository, stack or runtime implementation decision follows automatically.

## 11. Fresh-session mandatory orientation

After `AGENTS.md`, a new technical-architecture session must read:

1. this STATUS;
2. `docs/design/AURORA-TECHNICAL-ARCHITECTURE-BASELINE-MAP.md`;
3. `docs/superpowers/plans/2026-08-12-aurora-technical-architecture-baseline.md`;
4. `docs/reviews/2026-08-12-technical-architecture-baseline-map-review.md`;
5. `docs/design/AURORA-SYSTEM-ARCHITECTURE-DECISION-LANDSCAPE.md`;
6. relevant Blueprint sections and scoped ADRs;
7. current TA-01/TA-02 artifacts.

The session must not restart product discovery, continue Presence micro-policy by default, choose frameworks first or implement code.
