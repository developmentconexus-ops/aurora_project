---
id: RESEARCH-AURORA-AGENT-FRAMEWORKS-RUNTIMES-V1
title: Aurora Research — Agent Frameworks and Specialized Runtimes
document_type: research_report
form: explanation
authority: research
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - agent framework landscape research through 2026-08-05
related:
  - DOC-AURORA-BLUEPRINT-05
  - DOC-AURORA-BLUEPRINT-07
  - DOC-AURORA-BLUEPRINT-12
source_manifest: AURORA-RESEARCH-AGENT-FRAMEWORKS-RUNTIMES-v1.sources.json
review_triggers:
  - first reference Harness runtime selection
  - framework major release
  - SPK-008 result
last_reviewed: 2026-08-05
---

# Aurora Research — Agent Frameworks and Specialized Runtimes

## 1. Research question

Which current agent/workflow frameworks can be useful **inside** Aurora Harnesses, and what should remain outside them?

Reviewed:

- Mastra;
- LangGraph;
- Pi;
- OpenHands;
- OpenAI Agents SDK;
- Langflow.

The report does not select the Aurora Core framework or first Harness runtime.

---

## 2. Executive finding

Each candidate is optimized for a different problem. None should become the constitutional architecture of Aurora.

```text
Mastra
→ TypeScript agents/workflows/memory/evals/tooling

LangGraph
→ low-level stateful graph runtime and human interrupts

Pi
→ minimal coding-agent runtime/SDK/RPC; strong MNFS relevance

OpenHands
→ software-engineering agents, workspaces and sandbox-oriented execution

OpenAI Agents SDK
→ concise agent/tool/handoff/guardrail/tracing patterns

Langflow
→ visual flow prototyping and MCP-enabled components
```

Aurora should select framework per Harness, map through the Contract Model and validate framework neutrality with two implementations of the same Capability.

---

## 3. Evaluation criteria

- intended domain;
- language/ecosystem;
- agent/tool abstraction;
- deterministic workflow support;
- adaptive loops/branching;
- persistence/checkpointing;
- human-in-the-loop;
- long-running durability;
- model/provider flexibility;
- MCP/A2A/tool integration;
- observability/evals;
- sandbox/security boundary;
- deployment/operational burden;
- ability to hide behind AHDK;
- lock-in risk;
- maturity/current development.

---

## 4. Mastra

Mastra is a TypeScript AI framework with agents, tools, workflows, memory, observability/evals and MCP-related integration [S01]. Its workflows support sequential/parallel/branch/loop patterns and suspend/resume. Mastra also integrates with Temporal for stronger multi-hour/day durability [S02].

### Strengths

- TypeScript-first;
- cohesive agents/workflows/tools;
- explicit workflow control;
- memory/eval/observability ecosystem;
- good fit with web/Node tooling;
- potential fit with Pi/MNFS-adjacent TypeScript ecosystem;
- rapid reference Harness development.

### Risks

- framework concepts can leak into Aurora domain;
- built-in memory is not the Aurora memory architecture;
- suspend/resume does not automatically solve global durability;
- security depends on execution environment/tools;
- version churn and provider-specific abstractions;
- a large all-in-one framework can invite using it as Core.

### Suitable Aurora roles

- Research Harness prototype;
- evaluation/workflow Harness;
- first TypeScript reference provider;
- internal agent/workflow runtime behind AHDK.

---

## 5. LangGraph

LangGraph is a low-level orchestration runtime for long-running, stateful agents with graph state, persistence/checkpoints, interrupts, streaming and human-in-the-loop [S03]. It can be used without all high-level LangChain abstractions.

### Strengths

- explicit graph/state model;
- adaptive/deterministic workflow composition;
- checkpointing and replay/time-travel;
- mature Python ecosystem and TypeScript availability;
- human interrupts;
- granular control.

### Risks

- graph/state can become accidental Aurora Mission model;
- external effect idempotency remains application responsibility;
- checkpointer/store choices and framework version matter;
- Python ecosystem may not align with first Core/AHDK language;
- graph complexity can become difficult to reason about;
- cross-Harness authority remains outside.

### Suitable roles

- complex research/evaluation Harness;
- self-improvement causal workflow;
- Python engineering/AI Harness;
- SPK-008 framework-neutrality comparison.

---

## 6. Pi

Pi is a minimal coding-agent ecosystem with model abstraction, coding agent, extensions, SDK and RPC mode [S04][S05]. It is the first runtime reference for MNFS.

### Strengths

- focused, extensible coding-agent core;
- TypeScript;
- in-process SDK and process-isolated RPC options;
- model/provider flexibility;
- minimal surface suitable for custom Harness control plane;
- existing relevance to MNFS.

### Risks

- coding domain, not personal/laboratory Core;
- host permissions unless externally sandboxed;
- no complete Aurora Mission/authority/memory architecture;
- minimalism shifts control-plane work to MNFS/Aurora adapters;
- current project/repository maturity must be monitored.

### Suitable roles

- MNFS internal runtime;
- software/coding Harness;
- reference for SDK versus RPC integration;
- not Aurora Core.

---

## 7. OpenHands

OpenHands provides an SDK/platform for software development agents, including agents, conversations, tools, workspaces, events, remote agent servers and security/sandbox concepts [S06].

### Strengths

- strong software-engineering domain;
- explicit workspace/tool execution;
- local/remote agent server patterns;
- event model;
- sandbox awareness;
- broader ready-made coding capabilities.

### Risks

- large domain/platform surface;
- internal semantics may overlap MNFS/Aurora;
- security policies do not replace sandboxing;
- adaptation to evidence/authority contracts may be nontrivial;
- not personal memory, presence or device architecture.

### Suitable roles

- external/alternative software Harness;
- source of workspace/sandbox/agent-server patterns;
- comparative coding capability;
- not global Core.

---

## 8. OpenAI Agents SDK

OpenAI Agents SDK offers agents, tools, handoffs, guardrails, sessions and tracing with a relatively small abstraction surface [S07]. It supports manager-style orchestration and handoffs.

### Strengths

- concise and approachable;
- tool and agent composition;
- tracing;
- guardrails;
- Python and TypeScript ecosystems;
- useful for simple internal Harness agents.

### Risks

- provider/product evolution;
- guardrails do not cover every external effect path automatically;
- handoff semantics can transfer conversation/control in ways different from Aurora hierarchy;
- session/tracing are not global state/evidence;
- hosted tools/provider coupling.

### Suitable roles

- bounded internal agent teams;
- lightweight reference Harness;
- model/tool orchestration behind AHDK;
- not authority/durability system.

---

## 9. Langflow

Langflow is a visual builder/runtime for AI flows with component ecosystem and MCP capabilities [S08].

### Strengths

- rapid visual prototyping;
- accessible flow inspection;
- component reuse;
- useful for experiments and demonstrations;
- MCP integration.

### Risks

- visual graph may become noncanonical hidden architecture;
- production lifecycle/versioning/testing must be assessed;
- weaker fit for exact authority/evidence contracts without custom components;
- generated/exported flow portability;
- not a durable global control plane.

### Suitable roles

- design/prototype laboratory;
- visual workflow comparison;
- noncritical Harness prototyping;
- not current Core hypothesis.

---

## 10. Comparative matrix

| Dimension | Mastra | LangGraph | Pi | OpenHands | OpenAI Agents SDK | Langflow |
|---|---|---|---|---|---|---|
| Primary domain | general TS agents/workflows | low-level stateful agents | coding | software agents | general agent composition | visual flows |
| Main language | TypeScript | Python/TS | TypeScript | Python/platform | Python/TS | Python/platform |
| Workflow control | strong explicit | strong graph | custom/extensions | conversations/agent loops | manager/handoffs | visual graph |
| Checkpoint/HITL | yes | strong | custom/MNFS | platform-specific | sessions/guardrails, custom durability | flow/runtime-specific |
| Strong durable engine | Temporal integration | checkpointer; external effects need care | no global engine | platform-specific | no global engine | no global engine |
| Coding specialization | moderate | customizable | strong | strong | customizable | moderate |
| Visual design | Studio/observability | Studio/platform ecosystem | terminal/RPC | UI/platform | tracing | strong |
| Security boundary | external needed | external needed | external needed | sandbox-focused but external verification | external needed | external needed |
| Aurora lock-in risk | medium/high if Core | medium if graph becomes domain | high if generalized beyond coding | high if platform becomes Core | medium if handoffs define domain | high if flows become source of truth |

This table is qualitative research, not a ranking or final selection.

---

## 11. Manager versus handoff

Aurora's accepted hierarchy is closer to a manager/control-plane pattern:

- Aurora retains global Mission and user relationship;
- Harness receives scoped Delegation;
- Harness may use internal handoffs;
- cross-Harness request returns to Aurora.

Framework handoffs that transfer the whole conversation/authority to a specialist should be adapted carefully. Aurora should not disappear as global owner.

---

## 12. Framework memory

Framework-provided memory can serve local Harness needs, but:

- canonical personal/project memory remains Aurora-owned;
- Context Pack is explicit;
- local memory scope is declared;
- durable provider state is queryable/recoverable;
- no automatic global promotion;
- deletion/data policy applies.

Mastra/Letta/LangGraph memory choices should not silently define the Core memory model.

---

## 13. Framework observability

Each framework offers different traces/events. AHDK adapter should map to Aurora/OpenTelemetry conventions:

- provider run;
- model/tool call;
- local graph/node;
- error/retry;
- artifact/evidence;
- effect request.

Framework-specific traces remain diagnostic detail. Aurora global trace and domain events remain stable.

---

## 14. Framework security

Framework guardrails and tool policies can improve behavior but do not replace:

- sandbox;
- network egress controls;
- credentials broker;
- effect gateway;
- data minimization;
- provider identity/build trust;
- audit;
- physical interlock.

A framework can be fully conformant and still unsafe if launched with broad host privileges.

---

## 15. Selection by Harness

A future Harness selection process should ask:

- Does the domain need agent reasoning or deterministic workflow?
- Is TypeScript or Python ecosystem important?
- Does it need complex adaptive graphs?
- Does it require coding workspace/sandbox?
- Does it need local embedded/device runtime?
- What durability level?
- What provider/model restrictions?
- Can AHDK integration remain clean?
- What eval/debug tooling?
- What operational burden?

Do not select one framework for organizational consistency if it harms a domain materially. AHDK/Contract Model supplies consistency at the boundary.

---

## 16. SPK-008

Implement the same capability in two different runtimes, e.g.:

- Mastra TypeScript;
- LangGraph Python or native deterministic implementation.

Capability example:

```text
analyze an artifact
→ request one bounded tool
→ publish report
→ record evidence
→ wait for one decision
```

Prove:

- same Delegation/Context/Authority;
- same artifact/evidence contracts;
- same global lifecycle;
- framework-specific state remains internal;
- both pass conformance;
- compare implementation effort, latency, debugging and recovery.

---

## 17. Decision implications

### Supported

- frameworks remain internal Harness choices;
- first real/provider runtime chosen by domain/readiness;
- Pi remains specific to MNFS/coding path;
- Mastra and LangGraph are strong reference candidates;
- OpenHands is a coding provider/platform candidate;
- OpenAI Agents SDK is useful for lightweight internal teams;
- Langflow is a prototyping/visual candidate;
- AHDK normalizes boundaries/observability;
- SPK-008 proves neutrality.

### Not decided

- first reference Harness framework;
- Core language/runtime;
- first AHDK language;
- MNFS versus another first real Harness;
- use of visual studios in production;
- framework-specific memory.

---

## 18. Limitations

- Frameworks release rapidly; current APIs may change.
- Official docs describe features, not comparative reliability under Aurora scenarios.
- Security and durability claims require fault/adversarial tests.
- Operational cost depends on deployment and provider mix.
- This report does not include every emerging framework; it covers the candidates already relevant to the project discussion.

---

## 19. Conclusion

Aurora should not answer “Mastra or LangGraph?” globally.

The correct sequence is:

```text
specify capability and Harness boundary
→ identify domain/runtime requirements
→ shortlist frameworks
→ implement behind AHDK
→ run conformance/evals/faults
→ choose for that Harness
```

Framework diversity is acceptable because Aurora standardizes the external contract, authority, evidence and observability—not the internal reasoning engine.
