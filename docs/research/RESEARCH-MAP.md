---
id: DOC-AURORA-RESEARCH-MAP
title: Aurora Research Map
document_type: research_map
form: reference
authority: research
status: current
version: 0.2.0
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
last_reviewed: 2026-08-06
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

## 3. Current A0 reports

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

Current totals are tracked by CI; at the remediation checkpoint the focused program contains 9 reports/manifests and 92 primary-source entries.

## 4. Interpretation boundaries

### Memory

The research validates a hybrid and evaluated memory direction, not one selected storage engine or retrieval algorithm.

Still open:

- exact strata implementation;
- observational memory role;
- structured store versus event log/graph/vector mix;
- consolidation policy;
- deletion semantics across derived representations;
- Aurora-specific long-horizon evaluation thresholds.

### Interoperability

The research supports:

```text
MCP
→ strong candidate for tools/resources and bounded calls

A2A
→ strong candidate for remote opaque task-oriented providers

Native AHDK / local RPC
→ strong candidate for first-party local integration
```

It does not accept a binding or mapping until contract and conformance spikes pass.

### AHDK

The research supports a mandatory first-party Golden Path by policy while preserving contract/SDK independence.

Still open:

- first language;
- source schema and code-generation toolchain;
- SDK module boundaries;
- version matrix;
- scaffolder implementation.

### Durability

Temporal, DBOS, Restate and other mechanisms remain candidates. The accepted architectural direction is a `DurableExecutionPort`, not a selected engine.

### Authority and security

Cedar, OPA, OAuth Token Exchange and SPIFFE provide useful models. No product mechanism is selected. The constitutional boundary is separation of identity, decision, enforcement, containment and audit.

### Frameworks

Frameworks may implement Harness internals. None currently owns global Aurora state, identity, authority, memory governance or physical safety.

## 5. Required research before technical commitment

| Product area | Required investigation before R4 | Expected evidence |
|---|---|---|
| Sovereign Core storage and recovery | local-first stores, event/state ownership, backup/restore | crash/restart/restore spike |
| Memory and Context Builder | retrieval, consolidation, temporal/authority conflict, scale | benchmark/eval suite and adversarial journeys |
| AHDK source model | schema/codegen/SDK/conformance alternatives | same capability with SDK and direct implementation |
| Harness protocol binding | MCP/A2A/native mappings and version maturity | official TCK plus Aurora conformance |
| Durable execution | restart, idempotency, timers, cancellation, operational burden | comparative spike |
| Authority and effects | delegated identity, token lifetime, revocation, gateway and broker | denied/allowed/revoked effect drill |
| Presence and device trust | pairing, attestation, handoff, environment privacy | multi-presence threat model and prototype |
| Voice/vision | latency, local/cloud privacy, interruption and modality failure | user journey/evals |
| Laboratory integration | device identity, telemetry, calibration and interlocks | simulator/HIL and safety drill |
| Self-improvement | causal investigation, holdouts, canary, rollback | closed-loop candidate evaluation |

## 6. Freshness policy

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

A stale report remains discoverable but cannot support a new technical commitment without revalidation.

## 7. Research gaps beyond current A0

These are intentional future programs, not untracked omissions:

- Presence Fabric and device/workload trust;
- local-first synchronization and backup architecture;
- multimodal voice, vision and spatial interaction;
- laboratory instrument protocols and hardware safety standards;
- model routing, capability matching and cost/latency optimization;
- Artifact/Evidence Store architecture;
- operational data model and event/state reconciliation;
- private model execution and confidential computing where relevant;
- privacy, retention and third-party data handling in ambient environments;
- physical test design, calibration and measurement uncertainty.

## 8. Promotion path

```text
Research question
→ focused report + sources manifest
→ alternatives and limitations
→ Architecture Spike when necessary
→ proposed ADR / Capability requirement
→ review
→ accepted or rejected decision
```

A report should explicitly name which decision it can inform and which claims remain insufficient. Popularity, number of GitHub stars or framework convenience is never enough to promote a product decision.
