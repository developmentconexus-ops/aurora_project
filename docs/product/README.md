---
id: DOC-AURORA-PRODUCT-INDEX
title: Aurora Product Documentation
document_type: product_index
form: reference
authority: constitutional
status: accepted
accepted_at: 2026-08-23
acceptance_evidence: DOC-AURORA-MR-01-OPERATOR-RATIFICATION
version: 0.3.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - product documentation entrypoint
  - Product Blueprint read order
  - blueprint-to-build entrypoints
related:
  - DOC-AURORA-DOCUMENTATION-MAP
  - DOC-AURORA-PRODUCT-BLUEPRINT
  - DOC-AURORA-BLUEPRINT-14
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DOC-AURORA-REQUIREMENTS-TRACEABILITY
last_reviewed: 2026-08-23
---

# Aurora Product Documentation

The Product Blueprint defines **what Aurora is**, which invariants survive implementation, and how the long-horizon Product vision remains distinct from current technical commitment.

The repository-program state is not owned here. For the current gate, implementation permission, blockers and exact next action use [`../roadmap.md`](../roadmap.md).

## 1. Two horizons

```text
Long-term Product vision
→ complete, constitutional and evolutionary

Current executable horizon
→ progressively researched, specified, contracted and proven
```

The vision is not reduced to an MVP. The implementation is not allowed to pretend distant technical choices are already known.

## 2. Canonical Product Blueprint sources

| Section | Canonical source | Governs |
|---:|---|---|
| 01 | [Product Vision](blueprint/01-product-vision.md) | definition, North Star, scope, principles and success |
| 02 | [Human–Aurora Relationship](blueprint/02-human-aurora-relationship.md) | relationship, personality, proactivity and trust |
| 03 | [Domain and World Model](blueprint/03-domain-world-model.md) | entities, identities, relationships, epistemic and temporal model |
| 04 | [Cognitive Lifecycle and Journeys](blueprint/04-cognitive-lifecycle-journeys.md) | cognitive loop and end-to-end scenarios |
| 05 | [Capability System](blueprint/05-capability-system.md) | capabilities, providers, manifests, trust, AHDK and conformance |
| 06 | [Memory, Knowledge and Context](blueprint/06-memory-knowledge-context.md) | memory strata, promotion, supersession, Context Builder and evaluation |
| 07 | [Harness Orchestration](blueprint/07-harness-orchestration.md) | Aurora–Harness boundary, Delegations, artifacts and recovery |
| 08 | [Interaction, Multimodality and Presence](blueprint/08-interaction-multimodality-presence.md) | Presence Fabric, surfaces, handoff and degraded operation |
| 09 | [Tools, Devices and Laboratory](blueprint/09-tools-devices-laboratory.md) | devices, instruments, telemetry and physical progression |
| 10 | [Autonomy, Authority and Safety](blueprint/10-autonomy-authority-safety.md) | authority, effects, interlocks, revocation and emergency boundaries |
| 11 | [Security, Privacy and Sovereignty](blueprint/11-security-privacy-sovereignty.md) | local-first control, data classes, identity, credentials and threat model |
| 12 | [System Architecture](blueprint/12-system-architecture.md) | logical architecture, state ownership, topology and evolution |
| 13 | [Reliability, Observability and Evaluation](blueprint/13-reliability-observability-evaluation.md) | evidence, evals, traces, incidents and self-improvement |
| 14 | [Product Capability Roadmap](blueprint/14-capability-roadmap.md) | Product Milestones, Golden Proofs and replan triggers |
| 15 | [Documentation and Research Governance](blueprint/15-documentation-research-governance.md) | authority, lifecycle, promotion, repository memory and continuity |

## 3. Generated publication versus repository roadmap

- [Complete Product Blueprint publication](PRODUCT-BLUEPRINT.md) is a generated read-only projection of Blueprint 01–15.
- [Blueprint 14](blueprint/14-capability-roadmap.md) is the Product capability-roadmap authority.
- [`docs/roadmap.md`](../roadmap.md) is deliberately different: it is the hand-maintained repository-program stage/status/permission/next-action authority.

Generate/check only the Product aggregate:

```bash
python scripts/generate_docs.py
python scripts/generate_docs.py --check
```

Never edit the Product aggregate directly and never generate `docs/roadmap.md` from Blueprint 14.

## 4. Blueprint-to-build method

- [Aurora Capability Realization Method](../development/capability-realization.md)
- [Aurora Planning and Implementation-Readiness](../development/planning-readiness.md)
- [Constitutional Requirements and Traceability](REQUIREMENTS-TRACEABILITY.md)

Global cross-system readiness closes foundational ambiguity; ACRM R0–R8 realizes one selected Product Milestone / Capability / Mission without replaying the whole global architecture program.

```text
Blueprint / accepted architecture
→ cross-system readiness authority
→ applicability / requirements
→ research / spikes / decisions
→ Capability Spec
→ Mission Contract
→ implementation-design readiness
→ separately authorized execution
→ Receipts / Evidence / Verdict
→ Product Milestone closeout
```

No artifact silently substitutes another. Research informs but does not decide; green implementation is not acceptance; a local component is not a Product Golden Proof.

## 5. Research and decisions

- [Research Map](../research/RESEARCH-MAP.md)
- [Decision Register](../decisions/index.md)
- [ADR Index](../decisions/adr/README.md)
- [Architecture Spike Portfolio](../reference/architecture-spikes.md)

Technical mechanisms remain open until their owning readiness stage or Capability R4 decision is satisfied. Research candidates and framework features are not stack choices by existence.

## 6. Historical, Evidence and traceability material

- [Origin and Discovery Record](../history/2026-08-05-aurora-origin-and-discovery-record.md)
- [A0 Discovery/Documentation Coverage Evidence](../evidence/a0-documentation-coverage.md)
- [A0 Adversarial Documentation Review](../evidence/reviews/2026-08-05-a0-adversarial-documentation-review.md)

History and Evidence preserve motivation, provenance and observed proof. Current Product meaning remains with its current owner.

## 7. Recommended Product read paths

### Understand the Product

```text
01 Product Vision
→ 02 Human–Aurora Relationship
→ 03 Domain and World Model
→ 04 Cognitive Lifecycle
→ 14 Product Capability Roadmap
```

### Understand the cognitive core

```text
03 Domain and World Model
→ 06 Memory, Knowledge and Context
→ 12 System Architecture
→ 13 Reliability and Evaluation
```

### Understand the capability ecosystem

```text
05 Capability System
→ 07 Harness Orchestration
→ 10 Autonomy and Authority
→ 11 Security and Sovereignty
```

### Plan scoped realization

```text
docs/roadmap.md
→ Planning Readiness / ACRM as applicable
→ exact Product / architecture / decision owners
→ exact Capability / Contract
→ proof obligations
```

Do not use this Product index as a second mutable project-status surface.
