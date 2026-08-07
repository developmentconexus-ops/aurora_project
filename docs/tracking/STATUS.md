---
id: DOC-AURORA-STATUS
title: Aurora Project Status
document_type: project_status
form: reference
authority: tracking
status: current
version: 0.8.0
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
  - DOC-AURORA-A0-OPERATOR-ACCEPTANCE
  - DOC-AURORA-GP-A0-FRESH-001-EVALUATION
  - DOC-AURORA-M0-OPERATOR-SELECTION
last_reviewed: 2026-08-06
---

# Aurora Project Status

## 1. Current summary

- **Project:** Projeto Aurora
- **Canonical branch:** `main`
- **A0:** ACCEPTED and MERGED
- **A0 merge commit:** `f22085d97e198d99e89d52221b7b26d59d49bc12`
- **ADR-0001:** ACCEPTED
- **ADR-0002:** ACCEPTED
- **Independent Fresh-Session Golden Proof:** PASS — 100/100, zero hard failures
- **First Product Milestone:** `M0 — Sovereign Core Walking Skeleton` — SELECTED by operator
- **Current readiness gate:** ACRM R0 — Constitutional Baseline
- **R0 execution:** AUTHORIZED TO BEGIN IN A FRESH SESSION / NOT STARTED
- **Stack decisions:** none
- **Runtime implementation:** not started and not authorized

## 2. M0 selection

After A0 closeout, the accepted roadmap was reviewed against the main plausible first-milestone alternatives:

- M0 — Sovereign Core Walking Skeleton;
- M1 — Governed Conversation, Project Context and Memory;
- M2 — Capability Registry, AHDK Kernel and Reference Provider.

The analysis concluded that M0 should remain first because it retires the foundational risk that Aurora is merely a running session whose identity, operational state and authority disappear when the process ends. M1 depends on a sovereign state/authority foundation. M2 must not turn the provider/SDK ecosystem into the de facto Aurora Core.

The operator explicitly approved M0 as the first Product Milestone on 2026-08-06.

Evidence: `docs/acceptance/2026-08-06-m0-operator-selection.md`.

## 3. Selected M0 outcome

As defined by the accepted Capability Roadmap:

> A minimal Aurora Core persists identity, project state, authority snapshot and one interaction lifecycle across restart without depending on an external model or Harness as authority.

Directional Golden Proof:

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

R0–R6 may refine the executable contract and verification details but must not silently change the accepted product intent.

## 4. Current authorization boundary

```text
A0 baseline:                    ACCEPTED / MERGED
ADR-0001 / ADR-0002:           ACCEPTED
First Product Milestone:        M0 SELECTED
ACRM R0 for M0:                 AUTHORIZED / NOT STARTED
ACRM R1+:                       NOT AUTHORIZED BY IMPLICATION
Architecture Spike planning:    NOT STARTED
Architecture Spike execution:   PROHIBITED
Capability implementation:      PROHIBITED
Aurora Core implementation:     PROHIBITED
AHDK implementation:            PROHIBITED
MNFS integration:               PROHIBITED
Stack selection:                NOT PERFORMED
Mission Contract:               NOT STARTED
Microdesign/Implementation Plan: NOT STARTED
Automatic implementation:       NOT AUTHORIZED
```

Milestone selection does not skip any ACRM gate and does not authorize technical choices or runtime work.

## 5. Deliberately open M0-relevant technical decisions

The following remain open and must not be decided during R0 merely to make progress:

- Aurora Core implementation language;
- initial process/deployment topology;
- operational state storage mechanism;
- state-versus-event model;
- schema representation;
- migration strategy details;
- backup/restore mechanism details;
- authority snapshot representation;
- audit/event mechanism;
- any durable execution engine beyond what M0 actually requires.

Broader open decisions from A0 also remain open unless later M0 applicability makes them material.

## 6. R0 mandate

The fresh R0 session must answer only:

> Is the accepted constitutional intent required for M0 coherent, discoverable, sufficiently owned and authorized to proceed to applicability analysis?

Required inputs include at minimum:

- `AGENTS.md`;
- this `STATUS.md`;
- `docs/DOCUMENTATION-MAP.md`;
- accepted roadmap M0 definition;
- `docs/product/CAPABILITY-REALIZATION-METHOD.md`;
- relevant accepted Blueprint sections;
- accepted ADRs and requirements traceability where needed.

R0 output:

```text
R0 PASS | FAIL | BLOCKED
```

with repository-path citations, identified contradictions/gaps, and the exact next authorized action.

## 7. Current blocker/gate

There is no remaining A0 blocker and no milestone-selection blocker.

The intentional gate is now:

```text
M0_ACRM_R0_CONSTITUTIONAL_BASELINE
```

No implementation blocker exists because implementation is not authorized work yet.

## 8. Immediate next action

```text
start a fresh repository-only session
→ read AGENTS.md and current STATUS
→ execute ACRM R0 for M0 only
→ produce R0 verdict and findings
→ stop before R1 unless separately authorized
```

The next action is constitutional readiness analysis, not implementation, stack selection, Architecture Spike execution or Mission planning.
