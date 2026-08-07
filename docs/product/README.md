---
id: DOC-AURORA-PRODUCT-INDEX
title: Aurora Product Documentation
document_type: product_index
form: reference
authority: constitutional
status: accepted
accepted_at: 2026-08-06
acceptance_evidence: DOC-AURORA-A0-OPERATOR-ACCEPTANCE
version: 0.2.0
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
  - DOC-AURORA-ROADMAP
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DOC-AURORA-REQUIREMENTS-TRACEABILITY
last_reviewed: 2026-08-06
---

# Aurora Product Documentation

The Product Blueprint defines **what Aurora is**, which invariants must survive implementation and how the long-term vision is constrained by authority, safety, evidence and sovereignty.

The Aurora Capability Realization Method defines **how one approved part of that intent becomes an implemented, verified and accepted capability**.

## 1. Two approved horizons

```text
Long-term vision
→ complete, constitutional and evolutionary

Current executable horizon
→ narrow, researched, specified, contracted and proven
```

The vision is not reduced to an MVP. The implementation is not allowed to pretend that distant technical choices are already known.

## 2. Canonical Product Blueprint sources

| Section | Canonical source | Governs |
|---:|---|---|
| 01 | [Product Vision](blueprint/01-product-vision.md) | definition, North Star, scope, principles and success |
| 02 | [Human–Aurora Relationship](blueprint/02-human-aurora-relationship.md) | copilot relationship, personality, proactivity and trust |
| 03 | [Domain and World Model](blueprint/03-domain-world-model.md) | entities, identities, relationships, epistemic and temporal model |
| 04 | [Cognitive Lifecycle and Journeys](blueprint/04-cognitive-lifecycle-journeys.md) | perceive–understand–act–observe loop and end-to-end scenarios |
| 05 | [Capability System](blueprint/05-capability-system.md) | capabilities, providers, manifests, trust, AHDK and conformance |
| 06 | [Memory, Knowledge and Context](blueprint/06-memory-knowledge-context.md) | memory strata, promotion, supersession, context construction and evaluation |
| 07 | [Harness Orchestration](blueprint/07-harness-orchestration.md) | Aurora–Harness boundary, Delegations, events, artifacts and recovery |
| 08 | [Interaction, Multimodality and Presence](blueprint/08-interaction-multimodality-presence.md) | voice/vision/surfaces, Presence Fabric, handoff and degraded operation |
| 09 | [Tools, Devices and Laboratory](blueprint/09-tools-devices-laboratory.md) | devices, instruments, telemetry, protocols and physical safety progression |
| 10 | [Autonomy, Authority and Safety](blueprint/10-autonomy-authority-safety.md) | grants, autonomy envelopes, effects, interlocks, revocation and emergency authority |
| 11 | [Security, Privacy and Sovereignty](blueprint/11-security-privacy-sovereignty.md) | local-first control, data classes, workload identity, credentials and threat model |
| 12 | [System Architecture](blueprint/12-system-architecture.md) | logical components, ports, state ownership, local/cloud topology and evolution |
| 13 | [Reliability, Observability and Evaluation](blueprint/13-reliability-observability-evaluation.md) | evidence, evals, traces, incidents, Failure Intelligence and self-improvement |
| 14 | [Capability Roadmap](blueprint/14-capability-roadmap.md) | cumulative Product Milestones, Golden Proofs and replan triggers |
| 15 | [Documentation and Research Governance](blueprint/15-documentation-research-governance.md) | authority, lifecycle, promotion, validation, storage and session continuity |

## 3. Generated publication

- [Complete Product Blueprint publication](PRODUCT-BLUEPRINT.md)
- [Capability Roadmap projection](../roadmap.md)

These files are generated projections. Edit modular sources and run the documentation generator; never edit projections directly.

```bash
python scripts/generate_docs.py
python scripts/generate_docs.py --check
```

The source files remain canonical because they support focused ownership, review and context loading. The aggregate supports full reading, export and adversarial review.

## 4. Blueprint-to-build method

- [Aurora Capability Realization Method](CAPABILITY-REALIZATION-METHOD.md)
- [Constitutional Requirements and Traceability](REQUIREMENTS-TRACEABILITY.md)

The realization chain is:

```text
Blueprint
→ applicable requirements
→ focused research / Architecture Spikes
→ ADRs
→ Capability Spec
→ Mission Contract
→ Microdesign / Implementation Plan
→ implementation
→ Receipts and Evidence
→ Verdict and Product Milestone Closeout
```

No artifact silently substitutes another:

- research informs but does not decide;
- an ADR decides but does not implement;
- a Capability Spec defines reusable behavior but is not a scoped commitment;
- a Mission Contract commits scope but does not authorize every external effect;
- an implementation claim is not acceptance;
- a local green component is not the Product Milestone Golden Proof.

## 5. Research and decisions

- [Research Map](../research/RESEARCH-MAP.md)
- [ADR Index](../adr/README.md)
- [Architecture Spikes](../design/ARCHITECTURE-SPIKES.md)

Technical mechanisms remain open until the relevant readiness gates are satisfied. Current research candidates must not be read as stack choices.

## 6. Historical and traceability material

- [Origin and Discovery Record](../history/2026-08-05-aurora-origin-and-discovery-record.md)
- [Discovery and Documentation Coverage](../tracking/DOCUMENTATION-COVERAGE.md)
- [A0 Adversarial Documentation Review](../reviews/2026-08-05-a0-adversarial-documentation-review.md)

The history preserves original motivation, examples and decision reasoning. It is not a competing product authority.

## 7. Recommended read paths

### Understand the product

```text
01 Product Vision
→ 02 Human–Aurora Relationship
→ 03 Domain and World Model
→ 04 Cognitive Lifecycle
→ 14 Roadmap
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
→ harness research and ADRs
```

### Understand physical-world evolution

```text
08 Interaction and Presence
→ 09 Tools, Devices and Laboratory
→ 10 Autonomy and Safety
→ 11 Security
→ M8–M10 in the Roadmap
```

### Plan future implementation

```text
STATUS
→ current Product Milestone
→ ACRM R0–R8
→ applicable Blueprint sections
→ requirements
→ research/spikes/ADRs
→ Capability Spec
→ Mission Contract
→ Microdesign and plan
```

## 8. Current authority state

All A0 constitutional sources were explicitly accepted by the operator on 2026-08-06. Merge is a separate repository action and was authorized by the same operator decision.

```text
A0 baseline: ACCEPTED
Constitutional sources: ACCEPTED
ADR-0001 / ADR-0002: ACCEPTED
PR #1 merge: AUTHORIZED
Architecture Spikes: NOT AUTHORIZED
Aurora runtime implementation: PROHIBITED
AHDK implementation: PROHIBITED
MNFS integration: PROHIBITED
```

A merged file is not automatically accepted unless its lifecycle and operator decision say so. A0 acceptance is a separate explicit gate.
