---
id: REVIEW-AURORA-M0-R4-ARCHITECTURE-DECISION-READINESS-2026-08-07
title: M0 ACRM R4 Architecture/Decision Readiness Review
document_type: adversarial_review
form: explanation
authority: evidence
status: current
version: 1.0.1
owners:
  - developmentconexus-ops
source_of_truth_for:
  - M0 R4 documentary architecture review observations and verdict at the pre-spike checkpoint
related:
  - DOC-AURORA-STATUS
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DOC-AURORA-M0-R4-OPERATOR-AUTHORIZATION
  - DESIGN-AURORA-M0-R4-DECISION-LANDSCAPE
  - DOC-AURORA-CAP-SOVEREIGN-CORE-R4-DECISION-COVERAGE
  - DOC-AURORA-ADR-INDEX
  - DOC-AURORA-RESEARCH-MAP
  - DESIGN-AURORA-M0-SOVEREIGN-STORE-SPIKE-001
  - DESIGN-AURORA-M0-OWNER-TRUST-SPIKE-002
  - REVIEW-AURORA-M0-R4-ARCHITECTURE-DECISION-READINESS-RERUN-2026-08-07
source_revision: d00cc1abfc2a41ac7e81e1f3478e188b3c5e9e52
r4_documentary_package_revision: ef7cdcc31bf09a0c91ea88eee5ef6f501c9eeff7
reviewed_at: 2026-08-07
last_reviewed: 2026-08-07
---

# M0 ACRM R4 — Architecture/Decision Readiness Review

## 1. Executive verdict

```text
R4 BLOCKED
```

This is **not** an R4 failure of the architecture direction.

The documentary R4 work is sufficiently complete to expose the remaining material uncertainty precisely:

- all 15 R3 architecture questions are accounted for;
- four focused current research reports and source manifests exist;
- six ADRs are proposed, none silently accepted;
- two exact Architecture Spikes are specified;
- four ADRs are documentary-ready for operator review;
- two ADRs remain explicitly spike-blocked;
- no spike execution or implementation is authorized.

R4 cannot PASS because the accepted research governance explicitly requires crash/restart/restore executable evidence before Sovereign Core storage/recovery commitment, and that evidence does not yet exist. The material ADRs also remain `proposed` pending operator review.

R5 is not authorized and MUST NOT begin.

---

## 2. Gate used

The accepted Capability Realization Method asks at R4:

> Have material technical uncertainties and choices been researched, proven and decided enough for the scoped Mission?

Gate conditions require:

- no unresolved material architecture choice in current scope;
- accepted ADRs where needed;
- decisions compatible with Blueprint;
- migration/rollback considered;
- required spike evidence reviewed.

At this checkpoint, research and proposed decisions are mature enough to narrow the uncertainty, but required spike evidence is intentionally absent because Architecture Spike execution has not been authorized.

---

## 3. Fixed baseline and package reviewed

R4 source baseline:

```text
d00cc1abfc2a41ac7e81e1f3478e188b3c5e9e52
```

Clean documentary package reviewed:

```text
ef7cdcc31bf09a0c91ea88eee5ef6f501c9eeff7
```

The package includes:

```text
docs/acceptance/2026-08-07-m0-r4-operator-authorization.md
docs/design/M0-R4-DECISION-LANDSCAPE.md

docs/research/AURORA-RESEARCH-M0-RUNTIME-PERSISTENCE-R4-v1.md
...sources.json
docs/research/AURORA-RESEARCH-M0-PORTABILITY-INTEGRITY-R4-v1.md
...sources.json
docs/research/AURORA-RESEARCH-M0-OWNER-AUTHORITY-RECOVERY-R4-v1.md
...sources.json
docs/research/AURORA-RESEARCH-M0-OBSERVABILITY-R4-v1.md
...sources.json

docs/adr/0003-m0-go-core-runtime.md
docs/adr/0004-m0-local-state-execution-shape.md
docs/adr/0005-m0-portable-state-export.md
docs/adr/0006-m0-observability-boundary.md
docs/adr/0007-m0-sqlite-operational-store.md
docs/adr/0008-m0-owner-root-recovery-trust.md

docs/design/SPK-AURORA-M0-SOVEREIGN-STORE-001.md
docs/design/SPK-AURORA-M0-OWNER-TRUST-002.md

docs/capabilities/CAP-SOVEREIGN-CORE/R4-DECISION-COVERAGE.md
```

---

## 4. Decision philosophy review

The operator approved:

```text
long-horizon exploration
+ evidence-bounded commitment
```

The package follows that rule rather than optimizing only M0:

- H0 evaluates the M0 Golden Proof;
- H1 evaluates foreseeable M1–M4 pressure;
- H2 checks constitutional long-term direction;
- decisions are classified by reversibility/lock-in;
- exit/migration paths are recorded;
- mechanisms are not introduced merely because they may be useful later.

No current proposal requires M1 memory architecture, M2 Registry/AHDK, M3 Delegation or M4 durable execution to exist in M0.

---

## 5. Fifteen-question coverage

`R4-DECISION-COVERAGE.md` contains exactly one disposition for every R3 open question:

```text
R4-Q-CORE-001
R4-Q-STORE-001
R4-Q-STATE-001
R4-Q-SCHEMA-001
R4-Q-ATOMIC-001
R4-Q-INTEGRITY-001
R4-Q-TIME-001
R4-Q-AUTHN-001
R4-Q-EXPORT-001
R4-Q-MIGRATE-001
R4-Q-AUDIT-001
R4-Q-TELEM-001
R4-Q-TOPOLOGY-001
R4-Q-ENGINE-001
R4-Q-RESTORE-001
```

No choice has been silently delegated to an eventual implementer.

The current dispositions divide cleanly into:

### Documentary-ready proposals

- Core runtime → ADR-0003;
- state/event/topology/durable-engine non-selection → ADR-0004;
- portable schema/export/migration → ADR-0005;
- observability boundary → ADR-0006.

### Executable-evidence-dependent proposals

- operational store and crash-consistent transaction boundary → ADR-0007 / SPK-001;
- owner root, authenticated rollback/time anchor and restore freshness → ADR-0008 / SPK-002.

### Physical mechanisms partially dependent on SPK-001/002

- operational backup;
- physical audit co-location/transaction behavior;
- migration fixture behavior;
- authenticated integrity/trust-anchor behavior.

This is an acceptable R4 BLOCKED state because the missing evidence has exact owners and proof procedures.

---

## 6. Runtime and persistence review

### Go proposal

ADR-0003 chooses Go for the initial Core while preserving language-neutral Aurora boundaries.

The decision is not based only on simplicity. Current M0 risk is dominated by state/authority/recovery correctness; the proposed Core does not need hard-real-time or one universal implementation language. The research explicitly keeps Rust available for future lower-level Capabilities.

No material language uncertainty remains that requires a standalone Go-vs-Rust prototype before operator review.

### SQLite proposal

ADR-0007 intentionally does **not** claim SQLite is already decided.

SQLite is the strongest scope-fit candidate because M0 is one local Core with low writer concurrency and no network database requirement. PostgreSQL is preserved as the fallback/migration class; low-level KV stores were rejected for the current scope because they transfer more invariant/query/migration ownership into Aurora without a named benefit.

The review agrees that the real remaining uncertainty is not “can SQLite store rows?” but the exact:

```text
Go binding
+ SQLite configuration
+ filesystem/runtime behavior
+ crash boundary
+ online backup/restore behavior
```

Therefore ADR-0007 correctly remains spike-blocked.

---

## 7. State/event/topology/durable-engine review

ADR-0004 is internally consistent with R3:

```text
one local modular Core process
+ explicit current governing state
+ immutable accepted state revisions
+ logically distinct audit/domain events
```

Full event sourcing is not selected because it would create a new permanent replay/event-migration authority problem not required by M0.

A durable workflow engine is intentionally `NOT_YET_A_DECISION` for M0. This is not accidental deferral: M0 contains no long-lived Mission waits/timers/provider/external-effect lifecycle that currently justifies a second durable history system. The reconsideration trigger at M4/current requirement is explicit.

The review finds no current local-maximum problem in this non-selection because the R3 ports preserve later introduction without rewriting Project/Authority domain meaning.

---

## 8. Portability/export/migration review

ADR-0005 separates:

```text
physical DB schema
portable Aurora logical schema
canonical bytes for digest/MAC
outer encrypted export package
```

This materially reduces store lock-in.

The proposed choices—JSON Schema 2020-12, JSON logical interchange, RFC 8785 JCS for deterministic hash/MAC input, SHA-256 content digests and age for normal SENSITIVE outer export encryption—are reversible mechanisms rather than product semantics.

The review agrees that operational database backup MUST remain distinct from the portable sovereignty export.

Application-owned version-pair migration also prevents SQLite/PostgreSQL physical migrations from silently changing identity/authority meaning.

No extra architecture spike is required to decide this boundary, although SPK-001/002 will provide end-to-end evidence for operational backup and recovery usage.

---

## 9. Owner-root/time/restore review

ADR-0008 proposes a random Owner Root Key wrapped by an Argon2id-passphrase-derived key, purpose-separated subkeys, authenticated governing descriptors and a small Owner Trust Store outside the operational database.

The architecture solves an important conceptual error: a MAC stored with its key in the same database would not establish an independent trust boundary, and a valid old backup MAC would prove historical authenticity but not current authority freshness.

The proposed states:

```text
DB generation == anchor → verify normally
DB generation < anchor  → rollback suspected / fail closed
DB generation > anchor  → possible post-commit anchor lag / authenticated reconciliation
meaningful backward time → TIME_UNTRUSTED
historical restore       → REVALIDATION_REQUIRED
```

are coherent with R3.

However, the cross-file DB/trust-anchor crash protocol is exactly the sort of mechanism that cannot be accepted from documentary reasoning alone. ADR-0008 correctly remains blocked by SPK-002.

The review also accepts the explicitly stated residual-risk boundary: purely local M0 does not claim resistance to an attacker capable of replaying every local trust file under total owner/admin/root compromise.

---

## 10. Observability review

ADR-0006 keeps observability replaceable:

```text
OpenTelemetry traces/metrics
Go slog structured logs
optional OTLP/export backend
```

The telemetry backend is not required for canonical state, audit or evidence correctness. This satisfies the R3 ownership boundary and avoids making a beta/remote observability pipeline foundational.

No R4 Architecture Spike is required for this choice; R7 verification can prove exporter-failure independence and redaction.

---

## 11. Spike decomposition review

### SPK-AURORA-M0-SOVEREIGN-STORE-001

Scope is appropriately narrow:

- Go + SQLite only;
- modernc versus mattn binding;
- Linux/Windows controlled matrix;
- kill points around one state transaction;
- WAL/restart;
- supported backup/restore;
- corruption/incompatibility;
- migration fixture;
- operational/build burden.

PostgreSQL is not implemented unless both SQLite candidates fail or evidence exposes a current requirement for the client/server class.

This avoids a broad benchmark contest while preserving a falsification path.

### SPK-AURORA-M0-OWNER-TRUST-002

Correctly depends on the reviewed SPK-001 result.

It tests only the owner-root/trust protocol on the viable store:

- unlock/rotation;
- direct DB mutation;
- DB rollback;
- DB/anchor crash gap;
- backward time;
- historical restore;
- self-revalidation denial;
- owner revalidation;
- fresh-machine recovery.

Running SPK-001 and SPK-002 concurrently would confound storage and trust failures. Sequential execution is therefore required.

Both spike specs declare `DISCARD`; spike code cannot become production by success alone.

---

## 12. Documentary package validation

The dedicated R4 validator initially caught a validator/index identifier mismatch; the validator was corrected to distinguish internal ADR document IDs from public ADR numbers.

The repository validator then caught a real metadata issue: conceptual `SPK-*` identifiers are not valid canonical document IDs under Aurora's document-ID grammar. The package was corrected to preserve:

```text
document id: DESIGN-AURORA-...-SPIKE-...
spike_id:    SPK-AURORA-...
```

and all `related:` references use canonical document IDs.

After those corrections, GitHub Actions run `31200153496` completed successfully through:

- R4 documentary-package validation;
- generated documentation projection;
- repository documentation validation;
- generated-projection freshness check;
- helper cleanup and commit.

The clean package revision is `ef7cdcc31bf09a0c91ea88eee5ef6f501c9eeff7`.

---

## 13. Findings

### R4-F01 — validator compared internal ADR IDs to public index numbers

Status: `RESOLVED`.

No product artifact was wrong. The one-shot validator expected `ADR-AURORA-0003` while the ADR index correctly renders `ADR-0003`.

### R4-F02 — conceptual spike IDs were used as canonical document IDs

Status: `RESOLVED`.

Aurora's document-ID grammar permits `DESIGN-*` but not `SPK-*`. The solution preserves both identities explicitly and resolves all frontmatter relations.

### R4-F03 — required storage/recovery executable evidence does not exist

Status: `OPEN / GATE BLOCKER`.

Owner: R4 / `SPK-AURORA-M0-SOVEREIGN-STORE-001`.

Architecture Spike execution is currently not authorized.

### R4-F04 — owner-root/trust protocol lacks executable crash/rollback/time evidence

Status: `OPEN / GATE BLOCKER`.

Owner: R4 / `SPK-AURORA-M0-OWNER-TRUST-002`, sequenced after reviewed SPK-001.

### R4-F05 — material proposed ADRs are not accepted

Status: `OPEN / GATE BLOCKER`.

ADR-0003 through ADR-0006 are ready for operator review based on documentary evidence.

ADR-0007 and ADR-0008 MUST remain proposed until their required spike evidence is complete/reviewed.

---

## 14. Gate checklist

| R4 condition | Result |
|---|---|
| all current material questions identified | PASS — 15/15 |
| focused current research exists | PASS |
| alternatives and operational burden considered | PASS |
| migration/exit path considered | PASS |
| Blueprint/R2/R3 compatibility | PASS at documentary level |
| no hidden implementation authorization | PASS |
| required ADRs accepted | **BLOCKED** |
| required crash/restart/restore evidence reviewed | **BLOCKED** |
| owner-root/trust executable evidence reviewed | **BLOCKED** |
| no unresolved material architecture choice | **BLOCKED pending exact spike outcomes** |

Therefore:

```text
R4 BLOCKED
```

is the only defensible current verdict.

---

## 15. Exact next actions

The safe sequence is:

```text
1. operator reviews ADR-0003..0006
2. operator accepts/rejects/revises those documentary-ready decisions
3. operator separately authorizes execution of exact SPK-AURORA-M0-SOVEREIGN-STORE-001
4. execute/review SPK-001
5. use evidence to accept/reject/revise ADR-0007 and select exact store binding/configuration
6. then authorize SPK-AURORA-M0-OWNER-TRUST-002 against that reviewed store result
7. execute/review SPK-002
8. accept/reject/revise ADR-0008
9. re-run R4 review
10. only if all material choices are accepted/evidenced → R4 PASS
11. STOP before R5 and await separate authorization
```

No spike execution, Mission Contract, Microdesign or production implementation is authorized by this review.
