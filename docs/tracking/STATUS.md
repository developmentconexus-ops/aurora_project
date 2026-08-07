---
id: DOC-AURORA-STATUS
title: Aurora Project Status
document_type: project_status
form: reference
authority: tracking
status: current
version: 0.16.0
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
  - REVIEW-AURORA-M0-R0-CONSTITUTIONAL-BASELINE-2026-08-06
  - DOC-AURORA-M0-R0-REMEDIATION-AUTHORIZATION
  - DOC-AURORA-M0-R0-REMEDIATION-OPERATOR-ACCEPTANCE
  - DOC-AURORA-M0-R0-REMEDIATION-MERGE-CLOSEOUT
  - DOC-AURORA-M0-R0-RERUN-TARGET-FINDING
  - REVIEW-AURORA-M0-R0-CONSTITUTIONAL-BASELINE-RERUN-2026-08-07
  - DOC-AURORA-M0-R1-OPERATOR-AUTHORIZATION
  - DOC-AURORA-CAP-SOVEREIGN-CORE-APPLICABILITY
  - REVIEW-AURORA-M0-R1-APPLICABILITY-2026-08-07
  - DOC-AURORA-M0-R2-OPERATOR-AUTHORIZATION
  - DOC-AURORA-CAP-SOVEREIGN-CORE-REQUIREMENTS
  - DOC-AURORA-CAP-SOVEREIGN-CORE-R2-COVERAGE
  - REVIEW-AURORA-M0-R2-REQUIREMENTS-2026-08-07
  - DOC-AURORA-M0-R3-OPERATOR-AUTHORIZATION
  - DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC
  - DOC-AURORA-CAP-SOVEREIGN-CORE-THREAT-MODEL
  - DOC-AURORA-CAP-SOVEREIGN-CORE-TEST-PLAN
  - DOC-AURORA-CAP-SOVEREIGN-CORE-R3-COVERAGE
  - REVIEW-AURORA-M0-R3-RESEARCH-FRESHNESS-2026-08-07
  - REVIEW-AURORA-M0-R3-CAPABILITY-READINESS-2026-08-07
  - DOC-AURORA-M0-R4-OPERATOR-AUTHORIZATION
  - DESIGN-AURORA-M0-R4-DECISION-LANDSCAPE
  - DOC-AURORA-CAP-SOVEREIGN-CORE-R4-DECISION-COVERAGE
  - ADR-AURORA-0003
  - ADR-AURORA-0004
  - ADR-AURORA-0005
  - ADR-AURORA-0006
  - ADR-AURORA-0007
  - ADR-AURORA-0008
  - DESIGN-AURORA-M0-SOVEREIGN-STORE-SPIKE-001
  - DESIGN-AURORA-M0-OWNER-TRUST-SPIKE-002
  - REVIEW-AURORA-M0-R4-ARCHITECTURE-DECISION-READINESS-2026-08-07
last_reviewed: 2026-08-07
---

# Aurora Project Status

## 1. Current summary

- **Project:** Projeto Aurora
- **Canonical branch:** `main`
- **Canonical baseline reviewed by initial M0 R0:** `1da990f368a1bc693c09191c41d30a3db454d11e`
- **A0:** ACCEPTED and MERGED
- **A0 merge commit:** `f22085d97e198d99e89d52221b7b26d59d49bc12`
- **ADR-0001:** ACCEPTED
- **ADR-0002:** ACCEPTED
- **First Product Milestone:** `M0 — Sovereign Core Walking Skeleton` — SELECTED by operator
- **Current readiness gate:** ACRM R4 — Architecture/Decision Readiness — BLOCKED
- **Initial R0 verdict:** FAIL
- **R0 re-run target:** `6054f84d007347c0aa9eef9e71317134b1047d3c`
- **R0 re-run verdict:** PASS
- **Canonical R0 remediation merge:** `d0ddfb794296e599ac96bb73bf3772937d371bf9`
- **R0 remediation:** ACCEPTED AND MERGED — `d0ddfb794296e599ac96bb73bf3772937d371bf9`
- **R1 source baseline:** `735f269025e2cc317424e4931f3a5cd414cd6f2a`
- **R1 applicability artifact:** `7f10734ba6018154f196557de6c5735719046253` — 294/294 classified
- **R1 review:** `fbbae69d529a53532e5858693394747081e11d0f` — PASS
- **R1 active constitutional sources for R2:** 127 (`78 APPLIES + 49 PARTIALLY_APPLIES`)
- **R2 source baseline:** `495b712142d7c3d722da2298f7a0b060707f9f5e`
- **R2 reviewed requirements package:** `a8ffbbe22995b8e683d9d49ad06f487c745709f9`
- **R2 derived requirements:** 122 proposed atomic requirements; coverage 127/127 active sources
- **R2 canonical integration:** `9bfab2b30eaccb92ddb55852f97735653172f064`
- **R2 verdict:** PASS
- **R3 source baseline:** `9ea8adf5c115f54071d7e36e312695d19420d8b0`
- **R3 reviewed clean package:** `4b8558b724f28310fd8fbc6884944f7f59f16ea6`
- **R3 canonical integration:** `58a7946b62f27d8b8784169e7e3741eec24ecc95`
- **R3 requirement allocation:** 122/122 R2 requirements allocated to Spec mechanisms and planned verification
- **R3 test-plan baseline:** 84 planned test IDs; 80 referenced directly by requirement coverage
- **R3 research freshness:** SUFFICIENT for boundary reasoning; R4 mechanism/version revalidation required
- **R3 verdict:** PASS
- **R4 source baseline:** `d00cc1abfc2a41ac7e81e1f3478e188b3c5e9e52`
- **R4 documentary package/review integration:** `71f64bab2a82c2a7781d28274224f60abc277b2c`
- **R4 architecture questions:** 15/15 mapped to research, proposed ADRs and/or exact spike evidence
- **R4 focused research:** 4 current reports + 4 source manifests
- **R4 ADRs:** ADR-0003..0008 all PROPOSED; 0003..0006 documentary-ready for operator review; 0007..0008 spike-blocked
- **SPK-AURORA-M0-SOVEREIGN-STORE-001:** PROPOSED / EXECUTION NOT AUTHORIZED
- **SPK-AURORA-M0-OWNER-TRUST-002:** PROPOSED / EXECUTION NOT AUTHORIZED / depends on reviewed SPK-001 result
- **R4 verdict:** BLOCKED — required executable spike evidence and operator decision acceptance are not yet complete
- **R5 — Contract Readiness:** NOT AUTHORIZED
- **R6 and later gates:** NOT AUTHORIZED BY IMPLICATION
- **Accepted stack decisions:** none; R4 technical choices remain proposed pending required review/evidence
- **Runtime implementation:** not started and not authorized

## 2. M0 selection

The operator explicitly selected `M0 — Sovereign Core Walking Skeleton` as the first Product Milestone on 2026-08-06 after comparing M0 against M1 and M2.

Accepted M0 outcome:

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

Evidence: `docs/acceptance/2026-08-06-m0-operator-selection.md`.

## 3. Initial M0 R0 result

A fresh repository-only R0 review was executed against:

```text
1da990f368a1bc693c09191c41d30a3db454d11e
```

Verdict:

```text
R0 FAIL
```

Gate-failing findings:

- **R0-F01 — Product Milestone anatomy divergence:** Blueprint 14 required a complete executable-horizon milestone anatomy, while selected M0 lacked Architecture Spikes, Exit Criteria, Telemetry Baseline, Dependencies and Promotion/Authority Boundary; §14.5 also failed to distinguish selected executable milestones from intentionally directional future milestones.
- **R0-F02 — ADR status divergence:** `docs/adr/README.md`, despite owning ADR status discovery, still reported ADR-0001/0002 as proposed while their accepted ADR files and operator evidence reported ACCEPTED.
- **R0-F03 — mutable-state duplication/drift:** bootstrap/index/constitutional documents retained pre-A0/pre-M0 coordination snapshots even though `STATUS.md` and operator evidence had advanced.

Review record: `docs/reviews/2026-08-06-m0-r0-constitutional-baseline-review.md`.

## 4. Remediation boundary

The operator authorized remediation after the R0 FAIL. The authorized work is documentary/constitutional only:

- repair M0 roadmap anatomy without choosing technical mechanisms;
- make §14.5 consistent with the constitutional/executable two-horizon model;
- align ADR status discovery with accepted ADR owners/evidence;
- remove mutable-current-state ownership from durable constitutional/index documents and point it to `STATUS.md`;
- regenerate generated projections;
- run documentation validation;
- present the corrected revision for explicit operator acceptance;
- re-run M0 R0 only after the corrected constitutional revision is accepted/canonical.

This remediation does **not** authorize R1, Architecture Spike execution, Capability implementation, Aurora Core implementation, AHDK, MNFS integration, stack selection, Mission Contract or Microdesign.

## 5. Deliberately open M0-relevant technical decisions

R4 has now researched and proposed concrete dispositions for the M0 mechanism questions below, but none of the new ADRs is accepted yet. Store/atomicity and owner-root/time/restore mechanisms additionally require reviewed Architecture Spike evidence before acceptance:

- Aurora Core implementation language/runtime;
- operational-state storage mechanism;
- state-versus-event persistence pattern;
- schema/serialization representation;
- crash-consistent commit/atomicity mechanism;
- integrity mechanism;
- time/rollback semantics used for authority expiry;
- local owner authentication/bootstrap mechanism;
- export/backup format and topology;
- migration mechanism/tooling;
- audit/event physical mechanism;
- telemetry backend/transport;
- initial process/deployment topology;
- durable execution engine applicability, only if M0 need proves it proportionate;
- authority freshness/revalidation mechanism after restore.

## 6. Current authorization boundary

```text
A0 baseline:                    ACCEPTED / MERGED
ADR-0001 / ADR-0002:           ACCEPTED
First Product Milestone:        M0 SELECTED
Initial M0 R0:                 FAIL — historical
R0 documentary remediation:    ACCEPTED / MERGED
M0 ACRM R0 re-run:             PASS — target 6054f84d007347c0aa9eef9e71317134b1047d3c
M0 ACRM R1 — Applicability:    PASS — source baseline 735f269025e2cc317424e4931f3a5cd414cd6f2a
R1 applicability coverage:     294/294 classified; 127 active sources
M0 ACRM R2 — Requirements:     PASS — source baseline 495b712142d7c3d722da2298f7a0b060707f9f5e
R2 requirements baseline:       122 proposed atomic requirements; coverage 127/127
M0 ACRM R3 — Capability Readiness: PASS — source baseline 9ea8adf5c115f54071d7e36e312695d19420d8b0
R3 proposed Capability package: Spec + threat model + test plan + 122/122 allocation
M0 ACRM R4 — Architecture/Decision Readiness: BLOCKED — documentary baseline d00cc1abfc2a41ac7e81e1f3478e188b3c5e9e52
R4 decision coverage:             15/15 mapped; 6 proposed ADRs; 2 required spikes unexecuted
ACRM R5 — Contract Readiness:    NOT AUTHORIZED
ACRM R6+:                       NOT AUTHORIZED
Architecture Spike execution:   PROHIBITED / AWAITING EXPLICIT OPERATOR AUTHORIZATION
Capability implementation:      PROHIBITED
Aurora Core implementation:     PROHIBITED
AHDK implementation:            PROHIBITED
MNFS integration:               PROHIBITED
Stack selection:                NOT PERFORMED
Mission Contract:               NOT STARTED
Microdesign/Implementation Plan: NOT STARTED
```

## 7. Current blocker/gate

M0 ACRM R4 has completed the documentary architecture/decision investigation against fixed source baseline `d00cc1abfc2a41ac7e81e1f3478e188b3c5e9e52`. All 15 R3-open architecture questions now have explicit dispositions. Four current focused research reports support six proposed ADRs: ADR-0003 through ADR-0006 are documentary-ready for operator review, while ADR-0007 (SQLite operational store) and ADR-0008 (Owner Root/recovery trust) are explicitly blocked by executable evidence. Two minimal sequential Architecture Spikes have been specified: `SPK-AURORA-M0-SOVEREIGN-STORE-001` first, then `SPK-AURORA-M0-OWNER-TRUST-002` only after the first spike has a reviewed viable store result.

The independent R4 review returned `R4 BLOCKED`, not FAIL and not PASS. The blocker is intentional and concrete: the accepted Research Map requires crash/restart/restore evidence before Sovereign Core storage/recovery commitment; Architecture Spike execution remains prohibited because no separate operator authorization exists; and all R4 ADRs remain `proposed`. No new technical choice is governing yet. R5, Mission Contract, Microdesign and implementation remain unauthorized.

## 8. Immediate next action

```text
R4 documentary investigation reviewed
→ R4 BLOCKED recorded
→ operator reviews ADR-0003..0006
→ separately authorize exact SPK-AURORA-M0-SOVEREIGN-STORE-001 execution if approved
→ SPK-002 remains blocked until SPK-001 evidence is reviewed
→ do not begin R5
```
