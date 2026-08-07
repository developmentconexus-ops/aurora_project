---
id: ADR-AURORA-0004
title: M0 Local State and Execution Shape
document_type: adr
form: explanation
authority: decision
status: proposed
version: 0.2.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - proposed M0 process topology, canonical-state persistence pattern and durable-engine applicability
related:
  - DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC
  - DESIGN-AURORA-M0-R4-DECISION-LANDSCAPE
  - RESEARCH-AURORA-M0-RUNTIME-PERSISTENCE-R4-V1
  - RESEARCH-AURORA-MASTRA-COGNITIVE-HARNESS-R4-V1
  - ADR-AURORA-0009
supersedes: []
superseded_by: null
last_reviewed: 2026-08-07
---

# ADR-0004 — M0 Local State and Execution Shape

## Context

R3 left topology, state-versus-event persistence and durable-workflow-engine applicability open. M0 is one local, single-user walking skeleton with short internal state operations, not a long-running Mission engine.

The later Mastra-focused R4 research exposed a rich future runtime for workflows, durable agents, background tasks and provider-local persistence. That finding does not change M0's scope: no Mastra process or other Harness/runtime is required to prove Sovereign Core continuity.

## Decision drivers

- current state must remain canonical and process-independent;
- events/audit cannot become sole state authority;
- logical modularity precedes physical distribution;
- partial accepted-state/audit updates are forbidden;
- M0 has no durable timers, provider lifecycle or external-effect workflow;
- near-horizon evolution must not require domain rewrite;
- future Mastra workflow/durable state must remain provider-local rather than becoming Core truth.

Affected requirements include `REQ-010..020`, `024..031`, `046..055`, `077..088`, `096`, `097`, `100`, `102`, `105`, `107`.

## Options

1. modular local Core + explicit current relational state/revisions + separate audit records;
2. full event sourcing with derived current projections;
3. durable workflow engine as primary state/history owner;
4. distributed/microservice Core from M0.

## Decision

**Proposed:** M0 uses **one local modular Sovereign Core process** with:

```text
explicit current governing state
+ immutable accepted state revisions
+ logically distinct audit/domain-event records
+ one transactional mutation boundary
```

Full event sourcing is **not selected** for M0.

Audit/event records MAY be physically co-located in the same relational operational store to share the mutation transaction, while remaining logically distinct and non-authoritative for current state.

A durable workflow engine is **NOT_YET_A_DECISION** for M0 and MUST NOT be introduced. Reconsider at M4 or earlier only if an accepted current requirement cannot safely be implemented as a bounded transaction/recovery operation.

Mastra Workflows, Durable Agents, Background Tasks or a Temporal-backed Mastra runtime MAY later implement Harness-local durability or an adapter behind Aurora's `DurableExecutionPort`. Their snapshots/run history MUST NOT automatically become canonical Project/Mission/Authority state.

Physical service distribution is not selected for M0. A later separate Mastra service is compatible with this ADR because it belongs to a provider/Harness boundary, not to the M0 Core topology.

## Consequences

### Positive

- direct alignment with R3 current-state semantics;
- fewer competing durable histories in M0;
- atomic state/audit boundary is tractable;
- lower operations burden;
- easier migration to a future networked store behind the Durable State Port;
- Mastra can be adopted later without rewriting M0 state ownership.

### Negative

- M0 does not gain workflow-engine timers/signals/replay;
- later M4/Mastra-backed capabilities may introduce provider-local execution histories that require explicit reconciliation with Aurora global state.

### Risks

Application code could accidentally treat audit history or future Mastra workflow/session history as replay authority. Static architecture tests, explicit store APIs and provider contracts must preserve the distinction.

## Compatibility / migration / rollback

A future durable engine or Mastra runtime may be added behind its own port/provider contract without changing Project/Authority state ownership. A future store may migrate revisions/current pointers/audit records through explicit logical migrations.

Loss or corruption of provider-local Mastra state may degrade/fail a provider run but cannot erase Aurora identity, canonical Project state or Authority.

## Validation

Documentary evidence is sufficient for the M0 topology/state/engine non-selection. Exact store atomicity remains blocked by `SPK-AURORA-M0-SOVEREIGN-STORE-001`.

No Mastra integration spike is required for M0 because M0 does not consume Mastra. The first Mastra-backed Capability must prove its provider boundary when it reaches its own implementation horizon.

## Reconsideration triggers

- multi-process concurrent Core writers become current scope;
- long-running waits/timers/external effects enter an accepted milestone;
- SQLite/store spike shows the co-located transactional boundary cannot satisfy M0 safely;
- M4 enters execution horizon;
- first Mastra-backed Capability requires durable-run reconciliation semantics.

## References

- `RESEARCH-AURORA-M0-RUNTIME-PERSISTENCE-R4-V1`
- `RESEARCH-AURORA-MASTRA-COGNITIVE-HARNESS-R4-V1`
- ADR-0009
- `CAP-SOVEREIGN-CORE/SPEC.md`
