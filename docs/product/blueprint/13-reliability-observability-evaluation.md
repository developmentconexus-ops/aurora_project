---
id: DOC-AURORA-BLUEPRINT-13
title: Confiabilidade, Observabilidade, Avaliação e Autoaperfeiçoamento
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
  - reliability principles
  - observability and evidence model
  - evaluation and calibration
  - Failure Intelligence and self-improvement lifecycle
related:
  - DOC-AURORA-BLUEPRINT-04
  - DOC-AURORA-BLUEPRINT-06
  - DOC-AURORA-BLUEPRINT-07
  - DOC-AURORA-BLUEPRINT-10
  - DOC-AURORA-BLUEPRINT-11
  - DOC-AURORA-BLUEPRINT-12
  - RESEARCH-AURORA-HARNESS-ARCHITECTURE-V1
review_triggers:
  - acceptance or evidence model changes
  - self-improvement promotion boundary changes
  - telemetry semantic changes
  - reliability objectives changes
last_reviewed: 2026-08-06
---

# 13. Confiabilidade, Observabilidade, Avaliação e Autoaperfeiçoamento

## 13.1 Propósito

Aurora will be probabilistic in reasoning but must be reliable as a system.

Reliability does not mean never being wrong. It means the system can:

- expose uncertainty;
- preserve authoritative state;
- observe the result of actions;
- detect divergence and failure;
- recover without inventing progress;
- verify claims with appropriate evidence;
- explain decisions and effects;
- learn from repeated patterns;
- change itself only through governed evaluation and rollback.

The product cannot judge quality by:

- fluent answers;
- amount of activity;
- number of tool calls;
- long plans;
- self-declared completion;
- a single universal benchmark;
- user approval obtained without presenting risk.

Reliability is a combination of domain correctness, operational durability, security, safety, evidence and calibrated interaction.

---

## 13.2 Reliability dimensions

### Continuity

Can Aurora resume the right project, mission and authority after interruption?

### Correctness

Are facts, calculations, code, decisions and measurements accurate enough for their purpose?

### Context quality

Did Aurora retrieve the right sources and ignore stale/irrelevant information?

### Authority correctness

Did every effect remain inside active permission?

### Operational reliability

Did processes, providers, stores and channels behave across failure and restart?

### Evidence quality

Can conclusions be reproduced and traced to observations?

### Safety

Did digital/physical guardrails and interlocks work?

### Security/privacy

Were data and credentials contained and minimized?

### Interaction quality

Did Aurora communicate uncertainty, urgency and next action correctly?

### Efficiency

Did Aurora achieve the outcome without unnecessary tokens, cost, experiments, interruptions or complexity?

One dimension cannot hide failure in another.

---

## 13.3 Claim, Receipt, Evidence, Verdict

```text
CLAIM
→ actor states that something is true or complete

RECEIPT
→ controlled mechanism records an action or verification result

EVIDENCE
→ one or more observations/artifacts support a criterion or hypothesis

VERDICT
→ permitted authority evaluates the evidence

OUTCOME
→ global result consolidates verdicts, limitations and remaining risk
```

Examples:

- a coding harness claims tests pass;
- CI receipt records the command and result;
- logs and test report form evidence for criteria;
- reviewer verdict accepts or rejects;
- Aurora composes mission outcome.

For physical work:

- Laboratory Harness claims stable output;
- instrument artifacts and calibration records form evidence;
- evaluation verifies metrics;
- safety authority confirms boundary compliance.

---

## 13.4 Evidence properties

Material evidence preserves:

- identity;
- criterion/hypothesis addressed;
- producer;
- verifier;
- method;
- environment;
- input versions;
- artifact references;
- observation time;
- result;
- uncertainty;
- limitations;
- integrity/hash;
- authority;
- retention.

Evidence may be:

```text
deterministic
statistical
observational
comparative
manual
instrumental
adversarial
historical
```

The required evidence profile depends on risk and claim.

---

## 13.5 Acceptance criteria hierarchy

Acceptance may exist at:

- product milestone;
- capability;
- mission;
- delegation;
- experiment;
- effect;
- self-improvement candidate.

Parents require composed evidence, not only child completion.

Example:

```text
all firmware variants built
≠ thermal campaign successful

all harness delegations completed
≠ global project objective achieved
```

Aurora evaluates global coherence and unresolved risk.

---

## 13.6 Observability model

Observability provides signals for decisions, not activity theater.

### Traces

Causal/correlated path across:

```text
presence
→ Aurora intent
→ mission/delegation
→ provider/harness
→ tool/effect gateway
→ artifact/evidence
```

### Metrics

Aggregated measurements such as:

- success by capability;
- latency;
- cost;
- retry/recovery;
- context retrieval quality;
- budget use;
- notification load;
- safety triggers;
- provider availability.

### Logs

Detailed diagnostic events with controlled sensitivity.

### Domain events

Durable meaningful changes, separate from telemetry.

### Receipts/audit

Security, authority and effect accountability.

OpenTelemetry is the current vendor-neutral baseline hypothesis for telemetry. Domain semantics remain Aurora-owned.

---

## 13.7 Semantic conventions

Initial conceptual span/event names:

```text
aurora.interaction
aurora.context.build
aurora.memory.retrieve
aurora.mission
aurora.delegation
aurora.provider.select
aurora.harness.run
aurora.capability.invoke
aurora.tool.call
aurora.effect.request
aurora.effect.execute
aurora.artifact.publish
aurora.evidence.record
aurora.decision.request
aurora.presence.handoff
aurora.device.command
aurora.experiment.run
aurora.improvement.evaluate
```

Common attributes may include:

- project/mission/delegation IDs;
- actor chain;
- provider/build;
- model;
- environment;
- data classification;
- authority grant reference;
- budget consumed;
- attempt;
- artifact/evidence links;
- error class;
- recovery action;
- redaction policy.

Sensitive content should not be copied into general telemetry.

---

## 13.8 Signal-to-decision rule

Every retained metric should answer:

```text
Which decision does this inform?
What threshold or pattern triggers action?
What is its coverage?
What can it not prove?
Who owns response?
```

Examples:

### Context retrieval precision

Informs whether Context Builder is selecting authoritative sources.

Cannot alone prove task correctness.

### Tool-call success

Informs adapter reliability.

Cannot prove the tool produced the desired product outcome.

### Notification acknowledgment

Informs delivery, not whether interruption was relevant.

### Token/cost

Informs efficiency, not quality.

Aurora should not produce a universal productivity score.

---

## 13.9 Reliability objectives

Future capabilities may define SLO-like objectives appropriate to scope.

Examples:

- recovery from Core restart;
- maximum lost durable state;
- provider status freshness;
- effect-receipt completeness;
- critical alert latency;
- memory retrieval quality;
- context contamination rate;
- budget enforcement;
- experiment reproducibility;
- interlock reaction time.

Do not define arbitrary numbers before architecture and risk are known. Each target must include measurement method and rationale.

---

## 13.10 Error taxonomy

### Cognitive

- intent misinterpretation;
- unsupported inference;
- hallucinated fact;
- wrong hypothesis;
- poor plan;
- provider selection error.

### Context/memory

- missing authoritative source;
- stale memory;
- irrelevant retrieval;
- cross-project contamination;
- false memory;
- supersession failure.

### Contract

- schema violation;
- incompatible version;
- invalid transition;
- missing required evidence.

### Operational

- process crash;
- provider unavailable;
- timeout;
- duplicate event;
- lost heartbeat;
- checkpoint failure;
- storage failure.

### Effect

- unauthorized request;
- denial;
- ambiguous outcome;
- duplicate effect;
- partial external action;
- compensation failure.

### Security/safety

- data leakage;
- credential exposure;
- prompt injection;
- provider compromise;
- interlock failure;
- unsafe command;
- privacy violation.

### Evaluation

- metric gaming;
- leaked holdout;
- insufficient sample;
- non-reproducible result;
- reviewer conflict;
- regression.

Error class determines recovery. Generic retries are prohibited for material errors.

---

## 13.11 Recovery and resilience

### Retry

Allowed only when:

- error is classified retryable;
- operation is idempotent or protected;
- budget permits;
- new attempt remains inside authority;
- retry cannot hide systematic failure.

### Resume

Continue from checkpoint with exact version and state.

### Reconcile

Compare Aurora global state, provider snapshot and target-system evidence.

### Substitute

Choose another approved provider while preserving contract and documenting differences.

### Compensate

Attempt a defined reversal for an effect when true atomicity is unavailable.

### Block/escalate

Required when state is ambiguous, risk material or decision belongs to Leandro.

### Degrade

Continue with reduced capability while stating limitations.

---

## 13.12 Evaluation system

Aurora requires evaluation at several layers.

### Component eval

Memory retrieval, provider selection, policy decision, schema mapping.

### Capability eval

Can a provider produce the required outcome under representative conditions?

### Journey eval

Can the complete North Star or campaign flow succeed?

### Safety/security eval

Can the system resist forbidden effects and contain faults?

### Interaction eval

Does Aurora communicate clearly, challenge appropriately and avoid unnecessary interruption?

### Longitudinal eval

Does performance remain useful as projects, memories and providers grow over months?

---

## 13.13 Evaluation dataset governance

Evaluation sets may be:

```text
investigation
→ understand failure

development
→ build candidate

validation
→ compare candidates

holdout
→ protected final estimate

adversarial
→ attack assumptions and boundaries

production shadow
→ observe without governing output
```

Controls:

- versioned dataset;
- source/provenance;
- data classification;
- leakage prevention;
- immutable scoring rules during a campaign;
- representative coverage;
- known limitations;
- human review where required.

Aurora cannot change the evaluation to favor its candidate.

---

## 13.14 Multi-objective evaluation

A candidate can improve accuracy while worsening:

- latency;
- cost;
- privacy;
- token use;
- notification burden;
- safety;
- personality consistency;
- robustness;
- explainability.

Evaluation should use a Pareto view rather than a single scalar when trade-offs are material.

Example:

```text
Candidate A
+12% context accuracy
+35% token cost
-18% response speed
+9% false conflict alerts
```

Aurora presents the trade-off and recommendation. Leandro decides material product priorities.

---

## 13.15 Calibration

Aurora should calibrate confidence and escalation.

Calibration questions:

- when Aurora says 80% confidence, how often is the class correct?
- does confidence reflect source authority or only model certainty?
- does Aurora over-escalate low-risk uncertainty?
- does she understate physical risk?
- are provider trust estimates current?

Confidence should be decomposable where possible:

```text
intent confidence
source confidence
freshness
measurement quality
model confidence
verification coverage
```

One number should not hide incompatible sources of uncertainty.

---

## 13.16 Human evaluation

Leandro's feedback may indicate:

- factual error;
- poor relevance;
- wrong tone;
- excessive interruption;
- missing context;
- incorrect authority assumption;
- disagreement with a trade-off;
- personal preference.

Feedback must be classified before learning.

A negative reaction to justified disagreement should not automatically train Aurora to agree. A correction to a factual project state should update the appropriate source.

---

## 13.17 Failure Intelligence

`Failure Intelligence` is the subsystem/process that turns incidents and evaluation into causal improvement.

It maintains relationships between:

- incident;
- symptom;
- impact;
- context;
- affected mechanism;
- causal hypothesis;
- contributing factor;
- missing control;
- reproduction;
- candidate;
- evaluation;
- promotion;
- regression;
- learning.

### Principle

> Investigate which mechanism allowed the failure, not only which answer was wrong.

---

## 13.18 Structured incident example

```yaml
incident:
  id: INC-0031
  type: AUTHORITY_CONTEXT_OMISSION
  project: PRJ-AURORA
  observed_behavior: recommended a rejected provider
  expected_behavior: retrieve accepted provider decision
  impact: medium
  occurred_at: ...
  available_sources:
    - ADR-AURORA-0012
    - current-provider-approval
  sources_used:
    - conversational_memory
  sources_ignored:
    - ADR-AURORA-0012
  initial_hypotheses:
    - Context Builder failed to retrieve authority snapshot
    - retrieved source lost priority during compilation
  evidence:
    - trace-context-build-008
    - response-artifact-031
```

The record preserves observation without prematurely declaring root cause.

---

## 13.19 Correlation and causal graph

Seemingly separate incidents:

```text
INC-014 forgot a decision
INC-022 recommended rejected tool
INC-031 ignored roadmap gate
INC-038 repeated an answered question
```

Potential graph:

```text
Context assembly failure
├── authority source not retrieved
│   ├── rejected tool recommended
│   └── roadmap gate ignored
├── project memory not scoped
│   └── repeated question
└── old conversational summary overweighted
    └── decision forgotten
```

Correlation uses evidence such as shared component, trace pattern, source omission and timing. It does not merge incidents merely because text appears similar.

---

## 13.20 Causal investigation

Steps:

1. define observed failure;
2. assess impact and urgency;
3. reproduce when possible;
4. inspect exact context, versions and traces;
5. list competing hypotheses;
6. identify discriminating tests;
7. test cheap/high-information hypotheses first;
8. record supporting and contradicting evidence;
9. classify confirmed, supported, refuted or inconclusive;
10. select intervention at the lowest correct causal layer.

### Unreproducible failure

Aurora should not invent a confident cause. She may:

- add instrumentation;
- preserve evidence;
- watch recurrence;
- create a bounded hypothesis;
- avoid risky change.

---

## 13.21 Improvement Candidate

A candidate contains:

- identity/version;
- causal hypothesis;
- targeted mechanism;
- incident set;
- change description;
- expected benefit;
- risk;
- protected areas;
- test plan;
- evaluation sets;
- rollback;
- promotion scope;
- owner;
- status.

```text
PROPOSED
→ AUTHORIZED_FOR_EXPERIMENT
→ EXPERIMENTING
→ EVALUATED
→ REVIEWED
→ PROMOTION_PROPOSED
→ SHADOW
→ CANARY
→ PROMOTED | REJECTED | ROLLED_BACK | DEFERRED
```

---

## 13.22 Broad evaluation requirement

A correction is not improvement if it only memorizes the original case.

Required categories:

### Original

Does it resolve the motivating incident?

### Neighbor

Does it improve related symptoms of the same cause?

### Contrary

Does it avoid triggering when the condition is absent?

### Historical regression

Does it preserve previously correct behavior?

### Unseen/holdout

Does it generalize beyond cases used to create it?

### Adversarial

Can inputs bypass, game or overconstrain the new rule?

### Systemic

What happens to cost, latency, safety, privacy, personality and other capabilities?

---

## 13.23 Independent review

The proposer is not the sole verifier or promoter.

Independent review can combine:

- deterministic tests;
- separate evaluation harness;
- reviewer model with isolated context;
- human review;
- security review;
- physical/safety verification;
- blind comparison.

Independence is risk-proportional, not necessarily a separate organization.

---

## 13.24 Shadow, canary and rollback

### Shadow

Candidate observes or produces parallel output without governing action.

### Canary

Candidate governs a narrow approved scope:

- one project;
- one capability;
- percentage of interactions;
- limited time window;
- non-critical effects.

### Promotion

Requires evidence and explicit authority according to class.

### Rollback

Must preserve:

- prior version;
- compatible state;
- configuration;
- migration reversal or safe fallback;
- trigger conditions;
- evidence of rollback test.

Aurora never deletes the only known working version before replacement is proven.

---

## 13.25 Protected constitutional areas

Aurora cannot autonomously promote changes to:

- Leandro's final authority;
- constitutional identity and principles;
- security and safety invariants;
- promotion rules;
- audit requirements;
- revocation and shutdown;
- ability to grant authority;
- evidence independence requirements.

She may research and propose constitutional changes for explicit review.

---

## 13.26 Continuous detection, governed experimentation

Aurora may continuously:

- collect incidents and corrections;
- detect recurring failure clusters;
- identify cost/latency degradation;
- compare providers;
- notice poor context retrieval;
- propose new instrumentation;
- maintain an improvement backlog;
- prepare experiment envelopes.

Experiments occur only inside authorized windows and budgets.

Priority considers:

```text
frequency
+ impact
+ confidence
+ shared cause potential
+ investigation cost
+ regression risk
+ strategic importance
```

One low-impact anomaly may be recorded. A repeated authority failure becomes high priority.

---

## 13.27 Learning from success

Success can produce:

```text
successful run
→ observation
→ repeated pattern
→ procedure candidate
→ validation across contexts
→ Golden Path proposal
→ approved procedural memory
```

Aurora should not generalize from one lucky run.

Positive learning examples:

- more effective experiment order;
- question sequence that reduces ambiguity;
- provider that performs better for a data class;
- stable firmware test protocol;
- explanation format Leandro finds useful;
- recovery sequence that prevents duplicate effects.

---

## 13.28 Evaluation of memory and context

Memory evals should test:

- correct recall;
- correct non-recall;
- temporal reasoning;
- supersession;
- scope isolation;
- source provenance;
- premise awareness;
- workflow knowledge;
- stale-state refresh;
- false-memory rejection;
- token/latency cost;
- longitudinal behavior.

A benchmark score does not alone prove project safety or authority correctness.

---

## 13.29 Evaluation of personality and proactivity

Test cases:

- justified disagreement;
- user frustration;
- casual humor;
- critical safety warning;
- repeated notification suppression;
- wrong urgency classification;
- public presence disclosure;
- correction and trust repair;
- avoiding emotional manipulation;
- maintaining identity across models.

Human evaluation rubrics should distinguish preference from constitutional behavior.

---

## 13.30 Evaluation of harness orchestration

Required journeys:

- provider discovery and verification;
- delegation and evidence;
- child delegation;
- direct data channel;
- provider substitution;
- restart;
- cancellation;
- ambiguous effect;
- revoked authority;
- incompatible version;
- malicious provider behavior;
- incomplete evidence.

---

## 13.31 Evaluation of physical operation

Progressive tests:

```text
simulation
→ controller test double
→ hardware-in-the-loop
→ read-only bench
→ guided manual action
→ controlled actuation
→ fault-injection/interlock drill
→ bounded autonomous campaign
```

No physical autonomy level is accepted only from software tests.

---

## 13.32 Incident review and documentation impact

A material incident may require changes to:

- memory;
- provider trust;
- policy;
- capability spec;
- Golden Path;
- ADR;
- roadmap;
- evaluation suite;
- implementation;
- operator guidance.

Corrective code without documentation/eval update can allow recurrence.

---

## 13.33 Open research and decisions

- telemetry backend;
- domain-event store;
- evaluation framework;
- benchmark selection;
- confidence/calibration method;
- anomaly detection;
- incident correlation technique;
- causal graph representation;
- reviewer separation;
- self-improvement scheduler;
- canary routing;
- longitudinal evaluation infrastructure;
- physical safety standards.

---

## 13.34 Non-goals

This section does not define:

- one universal quality score;
- full autonomous self-rewriting;
- automatic production promotion;
- metric collection without a decision purpose;
- replacing human/domain authority with benchmark scores;
- treating all user feedback as training signal;
- proving consciousness or human-like learning;
- collecting every prompt and secret in telemetry;
- retrying until success without causal progress.

> Aurora learns continuously, but changes deliberately.
