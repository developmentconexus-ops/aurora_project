---
id: DOC-AURORA-BLUEPRINT-15
title: Governança Documental e de Pesquisa
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
  - documentation lifecycle
  - research lifecycle
  - source-of-truth protocol
related:
  - DOC-AURORA-DOCUMENTATION-MAP
  - DOC-AURORA-RESEARCH-MAP
last_reviewed: 2026-08-05
---

# 15. Governança Documental e de Pesquisa

## 15.1 Propósito

Impedir que:

- conversa vire arquitetura implícita;
- research vire decisão;
- status vire doutrina;
- código altere o produto sem declaração;
- documentos contraditórios governem o mesmo conceito;
- conhecimento rejeitado desapareça;
- sessões precisem reconstruir o projeto;
- ferramentas atuais definam invariantes permanentes.

## 15.2 Propriedades obrigatórias

Documento canônico possui:

- identity;
- authority;
- lifecycle;
- ownership;
- source-of-truth scope;
- relations;
- review triggers;
- validation.

## 15.3 Pesquisa

Research report registra:

- pergunta;
- escopo;
- data;
- fontes primárias;
- versões;
- método;
- achados;
- limitações;
- divergências;
- implicações;
- candidatos a spike;
- decisões que pode informar.

Research é A9 e não governa implementação.

## 15.4 Manifest de fontes

Toda pesquisa material possui `.sources.json` com:

- source ID;
- title;
- URL;
- publisher;
- type;
- accessed date;
- version/date quando aplicável;
- claims supported;
- notes.

## 15.5 Promoção de pesquisa para decisão

```text
Research
→ alternatives
→ recommendation
→ architecture spike, quando necessário
→ ADR proposed
→ review
→ ADR accepted/rejected
```

Não se cria ADR apenas para registrar popularidade.

## 15.6 Modularidade

O Blueprint é modular. Um agregado gerado somente será criado quando houver generator e check de freshness.

Não criar arquivo generated manualmente.

## 15.7 Tracking

`STATUS.md` declara:

- fase;
- documentos aceitos/propostos;
- autorização;
- bloqueios;
- branch/PR;
- verificação;
- próxima ação.

`WORKLOG.md` preserva história.  
`DECISIONS.md` indexa decisões.  
`BACKLOG.md` captura ideias sem compromisso.

## 15.8 Supersession

Documentos aceitos nunca são reescritos de forma a apagar a história. Mudança material:

- incrementa versão;
- registra rationale;
- liga supersedes/superseded_by;
- atualiza relações;
- preserva Git history;
- reavalia documentos dependentes.

## 15.9 Documentation impact

Toda mudança material futura declara:

```yaml
documentation_impact:
  status: NONE | UPDATED | FOLLOW_UP_REQUIRED
  affected: []
  rationale: ""
  follow_up: null
```

`NONE` exige justificativa específica.

## 15.10 Validações futuras

- Markdown;
- links;
- frontmatter;
- IDs únicos;
- relações;
- owners;
- allowed statuses;
- ADR index;
- supersession;
- source manifests;
- stale research dates;
- generated-file freshness;
- placeholder scan;
- accepted-doc ambiguity scan.

## 15.11 Regra de continuidade

Uma nova sessão, usando somente o repositório, deve:

1. explicar o que Aurora é;
2. identificar a North Star;
3. distinguir visão e compromisso;
4. localizar decisões;
5. reconhecer pesquisa aberta;
6. dizer a fase atual;
7. respeitar autorizações;
8. continuar pela próxima ação exata.

Essa é a Golden Proof documental de A0.
