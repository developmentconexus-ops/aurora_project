---
id: DOC-AURORA-BLUEPRINT-10
title: Autonomia, Autoridade e Segurança
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
  - autonomy model
  - authority model
  - self-improvement limits
  - safety principles
related:
  - DOC-AURORA-BLUEPRINT-02
  - DOC-AURORA-BLUEPRINT-05
  - DOC-AURORA-BLUEPRINT-07
last_reviewed: 2026-08-05
---

# 10. Autonomia, Autoridade e Segurança

## 10.1 Modelo

Aurora opera por **autoridade progressiva e autonomia delegada**.

```text
observar
→ analisar
→ recomendar
→ preparar
→ executar effect autorizado
→ confirmar ação material
→ interromper emergência
```

A autonomia pode ser concedida por ação, workflow ou campanha.

## 10.2 Níveis

### N1 — Ação autorizada

Executa uma operação específica.

### N2 — Workflow autorizado

Conduz sequência conhecida dentro de limites.

### N3 — Campanha autônoma

Recebe objetivo e adapta hipóteses, testes e estratégia.

### N4 — Programa contínuo

Monitora um sistema, identifica oportunidades e executa campanhas previamente autorizadas. Futuro; não é comportamento padrão atual.

## 10.3 Envelope de Autonomia

Toda campanha material possui:

- objetivo;
- baseline;
- espaço modificável;
- ambiente;
- authority;
- budget;
- critérios de avaliação;
- guardrails;
- stop conditions;
- escalation rules;
- evidence requirements;
- promotion rules;
- rollback.

> Aurora é autônoma dentro do envelope e conservadora na fronteira.

## 10.4 Guardrails determinísticos

Em effects físicos e digitais críticos, limites não podem depender apenas do julgamento do modelo.

Exemplos:

- corrente;
- tensão;
- temperatura;
- movimento;
- timeouts;
- gastos;
- rede;
- credentials;
- filesystem;
- produção.

A harness pode decidir dentro do espaço seguro. Não redefine o espaço seguro durante a execução.

## 10.5 Policy e enforcement

Camadas:

```text
Constitution
→ invariantes

Authority Grant
→ permissão delimitada

Policy Decision Point
→ decide

Effect Gateway
→ aplica no ponto da ação

Sandbox / OS boundary
→ contém

Audit / Receipt
→ prova
```

SDK e prompt orientam, não aplicam a boundary.

## 10.6 Delegação de identidade

O sistema distingue:

- Leandro, autoridade original;
- Aurora, ator delegado;
- harness, executor delegado;
- worker, ator operacional;
- dispositivo, presença e boundary;
- delegation, escopo concreto.

Tokens e credentials devem ser curtos, escopados e revogáveis. Segredos não são entregues diretamente quando um broker ou reference resolve.

## 10.7 Effects

Categorias iniciais:

- read;
- compute;
- write local;
- network;
- credential use;
- repository change;
- external communication;
- purchase/financial;
- deployment;
- device control;
- physical actuation.

Risk e reversibilidade informam approval e enforcement.

## 10.8 Autoaperfeiçoamento

Aurora pode:

- observar falhas e acertos;
- correlacionar padrões;
- investigar causas;
- formular hipóteses concorrentes;
- criar candidatas;
- executar evals em sandbox;
- preparar promoção.

Aurora não pode promover autonomamente mudanças materiais em:

- identidade constitucional;
- autoridade de Leandro;
- guardrails;
- security policies;
- promotion rules;
- revogação;
- audit;
- shutdown;
- capacidade de conceder authority a si mesma.

## 10.9 Failure Intelligence

Ciclo:

```text
observe
→ record
→ correlate
→ causal analysis
→ reproduce
→ competing hypotheses
→ candidate
→ broad evaluation
→ independent review
→ promotion
→ canary
→ rollback or graduate
```

Uma correção não é melhoria se apenas memoriza o caso original.

## 10.10 Avaliação de candidatas

- incidente original;
- casos vizinhos;
- casos contrários;
- regressão histórica;
- casos inéditos;
- casos adversariais;
- custo;
- latência;
- segurança;
- personalidade;
- memória;
- interrupções;
- explicabilidade.

O sistema que propõe a melhoria não é a única autoridade que a aceita.

## 10.11 Detecção contínua, mudança deliberada

Aurora pode registrar oportunidades continuamente. Experimentos exigem janela e envelope autorizados. Promoção segue gate.

> Aurora aprende continuamente, mas muda deliberadamente.

## 10.12 Revogação e contenção

Revogar authority deve:

- impedir novas ações;
- cancelar quando possível;
- fechar channels;
- revogar tokens;
- bloquear credentials;
- registrar receipts;
- preservar evidence;
- reconciliar effects já ocorridos.

## 10.13 Emergência

Ação automática de segurança pode ser autorizada previamente, por exemplo:

- desenergizar ao exceder corrente;
- interromper workflow que excede budget;
- bloquear exfiltração;
- revogar provider comprometido.

Emergency authority é estreita, observável e testada.
