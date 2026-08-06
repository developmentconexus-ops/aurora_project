---
id: DOC-AURORA-BLUEPRINT-07
title: Orquestração de Harnesses
document_type: product_blueprint_section
form: explanation
authority: constitutional
status: proposed
version: 0.1.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - Aurora-harness boundary
  - hierarchical orchestration
  - delegation model
  - protocol neutrality
related:
  - DOC-AURORA-BLUEPRINT-05
  - DOC-AURORA-BLUEPRINT-10
  - ADR-AURORA-0001
  - ADR-AURORA-0002
  - DESIGN-AURORA-CAPABILITY-HARNESS-ARCHITECTURE
last_reviewed: 2026-08-05
---

# 7. Orquestração de Harnesses

## 7.1 Definição

Uma harness é um sistema especializado que pode conter:

- agentes;
- workflows;
- tools;
- máquinas de estado;
- runtime;
- storage;
- ambientes;
- políticas locais;
- avaliações;
- recuperação;
- interfaces humanas;
- metodologia própria.

Aurora não administra cada prompt interno. Ela delega um resultado esperado dentro de um contrato.

## 7.2 Fronteira de autoridade

### Aurora possui

- intenção original;
- objetivo global;
- identidade e memória global;
- modelo de projeto;
- decomposição entre domínios;
- seleção de capabilities;
- authority grants;
- budget global;
- dependências;
- estado global da missão;
- composição de artefatos;
- comunicação com Leandro;
- acceptance global.

### Harness possui

- estratégia especializada;
- agentes e workflows internos;
- ferramentas do domínio;
- plano operacional local;
- retries locais permitidos;
- artefatos intermediários;
- critérios técnicos locais;
- estado detalhado de execução.

### Fronteira compartilhada

- delegation contract;
- context pack;
- authority grant;
- budget;
- lifecycle;
- eventos;
- pedidos de decisão;
- artefatos;
- evidências;
- outcome;
- cancelamento e recovery.

> Aurora governa o porquê, o quê, os limites e a composição. A harness governa como produzir o resultado especializado.

## 7.3 Orquestração hierárquica

Harnesses podem solicitar colaboração, mas Aurora governa toda delegação entre boundaries.

```text
Harness A pede capability B
→ Aurora avalia necessidade
→ seleciona provider
→ cria Context Pack mínimo
→ concede authority própria
→ acompanha a delegação
→ entrega o resultado relevante a A
```

Authority não é herdada silenciosamente.

## 7.4 Control plane e data plane

O control plane é centralmente governado:

- identidade;
- contratos;
- state;
- authority;
- budget;
- audit;
- decisions.

O data plane pode ser direto quando autorizado:

```text
Laboratory Harness
──── telemetria autorizada ────>
Evaluation Harness
```

Aurora registra channel, schema, limites, retention e provenance. Um canal direto não cria authority nova.

## 7.5 Delegation Envelope

Uma delegação contém:

### Identidade

- mission ID;
- delegation ID;
- parent delegation;
- capability;
- provider instance;
- contract version.

### Intenção

- objetivo;
- expected outcome;
- scope;
- non-goals;
- assumptions;
- dependencies.

### Contexto

- references;
- authority sources;
- project snapshot;
- authorized memories;
- data classification;
- freshness requirements.

### Controle

- authority grant;
- allowed/prohibited effects;
- environment;
- budget;
- deadline/window;
- guardrails;
- stop conditions;
- escalation rules.

### Qualidade

- acceptance criteria;
- evidence requirements;
- output schemas;
- reproducibility;
- reviewer requirements.

## 7.6 Context Pack

Context Pack não é o histórico integral.

Ele contém apenas:

- objetivo;
- estado relevante;
- fontes canônicas;
- memórias autorizadas;
- schemas;
- constraints;
- authority snapshot;
- artifact references;
- freshness metadata.

Contexto sensível pode ser entregue por referência ou token de acesso estreito.

## 7.7 Lifecycle proposto

```text
PROPOSED
→ AUTHORIZED
→ QUEUED
→ RUNNING
→ WAITING_FOR_INPUT | BLOCKED
→ CANCEL_REQUESTED
→ COMPLETED | FAILED | CANCELED | REJECTED
```

Estados exatos serão validados por spikes. Invariantes:

- transições são estruturadas;
- terminal é explícito;
- `success: true` não basta;
- cancelamento é cooperativo quando necessário;
- restart não perde o estado global;
- retries não duplicam effects.

## 7.8 Eventos

Eventos significativos:

- delegation authorized;
- execution started;
- checkpoint created;
- progress material;
- hypothesis changed;
- decision requested;
- budget threshold;
- warning;
- block;
- effect requested/executed/denied;
- artifact published;
- evidence recorded;
- completion/failure/cancel.

Mensagem transitória não é memória nem entrega confiável. Informação crítica deve existir em estado ou artefato recuperável.

## 7.9 Artefatos e evidência

Distinções:

```text
Message
→ comunicação

Artifact
→ produto ou dado entregue

Observation
→ algo observado

Claim
→ afirmação do executor

Receipt
→ evidência de uma verificação ou effect

Evidence
→ suporte ligado a critério

Verdict
→ decisão da autoridade permitida

Outcome
→ resultado consolidado
```

A harness não certifica sozinha a própria conclusão quando o risco exige verificação independente.

## 7.10 Recovery

Toda harness declara:

- resumable;
- checkpoint semantics;
- idempotency;
- cancellation;
- timeout;
- heartbeat;
- already-produced effects;
- reconciliation behavior;
- retention;
- failure taxonomy.

Aurora mantém o estado global e reconcilia o estado da harness após restart ou perda de conexão.

## 7.11 Protocol bindings

### Native SDK / in-process

First-party, baixo overhead, tipagem e ergonomia.

### Local RPC

Isolamento de processo, interoperabilidade entre linguagens e restart independente.

### A2A

Candidato para aplicações agênticas remotas, opacas e long-running.

### MCP

Candidato para tools, resources e operações delimitadas; Tasks pode apoiar trabalho assíncrono, mas não substitui o Delegation Model.

### HTTP/gRPC/event transport

Bindings adicionais quando necessários.

> Transporte carrega semântica; não a define.

## 7.12 Framework neutrality

Uma harness pode usar:

- Mastra;
- LangGraph;
- Pi;
- OpenHands;
- OpenAI Agents SDK;
- workflow engine;
- código determinístico;
- firmware.

Nenhum runtime se torna autoridade do domínio Aurora.

## 7.13 Durable Execution Port

Campanhas longas não dependem da vida de um processo agêntico.

A porta abstrata cobre:

- scheduling;
- durable timers;
- checkpoints;
- retries;
- leases;
- resumable steps;
- event delivery;
- recovery.

Temporal, DBOS, Restate e outros permanecem candidatos. A escolha exige spike.

## 7.14 Integração futura com MNFS

MNFS será provider de capabilities de engenharia de software. Sua prontidão, runtime e modelo interno não definem Aurora.

A integração ocorrerá por adapter/SDK após:

- contratos do Aurora estabilizarem;
- MNFS oferecer uma boundary adequada;
- conformance ser demonstrada;
- authority e recovery serem mapeados.

## 7.15 Não objetivos atuais

- delegação peer-to-peer irrestrita;
- propagação transitiva de credentials;
- estado global dentro de uma harness;
- tráfego obrigatoriamente roteado pelo Core;
- um único framework para todos os domínios;
- coordenação baseada em texto livre;
- integração prematura com MNFS.
