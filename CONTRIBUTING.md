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

## Estado desta baseline

A branch `docs/architecture-baseline` é uma proposta para revisão do operador. Nenhum arquivo nela se torna aceito apenas por existir.
