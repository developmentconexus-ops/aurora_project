---
id: DOC-AURORA-ROADMAP
title: Aurora Capability Roadmap
document_type: product_roadmap
form: reference
authority: constitutional
status: proposed
version: 0.1.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - capability sequence
  - product milestone intent
related:
  - DOC-AURORA-BLUEPRINT-01
  - DESIGN-AURORA-CAPABILITY-HARNESS-ARCHITECTURE
last_reviewed: 2026-08-05
---

# Aurora Capability Roadmap

## 1. Princípio

O roadmap é uma sequência de capacidades cumulativas, cada uma encerrada por uma prova ponta a ponta que reduz um risco material.

Não é:

- cronograma;
- lista fixa de frameworks;
- backlog de features;
- promessa de datas;
- arquitetura detalhada de marcos distantes.

## 2. Unidade

Cada Product Milestone possui:

```text
Outcome
Operator-visible value
Entry Criteria
Capabilities
Golden Proof
Exit Criteria
Non-goals
Dependencies
Replan Triggers
```

## A0 — Product and Architecture Baseline

### Outcome

Visão, authority, memory, capability system, harness boundary, research governance e roadmap tornam-se canônicos e retomáveis por sessões futuras.

### Golden Proof

```text
nova sessão
→ lê somente o repositório
→ explica Aurora corretamente
→ identifica fase, decisões e lacunas
→ respeita implementação proibida
→ aponta próxima ação
```

### Estado

```text
IN_REVIEW
```

## M0 — Sovereign Core Walking Skeleton

### Outcome

Um processo mínimo carrega identidade, projeto, authority snapshot e estado durável após restart.

### Golden Proof

```text
initialize
→ record project state
→ terminate
→ fresh process recovers exact state
```

### Non-goals

- LLM orchestration completa;
- multi-device;
- harness real;
- cloud.

## M1 — Governed Context and Conversation

### Outcome

Aurora mantém continuidade conversacional e de projeto com provenance, scopes e correction.

### Golden Proof

```text
conversation A creates approved decision
→ session ends
→ session B resumes project
→ retrieves correct source
→ ignores superseded candidate
→ explains provenance
```

## M2 — Capability Registry and Reference Provider

### Outcome

Aurora descobre, inspeciona, sandboxa e aprova uma capability de referência.

### Golden Proof

```text
manifest
→ discovered
→ conformance
→ approved for narrow scope
→ exact build recorded
```

## M3 — Contractual Delegation and Evidence

### Outcome

Aurora delega um trabalho simples e recebe artifact/evidence estruturados.

### Golden Proof

```text
intent
→ Delegation Envelope
→ authorized execution
→ artifact
→ verification
→ outcome
```

## M4 — Durable Delegation and Recovery

### Outcome

Delegation sobrevive a Core/provider restart, sem duplicar effects.

### Golden Proof

```text
start
→ effect
→ checkpoint
→ kill both sides
→ recover
→ no duplicate effect
→ complete
```

## M5 — Hierarchical Multi-Harness Composition

### Outcome

Duas harnesses colaboram por child delegation mediada pelo Aurora.

### Golden Proof

```text
Harness A requests capability B
→ Aurora creates child grant/context
→ direct authorized data channel optional
→ results reconcile
→ global criteria close
```

## M6 — First Real Engineering Harness

### Outcome

Uma harness real de engenharia entra pelo mesmo contract model.

Candidatos dependem de readiness:

- research;
- evaluation;
- MNFS;
- firmware;
- hardware.

MNFS não é dependência obrigatória deste marco.

## M7 — Delegated Experimental Campaigns

### Outcome

Aurora conduz uma campanha adaptativa com budget, guardrails, stop conditions e escalation.

### Golden Proof

```text
baseline
→ hypotheses
→ variants
→ evaluations
→ budget respected
→ no unauthorized promotion
→ evidence report
```

## M8 — Multi-Presence Continuity

### Outcome

Mesma Aurora continua uma atividade entre computador e outra presença, com proteção contextual.

### Golden Proof

```text
private workstation context
→ handoff
→ public/limited device
→ safe summary
→ re-auth
→ authorized continuation
```

## M9 — Laboratory Observation

### Outcome

Aurora recebe telemetria de dispositivos e acompanha protocolo sem physical actuation.

### Golden Proof

```text
identified device
→ authorized telemetry
→ anomaly
→ evidence
→ no effect beyond observe
```

## M10 — Controlled Physical Actuation

### Outcome

Aurora executa ações físicas delimitadas com independent safety interlocks.

### Golden Proof

```text
approved protocol
→ deterministic limits
→ controlled action
→ forced safety condition
→ automatic containment
→ trace and receipt
```

## M11 — Self-Improvement Campaign

### Outcome

Aurora detecta classe de falha, cria candidata, avalia além do incidente e propõe promoção supervisionada.

### Golden Proof

```text
multiple incidents
→ shared root cause
→ sandbox candidate
→ holdout/regression
→ independent review
→ canary
→ rollback proven
```

## M12 — Continuous Engineering Companion

### Outcome

Aurora coordena pesquisa, projeto, construção, teste, observação, documentation e learning em um programa contínuo.

Este marco é direção estratégica, não commitment técnico atual.

## 3. Replan triggers

- pesquisa contradiz premise;
- spike falha;
- protocol muda materialmente;
- security model insuficiente;
- operational burden desproporcional;
- memória/contexto não atinge quality;
- MNFS readiness muda;
- physical safety exige arquitetura separada;
- novo requisito pessoal altera constitutional scope.
