---
id: SPK-AURORA-M0-OWNER-TRUST-002
title: M0 Owner Root, Rollback, Time and Restore Freshness Spike
document_type: architecture_spike
form: reference
authority: design
status: proposed
version: 0.1.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - proposed executable investigation of M0 owner root and local trust-anchor protocol
related:
  - DOC-AURORA-M0-R4-OPERATOR-AUTHORIZATION
  - DESIGN-AURORA-M0-R4-DECISION-LANDSCAPE
  - RESEARCH-AURORA-M0-OWNER-AUTHORITY-RECOVERY-R4-V1
  - RESEARCH-AURORA-M0-PORTABILITY-INTEGRITY-R4-V1
  - SPK-AURORA-M0-SOVEREIGN-STORE-001
  - DOC-AURORA-CAP-SOVEREIGN-CORE-THREAT-MODEL
source_revision: d00cc1abfc2a41ac7e81e1f3478e188b3c5e9e52
last_reviewed: 2026-08-07
---

# SPK-AURORA-M0-OWNER-TRUST-002

## 1. Authorization state

```text
Specification: PROPOSED
Execution: NOT AUTHORIZED
Dependency: SPK-AURORA-M0-SOVEREIGN-STORE-001 reviewed result
Production promotion: PROHIBITED
```

This spike MUST NOT execute until separately authorized. It should run against the viable store/binding selected for decision evaluation by SPK-001; it must not independently build a second Core.

## 2. Exact uncertainty

Can the proposed local Owner Root + external trust-anchor protocol distinguish owner authority from raw technical access and remain fail-closed/recoverable across passphrase errors, DB rollback, process crash, backward wall-clock movement and historical backup restore without requiring cloud/hardware trust?

## 3. Decisions informed

- `R4-Q-AUTHN-001` owner authentication/bootstrap;
- `R4-Q-INTEGRITY-001` authenticated governing-state integrity;
- `R4-Q-TIME-001` time rollback semantics;
- `R4-Q-RESTORE-001` authority freshness/revalidation;
- security portion of `R4-Q-EXPORT-001`.

## 4. Candidate protocol

```text
Owner Root Key (ORK): random 256 bits
Owner passphrase → Argon2id KEK
ORK wrapped at rest with AES-256-GCM
ORK → HKDF-SHA-256 purpose subkeys
DB protected descriptor → HMAC-SHA-256
Owner Trust Store outside operational DB → generation/time high-water + HMAC
historical restore → REVALIDATION_REQUIRED
```

Baseline Argon2id spike profile:

```text
memory = 64 MiB
time/iterations = 3
parallelism = 4 when target CPU permits
KEK = 32 bytes
```

Any adjustment must be recorded as evidence and versioned.

## 5. Alternatives considered

### A — random ORK wrapped by passphrase-derived KEK

Primary candidate; supports passphrase rotation without changing ORK lineage.

### B — passphrase directly derives long-lived integrity root

Reference alternative. Implement only the minimal path needed to confirm whether key rotation/recovery complexity is materially lower/higher; do not build duplicate full flows if documentary analysis plus Candidate A evidence is conclusive.

### C — mandatory OS keychain/hardware root

Fallback class, not initially implemented. Escalate only if Candidate A cannot satisfy the intended local threat boundary without unacceptable fragility.

## 6. Controlled environment

Use the store/binding winner from SPK-001 on at least the same Linux/Windows CI matrix if that winner supports both.

Capture:

- exact Go version;
- x/crypto version;
- crypto primitives/APIs;
- selected store/binding/config;
- filesystem;
- injected wall-clock fixture mechanism;
- all KDF parameters.

Never log passphrases, ORK, KEK or derived subkeys.

Use deterministic non-secret fixture values only.

## 7. Minimal prototype

Implement only:

- owner bootstrap/unlock;
- wrapped ORK envelope;
- purpose-subkey derivation;
- small Owner Trust Store;
- protected current-state descriptor in selected DB;
- state generation increment;
- authority expiry timestamp fixture;
- rollback/time/restore state evaluator;
- explicit owner revalidation command/path.

No external effects, OAuth, workload identity, cloud KMS, hardware token, UI framework or full policy engine.

## 8. Test scenarios

### S01 — bootstrap/unlock

- create ORK using secure RNG;
- wrap under Argon2id-derived KEK;
- correct passphrase unlocks;
- wrong passphrase fails without leaking ORK;
- on-disk scan contains no plaintext ORK/passphrase.

### S02 — passphrase rotation

- unlock ORK;
- rewrap under new salt/passphrase-derived KEK;
- ORK identity/digest remains same;
- old passphrase stops unlocking new envelope;
- governing DB integrity still verifies with ORK-derived keys.

### S03 — raw technical DB mutation

Modify protected governing descriptor without ORK-derived MAC. Expected: startup/unlock rejects governing use.

### S04 — DB-only rollback

Create generations N then N+1; preserve trust anchor at N+1; replace DB with valid historical N copy. Expected:

```text
DB generation < anchor
→ STATE_ROLLBACK / fail closed
```

### S05 — crash after DB commit, before anchor update

Inject process death after durable DB generation N+1 commit but before trust anchor update from N to N+1. Expected:

```text
DB generation > anchor
+ valid state HMAC
→ ANCHOR_LAG
→ no silent rollback
→ authenticated owner reconciliation can advance anchor
```

### S06 — anchor ahead / DB behind

Force anchor N+1 while DB is N. Expected: fail closed; no automatic anchor rollback.

### S07 — backward wall clock

Set trusted high-water, then simulate current wall time meaningfully earlier. Expected:

```text
TIME_UNTRUSTED
→ expiry-dependent authority is non-permitting
```

In-process monotonic timing behavior should remain unaffected for elapsed-duration checks.

### S08 — historically authentic old backup restore

Backup at ACTIVE grant, later revoke in live state, then restore old backup into fresh environment with valid historical cryptographic metadata. Expected:

```text
historical state authenticity may verify
BUT authority freshness != current
→ REVALIDATION_REQUIRED
```

### S09 — self-revalidation attack

Attempt to use restored historical grant itself to authorize revalidation. Expected: denied.

### S10 — owner revalidation

After correct owner unlock, explicitly revalidate intended current authority. Expected:

- new authority revision;
- new local trust-anchor generation;
- next safe action available only after the new revision is current.

### S11 — fresh-machine root recovery

Restore logical export containing encrypted root recovery envelope but not a current trust-anchor high-water. Correct owner secret unlocks root; state stays `REVALIDATION_REQUIRED` until explicit revalidation.

### S12 — missing root/recovery material

Expected: explicit owner-root recovery failure; no fabricated authority or replacement owner identity.

## 9. Crash protocol proof

The spike MUST record the exact DB/anchor ordering and inject failure between every durable boundary.

Acceptance requires a deterministic classification for every observed state pair:

```text
DB generation == anchor
DB generation > anchor
DB generation < anchor
invalid DB MAC
invalid anchor MAC
missing anchor
missing wrapped root key
```

No pair may be interpreted as success through best-effort guessing.

## 10. Performance/usability measurements

Measure:

- Argon2id unlock latency and peak memory on target runners;
- bootstrap/rotation time;
- trust-anchor update latency;
- recovery/reconciliation steps;
- user-visible number of required secret prompts in Golden Proof paths;
- size of trust artifact;
- failure diagnostics quality.

Security parameters must not be silently weakened just to improve latency.

## 11. Golden Proof

```text
bootstrap owner root
→ unlock
→ write generation N
→ write generation N+1
→ detect DB rollback to N
→ recover crash-induced ANCHOR_LAG safely
→ mark backward time as untrusted
→ restore authentic old ACTIVE backup after live revocation
→ restored grant cannot authorize itself
→ owner explicitly revalidates
→ new authority revision becomes current
→ passphrase rotates without changing ORK identity
```

## 12. Gate criteria

Candidate protocol passes only if:

- technical DB access alone cannot mint a valid governing descriptor;
- DB-only rollback is detected;
- crash between DB and trust-anchor update has one safe/recoverable classification;
- time rollback cannot revive expired permission;
- historical restore never automatically restores current permission;
- only authenticated owner recovery can create post-restore authority revision;
- root key is encrypted at rest;
- passphrase rotation does not require changing domain/owner identity;
- target resource/latency burden is acceptable and documented;
- no secret appears in logs/evidence artifacts.

## 13. Known residual risk to document

Purely local M0 cannot prove rollback resistance against an attacker with authority to capture and replay **all** trust files plus secrets/runtime state. This spike validates the narrower accepted threat boundary: technical DB/state compromise or accidental rollback without total owner/admin trust-root compromise.

If evidence shows this boundary is insufficient for R3 requirements, the spike must FAIL/BLOCK and recommend a stronger hardware/remote trust anchor rather than overclaim protection.

## 14. Evidence artifacts

Required:

- exact source/dependency revisions;
- machine-readable scenario matrix;
- sanitized logs;
- KDF timing/memory measurements;
- trust-store/DB state diagrams before/after faults (without secret values);
- recovery classifications;
- proof that old grants remain blocked after restore;
- limitations/residual-risk report;
- reviewer verdict.

## 15. Stop conditions

Stop if:

- the candidate requires secret logging or embedding ORK in canonical DB;
- safe crash reconciliation cannot distinguish rollback from anchor lag;
- meaningful backward-time test can reactivate expired authority;
- old backup can become current authority without explicit owner action;
- prototype scope expands into future OAuth/PDP/effect infrastructure.

## 16. Disposal rule

```text
DISCARD
```

The spike is evidence only. Production security code requires later Contract/Microdesign/implementation gates.

## 17. Decision implication

A passing result may support acceptance of the M0 owner-root/trust architecture ADR. It does not authorize implementation or claim protection against total local-root compromise.
