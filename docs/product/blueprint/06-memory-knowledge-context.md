---
id: DOC-AURORA-BLUEPRINT-06
title: Memória, Conhecimento e Contexto
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
  - memory principles
  - context boundaries
  - memory promotion policy
related:
  - DOC-AURORA-BLUEPRINT-01
  - DOC-AURORA-BLUEPRINT-10
last_reviewed: 2026-08-05
---

# 6. Memória, Conhecimento e Contexto

## 6.1 Princípio

> A inteligência observável do Aurora dependerá tanto da qualidade do contexto construído quanto do modelo que raciocina.

Memória não é um banco homogêneo nem o transcript inteiro. Aurora precisa de mecanismos diferentes para funções diferentes.

## 6.2 Distinções

| Conceito | Significado |
|---|---|
| Histórico | registro bruto do que aconteceu |
| Memória | informação preservada para uso futuro |
| Conhecimento | fontes e modelos sobre o mundo e domínios |
| Estado operacional | realidade atual de execuções e dispositivos |
| Contexto ativo | seleção limitada usada na decisão presente |
| Fonte de verdade | autoridade que governa um fato ou decisão |

> Memória orienta o raciocínio; autoridade, evidência e estado ao vivo determinam a verdade operacional.

## 6.3 Escopos de memória

### Memória de trabalho

Objetivo, mensagens recentes, ferramentas, hipótese e próximo passo da atividade atual.

### Memória conversacional

Continuidade de uma conversa específica, sem exigir carregamento integral.

### Memória observacional

Observações temporais densas derivadas de conversas, ações e resultados.

### Memória episódica

Eventos e experiências: teste, falha, campanha, decisão, incidente ou sessão de laboratório.

### Memória de projeto

Visão, estado, arquitetura, decisões, hipóteses, roadmap, tarefas, bloqueios, riscos, dispositivos, artefatos e próximos passos.

### Memória global pessoal

Identidade, competências, objetivos, preferências estáveis e padrões de trabalho de Leandro.

### Memória relacional

Relações entre pessoa, projeto, harness, dispositivo, artefato, decisão e incidente.

### Memória procedural

Protocolos e Golden Paths comprovados.

### Memória de falhas e aprendizagem

Sintomas, hipóteses causais, correções, regressões e abordagens que não devem ser repetidas.

### Memória operacional

Checkpoints, budgets, grants, execuções e dispositivos. Deve existir em estado estruturado, não apenas em linguagem natural.

## 6.4 Metadados mínimos

Uma memória material preserva conceitualmente:

```text
identity
type
scope
content
provenance
created_at / observed_at
epistemic status
confidence
authority
sensitivity
validity interval
supersession
relations
retention policy
```

## 6.5 Status epistêmico

Aurora diferencia:

- Leandro afirmou;
- Leandro aprovou;
- sistema observou;
- instrumento mediu;
- documento estabelece;
- Aurora inferiu;
- fonte externa reporta;
- hipótese está em teste.

Essas categorias não recebem a mesma autoridade.

## 6.6 Promoção por risco e alcance

### Promoção automática permitida

- estado de tarefa;
- eventos de projeto;
- resultados ligados a evidência;
- decisão explícita;
- hipótese marcada;
- autorização e expiração;
- correção factual explícita;
- preferência local de projeto.

### Promoção condicionada

- preferência global;
- inferência pessoal;
- informação sensível;
- padrão comportamental;
- regra cross-project;
- mudança constitucional;
- autoridade permanente.

Princípio:

> Quanto maior alcance, sensibilidade ou autoridade, maior o requisito de promoção.

## 6.7 Lifecycle

```text
OBSERVED
→ CANDIDATE
→ ACCEPTED
→ CONFIRMED
→ SUPERSEDED | EXPIRED | REMOVED
```

Contradição não é resolvida por overwrite silencioso.

## 6.8 Context Builder

O Context Builder seleciona:

- escopo;
- autoridade;
- recência;
- relevância;
- sensibilidade;
- orçamento de tokens;
- fonte original;
- necessidade de leitura ao vivo.

Ele deve:

- priorizar fontes canônicas;
- identificar supersession;
- excluir memória irrelevante;
- evitar contaminação entre projetos;
- explicar proveniência;
- consultar estado atual quando a memória pode estar stale.

## 6.9 Esquecimento e exclusão

Aurora precisa:

- expirar detalhes transitórios;
- resumir ruído;
- marcar stale;
- superseder;
- preservar histórico fora do contexto normal;
- apagar por ordem de Leandro;
- respeitar retenção;
- impedir uso em escopo proibido.

Guardar mais não é lembrar melhor.

## 6.10 Soberania

Aurora é local-first e cloud-assisted.

- memória canônica, identidade, policies e audit ficam sob controle de Leandro;
- provedores externos recebem contexto mínimo autorizado;
- segredos não entram diretamente no prompt;
- toda transferência registra provider, finalidade, classes de dados e resultado;
- exclusão, exportação, migração e revogação são requisitos.

## 6.11 Arquitetura técnica aberta

Candidatos futuros incluem:

- compaction;
- observational memory;
- RAG;
- busca textual;
- embeddings;
- storage estruturado;
- event log;
- knowledge graph;
- contextos hierárquicos;
- políticas aprendidas.

Nenhuma abordagem está escolhida. A capability exigirá pesquisa, protótipos e evals específicos.
