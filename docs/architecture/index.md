---
id: DOC-AURORA-ARCHITECTURE-INDEX
title: Aurora Architecture Index
document_type: architecture_index
form: reference
authority: reference
status: current
version: 1.1.0
owners:
  - developmentconexus-ops
related:
  - DESIGN-AURORA-SYSTEM-ARCHITECTURE-REBASELINE
  - DESIGN-AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY
  - DESIGN-AURORA-TA-03-CROSS-SYSTEM-OPERATION-SURFACE
  - DOC-AURORA-PLANNING-READINESS
last_reviewed: 2026-09-05
---

# Aurora Architecture Index

| Question | Current owner |
|---|---|
| Product-level logical architecture/invariants | `docs/product/blueprint/12-system-architecture.md` |
| System-architecture rebaseline/boundaries | `system-architecture.md` |
| Canonical modules, data ownership, dependency direction, Stage-A/B runtime topology | `module-runtime-topology.md` |
| Cross-system operation admission and mutation law | [Cross-System Operation Surface](cross-system-operation-surface.md) — proposed design; partial operator review is not canonical promotion |
| Stage-A availability / activation / locked-workstation constraints | `stage-a-availability-activation.md` |
| Open cross-system architecture decision landscape | `system-decision-landscape.md`, interpreted through current Planning Readiness |
| Capability/Harness responsibility boundary | `capability-harness-boundary.md` + Blueprint 07 |
| Cross-system remaining stage order | `../development/planning-readiness.md` |
| M0-specific old microdesign/spikes | `../phases/m0/` — Evidence/history, not target architecture by existence |
| Architecture spike portfolio | `../reference/architecture-spikes.md` |

Old Technical Architecture Baseline ordering is preserved as a superseded snapshot under `../phases/technical-architecture-baseline-map.md`; TA-01/TA-02 results remain current.
