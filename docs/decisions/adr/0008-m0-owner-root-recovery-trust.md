---
id: ADR-AURORA-0008
title: M0 Owner Root and Recovery Trust Boundary
document_type: adr
form: explanation
authority: decision
status: accepted
accepted_at: 2026-08-07
acceptance_evidence: DOC-AURORA-M0-R4-ADR0008-ACCEPTANCE
version: 0.2.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - M0 local owner authentication, integrity-anchor, time rollback and restore revalidation mechanism
related:
  - DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC
  - DOC-AURORA-CAP-SOVEREIGN-CORE-THREAT-MODEL
  - RESEARCH-AURORA-M0-OWNER-AUTHORITY-RECOVERY-R4-V1
  - RESEARCH-AURORA-M0-PORTABILITY-INTEGRITY-R4-V1
  - DESIGN-AURORA-M0-OWNER-TRUST-SPIKE-002
  - DOC-AURORA-M0-R4-SPK002-EVIDENCE-RECEIPT
  - REVIEW-AURORA-M0-R4-SPK002-OWNER-TRUST-2026-08-07
  - DOC-AURORA-M0-R4-ADR0008-ACCEPTANCE
  - REVIEW-AURORA-M0-R4-ARCHITECTURE-DECISION-READINESS-RERUN-2026-08-07
supersedes: []
superseded_by: null
last_reviewed: 2026-08-07
---

# ADR-0008 — M0 Owner Root and Recovery Trust Boundary

## Context

M0 must distinguish owner authority from raw technical access, authenticate governing state independently of the operational DB, detect DB-only rollback, prevent backward wall-clock movement from reviving expired permission and ensure historical backups cannot silently restore current authority.

A root/key stored only inside the operational DB cannot provide an independent trust boundary for that DB.

The documentary R4 research proposed a random Owner Root Key plus an external authenticated trust high-water. The required executable uncertainty—two durable trust domains under crash/rollback/time/restore behavior—was tested by `SPK-AURORA-M0-OWNER-TRUST-002` and independently reviewed.

## Decision drivers

- local-first/no cloud identity requirement;
- owner authority distinct from DB/filesystem technical access;
- passphrase rotation without redefining owner/domain identity;
- authenticated governing-state integrity;
- DB-only rollback detection within the accepted local threat boundary;
- deterministic DB/anchor crash classification;
- fail-closed backward-time handling;
- explicit post-restore authority revalidation;
- recovery to a fresh machine;
- bounded operational complexity compatible with M0.

Affected requirements include `REQ-032..045`, `058..076`, `091`, `095`, `104`, `119`.

## Executable evidence

Evidence owners:

```text
DOC-AURORA-M0-R4-SPK002-EVIDENCE-RECEIPT
REVIEW-AURORA-M0-R4-SPK002-OWNER-TRUST-2026-08-07
```

Final disposable SPK-002 execution:

```text
source baseline: ba7a211b40e21bfcd9aa1b90e5ca3c21cd229d1b
spec blob: 0ffb6fa2b35014e34b4301365dc2d5a8d96f021d
execution revision: c76b96fee36878f15c54028b4ba1896f84ebdeca
workflow: 31219882882
matrix: Ubuntu 24.04 + Windows latest
result: S01-S12 PASS + recovery classifications PASS + secret hygiene PASS
```

The final evidence also includes a post-green adversarial hardening: ordinary governing mutations now preflight and reject `STATE_ROLLBACK`, `ANCHOR_LAG`, `TIME_UNTRUSTED` and `REVALIDATION_REQUIRED` rather than relying on callers to classify first.

## Options

### A — random ORK wrapped by passphrase-derived KEK

```text
random ORK
→ stable owner/domain root lineage
passphrase-derived KEK
→ only wraps/unlocks ORK
```

Executable SPK-002 candidate. PASS.

### B — passphrase directly derives long-lived integrity root

Rejected for M0. Candidate A evidence plus documentary analysis is conclusive enough that the canonical spike did not require a duplicate full B flow.

Direct derivation couples passphrase changes to the integrity root. Rotation would therefore either change root/key lineage or require a broader rekey/continuity ceremony across protected material. Candidate A demonstrates lower semantic coupling: passphrase/KEK can rotate while ORK and ORK-derived domain lineage stay stable.

### C — mandatory OS keychain/hardware root

Not selected for M0. This remains a strengthening/fallback class if the threat model expands to hostile total-local-admin replay or hardware-backed trust becomes a current requirement.

## Decision

**Accepted by the operator on 2026-08-07:** use a random **256-bit Owner Root Key (ORK)** as the stable local owner/domain integrity root.

### Root custody and unlock

```text
owner passphrase
→ Argon2id
→ 32-byte KEK
→ AES-256-GCM unwrap of random ORK
```

Evidence-qualified M0 Argon2id baseline:

```text
memory:       64 MiB
iterations:   3
parallelism:  4
KEK:          32 bytes
```

The root envelope MUST be versioned and production decoding MUST bound/allowlist KDF parameters before invoking Argon2id.

Changing the owner passphrase rewraps the same ORK under a newly derived KEK/salt. Passphrase rotation MUST NOT redefine Aurora owner/domain identity or ORK lineage.

### Purpose subkeys

ORK-derived purpose keys use:

```text
HKDF-SHA-256
```

Distinct purpose labels separate at least governing-state integrity and trust-anchor integrity.

### Governing-state authenticated integrity

The protected governing descriptor uses:

```text
HMAC-SHA-256
```

and binds the material authority/current-state fields needed to prevent raw DB mutation from becoming governing without ORK-derived integrity capability.

### External Owner Trust Store

A small authenticated trust artifact is physically separate from the operational SQLite DB and contains at least:

```text
owner/root identity binding
governing generation high-water
observed wall-time high-water
HMAC-SHA-256 authentication
```

“Observed wall-time high-water” is deliberate wording. M0 does not claim an external trusted-time oracle.

### Normal mutation preflight

An ordinary governing mutation MUST fail before writing unless:

```text
wrapped owner/root state is unlocked as required
DB governing descriptor HMAC verifies
trust-anchor HMAC verifies
no restore-revalidation marker/state is pending
DB generation == anchor generation
current wall time >= authenticated observed wall-time high-water
```

An anomalous durable pair cannot be treated as normal merely because an owner session exists.

### Crash and rollback classification

```text
DB generation == anchor
→ NORMAL, subject to authority status/expiry

DB generation > anchor
+ valid DB HMAC
→ ANCHOR_LAG
→ non-permitting
→ authenticated owner may explicitly reconcile anchor forward

DB generation < anchor
→ STATE_ROLLBACK
→ fail closed
→ no implicit anchor rollback

invalid DB HMAC
→ INVALID_DB_MAC
→ fail closed

invalid anchor HMAC
→ INVALID_ANCHOR_MAC
→ fail closed

missing current anchor outside restore flow
→ MISSING_ANCHOR
→ fail closed
```

The evidence-qualified ordinary mutation ordering is:

```text
trust preflight
→ durable SQLite governing mutation commit
→ if process dies before anchor publication: ANCHOR_LAG
→ write/fsync next trust-anchor temp
→ publish next anchor
```

Exact physical path names and production filesystem wrapper are implementation-design details; the two independent durable trust domains and classification semantics are not.

### Time rollback

```text
process-local elapsed duration
→ monotonic time where available

cross-restart authority expiry
→ UTC wall time + authenticated observed high-water

observed current wall time < authenticated high-water
→ TIME_UNTRUSTED
→ expiry-dependent authority non-permitting
→ ordinary governing mutation fail closed
```

This detects backward movement relative to previously authenticated observations. It does not prove objectively correct global time against total runtime/admin compromise.

### Historical/fresh-machine restore

Recovery/export may carry an **encrypted ORK recovery envelope** but MUST NOT import a historical trust high-water as proof of current freshness.

```text
restore authentic historical state + encrypted root envelope
→ owner may unlock/recover ORK lineage
→ current trust high-water is absent/not trusted as current
→ REVALIDATION_REQUIRED
→ no next safe action from historical grant
```

The restored historical grant MUST NOT authorize its own revalidation.

Only the authenticated owner recovery path may create a **new** post-restore authority revision/generation and establish a new current trust anchor. Only after that new revision is current may normal next-safe-action projection resume.

## Consequences

### Positive

- raw operational DB access alone is insufficient to mint valid governing state;
- DB-only rollback is detectable against an independent authenticated generation high-water;
- DB→anchor crash has an explicit recoverable `ANCHOR_LAG` state;
- passphrase rotation does not re-key owner/domain identity;
- old backups cannot silently revive permission;
- fresh-machine root recovery preserves ORK lineage without pretending historical freshness is current;
- future platform/hardware root protection can wrap/protect the same ORK without changing Core authority semantics.

### Negative

- M0 has a second small durable trust artifact outside the operational DB;
- DB and trust-anchor publication require explicit crash protocol/reconciliation;
- owner recovery introduces intentional unlock/revalidation steps;
- Argon2id intentionally consumes memory/latency;
- local high-water adds filesystem writes to authority-sensitive flows.

## Security and reliability limits

### Total local compromise

Purely local M0 cannot defeat an attacker able to capture/replay **all** local trust artifacts while also compromising owner secrets/runtime state.

The supported claim is narrower: DB compromise alone, DB-only rollback, partial trust rollback/loss, process crash, historical restore and backward time below authenticated high-water fail closed rather than silently minting/reviving current authority.

If hostile total-admin replay becomes a current requirement, revisit hardware/remote trust anchoring rather than stretching this ADR’s claim.

### KDF parameter parsing

Production implementation MUST reject unsupported or extreme Argon2id envelope parameters before KDF allocation. The evidence-qualified profile above is the initial allowed baseline; parameter changes require versioned evidence.

### Physical power-loss semantics

SPK-002 used real process kill on GitHub-hosted Linux/Windows filesystems. It does not emulate every storage-controller/write-cache/power-loss failure. Production/R7 environment proof must verify the chosen file publication + fsync/directory-sync semantics to the level claimed by the target operating environment.

## Performance/usability evidence

Final observed Argon2id derivation:

```text
Linux:   ~52.1 ms, ~64 MiB observed Go heap delta
Windows: ~53.4 ms, ~64 MiB observed Go heap delta
```

Bootstrap was ~62 ms Linux / ~131 ms Windows and passphrase rotation ~106 ms on both final hosted runners. These measurements are acceptable for the current owner unlock/recovery path and are not throughput SLOs.

## Compatibility / migration / rollback

The ORK is an Aurora security-domain primitive, not a storage-engine identity. The operational DB remains behind the Durable State Port.

Logical recovery/export may carry a versioned encrypted ORK envelope while deliberately excluding current trust-high-water freshness.

Future OS keychain, TPM, YubiKey or remote wrapping may protect the same ORK if justified, preserving owner/domain semantics.

## Validation status

SPK-002 requirement is satisfied:

```text
SPK-AURORA-M0-OWNER-TRUST-002
→ PASS
→ EVIDENCE_COMPLETE
→ REVIEWED
→ DECISION_INFORMED
→ CLOSED
```

ADR-0008 v0.2.0 was explicitly accepted by the operator through `DOC-AURORA-M0-R4-ADR0008-ACCEPTANCE`. It is now a governing R4 architecture decision. Acceptance does not authorize R5, Microdesign or production implementation.

## Reconsideration triggers

- local threat model expands to hostile administrator/root replay of all trust artifacts;
- hardware/remote trust becomes required;
- owner unlock/revalidation burden becomes operationally unacceptable;
- KDF profile becomes inappropriate for target devices;
- a later cross-platform fault suite contradicts SPK-002;
- target filesystem/power-loss model cannot support the required publication guarantees;
- owner identity model expands beyond the local M0 owner-root boundary.

## References

- `RESEARCH-AURORA-M0-OWNER-AUTHORITY-RECOVERY-R4-V1`
- `RESEARCH-AURORA-M0-PORTABILITY-INTEGRITY-R4-V1`
- `DOC-AURORA-M0-R4-SPK002-EVIDENCE-RECEIPT`
- `REVIEW-AURORA-M0-R4-SPK002-OWNER-TRUST-2026-08-07`
- `DOC-AURORA-M0-R4-ADR0008-ACCEPTANCE`
- `REVIEW-AURORA-M0-R4-ARCHITECTURE-DECISION-READINESS-RERUN-2026-08-07`
- `SPK-AURORA-M0-OWNER-TRUST-002`
- RFC 9106, RFC 5869, NIST GCM/HMAC/SHA sources captured by the R4 research manifests
