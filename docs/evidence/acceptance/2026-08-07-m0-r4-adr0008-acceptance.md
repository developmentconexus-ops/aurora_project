---
id: DOC-AURORA-M0-R4-ADR0008-ACCEPTANCE
title: M0 R4 ADR-0008 Operator Acceptance
document_type: operator_acceptance
form: evidence
authority: operator
status: accepted
accepted_at: 2026-08-07
version: 1.0.0
owners:
  - operator
source_of_truth_for:
  - operator acceptance of ADR-0008 revision 0.2.0
related:
  - ADR-AURORA-0008
  - DOC-AURORA-M0-R4-SPK002-EVIDENCE-RECEIPT
  - REVIEW-AURORA-M0-R4-SPK002-OWNER-TRUST-2026-08-07
  - DOC-AURORA-M0-R4-OPERATOR-AUTHORIZATION
  - DOC-AURORA-STATUS
canonical_pre_acceptance_revision: 35ce98fb2ddde16133c01a4da2f3545c8ae5e308
accepted_adr_blob: 2a1497f8311ba9d04cd61f5025d7eae2af2fc57f
accepted_adr_version: 0.2.0
last_reviewed: 2026-08-07
---

# M0 R4 — ADR-0008 Operator Acceptance

## 1. Operator decision

On 2026-08-07 the operator explicitly stated:

> Aceito ADR-0008 v0.2.0.

This accepts exactly:

```text
ADR-AURORA-0008
revision: 0.2.0
blob: 2a1497f8311ba9d04cd61f5025d7eae2af2fc57f
canonical pre-acceptance main: 35ce98fb2ddde16133c01a4da2f3545c8ae5e308
```

No later or materially modified revision is accepted by implication.

## 2. Accepted decision boundary

The accepted ADR establishes the M0 owner-root and recovery trust architecture informed by reviewed SPK-002 evidence, including:

- random 256-bit Owner Root Key as stable local owner/domain integrity root;
- Argon2id-derived KEK for owner-passphrase unlock;
- AES-256-GCM wrapped ORK at rest;
- HKDF-SHA-256 purpose-separated subkeys;
- HMAC-SHA-256 governing-state and external trust-anchor integrity;
- authenticated generation and observed wall-time high-water outside the operational database;
- fail-closed `STATE_ROLLBACK`, `ANCHOR_LAG`, `TIME_UNTRUSTED` and `REVALIDATION_REQUIRED` semantics;
- authenticated-owner-only reconciliation/revalidation paths;
- historical restore that cannot silently revive current authority.

The residual local-threat boundary and implementation/evidence obligations documented in ADR-0008 and the SPK-002 review remain part of the accepted decision.

## 3. What this acceptance authorizes

This acceptance authorizes ADR-0008 to move from `proposed / evidence-ready` to `accepted` and allows the already-authorized M0 R4 readiness review to be rerun against the now-complete architecture decision set.

## 4. What this acceptance does not authorize

This acceptance does **not** authorize by implication:

```text
ACRM R5
Mission Contract creation
ACRM R6
Microdesign / Implementation Plan
production implementation
promotion of SPK-002 code
Aurora Core runtime code
Mastra implementation
AHDK implementation
MNFS integration
additional Architecture Spikes
```

Any next ACRM gate still requires its own operator authorization under the Capability Realization Method.

## 5. Required next action

```text
record ADR-0008 v0.2.0 as accepted
→ rerun M0 ACRM R4 against the complete accepted/evidenced decision set
→ PASS | FAIL | BLOCKED
→ STOP
→ if R4 PASS, await separate operator authorization before R5
```
