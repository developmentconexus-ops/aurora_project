---
id: DOC-AURORA-BLUEPRINT-06
title: Memória, Conhecimento e Construção de Contexto
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
  - memory principles and strata
  - context construction boundaries
  - memory promotion, supersession, retention and deletion
  - epistemic and temporal memory model
related:
  - DOC-AURORA-BLUEPRINT-01
  - DOC-AURORA-BLUEPRINT-03
  - DOC-AURORA-BLUEPRINT-04
  - DOC-AURORA-BLUEPRINT-11
  - DOC-AURORA-BLUEPRINT-12
  - DOC-AURORA-BLUEPRINT-13
review_triggers:
  - memory taxonomy changes
  - context precedence changes
  - memory promotion/deletion policy changes
  - technical memory architecture decision
last_reviewed: 2026-08-06
---

# 6. Memória, Conhecimento e Construção de Contexto

## 6.1 Propósito

Memory is one of the capabilities with the greatest impact on Aurora's observed intelligence.

A powerful model with stale, irrelevant or incorrectly authoritative context can perform worse than a smaller model with the correct state and evidence.

The memory problem is not:

> “How do we store all conversations?”

It is:

> **How does Aurora preserve durable experience, distinguish what kind of knowledge it is, govern its authority and sensitivity, and construct the smallest correct context for the current situation?**

A robust system must support:

- conversational continuity;
- project continuity across months/years;
- personal/global context;
- temporal and relational knowledge;
- exact historical recovery;
- operational mission recovery;
- source authority and provenance;
- correction and supersession;
- forgetting and deletion;
- local-first sovereignty;
- evaluation of recall and non-recall;
- growth without loading everything into a model.

The technical implementation remains deliberately open. “Use a vector database,” “use observational memory,” or “use a knowledge graph” is not a complete architecture decision.

---

## 6.2 Foundational distinction

| Concept | Meaning | Typical authority |
|---|---|---|
| Raw history | exact record of conversation, tool call or event | historical evidence |
| Memory | preserved information to guide future context | supporting unless promoted |
| Knowledge Source | document, code, paper, database, device or measurement | source-specific |
| Operational state | what is active and true now for missions/devices | runtime authority |
| Active context | selected information supplied for one decision | compiled projection |
| World/project model | structured entities and relationships | composite references |
| Source of truth | owner that governs a fact/decision/state | explicit authority |

> **Memory guides reasoning. Authority, evidence and live state determine operational truth.**

### Example

Observational memory says:

```text
The firmware test was completed successfully.
```

Current operational state says:

```text
RUN-024 = FAILED
reason = telemetry_gap
```

The run remains failed. Aurora may inspect the original session to understand why the summary became wrong, but she cannot close the criterion from memory.

---

## 6.3 Memory is a managed lifecycle

Memory requires three major operations:

```text
WRITE
→ what becomes a candidate, at which scope and with which provenance?

MANAGE
→ consolidate, relate, correct, supersede, expire, archive, delete

READ
→ retrieve, rank, verify, minimize and compile context
```

Storage is only one implementation concern inside this lifecycle.

---

## 6.4 Conceptual architecture

```text
RAW EXPERIENCE
conversations • actions • tools • telemetry • documents
        │
        ▼
CAPTURE AND OBSERVATION
what happened • source • time • scope • quality
        │
        ▼
CANDIDATES
facts • events • preferences • decisions • hypotheses • patterns
        │
        ▼
CONSOLIDATION AND GOVERNANCE
classify • deduplicate • relate • validate • supersede • expire
        │
        ├──────────────┬──────────────┬───────────────┐
        ▼              ▼              ▼               ▼
Personal/Global    Project/Domain   Episodes       Procedures/Failures
        └──────────────┴──────────────┴───────────────┘
                               │
                               ▼
                         CONTEXT BUILDER
scope • authority • recency • relevance • sensitivity • budget
                               │
                               ▼
                      ACTIVE REASONING CONTEXT
                               │
                               ▼
                    ACTION • OBSERVATION • LEARNING
```

Each stage can have different mechanisms and stores.

---

## 6.5 Memory strata

Aurora should distinguish logical strata even if initial implementation uses fewer physical stores.

```text
L0 — Authoritative Sources and Operational State
L1 — Current Authority/Project Snapshots
L2 — Governed Durable Memory
L3 — Observational and Session Synthesis
L4 — Exact Interaction/Tool/Event History
L5 — Ephemeral Active Context and Transport
```

### L0 — Authoritative Sources and Operational State

Examples:

- Product Blueprint;
- accepted ADR;
- current project status;
- active mission/delegation;
- Authority Grant;
- source code commit;
- device telemetry;
- calibration record;
- effect receipt;
- evidence artifact.

L0 determines current action according to domain authority.

### L1 — Current snapshots

Compiled projections such as:

- current project snapshot;
- current authority snapshot;
- active blockers;
- permitted next actions;
- current device snapshot;
- mission checkpoint.

L1 is generated from authoritative sources and includes freshness/hash references.

### L2 — Governed durable memory

Personal, project, episodic, relational, procedural and failure memories with metadata/lifecycle.

### L3 — Observational/session synthesis

Compressed observations and reflections supporting continuity. Probabilistic and replaceable.

### L4 — Exact history

Raw conversation, tool calls, outputs, sensor events or session ledger, retrieved on demand.

### L5 — Ephemeral context/transport

Prompt context, message queues, streams and temporary buffers. Not memory or durable state.

---

# 6.6 Memory classes

## 6.6.1 Working memory

Scope:

```text
current turn/task/short activity
```

Contains:

- immediate objective;
- recent messages;
- current hypothesis;
- pending tool calls;
- temporary variables;
- next step;
- unresolved ambiguity.

Properties:

- high relevance;
- small budget;
- short retention;
- may be reconstructed from session/mission state;
- should not become durable automatically.

Failure modes:

- losing a constraint during a long response;
- mixing two active threads;
- retaining transient assumption as fact;
- overflow/compaction removing the current objective.

Evaluation:

- task continuity;
- constraint adherence;
- correct thread separation;
- recovery after compaction.

---

## 6.6.2 Conversational memory

Preserves continuity within and across conversations:

- what was asked;
- answers and corrections;
- alternatives discussed;
- unresolved questions;
- tone/mode;
- references to artifacts;
- explicit approvals.

It combines:

- exact history;
- summaries;
- observational notes;
- durable promoted items.

The raw transcript is not automatically current product truth.

Example:

A conversation says “we may use PostgreSQL.” Later an ADR selects SQLite for the milestone. Conversational memory preserves the earlier alternative but Context Builder marks the ADR as current authority.

---

## 6.6.3 Observational memory

Observational memory transforms ongoing history into time-aware, dense observations.

Example:

```text
2026-08-05
- Leandro approved Aurora as Leandro-first and single-user in the current horizon.
- Engineering is the first deep operational domain.
- Aurora will use hierarchical contractual orchestration.
```

Potential operations:

- Observer extracts material events/constraints;
- Reflector reorganizes or consolidates when observations grow;
- source references remain available;
- recent raw history remains active;
- old detail moves out of prompt while remaining retrievable.

Benefits to test:

- long conversational continuity;
- temporal event tracking;
- lower context size;
- preservation of decisions/constraints.

Limits:

- observations may omit critical details;
- summaries may be wrong;
- reflections can overgeneralize;
- source authority may be lost;
- benchmarks often measure conversation recall, not project/physical correctness.

Observational memory is a candidate mechanism for L3, not the entire Aurora memory system.

---

## 6.6.4 Episodic memory

Represents meaningful events/experiences.

Examples:

- firmware test run;
- architecture review;
- failed deployment;
- laboratory incident;
- successful project session;
- delegated overnight campaign;
- device handoff.

Conceptual structure:

```yaml
episode:
  id: EPI-POWER-024
  type: firmware_thermal_test
  project: PRJ-POWER-SUPPLY
  started_at: ...
  ended_at: ...
  participants:
    - LEANDRO
    - AURORA
    - LAB-HARNESS
  objective: compare control strategy C
  outcome: INCONCLUSIVE
  key_observations:
    - temperature plateau not reached
    - telemetry gap after minute 42
  artifacts:
    - waveform-024
    - telemetry-run-024
  decisions:
    - repeat with corrected logger
```

Episodes link to raw artifacts rather than embedding every detail.

---

## 6.6.5 Project memory

Each project has a bounded continuity universe:

- vision;
- goals;
- architecture;
- decisions;
- status;
- roadmap;
- tasks/missions;
- hypotheses;
- experiments;
- devices;
- repositories;
- people;
- risks;
- incidents;
- evidence;
- next actions.

Project memory should be the primary context when Leandro says:

> “Vamos continuar o projeto da fonte.”

It should not load unrelated Metal Nobre or Aurora details unless relevant.

Project memory combines references to authoritative documents/state and durable supporting memory.

---

## 6.6.6 Global personal memory

Knowledge relatively stable across projects:

- identity;
- engineering background;
- skills;
- long-term objectives;
- communication preferences;
- preferred research rigor;
- work patterns;
- stable constraints;
- relationship preferences.

This scope has high impact and privacy risk.

Example:

```text
Observed:
Leandro selected Go in three projects.

Not sufficient for:
"Leandro always prefers Go for all backends."
```

A global preference needs explicit confirmation or strong repeated evidence and should remain a preference, not mandatory policy.

---

## 6.6.7 Relational memory

Preserves relationships:

```text
Leandro CREATES Aurora
Aurora USES MNFS as a future provider
Project PowerSupply HAS Device PCB-REV-B
PCB-REV-B RUNS Firmware FW-014
Hypothesis HYP-017 TESTED_BY Experiment EXP-024
```

Relationships need:

- scope;
- time;
- provenance;
- confidence/authority;
- supersession.

Relational memory may be implemented using graph techniques, relational tables, documents or hybrid projections. Technology remains open.

Avoid creating a relation for every sentence. Only durable, useful relationships should enter the model.

---

## 6.6.8 Procedural memory

Preserves how to perform recurring work:

- first power-up protocol;
- firmware test sequence;
- research methodology;
- project handoff procedure;
- provider verification;
- rollback;
- incident containment.

Promotion path:

```text
successful execution
→ repeated pattern
→ procedure candidate
→ validation across contexts
→ approved protocol/Golden Path
```

A single success is observation, not procedure.

Procedural memory should link to versioned Standard/Golden Path when normative.

---

## 6.6.9 Failure and learning memory

Preserves:

- incidents;
- symptoms;
- causal hypotheses;
- failed approaches;
- corrections;
- regressions;
- evaluation cases;
- improvement candidates;
- lessons.

Purpose:

- prevent rediscovering the same failure;
- correlate shared causes;
- inform future plans;
- preserve why an approach was rejected.

A failed attempt must not dominate future context when the underlying condition changed. Applicability and version matter.

---

## 6.6.10 Operational memory/state

Includes:

- active mission/delegation;
- checkpoint;
- budget consumed;
- grants;
- provider run;
- pending decision;
- device lease;
- current firmware;
- stop condition;
- effect receipts.

This must be structured and authoritative in runtime state. It is called memory colloquially but cannot depend on model recall.

Example:

Aurora cannot “remember” a US$20 campaign budget from conversation; a deterministic budget record must enforce it.

---

## 6.7 Memory Item model

Conceptual example:

```yaml
memory:
  id: MEM-AURORA-000421
  type: DECISION_SUMMARY
  scope:
    level: PROJECT
    project: PRJ-AURORA
  content:
    statement: "First-party harnesses use AHDK by engineering policy."
    structured_refs:
      - ADR-AURORA-0002

  provenance:
    source_type: accepted_decision
    source_ref: ADR-AURORA-0002
    observed_at: 2026-08-05T22:00:00-03:00
    recorded_at: 2026-08-05T22:01:00-03:00
    extractor: aurora-memory-writer@version

  epistemic:
    kind: APPROVED_DECISION
    confidence: 1.0
    authority: OPERATOR_ACCEPTED

  temporal:
    valid_from: 2026-08-05
    valid_until: null
    freshness: source_on_use
    superseded_by: null

  governance:
    sensitivity: INTERNAL
    retention: project_lifetime
    editable: true
    deletable: conditional_on_source
    allowed_domains:
      - PRJ-AURORA
    cloud_policy: minimized_reference_only

  relationships:
    - type: SUMMARIZES
      target: ADR-AURORA-0002
```

Exact schema remains open. Required semantics should not be omitted by a simpler storage implementation.

---

## 6.8 Epistemic status

Aurora distinguishes at least:

```text
USER_STATED
USER_APPROVED
DOCUMENT_ESTABLISHED
SYSTEM_OBSERVED
INSTRUMENT_MEASURED
PROVIDER_REPORTED
AURORA_INFERRED
HYPOTHESIS
VERIFIED_RESULT
HISTORICAL
```

### Examples

**User stated:**

> “Acho que prefiro Go.”

**User approved:**

> “Aprovo Go para o serviço X.”

**Aurora inferred:**

> “Leandro pode preferir linguagens compiladas para control planes.”

These must not receive equal authority or scope.

### Confidence versus authority

A high-confidence inference remains an inference. A low-confidence measurement may still be the only live observation but cannot close a strict criterion.

---

## 6.9 Temporal model

Memory should answer:

- when was it observed?
- when was it recorded?
- during which interval was it valid?
- when was the source last verified?
- which item superseded it?
- does the current question require live refresh?

### Example

```text
2026-08-01: Firmware FW-013 running on PCB-B
2026-08-05: FW-014 flashed and verified
```

Context for current test uses FW-014. Historical question about incident on August 1 uses FW-013.

Temporal retrieval should not erase the old relationship; it should mark validity.

---

## 6.10 Scope model

Possible scopes:

```text
TURN
SESSION
TASK
DELEGATION
MISSION
PROJECT
DOMAIN
PERSONAL_GLOBAL
PRESENCE
DEVICE
ENVIRONMENT
```

A memory can have primary and secondary scopes.

Isolation rule:

- project-specific decisions do not enter another project by default;
- presence-local sensor observation does not become global personal memory automatically;
- provider-local state does not become Aurora global state;
- global preference can influence projects only where applicable.

---

## 6.11 Promotion policy

The selected constitutional policy is **promotion by risk, authority and scope**.

### Automatic low-risk promotion

- current task state;
- project event;
- tool/test result linked to evidence;
- explicit decision by Leandro within a project;
- clearly labeled hypothesis;
- temporary authority and expiry;
- explicit factual correction;
- project-local preference;
- generated project summary with source references.

### Conditioned promotion

- global personal preference;
- behavioral inference;
- sensitive information;
- third-party information;
- cross-project rule;
- constitutional change;
- permanent authority;
- mental/emotional/health inference;
- long-retention audio/video observation.

Condition can be:

- explicit confirmation;
- repeated evidence;
- source authority;
- review;
- policy;
- domain-specific consent.

> The greater the reach, sensitivity or authority, the stronger the promotion requirement.

---

## 6.12 Memory lifecycle

```text
OBSERVED
→ CANDIDATE
→ ACCEPTED_FOR_SCOPE
→ CONFIRMED
→ SUPERSEDED | EXPIRED | ARCHIVED | REMOVED
```

### OBSERVED

Captured event/statement; not necessarily durable.

### CANDIDATE

Potential future value identified; awaits policy/evidence.

### ACCEPTED_FOR_SCOPE

Usable in a bounded scope with known epistemic status.

### CONFIRMED

Strong source or explicit approval.

### SUPERSEDED

Replaced by newer/current information; retained historically if policy allows.

### EXPIRED

Validity/retention ended.

### ARCHIVED

Preserved outside normal retrieval.

### REMOVED

No longer retained/used; derived copies handled according to deletion policy.

---

## 6.13 Conflict and supersession

Conflict examples:

### Decision conflict

Memory:

```text
Project will use SQLite.
```

New user statement:

```text
Let's switch to PostgreSQL.
```

Aurora should say:

> “SQLite is currently an accepted decision. Your statement is a proposed architecture change. I can analyze impact and prepare an ADR supersession; I will not silently overwrite the current decision.”

### Preference conflict

Global memory says deep explanations preferred. Current user says “just give me the command.”

Current interaction preference applies locally; global preference remains.

### Live-state conflict

Memory says board disconnected. Device Registry sees authenticated active board. Live state governs current operation, and memory may be updated.

Conflict resolution should record:

- compared sources;
- authority;
- time;
- chosen current interpretation;
- unresolved ambiguity;
- supersession/action.

---

## 6.14 False and plausible memory

One of the highest-risk failures is a memory that sounds reasonable but was never established.

Example:

> “Leandro always prefers Go for backend.”

Possible causes:

- inference promoted as fact;
- repeated provider summary;
- source removed;
- project-local choice generalized globally;
- model hallucination during consolidation.

Controls:

- epistemic labels;
- source references;
- scope;
- promotion thresholds;
- user inspection;
- evaluation with false-premise cases;
- source-on-use for material decisions;
- no silent orphan memory.

Aurora should express:

> “Você usou Go em três serviços, mas não há preferência global confirmada.”

---

## 6.15 Deduplication and consolidation

Multiple memory items may represent the same event or evolving understanding.

Consolidation can:

- merge redundant observations;
- create a higher-level summary;
- retain source links;
- preserve exceptions;
- supersede stale synthesis;
- avoid repeated prompt content.

It must not:

- delete contradictory evidence silently;
- turn repeated rumor into fact;
- broaden scope;
- remove temporal distinction;
- lose original source.

The consolidated memory has its own provenance and version.

---

## 6.16 Forgetting, expiry and archive

Good memory requires controlled forgetting.

Candidates for expiry/compaction:

- transient tool outputs;
- repeated low-value conversation detail;
- temporary task variables;
- expired presence state;
- superseded summaries;
- old provider availability;
- temporary preferences.

Candidates for long retention:

- accepted decisions;
- project outcomes;
- safety incidents;
- experiment evidence;
- historical rationale;
- important corrections;
- constitutional approvals.

Forgetting mechanisms:

- delete;
- expire;
- archive outside normal retrieval;
- retain source but remove derived summary;
- compact raw detail into referenced artifact;
- reduce priority.

“Stored” and “eligible for active context” are separate.

---

## 6.17 Deletion and derived data

Deletion may need to affect:

- canonical memory row/document;
- observational summaries;
- embeddings/indexes;
- graph relationships;
- caches;
- local device copies;
- backups according to policy;
- external provider stores where possible;
- evaluation datasets.

The system should record a deletion/tombstone or verification without retaining the sensitive content itself when necessary for audit.

If the original authoritative document remains, deleting a memory summary does not delete the document. Aurora must explain the scope.

---

## 6.18 Context Builder

Context Builder compiles the decision-specific context.

### Inputs

- current interaction and intent;
- presence/environment;
- project/mission/delegation;
- authority snapshot;
- risk;
- capability/provider requirements;
- token/data budget;
- freshness requirements;
- model/provider policy.

### Retrieval sources

- current operational state;
- Product Blueprint/ADRs/contracts;
- project state;
- knowledge sources;
- durable memory;
- observational memory;
- exact history;
- live device/provider state.

### Operations

```text
determine scope
→ load authority first
→ retrieve candidates
→ verify freshness/source
→ resolve conflicts
→ rank relevance
→ minimize and redact
→ preserve provenance
→ allocate context budget
→ compile structured Context Pack
→ validate against provider policy
```

### Output

A Context Pack can contain:

- objective;
- current state;
- authority sources;
- relevant facts/memories;
- hypotheses;
- constraints;
- artifact references;
- permitted actions;
- data classification;
- freshness;
- omissions/limitations.

---

## 6.19 Context precedence

For current action, a default conceptual order:

```text
1. immediate physical/safety state
2. active authority, policy and guardrails
3. current mission/project/device operational state
4. accepted decisions/contracts/specs
5. verified evidence
6. current knowledge sources
7. governed durable memory
8. observational/session synthesis
9. exact historical transcript for context
10. external suggestions/unverified sources
```

This is not universal ranking. A historical question may prioritize exact transcript. A technical fact may prioritize current official documentation. The Context Builder must identify the question type.

---

## 6.20 Context minimization

More context can degrade performance, cost and privacy.

Minimization strategies:

- project scope;
- source references rather than full content;
- extract exact sections;
- retain recent window plus observations;
- structured authority snapshot;
- omit irrelevant personal data;
- provider-specific redaction;
- retrieve live source on demand;
- progressive disclosure/tool retrieval.

Aurora should not send complete personal memory to a coding or research provider unless explicitly necessary and approved.

---

## 6.21 Context contamination

Failure modes:

- memory from wrong project;
- old decision;
- another provider's internal state;
- hypothetical future treated as current;
- research recommendation treated as ADR;
- user quote treated as preference;
- malicious document instruction;
- third-party data exposed.

Controls:

- stable IDs;
- scope filters;
- authority labels;
- temporal validity;
- content trust classification;
- cross-project negative tests;
- source citation;
- context inspection/audit.

---

## 6.22 Live-state refresh

Context Builder must know when memory is insufficient.

Refresh triggers:

- device command;
- production/deploy state;
- current provider availability;
- current branch/commit;
- active grant/budget;
- calibration validity;
- time-sensitive external information;
- stale source threshold;
- conflict.

Example:

Before flashing firmware, Aurora reads live device identity and current running firmware. It does not rely on yesterday's project memory.

---

## 6.23 Local-first and cloud-assisted memory

Canonical memory, identity, policy and audit remain under Leandro's control.

External providers can receive minimized context under policy.

Flow:

```text
canonical local memory/state
→ Context Builder
→ data classification
→ minimization/redaction/reference
→ approved provider/model
→ result and provenance return
→ local recording/promotion policy
```

Provider output does not automatically become durable memory.

External memory services, embeddings or sync require:

- data class policy;
- retention/training understanding;
- export/delete;
- encryption/access;
- fallback;
- migration.

---

## 6.24 Memory quality dimensions

### Recall

Retrieve needed information.

### Precision/non-recall

Do not retrieve irrelevant or forbidden memory.

### Temporal accuracy

Use correct version/time.

### Epistemic accuracy

Distinguish fact, inference and hypothesis.

### Authority awareness

Prioritize current canonical source.

### Scope isolation

Avoid cross-project/domain leakage.

### Provenance

Explain source.

### Correction and deletion

Update/remove effectively.

### Efficiency

Latency, tokens, storage and compute.

### Longitudinal stability

Continue working as memory grows.

No single recall benchmark proves all dimensions.

---

## 6.25 Evaluation scenarios

### Session continuation

Remember unresolved question and explicit decision after restart.

### Supersession

Use new ADR while retaining old historical decision.

### False premise

Reject “as we decided to use X” when no decision exists.

### Cross-project isolation

Do not retrieve Metal Nobre data into Aurora project task.

### Temporal

Identify firmware used at incident time versus current firmware.

### Sensitive non-recall

Do not supply personal memory to public provider.

### Deletion

Deleted item no longer appears through text/vector/graph paths according to policy.

### Live refresh

Detect stale device memory and query current state.

### Observational compression

Preserve material decisions across long history while allowing exact source retrieval.

### Procedural

Use approved first-power protocol, not one-off successful notes.

---

## 6.26 Memory research program

Future focused research must compare:

### Approaches

- raw history and compaction;
- observational memory;
- RAG/vector retrieval;
- full-text search;
- structured relational storage;
- temporal knowledge graphs;
- event logs;
- hierarchical context/virtual memory;
- memory operating systems;
- learned management policies;
- hybrid architecture.

### Evidence

- official implementations;
- papers;
- benchmarks;
- reproducibility;
- open-source code;
- real project prototypes;
- privacy/deletion behavior;
- operational cost.

### Spikes

- conversational observational memory;
- project source authority retrieval;
- temporal/supersession;
- cross-project isolation;
- context assembly under token limits;
- deletion across indexes;
- longitudinal growth.

---

## 6.27 Technical architecture remains open

Potential logical components:

```text
Session Ledger
Observation Pipeline
Memory Writer
Consolidator/Reflector
Memory Stores by scope/type
Knowledge Source Index
Relationship/Temporal Store
Context Builder
Provenance Service
Retention/Deletion Worker
Evaluation Harness
```

Potential physical choices are not yet accepted.

The first implementation may intentionally support fewer memory classes while preserving the canonical model and migration path.

---

## 6.28 Failure modes

### Store everything

Noise, cost and privacy grow; relevant context degrades.

### One vector store as truth

Similarity ignores authority, time and exact state.

### Summary without source

Plausible memory cannot be verified.

### Automatic global promotion

Project-local behavior becomes personal rule.

### Silent overwrite

Decision history and causality disappear.

### Stale live state

Device/provider action uses old memory.

### Cross-project leakage

Sensitive context enters wrong task/provider.

### Deletion theater

Original row deleted while embeddings/summaries remain.

### Reflection overreach

Consolidator invents a general rule.

### Model-specific memory

Changing provider loses or changes the system's identity/context semantics.

---

## 6.29 Non-goals

This section does not:

- select a vector database;
- select a graph database;
- declare observational memory sufficient;
- retain all audio/video;
- make every sentence permanent;
- infer personal traits broadly;
- use memory to override authoritative sources;
- require one physical store per memory class;
- guarantee perfect recall;
- make embeddings or summaries the only history;
- permit provider-specific memory to become Aurora canonical memory automatically.
