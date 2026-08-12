---
id: DESIGN-AURORA-SYSTEM-ARCHITECTURE-REBASELINE
title: Aurora System Architecture Rebaseline Design
document_type: system_design
form: explanation
authority: design
status: proposed
version: 0.1.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - proposed program-level System Architecture Rebaseline
  - proposed treatment of the frozen M0 R7 implementation candidate during rebaseline
related:
  - DOC-AURORA-STATUS
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DOC-AURORA-BLUEPRINT-12
  - DOC-AURORA-BLUEPRINT-14
  - DOC-AURORA-BLUEPRINT-15
  - DOC-AURORA-DECISIONS
  - DESIGN-AURORA-M0-R4-DECISION-LANDSCAPE
source_revision: e7ca5ffb652fbbd68b35d4434506c58d26daf0e1
last_reviewed: 2026-08-12
---

# Aurora System Architecture Rebaseline Design

## 1. Decision requested

Projeto Aurora already has a strong accepted product constitution, a logical architecture, a domain vocabulary, an R0–R8 Capability Realization Method and a deeply specified M0 Sovereign Core slice.

The current gap is not absence of product meaning. It is the missing program-level bridge between:

```text
accepted long-term product constitution
→ coherent cross-system technical architecture
→ capability-specific R0–R6 readiness
→ implementation
```

This design proposes a **System Architecture Rebaseline** before any further Aurora implementation or Product Milestone expansion.

The rebaseline is a program-level architecture activity inside the existing Aurora Capability Realization Method. It is **not a new ACRM gate, lifecycle, score or parallel governance framework**.

## 2. Why the rebaseline is necessary

Aurora is a system of systems. Its accepted direction includes:

- sovereign identity, state and authority;
- a persistent cognitive Core;
- governed memory and context construction;
- replaceable models and cognitive runtimes;
- capability and provider contracts;
- specialized Harnesses;
- AHDK and independent conformance;
- digital and physical Effect boundaries;
- artifacts, evidence and verdicts;
- multiple presences and multimodality;
- voice, vision, devices and laboratory progression;
- durable campaigns and governed self-improvement.

M0 was intentionally narrow. Its R4–R6 work answered how to implement one local Sovereign Core walking skeleton. It did not and should not have selected a universal database, authentication product, event broker, API style, voice provider, model router, deployment topology or full-system module topology for every later Aurora capability.

Without a program-level architecture baseline, later capabilities could each invent incompatible meanings for:

- identity;
- project and mission state;
- authentication and authorization;
- API errors and idempotency;
- event envelopes;
- data ownership;
- memory and knowledge;
- secrets and credentials;
- observability and audit;
- runtime and sandbox boundaries;
- provider and model integration.

The likely result would be either a monolithic superagent or an accidental distributed platform with overlapping sources of truth.

## 3. What remains accepted

The rebaseline does not reopen accepted constitutional meaning without a material Finding.

The following remain governing:

- Aurora is a Leandro-first personal cognitive operating system and global capability control plane;
- Aurora owns global identity, project/mission meaning, authority, budgets, governed context and global outcome;
- Harnesses own specialized local methodology and execution inside Delegations;
- Aurora owns language- and framework-neutral Contract Model semantics;
- protocols and runtimes are replaceable bindings;
- first-party Harnesses use AHDK by policy unless explicitly waived;
- conformance remains independent from AHDK;
- memory is distinct from authoritative state, evidence and live sources;
- security enforcement for material effects remains outside probabilistic model judgment;
- logical modularity precedes physical distribution;
- local-first/cloud-assisted sovereignty remains the direction;
- Product Milestones close through end-to-end Golden Proofs and operator Verdicts.

Accepted ADRs remain accepted for their stated scope. In particular:

```text
Go selected for the M0 Sovereign Core
≠ Go selected for every Aurora component

SQLite selected for M0 operational state
≠ SQLite selected as Aurora's universal database

Mastra preferred-first for future first-party agentic Harnesses
≠ Mastra selected as Aurora Sovereign Core or global state owner
```

## 4. Decision

Before further Aurora production implementation, the program will establish a **System Architecture Rebaseline** that answers enough cross-system questions to make later capability decisions coherent, while deferring choices that do not affect the next evidence-supported step.

The governing rule is:

> Plan deeply where a decision creates structural coupling or is expensive to reverse. Preserve explicit boundaries and defer replaceable mechanisms until a real consumer makes the decision material.

The rebaseline will:

1. consolidate accepted product and architecture constraints;
2. map logical subsystems, module boundaries and ownership;
3. identify cross-system contracts and data flows;
4. identify trust, execution and deployment boundaries;
5. create a global Architecture Decision Landscape;
6. classify each open question as `DECIDE`, `RESEARCH`, `SPIKE` or `DEFER`;
7. identify the earliest Product Milestone/capability that consumes each decision;
8. prevent local milestone choices from becoming accidental global architecture;
9. feed capability-specific R0–R6 work without replacing it;
10. preserve explicit non-implementation authority until later gates.

## 5. Classification model

### `DECIDE`

Current accepted evidence is sufficient and the choice is required before the next architecture or implementation commitment.

Required output:

- owning ADR, Specification, Standard or accepted architecture artifact;
- alternatives and consequences;
- scope and compatibility;
- reconsideration triggers.

### `RESEARCH`

The question is material to an upcoming decision, but current evidence is insufficient or temporally stale.

Required output:

- exact research question;
- source and freshness policy;
- decision that the research can inform;
- explicit limitations.

### `SPIKE`

The required property depends on observed runtime, operational, security, latency, compatibility or recovery behavior that documentary analysis cannot establish.

Required output:

- exact uncertainty;
- minimal disposable experiment;
- measurements and Golden Proof;
- failure cases;
- disposal/promotion rule;
- separate execution authorization.

### `DEFER`

The answer does not change the next architecture/build decision.

Required output:

- earliest known consumer;
- reconsideration trigger;
- constraints that must be preserved now;
- explicit prohibition against speculative implementation.

`DEFER` is a governed decision not to decide yet. It is not an undocumented omission.

## 6. Architecture Landscape record

Each material question will carry:

| Field | Purpose |
|---|---|
| Architecture area | subsystem or cross-cutting concern |
| Product constraint | accepted Blueprint/requirement that bounds the answer |
| Existing decision | applicable accepted ADR or explicit non-decision |
| Current hypothesis | non-governing working model |
| Open question | concrete decision to resolve |
| Dependencies | decisions and capabilities affected |
| Lock-in/risk | cost and consequence of a wrong decision |
| Earliest consumer | first milestone/capability requiring commitment |
| Treatment | `DECIDE`, `RESEARCH`, `SPIKE` or `DEFER` |
| Evidence needed | sufficient proof for promotion |
| Decision owner | Blueprint, ADR, Spec, Standard or Contract |
| Reconsideration trigger | condition that reopens the result |

## 7. Initial architecture clusters

The first landscape will cover these clusters without implying one ADR or one service per cluster:

1. system context and external boundaries;
2. logical modules and state ownership;
3. Contract Model and compatibility;
4. API architecture and error/idempotency conventions;
5. events, messaging and synchronization;
6. human, Aurora, service, provider and device identity;
7. authentication;
8. authorization, policy and Effect enforcement;
9. secrets and credential brokering;
10. data categories, ownership and lifecycle;
11. operational storage;
12. backup, export, migration and portability;
13. memory, knowledge, retrieval and Context Builder;
14. model and inference architecture;
15. Brain/cognitive runtime boundaries;
16. Harness, AHDK and conformance integration;
17. sandbox and execution environments;
18. durable work, timers and scheduling;
19. artifacts, evidence and provenance;
20. observability, audit and evaluation;
21. voice and real-time interaction;
22. vision and multimodality;
23. Presence and device trust;
24. deployment and topology;
25. configuration and environment management;
26. networking and local/cloud boundaries;
27. software supply chain and build provenance;
28. laboratory and physical-device boundaries.

Many clusters will remain `DEFER` until a named milestone enters its executable horizon.

## 8. Relationship to ACRM R0–R8

The System Architecture Rebaseline does not replace capability realization.

```text
Product Constitution
        ↓
System Architecture Rebaseline
        ↓
Product Milestone / Capability
        ↓
R0 → R1 → R2 → R3 → R4 → R5 → R6
        ↓
R7 implementation and evidence
        ↓
R8 Product Milestone closeout
```

Capability-specific behavior remains owned by Capability Specs and Mission Contracts.

The rebaseline contributes:

- current cross-system constraints to R0/R1;
- ownership and boundary hypotheses to R2/R3;
- decision dependencies and evidence obligations to R4;
- accepted versions and non-goals to R5;
- implementation boundaries to R6.

Every consuming R4 must revalidate the global landscape against current evidence. A landscape entry cannot substitute for an accepted ADR or Capability Spec when one is required.

## 9. Treatment of the M0 R7 candidate

The non-canonical R7 branch is preserved as an executable candidate and evidence source.

It is not:

- merged into `main`;
- an R7 independent Verdict;
- an R8 closeout;
- proof that the whole Aurora architecture is production-ready;
- authority to continue TASK-13 or implement M1+;
- a universal technology baseline.

During the rebaseline:

```text
feat/m0-r7-sovereign-core-20260810
→ FROZEN / PRESERVED
→ may inform Findings and architecture evidence
→ must not receive implementation expansion
→ promotion requires renewed review against the rebaseline
```

The branch is not rejected or discarded. Any future continuation must first determine whether its M0-scoped decisions remain compatible with the program architecture baseline.

## 10. Development Harness boundary

Leandro's software-development Harness may later be the primary factory that researches, plans, implements, tests, reviews and packages evidence for Aurora.

The relationship is:

```text
Development Harness
        │ builds / verifies
        ▼
      Aurora
```

Aurora runtime must not depend on the Development Harness runtime for sovereign operation.

A future `Aurora-ready` threshold for the Development Harness may require:

- durable task/plan continuity;
- isolated workspaces/sandboxes;
- current primary-source research;
- TDD and full-target verification;
- independent review;
- requirement/decision traceability;
- evidence packaging;
- architectural Finding and replan behavior;
- safe authority/effect boundaries.

That threshold is program execution planning, not Aurora product architecture.

## 11. Documentation ownership

The rebaseline uses existing owners:

- `docs/product/CAPABILITY-REALIZATION-METHOD.md` owns the method integration;
- `docs/design/AURORA-SYSTEM-ARCHITECTURE-DECISION-LANDSCAPE.md` owns the proposed architecture question/dependency map;
- accepted ADRs own promoted specific decisions;
- Capability Specs own reusable behavior;
- `docs/tracking/STATUS.md` owns current pause/authorization/next action;
- `docs/tracking/DECISIONS.md` indexes accepted and open decisions;
- `docs/tracking/WORKLOG.md` preserves chronology;
- an operator-direction record preserves the explicit rebaseline instruction.

Blueprint 12 already owns logical architecture and should not receive implementation technology choices during this change. It is revised only if the rebaseline finds a material constitutional architecture defect.

## 12. Non-goals

This change does not:

- implement Aurora;
- continue M0 R7 TASK-13;
- merge or promote the R7 branch;
- authorize R8;
- select Keycloak, Zitadel, Ory, SPIFFE, Cedar, OPA, Vault or another identity/policy mechanism;
- select PostgreSQL, graph, vector, object, event or telemetry stores;
- select REST, gRPC, A2A, MCP, CloudEvents, AsyncAPI or a broker;
- select voice, vision or model providers;
- implement AHDK, MNFS or a Mastra adapter;
- create a new ACRM gate or parallel lifecycle;
- redesign accepted product meaning;
- generalize M0 implementation choices into global mandates.

## 13. Adversarial failure modes

The rebaseline fails if it becomes:

### Architecture theater

Large diagrams and inventories without decision owners, consumers or evidence paths.

### Universal stack selection

Choosing products before data, identity, authority and runtime requirements are understood.

### Infinite planning

Researching distant mechanisms with no named consumer or reconsideration trigger.

### Hidden implementation

Creating abstractions, adapters, schemas, services or spike code under the label of planning.

### Local-decision globalization

Treating M0 Go/SQLite or another milestone mechanism as a universal platform decision.

### Duplicate governance

Creating another lifecycle, score, authority hierarchy or source of truth beside ACRM.

### Framework capture

Allowing a model, SDK, agent runtime, protocol or workflow engine to define Aurora canonical semantics.

### Frozen-code authority

Treating the R7 candidate as accepted merely because its CI or Golden Proof was green.

## 14. Acceptance criteria

This design is ready for promotion when the operator confirms:

1. implementation remains paused;
2. R7 candidate remains preserved but non-canonical;
3. no R8 authority is implied;
4. ACRM R0–R8 remains the only realization lifecycle;
5. System Architecture Rebaseline is a program-level input to capability gates;
6. `DECIDE / RESEARCH / SPIKE / DEFER` is the decision-treatment vocabulary;
7. the global landscape records earliest consumer and evidence needed;
8. accepted ADRs retain explicit scope;
9. Blueprint 12 is not overloaded with technology choices;
10. Development Harness builds Aurora but does not become a runtime dependency;
11. later implementation resumes only after architecture review and explicit authorization.

## 15. Immediate next action after acceptance

```text
record operator direction
→ update ACRM method
→ create initial global Architecture Decision Landscape
→ repair STATUS / DECISIONS / WORKLOG continuity
→ repair accepted-ADR lifecycle wording drift
→ run documentation validation
→ perform adversarial review
→ present fixed revision for operator review
```

No Aurora runtime implementation follows from acceptance of this design.
