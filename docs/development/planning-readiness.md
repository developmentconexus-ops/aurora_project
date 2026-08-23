---
id: DOC-AURORA-PLANNING-READINESS
title: Aurora Planning and Implementation-Readiness Standard
document_type: planning_readiness_standard
form: reference
authority: standard
status: proposed
version: 0.1.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - proposed Aurora cross-system planning and implementation-readiness lifecycle
  - proposed relationship between global readiness and ACRM
  - proposed technology-research and Paved-Road decision law
related:
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DOC-AURORA-BLUEPRINT-15
  - DOC-AURORA-METHODOLOGY-REPOSITORY-REBASELINE
  - DESIGN-AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY
review_triggers:
  - planning stage dependency changes
  - ACRM lifecycle changes
  - implementation agents repeatedly invent material cross-system decisions
  - Paved Road ownership classes prove insufficient
last_reviewed: 2026-08-23
---

# Aurora Planning and Implementation-Readiness Standard

> **PROPOSED.** This standard becomes active only after MR-01 final ratification and canonical integration.

## 1. Purpose

Aurora is a long-horizon system of systems. Correct local Capability design is necessary but insufficient when implementation can still encounter unresolved cross-system architecture, operation, contract, data, security, cognitive-runtime, repository or deployment decisions.

This standard defines the cross-system planning path that must reduce those material degrees of freedom before Product implementation begins.

The target property is:

> A future implementation actor should spend its architectural judgment on local reversible realization choices, not on deciding Aurora's missing owners, contracts, trust boundaries, stores, protocols, foundational stack or user-facing semantics while coding.

## 2. Governing authorities

This standard specializes, but does not copy or override:

```text
DevelopmentConexus Engineering Method v1.0.0
→ engineering reasoning and material-decision law

DevelopmentConexus Repository Standard v1.0.0
→ repository operating envelope

Aurora Product Blueprint
→ Product meaning and constitutional invariants

accepted Aurora architecture / decisions
→ exact current structural and technical constraints
```

If this standard conflicts with accepted Product/architecture authority, the conflict is a Finding. This standard cannot silently repair Product meaning.

## 3. Relationship to ACRM

Global readiness and ACRM answer different questions.

```text
GLOBAL PLANNING / READINESS
What cross-system structure must exist so a bounded capability can be realized without inventing foundational architecture?

ACRM R0–R8
How does one selected Product Milestone / Capability / Mission move from accepted intent through applicability, requirements, decisions, contract, implementation design, execution Evidence and closeout?
```

A capability consumes current cross-system authority through R0–R6. It does not replay every global TA stage. Conversely, a global TA stage must not become a second Capability Spec for every internal method.

## 4. Global stage graph

```text
TA-01  Logical Modules & Canonical Ownership
TA-02  Process, Runtime & Evolutionary Topology
TA-03  Cross-System Operation Surface
TA-04  Executable Contract Model & Wire
TA-05  Human Interaction & Presence Realization
TA-06  State, Data, Memory, Knowledge, Artifact & Evidence
TA-07  Identity, Authentication, Policy, Secrets & Effects
TA-08  Cognitive Runtime, Models, Context & Harnesses
TA-09  Paved Road, Repository, Source & Build Realization
TA-10  Runtime, Deployment & Operations
TA-TX  Transition / Cutover — conditional
TA-11  Whole-System Coherence & Golden Flows
TA-12  Implementation Program & Execution Graph
TA-13  Adversarial Implementation Readiness
→ explicit operator Product execution grant
```

TA-01 and TA-02 are already accepted/canonical at the time of this proposal and are not reopened by naming this graph.

## 5. Tailoring law

The graph is dependency-ordered, not ceremony-maximizing.

For every stage:

```text
real consumer / protected property exists
→ do the smallest work that closes the material ambiguity

no current consumer and seam can be added later safely
→ DEFER SAFELY with owner + trigger

adjacent questions share one coherent authority/proof boundary
→ one gate/PR may close them together when explicitly scoped

later Capability consumes accepted global authority
→ do not replay project inception
```

No empty artifact is required merely to preserve numbering.

## 6. Stage exit law

A stage may close only when:

- its semantic owner and non-owner boundaries are explicit;
- material Known/Inferred/Unknown/Deferred states are honest;
- credible alternatives were compared when real ambiguity existed;
- the preferred candidate survived proportional adversarial challenge;
- a proof/falsifier strategy exists for every material runtime/external claim deferred to later execution;
- downstream consumers can proceed without silently inventing a decision owned here;
- reopen triggers are explicit.

## 7. Technology and mechanism selection

Material realization decisions follow:

```text
protected property / failure class / real consumer
→ current Aurora authority
→ normative standard when one exists
→ current official exact-version docs/source/security advisories when version-sensitive
→ credible alternatives
→ operational/replacement/supply-chain cost
→ ADOPT | ADAPT | BUILD | DEFER | STOP
→ proof before implementation
```

Definitions:

- `ADOPT` — use a standard/native/mature mechanism without moving Aurora semantic authority;
- `ADAPT` — use a proven mechanism behind a narrow Aurora-owned boundary;
- `BUILD` — custom implementation only after documenting the exact unresolved gap;
- `DEFER` — no current consumer requires the mechanism and later addition does not dismantle current authority;
- `STOP` — correctness depends on an unproven prerequisite.

Popularity, framework availability, legacy code shape or another repository's choice are not sufficient reasons.

## 8. Research law

Research is decision-linked and bounded.

Use the strongest source required by the claim:

1. current Aurora authority/current executable fact;
2. normative standard/specification;
3. official exact-version documentation;
4. official repository/source/release/security advisory;
5. mature first-party/production reference where materially comparable;
6. community sources for discovery/corroboration only.

Stop researching when the property/constraints, credible alternatives, material failure modes and proportional falsifier are clear enough to decide. Reference collection without decision impact is research theater.

## 9. Proof law

Proof strength matches the claim.

```text
pure/local semantic claim      → static/contract/property/negative analysis
schema/type/codegen claim      → executable generation/compatibility proof
DB/concurrency claim           → real selected DB/transaction proof
browser/session claim          → real browser/server integration where claimed
external provider claim        → controlled real dependency Evidence when required
restart/recovery claim         → restart/loss/ambiguity proof at the owning stage
production claim               → production-stage Evidence only
```

A control counts only when it can be demonstrated capable of firing.

## 10. Cross-system operation admission

TA-03 admits an operation only when it crosses a material boundary such as:

- semantic owner boundary;
- process/runtime/provider boundary;
- actor/Presence boundary;
- effect/credential boundary;
- external-system boundary;
- durable cross-owner state consequence.

Do not create a global operation entry for ordinary private methods inside one owner.

Each admitted operation records, as applicable:

```text
identity / name
caller / consumer class
semantic owner
read | command | proposal | effect | callback/observation class
input/output semantic contract
authority/preconditions
state consequence
failure / ambiguity
concurrency / idempotency
cancellation / deadline
audit / exact-history implication
Evidence / receipt implication
explicit non-owner / forbidden shortcut
```

## 11. Executable contract law

TA-04 turns accepted semantics into one exact source per machine-readable meaning.

It must separate:

```text
semantic contract version
schema version
binding/protocol version
SDK/generated projection version
provider capability/build identity
```

Generated client/server/provider projections never become parallel semantic authorities.

## 12. Human/Presence realization law

A human-facing interaction must trace bidirectionally:

```text
accepted Aurora meaning/operation
→ user/Presence interaction home

interaction/control
→ admitted operation/read truth
→ semantic owner
→ accepted Product meaning
```

Presence/UI does not gain business/system authority from visibility or convenience.

When a real visual Product surface enters scope, the project may deliberately adopt the current reusable Frontend Product Experience Planning Method as a specialized method. It is not automatically imported as Aurora authority before that consumer exists.

## 13. Data/store selection law

No store is selected until logical data classes and guarantees exist.

At minimum distinguish:

```text
canonical operational state
governed durable memory
knowledge/document sources
interaction/exact history
artifacts
Evidence/receipts
audit
telemetry
derived indexes
ephemeral active context
provider-local runtime state
secret/credential references
```

Each physical store requires a named role, owner, backup/restore/deletion story and migration/replacement path.

## 14. Paved Road law

TA-09 defines the smallest default realization path that removes repeated foundational invention from coding agents.

Three ownership classes are canonical candidates:

```text
GENERATED
AURORA-FOUNDATION
MODULE-OWNED
```

### GENERATED

Reproducible from accepted machine-readable authority. Hand-owned divergent semantics are forbidden.

### AURORA-FOUNDATION

Versioned Aurora-owned infrastructure/seams protecting repeated cross-cutting properties. Consuming implementation slices may use but may not silently weaken or redefine them.

### MODULE-OWNED

Code and configuration whose Product/domain semantics belong to one exact Aurora owner/capability. This is the primary legitimate implementation surface.

A scaffold is **infrastructure-rich / Product-feature-poor**: it contains only mechanisms with real repeated consumers/protected properties.

## 15. Escape-hatch law

A Paved-Road escape hatch is explicit, never an undocumented bypass.

It must name:

- the unmet current property/consumer;
- why the default road is insufficient;
- protected Aurora invariants that remain non-degradable;
- smallest additional mechanism;
- proof/falsifier;
- removal/rejoin trigger;
- resulting Baseline/version impact.

## 16. Golden-flow law

TA-11 selects representative flows by defect class, not operation count or module count.

A flow earns a place only when removing it would leave a material accepted invariant without a representative cross-boundary falsifier.

Architecture-level Golden Flows define acceptance targets. They do not claim implementation conformance until the real runtime executes the required proof.

## 17. Implementation-slice contract

TA-12 must produce slices where foundational decisions are already explicit.

Each slice carries, as applicable:

```text
outcome
exact owners
admitted operations/consumers
contract revisions/generated projections
allowed GENERATED/AURORA-FOUNDATION/MODULE-OWNED surfaces
admitted dependencies/version pins or admission gates
allowed persistence/migration/config/process surfaces
security/effect constraints
proof/falsifier obligations
prerequisites
stop/reopen triggers
completion criteria
```

A slice may not independently choose a foundational stack already owned by a global readiness decision.

## 18. Final readiness test

TA-13 closes only when a fresh challenger cannot identify a material implementation decision that remains ownerless or convenience-driven.

The review attacks at least:

```text
consumer without operation
operation without consumer
operation without owner
owner duplication
parallel contract/DTO authority
screen-shaped API
provider/framework semantics promoted to Aurora meaning
store without invariant/owner
identity proof confused with authority
policy without enforcement
secret without credential lifecycle
Paved-Road silent bypass
runtime mechanism without consumer
recovery with ambiguous owner
unfalsifiable Golden Flow
implementation slice that still requires architecture
```

Only after TA-13 and a separate explicit operator Product execution grant does Product implementation become eligible.
