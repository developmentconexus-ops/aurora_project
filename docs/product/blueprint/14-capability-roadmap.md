---
id: DOC-AURORA-BLUEPRINT-14
title: Roadmap de Capacidades e Ordem de Realização
document_type: product_blueprint_section
form: explanation

authority: constitutional
status: accepted
accepted_at: 2026-08-06
acceptance_evidence: DOC-AURORA-A0-OPERATOR-ACCEPTANCE
version: 0.2.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - product capability sequence
  - milestone anatomy
  - long-horizon versus commitment boundary
related:
  - DOC-AURORA-BLUEPRINT-01
  - DOC-AURORA-BLUEPRINT-04
  - DOC-AURORA-BLUEPRINT-05
  - DOC-AURORA-BLUEPRINT-12
  - DOC-AURORA-BLUEPRINT-13
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DOC-AURORA-ROADMAP
review_triggers:
  - milestone sequence changes
  - Golden Proof or product risk changes
  - capability readiness evidence changes
last_reviewed: 2026-08-06
---

# 14. Roadmap de Capacidades e Ordem de Realização

## 14.1 Propósito

Aurora is a long-horizon project that can easily become either:

- an endless architecture exercise;
- a collection of impressive but disconnected demos;
- a chat UI built too early;
- a laboratory control system without a reliable cognitive Core;
- a framework platform that never becomes the personal intelligence Leandro wants.

The roadmap prevents these outcomes by defining a cumulative sequence of capabilities. Each Product Milestone must:

- deliver operator-visible value;
- traverse the necessary layers end to end;
- retire a named product risk;
- close with evidence;
- preserve future direction without implementing distant complexity prematurely.

> The roadmap is a sequence of cumulative capabilities, each closed by a Golden Proof that demonstrates a real behavior and reduces a material risk.

---

## 14.2 Two horizons

### Constitutional horizon

Documents the complete intended direction:

- personal intelligence;
- memory;
- capability fabric;
- multiple presences;
- laboratory observation and actuation;
- delegated campaigns;
- self-improvement;
- continuous engineering companion.

This horizon protects the destination from local architecture traps.

### Executable horizon

Details only:

- current Product Milestone;
- prerequisites;
- relevant capabilities;
- research;
- ADRs;
- architecture spikes;
- contracts;
- test/evidence plan.

Distant milestones remain directional until promoted by readiness.

> Complete long-term vision; progressively detailed technical commitment.

---

## 14.3 What the roadmap is not

- calendar commitment;
- promise of release dates;
- feature backlog;
- list of frameworks;
- fixed waterfall;
- permission to implement future milestones;
- justification for creating empty abstractions;
- substitute for Capability Specs and Mission Contracts.

---

## 14.4 Roadmap units

### Product Milestone

Delivers reusable Aurora capability and operator-visible behavior.

### Capability

Reusable product behavior specified independently of one milestone.

### Architecture Spike

Investigates a material uncertainty through executable evidence.

### Implementation Mission

Scoped work contract that implements part of a milestone/capability.

### Acceptance/Closeout

Evidence that the Product Milestone's own criteria are satisfied.

Example:

```text
Product Milestone:
M2 — Capability Registry and Reference Provider

Capabilities:
CAP-CAPABILITY-REGISTRY
CAP-PROVIDER-CONFORMANCE

Spikes:
SPK-001 Contract/AHDK independence
SPK-007 Build-bound trust

Implementation Missions:
future MIS-...
```

Product Milestone is not the same as an internal mission milestone.

---

## 14.5 Required milestone anatomy

Every Product Milestone must define:

```text
Outcome
Operator-visible value
Risk retired
Entry criteria
Capabilities involved
Architecture spikes
Golden Proof
Evidence requirements
Exit criteria
Telemetry baseline
Non-goals
Dependencies
Replan triggers
Promotion/authority boundary
```

A milestone cannot close because files exist or a demo looks plausible.

---

## 14.6 Roadmap principles

### Walking skeleton first

Prove the complete path with minimum depth.

### Vertical slices

A milestone crosses domain, storage, interaction, authority and evidence layers as needed.

### Sovereign Core before broad integration

Identity, state and authority must not depend on an external harness or provider.

### Context before autonomy

Aurora should not act broadly before she can recover authoritative project context.

### Observe before actuate

Laboratory progression begins read-only.

### One reference provider before ecosystem

Prove contracts and conformance before catalog growth.

### One AHDK before multiple languages

Stabilize semantics and Golden Path before parallel SDK maintenance.

### Durability before overnight campaigns

Long-running autonomy requires recovery, budgets and idempotency.

### Evidence before self-improvement promotion

Improvement can be researched earlier; production promotion waits for evaluation and rollback.

### Prove before generalize

No generic platform abstraction without a present consumer and proof.

---

# 14.7 A0 — Product, Discovery and Architecture Baseline

## Outcome

A complete, reviewable and session-resilient constitutional package defines Aurora's product identity, relationship, domain, lifecycle, memory, capability/harness architecture, presence, laboratory direction, authority, security, system architecture, reliability, roadmap and documentation governance.

## Operator-visible value

Leandro can leave the current chat and a new session can recover not only the decisions but the reasoning, examples, boundaries and current gate.

## Risk retired

```text
Project intent exists only in conversation and is lost across sessions.
```

## Entry criteria

- repository exists;
- initial discovery dialogue completed;
- MNFS available as documentation benchmark;
- operator authorizes documentation and research.

## Capabilities involved

- documentation governance;
- research governance;
- product modeling;
- decision traceability;
- status/handoff.

## Architecture spikes

None required to write the constitution. Technical claims remain proposed until later spikes.

## Golden Proof

```text
1. Start a fresh session with repository access only.
2. Read AGENTS, Documentation Map, Status and relevant Blueprint sections.
3. Explain Aurora's North Star, scope, relationship and autonomy model.
4. Explain memory classes and authority precedence.
5. Explain Aurora–harness boundary and AHDK policy.
6. Identify open technical decisions.
7. Refuse to implement because A0 remains in review.
8. Point to the exact next approved action.
```

## Evidence requirements

- complete fifteen-section Blueprint;
- discovery history;
- research reports and source manifests;
- ADR proposals;
- Capability Realization Method;
- requirements coverage;
- link/ID/metadata validation;
- fresh-session review record.

## Exit criteria

- operator accepts constitutional content;
- gaps and contradictions resolved;
- all approved discovery decisions mapped;
- no silent stack selection;
- A0 closeout recorded.

## Telemetry baseline

Documentation validation only:

- file/ID/link coverage;
- source manifest integrity;
- decision coverage;
- fresh-session comprehension findings.

## Non-goals

- Core implementation;
- AHDK implementation;
- running architecture spikes;
- MNFS integration;
- hardware control.

## Replan triggers

- operator finds missing intent;
- new discovery changes constitutional scope;
- research invalidates global architecture direction.

---

# 14.8 M0 — Sovereign Core Walking Skeleton

## Outcome

A minimal Aurora Core persists identity, project state, authority snapshot and one interaction lifecycle across restart without depending on an external model or harness as authority.

## Operator-visible value

Leandro initializes Aurora, records a project state, restarts the process and receives the exact current state and permitted next action.

## Risk retired

```text
Aurora is merely a running session; restart destroys identity and state.
```

## Entry criteria

- A0 accepted;
- Core boundaries approved;
- minimal domain/entity spec;
- storage and language spikes complete enough for one local implementation;
- backup/restore and migration strategy for the slice.

## Capabilities

- sovereign identity;
- project registry;
- operational state;
- authority snapshot;
- event/audit minimum;
- CLI or simple interface.

## Golden Proof

```text
initialize Aurora instance
→ create project
→ record accepted state and next action
→ terminate all Aurora processes
→ start fresh process
→ recover same identities/state
→ reject an invalid transition
→ export and restore state
```

## Evidence

- state hashes/IDs;
- restart receipt;
- invalid transition test;
- backup/restore result;
- no transcript dependency.

## Non-goals

- conversational memory;
- model routing;
- harness registry;
- voice;
- multi-device;
- cloud;
- physical devices.

## Replan triggers

- store cannot preserve required state simply;
- domain model proves too broad for slice;
- operational burden exceeds single-user baseline.

---

# 14.9 M1 — Governed Conversation, Project Context and Memory

## Outcome

Aurora continues a project across sessions using governed memory, provenance, authority precedence, correction and supersession.

## Operator-visible value

Leandro can stop and resume a project without re-explaining accepted decisions, while inspecting and correcting what Aurora remembers.

## Risk retired

```text
A capable model still behaves statelessly or retrieves plausible but wrong context.
```

## Entry criteria

- M0 state accepted;
- memory research current;
- CAP-MEMORY-AND-CONTEXT spec;
- evaluation corpus for project/session/supersession;
- privacy and retention policy for slice.

## Capabilities

- session ledger;
- project memory;
- observational memory candidate;
- Context Builder;
- provenance;
- memory promotion;
- edit/delete/supersede;
- model adapter minimum.

## Architecture spikes

- memory strategy comparison;
- context-building evaluation;
- deletion/derived-index behavior.

## Golden Proof

```text
Session A:
  discuss and explicitly approve a project decision
  record a competing hypothesis
  correct a mistaken preference

terminate session

Session B:
  resume project
  retrieve accepted decision with source
  retain hypothesis as unconfirmed
  exclude superseded preference
  explain provenance
  refuse cross-project memory
  delete selected memory and verify non-retrieval
```

## Evidence

- retrieval traces;
- authority resolution;
- temporal/supersession tests;
- cross-project isolation;
- deletion verification;
- quality/cost/latency baseline.

## Non-goals

- broad global life memory;
- ambient audio/video;
- self-improvement;
- multiple harnesses;
- autonomous effects.

---

# 14.10 M2 — Capability Registry, AHDK Kernel and Reference Provider

## Outcome

Aurora can describe a capability, ingest an exact provider manifest, run conformance/sandbox verification and approve one reference provider for a narrow scope.

## Operator-visible value

Leandro sees what Aurora can do, which provider offers it, what was verified and what remains forbidden.

## Risk retired

```text
Every integration is a one-off tool call and technical availability becomes implicit trust.
```

## Entry criteria

- Contract Model spec;
- ADR-0001/0002 accepted;
- SPK-001 and SPK-007 complete;
- one AHDK language selected by evidence;
- provider trust model specified.

## Capabilities

- Capability Definition;
- Provider Manifest;
- Registry;
- AHDK kernel;
- Conformance Kit;
- build identity/provenance;
- sandbox verification;
- scoped approval.

## Golden Proof

```text
create trivial provider using AHDK
→ publish exact manifest/build
→ Registry marks DISCOVERED
→ incompatible version rejected
→ sandbox and conformance pass
→ provider approved only for synthetic/internal data
→ rebuild changes digest
→ previous approval is not silently inherited
```

## Evidence

- schemas;
- conformance report;
- build/provenance record;
- approval scope;
- revocation test;
- AHDK versus direct implementation comparison from spike.

## Non-goals

- public marketplace;
- multiple SDK languages;
- real engineering harness;
- autonomous mission;
- remote federation.

---

# 14.11 M3 — Contractual Delegation, Artifacts and Evidence

## Outcome

Aurora converts an objective into a narrow Delegation, supplies a governed Context Pack and Authority Grant, receives structured progress, artifact, evidence and outcome.

## Operator-visible value

Leandro can ask Aurora to perform a bounded task and see exactly what was delegated, allowed, produced and verified.

## Risk retired

```text
Provider execution is opaque text with no stable contract, authority or evidence.
```

## Entry criteria

- M2 provider accepted;
- Delegation/Context/Authority specs;
- Effect Gateway minimum for slice;
- Artifact/Evidence metadata store;
- cancellation/error model.

## Capabilities

- mission/delegation;
- Context Pack;
- Authority Grant;
- provider selection;
- event ingestion;
- decision request;
- artifact/evidence;
- verdict/outcome;
- cancellation.

## Golden Proof

```text
Leandro requests transformation of an input artifact
→ Aurora selects approved provider
→ creates minimized Context Pack
→ provider requests one allowed write effect
→ disallowed network effect denied
→ provider publishes artifact and claim
→ independent check creates receipt/evidence
→ Aurora closes outcome
→ project memory records result by reference
```

## Non-goals

- long-running restart;
- child delegation;
- real code/lab changes;
- broad provider catalog.

---

# 14.12 M4 — Durable Delegation, Budget and Recovery

## Outcome

A Delegation survives Core and provider restart, timers and human waits, while preserving budgets and avoiding duplicate effects.

## Operator-visible value

Leandro can leave an authorized job running and trust that restart or network loss will not erase progress or repeat an external action blindly.

## Risk retired

```text
Long-running Aurora work depends on one process and ambiguous retry can duplicate effects.
```

## Entry criteria

- SPK-004 durable-engine result;
- effect idempotency/reconciliation spec;
- backup/recovery;
- budget model;
- exact version binding;
- fault-injection environment.

## Capabilities

- DurableExecutionPort;
- checkpoint;
- heartbeat;
- timers;
- resume;
- cancellation;
- idempotency;
- budget enforcement;
- reconciliation;
- degraded state.

## Golden Proof

```text
start delegation
→ perform one external effect with idempotency key
→ checkpoint
→ terminate Core and provider
→ restart both
→ recover state
→ prove effect not duplicated
→ hit budget warning
→ wait for decision
→ continue and complete
```

## Non-goals

- high availability cluster;
- overnight adaptive campaign;
- multi-harness collaboration;
- cloud-first execution.

---

# 14.13 M5 — Hierarchical Multi-Harness Composition

## Outcome

Aurora coordinates two reference providers with child Delegation, separate authority and optional governed direct data channel.

## Operator-visible value

Leandro gives one cross-domain objective and Aurora composes specialized work without manually copying context between systems.

## Risk retired

```text
Harnesses either become isolated silos or share context/credentials peer-to-peer without global control.
```

## Entry criteria

- M4 durability;
- child-delegation contract;
- data-channel contract;
- capability request flow;
- global acceptance criteria;
- provider substitution rules.

## Capabilities

- global decomposition;
- child capability request;
- child grant/context;
- dependency graph;
- direct data channel;
- artifact handoff;
- global outcome;
- cross-provider trace.

## Golden Proof

```text
Provider A receives delegation
→ requests capability offered by Provider B
→ Aurora creates child delegation with narrower authority
→ B produces artifact/data
→ authorized direct channel transfers bounded data
→ channel revokes at end
→ A consumes result
→ Aurora closes global criteria
→ no authority is inherited transitively
```

## Non-goals

- peer-to-peer autonomous federation;
- arbitrary swarm;
- real MNFS/lab integration;
- unlimited concurrency.

---

# 14.14 M6 — First Real Engineering Harness

## Outcome

A real engineering domain provider integrates through the same Contract Model and evidence requirements proven by reference providers.

## Candidate selection

Candidate is chosen by readiness and risk, not prestige:

- Research Harness;
- Evaluation Harness;
- MNFS;
- Firmware Harness;
- Hardware analysis provider.

MNFS is a strong future candidate but is not architecturally mandatory or assumed ready.

## Operator-visible value

Aurora performs useful real engineering work while preserving context, authority and evidence.

## Risk retired

```text
The architecture works only with toy providers.
```

## Entry criteria

- provider domain spec;
- stable external boundary;
- conformance;
- known effects;
- domain evals;
- operational owner;
- rollback/containment.

## Golden Proof

Defined after provider selection. It must include a real outcome, not just successful connection.

## Non-goals

- multiple real harnesses at once;
- broad autonomy;
- physical actuation unless the selected domain requires only observation.

---

# 14.15 M7 — Delegated Experimental Campaigns

## Outcome

Aurora conducts a multi-cycle adaptive campaign under objective, baseline, immutable evaluation, budgets, guardrails, stop and escalation rules.

## Operator-visible value

Leandro can authorize overnight improvement work without confirming every experiment.

## Risk retired

```text
Autonomy is either too passive to be useful or too broad to be trusted.
```

## Entry criteria

- M4 durability;
- M6 real provider;
- experiment model;
- evaluation and holdout governance;
- campaign budget;
- no automatic production promotion;
- failure/recovery drills.

## Golden Proof

```text
approved baseline and mutable space
→ Aurora formulates competing hypotheses
→ runs multiple variants
→ adapts next tests from evidence
→ respects hard budget
→ stops on convergence/no-progress
→ protects holdout and production
→ returns reproducible report
→ requires approval for promotion
```

## Non-goals

- continuous self-directed programs;
- constitutional self-modification;
- unattended hazardous physical campaign.

---

# 14.16 M8 — Multi-Presence Continuity

## Outcome

Same Aurora preserves one activity between workstation and a second trusted presence while adapting disclosure, input/output and authority to environment.

## Operator-visible value

Leandro can leave the computer and continue through mobile/wearable without restarting or exposing the whole project.

## Risk retired

```text
Aurora remains an application tied to one screen or leaks context during handoff.
```

## Entry criteria

- Presence Fabric spec;
- device trust and revoke;
- environment/privacy model;
- contextual handoff;
- local degraded mode;
- sensor indicators;
- mobile/wearable prototype.

## Golden Proof

```text
private workstation has confidential context
→ handoff to second presence in public/uncertain environment
→ safe summary only
→ critical campaign alert delivered
→ step-up authentication
→ authorized detail resumes
→ device goes offline and declares limitation
→ events reconcile on reconnect
→ remote revoke blocks access
```

## Non-goals

- continuous camera recording;
- multiple users;
- fully custom glasses hardware;
- physical actuation from wearable by default.

---

# 14.17 M9 — Laboratory Observation

## Outcome

Aurora identifies a device, ingests governed telemetry and accompanies a protocol without commanding physical actuation.

## Operator-visible value

Aurora can watch and explain a real experiment while preserving measurements, evidence and project continuity.

## Risk retired

```text
Physical integration begins with unsafe control before identity, telemetry and evidence are trustworthy.
```

## Entry criteria

- Device Registry/manifest;
- controller sandbox;
- telemetry schema/time model;
- calibration record;
- laboratory protocol;
- read-only authority;
- privacy/security review.

## Golden Proof

```text
identify exact board/instrument
→ verify firmware and calibration
→ ingest telemetry with units/time
→ correlate protocol step and observations
→ detect injected anomaly
→ preserve artifact/evidence
→ prove command path remains unavailable
```

## Non-goals

- changing source/load settings;
- autonomous flash;
- high-rate video retention;
- safety-critical control.

---

# 14.18 M10 — Controlled Physical Actuation

## Outcome

Aurora executes a narrow physical action through authority, device gateway and independent interlocks.

## Operator-visible value

Aurora can safely configure and run a bounded laboratory protocol rather than only observe.

## Risk retired

```text
Physical commands rely on model judgment or ungoverned device APIs.
```

## Entry criteria

- M9 accepted;
- risk analysis;
- device gateway;
- deterministic command validation;
- interlock and emergency stop;
- guided-manual predecessor;
- physical drills;
- exact recovery/cleanup.

## Golden Proof

```text
approved protocol and grant
→ preflight verifies device, wiring, calibration and interlock
→ Aurora requests bounded action
→ gateway enforces range
→ physical system responds
→ forced fault triggers independent containment
→ cloud/model loss does not prevent safe state
→ receipts and telemetry reconstruct timeline
```

## Non-goals

- broad autonomous lab;
- hazardous/high-voltage unattended operation;
- disabling interlocks;
- changing absolute limits during campaign.

---

# 14.19 M11 — Governed Self-Improvement Campaign

## Outcome

Aurora correlates multiple failures, investigates a shared cause, creates a candidate, evaluates broadly and proposes supervised promotion with canary/rollback.

## Operator-visible value

Aurora becomes measurably better without accumulating narrow prompt patches or silently rewriting herself.

## Risk retired

```text
Self-improvement overfits one incident, validates itself and changes production without independent evidence.
```

## Entry criteria

- Failure Intelligence;
- evaluation datasets and holdout;
- candidate versioning;
- independent reviewer;
- shadow/canary routing;
- rollback;
- protected constitutional areas;
- experiment authorization.

## Golden Proof

```text
four incidents share suspected cause
→ reproduce and compare competing hypotheses
→ candidate targets causal mechanism
→ original/neighbor/contrary/history/unseen/adversarial tests
→ multi-objective evaluation
→ independent review
→ shadow then narrow canary
→ forced regression triggers rollback
→ constitutional policy unchanged
```

## Non-goals

- autonomous constitutional changes;
- automatic broad promotion;
- model-weight training requirement;
- treating user satisfaction as sole objective.

---

# 14.20 M12 — Continuous Engineering Companion

## Outcome

Aurora coordinates ongoing research, design, software, hardware, firmware, laboratory experimentation, documentation, memory and learning across presences.

## Operator-visible value

The North Star becomes routine:

> Leandro enters the laboratory and continues any project with shared context, coordinated capabilities, observed evidence and a clear next decision.

## Directional capabilities

- mature global/project memory;
- multiple real harnesses;
- delegated campaigns;
- multimodal presence;
- laboratory observation and controlled actuation;
- proactive attention management;
- longitudinal evaluation;
- governed self-improvement.

## Status

```text
STRATEGIC_DIRECTION
```

M12 is not an implementation commitment. It exists to evaluate whether earlier architecture preserves the destination.

---

## 14.21 Architecture spike portfolio

Initial spikes before or during early milestones:

### SPK-001 Contract/AHDK independence

Prove same capability through official SDK and direct protocol implementation.

### SPK-002 MCP mapping

Test tool/resource/task semantics, auth and conformance.

### SPK-003 A2A mapping

Test Agent Card, task, artifacts, streaming, cancel and TCK.

### SPK-004 Durable execution

Compare minimal baseline, DBOS, Restate and Temporal against local-first requirements.

### SPK-005 Authority and Effect Gateway

Prove scoped token/grant, denial, revocation and receipt.

### SPK-006 Distributed trace

Trace Core → provider → tool → gateway → artifact/evidence.

### SPK-007 Registry and provenance

Prove approval is bound to exact build/provenance.

### SPK-008 Framework neutrality

Same capability in two internal runtimes without changing Aurora mission semantics.

Additional spikes require explicit uncertainty and removal conditions.

---

## 14.22 Milestone status vocabulary

```text
DIRECTIONAL
DISCOVERY
READY_FOR_SPEC
READY_FOR_PLAN
AUTHORIZED
IN_PROGRESS
VERIFYING
ACCEPTED
BLOCKED
SUPERSEDED
CANCELED
```

A distant milestone remains `DIRECTIONAL`. Documentation detail does not authorize it.

---

## 14.23 Readiness progression

```text
R0 Constitutional alignment
R1 Research/applicability
R2 Requirements derived
R3 Capability Spec ready
R4 Architecture decisions/spikes ready
R5 Mission contract ready
R6 Implementation plan/microdesign ready
R7 Execution and evidence
R8 Product Milestone closeout
```

The Capability Realization Method owns exact gates and artifacts.

---

## 14.24 Roadmap change

Roadmap changes require:

- reason;
- new evidence;
- impact on earlier milestones;
- dependency changes;
- risks;
- supersession history;
- operator approval when material.

Changing order because a framework is popular is not sufficient.

---

## 14.25 Global replan triggers

- research contradicts a constitutional assumption;
- memory quality cannot meet continuity needs;
- security/safety boundary proves insufficient;
- protocol maturity changes materially;
- provider readiness changes;
- operational burden exceeds single-user value;
- physical integration requires separate architecture;
- AHDK semantics do not generalize across two providers;
- durable engine creates unacceptable lock-in;
- Leandro changes product scope;
- Golden Proof cannot be demonstrated without hidden assumptions.

---

## 14.26 Non-goals

The roadmap does not:

- choose implementation dates;
- authorize all milestones;
- require MNFS as M6 provider;
- require physical autonomy;
- require cloud deployment;
- require multi-tenancy;
- guarantee M12 exactly as imagined;
- prevent evidence-based reordering;
- allow skipping foundational safety and context capabilities for a more impressive demo.
