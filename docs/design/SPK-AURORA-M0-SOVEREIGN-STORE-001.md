---
id: SPK-AURORA-M0-SOVEREIGN-STORE-001
title: M0 Sovereign Store Crash, Restart and Restore Spike
document_type: architecture_spike
form: reference
authority: design
status: proposed
version: 0.1.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - proposed executable investigation of M0 operational store and Go SQLite binding
related:
  - DOC-AURORA-M0-R4-OPERATOR-AUTHORIZATION
  - DESIGN-AURORA-M0-R4-DECISION-LANDSCAPE
  - RESEARCH-AURORA-M0-RUNTIME-PERSISTENCE-R4-V1
  - DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC
  - DOC-AURORA-CAP-SOVEREIGN-CORE-TEST-PLAN
source_revision: d00cc1abfc2a41ac7e81e1f3478e188b3c5e9e52
last_reviewed: 2026-08-07
---

# SPK-AURORA-M0-SOVEREIGN-STORE-001

## 1. Authorization state

```text
Specification: PROPOSED
Execution: NOT AUTHORIZED
Production promotion: PROHIBITED
```

This spike MUST NOT execute until the operator separately authorizes this exact spike/revision.

## 2. Exact uncertainty

Can a minimal **Go + SQLite** persistence stack satisfy the M0 Sovereign Core crash/restart/restore and atomicity requirements with lower operational burden than introducing a client/server database, and which current Go SQLite binding gives the better first implementation baseline?

## 3. Decisions informed

- `R4-Q-STORE-001` operational-state storage mechanism;
- `R4-Q-ATOMIC-001` crash-consistent commit mechanism;
- SQLite-specific portion of `R4-Q-AUDIT-001`;
- SQLite operational-backup portion of `R4-Q-EXPORT-001`;
- driver/build implications of `R4-Q-CORE-001`.

It does **not** decide logical export schema, owner-root security or R5 implementation scope.

## 4. Alternatives

### Candidate A

```text
Go 1.26.5
+ database/sql
+ modernc.org/sqlite current pinned compatible release set
+ SQLite WAL
+ synchronous=FULL
```

### Candidate B

```text
Go 1.26.5
+ database/sql
+ mattn/go-sqlite3 current pinned release
+ SQLite WAL
+ synchronous=FULL
```

### Fallback class, not initially implemented

```text
PostgreSQL 18
```

PostgreSQL is expanded into an executable candidate only if both SQLite candidates fail a gate criterion or if the spike exposes a current requirement that inherently needs client/server or multi-writer semantics.

## 5. Controlled environment

Minimum reproducible matrix:

```text
ubuntu-24.04 amd64 GitHub Actions
windows-latest amd64 GitHub Actions
```

Exact runner images, Go patch release, module versions, SQLite engine version, compiler/C toolchain for CGO, filesystem, PRAGMAs and environment variables MUST be captured in evidence.

A later operator-machine reproduction may supplement the result but is not a substitute for the controlled CI matrix.

No production credentials, network databases or user data are permitted.

## 6. Minimal prototype

Disposable spike code models only:

- Aurora identity;
- one Project identity;
- immutable state revisions;
- current revision pointer;
- minimal authority revision reference;
- transition attempt/audit row;
- evidence reference;
- export/backup attempt metadata.

No LLM, Harness, AHDK, UI framework, network service, memory subsystem or durable workflow engine.

## 7. Required transaction boundary

One accepted transition candidate must attempt in one DB transaction:

```text
read/validate expected current revision
→ insert immutable state revision N+1
→ update current revision pointer
→ record transition accepted result
→ record required audit metadata
→ record evidence reference
→ commit
```

Rejected/stale transitions must leave governing state unchanged.

## 8. Fault injection procedure

For both candidates execute deterministic injection points:

1. before transaction begin;
2. after revision validation;
3. after new revision insert but before pointer update;
4. after pointer update but before audit/evidence writes;
5. immediately before `Commit`;
6. immediately after `Commit` returns success;
7. during/around WAL checkpoint;
8. during live backup operation where safely reproducible.

Kill the process abruptly at each point, start a fresh process and inspect the governing state.

Expected invariant after recovery:

```text
exactly old complete state
OR
exactly new complete committed state
NEVER a partial governing combination
```

A test may use explicit process-kill fault injection; it must not claim to emulate all physical power-loss/storage-controller failures.

## 9. Recovery cases

Required cases:

- fresh restart after clean commit;
- repeated restart;
- stale transition rejection after restart;
- WAL present at restart;
- interrupted checkpoint;
- invalid/corrupt/truncated test database;
- supported consistent online backup while DB is open;
- restore backup into a fresh directory/process;
- identity collision on restore fixture;
- unsupported logical schema version fixture;
- one explicit migration fixture preserving identity/current revision semantics.

Use SQLite-supported backup/snapshot APIs; intentionally demonstrate that naïve live file copying is not accepted as Aurora backup behavior.

## 10. Golden Proof

```text
initialize disposable Aurora
→ create Project
→ commit accepted state + audit/evidence refs
→ kill all spike processes at controlled points
→ fresh process recovers one coherent governing revision
→ stale/invalid transition is rejected with zero governing mutation
→ create supported online backup
→ destroy original working directory
→ restore in fresh environment
→ recover same Aurora/Project identities and state
→ run migration fixture
→ invariants remain true
```

Both candidates must run the same black-box scenario definitions.

## 11. Measurements

Collect for each candidate:

- pass/fail by fault case;
- binary size;
- build time on Linux/Windows;
- external build-tool requirements;
- dependency/module count and notable native dependencies;
- startup/open time;
- transaction latency distribution for representative tiny M0 commits;
- backup/restore duration for small/medium deterministic fixtures;
- database/WAL file behavior;
- cross-platform differences;
- debugging clarity;
- upgrade/pinning burden;
- operational steps required for install/run/backup/restore.

Performance is secondary to correctness and operational simplicity. No microbenchmark alone can select the winner.

## 12. Gate criteria

A candidate remains viable only if:

- all required crash/restart cases preserve coherent governing state;
- invalid transition never mutates governing state;
- supported backup/restore reproduces identities/state;
- corruption/incompatibility fails explicitly rather than fabricating state;
- normal M0 operation requires no separate DB server;
- build/reproduction succeeds on the required environment matrix;
- dependency/pinning burden is documented and judged acceptable;
- the implementation keeps SQL/store semantics behind the Durable State Port.

## 13. Decision rule

If both pass correctness:

1. prefer lower operational/build burden;
2. prefer fewer fragile dependencies/toolchains;
3. prefer clearer cross-platform reproduction;
4. do not trade durability correctness for binary size or benchmark speed.

If modernc passes but its exact `libc` pinning proves operationally fragile, mattn may win despite CGO.

If mattn's CGO cross-platform burden is material and modernc remains reproducible, modernc may win.

If both fail a correctness criterion, broaden to PostgreSQL rather than weakening the criterion.

## 14. Evidence artifacts

Required:

- exact source commit;
- dependency lock/module files;
- workflow/commands;
- stdout/stderr logs;
- machine-readable case matrix;
- recovered-state dumps in logical form;
- DB integrity-check output where applicable;
- backup/restore hashes;
- binary/build metrics;
- limitations statement;
- comparison report;
- reviewer verdict.

## 15. Stop conditions

Stop and return `SPIKE_BLOCKED/FAILED` if:

- required target environment cannot reproduce a candidate;
- a candidate needs undocumented behavior to pass;
- crash behavior is nondeterministically ambiguous after repeated controlled runs;
- backup requires unsafe raw copy semantics;
- spike scope begins implementing full Aurora Core rather than the disposable state fixture.

## 16. Disposal rule

```text
DISCARD
```

Spike code is disposable evidence. Any later production implementation must be created through R5/R6/R7 and may reuse knowledge, not automatically promote spike code.

## 17. Decision implication

A passing reviewed result may support a proposed/accepted ADR selecting:

```text
Go + SQLite + exact binding/configuration
```

for M0 operational state. It cannot by itself authorize implementation.
