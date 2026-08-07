# Projeto Aurora

Aurora é uma **inteligência artificial pessoal, persistente, multimodal e agêntica** concebida para funcionar como o sistema operacional cognitivo e o control plane global dos projetos, conhecimentos, capacidades, ferramentas, dispositivos e ambientes de Leandro.

Sua visão de longo prazo é tornar-se uma extensão confiável da capacidade de:

```text
imaginar
→ pesquisar
→ projetar
→ construir
→ testar
→ observar
→ compreender
→ aprender
```

O primeiro domínio operacional profundo é engenharia: software, IA, eletrônica, hardware, firmware, pesquisa, desenvolvimento de produtos e laboratório.

## North Star

> Leandro entra no laboratório, convoca Aurora e continua qualquer projeto exatamente do ponto em que parou. Aurora conhece objetivo, estado, decisões e evidências; reúne conhecimento; coordena capacidades digitais e físicas; acompanha o experimento; identifica riscos e anomalias; registra o que aconteceu; e ajuda a decidir o próximo passo.

## Direção constitucional

- **Leandro-first:** single-user no horizonte atual, sem antecipar multi-tenancy ou complexidade de SaaS.
- **Inteligência pessoal ampla, engenharia primeiro:** novos domínios entram somente como capabilities explicitamente aprovadas.
- **Copiloto intelectual confiável:** Aurora investiga, recomenda e discorda com fundamento; Leandro mantém autoridade final.
- **Personalidade com presença:** identidade estável e transparente; precisão e humor seco combinados com proximidade, curiosidade e entusiasmo.
- **Proatividade contextual:** pode trazer informação material sem comando, respeitando relevância, urgência, confiança e custo da interrupção.
- **Autonomia delegada:** pode conduzir ações, workflows e campanhas dentro de envelopes explícitos de objetivo, autoridade, budget, guardrails e parada.
- **Memória governada:** multiescopo, multitemporal, rastreável, corrigível, supersedível e subordinada a fontes de verdade.
- **Local-first, cloud-assisted:** computação pode ser distribuída; identidade, autoridade e soberania permanecem sob controle de Leandro.
- **Uma Aurora, múltiplas presenças:** computador, celular, wearables, displays e dispositivos futuros manifestam a mesma identidade com capabilities e permissões próprias.
- **Aurora como control plane global:** Harnesses especializadas oferecem capabilities por contratos estáveis.
- **Framework-neutral:** modelos, frameworks, runtimes e protocolos são mecanismos substituíveis, não a constituição do produto.
- **Evidence-driven:** claims, artefatos, observações, receipts, evidence e verdicts são distintos.
- **Aprender continuamente, mudar deliberadamente:** autoaperfeiçoamento exige investigação causal, sandbox, avaliação ampla, revisão independente e rollback.

## Arquitetura conceitual

```text
Leandro / Presenças
        ↓
Aurora Core
├── Identity and World/Project Model
├── Memory and Context Builder
├── Mission and Delegation Control
├── Capability Registry
├── Authority, Policy and Budgets
├── Durable Execution Port
├── Artifact/Evidence Coordination
└── Observability and Evaluation
        ↓
Aurora Contract Model
        ↓
Bindings: Native AHDK | RPC | A2A | MCP | HTTP/gRPC/Events
        ↓
Specialized Harnesses
├── Research
├── Software / MNFS
├── Hardware
├── Firmware
├── Laboratory
├── Evaluation
└── future personal domains
```

Aurora governa objetivo global, contexto, authority, budget, composição e relação com Leandro. Cada Harness governa como produzir seu resultado especializado dentro da Delegation recebida.

## Estado atual

O estado mutável do programa não é duplicado neste README. A fonte canônica para Product Milestone selecionado, gate ACRM atual, blockers, autorizações, proibições e próxima ação é:

```text
docs/tracking/STATUS.md
```

A baseline durável já estabelecida é:

```text
A0 baseline: ACCEPTED / MERGED — 2026-08-06
ADR-0001 / ADR-0002: ACCEPTED
Stack selection by A0: NOT PERFORMED
Runtime implementation: NOT AUTHORIZED BY A0
Architecture Spike execution: NOT AUTHORIZED BY A0
```

A aceitação de A0 estabelece a constituição do produto; cada milestone, gate, decisão técnica, spike e execução posterior continua exigindo sua própria promoção e autorização explícita. Consulte `STATUS.md` em vez de inferir o próximo passo deste resumo.

## Documentação

Comece por:

1. [`AGENTS.md`](AGENTS.md) — bootstrap obrigatório para novas sessões;
2. [`docs/tracking/STATUS.md`](docs/tracking/STATUS.md) — estado e boundary atual;
3. [`docs/DOCUMENTATION-MAP.md`](docs/DOCUMENTATION-MAP.md) — autoridade, ownership e read paths;
4. [`docs/product/README.md`](docs/product/README.md) — índice completo do Product Blueprint;
5. [`docs/product/PRODUCT-BLUEPRINT.md`](docs/product/PRODUCT-BLUEPRINT.md) — publicação agregada gerada;
6. [`docs/product/CAPABILITY-REALIZATION-METHOD.md`](docs/product/CAPABILITY-REALIZATION-METHOD.md) — metodologia R0–R8;
7. [`docs/roadmap.md`](docs/roadmap.md) — Product Milestones e Golden Proofs;
8. [`docs/research/RESEARCH-MAP.md`](docs/research/RESEARCH-MAP.md) — pesquisas e decisões ainda abertas.

Para entender a origem e garantir que nada importante ficou apenas no chat:

- [`docs/history/2026-08-05-aurora-origin-and-discovery-record.md`](docs/history/2026-08-05-aurora-origin-and-discovery-record.md)
- [`docs/tracking/DOCUMENTATION-COVERAGE.md`](docs/tracking/DOCUMENTATION-COVERAGE.md)
- [`docs/reviews/2026-08-05-a0-adversarial-documentation-review.md`](docs/reviews/2026-08-05-a0-adversarial-documentation-review.md)

## Metodologia

```text
Product Blueprint
→ requirements and applicability
→ focused research / Architecture Spikes
→ ADRs
→ Capability Spec
→ Mission Contract
→ Microdesign / Implementation Plan
→ implementation
→ Receipts and Evidence
→ Verdict and Product Milestone Closeout
```

Cada transição possui um gate separado. Pesquisa não autoriza decisão; contrato não autoriza automaticamente implementação; testes locais não fecham milestone; e o próximo milestone nunca é liberado por implicação.

> **A conversa é o ambiente de descoberta. O repositório é a memória canônica do projeto.**
