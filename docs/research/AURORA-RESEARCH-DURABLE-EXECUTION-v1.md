---
id: RESEARCH-AURORA-DURABLE-EXECUTION-V1
title: Aurora Research — Durable Execution, Recovery and Long-Running Missions
document_type: research_report
form: explanation
authority: research
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - durable execution research through 2026-08-05
related:
  - DOC-AURORA-BLUEPRINT-04
  - DOC-AURORA-BLUEPRINT-07
  - DOC-AURORA-BLUEPRINT-10
  - DOC-AURORA-BLUEPRINT-12
source_manifest: AURORA-RESEARCH-DURABLE-EXECUTION-v1.sources.json
review_triggers:
  - M4 readiness
  - SPK-004 result
  - selected engine major release
last_reviewed: 2026-08-05
---

# Aurora Research — Durable Execution, Recovery and Long-Running Missions

## 1. Research question

Which architecture can let Aurora run hours- or days-long missions, wait for human decisions, survive process/network failure, preserve budgets and prevent duplicate external effects—while remaining local-first and independent of one agent framework?

Candidates reviewed:

- Temporal;
- DBOS;
- Restate;
- Inngest;
- LangGraph checkpointing/durable execution;
- Mastra integration with Temporal;
- a minimal Aurora-owned local baseline.

---

## 2. Executive finding

Durable execution is a layer distinct from agent reasoning and Harness runtime.

```text
Agent/Harness runtime
→ decides local next work

Durable execution engine
→ persists progress, timers, waits, retries and recovery

Aurora Operational State
→ owns global Mission/Delegation semantics, authority and acceptance
```

No engine should become Aurora's domain model.

The initial architecture should define a `Durable Execution Port` and run a comparative spike rather than selecting the most feature-rich engine immediately. Temporal offers the most established workflow-history model but has higher operational/conceptual weight. DBOS offers a PostgreSQL-centered durable model and potentially simpler local operation. Restate combines durable state, journal, messaging and workflows. Inngest emphasizes event-driven step-based functions and managed/serverless operation. LangGraph and Mastra provide agent workflow durability, but their state is still framework-specific and may rely on external durable engines for stronger guarantees.

---

## 3. Aurora requirements

A candidate must address:

### Durable state

- active Mission/Delegation;
- phase/checkpoint;
- exact contract/build;
- pending signal/decision;
- budget;
- completed effects/artifacts;
- cancellation/revocation.

### Failure

- Core crash;
- Harness crash;
- worker crash;
- network partition;
- duplicate event;
- timeout;
- stale provider;
- incompatible code version;
- storage restart.

### Human-in-the-loop

- wait for approval/decision;
- preserve context;
- deadline/default;
- resume exact state;
- no busy process.

### Effects

- idempotency;
- at-least-once execution reality;
- activity/step boundaries;
- receipts;
- compensation/reconciliation;
- no blind retry.

### Single-user/local-first

- reasonable installation and operations;
- local development;
- backup/restore;
- no mandatory large cluster;
- future remote workers.

### Versioning

- workflow code changes;
- in-flight runs;
- provider/AHDK contract version;
- deterministic replay constraints;
- migration/rollback.

---

## 4. Temporal

Temporal uses durable Workflow Execution with persisted event history. Workflow code must remain deterministic under replay, while Activities perform side effects [S01].

### Strengths

- mature durable workflow semantics;
- timers and signals;
- long-running execution;
- retries/timeouts;
- workflow versioning patterns;
- strong operational visibility;
- multiple SDK languages;
- worker restart/resume;
- explicit Activities boundary.

### Fit for Aurora

- long-running campaigns;
- waiting for Leandro;
- multi-step Delegations;
- remote workers;
- robust retry/timeouts;
- durable orchestration port.

### Risks/costs

- operational services and learning curve;
- replay determinism constraints can leak into design;
- workflow history is not automatically Aurora domain state;
- incorrect activity idempotency still duplicates effects;
- may be excessive for first single-user milestones;
- language SDK choice influences implementation.

### Key implication

If selected, Aurora Mission/Delegation entities remain in Core. Temporal workflow IDs/history implement execution, not product authority.

---

## 5. DBOS

DBOS provides durable execution backed by PostgreSQL, recording workflow/step progress and supporting recovery after interruption [S02].

### Strengths

- PostgreSQL-centered architecture;
- potentially smaller operational footprint where PostgreSQL is already used;
- durable functions/workflows;
- transactions integrated with application state;
- recovery and queues.

### Fit

- local-first modular Core;
- one database for operational state and durable metadata may simplify M4;
- first-party code-centric workflows.

### Risks/questions

- maturity/ecosystem relative to Temporal;
- language/platform support;
- separation of Aurora domain tables and engine internals;
- scaling/remote worker behavior;
- long waits/signals/versioning details;
- effect idempotency still application responsibility.

DBOS deserves a spike if Aurora chooses PostgreSQL or values minimal infrastructure.

---

## 6. Restate

Restate provides durable execution with a journal, durable state, timers, messaging and service/workflow abstractions [S03].

### Strengths

- durable service calls and virtual objects;
- state plus communication;
- retries and exactly-once-oriented abstractions over durable invocation;
- support for long-running handlers/workflows;
- potentially good fit for provider/instance objects and signals.

### Fit

- cross-provider durable communication;
- per-entity state such as provider/device coordination;
- resilient orchestration without a full generic broker stack.

### Risks/questions

- operational/runtime maturity for Aurora's environment;
- conceptual coupling to Restate service/object model;
- language support;
- evidence for complex human/physical campaigns;
- how cancellation/version migration behaves;
- avoiding Restate state becoming canonical domain accidentally.

---

## 7. Inngest

Inngest implements event-driven durable functions using steps that checkpoint progress and resume on events/time [S04].

### Strengths

- developer-friendly event functions;
- step retries and state;
- waits and scheduling;
- strong cloud/serverless experience;
- observability and concurrency controls.

### Fit

- cloud-assisted event-driven capabilities;
- notifications and background jobs;
- future hosted workers.

### Risks/questions

- managed/cloud orientation may conflict with sovereign local-first Core;
- self-hosting/operational assumptions;
- domain/workflow versioning;
- physical/device local availability;
- external service dependency for critical operation.

Likely more relevant to cloud-assisted capability than first Core baseline unless self-hosted evidence is compelling.

---

## 8. LangGraph durability

LangGraph provides checkpoints, threads, persistence, interrupts, time travel and durable execution modes for graph-based agents [S05].

### Strengths

- agent-specific state graph;
- human-in-the-loop interrupts;
- checkpointed graph execution;
- replay/fork for debugging;
- Python/TypeScript ecosystems.

### Fit

- internal runtime of a Harness with adaptive graph;
- experimentation/review workflows;
- local stateful agent flow.

### Limits relative to Aurora

- graph state is framework-specific;
- Core Mission semantics should not become LangGraph state;
- cross-Harness durability and effect governance remain external;
- exact external side effects require idempotency and careful node boundaries;
- local persistence choice still matters.

LangGraph may use/implement a Harness local plan while Aurora's durable layer remains above or alongside it.

---

## 9. Mastra and Temporal

Mastra supports workflows, suspend/resume and persistence, and has published Temporal integration for workflows requiring multi-hour/day durability and worker resilience [S06].

This is important architectural evidence:

> An agent/workflow framework can still need a dedicated durable engine for stronger execution guarantees.

Aurora should not assume that selecting Mastra or another agent framework resolves Core durability.

---

## 10. Minimal Aurora-owned baseline

A spike should include a minimal local baseline rather than only third-party engines.

Possible baseline:

```text
operational database
+ explicit state machine
+ append-only material events
+ lease/heartbeat
+ idempotency/effect receipts
+ timer scheduler
+ restart reconciler
```

Purpose:

- establish minimum required complexity;
- measure third-party value;
- validate Aurora domain before engine mapping;
- potentially support M0–M3 before M4.

It should not grow into a home-built general workflow engine if a mature engine clearly fits.

---

## 11. Determinism and replay

Engines differ in replay constraints.

Aurora must isolate:

### Deterministic orchestration

- transition choice from recorded inputs;
- scheduling/waits;
- references to completed effects;
- state projection.

### Non-deterministic/side-effectful work

- model call;
- web request;
- filesystem;
- repository;
- device command;
- current time/randomness;
- telemetry query.

These occur through Activities/Steps/Effect Gateways with recorded results.

A model response should not be recomputed during replay and assumed identical.

---

## 12. Effect semantics

No engine can guarantee exactly-once external reality when the external system lacks idempotency/transaction support.

Aurora needs:

```text
Effect Request ID / idempotency key
→ record intent before send where appropriate
→ execute through gateway
→ capture external reference/receipt
→ reconcile after timeout/restart
→ compensate or mark ambiguous
```

Engine retries are allowed only around a correctly designed effect boundary.

---

## 13. Checkpoint ownership

Potential checkpoints:

- Core Mission state;
- durable engine workflow history;
- Harness local snapshot;
- model conversation/thread state;
- artifact/evidence progress;
- device state.

A `Recovery Snapshot` should reference all relevant owners rather than copy them blindly.

Example:

```yaml
recovery_snapshot:
  delegation: DEL-024
  core_state_version: 188
  engine_execution: temporal://workflow/...
  provider_snapshot: provider://.../checkpoint/12
  context_pack: CTX-024-R3
  authority_grant: GRANT-024-R2
  completed_effects:
    - EFFECT-1001
  published_artifacts:
    - ART-301
```

---

## 14. Waiting for human decision

Requirements:

- durable wait without active compute;
- decision identity and alternatives;
- deadline/default-safe action;
- notification/attention policy;
- authority expiration during wait;
- Context Pack refresh on resume;
- exact code/contract version;
- cancellation.

A long wait can make assumptions stale. Resume should revalidate:

- project state;
- provider availability;
- grant;
- budget;
- device/environment;
- source freshness.

---

## 15. Versioning in-flight work

Possible strategies:

- pin old workflow/provider build until completion;
- explicit version marker/patching;
- migrate state to new version;
- cancel/replan;
- keep compatibility adapter.

Requirements:

- no silent code replacement;
- old artifact/provider build remains available or migration is explicit;
- authority/policy changes can still revoke old run;
- security critical patch may force containment;
- history preserves which code produced which effect.

---

## 16. Operational burden criteria

Compare candidates on:

- services/processes required;
- local install and backup;
- monitoring;
- database/cluster;
- resource use;
- upgrades;
- developer debugging;
- worker deployment;
- language SDK;
- failure recovery;
- export/migration;
- vendor/cloud dependency.

A single-user system should not accept distributed-system complexity without a Golden Proof.

---

## 17. SPK-004 design

### Scenario

```text
1. Start one durable Delegation.
2. Execute an idempotent external effect.
3. Record artifact and checkpoint.
4. Kill Core and provider processes.
5. Restart.
6. Recover exact state and build.
7. Prove the effect is not duplicated.
8. Wait for a human decision.
9. Expire/revoke authority while waiting.
10. Resume and prove the effect is denied until new grant.
11. Complete and produce evidence.
```

### Candidates

- minimal local baseline;
- DBOS;
- Restate;
- Temporal.

Inngest may be included as cloud-oriented comparison if current local/self-host profile is relevant.

### Measurements

- implementation complexity;
- operational components;
- recovery correctness;
- effect duplication;
- wait/signal;
- versioning;
- observability;
- latency/throughput relevant to single-user;
- failure diagnosis;
- migration/export;
- language fit.

### Decision outcome

ADR may choose:

- one engine;
- baseline for early milestones and engine later;
- different engines for Core and Harness internal workflows;
- rejection of candidates with reasons.

---

## 18. Decision implications

### Supported

- define Durable Execution Port independent from engine;
- keep Aurora operational state canonical;
- separate model calls/effects from deterministic workflow history;
- use idempotency/receipts/reconciliation;
- compare operational burden against minimal baseline;
- do not claim agent framework durability equals global durability.

### Not decided

- engine selection;
- database;
- workflow language;
- deployment model;
- whether M0–M3 use baseline only;
- whether Harnesses may choose different internal engines.

---

## 19. Limitations

- Documentation describes intended guarantees; exact behavior depends on versions and deployment.
- Benchmarks do not replace kill/partition/effect tests.
- The correct choice depends on Core language/store, which remains open.
- Physical device recovery has additional local controller/interlock requirements.
- “Exactly once” marketing must be interpreted carefully at external effect boundaries.

---

## 20. Conclusion

Aurora should not choose a durable engine as the first technical decision. She should first stabilize Mission/Delegation/effect semantics, then use SPK-004 to compare:

```text
minimum local implementation
versus
mature durable engines
```

The selected engine becomes an adapter behind the Durable Execution Port, never the owner of Aurora's global product state.
