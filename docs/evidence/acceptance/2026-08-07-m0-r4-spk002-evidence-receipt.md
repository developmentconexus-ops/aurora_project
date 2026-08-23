---
id: DOC-AURORA-M0-R4-SPK002-EVIDENCE-RECEIPT
title: M0 R4 SPK-002 Owner Trust Evidence Receipt
document_type: evidence_receipt
form: reference
authority: evidence
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - fixed SPK-002 execution revision and workflow evidence
  - cross-platform owner-root/trust protocol measurements
  - SPK-002 artifact digests and execution limitations
related:
  - DOC-AURORA-M0-R4-ADR0007-ACCEPTANCE-SPK002-AUTHORIZATION
  - DESIGN-AURORA-M0-OWNER-TRUST-SPIKE-002
  - ADR-AURORA-0008
  - REVIEW-AURORA-M0-R4-SPK002-OWNER-TRUST-2026-08-07
source_revision: ba7a211b40e21bfcd9aa1b90e5ca3c21cd229d1b
spike_execution_revision: c76b96fee36878f15c54028b4ba1896f84ebdeca
workflow_run: 31219882882
last_reviewed: 2026-08-07
---

# M0 R4 — SPK-002 Owner Trust Evidence Receipt

## 1. Fixed execution identity

```text
spike: SPK-AURORA-M0-OWNER-TRUST-002
canonical source baseline: ba7a211b40e21bfcd9aa1b90e5ca3c21cd229d1b
canonical specification blob: 0ffb6fa2b35014e34b4301365dc2d5a8d96f021d
specification version: 0.1.0
final disposable execution revision: c76b96fee36878f15c54028b4ba1896f84ebdeca
final workflow run: 31219882882
workflow conclusion: SUCCESS
```

The final execution revision contains disposable spike code only on `spike/m0-owner-trust-002`. It is not production code and MUST NOT be promoted to `main` as Aurora Core implementation.

## 2. Evidence-qualified dependency/runtime baseline

```text
Go: 1.26.5
CGO_ENABLED: 0
golang.org/x/crypto: v0.54.0
modernc.org/sqlite: v1.54.0
modernc.org/libc: v1.74.1 exact compatibility pin
store class: SQLite via database/sql
journal_mode: WAL
synchronous: FULL
foreign_keys: ON
```

The `modernc.org/libc` compatibility pin is deliberately explicit because the accepted ADR-0007 requires exact reproducible binding/module compatibility rather than relying on an accidental transitive resolution.

## 3. Candidate A implemented

```text
Owner Root Key (ORK): random 256-bit secret
owner passphrase → Argon2id KEK
ORK at rest → AES-256-GCM wrapped envelope
ORK purpose keys → HKDF-SHA-256
DB governing descriptor → HMAC-SHA-256
external Owner Trust Store → authenticated generation + observed wall-time high-water
historical/fresh-machine restore → REVALIDATION_REQUIRED
```

Evidence-qualified Argon2id profile:

```text
memory:       64 MiB
iterations:   3
parallelism:  4
KEK:          32 bytes
```

## 4. Final scenario result

The final Linux/Windows matrix passed the canonical S01–S12 scenario contract:

| Scenario | Required behavior | Final result |
|---|---|---|
| S01 | random ORK; correct unlock; wrong passphrase fails; no plaintext passphrase/ORK on durable artifacts | PASS |
| S02 | passphrase rotation preserves ORK identity; old passphrase fails | PASS |
| S03 | raw DB mutation without ORK-derived MAC cannot become governing | PASS |
| S04 | DB-only rollback below trust high-water → `STATE_ROLLBACK` | PASS |
| S05 | process kill across DB→anchor durable gap → deterministic old/`ANCHOR_LAG`/new classification + owner reconciliation | PASS |
| S06 | anchor ahead / DB behind fails closed; no automatic anchor rollback | PASS |
| S07 | backward wall clock → `TIME_UNTRUSTED`; expired permission does not revive | PASS |
| S08 | authentic historical ACTIVE restore → `REVALIDATION_REQUIRED` | PASS |
| S09 | restored historical grant cannot self-revalidate | PASS |
| S10 | authenticated owner creates new post-restore authority revision before next safe action | PASS |
| S11 | fresh-machine encrypted root recovery remains `REVALIDATION_REQUIRED` without current trust high-water | PASS |
| S12 | missing wrapped root/recovery material fails explicitly; no fabricated authority/owner | PASS |

## 5. Recovery classification matrix

The final test suite also passed every required DB/anchor state classification:

```text
valid DB generation N / valid anchor N     → NORMAL
valid DB generation N+1 / valid anchor N   → ANCHOR_LAG / non-permitting
valid DB generation N / valid anchor N+1   → STATE_ROLLBACK / non-permitting
invalid governing-state HMAC               → INVALID_DB_MAC / non-permitting
invalid trust-anchor HMAC                   → INVALID_ANCHOR_MAC / non-permitting
missing current trust anchor                → MISSING_ANCHOR / non-permitting
missing wrapped root                        → MISSING_WRAPPED_ROOT / non-permitting
backward time below authenticated high-water→ TIME_UNTRUSTED / non-permitting
historical/fresh restore                    → REVALIDATION_REQUIRED / non-permitting
```

No anomaly is mapped to permissive best-effort behavior.

## 6. Mutation-boundary hardening finding

An adversarial review after the first complete green matrix found that ordinary `Advance` mutations could theoretically be called directly from an already-unlocked owner session while the durable pair was classified as `STATE_ROLLBACK`, `ANCHOR_LAG`, `TIME_UNTRUSTED` or restore-pending.

The spike was hardened before final evidence:

```text
ordinary governing mutation
→ preflight wrapped-root/session + DB HMAC + anchor HMAC
→ require no restore marker
→ require DB generation == anchor generation
→ require wall clock >= authenticated observed high-water
→ only then enter DB mutation transaction
```

Explicit recovery operations remain separate:

```text
ANCHOR_LAG → ReconcileAnchor
REVALIDATION_REQUIRED → authenticated-owner Revalidate
STATE_ROLLBACK → fail closed; no implicit repair
TIME_UNTRUSTED → fail closed for ordinary mutation
```

The focused RED initially failed on absent anomaly errors. The hardening test then passed, and the complete hardened Linux/Windows final matrix passed again at run `31219882882`.

## 7. Final performance/usability measurements

These are evidence observations, not product SLOs.

| Measurement | Linux | Windows |
|---|---:|---:|
| Argon2id KEK derivation | 52.1 ms | 53.4 ms |
| observed Go heap delta during Argon2id | ~64.0 MiB | ~64.0 MiB |
| bootstrap | 62.1 ms | 130.7 ms |
| passphrase rotation | 105.8 ms | 106.1 ms |
| governing DB commit | 2.08 ms | 173.80 ms |
| trust-anchor publish | 0.35 ms | 4.37 ms |
| restore posture | `REVALIDATION_REQUIRED` | `REVALIDATION_REQUIRED` |
| post-owner-revalidation | `NORMAL / permitting` | `NORMAL / permitting` |

The KDF burden is compatible with an M0 local owner-unlock/recovery path. Performance differences in hosted runners are secondary to correctness and do not define an M0 throughput target.

## 8. Sanitized evidence and artifact digests

Final artifact set:

```text
spk002-ubuntu-24.04
  artifact id: 9009957252
  sha256: fc58cb6c24bf9f269ba3dd5ecc6983d53a9853792f9dfcbf8540ed8e3816f8b7

spk002-windows-latest
  artifact id: 9010002102
  sha256: 25f1bac4f8f93a4ff88ac655561ea795396ed545b967ec43d1da8e3985535453

spk002-aggregate
  artifact id: 9010007284
  sha256: f4efd57740ecc22c8b34ce891bcd1128812de427d5b50263f9a0ca3639745a8e
```

Both platform jobs passed the final evidence-secret-hygiene step. The evidence set contains deterministic non-secret fixtures, sanitized logs/module locks/measurements and no ORK/KEK/derived keys.

## 9. Evidence-hygiene debugging finding

An earlier otherwise-green Ubuntu run put the evidence executable itself inside the artifact directory. Because compiled fixture strings are naturally embedded in a binary, the evidence scanner correctly failed that packaging layout.

Disposition:

```text
EVIDENCE_PACKAGING_FINDING
→ protocol unaffected
→ measurement executable moved to RUNNER_TEMP
→ executable deleted before artifact upload
→ final Linux + Windows secret-hygiene checks PASS
```

The earlier artifact is not used as final SPK-002 evidence.

## 10. Candidate B/C disposition

Candidate B—deriving the long-lived integrity root directly from the passphrase—was not expanded into a duplicate full prototype. The canonical spike explicitly permits documentary analysis when Candidate A evidence is conclusive.

Candidate A proved the decision-driving property:

```text
passphrase changes
BUT
ORK/domain-root identity does not change
```

A direct passphrase-root design would either change the integrity root when the passphrase changes or require an additional rekey/continuity procedure across all protected material. That is not lower recovery/rotation complexity for M0.

Candidate C—mandatory OS keychain/hardware root—was not escalated because Candidate A satisfied the intended local threat boundary. It remains a future strengthening class if the threat model later expands to hostile total-local-admin replay.

## 11. Required residual-risk statement

The final evidence does **not** claim protection against an attacker who can capture/replay every local trust artifact **and** compromise owner secrets/runtime state.

The proved M0 boundary is narrower:

```text
raw operational DB access alone
or DB-only rollback
or accidental/local partial rollback
or process crash between DB and trust anchor
or historical backup restore
or backward wall-clock movement

→ cannot silently mint/revive current authority
```

This is consistent with the canonical SPK-002 residual-risk boundary.

## 12. Additional hardening reserved for implementation design

The production decoder MUST validate/allowlist the accepted Argon2id profile and bounded KDF parameters before invoking the KDF. The disposable spike reads versioned envelope parameters directly; a malicious extreme parameter value could otherwise create an availability/DoS condition. This does not create authority escalation and does not invalidate the R4 architecture, but it is a required R6 hardening constraint.

Likewise, exact directory-fsync/rename behavior for literal power-loss durability remains an environment-specific implementation-design validation. The spike proves real process-kill behavior on GitHub-hosted Linux/Windows filesystems, not every storage controller/power-failure mode.

## 13. Lifecycle implication

The executable evidence supports:

```text
SPK-AURORA-M0-OWNER-TRUST-002
→ EVIDENCE_COMPLETE
```

Closure requires the independent evidence review. If that review remains PASS, the spike may become:

```text
REVIEWED
→ DECISION_INFORMED
→ CLOSED
```

This receipt does not accept ADR-0008 and does not authorize R5 or production implementation.
