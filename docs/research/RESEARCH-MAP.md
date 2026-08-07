---
id: DOC-AURORA-RESEARCH-MAP
title: Aurora Research Map
document_type: research_map
form: reference
authority: research
status: current
version: 0.4.1
owners:
  - developmentconexus-ops
source_of_truth_for:
  - research discovery
  - research freshness and scope
  - current research gaps
related:
  - DOC-AURORA-DOCUMENTATION-MAP
  - DOC-AURORA-PRODUCT-INDEX
  - RESEARCH-AURORA-HARNESS-ARCHITECTURE-V1
  - DOC-AURORA-M0-R4-OPERATOR-AUTHORIZATION
  - RESEARCH-AURORA-MASTRA-COGNITIVE-HARNESS-R4-V1
last_reviewed: 2026-08-07
---

# Aurora Research Map

## 1. Authority notice

Research is evidence, not product authority.

A report may:

- describe the state of a field;
- compare mechanisms;
- identify risks and gaps;
- recommend experiments;
- inform a proposed ADR.

It may not, by itself:

- select the Aurora stack;
- authorize implementation;
- grant authority;
- change a Product Blueprint invariant;
- promote experimental code to production.

## 2. Research method

Material reports use:

- primary specifications and standards;
- official documentation and repositories;
- official release notes;
- research papers for scientific claims;
- explicit access dates and versions;
- a matching `*.sources.json` manifest;
- claim-to-source references such as `[S01]`;
- limitations and Aurora-specific implications;
- architecture spike recommendations when documentary evidence is insufficient.

The report owns interpretation. The source manifest owns source identity and supported-claim metadata.

Materiality rule for implementation-near work:

```text
Investigate deeply when the answer changes the next architecture/build decision.
Otherwise record the future trigger and keep building.
```

## 3. A0 baseline reports

| Research ID | Report | Primary question | Role | Reviewed |
|---|---|---|---|---|
| RESEARCH-AURORA-HARNESS-ARCHITECTURE-V1 | [Harness Architecture synthesis](AURORA-RESEARCH-HARNESS-ARCHITECTURE-v1.md) | What global architecture can coordinate heterogeneous Harnesses without framework lock-in? | cross-topic synthesis; focused reports are preferred for detailed decisions | 2026-08-05 |
| RESEARCH-AURORA-MEMORY-CONTEXT-V1 | [Memory and Context](AURORA-RESEARCH-MEMORY-CONTEXT-v1.md) | How should Aurora write, consolidate, retrieve, supersede, forget and evaluate memory? | informs Blueprint 06 and future CAP-MEMORY-CONTEXT | 2026-08-05 |
| RESEARCH-AURORA-HARNESS-INTEROPERABILITY-V1 | [Harness Interoperability](AURORA-RESEARCH-HARNESS-INTEROPERABILITY-v1.md) | How should Aurora map its Delegation semantics to MCP, A2A, RPC and direct data channels? | informs Blueprint 07 and ADR-0001 | 2026-08-05 |
| RESEARCH-AURORA-AHDK-CONFORMANCE-GOLDEN-PATHS-V1 | [AHDK, Conformance and Golden Paths](AURORA-RESEARCH-AHDK-CONFORMANCE-GOLDEN-PATHS-v1.md) | How can first-party Harness creation become easy without making the SDK the specification or security boundary? | informs Blueprint 05 and ADR-0002 | 2026-08-05 |
| RESEARCH-AURORA-DURABLE-EXECUTION-V1 | [Durable Execution](AURORA-RESEARCH-DURABLE-EXECUTION-v1.md) | How can Missions and Delegations survive process failure without duplicate effects? | informs Blueprint 07/12 and SPK-004 | 2026-08-05 |
| RESEARCH-AURORA-AUTHORITY-IDENTITY-EFFECTS-V1 | [Authority, Identity and Effects](AURORA-RESEARCH-AUTHORITY-IDENTITY-EFFECTS-v1.md) | How should subject, actor, executor, policy decision, credentials and effect enforcement compose? | informs Blueprint 10/11/12 and SPK-005 | 2026-08-05 |
| RESEARCH-AURORA-EVENTS-OBSERVABILITY-SCHEMAS-V1 | [Events, Observability and Schemas](AURORA-RESEARCH-EVENTS-OBSERVABILITY-SCHEMAS-v1.md) | How should events, schemas, traces and compatibility remain portable across runtimes? | informs Blueprint 05/07/13 and SPK-006 | 2026-08-05 |
| RESEARCH-AURORA-AGENT-FRAMEWORKS-RUNTIMES-V1 | [Agent Frameworks and Runtimes](AURORA-RESEARCH-AGENT-FRAMEWORKS-RUNTIMES-v1.md) | What roles can Mastra, LangGraph, Pi, OpenHands, OpenAI Agents SDK and Langflow play inside Harnesses? | prevents a framework from becoming the Aurora constitution | 2026-08-05 |
| RESEARCH-AURORA-EVALUATION-SELF-IMPROVEMENT-V1 | [Evaluation and Self-Improvement](AURORA-RESEARCH-EVALUATION-SELF-IMPROVEMENT-v1.md) | How can Aurora distinguish real improvement from local patching, overfitting and self-approval? | informs Blueprint 13 and future self-improvement capability | 2026-08-05 |

The accepted A0 baseline contains 9 reports/manifests and 92 primary-source entries as recorded by validation at A0/R3 checkpoints.

## 4. Current M0 R4 focused reports

These reports refresh implementation-near evidence for `M0 — Sovereign Core Walking Skeleton` and material cross-horizon findings discovered while making M0 technical decisions. They remain research evidence and do not make their recommendations governing.

| Research ID | Report | R4 decisions informed | Evidence boundary | Reviewed |
|---|---|---|---|---|
| RESEARCH-AURORA-M0-RUNTIME-PERSISTENCE-R4-V1 | [Runtime, Persistence and Durable Execution](AURORA-RESEARCH-M0-RUNTIME-PERSISTENCE-R4-v1.md) | Core runtime, store, state/event, atomicity, audit placement, topology, durable-engine applicability | documentary support for Go/local-state shape; SPK-AURORA-M0-SOVEREIGN-STORE-001 has now closed PASS and informs ADR-0007 v0.2.0 | 2026-08-07 |
| RESEARCH-AURORA-M0-PORTABILITY-INTEGRITY-R4-V1 | [Portability, Schema, Integrity and Export](AURORA-RESEARCH-M0-PORTABILITY-INTEGRITY-R4-v1.md) | logical schema, export, migration, digest/integrity boundary | documentary support for JSON Schema/JCS/age; authenticated local trust-anchor key custody remains coupled to SPK-002 | 2026-08-07 |
| RESEARCH-AURORA-M0-OWNER-AUTHORITY-RECOVERY-R4-V1 | [Owner Root, Authority Freshness and Recovery Trust](AURORA-RESEARCH-M0-OWNER-AUTHORITY-RECOVERY-R4-v1.md) | owner auth/bootstrap, authenticated integrity, time rollback, restore freshness | cryptographic architecture supported; cross-file crash/rollback protocol requires SPK-AURORA-M0-OWNER-TRUST-002 | 2026-08-07 |
| RESEARCH-AURORA-M0-OBSERVABILITY-R4-V1 | [Observability Mechanism](AURORA-RESEARCH-M0-OBSERVABILITY-R4-v1.md) | telemetry mechanism/backend boundary | documentary support sufficient; backend/export remains optional | 2026-08-07 |
| RESEARCH-AURORA-MASTRA-COGNITIVE-HARNESS-R4-V1 | [Mastra as Cognitive and Harness Runtime](AURORA-RESEARCH-MASTRA-COGNITIVE-HARNESS-R4-v1.md) | scope of Go Core decision, future Harness/cognitive runtime posture, memory/workflow/interoperability reuse | strong cross-horizon fit; deliberately no M0 integration spike because M0 does not consume Mastra | 2026-08-07 |

## 5. Interpretation boundaries

### Memory

The research validates a hybrid and evaluated memory direction, not one selected storage engine or retrieval algorithm.

The focused Mastra study now adds a concrete reuse posture:

```text
L3 observational/session synthesis
→ Mastra Observational Memory is a strong default candidate

L4 exact thread/tool history
→ Mastra is a strong provider-local candidate

L2 governed durable memory
→ Mastra may extract/store candidates, but Aurora owns provenance, epistemic/temporal governance and promotion

L0/L1 authoritative state/snapshots
→ never Mastra-owned merely because they enter model context
```

Still capability-specific/future:

- exact governed-memory physical store;
- consolidation/promotion policy implementation;
- deletion semantics across provider-local and governed representations;
- Aurora-specific long-horizon evaluation thresholds.

### Interoperability

The research supports:

```text
MCP
→ strong candidate for tools/resources and bounded calls

A2A
→ strong candidate for remote opaque task-oriented providers

ACP
→ strong candidate for coding Harness bindings

Native AHDK / local RPC / HTTP
→ strong candidates for first-party local integration
```

Mastra now provides current A2A/ACP/SDK-subagent implementation surfaces, making it a strong binding substrate. It still cannot define Aurora Delegation authority semantics.

### AHDK / Harness runtime

The research supports a mandatory first-party Golden Path by policy while preserving contract/SDK independence.

Accepted ADR-0009 establishes Mastra as the preferred-first agentic Harness substrate to evaluate, not as the specification. AHDK remains Aurora-owned and may wrap Mastra provider/runtime APIs.

Still open until the first consuming Capability:

- exact AHDK first language/package split;
- Go↔Mastra provider binding shape;
- generated contract mapping;
- conformance suite implementation;
- pinned Mastra version matrix.

### Durability

Temporal, DBOS, Restate and other mechanisms remain candidates for future long-running execution. Mastra now adds Durable Agents, workflow snapshots and Temporal integration as concrete future adapter candidates.

R4 M0 research still finds no current requirement justifying a durable workflow engine in M0. That is a proposed non-selection, not a permanent rejection.

### Authority and security

Cedar, OPA, OAuth Token Exchange and SPIFFE remain future effect/delegation models. Mastra Tool Hooks/AgentController permissions are useful local enforcement integration points but cannot mint Aurora authority.

M0 focused research addresses only the local owner-root/current-state recovery slice and does not implement the later effect plane.

### Frameworks

Mastra is now the strongest current preferred candidate for generic first-party agentic Harness infrastructure, but no framework owns global Aurora state, identity, authority, memory governance, global verdict or physical safety.

## 6. Required research before technical commitment

| Product area | Required investigation before R4 commitment | Expected evidence | Current M0 status |
|---|---|---|---|
| Sovereign Core storage and recovery | local-first stores, event/state ownership, backup/restore | crash/restart/restore spike | `SPK-AURORA-M0-SOVEREIGN-STORE-001` PASS/CLOSED; SQLite + modernc evidence informs ADR-0007 v0.2.0 |
| M0 owner root / rollback / restore freshness | key custody, trust anchor, backward time, historical authority restore | owner-root/recovery fault spike | documentary research complete; `SPK-AURORA-M0-OWNER-TRUST-002` specified, execution not authorized |
| Mastra-backed first-party Harness boundary | Core/provider state ownership, authority interception, restart/reconciliation, contract/version isolation | bounded conformance proof in first consuming Capability | research complete; intentionally **not an M0 blocker** |
| Memory and Context Builder | retrieval, consolidation, temporal/authority conflict, scale | benchmark/eval suite and adversarial journeys | future M1; Mastra is now a strong substrate candidate |
| AHDK source model | schema/codegen/SDK/conformance alternatives | same capability with SDK and direct implementation | future M2 |
| Harness protocol binding | MCP/A2A/ACP/native mappings and version maturity | official TCK plus Aurora conformance | future consuming capability |
| Durable execution | restart, idempotency, timers, cancellation, operational burden | comparative spike | not required by M0; reconsider at M4/current requirement trigger |
| Authority and effects | delegated identity, token lifetime, revocation, gateway and broker | denied/allowed/revoked effect drill | future effect plane |
| Presence and device trust | pairing, attestation, handoff, environment privacy | multi-presence threat model and prototype | future M8 |
| Voice/vision | latency, local/cloud privacy, interruption and modality failure | user journey/evals | future |
| Laboratory integration | device identity, telemetry, calibration and interlocks | simulator/HIL and safety drill | future M9/M10 |
| Self-improvement | causal investigation, holdouts, canary, rollback | closed-loop candidate evaluation | future M11 |

## 7. Freshness policy

Revalidate a report when any of the following occurs:

- a cited specification or stable protocol version changes materially;
- an official SDK reaches or leaves stable support;
- a framework changes its persistence, memory, security or runtime model;
- a source is withdrawn or contradicted;
- an Architecture Spike produces conflicting evidence;
- an ADR is proposed for implementation;
- the relevant Product Milestone enters R3/R4;
- six months pass for a fast-moving agent/protocol topic;
- twelve months pass for a slower standard, unless implementation risk requires earlier review.

For Mastra specifically, the first consuming Capability must re-pin and revalidate the exact core/memory/server packages because the framework is evolving rapidly.

A stale report remains discoverable but cannot support a new technical commitment without revalidation.

## 8. Research gaps beyond current M0

These are intentional future programs, not untracked omissions:

- Presence Fabric and device/workload trust;
- local-first multi-node synchronization;
- multimodal voice, vision and spatial interaction;
- laboratory instrument protocols and hardware safety standards;
- model routing, capability matching and cost/latency optimization;
- Artifact/Evidence Store architecture beyond M0 metadata/references;
- operational event/state reconciliation at distributed scale;
- private model execution and confidential computing where relevant;
- privacy, retention and third-party data handling in ambient environments;
- physical test design, calibration and measurement uncertainty.

## 9. Promotion path

```text
Research question
→ focused report + sources manifest
→ alternatives and limitations
→ Architecture Spike/conformance proof only when materially required
→ proposed ADR / Capability requirement
→ review
→ accepted or rejected decision
```

A report should explicitly name which decision it can inform and which claims remain insufficient. Popularity, number of GitHub stars or framework convenience is never enough to promote a product decision.
