---
id: REVIEW-AURORA-M0-R4-SPK002-OWNER-TRUST-2026-08-07
title: M0 R4 SPK-002 Owner Trust Evidence Review
document_type: adversarial_review
form: explanation
authority: evidence
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - reviewed SPK-002 evidence and decision-informed recommendation
related:
  - DOC-AURORA-M0-R4-SPK002-EVIDENCE-RECEIPT
  - DESIGN-AURORA-M0-OWNER-TRUST-SPIKE-002
  - ADR-AURORA-0008
  - DOC-AURORA-STATUS
spike_execution_revision: c76b96fee36878f15c54028b4ba1896f84ebdeca
workflow_run: 31219882882
reviewed_at: 2026-08-07
last_reviewed: 2026-08-07
---

# M0 R4 — SPK-002 Owner Trust Evidence Review

## 1. Verdict

```text
SPK-AURORA-M0-OWNER-TRUST-002: PASS

Lifecycle:
AUTHORIZED
→ EXECUTING
→ EVIDENCE_COMPLETE
→ REVIEWED
→ DECISION_INFORMED
→ CLOSED
```

The evidence is sufficient to answer the authorized R4 uncertainty within the explicitly bounded M0 local threat model.

Candidate A is recommended:

```text
random 256-bit Owner Root Key (ORK)
+ owner passphrase → Argon2id KEK
+ AES-256-GCM wrapped ORK at rest
+ HKDF-SHA-256 purpose subkeys
+ HMAC-SHA-256 governing descriptor and external trust anchor
+ authenticated generation / observed wall-time high-water outside the operational DB
+ fail-closed historical restore → REVALIDATION_REQUIRED
+ authenticated-owner-only post-restore authority revision
```

This review does **not** accept ADR-0008. ADR-0008 remains `proposed` until the operator accepts, rejects or revises the evidence-informed decision.

## 2. Gate-by-gate review

| Canonical gate | Reviewed evidence | Verdict |
|---|---|---|
| DB access alone cannot mint valid governing state | S03 raw DB mutation invalidates state HMAC | PASS |
| DB-only rollback is detected | S04 historical DB N against authenticated anchor N+1 → `STATE_ROLLBACK` | PASS |
| DB→anchor crash gap has one safe/recoverable classification | S05 real process kill before DB commit / after DB commit / after anchor temp sync / after anchor publish | PASS |
| anchor lag can be explicitly reconciled | valid DB N+1 + valid anchor N → `ANCHOR_LAG`, then authenticated `ReconcileAnchor` | PASS |
| anchor ahead / DB behind does not lower anchor | S06 → fail closed / `STATE_ROLLBACK`; reconciliation refuses rollback | PASS |
| backward time cannot revive expired permission | S07 authenticated high-water + earlier wall clock → `TIME_UNTRUSTED` / non-permitting | PASS |
| historical authentic restore does not become current permission | S08 → `REVALIDATION_REQUIRED`, current anchor not imported | PASS |
| restored grant cannot authorize itself | S09 `ActorRestoredGrant` denied | PASS |
| only authenticated owner creates post-restore revision | S10 authenticated owner → new authority revision/generation; only then next safe action | PASS |
| fresh-machine root recovery remains non-current | S11 correct owner secret recovers ORK but state remains `REVALIDATION_REQUIRED` | PASS |
| missing root cannot fabricate identity/authority | S12 explicit `MISSING_WRAPPED_ROOT`; no replacement root created | PASS |
| root encrypted at rest | S01 on-disk scan finds neither plaintext passphrase nor plaintext ORK | PASS |
| passphrase rotation preserves domain root | S02 old passphrase rejected; new passphrase unlocks same ORK fingerprint | PASS |
| KDF cost documented and usable | 64 MiB / t=3 / p=4; ~52–53 ms KDF on final hosted runners | PASS |
| no secret in final evidence artifacts | Linux and Windows evidence-hygiene jobs PASS; measurement binary excluded from artifacts | PASS |
| cross-platform reproduction | final Ubuntu and Windows jobs PASS under Go 1.26.5 / CGO=0 | PASS |

## 3. Crash protocol review

The final prototype uses this ordering for ordinary accepted governing mutation:

```text
preflight normal trust state
→ begin SQLite transaction
→ validate current authenticated governing state
→ write new authority revision + governing descriptor HMAC
→ SQLite COMMIT
→ classify crash here as ANCHOR_LAG if anchor still N
→ write next anchor to temporary file
→ fsync temporary anchor file
→ atomic publish/rename as current anchor
```

The fault suite kills a child process at:

```text
before DB commit
after DB commit
after anchor temp fsync
after anchor publish
```

Observed outcomes remain deterministic:

- pre-commit death leaves the previous coherent generation;
- post-DB/pre-anchor death leaves a valid DB generation ahead of anchor and is classified `ANCHOR_LAG`;
- authenticated owner reconciliation advances the anchor rather than silently rolling state back;
- after anchor publication the new generation is coherent.

This is the intended M0 two-durable-domain recovery protocol.

## 4. Adversarial mutation-boundary finding

The first complete green matrix exposed an architectural review question not covered strongly enough by the original happy-path API surface: an already-unlocked owner session could invoke the ordinary mutation function directly while the DB/anchor pair was anomalous.

That would have made classification correct but enforcement too dependent on a caller remembering to evaluate first.

A RED was added for four conditions:

```text
STATE_ROLLBACK
ANCHOR_LAG
TIME_UNTRUSTED
REVALIDATION_REQUIRED
```

The final implementation requires ordinary `Advance` to preflight and reject all four. Explicit recovery remains separate:

```text
ANCHOR_LAG            → authenticated ReconcileAnchor
REVALIDATION_REQUIRED → authenticated-owner Revalidate
STATE_ROLLBACK        → no implicit repair
TIME_UNTRUSTED        → no ordinary mutation while time is behind high-water
```

The focused hardening test passed, then the full hardened Ubuntu/Windows matrix passed again. This finding is therefore:

```text
SPK002-F01 — operational mutation could bypass anomaly classification
severity: material before fix
status: RESOLVED IN SPIKE EVIDENCE
implication: ADR-0008 should require enforcement at the mutation boundary, not merely diagnostic classification
```

## 5. Evidence-hygiene finding

An earlier evidence run built the sanitized measurement executable inside the artifact directory. Compiled deterministic fixture strings caused the scanner to fail that artifact set.

No production/user secret was involved, and the protocol tests themselves remained green, but a failing hygiene gate is not acceptable final evidence.

The executable was moved to the runner temporary directory and deleted before artifact upload. The complete matrix was rerun and both platform hygiene checks passed.

Disposition:

```text
SPK002-F02 — measurement binary contaminated evidence artifact with fixture strings
status: RESOLVED
final evidence: run 31219882882 only
```

## 6. Candidate B review

Candidate B directly derives the long-lived integrity root from the owner passphrase.

The canonical spike says not to build a duplicate full Candidate B flow if documentary analysis plus Candidate A evidence is conclusive. That condition is met.

The primary architectural requirement is passphrase rotation without redefining owner/domain identity. Candidate A demonstrates:

```text
passphrase-derived KEK changes
→ ORK remains unchanged
→ all ORK-derived integrity/key lineage remains stable
```

With Candidate B, changing the passphrase changes the directly derived root unless the system performs a broader rekey/migration/continuity ceremony across protected material. Preserving an old passphrase-derived root would instead defeat the intended rotation semantics.

Therefore Candidate B is not a simpler recovery/rotation architecture for M0 and no duplicate executable flow is justified.

## 7. Candidate C review

Mandatory platform keychain/hardware root was a fallback only if Candidate A could not satisfy the intended local threat boundary.

Candidate A passes all required gates without cloud/hardware trust. Candidate C is therefore not escalated in M0.

Future OS-keystore/TPM/YubiKey protection may wrap/protect the same ORK if the threat model later changes, without changing the semantic distinction between owner root and operational DB state.

## 8. Time semantics review

The evidence supports an **authenticated observed wall-time high-water**, not an externally trusted time oracle.

What is proven:

```text
once a later wall time T has been authenticated into the trust high-water,
a later process observing wall time < T fails closed for expiry-dependent permission.
```

What is not proven:

- the local wall clock is objectively correct before a later value has ever been observed;
- a fully compromised runtime/administrator cannot lie consistently about time and all trust artifacts;
- network/external time attestation.

The ADR must use this narrower language to avoid overstating local trust.

## 9. KDF/resource review

Final evidence-qualified profile:

```text
Argon2id
memory = 64 MiB
iterations = 3
parallelism = 4
KEK = 32 bytes
```

Observed KDF derivation:

```text
Linux:   ~52.1 ms / ~64 MiB observed Go heap delta
Windows: ~53.4 ms / ~64 MiB observed Go heap delta
```

Bootstrap and rotation remain sub-second on both hosted target runners. For an owner unlock/recovery path with no current high-frequency SLO, the measured burden is acceptable.

Production implementation must reject unsupported/extreme envelope KDF parameters before invoking Argon2id; otherwise a corrupted/malicious envelope could turn parameter parsing into local resource exhaustion. This is an R6 hardening constraint, not an authority-bypass finding against the R4 architecture.

## 10. Required residual risk

Pure local trust cannot provide absolute rollback resistance against a hostile actor who can replay **all** local trust artifacts while also controlling owner secrets/runtime state.

The evidence supports the narrower M0 claim:

```text
DB compromise alone
DB-only rollback
partial trust-file loss/rollback
process crash between DB and anchor
historical backup restore
backward wall-clock movement below authenticated high-water

→ fail closed rather than silently mint/revive current authority
```

This limitation is material and must remain in ADR-0008 and future threat-model/operations documentation.

## 11. Physical durability limitation

The final fault injection uses real process termination on GitHub-hosted Linux and Windows filesystems. It does not emulate every literal power-loss, kernel, filesystem, storage-controller or write-cache failure.

The spike fsyncs the temporary trust-anchor content before rename. Exact directory-sync/publish durability for the production target remains an implementation/environment validation obligation.

This residual does not block the R4 architecture because the canonical spike explicitly scopes its fault model to controlled process kill, but R7 Golden Proof must not overclaim hardware-level persistence guarantees without matching evidence.

## 12. Recommended R4 decision

Evidence supports ADR-0008 revision 0.2.0 with the following architecture:

```text
Owner identity root:
  random 256-bit ORK

owner unlock / root-at-rest:
  Argon2id-derived KEK
  evidence-qualified baseline = 64 MiB / t=3 / p=4 / 32-byte KEK
  AES-256-GCM wrapped ORK
  versioned/bounded KDF profile

purpose keys:
  HKDF-SHA-256

governing integrity:
  HMAC-SHA-256 descriptor bound to generation/owner/authority state

local trust high-water:
  physically outside operational DB
  HMAC-SHA-256 authenticated
  generation + observed wall-time high-water

normal mutation boundary:
  must preflight DB MAC / anchor MAC / generation equality / time high-water / restore state

DB ahead of anchor:
  ANCHOR_LAG
  non-permitting until authenticated owner reconciliation

DB behind anchor:
  STATE_ROLLBACK
  fail closed; no implicit anchor rollback

backward observed time:
  TIME_UNTRUSTED
  expiry-dependent permission and ordinary governing mutation fail closed

historical/fresh restore:
  encrypted root recovery envelope may restore ORK lineage
  current trust high-water is not imported
  state becomes REVALIDATION_REQUIRED
  restored grant cannot self-authorize
  authenticated owner creates a new authority revision/generation
```

## 13. R4 implication

SPK-002 is no longer an executable-evidence blocker:

```text
SPK-AURORA-M0-OWNER-TRUST-002
→ PASS
→ EVIDENCE_COMPLETE
→ REVIEWED
→ DECISION_INFORMED
→ CLOSED
```

ADR-0008 remains `proposed` and is now evidence-ready for operator decision.

Current R4 state after this review:

```text
all required M0 architecture spikes: CLOSED with reviewed PASS
ADR-0003/0004/0005/0006/0007 accepted
ADR-0009 accepted cross-horizon
ADR-0008 evidence-ready, operator decision pending
R4: BLOCKED only on ADR-0008 operator decision
R5: NOT AUTHORIZED
```

No R5, Mission Contract, Microdesign or production security implementation is authorized by this review.
