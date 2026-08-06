---
id: DOC-AURORA-DOCUMENTATION-COVERAGE
title: Aurora Discovery and Documentation Coverage
document_type: traceability_matrix
form: reference
authority: tracking
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - current documentation coverage
related:
  - REVIEW-AURORA-A0-DOCUMENTATION-2026-08-05
  - DOC-AURORA-DOCUMENTATION-MAP
last_reviewed: 2026-08-05
---

# Aurora Discovery and Documentation Coverage

## 1. Purpose

This matrix prevents approved discovery work from remaining only in chat history.

Coverage states:

```text
FULL        mechanism, examples, boundaries and failure behavior preserved
PARTIAL     conclusion preserved but depth missing
MISSING     no canonical owner or material treatment
MISPLACED   present only in research/design/tracking instead of canonical owner
OPEN        intentionally unresolved technical decision
```

Tracking is not architecture. The canonical owner column determines where the repaired content must live.

## 2. Product identity and scope

| Discovery decision or concept | Before remediation | Canonical owner | Required depth |
|---|---:|---|---|
| Aurora as personal cognitive operating system | PARTIAL | Blueprint 01 | definition, role, boundaries, examples |
| Aurora is more than chatbot/voice/model/memory/device | PARTIAL | Blueprint 01 | component distinction and failure examples |
| North Star laboratory continuation experience | PARTIAL | Blueprint 01 / 04 | complete journey and evidence flow |
| Five verbs: understand, remember, reason, act, observe | PARTIAL | Blueprint 01 / 04 | closed-loop lifecycle |
| Leandro-first, single-user, no current multi-tenancy | FULL | Blueprint 01 | implications and future boundary |
| Personal intelligence broad; engineering first domain | PARTIAL | Blueprint 01 | domain expansion protocol |
| Vision complete, technical commitment short-horizon | PARTIAL | Blueprint 01 / 14 | two-horizon planning model |
| One Aurora independent of model/provider/interface | PARTIAL | Blueprint 01 / 12 | identity and replacement boundaries |

## 3. Human–Aurora relationship

| Discovery decision or concept | Before remediation | Canonical owner | Required depth |
|---|---:|---|---|
| Copilot intellectual confiável | PARTIAL | Blueprint 02 | obligations, disagreement, repair |
| Lealdade ao objetivo, não à premissa momentânea | PARTIAL | Blueprint 02 | examples and escalation |
| Leandro retains final authority | FULL | Blueprint 02 / 10 | decision categories and exceptions |
| Stable transparent AI identity | PARTIAL | Blueprint 02 | identity continuity and model substitution |
| Personality B+: JARVIS precision + EVIE proximity | PARTIAL | Blueprint 02 | traits, examples, anti-patterns |
| Personality with presence, not performance | PARTIAL | Blueprint 02 | behavior modes and quality criteria |
| Contextual humor | PARTIAL | Blueprint 02 | safe/unsafe contexts |
| Proactivity contextual and controlled | PARTIAL | Blueprint 02 | attention budget and notification policy |
| Initiative must be explainable | PARTIAL | Blueprint 02 / 13 | reason codes and audit |
| No emotional manipulation or simulated humanity | PARTIAL | Blueprint 02 / 11 | boundaries and incidents |

## 4. Presence and multimodality

| Discovery decision or concept | Before remediation | Canonical owner | Required depth |
|---|---:|---|---|
| Persistent event-oriented Core | PARTIAL | Blueprint 08 / 12 | lifecycle and topology |
| Availability does not mean surveillance | PARTIAL | Blueprint 08 / 11 | sensor modes and enforcement |
| One Aurora, multiple presences | PARTIAL | Blueprint 08 | complete presence model |
| Computer-to-glasses handoff | PARTIAL | Blueprint 08 | journey, authentication, context pack |
| Contextual continuity protected by environment | PARTIAL | Blueprint 08 / 11 | public/private scenarios |
| Device capability declaration | PARTIAL | Blueprint 08 / 09 | manifest example and permissions |
| Offline/degraded presence | MISSING | Blueprint 08 / 12 | local capabilities and reconciliation |
| Visible sensor indicators and retention | PARTIAL | Blueprint 08 / 11 | policy and audit |
| Adapt interface by device capabilities | MISSING | Blueprint 08 | presentation negotiation |

## 5. Memory, knowledge and context

| Discovery decision or concept | Before remediation | Canonical owner | Required depth |
|---|---:|---|---|
| Memory is a core subsystem | PARTIAL | Blueprint 06 | architecture and requirements |
| Memory is not context/history/knowledge/state | PARTIAL | Blueprint 06 | ownership and precedence |
| Working memory | PARTIAL | Blueprint 06 | write/read/expiry/evaluation |
| Conversational memory | PARTIAL | Blueprint 06 | exact history and compaction |
| Observational memory | PARTIAL | Blueprint 06 + Research | mechanism, evidence, limits |
| Episodic memory | PARTIAL | Blueprint 06 | event model and examples |
| Project memory | PARTIAL | Blueprint 06 | project universe and authority |
| Global personal memory | PARTIAL | Blueprint 06 | scope, sensitivity, confirmation |
| Relational memory/world graph | PARTIAL | Blueprint 03 / 06 | relationships and temporal validity |
| Procedural memory/Golden Paths | PARTIAL | Blueprint 06 / Capability Method | promotion and validation |
| Failure and learning memory | PARTIAL | Blueprint 06 / 13 | causal relationship |
| Operational memory must be structured | PARTIAL | Blueprint 06 / 12 | storage owner and recovery |
| Memory metadata example | MISSING | Blueprint 06 | complete structured example |
| Epistemic distinctions | PARTIAL | Blueprint 06 | conflict behavior and confidence |
| Promotion by risk, authority and scope | PARTIAL | Blueprint 06 | matrix and examples |
| Automatic vs conditioned promotion | PARTIAL | Blueprint 06 | policy scenarios |
| Supersession instead of silent overwrite | PARTIAL | Blueprint 06 | lifecycle and conflict examples |
| Forgetting, expiry, deletion and archive | PARTIAL | Blueprint 06 / 11 | policies and tests |
| Plausible but false memory | MISSING | Blueprint 06 / 13 | failure detection |
| Cross-project contamination | MISSING | Blueprint 06 / 11 | isolation test |
| Context Builder | PARTIAL | Blueprint 06 / 12 | pipeline, ranking and live checks |
| Local-first memory, cloud-assisted inference | PARTIAL | Blueprint 06 / 11 | data flow and audit |
| Memory evaluation research | MISSING | Research Memory | benchmarks and limits |

## 6. Autonomy, authority and campaigns

| Discovery decision or concept | Before remediation | Canonical owner | Required depth |
|---|---:|---|---|
| Progressive authority | PARTIAL | Blueprint 10 | action matrix and grants |
| Delegated autonomy by mission | PARTIAL | Blueprint 10 | complete campaign contract |
| N1 action, N2 workflow, N3 campaign, N4 continuous program | PARTIAL | Blueprint 10 | promotion criteria |
| Overnight AI workflow optimization | MISSING | Blueprint 04 / 10 | end-to-end scenario |
| Firmware variant campaign | MISSING | Blueprint 04 / 09 / 10 | end-to-end scenario |
| Autonomy Envelope | PARTIAL | Blueprint 10 | structured example |
| Baseline and immutable evaluation criteria | MISSING | Blueprint 10 / 13 | anti-gaming controls |
| Budget: time, cost, tokens, cycles, energy, wear | MISSING | Blueprint 10 | units and enforcement |
| Stop and escalation conditions | PARTIAL | Blueprint 10 | taxonomy and examples |
| Autonomous inside envelope, conservative at boundary | FULL | Blueprint 10 | invariants |
| Deterministic physical guardrails | PARTIAL | Blueprint 09 / 10 | interlocks and fault cases |
| Emergency authority | PARTIAL | Blueprint 10 | preauthorization and drills |
| Revocation has operational effect | PARTIAL | Blueprint 10 / 11 | partition and containment |

## 7. Self-improvement and learning

| Discovery decision or concept | Before remediation | Canonical owner | Required depth |
|---|---:|---|---|
| Experimental self-improvement with supervised promotion | PARTIAL | Blueprint 13 | full lifecycle |
| Continuous detection, authorized experiments | PARTIAL | Blueprint 13 | triage and prioritization |
| Symptom versus root cause | PARTIAL | Blueprint 13 | causal analysis method |
| Multiple errors may share one cause | MISSING | Blueprint 13 | causal graph example |
| Structured incident record | MISSING | Blueprint 13 | schema example |
| Correlation before patching | PARTIAL | Blueprint 13 | thresholds and evidence |
| Competing hypotheses | PARTIAL | Blueprint 13 | falsification and selection |
| Reproduction before correction | PARTIAL | Blueprint 13 | unreproducible incident policy |
| Candidate version, sandbox and rollback | PARTIAL | Blueprint 13 | lifecycle and artifacts |
| Original, neighbor, contrary, history, unseen, adversarial tests | PARTIAL | Blueprint 13 | evaluation suites |
| Holdout prevents self-selected evaluation | MISSING | Blueprint 13 | governance |
| Independent reviewer | PARTIAL | Blueprint 13 | authority and separation |
| Canary and shadow modes | PARTIAL | Blueprint 13 | rollout states |
| Constitutional areas protected | PARTIAL | Blueprint 10 / 13 | immutable boundary |
| Learn from successes into candidate Golden Paths | MISSING | Blueprint 13 / Capability Method | promotion stages |

## 8. Capability and harness model

| Discovery decision or concept | Before remediation | Canonical owner | Required depth |
|---|---:|---|---|
| Tool/agent/workflow/runtime/harness distinctions | PARTIAL | Blueprint 03 / 05 | examples and ownership |
| Harness as specialized complete system | PARTIAL | Blueprint 05 / 07 | domain boundary |
| Aurora as global control plane | PARTIAL | Blueprint 07 / 12 | component and state model |
| MNFS as future software provider | FULL | Blueprint 07 | readiness and adapter criteria |
| Aurora owns why/what/limits; harness owns local how | FULL | Blueprint 07 | invariants |
| Hierarchical orchestration | PARTIAL | Blueprint 07 | multi-harness journeys |
| Harness asks Aurora for another capability | PARTIAL | Blueprint 07 | child-delegation schema |
| Central control plane, authorized direct data plane | PARTIAL | Blueprint 07 / 12 | channel lifecycle |
| Capability Manifest | PARTIAL | Blueprint 05 | full example and validation |
| Discovery ≠ compatibility ≠ trust ≠ authority ≠ execution | FULL | Blueprint 05 | state machine |
| Multidimensional trust | PARTIAL | Blueprint 05 / 11 | evidence model |
| Trust bound to version/build/environment | PARTIAL | Blueprint 05 / 11 | provenance and upgrade |
| Dynamic discovery with governed activation | PARTIAL | Blueprint 05 | threat and fallback |
| Provider selection by fit, sensitivity, cost, latency and trust | PARTIAL | Blueprint 05 | explainability and fallback |
| Provider suspension/revocation/retirement | PARTIAL | Blueprint 05 | incident flows |

## 9. Contracts, SDK and interoperability

| Discovery decision or concept | Before remediation | Canonical owner | Required depth |
|---|---:|---|---|
| Aurora-owned semantics | PARTIAL | Blueprint 03 / 07 / ADR-0001 | complete entity definitions |
| Contracts mandatory | PARTIAL | Blueprint 05 / Capability Method | specification lifecycle |
| AHDK mandatory for first-party by policy | FULL | Blueprint 05 / ADR-0002 | waiver path |
| Contract independent from SDK | FULL | Blueprint 05 / ADR-0002 | conformance example |
| SDKs optional technically, not organizationally | PARTIAL | Blueprint 05 | policy and exception |
| AHDK as development kit, not client library | PARTIAL | Blueprint 05 / Research SDK | module architecture |
| Generated types | PARTIAL | Blueprint 05 | source and compatibility |
| Manifest builder | PARTIAL | Blueprint 05 | code example |
| Lifecycle API | PARTIAL | Blueprint 05 | usage example |
| Context API | PARTIAL | Blueprint 05 / 06 | access boundaries |
| Effect client does not authorize | PARTIAL | Blueprint 05 / 10 | gateway flow |
| Artifact/Evidence API | PARTIAL | Blueprint 05 / 07 | data types |
| Cancellation/checkpoint/heartbeat/budget | PARTIAL | Blueprint 05 / 07 | API behaviors |
| Automatic OpenTelemetry | PARTIAL | Blueprint 13 / Research Observability | semantic conventions |
| Simulator, mocks and fault injection | PARTIAL | Blueprint 05 / Capability Method | conformance path |
| Scaffolder and Golden Paths | PARTIAL | Blueprint 05 | template anatomy |
| MCP for tools/resources | PARTIAL | Blueprint 07 / Research Interop | mapping and limits |
| A2A for remote opaque tasks | PARTIAL | Blueprint 07 / Research Interop | mapping and maturity |
| Native SDK/RPC for first-party | PARTIAL | Blueprint 07 / Research Interop | trade-offs |
| Frameworks remain internal | FULL | Blueprint 07 | comparison matrix |
| Mastra/LangGraph/Pi/OpenHands/OpenAI Agents roles | MISPLACED | Research Frameworks | detailed comparison |

## 10. Durability, security and observability

| Discovery decision or concept | Before remediation | Canonical owner | Required depth |
|---|---:|---|---|
| Durable execution separate from agent runtime | PARTIAL | Blueprint 12 / Research Durability | component port and semantics |
| Temporal/DBOS/Restate/Inngest candidates | MISPLACED | Research Durability | focused comparison and spikes |
| SDK is not security boundary | FULL | Blueprint 10 / 11 | bypass examples |
| Policy decision separate from enforcement | PARTIAL | Blueprint 10 / 11 | architecture and receipts |
| Effect Gateways | PARTIAL | Blueprint 10 / 12 | gateway catalog |
| Sandbox/OS containment | PARTIAL | Blueprint 11 / 12 | threat model |
| Credential Broker | PARTIAL | Blueprint 11 / 12 | secret reference flow |
| Subject/actor/executor identity | PARTIAL | Blueprint 10 / 11 | token and audit example |
| Cedar/OPA/SPIFFE/RFC8693 candidates | MISPLACED | Research Authority | detailed comparison |
| OpenTelemetry baseline | PARTIAL | Blueprint 13 / Research Observability | signals, spans, redaction |
| CloudEvents/AsyncAPI/JSON Schema/Protobuf candidates | MISPLACED | Research Events | boundaries and compatibility |
| SLSA provenance for provider trust | PARTIAL | Blueprint 05 / 11 | supply-chain lifecycle |
| Distributed trace across Core/harness/tool/gateway | PARTIAL | Blueprint 13 | Golden Proof |

## 11. Documentation and roadmap

| Discovery decision or concept | Before remediation | Canonical owner | Required depth |
|---|---:|---|---|
| Documentation is project memory and governance | PARTIAL | Blueprint 15 | full lifecycle and checks |
| Blueprint before implementation | FULL | Blueprint 15 / STATUS | gate |
| Research before technical choice | FULL | Blueprint 15 | freshness and source rules |
| Two horizons: constitutional vision and next executable detail | PARTIAL | Blueprint 01 / 14 | planning model |
| Roadmap by cumulative capabilities and proofs | PARTIAL | Blueprint 14 | full milestone anatomy |
| Fresh-session Golden Proof | PARTIAL | Blueprint 14 | exact test procedure |
| Status for session handoff | PARTIAL | Blueprint 15 / STATUS | required fields |
| One concept, one canonical owner | FULL | Blueprint 15 / Documentation Map | conflict protocol |
| Conversation is discovery, repository is canonical memory | PARTIAL | Blueprint 15 / History Record | promotion path |
| Aggregate Blueprint publication | MISSING | Product Index | generated projection |
| Capability Realization Method | MISSING | Product Method | gates and traceability |
| Requirements traceability | MISSING | Product Matrix | bidirectional mapping |

## 12. Exit condition

This matrix becomes `FULL` for all approved discovery decisions before A0 can be accepted. Items marked `OPEN` may remain open only when their decision is deliberately deferred to research or an architecture spike.
