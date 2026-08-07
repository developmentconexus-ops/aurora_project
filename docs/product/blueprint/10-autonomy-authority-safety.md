---
id: DOC-AURORA-BLUEPRINT-10
title: Autonomia, Autoridade e Segurança Operacional
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
  - autonomy levels and delegation model
  - Authority Grant and effect principles
  - autonomous mission envelope
  - guardrails, interlocks and emergency authority
  - self-improvement protected boundaries
related:
  - DOC-AURORA-BLUEPRINT-02
  - DOC-AURORA-BLUEPRINT-03
  - DOC-AURORA-BLUEPRINT-04
  - DOC-AURORA-BLUEPRINT-05
  - DOC-AURORA-BLUEPRINT-07
  - DOC-AURORA-BLUEPRINT-09
  - DOC-AURORA-BLUEPRINT-11
  - DOC-AURORA-BLUEPRINT-13
review_triggers:
  - authority or autonomy scope changes
  - new effect category
  - emergency behavior changes
  - self-improvement promotion changes
last_reviewed: 2026-08-06
---

# 10. Autonomia, Autoridade e Segurança Operacional

## 10.1 Propósito

Aurora is intended to become meaningfully autonomous, not merely a system that asks confirmation before every tool call.

Leandro explicitly wants to be able to delegate work such as:

- run many AI workflow variants overnight;
- compile and test firmware strategies;
- adapt the next experiment based on results;
- stop only when success, convergence, budget or guardrail conditions are reached;
- return evidence and recommendation without continuous human interaction.

This autonomy is valuable only if the system can distinguish:

- technical capability from permission;
- local operational choice from material product decision;
- reversible action from irreversible effect;
- model guardrail from enforceable safety boundary;
- pre-authorized emergency containment from general authority;
- experimental self-improvement from production promotion.

The constitutional model is:

> **progressive authority plus mission-scoped delegated autonomy.**

---

## 10.2 Principles

### P1 — Access is not authority

A configured tool, credential, API or device does not imply permission.

### P2 — Authority is explicit, scoped, expiring and revocable

No broad permanent administrator authority by default.

### P3 — Autonomous inside the envelope, conservative at the boundary

Aurora can adapt local strategy but cannot redefine objective, absolute safety limit, protected data or promotion rules.

### P4 — Reversibility informs approval, not truth

A reversible action may receive broader preauthorization, but it still needs audit and scope.

### P5 — Physical safety is independent of model judgment

Deterministic interlocks remain local and testable.

### P6 — Authority cannot be inherited transitively

A Harness cannot pass its grant to another provider.

### P7 — Revocation must have operational effect

Changing a database flag without blocking credentials/gateways is insufficient.

### P8 — Ambiguous effect blocks blind retry

External state must be reconciled.

### P9 — Emergency authority is narrow and pre-defined

Containment is not permission to continue the mission.

### P10 — Self-improvement does not control its own constitution

Protected boundaries require independent approval.

---

## 10.3 Authority versus autonomy

### Authority

What actions may be performed, on which resources, by which actor, under what conditions.

### Autonomy

How much decision-making freedom the actor has inside that authority.

Examples:

- Aurora may be authorized to read telemetry but not autonomously change settings.
- Aurora may be authorized to execute 100 sandbox experiments and choose them adaptively.
- Aurora may be authorized to stop a test automatically but not restart it.

Authority and autonomy are orthogonal.

---

## 10.4 Progressive autonomy levels

### N0 — Observe only

Read state, telemetry and artifacts. No external change.

### N1 — Recommend/prepare

Analyze and prepare plan/configuration/candidate without execution.

### N2 — Authorized action

Execute one specific bounded effect.

Example:

> “Run the approved test command in the sandbox.”

### N3 — Authorized workflow

Execute a known sequence with defined branching and checks.

Example:

> “Build, flash, run the fixed protocol and create the report.”

### N4 — Adaptive campaign

Receive an objective and adapt hypotheses, variants, order and local strategy.

Example:

> “Find a better workflow configuration within this budget and immutable eval.”

### N5 — Continuous program

Monitor an authorized system, identify opportunities and execute predefined classes of campaigns without individual initiation.

This is future/directional and requires mature evaluation, attention, cost and governance.

The prior dialogue called action/workflow/campaign N1/N2/N3. This expanded model adds observe/prepare while preserving the substantive progression.

---

## 10.5 Promotion between autonomy levels

A capability/provider earns broader autonomy through evidence.

Criteria may include:

- conformance;
- functional evaluations;
- error/recovery history;
- effect correctness;
- authority compliance;
- incident record;
- budget predictability;
- rollback;
- scope maturity;
- Leandro approval.

Promotion is:

```text
capability-specific
provider/build-specific
environment-specific
effect-specific
risk-specific
time-bounded
```

Success in public research does not grant autonomy over repositories or devices.

---

## 10.6 Autonomous Mission Envelope

Every material adaptive campaign requires a structured envelope.

### 10.6.1 Objective

What result is sought.

### 10.6.2 Baseline

Current state/metrics against which improvement is measured.

### 10.6.3 Mutable search space

What Aurora may change.

### 10.6.4 Protected space

What cannot be changed.

### 10.6.5 Environment

Where work may run and what isolation exists.

### 10.6.6 Authority

Allowed effects/resources/providers/credentials.

### 10.6.7 Budgets

Money, tokens, model calls, compute, time, experiments, energy, storage, device cycles and attention.

### 10.6.8 Evaluation

Metrics, datasets, holdouts, acceptance and trade-offs fixed before the campaign.

### 10.6.9 Guardrails and interlocks

Policy and deterministic boundaries.

### 10.6.10 Stop conditions

Success, convergence, no progress, budget, risk, incident or time.

### 10.6.11 Escalation

When a decision returns to Aurora/Leandro.

### 10.6.12 Evidence

Experiment ledger, artifacts, receipts, failures and recommendation.

### 10.6.13 Promotion and rollback

What may happen after a candidate exists.

---

## 10.7 Complete campaign example — AI workflow

```yaml
autonomous_mission:
  id: MIS-WORKFLOW-IMPROVEMENT-01

  objective:
    statement: improve answer accuracy without unacceptable cost/latency regression
    baseline: EVAL-BASE-017

  mutable:
    - prompt_templates
    - workflow_step_order
    - retrieval_parameters
    - approved_model_parameters
    - provider_from_approved_list

  protected:
    - production_environment
    - holdout_dataset
    - safety_filters
    - constitutional_prompts
    - scoring_rules

  environment: isolated-ai-sandbox

  authority:
    allow:
      - read_eval_dataset
      - run_approved_models
      - write_candidate_artifacts
      - execute_eval
    deny:
      - deploy_production
      - modify_holdout
      - access_unrelated_memory

  budget:
    usd_hard: 20
    model_calls_hard: 500
    experiments_hard: 100
    wall_time_hard: 8h
    warn_at_percent: 80

  evaluation:
    metrics:
      - correctness
      - hallucination_rate
      - latency_p95
      - cost_per_task
    holdout: HOLDOUT-003
    promotion: supervised

  stop:
    - accepted_target_reached
    - budget_exhausted
    - no_progress_after: 15
    - material_security_or_privacy_issue
    - environment_degraded

  escalation:
    - requires_new_data_class
    - requires_nonapproved_provider
    - objective_or_metric_change
    - production_effect
```

Aurora can choose hypotheses and experiments inside this envelope. She cannot change the holdout or promote automatically.

---

## 10.8 Campaign loop

```text
load baseline and constraints
→ analyze failure classes
→ formulate competing hypotheses
→ prioritize experiment
→ create candidate variant
→ execute/evaluate
→ record result and cost
→ eliminate regression or update belief
→ choose next experiment
→ stop/escalate
→ independent review
→ report/promotion proposal
```

Each change in strategy remains attributable. A campaign does not need human approval for every local variant.

---

## 10.9 Firmware campaign example

Objective:

> Compare five control strategies for thermal stability.

Allowed autonomous choices:

- variant order;
- parameter values inside approved range;
- repeat inconclusive runs;
- abandon unsafe/poor variant;
- allocate remaining experiment budget;
- select next discriminating test.

Fixed constraints:

- exact board revisions;
- firmware build identity;
- voltage/current/temperature limits;
- ramp and cool-down;
- flash cycle budget;
- calibration requirements;
- interlocks;
- no production/default promotion.

The model can select test order. The current limiter prevents unsafe current independently.

---

## 10.10 Authority Grant model

Conceptual fields:

```yaml
authority_grant:
  id: GRANT-LAB-021
  grantor: LEANDRO
  subject: AURORA
  actor: LAB-HARNESS@build
  delegation: DEL-LAB-024

  actions:
    - device.read_telemetry
    - instrument.configure
    - source.set_output

  resources:
    - DEV-SOURCE-01
    - DEV-ELOAD-01
    - DEV-PCB-POWER-B-01

  constraints:
    environment: LAB-BENCH-01
    current_max_a: 0.5
    voltage_max_v: 24
    temperature_max_c: 80
    valid_from: ...
    valid_until: ...

  budget:
    energy_wh: 100
    duration: 60m

  escalation:
    - limit_change
    - unknown_device
    - interlock_unavailable
```

Exact token/schema is future work.

---

## 10.11 Actor chain

A material effect preserves:

```text
subject / authority origin: Leandro
actor: Aurora
executor: Harness/Worker
presence: originating interface/device
mission/delegation: purpose
resource/action: effect
```

This allows policy and audit to distinguish:

- Leandro directly pressing an emergency stop;
- Aurora stopping under preauthorization;
- Harness configuring a device inside a Delegation;
- compromised worker attempting unrelated action.

---

## 10.12 Effect taxonomy

### E0 — Internal reasoning

No external state effect beyond compute/logging.

### E1 — Read

Read file, query status, read telemetry.

### E2 — Compute/Artifact

Generate report, simulation, candidate in sandbox.

### E3 — Local reversible write

Modify temporary file/branch/config inside isolated environment.

### E4 — Network/external data

Send request or context outside local boundary.

### E5 — Credential-mediated operation

Use secret/token for repository/service.

### E6 — External communication

Email/message/publication.

### E7 — Repository/infrastructure

Commit, PR, merge, deployment, database migration.

### E8 — Financial/legal

Purchase, payment, agreement or material business action.

### E9 — Device/physical actuation

Power, motion, thermal/mechanical/electrical effects.

### E10 — Emergency containment

Preauthorized narrow stop/isolation/revoke action.

Risk is contextual. A repository write in a disposable branch differs from production merge.

---

## 10.13 Reversibility and materiality

Properties:

- reversible automatically;
- reversible with cost;
- partially reversible;
- irreversible;
- safety-critical;
- externally visible;
- legally/financially material;
- affects third party;
- changes authority/security;
- changes production.

Approval policy considers consequence, not only technical effect type.

Example:

- writing a sandbox prompt candidate: low materiality;
- deleting canonical memory: high personal impact despite local write;
- publishing an email: external irreversible communication;
- energizing low-voltage board: physical but bounded;
- disabling interlock: prohibited/critical.

---

## 10.14 Policy and enforcement stack

```text
Constitution and Product Policy
        ↓
Capability/Environment Standards
        ↓
Mission and Delegation Contract
        ↓
Authority Grant
        ↓
Policy Decision Point
        ↓
Effect Gateway / Credential Broker
        ↓
Sandbox / OS / Device Controller
        ↓
Target
        ↓
Receipt, Observation and Audit
```

### Constitution

Defines non-negotiable boundaries.

### Contract/Grant

Defines current permitted scope.

### PDP

Evaluates request context.

### Gateway

Applies allow/deny/modify/confirm at the effect.

### Sandbox/device controller

Prevents bypass and contains process.

### Receipt

Records observable result.

AHDK and prompts help form requests; they are not the enforcement boundary.

---

## 10.15 Policy decision outcomes

Possible decisions:

```text
ALLOW
DENY
REQUIRE_CONFIRMATION
ALLOW_WITH_MODIFICATIONS
DEFER
EMERGENCY_ALLOW
```

Examples:

- allow read telemetry;
- modify requested current from 0.6 A to granted max 0.5 A only if contract permits clamping and reports it;
- require confirmation for merge;
- deny public provider confidential context;
- emergency allow stop relay.

Modification must never silently change the meaning of the requested experiment.

---

## 10.16 Credential brokering

Harness should receive references or short-lived scoped credentials.

```text
provider effect request
→ PDP allows
→ Broker exchanges/creates scoped credential
→ Gateway uses or passes securely
→ credential expires/revokes
→ receipt stores reference, not secret
```

A parent Harness cannot copy a secret to a child.

Credential scope can bind:

- actor;
- action;
- resource;
- time;
- environment;
- delegation;
- nonce/idempotency.

---

## 10.17 Budget model

Budgets are deterministic constraints with units and reconciliation.

Possible dimensions:

- currency;
- tokens/model calls;
- compute seconds;
- wall time;
- experiment count;
- storage;
- network bytes;
- energy;
- device/flash cycles;
- instrument occupancy;
- attention interruptions.

### Hard versus soft

Soft threshold generates warning/replan opportunity.

Hard threshold blocks new consumption or stops according to contract.

### Accounting ambiguity

If provider and Aurora disagree on consumed budget:

- new effects may pause;
- reconcile receipts/provider records;
- do not assume the lower value.

---

## 10.18 Guardrails

Guardrails include:

- allowed search/mutable space;
- prohibited resources;
- data policy;
- model/provider allowlist;
- output/evidence requirements;
- parameter ranges;
- maximum concurrency;
- retry rules;
- promotion rules.

Guardrails can be:

- constitutional;
- policy;
- contract;
- runtime;
- physical.

A prompt instruction is the weakest form and should not carry critical enforcement alone.

---

## 10.19 Deterministic interlocks

Physical/digital critical limits require independent mechanisms.

Examples:

- hardware current limit;
- over-temperature cutoff;
- watchdog;
- physical emergency stop;
- gateway maximum range;
- filesystem sandbox;
- network egress firewall;
- credential scope;
- deployment environment protection;
- database transaction/constraint.

Properties:

- testable;
- state observable;
- fail-safe where practical;
- cannot be disabled by ordinary campaign authority;
- triggers incident/receipt;
- survives model/cloud failure.

---

## 10.20 Stop conditions

A campaign stops when:

- success criterion reached;
- budget exhausted;
- deadline;
- convergence;
- no progress after defined cycles;
- all hypotheses refuted/insufficient;
- required capability unavailable;
- guardrail/interlock trigger;
- evidence invalid;
- environment degraded;
- authority revoked;
- incident;
- Leandro cancels.

Stop means no new work/effects. Cleanup and reconciliation may continue under separate narrow authority.

---

## 10.21 Escalation conditions

Escalate when:

- objective/scope must change;
- new data class/provider/effect required;
- budget extension;
- material architecture choice;
- irreversible effect;
- safety boundary issue;
- conflicting evidence with no safe default;
- evaluation metric ambiguity;
- legal/financial/third-party impact;
- protected self-improvement area;
- no viable hypothesis remains.

Decision Request should include recommendation and consequence of waiting.

---

## 10.22 Revocation

Revocation can target:

- grant;
- provider/build;
- presence/device;
- credential;
- data channel;
- effect category;
- mission/delegation;
- environment.

Operational response:

```text
block new effect at gateway
→ invalidate token/credential
→ notify provider/cancel
→ close/revoke channels
→ enter safe state
→ preserve receipts/evidence
→ reconcile effects already started
→ update Registry/mission state
→ inform Leandro according to severity
```

### Partition case

If provider is disconnected, local authority must expire and gateway/device controls must enforce expiry. A notification delivered later is not sufficient.

---

## 10.23 Ambiguous effects

Example:

A purchase API or device command times out after request was sent.

State:

```text
REQUESTED
→ SENT
→ RESULT_UNKNOWN
```

Aurora must:

- use idempotency/external reference;
- query target or receipt store;
- block duplicate retry;
- escalate if unreconcilable;
- preserve ambiguity in outcome.

“Try again” is not a safe generic recovery.

---

## 10.24 Emergency authority

Emergency authority is preauthorized for narrow containment.

Examples:

- disable source output on overcurrent;
- stop motor on safety sensor;
- cancel campaign on cost runaway;
- revoke compromised provider;
- block data exfiltration;
- isolate environment.

Properties:

- smallest action necessary;
- independent trigger where possible;
- no mission continuation;
- immediate receipt/notification;
- tested through drills;
- manual override/emergency stop available;
- post-incident review.

Aurora cannot label normal optimization as emergency to bypass approval.

---

## 10.25 Human confirmation design

Confirmation should state:

- exact action;
- target;
- material consequence;
- reversible/irreversible;
- authority scope;
- current safeguards;
- alternatives;
- timeout/default.

Bad:

> “Confirmar?”

Better:

> “Confirmar o merge do commit `abc123` na branch `main` e iniciar o deploy de produção? Esta ação altera o ambiente externo; rollback está disponível pela release anterior.”

Voice confirmation for critical values may require read-back and a deterministic control.

---

## 10.26 Self-improvement authority

Aurora may autonomously within an approved experiment environment:

- collect incidents;
- formulate causal hypotheses;
- modify non-constitutional prompts/workflows/retrieval;
- build candidates;
- run evaluation;
- compare results;
- prepare promotion.

She may not autonomously promote material changes to:

- constitutional identity;
- Leandro's authority;
- guardrails/interlocks;
- security/data policy;
- promotion rules;
- audit;
- revocation/shutdown;
- ability to grant authority;
- production memory semantics;
- protected holdout/scoring.

She can propose changes with research and impact analysis.

---

## 10.27 Protected area enforcement

Protected areas should be enforced through:

- separate repository/path ownership;
- policy;
- signing/approval;
- environment restrictions;
- deployment gate;
- immutable evaluation references;
- independent review;
- audit.

A prompt saying “do not change constitution” is not sufficient.

---

## 10.28 Safety cases by domain

### Research

Main risks: data leakage, bad sources, misinformation, cost.

### Software

Repository modification, credential use, deployment, data migration.

### Firmware

Target compatibility, bricking, cycle wear, safety behavior.

### Laboratory

Electrical/thermal/mechanical danger, wrong device, calibration, interlock.

### Personal operations

Privacy, communication, financial/legal effects.

Each Capability Spec defines domain-specific authority and evidence.

---

## 10.29 Authority incidents

Examples:

- provider attempted undeclared network call;
- expired grant used;
- child provider received parent context;
- budget exceeded;
- device command outside range;
- action occurred after cancellation;
- emergency stop unavailable;
- confirmation did not match actual target;
- model classified irreversible action as reversible.

Incident response includes containment, evidence, trust downgrade and causal review.

---

## 10.30 Evaluation requirements

Future implementation must prove:

1. technical access without grant is denied;
2. provider receives only scoped actions/resources;
3. child Delegation cannot inherit parent token;
4. soft/hard budgets warn/block correctly;
5. campaign adapts experiments without repeated confirmation;
6. campaign cannot alter protected evaluation/production;
7. stop condition prevents new work;
8. revocation blocks new effects during provider partition;
9. ambiguous effect does not duplicate;
10. emergency containment works without model/cloud;
11. interlock prevents command beyond absolute range;
12. confirmation displays exact target/consequence;
13. self-improvement candidate cannot promote itself;
14. protected constitutional path is enforced;
15. all material effects produce receipt and actor chain;
16. a low-risk provider can be promoted without granting unrelated scope;
17. authority can be inspected and revoked by Leandro.

---

## 10.31 Open technical decisions

- policy engine;
- Authority Grant schema;
- token/delegation mechanism;
- Credential Broker;
- effect gateway architecture;
- sandbox/environment profiles;
- budget ledger;
- risk classification;
- confirmation UX;
- emergency signaling;
- device interlock implementation;
- revocation under offline operation;
- constitutional deployment gate.

---

## 10.32 Non-goals

- confirm every action;
- broad permanent root access;
- autonomous production promotion;
- peer-to-peer authority propagation;
- prompt-only safety;
- one scalar risk score;
- model-controlled interlock;
- unattended hazardous physical operation in current horizon;
- emergency authority used for normal progress;
- silent budget extension;
- changing evaluation criteria during campaign;
- treating access to a tool as consent.
