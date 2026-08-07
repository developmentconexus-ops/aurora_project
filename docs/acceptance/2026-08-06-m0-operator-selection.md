---
id: DOC-AURORA-M0-OPERATOR-SELECTION
title: M0 Product Milestone Operator Selection
document_type: operator_decision_evidence
form: evidence
authority: evidence
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - operator selection of the first Product Milestone after A0
related:
  - DOC-AURORA-STATUS
  - DOC-AURORA-ROADMAP
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DOC-AURORA-DECISIONS
last_reviewed: 2026-08-06
---

# M0 Product Milestone Operator Selection

## 1. Context

After A0 was accepted and merged, the roadmap was reviewed adversarially to determine whether the first Product Milestone should remain M0 or whether M1/M2 should be promoted earlier.

The analysis compared:

- **M0 — Sovereign Core Walking Skeleton**;
- **M1 — Governed Conversation, Project Context and Memory**;
- **M2 — Capability Registry, AHDK Kernel and Reference Provider**.

The review concluded that M0 should remain first because it retires the foundational risk that Aurora is merely a running session whose identity, operational state and authority disappear on process restart. M1 depends on a sovereign state/authority foundation, while M2 should not cause the provider/SDK ecosystem to become the de facto Aurora Core.

## 2. Operator decision

On 2026-08-06, after receiving the milestone analysis, Leandro responded:

> “Aprovo”

The response directly answered the explicit question whether `M0 — Sovereign Core Walking Skeleton` should be selected as the first Product Milestone.

## 3. Verdict

```text
First Product Milestone: M0 — Sovereign Core Walking Skeleton
Selection: ACCEPTED
ACRM R0: AUTHORIZED TO BEGIN IN A NEW SESSION
R0 execution in the selecting session: NOT STARTED
```

## 4. What this decision establishes

The next readiness subject is M0 as defined by the accepted Capability Roadmap.

The M0 product outcome remains:

> A minimal Aurora Core persists identity, project state, authority snapshot and one interaction lifecycle across restart without depending on an external model or Harness as authority.

The accepted roadmap Golden Proof remains the directional proof target until R0–R6 refine the executable contract:

```text
initialize Aurora instance
→ create project
→ record accepted state and next action
→ terminate all Aurora processes
→ start fresh process
→ recover same identities/state
→ reject an invalid transition
→ export and restore state
```

## 5. What is not decided or authorized

This selection does **not** decide or authorize:

- implementation of Aurora Core;
- a programming language;
- database/storage technology;
- process topology;
- schema format;
- event-sourcing strategy;
- durable execution engine;
- policy engine;
- AHDK implementation;
- Architecture Spike execution;
- MNFS integration;
- Mission Contract scope;
- Microdesign or Implementation Plan;
- R1 or later gates by implication.

## 6. Next exact action

A fresh session must:

```text
read AGENTS.md and current STATUS
→ read the accepted M0 roadmap definition and ACRM
→ execute only ACRM R0 — Constitutional Baseline for M0
→ identify applicable constitutional sources, contradictions, missing ownership and readiness blockers
→ produce an R0 PASS | FAIL | BLOCKED verdict with repository citations
→ stop before R1 unless separately authorized
```

The purpose of the fresh session is to prove that M0 readiness can begin from repository state alone, not from this conversation.
