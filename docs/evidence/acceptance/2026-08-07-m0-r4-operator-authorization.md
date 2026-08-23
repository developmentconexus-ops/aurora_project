---
id: DOC-AURORA-M0-R4-OPERATOR-AUTHORIZATION
title: M0 ACRM R4 Operator Authorization
document_type: operator_authorization
form: evidence
authority: evidence
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - operator authorization to execute M0 ACRM R4 Architecture/Decision Readiness
related:
  - DOC-AURORA-STATUS
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC
  - REVIEW-AURORA-M0-R3-CAPABILITY-READINESS-2026-08-07
source_revision: d00cc1abfc2a41ac7e81e1f3478e188b3c5e9e52
authorized_at: 2026-08-07
last_reviewed: 2026-08-07
---

# M0 ACRM R4 — Operator Authorization

## 1. Authorization

After M0 ACRM R3 — Capability Readiness returned PASS and the project stopped at the R3 boundary, the operator stated:

> Vamos para próximo passo então

In direct continuity from the exact STATUS next action, this statement authorizes:

```text
M0 ACRM R4 — Architecture/Decision Readiness
```

only.

The operator subsequently reviewed and approved the R4 decision philosophy:

> Aprovado

The approved decision philosophy is:

```text
long-horizon exploration
+ evidence-bounded commitment
```

Interpretation:

- explore the long-term Aurora architectural horizon, including foreseeable M1+ constraints, to avoid local maxima;
- commit only to mechanisms whose current adequacy is supported by evidence;
- distinguish no-regret decisions, reversible mechanisms, high-lock-in decisions and not-yet-a-decision outcomes;
- prefer the smallest mechanism that survives the horizon analysis rather than the simplest or most feature-rich mechanism by default;
- preserve explicit migration/exit paths and option value where uncertainty remains material.

## 2. Authorized R4 work

This authorization permits:

- map the 15 open R4 decisions/uncertainties left by R3;
- analyze dependencies and lock-in among them;
- refresh focused research using current primary sources;
- compare candidate mechanisms against the accepted Blueprint, R2 requirements, R3 Spec, threat model and test plan;
- assess migration/rollback and operational burden;
- draft or accept/reject architecture decisions where documentary evidence is sufficient;
- specify exact Architecture Spikes where executable evidence remains necessary;
- perform adversarial R4 review and return `R4 PASS | FAIL | BLOCKED`.

## 3. Explicitly not authorized by this statement

This authorization does **not** authorize by implication:

- Architecture Spike execution;
- production or experimental implementation code;
- Capability implementation;
- Aurora Core implementation;
- AHDK implementation;
- MNFS integration;
- Mission Contract creation or acceptance (R5);
- Microdesign / Implementation Plan (R6);
- R7 execution/evidence;
- R8 milestone closeout.

If R4 determines that an Architecture Spike is materially required, the spike specification may be produced, but execution remains blocked until separately authorized by the operator.

## 4. Fixed baseline

R4 begins from canonical `main` at:

```text
d00cc1abfc2a41ac7e81e1f3478e188b3c5e9e52
```

No later repository state may be silently substituted as the R4 source baseline.
