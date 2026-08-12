---
id: DOC-AURORA-CAPABILITY-REALIZATION-METHOD
title: Aurora Capability Realization Method
document_type: capability_realization_method
form: reference

authority: standard
status: accepted
accepted_at: 2026-08-06
acceptance_evidence: DOC-AURORA-A0-OPERATOR-ACCEPTANCE
version: 0.2.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - blueprint-to-build realization process
  - program-level System Architecture Rebaseline integration
  - readiness gates
  - requirements traceability
  - implementation and evidence promotion
  - Product Milestone closeout
related:
  - DOC-AURORA-PRODUCT-INDEX
  - DOC-AURORA-BLUEPRINT-14
  - DOC-AURORA-BLUEPRINT-15
  - DOC-AURORA-REQUIREMENTS-TRACEABILITY
  - DESIGN-AURORA-SYSTEM-ARCHITECTURE-REBASELINE
  - DOC-AURORA-SYSTEM-ARCHITECTURE-REBASELINE-OPERATOR-DIRECTION
  - DESIGN-AURORA-SYSTEM-ARCHITECTURE-DECISION-LANDSCAPE
review_triggers:
  - readiness gate changes
  - System Architecture Rebaseline method changes
  - implementation methodology changes
  - evidence or closeout model changes
last_reviewed: 2026-08-12
---

# Aurora Capability Realization Method

## 1. Purpose

The Product Blueprint defines what Aurora is and which invariants must survive implementation.

The Capability Realization Method defines how one part of that intent becomes:

```text
research
→ explicit requirements
→ decisions and executable uncertainty reduction
→ reusable Capability Spec
→ scoped Mission Contract
→ implementation plan
→ code/configuration
→ verification and evidence
→ accepted Product Milestone
```

The method prevents:

- implementation beginning from a conversation summary;
- a framework chosen before the requirement is understood;
- code existing without a product requirement;
- requirements disappearing between Blueprint and task list;
- tests proving implementation details but not the user journey;
- a Harness self-declaring completion;
- architecture spikes becoming production by accident;
- distant roadmap content becoming current commitment;
- evidence existing without a criterion;
- a milestone closing because all tasks were checked.

This method is intentionally rigorous for material capabilities. It must be applied proportionally to scope/risk.

---

## 2. Core principles

### M1 — Every durable implementation traces to approved intent

No material code path without a Blueprint, ADR, Capability requirement, Contract criterion or accepted defect/finding.

### M2 — Every approved requirement traces forward to evidence or explicit deferral

No orphan requirement.

### M3 — Research, decision, specification, contract and evidence are separate artifacts

They may reference each other but do not collapse.

### M4 — Uncertainty is reduced before commitment

Research and Architecture Spikes answer material unknowns before production design is frozen.

### M5 — Capability Spec owns reusable product behavior

Mission Contract owns one scoped implementation commitment.

### M6 — Implementation plan cannot change approved product intent

Material change triggers replan/ADR/spec revision.

### M7 — Claims are not acceptance

Implementer output needs receipts/evidence/verdict appropriate to risk.

### M8 — Product Milestone closes with an end-to-end Golden Proof

Local components green are insufficient.

### M9 — Authorization is explicit per gate

A completed artifact does not imply authority to move to the next stage.

### M10 — Rigor is proportional to risk but traceability remains

A small reversible change may use a compact path; security, memory, authority and physical effects need full gates.

---

## 3. Artifact model

### Product Blueprint

Constitutional product intent and long-horizon principles.

### Roadmap/Product Milestone

Capability sequence and Golden Proof.

### Research Report

Evidence landscape, alternatives, limitations and decision implications.

### Architecture Spike

Executable investigation of a material uncertainty.

### ADR

Accepted/rejected specific decision.

### Capability Spec

Complete reusable behavior, contracts, lifecycle, security and evaluation.

### Capability Requirement

Atomic normative statement derived from product intent/research/decision.

### Mission Contract

Scoped implementation commitment: exact outcome, criteria, authority, versions and non-goals.

### Implementation Plan / Microdesign

Concrete files, interfaces, algorithms, migrations and tests for one contract scope.

### Code / Configuration / Schema

Implemented mechanism.

### Claim

Actor states requirement/criterion is fulfilled.

### Receipt

Controlled check/effect result.

### Evidence

Artifacts/observations linked to requirement/criterion.

### Verdict

Permitted reviewer/authority accepts/rejects evidence.

### Product Milestone Closeout

Global proof, limitations, metrics, incidents and accepted state.

---

## 4. Identity hierarchy

Example identities:

```text
Blueprint principle:
  PB-06-P06 — Memory guides; authority and evidence determine

Product Milestone:
  PM-M1 — Governed Conversation, Project Context and Memory

Capability:
  CAP-MEMORY-CONTEXT

Capability requirement:
  CAP-MEMORY-CONTEXT-REQ-017

Mission:
  MIS-M1-MEMORY-SLICE-001

Mission criterion:
  MIS-M1-MEMORY-SLICE-001-CRIT-008

Implementation unit:
  TASK-...

Evidence:
  EVID-M1-MEMORY-SUPERSESSION-001
```

Exact syntax may be refined, but IDs must be stable and unambiguous.

---

## 5. Requirement qualities

A Capability Requirement should be:

- atomic enough to verify;
- normative (`MUST`, `MUST NOT`, `SHOULD` with rationale);
- scoped;
- attributable to source intent;
- implementation-neutral unless decision already accepted;
- test/evidence-aware;
- explicit about conditions;
- free of vague adjectives such as “robust” without measurable behavior.

### Bad

> Memory should be good and safe.

### Better

> When a current accepted ADR conflicts with a conversational memory, Context Builder MUST mark the memory as non-governing and include the ADR as the current authority for action.

### Requirement metadata

```yaml
id: CAP-MEMORY-CONTEXT-REQ-017
statement: ...
source:
  - DOC-AURORA-BLUEPRINT-06#6.13
rationale: prevent stale decision from governing
risk: high
verification:
  - contract_test
  - adversarial_journey
status: proposed
allocation: []
evidence: []
```

---

## 6. Applicability

Not every Blueprint statement applies to every Capability.

A Capability applicability analysis classifies each relevant requirement:

```text
APPLIES
PARTIALLY_APPLIES
NOT_APPLICABLE
DEFERRED_BY_ROADMAP
CONFLICT_REQUIRES_DECISION
```

A `NOT_APPLICABLE` result requires rationale.

Example for a read-only Research Harness:

- physical interlock → `NOT_APPLICABLE`;
- context minimization → `APPLIES`;
- provider trust/build → `APPLIES`;
- autonomous production promotion → `NOT_APPLICABLE` but prohibition may still apply;
- durable resume → `PARTIALLY_APPLIES` depending on capability profile.

Applicability prevents both missing cross-cutting requirements and blindly importing all product complexity.

---

## 6A. Program-level System Architecture Rebaseline

### Purpose

Aurora is a system of systems. Capability-specific readiness alone can produce locally valid implementations that still conflict across identity, authority, data ownership, contracts, memory, execution, observability, Presence or deployment.

A **System Architecture Rebaseline** establishes the current cross-system map before further multi-subsystem implementation expansion. It supplies coherent constraints to later capability gates without freezing every distant mechanism.

It is a program-level architecture activity inside this method. It is **not**:

- a new readiness gate;
- a second lifecycle or FSM;
- a numerical maturity score;
- a substitute for Capability Specs, ADRs or Mission Contracts;
- implementation authority.

### Triggers

A rebaseline is required when one or more of the following becomes material:

- an upcoming executable horizon spans multiple major subsystems;
- different capabilities could create overlapping sources of truth or incompatible contracts;
- a fixed implementation/spike exposes a cross-system architecture assumption not owned by its local scope;
- accepted technology decisions risk being generalized beyond their stated scope;
- a new trust, data, execution, Presence, model or physical boundary is entering the executable horizon;
- material architecture drift or documentation divergence is found;
- current research contradicts a global architecture assumption;
- the operator explicitly pauses implementation for architecture alignment.

### Inputs

- accepted Product Blueprint and roadmap;
- Documentation Map and ownership hierarchy;
- accepted ADRs with exact scope;
- current Capability Specs and Mission Contracts where relevant;
- current Research Map and source freshness;
- fixed implementation, Architecture Spike and operational evidence;
- current decision/open-question index;
- current authorization and blockers from `STATUS.md`.

### Required outputs

A material rebaseline produces or refreshes:

1. system context and external boundaries;
2. logical modules and canonical state/data ownership;
3. trust, authority, execution and deployment boundaries;
4. cross-system contracts and material data flows;
5. a global Architecture Decision Landscape;
6. earliest consumer for each open decision;
7. evidence needed to promote each decision;
8. explicit Findings and replan obligations;
9. implementation pause/authorization boundary;
10. reconsideration triggers for every deliberate deferral.

### Decision treatment vocabulary

#### `DECIDE`

Current evidence is sufficient and the choice is required before the next architecture or implementation commitment.

The result must be promoted to the correct owner—normally an ADR, accepted Specification or Standard—and include scope, alternatives, consequences, compatibility and reconsideration triggers.

#### `RESEARCH`

The question is material to an upcoming decision, but current evidence is insufficient or temporally stale.

Research must name the exact decision it can inform, use current primary sources and state limitations. Research does not decide.

#### `SPIKE`

The required property depends on observed runtime, operational, security, compatibility, latency or recovery behavior that documentary analysis cannot establish.

The spike must be minimal, disposable by default, evidence-producing and separately authorized before execution.

#### `DEFER`

The answer does not change the next evidence-supported architecture/build decision.

A valid deferral records:

- the earliest known consumer;
- constraints that must be preserved now;
- the exact reconsideration trigger;
- prohibition against speculative implementation.

`DEFER` is a governed non-selection, not an undocumented omission.

### Landscape record

Each material entry should identify:

```text
architecture area
accepted product constraint
existing accepted decision or explicit non-decision
current non-governing hypothesis
concrete open question
dependencies and lock-in/risk
earliest consumer
DECIDE | RESEARCH | SPIKE | DEFER
evidence needed
decision owner
reconsideration trigger
```

### Relationship to R0–R8

The rebaseline feeds the existing gates:

```text
Product Constitution
        ↓
System Architecture Rebaseline
        ↓
Product Milestone / Capability
        ↓
R0 → R1 → R2 → R3 → R4 → R5 → R6
        ↓
R7 Execution and Evidence
        ↓
R8 Product Milestone Closeout
```

- R0 checks whether applicable global architecture constraints are current and discoverable.
- R1 classifies those cross-system constraints for the Capability.
- R2 derives verifiable requirements rather than importing a technology hypothesis.
- R3 allocates the requirements to complete reusable behavior and boundaries.
- R4 consumes and revalidates the landscape, performs required research/spikes and promotes material choices to accepted owners.
- R5 binds the exact accepted versions and scope into one Mission commitment.
- R6 prevents implementers from inventing remaining material architecture.
- R7 follows the accepted Contract/design; a deviation opens a Finding and replan.
- R8 evaluates the Product Milestone outcome, not the existence of the rebaseline.

A landscape entry is not an accepted decision merely because it is detailed. A capability R4 still requires accepted ADR/Specification authority where material.

### Guardrails

The rebaseline must not:

- choose a universal stack by popularity;
- research distant mechanisms without a named consumer;
- create code, schemas, services or adapters as hidden implementation;
- turn a milestone-scoped ADR into a global mandate;
- allow a framework/protocol/runtime to own Aurora canonical semantics;
- rewrite accepted historical decisions to match current preference;
- authorize the next gate by implication.

Completion of a rebaseline only establishes architecture readiness for further scoped decision work. It does not authorize implementation, R7 continuation, merge/promotion or Product Milestone closeout.

---

## 7. Readiness gates R0–R8

```text
R0 — Constitutional Baseline
R1 — Applicability
R2 — Requirements
R3 — Capability Readiness
R4 — Architecture/Decision Readiness
R5 — Contract Readiness
R6 — Implementation Design Readiness
R7 — Execution and Evidence
R8 — Product Milestone Closeout
```

No gate advances solely because a document exists.

---

# 8. R0 — Constitutional Baseline

## Question

Is the product intent required for this capability accepted, coherent and discoverable?

## Inputs

- current Product Blueprint;
- Documentation Map;
- current roadmap;
- accepted/rejected decisions;
- discovery coverage;
- current authorization.

## Checks

- relevant Blueprint sections accepted/current;
- no unresolved constitutional divergence;
- Product Milestone exists and has Golden Proof;
- scope/non-goals clear;
- current status authorizes readiness work;
- new session can find sources.

## Output

```text
R0 PASS | FAIL | BLOCKED
```

## Failure example

Trying to implement device actuation before Blueprint 09/10/11 or M10 acceptance design is complete.

---

# 9. R1 — Applicability

## Question

Which constitutional, security, memory, authority, reliability and documentation requirements apply to this Capability?

## Process

1. list source sections/principles;
2. classify applicability;
3. record rationale;
4. identify cross-capability dependencies;
5. identify conflicts/open research;
6. assign owner.

## Output

```text
CAP-*/APPLICABILITY.md
```

or equivalent structured table.

## Gate condition

- every known cross-cutting source considered;
- no unjustified exclusions;
- high-risk dependencies identified.

---

# 10. R2 — Requirements

## Question

Are all applicable product statements transformed into verifiable requirements?

## Process

- derive requirements;
- assign IDs;
- link sources/rationale;
- classify risk;
- define verification method;
- identify open decision/spike;
- detect duplicates/conflicts;
- create coverage matrix.

## Verification methods

```text
STATIC_ANALYSIS
SCHEMA_VALIDATION
UNIT_TEST
CONTRACT_TEST
CONFORMANCE
INTEGRATION
FAULT_INJECTION
SECURITY_TEST
SIMULATION
HIL
PHYSICAL_DRILL
BENCHMARK/EVAL
USER_JOURNEY
DOCUMENT_REVIEW
OPERATOR_VERDICT
```

## Gate condition

- no applicable source without requirement or rationale;
- no vague/high-risk requirement without verification direction;
- contradictions resolved or block.

---

# 11. R3 — Capability Readiness

## Question

Is the reusable Capability design complete enough to support a scoped implementation contract?

## Required Capability Spec sections

```text
Purpose and use cases
Goals / non-goals
Applicability
Domain model
Contracts / schemas
Lifecycle / state
Architecture and boundaries
Context / memory
Authority / effects
Security / privacy / threat model
Failure / recovery
Observability
Evaluation / evidence
Rollout / graduation
Compatibility / migration
Open questions
Requirement coverage
```

## Capability test plan

- requirement-to-test mapping;
- representative journeys;
- adversarial/fault cases;
- evaluation datasets;
- success thresholds/rationale;
- evidence format;
- graduation levels.

## Gate condition

- all R2 requirements allocated to Spec mechanisms and tests;
- no architecture placeholder for current scope;
- open questions explicitly outside current contract or blocking;
- owner/reviewer defined.

---

# 12. R4 — Architecture/Decision Readiness

## Question

Have material technical uncertainties and choices been researched, proven and decided enough for the scoped Mission?

## Inputs

- current System Architecture Decision Landscape entries applicable to the scope;
- focused research;
- architecture spikes;
- ADRs;
- threat model;
- compatibility/migration analysis;
- operational burden analysis.

## Architecture Spike lifecycle

```text
PROPOSED
→ AUTHORIZED
→ EXECUTING
→ EVIDENCE_COMPLETE
→ REVIEWED
→ DECISION_INFORMED
→ CLOSED
```

### Spike requirements

- exact uncertainty;
- alternatives;
- minimal prototype;
- controlled environment;
- measurements;
- Golden Proof;
- evidence artifacts;
- removal/disposal conditions;
- decision implications.

Spike code is not production by default.

## Gate condition

- applicable landscape entries were revalidated against current evidence;
- current scope has no unresolved material architecture choice;
- accepted ADRs exist where needed;
- decisions compatible with Blueprint;
- migration/rollback considered;
- spike evidence reviewed.

---

# 13. R5 — Contract Readiness

## Question

Is one scoped Mission commitment exact, reviewable, authorized and traceable?

## Mission Contract must define

### Identity and baseline

- Mission ID/revision;
- exact repository/system baseline;
- relevant contract/spec versions;
- environment;
- authority source.

### Outcome

- objective;
- operator-visible value;
- acceptance criteria;
- Golden Proof contribution.

### Scope

- included;
- non-goals;
- assumptions;
- dependencies;
- external effects;
- risks.

### Decomposition

- milestones/features/tasks at contract level;
- ownership;
- dependencies;
- parallelism/isolation if relevant.

### Requirements allocation

Every criterion maps to Capability requirements.

### Authority and budget

- what can be changed/executed;
- prohibited actions;
- credentials/effects;
- time/resource limits;
- approval gates.

### Evidence

- test/eval/drill artifacts;
- independent review;
- closeout.

### Change/replan

- triggers;
- immutable approved revision/hash;
- supersession path.

## Gate condition

- exact contract approved by required authority;
- all in-scope requirements allocated;
- no hidden scope;
- implementation authorization still separate if policy requires.

---

# 14. R6 — Implementation Design Readiness

## Question

Can implementers execute the approved contract without inventing material architecture or product behavior?

## Implementation Plan/Microdesign includes

- exact files/modules;
- current code references;
- interfaces/types/schemas;
- data migrations;
- algorithms/state transitions;
- effect/security integration;
- test-first steps;
- fault/recovery;
- observability;
- documentation changes;
- rollback;
- verification commands;
- commit/PR strategy.

## Quality criteria

- steps small and verifiable;
- no ambiguous “implement feature” steps;
- existing patterns checked;
- planned tests prove contract criteria;
- cross-layer wiring included;
- no unrelated refactoring;
- environment/reproduction clear.

## Gate condition

- adversarial plan review;
- all tasks trace to criteria/requirements;
- no material design decision delegated to implementer;
- implementation explicitly authorized.

---

# 15. R7 — Execution and Evidence

## Execution rules

- follow approved contract/design;
- TDD/verification per project standard;
- preserve exact baseline and revisions;
- claims linked to criteria;
- external effects follow authority;
- deviations trigger finding/replan;
- work-in-progress status durable;
- docs impact updated.

## Evidence flow

```text
Implementer Claim
→ deterministic Receipts
→ independent Review/QA where needed
→ Evidence Bundle
→ Verdict by allowed authority
```

## Evidence Bundle

May include:

- code diff/commit;
- build/test output;
- conformance report;
- benchmark/eval;
- trace;
- artifact hashes;
- screenshots/recording;
- security/fault result;
- physical measurement/calibration;
- reviewer findings/resolution;
- documentation validation.

## Gate condition

- every criterion has sufficient evidence or explicit rejection/deferral;
- implementation and docs align;
- security/effects reviewed;
- integrated journey tested;
- no unresolved critical finding;
- automatic merge only if separately authorized.

---

# 16. R8 — Product Milestone Closeout

## Question

Did the Product Milestone deliver its global capability and retire the named risk?

## Closeout contents

- milestone/version/date;
- accepted outcome;
- Golden Proof procedure/results;
- capability versions;
- evidence index;
- operational metrics;
- security/safety incidents;
- unresolved limitations;
- deviations;
- lessons;
- documentation impact;
- rollback/support plan;
- next milestone readiness implications;
- explicit operator acceptance.

## Gate condition

- Golden Proof end-to-end;
- operator-visible value demonstrated;
- named risk reduced with evidence;
- not merely local component completion;
- status/roadmap updated;
- next milestone not auto-authorized.

---

## 17. Risk-proportional paths

### Compact path

For low-risk, reversible, local changes within accepted design:

```text
existing R0–R4 baseline
→ compact Contract/issue criterion
→ implementation plan/TDD
→ deterministic evidence
→ review/merge
```

### Standard capability path

R0–R8 with Capability Spec and ADRs as needed.

### Critical path

For:

- memory semantics;
- authority/security;
- credentials;
- provider trust;
- production effects;
- physical actuation;
- self-improvement promotion.

Adds:

- threat model;
- independent/adversarial review;
- fault injection;
- staged rollout/canary;
- drill/rollback;
- explicit operator gates.

Rigor affects artifact depth, not permission to skip traceability.

---

## 18. Traceability graph

```text
Blueprint section/principle
        ↓ derives
Capability Requirement
        ↓ allocated to
Capability Spec mechanism
        ↓ committed by
Mission Criterion
        ↓ implemented by
Task / Code / Config
        ↓ verified by
Receipt / Evidence
        ↓ accepted by
Verdict / Closeout
```

Reverse trace:

```text
Code line/module
→ implementation task
→ mission criterion
→ capability requirement
→ Blueprint/ADR source
```

This permits impact analysis when product intent changes.

---

## 19. Orphan detection

### Intent orphan

Blueprint statement with no requirement/explicit deferral.

### Requirement orphan

Requirement with no Spec/criterion allocation.

### Contract orphan

Criterion without source requirement.

### Implementation orphan

Code/feature with no accepted criterion/defect.

### Evidence orphan

Artifact/test not linked to criterion.

### Decision orphan

ADR with no research/requirement impact.

### Documentation orphan

Normative document with no source-of-truth scope or consumer.

Checks should become mechanical where possible.

---

## 20. Change and replan triggers

Material triggers:

- objective/scope changes;
- accepted requirement changes;
- architecture assumption fails;
- provider/framework/version changes semantics;
- security/threat finding;
- effect/data boundary change;
- budget/risk changes;
- evaluation metric/holdout changes;
- implementation cannot satisfy criterion;
- physical environment/device change;
- new evidence contradicts decision.

Response:

```text
stop affected work
→ open Finding
→ identify owning artifact
→ revise research/ADR/Spec/Contract
→ re-run required gates
→ approve new exact revision
```

Do not mutate approved Contract in place.

---

## 21. Decisions during implementation

### Local implementation decision

Can be made by implementer when it does not alter:

- contract behavior;
- architecture boundary;
- authority/security;
- public/internal contract;
- performance/risk materially;
- documentation truth.

Record in code/plan as appropriate.

### Material decision

Requires ADR/Spec/Contract revision and appropriate approval.

Examples:

- replacing storage engine;
- exposing new external effect;
- changing memory precedence;
- adding provider data access;
- changing an interlock;
- weakening evidence.

---

## 22. Evidence independence

Risk classes determine reviewer independence.

### Low

Automated checks + peer/code review.

### Medium

Separate reviewer context/model and integration journey.

### High

Independent security/safety review, adversarial tests, operator verdict, canary/drill.

The same agent may help create and evaluate low-risk drafts, but cannot be the sole authority for material self-improvement or physical safety.

---

## 23. Documentation impact at every gate

At R2–R8, assess:

```yaml
documentation_impact:
  status: NONE | UPDATED | FOLLOW_UP_REQUIRED
  affected: []
  rationale: ""
  follow_up: null
```

Examples:

- schema change → reference/spec/generated types;
- runtime behavior → Capability Spec/ADR;
- new operation → how-to/runbook;
- accepted evidence → acceptance report/status;
- failed spike → research/decision history;
- incident → threat model/eval/Golden Path.

---

## 24. Historical A0 application

A0 was the first product-level application of this method:

```text
Discovery conversations
→ Product Blueprint requirements
→ research/ADR proposals
→ documentation coverage
→ adversarial review
→ fresh-session Golden Proof
→ operator acceptance
```

This section is historical. It does not own the current Product Milestone, readiness gate, implementation state or next action; those values belong to `docs/tracking/STATUS.md`.

---

## 25. A0 acceptance checklist

- [x] 15 Blueprint sections complete.
- [x] Discovery History preserves original examples/reasoning.
- [x] Documentation Coverage maps all discovery decisions.
- [x] Focused research reports/source manifests exist for current architecture claims.
- [x] ADR proposals are consistent with research and Blueprint.
- [x] Requirements Traceability maps constitutional requirements.
- [x] Product aggregate and roadmap projection are current.
- [x] IDs, links, metadata and source references validate.
- [x] Adversarial review findings are resolved.
- [x] Fresh-session Golden Proof passes.
- [x] Operator explicitly accepts A0.
- [x] Status records next gate and prohibitions.

---

## 26. Non-goals

- require a full Capability Spec for every tiny local change;
- make documentation volume a gate by itself;
- authorize implementation from research;
- let a plan override an ADR/Blueprint;
- accept a provider claim as evidence;
- automatically merge after checks;
- freeze all architecture before any spike;
- use one workflow engine for the realization process;
- replace engineering judgment with mechanical coverage numbers;
- allow distant roadmap milestones to enter execution without readiness.
