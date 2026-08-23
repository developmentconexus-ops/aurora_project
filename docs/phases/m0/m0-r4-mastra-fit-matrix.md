---
id: DESIGN-AURORA-M0-R4-MASTRA-FIT-MATRIX
title: M0 R4 Mastra Fit Matrix
document_type: architecture_fit_matrix
form: reference
authority: design
status: proposed
version: 0.1.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - proposed Mastra USE ADAPT WRAP DO_NOT_USE FUTURE mapping
related:
  - RESEARCH-AURORA-MASTRA-COGNITIVE-HARNESS-R4-V1
  - DESIGN-AURORA-M0-R4-DECISION-LANDSCAPE
  - ADR-AURORA-0003
  - ADR-AURORA-0004
source_revision: 0efcfd6fbef1f4dd77ad23f7f62db27dacc67567
last_reviewed: 2026-08-07
---

# M0 R4 — Mastra Fit Matrix

## 1. Decision rule

```text
USE
→ reuse directly as Harness-local mechanism

ADAPT
→ useful primitive, but Aurora semantics differ

WRAP
→ use only behind an Aurora-owned contract/authority/governance adapter

DO_NOT_USE_AS_OWNER
→ mechanism may participate, but cannot own this concept

FUTURE
→ promising but not material to the current M0 build
```

The goal is to prevent both reinvention and framework capture.

## 2. Capability matrix

| Mastra capability | Aurora category | Posture | Ownership rule | Current M0 effect |
|---|---|---|---|---|
| Agent runtime/model routing | Harness cognition | `USE` | Harness-local | none |
| AgentController | Harness interactive runtime | `USE` | sessions/modes/tasks are provider-local | none |
| exact thread/message/tool history | L4 history | `USE` | historical/supporting, not project truth | none |
| Observational Memory | L3 synthesis | `USE` | probabilistic/supporting | none |
| Semantic Recall | retrieval | `USE` | relevance is not authority | none |
| Mastra Working Memory | persistent structured context | `ADAPT` | do not map name directly to Aurora short-lived Working Memory | none |
| Memory Extractors | memory candidate production | `WRAP` | extracted values enter Aurora governance before promotion | none |
| RAG | knowledge retrieval | `USE/WRAP` | Aurora supplies source scope/freshness/authority policy | none |
| Workflows + suspend/resume | Harness-local execution | `USE` | workflow snapshot is not global Mission state | none |
| Stored declarative workflows | Harness definition/runtime | `USE` | versioned provider-local definition | none |
| Durable Agents | Harness-local durable run | `FUTURE` | may implement later execution port | none |
| Temporal integration | stronger durable execution | `FUTURE` | candidate adapter only when requirement justifies it | none |
| Background Tasks | Harness-local async work | `USE` | local task state only | none |
| Signals | reactive context/notification | `ADAPT` | input, not authority/truth | none |
| Schedules | recurring execution trigger | `ADAPT` | execution-time authority remains Aurora-owned | none |
| Goals | local objective loop | `ADAPT` | LLM judge cannot close global Mission | none |
| Tool Hooks | effect interception | `WRAP` | call Aurora PDP/Authority boundary for material effects | none |
| Workspaces | file/search/command environment | `USE` | containment only; not equivalent to OS sandbox | none |
| Skills | provider methodology/context | `USE` | local Harness procedure; normative Aurora standards remain external | none |
| ACP | coding-harness binding | `WRAP` | new authority boundary requires Aurora child Delegation | none |
| A2A | remote-agent binding | `WRAP` | transport does not define Delegation semantics | none |
| SDK subagents | provider composition | `WRAP` | local subagent permitted only inside granted envelope | none |
| MCP | tool/resource binding | `WRAP` | binding only; Aurora contracts govern semantics | none |
| Mastra traces/metrics/logs | provider observability | `USE` | evidence input, not canonical evidence/verdict | none |
| Evals/Gates/Verdicts | local quality evidence | `ADAPT` | Aurora acceptance remains independent | none |
| Factory | Software Harness composition/reference | `FUTURE/USE` | specialized Harness, not Aurora Mission Control | none |
| Server/OpenAPI | Go↔Mastra binding | `WRAP` | Aurora-owned provider contract outside Mastra API types | none |
| Composite Storage | Mastra-local persistence | `USE` | cannot become Core operational state owner | none |
| Managed Platform/EE features | deployment convenience | `FUTURE` | optional only; no constitutional dependency | none |

## 3. Hard ownership boundary

### Aurora owns

```text
Aurora identity
Project current state
Mission / Delegation global state
Authority Grant
budgets and global guardrails
provider approval/trust
governed memory semantics
artifact/evidence identity
global Verdict / Outcome
Effect permission
```

### Mastra may own locally

```text
agent loop
thread/session
modes/tasks
local workflow snapshot
local durable run
messages/tool history
observational synthesis
retrieval machinery
workspace/skills
provider-local schedule/goal/background task
provider traces/evals
```

## 4. Candidate process boundary

```text
┌────────────────────────────────────────────┐
│ AURORA SOVEREIGN CORE — Go                │
│ identity • project • mission • authority   │
│ governance • evidence • acceptance         │
└───────────────────┬────────────────────────┘
                    │ Aurora Provider Contract
                    ▼
             MastraProviderAdapter
                    │ HTTP/OpenAPI
                    ▼
┌────────────────────────────────────────────┐
│ MASTRA RUNTIME — TypeScript / Node         │
│ AgentController • memory • workflows       │
│ RAG • workspace • skills • subagents       │
│ signals • local evals/observability         │
└───────────────────┬────────────────────────┘
                    ▼
            models / ACP / A2A / MCP
```

## 5. Failure-domain invariant

A future Mastra-backed Capability is architecturally valid only if this remains true:

```text
kill Mastra
→ Aurora identity remains
→ Project truth remains
→ Authority remains
→ global Mission state remains queryable
→ provider becomes unavailable/degraded/reconciling

NOT
→ Aurora loses identity or governing state
```

Deleting or corrupting Mastra-local storage may lose local history/execution progress, but cannot erase Aurora's canonical identity/authority/project truth.

## 6. Materiality decision

Mastra is material to the **long-horizon runtime strategy** but is not consumed by M0.

Therefore:

```text
record architecture direction now
→ do not build generic equivalents from scratch casually
→ do not execute a Mastra integration spike during M0
→ continue M0 Sovereign Core evidence path
→ prove Mastra boundary when the first Mastra-backed capability enters implementation horizon
```

This prevents both a local maximum from rebuilding Mastra-like infrastructure and an analysis maximum from delaying the product before Mastra is actually required.
