---
id: DOC-AURORA-M0-R2-OPERATOR-AUTHORIZATION
title: M0 ACRM R2 Operator Authorization
document_type: operator_authorization
form: evidence
authority: evidence
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - operator authorization to execute M0 ACRM R2 Requirements only
related:
  - DOC-AURORA-STATUS
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DOC-AURORA-CAP-SOVEREIGN-CORE-APPLICABILITY
  - REVIEW-AURORA-M0-R1-APPLICABILITY-2026-08-07
last_reviewed: 2026-08-07
---

# M0 ACRM R2 — Operator Authorization

## Authorization context

Immediately before this authorization, the canonical project state was:

```text
M0 ACRM R0 — PASS
M0 ACRM R1 — PASS
M0 ACRM R2 — NOT AUTHORIZED / awaiting explicit operator authorization
R3 and later — NOT AUTHORIZED BY IMPLICATION
```

The assistant explained that R2 would transform the 127 constitutional source rows active after R1 (`78 APPLIES + 49 PARTIALLY_APPLIES`) into precise, atomic and verifiable requirements for `CAP-SOVEREIGN-CORE`, while leaving implementation mechanisms, stack, Architecture Spikes, Capability/System design, Mission Contract, Microdesign and implementation outside the authorization.

## Exact operator statement

On 2026-08-07, Leandro stated:

> Entendi continue: Autorizo o M0 ACRM R2 — Requirements.

## Interpretation

This statement explicitly authorizes execution and closeout of:

```text
M0 ACRM R2 — Requirements
```

for:

```text
CAP-SOVEREIGN-CORE
```

It authorizes:

- derivation of atomic Capability requirements from the R1-active constitutional sources;
- requirement IDs, statements, source traceability, rationale and risk classification;
- verification direction for every requirement;
- identification of open decisions/spikes without selecting or executing them;
- duplicate/conflict analysis and a source-to-requirement coverage matrix;
- adversarial R2 review and R2 PASS | FAIL | BLOCKED verdict;
- documentation validation and tracking/evidence updates required to record the gate.

It does **not** authorize:

- ACRM R3 or later gates;
- a Capability/System Spec or detailed domain/system design;
- threat-model execution beyond recording R2 security obligations and R3 dependency;
- Architecture Spike execution;
- stack, language, runtime, database, storage, schema, protocol, topology or backend selection;
- Mission Contract;
- Microdesign/Implementation Plan;
- Aurora Core implementation;
- AHDK implementation;
- MNFS integration.

## Fixed source baseline

R2 begins from the canonical `main` revision that recorded the R1 PASS closeout:

```text
495b712142d7c3d722da2298f7a0b060707f9f5e
```

All R2 source interpretation must remain traceable to this fixed baseline unless a material divergence is explicitly raised and handled through its owning artifact.
