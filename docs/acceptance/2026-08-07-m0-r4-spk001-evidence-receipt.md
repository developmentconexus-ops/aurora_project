---
id: DOC-AURORA-M0-R4-SPK001-EVIDENCE-RECEIPT
title: M0 R4 SPK-001 Sovereign Store Executable Evidence Receipt
document_type: architecture_spike_evidence
form: evidence
authority: evidence
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - executable evidence identity and measured results for SPK-AURORA-M0-SOVEREIGN-STORE-001
related:
  - DOC-AURORA-M0-R4-ADR-ACCEPTANCE-SPK001-AUTHORIZATION
  - DESIGN-AURORA-M0-SOVEREIGN-STORE-SPIKE-001
  - ADR-AURORA-0007
  - DOC-AURORA-STATUS
authorization_baseline: a3192afad3dba9c6e699588c07ca2bcaac1161fd
spike_spec_revision: 36f46956bc275d0aec32b7e3ea4d959010fa9dcb
spike_spec_blob: 6ad7397d46208a0a9c762073d2c5239ceff4e056
spike_execution_revision: 4242342486f512320f12e0b603f052166264c4ea
workflow_run: 31213792366
recorded_at: 2026-08-07
last_reviewed: 2026-08-07
---

# M0 R4 — SPK-001 Sovereign Store Executable Evidence Receipt

## 1. Lifecycle result

```text
SPK-AURORA-M0-SOVEREIGN-STORE-001
PROPOSED
→ AUTHORIZED
→ EXECUTING
→ EVIDENCE_COMPLETE
```

This receipt records executable observations. The independent review owns `REVIEWED / DECISION_INFORMED / CLOSED` disposition.

The spike code is disposable evidence and is **not** production Aurora Core code.

## 2. Fixed identity

```text
canonical authorization baseline:
  a3192afad3dba9c6e699588c07ca2bcaac1161fd

authorized spike specification:
  docs/design/SPK-AURORA-M0-SOVEREIGN-STORE-001.md
  revision: 36f46956bc275d0aec32b7e3ea4d959010fa9dcb
  blob:     6ad7397d46208a0a9c762073d2c5239ceff4e056

disposable execution branch:
  spike/m0-sovereign-store-001

final evidence revision:
  4242342486f512320f12e0b603f052166264c4ea

final workflow:
  run 31213792366 — SPK-001 Final Evidence One-Shot — SUCCESS
```

The final one-shot workflow was used because the normal final-matrix workflow entered GitHub Actions `pending` before allocating any job. The independent one-shot used the same branch revision, candidate matrix, tests, metrics runner and aggregate gate without the stalled concurrency configuration.

## 3. Candidate matrix

Pinned candidate classes:

```text
Go toolchain: 1.26.5

Candidate A — modernc
  database/sql
  modernc.org/sqlite v1.54.0
  CGO_ENABLED=0
  observed SQLite runtime: 3.53.3

Candidate B — mattn
  database/sql
  github.com/mattn/go-sqlite3 v1.14.49
  CGO_ENABLED=1
  observed SQLite runtime: 3.53.4

Both:
  SQLite WAL
  synchronous=FULL
  foreign_keys=ON
  one transactional accepted-state/current-pointer/audit/evidence mutation boundary
```

Controlled matrix:

```text
Ubuntu 24.04 amd64 × modernc
Ubuntu 24.04 amd64 × mattn
Windows latest amd64 × modernc
Windows latest amd64 × mattn
```

All four final matrix jobs completed successfully:

| Matrix case | Job | Result |
|---|---:|---|
| Ubuntu 24.04 / modernc | `92982500952` | PASS |
| Ubuntu 24.04 / mattn | `92982501031` | PASS |
| Windows / modernc | `92982500912` | PASS |
| Windows / mattn | `92982501110` | PASS |
| aggregate | `92982998192` | PASS |

Final aggregate asserted:

```text
complete: True
all correctness cases passed: True
all required evidence receipts complete: True
medium deterministic fixture: 4 MiB retained state revision
```

## 4. Correctness and fault evidence

The same black-box store semantics were exercised against both bindings.

### 4.1 Restart and identity continuity

- bootstrap one Aurora identity and one Project identity;
- close/reopen through fresh database handles repeatedly;
- preserve schema identity, Aurora ID, Project ID, current revision, authority revision and state summary;
- five repeated reopen inspections preserve the same governing snapshot.

Result: **PASS in all four matrix cases**.

### 4.2 Revision-bound transition safety

One accepted transition executes:

```text
validate predecessor
→ insert immutable state revision
→ compare-and-set current pointer
→ record accepted transition
→ record audit
→ record evidence reference
→ COMMIT
```

Verified:

- accepted transition advances exactly one revision;
- stale expected revision is rejected before durable mutation;
- structurally invalid revision jump is explicitly rejected;
- stale/invalid attempts leave governing snapshot and durable row counts unchanged.

Result: **PASS in all four matrix cases**.

### 4.3 Real process-kill fault injection

A child process writes an out-of-database fault marker and blocks. The parent uses the operating-system process-kill primitive; recovery does not rely on child `defer` cleanup.

Transition kill points:

```text
before_tx
after_validation
after_revision_insert
after_pointer_update
before_commit
after_commit
```

Observed invariant:

```text
pre-commit kill
→ old complete governing state
→ no partial revision/audit/evidence rows

after-commit kill
→ new complete governing state
→ revision/current-pointer/audit/evidence coherent
```

Result: **PASS in all four matrix cases**.

### 4.4 WAL/checkpoint recovery

- automatic WAL checkpoint was disabled for the controlled fixture to make WAL behavior observable;
- the test proves a non-empty WAL exists after committed revision before process kill;
- main-file-only naïve copy does not represent the current committed state while the current revision lives in WAL;
- after killing the process, fresh inspection of the original store recovers the committed revision;
- process kill immediately before and after explicit checkpoint preserves governing state.

Result: **PASS in all four matrix cases**.

`wal_autocheckpoint=0` is a spike fixture control, not a production policy decision.

### 4.5 Supported online backup and restore

The disposable implementation used SQLite `VACUUM INTO` as the supported consistent online backup mechanism for the common test surface:

```text
write temporary standalone backup
→ validate it
→ publish by rename only after validation
```

Verified:

- source database remains open while backup is created;
- kill before backup publication leaves no final backup that could masquerade as valid;
- kill after SQLite backup creation but before publication leaves no final published backup;
- kill after publication leaves a valid final backup;
- original working directory is destroyed after backup;
- restore into a fresh directory preserves Aurora ID, Project ID and governing state;
- existing target with a different Aurora identity fails as explicit identity collision and is not replaced.

Result: **PASS in all four matrix cases**.

### 4.6 Corruption, compatibility and migration

Verified:

- truncated/corrupt database is not accepted as governing state;
- unsupported logical schema version fails explicitly;
- explicit v1→v2 migration preserves Aurora/Project IDs, current revision, state semantics and transition/audit/evidence counts.

Result: **PASS in all four matrix cases**.

### 4.7 Integrity and recovery receipts

Each final matrix artifact contains:

- environment/toolchain receipt;
- complete `go list -m all` graph;
- `go.sum` lock receipt;
- machine-readable `go test -json` output;
- build metrics;
- small-fixture runtime metrics;
- 4 MiB medium-fixture backup/restore metrics;
- explicit `PRAGMA integrity_check = ok` receipt;
- SHA-256 of the backup database;
- SHA-256 of the restored database;
- recovered logical snapshot dump.

The aggregate gate fails if these receipts are missing.

## 5. Final measured comparison

Final aggregate from job `92982998192`:

| OS | Candidate | Correctness | Receipts | CGO | SQLite | Build s | Binary MiB | Go deps | Small Tx p95 ms | Medium MiB | Medium Backup ms | Medium Restore ms |
|---|---|---|---|---:|---|---:|---:|---:|---:|---:|---:|---:|
| Linux | modernc | PASS | YES | 0 | 3.53.3 | 0.389 | 9.65 | 140 | 3.087 | 4.06 | 26.386 | 14.895 |
| Linux | mattn | PASS | YES | 1 | 3.53.4 | 0.859 | 7.42 | 115 | 1.959 | 4.06 | 13.207 | 6.354 |
| Windows | modernc | PASS | YES | 0 | 3.53.3 | 0.885 | 9.98 | 150 | 36.073 | 4.06 | 206.007 | 58.307 |
| Windows | mattn | PASS | YES | 1 | 3.53.4 | 1.188 | 11.95 | 117 | 43.338 | 4.06 | 192.010 | 59.714 |

Cross-platform means from the same aggregate:

| Metric | modernc | mattn |
|---|---:|---:|
| all platforms correct | yes | yes |
| all receipts complete | yes | yes |
| CGO | no | yes |
| build time mean | 0.637 s | 1.023 s |
| binary bytes mean | 10,291,248 | 10,153,368.5 |
| Go dependency packages mean | 145 | 116 |
| small transition p95 mean | 19.580 ms | 22.649 ms |
| 4 MiB backup mean | 116.196 ms | 102.608 ms |
| 4 MiB restore mean | 36.601 ms | 33.034 ms |

These timing values are comparative runner observations, not product SLOs or general database benchmarks.

## 6. Artifact identities

Final workflow artifacts:

| Artifact | ID | SHA-256 digest |
|---|---:|---|
| aggregate | `9007749571` | `7fdc58f2cfd7fb898049da1036ec2481275aba7ee3e5352769a4aaa0022b8fcd` |
| Windows / mattn | `9007744317` | `2c733479244fe0b66fedb21341c514affc794066d41807f0301d3f0081932179` |
| Windows / modernc | `9007726989` | `708b754c631616424061df15dd3e83fc9e67e232d2ef6ef66c2c285efe691c3a` |
| Ubuntu / mattn | `9007718781` | `e84ec8b0f6543177f770f24c1c8979d02f5db26c8d7884bfeb64bedec3f5e772` |
| Ubuntu / modernc | `9007705578` | `5cea38ab644228b3a53163654eafc79170ce79c8908d8707a8cbbcc07bcf36fc` |

Artifacts are retained by GitHub Actions for the configured retention period. This canonical receipt preserves their identities/digests after expiry.

## 7. Material debugging finding

An intermediate Windows/modernc process-kill test failed with `TerminateProcess: Access is denied` after the child had already written the fault marker.

Systematic debugging found the harness used `select {}` as its indefinite block. That can allow Go's deadlock detector to terminate the helper before the parent performs the OS kill, making the failure a test-harness artifact rather than database evidence.

The helper was changed only to a timer-backed blocking loop. The same Windows/modernc candidate then passed the complete fault suite, and the final run passed again with the expanded evidence receipts.

This finding is preserved because hiding a false negative would weaken confidence in the experiment method.

## 8. Evidence limitations

This spike proves **process-kill/crash recovery in controlled GitHub-hosted Ubuntu/Windows filesystems**. It does not emulate:

- abrupt physical power removal during storage-controller writeback;
- lying/broken filesystem or storage hardware that violates fsync expectations;
- total host/root compromise;
- network filesystem behavior;
- high-concurrency multi-process writers;
- production load/SLO behavior.

Those are not current M0 requirements. The selected SQLite durability posture must still retain `synchronous=FULL`, supported backup mechanisms, explicit integrity/version checks and the storage-neutral domain boundary.

## 9. Evidence implication

Both SQLite bindings satisfy the tested M0 correctness gates.

The evidence therefore activates the spike's tie-break rule rather than a correctness rejection:

```text
correctness tie
→ compare operational/build burden
→ performance remains secondary
```

The independent review determines the decision-informed recommendation and residual risks.
