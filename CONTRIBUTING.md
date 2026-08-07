# Contributing to Projeto Aurora

Projeto Aurora está em fase de arquitetura e documentação.

## Antes de contribuir

Leia:

1. `AGENTS.md`
2. `docs/DOCUMENTATION-MAP.md`
3. `docs/tracking/STATUS.md`
4. os documentos canônicos do escopo.

## Mudanças normativas

- Não altere uma invariante por meio de tracking ou research.
- Decisão específica material exige ADR.
- Research precisa de sources manifest.
- Mudança aceita preserva história e supersession.
- Implementação futura exige design e plano aprovados.

## Pull requests

Toda PR material deve declarar:

```yaml
documentation_impact:
  status: NONE | UPDATED | FOLLOW_UP_REQUIRED
  affected: []
  rationale: ""
  follow_up: null
```

## Estado e branches

`main` é a branch canônica do repositório. Branches não canônicas são superfícies de proposta/revisão e não promovem conteúdo por existência, commit, CI verde ou permissão de escrita.

A branch `docs/architecture-baseline` pode ser reutilizada pelo workflow documental para regeneração e revisão de mudanças constitucionais, mas seu conteúdo só se torna canônico após a autoridade aplicável aceitar a revisão e a mudança ser integrada a `main`.

O gate, as autorizações, os blockers e a próxima ação correntes são sempre lidos de `docs/tracking/STATUS.md`.
