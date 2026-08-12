---
id: DOC-AURORA-STATUS
title: Aurora Project Status
document_type: project_status
form: reference
authority: tracking
status: current
version: 0.33.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - current Aurora project phase
  - current authorization boundary
  - current blockers and immediate next action
related:
  - DOC-AURORA-DOCUMENTATION-MAP
  - DOC-AURORA-PRODUCT-INDEX
  - DOC-AURORA-ROADMAP
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DESIGN-AURORA-SYSTEM-ARCHITECTURE-REBASELINE
  - DESIGN-AURORA-SYSTEM-ARCHITECTURE-DECISION-LANDSCAPE
  - DESIGN-AURORA-TECHNICAL-ARCHITECTURE-BASELINE-MAP
  - PLAN-AURORA-TECHNICAL-ARCHITECTURE-BASELINE
  - DESIGN-AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY
  - DOC-AURORA-DECISIONS
  - DOC-AURORA-WORKLOG
last_reviewed: 2026-08-12
---

# Aurora Project Status

## 1. Current summary

- **Canonical branch:** `main`
- **Canonical revision at TA-01/TA-02 start:** `9cbce1efe4f742f90623b894c0c1ba2eaa3cebcc`
- **Current documentary branch:** `docs/ta-01-02-module-runtime-topology-20260812`
- **A0:** ACCEPTED / MERGED
- **ADR-0001..0009:** ACCEPTED within exact scope
- **M0 R0–R6:** PASS
- **M0 R7 candidate:** FROZEN / PRESERVED / NON-CANONICAL
- **M0 R7 Verdict:** NOT ISSUED
- **M0 R8:** NOT AUTHORIZED / NOT PERFORMED
- **Technical Architecture Baseline map:** ACCEPTED / MERGED
- **Current tranche:** `TA-01 + TA-02`
- **Current tranche state:** PROPOSED DESIGN PRODUCED / REVIEW IN PROGRESS
- **Aurora implementation:** PAUSED

## 2. Current objective

The current work converts accepted Aurora product architecture into a coherent technical structure by deciding:

```text
TA-01
→ minimum technical component set
→ canonical responsibility/data ownership
→ allowed and forbidden dependencies

TA-02
→ Stage A deployables/processes/runtimes
→ always-active versus on-demand placement
→ failure and restart ownership
→ bounded Stage B evolution
```

The proposed owner is:

```text
docs/design/AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY.md
```

It remains `status: proposed` until operator acceptance.

## 3. Proposed architecture under review

The proposal compares three complete approaches:

```text
A — Core-centric single application
B — early service decomposition
C — Evolutionary Sovereign Host with one early provider-runtime seam
```

Current recommendation:

```text
Approach C

one persistent Go Aurora Sovereign Host
→ multiple logically independent Aurora-owned domain modules
→ application coordination and canonical state/recovery
→ thin local Presence/activation adapter initially co-packaged

one independently replaceable cognitive/provider runtime boundary
→ starts when first consumed
→ Mastra/TypeScript remains preferred-first to evaluate
→ provider-local state only

specialized Harnesses remain independent providers

all other process splits require a named lifecycle, runtime,
security, fault, resource, remote-node or consumer trigger
```

The recommendation does not select the exact local transport, repository strategy, store, authentication product, service manager or Mastra implementation.

## 4. Proposed canonical domain owners

The design proposes:

1. Identity and Relationship;
2. Project and World State;
3. Mission and Delegation Control;
4. Authority and Policy;
5. Capability and Provider Registry;
6. Memory, Knowledge and Context;
7. Artifact and Evidence;
8. Presence and Interaction;
9. Attention and Proactivity — future/deferred;
10. Failure Intelligence and Evaluation — future/deferred.

It also separates application coordinators, deployables, enforcement boundaries, provider runtimes and mechanism adapters from canonical domain ownership.

A provider, model, Harness, Presence, memory index, telemetry backend or store adapter cannot write canonical Aurora state directly.

## 5. Validation evidence

Initial design commit:

```text
2192f182c24eed0c8406062957e035a0db6e88a5
```

Documentation workflow:

```text
run: 31623255539
result: SUCCESS
```

The green workflow proves documentary structural consistency only. Adversarial architecture review and operator acceptance remain required.

## 6. Current authorized work

Authorized:

- inspect accepted Blueprint, ADR, research, M0 and frozen-R7 evidence relevant to TA-01/TA-02;
- revise the proposed module/ownership/runtime design;
- perform duplicate-owner, framework-capture, God-process and service-sprawl review;
- compare Stage A approaches and Stage B evolution;
- classify unresolved questions as `DECIDE`, `RESEARCH`, `SPIKE` or `DEFER`;
- prepare a draft PR and fixed review revision;
- use current primary sources only if they change a near TA-01/TA-02 decision;
- specify but not execute an Architecture Spike;
- present the proposal for operator review.

## 7. Explicitly prohibited

Not authorized:

- Aurora runtime implementation;
- modification, merge or promotion of the frozen M0 R7 candidate;
- an M0 R7 Verdict or M0 R8 closeout;
- M1+ implementation;
- Architecture Spike execution;
- TA-03 repository strategy finalization;
- production repository creation/restructuring;
- choosing monorepo/polyrepo;
- selecting HTTP, gRPC, Connect, MCP, A2A, event transport or local IPC;
- selecting new databases/stores;
- selecting Keycloak, Zitadel, Authentik, Ory, SPIFFE, OPA, Cedar, Vault or equivalents;
- implementing AHDK, MNFS, Mastra adapter, Brain, memory, Voice, Presence, model routing or observability systems;
- treating Go/SQLite/JSON-JCS/OTel M0 choices as universal;
- creating one process/service per conceptual module;
- returning to Presence/session micro-policy unless structurally material.

## 8. Current architecture blockers

```text
TA-01/TA-02 OPERATOR ACCEPTANCE: NOT ISSUED
TA-03 FINALIZATION: BLOCKED
AURORA IMPLEMENTATION: BLOCKED
ARCHITECTURE SPIKE EXECUTION: NOT AUTHORIZED
```

Open review questions include:

- whether the proposed domain-owner set is minimal and complete;
- whether any global entity has duplicate or missing ownership;
- whether the Cognitive Runtime process boundary should be fixed structurally before the first M1 consumer;
- whether the thin local Presence may remain co-packaged in Stage A;
- whether the recommended Sovereign Host can remain modular without becoming a God process.

## 9. Immediate next action

```text
adversarially review the proposed TA-01/TA-02 design
→ remediate admitted findings
→ open/update one draft PR
→ validate the fixed revision
→ present Approach C and alternatives to the operator
→ ACCEPT | REVISE | REJECT
→ STOP before TA-03 finalization or implementation
```

No technology implementation follows automatically from accepting the topology.

## 10. Fresh-session read order

After `AGENTS.md`, a new TA-01/TA-02 session must read:

1. this `STATUS.md`;
2. `docs/DOCUMENTATION-MAP.md`;
3. `docs/product/README.md`;
4. `docs/roadmap.md`;
5. `docs/design/AURORA-TECHNICAL-ARCHITECTURE-BASELINE-MAP.md`;
6. `docs/superpowers/plans/2026-08-12-aurora-technical-architecture-baseline.md`;
7. `docs/design/AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY.md`;
8. `docs/design/AURORA-SYSTEM-ARCHITECTURE-DECISION-LANDSCAPE.md`;
9. relevant Product Blueprint sections and accepted ADRs;
10. the frozen R7 candidate only as evidence when needed.

The session must not restart broad product discovery, choose products before boundaries, or implement code.
