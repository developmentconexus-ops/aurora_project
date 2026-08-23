---
id: DOC-AURORA-TA-01-02-OPERATOR-ACCEPTANCE
title: TA-01/TA-02 Module and Runtime Topology Operator Acceptance
document_type: operator_acceptance
form: reference
authority: evidence
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - operator acceptance of TA-01 logical module and canonical ownership baseline
  - operator acceptance of TA-02 Stage A and Stage B runtime topology direction
  - operator acceptance of Approach C, G01, C01-C12, A05 and B01 semantics
related:
  - DESIGN-AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY
  - REVIEW-AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY-2026-08-12
  - DESIGN-AURORA-TECHNICAL-ARCHITECTURE-BASELINE-MAP
  - PLAN-AURORA-TECHNICAL-ARCHITECTURE-BASELINE
  - DOC-AURORA-STATUS
  - DOC-AURORA-DECISIONS
  - DOC-AURORA-WORKLOG
accepted_design_version: 0.5.0
accepted_design_commit: 965211ad421f13994285031bbcf04b7e943cf75e
operator_decision: ACCEPT
recorded_at: 2026-08-13
last_reviewed: 2026-08-13
---

# TA-01/TA-02 Module and Runtime Topology — Operator Acceptance

## 1. Decision

After the reviewed TA-01/TA-02 package was presented with the explicit gate `ACCEPT | REVISE | REJECT`, the operator responded:

> `Accept`

This is recorded as explicit operator **ACCEPT** of `DESIGN-AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY` v0.5.0, semantically fixed at:

```text
965211ad421f13994285031bbcf04b7e943cf75e
```

## 2. Accepted architecture

The acceptance covers the reviewed TA-01/TA-02 semantic design, including:

```text
G01 Contract Model Governance

C01 Identity and Relationship
C02 Project, World and Experiment State
C03 Intent, Mission and Delegation Control
C04 Authority and Policy
C05 Capability and Provider Registry
C06 Memory, Knowledge and Context
C07 Artifact, Observation and Evidence
C08 Presence and Interaction
C09 Attention and Proactivity                [future/deferred]
C10 Failure Intelligence and Evaluation      [future/deferred]
C11 Environment and Device Registry          [future/deferred]
C12 Audit and Exact History
```

and the application/runtime boundaries:

```text
A01 Core Application Coordination
A02 Cognitive Coordination
A03 Capability Fabric / Harness Integration
A04 Durable Execution Port
A05 Runtime Lifecycle Coordination
B01 Transport-neutral Provider Runtime Boundary Profile
E01 Effect Gateway family
E02 Credential Broker boundary
```

The accepted Stage A topology direction is **Approach C — Evolutionary Sovereign Host with one early provider-runtime seam**.

## 3. Acceptance meaning

The decision accepts:

- one small persistent Sovereign Host as the Stage A composition boundary;
- canonical ownership remaining modular even when physically co-located;
- a separate on-demand Cognitive Runtime at the first real consumer;
- provider-local state remaining non-canonical;
- G01 as the owner of cross-system contract semantics;
- C03 as owner of validated interpreted Intent before explicit Mission promotion;
- C12 as owner of Audit and L4 exact history;
- A05 as Aurora-side runtime lifecycle-policy owner;
- B01 lifecycle, identity, idempotency, cancellation, retry and reconciliation semantics before transport selection;
- Stage B evolution without changing Aurora domain identity or canonical ownership;
- environment-bound provider instances being re-registered/re-approved after migration.

## 4. Explicit non-authorization

This acceptance does **not** authorize:

- merging PR #5 by implication;
- Aurora runtime implementation;
- continuation or promotion of the frozen M0 R7 candidate;
- M0 R8;
- M1+ implementation;
- Architecture Spike execution;
- TA-03 repository strategy finalization;
- TA-04 protocol/binding finalization;
- TA-05 storage finalization;
- TA-06 authentication/policy/secrets product selection;
- TA-07 Mastra/Cognitive Runtime implementation;
- TA-08 supervisor/deployment/observability finalization;
- monorepo/polyrepo, universal language, transport, store, authentication, policy, secrets, supervisor, model or Voice product selection.

## 5. Promotion boundary

The accepted semantic revision may now receive lifecycle metadata, tracking updates and promotion evidence without changing the accepted architecture meaning.

```text
operator ACCEPT
→ lifecycle promotion on review branch
→ documentation validation
→ separate explicit merge/promotion authorization
→ canonical merge only after that authorization
```

Until merge authorization is separately given, PR #5 remains open and `main` remains canonical without this accepted package.
