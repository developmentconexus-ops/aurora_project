---
id: ADR-AURORA-0001
title: Aurora-owned Contract Model and Replaceable Bindings
document_type: adr
form: explanation
authority: decision
status: proposed
version: 0.1.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - ownership of cross-harness semantics
  - protocol binding policy
related:
  - DOC-AURORA-BLUEPRINT-05
  - DOC-AURORA-BLUEPRINT-07
  - RESEARCH-AURORA-HARNESS-ARCHITECTURE-V1
supersedes: []
superseded_by: null
last_reviewed: 2026-08-05
---

# ADR-0001 — Aurora-owned Contract Model and Replaceable Bindings

## Context

Aurora precisa coordenar harnesses que podem usar Mastra, LangGraph, Pi, OpenHands, código determinístico, firmware ou runtimes futuros. MCP e A2A cobrem partes importantes, mas não possuem todos os conceitos de identidade, memória, authority, budget, evidence e acceptance do Aurora.

Escolher um framework ou protocolo como modelo canônico criaria lock-in e forçaria domínios heterogêneos a uma semântica externa.

## Decision drivers

- estabilidade da visão;
- substituição de frameworks;
- interoperability;
- local-first;
- support a harnesses não agênticas;
- segurança;
- evolution;
- conformance;
- YAGNI.

## Options

### Framework-owned domain

Aurora adota o modelo de um framework.

Rejeitado por lock-in e cobertura insuficiente.

### Protocol-owned domain

Aurora usa MCP ou A2A como constituição.

Rejeitado porque são bindings úteis, mas não cobrem o domínio completo.

### Aurora-owned domain with replaceable bindings

Aurora define semantics e contracts; protocols e adapters transportam.

Recomendado.

## Decision

Aurora possuirá um Contract Model independente de linguagem, framework e transport para:

- capabilities;
- providers;
- manifests;
- delegations;
- Context Packs;
- Authority Grants;
- budgets;
- events;
- effects;
- artifacts;
- evidence;
- outcomes;
- recovery.

MCP, A2A, native RPC, HTTP/gRPC e futuros protocolos serão bindings.

Extensão proprietária só será criada após gap analysis e spike.

## Consequences

### Positive

- preserva o produto;
- permite múltiplos runtimes;
- facilita conformance;
- separa semantics e transport;
- permite migração;
- suporta software e hardware.

### Negative

- exige schemas e adapters próprios;
- cria responsibility de versioning;
- requer conformance suite;
- pode duplicar conceitos se boundaries forem mal desenhadas.

### Risks

- overengineering;
- semantic drift;
- adapter explosion.

Mitigação: implementar apenas contratos usados por reference slices, code generation e compatibility checks.

## Validation

Architecture spikes SPK-001, SPK-002, SPK-003 e SPK-008.

## Reconsideration triggers

- padrão aberto cobre integralmente o domínio sem perda de invariantes;
- adapters excedem benefício;
- conformance demonstra incompatibilidade estrutural;
- architecture spike mostra custo desproporcional.
