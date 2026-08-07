---
id: DOC-AURORA-M0-R3-OPERATOR-AUTHORIZATION
title: M0 ACRM R3 Operator Authorization
document_type: operator_authorization
form: evidence
authority: evidence
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - operator authorization to execute M0 ACRM R3 Capability Readiness
related:
  - DOC-AURORA-STATUS
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DOC-AURORA-CAP-SOVEREIGN-CORE-REQUIREMENTS
  - REVIEW-AURORA-M0-R2-REQUIREMENTS-2026-08-07
source_revision: 9ea8adf5c115f54071d7e36e312695d19420d8b0
recorded_at: 2026-08-07
last_reviewed: 2026-08-07
---

# M0 ACRM R3 — Operator Authorization

## 1. Operator statement

On 2026-08-07, immediately after the repository recorded `M0 ACRM R2 — Requirements — PASS` and stopped at the R2 boundary, the operator stated:

> Autorizo o M0 ACRM R3 — Capability Readiness.

This record preserves that authorization as repository evidence.

## 2. Authorized scope

The statement authorizes execution of **M0 ACRM R3 — Capability Readiness** for:

```text
Capability: CAP-SOVEREIGN-CORE
Product Milestone: M0 — Sovereign Core Walking Skeleton
R3 source baseline: 9ea8adf5c115f54071d7e36e312695d19420d8b0
```

Authorized R3 work includes:

- produce the reusable Capability/System Spec required by ACRM;
- define the M0 domain model and logical ownership boundaries;
- define semantic contracts/schemas without selecting serialization or storage technology;
- define lifecycle/state semantics;
- define context/memory boundaries for the M0 slice;
- define authority semantics and the explicit absence of an M0 external-effect plane;
- produce the R3 threat model and privacy/data-classification analysis;
- define failure/recovery behavior;
- define observability/evidence semantics;
- define compatibility/migration behavior;
- produce a Capability test plan;
- allocate every R2 requirement to one or more Spec mechanisms and tests;
- identify R4 technical uncertainties/open decisions without selecting their mechanisms;
- conduct adversarial review, validation and emit `R3 PASS | FAIL | BLOCKED`;
- update tracking after the R3 verdict.

## 3. Explicitly not authorized

This authorization does **not** authorize:

- M0 ACRM R4 — Architecture/Decision Readiness;
- execution of any Architecture Spike;
- selection of Core language/runtime;
- selection of database, storage engine or state/event implementation pattern;
- selection of schema/serialization/code-generation stack;
- selection of process/deployment topology;
- selection of event/audit or telemetry backend;
- selection of backup/restore/migration technology;
- selection of a durable execution engine;
- creation/approval of the M0 Mission Contract (R5);
- Microdesign/Implementation Plan (R6);
- Aurora Core implementation (R7);
- AHDK implementation;
- MNFS integration.

Completing or passing R3 does not authorize R4 by implication.

## 4. Stop boundary

```text
R3 Capability Readiness
→ R3 PASS | FAIL | BLOCKED
→ STOP
→ await separate explicit operator authorization for R4
```
