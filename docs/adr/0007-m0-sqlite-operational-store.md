---
id: ADR-AURORA-0007
title: SQLite as the M0 Operational State Store
document_type: adr
form: explanation
authority: decision
status: accepted
accepted_at: 2026-08-07
acceptance_evidence: DOC-AURORA-M0-R4-ADR0007-ACCEPTANCE-SPK002-AUTHORIZATION
version: 0.2.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - M0 operational-state database class, Go binding and durability posture
related:
  - DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC
  - RESEARCH-AURORA-M0-RUNTIME-PERSISTENCE-R4-V1
  - DESIGN-AURORA-M0-SOVEREIGN-STORE-SPIKE-001
  - DOC-AURORA-M0-R4-SPK001-EVIDENCE-RECEIPT
  - REVIEW-AURORA-M0-R4-SPK001-SOVEREIGN-STORE-2026-08-07
  - ADR-AURORA-0003
  - ADR-AURORA-0004
supersedes: []
superseded_by: null
last_reviewed: 2026-08-07
---

# ADR-0007 — SQLite as the M0 Operational State Store

## Context

M0 needs process-independent transactional state, low write concurrency, local sovereignty and safe backup/recovery. It does not currently need a network database or many concurrent writers.

The accepted Research Map required a crash/restart/restore Architecture Spike before committing to the Sovereign Core operational store. `SPK-AURORA-M0-SOVEREIGN-STORE-001` has now completed and been reviewed against the authorized fixed specification.

## Decision drivers

- transaction/crash consistency;
- one local Core/single user;
- zero/minimal administration;
- backup/restore;
- Windows/Linux reproducibility;
- migration path to PostgreSQL or another relational store;
- driver/toolchain burden;
- current state/audit transaction boundary;
- reproducible dependency/version management.

Affected requirements include `REQ-024..028`, `046..066`, `067`, `095..100`, `104`, `107`.

## Evidence

`RESEARCH-AURORA-M0-RUNTIME-PERSISTENCE-R4-V1` compared SQLite, PostgreSQL and RocksDB using current primary sources.

The executable evidence is owned by:

```text
DOC-AURORA-M0-R4-SPK001-EVIDENCE-RECEIPT
REVIEW-AURORA-M0-R4-SPK001-SOVEREIGN-STORE-2026-08-07
```

Final SPK-001 workflow:

```text
run: 31213792366
revision: 4242342486f512320f12e0b603f052166264c4ea
matrix: Ubuntu/Windows × modernc/mattn
result: 4/4 correctness PASS + complete evidence receipts
```

Both SQLite bindings passed the same restart, stale/invalid transition, process-kill, WAL/checkpoint, online backup, interrupted backup, restore, collision, corruption, compatibility and migration cases.

The correctness tie activates the authorized operational-burden tie-break rule.

## Options

### A — SQLite + `modernc.org/sqlite`

Observed evidence baseline:

```text
modernc.org/sqlite v1.54.0
CGO_ENABLED=0
SQLite runtime 3.53.3 in the final matrix
```

Pros:

- passed all tested correctness/fault gates on Ubuntu and Windows;
- removes C compiler/CGO as an operational build dependency;
- faster mean build time in the final matrix;
- same `database/sql` boundary used by the alternate candidate;
- strong fit for one local low-writer-concurrency Core.

Cons:

- larger Go dependency graph in the fixture;
- official package documentation identifies a fragile exact-version relationship with `modernc.org/libc`;
- binding/runtime upgrades require deliberate compatible pinning and regression evidence.

### B — SQLite + `mattn/go-sqlite3`

Observed evidence baseline:

```text
github.com/mattn/go-sqlite3 v1.14.49
CGO_ENABLED=1
SQLite runtime 3.53.4 in the final matrix
```

Pros:

- passed all tested correctness/fault gates on Ubuntu and Windows;
- mature direct SQLite C binding;
- fewer Go dependency packages in the fixture;
- slightly lower mean 4 MiB backup/restore time in the final run.

Cons:

- requires CGO and a functioning native C toolchain;
- increases cross-platform build and distribution surface for the Sovereign Core;
- benchmark advantages are not material to any current M0 SLO.

### C — PostgreSQL 18

Remains the documented client/server fallback/migration class when independent concurrent writers, network-shared access or another current requirement makes that topology necessary.

The SPK-001 fallback trigger did not fire because both SQLite candidates passed all correctness gates and no current client/server requirement emerged.

### D — RocksDB / embedded KV

Still rejected for current scope because it shifts more relational constraint/query/migration responsibility into Aurora without a demonstrated M0 benefit.

## Decision

**Accepted by the operator on 2026-08-07:** use **SQLite** as the M0 operational-state store with **`database/sql` + `modernc.org/sqlite` as the initial Go binding baseline**.

The R4 evidence-qualified binding release is:

```text
modernc.org/sqlite v1.54.0
```

The implementation baseline must pin an exact compatible dependency/module lock rather than treating one transitive dependency set as permanently valid product meaning.

Required durability posture:

```text
journal_mode = WAL
synchronous = FULL
foreign_keys = ON
accepted state + current pointer + audit + evidence references
  committed in one database transaction
structural/version/integrity validation before governing use
supported consistent SQLite backup mechanism
no naïve live main-file copy as backup
```

The spike used `wal_autocheckpoint=0` only to make WAL recovery deterministic and observable. This ADR does **not** select that as the permanent production checkpoint policy.

Exact connection pool values, checkpoint cadence, table/index naming and production backup wrapper are R6 implementation-design concerns provided they preserve the accepted R3/R4 semantics and do not weaken durability.

This ADR is `accepted`. Acceptance selects the R4 architecture decision; it does not authorize production implementation, R5, R6 or promotion of spike code.

## Explicit fallback

If the selected binding later fails an accepted conformance/recovery suite, loses maintainable support, or its exact dependency compatibility becomes operationally unacceptable:

1. test the current supported `mattn/go-sqlite3` path against the same store contract;
2. only broaden to PostgreSQL when a current requirement or SQLite-class failure justifies client/server topology;
3. never weaken durability, backup or cross-platform requirements merely to retain the original choice.

## Consequences

### Positive

- no database server required for M0;
- no C compiler/CGO dependency in the proposed initial Sovereign Core store binding;
- relational constraints/transactions fit revision/authority/audit model;
- tested crash/restart/restore behavior on Ubuntu and Windows;
- portable logical export keeps future store migration viable;
- supported backup remains distinct from portable sovereignty export.

### Negative

- one-writer SQLite characteristics remain part of the current topology envelope;
- `modernc` brings a larger Go dependency graph than the tested `mattn` path;
- exact `modernc.org/libc` compatibility requires disciplined pinning;
- future networked/multi-writer growth can trigger store migration.

### Risks

The executable spike used GitHub-hosted filesystems and real process kill, not physical power-cut/storage-controller emulation. `synchronous=FULL` and supported backup semantics therefore remain mandatory; a future environment with materially different persistence guarantees must revalidate the failure model.

## Compatibility / migration / rollback

All SQL/store access remains behind the Durable State Port. Logical exports and migrations are database-independent. PostgreSQL remains the documented migration class when multi-process/network/concurrency needs become current.

The selected Go binding is an adapter choice, not an Aurora domain type. Replacing it must not change Aurora/Project/Authority meaning.

## Validation

SPK-001 requirement is satisfied:

```text
SPK-AURORA-M0-SOVEREIGN-STORE-001
→ EVIDENCE_COMPLETE
→ REVIEWED
→ DECISION_INFORMED
→ CLOSED
```

ADR-0007 v0.2.0 was accepted by the operator through `DOC-AURORA-M0-R4-ADR0007-ACCEPTANCE-SPK002-AUTHORIZATION`. No further M0 store Architecture Spike is required unless new material evidence contradicts SPK-001. Implementation still requires the later separately authorized ACRM gates.

## Reconsideration triggers

- concurrent independent writer processes become current scope;
- network-shared DB becomes required;
- selected SQLite binding or critical dependency loses maintainable support;
- selected binding/runtime version materially changes;
- exact modernc/libc compatible pin cannot be maintained reproducibly;
- a later fault suite contradicts SPK-001;
- M4 durable-execution architecture materially changes persistence needs.

## References

- `RESEARCH-AURORA-M0-RUNTIME-PERSISTENCE-R4-V1`
- `DOC-AURORA-M0-R4-SPK001-EVIDENCE-RECEIPT`
- `REVIEW-AURORA-M0-R4-SPK001-SOVEREIGN-STORE-2026-08-07`
- SQLite official documentation cited by the research manifest
- `SPK-AURORA-M0-SOVEREIGN-STORE-001`
- `DOC-AURORA-M0-R4-ADR0007-ACCEPTANCE-SPK002-AUTHORIZATION`
