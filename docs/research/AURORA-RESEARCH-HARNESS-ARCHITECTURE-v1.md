---
id: RESEARCH-AURORA-HARNESS-ARCHITECTURE-V1
title: Harness Architecture, Interoperability, SDK, Durability and Authority
document_type: research_report
form: explanation
authority: research
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - research findings for capability and harness architecture
related:
  - RESEARCH-AURORA-HARNESS-ARCHITECTURE-V1-SOURCES
  - DOC-AURORA-BLUEPRINT-05
  - DOC-AURORA-BLUEPRINT-07
  - DOC-AURORA-BLUEPRINT-10
last_reviewed: 2026-08-05
---

# Harness Architecture Research v1

## 1. Pergunta

Como Aurora pode coordenar harnesses heterogêneas — agênticas, determinísticas, locais, remotas e físicas — com contracts, SDKs, authority, durability, observability e recovery, sem:

- ficar presa a um framework;
- reinventar protocolos maduros;
- permitir que protocol ou SDK vire a constituição;
- confiar em manifests sem verificação;
- depender de processos efêmeros;
- confundir guardrails de agente com security boundary?

## 2. Método

A pesquisa consultou apenas:

- especificações oficiais;
- documentação oficial;
- repositórios oficiais;
- standards;
- release notes dos mantenedores.

Data de corte: **2026-08-05**.

As fontes são identificadas como `[Sxx]` e detalhadas no manifest associado.

## 3. Conclusão executiva

A direção aprovada é sólida:

1. Aurora deve possuir domínio e contracts próprios.
2. Harnesses devem ser relativamente opacas e especializadas.
3. Protocols devem ser bindings substituíveis.
4. First-party harnesses devem seguir um Harness Development Kit por política.
5. Contract e conformance devem permanecer independentes do SDK.
6. MCP e A2A são complementares, não alternativas totais.
7. Durable execution é uma camada separada do agent framework.
8. Policy decision, effect enforcement e sandbox são responsabilidades separadas.
9. OpenTelemetry é a melhor baseline vendor-neutral para observability.
10. Build identity e provenance são parte da confiança de uma capability.

Nenhuma fonte sustenta a ideia de escolher um único framework como núcleo universal de uma plataforma com esse escopo.

## 4. Cinco camadas que não devem ser confundidas

### 4.1 Semântica do Aurora

Define Capability, Harness, Delegation, Authority Grant, Context Pack, Effect, Artifact, Evidence, Outcome e Recovery.

Deve permanecer sob governança do Product Blueprint, ADRs e Specs do Aurora.

### 4.2 Contracts e schemas

Representam a semântica em formatos legíveis por máquinas.

Precisam de:

- versioning;
- compatibility;
- breaking-change detection;
- code generation;
- validation;
- stable IDs.

### 4.3 Development Kit

Fornece experiência de alto nível e Golden Paths.

Não governa semântica nem segurança.

### 4.4 Protocol bindings

MCP, A2A, local RPC, HTTP, gRPC e event transports.

Transportam contracts.

### 4.5 Runtime interno

Mastra, LangGraph, Pi, OpenHands, OpenAI Agents SDK, durable engine, código determinístico ou firmware.

Implementam a harness.

## 5. MCP

A versão `2026-07-28` transformou o núcleo do MCP em stateless request/response, adicionou discovery opcional, routing por headers, caching, authorization hardening e extensions formais. Tasks foi movido para uma extension com durable handles e polling [S01][S03].

A conformance suite oficial testa clients e servers por scenarios e versões, incluindo core, extensions, auth e draft [S02].

### Adequação ao Aurora

MCP é forte para:

- tools;
- resources;
- operações delimitadas;
- discovery de interfaces;
- integração com ecossistema amplo;
- calls locais ou remotas;
- Tasks quando uma tool call precisa durar.

### Limite

MCP não define, por si só:

- objetivo global;
- decomposition entre domínios;
- authority envelope do Aurora;
- budget global;
- acceptance criteria;
- evidence model;
- trust lifecycle;
- cross-harness composition;
- constitutional boundaries.

Mesmo Tasks resolve lifecycle de uma operação MCP, não o domínio completo de uma Delegation Aurora.

### Veredito

**Adotar como binding candidato de tools/resources.**  
Não adotar como modelo canônico de missão ou harness.

## 6. A2A

A2A 1.0 descreve colaboração entre aplicações agênticas opacas. Agent Cards declaram skills, interfaces e capabilities. Tasks possuem state, artifacts, polling, streaming, push e cancellation [S04][S05].

A spec separa:

- Message: comunicação;
- Task: unidade stateful;
- Artifact: deliverable.

Ela alerta que mensagens de stream não constituem entrega confiável de informação crítica; critical information precisa estar em state/artifact recuperável [S05].

Agent Cards podem ser descobertos por well-known URI, registry ou configuração direta e podem ser assinados [S05]. O projeto oferece TCK [S06].

A maturidade dos SDKs varia por linguagem na data de corte: Python declara 1.0 completo, enquanto JavaScript ainda mantém v0.3 como linha estável e 1.0 em alpha [S07][S08].

### Adequação ao Aurora

A2A é forte para:

- harness remota;
- execução opaca;
- trabalho assíncrono;
- tasks e artifacts;
- streaming/polling/push;
- Agent Card discovery;
- multi-turn.

### Limite

A2A não define:

- classes de memória do Aurora;
- authority grants detalhados;
- effect policy;
- global budget;
- causal evidence;
- provider trust assessment;
- constitutional protection;
- acceptance global.

### Veredito

**Investigar como binding principal para harnesses remotas e opacas**, envolvido por Aurora Delegation Envelope e submetido ao TCK e à conformance própria.

## 7. MCP e A2A juntos

Modelo provável:

```text
Aurora ──A2A/Native RPC──> Harness
Harness ──MCP────────────> Tools/Resources
```

Não é obrigatório que toda harness use ambos.

- uma harness local first-party pode usar native SDK/RPC;
- uma harness externa pode usar A2A;
- uma tool estreita pode usar MCP;
- um device controller pode usar protocolo determinístico e adapter.

## 8. Contract-first, SDK e conformance

OpenTelemetry demonstra uma separação madura entre specification, API, SDK, Semantic Conventions e contribuições. API e SDK possuem versões independentes da spec e alternative implementations são possíveis [S09][S10].

MCP e A2A mantêm conformance/TCK próprios [S02][S06].

### Implicação

O Aurora deve adotar:

```text
Canonical Contract Model
→ language-neutral schemas
→ generated API/types
→ official AHDK implementation
→ black-box Conformance Kit
```

### Política first-party

Harnesses first-party usam AHDK por padrão obrigatório.

Waiver exige:

- razão material;
- owner;
- escopo;
- risk assessment;
- exit condition;
- conformance integral.

External/legacy providers podem usar adapters ou implementação direta.

### Por que não tornar o SDK a spec

- impedir lock-in de linguagem;
- permitir adapters;
- suportar external providers;
- testar o próprio SDK;
- separar semantic versioning;
- preservar migração;
- evitar que bug da biblioteca vire regra de domínio.

## 9. Conteúdo do AHDK

O Development Kit deve incluir:

1. generated types;
2. manifest builder;
3. lifecycle client;
4. Context Pack reader;
5. decision request API;
6. artifact/evidence API;
7. effect request client;
8. cancellation/deadline;
9. checkpoint/resume;
10. OpenTelemetry instrumentation;
11. simulator;
12. mocks;
13. fault injection;
14. scaffolder;
15. conformance runner;
16. compatibility report;
17. build metadata/provenance hooks.

Backstage Software Templates demonstra o valor de scaffolding com inputs, review, task execution, dry-run, cancellation e padrões organizacionais incorporados [S31].

### Veredito

AHDK é mais que client library. É um **Harness Development Kit** e Golden Path.

## 10. Schemas e events

### JSON Schema

Draft 2020-12 é forte para:

- manifests;
- configurations;
- JSON artifacts;
- portable validation;
- human-readable contracts [S13].

### Protobuf

Forte para:

- typed RPC;
- code generation;
- compact wire format;
- compatibility checks.

Buf oferece breaking-change detection automatizada [S14].

### CloudEvents

Fornece envelope protocol-neutral para event metadata [S11].

### AsyncAPI

Documenta channels, messages, operations e bindings de APIs event-driven [S12].

### Veredito

Não escolher um único formato para tudo.

Hipótese a testar:

```text
JSON Schema
→ manifests e contracts portáveis

CloudEvents
→ envelope externo de eventos

AsyncAPI
→ documentação de channels

Protobuf/gRPC
→ bindings de alta performance quando necessário
```

## 11. Observability

OpenTelemetry separa API, SDK e Semantic Conventions, permitindo instrumentação vendor-neutral e exporters substituíveis [S09][S10].

Aurora precisa de semantic conventions próprias:

```text
aurora.mission
aurora.delegation
aurora.harness.run
aurora.capability.invoke
aurora.effect.request
aurora.effect.execute
aurora.artifact.publish
aurora.evidence.record
aurora.decision.request
```

Cada trace deve preservar:

- Leandro/Aurora/harness identities apropriadas;
- mission/delegation IDs;
- model/provider;
- context classification;
- budget;
- effect;
- artifact/evidence links;
- error/recovery;
- privacy redaction.

### Veredito

OpenTelemetry é baseline recomendada. Storage e backend não estão escolhidos.

## 12. Durable execution

Campanhas podem durar horas ou dias e sobreviver a restart. Process memory e transcript são insuficientes.

### Temporal

Event History, deterministic replay e Activities para external I/O fornecem forte durability, mas trazem deployment e programming model próprios [S15].

### DBOS

Biblioteca sobre PostgreSQL, com checkpoints de workflows/steps e menor infraestrutura inicial; produção de alta disponibilidade pode adicionar control plane [S16].

### Restate

Journal, durable steps, state, timers e reliable communication em um runtime dedicado [S17].

### Inngest

Functions e steps checkpointados, retries independentes e resume a partir de resultados memoizados [S18].

### Evidência de separação de camadas

Mastra adicionou integração com Temporal para workflows que precisam durar horas/dias e sobreviver a worker restart [S19]. Isso mostra que agent/workflow framework e durable execution engine são responsabilidades distintas.

### Veredito

Criar `DurableExecutionPort`.

Não selecionar engine por pesquisa documental. Executar spike comparativo incluindo:

- baseline mínima;
- DBOS;
- Restate;
- Temporal;
- outro candidato somente se requisito justificar.

Critérios:

- local-first;
- operational burden;
- restart;
- idempotency;
- timers;
- HITL;
- cancellation;
- versioning;
- observability;
- language fit;
- backup/restore;
- portability.

## 13. Identity, delegation e policy

### Cedar

Cedar modela principal, action, resource e context; documentação oficial discute agentes agindo em nome de usuários e duas formas de representar user/agent [S20].

### OPA

OPA separa policy decision de enforcement [S21].

### OAuth Token Exchange

RFC 8693 diferencia delegation e impersonation e representa subject/actor [S22].

### SPIFFE

SVIDs são referência para workload identity verificável e de curta duração [S23].

### Implicação

Aurora precisa representar:

```text
subject: Leandro
actor: Aurora
executor: Harness/Worker
action: Effect
resource: target
context: delegation, device, environment, risk
```

### Veredito

Criar abstrações:

- WorkloadIdentity;
- AuthorityGrant;
- DelegationToken;
- PolicyDecisionPoint;
- EffectGateway;
- CredentialBroker;
- EffectReceipt.

Cedar versus OPA e adoção imediata de SPIFFE permanecem abertos a spike.

## 14. SDK não é security boundary

Uma harness pode ignorar uma library e usar filesystem, network ou process APIs diretamente.

Logo:

```text
AHDK
→ ergonomia e conformidade

PDP
→ decisão

Effect Gateway
→ enforcement

Sandbox/OS
→ containment

Audit/Receipt
→ evidence
```

Prompts, guardrails e manifests não substituem enforcement.

## 15. Supply-chain trust

SLSA provenance registra como, onde e a partir de quais materiais um artifact foi construído [S24].

Confiança de provider deve se ligar a:

- publisher;
- source revision;
- build digest;
- provenance;
- contract version;
- SDK version;
- conformance result;
- environment;
- approval scope.

Versão textual não basta.

## 16. Frameworks

### LangGraph

Runtime de baixo nível para agentes stateful e long-running, com durable execution, streaming e HITL [S25].

### Mastra

Framework TypeScript com agents, typed workflows, loops, branches, suspend/resume, memory e MCP; pode trocar engine em certos cenários [S19][S26].

### Pi

Runtime/coding agent com SDK e RPC. SDK é natural em TypeScript in-process; RPC atende cross-language/process isolation [S27][S28].

### OpenHands

Especializado em software engineering agents, workspaces, tools e agent server [S29].

### OpenAI Agents SDK

Agents, tools, handoffs, guardrails, sessions e tracing [S30].

### Veredito

Todos são candidatos internos de harness. Nenhum cobre, sozinho, identidade pessoal, memory governance, device presence, global authority, multi-harness contracts e physical safety.

## 17. Arquitetura recomendada

```text
Product Constitution
        ↓
Canonical Contract Model
        ↓
AHDK + Conformance Kit
        ↓
Bindings: Native | RPC | A2A | MCP | HTTP/gRPC
        ↓
Harness runtimes heterogêneos
```

Cross-cutting:

- Capability Registry;
- Durable Execution Port;
- PDP;
- Effect Gateways;
- Credential Broker;
- Artifact/Evidence Stores;
- OpenTelemetry;
- provenance.

## 18. Riscos

### Overengineering

Mitigação: spikes, one SDK first, reference harnesses mínimas.

### Protocol churn

Mitigação: adapters, conformance por version, contract Aurora independente.

### Semantic duplication

Mitigação: canonical definitions e code generation.

### Security theater

Mitigação: enforcement fora do SDK e testes adversariais.

### Framework lock-in

Mitigação: same capability in two runtimes no spike.

### State split-brain

Mitigação: owner explícito de global/local state e reconciliation protocol.

### Event loss

Mitigação: critical state/artifacts duráveis; stream não é memória.

### Trust inheritance

Mitigação: approvals por exact build/scope.

## 19. Architecture spikes recomendados

1. Contract + SDK + direct protocol implementation.
2. MCP binding.
3. A2A binding e TCK.
4. Durable execution comparison.
5. Authority + Effect Gateway.
6. Distributed trace.
7. Registry + provenance.
8. Framework neutrality.

## 20. Decisões informadas, não encerradas

A pesquisa sustenta:

- domain-owned contracts;
- AHDK first-party;
- conformance universal;
- protocol neutrality;
- MCP/A2A complementarity;
- separate durability;
- separate policy/enforcement;
- OTel baseline;
- build-bound trust.

Permanece aberto:

- language;
- schema split;
- local RPC binding;
- durable engine;
- PDP implementation;
- workload identity implementation;
- registry storage;
- first reference runtime.

## 21. Limitações

- Nenhum benchmark local foi executado.
- Nenhum protocol binding foi prototipado.
- Documentação de frameworks muda rapidamente.
- A2A SDK maturity varia.
- MCP `2026-07-28` é recente.
- Security claims precisam de adversarial testing.
- Local-first operational burden só será conhecido por spike.

Portanto, este report autoriza design e spikes, não uma escolha de stack.
