---
id: DOC-AURORA-BLUEPRINT-03
title: Modelo de Domínio e de Mundo
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
  - canonical domain vocabulary
  - durable entity boundaries
  - world-model principles
  - identity and relationship semantics
related:
  - DOC-AURORA-BLUEPRINT-01
  - DOC-AURORA-BLUEPRINT-04
  - DOC-AURORA-BLUEPRINT-05
  - DOC-AURORA-BLUEPRINT-06
  - DOC-AURORA-BLUEPRINT-07
  - DOC-AURORA-BLUEPRINT-10
review_triggers:
  - canonical entity added, removed or materially redefined
  - source-of-truth ownership changes
  - project, memory, presence or delegation hierarchy changes
last_reviewed: 2026-08-05
---

# 3. Modelo de Domínio e de Mundo

## 3.1 Propósito

Aurora precisa representar uma realidade que atravessa:

- Leandro e sua identidade;
- projetos de software, hardware e firmware;
- conversas e decisões;
- documentos e conhecimento externo;
- agentes, tools, workflows e harnesses;
- computadores, wearables, sensores e instrumentos;
- objetivos, missões, experimentos e evidências;
- autoridade, risco e efeitos reais;
- tempo, mudança, incerteza e memória.

Sem um modelo canônico, cada sessão ou framework atribuiria significados diferentes aos mesmos termos. Um runtime poderia chamar tudo de `task`; outro poderia tratar uma mensagem como estado; uma memória poderia ser confundida com decisão; um processo poderia ser confundido com uma harness; um dispositivo acessível poderia ser interpretado como autorizado.

O modelo de domínio existe para garantir que:

- identidade seja estável mesmo quando processos, modelos e dispositivos mudam;
- fatos, hipóteses, decisões e observações não sejam misturados;
- cada conceito possua um owner e uma fonte de verdade;
- relações sejam temporais e rastreáveis;
- o Context Builder possa selecionar informação por escopo e autoridade;
- delegações possam atravessar runtimes sem perder significado;
- efeitos digitais e físicos permaneçam atribuíveis;
- a arquitetura evolua sem reescrever a semântica do produto.

Este documento descreve o **produto-alvo conceitual**. Não afirma que todas as entidades já possuem schema ou implementação.

---

## 3.2 Princípios de modelagem

### 3.2.1 Um conceito, uma responsabilidade

Aurora não utilizará um objeto genérico chamado `Task`, `Memory` ou `Agent` para representar entidades com autoridades e lifecycles diferentes.

Exemplos:

- `Goal` expressa um resultado desejado;
- `Mission` organiza trabalho global;
- `Delegation` atribui trabalho especializado;
- `Experiment` compara uma hipótese com observações;
- `Effect Request` solicita uma mudança externa;
- `Artifact` é um produto materializado;
- `Evidence` sustenta um critério;
- `Memory Item` preserva contexto futuro.

Eles podem estar relacionados. Não são intercambiáveis.

### 3.2.2 Identidade não é instância nem estado

A identidade de um projeto, missão, dispositivo ou harness não muda automaticamente quando:

- uma sessão reinicia;
- um modelo é substituído;
- um processo morre;
- uma conexão cai;
- uma nova tentativa inicia;
- um firmware é atualizado;
- uma presença muda do computador para os óculos.

A instância concreta e o estado atual possuem identidades próprias.

### 3.2.3 Estado atual não é inferido de narrativa

Aurora não conclui que uma missão terminou porque uma harness escreveu “finalizado”, que uma placa está segura porque o último log parecia normal ou que uma preferência é global porque apareceu em três conversas.

Estado autoritativo exige:

- owner conhecido;
- transição válida;
- timestamp;
- ator;
- referência à evidência ou fonte aplicável.

### 3.2.4 O modelo de mundo é uma representação, não a realidade

O World Model contém descrições, relações, observações e inferências sobre o mundo.

Ele pode estar:

- incompleto;
- stale;
- contraditório;
- baseado em sensor impreciso;
- afetado por uma inferência errada;
- desatualizado em relação ao dispositivo real.

Aurora deve saber quando consultar uma fonte ao vivo, solicitar confirmação ou declarar incerteza.

### 3.2.5 Tempo é parte do significado

A relação:

```text
PCB-REV-B → uses firmware FW-014
```

pode ser verdadeira somente durante uma janela.

Entidades e relações materiais precisam suportar:

- `observed_at`;
- `valid_from`;
- `valid_until`;
- `recorded_at`;
- `superseded_by`;
- versão ou revision;
- freshness requirement.

### 3.2.6 Autoridade e confiança são dimensões diferentes

Uma fonte pode ser altamente confiável e não possuir autoridade para uma decisão.

Exemplo:

- um paper pode ser tecnicamente confiável;
- ele não decide qual arquitetura o Aurora adotou;
- um ADR aceito possui autoridade decisória;
- ele pode posteriormente ser superseded por evidência nova e nova decisão.

### 3.2.7 Relação não implica permissão

Conhecer que um dispositivo pertence ao laboratório não autoriza controlá-lo.

```text
Device relationship
≠ trust
≠ authority
≠ active grant
≠ executed effect
```

### 3.2.8 Eventos registram fatos; projeções apresentam estado

Eventos ajudam auditoria e reconstrução. Projeções estruturadas respondem “qual é o estado atual?”. Nenhum dos dois deve ser substituído por resumo conversacional.

---

## 3.3 Visão geral do domínio

```text
Leandro
│
├── Personal Identity and Preferences
├── Projects
│   ├── Goals
│   ├── Decisions
│   ├── Hypotheses
│   ├── Missions
│   │   ├── Delegations
│   │   ├── Budgets
│   │   ├── Authority Grants
│   │   ├── Artifacts
│   │   └── Evidence
│   ├── Experiments
│   │   ├── Variants
│   │   ├── Observations
│   │   └── Results
│   ├── Devices and Environments
│   └── Knowledge Sources and Memories
│
Aurora Identity
├── Constitutional Profile
├── Interaction Profile
├── Capability Registry
├── Presence Fabric
├── Context Builder
├── Global Mission Control
├── Authority and Effect Plane
└── Failure Intelligence

Capability Ecosystem
├── Capability
├── Provider
│   ├── Harness
│   ├── Service
│   └── Device Controller
├── Provider Instance
├── Manifest
├── Verification
├── Approval
└── Delegation

Physical and Digital World
├── Environment
├── Presence
├── Device
├── Sensor
├── Actuator
├── Instrument
├── Software System
├── Repository
└── External Service
```

---

## 3.4 Person

`Person` representa um ser humano relevante para o Aurora.

No horizonte atual, Leandro é:

```text
Primary User
Operator
Domain Authority
Data Subject
Authority Grantor
```

Esses papéis não devem ser reduzidos a um único campo `user_id`, porque possuem responsabilidades diferentes.

### 3.4.1 Leandro Profile

O perfil pode relacionar:

- identidade;
- competências;
- objetivos de longo prazo;
- preferências estáveis;
- projetos;
- dispositivos confiáveis;
- políticas pessoais;
- horários e atenção;
- decisões explícitas;
- memória global.

Inferências sobre Leandro permanecem distintas de fatos ou preferências confirmadas.

### 3.4.2 Third Party

Pessoas que aparecem em conversas, imagens, áudio, documentos ou projetos não se tornam automaticamente usuários do Aurora.

Dados de terceiros exigem:

- finalidade;
- classificação;
- retenção;
- acesso;
- políticas de gravação e divulgação.

---

## 3.5 Aurora Identity

`Aurora Identity` representa a continuidade do sistema através de modelos, versões, máquinas e presenças.

Inclui conceitualmente:

- nome;
- missão;
- princípios constitucionais;
- relationship contract;
- personality profile;
- current product version;
- capability boundaries;
- protected areas;
- allowed evolution paths.

Não inclui como identidade:

- a personalidade específica de um modelo;
- um transcript;
- um processo;
- um hostname;
- uma chave de API;
- uma única base de dados.

Uma mudança de modelo não cria outra Aurora. Uma mudança constitucional material exige versionamento e aprovação explícita.

---

## 3.6 Presence

`Presence` é a manifestação contextual do Aurora em uma interface ou dispositivo.

Exemplos:

```text
PRS-LAB-DESKTOP-01
PRS-MOBILE-01
PRS-GLASSES-01
PRS-LAB-DISPLAY-01
```

Uma Presence declara:

- device identity;
- authenticated person;
- input/output capabilities;
- environment classification;
- privacy mode;
- local storage limits;
- active session;
- context exposure policy;
- effective authority;
- connectivity state.

Presence não é uma cópia da Aurora Identity.

---

## 3.7 Environment

`Environment` representa o espaço operacional no qual uma ação, presença ou dispositivo existe.

Exemplos:

- laboratório privado;
- computador WSL2;
- sandbox de avaliação;
- rede doméstica;
- ambiente de produção;
- ambiente público usando óculos;
- bancada de alta tensão;
- dispositivo desconectado.

Atributos relevantes:

- trust zone;
- physical safety class;
- network policy;
- data policy;
- available effect gateways;
- emergency mechanisms;
- observability;
- owner;
- current health.

A mesma capability pode possuir authority diferente em ambientes diferentes.

---

## 3.8 Project

`Project` é o principal limite de continuidade de trabalho.

Exemplos:

```text
PRJ-AURORA
PRJ-MNFS
PRJ-PROGRAMMABLE-POWER-SUPPLY
PRJ-METAL-NOBRE-DATA-BRAIN
```

Um Project possui:

- vision;
- objectives;
- boundaries;
- status;
- roadmap;
- decisions;
- knowledge sources;
- memories;
- people;
- repositories;
- devices;
- environments;
- active missions;
- experiments;
- risks;
- incidents;
- next actions.

### 3.8.1 Project Snapshot

Uma projeção destinada a um Context Pack pode conter:

```text
project identity
current objective
current phase
accepted decisions
open hypotheses
active blockers
relevant resources
permitted next actions
freshness timestamp
```

Snapshot é uma projeção, não o owner de cada conteúdo.

---

## 3.9 Goal, Objective and Outcome

### Goal

Direção desejada, possivelmente ampla.

> “Construir uma fonte programável confiável.”

### Objective

Resultado delimitado e verificável dentro de um horizonte.

> “Validar estabilidade térmica da revisão B até 5 A.”

### Outcome

Resultado observado e consolidado após execução.

> “Estratégia C permaneceu estável por 60 minutos; estratégia D excedeu o limite térmico em 18 minutos.”

Aurora não deve confundir intenção com resultado.

---

## 3.10 Mission

`Mission` é uma unidade global de intenção coordenada pelo Aurora.

Pode atravessar vários domínios e harnesses.

Exemplo:

> “Desenvolver e validar um novo módulo de controle para a fonte programável.”

Uma Mission pode conter:

- objective;
- expected outcomes;
- acceptance criteria;
- scope/non-goals;
- assumptions;
- risks;
- dependencies;
- budgets;
- authority envelope;
- delegations;
- decisions;
- evidence;
- consolidated outcome;
- closeout.

Mission não é um processo nem uma conversa.

Dimensões de estado podem ser separadas:

```text
lifecycle = PROPOSED | ACTIVE | CLOSED | CANCELED
phase = INTAKE | PLANNING | AUTHORIZED | EXECUTING | VERIFYING | CLOSING
attention = NORMAL | BLOCKED | NEEDS_LEANDRO | DEGRADED | INCIDENT
```

Estados exatos permanecem sujeitos à Capability Realization Method e aos architecture spikes.

---

## 3.11 Delegation

`Delegation` é uma atribuição concreta de resultado a uma Capability/Provider dentro de uma Mission.

Exemplo:

```text
Mission:
  validate control module

Delegation DEL-001:
  research candidate topologies

Delegation DEL-002:
  simulate selected circuits

Delegation DEL-003:
  implement firmware variants

Delegation DEL-004:
  execute laboratory protocol
```

Uma Delegation possui identidade distinta de:

- provider instance;
- attempt;
- worker run;
- chat/session;
- effect request.

Ela é a boundary principal para:

- context;
- authority;
- budget;
- expected outcome;
- lifecycle;
- artifacts;
- evidence;
- escalation.

---

## 3.12 Capability

`Capability` descreve um resultado reutilizável que pode ser solicitado.

Exemplos:

```text
technical_research
software_delivery
firmware_build_and_flash
schematic_review
laboratory_measurement
memory_consolidation
evaluation_campaign
```

Capability não representa:

- a harness específica;
- o endpoint;
- a instância;
- a autorização;
- uma execução.

Ela define:

- semantic identity;
- input/output contracts;
- expected effect classes;
- evidence profile;
- compatibility and versioning.

---

## 3.13 Provider, Harness and Instance

### Provider

Entidade que oferece uma capability.

Pode ser:

- harness;
- deterministic service;
- external agent application;
- device controller;
- human workflow adapter.

### Harness

Sistema especializado completo que governa como executar um domínio.

Pode conter:

- agents;
- workflows;
- tools;
- runtimes;
- storage;
- local state machines;
- evaluation;
- recovery;
- environments;
- interfaces.

### Provider Instance

Execução concreta de um Provider em um ambiente e build exatos.

Exemplo:

```text
Provider: research-harness
Version: 1.2.0
Instance: research-local-01
Build digest: sha256:...
Environment: local-restricted
```

Trust e authority podem estar associados à instância exata.

---

## 3.14 Tool, Agent, Workflow and Runtime

### Tool

Operação delimitada.

Exemplos:

- read repository file;
- query instrument;
- run simulator;
- create branch;
- retrieve document.

### Agent

Unidade probabilística que interpreta contexto, decide e utiliza tools.

Agent não é automaticamente:

- state owner;
- authority grantor;
- verifier;
- harness;
- Aurora Identity.

### Workflow

Coordenação de etapas conhecidas ou adaptativas.

Pode ser:

- deterministic;
- state-machine based;
- graph based;
- agent-directed;
- durable.

### Runtime

Mecanismo que executa agents/workflows.

Exemplos candidatos:

- Pi;
- Mastra;
- LangGraph;
- OpenHands;
- OpenAI Agents SDK;
- durable workflow engine;
- code/firmware deterministic runtime.

Runtime não governa a semântica do Aurora.

---

## 3.15 Knowledge Source

`Knowledge Source` representa conteúdo consultável com identidade e provenance.

Exemplos:

- Product Blueprint;
- ADR;
- datasheet;
- paper;
- schematic;
- source code;
- manual;
- measurement report;
- website;
- database view;
- live sensor.

Atributos:

- source type;
- owner/publisher;
- authority class;
- version;
- observed/accessed time;
- freshness;
- sensitivity;
- integrity;
- citation/reference;
- access policy.

Uma Knowledge Source pode produzir Memory Items, mas permanece distinta deles.

---

## 3.16 Memory Item

`Memory Item` é uma unidade preservada para orientar contexto futuro.

Possui:

- type;
- scope;
- content or structured payload;
- provenance;
- epistemic status;
- authority;
- confidence;
- validity;
- sensitivity;
- retention;
- relationships;
- lifecycle.

Uma memória pode apontar para uma Knowledge Source em vez de duplicar o conteúdo.

---

## 3.17 Decision

`Decision` registra uma escolha material, sua autoridade e consequências.

Tipos incluem:

- product decision;
- architecture decision;
- operational decision;
- experiment decision;
- safety decision;
- promotion decision;
- temporary delegation decision.

Uma Decision possui:

- question;
- alternatives;
- evidence;
- decider/authority;
- chosen option;
- rationale;
- consequences;
- scope;
- effective time;
- review triggers;
- supersession.

Uma fala em conversa pode ser evidence de aprovação, mas decisões duráveis materiais precisam de promoção à fonte canônica aplicável.

---

## 3.18 Hypothesis

`Hypothesis` é uma explicação ou expectativa ainda não confirmada.

Exemplo:

> “A oscilação acima de 2,7 A é causada pela compensação inadequada da malha.”

Possui:

- claim;
- supporting observations;
- contradicting observations;
- confidence;
- test method;
- scope;
- status;
- related hypotheses;
- result.

Status possível:

```text
PROPOSED
TESTABLE
UNDER_TEST
SUPPORTED
REFUTED
INCONCLUSIVE
SUPERSEDED
```

Aurora não deve armazenar hipótese como fato.

---

## 3.19 Experiment, Variant and Run

### Experiment

Estrutura que testa uma ou mais hypotheses.

### Variant

Configuração candidata submetida à comparação.

### Experiment Run

Execução concreta em ambiente, versão e condições específicas.

Um experimento material registra:

- baseline;
- variables;
- controls;
- immutable evaluation criteria;
- environment;
- variants;
- runs;
- randomization/repetition where applicable;
- observations;
- artifacts;
- cost;
- safety limits;
- conclusion authority.

---

## 3.20 Observation and Measurement

### Observation

Algo percebido por pessoa, modelo, software ou sensor.

### Measurement

Observation quantitativa ligada a instrumento, unidade, calibration e uncertainty.

Exemplo:

```yaml
measurement:
  signal: output_voltage_ripple
  value: 92
  unit: mVpp
  instrument: oscilloscope-01
  calibration_ref: CAL-2026-04
  observed_at: 2026-09-14T22:13:04-03:00
  conditions:
    load_current_a: 2.7
    ambient_temperature_c: 27.4
```

Observation não é automaticamente evidence suficiente. Quality e provenance importam.

---

## 3.21 Artifact, Claim, Receipt, Evidence, Verdict and Outcome

```text
Artifact
→ materialized output: code, report, dataset, firmware, trace, image

Claim
→ executor states that a criterion was satisfied

Receipt
→ controlled system records a verification or effect

Evidence
→ support linked to an explicit criterion or hypothesis

Verdict
→ permitted authority evaluates evidence

Outcome
→ consolidated result of mission/delegation/experiment
```

Essa separação impede que “a harness disse que terminou” se torne aceitação global.

---

## 3.22 Authority Grant

`Authority Grant` representa permissão delegada, limitada e revogável.

Possui:

- grantor;
- subject;
- actor/executor;
- allowed actions/effects;
- resources;
- environment;
- validity window;
- budget;
- guardrails;
- escalation;
- revocation status;
- credential references;
- policy and contract hashes.

Uma Capability tecnicamente disponível não cria um Authority Grant.

---

## 3.23 Effect Request and Effect Receipt

### Effect Request

Solicitação de alteração em um sistema externo ou no mundo físico.

### Effect Receipt

Registro do que o gateway conseguiu executar, negar, falhar ou reconciliar.

Efeitos incluem:

- filesystem write;
- network call;
- credential use;
- repository change;
- message or email;
- deployment;
- purchase;
- device command;
- physical actuation.

Campos críticos:

- idempotency key;
- actor;
- target resource;
- action;
- request hash;
- policy decision;
- start/end time;
- result;
- external reference;
- ambiguity status.

---

## 3.24 Budget

`Budget` governa recursos consumíveis.

Dimensões possíveis:

- money;
- tokens;
- model calls;
- compute time;
- wall-clock time;
- storage;
- network;
- energy;
- number of experiments;
- firmware flash cycles;
- mechanical cycles;
- instrument occupancy;
- attention/interruption.

Budgets possuem:

- units;
- hard/soft limits;
- warning thresholds;
- scope;
- consumption source;
- reconciliation;
- extension authority.

---

## 3.25 Guardrail and Interlock

### Guardrail

Regra que limita decisão ou ação.

### Interlock

Mecanismo determinístico e independente que impede ou interrompe condição perigosa.

Exemplo:

```text
Aurora policy:
  do not exceed 500 mA during first power-up

Bench interlock:
  hardware/current limiter physically prevents > 600 mA
```

Guardrail de prompt não é interlock.

---

## 3.26 Incident, Finding and Improvement Opportunity

### Incident

Evento que produziu ou quase produziu impacto material.

### Finding

Problema, divergência ou risco identificado durante review, verification ou operation.

### Improvement Opportunity

Padrão candidato a investigação e melhoria.

### Improvement Candidate

Versão concreta de uma mudança proposta, ligada a:

- causal hypothesis;
- affected components;
- experiment plan;
- evaluation sets;
- risk;
- rollback;
- promotion status.

Um único Incident não obriga uma mudança. Correlation e causal analysis precedem generalização.

---

## 3.27 Relações canônicas

Exemplos:

```text
Person OWNS Project
Person AUTHORIZES AuthorityGrant
Aurora COORDINATES Mission
Mission CONTAINS Delegation
Delegation REQUESTS Capability
Provider OFFERS Capability
ProviderInstance EXECUTES Delegation
Delegation RECEIVES ContextPack
Delegation OPERATES_UNDER AuthorityGrant
Harness USES Tool
Tool PRODUCES EffectRequest
EffectGateway EMITS EffectReceipt
Experiment TESTS Hypothesis
ExperimentRun PRODUCES Observation
Artifact SUPPORTS Evidence
Evidence ADDRESSES AcceptanceCriterion
Decision SUPERSEDES Decision
MemoryItem DERIVES_FROM KnowledgeSource
Presence RUNS_ON Device
Device EXISTS_IN Environment
```

Relações materiais carregam temporalidade, provenance e authority quando aplicável.

---

## 3.28 Source-of-truth ownership

| Conceito | Owner canônico futuro |
|---|---|
| Aurora constitutional identity | Product Blueprint / accepted constitutional changes |
| Current project status | Project state store |
| Accepted architecture decision | ADR |
| Raw source code | Git repository |
| Current device telemetry | Device/telemetry source |
| Device inventory | Device Registry |
| Capability/provider trust | Capability Registry |
| Active mission/delegation state | Aurora operational state |
| Harness internal execution | Harness local state, reconciled globally |
| Authority | Authority/Policy plane |
| External effect result | Effect Receipt + target-system reference |
| Historical conversation | Session ledger/archive |
| Memory synthesis | Memory system, never above original authority |
| Experiment result | Experiment/Evidence store |

Exact storage technology remains open.

---

## 3.29 World Model views

Aurora should not load one universal graph for every question. It may create views:

### Personal view

Leandro, preferences, goals, relationships, attention and permissions.

### Project view

Decisions, artifacts, people, devices, missions, hypotheses and blockers.

### Operational view

Active delegations, budgets, grants, incidents and permitted actions.

### Physical environment view

Devices, topology, sensors, actuators, location, state and safety class.

### Knowledge view

Sources, authority, citations, freshness and contradictions.

### Failure view

Incidents, symptoms, causal hypotheses, candidates and evaluations.

Views are projections over canonical entities and sources.

---

## 3.30 Example: programmable power supply

```text
Person LEANDRO
  OWNS Project PRJ-POWER-SUPPLY

Project
  HAS Device PCB-REV-B
  HAS Repository FW-POWER-SUPPLY
  HAS Decision ADR-POWER-004
  HAS Hypothesis HYP-017
  HAS Mission MIS-VALIDATE-STABILITY

Hypothesis HYP-017
  "loop compensation causes oscillation above 2.7 A"

Mission
  CONTAINS Delegation DEL-LAB-TEST
  CONTAINS Delegation DEL-FW-VARIANTS

DEL-LAB-TEST
  REQUESTS Capability laboratory_measurement
  OPERATES_UNDER Grant GRANT-LAB-021
  PRODUCES Observation OBS-RIPPLE-092
  PUBLISHES Artifact waveform-014

Observation
  SUPPORTS Evidence EVID-HYP-017-03

Verdict
  HYP-017 remains INCONCLUSIVE
```

This model lets Aurora say what is known, how it is known and what remains undecided.

---

## 3.31 Failure modes

### Generic-task collapse

Everything becomes a task and loses authority/lifecycle distinctions.

### Identity aliasing

Provider, instance, process and delegation share one identifier, making recovery ambiguous.

### Memory as truth

Compressed observation overrides current state.

### Graph overreach

Every sentence becomes a relationship, generating noise and false certainty.

### Stale world model

Aurora acts on last-known device state without live verification.

### Authority smuggling

A relationship or parent delegation is treated as permission.

### Hidden time

Old firmware/device relation remains active after an update.

### Evidence flattening

Claim, receipt and verdict are stored as one “result”.

### Cross-project contamination

An entity from one project enters another Context Pack because IDs or scopes are weak.

---

## 3.32 Evaluation requirements

Future implementation must prove:

1. identities survive restart and provider replacement;
2. distinct concepts are not collapsed in schemas;
3. temporal relationships supersede correctly;
4. hypotheses never become facts without verdict;
5. memory cannot overwrite a higher-authority source;
6. active authority is independently resolvable;
7. Context Builder can create project-scoped views;
8. stale device state triggers live verification;
9. cross-project isolation prevents contamination;
10. every material artifact/effect is attributable to an actor, delegation and version.

---

## 3.33 Non-goals

This section does not choose:

- relational versus graph database;
- event sourcing architecture;
- object model implementation language;
- one universal ontology;
- public multi-tenant identity;
- automatic extraction of every entity from conversation;
- a semantic-web standard;
- exact IDs or schema version syntax.

Those choices require specs, research and spikes after the domain has been reviewed.
