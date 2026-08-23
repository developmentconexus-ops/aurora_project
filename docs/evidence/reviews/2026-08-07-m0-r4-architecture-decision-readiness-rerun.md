---
id: REVIEW-AURORA-M0-R4-ARCHITECTURE-DECISION-READINESS-RERUN-2026-08-07
title: M0 ACRM R4 Architecture/Decision Readiness Rerun
document_type: adversarial_review
form: explanation
authority: evidence
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - final M0 R4 Architecture/Decision Readiness verdict after ADR-0008 operator acceptance
related:
  - DOC-AURORA-STATUS
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DOC-AURORA-M0-R4-OPERATOR-AUTHORIZATION
  - REVIEW-AURORA-M0-R4-ARCHITECTURE-DECISION-READINESS-2026-08-07
  - DOC-AURORA-CAP-SOVEREIGN-CORE-R4-DECISION-COVERAGE
  - DOC-AURORA-M0-R4-SPK001-EVIDENCE-RECEIPT
  - REVIEW-AURORA-M0-R4-SPK001-SOVEREIGN-STORE-2026-08-07
  - DOC-AURORA-M0-R4-SPK002-EVIDENCE-RECEIPT
  - REVIEW-AURORA-M0-R4-SPK002-OWNER-TRUST-2026-08-07
  - DOC-AURORA-M0-R4-ADR0008-ACCEPTANCE
  - ADR-AURORA-0003
  - ADR-AURORA-0004
  - ADR-AURORA-0005
  - ADR-AURORA-0006
  - ADR-AURORA-0007
  - ADR-AURORA-0008
  - ADR-AURORA-0009
pre_acceptance_canonical_revision: 35ce98fb2ddde16133c01a4da2f3545c8ae5e308
accepted_adr0008_blob: 2a1497f8311ba9d04cd61f5025d7eae2af2fc57f
reviewed_at: 2026-08-07
last_reviewed: 2026-08-07
---

# M0 ACRM R4 — Architecture/Decision Readiness Rerun

## 1. Executive verdict

```text
R4 PASS
```

The final blocker identified by the prior R4 review has been removed: the operator explicitly accepted ADR-0008 revision 0.2.0 after reviewed SPK-002 evidence.

The M0 architecture decision set is now sufficiently researched, evidenced and accepted for the scoped Mission-contract stage. This verdict does **not** authorize R5 or implementation.

```text
R4 PASS
→ STOP
→ R5 NOT AUTHORIZED
→ await separate operator authorization
```

## 2. Gate used

The accepted Capability Realization Method defines R4 as:

> Have material technical uncertainties and choices been researched, proven and decided enough for the scoped Mission?

The gate requires:

- current scope has no unresolved material architecture choice;
- accepted ADRs exist where needed;
- decisions are compatible with the Blueprint;
- migration/rollback has been considered;
- required spike evidence has been reviewed.

This rerun evaluates exactly those conditions after all required R4 evidence and operator decisions.

## 3. Inputs reviewed

The rerun uses the accumulated R4 package, including:

- 15-question R4 Decision Landscape and current decision coverage;
- focused runtime/persistence, portability/integrity, owner-authority/recovery and observability research;
- Mastra cross-horizon materiality analysis;
- accepted ADR-0003 through ADR-0007;
- accepted cross-horizon ADR-0009;
- SPK-001 `PASS / REVIEWED / CLOSED` evidence;
- SPK-002 `PASS / REVIEWED / CLOSED` evidence;
- ADR-0008 v0.2.0 plus explicit operator acceptance `DOC-AURORA-M0-R4-ADR0008-ACCEPTANCE`.

## 4. Resolution of prior R4 blockers

The original R4 review identified three gate-blocking classes after its documentary phase.

### R4-F03 — storage/recovery executable evidence absent

```text
prior status: OPEN / GATE BLOCKER
current status: RESOLVED
```

`SPK-AURORA-M0-SOVEREIGN-STORE-001` completed with reviewed cross-platform PASS evidence. The resulting SQLite + `database/sql` + `modernc.org/sqlite` decision is accepted in ADR-0007.

### R4-F04 — owner-root/trust executable evidence absent

```text
prior status: OPEN / GATE BLOCKER
current status: RESOLVED
```

`SPK-AURORA-M0-OWNER-TRUST-002` completed with hardened Linux/Windows PASS evidence covering unlock/rotation, DB integrity, rollback, DB/anchor crash gap, backward time, historical restore, self-revalidation denial, owner revalidation, fresh-machine recovery and secret hygiene.

The adversarial review additionally forced normal governing mutations to reject anomalous trust states before writing. That material finding was resolved and the full matrix passed again.

### R4-F05 — required ADRs not accepted

```text
prior status: OPEN / GATE BLOCKER
current status: RESOLVED
```

The current accepted architecture decisions are:

```text
ADR-0003 — Go as initial Sovereign Core runtime
ADR-0004 — local Core/state/execution shape
ADR-0005 — portable logical state/export/migration boundary
ADR-0006 — observability boundary
ADR-0007 — SQLite operational-state store + modernc binding baseline
ADR-0008 — Owner Root and recovery trust boundary
ADR-0009 — Mastra preferred-first cross-horizon Harness substrate
```

ADR-0008 acceptance is bound specifically to v0.2.0 / blob `2a1497f8311ba9d04cd61f5025d7eae2af2fc57f`.

## 5. Fifteen-question closure

All 15 implementation-blocking M0 R4 questions now have a governing disposition:

```text
R4-Q-CORE-001      decided / accepted
R4-Q-STORE-001     decided / accepted + evidenced
R4-Q-STATE-001     decided / accepted + evidenced
R4-Q-SCHEMA-001    decided / accepted
R4-Q-ATOMIC-001    decided / accepted + evidenced
R4-Q-INTEGRITY-001 decided / accepted + evidenced
R4-Q-TIME-001      decided / accepted + evidenced
R4-Q-AUTHN-001     decided / accepted + evidenced
R4-Q-EXPORT-001    decided / accepted + evidenced
R4-Q-MIGRATE-001   decided / accepted + evidenced
R4-Q-AUDIT-001     decided / accepted + evidenced
R4-Q-TELEM-001     decided / accepted
R4-Q-TOPOLOGY-001  decided / accepted + evidenced
R4-Q-ENGINE-001    decided / intentional M0 non-selection
R4-Q-RESTORE-001   decided / accepted + evidenced
```

There is no implementation-blocking M0 architecture choice silently delegated to R5/R6.

## 6. Blueprint and horizon compatibility

The accepted architecture remains compatible with the established Aurora sovereignty boundaries:

- Aurora Core owns canonical identity, project state, authority and governance;
- storage implements a replaceable durable-state boundary rather than owning domain meaning;
- current state/revisions remain primary while audit/events remain distinct;
- portable logical state is separated from the physical SQLite representation;
- Mastra remains a preferred-first agentic Harness substrate, not Sovereign Core authority;
- no durable workflow engine is introduced into M0 without a current requirement;
- future OS/hardware root protection can strengthen ORK custody without changing owner/authority semantics.

The long-horizon analysis therefore does not expose a current local-maximum blocker requiring another M0 mechanism.

## 7. Migration, rollback and recovery

R4 has explicit exit/recovery paths rather than mechanism lock-in by omission:

- logical export and schema versioning are independent of SQLite physical representation;
- explicit application-owned migrations preserve authority semantics;
- operational backup is distinct from sovereignty export;
- SQLite binding is replaceable behind `database/sql` and the Durable State Port;
- DB-only rollback and DB/anchor disagreement fail closed under ADR-0008;
- historical restore enters `REVALIDATION_REQUIRED` rather than silently restoring current permission;
- future stronger OS/TPM/hardware custody may wrap the same ORK lineage.

The gate therefore has sufficient migration/rollback analysis for M0 contract readiness.

## 8. Carry-forward constraints — not R4 blockers

The following obligations remain material but belong to later authorized gates/evidence, not to architecture choice resolution:

### R6 implementation-design constraints

- production root-envelope decoding must allowlist/bound Argon2id parameters before allocating KDF resources;
- exact filesystem publish/directory-sync behavior must be designed explicitly for the supported production targets;
- production APIs must preserve the SPK-002 rule that anomaly classification is enforced at the mutation boundary;
- exact module/toolchain pins must be revalidated at implementation baseline.

### R7 evidence constraints

- Golden Proof must not overclaim literal hardware/power-loss durability beyond the tested fault model;
- target-environment crash/restart/restore behavior must be verified against the implemented filesystem/storage wrapper;
- observability/export failure must remain non-authoritative and redaction-safe;
- the local-trust residual risk must remain explicit: total replay plus owner-secret/runtime compromise is outside the M0 guarantee.

These are verification/hardening obligations with an accepted architecture to implement, not unresolved architecture alternatives.

## 9. R4 gate checklist

| R4 condition | Final result |
|---|---|
| all current material questions identified | PASS — 15/15 |
| focused current research exists | PASS |
| alternatives and operational burden considered | PASS |
| migration/exit path considered | PASS |
| Blueprint/R2/R3 compatibility | PASS |
| required ADRs accepted | PASS |
| required storage crash/restart/restore evidence reviewed | PASS — SPK-001 |
| required owner-root/trust evidence reviewed | PASS — SPK-002 |
| no unresolved material architecture choice | PASS |
| no hidden implementation authorization | PASS |

Therefore:

```text
M0 ACRM R4 — PASS
```

## 10. Authorization boundary after PASS

R4 PASS means the architecture is ready to support a scoped Mission Contract. It does not create that contract and does not authorize its gate.

Still prohibited until separately authorized:

```text
R5 Contract Readiness work
Mission Contract creation
R6 Microdesign / Implementation Plan
production code
promotion of spike code
Aurora Core runtime implementation
AHDK implementation
MNFS integration
additional spikes not separately authorized
```

## 11. Exact next action

```text
R4 PASS
→ STOP
→ await explicit operator authorization for M0 ACRM R5 — Contract Readiness
```
