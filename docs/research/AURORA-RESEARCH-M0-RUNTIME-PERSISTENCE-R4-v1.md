---
id: RESEARCH-AURORA-M0-RUNTIME-PERSISTENCE-R4-V1
title: Aurora M0 R4 Research — Runtime, Persistence and Durable Execution
_document_type_note: focused R4 architecture research
document_type: research_report
form: explanation
authority: research
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - M0 R4 runtime and persistence research through 2026-08-07
related:
  - DOC-AURORA-M0-R4-OPERATOR-AUTHORIZATION
  - DESIGN-AURORA-M0-R4-DECISION-LANDSCAPE
  - DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC
  - DOC-AURORA-CAP-SOVEREIGN-CORE-THREAT-MODEL
  - DOC-AURORA-CAP-SOVEREIGN-CORE-TEST-PLAN
source_manifest: AURORA-RESEARCH-M0-RUNTIME-PERSISTENCE-R4-v1.sources.json
review_triggers:
  - Architecture Spike evidence for M0 storage/recovery
  - Go or Rust major/runtime support change
  - selected database or driver major change
  - M1–M4 requirement changes that alter write concurrency or execution lifecycle
last_reviewed: 2026-08-07
---

# Aurora M0 R4 — Runtime, Persistence and Durable Execution

## 1. Research question

Which runtime and persistence direction can implement the proposed `CAP-SOVEREIGN-CORE` with the smallest operational footprint while preserving a credible evolution path and proving:

- stable identities;
- revision-bound canonical state;
- atomic accepted-state/audit updates;
- process-independent restart/recovery;
- local sovereignty;
- safe backup/export/restore;
- explicit migration;
- future adapter replacement without domain rewrite?

The research also asks whether M0 needs a durable workflow engine in addition to a durable operational-state store.

## 2. Method

Evaluation uses the approved R4 philosophy:

```text
H0 — pass M0 correctly
H1 — avoid foreseeable conflict with M1–M4
H2 — preserve constitutional long-term direction
```

A mechanism is not preferred because it is popular, minimal or feature-rich. It is preferred when it is the smallest current mechanism that survives the horizon analysis and has a contained exit path.

Primary/current sources only are used for technical claims. Runtime behavior that depends on driver/configuration/filesystem integration is explicitly reserved for Architecture Spike evidence rather than inferred from documentation.

## 3. Decision drivers from R3

The most constraining R3 properties are:

1. one current canonical Project state revision exists independently from process memory;
2. current state and audit/events remain logically distinct;
3. accepted mutation cannot be reported successful after a partial commit;
4. restart cannot rely on transcript/model/Harness state;
5. restore must validate identity, version, integrity and authority safety;
6. local-first/single-user is the current scope;
7. domain meaning cannot be owned by database/runtime semantics;
8. migration/exit must be explicit;
9. no durable workflow engine may be introduced merely because state must survive process death.

## 4. Runtime candidates

### 4.1 Go

Current line: Go 1.26, with the current 1.26.x stable line documented by the Go project [S01][S02].

Relevant properties:

- strong compatibility policy across releases [S01][S02];
- standard `database/sql` abstraction supports SQLite and PostgreSQL through drivers [S03];
- `sql.Tx` gives a simple application boundary for grouped commit/rollback [S04];
- simple static-binary operational model when dependencies permit;
- mature cross-platform standard library;
- garbage-collected memory model, acceptable for a control/state Core whose dominant risk is not hard-real-time latency;
- no requirement that future Harness, device or firmware capabilities share the Core language because ADR-0001 preserves language-neutral boundaries.

The main SQLite caveat is not Go itself but the binding choice:

- `modernc.org/sqlite` is CGo-free and currently covers major Linux/macOS/Windows architectures, but its own documentation warns that `modernc.org/libc` must match exactly because of a fragile dependency [S15];
- `mattn/go-sqlite3` is a mature `database/sql` driver but requires CGO/C toolchain [S16];
- `ncruces/go-sqlite3` remains interesting, but its own release notes explicitly call the recent `wasm2go` transition comparatively new and suggest older versions for users prioritizing stability [S17].

Therefore `Go` can be evaluated separately from the final SQLite driver.

### 4.2 Rust

Current stable line: Rust 1.97.1 [S05].

Relevant properties:

- ownership/borrowing provides compile-time memory-safety guarantees without GC [S06];
- strong type system is attractive for state-machine and security-sensitive code;
- `rusqlite` offers an ergonomic SQLite binding and a bundled SQLite mode that reduces dependency on host SQLite [S07];
- Cargo provides strong build/test tooling.

However, for M0 the dominant failure modes are stale revisions, partial durable state, unsafe restore, rollback, authority freshness and audit/evidence consistency. Rust memory safety reduces an important class of implementation bugs but does not remove those system-level correctness obligations.

`rusqlite` ultimately wraps SQLite's C API; choosing Rust therefore does not eliminate native SQLite integration. It shifts the language/runtime cost without eliminating the required crash/recovery proof.

### 4.3 Runtime comparison

| Driver | Go | Rust | R4 interpretation |
|---|---|---|---|
| stable current runtime | strong | strong | tie |
| memory safety without GC | no | yes | Rust advantage |
| simple relational std abstraction | `database/sql` in stdlib | ecosystem crate | Go advantage |
| local operational simplicity | high | high | roughly tie |
| SQLite integration | multiple drivers; pure-Go or CGO | bundled C SQLite via rusqlite | both need proof |
| observability ecosystem for proposed Core | OpenTelemetry Go traces/metrics stable | Rust OTel maturity currently lower | Go advantage for current horizon |
| complexity for M0 control/state Core | lower language/tooling burden | higher ownership/lifetime cognitive burden | Go advantage |
| future hardware/low-level code | can coexist via contracts | excellent | not a Core-runtime reason |
| lock-in containment | strong if domain ports retained | strong if domain ports retained | tie |

### 4.4 Runtime research recommendation

```text
RECOMMENDED R4 DIRECTION: Go for Aurora Core
CONFIDENCE: high documentary confidence
SPIKE REQUIRED FOR LANGUAGE ITSELF: no
SPIKE REQUIRED FOR SELECTED SQLite BINDING/RECOVERY STACK: yes
```

The recommendation is not “Go because it is simpler.” It is that Rust's major differentiating guarantee—compile-time memory safety without GC—does not dominate the actual M0 risk profile strongly enough to justify making the Core implementation and its entire application ecosystem more complex. Language-neutral Aurora boundaries keep Rust available for later capabilities where low-level safety/performance is a primary requirement.

No third runtime candidate is added because no accepted M0 requirement exposes a gap that Go/Rust cannot cover; adding another candidate would increase research breadth without reducing a named uncertainty.

## 5. Persistence candidates

### 5.1 SQLite

SQLite officially positions itself as local application/device storage emphasizing economy, independence, reliability and simplicity [S08]. Its documentation says client/server databases are preferable when many clients across a network or many concurrent writers must access the same database, while device-local low-writer-concurrency storage is a canonical SQLite use case [S08].

That matches M0 closely:

```text
one local Core
single user
low writer concurrency
state small relative to SQLite limits
no network-shared DB
no horizontal write scaling requirement
```

Important correctness properties:

- transactions and crash recovery are built into SQLite;
- WAL mode with `synchronous=FULL` is documented as ACID across power loss, while `NORMAL` can lose recent commits after power loss [S09];
- safe live backups exist through the Online Backup API [S10];
- SQLite explicitly documents that naïvely copying a database file during an active transaction can yield invalid backup state, and hot journals/WAL must not be separated from the database [S11].

R4 implication: if SQLite wins, Aurora must use supported snapshot/backup APIs and must not define “backup” as arbitrary filesystem copy.

### 5.2 PostgreSQL 18

PostgreSQL 18 provides a mature client/server durability model using WAL and data-page checksums [S12]. Its backup system supports SQL dumps, filesystem-level backups and continuous archiving/PITR [S13].

It becomes strongly preferable when:

- multiple independent processes need concurrent writes;
- database access must be networked;
- centralized administration or replication becomes a current product need;
- future durable-execution technology specifically requires Postgres.

Those are not current M0 requirements. Selecting PostgreSQL now creates:

- a long-running server service;
- install/start/upgrade/backup administration distinct from Aurora;
- credentials/network/socket configuration;
- a larger recovery topology than the M0 Golden Proof requires.

The R3 `Durable State Port` means a future migration to PostgreSQL can be engineered as a store replacement without changing domain meaning, provided schema/migration ownership remains application-level.

### 5.3 RocksDB / embedded key-value class

RocksDB is an embedded persistent key/value LSM engine oriented to high-throughput server workloads and flexible storage tuning [S14]. It supports WAL, backups and ACID variants.

The mismatch is not reliability; it is abstraction level. Aurora would need to own more of:

- relational constraints;
- queries over revisions/authority/audit;
- secondary indexes;
- migration semantics;
- transaction/domain mapping;
- operational compaction/tuning choices.

For M0 there is no evidence that this additional control solves a problem that SQLite cannot solve. Therefore embedded KV engines are removed from the M0 shortlist unless future spike evidence falsifies the relational-store premise.

## 6. Persistence horizon analysis

### H0 — M0

SQLite has the strongest scope match. PostgreSQL is correct but operationally larger. RocksDB is lower-level than needed.

### H1 — M1–M4

Foreseeable pressures:

- M1 memory/context may add more read/write structures but does not automatically require multi-writer network access;
- M2 Registry/AHDK adds metadata/trust state that remains naturally relational;
- M3 Delegation adds more state transitions but can remain within a single local Core;
- M4 is the first milestone where long-running durable execution and duplicate-effect prevention can materially change the execution architecture.

SQLite therefore does not show an obvious structural dead end through the near horizon as long as:

- Aurora domain/state APIs remain storage-neutral;
- migrations are application-owned;
- direct SQL does not leak across module boundaries;
- future multi-process write requirements trigger reconsideration.

### H2 — long-term Aurora

A future distributed personal system may use a networked or replicated store. R3 already requires adapter evolution rather than domain rewrite. The long-term requirement is therefore **migratability**, not premature adoption of a distributed database in M0.

## 7. State-versus-event model

R3 already fixes the product semantics strongly enough for an R4 no-regret decision:

```text
canonical current state
+ immutable accepted state revisions
+ logically distinct audit/domain events
+ telemetry separate
```

Full event sourcing is not required.

Choosing event sourcing now would create new obligations—event replay compatibility, projection rebuild semantics, event-schema permanence, event migration and event-log authority—that are not needed for M0 and could conflict with the explicit rule that events are not the sole canonical owner of current state.

Research recommendation:

```text
DO NOT use full event sourcing for M0.
Persist explicit current-state/revision structures transactionally.
Persist audit/domain-event records as distinct records, potentially in the same physical relational database.
```

## 8. Atomicity model to prove

The candidate implementation should use one database transaction for the state mutation boundary containing at minimum:

```text
validate current revision
validate current authority inputs
insert immutable new state revision
update current-state pointer/revision
record transition attempt/result
record required audit metadata
record evidence metadata/reference needed to explain acceptance
commit
```

External immutable artifact content, when applicable, can be staged before this transaction and referenced by stable hash/ID; the transaction must not claim success until all required references are valid.

The exact SQL/schema belongs to R6, but the atomic boundary is an R4 architectural commitment because crash behavior depends on it.

## 9. Backup and recovery implications

Two different artifacts are needed:

1. **operational database backup/snapshot** — mechanism-specific operational recovery;
2. **Aurora logical export** — portable sovereignty artifact governed by Aurora schema/version/integrity rules.

They must not be conflated.

For SQLite, Online Backup API or another documented consistent snapshot method is the candidate operational-backup mechanism [S10][S11].

A logical export remains necessary even if raw SQLite backup is easy because M0 requires versioned portable restore/migration semantics independent of one storage product.

## 10. Durable workflow engine applicability

DBOS checkpoints workflow/step execution to PostgreSQL and recovers interrupted workflows [S18]. Restate journals execution and supports durable state/timers [S19]. Temporal, including Mastra's current integration, is explicitly aimed at workflows that call external APIs, run for hours/days or must outlive workers [S20].

M0 currently has none of the following:

- hours/days-long Mission workflow;
- durable timer;
- wait for human input inside a Mission;
- external effect retry/idempotency chain;
- distributed worker execution;
- provider recovery lifecycle.

M0 only requires short internal state operations whose **results** must be durable across process restart.

Introducing a workflow engine now would create a second durable execution/history system whose relation to Aurora canonical state would itself need recovery/versioning/backup rules.

R4 recommendation:

```text
R4-Q-ENGINE-001 = NOT_YET_A_DECISION for M0
Do not introduce a durable workflow engine.
Reconsider at M4 or earlier only if a current accepted requirement cannot be safely implemented as a bounded transactional operation.
```

## 11. Operational burden matrix

| Candidate | Processes/services | Backup mechanics | Write concurrency | M0 burden | Exit path |
|---|---:|---|---|---|---|
| SQLite | Core only | file + supported snapshot API | one writer at a time | lowest | logical export + Durable State Port |
| PostgreSQL | Core + DB server | mature server tools | high | materially higher | standard logical migration |
| RocksDB | Core only | engine-specific backup/checkpoint | engine-dependent | medium/high due low-level ownership | app-specific KV migration |
| DBOS + PostgreSQL | Core + Postgres, optionally control plane | DB + workflow history | high | too high for current scope | workflow rewrite/removal cost |
| Restate | Core/service + Restate runtime | runtime+journal state | distributed-oriented | too high for current scope | execution-runtime migration |

## 12. Documentary recommendations

### Strong enough for proposed ADR without spike

- `R4-Q-CORE-001`: **Go** as Core runtime.
- `R4-Q-STATE-001`: explicit current state + immutable revisions + distinct audit/events; **no full event sourcing**.
- `R4-Q-TOPOLOGY-001`: **one local modular Core process** for M0.
- `R4-Q-ENGINE-001`: **no durable workflow engine in M0**; reconsider at M4/requirement trigger.

### Strong candidate but executable evidence still required

- `R4-Q-STORE-001`: **SQLite** as M0 operational store.
- `R4-Q-ATOMIC-001`: single SQLite transaction for current-revision/state/audit boundary.
- driver: shortlist **modernc.org/sqlite** versus **mattn/go-sqlite3**; current evidence favors modernc operationally because it removes CGO, but its own fragile-libc warning makes a comparative spike prudent.
- WAL/durability configuration: candidate `journal_mode=WAL` + `synchronous=FULL`, to be verified under fault injection rather than merely copied from documentation.

## 13. Required Architecture Spike

Accepted research governance already requires a crash/restart/restore spike for Sovereign Core storage/recovery. Documentary research cannot close this uncertainty.

The minimum spike should compare, on the exact selected Go version:

```text
A — Go + SQLite via modernc.org/sqlite
B — Go + SQLite via mattn/go-sqlite3
```

PostgreSQL does not need to be implemented in the first spike unless SQLite fails a gate criterion; its documented semantics and higher operational burden make it the fallback class, not an equal current frontrunner.

The spike must measure/prove:

- build/reproduction on target environments;
- transaction/crash consistency;
- kill before/after commit;
- WAL/checkpoint restart;
- same identities/state after fresh process;
- invalid transition no mutation;
- safe live backup;
- restore into fresh environment;
- corrupt/truncated state detection;
- migration fixture;
- operational footprint and dependency burden.

## 14. Decision implications

If the spike passes for SQLite:

```text
Go + SQLite
→ candidate R4 storage/runtime baseline
→ logical export remains independent
→ PostgreSQL remains documented future migration target class
→ no durable engine in M0
```

If both SQLite bindings fail materially:

```text
expand spike to PostgreSQL
or revisit runtime/binding assumptions
```

Do not silently weaken durability, backup or cross-platform criteria to preserve SQLite.

## 15. Limitations

This report does not:

- execute fault injection;
- prove driver-specific fsync/WAL behavior on Aurora's real environment;
- establish final target OS support matrix;
- accept a stack;
- authorize implementation.

Its output is R4 evidence and candidate-decision input only.
