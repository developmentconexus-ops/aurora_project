---
id: ADR-AURORA-0002
title: First-party Harness Development Kit and Universal Conformance
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
  - first-party SDK policy
  - conformance policy
related:
  - DOC-AURORA-BLUEPRINT-05
  - RESEARCH-AURORA-HARNESS-ARCHITECTURE-V1
supersedes: []
superseded_by: null
last_reviewed: 2026-08-05
---

# ADR-0002 — First-party Harness Development Kit and Universal Conformance

## Context

Aurora e Leandro criarão a maioria das harnesses first-party. Reimplementar manifests, lifecycle, telemetry, authority calls, artifacts e recovery manualmente gera inconsistência e dificulta que a própria Aurora construa novas capabilities.

Tornar uma library a única integração possível, porém, aprisiona contracts à linguagem e impede testar a implementação.

## Options

### Runtime e SDK obrigatórios universalmente

Uniforme, porém incompatível com external/legacy providers e múltiplas linguagens.

### SDK opcional para todos

Flexível, mas permite implementações artesanais first-party.

### SDK obrigatório por policy para first-party; contracts e conformance universais

Fornece Golden Path sem transformar SDK em spec.

## Decision

Criar um Aurora Harness Development Kit.

Harnesses first-party **MUST** usar o AHDK, salvo waiver explícito.

External/legacy providers **MAY**:

- usar adapter;
- implementar binding diretamente;
- expor MCP/A2A;
- usar outro SDK.

Toda integração **MUST** passar pela mesma Conformance Suite.

O AHDK não:

- define authority normativa;
- possui state global;
- decide policy;
- substitui sandbox;
- dispensa conformance.

## Consequences

### Positive

- caminho correto mais fácil;
- code generation;
- observability by default;
- menos boilerplate;
- segurança e errors consistentes;
- harness creation por Aurora torna-se viável.

### Negative

- manutenção do kit;
- necessidade de version matrix;
- risco de SDK crescer demais;
- language choice inicial ganha influência.

### Risks

- hidden coupling;
- falsa segurança;
- SDK-first design.

Mitigação: contract-first, reference implementation alternativa no SPK-001 e black-box tests.

## Waiver

Waiver declara:

- razão;
- owner;
- scope;
- risks;
- compensating controls;
- expiration/removal condition;
- conformance result.

## Validation

SPK-001 deve implementar a mesma capability:

1. com AHDK;
2. diretamente por protocol.

Ambas passam pela mesma suite.

## Reconsideration triggers

- AHDK adiciona mais complexidade que remove;
- direct implementations tornam-se padrão;
- uma linguagem domina todos os casos comprovadamente;
- conformance não detecta divergência relevante.
