---
id: DOC-AURORA-METHODOLOGY-REPOSITORY-REBASELINE
title: Aurora Methodology and Repository Rebaseline
document_type: architecture_methodology_decision
form: reference
authority: decision
status: proposed
version: 0.1.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - proposed MR-01 methodology and repository rebaseline decision
  - proposed disposition of the old TA-03+ planning sequence
  - proposed DevelopmentConexus Method and Repository Standard adoption boundary
related:
  - DOC-AURORA-BLUEPRINT-15
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DESIGN-AURORA-TECHNICAL-ARCHITECTURE-BASELINE-MAP
  - DESIGN-AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY
review_triggers:
  - DevelopmentConexus Method material amendment
  - Repository Standard material amendment
  - implementation-readiness graph proves incomplete or ceremonial
  - repository migration cannot preserve current Aurora authority/provenance
  - accepted Product or TA-01/TA-02 contradiction
last_reviewed: 2026-08-23
---

# Aurora Methodology and Repository Rebaseline

> **PROPOSED / NOT YET RATIFIED.** The operator approved the MR-01 design direction on 2026-08-23. This durable candidate still requires exact constitutional reconciliation, migration planning, independent challenge and final operator ratification before it can replace current `main` authority.

## 1. Decision question

How should Projeto Aurora incorporate the engineering and repository discipline that matured across DevelopmentConexus projects so implementation becomes constrained realization of accepted authority rather than a second architecture/design phase performed while coding?

## 2. Proposed decision

Adopt, by reference and without local duplication:

```text
developmentconexus-ops/conexus-methodology/METHOD.md v1.0.0
→ cross-repository engineering reasoning authority

developmentconexus-ops/conexus-methodology/REPOSITORY-STANDARD.md v1.0.0
→ repository operating-envelope authority
```

Preserve Aurora-owned Product and architecture authority. Repository standards do not define Aurora Product meaning, system ownership, technology or stage semantics.

Refine Aurora-specific methods into non-overlapping layers:

```text
DevelopmentConexus Engineering Method
→ reasoning and decision quality

DevelopmentConexus Repository Standard
→ repository navigation, authority routing, temporary work, Git/review/verification envelope

Aurora Product / Architecture
→ Aurora meaning, invariants, owners and target boundaries

Aurora Planning & Implementation-Readiness
→ cross-system prerequisites that must be closed before Product execution

specialized methods
→ frontend, research, qualification, transition/cutover only when a named consumer triggers them

Aurora Capability Realization Method (ACRM)
→ capability/slice applicability, requirements, design/contract, execution Evidence and milestone closeout

Implementation Program / Execution Graph
→ bounded executable slices after architecture/readiness closure
```

## 3. Target invariant

> Product implementation MUST NOT be the stage where material Product, architecture, contract, security, data, runtime or user-interaction decisions are invented by convenience.

When a required material decision is missing:

```text
STOP
→ identify the smallest owning authority
→ reopen only that authority
→ recompile affected downstream artifacts
→ resume only after the missing authority is accepted
```

Do not substitute a plausible implementation default.

## 4. Preserve / refine / reopen dispositions

| Current Aurora surface | Disposition | Proposed treatment |
| --- | --- | --- |
| Product Blueprint 01–14 Product meaning | PRESERVE | no Product redesign by MR-01 |
| Blueprint 15 principles | PRESERVE | one owner, repository-as-memory, research≠authority, minimal sufficient context, gated execution remain binding |
| Blueprint 15 repository/status/layout clauses | BOUNDED REOPEN | reconcile to Repository Standard v1 without changing unrelated Product meaning |
| TA-01 | PRESERVE | accepted canonical module/ownership baseline |
| TA-02 | PRESERVE | accepted canonical process/runtime topology baseline |
| old TA-03+ sequence | REOPEN / SUPERSEDE IF RATIFIED | replace with dependency order in §5 |
| ACRM R0–R8 | PRESERVE + REFINE SCOPE | capability/slice realization and Evidence lifecycle; not repository operating model |
| M0 R7 candidate | PRESERVE AS EVIDENCE | frozen / non-canonical; no Verdict or R8 implication |
| permanent `docs/superpowers/**` | RETIRE AFTER CONSOLIDATION | temporary plans/work belong in branch-only `docs/work/**`; durable semantics move to owning docs |
| `docs/tracking/STATUS.md` as mutable authority | SUPERSEDE AFTER MIGRATION | `docs/roadmap.md` becomes sole mutable stage/status/allowed-work/next-action authority |
| MNFS as current software-Harness name | REFINE | current name is Conexus OS; historical artifacts may retain MNFS for provenance |

## 5. Proposed Aurora pre-implementation planning graph

```text
TA-01  Logical Modules & Canonical Ownership                CLOSED / PRESERVE
TA-02  Process, Runtime & Evolutionary Topology              CLOSED / PRESERVE
TA-03  Cross-System Operation Surface                       NEXT AFTER MR-01 MIGRATION
TA-04  Executable Contract Model & Wire
TA-05  Human Interaction & Presence Realization
TA-06  State, Data, Memory, Knowledge, Artifact & Evidence
TA-07  Identity, Authentication, Policy, Secrets & Effects
TA-08  Cognitive Runtime, Models, Context & Harnesses
TA-09  Paved Road, Repository, Source & Build Realization
TA-10  Runtime, Deployment & Operations
TA-TX  Transition / Cutover                                 CONDITIONAL
TA-11  Whole-System Coherence & Golden Flows
TA-12  Implementation Program & Execution Graph
TA-13  Adversarial Implementation Readiness
        ↓
explicit operator Product execution grant
        ↓
Product implementation eligible
```

The dependency order is binding when one stage owns assumptions needed by a later stage. Ceremony is not binding: adjacent work may share one coherent gate when its owners and proof remain explicit, and a stage with no real consumer produces no empty artifact.

## 6. Stage intent

### TA-03 — Cross-System Operation Surface

Define only material operations crossing owner/process/actor/effect boundaries. Each admitted operation names caller/consumer class, semantic owner, input/output meaning, authority, state consequence, failure/ambiguity, concurrency/idempotency where relevant, audit/Evidence implications and explicit non-owner boundaries.

TA-03 is not an endpoint inventory and does not duplicate internal Capability Spec methods.

### TA-04 — Executable Contract Model & Wire

Project accepted semantics into exact executable contract families: schemas/types, error taxonomy, compatibility/versioning, deadlines/cancellation, correlation/causation, idempotency/reconciliation, streaming/event profiles and generated projections. Concrete bindings are selected per boundary only after requirements are known.

### TA-05 — Human Interaction & Presence Realization

Prove that material Leandro-facing interactions, session/intent progression, progress/decision/evidence presentation and failure/recovery paths have accepted homes without creating UI-owned Product authority. The reusable frontend planning method is activated only when a material visual/browser consumer exists.

### TA-06 — State, Data, Memory, Knowledge, Artifact & Evidence

Classify canonical/derived/ephemeral/provider-local data, owner/writer rules, consistency/freshness, retention/deletion, backup/export/restore, query/volume/security needs and rebuildability before selecting physical stores.

### TA-07 — Identity, Authentication, Policy, Secrets & Effects

Define actor identity, authentication proof, sessions/tokens/credentials, authority/delegation, policy decisions, Effect Gateway enforcement, credential last-mile, revocation/recovery and audit before selecting IAM/policy/secret products.

### TA-08 — Cognitive Runtime, Models, Context & Harnesses

Close deterministic-versus-model responsibilities, Context Builder ownership, model/provider substitution, memory interaction, tool/effect path, Delegation lifecycle, provider-local reconciliation and the exact Mastra/Conexus OS boundary for real consumers.

### TA-09 — Paved Road, Repository, Source & Build Realization

Only after the preceding consumers are known, choose repository/source/build topology, language/workspace placement, generated-contract custody, dependency enforcement, scaffold/profile, code-generation, CI layers, versioning/release and developer/agent Paved Road.

### TA-10 — Runtime, Deployment & Operations

Select startup/readiness/shutdown, supervision, configuration, secret injection, packaging/install, observability, migrations, update/rollback, backup/recovery, Stage A/B operation and supply-chain mechanisms.

### TA-TX — Transition / Cutover

Open only when Aurora has a real current installation/state/runtime that must survive migration. Do not manufacture compatibility/cutover machinery while no consumer exists.

### TA-11 — Whole-System Coherence & Golden Flows

Compose a smallest irreducible set of positive/negative/recovery flows selected by defect class: removing one flow must leave a material accepted invariant without a representative falsifier.

### TA-12 — Implementation Program & Execution Graph

Compile accepted authority into bounded slices with exact owners, operations, contracts, code-ownership classes, dependencies, allowed state/config/migration surfaces, proof obligations, stop/reopen triggers and completion criteria.

### TA-13 — Adversarial Implementation Readiness

Attack the complete implementation envelope for hidden material decisions, duplicate/missing owners, consumer/operation gaps, parallel contract authorities, screen-shaped APIs, unowned persistence, framework leakage, unauthorized bypasses, unverifiable recovery and slices that still require architecture invention.

## 7. Realization decision protocol

For every material mechanism/technology decision:

```text
protected property / real consumer
→ current Aurora authority
→ normative standard where applicable
→ current official exact-version documentation/source/security evidence
→ credible alternatives
→ ADOPT | ADAPT | BUILD | DEFER | STOP
→ proof strategy before implementation
→ exact owner / replacement and reopen triggers
```

A framework capability does not become Aurora capability merely because it exists.

## 8. Code ownership classes for the future Paved Road

```text
GENERATED
→ reproducible projection from accepted machine-readable authority; never hand-owned divergent semantics

AURORA-FOUNDATION
→ versioned Aurora-owned infrastructure/seam protecting repeated cross-cutting invariants; implementation slices consume it but may not silently weaken/redefine it

MODULE-OWNED
→ Product/domain logic owned by the exact Aurora module/capability; legitimate primary coding surface
```

`AURORA-FOUNDATION` is deliberately not named `AURORA-CONTRACT` so it cannot be confused with G01 Aurora Contract Model semantics.

## 9. Repository operating target

The repository target follows the Repository Standard:

```text
README.md                       landing only
AGENTS.md                       compact bootstrap / local hard stops
docs/index.md                   task/intention router
docs/roadmap.md                 sole mutable stage/status/allowed-work/next-action authority
1–2 routed owning docs          normal task-specific authority
```

Normal fresh-actor work targets five files or fewer. Research, Evidence, Git history, old phase chronology and qualification are opt-in by concrete need.

Target semantic surfaces are created only for real consumers:

```text
docs/product/
docs/architecture/
docs/decisions/
docs/phases/
docs/development/
docs/capabilities/
docs/reference/
docs/research/
docs/evidence/
docs/diagrams/
docs/work/        branch-only temporary
```

The exact migration of current Aurora paths is a separate ratification prerequisite; no live authority is deleted until surviving semantics and required provenance have a destination.

## 10. Conexus OS terminology boundary

Current authoritative references to the software-development Harness use **Conexus OS**. Historical evidence may retain **MNFS** when that was the name at the time.

The architectural boundary is unchanged:

```text
Conexus OS may plan/build/test/review/package Evidence for Aurora within an authorized Delegation
≠
Conexus OS owns Aurora Product meaning, sovereign state, authority or required runtime availability
```

## 11. Explicit non-decisions

MR-01 does not select:

- monorepo/polyrepo;
- universal language/runtime;
- concrete Mastra version/integration;
- API/RPC/event binding;
- database/vector/object/telemetry store;
- IAM/policy/secrets product;
- service supervisor/container/orchestrator;
- model/provider;
- Voice stack;
- first AHDK language;
- Product code structure beyond future ownership-class law.

## 12. Ratification prerequisites

Before this decision may become current authority:

1. exact Blueprint 15 constitutional amendment is prepared;
2. repository migration plan covers current live surfaces and provenance;
3. generator/validator/CI migration is specified, including negative controls;
4. fresh-session target route is mechanically testable;
5. independent challenger reviews the exact candidate;
6. material findings are adjudicated;
7. operator explicitly ratifies the final candidate;
8. merge authorization remains separate.

## 13. Reopen triggers

Reopen this decision only when material Evidence shows:

- stage ordering forces a downstream decision before its owner exists;
- ACRM and global readiness duplicate authority in practice;
- repository routing cannot keep a normal task inside the bounded context target;
- implementation still requires repeated material architecture invention;
- a required stage has no recurring consumer and creates systematic ceremony;
- a named new Aurora boundary materially changes the dependency graph;
- the organizational Method/Repository Standard changes in a way that conflicts with Aurora constraints.
