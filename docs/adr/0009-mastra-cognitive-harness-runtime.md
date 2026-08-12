---
id: ADR-AURORA-0009
title: Mastra as Preferred First-Party Cognitive and Harness Runtime Substrate
document_type: adr
form: explanation
authority: decision
status: accepted
accepted_at: 2026-08-07
acceptance_evidence: DOC-AURORA-M0-R4-ADR-ACCEPTANCE-SPK001-AUTHORIZATION
version: 0.1.1
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - proposed default runtime posture for first-party agentic Harnesses
related:
  - DOC-AURORA-BLUEPRINT-06
  - DOC-AURORA-BLUEPRINT-07
  - DOC-AURORA-BLUEPRINT-10
  - DOC-AURORA-BLUEPRINT-12
  - RESEARCH-AURORA-MASTRA-COGNITIVE-HARNESS-R4-V1
  - DESIGN-AURORA-M0-R4-MASTRA-FIT-MATRIX
  - ADR-AURORA-0003
  - ADR-AURORA-0004
supersedes: []
superseded_by: null
last_reviewed: 2026-08-12
---

# ADR-0009 — Mastra as Preferred First-Party Cognitive and Harness Runtime Substrate

## Context

Aurora needs rich agentic infrastructure across future Capabilities: long-session context, agent loops, local workflows, human-in-the-loop, workspace/tool execution, skills, RAG, subagents, external Harness bindings, scheduling, signals, observability and evaluation.

Building generic versions of all of these inside Aurora would consume substantial engineering effort while creating a second framework surface to maintain.

Current Mastra research shows that the open-source framework now provides a broad, coherent implementation surface for these concerns while remaining deployable as a separate TypeScript/Node runtime with HTTP/OpenAPI exposure and replaceable storage/provider integrations.

At the same time, accepted Aurora Blueprints require identity, Project/Mission state, authority, global budgets, provider trust, governed memory semantics and acceptance to remain Aurora-owned and framework-replaceable.

## Decision drivers

- maximize reuse of mature/open generic agent infrastructure;
- avoid rebuilding AgentController/memory/workflow/RAG/workspace/interoperability plumbing without a requirement-driven reason;
- preserve Aurora's sovereign truth/authority/evidence boundary;
- preserve provider/runtime replaceability;
- allow different runtimes where a specific Capability has stronger requirements;
- avoid delaying M0 for infrastructure that M0 does not consume;
- contain Mastra's rapid release cadence behind a pinned adapter boundary;
- avoid mandatory dependence on Enterprise Edition or hosted Platform services.

## Options

### A — build Aurora-native generic agent runtime infrastructure

Pros: maximum internal control.

Cons: duplicates large amounts of undifferentiated infrastructure; higher maintenance burden; slower product delivery; likely reproduces concepts already implemented and evolving in Mastra.

### B — use Mastra selectively as isolated libraries with no preferred posture

Pros: low commitment.

Cons: encourages repeated per-Capability framework selection and may still lead to duplicated Harness/memory/workflow infrastructure.

### C — prefer Mastra as first-party cognitive/Harness runtime substrate behind Aurora contracts

Pros: high reuse; one coherent agentic substrate; strong memory/workflow/workspace/interoperability surface; preserves Core sovereignty if boundaries are enforced.

Cons: framework velocity and Node/TypeScript operational dependency; local state overlap must be governed; exact version upgrades require conformance testing.

### D — make Mastra the Aurora Core/control plane

Pros: maximal framework reuse.

Cons: violates accepted ownership boundaries by allowing framework state/semantics to become global product authority; creates unacceptable lock-in and recovery ambiguity.

## Decision

**Accepted decision:** choose **Option C**.

For first-party **agentic Harnesses and cognitive execution capabilities**, Mastra becomes the **preferred default substrate to evaluate first** before Aurora builds generic equivalent infrastructure.

This means future Capability work should start with the question:

> Can the required local agent/runtime behavior be implemented safely using current OSS Mastra primitives behind Aurora contracts?

If yes, reuse them. If no, record the capability-specific reason and select another mechanism.

Mastra is **not** the Aurora Sovereign Core and does not become mandatory for non-agentic, device, hard-real-time, firmware or other capabilities where its runtime profile is not appropriate.

## Required ownership boundary

Mastra MAY own/provider-localize:

```text
agent loops and model routing
AgentController sessions/modes/tasks
thread/message/tool history
observational memory
semantic recall/retrieval machinery
local workflow snapshots
local durable/background runs
workspaces and skills
signals/schedules/goals
provider-local tracing/evals
ACP/A2A/MCP execution bindings
```

Mastra MUST NOT own Aurora's canonical:

```text
identity
Project current state
Mission/Delegation global state
Authority Grant
global budgets/guardrails
provider approval/trust
governed memory authority semantics
artifact/evidence identity
global Verdict/Outcome
effect permission
```

A Mastra-local `passed`, `approved`, `goal complete`, `thread state`, `workflow snapshot` or `tool permission` is never automatically an Aurora global decision.

## Integration boundary

Preferred direction:

```text
Go Sovereign Core
→ Aurora-owned provider contract
→ MastraProviderAdapter
→ HTTP/OpenAPI or another accepted binding
→ version-pinned Mastra runtime
```

Mastra TypeScript/internal storage types MUST NOT leak into Aurora Core domain contracts.

Material tool/effect calls should use Mastra interception points only as adapters into Aurora's current Authority/PDP/effect boundary.

## Memory interpretation

Mastra Observational Memory is a strong implementation candidate for Aurora L3 observational/session synthesis. Message/tool history is a strong L4 candidate.

Mastra Memory Extractors may produce structured `MemoryCandidate` inputs, but Aurora governance owns provenance, epistemic status, temporal validity, scope, sensitivity, supersession and promotion into governed memory.

Mastra Working Memory terminology must not be imported directly into Aurora because the products use different conceptual scopes.

## Durability interpretation

Mastra workflows, Durable Agents, background tasks and Temporal integration MAY later implement Harness-local durability or adapters behind Aurora's `DurableExecutionPort`.

Their execution history never becomes canonical Aurora Project/Mission state automatically.

M0 continues to require no Mastra process or Mastra durable engine.

## Licensing / deployment rule

Aurora's foundational Mastra usage should rely on OSS code under the repository's Apache-2.0 boundary. EE or hosted Platform features may be used only as optional replaceable conveniences under their applicable terms.

No constitutional continuity, identity, authority or recovery path may depend exclusively on a paid/EE/platform feature.

## Validation before first implementation

This ADR does **not** require a Mastra integration spike during M0 because M0 does not consume Mastra.

Before the first Mastra-backed Capability is authorized for implementation, that Capability's readiness path must prove the relevant boundary, including at minimum where applicable:

- Core canonical state survives Mastra process/storage loss;
- provider-local run/state reconciles explicitly after restart;
- tool/effect permission consults Aurora authority rather than Mastra memory alone;
- cross-authority subagent/provider requests create Aurora child Delegations;
- Mastra version/API is pinned and contract-tested;
- any EE/platform dependency is explicitly classified and optional or accepted.

The exact proof should be scoped to the first consuming Capability rather than built speculatively now.

## Consequences

### Positive

- dramatically reduces generic agent-infrastructure reimplementation risk;
- provides a coherent Golden Path for first-party Harnesses;
- keeps Mastra's rapid innovation available without making it constitutional truth;
- makes Go Core + TypeScript agent runtime an intentional polyglot boundary rather than accidental architecture;
- preserves alternate runtimes for capability-specific needs.

### Negative

- adds a Node/TypeScript runtime and dependency lifecycle when Mastra-backed capabilities begin;
- requires explicit state ownership classification to avoid duplicated truth;
- version upgrades need conformance/compatibility checks;
- developers must resist importing Mastra terminology/types into Aurora domain semantics.

### Risks

The largest risk is not Mastra failure; it is allowing convenient Mastra persisted state to gradually become de facto global Aurora state. Architecture tests, contracts and review must keep the boundary visible.

## Compatibility / migration / rollback

Because Aurora owns the provider contract, a later accepted decision may replace Mastra for one or all Harness classes without changing Core domain meaning.

Provider-local historical/memory data may require migration or may be intentionally non-portable depending on its class. Any data promoted into governed Aurora memory/artifacts/evidence must use Aurora-owned identity/provenance before it becomes a sovereign dependency.

## M0 materiality

This proposal **does not add a new M0 blocker**.

The current M0 storage/recovery proof remains the next executable architecture experiment if separately authorized. Mastra should not be introduced into the M0 Golden Proof merely to validate a future substrate.

## Reconsideration triggers

- Mastra's OSS/license boundary materially changes;
- a Mastra major/runtime change removes required replaceability or server boundary;
- first consuming Capability exposes a material conformance/security/recovery failure;
- another runtime demonstrates substantially better fit for the same reusable Harness class;
- Mastra types/state begin leaking into Aurora canonical contracts;
- operational burden becomes disproportionate to reuse value.

## References

- `RESEARCH-AURORA-MASTRA-COGNITIVE-HARNESS-R4-V1`
- `DESIGN-AURORA-M0-R4-MASTRA-FIT-MATRIX`
- Blueprint 06 — Memory, Knowledge and Context
- Blueprint 07 — Harness Orchestration and Delegation
- Blueprint 10 — Autonomy, Authority and Safety
- Blueprint 12 — System Architecture and Component Boundaries
