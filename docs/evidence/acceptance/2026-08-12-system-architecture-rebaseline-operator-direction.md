---
id: DOC-AURORA-SYSTEM-ARCHITECTURE-REBASELINE-OPERATOR-DIRECTION
title: Aurora System Architecture Rebaseline Operator Direction
document_type: operator_direction
form: reference
authority: evidence
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - operator direction to pause Aurora implementation expansion and perform a program-level System Architecture Rebaseline
  - operator treatment of the non-canonical M0 R7 implementation candidate during the rebaseline
related:
  - DOC-AURORA-STATUS
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DESIGN-AURORA-SYSTEM-ARCHITECTURE-REBASELINE
  - DOC-AURORA-DECISIONS
source_revision: e7ca5ffb652fbbd68b35d4434506c58d26daf0e1
proposal_commit: cb6431883bf54cc5428e77e5202955aa9a0646e2
proposal_blob: 6ec1b39c7899593d4d1cd3ac5df033d6d4ce8dc2
recorded_at: 2026-08-12
last_reviewed: 2026-08-12
---

# Aurora System Architecture Rebaseline — Operator Direction

## 1. Context

After returning to Projeto Aurora following a pause, the operator identified a material program risk:

- Aurora is a large system of systems rather than one ordinary application;
- the accepted product constitution is broad and coherent, but later implementation could still fragment if cross-system technical boundaries are not mapped before expansion;
- important concerns such as APIs, authentication, authorization, data ownership, storage, memory, models, voice, observability, sandboxes, deployment and module communication must be connected before substantial code growth;
- planning must be deep enough to prevent architectural drift without becoming enterprise theater or indefinite overengineering;
- the software-development Harness should mature enough to help build Aurora, but Aurora runtime must not depend on that development Harness.

The operator explicitly stated that new Aurora implementation should not continue now. Controlled research and Architecture Spikes remain appropriate only when needed to decide a material question.

## 2. Design reviewed

The operator reviewed the complete proposed design recorded as:

```text
docs/superpowers/specs/2026-08-12-aurora-system-architecture-rebaseline-design.md
proposal commit: cb6431883bf54cc5428e77e5202955aa9a0646e2
proposal blob:   6ec1b39c7899593d4d1cd3ac5df033d6d4ce8dc2
version:         0.1.0
```

The design proposed:

- preserving accepted A0 product meaning and accepted ADRs within their stated scopes;
- freezing and preserving the existing non-canonical M0 R7 candidate;
- not authorizing R8;
- adding a program-level System Architecture Rebaseline inside the existing ACRM rather than creating another gate/lifecycle;
- creating a global Architecture Decision Landscape;
- classifying questions as `DECIDE`, `RESEARCH`, `SPIKE` or `DEFER`;
- identifying earliest consumers, evidence needs and decision owners;
- repairing tracking and accepted-ADR lifecycle wording drift;
- keeping Blueprint 12 free of premature technology choices;
- resuming implementation only after architecture review and explicit authority.

## 3. Operator decision

In direct response to the complete design, the operator stated:

```text
Aprovado
```

This records acceptance of the proposal semantics at the exact proposal blob above.

## 4. Authorized scope

The current program may:

- promote the approved rebaseline design lifecycle without changing its semantics;
- update `docs/product/CAPABILITY-REALIZATION-METHOD.md` to incorporate the program-level rebaseline inside R0–R8;
- create the global Architecture Decision Landscape;
- map accepted constraints, existing decisions, open questions, dependencies, earliest consumers and evidence needs;
- classify questions as `DECIDE`, `RESEARCH`, `SPIKE` or `DEFER`;
- update `STATUS`, `DECISIONS` and `WORKLOG` for truthful continuity;
- correct stale `Proposed` wording inside already accepted ADRs without changing the accepted decisions;
- run documentation generation/validation and adversarial review;
- prepare research or Architecture Spike proposals when the landscape finds that they are necessary.

## 5. Explicitly not authorized

This direction does **not** authorize:

- continuation of M0 R7 TASK-13;
- modification or expansion of the existing R7 implementation candidate;
- merge/promotion of the R7 branch to `main`;
- R7 independent Verdict or R8 closeout;
- new Aurora runtime implementation;
- M1+ implementation;
- AHDK, MNFS or Mastra adapter implementation;
- voice, memory, model-router, Presence, device or laboratory implementation;
- selection of authentication, policy, database, API, broker, observability, voice or model products merely because they are candidates;
- execution of an Architecture Spike without separate explicit authorization;
- creation of a new readiness gate, score, FSM or parallel governance method;
- reopening accepted constitutional meaning without a material Finding.

## 6. M0 R7 candidate treatment

The branch:

```text
feat/m0-r7-sovereign-core-20260810
head observed: 7ec999b093205a9d82eef2802eca60330d96e14d
```

is classified for current coordination as:

```text
FROZEN / PRESERVED / NON-CANONICAL
```

It may be inspected as executable evidence. It is not rejected, accepted, merged or authorized for expansion.

Any future continuation or promotion requires:

1. review against the completed System Architecture Rebaseline;
2. reconciliation of any material Finding;
3. a fixed target revision;
4. renewed explicit execution/closeout authority as applicable.

## 7. Stop boundary

```text
approved design promotion
→ documentary rebaseline package
→ mechanical validation
→ adversarial review
→ fixed revision presented to operator
→ STOP
```

Completion of the documentary package does not authorize Aurora implementation. The next substantive program work is architecture mapping and decision reduction, not code production.
