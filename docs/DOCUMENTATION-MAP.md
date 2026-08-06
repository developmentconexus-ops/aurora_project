---
id: DOC-AURORA-DOCUMENTATION-MAP
title: Aurora Documentation Map
document_type: documentation_map
form: reference
authority: constitutional
status: proposed
version: 0.2.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - documentation discovery
  - authority hierarchy
  - canonical ownership map
  - human and agent read paths
  - documentation storage boundaries
related:
  - DOC-AURORA-PRODUCT-INDEX
  - DOC-AURORA-PRODUCT-BLUEPRINT
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DOC-AURORA-REQUIREMENTS-TRACEABILITY
  - DOC-AURORA-RESEARCH-MAP
  - DOC-AURORA-ROADMAP
  - DOC-AURORA-STATUS
last_reviewed: 2026-08-06
---

# Aurora Documentation Map

## 1. Current phase

```text
A0 — Product, Discovery and Architecture Baseline
```

A0 exists to transform the initial Aurora discovery dialogue into a complete, research-backed and resumable product constitution before runtime implementation begins.

Current work includes:

- Product Blueprint definition;
- preservation of origin, examples and decision reasoning;
- focused research with primary-source manifests;
- proposed architecture decisions;
- constitutional requirement derivation;
- a Blueprint-to-build realization method;
- roadmap and Golden Proof definition;
- mechanical and adversarial documentation validation.

Current work does **not** authorize:

- Aurora Core implementation;
- AHDK implementation;
- architecture spike execution;
- stack selection;
- MNFS integration;
- device or laboratory control.

The authoritative current boundary is always recorded in [`docs/tracking/STATUS.md`](tracking/STATUS.md).

---

## 2. Governance principle

> **One durable concept has one canonical owner. Other artifacts may explain, summarize, research, project or apply it; they cannot silently redefine it.**

Examples:

| Concept | Canonical owner |
|---|---|
| Aurora definition and North Star | Product Blueprint 01 |
| relationship and personality | Product Blueprint 02 |
| Memory promotion and precedence | Product Blueprint 06 |
| specific protocol choice | accepted ADR |
| complete reusable memory capability | Capability Spec |
| one implementation commitment | approved Mission Contract |
| current runtime state | future operational store |
| what happened in one verification | Evidence/Receipt |
| current coordination | STATUS/tracking |
| original discussion and motivation | Discovery History |

A conversation can originate a decision. It becomes durable only after promotion to its owning artifact.

---

## 3. Authority classes

| Class | Authority | Purpose and allowed ownership |
|---:|---|---|
| A0 | Constitutional | product identity, invariants, boundaries, long-horizon direction and authority hierarchy |
| A1 | Decision | one accepted/rejected choice, alternatives, rationale, consequences and reconsideration |
| A2 | Specification | complete reusable capability, protocol, schema, evaluation or component behavior |
| A3 | Contract | exact scoped commitment, criteria, versions, authority and execution boundary |
| A4 | Standard / Policy | method, engineering rule, Golden Path, enforcement profile and waiver |
| A5 | Reference | exact current machinery, syntax, states, schemas and compatibility |
| A6 | Guidance | tutorials, how-to, runbooks and recommended operational sequences |
| A7 | Evidence | observed result, benchmark, test, trace, measurement, acceptance and closeout |
| A8 | Tracking | current coordination, blockers, progress, backlog and work history |
| A9 | Research / Historical | investigation, comparison, source maps, rejected proposals and discovery record |
| A10 | Generated Projection | convenient publication derived from canonical sources; no independent authority |

Reader form and authority are independent. A polished tutorial does not outrank a contract. A large research report does not become a decision because it is detailed.

---

## 4. Conflict precedence

When sources appear to conflict:

```text
1. current constitutional invariant
2. accepted specific ADR compatible with that constitution
3. current accepted Capability/System Specification
4. current approved Contract for the exact scope
5. current Standard / Policy / Golden Path / Profile
6. current implementation-derived Reference
7. Guidance
8. Evidence for the exact observation only
9. Tracking
10. Research / Historical
11. Generated Projection follows its canonical source
```

This hierarchy does not authorize silent conflict resolution.

A material conflict produces:

```text
DOCUMENTATION_DIVERGENCE
```

Required response:

1. identify the concept owner;
2. stop affected planning or implementation when material;
3. determine whether the owner, dependent document or implementation is stale;
4. open a Finding or decision proposal;
5. update/supersede the correct artifact;
6. re-run affected readiness gates and validation.

An ordinary ADR cannot override a constitutional invariant. Constitutional change requires explicit Blueprint revision, impact analysis and operator approval.

---

## 5. Canonical entrypoints

| Path | Audience | Purpose | Authority |
|---|---|---|---|
| `README.md` | humans | concise project introduction and current orientation | guidance/index |
| `AGENTS.md` | agents and new sessions | mandatory bootstrap, gates and read order | guidance/index |
| `docs/DOCUMENTATION-MAP.md` | humans and agents | authority, ownership, discovery and read paths | constitutional reference |
| `docs/product/README.md` | product/architecture readers | complete Product Blueprint and methodology index | constitutional index |
| `docs/product/PRODUCT-BLUEPRINT.md` | full reviewers/export | generated complete constitutional publication | generated projection |
| `docs/product/CAPABILITY-REALIZATION-METHOD.md` | planners and leads | R0–R8 Blueprint-to-build method | standard |
| `docs/product/REQUIREMENTS-TRACEABILITY.md` | architecture and reviewers | proposed constitutional requirements and proof intent | specification |
| `docs/roadmap.md` | operator and contributors | generated Product Milestone sequence and Golden Proofs | generated projection |
| `docs/research/RESEARCH-MAP.md` | researchers and architects | current reports, freshness and open investigations | research index |
| `docs/adr/README.md` | architecture | decision lifecycle and accepted/proposed ADR index | decision index |
| `docs/tracking/STATUS.md` | every working session | exact phase, authorization, evidence, blockers and next action | tracking |

---

## 6. Product Blueprint ownership

| Section | Path | Owns |
|---:|---|---|
| 01 | `docs/product/blueprint/01-product-vision.md` | definition, North Star, scope, principles, success and non-goals |
| 02 | `docs/product/blueprint/02-human-aurora-relationship.md` | relationship, identity, personality, disagreement, proactivity and trust |
| 03 | `docs/product/blueprint/03-domain-world-model.md` | canonical entities, identity, relationships, temporal and epistemic world model |
| 04 | `docs/product/blueprint/04-cognitive-lifecycle-journeys.md` | cognitive loop, work lifecycles and end-to-end reference journeys |
| 05 | `docs/product/blueprint/05-capability-system.md` | capabilities, providers, manifests, Registry, trust, AHDK and conformance |
| 06 | `docs/product/blueprint/06-memory-knowledge-context.md` | memory strata, metadata, promotion, supersession, forgetting and Context Builder |
| 07 | `docs/product/blueprint/07-harness-orchestration.md` | Aurora–Harness boundary, Delegations, Context Packs, events, artifacts and recovery |
| 08 | `docs/product/blueprint/08-interaction-multimodality-presence.md` | interaction surfaces, Presence Fabric, handoff, sensors and degraded operation |
| 09 | `docs/product/blueprint/09-tools-devices-laboratory.md` | tools, devices, instruments, telemetry, protocols and physical progression |
| 10 | `docs/product/blueprint/10-autonomy-authority-safety.md` | authority, effects, campaigns, guardrails, revocation and emergency action |
| 11 | `docs/product/blueprint/11-security-privacy-sovereignty.md` | data sovereignty, privacy, identity, credentials, threat model and containment |
| 12 | `docs/product/blueprint/12-system-architecture.md` | components, ports, state ownership, topology and evolutionary boundaries |
| 13 | `docs/product/blueprint/13-reliability-observability-evaluation.md` | reliability, telemetry, evals, Failure Intelligence and self-improvement evidence |
| 14 | `docs/product/blueprint/14-capability-roadmap.md` | Product Milestones, Golden Proofs, dependencies, non-goals and replan triggers |
| 15 | `docs/product/blueprint/15-documentation-research-governance.md` | document/research lifecycle, promotion, validation, ownership and continuity |

Generated aggregate:

```text
docs/product/PRODUCT-BLUEPRINT.md
```

The aggregate is built deterministically from the fifteen sources and includes their hashes. It must never be edited directly.

---

## 7. Blueprint-to-build artifacts

### Product definition

```text
Product Blueprint
→ Constitutional Requirements
→ Product Roadmap Milestone
```

### Uncertainty reduction and reusable design

```text
Focused Research
→ Architecture Spike when evidence must be executable
→ ADR
→ Capability Spec and Test Plan
```

### Scoped commitment and execution

```text
Mission Contract
→ Microdesign / Implementation Plan
→ Code / Configuration / Schema
→ Claim
→ Receipt / Evidence
→ Verdict
→ Product Milestone Closeout
```

The full ownership and gates are defined in:

```text
docs/product/CAPABILITY-REALIZATION-METHOD.md
```

A plan cannot change a Capability Spec. A Contract cannot silently override an ADR. A test result does not alter the Product Blueprint. Material changes return to the owning artifact and readiness gate.

---

## 8. Research system

Research lives under:

```text
docs/research/
```

Each material report contains:

- exact question and scope;
- date and versions;
- source-selection method;
- primary-source findings;
- alternatives and disagreements;
- evidence limits;
- implications for Aurora;
- architecture spike requirements;
- decision candidates;
- matching `*.sources.json` manifest.

Current thematic areas include:

- memory and context;
- Harness interoperability;
- AHDK, conformance and Golden Paths;
- durable execution;
- authority, identity and effects;
- events, observability and schemas;
- agent frameworks and runtimes;
- evaluation and self-improvement;
- the original synthesis report.

Research is refreshed when a material specification/version changes, a spike contradicts it, implementation approaches, or the freshness threshold is exceeded.

---

## 9. ADR system

Canonical paths:

```text
docs/adr/README.md
docs/adr/template.md
docs/adr/0001-*.md
```

An ADR owns one decision and records:

- context and problem;
- decision drivers;
- researched alternatives;
- decision;
- implications and trade-offs;
- implementation constraints;
- validation/spikes;
- migration/rollback where applicable;
- reconsideration triggers;
- supersession history.

`PROPOSED` ADRs do not govern implementation. Acceptance requires explicit operator review under the current phase gate.

---

## 10. Design, spikes and implementation plans

### System/capability design

```text
docs/design/
docs/superpowers/specs/
```

### Architecture Spikes

Spikes investigate uncertainty. They must declare disposal/promotion rules; experimental code is not production by default.

### Implementation plans

```text
docs/superpowers/plans/
```

A plan is executable only when:

- its governing Blueprint/requirements are current;
- applicable Capability Spec and ADRs are accepted;
- exact Contract is approved;
- the design was adversarially reviewed;
- STATUS explicitly authorizes execution.

---

## 11. Tracking and historical material

### Tracking

```text
docs/tracking/STATUS.md
→ current phase, authorization, evidence, blockers and next action

docs/tracking/WORKLOG.md
→ chronological history of material work

docs/tracking/DECISIONS.md
→ index pointing to canonical owners

docs/tracking/BACKLOG.md
→ ideas without commitment

docs/tracking/DOCUMENTATION-COVERAGE.md
→ discovery-to-canonical coverage
```

Tracking never owns product architecture.

### Historical discovery

```text
docs/history/2026-08-05-aurora-origin-and-discovery-record.md
```

This preserves the original motivation, scenarios, alternatives and approvals. It helps answer “why did this concept emerge?” but current meaning belongs to canonical owners.

### Reviews

```text
docs/reviews/
```

Reviews record findings against fixed scope/revision. They are evidence and do not automatically update normative content.

---

## 12. Storage boundaries

## Git

Canonical human-readable product knowledge:

- Blueprint;
- requirements;
- ADRs;
- Capability Specs;
- Contracts intended to travel with source;
- methods and standards;
- research;
- designs/plans;
- selected evidence and closeouts;
- generated documentation sources and projections;
- tracking.

## Future operational store

Canonical current runtime state:

- Missions and Delegations;
- grants and budgets;
- provider instances and trust status;
- checkpoints and events;
- effects and receipts;
- artifacts/evidence references;
- incidents/findings;
- device and Presence state.

Markdown must not be used to simulate live operational truth.

## Future Artifact/Evidence Store

Potentially large or generated data:

- logs and traces;
- prompts and model records under policy;
- datasets/eval results;
- screenshots/recordings;
- firmware and binaries;
- measurements and waveforms;
- tool outputs;
- temporary evidence.

Repository-owned evidence is promoted by verifiable references/hashes when necessary.

## GitHub

- Issue: discussion and work container;
- PR: proposed change, review and checks;
- merged/source files: repository state;
- PR comment: not the canonical result by itself.

## Session/transcript

- supports discovery and continuity;
- may preserve exact conversational history;
- does not govern the product until promotion.

---

## 13. Human read paths

### New reader

```text
README
→ Documentation Map
→ Product Index
→ Product Vision
→ Roadmap
```

### Full A0 reviewer

```text
STATUS
→ Discovery History
→ complete generated Product Blueprint
→ Requirements Traceability
→ thematic Research Map and reports
→ ADRs
→ ACRM
→ Documentation Coverage
→ Adversarial Review
```

### Capability designer

```text
STATUS
→ current Product Milestone
→ ACRM R0–R4
→ applicable Blueprint sections
→ applicable requirements
→ research/spikes/ADRs
→ Capability Spec template and test plan
```

### Implementation reviewer

```text
STATUS
→ accepted Capability Spec
→ accepted ADRs
→ exact Mission Contract/revision
→ approved Microdesign/Plan
→ fixed diff/commit
→ criteria and Evidence Bundle
```

---

## 14. Agent read paths

### New Lead/session

```text
AGENTS.md
→ STATUS
→ Documentation Map
→ immediate next action
→ only relevant canonical sources
```

### Research actor

```text
Research Map
→ owning Blueprint requirement
→ current sources/manifests
→ exact question and freshness window
→ report with limitations
```

### Planner

```text
ACRM
→ current gate
→ applicable requirements
→ accepted decisions/specs/contracts
→ unresolved findings
```

### Implementer

```text
exact Contract
→ approved Microdesign/Plan
→ allocated requirements/criteria
→ repository commands and interfaces
```

### Reviewer/QA

```text
fixed target revision
→ criteria and proof requirements
→ independent context
→ relevant standards/threat model/evals
```

Agents must not load the entire Blueprint by default when a smaller authoritative Context Pack is sufficient. Full reading is required for constitutional/A0 review or when cross-domain impact is material.

---

## 15. Document lifecycle

General normative lifecycle:

```text
DRAFT
→ PROPOSED
→ ACCEPTED
→ SUPERSEDED | REJECTED | WITHDRAWN
```

Research:

```text
CURRENT
→ STALE | HISTORICAL | SUPERSEDED
```

Tracking:

```text
CURRENT
→ ARCHIVED
```

Generated projection:

```text
GENERATED
```

Acceptance requires the authority named by the document and current phase. Merge alone does not silently transform `proposed` into `accepted`.

---

## 16. Documentation generation and validation

Canonical commands:

```bash
python scripts/generate_docs.py
python scripts/generate_docs.py --check
python scripts/validate_docs.py --generated-root <generated-root> --report <report-path>
```

CI checks include:

- required file coverage;
- frontmatter shape;
- stable/unique IDs;
- related-ID resolution;
- internal links;
- research source manifests and citations;
- requirement uniqueness and coverage threshold;
- unresolved documentation gaps;
- normative placeholder detection;
- generated Blueprint/Roadmap freshness.

Future checks may add:

- anchors and section IDs;
- source-to-requirement coverage;
- lifecycle/status compatibility;
- owner existence;
- ADR numbering/index consistency;
- supersession graph;
- Architecture Spike and Capability Spec schemas;
- documentation-impact declarations;
- fresh-session Golden Proof automation.

---

## 17. Documentation impact

Every future material change declares:

```yaml
documentation_impact:
  status: NONE | UPDATED | FOLLOW_UP_REQUIRED
  affected: []
  rationale: ""
  follow_up: null
```

A material change cannot use `NONE` without a specific rationale.

Examples:

- schema or lifecycle change → Spec/Reference/generated types;
- new effect → authority/security docs and tests;
- implementation change affecting behavior → Capability Spec/Contract/Reference;
- new finding → research, ADR or replan;
- accepted milestone → evidence, closeout, roadmap and status;
- device/safety change → threat model, protocol and drill evidence.

---

## 18. Fresh-session continuity proof

A session with no access to the discovery chat must be able to read only the repository and correctly state:

1. what Aurora is and is not;
2. the North Star and first deep domain;
3. the Human–Aurora relationship and authority hierarchy;
4. how memory, capabilities, Harnesses, devices and evidence relate;
5. which technical mechanisms remain open;
6. the current phase and exact prohibitions;
7. the proposed ADRs and research evidence;
8. the next action required to close A0.

It must not:

- start implementation;
- treat research candidates as decisions;
- make MNFS the architectural center;
- infer acceptance from file existence;
- invent missing stack choices.

This is the A0 documentation Golden Proof.

---

## 19. Immediate next action

```text
complete mechanical validation
→ refresh adversarial review against the repaired package
→ execute fresh-session Golden Proof
→ present A0 for operator review
```

Implementation remains prohibited until A0 is explicitly accepted and the first Product Milestone passes its own ACRM readiness gates.
