---
id: DOC-AURORA-DOCUMENTATION-MAP
title: Aurora Documentation Map
document_type: documentation_map
form: reference
authority: constitutional
status: proposed
version: 0.1.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - documentation discovery
  - authority hierarchy
  - canonical read paths
related:
  - DOC-AURORA-PRODUCT-INDEX
  - DOC-AURORA-STATUS
last_reviewed: 2026-08-05
---

# Aurora Documentation Map

## 1. Fase atual

```text
A0 — Product and Architecture Baseline
```

Objetivo: transformar a visão aprovada em documentação canônica, registrar a base de pesquisa e submeter a arquitetura de capabilities e harnesses à revisão antes de qualquer implementação.

## 2. Princípio de governança

> Um conceito durável possui uma fonte proprietária. Outros documentos podem resumir, explicar ou aplicar; não podem redefinir silenciosamente.

## 3. Classes de autoridade

| Classe | Tipo | Governa |
|---|---|---|
| A0 | Constitutional | identidade do produto, invariantes, limites e modelo de autoridade |
| A1 | Decision | uma escolha específica, alternativas, consequências e supersession |
| A2 | Specification | design reutilizável de capability, protocolo, schema ou componente |
| A3 | Contract | compromisso delimitado de missão, delegação, ambiente ou API |
| A4 | Standard / Policy | regra aplicável, Golden Path, enforcement e exceções |
| A5 | Reference | descrição exata da machinery atual |
| A6 | Guidance | tutorial, how-to, runbook e orientação operacional |
| A7 | Evidence | resultado observado, teste, benchmark, acceptance e provenance |
| A8 | Tracking | coordenação e estado presente |
| A9 | Research / Historical | investigação, comparação, proposta rejeitada e histórico |
| A10 | Generated Projection | publicação derivada sem autoridade independente |

## 4. Precedência

Quando houver conflito:

```text
Constituição atual
→ ADR específico aceito, desde que constitucionalmente compatível
→ Capability/System Spec aceita
→ Contrato aprovado para o escopo
→ Standard/Policy/Golden Path
→ Reference atual
→ Guidance
→ Tracking
→ Research/Historical
→ Generated Projection segue sua fonte
```

Um ADR ordinário não pode alterar uma invariante constitucional. Uma mudança constitucional exige proposta explícita, revisão de impacto e atualização do Blueprint.

Conflitos materiais produzem `DOCUMENTATION_DIVERGENCE`; não se escolhe silenciosamente o primeiro arquivo lido.

## 5. Fontes canônicas por assunto

| Assunto | Fonte proprietária |
|---|---|
| visão e princípios | `docs/product/blueprint/01-product-vision.md` |
| relação Leandro–Aurora | `docs/product/blueprint/02-human-aurora-relationship.md` |
| sistema de capabilities | `docs/product/blueprint/05-capability-system.md` |
| memória e contexto | `docs/product/blueprint/06-memory-knowledge-context.md` |
| orquestração de harnesses | `docs/product/blueprint/07-harness-orchestration.md` |
| autonomia, autoridade e segurança | `docs/product/blueprint/10-autonomy-authority-safety.md` |
| sequência de evolução | `docs/roadmap.md` e futura seção 14 |
| governança documental e pesquisa | `docs/product/blueprint/15-documentation-research-governance.md` |
| decisão específica | `docs/adr/NNNN-*.md` |
| pesquisa atual | `docs/research/` |
| design em revisão | `docs/superpowers/specs/` e `docs/design/` |
| estado presente | `docs/tracking/STATUS.md` |
| histórico de trabalho | `docs/tracking/WORKLOG.md` |

## 6. Storage boundaries

### Git

Conhecimento humano canônico:

- Product Blueprint;
- ADRs;
- specs;
- roadmap;
- standards;
- Golden Paths;
- research reports;
- designs;
- evidence selecionada;
- tracking.

### Runtime futuro

Estado operacional não deverá ser inferido apenas de Markdown. Um futuro storage operacional possuirá:

- execuções;
- delegações;
- grants;
- eventos;
- checkpoints;
- budgets;
- artefatos;
- receipts;
- findings;
- estado de dispositivos.

### Artifact Store futuro

Logs, traces, outputs, datasets, screenshots, medições e artefatos volumosos. Quando uma evidência precisar ser preservada pelo repositório, será promovida por referência verificável.

### GitHub

- issue: discussão e tracking;
- PR: proposta, revisão e CI;
- arquivos merged: resultado canônico.

## 7. Entrypoints

| Caminho | Público | Propósito |
|---|---|---|
| `README.md` | humanos | visão resumida e orientação |
| `AGENTS.md` | agentes | bootstrap, autoridade e limites |
| `docs/DOCUMENTATION-MAP.md` | ambos | mapa canônico e ordem de leitura |
| `docs/product/README.md` | arquitetura | índice do Blueprint |
| `docs/roadmap.md` | operador | evolução orientada a capacidades |
| `docs/tracking/STATUS.md` | ambos | situação e próxima ação |
| `docs/research/RESEARCH-MAP.md` | arquitetura | pesquisas, validade e lacunas |

## 8. Read paths

### Nova sessão

```text
README
→ DOCUMENTATION-MAP
→ STATUS
→ roadmap
→ seções relevantes do Blueprint
→ ADRs/Research/Design do escopo
```

### Pesquisa

```text
RESEARCH-MAP
→ pergunta e escopo
→ report
→ sources manifest
→ implications
→ ADR, se houver decisão
```

### Implementação futura

```text
AGENTS
→ STATUS
→ accepted design/spec
→ accepted ADRs
→ approved implementation plan
→ code
→ evidence
→ documentation impact
```

## 9. Estados documentais

```text
DRAFT
→ PROPOSED
→ ACCEPTED
→ SUPERSEDED | REJECTED | WITHDRAWN
```

Tracking usa `CURRENT` ou `ARCHIVED`. Research pode usar `CURRENT`, `STALE` ou `HISTORICAL`.

## 10. Regras de escrita

- IDs são estáveis e únicos.
- Todo documento normativo declara owner, authority, status e source-of-truth scope.
- Research sempre declara limitações e não é normativo.
- Projeções geradas não são editadas diretamente.
- Não criar diretórios vazios por estética.
- Mudança material registra impacto documental.
- Informação temporalmente instável inclui data de verificação.
