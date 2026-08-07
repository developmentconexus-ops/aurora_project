---
id: RESEARCH-AURORA-M0-OWNER-AUTHORITY-RECOVERY-R4-V1
title: Aurora M0 R4 Research — Owner Root, Authority Freshness and Recovery Trust
_document_type_note: focused security architecture research
document_type: research_report
form: explanation
authority: research
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - M0 R4 owner-root, rollback/time and restore-freshness research through 2026-08-07
related:
  - DOC-AURORA-M0-R4-OPERATOR-AUTHORIZATION
  - DESIGN-AURORA-M0-R4-DECISION-LANDSCAPE
  - DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC
  - DOC-AURORA-CAP-SOVEREIGN-CORE-THREAT-MODEL
source_manifest: AURORA-RESEARCH-M0-OWNER-AUTHORITY-RECOVERY-R4-v1.sources.json
review_triggers:
  - owner credential mechanism change
  - integrity key custody change
  - restore-freshness or rollback finding
  - target host/security boundary change
last_reviewed: 2026-08-07
---

# Aurora M0 R4 — Owner Root, Authority Freshness and Recovery Trust

## 1. Research question

How can a single-user local Aurora establish that **Leandro the owner**—not merely a process with filesystem/database access—authorized sensitive Core operations, while also detecting common rollback/corruption scenarios and preventing stale backup authority from silently becoming current?

The mechanism must remain local-first and must not require cloud identity, a model provider or a future effect-plane capability.

## 2. Threat-boundary interpretation

R3 already excludes total compromise of the host/administrator trust root from M0's solved threat set. Therefore M0 does not need to defeat an owner/root adversary who can snapshot and replay every local trust file and runtime secret.

It does need to defend against:

- ordinary application/database access being treated as owner authority;
- an untrusted Project/runtime path directly changing authority;
- accidental or isolated rollback of the canonical database;
- state corruption/tampering without access to the owner root key;
- wall-clock rollback making expired authority appear valid;
- pre-revocation backup restore resurrecting permission;
- crash between canonical DB commit and external trust-anchor update;
- owner passphrase rotation without redefining Aurora identity or re-keying all domain meaning.

This boundary is important: demanding secure rollback resistance against complete local-root compromise would require hardware/remote monotonic trust and would materially change M0 scope.

## 3. Owner authentication alternatives

### A — OS login/user identity alone

Rejected for M0 owner authority.

Reason: being able to launch a process or access a database is explicitly not equivalent to authority. OS login can remain an environmental signal but cannot be the sole owner-authentication proof.

### B — passphrase directly derives every long-lived integrity key

Technically feasible with Argon2id [S01][S02].

Problems:

- passphrase rotation changes every derived long-lived key;
- old backups become coupled to old passphrase/key versions;
- integrity-key identity becomes the human password itself;
- recovery/key-version bookkeeping grows quickly.

### C — random Owner Root Key wrapped by a passphrase-derived key

Recommended.

Model:

```text
bootstrap
→ generate random 256-bit Owner Root Key (ORK)
→ derive Key-Encryption Key (KEK) from owner passphrase with Argon2id
→ wrap ORK with authenticated encryption
→ store only wrapped ORK + KDF parameters/salt at rest

unlock
→ derive KEK from entered passphrase
→ authenticated-decrypt ORK
→ derive purpose-specific subkeys from ORK using HKDF
```

Go's `crypto/rand` provides the random-root source [S03]. Argon2id is the password KDF [S01][S02]. AES-GCM is a standard authenticated-encryption mechanism and is directly supported by Go's standard library [S04][S05][S06]. HKDF provides independent labeled subkeys [S07][S08].

The key architectural benefit is that **passphrase rotation rewraps ORK rather than changing ORK**. Domain identity/integrity key lineage can therefore remain stable while the human unlock secret changes.

## 4. Proposed Owner Trust Store

M0 should maintain a small durable trust artifact physically separate from the operational-state database:

```text
OwnerTrustStore
- trust_store_version
- owner_id
- owner_root_key_version
- argon2id salt + parameters
- wrapped_owner_root_key + AEAD nonce/metadata
- highest_accepted_state_generation
- highest_accepted_authority_revision or authority generation
- highest_trusted_wall_time_utc
- anchor_payload_version
- anchor_mac
```

This is **security metadata**, not Project canonical state and not conversational memory.

The trust store is protected by filesystem permissions appropriate to the target host, but its cryptographic authenticity comes from HMAC under a subkey derived from ORK, not from permissions alone [S09][S11].

The exact path/permissions/locking API is implementation design, not R4 domain meaning.

## 5. Key separation

After ORK unlock, derive independent keys with HKDF-SHA-256 [S07][S08]:

```text
K_anchor    = HKDF(ORK, info="aurora/m0/trust-anchor/v1")
K_state_mac = HKDF(ORK, info="aurora/m0/state-integrity/v1")
K_recovery  = HKDF(ORK, info="aurora/m0/recovery/v1")
```

Do not use raw ORK directly as an HMAC/AES key across all purposes.

This keeps future key rotation/versioning explicit.

## 6. Authenticated state-generation model

The operational DB should expose one monotonically increasing logical `state_generation` independent from human-readable state revision names.

For every material accepted commit, the transaction stores a protected descriptor such as:

```text
ProtectedStateDescriptor
- aurora_id
- project_id/current revision references
- current authority-state revision
- state_generation
- relevant logical schema versions
- canonical digest references
- state_mac_algorithm/version
- state_mac
```

`state_mac` is HMAC-SHA-256 over the defined JCS-canonical protected descriptor using `K_state_mac` [S09][S11][S12].

This is not intended to MAC every database byte. It authenticates the **governing invariants needed to decide whether the loaded current state can be trusted**.

## 7. External rollback anchor

If both `state_generation` and its MAC live only in the operational DB, replaying an old valid database also replays its old valid MAC.

Therefore the Owner Trust Store keeps the highest accepted generation outside that database.

At startup after owner unlock:

```text
DB generation == trust-anchor generation
→ verify MACs
→ normal recovery may proceed

DB generation < trust-anchor generation
→ database rollback or loss suspected
→ fail closed as STATE_ROLLBACK / STATE_UNTRUSTED

DB generation > trust-anchor generation
→ possible crash after DB commit but before anchor update
→ verify DB HMAC and committed-state invariants
→ enter ANCHOR_LAG reconciliation
→ only authenticated owner recovery may advance anchor
```

This provides useful rollback detection against database-only rollback while preserving a safe crash-recovery path.

### Residual limitation

An actor who can replay **both** the operational DB and the complete Owner Trust Store can defeat this purely local high-water comparison. Preventing that under hostile local-root compromise would require a stronger external/hardware/remote monotonic anchor, outside current M0 scope.

This residual risk must remain explicit rather than hidden.

## 8. Cross-file crash ordering

The operational database transaction and Owner Trust Store cannot be assumed to share one atomic filesystem transaction.

Recommended order:

```text
1. validate owner/current state/authority
2. compute new protected descriptor + HMAC
3. commit canonical DB transaction with state_generation = N
4. fsync/durability semantics supplied by selected store
5. update authenticated Owner Trust Store high-water to N
6. durably replace/fsync trust-store artifact
7. only then report trust-anchor synchronization complete
```

If process death occurs after step 3 but before step 6, the next unlock sees `DB generation > anchor` and follows the `ANCHOR_LAG` path.

This protocol is a **candidate** until the security/recovery Architecture Spike proves it under actual filesystem/runtime behavior.

## 9. Owner passphrase KDF

RFC 9106 recommends Argon2id and gives a 64 MiB / time=3 / parallelism=4 profile for environments where the 2 GiB profile is unsuitable [S01]. Go's current `x/crypto/argon2` exposes that same guidance [S02].

R4 recommendation:

```text
baseline profile for spike:
Argon2id
memory = 64 MiB
iterations/time = 3
parallelism = 4 (subject to target CPU availability)
output KEK = 32 bytes
random salt per wrapped-key version
```

The Architecture Spike must measure actual unlock latency/resource use on the target environment. If parameters must be adjusted, security rationale and resulting profile version must be recorded rather than silently reduced.

## 10. ORK wrapping

Recommended current mechanism:

```text
KEK = Argon2id(passphrase, salt, versioned parameters)
ORK = random 32-byte key
wrapped_ORK = AES-256-GCM(KEK, random nonce, ORK, associated metadata)
```

Associated authenticated data should bind at least the envelope format version and owner identity reference so encrypted key material cannot be silently transplanted between unrelated owner roots.

AES-GCM provides authenticated encryption [S04][S05][S06]. Nonce generation uses cryptographically secure randomness [S03].

## 11. Time and expiry semantics

Go `time.Now()` can carry a process-local monotonic reading, but serialized `Time` values do not preserve a cross-process monotonic clock [S10]. Therefore two time modes are required:

### In-process durations

Use Go monotonic-time behavior for elapsed durations and deadlines within the same process where applicable.

### Cross-restart authority expiry

Persist UTC wall-clock expiry as domain data and maintain an authenticated `highest_trusted_wall_time_utc` in the Owner Trust Store.

On startup/unlock:

```text
if current_wall_time + allowed_skew < highest_trusted_wall_time
→ TIME_UNTRUSTED
→ expiry-dependent authority fails closed
→ owner must correct/revalidate time before permission resumes
```

R4 should not invent a global distributed clock. The allowed skew/tolerance is an implementation/configuration parameter that the spike/test plan must justify; the security invariant is simply that a meaningful backward step cannot make expired authority valid again.

After a normal trustworthy later time is observed, the high-water may advance but never move backward automatically.

## 12. Restore freshness

A backup/export represents a historical snapshot. Its valid old HMAC proves only that it was once authentic, not that its grants are still current.

Therefore R3's default remains mandatory:

```text
restore historical authority
→ REVALIDATION_REQUIRED
→ no restored grant can authorize its own revalidation
→ authenticated owner unlock/recovery
→ operator explicitly revalidates current intended authority
→ create a new authority-state revision
→ establish new local trust-anchor high-water
```

The restored package may include the encrypted ORK recovery envelope and KDF metadata needed to authenticate the owner on a replacement machine. It must **not** import an old trust-anchor high-water as proof of current freshness.

This is the core distinction between **authentic historical state** and **current authority**.

## 13. Recovery/key-loss model

M0 should treat owner-root loss as a material sovereignty incident.

The logical export/recovery package must preserve enough encrypted owner-root recovery material for migration to a fresh machine. The secret unlock passphrase itself is never stored in the export.

If neither a usable encrypted ORK envelope nor the correct owner recovery secret exists, Aurora cannot fabricate a replacement authority history. Recovery requires an explicit new root/bootstrap path whose relationship to the old Aurora identity must be governed and audited; exact emergency-root procedure can be deferred if it is outside the first implementation contract, but silent root replacement is prohibited.

## 14. Alternatives rejected/deferred

### OS keychain as mandatory M0 authority root

Deferred.

OS credential stores can improve user experience later, but choosing Windows/macOS/Linux-specific keychain semantics now would couple the first Core security model to a host platform. The wrapped-ORK architecture allows a future keychain/hardware module to protect the same ORK without changing domain meaning.

### Hardware TPM/YubiKey mandatory for M0

Deferred.

Useful future hardening, especially against local key theft/rollback, but not required to prove the current single-user walking skeleton and would create hardware/platform setup not present in M0 requirements.

### automatic authority freshness after restore

Rejected for M0.

Without a current external revocation source, restored state cannot prove that a historical grant was not revoked after backup creation. Default explicit owner revalidation is safer and simpler.

## 15. Proposed R4 decisions

Documentary evidence is strong enough to recommend the architecture shape:

```text
R4-Q-AUTHN-001
→ explicit owner unlock independent from DB access
→ random 256-bit Owner Root Key
→ Argon2id passphrase-derived KEK
→ AES-256-GCM wrapped ORK
→ HKDF-derived purpose-specific subkeys

R4-Q-TIME-001
→ process-local monotonic time for elapsed durations
→ authenticated wall-time high-water across restarts
→ meaningful backward step => TIME_UNTRUSTED / fail closed

R4-Q-RESTORE-001
→ restored historical grants always REVALIDATION_REQUIRED unless a future stronger freshness mechanism is separately proven
→ owner-authenticated explicit revalidation creates new authority revision

R4-Q-INTEGRITY-001
→ SHA-256 for non-secret content fingerprints
→ HMAC-SHA-256 for authenticated governing descriptors/trust anchors
→ Owner Trust Store physically outside operational DB
```

However, the **cross-file trust-anchor update/crash protocol and actual rollback/time behavior are implementation-dependent and must be proven by an Architecture Spike before these high-lock-in security decisions become accepted**.

## 16. Required security/recovery Architecture Spike

The spike must implement only enough to exercise the candidate owner-root/trust protocol and prove:

- correct passphrase unlock succeeds; wrong passphrase cannot produce valid ORK;
- root key remains encrypted at rest;
- passphrase rotation rewraps ORK without changing root-key identity;
- direct DB mutation without ORK cannot create a valid protected descriptor;
- rollback DB while trust anchor remains current is detected;
- crash after DB commit before trust-anchor update produces recoverable `ANCHOR_LAG`, not rollback/success ambiguity;
- anchor-ahead/DB-behind fails closed;
- wall-clock rollback triggers `TIME_UNTRUSTED` for expiry-dependent authority;
- old backup restore becomes `REVALIDATION_REQUIRED` even when historically authentic;
- only authenticated owner can create the new post-restore authority revision;
- fresh-machine recovery from encrypted root envelope works;
- losing root/recovery material fails explicitly.

## 17. Decision implications

If the spike validates the candidate protocol:

```text
owner authority remains independent from raw technical access
+ state rollback becomes detectable in the intended threat boundary
+ time rollback fails closed
+ historical backups cannot resurrect permission
+ passphrase rotation does not redefine owner/root key identity
```

If the local trust-anchor protocol cannot be made crash-safe and operationally reasonable, R4 must reconsider a stronger OS/hardware/remote trust anchor rather than weakening rollback/authority requirements.

## 18. Limitations

This report does not claim protection against:

- hostile root/admin able to snapshot/replay every local file and memory secret;
- compromised kernel/hardware RNG;
- physical extraction from an unlocked process;
- operator intentionally reauthorizing unsafe historical state;
- future remote/multi-Presence identity threats.

Those are outside M0 or require later hardening.
