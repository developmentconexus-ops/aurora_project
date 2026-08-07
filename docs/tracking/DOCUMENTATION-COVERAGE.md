---
id: DOC-AURORA-DOCUMENTATION-COVERAGE
title: Aurora Discovery and Documentation Coverage
document_type: traceability_matrix
form: reference
authority: tracking
status: current
version: 2.0.1
owners:
  - developmentconexus-ops
source_of_truth_for:
  - current documentation coverage
  - discovery-to-canonical traceability
related:
  - REVIEW-AURORA-A0-DOCUMENTATION-2026-08-05
  - DOC-AURORA-DOCUMENTATION-MAP
  - DOC-AURORA-REQUIREMENTS-TRACEABILITY
  - HISTORY-AURORA-ORIGIN-DISCOVERY-2026-08-05
last_reviewed: 2026-08-06
---

# Aurora Discovery and Documentation Coverage

## 1. Purpose

This matrix verifies that the initial Aurora discovery dialogue has been promoted into durable, reviewable repository artifacts.

It does not make the tracked material normative. The **canonical owner** column identifies which Product Blueprint section, research report, ADR or method owns the current meaning.

## 2. Coverage states

```text
FULL
→ definition, rationale, mechanisms, scenarios, boundaries, failure modes and proof intent are preserved

OPEN_RESEARCH
→ the required product behavior is documented, but the implementation mechanism needs focused research or a spike

OPEN_DECISION
→ alternatives and decision drivers are documented; no option has been promoted to an accepted ADR

DEFERRED_BY_ROADMAP
→ constitutionally preserved direction that is intentionally outside the current executable horizon
```

A0 cannot close with `MISSING`, `UNMAPPED` or `OPEN_GAP` coverage.

---

## 3. Product identity, scope and North Star

| Discovery topic | Current coverage | Canonical owner | Evidence of preservation |
|---|---|---|---|
| Aurora as personal cognitive operating system and global control plane | FULL | Blueprint 01, 12 | definition, layer diagram, component boundaries and non-goals |
| More than chatbot, voice, model, memory or device | FULL | Blueprint 01 | component distinction and failure patterns |
| Laboratory continuation North Star | FULL | Blueprint 01, 04, 09 | end-to-end journey from project recovery to experiment evidence |
| Understand, remember, reason, act and observe | FULL | Blueprint 01, 04 | expanded closed cognitive loop with verification, recording and learning |
| Leandro-first and single-user in current horizon | FULL | Blueprint 01 | implications, excluded SaaS concerns and future boundary |
| Broad personal intelligence with engineering first | FULL | Blueprint 01, 03 | initial domain map and controlled domain expansion |
| Complete vision with short-horizon technical commitment | FULL | Blueprint 01, 14, ACRM | two-horizon model and readiness gates |
| One Aurora independent of model, provider, process and interface | FULL | Blueprint 01, 08, 12 | identity continuity and replaceable engines/presences |

---

## 4. Human–Aurora relationship and personality

| Discovery topic | Current coverage | Canonical owner | Evidence of preservation |
|---|---|---|---|
| Trusted intellectual copilot | FULL | Blueprint 02 | duties, disagreement protocol, escalation and trust repair |
| Loyalty to objective, not momentary premise | FULL | Blueprint 02 | direct examples and decision categories |
| Leandro retains final authority | FULL | Blueprint 02, 10 | authority hierarchy and prohibited self-escalation |
| Stable and transparent AI identity | FULL | Blueprint 02, 08 | model substitution and anti-deception rules |
| J.A.R.V.I.S. precision plus E.V.I.E. proximity | FULL | Blueprint 02, History | trait model, inspiration boundary and original rationale |
| Personality with presence, not performance | FULL | Blueprint 02 | mode table, positive examples, anti-patterns and evaluation criteria |
| Contextual humor | FULL | Blueprint 02 | casual, engineering, laboratory, incident and high-risk behavior |
| Contextual, controlled proactivity | FULL | Blueprint 02, 13 | attention budget, notification intensity and reason codes |
| No simulated humanity or emotional manipulation | FULL | Blueprint 02, 11 | prohibited behavior and incident handling |

---

## 5. Presence, interaction and multimodality

| Discovery topic | Current coverage | Canonical owner | Evidence of preservation |
|---|---|---|---|
| Persistent event-oriented Core | FULL | Blueprint 08, 12 | topology and service/presence distinction |
| Availability does not imply surveillance | FULL | Blueprint 08, 11 | sensor modes, consent, indicators and retention |
| One Aurora, multiple presences | FULL | Blueprint 08 | Presence, Device and Session entities plus architecture diagram |
| Computer-to-glasses handoff | FULL | Blueprint 04, 08 | authentication, safe Context Pack and continuation scenario |
| Environment-protected continuity | FULL | Blueprint 08, 11 | private/public/unknown environment behavior |
| Device capability declaration | FULL | Blueprint 08, 09 | manifest example, trust and effective-authority calculation |
| Offline and degraded presence | FULL | Blueprint 08, 12 | local fallback, queued observations and reconciliation |
| Interface adaptation by capability | FULL | Blueprint 08 | modality negotiation and safe degradation |
| Voice, vision, screen and spatial input | DEFERRED_BY_ROADMAP | Blueprint 08, 14 | constitutional behavior defined; implementation research deferred |

---

## 6. Memory, knowledge and context

| Discovery topic | Current coverage | Canonical owner | Evidence of preservation |
|---|---|---|---|
| Memory as a core subsystem | FULL | Blueprint 06 | full architecture, lifecycle and evaluation model |
| Memory distinct from context, history, knowledge and operational state | FULL | Blueprint 03, 06 | ownership, precedence and source-of-truth table |
| Working, conversational and observational memory | FULL | Blueprint 06, Memory Research | write/read/compaction/expiry roles and limitations |
| Episodic and project memory | FULL | Blueprint 06 | schemas, examples and project-universe behavior |
| Global personal memory | FULL | Blueprint 06, 11 | risk-based promotion, sensitivity and correction |
| Relational/world-model memory | FULL | Blueprint 03, 06 | temporal relationships and provenance |
| Procedural memory and Golden Paths | FULL | Blueprint 06, ACRM | candidate-to-validated-procedure lifecycle |
| Failure and learning memory | FULL | Blueprint 06, 13 | incident, causal hypothesis, correction and regression links |
| Operational state is structured, not model memory | FULL | Blueprint 06, 12 | state-owner separation and recovery |
| Memory metadata | FULL | Blueprint 06 | complete structured example with epistemic and governance metadata |
| Fact/decision/observation/inference/hypothesis distinction | FULL | Blueprint 03, 06 | epistemic classes and conflict rules |
| Promotion by risk, scope and authority | FULL | Blueprint 06 | automatic/conditioned promotion matrix and scenarios |
| Supersession rather than silent overwrite | FULL | Blueprint 06 | lifecycle, conflict example and evaluation journey |
| Forgetting, expiry, deletion and archive | FULL | Blueprint 06, 11 | retention policy and proof requirements |
| Plausible but false memory | FULL | Blueprint 06, 13 | incident class and evaluation cases |
| Cross-project contamination | FULL | Blueprint 06, 11 | isolation invariant and adversarial test intent |
| Context Builder pipeline | FULL | Blueprint 06, 12 | retrieval, authority, temporal, sensitivity and live-state checks |
| Memory mechanism selection | OPEN_RESEARCH | Memory Research | observational memory, RAG, graphs, event logs and hybrid candidates |
| Memory evaluation at scale | OPEN_RESEARCH | Memory Research, Blueprint 13 | benchmark limitations and proposed Aurora-specific eval program |

---

## 7. Autonomy, campaigns and physical safety

| Discovery topic | Current coverage | Canonical owner | Evidence of preservation |
|---|---|---|---|
| Progressive authority | FULL | Blueprint 10 | observe-to-effect ladder and grant model |
| Delegated autonomy by mission | FULL | Blueprint 04, 10 | delegation/campaign lifecycle and authority envelope |
| Action, workflow, campaign and continuous-program levels | FULL | Blueprint 10 | N1–N4 definitions and graduation conditions |
| Overnight AI workflow improvement | FULL | Blueprint 04, 10, 13 | complete scenario, immutable eval criteria and promotion boundary |
| Firmware variant campaign | FULL | Blueprint 04, 09, 10 | compilation, flash, telemetry, safety and evidence journey |
| Autonomy Envelope | FULL | Blueprint 10 | structured schema and invariants |
| Budget across money, tokens, time, energy, cycles and wear | FULL | Blueprint 10 | budget dimensions, thresholds and enforcement |
| Stop and escalation conditions | FULL | Blueprint 10 | taxonomy, examples and ambiguous-boundary behavior |
| Deterministic physical guardrails | FULL | Blueprint 09, 10 | interlocks, fail-safe state and fault drills |
| Emergency authority | FULL | Blueprint 10 | narrow preauthorization and receipts |
| Revocation has operational effect | FULL | Blueprint 10, 11 | token/channel/credential containment and reconciliation |
| Physical actuation implementation | DEFERRED_BY_ROADMAP | Blueprint 09, 10, 14 | M10 only after observation, simulation and safety readiness |

---

## 8. Self-improvement and learning

| Discovery topic | Current coverage | Canonical owner | Evidence of preservation |
|---|---|---|---|
| Experimental self-improvement with supervised promotion | FULL | Blueprint 10, 13 | candidate, eval, review, canary, rollback and promotion gate |
| Continuous detection with authorized experiments | FULL | Blueprint 13 | opportunity lifecycle and prioritization |
| Symptom versus root cause | FULL | Blueprint 13 | causal-analysis method and anti-patch rule |
| Multiple incidents sharing one cause | FULL | Blueprint 13, History | causal graph scenario preserved |
| Structured incident record | FULL | Blueprint 13 | incident schema and required evidence |
| Correlation before patching | FULL | Blueprint 13 | thresholds and evidence requirements |
| Competing hypotheses and reproduction | FULL | Blueprint 13 | falsification and unreproducible-incident handling |
| Candidate versions, sandbox and rollback | FULL | Blueprint 13 | lifecycle and artifacts |
| Original, neighboring, contrary, historical, unseen and adversarial cases | FULL | Blueprint 13 | evaluation suite anatomy |
| Holdout governance | FULL | Blueprint 13 | protected evaluation data and anti-gaming rules |
| Independent reviewer | FULL | Blueprint 10, 13 | separation of proposer and acceptance authority |
| Canary and shadow modes | FULL | Blueprint 13 | rollout states and rollback triggers |
| Constitutional areas protected | FULL | Blueprint 10, 13 | immutable/autonomously non-promotable set |
| Learning successes into candidate Golden Paths | FULL | Blueprint 06, 13, ACRM | observation → pattern → validation → preferred path |
| Self-improvement evaluation mechanisms | OPEN_RESEARCH | Evaluation Research | benchmark, causal and statistical design requires spikes |

---

## 9. Capability, harness, contracts and AHDK

| Discovery topic | Current coverage | Canonical owner | Evidence of preservation |
|---|---|---|---|
| Tool/Agent/Workflow/Runtime/Harness distinctions | FULL | Blueprint 03, 05 | definitions, examples and ownership |
| Harness as complete specialized system | FULL | Blueprint 05, 07 | internal autonomy and external contract boundary |
| Aurora as global control plane | FULL | Blueprint 07, 12 | state, authority, composition and operator interaction |
| MNFS as future provider, not foundation | FULL | Blueprint 07, 14 | readiness and adapter criteria |
| Aurora owns why/what/limits; Harness owns local how | FULL | Blueprint 07 | responsibility matrix and examples |
| Hierarchical orchestration | FULL | Blueprint 07 | child delegation and multi-harness journeys |
| Authorized direct data plane | FULL | Blueprint 07, 12 | channel contract, telemetry example and no authority propagation |
| Capability Manifest | FULL | Blueprint 05 | complete structured example and validation lifecycle |
| Discovery/compatibility/trust/authority/execution separation | FULL | Blueprint 05 | separate state machines and failure behavior |
| Multidimensional, build-bound trust | FULL | Blueprint 05, 11 | provenance, version, environment and scope |
| Dynamic discovery with governed activation | FULL | Blueprint 05 | fail-closed activation and sandbox verification |
| Provider selection and fallback | FULL | Blueprint 05 | fit, sensitivity, cost, latency, recovery and explainability |
| Provider suspension/revocation/retirement | FULL | Blueprint 05, 11 | lifecycle and incident response |
| Aurora-owned semantic contracts | FULL | Blueprint 03, 05, 07, ADR-0001 | entities, versioning and binding neutrality |
| AHDK mandatory for first-party by policy | FULL | Blueprint 05, ADR-0002 | waiver and organizational policy |
| Contract independent from SDK | FULL | Blueprint 05, ADR-0002 | direct implementation/conformance proof |
| AHDK as development kit | FULL | Blueprint 05, AHDK Research | module architecture and Golden Path |
| Generated types, manifest builder, lifecycle/context/effect/artifact APIs | FULL | Blueprint 05 | conceptual interfaces and responsibilities |
| Cancellation, checkpoints, heartbeat and budgets | FULL | Blueprint 05, 07 | API/lifecycle behavior |
| Simulator, mocks, fault injection and scaffolder | FULL | Blueprint 05, ACRM | conformance and creation path |
| MCP/A2A/native binding roles | FULL | Blueprint 07, Interoperability Research | mapping, limits and maturity |
| Framework roles | FULL | Framework Research | Mastra, LangGraph, Pi, OpenHands, OpenAI Agents SDK and Langflow |
| Exact protocol/schema/runtime choices | OPEN_DECISION | ADRs/spikes | no mechanism silently promoted |

---

## 10. Durability, security, observability and provenance

| Discovery topic | Current coverage | Canonical owner | Evidence of preservation |
|---|---|---|---|
| Durable execution separate from agent runtime | FULL | Blueprint 12, Durability Research | DurableExecutionPort and comparison criteria |
| Temporal, DBOS, Restate and Inngest | OPEN_RESEARCH | Durability Research | focused alternatives and spike plan |
| SDK is not a security boundary | FULL | Blueprint 10, 11, ADR-0002 | bypass examples and external enforcement chain |
| Policy decision separate from enforcement | FULL | Blueprint 10, 11, 12 | PDP → gateway → receipt flow |
| Effect Gateways and Credential Broker | FULL | Blueprint 10–12 | gateway catalog and reference-based secret flow |
| Subject, actor, executor and device identity | FULL | Blueprint 03, 10, 11 | delegated identity and audit chain |
| Cedar, OPA, SPIFFE and RFC 8693 | OPEN_RESEARCH | Authority Research | candidates remain unselected |
| OpenTelemetry baseline | FULL | Blueprint 13, Events Research | signal model and Aurora semantic conventions |
| CloudEvents, AsyncAPI, JSON Schema and Protobuf | OPEN_RESEARCH | Events Research | boundary hypotheses and compatibility tests |
| SLSA-style provenance | FULL | Blueprint 05, 11 | build-bound provider trust lifecycle |
| Distributed trace across Core/Harness/Tool/Gateway | FULL | Blueprint 13 | explicit Golden Proof |
| Product storage/topology selection | OPEN_DECISION | Blueprint 12, future spikes | architecture ports defined without stack lock-in |

---

## 11. Documentation, methodology and roadmap

| Discovery topic | Current coverage | Canonical owner | Evidence of preservation |
|---|---|---|---|
| Documentation as project memory and governance | FULL | Blueprint 15, Documentation Map | authority/lifecycle/storage/read paths |
| Blueprint before implementation | FULL | Blueprint 15, ACRM, STATUS | R0 gate and implementation prohibition |
| Research before material technical choice | FULL | Blueprint 15, Research Map | primary-source and freshness protocol |
| Two planning horizons | FULL | Blueprint 01, 14 | constitutional direction vs executable commitment |
| Roadmap by cumulative capability and proof | FULL | Blueprint 14, generated roadmap | full milestone anatomy and Golden Proofs |
| Fresh-session Golden Proof | FULL | Blueprint 14, 15 | exact read path and expected output |
| Status as handoff | FULL | Blueprint 15, STATUS | phase, authorization, evidence, blockers and next action |
| One concept, one canonical owner | FULL | Blueprint 15, Documentation Map | conflict and divergence protocol |
| Conversation as discovery, repository as canonical memory | FULL | Blueprint 15, History | promotion and historical preservation |
| Aggregate Blueprint publication | FULL | Product Index, generated aggregate | deterministic source hashes and CI freshness |
| Capability Realization Method | FULL | ACRM | R0–R8 gates and artifact model |
| Requirements traceability | FULL | Requirements Traceability | 294 accepted A0 constitutional requirements |
| Mechanical documentation validation | FULL | scripts and CI | structure, IDs, relations, links, manifests, requirements and freshness |

---

## 12. Quantitative baseline

At A0 acceptance:

```text
Blueprint sections:          15
Constitutional requirements: 294
Focused + synthesis reports: 9
Primary-source entries:      92
Generated projections:       Product Blueprint + Roadmap
```

Counts are evidence of coverage, not proof of quality. Quality is determined by coherence, authority, mechanisms, scenarios, failure behavior and review.

---

## 13. Remaining open material

The following remain deliberately open and do not constitute documentation gaps:

- Aurora Core language and deployment shape;
- first AHDK language;
- schema representation per boundary;
- local RPC binding;
- durable execution engine;
- policy engine and workload identity implementation;
- canonical operational storage;
- Artifact/Evidence Store;
- event transport/backend;
- memory mechanism mix;
- first reference Harness runtime;
- exact first Mission Contract for the selected M0 milestone.

Each requires focused readiness work under the Aurora Capability Realization Method.

---

## 14. Exit condition

Discovery coverage is considered complete for A0 when:

1. this matrix has no missing or unmapped topic;
2. generated projections are current;
3. metadata, relations, links and source manifests validate;
4. the adversarial review is refreshed against the repaired package;
5. a fresh session passes the documentation Golden Proof;
6. Leandro explicitly accepts A0.

A0 exit state:

```text
CONTENT REMEDIATION COMPLETE
MECHANICAL + ADVERSARIAL VALIDATION COMPLETE
INDEPENDENT FRESH-SESSION GOLDEN PROOF: PASS
OPERATOR ACCEPTANCE: GRANTED — 2026-08-06
A0 MERGE: COMPLETE
```

This coverage matrix does not own post-A0 milestone/gate coordination. Current readiness state is recorded only in `docs/tracking/STATUS.md`.
