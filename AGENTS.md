# AGENTS.md

Este arquivo é o bootstrap mínimo para qualquer agente ou nova sessão que trabalhe no Projeto Aurora.

## Ordem obrigatória de leitura

1. `docs/DOCUMENTATION-MAP.md`
2. `docs/tracking/STATUS.md`
3. `docs/roadmap.md`
4. somente as seções do Product Blueprint relacionadas ao trabalho atual;
5. ADRs, pesquisas, specs e designs referenciados pelo escopo.

## Regras de autoridade

- Conversas, issues, backlog e research reports não redefinem arquitetura por conta própria.
- O Product Blueprint governa invariantes constitucionais.
- ADRs aceitos governam decisões específicas dentro dos limites constitucionais.
- Specs governam capabilities reutilizáveis.
- Contratos governam uma execução delimitada.
- Tracking informa estado atual, nunca doutrina.
- Código futuro deve declarar impacto documental.

## Limites da fase atual

```text
Fase: A0 — Product and Architecture Baseline
Implementação do Aurora Core: PROIBIDA
Architecture spikes: PROPOSTOS, NÃO AUTORIZADOS
Escolha de stack: NÃO REALIZADA
Documento em revisão: Capability System and Harness Architecture
```

Não implemente runtime, SDK, protocol binding, banco, UI ou integração com MNFS antes de:

1. a proposta documental ser revisada por Leandro;
2. decisões materiais serem promovidas a ADRs aceitos;
3. existir um plano de implementação aprovado.

## Conduta de pesquisa

- Use fontes primárias e atuais.
- Registre data, versão, limitações e origem.
- Separe fato, inferência, recomendação e decisão.
- Uma tecnologia estudada é candidata, não escolha.
- Diante de informação temporalmente instável, verifique novamente.

## Regra de continuidade

Ao terminar trabalho material:

- atualize `docs/tracking/STATUS.md`;
- registre o que ocorreu em `docs/tracking/WORKLOG.md`;
- atualize `docs/tracking/DECISIONS.md` quando aplicável;
- declare divergências, bloqueios e próxima ação exata;
- não deixe decisões apenas no transcript.
