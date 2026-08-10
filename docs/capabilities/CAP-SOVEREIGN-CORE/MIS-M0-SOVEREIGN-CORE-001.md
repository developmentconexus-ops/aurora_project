---
id: DOC-AURORA-MIS-M0-SOVEREIGN-CORE-001
title: MIS-M0-SOVEREIGN-CORE-001 Mission Contract
document_type: mission_contract
form: reference
authority: contract
status: accepted
accepted_at: 2026-08-09
acceptance_evidence: DOC-AURORA-M0-R5-A2-CONTRACT-OPERATOR-ACCEPTANCE
accepted_from_blob: 1db39012874828f54f293bf76571259494ba5a79
version: 0.1.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - exact scoped implementation commitment for M0 Sovereign Core Walking Skeleton
related:
  - DOC-AURORA-M0-R5-OPERATOR-AUTHORIZATION
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DOC-AURORA-CAP-SOVEREIGN-CORE-REQUIREMENTS
  - DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC
  - DOC-AURORA-CAP-SOVEREIGN-CORE-THREAT-MODEL
  - DOC-AURORA-CAP-SOVEREIGN-CORE-TEST-PLAN
  - DOC-AURORA-CAP-SOVEREIGN-CORE-R3-COVERAGE
  - DOC-AURORA-CAP-SOVEREIGN-CORE-R4-DECISION-COVERAGE
  - ADR-AURORA-0003
  - ADR-AURORA-0004
  - ADR-AURORA-0005
  - ADR-AURORA-0006
  - ADR-AURORA-0007
  - ADR-AURORA-0008
  - ADR-AURORA-0009
  - DOC-AURORA-M0-R5-A2-CONTRACT-OPERATOR-ACCEPTANCE
mission_id: MIS-M0-SOVEREIGN-CORE-001
contract_revision: 0.1.0
source_revision: 74167bd1404d9076423ffdbae20f97958283527c
last_reviewed: 2026-08-07
---

# MIS-M0-SOVEREIGN-CORE-001 — Mission Contract

## 1. Status and contract boundary

```text
Mission:          MIS-M0-SOVEREIGN-CORE-001
Contract:         v0.1.0
Capability:       CAP-SOVEREIGN-CORE
Product Milestone: M0 — Sovereign Core Walking Skeleton
Baseline:         74167bd1404d9076423ffdbae20f97958283527c
Status:           APPROVED
```

This Contract commits one **vertical M0 walking skeleton**. It does not split Identity, Project State, Authority, Recovery or Audit/Evidence into independently productized services or Missions.

The Contract was explicitly approved by the operator through `DOC-AURORA-M0-R5-A2-CONTRACT-OPERATOR-ACCEPTANCE` and is the governing scoped M0 implementation commitment. Approval does not itself authorize R6, production implementation or merge/promotion.

## 2. Outcome

### Objective

Build and verify the smallest real Aurora Sovereign Core that preserves its own identity, one Project's current accepted operational state, current M0 authority semantics and a derived next-safe-action projection across fresh-process restart, while safely supporting invalid-transition rejection, export, restore and explicit migration.

### Operator-visible value

Leandro can initialize Aurora, create one Project, record accepted state, stop every Aurora process, start again and observe the **same** Aurora/Project identity and governing state rather than a new session pretending to be continuity.

### Product risk retired

```text
Aurora is not merely a running model/session whose identity, Project truth or authority disappears or silently changes when the process dies.
```

## 3. Golden Proof contribution

The Mission must implement and produce evidence for the full M0 Product Milestone journey:

```text
initialize Aurora instance
→ create Project
→ record accepted state and proposed next action
→ evaluate current authority and show permitted/blocked next safe action
→ terminate all Aurora processes
→ start a fresh process
→ recover the same Aurora/Project identities, current state and authority result
→ reject an invalid/stale transition with zero governing mutation
→ export logical sovereignty state
→ restore into a fresh context
→ enforce restore authority freshness/revalidation semantics
→ demonstrate recovered current state and permitted next action only after valid authority
```

R8 remains the owner of final Product Milestone acceptance. Passing this Mission does not automatically close M0.

## 4. Contract criteria

### `MIS-M0-SOVEREIGN-CORE-001-CRIT-001` — Stable identity and bounded M0 scope

The implementation MUST initialize one stable Aurora identity and one stable Project identity, preserve them across restart/restore, remain Leandro-first/single-user and demonstrate that model, Harness, database, process and proof UI are not identity owners.

Primary requirement allocation: `REQ-001..009`.

### `MIS-M0-SOVEREIGN-CORE-001-CRIT-002` — Canonical state ownership

The Core MUST own exactly one current accepted Project state revision and the M0 authority/current-next-action state; projections, audit, telemetry, narrative, model/Harness state and database-specific representations MUST remain non-authoritative as defined by the Capability Spec.

Primary requirement allocation: `REQ-010..020`.

### `MIS-M0-SOVEREIGN-CORE-001-CRIT-003` — Revision-bound transition lifecycle

A valid authorized transition MUST create exactly one attributable immutable accepted revision and advance the current pointer coherently. Stale, malformed, identity-mutating or unauthorized attempts MUST be rejected without changing governing Project state, authority state or derived next action. Success MUST be based on observed durable result, not dispatch.

Primary requirement allocation: `REQ-021..030`.

### `MIS-M0-SOVEREIGN-CORE-001-CRIT-004` — Authority and next-safe-action correctness

Technical access MUST NOT become authority. Active, scoped, time-valid authority may permit the matching internal M0 operation; wrong-scope, expired, revoked, superseded, missing, corrupt or unproven restored authority MUST fail closed. Next safe action MUST remain a projection derived from current state + current authority, never a grant.

Primary requirement allocation: `REQ-032..045`.

### `MIS-M0-SOVEREIGN-CORE-001-CRIT-005` — Fresh-process continuity and recovery

After terminating all Aurora processes, a fresh process MUST recover canonical identities/current state and re-evaluate authority from durable state plus required local trusted inputs, without transcript/model/Harness/runtime-memory assistance. Missing/corrupt/ambiguous state MUST produce explicit degraded/failure classification rather than fabricated progress, and unsafe blind retry MUST be prevented.

Primary requirement allocation: `REQ-046..055`.

### `MIS-M0-SOVEREIGN-CORE-001-CRIT-006` — Export, restore and migration

The Core MUST produce a versioned portable logical export distinct from physical SQLite backup, verify integrity/compatibility before restore, detect identity collision, preserve stable semantics through supported migration, and ensure historical/fresh-machine restore cannot silently revive current permission.

Primary requirement allocation: `REQ-056..066`.

### `MIS-M0-SOVEREIGN-CORE-001-CRIT-007` — Security, sovereignty and secret hygiene

M0 data classes, local-first ownership, untrusted-content isolation, owner-authentication boundary, integrity controls and secret hygiene MUST satisfy the accepted Threat Model and ADR-0008. Ordinary logs/telemetry MUST NOT carry secrets or sensitive payload merely for correlation.

Primary requirement allocation: `REQ-067..076`.

### `MIS-M0-SOVEREIGN-CORE-001-CRIT-008` — Audit, evidence and telemetry separation

Material initialize/transition/recovery/export/restore/migration operations MUST produce attributable audit/evidence metadata with stable correlation to governing identities/revisions. Claim, Receipt, Evidence, Verdict, domain audit and optional telemetry MUST remain distinct; telemetry failure MUST NOT redefine canonical state.

Primary requirement allocation: `REQ-077..088`.

### `MIS-M0-SOVEREIGN-CORE-001-CRIT-009` — Reliability and fault containment

Material operations MUST have explicit deterministic verification, fault/recovery classification and bounded retry semantics. Ambiguous mutation/restore/migration MUST be reconciled before retry; component green status alone MUST NOT be presented as M0 completion.

Primary requirement allocation: `REQ-089..095`.

### `MIS-M0-SOVEREIGN-CORE-001-CRIT-010` — Accepted architecture guards

Implementation MUST conform to accepted ADR-0003..0008 and the intentional M0 non-selections: Go Sovereign Core; one local modular process; SQLite/`database/sql`/`modernc.org/sqlite`; current-state-first revisions; portable JSON Schema/JSON/JCS boundary; OTel+slog non-authoritative observability; accepted owner-root/trust protocol; no full event sourcing, durable workflow engine, Mastra runtime, AHDK, M1 memory, M2 registry or distributed platform introduced merely for M0.

Primary requirement allocation: `REQ-096..107`.

### `MIS-M0-SOVEREIGN-CORE-001-CRIT-011` — Documentation, traceability and readiness integrity

Every durable implementation/evidence claim MUST trace through current Blueprint/requirement/Spec/ADR/Contract authority; fixed revisions, limitations, documentation impact and non-goals MUST remain discoverable. Implementation MAY NOT silently advance R6/R7/R8 or future roadmap scope.

Primary requirement allocation: `REQ-108..122`.

### `MIS-M0-SOVEREIGN-CORE-001-CRIT-012` — Complete M0 Golden Proof

The integrated implementation MUST pass the full operator-visible Golden Proof end to end on a fixed revision and produce sufficient evidence for independent R7 review and later R8 operator verdict.

Primary requirement allocation: `REQ-031` plus composition of `CRIT-001..011`.

## 5. Requirement allocation rule

`R5-COVERAGE.md` MUST contain exactly one primary Mission criterion for every `CAP-SOVEREIGN-CORE-REQ-001..122`.

The existing `R3-COVERAGE.md` remains the detailed requirement→Spec mechanism→planned test owner. R5 does not fork or duplicate that test catalog.

No requirement may be silently deferred from this M0 Mission. If an accepted requirement is discovered to be non-implementable or incorrectly scoped, the Mission MUST enter replan rather than omit it.

## 6. Scope

### Included

- one Aurora identity and one logical owner/operator identity reference;
- one or more Projects sufficient to prove stable identity/collision behavior;
- revisioned Project accepted-state envelope/current pointer;
- transition attempt lifecycle;
- minimum M0 authority grant/state/snapshot semantics;
- next-safe-action projection;
- local owner bootstrap/unlock and accepted Owner Root/trust-high-water semantics;
- local durable operational state;
- logical export + restore + explicit migration path;
- audit/evidence metadata sufficient for criteria;
- minimal operator proof/control surface;
- deterministic tests, fault injection and security cases required by the accepted Test Plan;
- documentation/reference updates required to keep repository truth aligned.

### Explicit non-goals

- LLM/model dependency or model routing;
- Mastra runtime usage in M0;
- conversational/project memory or Context Builder;
- Capability Registry, provider trust, AHDK or MNFS integration;
- Mission/Delegation runtime beyond this development Contract artifact;
- external Effect Gateway/PDP/Credential Broker;
- cloud service dependency;
- multi-node synchronization;
- multi-Presence, voice, vision or ambient sensing;
- devices/laboratory control;
- multi-tenancy;
- generic workflow/event/authorization platform;
- long-running timers/campaigns/budgets;
- self-improvement runtime;
- production deployment or customer-facing service.

## 7. Accepted architecture bindings

This Mission MUST conform to the currently accepted R4 decisions:

```text
Core runtime:        Go
Core topology:       one local modular Sovereign Core process
state model:         explicit current governing state + immutable accepted revisions
full Event Sourcing: not selected for M0
durable engine:      not selected for M0
operational store:   SQLite via database/sql + modernc.org/sqlite
SQLite posture:      WAL + synchronous=FULL; foreign key enforcement enabled where schema uses it
portable state:      JSON Schema 2020-12 + JSON logical form + RFC 8785 JCS where deterministic bytes are required
content digest:      SHA-256
sensitive export:    age-protected outer envelope for normal SENSITIVE portability use
observability:       OpenTelemetry traces/metrics + Go slog; exporter/backend optional
owner root:          random 256-bit ORK
owner unlock:        Argon2id-derived 32-byte KEK
ORK at rest:         AES-256-GCM wrapped
purpose subkeys:     HKDF-SHA-256
state/trust auth:    HMAC-SHA-256
trust anchor:        authenticated external generation + observed-wall-time high-water
historical restore:  REVALIDATION_REQUIRED
```

ADR-0009 is cross-horizon only. It creates no Mastra dependency in this Mission.

## 8. Environment and reproducibility baseline

### Canonical implementation/evidence class

```text
local single-user developer-controlled environment
primary CI/reference OS: Ubuntu 24.04 amd64
required storage/trust compatibility evidence: Windows amd64 runner class
network/cloud runtime dependency: none
CGO requirement for selected SQLite binding: none
```

### Evidence-qualified dependency baseline from R4

```text
Go:                     1.26.5
modernc.org/sqlite:     v1.54.0
modernc.org/libc:       v1.74.1 compatible pin
observed SQLite:        3.53.3
golang.org/x/crypto:   v0.54.0
CGO_ENABLED:            0
```

These versions are the R4 evidence baseline, not permission to use stale dependencies indefinitely. R6 MUST revalidate the exact implementation pins. A material version change affecting semantics, durability, crypto or cross-platform behavior triggers Contract replan or renewed evidence; an evidence-neutral patch may be accepted through the later authorized implementation-design process.

## 9. Contract-level decomposition

The Mission may be decomposed at R6 into small implementation units, but the contract-level workstreams are fixed as:

1. **Bootstrap and proof surface** — initialize/unlock/inspect without making the adapter authoritative.
2. **Identity + Project state** — durable identity, accepted revisions, current pointer and projections.
3. **Transition + authority** — revision-bound mutations, grant/state evaluation, invalid transition containment.
4. **Owner trust** — ORK custody, authenticated trust high-water, time/rollback classification and explicit recovery paths.
5. **Portability** — logical export, protected envelope, restore, collision handling and migration.
6. **Audit/evidence/observability** — required attributable records + optional non-authoritative telemetry.
7. **Integrated proof** — fresh-process, negative/security/fault journeys and complete M0 Golden Proof.

R6 owns file/module/interface/task decomposition. No contract workstream authorizes independent service extraction.

## 10. Dependencies

Required inputs:

- accepted M0 Product Milestone and Golden Proof;
- R0–R4 `PASS`;
- accepted ADR-0003..0008;
- current R2/R3 Capability package;
- reviewed SPK-001/SPK-002 evidence and limitations;
- documentation validator/generator.

No runtime dependency on Mastra, MNFS, model providers, cloud services, message brokers or durable workflow engines is permitted.

## 11. Authority and prospective implementation envelope

This Contract defines the bounds of later implementation but does **not** grant execution authority by itself.

If R6/R7 are separately authorized, work MAY be limited to:

- create/modify CAP-SOVEREIGN-CORE M0 source, tests, schemas/config and supporting reference/docs inside this repository;
- add only dependencies necessary to realize accepted ADRs and this Contract;
- run local/CI builds, deterministic tests, fault injection and security verification;
- create disposable fixtures and evidence artifacts;
- use isolated branches/worktrees and GitHub Actions for controlled verification.

Still prohibited unless separately authorized:

- production deployment;
- external communication or financial/legal effects;
- access to unrelated personal/company data;
- secrets in repository/logs/evidence;
- modifying constitutional product intent to make tests pass;
- widening authority, threat model or data scope;
- implementing future M1+ systems;
- promoting Architecture Spike code as production by copy/merge;
- auto-advancing R6/R7/R8.

## 12. Resource and complexity budget

M0's budget is primarily **scope/complexity**, not an arbitrary performance target:

```text
Core process class:      1 local modular process
operational DB class:    1 embedded SQLite store
local trust domains:     operational DB + small external Owner Trust Store/root envelope
runtime cloud services:  0 required
runtime model calls:     0 required
message brokers:         0
durable workflow engine: 0
microservices required:  0
Mastra/AHDK required:    0
```

No arbitrary latency/SLO number is introduced at R5. R6 MUST define deterministic test timeouts and fixture bounds; R7 records actual measurements and failures. Material resource pressure that would force an architecture change triggers replan.

## 13. Evidence contract

Every criterion must receive fixed-revision evidence appropriate to its R3 test allocation.

Minimum R7 evidence bundle must contain:

- exact source commit and dependency/toolchain lock;
- build/static/schema verification receipts;
- unit/contract/integration results as allocated;
- real fresh-process restart evidence;
- fault-injection evidence around accepted-state/store and trust-anchor boundaries;
- authority/security negative cases;
- export/restore/migration fixtures and integrity results;
- cross-platform storage/trust evidence on Ubuntu + Windows runner classes;
- proof that telemetry/exporter absence does not become canonical-state loss;
- secret/redaction checks;
- complete requirement→criterion→test/evidence coverage;
- unresolved limitation/residual-risk report;
- independent/adversarial review;
- complete M0 Golden Proof receipt.

A command exit code, implementer claim or local component green status alone is insufficient.

## 14. Acceptance thresholds

The Mission cannot receive a successful R7 verdict unless all are true:

1. `122/122` in-scope requirements have criterion and evidence disposition.
2. Zero accepted invalid/stale/unauthorized transition changes governing state.
3. Zero expired/revoked/unproven-restored authority is treated as current permission.
4. Zero silent Aurora/Project identity replacement or restore collision acceptance.
5. Missing/corrupt/unverifiable canonical state or trust fails explicitly closed.
6. DB/anchor anomaly states are enforced at the mutation boundary, not merely diagnosed.
7. Secret-hygiene checks pass for repository/log/evidence surfaces used by the Mission.
8. Export/restore/migration preserve required identities/semantics or fail explicitly.
9. Required audit/evidence relationships are present while optional telemetry remains non-authoritative.
10. Full M0 Golden Proof passes end to end on one fixed revision.
11. No unresolved critical/high finding contradicts a Contract criterion.
12. Documentation validation passes and current state is promoted to canonical repository memory.

## 15. Risks and accepted limitations

### Local trust boundary

M0 does not claim protection from an attacker who can replay all local trust artifacts while also controlling owner secrets/runtime. This residual must remain visible.

### Physical power-loss scope

R4 process-kill evidence does not prove every storage-controller/write-cache/power-loss failure. R6/R7 must explicitly validate the supported target's filesystem publication/fsync semantics to the level actually claimed.

### KDF envelope resource use

Production decoding must allowlist/bound supported Argon2id parameters before allocation.

### Dependency evolution

The exact versions observed in R4 are evidence-qualified starting points; they require implementation-baseline revalidation.

### Proof-surface simplicity

The operator adapter is intentionally a proof/control surface, not Presence Fabric or a product UI architecture commitment.

## 16. Change and replan triggers

This Mission MUST stop and create a finding/replan before continuing if execution would require any of the following:

- change to accepted M0 outcome/Golden Proof;
- change or weakening of an accepted `CAP-SOVEREIGN-CORE-REQ-*` requirement;
- material change to accepted ADR-0003..0008;
- new runtime cloud/model/Harness dependency;
- replacement of SQLite/store class, current-state model or owner-root/trust protocol;
- new external effect/credential/provider boundary;
- multi-user/multi-node semantics;
- inability to meet a critical/high criterion without scope expansion;
- target filesystem behavior incompatible with required durability semantics;
- security finding that invalidates the accepted local threat boundary;
- contract criterion discovered to be untestable or contradictory.

Replan produces a new Contract revision or explicit supersession. An implementation plan cannot silently rewrite this Contract.

## 17. Revision and supersession

Before operator approval, this document is `proposed` and may be revised through R5 review.

After approval:

```text
approved contract revision/hash becomes immutable authority for that Mission scope
material change → proposed new revision
→ operator review when required
→ supersede or reject
```

Implementation/evidence must always identify the exact approved Contract revision.

## 18. R5/R6 boundary

This Contract deliberately does **not** define:

- source directory/package names;
- Go interfaces/types/functions;
- exact SQL DDL/migration files;
- CLI command syntax;
- filesystem path names;
- exact crypto wrapper APIs;
- test framework/helpers;
- commit sequence;
- detailed fault-injection harness implementation.

Those are R6 Microdesign/Implementation Plan responsibilities constrained by this Contract, the accepted Capability package and ADRs.
