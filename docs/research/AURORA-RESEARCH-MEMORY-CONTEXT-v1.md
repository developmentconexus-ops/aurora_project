---
id: RESEARCH-AURORA-MEMORY-CONTEXT-V1
title: Aurora Research — Memory, Knowledge and Context Architecture
document_type: research_report
form: explanation
authority: research
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - research findings on agent memory and context through 2026-08-05
related:
  - DOC-AURORA-BLUEPRINT-06
  - DOC-AURORA-BLUEPRINT-13
  - DOC-AURORA-RESEARCH-MAP
source_manifest: AURORA-RESEARCH-MEMORY-CONTEXT-v1.sources.json
review_triggers:
  - major memory benchmark or implementation release
  - Aurora M1 readiness
  - memory architecture spike result
last_reviewed: 2026-08-05
---

# Aurora Research — Memory, Knowledge and Context Architecture

## 1. Research question

Which current memory mechanisms, architectures and evaluation methods can inform Aurora's need for:

- long conversational continuity;
- project memory;
- personal/global memory;
- temporal fact updates and supersession;
- source authority;
- workflow and environment experience;
- privacy, deletion and audit;
- context construction under token/latency budgets;
- future multimodal and embodied memory?

This report does **not** select a production store or implementation framework.

---

## 2. Method

The review prioritizes:

1. current official implementations and technical descriptions;
2. primary research papers;
3. open-source repositories/evaluation kits;
4. benchmarks that test multi-session or environment experience;
5. security/governance analysis across the memory lifecycle.

The report separates:

- mechanisms that preserve conversational history;
- systems that manage agent state and editable memory;
- project/institutional knowledge architectures;
- evaluation evidence;
- gaps relative to Aurora.

---

## 3. Executive finding

There is no single current technique that satisfies Aurora's memory requirements.

The evidence supports a **hybrid, stratified architecture**:

```text
exact history and event ledger
+ observational/reflective conversational synthesis
+ structured project and operational state
+ authoritative knowledge/document retrieval
+ temporal/relational memory
+ Context Builder with authority, recency and privacy
+ evaluation across recall, workflow, supersession and scale
```

Observational memory is a strong candidate for conversational continuity, but it should not own project decisions, live device state or operational budgets. Vector/RAG retrieval remains useful but similarity alone is insufficient for temporal and authority-sensitive memory. Files and coding-agent retrieval can perform well on environment experience but introduce latency. Self-editing memory systems demonstrate persistence and portability, while also increasing the importance of versioning, security, evaluation and protected promotion.

---

## 4. Memory as write–manage–read

A 2026 survey formalizes modern agent memory as a `write → manage → read` loop coupled to perception and action, rather than a static store [S08]. This framing maps directly to Aurora:

### Write

- capture event/statement/tool result;
- determine candidate type and scope;
- preserve source and time;
- avoid premature promotion.

### Manage

- consolidate;
- relate;
- update;
- supersede;
- forget;
- protect;
- evaluate.

### Read

- select by intent/scope;
- prioritize authority/freshness;
- retrieve evidence;
- minimize for provider;
- compile active context.

This supports the Blueprint decision that storage technology is subordinate to lifecycle semantics.

---

## 5. Observational memory

Mastra's Observational Memory uses an Observer to convert conversation/tool history into dated observations and a Reflector to reorganize observations as they grow [S01]. Mastra reports high LongMemEval results with a stable context window rather than dynamically injecting arbitrary retrieved fragments every turn.

Potential Aurora fit:

- L3 observational/session synthesis;
- long-running conversation compaction;
- temporal event summaries;
- project session handoff;
- prompt-cache stability;
- source-aware recent observations.

Important limitations:

- benchmark success focuses primarily on conversational memory;
- an observation remains model-generated synthesis;
- project authority and live state need external sources;
- omission or reflection errors can become persistent;
- token-stable context can still carry stale or overgeneralized content;
- deletion and derived observation handling need explicit governance.

Decision implication:

> Investigate Observational Memory as one write/manage mechanism, not as the canonical memory architecture.

Required spike:

- compare raw-window + observations versus retrieval/hybrid on Aurora project continuation, supersession and false-premise tests.

---

## 6. Layered institutional context

OpenAI's in-house data agent describes multiple context layers: table usage, human annotations, code-derived enrichment, institutional knowledge, memory and runtime context [S02]. It retrieves relevant enriched context and performs live queries when information is missing or stale.

This provides strong evidence for Aurora's distinction between:

```text
memory
≠ knowledge source
≠ runtime/live context
```

Relevant lessons:

- domain knowledge should be curated/enriched, not inferred solely from schema/history;
- memory can be scoped globally and personally;
- users need memory creation/editing;
- live systems must be queried when stored information is stale;
- less, high-quality context can outperform broad history;
- meaning often lives in code and operational systems.

Aurora extends this pattern to projects, devices, authority and physical experiments.

---

## 7. Temporal revision and “dreaming”

OpenAI's 2026 memory work describes asynchronous memory revision that updates time-sensitive statements as time passes [S03]. The key conceptual value is not the product implementation but the recognition that memory needs ongoing maintenance rather than append-only facts.

Aurora implications:

- planned future event can become completed historical event;
- current preference/state may supersede older value;
- consolidation jobs can detect stale temporal expressions;
- source and validity interval should remain;
- asynchronous maintenance must not invent completion from time alone.

Example:

```text
"Test will run tonight"
```

cannot become:

```text
"Test completed successfully"
```

without operational evidence. Time can update grammatical status; it cannot manufacture outcome.

---

## 8. Stateful/self-editing agent systems

Letta (formerly MemGPT) demonstrates stateful agents with editable memory blocks, message history, model-agnostic operation and self-improvement mechanisms [S04]. Letta Code emphasizes cohesive identity across models and stores memory/context in a Git-tracked MemFS [S05]. Agent File proposes portable serialization of agent state including history, system prompt, memory blocks and tool rules [S06].

Relevant architectural lessons:

- model identity can be separated from durable agent identity;
- memory can be explicitly editable and versioned;
- file/Git representations improve inspectability and portability;
- context can include protected/core blocks and searchable history;
- self-editing creates a need for audit, protected boundaries and rollback;
- serialization formats do not guarantee semantic equivalence across frameworks.

Aurora implications:

- a portable/exportable memory representation is desirable;
- version control may suit selected human-readable memory/procedure artifacts;
- operational state should not be represented only as editable prompt files;
- protected constitutional memory must remain outside ordinary self-editing;
- provider/framework export requires semantic adapters, not only file copying.

---

## 9. Storage → Reflection → Experience

A 2026 survey frames the evolution of agent memory as:

```text
Storage
→ preserve trajectories

Reflection
→ refine and summarize trajectories

Experience
→ abstract reusable strategies across trajectories
```

[S07]

This maps to Aurora's classes:

- raw/session/episodic history;
- observational consolidation;
- procedural memory and Golden Paths;
- Failure Intelligence and cross-incident causal learning.

The framework reinforces that “remembering” is not only recalling facts. A mature engineering companion needs to learn:

- workflows;
- environment gotchas;
- failed approaches;
- discriminating tests;
- validated procedures.

However, experience abstraction is also where overgeneralization risk is highest. Promotion requires repeated evidence and explicit scope.

---

## 10. Memory security lifecycle

A 2026 long-term memory security survey identifies lifecycle phases:

```text
Write
Store
Retrieve
Execute
Share/Propagate
Forget/Rollback
```

and security objectives across integrity, confidentiality, availability and governance [S09].

This directly supports Aurora's requirement that security cannot be added only at retrieval time.

Threat examples:

- malicious or incorrect memory written persistently;
- poisoned relationship/identity memory;
- unauthorized retrieval across projects;
- memory-driven unsafe tool execution;
- propagation to external providers or child harnesses;
- incomplete deletion from summaries/embeddings/backups;
- rollback restoring revoked grants or stale facts.

Required Aurora controls:

- provenance at write;
- versioning and supersession;
- scope/data classification;
- policy-aware retrieval;
- source validation before effects;
- derived-data tracking;
- audited deletion/rollback;
- protected memory classes.

---

## 11. Evaluation beyond conversational recall

### LongMemEval-V2

LongMemEval-V2 evaluates environment-specific experience across:

- static state recall;
- dynamic state tracking;
- workflow knowledge;
- environment gotchas;
- premise awareness [S10].

This is closer to Aurora than pure personal-chat recall because projects and tools require workflow and failure knowledge.

The reported strong result of a coding-agent/file-based gatherer also shows a trade-off: richer active search over files can improve accuracy but at high latency.

Aurora implication:

- evaluate different retrieval interfaces, including structured index, file workspace and coding-agent gathering;
- measure accuracy/latency/cost Pareto;
- include premise-awareness tests so Aurora rejects false “we already decided X” claims.

### Scale-conditioned evaluation

Recent work shows memory reliability can degrade as irrelevant sessions grow even when relevant evidence remains unchanged [S11]. Fixed benchmark accuracy therefore hides usable-scale limits.

Aurora implication:

- longitudinal tests must add irrelevant project/session noise;
- report budget-compliant reliability and tail memory calls;
- test per model/interface combination;
- define a usable-scale boundary rather than claim “supports long memory.”

### Supersession gap

2026 work isolates a significant gap in using current rather than superseded facts, and finds that simply increasing memory size does not solve it [S12].

Aurora implication:

- supersession is a first-class lifecycle/benchmark;
- stronger models alone do not remove stale-memory risk;
- use linked temporal chains/current projections;
- include explicit current-value reward/eval cases;
- never rely on semantic similarity alone to resolve updates.

---

## 12. Retrieval mechanisms

### Raw/full history

Benefits:

- exact source;
- no summary omission.

Limits:

- context ceiling;
- cost/latency;
- irrelevant noise;
- privacy exposure;
- poor scale.

### Compaction/summary

Benefits:

- simple;
- low context cost.

Limits:

- information loss;
- false synthesis;
- temporal flattening;
- difficult source recovery.

### Vector/RAG

Benefits:

- scalable semantic retrieval;
- mature tooling;
- suitable for documents/history.

Limits:

- similarity is not authority or currentness;
- poor exact/state/relationship queries;
- deletion/index governance;
- retrieval can miss workflow or distributed evidence.

### Full-text/metadata filtering

Benefits:

- exact terms/IDs;
- deterministic filters;
- useful with schemas and dates.

Limits:

- weak semantic match;
- requires structured metadata.

### Relational/structured state

Benefits:

- exact entities, scopes, constraints and current projections;
- deterministic queries and transactions.

Limits:

- schema evolution;
- less flexible unstructured knowledge;
- extraction cost.

### Graph/temporal relationships

Benefits:

- relation traversal;
- provenance/temporal chains;
- project/world model.

Limits:

- noisy extraction;
- ontology/maintenance cost;
- graph retrieval alone does not solve narrative context.

### File/workspace + agent search

Benefits:

- inspectable;
- code/file tools can gather complex evidence;
- Git/versioning.

Limits:

- latency;
- active-agent variability;
- security/tooling complexity.

### Learned/agent-managed memory

Emerging work treats memory operations as agent actions [S08].

Potential:

- adaptive write/read/forget;
- task-specific management.

Risk:

- policy learned from benchmark may violate governance;
- difficult predictability;
- requires deterministic ceilings and protected stores.

---

## 13. Recommended logical architecture for spikes

```text
Session Ledger
├── exact messages/tool events

Observation Pipeline
├── Observer
└── Reflector/Consolidator

Structured Project/Operational Stores
├── current state
├── decisions/hypotheses
├── temporal relations
└── authority references

Knowledge Source Index
├── documents/code/research
├── full-text/metadata
└── semantic retrieval

Memory Stores
├── personal/global
├── project
├── episodic
├── procedural
└── failure/learning

Context Builder
├── scope
├── authority/freshness
├── retrieval
├── conflict resolution
├── minimization/redaction
└── provenance

Retention/Deletion
└── derived data and audit

Evaluation Harness
└── recall/non-recall/supersession/workflow/scale/security
```

This is a logical decomposition, not a microservice mandate.

---

## 14. Decision implications

### Supported

- memory must be stratified and hybrid;
- active context must be compiled, not equal to memory store;
- live/runtime context is distinct from memory;
- source provenance and editable memory are important;
- observational memory deserves an architecture spike;
- temporal supersession must be explicit;
- evaluation must include workflow, premise awareness, scale and security;
- local-first canonical memory is justified by sovereignty and lifecycle control.

### Not supported yet

- selecting Mastra Observational Memory as the full solution;
- selecting Letta as Aurora Core;
- choosing vector, graph or relational database;
- claiming benchmark scores transfer to Aurora projects/laboratory;
- allowing autonomous global memory promotion;
- using model-managed memory without deterministic governance.

---

## 15. Required architecture spikes

### MEM-SPK-001 — Conversational observation

Compare:

- sliding raw history + summary;
- Observational Memory approach;
- semantic retrieval;
- hybrid.

Test long discovery conversation with decisions, corrections and temporal updates.

### MEM-SPK-002 — Project authority

Given contradictory chat, research and ADR, recover current decision and provenance.

### MEM-SPK-003 — Supersession

Update changing facts/preferences/firmware and answer current versus historical questions.

### MEM-SPK-004 — Scope isolation

Add similar content across multiple projects and prove no contamination.

### MEM-SPK-005 — Usable-scale

Keep relevant evidence fixed while adding irrelevant sessions/projects; measure reliability, latency, calls and tokens.

### MEM-SPK-006 — Deletion

Delete a memory/source and verify raw, observational, vector, graph, cache and backup behavior according to policy.

### MEM-SPK-007 — Workflow experience

Test whether memory helps with project workflows, environment gotchas and recovery—not only facts.

### MEM-SPK-008 — False-memory/premise awareness

Inject plausible but unsupported premise and verify rejection.

---

## 16. Evaluation matrix

| Dimension | Example proof |
|---|---|
| Recall | retrieve accepted decision after session restart |
| Non-recall | do not expose unrelated project memory |
| Temporal | use current firmware but answer historical incident correctly |
| Supersession | old preference excluded from current profile |
| Authority | ADR beats conversation summary |
| Provenance | show exact source/ref |
| Premise awareness | reject fabricated prior decision |
| Workflow | remember correct multi-step environment procedure |
| Gotcha | avoid previously observed failure |
| Scale | maintain reliability as irrelevant history grows |
| Efficiency | latency/token/storage budget |
| Security | poisoned memory does not trigger effect |
| Deletion | derived copies no longer retrieve |
| Longitudinal | operate after months/projects of growth |

---

## 17. Limitations

- Many memory benchmarks still focus on QA over conversations rather than real autonomous projects.
- Vendor-reported benchmark results may depend on model, prompts, dataset handling and cost.
- 2026 papers are recent and may not yet have broad replication.
- Physical/multimodal memory is less mature than text memory.
- Privacy/deletion claims require implementation-level testing.
- Aurora's single-user, authority-rich, engineering environment is not directly represented by one benchmark.

---

## 18. Research conclusion

The evidence validates treating memory as a first-class governed subsystem and Context Builder as a critical product component.

Aurora should not select “a memory framework” as one decision. She needs a sequence of narrower decisions:

1. exact history/event representation;
2. observational synthesis;
3. structured project and operational state;
4. knowledge retrieval;
5. temporal/relational model;
6. promotion and epistemic governance;
7. Context Builder;
8. privacy/deletion;
9. evaluation and usable-scale limits.

Only after the memory spikes should a storage/runtime ADR be proposed.
