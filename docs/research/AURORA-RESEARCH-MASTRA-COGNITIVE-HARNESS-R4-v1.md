---
id: RESEARCH-AURORA-MASTRA-COGNITIVE-HARNESS-R4-V1
title: Aurora R4 Research — Mastra as Cognitive and Harness Runtime
document_type: research_report
form: explanation
authority: research
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - Mastra architecture-fit research through 2026-08-07
  - evidence for Mastra cognitive/harness-runtime boundary decisions
related:
  - DOC-AURORA-BLUEPRINT-06
  - DOC-AURORA-BLUEPRINT-07
  - DOC-AURORA-BLUEPRINT-10
  - DOC-AURORA-BLUEPRINT-12
  - DESIGN-AURORA-M0-R4-DECISION-LANDSCAPE
  - ADR-AURORA-0003
  - ADR-AURORA-0004
source_manifest: AURORA-RESEARCH-MASTRA-COGNITIVE-HARNESS-R4-v1.sources.json
review_triggers:
  - first Mastra-backed Aurora Capability enters implementation horizon
  - Mastra major/minor API change affecting AgentController, memory, workflows, storage or server boundary
  - Mastra license boundary changes
  - accepted Aurora state/authority ownership changes
last_reviewed: 2026-08-07
---

# Aurora R4 — Mastra as Cognitive and Harness Runtime

## 1. Research question

Can Aurora reuse Mastra aggressively for agentic capabilities without making Mastra the owner of Aurora identity, canonical project/mission state, authority, global acceptance or sovereignty?

This investigation was triggered after the initial M0 R4 documentary checkpoint because Mastra's 2026 feature surface had expanded materially beyond the earlier generic framework comparison. The question is therefore not:

> “Does Mastra have agents, memory and workflows?”

It is:

> **Which Aurora responsibilities can safely be delegated to Mastra, which require an Aurora-owned adapter/governance layer, and which must remain entirely outside Mastra?**

The investigation also asks whether this finding should delay the M0 Sovereign Core path.

## 2. Method and materiality rule

The approved R4 philosophy remains:

```text
Explore long horizon.
Commit only to evidence-supported irreversible steps.
```

A second rule is applied to avoid analysis paralysis:

```text
Materiality before investigation.

If A versus B does not change the next architecture/build decision,
record it and move on.
```

Therefore this report is deep on ownership boundaries and feature classes that could replace Aurora-built infrastructure, but intentionally does not attempt to benchmark every Mastra provider, integration or UI feature.

Technical claims use current primary Mastra documentation, release metadata, repository licensing and current accepted Aurora Blueprints. Exact implementation/API details must be pinned to the chosen Mastra release at the future implementation gate.

## 3. Current maturity and product direction

The latest stable repository release observed during this research is `@mastra/core@1.56.0`, published 2026-08-06 [S01]. The release includes persistable declarative workflows, stored-workflow HTTP APIs and persistence, richer eval controls and observability/streaming changes [S01].

This matters because Mastra's current scope is no longer accurately described as only an agent-call abstraction. The 2026 surface now spans:

```text
agents and model routing
AgentController interactive runtime
memory and observational memory
workflows and persisted snapshots
background tasks
durable agents
signals and notifications
schedules
goals
tool hooks
workspaces and skills
RAG
A2A / ACP / SDK subagents
storage abstractions
server/OpenAPI exposure
observability and evals
software-factory compositions
```

Mastra itself describes `AgentController` as a control layer for stateful interactive agent experiences and explicitly states that Mastra broadly already supplies the runtime around agents: tools, memory, workflows, storage, human-in-the-loop and signals [S04]. Mastra Factory demonstrates those primitives composed into a governed software-delivery loop [S26].

### 3.1 Licensing boundary

The repository root license states that code outside `ee/` restrictions is Apache-2.0, while code under `ee/` is governed separately [S02]. The Enterprise Edition license restricts production use of EE code without a written agreement [S03].

R4 implication:

```text
Aurora foundational dependency
→ prefer OSS/Apache-2.0 Mastra primitives

EE/platform convenience
→ optional adapter/infrastructure only
→ never required for constitutional continuity
```

The distinction must be checked by exact package/path before a production dependency is accepted.

## 4. AgentController and Harness fit

AgentController provides sessions, modes, threads, permissions, model switching, subagents, persisted state and UI events [S04]. Its predecessor was called `Harness`; the rename reflects that Mastra itself is a broader agent harness while AgentController specifically controls the interactive agent experience [S04].

This maps unusually well to Aurora's accepted Harness definition in Blueprint 07:

```text
Aurora Harness local concerns
↔
Mastra AgentController / agent runtime

local session/thread
local execution mode
local plan/task state
local model/provider choice
local subagents
local approvals
local event/UI state
local recovery
```

### R4 interpretation

`AgentController` is a strong candidate to avoid building a generic first-party interactive Harness runtime from scratch.

It must remain **Harness-local**. Its persisted state cannot become:

- Aurora Project current state;
- Aurora Mission/Delegation source of truth;
- Authority Grant source;
- global budget authority;
- provider approval/trust source;
- global acceptance/outcome.

A provider may lose or rebuild AgentController state while Aurora remains able to identify the provider run and classify it as recoverable, degraded, failed or requiring reconciliation.

## 5. Memory fit

Mastra supports conversation/message history, working memory, semantic recall and observational memory. Observational Memory uses Observer/Reflector agents to convert growing message history into dense observations and maintain a stable context shape [S05][S06].

Mastra currently treats Observational Memory as a primary long-horizon memory direction [S05]. Its published LongMemEval results demonstrate strong conversation-recall performance, but those benchmarks do not test Aurora's product-state, authority or physical correctness semantics [S06].

### 5.1 Mapping to Aurora strata

| Aurora stratum | Mastra fit | R4 posture |
|---|---|---|
| `L0` authoritative sources/state | poor as owner | **DO NOT USE AS OWNER** |
| `L1` current authoritative snapshots | useful as injected projection only | **WRAP** |
| `L2` governed durable memory | useful extraction/storage substrate, incomplete governance | **ADAPT/WRAP** |
| `L3` observational/session synthesis | excellent direct fit | **USE** |
| `L4` exact interaction/tool history | strong fit | **USE** |
| `L5` active context/transport | strong fit | **USE** |

### 5.2 Important terminology mismatch

Mastra Working Memory is persistent structured memory that can survive threads/resources. Aurora's accepted conceptual `working memory` is short-lived current-turn/task state and should not automatically become durable.

Therefore:

```text
Aurora.WorkingMemory != Mastra.WorkingMemory
```

The mapping must use explicit Aurora names/scopes rather than importing Mastra terminology into the domain.

### 5.3 Memory Extractors

Memory Extractors can derive structured values during observational-memory processing and invoke `onExtracted` callbacks for downstream systems [S07].

This is a particularly strong fit for Aurora's memory lifecycle if treated as:

```text
conversation/history
→ Mastra Observer/Reflector
→ structured extraction
→ MEMORY CANDIDATE
→ Aurora governance
   provenance
   epistemic status
   scope
   temporal validity
   sensitivity
   dedupe/supersession
→ promoted governed memory or rejected candidate
```

The extractor is a **candidate producer**, not a memory authority.

## 6. Workflows, HITL and durable execution

Mastra workflows persist snapshots used for suspend/resume and human-in-the-loop continuation [S12]. The current release adds JSON-persistable declarative workflow graphs and HTTP/storage lifecycle for workflow definitions [S01]. Mastra also supports Temporal-backed workflows for stronger durable execution when required [S13].

Mastra Durable Agents add resumable run/stream behavior using `runId` and durable/cache/evented mechanisms [S11]. Background Tasks persist long-running tool execution state and can resume across server restarts [S10].

### R4 interpretation

These are strong mechanisms for **Harness-local execution** and later for implementations behind Aurora's `DurableExecutionPort`.

They do not alter M0's state ownership rule:

```text
Mastra workflow snapshot / durable run / background task
!=
Aurora canonical Project or Mission state
```

For M0 specifically, no Mastra workflow/durable engine is required. M0 must recover identity/state/authority even when every Mastra process is absent.

The current M0 `NOT_YET_A_DECISION` posture for a durable workflow engine remains correct.

## 7. Signals, schedules, goals and proactivity

### 7.1 Signals

Agent Signals support multiple clients observing/steering one thread, reactive guidance, state signals and externally generated notification signals; they can wake idle agents and operate across multiple processes when Pub/Sub is configured [S14].

This is a strong future primitive for Presence/proactivity and Harness-local reactive context.

But a Signal is:

```text
input / notification / provider-local context
```

not:

```text
authenticated authority
canonical global truth
```

Developer-defined signal attributes such as role/source are context metadata and cannot substitute Aurora actor authentication/Authority Grant checks.

### 7.2 Schedules

Mastra can persist schedules for agents/workflows [S15]. This is useful for recurring Harness work, but schedule existence cannot grant permission to execute an effect. Aurora must separately govern whether the scheduled action remains authorized at execution time.

### 7.3 Goals

Mastra Goals persist thread-scoped objectives and use an LLM judge to determine whether a goal loop passes or continues [S16].

Useful mapping:

```text
Mastra Goal
→ local Harness objective / progress-control primitive
```

Invalid mapping:

```text
Mastra goal judge passed
→ Aurora Mission accepted
```

Global Aurora acceptance continues to require authoritative criteria/evidence and operator/product governance.

## 8. Tools, Workspaces, Skills and effect boundaries

Mastra Workspaces provide filesystem access, command execution, search and Skills with local permission controls [S08]. First-party Skills can be attached to agents/workspaces and stored/versioned through Mastra storage surfaces. The official Agent Harness template demonstrates local filesystem/tool approval/scheduling/memory composition and explicitly warns that `LocalSandbox` is not operating-system isolation by default [S09].

Tool Hooks can observe/intercept calls and block execution before a tool runs [S17]. This gives Aurora a natural integration point:

```text
Mastra agent selects tool
→ beforeToolCall
→ Aurora Effect/Authority Adapter
→ current Authority Grant / policy decision
→ ALLOW | DENY | REQUIRE_CONFIRMATION | MODIFY
→ tool executes only if permitted
→ receipt/audit returned to Aurora
```

Tool Hooks are therefore valuable **enforcement adapters**, but the authoritative policy state cannot live only in Mastra memory, prompt state or provider-local approvals.

Workspaces reduce the amount of generic file/search/command/skill infrastructure Aurora Harnesses need to build, while OS/container/device isolation remains a separate deterministic boundary.

## 9. Interoperability and external Harnesses

Mastra supports A2A remote agents [S18], ACP-compatible coding harnesses including Claude Agent, Codex CLI, Cursor and Gemini CLI [S19], and SDK-wrapped Claude/Cursor/Codex agents that inherit Mastra's composition/observability surface [S20].

This makes Mastra a strong candidate implementation substrate for a large portion of the future Aurora Capability Fabric.

However, Blueprint 07's non-transitive authority rule remains decisive.

Bad flow:

```text
Mastra Harness A has grant
→ directly gives context/permission to Codex child
```

Required Aurora flow for a new authority boundary:

```text
Harness A emits CapabilityRequest
→ Aurora evaluates scope/provider/data/budget
→ Aurora creates child Delegation + Context Pack + Authority Grant
→ adapter instructs Mastra to invoke ACP/A2A/SDK provider
→ provider-local run executes
```

Mastra may perform the transport and runtime composition. Aurora owns the delegation semantics.

## 10. RAG and knowledge retrieval

Mastra RAG provides document chunking, embeddings, multiple vector stores, retrieval, metadata filtering, reranking, Graph RAG/ReAG and agentic retrieval patterns [S21].

This is sufficiently rich that Aurora should not build a generic RAG framework from scratch without a capability-specific reason.

Recommended boundary:

```text
Aurora Context Builder
→ query + source scope + freshness + classification + authority precedence
→ Mastra retrieval adapter
→ retrieved chunks + source refs + scores + limitations
→ Aurora Context Builder compiles active context
```

The vector/retrieval system does not own document authority, current project truth or source freshness.

## 11. Storage architecture

Mastra 1.0 introduced composite storage so domains such as memory, workflows and scores can use different backends [S22]. Current Mastra storage surfaces support multiple persistent backend classes; the exact provider can evolve independently.

This is positive for Aurora because Mastra-local data does not need to share the Sovereign Core database.

Recommended physical separation:

```text
Aurora Core store
→ canonical identity/project/authority/mission governance

Mastra storage
→ threads/messages/observations/workflow snapshots/tasks/schedules/traces
```

They MAY later share one database engine for operational convenience, but logical ownership and schemas must remain distinct. A Mastra storage migration cannot be treated as an Aurora domain migration automatically.

## 12. Server and Go integration boundary

Mastra Server Adapters expose agents, workflows, tools and MCP endpoints over HTTP and can generate OpenAPI [S23]. This supports Leandro's proposed architecture naturally:

```text
Go Sovereign Core
      │
      │ Aurora-owned provider contract
      ▼
MastraProviderAdapter
      │
      │ HTTP/OpenAPI / version-pinned Mastra API
      ▼
Mastra service/runtime
```

The Go domain should not import Mastra TypeScript types or internal storage models. The adapter translates Aurora contracts to the selected Mastra server API.

This separate-process shape is attractive because:

- Go Core and Node/TypeScript Mastra have independent lifecycles;
- Mastra can restart/upgrade without rewriting Core;
- a future replacement runtime remains possible;
- failure-domain testing is direct (`kill Mastra` while Core survives);
- Mastra can evolve faster behind a pinned adapter.

This is a **cross-horizon architecture direction**, not an M0 requirement. M0 itself does not need to start a Mastra process.

## 13. Observability, evals and software factory

Mastra Observability provides logs, traces and metrics; its platform uses OTel-shaped signals and a ClickHouse-based backend [S24]. Gates/Scorers/Verdicts support deterministic and model-based evaluation in normal test/CI flows [S25]. Mastra Factory demonstrates specialized agents linked by explicit workflow gates, memory, scheduling, tools and observability [S26].

Recommended mapping:

```text
Mastra trace/eval
→ provider-local observability/evidence source

Aurora Evidence
→ references relevant trace/eval artifacts
→ records producer/version/environment/limitations

Aurora Verdict/Outcome
→ remains independent
```

Mastra Factory should be treated as a valuable reference and possible substrate for specialized software-engineering Harnesses, not as Aurora global Mission Control.

## 14. Ownership matrix

| Concern | Recommended owner | Mastra role |
|---|---|---|
| agent loop/model routing | Harness/Mastra | `USE` |
| AgentController session/modes/tasks | Harness/Mastra | `USE` |
| exact agent message/tool history | Harness/Mastra | `USE` |
| observational/session synthesis | Harness/Mastra | `USE` |
| semantic recall / retrieval mechanics | Harness/Mastra | `USE` |
| RAG pipeline | Harness/Mastra | `USE` |
| local workflows/HITL/snapshots | Harness/Mastra | `USE` |
| workspaces/skills/local tools | Harness/Mastra | `USE` |
| signals/schedules/background tasks | Harness/Mastra | `USE`, authority checked externally |
| local goal loops | Harness/Mastra | `ADAPT` |
| memory extraction | Mastra producer + Aurora governance | `WRAP` |
| tool permission/effect decision | Aurora authority + gateway | Mastra hook is `WRAP` point |
| cross-provider transport (A2A/ACP) | Aurora Delegation + Mastra binding | `WRAP` |
| local tracing/evals | Harness/Mastra | `USE` as evidence input |
| Aurora identity | Aurora Core | `DO NOT USE AS OWNER` |
| Project current state | Aurora Core | `DO NOT USE AS OWNER` |
| Mission/Delegation global state | Aurora Core | `DO NOT USE AS OWNER` |
| Authority Grant | Aurora Core | `DO NOT USE AS OWNER` |
| global budget/guardrails | Aurora Core | `DO NOT USE AS OWNER` |
| provider approval/trust | Aurora Core | `DO NOT USE AS OWNER` |
| governed durable memory semantics | Aurora Memory subsystem | `ADAPT/WRAP` |
| artifact/evidence identity | Aurora Core | `DO NOT USE AS OWNER` |
| global acceptance/verdict/outcome | Aurora/operator | `DO NOT USE AS OWNER` |

## 15. Main risks

### R1 — overlapping durable state

Mastra can persist threads, memory, workflow state, schedules, goals, tasks and traces. Aurora will persist project/mission/authority/evidence state. Without explicit classification, the system can become the accidental distributed platform prohibited by Blueprint 12.

Mitigation:

```text
every Mastra-persisted record classified as
LOCAL_EXECUTION_STATE | MEMORY | HISTORY | PROJECTION | CACHE | EVIDENCE_INPUT
never implicit GLOBAL_CANONICAL_STATE
```

### R2 — rapid framework evolution

The Harness→AgentController rename and the large feature/release cadence demonstrate meaningful API evolution [S01][S04].

Mitigation:

- pin exact compatible Mastra versions per Mission/implementation baseline;
- keep Mastra imports inside the TypeScript runtime/adapter;
- run contract/conformance tests before upgrades;
- preserve Aurora-owned external contracts.

### R3 — permissions confused with authority

AgentController/workspace/tool approvals are useful local controls but cannot mint Aurora authority.

Mitigation: current Core/PDP decision on material effects, with Mastra hooks as one enforcement integration point.

### R4 — probabilistic memory or goal results promoted to truth

Observational memory, extractors and goal judges are model-derived.

Mitigation: provenance/epistemic classification, separate authoritative state and independent acceptance.

### R5 — sandbox assumptions

Local workspace command controls do not equal OS isolation [S09].

Mitigation: future effect/security plane supplies actual sandbox/container/device boundaries where risk requires them.

### R6 — license/platform coupling

EE code has production restrictions [S03], and hosted Platform services add vendor dependency.

Mitigation: foundational path uses OSS primitives; platform/EE features remain replaceable optional infrastructure.

## 16. Horizon analysis

### H0 — M0 Sovereign Core

Mastra should **not** participate in M0 canonical continuity.

M0 success must remain possible with every Mastra/model/Harness process absent:

```text
Aurora identity survives
Project state survives
Authority survives
invalid transition remains rejected
export/restore remains valid
```

Therefore this Mastra finding does not replace or delay the required M0 storage/recovery spike.

### H1 — M1–M4

Mastra becomes highly relevant:

- M1 Memory/Context: OM, message history, extraction and retrieval can replace substantial custom infrastructure while Aurora retains governance;
- M2 Registry/AHDK/Harness: AgentController, Workspaces, Skills, A2A/ACP and server boundaries can accelerate the Golden Path;
- M3 Delegation: Mastra subagents/bindings can implement provider-local execution under Aurora-owned child Delegations;
- M4 Durable Execution: Mastra durable agents/workflows/Temporal integration become candidate `DurableExecutionPort` adapters.

### H2 — long-horizon Aurora

Mastra's signals, channels, workspaces, scheduling, observability and evolving factory primitives are compatible with Presence/proactivity and specialized Harnesses provided Aurora continues to own global identity/state/authority/evidence semantics.

## 17. Architecture recommendation

The strongest current posture is:

```text
Aurora Sovereign Core
→ Go
→ owns truth, identity, authority and governance

Mastra
→ preferred candidate/default substrate for first-party agentic Harnesses
→ owns cognition and local agent execution where fit is proven

Aurora Contract Boundary
→ prevents Mastra from becoming product ontology
```

This is not a recommendation to make every future Capability use Mastra. A capability may select a different runtime when its requirements justify it. It is a recommendation to **reuse Mastra by default before building generic agent infrastructure ourselves**.

## 18. R4 materiality disposition

This finding is **material to long-horizon architecture interpretation** but **not a new M0 gate blocker**.

Required R4 amendments:

1. narrow ADR-0003 explicitly to the **Sovereign Core** runtime, not all Aurora agentic runtimes;
2. reaffirm ADR-0004 that Mastra workflow/durable history is not M0 canonical state and is not introduced into M0;
3. record a proposed cross-horizon ADR for Mastra as the preferred first-party cognitive/Harness substrate;
4. classify a future Mastra boundary proof as required **when the first Mastra-backed capability enters implementation horizon**, not before M0 storage work;
5. keep `SPK-AURORA-M0-SOVEREIGN-STORE-001` as the next executable architecture experiment if separately authorized.

No new Mastra spike is required to continue M0.

## 19. Future proof trigger — intentionally not an M0 blocker

Before the first Mastra-backed Capability is promoted for implementation, a bounded conformance/boundary proof should demonstrate at least:

```text
Core creates authoritative Delegation/Authority
→ adapter invokes Mastra
→ Mastra persists local thread/workflow/memory
→ material tool request consults Aurora authority
→ child provider request crosses Aurora Delegation boundary
→ kill/restart Mastra
→ Core canonical state remains unchanged/available
→ provider-local state reconciles or fails explicitly
→ deleting Mastra-local state cannot erase Aurora identity/authority/project truth
```

This proof belongs to the first relevant Capability's R4/R6 path. Creating or executing it now would violate the materiality rule because M0 does not consume Mastra.

## 20. Final research finding

```text
MASTRA FIT: STRONG
ROLE: preferred candidate/default cognitive + Harness runtime substrate
CORE OWNERSHIP: REJECTED
M0 BLOCKER: NO
CUSTOM REIMPLEMENTATION POSTURE: reuse Mastra first for generic agent infrastructure
VERSION/LICENSE POSTURE: exact pinning + OSS boundary + adapter isolation
```

The architecture should exploit Mastra aggressively where it reduces undifferentiated agent infrastructure, while preserving Aurora's sovereign domain and deterministic authority/evidence boundaries.
