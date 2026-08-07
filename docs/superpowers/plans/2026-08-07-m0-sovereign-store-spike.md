---
id: PLAN-AURORA-M0-SOVEREIGN-STORE-SPIKE-001
title: SPK-001 M0 Sovereign Store Execution Plan
document_type: implementation_plan
form: guide
authority: design
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - execution plan for the authorized disposable SPK-AURORA-M0-SOVEREIGN-STORE-001 experiment
related:
  - DOC-AURORA-M0-R4-ADR-ACCEPTANCE-SPK001-AUTHORIZATION
  - DESIGN-AURORA-M0-SOVEREIGN-STORE-SPIKE-001
  - ADR-AURORA-0003
  - ADR-AURORA-0004
  - ADR-AURORA-0007
source_revision: a3192afad3dba9c6e699588c07ca2bcaac1161fd
spike_spec_revision: 36f46956bc275d0aec32b7e3ea4d959010fa9dcb
spike_spec_blob: 6ad7397d46208a0a9c762073d2c5239ceff4e056
last_reviewed: 2026-08-07
---

# SPK-001 — M0 Sovereign Store Execution Plan

> **DISPOSAL RULE: DISCARD.** Everything under `spikes/m0-sovereign-store-001/` is disposable executable evidence. Success does not promote it into the Aurora Core.

## 1. Goal

Answer the exact authorized R4 uncertainty with executable evidence:

```text
Can Go + SQLite satisfy M0 state/atomicity/crash/restart/backup/restore needs,
and which of the two pinned Go SQLite bindings is the better first baseline?
```

This plan does not implement the production Aurora Core and does not authorize SPK-002 or R5.

## 2. Fixed candidate set

```text
Go: 1.26.5

Candidate A:
  modernc.org/sqlite v1.54.0
  database/sql
  WAL
  synchronous=FULL

Candidate B:
  github.com/mattn/go-sqlite3 v1.14.49
  database/sql
  WAL
  synchronous=FULL

Fallback:
  PostgreSQL 18 only if both SQLite candidates fail a gate criterion or expose an inherent client/server requirement.
```

The module lock produced by `go mod tidy` is evidence and owns exact transitive dependency versions.

## 3. Minimal disposable design

One shared black-box suite drives both bindings through the same store surface:

```text
internal/sqlitedriver
  build-tag selected database/sql driver only

internal/store
  Open / Bootstrap / Inspect
  Transition
  IntegrityCheck
  SupportedBackup
  Restore
  MigrationFixture

internal/fault
  deterministic hook names
  child-process marker/block protocol

cmd/spike-runner
  environment/version reporting
  deterministic benchmark/backup fixtures
```

The fixture contains only Aurora identity, one Project, immutable state revisions, current revision pointer, minimal authority revision reference, transition/audit rows, evidence reference and backup metadata.

## 4. Test-first sequence

### Task 1 — RED: store contract does not exist

Create a behavioral test that requires a durable store to bootstrap one Aurora/Project and recover the same logical state after reopen. Run it and verify the expected failure before implementation.

### Task 2 — GREEN: common schema + both driver adapters

Implement only enough common store behavior to pass bootstrap/reopen for both build tags. Pin both bindings and record SQLite engine version.

### Task 3 — transaction invariants

Add tests first for:

- accepted revision mutation;
- stale expected revision rejection;
- state revision immutability;
- current pointer/audit/evidence coherence.

Then implement the one transaction boundary from the spike specification.

### Task 4 — real child-process fault injection

Add a self-reexec helper used by black-box tests. Parent process observes an out-of-DB marker and uses OS process kill. Execute transition fault points:

```text
before_tx
after_validation
after_revision_insert
after_pointer_update
before_commit
after_commit
```

For pre-commit points recovery must be exactly old state; after commit recovery must be exactly new state. No defers from the killed child are trusted.

### Task 5 — WAL/checkpoint recovery

Disable automatic WAL checkpoint for the fixture. Prove restart with a surviving WAL and kill/restart immediately around an explicit checkpoint boundary. Record the limitation that this is process-kill evidence, not a storage-controller/power-loss emulator.

### Task 6 — supported online backup and restore

Use one SQLite-supported consistent online backup mechanism common to both bindings. The implementation must write to a temporary destination and publish the final backup only after success so an interrupted backup cannot masquerade as valid.

Prove:

- online backup while DB remains open;
- kill around/during backup leaves either no published backup or a valid published backup;
- fresh-directory restore reproduces Aurora/Project identities and governing state;
- naïve live copy of only the main DB file while uncheckpointed WAL contains the current revision is demonstrably stale and is rejected as Aurora backup behavior.

### Task 7 — failure fixtures

Tests first for:

- truncated/corrupt database → explicit failure;
- unsupported schema version → explicit incompatibility;
- identity collision during restore → explicit collision;
- v1→v2 migration fixture preserves IDs/current revision/state semantics.

### Task 8 — measurements and controlled matrix

Run the same suite on:

```text
ubuntu-24.04 amd64
windows-latest amd64
×
modernc
mattn
```

Capture:

- exact runner and Go/toolchain environment;
- `go env` and module graph;
- SQLite runtime version;
- build success/time/binary size;
- CGO/compiler requirement;
- startup/open timing;
- tiny transition latency p50/p95/max;
- supported backup/restore timings for deterministic fixtures;
- WAL/database behavior and integrity checks;
- machine-readable case matrix;
- logs and limitations.

Correctness gates are mandatory; performance is secondary.

### Task 9 — evidence synthesis and adversarial review

Aggregate all four matrix artifacts into one comparison report. Review against every gate criterion in the canonical spike spec.

Decision rule:

```text
both correct
→ lower operational/build burden
→ fewer fragile dependencies/toolchains
→ clearer cross-platform reproduction

one fails correctness
→ reject that candidate

both fail correctness
→ do not weaken requirements; broaden R4 investigation to PostgreSQL
```

## 5. Stop conditions

Stop the spike and return `SPIKE_BLOCKED/FAILED` if any canonical stop condition is reached. Do not debug a non-material environment oddity indefinitely; capture it, determine whether it violates a gate criterion, and move according to the decision rule.

## 6. Canonical closeout boundary

The experimental branch may retain raw disposable code long enough to reproduce evidence. Canonical `main` receives only reviewed evidence/decision/tracking changes; spike code is not promoted as production implementation.

After SPK-001 review:

```text
update ADR-0007 based on evidence
→ record SPK-001 verdict
→ STOP
→ SPK-002 remains unauthorized until explicit operator authorization
```
