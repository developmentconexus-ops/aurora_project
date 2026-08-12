---
id: DOC-AURORA-TECHNICAL-ARCHITECTURE-MAP-OPERATOR-ACCEPTANCE
title: Aurora Technical Architecture Baseline Map Operator Acceptance
document_type: operator_acceptance
form: reference
authority: evidence
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - operator acceptance of the Aurora Technical Architecture Baseline Map
  - operator correction from Presence micro-policy detail to cross-system technical architecture
  - authorization to begin the TA-01 and TA-02 discovery/design tranche
related:
  - DESIGN-AURORA-TECHNICAL-ARCHITECTURE-BASELINE-MAP
  - PLAN-AURORA-TECHNICAL-ARCHITECTURE-BASELINE
  - DESIGN-AURORA-SYSTEM-ARCHITECTURE-REBASELINE
  - DESIGN-AURORA-SYSTEM-ARCHITECTURE-DECISION-LANDSCAPE
  - DESIGN-AURORA-SAR-A1-STAGE-A-AVAILABILITY-ACTIVATION
  - DOC-AURORA-STATUS
  - DOC-AURORA-DECISIONS
proposal_base_revision: a520df862ad39d8e17e8bd17a80da8b8b2f1a900
accepted_design_commit: b8f3588eaeb14dc4e15736ceda46178a81571456
recorded_at: 2026-08-12
last_reviewed: 2026-08-12
---

# Aurora Technical Architecture Baseline Map — Operator Acceptance

## 1. Context

During SAR-A1 discovery, the conversation began decomposing detailed Presence behavior such as activation while the workstation was locked and who might use an unlocked workstation session.

The operator identified that this had crossed the current materiality boundary. The program had already invested heavily in defining what Aurora is, its relationship model, authority, memory, Core, Harnesses, Presence direction, methodology and roadmap.

The operator clarified that the current goal is now:

> Build the technical map that explains how Aurora will be structured and constructed: components, modules, runtimes, repositories, contracts, APIs, data stores, authentication, security, Brain/model/Harness integration, observability and deployment.

The operator explicitly rejected spending the current architecture cycle on highly specific user/session behavior that does not change the next structural build decisions.

## 2. Accepted correction

The accepted correction is:

```text
preserve useful Stage A Presence constraints
→ stop further Presence/session micro-policy exploration
→ classify those details DEFER until a consuming Capability needs them
→ begin the cross-system Technical Architecture Baseline
```

The earlier Stage A decisions are not rejected. They remain downstream constraints for runtime topology and future Presence/Voice work.

## 3. Accepted technical architecture sequence

The operator reviewed and approved the following ordered map:

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

The sequence is dependency-ordered. It does not require every area to become globally final before bounded implementation can resume, but it prevents a later area from silently inventing an earlier area's ownership or topology.

## 4. Accepted decision method

For each technical area, the operator accepts this flow:

```text
boundary and responsibility
→ requirements and invariants
→ first real consumer
→ alternatives
→ current research where material
→ trade-offs and migration
→ DECIDE | RESEARCH | SPIKE | DEFER
→ owning ADR / Specification / Standard / Contract
→ resulting implementation impact
```

Technology must follow the problem definition.

Examples such as Keycloak, PostgreSQL, gRPC, OPA, Vault, vector stores, model providers and Voice frameworks are candidates to evaluate in their correct work area—not preselected stack commitments.

## 5. Current authorization

The operator stated:

```text
Aprovado, vamos seguir isso e esse mapa, documente tudo para não perder direção em uma outra sessão, por exemplo.
```

This accepts `DESIGN-AURORA-TECHNICAL-ARCHITECTURE-BASELINE-MAP` v0.1.0 and authorizes discovery/design for the first coupled tranche only:

```text
TA-01 — Logical modules and canonical ownership
+
TA-02 — Process, runtime and evolutionary topology
```

The tranche may:

- inspect accepted Blueprint, ADR, M0 and frozen-R7 evidence;
- define candidate technical components;
- compare 2–3 coherent module/runtime approaches;
- define canonical responsibility and data ownership;
- define allowed/forbidden dependencies;
- define Stage A and Stage B process/runtime hypotheses;
- identify research and Spike needs;
- produce an operator-reviewable technical architecture proposal.

## 6. Not authorized

This acceptance does not authorize:

- Aurora runtime implementation;
- continuation or promotion of the frozen M0 R7 candidate;
- M0 R8 closeout;
- M1+ implementation;
- Architecture Spike execution;
- repository creation or restructuring;
- choosing monorepo versus polyrepo before TA-01/TA-02;
- selecting a universal API protocol;
- selecting new databases or storage products;
- selecting authentication, policy or secrets products;
- implementing AHDK, MNFS or a Mastra adapter;
- implementing Brain, memory, Voice, Presence or observability systems;
- treating candidate technology examples as accepted decisions.

## 7. Fresh-session continuity requirement

Any new session must discover from `STATUS.md` and the accepted map that:

```text
current objective
→ technical architecture baseline, not more product-definition dialogue

current work
→ TA-01 + TA-02

current output
→ module ownership + process/runtime topology

current prohibitions
→ no implementation or stack choice by implication

next question
→ minimum coherent technical components and Stage A process boundaries
```

The map and current STATUS must remain more authoritative for current work priority than a historical conversation transcript.
