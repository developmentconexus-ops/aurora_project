---
id: REVIEW-AURORA-M0-R4-MASTRA-MATERIALITY-2026-08-07
title: M0 R4 Mastra Materiality and Scope Review
document_type: adversarial_review
form: explanation
authority: evidence
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - review of Mastra finding impact on current M0 R4 blockers and scope
related:
  - DOC-AURORA-STATUS
  - DESIGN-AURORA-M0-R4-DECISION-LANDSCAPE
  - DOC-AURORA-CAP-SOVEREIGN-CORE-R4-DECISION-COVERAGE
  - RESEARCH-AURORA-MASTRA-COGNITIVE-HARNESS-R4-V1
  - DESIGN-AURORA-M0-R4-MASTRA-FIT-MATRIX
  - ADR-AURORA-0003
  - ADR-AURORA-0004
  - ADR-AURORA-0009
source_revision: 0efcfd6fbef1f4dd77ad23f7f62db27dacc67567
reviewed_at: 2026-08-07
last_reviewed: 2026-08-07
---

# M0 R4 — Mastra Materiality and Scope Review

## 1. Verdict

```text
MASTRA FINDING: MATERIAL CROSS-HORIZON
NEW M0 BLOCKER: NO
M0 R4 VERDICT: REMAINS BLOCKED FOR EXISTING SPK-001/SPK-002 + ADR ACCEPTANCE REASONS
```

The focused Mastra research materially improves Aurora's long-horizon implementation strategy but does not justify delaying the Sovereign Core M0 path.

## 2. What changed

Current Mastra evidence shows a substantially richer substrate than the earlier generic framework comparison captured:

- AgentController for persistent interactive Harness experiences;
- Observational Memory and structured Memory Extractors;
- workflows with suspend/resume and persisted/stored workflow definitions;
- Durable Agents and background tasks;
- Signals and schedules;
- Workspaces and Skills;
- RAG;
- A2A, ACP and SDK subagents;
- tool hooks;
- tracing/evals;
- Software Factory composition.

This is enough to justify a new proposed cross-horizon default:

```text
Mastra
→ evaluate first for generic first-party agentic Harness infrastructure
```

It is not enough to make Mastra the Aurora Core or authority plane.

## 3. Constitutional compatibility

The proposal is compatible with accepted Blueprints only under this boundary:

```text
Aurora Sovereign Core
owns:
identity
Project/Mission/Delegation global state
authority
budgets/provider trust
governed memory semantics
artifact/evidence identity
global acceptance/verdict

Mastra Harness runtime
may own locally:
agent loop/session/thread
local tasks/modes
message/tool history
observational synthesis
retrieval
workflow snapshot/durable run
workspace/skills
signals/schedules/goals
provider-local tracing/evals
```

Mastra-local convenience state cannot become global truth by accumulation or repeated use.

## 4. Effect on ADR-0003

Finding: the earlier wording could be read too broadly even though the body already allowed other Capability languages.

Correction:

```text
Go
→ proposed runtime for the Aurora Sovereign Core
NOT
→ mandatory runtime for all Aurora capabilities
```

This strengthens rather than reverses the original decision.

A future Mastra/TypeScript runtime is an intentional provider/Harness boundary.

## 5. Effect on ADR-0004

Mastra now provides concrete future implementations for workflow/durable execution, but M0 still contains no accepted need for them.

Therefore:

```text
M0
→ one local modular Sovereign Core
→ no Mastra process required
→ no workflow engine required

future consuming capability
→ may use Mastra workflow/Durable Agent/Temporal adapter
→ provider-local execution history remains distinct from Aurora global state
```

The original M0 non-selection remains valid.

## 6. ADR-0009 assessment

ADR-0009 is appropriately **proposed**, not accepted.

It makes a reusable engineering-default claim:

> Before building generic first-party agent infrastructure, evaluate current OSS Mastra primitives behind Aurora contracts.

This is preferable to both extremes:

### Reinvention extreme

Build AgentController, observational memory, workflow/HITL, RAG, workspace, signals, scheduling, interoperability and agent observability from scratch.

### Framework-capture extreme

Let Mastra own global Aurora state, authority, memory governance or acceptance because its local runtime already persists related concepts.

ADR-0009 chooses reuse with sovereignty.

## 7. Why no Mastra spike now

A spike is justified only when its answer can change the current build/architecture decision.

M0's Golden Proof requires:

```text
identity/state/authority continuity
with no external model/Harness as authority
```

Mastra is not in the M0 execution path.

A Go↔Mastra integration spike now would answer a future implementation question while delaying the currently required store/recovery proof. That violates the agreed materiality rule.

Therefore the correct proof trigger is:

```text
first Mastra-backed Capability reaches implementation horizon
→ pin exact Mastra version
→ define exact Aurora provider contract
→ run bounded boundary/conformance proof
```

The minimum future proof should demonstrate that killing/deleting Mastra-local state cannot erase Aurora canonical identity/state/authority and that effect/subprovider requests consult Aurora authority/delegation boundaries.

## 8. Product-velocity check

The review explicitly checks for process overreach.

### Required now

- record focused Mastra research;
- record the ownership/fit matrix;
- refine Go/M0 ADR scope;
- record proposed preferred-substrate ADR;
- continue existing M0 evidence path.

### Not required now

- install/build a Mastra prototype;
- benchmark every Mastra memory/storage provider;
- choose the future Mastra storage backend;
- choose the first AHDK language;
- implement Go↔Mastra transport;
- test Temporal through Mastra;
- design M1/M2/M3 in advance;
- choose hosted Mastra Platform/EE products.

This is the stopping point for Mastra investigation in M0.

## 9. Gate effect

The 15 original M0 R4 questions remain 15/15 accounted for.

Mastra changes the interpretation of Core/runtime and future durable-execution choices but does not create another M0 implementation-blocking question.

Existing R4 blockers remain:

1. `SPK-AURORA-M0-SOVEREIGN-STORE-001` executable evidence;
2. `SPK-AURORA-M0-OWNER-TRUST-002` executable evidence after SPK-001;
3. required operator acceptance of sufficiently evidenced M0 ADRs.

Therefore:

```text
R4 BLOCKED
```

remains correct, but **Mastra is not a reason for the BLOCKED status**.

## 10. Exact next product-oriented sequence

```text
Mastra research/findings recorded
→ stop Mastra investigation for M0
→ operator reviews revised ADR-0003/0004 plus ADR-0005/0006
→ ADR-0009 may be reviewed separately as cross-horizon direction
→ separately authorize SPK-AURORA-M0-SOVEREIGN-STORE-001 when ready
→ execute the actual M0 persistence proof
→ continue toward R4 PASS
```

No R5, implementation or Mastra integration work is authorized by this review.
