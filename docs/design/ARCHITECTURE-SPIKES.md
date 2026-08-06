---
id: DESIGN-AURORA-ARCHITECTURE-SPIKES
title: Aurora Architecture Spike Portfolio
document_type: spike_portfolio
form: reference
authority: design
status: proposed
version: 0.1.0
owners:
  - developmentconexus-ops
related:
  - DESIGN-AURORA-CAPABILITY-HARNESS-ARCHITECTURE
  - RESEARCH-AURORA-HARNESS-ARCHITECTURE-V1
last_reviewed: 2026-08-05
---

# Architecture Spike Portfolio

Nenhum spike está autorizado nesta fase. Este documento define perguntas e provas, não implementação.

## SPK-001 — Contract, AHDK and Conformance

### Pergunta

O Contract Model permanece implementável sem o SDK? O AHDK reduz complexidade sem esconder semantics?

### Prova

Mesma `artifact_transform` capability:

- implementation A via AHDK;
- implementation B via direct protocol;
- same black-box suite;
- breaking schema test;
- cancellation;
- invalid transition;
- artifact integrity.

## SPK-002 — MCP Binding

Provar:

- discovery;
- tool/resource;
- JSON Schema;
- input_required;
- Tasks;
- cancellation;
- auth;
- trace propagation;
- official conformance.

## SPK-003 — A2A Binding

Provar:

- Agent Card;
- signed card candidate;
- Task;
- streaming;
- polling;
- artifact;
- input-required;
- cancel;
- reconnect;
- TCK;
- map to Delegation Envelope.

## SPK-004 — Durable Execution

Comparar uma baseline mínima com shortlist.

Golden proof:

```text
start
→ effect idempotent
→ checkpoint
→ kill process
→ restart
→ no duplicate effect
→ wait for input
→ resume
→ complete
```

Critérios:

- local-first;
- operational burden;
- languages;
- timers;
- cancellation;
- versioning;
- observability;
- backup;
- portability.

## SPK-005 — Authority and Effect Gateway

```text
Leandro grants
→ Aurora delegates
→ harness receives short-lived authority
→ allowed effect succeeds
→ denied effect fails
→ revoke during run
→ audit explains
```

## SPK-006 — Distributed Trace

Trace único:

```text
Aurora
→ remote harness
→ MCP tool
→ Effect Gateway
→ Artifact/Evidence
```

Verificar propagation, redaction, cost e cardinality.

## SPK-007 — Registry and Provenance

Registrar exact build, conformance e approval scope. Alterar digest e provar que trust não é herdado.

## SPK-008 — Framework Neutrality

Mesma capability em dois runtimes distintos. Aurora não muda mission/delegation model.

## Exit rule

Spike produz:

- question;
- environment;
- procedure;
- observations;
- artifacts;
- conclusion;
- limitation;
- recommendation;
- ADR impact.

Resultado negativo é evidence válida.
