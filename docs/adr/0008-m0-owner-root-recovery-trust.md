---
id: ADR-AURORA-0008
title: M0 Owner Root and Recovery Trust Boundary
document_type: adr
form: explanation
authority: decision
status: proposed
version: 0.1.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - proposed M0 local owner authentication, integrity-anchor, time rollback and restore revalidation mechanism
related:
  - DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC
  - DOC-AURORA-CAP-SOVEREIGN-CORE-THREAT-MODEL
  - RESEARCH-AURORA-M0-OWNER-AUTHORITY-RECOVERY-R4-V1
  - RESEARCH-AURORA-M0-PORTABILITY-INTEGRITY-R4-V1
  - DESIGN-AURORA-M0-OWNER-TRUST-SPIKE-002
supersedes: []
superseded_by: null
last_reviewed: 2026-08-07
---

# ADR-0008 — M0 Owner Root and Recovery Trust Boundary

## Context

M0 must distinguish owner authority from raw technical access, detect common state rollback/corruption, prevent clock rollback from reviving expired authority and ensure historical backups cannot silently restore current permission.

A key stored only in the operational DB cannot provide an independent trust boundary for that DB.

## Decision drivers

- local-first/no cloud identity requirement;
- owner authority distinct from DB/filesystem access;
- passphrase rotation without redefining owner/domain identity;
- governing-state authenticated integrity;
- DB rollback detection within the accepted local threat boundary;
- fail-closed backward-time handling;
- explicit post-restore authority revalidation;
- recovery to a fresh machine.

Affected requirements include `REQ-032..045`, `058..066`, `067..076`, `091`, `095`, `104`, `119`.

## Evidence

`RESEARCH-AURORA-M0-OWNER-AUTHORITY-RECOVERY-R4-V1` derives the candidate architecture from Argon2id, secure random generation, AES-GCM, HKDF and HMAC standards/current Go implementations.

The cryptographic primitives are mature. The remaining material uncertainty is the crash behavior of the **two durable trust domains**—operational DB and external Owner Trust Store—and recovery classifications under real runtime/filesystem behavior.

## Options

1. OS user identity alone;
2. passphrase directly derives every long-lived integrity key;
3. random Owner Root Key (ORK) wrapped by passphrase-derived KEK + external trust anchor;
4. mandatory platform keychain/hardware root.

## Decision

**Proposed candidate, acceptance blocked by SPK-002:**

```text
Owner Root Key = random 256-bit secret
passphrase KDF = Argon2id, versioned parameters/salt
ORK at rest = AES-256-GCM wrapped by KDF-derived KEK
purpose subkeys = HKDF-SHA-256 from ORK
governing descriptor/trust anchor = HMAC-SHA-256
Owner Trust Store = physically separate from operational DB
```

The Owner Trust Store holds authenticated high-water metadata for state generation and trusted wall time.

Proposed recovery classification:

```text
DB generation == anchor → normal verification
DB generation < anchor  → STATE_ROLLBACK / fail closed
DB generation > anchor  → ANCHOR_LAG / authenticated recovery reconciliation
```

For authority expiry:

```text
process-local elapsed time → monotonic time where available
cross-restart expiry → UTC wall time + authenticated high-water
meaningful backward step → TIME_UNTRUSTED / expiry-dependent authority non-permitting
```

For historical restore:

```text
restore old authentic state
→ REVALIDATION_REQUIRED
→ restored grant cannot authorize itself
→ authenticated owner explicitly creates new authority revision
```

An old trust-anchor high-water MUST NOT be imported as proof of current freshness.

This ADR MUST NOT move to `accepted` until `SPK-AURORA-M0-OWNER-TRUST-002` proves the crash/rollback/time/revalidation protocol against the selected M0 store.

## Consequences

### Positive

- owner authority is not equivalent to raw DB access;
- passphrase rotation rewraps ORK instead of re-keying domain identity;
- database-only rollback becomes detectable;
- old backups cannot silently revive permission;
- future OS keychain/hardware protection can wrap/protect the same ORK without changing Core semantics.

### Negative

- introduces a second small durable trust artifact;
- DB commit and trust-anchor update require explicit crash reconciliation;
- owner must unlock/recover authority-sensitive operations;
- purely local scheme cannot defeat replay of every trust file under total host/root compromise.

### Risks

The largest risk is false confidence in rollback protection. The accepted threat boundary excludes hostile total local-root replay; evidence and docs MUST state this limitation.

## Compatibility / migration / rollback

Logical exports may carry an encrypted ORK recovery envelope but do not import current trust-anchor freshness. On a new machine the owner unlocks/reestablishes local trust and explicitly revalidates authority.

Future TPM/YubiKey/OS keystore integration may protect ORK without changing authority semantics.

## Validation

Required before acceptance:

```text
SPK-AURORA-M0-OWNER-TRUST-002
→ proves unlock/rotation
→ DB rollback detection
→ ANCHOR_LAG crash recovery
→ TIME_UNTRUSTED behavior
→ historical restore REVALIDATION_REQUIRED
→ owner-only post-restore authority revision
```

## Reconsideration triggers

- local threat model expands to hostile administrator/root replay;
- hardware-backed trust becomes required;
- owner unlock burden becomes operationally unacceptable;
- spike cannot make cross-file crash recovery unambiguous;
- target host changes require a stronger root-custody mechanism.

## References

- `RESEARCH-AURORA-M0-OWNER-AUTHORITY-RECOVERY-R4-V1`
- `SPK-AURORA-M0-OWNER-TRUST-002`
- RFC 9106, RFC 5869, NIST GCM/HMAC/SHA standards cited in research manifest
