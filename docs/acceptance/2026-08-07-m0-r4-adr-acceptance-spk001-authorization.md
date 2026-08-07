---
id: DOC-AURORA-M0-R4-ADR-ACCEPTANCE-SPK001-AUTHORIZATION
title: M0 R4 ADR Acceptance and SPK-001 Execution Authorization
document_type: operator_authorization
form: evidence
authority: evidence
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - operator acceptance of ADR-0003, ADR-0004, ADR-0005, ADR-0006 and ADR-0009
  - operator authorization to execute SPK-AURORA-M0-SOVEREIGN-STORE-001
related:
  - DOC-AURORA-STATUS
  - ADR-AURORA-0003
  - ADR-AURORA-0004
  - ADR-AURORA-0005
  - ADR-AURORA-0006
  - ADR-AURORA-0009
  - DESIGN-AURORA-M0-SOVEREIGN-STORE-SPIKE-001
source_revision: 36f46956bc275d0aec32b7e3ea4d959010fa9dcb
spike_spec_blob: 6ad7397d46208a0a9c762073d2c5239ceff4e056
recorded_at: 2026-08-07
last_reviewed: 2026-08-07
---

# M0 R4 — ADR Acceptance and SPK-001 Execution Authorization

## 1. Operator statement

On 2026-08-07 the operator stated:

> Aceito ADR-0003, ADR-0004, ADR-0005, ADR-0006 e ADR-0009, e autorizo a execução do SPK-AURORA-M0-SOVEREIGN-STORE-001 conforme a especificação canônica atual.

## 2. Accepted decisions

The statement explicitly promotes these ADRs to `accepted`:

```text
ADR-0003 — Go as the Initial Aurora Sovereign Core Runtime
ADR-0004 — M0 Local State and Execution Shape
ADR-0005 — M0 Portable Logical State and Export Envelope
ADR-0006 — M0 Observability Boundary
ADR-0009 — Mastra as Preferred First-Party Cognitive and Harness Runtime Substrate
```

Acceptance is limited to each ADR's stated decision scope. It does not authorize Aurora Core production implementation, Mastra implementation or later gates.

## 3. Authorized spike

Execution is authorized for exactly:

```text
spike_id: SPK-AURORA-M0-SOVEREIGN-STORE-001
canonical specification path: docs/design/SPK-AURORA-M0-SOVEREIGN-STORE-001.md
specification source revision: 36f46956bc275d0aec32b7e3ea4d959010fa9dcb
specification blob: 6ad7397d46208a0a9c762073d2c5239ceff4e056
```

The authorized execution may perform only the disposable experiment defined by that specification, including candidate setup, fault injection, crash/restart/recovery, supported backup/restore, corruption/incompatibility, migration fixture, measurements, evidence capture and adversarial review.

## 4. Explicitly not authorized

This statement does **not** authorize:

- `SPK-AURORA-M0-OWNER-TRUST-002` execution;
- acceptance of ADR-0007 or ADR-0008 before their required evidence;
- M0 ACRM R5 or later gates;
- Mission Contract creation/approval;
- Microdesign/R6 production implementation planning;
- Aurora Core production implementation;
- AHDK implementation;
- MNFS integration;
- Mastra implementation;
- promotion of spike code into production.

## 5. Stop boundary

```text
record accepted ADRs + exact SPK-001 authority
→ execute SPK-001 only
→ review executable evidence
→ update ADR-0007 as evidence supports
→ STOP before SPK-002 unless separately authorized
```
