---
id: DOC-AURORA-BLUEPRINT-04
title: Ciclo Cognitivo e Jornadas Globais
document_type: product_blueprint_section
form: explanation
authority: constitutional
status: proposed
version: 0.2.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - global cognitive lifecycle
  - intent-to-outcome journeys
  - interruption and continuation model
  - human decision placement
related:
  - DOC-AURORA-BLUEPRINT-01
  - DOC-AURORA-BLUEPRINT-02
  - DOC-AURORA-BLUEPRINT-03
  - DOC-AURORA-BLUEPRINT-06
  - DOC-AURORA-BLUEPRINT-07
  - DOC-AURORA-BLUEPRINT-10
review_triggers:
  - cognitive lifecycle changes
  - global mission flow changes
  - decision or interruption boundary changes
last_reviewed: 2026-08-05
---

# 4. Ciclo Cognitivo e Jornadas Globais

## 4.1 Propósito

Aurora não pode ser definida apenas por componentes. A experiência depende de um ciclo que transforma percepção e intenção em resultado observado, verificável e lembrado.

O ciclo global é:

```text
PERCEBER
→ COMPREENDER
→ RECORDAR
→ FORMAR INTENÇÃO
→ PLANEJAR
→ SELECIONAR CAPACIDADES
→ AGIR
→ OBSERVAR
→ VERIFICAR
→ REGISTRAR
→ APRENDER
```

Essas etapas não significam uma pipeline rígida de onze chamadas. São responsabilidades que podem:

- ocorrer em loops;
- ser executadas por módulos diferentes;
- retornar a etapas anteriores;
- exigir decisão de Leandro;
- continuar durante horas;
- sobreviver a restart;
- produzir child delegations;
- operar com diferentes modelos e runtimes.

O objetivo desta seção é preservar o significado global para que frameworks internos não ditem o comportamento do produto.

---

## 4.2 Invariantes do ciclo

### 4.2.1 Toda ação possui intenção e contexto atribuíveis

Aurora deve conseguir responder:

- qual objetivo motivou a ação;
- qual projeto e missão estavam ativos;
- quais fontes e memórias foram usadas;
- qual authority permitiu;
- qual actor executou;
- qual resultado foi observado.

### 4.2.2 Percepção não é interpretação

Um sensor, mensagem ou tool result produz input. Aurora ainda precisa classificar:

- source;
- freshness;
- uncertainty;
- relevance;
- authority;
- relationship with current intent.

### 4.2.3 Planejamento não cria authority

Um plano pode recomendar uma ação sem torná-la permitida.

### 4.2.4 Execução não prova sucesso

Tool exit code, model claim ou completed state não substitui acceptance e evidence.

### 4.2.5 Observação deve fechar o loop

Efeitos importantes precisam gerar retorno observável. Quando a observação é impossível ou ambígua, a missão deve refletir essa incerteza.

### 4.2.6 Registro não significa promoção automática

Um evento pode ser preservado sem virar memória global, decisão ou Golden Path.

### 4.2.7 Aprendizado não altera automaticamente o sistema principal

Patterns e candidates seguem avaliação e promotion governance.

---

## 4.3 Perceber

Aurora recebe sinais de:

- Leandro por texto, voz, imagem ou gesto;
- conversations and interfaces;
- project repositories and documents;
- harness events;
- tools and services;
- device telemetry;
- cameras and microphones when authorized;
- schedules and timers;
- policy and security systems;
- environment health;
- self-observation and evaluation.

### 4.3.1 Perception Event

Um evento material deve preservar:

```text
source identity
channel
observed_at
received_at
payload reference
schema/version
sensitivity
integrity
confidence/quality
related project/environment
```

### 4.3.2 Sensor activation

O fato de uma Presence possuir câmera ou microfone não implica que Perceive possa utilizá-los.

Ativação requer:

- mode;
- purpose;
- policy;
- indicator;
- retention;
- effective authority.

### 4.3.3 Perception failure

Exemplos:

- stale telemetry;
- dropped event;
- ambiguous voice command;
- wrong device identity;
- tool result without schema;
- camera observation with insufficient quality;
- delayed event after cancellation.

Perception failures must not be silently converted into confident world-state updates.

---

## 4.4 Compreender

Understanding transforms signals into a situated interpretation.

Questions include:

- what does Leandro want?
- which project is active?
- is this a question, decision, hypothesis, command or casual remark?
- what terms are project-specific?
- is there ambiguity material to outcome or risk?
- which constraints already exist?
- is the request compatible with authority and current status?

### 4.4.1 Intent classes

Initial conceptual classes:

```text
ASK
EXPLORE
DECIDE
PLAN
DELEGATE
EXECUTE
OBSERVE
REVIEW
CORRECT
CANCEL
AUTHORIZE
REVOKE
REMEMBER
FORGET
```

These classes guide behavior but do not replace natural conversation.

### 4.4.2 Ambiguity handling

Aurora should not ask a question merely because any ambiguity exists. She evaluates:

```text
materiality
+ reversibility
+ available context
+ cost of wrong assumption
+ ability to proceed safely
```

Low-risk ambiguity may be resolved by a stated assumption. Material ambiguity triggers a decision request.

### 4.4.3 Understanding as a revisable hypothesis

Aurora may represent:

```text
interpreted intent
confidence
supporting cues
alternative interpretation
consequence if wrong
```

Execution feedback can invalidate the initial interpretation.

---

## 4.5 Recordar e construir contexto

Aurora does not “remember everything” into the prompt. She builds a context for the current decision.

Pipeline:

```text
intent and situation
→ determine scopes
→ load current authority snapshot
→ load project state
→ retrieve relevant memory/knowledge
→ verify freshness
→ resolve conflicts
→ minimize/redact
→ compile Context Pack
```

Precedence for action:

```text
live safety state
→ active authority and policy
→ current project/mission state
→ accepted decisions and contracts
→ verified evidence
→ relevant memory
→ historical conversation
→ external suggestions
```

The exact precedence varies by question. Historical questions may prefer the original transcript as evidence of what was said, without allowing it to govern current action.

---

## 4.6 Formar intenção operacional

User intent often needs transformation before execution.

Example:

> “Aurora, deixa esse workflow melhor durante a noite.”

Operational intent must resolve:

- definition of better;
- baseline;
- allowed mutable variables;
- evaluation dataset;
- budget;
- environment;
- protected components;
- stop and escalation conditions;
- promotion policy.

Aurora may prepare an Autonomous Mission Envelope and ask Leandro only about material missing decisions.

### 4.6.1 Intent preservation

The original intent remains recorded separately from derived plans. A plan cannot silently narrow or broaden the goal.

### 4.6.2 Intent drift detection

During long campaigns, Aurora compares proposed work against:

- original goal;
- accepted scope;
- non-goals;
- current evidence;
- remaining budget.

Material drift triggers replan or escalation.

---

## 4.7 Planejar

Planning establishes a proposed route from intent to evidence.

It may include:

- decomposition;
- hypotheses;
- capability needs;
- dependencies;
- order and parallelism;
- risks;
- decision gates;
- budget allocation;
- evidence plan;
- recovery;
- stop conditions.

### 4.7.1 Global versus local planning

Aurora plans global composition.

Harnesses plan specialized local execution.

```text
Aurora:
  research candidate approaches
  then simulate selected candidates
  then run firmware/lab validation

Hardware Harness:
  chooses topologies, calculations and simulator method
```

### 4.7.2 Plan quality

A plan is not good because it is long. It is good when it:

- preserves intent;
- exposes material choices;
- minimizes unnecessary work;
- identifies authority needs;
- establishes evidence;
- contains recovery;
- makes the next safe action clear.

### 4.7.3 Replanning

Triggers include:

- failed assumption;
- missing capability;
- new risk;
- budget variance;
- contradictory evidence;
- provider unavailable;
- framework/protocol limitation;
- decision by Leandro;
- incident;
- objective change.

Replan preserves previous plan history and reasons.

---

## 4.8 Selecionar capacidades

Aurora resolves what outcome is needed before choosing a Provider.

Flow:

```text
required outcome
→ capability identity
→ compatible providers
→ effective trust/authority
→ environment/data constraints
→ cost/latency/recovery fit
→ selected provider
→ fallback path
```

Selection must remain explainable.

Example:

> “Usei a Research Harness local porque o Context Pack contém informação confidencial. O provider cloud possui melhor benchmark, mas não está aprovado para essa classe de dados.”

When no provider is approved, Aurora may:

- prepare verification;
- request authority;
- use a lower-capability safe provider;
- ask Leandro;
- block.

She may not improvise an unregistered provider for a material effect.

---

## 4.9 Agir

Action may occur through:

- direct deterministic tool;
- local service;
- harness delegation;
- workflow;
- effect gateway;
- device controller;
- Leandro performing a requested manual action.

### 4.9.1 Action categories

```text
internal reasoning
read
compute
prepare
write local
network
credential use
repository change
external communication
deploy
financial
physical control
emergency stop
```

Each category has different authority and observation requirements.

### 4.9.2 Intent–Action–Observation

Material action records:

```text
Intent
→ why this action is needed

Action
→ exact request, actor, target and authority

Observation
→ what the system observed afterward
```

This model prevents an action request from being confused with its external result.

---

## 4.10 Observar

After acting, Aurora gathers relevant outcomes.

Examples:

- command exit code;
- repository diff;
- test report;
- model evaluation;
- effect receipt;
- instrument reading;
- physical interlock state;
- provider artifact;
- cost and latency;
- user correction.

### 4.10.1 Observation window

Some effects are immediate. Others require a monitoring window.

Example:

```text
firmware flashed
→ device boots
→ telemetry stable for 60 minutes
→ thermal plateau reached
→ no reset/error events
```

Completion cannot be declared before the required window.

### 4.10.2 Ambiguous observation

If a network request times out after an external purchase or device command, Aurora must not assume success or failure. The state becomes ambiguous until reconciled by idempotency reference, target query or human check.

---

## 4.11 Verificar

Verification evaluates whether observations and artifacts satisfy criteria.

Layers may include:

- deterministic checks;
- schema validation;
- unit/integration tests;
- independent model review;
- user journey;
- benchmark/eval;
- instrument measurement;
- safety drill;
- Leandro decision.

Distinctions:

```text
executor claim
≠ receipt
≠ evidence
≠ verdict
≠ global outcome
```

The required independence is proportional to risk.

---

## 4.12 Registrar

Aurora records different outputs in different stores or authority classes.

```text
operational state
→ runtime state store

large output
→ Artifact Store

verification
→ Evidence Store

durable product choice
→ ADR/spec/contract

conversation continuity
→ session and observational memory

project update
→ project state

historical narrative
→ history/worklog
```

A log line is not automatically a project decision. A memory summary is not automatically evidence.

---

## 4.13 Aprender

Learning identifies reusable knowledge without mutating Aurora impulsively.

Possible results:

- memory update;
- new project relation;
- procedure candidate;
- Golden Path candidate;
- failure correlation;
- evaluation case;
- new guardrail proposal;
- provider trust adjustment;
- improvement candidate;
- research question.

Promotion depends on scope, evidence and risk.

---

# 4.14 Global journey — resume a laboratory project

## Trigger

Leandro enters the laboratory:

> “Aurora, vamos continuar a fonte programável.”

## Flow

```text
1. Presence authenticates Leandro and identifies laboratory environment.
2. Understand resolves project PRJ-POWER-SUPPLY.
3. Context Builder loads current project snapshot and accepted decisions.
4. Live registry verifies connected devices and current firmware.
5. Aurora detects last incomplete experiment and safety prerequisites.
6. Aurora explains state, uncertainty and next proposed test.
7. Leandro authorizes the test envelope.
8. Laboratory Harness configures instruments inside deterministic limits.
9. Telemetry streams to evaluation.
10. Aurora monitors thresholds and progress.
11. Evidence is linked to hypothesis and acceptance criteria.
12. Outcome updates the project; next action is proposed.
```

## Failure cases

- wrong board detected;
- firmware version differs from memory;
- calibration expired;
- source current limit inconsistent;
- missing authority;
- telemetry stale;
- interlock unavailable;
- experiment history incomplete.

Aurora blocks or degrades based on materiality.

---

# 4.15 Global journey — overnight AI workflow campaign

## Request

> “Aurora, durante a noite tente melhorar a precisão deste workflow. Pode mudar prompts, ordem e parâmetros; use o eval aprovado; limite o custo a US$ 20; não altere produção.”

## Mission construction

```yaml
objective: improve workflow accuracy
baseline: evaluation run BASE-017
mutable_space:
  - prompts
  - step_order
  - authorized_model_parameters
environment: isolated-sandbox
budget:
  usd: 20
  max_runs: 100
protected:
  - production
  - evaluation_holdout
  - safety_filters
promotion:
  automatic: false
stop:
  - budget_exhausted
  - convergence
  - no_progress_after_n_cycles
  - material_risk
```

## Adaptive loop

```text
baseline analysis
→ failure clustering
→ competing hypotheses
→ variants
→ evaluation
→ eliminate regressions
→ repeat inconclusive results
→ change search direction
→ compare Pareto trade-offs
→ stop
→ independent review
→ report
```

## Required outcome

Aurora returns:

- baseline;
- experiment ledger;
- variants;
- cost;
- metrics;
- regressions;
- best candidate;
- uncertainty;
- promotion recommendation;
- reproducible artifacts.

She does not silently update production.

---

# 4.16 Global journey — firmware optimization campaign

## Request

> “Teste cinco estratégias de controle e determine qual oferece melhor estabilidade térmica.”

## Flow

```text
Firmware Harness builds signed variants
→ Device Gateway verifies target identity
→ flash under cycle budget
→ reset and health check
→ Laboratory Harness executes protocol
→ deterministic interlocks enforce current/temperature/time
→ Evaluation compares stability, ripple, latency and thermal behavior
→ inconclusive variants repeat within budget
→ unsafe variant is quarantined
→ artifacts and telemetry remain attributable to firmware build
```

## Safety distinction

Aurora may adapt test order and parameters inside the approved safe space. She cannot increase the absolute voltage/current/temperature boundary to finish the mission.

---

# 4.17 Global journey — multi-presence handoff

## Starting state

Leandro is viewing confidential project details on the laboratory computer.

## Handoff

> “Aurora, venha comigo.”

```text
Glasses presence authenticates device and Leandro
→ Core identifies active activity
→ environment classified as potentially public
→ Context Builder creates safe presence pack
→ Aurora resumes with a minimal summary
→ sensitive detail waits for private channel/re-auth
→ laboratory campaign continues independently
→ critical-only alerts selected
```

The handoff preserves intention and mission state without copying the complete project or credentials to the glasses.

---

# 4.18 Global journey — multi-harness collaboration

Mission:

> “Desenvolva e valide um novo módulo de controle.”

```text
Aurora creates global mission
├── Research Delegation
│   └── returns candidates and source evidence
├── Hardware Delegation
│   └── requests thermal capability
│       └── Aurora creates child delegation
├── Firmware Delegation
├── Laboratory Delegation
└── Evaluation Delegation
```

Aurora owns global dependencies and acceptance. Each harness owns local method.

A direct telemetry channel may connect Laboratory and Evaluation after Aurora authorizes:

- schema;
- endpoints;
- duration;
- data classification;
- rate/budget;
- retention;
- revocation.

---

# 4.19 Global journey — self-improvement

Incidents indicate that Aurora repeatedly ignores current project authority.

```text
incidents collected
→ symptoms clustered
→ possible shared Context Builder cause
→ reproduction case
→ competing hypotheses
→ candidate retrieval policy
→ original/neighbor/contrary/history/unseen/adversarial evals
→ independent reviewer
→ shadow mode
→ canary project scope
→ promotion proposal
→ monitoring and rollback proof
```

The conversation that identified the failure remains evidence. The candidate cannot rewrite the constitutional precedence rule to make its test easier.

---

## 4.20 Interruptions and attention

Interruptions are governed events.

### Categories

```text
INFORMATIONAL
DECISION_REQUIRED
AUTHORITY_REQUIRED
BUDGET_THRESHOLD
RISK_WARNING
INCIDENT
EMERGENCY
```

### Delivery policy

Aurora evaluates:

- severity;
- urgency;
- confidence;
- current presence;
- privacy;
- Leandro's focus state;
- available alternative such as summary;
- consequence of delay.

A notification must include why now and what happens if ignored.

---

## 4.21 Concurrency

Aurora may coordinate multiple active missions but must preserve:

- project isolation;
- budget attribution;
- authority separation;
- attention limits;
- device resource conflicts;
- instrument leases;
- priority and emergency preemption;
- deterministic reconciliation.

Starting more work is not progress if it increases confusion or resource collision.

---

## 4.22 Session and process replacement

A session is a presentation/reasoning context, not the mission.

After restart, Aurora reconstructs from:

```text
current authority
project state
active missions and delegations
checkpoints
recent material events
relevant memory
exact artifacts/evidence references
```

The fresh process may use a different model and must still preserve:

- IDs;
- authority;
- approved contracts;
- state;
- next safe action.

---

## 4.23 Failure taxonomy

### Interpretation failure

Wrong project, intent or scope.

### Context failure

Missing, stale, irrelevant or cross-contaminated context.

### Planning failure

Unnecessary work, missing dependency, no evidence plan.

### Selection failure

Wrong provider or unsupported capability.

### Authority failure

Action allowed without grant or valid action denied unexpectedly.

### Execution failure

Provider/tool/device cannot complete.

### Observation failure

Result unavailable, ambiguous or low quality.

### Verification failure

Claim lacks sufficient proof.

### Recording failure

State or evidence not durably persisted.

### Learning failure

Narrow patch, false generalization or ungoverned promotion.

Each class informs recovery rather than generic retry.

---

## 4.24 Recovery behavior

```text
interpretation error
→ restate intent and replan

stale context
→ refresh authoritative/live source

provider failure
→ resume, substitute or block based on contract

ambiguous effect
→ reconcile through receipt/target; no blind retry

observation loss
→ repeat measurement if safe and valid

budget exhaustion
→ stop or request extension

process restart
→ recover from durable state/checkpoints

security incident
→ contain, revoke, preserve evidence and escalate
```

---

## 4.25 Evaluation requirements

Future implementations must test complete journeys, not only individual APIs.

Required proof categories:

- correct intent resolution;
- authoritative context selection;
- project isolation;
- provider selection rationale;
- authority denial and approval;
- durable restart;
- child delegation;
- direct data channel revocation;
- observation and verification;
- safe presence handoff;
- campaign budget and stop;
- physical interlock;
- self-improvement holdout and rollback.

---

## 4.26 Non-goals

This section does not define:

- one universal planner;
- one model for every stage;
- a fixed workflow graph;
- implementation queues;
- exact state-machine schemas;
- always-on autonomous operation;
- peer-to-peer federation without Aurora;
- automatic promotion of all records to memory;
- exact voice or multimodal stack.
