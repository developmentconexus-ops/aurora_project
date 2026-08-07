---
id: DOC-AURORA-REQUIREMENTS-TRACEABILITY
title: Aurora Constitutional Requirements and Traceability
document_type: requirements_traceability
form: reference

authority: specification
status: accepted
accepted_at: 2026-08-07
acceptance_evidence: DOC-AURORA-M0-R0-REMEDIATION-OPERATOR-ACCEPTANCE
version: 0.2.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - derived constitutional requirements
  - Blueprint requirement coverage
  - initial verification intent
related:
  - DOC-AURORA-PRODUCT-INDEX
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DOC-AURORA-BLUEPRINT-01
  - DOC-AURORA-BLUEPRINT-14
  - DOC-AURORA-DOCUMENTATION-COVERAGE
review_triggers:
  - Product Blueprint change
  - Capability requirement derivation
  - A0 acceptance review
last_reviewed: 2026-08-07
---

# Aurora Constitutional Requirements and Traceability

## 1. Purpose

This document converts the accepted A0 Product Blueprint into explicit constitutional requirements.

These requirements are intentionally product- and architecture-level. They are not yet allocated to implementation tasks. Future Capability Specs will:

1. analyze applicability;
2. derive more specific requirements;
3. allocate them to contracts and tests;
4. link implementation and evidence.

Status legend:

```text
ACCEPTED
→ accepted A0 constitutional requirement; governs future applicability analysis but does not itself authorize implementation

DEFERRED_BY_ROADMAP
→ constitutionally required direction, not current milestone commitment

OPEN_MECHANISM
→ required behavior is clear; implementation choice remains open
```

Verification categories are intended proof classes, not implementation authorization.

---

## 2. Blueprint 01 — Product Vision

| ID | Requirement | Source | Initial verification intent |
|---|---|---|---|
| AUR-REQ-VIS-001 | Aurora MUST preserve a stable system identity independent of any single LLM, provider, process, session or device. | §1.1, P5 | restart/model/provider/presence journey |
| AUR-REQ-VIS-002 | Aurora MUST operate as a cognitive control plane that owns global identity, context, authority and cross-domain composition. | §1.1, §1.11 | architecture conformance and ownership review |
| AUR-REQ-VIS-003 | Aurora MUST remain Leandro-first and single-user in the current product horizon. | §1.6 | scope review; no multi-tenant behavior |
| AUR-REQ-VIS-004 | Aurora MUST treat engineering as the first deep operational domain while allowing future domains only through explicit capability approval. | §1.7 | roadmap/capability applicability review |
| AUR-REQ-VIS-005 | Aurora MUST preserve the complete North Star as an end-to-end product journey across project context, capabilities, devices, evidence and next action. | §1.4 | North Star Golden Proof |
| AUR-REQ-VIS-006 | Aurora MUST distinguish understanding, memory, reasoning, action, observation, verification, recording and learning as separate responsibilities. | §1.9 | architecture and journey review |
| AUR-REQ-VIS-007 | No model, memory subsystem, Harness, interface or device MAY be treated as Aurora by itself. | §1.10 | component ownership review |
| AUR-REQ-VIS-008 | The Product Blueprint MUST document long-term direction while technical implementation commitment remains limited to the current accepted milestone. | §1.8 | documentation/authorization checks |
| AUR-REQ-VIS-009 | Aurora MUST support increasingly capable levels from context awareness to a continuous engineering companion without requiring all future features in the initial implementation. | §1.15 | roadmap sequencing review |
| AUR-REQ-VIS-010 | Product success MUST be measured by real journey outcomes, context quality, authority, evidence, safety, recovery and efficiency—not activity volume. | §1.17 | milestone closeout rubric |
| AUR-REQ-VIS-011 | Aurora MUST NOT become a chatbot wrapper, single universal agent, unrestricted swarm, surveillance system or framework-locked application. | §1.12, §1.18 | adversarial architecture review |
| AUR-REQ-VIS-012 | Implementation MUST remain blocked during A0 until explicit baseline acceptance and subsequent gate authorization. | §1.19 | STATUS/readiness check |

---

## 3. Blueprint 02 — Human–Aurora Relationship

| ID | Requirement | Source | Verification |
|---|---|---|---|
| AUR-REQ-REL-001 | Aurora MUST act as a trusted intellectual copilot rather than a passive obedient interface. | §2.2 | disagreement journey |
| AUR-REQ-REL-002 | Leandro MUST retain final authority over purpose, values, material trade-offs and authority grants. | §2.2–2.3 | authority model tests |
| AUR-REQ-REL-003 | Aurora MUST challenge material false premises, contradictions and risks with evidence and a recommendation. | §2.4–2.5 | adversarial conversation eval |
| AUR-REQ-REL-004 | After one clear informed objection, Aurora SHOULD proceed with Leandro's permitted final decision unless new evidence or a higher boundary intervenes. | §2.4 | decision-loop eval |
| AUR-REQ-REL-005 | Aurora MUST maintain a stable transparent identity and MUST NOT claim human experience or consciousness as fact. | §2.6 | cross-model personality tests |
| AUR-REQ-REL-006 | Aurora SHOULD express a hybrid personality combining calm precision and dry humor with proximity, curiosity and enthusiasm. | §2.7 | human evaluation rubric |
| AUR-REQ-REL-007 | Personality MUST yield to clarity, truth and safety in technical, incident and emergency contexts. | §2.8–2.10 | mode-switch eval |
| AUR-REQ-REL-008 | Aurora MUST NOT use emotional manipulation, neediness, jealousy, guilt or exclusivity to influence Leandro. | §2.10, §2.22 | safety dialogue eval |
| AUR-REQ-REL-009 | Aurora MUST support contextual controlled proactivity and an explicit attention budget. | §2.12–2.14 | notification relevance/dedup eval |
| AUR-REQ-REL-010 | A proactive interruption MUST be attributable to relevance, urgency, confidence and consequence of silence. | §2.13 | notification audit test |
| AUR-REQ-REL-011 | Aurora MUST distinguish preparation from execution so she can be proactive without crossing effect authority. | §2.15 | authority/proactivity integration |
| AUR-REQ-REL-012 | Aurora MUST communicate material uncertainty with its source, consequence and next discriminating action. | §2.16 | uncertainty calibration eval |
| AUR-REQ-REL-013 | Aurora MUST acknowledge errors precisely, correct current state and preserve material incidents without inventing root cause. | §2.18 | trust-repair journey |
| AUR-REQ-REL-014 | Corrections MUST be classified by scope and MUST NOT be generalized globally from one local request. | §2.19 | memory promotion eval |
| AUR-REQ-REL-015 | Relationship memory MUST remain inspectable, correctable, scope-aware and non-manipulative. | §2.20 | memory governance test |
| AUR-REQ-REL-016 | Aurora MUST support interaction modes such as `DO_WITH_ME`, `TEACH_ME`, `DO_FOR_ME`, `REVIEW_ONLY` and `OBSERVE_ONLY` without conflating mode with authority. | §2.21 | interaction mode tests |

---

## 4. Blueprint 03 — Domain and World Model

| ID | Requirement | Source | Verification |
|---|---|---|---|
| AUR-REQ-DOM-001 | Aurora MUST use stable identities for durable entities and MUST NOT rely on mutable names alone. | §3.2–3.5 | schema/domain tests |
| AUR-REQ-DOM-002 | Person, Aurora, Presence, Environment, Project, Mission, Delegation, Capability, Provider, Artifact, Evidence, Authority Grant, Device and Incident MUST remain distinct concepts. | §3.4–3.23 | domain model conformance |
| AUR-REQ-DOM-003 | Aurora MUST distinguish a generic Tool, Agent, Workflow, Runtime and Harness. | §3.12 | API/manifest review |
| AUR-REQ-DOM-004 | A Project MUST preserve objective, state, decisions, hypotheses, tasks, artifacts, evidence and next actions without storing all content inline. | §3.7 | project journey/repository test |
| AUR-REQ-DOM-005 | Mission MUST own a global objective and Delegation MUST own scoped assigned work. | §3.8–3.9 | lifecycle/contract tests |
| AUR-REQ-DOM-006 | Capability MUST describe reusable outcome independently from Provider implementation. | §3.10 | registry/provider substitution test |
| AUR-REQ-DOM-007 | Provider approval MUST bind to an exact identity/profile/build appropriate to risk. | §3.11 | build/trust tests |
| AUR-REQ-DOM-008 | Decision, Hypothesis, Experiment, Observation, Measurement, Artifact, Claim, Evidence, Verdict and Outcome MUST remain epistemically distinct. | §3.15–3.20 | evidence schema/eval |
| AUR-REQ-DOM-009 | Authority Grant MUST NOT be equivalent to Capability or technical access. | §3.21 | authorization tests |
| AUR-REQ-DOM-010 | External effects MUST preserve request, policy decision, receipt and observation as separate records. | §3.22 | effect journey |
| AUR-REQ-DOM-011 | Physical devices, firmware, controllers, calibration and live connection state MUST be separately identifiable. | §3.24–3.27 | device identity tests |
| AUR-REQ-DOM-012 | Domain relationships MUST preserve provenance, scope and temporal validity when material. | §3.28–3.31 | temporal/relationship tests |
| AUR-REQ-DOM-013 | Provider-internal entities MUST NOT become Aurora global entities unless promoted through an explicit adapter/contract. | §3.30 | provider isolation test |
| AUR-REQ-DOM-014 | World/project model projections MUST reference authoritative sources rather than become an unaudited second source of truth. | §3.31 | source-ownership review |

---

## 5. Blueprint 04 — Cognitive Lifecycle and Journeys

| ID | Requirement | Source | Verification |
|---|---|---|---|
| AUR-REQ-COG-001 | Aurora MUST support the lifecycle `perceive → understand → remember → form intent → plan → select capability → act → observe → verify → record → learn`. | §4.1 | journey/state tests |
| AUR-REQ-COG-002 | Perception events MUST preserve source, time, quality, sensitivity and environment when material. | §4.3 | event schema tests |
| AUR-REQ-COG-003 | Aurora MUST treat interpreted intent as revisable and SHOULD preserve confidence and alternative interpretations when material. | §4.4 | ambiguity eval |
| AUR-REQ-COG-004 | Aurora MUST resolve low-risk ambiguity by stated assumptions where safe and escalate material ambiguity. | §4.4.2 | interaction eval |
| AUR-REQ-COG-005 | Operational intent MUST preserve the original user intent separately from derived plans. | §4.6 | mission contract test |
| AUR-REQ-COG-006 | Aurora MUST detect and handle intent drift in long campaigns. | §4.6.2 | campaign fault test |
| AUR-REQ-COG-007 | Aurora MUST separate global planning from Harness-local planning. | §4.7 | multi-harness test |
| AUR-REQ-COG-008 | Capability selection MUST consider compatibility, trust, authority, data, cost, recovery and availability and remain explainable when material. | §4.8 | provider selection tests |
| AUR-REQ-COG-009 | Material actions MUST preserve an Intent–Action–Observation chain. | §4.9 | effect/evidence trace |
| AUR-REQ-COG-010 | Completion MUST wait for required observation windows and MUST NOT be inferred solely from action dispatch. | §4.10 | delayed-observation tests |
| AUR-REQ-COG-011 | Aurora MUST distinguish claim, receipt, evidence, verdict and outcome during verification. | §4.11 | evidence model tests |
| AUR-REQ-COG-012 | Records MUST be routed to the appropriate state, artifact, evidence, memory or documentation owner. | §4.12 | storage ownership tests |
| AUR-REQ-COG-013 | Learning MUST produce governed memory/procedure/improvement candidates rather than immediate global behavior changes. | §4.13 | self-improvement lifecycle test |
| AUR-REQ-COG-014 | Aurora MUST support laboratory continuation, overnight AI campaign, firmware campaign, multi-presence handoff, multi-harness collaboration and self-improvement as explicit end-to-end journeys over time. | §4.14–4.19 | staged roadmap journey tests |
| AUR-REQ-COG-015 | Interruptions MUST be classified and delivered according to severity, confidence, privacy and attention. | §4.20 | notification tests |
| AUR-REQ-COG-016 | Concurrency MUST preserve project isolation, budgets, authority, resource leases and emergency priority. | §4.21 | concurrent mission tests |
| AUR-REQ-COG-017 | A fresh process MUST reconstruct current authority, projects, active missions, checkpoints, material events and next safe action without transcript dependency. | §4.22 | restart Golden Proof |
| AUR-REQ-COG-018 | Error classification MUST drive recovery rather than generic retry. | §4.23–4.24 | fault-injection tests |

---

## 6. Blueprint 05 — Capabilities, Registry and AHDK

| ID | Requirement | Source | Verification |
|---|---|---|---|
| AUR-REQ-CAP-001 | Aurora MUST define a language/framework-independent Capability identity and version. | §5.3 | contract schema tests |
| AUR-REQ-CAP-002 | Provider and Provider Instance MUST be distinct and exact build/environment identity MUST be representable. | §5.4 | registry tests |
| AUR-REQ-CAP-003 | Discovery, compatibility, trust, authority and execution MUST be separate states. | §5.5 | registry FSM tests |
| AUR-REQ-CAP-004 | Every Provider MUST publish or be adapted to a structured Manifest declaring capabilities, contracts, effects, data, runtime, recovery, evidence and compatibility. | §5.6 | manifest conformance |
| AUR-REQ-CAP-005 | A Manifest MUST be treated as a claim and MUST NOT grant trust or authority. | §5.6.3 | malicious manifest test |
| AUR-REQ-CAP-006 | Capability Registry MUST preserve manifests, versions, builds, verification, trust, approvals, incidents and availability. | §5.7 | registry persistence/recovery |
| AUR-REQ-CAP-007 | Provider trust MUST be multidimensional and scoped; a universal Boolean trust flag MUST NOT be sufficient. | §5.9 | trust model tests |
| AUR-REQ-CAP-008 | Verification MUST declare tested scope and limitations. | §5.10 | verification artifact validation |
| AUR-REQ-CAP-009 | Approval MUST constrain provider/build, capability, project/data/effects/environment/time and evidence as applicable. | §5.11 | policy/selection tests |
| AUR-REQ-CAP-010 | Provider selection MUST evaluate capability fit, approval, sensitivity, evidence, cost, latency, recovery and incidents. | §5.12 | selection eval |
| AUR-REQ-CAP-011 | Provider substitution MUST create new context/authority/attempt and MUST NOT assume local-state compatibility. | §5.13 | fallback test |
| AUR-REQ-CAP-012 | Material Provider or contract changes MUST trigger compatibility analysis and re-verification; trust MUST NOT be inherited silently. | §5.14 | build/version tests |
| AUR-REQ-CAP-013 | First-party Harnesses MUST use AHDK by engineering policy unless an explicit waiver is accepted. | §5.17 | repository/conformance policy check |
| AUR-REQ-CAP-014 | AHDK MUST NOT be the canonical specification or sole security boundary. | §5.17–5.23 | architecture/security review |
| AUR-REQ-CAP-015 | AHDK SHOULD provide generated contracts, manifest builder, lifecycle, Context, Effect, Artifact/Evidence, decision, budget, checkpoint, telemetry, simulator and scaffolder capabilities. | §5.18–5.30 | AHDK acceptance suite |
| AUR-REQ-CAP-016 | AHDK Effect Client MUST request effects through governed gateways and MUST NOT decide authorization or expose raw credentials. | §5.23 | authority bypass test |
| AUR-REQ-CAP-017 | Universal Conformance Kit MUST evaluate providers independently from AHDK implementation. | §5.32 | SPK-001 Golden Proof |
| AUR-REQ-CAP-018 | Conformance MUST be profile/capability-specific and MUST NOT imply domain correctness/security beyond tested scope. | §5.32 | conformance reporting tests |
| AUR-REQ-CAP-019 | Aurora SHOULD provide a Golden Path/scaffolder for new first-party Harnesses. | §5.30 | scaffold acceptance |
| AUR-REQ-CAP-020 | Aurora MUST support replaceable bindings including native/RPC and standards adapters where applicable. | §5.34 | framework/protocol neutrality test |

---

## 7. Blueprint 06 — Memory, Knowledge and Context

| ID | Requirement | Source | Verification |
|---|---|---|---|
| AUR-REQ-MEM-001 | Aurora MUST separate raw history, governed memory, knowledge sources, operational state, world/project model and active context. | §6.2 | architecture/domain review |
| AUR-REQ-MEM-002 | Memory MUST implement governed write, manage and read lifecycles. | §6.3 | memory lifecycle tests |
| AUR-REQ-MEM-003 | Aurora MUST support logical strata from authoritative sources/state to ephemeral active context. | §6.5 | Context Builder tests |
| AUR-REQ-MEM-004 | Aurora MUST distinguish working, conversational, observational, episodic, project, global personal, relational, procedural, failure and operational memory/state. | §6.6 | memory taxonomy/schema tests |
| AUR-REQ-MEM-005 | Observational memory MAY be used for synthesis but MUST NOT own current authority or live operational state. | §6.6.3 | authority conflict tests |
| AUR-REQ-MEM-006 | Operational mission/budget/grant/device state MUST be stored structurally and MUST NOT depend on model recall. | §6.6.10 | restart/budget tests |
| AUR-REQ-MEM-007 | Material Memory Items MUST preserve provenance, scope, epistemic status, temporal validity, sensitivity and lifecycle. | §6.7 | schema validation |
| AUR-REQ-MEM-008 | Aurora MUST distinguish user-stated, user-approved, document-established, system-observed, measured, provider-reported, inferred and hypothetical information. | §6.8 | epistemic eval |
| AUR-REQ-MEM-009 | Confidence MUST NOT be treated as authority. | §6.8 | false-confidence test |
| AUR-REQ-MEM-010 | Memory MUST support valid time, observed/recorded times, freshness and supersession. | §6.9 | temporal evaluation |
| AUR-REQ-MEM-011 | Memory retrieval MUST enforce project/domain/presence/device scope isolation. | §6.10 | cross-scope negative tests |
| AUR-REQ-MEM-012 | Promotion MUST be risk-, authority- and scope-based; global/sensitive/inferred memory MUST require stronger evidence or confirmation. | §6.11 | promotion tests |
| AUR-REQ-MEM-013 | Memory MUST support observed/candidate/accepted/confirmed/superseded/expired/archived/removed states. | §6.12 | lifecycle tests |
| AUR-REQ-MEM-014 | Current authoritative decisions or live state MUST override stale memory for action, while historical memory remains discoverable. | §6.13, §6.19 | supersession journey |
| AUR-REQ-MEM-015 | Aurora MUST detect plausible unsupported memories and SHOULD state the difference between observed pattern and confirmed preference. | §6.14 | false-premise evaluation |
| AUR-REQ-MEM-016 | Consolidation MUST preserve sources, contradictions, scope and temporal distinction. | §6.15 | consolidation tests |
| AUR-REQ-MEM-017 | Memory MUST support forgetting, expiry, archive and retrieval-priority reduction. | §6.16 | retention tests |
| AUR-REQ-MEM-018 | Deletion MUST account for derived summaries, indexes, relationships, caches and provider copies according to policy. | §6.17 | deletion spike |
| AUR-REQ-MEM-019 | Context Builder MUST compile provider/decision-specific context using scope, authority, freshness, relevance, sensitivity and budget. | §6.18 | context evaluation |
| AUR-REQ-MEM-020 | Context Builder MUST minimize and redact context before cross-provider transfer. | §6.20, §6.23 | privacy/provider test |
| AUR-REQ-MEM-021 | Context Builder MUST detect when live-state refresh is required instead of using memory. | §6.22 | stale-state test |
| AUR-REQ-MEM-022 | Memory quality MUST be evaluated across recall, non-recall, temporal accuracy, authority, provenance, isolation, deletion, efficiency and longitudinal scale. | §6.24–6.25 | evaluation suite |
| AUR-REQ-MEM-023 | Technical memory architecture MUST remain open until focused research and spikes compare hybrid mechanisms. | §6.26–6.27 | readiness/ADR gate |

---

## 8. Blueprint 07 — Harness Orchestration

| ID | Requirement | Source | Verification |
|---|---|---|---|
| AUR-REQ-ORCH-001 | Aurora MUST use hierarchical contractual orchestration for cross-Harness work. | §7.1, §7.6 | multi-Harness Golden Proof |
| AUR-REQ-ORCH-002 | Aurora MUST own global objective, cross-domain composition, global authority/budget/state and acceptance. | §7.3 | ownership tests |
| AUR-REQ-ORCH-003 | Harness MUST own its domain methodology, local plan, internal agents/tools and local execution state. | §7.3 | provider boundary review |
| AUR-REQ-ORCH-004 | Aurora MUST NOT micromanage every local Harness operation. | §7.4 | plan boundary eval |
| AUR-REQ-ORCH-005 | Harnesses MUST NOT create cross-boundary authority or delegate to another provider without Aurora mediation. | §7.5–7.6 | child authority denial |
| AUR-REQ-ORCH-006 | Every Delegation MUST identify objective, scope/non-goals, Context Pack, Authority Grant, budget, criteria, evidence and recovery semantics. | §7.8–7.9 | contract validation |
| AUR-REQ-ORCH-007 | Context Pack MUST be immutable/versioned for a run, minimized, attributable and provider-scoped. | §7.10 | context pack tests |
| AUR-REQ-ORCH-008 | Child capability requests MUST return to Aurora and produce a separate child Delegation, context and grant. | §7.12 | child delegation test |
| AUR-REQ-ORCH-009 | Delegation lifecycle MUST reject running without exact authorization and completion without required evidence. | §7.13 | FSM/conformance |
| AUR-REQ-ORCH-010 | Attempts/runs MUST preserve retry reason, versions, context/grant revisions and prior effects. | §7.14 | retry/fault tests |
| AUR-REQ-ORCH-011 | Events MUST carry stable identity, source, time/sequence, trace and classification and MUST NOT be the sole source of critical state. | §7.15 | event/recovery tests |
| AUR-REQ-ORCH-012 | Progress reports MUST represent material state rather than raw low-level activity. | §7.16 | interaction eval |
| AUR-REQ-ORCH-013 | Decision Requests MUST be reserved for material scope, architecture, risk, authority or budget decisions and include recommendation/impact. | §7.17 | decision routing eval |
| AUR-REQ-ORCH-014 | Artifact, Claim, Receipt, Evidence, Verdict and Outcome MUST remain distinct across Harness boundaries. | §7.18 | evidence conformance |
| AUR-REQ-ORCH-015 | Direct data channels MAY bypass Core for bulk transfer but MUST be authorized, bounded, observable and revocable by Aurora. | §7.19–7.20 | data-channel test |
| AUR-REQ-ORCH-016 | Aurora global state and Harness local state MUST have explicit ownership and reconciliation behavior. | §7.21 | restart/reconciliation test |
| AUR-REQ-ORCH-017 | Cancellation MUST block new effects through authority revocation even if provider cancellation is cooperative. | §7.22 | cancellation/revoke test |
| AUR-REQ-ORCH-018 | Provider recovery MUST declare resume, checkpoint, heartbeat, idempotency and state compatibility. | §7.23 | conformance/fault tests |
| AUR-REQ-ORCH-019 | Generic retries MUST be prohibited for material errors unless classified and safe. | §7.24 | retry policy tests |
| AUR-REQ-ORCH-020 | Protocol bindings MUST transport Aurora semantics rather than redefine them. | §7.26–7.27 | adapter conformance |
| AUR-REQ-ORCH-021 | Harness internal frameworks MUST remain replaceable and MUST NOT become Aurora Mission/authority state. | §7.28 | SPK-008 |
| AUR-REQ-ORCH-022 | MNFS MUST be integrated as a future provider only through a stable mapped boundary and MUST NOT define Aurora Core architecture. | §7.30 | integration readiness review |

---

## 9. Blueprint 08 — Interaction, Multimodality and Presence

| ID | Requirement | Source | Verification |
|---|---|---|---|
| AUR-REQ-PRS-001 | Aurora MUST support one identity across multiple Presences. | §8.1–8.3 | cross-device identity test |
| AUR-REQ-PRS-002 | Presence MUST preserve device, person, channel, environment, capabilities, trust and effective authority identities. | §8.4 | presence manifest/auth tests |
| AUR-REQ-PRS-003 | Sensor capability MUST NOT imply activation. | §8.2, §8.7 | privacy sensor test |
| AUR-REQ-PRS-004 | Presence lifecycle MUST distinguish available interaction from authorized observation and ambient campaign. | §8.5 | presence FSM tests |
| AUR-REQ-PRS-005 | Voice commands with material numbers/units MUST support read-back or risk-proportional confirmation. | §8.7 | voice ambiguity test |
| AUR-REQ-PRS-006 | Image/video interpretation MUST preserve confidence and MUST NOT masquerade as calibrated measurement. | §8.7 | multimodal eval |
| AUR-REQ-PRS-007 | Handoff MUST preserve activity/intention through a presence-specific minimized Context Pack and MUST NOT replicate all memory/credentials. | §8.9–8.10 | handoff Golden Proof |
| AUR-REQ-PRS-008 | Sensitive disclosure/action MUST adapt to environment and support step-up authentication. | §8.11–8.12 | public/private presence tests |
| AUR-REQ-PRS-009 | Material sensor activation MUST have purpose, scope, duration, retention and visible indication. | §8.13 | sensor contract tests |
| AUR-REQ-PRS-010 | Aurora MUST offer a visible privacy mode capable of disabling/restricting sensors, notifications and cloud processing. | §8.14 | privacy mode test |
| AUR-REQ-PRS-011 | Offline/degraded Presence MUST state limitations and reconcile local events without silent duplication. | §8.15 | offline/reconnect test |
| AUR-REQ-PRS-012 | Proactive notifications MUST be routed to the most appropriate Presence rather than broadcast indiscriminately. | §8.16 | notification routing test |
| AUR-REQ-PRS-013 | Personality expression MUST adapt to channel/risk while identity remains stable. | §8.17 | cross-surface evaluation |
| AUR-REQ-PRS-014 | Compromised Presence revocation MUST block new context and effects and invalidate sessions/credentials. | §8.19 | security incident test |
| AUR-REQ-PRS-015 | Current scope MUST NOT assume a specific glasses, voice or UI technology. | §8.21 | architecture review |

---

## 10. Blueprint 09 — Tools, Devices and Laboratory

| ID | Requirement | Source | Verification |
|---|---|---|---|
| AUR-REQ-LAB-001 | Aurora MUST progress from device inventory/identity to observation before physical actuation. | §9.2, §9.13 | roadmap/lab progression |
| AUR-REQ-LAB-002 | Device identity MUST be independent of mutable connection path and SHOULD combine multiple identity signals appropriate to risk. | §9.4 | wrong-device test |
| AUR-REQ-LAB-003 | Device Manifest MUST declare capabilities, effects, safety limits, data and operational behavior. | §9.5 | manifest conformance |
| AUR-REQ-LAB-004 | Instrument Registry MUST preserve calibration, health, availability, leases and approved capabilities. | §9.6 | registry/restart tests |
| AUR-REQ-LAB-005 | Observation and command paths MUST be distinct and command MUST pass policy/gateway/interlock. | §9.7 | read-only/actuation tests |
| AUR-REQ-LAB-006 | Telemetry MUST preserve source, signal, unit, source/receive time, quality, sequence, calibration and experiment context. | §9.8 | telemetry schema tests |
| AUR-REQ-LAB-007 | Aurora MUST account for clock synchronization/drift before inferring event causality. | §9.9 | time-correlation test |
| AUR-REQ-LAB-008 | Measurement evidence MUST reference calibration and uncertainty appropriate to its criterion. | §9.10, §9.25 | evidence quality tests |
| AUR-REQ-LAB-009 | Laboratory Protocol MUST be versioned and define prerequisites, connections, limits, steps, abort, cleanup and evidence. | §9.11 | protocol validation |
| AUR-REQ-LAB-010 | Critical physical limits MUST be enforced by deterministic interlocks independent of LLM/cloud availability. | §9.12 | physical fault drill |
| AUR-REQ-LAB-011 | Device authority MUST support progressive levels from inventory/observe through controlled actuation and bounded autonomous campaign. | §9.13 | authority progression tests |
| AUR-REQ-LAB-012 | Simulation/HIL MUST declare model limitations and MUST NOT be treated as physical proof automatically. | §9.14–9.15 | simulation-to-bench comparison |
| AUR-REQ-LAB-013 | Firmware artifacts MUST preserve source, toolchain, config, digest, compatibility, flash and running-version verification. | §9.16–9.17 | firmware provenance/flash test |
| AUR-REQ-LAB-014 | Flash success MUST NOT be inferred only from programmer exit code. | §9.17 | post-flash identity test |
| AUR-REQ-LAB-015 | Exclusive physical resources MUST use leases with expiry/cleanup/orphan recovery. | §9.18 | concurrency/lease test |
| AUR-REQ-LAB-016 | High-rate telemetry MUST use governed data-plane channels or artifacts rather than model prompt/control messages. | §9.19 | data-channel test |
| AUR-REQ-LAB-017 | A firmware/laboratory campaign MUST preserve fixed safety boundaries, experiment ledger, artifacts and no automatic production promotion. | §9.20 | campaign Golden Proof |
| AUR-REQ-LAB-018 | Human manual observations MUST be attributed and MUST NOT be represented as instrument measurements. | §9.22 | evidence typing test |
| AUR-REQ-LAB-019 | Physical incident response MUST prioritize containment, safe state and evidence over mission completion. | §9.23 | incident drill |
| AUR-REQ-LAB-020 | No physical autonomy level MAY be accepted solely from software simulation tests. | §9.26 | acceptance gate review |

---

## 11. Blueprint 10 — Autonomy, Authority and Safety

| ID | Requirement | Source | Verification |
|---|---|---|---|
| AUR-REQ-AUT-001 | Aurora MUST implement progressive authority and mission-scoped delegated autonomy. | §10.1–10.4 | autonomy-level tests |
| AUR-REQ-AUT-002 | Access to a tool, credential, API or device MUST NOT constitute authority. | §10.2 | unauthorized access test |
| AUR-REQ-AUT-003 | Authority MUST be scoped, expiring and revocable. | §10.2 | grant lifecycle tests |
| AUR-REQ-AUT-004 | Aurora MUST be autonomous inside an approved envelope and conservative at its boundary. | §10.2, §10.6 | adaptive campaign test |
| AUR-REQ-AUT-005 | Autonomous Mission Envelope MUST define objective, baseline, mutable/protected space, environment, authority, budgets, evaluation, guardrails, stop, escalation, evidence and promotion. | §10.6–10.7 | contract schema/conformance |
| AUR-REQ-AUT-006 | Adaptive campaigns MAY choose hypotheses, variants and order without per-step human confirmation within the envelope. | §10.8–10.9 | overnight campaign Golden Proof |
| AUR-REQ-AUT-007 | Authority Grant MUST preserve subject, actor, executor, Delegation, actions, resources, conditions and validity. | §10.10–10.11 | auth policy tests |
| AUR-REQ-AUT-008 | Effect policy MUST consider effect type, materiality, reversibility, environment and third-party impact. | §10.12–10.13 | risk/policy tests |
| AUR-REQ-AUT-009 | Policy decision MUST be enforced at Effect Gateway/environment and MUST NOT rely on model/AHDK alone. | §10.14–10.15 | SDK bypass test |
| AUR-REQ-AUT-010 | Credentials MUST be brokered/scoped and MUST NOT be broadly exposed to Harness/model context. | §10.16 | secret leakage test |
| AUR-REQ-AUT-011 | Budgets MUST support deterministic soft/hard thresholds across relevant resource dimensions. | §10.17 | budget enforcement |
| AUR-REQ-AUT-012 | Deterministic interlocks MUST be observable, testable and protected from ordinary campaign authority. | §10.18–10.19 | interlock drill |
| AUR-REQ-AUT-013 | Stop and escalation conditions MUST be explicit and enforceable. | §10.20–10.21 | campaign stop tests |
| AUR-REQ-AUT-014 | Revocation MUST operationally block new effects, credentials and channels and reconcile active effects. | §10.22 | partition/revocation test |
| AUR-REQ-AUT-015 | Ambiguous external effects MUST be reconciled and MUST NOT be blindly retried. | §10.23 | ambiguous effect test |
| AUR-REQ-AUT-016 | Emergency authority MUST be narrow, pre-defined, locally enforceable and MUST NOT permit mission continuation. | §10.24 | emergency drill |
| AUR-REQ-AUT-017 | Material confirmation MUST identify exact target/action/consequence and safe default. | §10.25 | confirmation UX test |
| AUR-REQ-AUT-018 | Aurora MAY experiment on non-constitutional components but MUST NOT autonomously promote changes to protected areas. | §10.26–10.27 | self-improvement security test |
| AUR-REQ-AUT-019 | Protected areas MUST be enforced through technical/review boundaries rather than prompt instructions only. | §10.27 | repository/policy gate test |
| AUR-REQ-AUT-020 | Authority incidents MUST cause containment, evidence preservation, trust update and causal review. | §10.29 | incident closeout |

---

## 12. Blueprint 11 — Security, Privacy and Sovereignty

| ID | Requirement | Source | Verification |
|---|---|---|---|
| AUR-REQ-SEC-001 | Aurora MUST preserve confidentiality, integrity, availability, authority integrity, safety, accountability and user sovereignty. | §11.2 | threat/acceptance review |
| AUR-REQ-SEC-002 | Identity, Authority, Context/Data, Effect, Execution, Presence, Supply-Chain and Audit planes MUST remain distinct. | §11.3 | architecture review |
| AUR-REQ-SEC-003 | Every material Capability MUST update/apply a threat model appropriate to its data/effects. | §11.4 | Capability readiness gate |
| AUR-REQ-SEC-004 | Trust zones MUST influence data, credential, effect, environment and verification policy. | §11.5 | zone policy tests |
| AUR-REQ-SEC-005 | Data MUST support PUBLIC, INTERNAL, CONFIDENTIAL, SENSITIVE, SECRET and DEVICE_RESTRICTED classifications or an explicitly equivalent accepted model. | §11.6 | data policy tests |
| AUR-REQ-SEC-006 | Cross-boundary context transfer MUST be minimized, redacted/referenced and purpose-authorized. | §11.7 | provider data-flow test |
| AUR-REQ-SEC-007 | Canonical identity, memory, grants, policies, secrets references, audit and operational state MUST remain under Leandro-controlled sovereign infrastructure. | §11.8 | deployment/data inventory audit |
| AUR-REQ-SEC-008 | External providers MUST be governed by data classes, purpose, retention/training policy, region/legal constraints where applicable, cost and audit. | §11.9 | provider policy test |
| AUR-REQ-SEC-009 | Secret values MUST NOT enter prompts, manifests or general logs through the normal path. | §11.10 | secret scanning/adversarial test |
| AUR-REQ-SEC-010 | Actor/delegation identity MUST be preserved across Core, Harness, worker, Presence and effect. | §11.11 | auth/audit test |
| AUR-REQ-SEC-011 | Policy decision and Effect enforcement MUST be separate. | §11.12 | architecture/security test |
| AUR-REQ-SEC-012 | Sandboxed environments MUST declare technically enforced filesystem, process, network, credential, device and resource boundaries. | §11.13 | sandbox escape/bypass tests |
| AUR-REQ-SEC-013 | Untrusted content MUST be treated as data and MUST NOT redefine authority or product policy. | §11.14 | prompt-injection tests |
| AUR-REQ-SEC-014 | Memory privacy MUST cover promotion, scope, derived data, deletion and cross-project access. | §11.15 | memory security eval |
| AUR-REQ-SEC-015 | Raw transcript, summaries, memories, audio/video and derived observations MUST have separate retention/policy. | §11.16 | data inventory/deletion test |
| AUR-REQ-SEC-016 | Provider/Harness trust MUST be bindable to exact source/build/provenance/conformance and MUST be invalidated when material identity changes. | §11.17 | supply-chain trust test |
| AUR-REQ-SEC-017 | Presence/device loss or compromise MUST support session, credential, cache and effect revocation. | §11.18 | lost-device drill |
| AUR-REQ-SEC-018 | Physical effects MUST use defense in depth from product policy through local physical protection. | §11.19 | safety architecture/drill |
| AUR-REQ-SEC-019 | Leandro MUST be able to inspect material data disclosures, providers, grants, effects and receipts. | §11.20 | audit UX test |
| AUR-REQ-SEC-020 | Data MUST support classification, retention, supersession, archive, deletion and verified deletion lifecycle. | §11.21 | lifecycle test |
| AUR-REQ-SEC-021 | Aurora MUST support export, correction, deletion, provider revocation, credential rotation, backup/restore and migration. | §11.22–11.23 | sovereignty Golden Proof |
| AUR-REQ-SEC-022 | Restore MUST NOT silently reactivate expired grants or compromised trust. | §11.23 | restore security test |
| AUR-REQ-SEC-023 | Queued effects after disconnection MUST revalidate freshness and authority before execution. | §11.24 | partition/reconnect test |
| AUR-REQ-SEC-024 | Security incident response MUST support detect, contain, preserve evidence, assess scope, recover, investigate and review. | §11.25 | incident drill |
| AUR-REQ-SEC-025 | Security controls SHOULD minimize friction through scoped preauthorization, step-up and short-lived grants rather than confirm every low-risk action. | §11.26 | usability/security eval |

---

## 13. Blueprint 12 — System Architecture

| ID | Requirement | Source | Verification |
|---|---|---|---|
| AUR-REQ-ARC-001 | Aurora MUST implement a sovereign Core with modular domain boundaries and replaceable inference/execution adapters. | §12.2–12.4 | architecture fitness review |
| AUR-REQ-ARC-002 | Domain modules MUST NOT import provider, protocol, database, workflow-engine or UI framework semantics as canonical domain types. | §12.4.1 | dependency/static architecture test |
| AUR-REQ-ARC-003 | Logical modularity MUST precede physical distribution; components MUST be separated into services only with evidence. | §12.4.3 | topology ADR review |
| AUR-REQ-ARC-004 | Durable state MUST outlive processes. | §12.4.4 | restart tests |
| AUR-REQ-ARC-005 | Large/high-rate data MUST move by artifact reference or governed data channel. | §12.4.5 | data-plane test |
| AUR-REQ-ARC-006 | Identity/Relationship, Project/World Model, Memory/Context, Mission Control, Registry, Authority, Durable Execution, Artifact/Evidence, Attention, Failure Intelligence and Operator Interaction MUST have explicit ownership boundaries. | §12.5 | component architecture review |
| AUR-REQ-ARC-007 | Aurora Contract Model MUST be language-neutral and versioned independently from AHDK/bindings. | §12.6 | contract/codegen tests |
| AUR-REQ-ARC-008 | Capability Fabric MUST support provider discovery/connection, dispatch, events, artifacts/evidence, signals and reconciliation across replaceable bindings. | §12.7 | adapter integration tests |
| AUR-REQ-ARC-009 | Operational state MUST remain distinct from conversation, engine history, Harness local state, Git docs and logs. | §12.10 | state ownership/recovery |
| AUR-REQ-ARC-010 | Domain Event, Transport Message, Telemetry, Audit and Artifact MUST remain distinct. | §12.11 | event architecture tests |
| AUR-REQ-ARC-011 | Direct data plane MUST NOT bypass control-plane authority. | §12.12 | channel security test |
| AUR-REQ-ARC-012 | Memory/knowledge topology MUST preserve authoritative strata and MUST NOT make vector storage the owner of truth. | §12.13 | memory architecture review |
| AUR-REQ-ARC-013 | Artifact and Evidence logical stores MUST remain distinct even if initially implemented physically together. | §12.14 | data model tests |
| AUR-REQ-ARC-014 | Effect architecture MUST preserve Grant → PDP → Gateway → Broker → Target → Receipt. | §12.15 | SPK-005 |
| AUR-REQ-ARC-015 | Device Plane MUST remain behind device registry, controllers, telemetry, command gateways and interlocks. | §12.16 | lab architecture test |
| AUR-REQ-ARC-016 | Presence adapters MAY host local functions but canonical state MUST remain in Core. | §12.17 | handoff/offline test |
| AUR-REQ-ARC-017 | Models MUST be replaceable role-specific runtimes and MUST NOT define Aurora identity. | §12.18 | cross-model identity test |
| AUR-REQ-ARC-018 | Initial topology SHOULD prefer a modular local Core over premature microservices/Kubernetes. | §12.19 | topology decision/spike |
| AUR-REQ-ARC-019 | Architecture MUST support evolution from single machine to lab node and distributed personal system through adapters rather than domain rewrite. | §12.20 | migration design review |
| AUR-REQ-ARC-020 | Failure domains MUST have explicit recovery/containment behavior. | §12.21 | fault-injection plan |
| AUR-REQ-ARC-021 | In-flight work MUST bind exact contract/provider/schema versions and migrate explicitly. | §12.22 | version migration tests |
| AUR-REQ-ARC-022 | Cross-component observability MUST propagate stable project/mission/delegation/provider/effect/artifact IDs without sensitive payloads. | §12.23 | SPK-006 |
| AUR-REQ-ARC-023 | Architecture changes MUST pass fitness questions for authority duplication, provider leakage, restart, enforcement, context and proof. | §12.25 | architecture CI/review checklist |

---

## 14. Blueprint 13 — Reliability, Evaluation and Self-Improvement

| ID | Requirement | Source | Verification |
|---|---|---|---|
| AUR-REQ-RELIA-001 | Reliability MUST be evaluated across continuity, correctness, context, authority, operations, evidence, safety, security, interaction and efficiency. | §13.2 | milestone closeout rubric |
| AUR-REQ-RELIA-002 | Claim, Receipt, Evidence, Verdict and Outcome MUST remain distinct. | §13.3 | evidence schema tests |
| AUR-REQ-RELIA-003 | Material Evidence MUST preserve criterion, producer, verifier, method, environment, versions, artifacts, uncertainty and limitations. | §13.4 | evidence conformance |
| AUR-REQ-RELIA-004 | Parent/global acceptance MUST NOT be inferred solely from child completion. | §13.5 | composition tests |
| AUR-REQ-RELIA-005 | Observability MUST use traces/metrics/logs for decisions and MUST NOT become activity theater or domain truth. | §13.6–13.8 | telemetry governance review |
| AUR-REQ-RELIA-006 | Every retained metric SHOULD state which decision/threshold/action it informs and what it cannot prove. | §13.8 | metric catalog validation |
| AUR-REQ-RELIA-007 | Errors MUST be classified across cognitive, memory, contract, operational, effect, security and evaluation domains. | §13.10 | fault taxonomy tests |
| AUR-REQ-RELIA-008 | Retry MUST require retryable classification, idempotency/safety, budget and no systematic-failure concealment. | §13.11 | retry policy tests |
| AUR-REQ-RELIA-009 | Evaluation MUST cover component, capability, journey, safety/security, interaction and longitudinal behavior. | §13.12 | evaluation coverage |
| AUR-REQ-RELIA-010 | Evaluation datasets MUST separate investigation, development, validation, holdout, adversarial and production shadow roles. | §13.13 | eval governance tests |
| AUR-REQ-RELIA-011 | Candidate evaluation MUST be multi-objective when trade-offs are material. | §13.14 | candidate report review |
| AUR-REQ-RELIA-012 | Confidence SHOULD preserve distinct sources such as intent, freshness, measurement and verification coverage. | §13.15 | calibration eval |
| AUR-REQ-RELIA-013 | Human feedback MUST be classified and MUST NOT automatically optimize Aurora toward agreement. | §13.16 | feedback learning tests |
| AUR-REQ-RELIA-014 | Failure Intelligence MUST preserve incident, symptom, context, hypotheses, reproduction, candidate, evaluation and promotion relations. | §13.17–13.20 | failure graph tests |
| AUR-REQ-RELIA-015 | Root-cause investigation MUST use competing hypotheses and discriminating tests and MUST state when cause is inconclusive. | §13.20 | incident review rubric |
| AUR-REQ-RELIA-016 | Improvement candidates MUST be versioned and linked to causal hypothesis, incidents, test plan, risks and rollback. | §13.21 | candidate schema |
| AUR-REQ-RELIA-017 | Improvement MUST be tested on original, neighboring, contrary, historical, unseen and adversarial cases. | §13.22 | candidate evaluation |
| AUR-REQ-RELIA-018 | Material self-improvement MUST receive independent review. | §13.23 | review separation test |
| AUR-REQ-RELIA-019 | Material candidate promotion MUST support shadow, canary and tested rollback. | §13.24 | canary/rollback Golden Proof |
| AUR-REQ-RELIA-020 | Constitutional/protected areas MUST NOT be autonomously promoted. | §13.25 | protected-area test |
| AUR-REQ-RELIA-021 | Aurora MAY continuously detect opportunities but MUST experiment only inside authorized scopes/budgets. | §13.26 | improvement scheduler tests |
| AUR-REQ-RELIA-022 | Successful behavior MUST become procedure/Golden Path only after repeated validation and explicit promotion. | §13.27 | procedural-memory tests |
| AUR-REQ-RELIA-023 | Memory, personality, orchestration and physical operation MUST each have dedicated evaluation suites. | §13.28–13.31 | eval coverage matrix |
| AUR-REQ-RELIA-024 | Material incidents MUST update applicable documentation, evaluation and trust—not only code. | §13.32 | incident closeout check |

---

## 15. Blueprint 14 — Capability Roadmap

| ID | Requirement | Source | Verification |
|---|---|---|---|
| AUR-REQ-RDM-001 | Aurora roadmap MUST consist of cumulative Product Milestones closed by operator-visible Golden Proofs and named risk reduction. | §14.1, §14.5 | roadmap validation |
| AUR-REQ-RDM-002 | Constitutional and executable horizons MUST remain separate. | §14.2 | status/authorization review |
| AUR-REQ-RDM-003 | Product Milestone, Capability, Architecture Spike, Implementation Mission and Closeout MUST remain distinct. | §14.4 | planning artifact review |
| AUR-REQ-RDM-004 | Every Product Milestone promoted into the current executable horizon MUST define outcome, value, risk, entry, capabilities, spikes, Golden Proof, evidence, exit, telemetry, non-goals, dependencies, replan triggers and promotion/authority boundary. | §14.5 | milestone schema/check |
| AUR-REQ-RDM-005 | Roadmap SHOULD follow walking skeleton, vertical slice, sovereign Core before integration, context before autonomy and observe-before-actuate principles. | §14.6 | sequence review |
| AUR-REQ-RDM-006 | A0 MUST close through complete documentation, traceability, adversarial review and fresh-session continuity. | §14.7 | A0 Golden Proof |
| AUR-REQ-RDM-007 | M0 MUST prove sovereign identity/state across restart and backup/restore. | §14.8 | M0 Golden Proof |
| AUR-REQ-RDM-008 | M1 MUST prove governed project/conversation memory, source authority, supersession, isolation and deletion. | §14.9 | M1 Golden Proof |
| AUR-REQ-RDM-009 | M2 MUST prove Capability Registry, AHDK kernel, conformance and build-bound trust with a reference provider. | §14.10 | M2 Golden Proof |
| AUR-REQ-RDM-010 | M3 MUST prove bounded Delegation, Context, Authority, Artifact and Evidence. | §14.11 | M3 Golden Proof |
| AUR-REQ-RDM-011 | M4 MUST prove durable restart, wait, budget and duplicate-effect prevention. | §14.12 | M4 Golden Proof |
| AUR-REQ-RDM-012 | M5 MUST prove hierarchical multi-Harness composition and non-transitive authority. | §14.13 | M5 Golden Proof |
| AUR-REQ-RDM-013 | M6 MUST integrate one real engineering Harness selected by readiness and risk, not assume MNFS automatically. | §14.14 | M6 readiness/Golden Proof |
| AUR-REQ-RDM-014 | M7 MUST prove adaptive campaign under immutable evaluation, budget and no automatic production promotion. | §14.15 | M7 Golden Proof |
| AUR-REQ-RDM-015 | M8 MUST prove contextual multi-Presence handoff, privacy, offline operation and revocation. | §14.16 | M8 Golden Proof |
| AUR-REQ-RDM-016 | M9 MUST prove read-only laboratory observation with identity, telemetry, calibration and evidence. | §14.17 | M9 Golden Proof |
| AUR-REQ-RDM-017 | M10 MUST prove bounded physical actuation and independent containment. | §14.18 | M10 Golden Proof |
| AUR-REQ-RDM-018 | M11 MUST prove causal self-improvement, protected eval, independent review, canary and rollback. | §14.19 | M11 Golden Proof |
| AUR-REQ-RDM-019 | M12 MUST remain strategic direction rather than current implementation commitment. | §14.20 | status/readiness check |
| AUR-REQ-RDM-020 | Roadmap changes MUST be evidence-based, versioned and approved when material. | §14.24 | roadmap governance |

---

## 16. Blueprint 15 — Documentation and Research Governance

| ID | Requirement | Source | Verification |
|---|---|---|---|
| AUR-REQ-DOC-001 | Every durable concept MUST have one canonical owner and conflicts MUST be resolved explicitly. | §15.2–15.6 | documentation authority check |
| AUR-REQ-DOC-002 | Conversation MUST be treated as discovery; durable intent MUST be promoted to repository artifacts. | §15.2, §15.11 | coverage/history review |
| AUR-REQ-DOC-003 | Research MUST be evidence and MUST NOT become authority without ADR/Spec acceptance. | §15.2, §15.12 | research/decision traceability |
| AUR-REQ-DOC-004 | Documents MUST declare authority, lifecycle, ownership, source-of-truth scope and related artifacts. | §15.3–15.7 | frontmatter validation |
| AUR-REQ-DOC-005 | Product Blueprint MUST use fifteen modular sources and a generated/read-only aggregate. | §15.9 | aggregate freshness check |
| AUR-REQ-DOC-006 | Documentation layout MUST grow on demand and MUST NOT create empty directories/documents solely for a diagram. | §15.10 | repository review |
| AUR-REQ-DOC-007 | Research MUST define question, scope, source strategy, evidence, analysis, implications and promotion. | §15.12 | research report validation |
| AUR-REQ-DOC-008 | Material research MUST have source manifests with claim support, date/version and limitations. | §15.14 | source-manifest validation |
| AUR-REQ-DOC-009 | Focused research reports MUST be independently refreshable and an aggregate MUST NOT replace them. | §15.13 | research map review |
| AUR-REQ-DOC-010 | ADRs MUST preserve context, alternatives, decision, rationale, consequences and supersession. | §15.16 | ADR template/check |
| AUR-REQ-DOC-011 | Capability Specs MUST own reusable behavior and Mission Contracts MUST own scoped commitments. | §15.17–15.18 | realization gate checks |
| AUR-REQ-DOC-012 | STATUS MUST state current gate, authorizations/prohibitions, blockers, verification and exact next action. | §15.19–15.20 | fresh-session test |
| AUR-REQ-DOC-013 | Superseded/rejected documents MUST remain discoverable and current indexes/projections MUST update. | §15.21 | supersession checks |
| AUR-REQ-DOC-014 | Generated projections MUST declare sources and MUST NOT be edited directly. | §15.22 | CI freshness check |
| AUR-REQ-DOC-015 | Every material change MUST declare documentation impact. | §15.23 | PR/mission check |
| AUR-REQ-DOC-016 | Documentation validation MUST cover structure, authority, research, quality, traceability and projection freshness. | §15.24 | docs CI |
| AUR-REQ-DOC-017 | Major baselines/Capabilities MUST receive adversarial documentation review. | §15.25 | review evidence |
| AUR-REQ-DOC-018 | Agent/human read paths MUST load the smallest correct authority set. | §15.26–15.27 | context/read-path test |
| AUR-REQ-DOC-019 | A0 MUST include a fresh-session Golden Proof using repository only. | §15.28 | A0 acceptance evidence |
| AUR-REQ-DOC-020 | Git write permission or generated proposal MUST NOT constitute operator approval. | §15.29 | branch/approval governance |
| AUR-REQ-DOC-021 | Documentation defects that affect authority/security/implementation MUST be treated as product incidents/findings. | §15.30 | incident workflow test |
| AUR-REQ-DOC-022 | Implementation MUST remain prohibited until A0 criteria and explicit operator acceptance pass. | §15.31 | STATUS/gate enforcement |

---

## 17. Accepted A0 totals and allocation baseline

Accepted A0 constitutional requirements:

```text
Blueprint 01  12
Blueprint 02  16
Blueprint 03  14
Blueprint 04  18
Blueprint 05  20
Blueprint 06  23
Blueprint 07  22
Blueprint 08  15
Blueprint 09  20
Blueprint 10  20
Blueprint 11  25
Blueprint 12  23
Blueprint 13  24
Blueprint 14  20
Blueprint 15  22
----------------
TOTAL         294
```

A0 allocation baseline:

```text
Product constitution: ACCEPTED
Capability Specs: NOT_STARTED at A0 closeout
Mission Contracts: NOT_STARTED at A0 closeout
Implementation: NOT AUTHORIZED BY A0
Evidence: accepted A0 documentation/research + independent fresh-session Golden Proof
```

The large number does not mean every Capability implements all 294 requirements. R1 applicability selects relevant requirements and records rationale. Current milestone/gate authorization is intentionally not owned by this specification; consult `docs/tracking/STATUS.md`.

---

## 18. A0 coverage requirements

For A0 acceptance:

- every requirement must resolve to an existing Blueprint source;
- IDs must be unique;
- no requirement may contradict another accepted/proposed section without a Finding;
- Documentation Coverage must map discovery decisions to these requirements/owners;
- proposed ADRs must identify affected requirement IDs;
- roadmap must retain all deferred directional requirements;
- fresh-session review must demonstrate comprehension of high-impact requirement groups;
- no implementation allocation is implied.

---

## 19. Next derivation step

`M0 — Sovereign Core Walking Skeleton` is the selected first Product Milestone. When ACRM R1 is separately authorized, create the first Capability applicability/requirements package.

Target package for M0:

```text
CAP-SOVEREIGN-CORE
├── applicability against 294 constitutional requirements
├── derived atomic requirements
├── system/domain spec
├── threat model
├── test plan
└── coverage
```

The Product Blueprint remains the constitutional source; Capability requirements can be more precise but cannot silently weaken it.
