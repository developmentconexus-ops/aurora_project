---
id: ADR-AURORA-0007
title: SQLite as the M0 Operational State Store
document_type: adr
form: explanation
authority: decision
status: proposed
version: 0.1.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - proposed M0 operational-state database class and durability posture
related:
  - DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC
  - RESEARCH-AURORA-M0-RUNTIME-PERSISTENCE-R4-V1
  - SPK-AURORA-M0-SOVEREIGN-STORE-001
  - ADR-AURORA-0003
  - ADR-AURORA-0004
supersedes: []
superseded_by: null
last_reviewed: 2026-08-07
---

# ADR-0007 — SQLite as the M0 Operational State Store

## Context

M0 needs process-independent transactional state, low write concurrency, local sovereignty and safe backup/recovery. It does not currently need a network database or many concurrent writers.

## Decision drivers

- transaction/crash consistency;
- one local Core/single user;
- zero/minimal administration;
- backup/restore;
- Windows/Linux reproducibility;
- migration path to PostgreSQL or another relational store;
- driver/toolchain burden;
- current state/audit transaction boundary.

Affected requirements include `REQ-024..028`, `046..066`, `067`, `095..100`, `104`, `107`.

## Evidence

`RESEARCH-AURORA-M0-RUNTIME-PERSISTENCE-R4-V1` compares SQLite, PostgreSQL and RocksDB using current primary sources.

SQLite is the best scope match. PostgreSQL is a robust future/fallback class but adds a server/admin topology not required by M0. RocksDB operates at a lower abstraction level and would push more query/index/migration semantics into Aurora.

The accepted Research Map explicitly requires a crash/restart/restore spike before Sovereign Core storage commitment.

## Options

- SQLite embedded relational store;
- PostgreSQL 18 client/server store;
- RocksDB/embedded KV;
- no durable DB / files only.

## Decision

**Proposed candidate, NOT YET ACCEPTABLE:** select **SQLite** as the M0 operational state store **only if** `SPK-AURORA-M0-SOVEREIGN-STORE-001` passes and identifies an acceptable exact Go binding/configuration.

Candidate durability posture to test:

```text
journal_mode = WAL
synchronous = FULL
accepted state/current pointer/audit/evidence refs committed in one DB transaction
supported SQLite backup/snapshot API for operational backup
```

Exact binding remains unresolved between the spike's current candidates:

```text
modernc.org/sqlite
mattn/go-sqlite3
```

This ADR MUST NOT move to `accepted` while that implementation-blocking choice and crash evidence remain unresolved.

## Explicit fallback

If both SQLite candidates fail a gate criterion, expand executable evaluation to PostgreSQL rather than weakening durability, backup or cross-platform requirements.

## Consequences

### Positive

- no DB server required for M0;
- strong fit for low-write local application state;
- relational constraints/transactions fit revision/authority/audit model;
- portable logical export keeps future store migration viable.

### Negative

- one writer at a time;
- network-shared/multi-process writer growth would trigger reconsideration;
- exact binding introduces either pinning complexity or CGO/toolchain complexity.

### Risks

Documentation cannot prove the actual Go binding + filesystem + WAL configuration behaves correctly under Aurora's crash boundary. That is why acceptance is spike-blocked.

## Compatibility / migration / rollback

All SQL/store access stays behind the Durable State Port. Logical exports and migrations are database-independent. PostgreSQL is the documented migration class when multi-process/network/concurrency needs become current.

## Validation

Required before acceptance:

```text
SPK-AURORA-M0-SOVEREIGN-STORE-001
→ EVIDENCE_COMPLETE
→ REVIEWED
→ decision identifies exact binding/configuration
```

## Reconsideration triggers

- concurrent independent writer processes become current scope;
- network-shared DB becomes required;
- SQLite driver/support degrades;
- spike fails crash/recovery/backup criteria;
- M4 durable-execution architecture materially changes persistence needs.

## References

- `RESEARCH-AURORA-M0-RUNTIME-PERSISTENCE-R4-V1`
- SQLite official documentation cited by the research manifest
- `SPK-AURORA-M0-SOVEREIGN-STORE-001`
