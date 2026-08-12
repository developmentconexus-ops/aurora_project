---
id: PLAN-AURORA-TECHNICAL-ARCHITECTURE-BASELINE
title: Aurora Technical Architecture Baseline Plan
document_type: implementation_plan
form: reference
authority: design
status: active
version: 0.1.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - ordered documentary execution plan for the Aurora Technical Architecture Baseline
  - review and stop boundaries for TA-01 through TA-08
related:
  - DESIGN-AURORA-TECHNICAL-ARCHITECTURE-BASELINE-MAP
  - DOC-AURORA-TECHNICAL-ARCHITECTURE-MAP-OPERATOR-ACCEPTANCE
  - DESIGN-AURORA-SYSTEM-ARCHITECTURE-DECISION-LANDSCAPE
  - DOC-AURORA-STATUS
last_reviewed: 2026-08-12
---

# Aurora Technical Architecture Baseline Plan

> **For agentic workers:** execute this plan as documentary architecture work. Do not implement Aurora runtime code, create production repositories, run Architecture Spikes, or promote technology choices without the authority defined in `STATUS.md`.

**Goal:** Produce a coherent, evidence-bounded technical architecture baseline that lets future Aurora capabilities be designed and implemented without inventing module ownership, runtime topology, repository structure, contracts, data roles, security boundaries or operational conventions.

**Architecture:** Work proceeds in dependency order from logical ownership to runtime placement, source/repository boundaries, contracts, data, security, cognitive integration and operations. Each area produces an operator-reviewable artifact and explicit `DECIDE / RESEARCH / SPIKE / DEFER` dispositions. Scoped accepted ADRs remain scoped; no universal stack is inferred.

**Tech Stack:** Markdown, YAML frontmatter, architecture diagrams, responsibility/ownership matrices, current primary-source research manifests where material, ADRs/Specifications/Standards only when decisions are promoted.

## Global constraints

- Product Blueprint remains the owner of product meaning.
- `DESIGN-AURORA-TECHNICAL-ARCHITECTURE-BASELINE-MAP` owns the work sequence and materiality rule.
- The frozen M0 R7 candidate is evidence input only.
- No new Aurora runtime implementation.
- No Architecture Spike execution without separate explicit authorization.
- No product/framework/database/protocol selection before the owning architecture question is defined.
- Do not create a second readiness lifecycle or score.
- Defer detail that does not change the next structural or implementation decision.
- Every material output records owner, consumer, evidence needed, migration impact and stop condition.
- Each work package ends with operator/reviewer checkpoint before the next package changes its assumptions.

---

## Work package status

| Work package | State | Current authorization |
|---|---|---|
| TA-01 Logical modules and canonical ownership | ACTIVE | discovery/design authorized |
| TA-02 Process, runtime and evolutionary topology | ACTIVE with TA-01 | discovery/design authorized |
| TA-03 Repository, source and build architecture | QUEUED | may not finalize before TA-01/TA-02 review |
| TA-04 Contracts, APIs, events and communication | QUEUED | may inventory only; no binding selection |
| TA-05 Data, storage, portability and lifecycle | QUEUED | may inventory only; no new store selection |
| TA-06 Identity, authentication, authorization, policy and secrets | QUEUED | may inventory only; no product selection |
| TA-07 Brain, models, memory and Harness integration | QUEUED | may inventory only; no implementation |
| TA-08 Configuration, observability, deployment and operation | QUEUED | may inventory only; no backend/deployment selection |

---

## Task 1 — Fix the technical architecture source baseline

**Files:**
- Read: `AGENTS.md`
- Read: `docs/tracking/STATUS.md`
- Read: `docs/design/AURORA-TECHNICAL-ARCHITECTURE-BASELINE-MAP.md`
- Read: `docs/design/AURORA-SYSTEM-ARCHITECTURE-DECISION-LANDSCAPE.md`
- Read: relevant Product Blueprint and accepted ADRs
- Inspect only as needed: `feat/m0-r7-sovereign-core-20260810`

- [ ] Record the exact canonical `main` revision and current work-branch revision.
- [ ] Confirm M0 R7 remains frozen/non-canonical.
- [ ] List accepted cross-system constraints and scoped technical decisions.
- [ ] List current open questions that can change TA-01/TA-02.
- [ ] Mark Presence/session details that do not affect module/runtime topology as `DEFER`.
- [ ] Stop if a material conflict exists between Blueprint, ADR, accepted design and current STATUS.

**Deliverable:** fixed-source preflight section in the TA-01/TA-02 design artifact.

---

## Task 2 — Build the TA-01 candidate component catalog

**Primary deliverable:** `docs/design/AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY.md`

- [ ] Derive candidate components from Product Blueprint 03, 05, 06, 07, 10, 11, 12 and 13.
- [ ] Include M0's actual logical owners as evidence, not universal truth.
- [ ] Define one-sentence responsibility for each candidate.
- [ ] Define the canonical entities/data each candidate may own.
- [ ] Define what each candidate explicitly must not own.
- [ ] Merge candidates that lack an independent invariant or consumer.
- [ ] Split candidates only when responsibility, security, lifecycle or runtime evidence requires it.
- [ ] Identify components that are future-only and mark them `DEFER` rather than making empty modules.

**Required matrix columns:**

```text
component
purpose
canonical responsibilities
owned state/data
consumed inputs
produced outputs/contracts
allowed dependencies
forbidden dependencies
first consumer
Stage A disposition
Stage B evolution
```

**Review test:** each global concept has one owner; no candidate is merely a renamed framework feature.

---

## Task 3 — Compare coherent TA-01/TA-02 topology approaches

Produce 2–3 complete alternatives, not isolated component debates.

### Required approaches to compare

```text
Approach A — single deployable modular Core with on-demand provider processes
Approach B — minimum sovereign service + separate Presence and Cognitive Runtime processes
Approach C — staged topology: one deployable initially, process seams fixed by ports and split triggers
```

The names may be revised after analysis, but comparison must cover equivalent concerns.

For each approach:

- [ ] draw Stage A runtime topology;
- [ ] identify always-active and on-demand processes;
- [ ] place Go and TypeScript/Mastra responsibilities;
- [ ] identify trust/failure boundaries;
- [ ] identify local communication crossings;
- [ ] describe startup/restart/recovery ownership;
- [ ] describe resource/operational burden;
- [ ] describe Stage B migration;
- [ ] identify lock-in and rollback;
- [ ] test against accepted Stage A Presence constraints;
- [ ] test against first likely consumers M1 and M2 without implementing them.

**Decision rule:** choose the smallest topology that preserves required ownership and real failure/security/runtime boundaries.

---

## Task 4 — Produce canonical ownership and dependency rules

- [ ] Create entity/state ownership matrix.
- [ ] Distinguish `read`, `propose`, `validate`, `commit` and `project` permissions.
- [ ] Define dependency direction between domain, application, ports/adapters and provider runtimes.
- [ ] Prohibit direct provider/Harness writes to canonical state.
- [ ] Prohibit model or memory-provider state from becoming authority by convenience.
- [ ] Define how artifacts/evidence reference large content without cross-module ownership leakage.
- [ ] Define which components may emit domain events and which may only emit telemetry.
- [ ] Identify cross-module transactions that are allowed, forbidden or require an application coordinator.

**Evidence:** adversarial duplicate-owner review with representative Project, Mission, Authority, Memory, Provider, Presence, Artifact and Evidence flows.

---

## Task 5 — Produce Stage A and Stage B runtime maps

### Stage A

- [ ] map workstation-hosted processes/deployables;
- [ ] map persistent minimum responsibilities;
- [ ] map on-demand cognition/model/Harness work;
- [ ] map local state and artifact access;
- [ ] map local Presence activation without expanding session micro-policy;
- [ ] map failure/restart cases;
- [ ] identify process boundaries deferred until measured evidence.

### Stage B

- [ ] move sovereign hosting to a persistent personal node hypothesis;
- [ ] convert workstation into a Presence without changing product semantics;
- [ ] identify new network/authentication boundaries;
- [ ] preserve canonical identity and authority ownership;
- [ ] identify data that should remain Presence-local before minimization;
- [ ] list migration/portability dependencies for TA-05/TA-06.

**Stop condition:** Stage B is an evolution path, not a second architecture designed in full.

---

## Task 6 — Classify TA-01/TA-02 open decisions

For every unresolved question, record:

```text
question
current hypothesis
dependency
first consumer
risk/lock-in
DECIDE | RESEARCH | SPIKE | DEFER
evidence needed
owning future artifact
reconsideration trigger
```

- [ ] Separate structural decisions from replaceable mechanisms.
- [ ] Identify any current primary-source research needed.
- [ ] Specify Spikes only for runtime properties that documentation cannot establish.
- [ ] Do not execute Spikes.
- [ ] Explicitly hand inputs to TA-03 repository architecture.

---

## Task 7 — Review and accept TA-01/TA-02

- [ ] Run documentation generation and validation.
- [ ] Run `git diff --check` through CI.
- [ ] Perform an adversarial review against the Technical Architecture Map.
- [ ] Test for duplicate owners, framework capture, premature services and hidden stack selection.
- [ ] Test that Stage A can remain operationally small.
- [ ] Test that Stage B does not require domain identity redesign.
- [ ] Resolve blocking/material findings.
- [ ] Present one fixed revision to the operator.
- [ ] Stop before TA-03 finalization unless separately approved by the operator/current STATUS.

---

# Queued work after TA-01/TA-02

## Task 8 — TA-03 Repository, source and build architecture

- [ ] derive source/release boundaries from accepted runtime topology;
- [ ] compare monorepo, polyrepo and staged strategies;
- [ ] define contract/schema ownership and generated types;
- [ ] define Go/TypeScript workspace boundaries;
- [ ] define CI, versioning, release and supply-chain rules;
- [ ] produce decision/review package.

Do not choose repository strategy merely from team fashion or tool preference.

---

## Task 9 — TA-04 Contracts, APIs, events and communication

- [ ] catalog contract families and process crossings;
- [ ] define errors, cancellation, deadlines, idempotency and correlation;
- [ ] distinguish events, messages, audit and telemetry;
- [ ] compare current binding candidates per boundary;
- [ ] define compatibility/deprecation rules;
- [ ] produce ADR/Standard proposals only for material current bindings.

---

## Task 10 — TA-05 Data, storage, portability and lifecycle

- [ ] inventory all logical data families;
- [ ] classify canonical, derived, ephemeral and provider-local data;
- [ ] define owner, query, consistency, retention, deletion and recovery needs;
- [ ] define portability/rebuildability classes;
- [ ] compare physical store candidates only after role definitions;
- [ ] specify benchmark/fault evidence where needed.

---

## Task 11 — TA-06 Identity, authentication, authorization, policy and secrets

- [ ] define identity/actor classes and trust chain;
- [ ] define authentication requirements per actor/topology;
- [ ] define authority and policy contracts;
- [ ] define effect enforcement and receipts;
- [ ] define secrets/credential-reference flows;
- [ ] compare products only after the model exists;
- [ ] produce scoped ADR/Specification proposals.

Do not return to highly specific workstation-user behavior unless it changes the current actor/topology design.

---

## Task 12 — TA-07 Brain, models, memory and Harness integration

- [ ] define deterministic versus cognitive responsibilities;
- [ ] define Context Builder and memory boundaries;
- [ ] define model role/provider contracts;
- [ ] define tool/Harness delegation and effect flow;
- [ ] define global versus provider-local state reconciliation;
- [ ] evaluate Mastra at the first real consuming boundary;
- [ ] define failure/restart/cancellation behavior;
- [ ] produce integrated cognitive architecture proposal.

---

## Task 13 — TA-08 Configuration, observability, deployment and operation

- [ ] define config/environment/secret precedence;
- [ ] define health, startup, recovery and migration behavior;
- [ ] define logging/tracing/metrics/audit/evidence semantics;
- [ ] define packaging, install, update and rollback;
- [ ] define Windows/Linux and Stage A/B operational posture;
- [ ] define CI/CD and artifact provenance;
- [ ] compare operational tools/backends only after requirements.

---

## Task 14 — Integrated Technical Architecture Baseline review

- [ ] assemble the integrated baseline deliverables listed in the accepted map;
- [ ] verify no concept has duplicate canonical ownership;
- [ ] verify every process crossing has a contract/failure classification;
- [ ] verify every store has a logical role and migration path;
- [ ] verify identity/authority/effect paths are complete;
- [ ] verify model/Harness loss cannot erase sovereign truth;
- [ ] verify operations can recover the system;
- [ ] identify the first bounded implementation horizon;
- [ ] re-run the appropriate ACRM readiness path before any implementation;
- [ ] obtain explicit operator acceptance.

---

## Fresh-session checkpoint

A new session must not restart product discovery or continue Presence micro-policy by default.

It must state:

```text
current program: Technical Architecture Baseline
current tranche: TA-01 + TA-02
current output: module ownership + process/runtime topology
current method: boundary → requirements → alternatives → evidence → decision
current implementation authority: none
next action: compare coherent component/runtime approaches
```
