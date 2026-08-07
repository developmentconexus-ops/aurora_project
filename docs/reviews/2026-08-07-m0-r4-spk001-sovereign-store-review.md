---
id: REVIEW-AURORA-M0-R4-SPK001-SOVEREIGN-STORE-2026-08-07
title: M0 R4 SPK-001 Sovereign Store Evidence Review
document_type: adversarial_review
form: explanation
authority: evidence
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - reviewed SPK-001 evidence and decision-informed recommendation
related:
  - DOC-AURORA-M0-R4-SPK001-EVIDENCE-RECEIPT
  - DESIGN-AURORA-M0-SOVEREIGN-STORE-SPIKE-001
  - ADR-AURORA-0007
  - DOC-AURORA-STATUS
spike_execution_revision: 4242342486f512320f12e0b603f052166264c4ea
workflow_run: 31213792366
reviewed_at: 2026-08-07
last_reviewed: 2026-08-07
---

# M0 R4 — SPK-001 Sovereign Store Evidence Review

## 1. Verdict

```text
SPK-AURORA-M0-SOVEREIGN-STORE-001: PASS

Lifecycle:
AUTHORIZED
→ EXECUTING
→ EVIDENCE_COMPLETE
→ REVIEWED
→ DECISION_INFORMED
→ CLOSED
```

The evidence is sufficient to answer the authorized R4 uncertainty:

1. **SQLite is viable for the M0 Sovereign Core operational-state store under the tested M0 topology and failure model.**
2. Both evaluated Go bindings satisfy the same tested correctness gates on Ubuntu and Windows.
3. Given the correctness tie, **`modernc.org/sqlite` is the recommended initial binding baseline** because it removes CGO/C-toolchain dependence from the Sovereign Core build while reproducing the required behavior on both platforms.
4. PostgreSQL expansion is not justified by current M0 evidence and is not executed.

This review does **not** accept ADR-0007. ADR-0007 remains `proposed` until the operator accepts/rejects/revises the evidence-informed decision.

## 2. Gate-by-gate review

| Canonical spike gate | Evidence | Review |
|---|---|---|
| bootstrap/reopen preserves stable identities | repeated fresh-handle recovery in all 4 matrix cases | PASS |
| accepted revision mutation is coherent | revision/current pointer/audit/evidence in one transaction | PASS |
| stale transition changes nothing | explicit stale predecessor case + unchanged snapshot/counts | PASS |
| invalid transition changes nothing | explicit invalid revision jump + unchanged snapshot/counts | PASS |
| pre-commit process kill cannot leave partial governing state | real child `Process.Kill()` at five pre-commit boundaries | PASS |
| post-commit process kill recovers complete new state | `after_commit` kill + exact new snapshot/counts | PASS |
| WAL-present restart recovers committed state | non-empty WAL observed before kill; fresh inspection recovers revision | PASS |
| interrupted checkpoint preserves governing state | kill before/after explicit checkpoint | PASS |
| supported live backup is consistent | database remains active; SQLite-supported backup creates validated standalone copy | PASS |
| interrupted backup cannot publish partial final backup | kill before publication and after backup SQL before publication | PASS |
| fresh restore survives loss of original working directory | original work directory deleted; backup restored to fresh directory | PASS |
| identity collision does not silently replace target | explicit collision error + unchanged existing target | PASS |
| corrupt/truncated state fails explicitly | corruption fixture rejected | PASS |
| unsupported logical schema fails explicitly | schema `999` rejected | PASS |
| migration preserves protected semantics | v1→v2 fixture preserves IDs/current state and audit/evidence counts | PASS |
| naïve live main-file copy is not Aurora backup | current revision retained in non-empty WAL; main-file-only copy does not represent current state | PASS |
| Linux and Windows reproducibility | 4/4 final matrix jobs green | PASS |
| evidence receipts complete | aggregate enforced tests/build/modules/go.sum/integrity/hashes/logical recovery/small+medium metrics | PASS |
| no client/server database required | complete M0 proof achieved with embedded SQLite | PASS |

No gate criterion requires broadening the experiment to PostgreSQL.

## 3. Candidate comparison

### 3.1 Correctness

There is no observed correctness winner:

```text
modernc: PASS on Linux + Windows
mattn:   PASS on Linux + Windows
```

Both survive the exact same transaction, process-kill, WAL, checkpoint, backup/restore, corruption, compatibility and migration fixtures.

Therefore performance is not allowed to override the architectural tie-break rule.

### 3.2 Build and operational burden

`modernc.org/sqlite`:

- `CGO_ENABLED=0` in the tested matrix;
- no C compiler/toolchain required for the binding path;
- mean measured build time in the final matrix: `0.637 s`;
- more Go dependency packages in this fixture: mean `145`;
- mean binary size approximately `9.82 MiB`;
- official package documentation identifies a fragile exact-version relationship with `modernc.org/libc`, so upgrades must preserve a tested compatible module lock.

`mattn/go-sqlite3`:

- `CGO_ENABLED=1` in the tested matrix;
- requires a functioning C compiler/toolchain and CGO-compatible cross-platform build environment;
- mean measured build time in the final matrix: `1.023 s`;
- fewer Go dependency packages in this fixture: mean `116`;
- mean binary size approximately `9.68 MiB`;
- mature direct SQLite C binding and slightly lower mean medium backup/restore times in this run.

For the M0 local Sovereign Core, eliminating the native build/toolchain dependency is the more material operational simplification. The additional Go-module graph and modernc/libc pinning caveat are real but can be controlled by exact module locks and the same cross-platform conformance suite on upgrades.

### 3.3 Performance interpretation

Final means:

```text
small transition p95:
  modernc ~19.58 ms
  mattn   ~22.65 ms

4 MiB backup:
  modernc ~116.20 ms
  mattn   ~102.61 ms

4 MiB restore:
  modernc ~36.60 ms
  mattn   ~33.03 ms
```

These differences are small relative to M0's product requirements and vary by hosted runner/platform. They are **not** SLOs and do not justify selecting the more operationally burdensome binding.

## 4. Recommended R4 decision

Evidence supports revising ADR-0007 to:

```text
M0 operational state store:
  SQLite

initial Go binding baseline:
  database/sql + modernc.org/sqlite

initial evidence-qualified release:
  modernc.org/sqlite v1.54.0

required durability posture:
  WAL
  synchronous=FULL
  foreign_keys=ON
  one atomic accepted-state/current-pointer/audit/evidence transaction
  integrity/version validation before governing use
  supported consistent SQLite backup mechanism
  no naïve live main-file copying as backup

version policy:
  exact dependency/module lock at implementation baseline
  rerun the cross-platform store suite when the SQLite binding/runtime is materially upgraded
```

The following remain implementation-design details, not new R4 blockers:

- exact connection-pool values;
- automatic versus explicit checkpoint cadence;
- exact production backup API wrapper, provided it preserves the proved consistent-snapshot semantics;
- physical table/index naming beyond required domain invariants.

The spike used `wal_autocheckpoint=0` only to expose WAL recovery deterministically. That is **not** recommended as a permanent product setting by this review.

## 5. Why PostgreSQL is not advanced now

The authorized fallback said PostgreSQL should be added if:

- both SQLite candidates fail a correctness gate; or
- evidence exposes a current requirement for the client/server class.

Neither occurred.

Adding PostgreSQL now would therefore violate the product-velocity rule by answering a question whose trigger did not fire. PostgreSQL remains the documented future migration/fallback class if independent concurrent writers, network-shared access or another current requirement materially changes the topology.

## 6. Material debugging finding

An intermediate Windows/modernc run produced:

```text
TerminateProcess: Access is denied
```

after the external fault marker existed.

The failure was not classified as database corruption. Investigation found the child helper used `select {}` as an artificial permanent block, allowing Go runtime deadlock behavior to interfere with the intended parent-driven kill on that environment.

Only the harness blocking method was changed to a timer-backed loop. Windows/modernc then passed the same fault cases, and the final expanded run passed again.

Disposition:

```text
TEST_HARNESS_FALSE_NEGATIVE
RESOLVED
```

This finding does not count as a candidate defect.

## 7. Residual risks and limitations

The spike did not emulate physical storage hardware violating persistence contracts or literal power removal during controller writeback. Evidence is process-kill based on GitHub-hosted Ubuntu/Windows filesystems.

Residuals that remain acceptable for this R4 decision:

- total host/root/storage-controller compromise is outside M0;
- no network filesystem is supported by this evidence;
- high-concurrency independent writers are outside M0;
- exact future modernc/libc version compatibility must be revalidated on upgrades;
- production workload performance has not been benchmarked because M0 has no throughput SLO requiring it.

These limits do not invalidate the M0 store decision because the current milestone is one local modular Core with low writer concurrency.

## 8. Disposal / non-promotion

The experimental branch:

```text
spike/m0-sovereign-store-001
```

is evidence only. It is not to be fast-forwarded or merged into `main` as Aurora Core implementation.

Canonical `main` receives only reviewed documentation/evidence/decision state. A later R5/R6/R7 implementation must be written against the accepted contract/design and may reuse **knowledge**, not spike code by automatic promotion.

## 9. R4 implication

SPK-001 is no longer an R4 evidence blocker.

Current next decision:

```text
operator reviews evidence-informed ADR-0007
→ accept / reject / revise
```

Even if ADR-0007 is accepted, R4 does not automatically pass because ADR-0008 still depends on separately unauthorized `SPK-AURORA-M0-OWNER-TRUST-002`.

No SPK-002 execution and no R5 work is authorized by this review.
