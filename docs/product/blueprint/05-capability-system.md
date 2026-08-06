---
id: DOC-AURORA-BLUEPRINT-05
title: Sistema de Capabilities
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
  - capability model
  - registry model
  - discovery trust authority separation
related:
  - DOC-AURORA-BLUEPRINT-07
  - ADR-AURORA-0001
  - ADR-AURORA-0002
  - RESEARCH-AURORA-HARNESS-ARCHITECTURE-V1
last_reviewed: 2026-08-05
---

# 5. Sistema de Capabilities

## 5.1 Propósito

O Capability System permite que Aurora compreenda **o que pode ser feito**, quais sistemas oferecem essa capacidade, sob quais condições eles podem ser confiados e em qual escopo podem agir.

Aurora não deve codificar cada integração no próprio raciocínio. Capabilities constituem a fronteira estável entre intenção global e execução especializada.

## 5.2 Vocabulário

| Conceito | Responsabilidade |
|---|---|
| Tool | operação delimitada |
| Agent | unidade probabilística de raciocínio e uso de tools |
| Workflow | coordenação de etapas conhecidas ou adaptativas |
| Runtime | ambiente que executa agentes/workflows |
| Harness | sistema especializado completo de um domínio |
| Capability | capacidade externa oferecida ao Aurora |
| Provider | harness ou serviço que oferece uma capability |
| Instance | execução concreta de um provider |
| Adapter | tradução entre contrato Aurora e interface externa |
| Binding | mapeamento para protocolo/transporte |
| Registry | catálogo canônico, confiança e aprovação |
| Delegation | trabalho concreto atribuído a uma capability |

## 5.3 Invariante

> Capability descreve o que pode ser entregue. Harness descreve quem e como oferece. Delegation descreve o trabalho concreto.

Duas harnesses podem oferecer a mesma capability. Uma harness pode oferecer várias capabilities.

## 5.4 Separações obrigatórias

```text
DISCOVERY
→ Aurora sabe que existe

COMPATIBILITY
→ Aurora compreende o contrato

TRUST
→ existe evidência sobre comportamento e identidade

AUTHORITY
→ uso foi permitido em um escopo

EXECUTION
→ uma delegação concreta está ativa
```

Nenhuma etapa implica automaticamente a seguinte.

## 5.5 Capability Manifest

O manifest declara, no mínimo:

### Identidade

- provider ID;
- capability IDs;
- versões;
- publisher;
- source revision;
- build digest;
- instance ID;
- provenance disponível.

### Contratos

- schemas de entrada;
- schemas de saída;
- eventos;
- artefatos;
- errors;
- compatibilidade com o Aurora Contract Model.

### Operação

- síncrona/assíncrona;
- streaming;
- cancelamento;
- resume;
- idempotência;
- limites;
- requisitos de rede, storage e compute.

### Dados e effects

- classes de dados requeridas;
- dados opcionais;
- tratamento de segredos;
- effects declarados;
- effects proibidos;
- ambientes suportados.

### Evidência

- tipos de artefato;
- receipts;
- cobertura de critérios;
- provenance;
- observabilidade.

O manifest é uma declaração, não uma prova.

## 5.6 Capability Registry

O Registry mantém:

- capabilities conhecidas;
- providers e instances;
- manifests por versão;
- bindings suportados;
- resultados de inspeção e conformidade;
- trust assessments;
- approvals por scope;
- incidentes;
- status operacional;
- suspensão, revogação e retirement;
- build provenance.

Entidades:

```text
Capability
→ Provider
→ Instance
→ Manifest
→ Verification
→ Approval
→ Delegation
```

## 5.7 Lifecycle de confiança

Proposta inicial:

```text
DISCOVERED
→ INSPECTED
→ SANDBOX_VALIDATED
→ VERIFIED
→ APPROVED_FOR_SCOPE
→ SUSPENDED | REVOKED | RETIRED
```

### DISCOVERED

Manifest ou endpoint foi encontrado. Execução bloqueada.

### INSPECTED

Identidade, schemas, effects e requisitos foram analisados.

### SANDBOX_VALIDATED

Execução restrita com dados sintéticos ou não sensíveis.

### VERIFIED

Evidência demonstra comportamento conforme o escopo testado.

### APPROVED_FOR_SCOPE

Uso permitido por capability, versão, domínio, data class, effect e environment.

### SUSPENDED / REVOKED / RETIRED

Uso temporariamente bloqueado, confiança removida ou provider aposentado.

## 5.8 Confiança multidimensional

- identidade;
- integridade do build;
- compatibilidade;
- correção funcional;
- segurança;
- recovery;
- evidência;
- custo;
- latência;
- disponibilidade;
- escopo aprovado.

Não existe `trusted: true` universal.

## 5.9 Seleção de provider

Aurora pode escolher entre providers considerando:

```text
capability fit
+ qualidade comprovada
+ data sensitivity
+ authority disponível
+ ambiente
+ custo
+ latência
+ disponibilidade
+ recovery
+ confiança
```

A decisão e seus fatores devem ser observáveis.

## 5.10 Evolução e compatibilidade

- versões materiais não herdam confiança automaticamente;
- breaking changes devem ser detectadas mecanicamente;
- manifests declaram versões de contrato e bindings;
- adapters podem preservar compatibilidade;
- capabilities obsoletas permanecem historicamente descobríveis;
- falha de compatibilidade é fail-closed.

## 5.11 Aurora Harness Development Kit

Harnesses first-party seguem, por política, o Aurora Harness Development Kit (AHDK), salvo waiver explícito.

O AHDK fornece:

- tipos e code generation;
- manifest builder;
- lifecycle API;
- context API;
- authority/effect client;
- artifact/evidence API;
- cancelamento e checkpoints;
- observabilidade automática;
- simulador local;
- mocks;
- fault injection;
- templates;
- Golden Paths;
- integração com a conformance suite.

O AHDK:

- não é a especificação;
- não é a fonte do estado global;
- não decide autorização;
- não é boundary de segurança;
- não torna conformidade opcional.

## 5.12 Conformance Kit

Toda integração, com ou sem SDK, deve passar por testes black-box de:

- schema;
- lifecycle;
- estados;
- erros;
- cancelamento;
- resume;
- idempotência;
- event ordering;
- artifact integrity;
- authority denial;
- effect receipts;
- budget enforcement;
- restart;
- compatibility.

## 5.13 Golden Path de criação

```text
aurora harness create
→ template aprovado
→ contratos gerados
→ instrumentation instalada
→ tests e conformance configurados
→ manifest produzido
→ sandbox local
→ registry candidate
```

Aurora não deve criar uma harness a partir de uma pasta vazia quando um caminho comprovado existe.

## 5.14 Não objetivos atuais

- marketplace público;
- multi-tenant registry;
- linguagem universal de agentes;
- runtime obrigatório;
- três SDKs antes de validar um;
- protocolo proprietário completo;
- trust automático de manifests assinados;
- autoativação de providers externos.
