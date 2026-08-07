---
id: DOC-AURORA-CAP-SOVEREIGN-CORE-THREAT-MODEL
title: CAP-SOVEREIGN-CORE Threat Model
document_type: threat_model
form: explanation
authority: specification
status: proposed
version: 0.1.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - proposed R3 security and threat analysis for CAP-SOVEREIGN-CORE
related:
  - DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC
  - DOC-AURORA-CAP-SOVEREIGN-CORE-REQUIREMENTS
  - DOC-AURORA-CAP-SOVEREIGN-CORE-TEST-PLAN
  - DOC-AURORA-CAP-SOVEREIGN-CORE-R3-COVERAGE
source_revision: 9ea8adf5c115f54071d7e36e312695d19420d8b0
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

This document is part of the proposed R3 Capability Spec package and MUST NOT select R4 technologies.

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

- authenticated operator context is established by an R4-selected mechanism;
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

- R4 must choose a time/rollback approach proportionate to M0;
- unverifiable time sufficient to affect authority must fail closed or surface `REVALIDATION_REQUIRED`.

### TB-04 — Core → Integrity Port

An integrity descriptor is only useful if the selected mechanism cannot be trivially forged by the same failure/adversary it is supposed to detect.

Required boundary property:

- mechanism and trust assumptions must be explicit in R4;
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
- a future R4 mechanism can provide some persistent local storage, operator-authentication boundary, time source and integrity check satisfying the logical contracts;
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

| ID | Threat | Impact | Required R3 mitigation/behavior | Residual/R4 dependency |
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

## 10. R4 security questions

The following are implementation-blocking technical decisions for R4, not unresolved R3 product semantics:

| Question | Why R4 must answer it |
|---|---|
| local operator authentication/bootstrap | prevent technical access from becoming owner authority |
| local state persistence and atomicity | preserve revision/current-pointer invariants across crash |
| storage rollback detection | detect/reduce silent reversion to older current state |
| authority time source/rollback behavior | ensure expiry cannot be reversed by stale clock |
| integrity descriptor mechanism | detect meaningful corruption/tampering |
| export confidentiality mechanism | protect `SENSITIVE` aggregate backup material |
| restore freshness/revalidation implementation | prevent pre-revocation backup from restoring active permission |
| migration execution/rollback mechanism | preserve protected semantics and prior governing state |
| audit/evidence durability | preserve required proof without making logs domain truth |
| telemetry redaction/correlation | propagate identifiers without sensitive payloads |

No candidate or winner is selected by this table.

---

## 11. Residual risks accepted only as R3-known constraints

R3 does not claim to solve:

- total compromise of the host/administrator trust root;
- physical theft/destruction of every copy of local state;
- confidentiality of backup media before R4 selects protection;
- cryptographic authenticity before R4 selects integrity/key mechanisms;
- trustworthy wall-clock semantics before R4 selects a strategy;
- accidental operator reauthorization of unsafe old authority after restore;
- future network/provider/device threats outside M0.

These limitations MUST remain visible in R4/R5 and later evidence. They are not reasons to weaken M0 fail-closed semantics.

---

## 12. Threat-model gate conclusion

The R3 threat model is complete enough to support architecture decision work because:

- protected assets/classes are explicit;
- trust boundaries are explicit;
- the six minimum R2 threat classes are covered;
- additional M0-specific rollback/time/atomicity risks are covered;
- required semantic mitigations are allocated;
- mechanism-dependent residual risks are named for R4;
- no security mechanism or vendor is selected.

This conclusion is subject to the independent R3 adversarial review and does not itself constitute the final R3 verdict.
