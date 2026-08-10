---
id: DOC-AURORA-CAP-SOVEREIGN-CORE-THREAT-MODEL
title: CAP-SOVEREIGN-CORE Threat Model
document_type: threat_model
form: explanation
authority: specification
status: accepted
accepted_at: 2026-08-09
acceptance_evidence: DOC-AURORA-M0-R5-A2-CONTRACT-OPERATOR-ACCEPTANCE
accepted_from_blob: 7e97f816d0c4966ba6b12cf0447c7a2210fbea34
version: 0.2.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - security and threat analysis for CAP-SOVEREIGN-CORE
related:
  - DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC
  - DOC-AURORA-CAP-SOVEREIGN-CORE-REQUIREMENTS
  - DOC-AURORA-CAP-SOVEREIGN-CORE-TEST-PLAN
  - DOC-AURORA-CAP-SOVEREIGN-CORE-R3-COVERAGE
  - DOC-AURORA-CAP-SOVEREIGN-CORE-R4-DECISION-COVERAGE
  - ADR-AURORA-0007
  - ADR-AURORA-0008
  - DOC-AURORA-MIS-M0-SOVEREIGN-CORE-001
  - DOC-AURORA-M0-R5-A2-CONTRACT-OPERATOR-ACCEPTANCE
source_revision: 9ea8adf5c115f54071d7e36e312695d19420d8b0
r4_alignment_revision: 74167bd1404d9076423ffdbae20f97958283527c
review_triggers:
  - R4 mechanism selection
  - new network/provider boundary
  - authentication or integrity mechanism change
  - storage/recovery topology change
  - security finding or incident
last_reviewed: 2026-08-07
---

# CAP-SOVEREIGN-CORE — Threat Model

## 1. Purpose and scope

This threat model covers the M0 sovereign Core slice defined by `SPEC.md`:

- stable Aurora/operator/Project identities;
- revisioned Project operational state;
- M0 authority state/snapshots and next-safe-action projection;
- restart/recovery;
- export/restore/migration;
- audit/evidence/telemetry boundaries;
- one local operator interaction boundary.

It does not model future cloud providers, Harness execution, external effects, credentials, multi-Presence, physical devices or multi-tenant operation. Introducing any of those boundaries requires threat-model revision in the applicable later gate.

The original R3 revision intentionally did not select mechanisms. This v0.2.0 preserves the threat semantics and records the accepted R4 mitigations from ADR-0007/ADR-0008; exact implementation controls remain R6 work.

---

## 2. Security objectives

M0 must preserve:

1. **Identity integrity** — Aurora/Project identities cannot change accidentally or through restore collision.
2. **Canonical-state integrity** — current accepted state cannot be silently overwritten, rolled back or fabricated.
3. **Authority integrity** — access is not authority; expiry/revocation/restore safety fail closed.
4. **Availability with explicit degradation** — unavailable/corrupt state yields diagnosable degraded/blocked failure, not invented truth.
5. **Sovereignty** — canonical state and backup/export remain under Leandro-governed control.
6. **Auditability/accountability** — material operations are attributable and evidence remains distinct from claims/logs.
7. **Confidentiality proportional to class** — project/authority/export data is not leaked through general telemetry or ordinary logs.
8. **Recoverability** — restart/restore/migration preserve protected semantics or fail safely.

---

## 3. Protected assets and data classes

| Asset | Classification floor | Security concern |
|---|---|---|
| `AuroraIdentity` | `INTERNAL` | substitution/collision destroys continuity |
| operator identity/authentication provenance | `SENSITIVE` | spoofing can mint owner authority |
| Project operational state | `CONFIDENTIAL` | tampering changes current truth; disclosure may expose private context |
| authority state/grants/snapshots | `SENSITIVE` | tampering/replay can enable unauthorized mutation |
| next-safe-action projection | `CONFIDENTIAL` | stale projection can misrepresent current permission |
| audit records | `CONFIDENTIAL` | tampering hides causal/accountability chain |
| evidence metadata/references | `CONFIDENTIAL` minimum | spoofing can create false proof; payload may inherit higher class |
| telemetry/correlation | `INTERNAL` | sensitive-payload leakage and false operational inference |
| export/backup package | `SENSITIVE` minimum | aggregate state/authority exposure and rollback risk |
| integrity descriptors | `INTERNAL` minimum | forgery can make corrupted state appear valid |

Export/backup inherits any higher classification of contained material.

---

## 4. Trust boundaries

### TB-01 — Operator Adapter → Sovereign Core

Untrusted assumption: possession of the CLI/UI/process endpoint does not prove operator identity or authority.

Required boundary property:

- authenticated operator context follows the accepted ADR-0008 owner-root/bootstrap boundary;
- ordinary content/arguments cannot redefine policy/authority;
- inspection and mutation are distinct command semantics.

### TB-02 — Core → Durable State Port

The persistence mechanism may fail, return stale/corrupt data or expose non-atomic outcomes.

Required boundary property:

- structural/version/integrity validation before governing use;
- revision/precondition-safe mutation;
- explicit ambiguous/failure handling;
- no storage-specific type becomes domain authority.

### TB-03 — Core → Time Source Port

Authority expiry depends on time. Local clock rollback or stale time can extend permission.

Required boundary property:

- accepted ADR-0008 uses an authenticated observed wall-time high-water plus fail-closed backward-time semantics;
- unverifiable time sufficient to affect authority must fail closed or surface `REVALIDATION_REQUIRED`.

### TB-04 — Core → Integrity Port

An integrity descriptor is only useful if the selected mechanism cannot be trivially forged by the same failure/adversary it is supposed to detect.

Required boundary property:

- accepted ADR-0008 defines ORK-derived authenticated governing/trust descriptors and their local threat assumptions;
- failed or unavailable required integrity validation blocks governing restore/state use.

### TB-05 — Core → Export/Backup Medium

Export leaves the live canonical store and may be copied, exposed, tampered with or restored later after authority changed.

Required boundary property:

- package classification/integrity/version metadata;
- safe restore validation;
- authority freshness reset/revalidation when later revocation cannot be proven absent.

### TB-06 — Core → Audit/Evidence/Telemetry physical sinks

These sinks may be unavailable, sampled, delayed or tampered with and must not become canonical current state.

Required boundary property:

- required audit/evidence semantics remain distinguishable from optional telemetry;
- missing telemetry cannot become loss of domain truth;
- sensitive payloads do not ride general correlation channels.

---

## 5. Assumptions

R3 assumes only:

- M0 is single-user/Leandro-first;
- no external effect plane is active;
- no cloud/provider/Harness is required for canonical recovery;
- accepted R4 mechanisms provide the M0 local SQLite store, owner-root authentication boundary, time-high-water behavior and authenticated integrity model needed by the logical contracts;
- physical compromise of all trusted local infrastructure is outside what pure application semantics can fully prevent.

These assumptions are explicit inputs to R4 and must be revisited if the implementation environment differs.

---

## 6. Threat sources

Potential sources include:

- accidental operator error;
- buggy Core or migration implementation;
- compromised/untrusted Project content;
- stale/corrupted persistence;
- malicious or accidental manipulation of export/backup files;
- local process/user with technical access but without product authority;
- compromised local host or administrator-level attacker;
- clock rollback/misconfiguration;
- telemetry/logging misconfiguration;
- future adapter/framework attempting to become an authority/state owner.

M0 does not assume an LLM is trustworthy for authority, recovery or integrity decisions.

---

## 7. Threat matrix

| ID | Threat | Impact | Required mitigation/behavior | Accepted R4 binding / later obligation |
|---|---|---|---|---|
| `TM-01` | canonical Project state is corrupted or tampered | false current truth | validate structure/version/integrity; block governing use on failure; preserve evidence | storage/integrity mechanism |
| `TM-02` | stale state replaces newer current state | rollback of accepted truth | revision/predecessor/current-pointer invariants; stale transition rejection; restore collision checks | rollback-detection/storage semantics |
| `TM-03` | revoked/expired authority becomes active after restart | unauthorized mutation | current authority evaluation against status/scope/time on every relevant decision | trustworthy time semantics |
| `TM-04` | old backup predating revocation is restored | silent authority resurrection | restored apparently-active grants become `REVALIDATION_REQUIRED` unless freshness beyond later revocation can be proven | restore-freshness mechanism |
| `TM-05` | restore package uses different Aurora/Project identity over existing state | identity takeover/collision | explicit identity collision check; no silent merge/overwrite | restore workflow mechanism |
| `TM-06` | export/backup is disclosed | private state/authority exposure | `SENSITIVE` minimum package classification, controlled storage, no unnecessary secret inclusion | encryption/access-control choice |
| `TM-07` | export/state package is modified | corrupt/forged restore | integrity descriptor and validation before apply | checksum/signature/key model |
| `TM-08` | ordinary Project content contains policy-like instructions | authority injection | treat content as data; authority changes only through governed authority command | parser/input isolation implementation |
| `TM-09` | local caller with technical access impersonates Leandro | unauthorized owner operations | explicit `OperatorIdentityRef` authentication boundary; access != authority | local auth/bootstrap mechanism |
| `TM-10` | clock rolls back or is stale | expired grant appears valid | explicit time-source boundary; fail closed/revalidation when trustworthy validity cannot be established | time/rollback strategy |
| `TM-11` | crash occurs around state mutation | partial/ambiguous accepted state | no success until mutually consistent commit/evidence observed; ambiguous result reconciled before retry | crash-consistent commit mechanism |
| `TM-12` | migration silently changes domain meaning | identity/state/authority corruption | explicit version-pair migration, protected invariant checks and evidence; incompatible state fails | migration tooling/schema choice |
| `TM-13` | audit/evidence record is spoofed or diverges from state | false proof/accountability | stable references to state/authority revisions; Evidence separate from Claim/Verdict; current state not owned by audit | integrity/evidence storage mechanism |
| `TM-14` | telemetry/logs contain sensitive state/authority payload | confidentiality leak | identifiers-only correlation by default; sensitive evidence separate from general telemetry | telemetry backend/redaction |
| `TM-15` | telemetry/audit backend unavailable | false belief that domain state is unavailable or unverified | telemetry not canonical truth; required evidence/audit failure explicitly blocks only criteria that require it | sink durability/availability |
| `TM-16` | adapter/framework/runtime internal state is used as recovery authority | sovereignty loss/lock-in | canonical state exclusively through Core owners/Durable State Port; adapter state non-authoritative | architecture/static boundaries |
| `TM-17` | retry repeats an ambiguous mutation/restore/migration | duplicate/corrupt state | retry only after classification and safe/idempotent proof; reconcile ambiguity first | idempotency/commit mechanism |
| `TM-18` | attacker or bug removes current state, leaving historical records | fabricated reconstruction risk | missing current owner/pointer is explicit failure; do not infer governing state from narrative/logs | backup/recovery mechanism |
| `TM-19` | authority snapshot or next-action cache is stale | unsafe permission presentation | projection carries source revisions/time; invalidate on state/authority/time change | caching implementation |
| `TM-20` | local host/admin fully compromised | broad confidentiality/integrity compromise | application-level provenance/fail-closed boundaries reduce accidental misuse but cannot guarantee protection from total host compromise | host hardening, OS/storage security; residual risk |

---

## 8. Abuse and adversarial cases required by R3

### AC-01 — stale transition replay

Replay a previously valid transition request after current state advances. Expected: reject on revision precondition, no canonical mutation.

### AC-02 — authority scope confusion

Use a grant for a different Project/action class. Expected: no permission and blocked next action.

### AC-03 — expired authority after restart

Persist an authority with a short validity boundary, terminate/restart after expiry. Expected: recovered grant record remains historical but effective status is `EXPIRED` and does not permit mutation.

### AC-04 — revoked authority after restart

Revoke authority, restart. Expected: revocation remains governing.

### AC-05 — pre-revocation backup restore

Export while authority is active, then revoke in live canonical state, then restore the older export into a fresh restore context. Expected: restored active-looking authority is not silently trusted; it becomes `REVALIDATION_REQUIRED` absent an R4-proven freshness mechanism.

### AC-06 — identity collision restore

Restore a package whose Aurora/Project identity conflicts with different current identity. Expected: reject or require explicit governed resolution; no silent merge.

### AC-07 — corrupt export

Modify material export bytes/state. Expected: integrity validation fails before state becomes current.

### AC-08 — untrusted content injection

Persist Project content containing instructions to grant permission or change policy. Expected: content remains data; no authority change.

### AC-09 — ambiguous crash during transition

Crash around a material state mutation. Expected: recovery either establishes one coherent accepted revision or reports ambiguous/failure; retry does not create a second accepted mutation.

### AC-10 — audit backend unavailable

Make non-canonical telemetry unavailable during inspection/recovery. Expected: domain truth still comes from canonical state; if required evidence cannot be produced, criterion is explicitly degraded/blocked rather than fabricated.

### AC-11 — time rollback

Simulate time moving backwards enough to affect expiry. Expected: the selected R4 time semantics must prevent expired authority from regaining permission or fail closed.

### AC-12 — migration semantic drift

Use a migration candidate that changes identity/authority meaning. Expected: invariant validation fails and prior governing state is not silently replaced.

---

## 9. Security controls expressed as semantic obligations

R3 requires these controls regardless of later implementation:

- stable immutable identities;
- explicit actor/source attribution;
- revision/precondition checks for mutation;
- owner-only authority administration semantics;
- scope/expiry/revocation evaluation;
- restore authority revalidation/freshness boundary;
- fail-closed authority/state handling;
- structural/version/integrity validation;
- identity collision detection;
- explicit migration compatibility;
- sensitive-data classification;
- no sensitive payload in general telemetry;
- audit/evidence separation from current state and verdict;
- no model/Harness/framework/database authority substitution;
- no blind retry after ambiguous mutation;
- explicit residual-risk reporting.

---

## 10. Accepted R4 security bindings

The R3 security questions are now resolved for M0 by accepted decisions and executable evidence:

| Concern | Accepted M0 binding | Later implementation/evidence obligation |
|---|---|---|
| local owner authentication/bootstrap | random 256-bit ORK; passphrase→Argon2id KEK; AES-256-GCM wrapped ORK | versioned envelope; bound/allowlisted KDF parameters |
| local state persistence/atomicity | SQLite + `database/sql` + `modernc.org/sqlite`, WAL, `synchronous=FULL` | exact schema/transaction wrapper and target-environment fault proof |
| storage rollback detection | external authenticated generation high-water + DB HMAC | enforce anomaly preflight at mutation boundary |
| authority time rollback | authenticated observed wall-time high-water; backward movement → `TIME_UNTRUSTED` | exact Time Source abstraction and diagnostics |
| governing integrity | HKDF-SHA-256 purpose keys + HMAC-SHA-256 | exact protected descriptor fields/canonicalization |
| portable export confidentiality | logical JSON/JCS package; normal SENSITIVE outer package protected with `age` | exact archive/library/recipient handling |
| restore freshness | restored active-looking authority → `REVALIDATION_REQUIRED`; owner-only new revision | exact recovery UX/API; never import historical high-water as current |
| migration | explicit version-pair application-owned migration | exact runner/rollback and invariant verification |
| audit/evidence | logically distinct; may be co-located transactionally for M0 | exact persistence/retention fields |
| telemetry | OTel traces/metrics + `slog`, exporter optional/non-authoritative | redaction and exporter-failure proof |

No new security mechanism choice is delegated silently to R6. R6 owns only the implementation details inside these accepted boundaries.

---

## 11. Residual risks and carry-forward constraints

Accepted R4 architecture does not eliminate every local threat.

M0 still does not claim to solve:

- total compromise of host/administrator plus replay of **all** local trust artifacts and owner secrets/runtime;
- physical destruction/theft of every state/recovery copy;
- objectively trusted global time — only backward movement below an authenticated observed high-water is detected;
- every literal storage-controller/write-cache/power-loss failure beyond the tested process-kill model;
- accidental owner decision to explicitly reauthorize unsafe historical intent after receiving the required revalidation state;
- future network/provider/device threats outside M0.

Accepted controls that must carry into R6/R7:

- `age`-protected normal SENSITIVE portability envelope and explicit export classification;
- ORK/HKDF/HMAC integrity boundary independent from the operational DB;
- mutation-boundary enforcement for rollback/anchor-lag/time/restore anomalies;
- bounded Argon2id parameter parsing before allocation;
- explicit filesystem publication/fsync/directory-sync design for claimed target durability;
- secret/redaction hygiene in logs/evidence.

These limitations MUST remain visible in the Mission Contract and evidence. They are not permission to weaken fail-closed semantics.

---

## 12. Threat-model gate conclusion

The R4-aligned threat model is complete enough to support the proposed M0 Mission Contract because:

- protected assets/classes are explicit;
- trust boundaries are explicit;
- the six minimum R2 threat classes are covered;
- additional M0-specific rollback/time/atomicity risks are covered;
- required semantic mitigations are allocated;
- accepted mechanism bindings and remaining implementation/evidence obligations are explicit;
- no unresolved M0 security mechanism choice is silently delegated to R6.

This conclusion remains subject to the accepted R3 review, accepted R4 ADRs/spike evidence and the R5 Contract Readiness review; the Threat Model does not authorize implementation.
