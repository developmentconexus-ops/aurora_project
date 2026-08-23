---
id: DOC-AURORA-DOCUMENTATION-MAP
title: Aurora Documentation Index
document_type: documentation_index
form: reference
authority: standard
status: accepted
accepted_at: 2026-08-23
acceptance_evidence: DOC-AURORA-MR-01-OPERATOR-RATIFICATION
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - task and intention routing
  - documentation authority entrypoints
  - smallest default read packs
related:
  - DOC-AURORA-BLUEPRINT-15
  - DOC-AURORA-REPOSITORY-ROADMAP
  - DOC-AURORA-DECISIONS
  - DOC-AURORA-PLANNING-READINESS
last_reviewed: 2026-08-23
---

# Aurora Documentation Index

This file routes a task to the smallest correct authority set. It does not own mutable program status or Product/architecture semantics.

| Intent | Start here | Do not read by default |
|---|---|---|
| Current gate, blockers, allowed work, exact next action | [`roadmap.md`](roadmap.md) | Product history, Evidence, Git history |
| Product / North Star / constitutional meaning | [`product/README.md`](product/README.md) → exact Blueprint owner | full Blueprint unless cross-domain impact is material |
| Logical modules, owners, Stage-A/B topology | [`architecture/index.md`](architecture/index.md) | M0 microdesign/history |
| Current decisions / open decisions / reopen triggers | [`decisions/index.md`](decisions/index.md) | review chronology |
| Cross-system planning/readiness | [`development/planning-readiness.md`](development/planning-readiness.md) | implementation plans |
| Capability/milestone realization R0–R8 | [`development/capability-realization.md`](development/capability-realization.md) | unrelated global TA stages |
| Repository/Git/review/verification rules | [`development/engineering-rules.md`](development/engineering-rules.md) | organizational Method copies |
| Capability-specific work | exact owner under [`capabilities/`](capabilities/) | other capabilities |
| Research | exact question/report via [`research/RESEARCH-MAP.md`](research/RESEARCH-MAP.md) | broad research corpus |
| Architecture-spike portfolio | [`reference/architecture-spikes.md`](reference/architecture-spikes.md) | spike execution unless authorized |
| M0 frozen design/evidence | [`phases/m0/`](phases/m0/) + exact Evidence | current source architecture inference |
| Acceptance/review proof | exact item under [`evidence/`](evidence/) when a current claim requires it | evidence chronology generally |
| Aurora origin/discovery | [`history/2026-08-05-aurora-origin-and-discovery-record.md`](history/2026-08-05-aurora-origin-and-discovery-record.md) | default task context |

## Authority law

```text
Product Blueprint / accepted structural owner / accepted decision
> Specification / Contract / Standard in its scope
> current machinery Reference
> Evidence for its observation
> research/history
```

`docs/roadmap.md` owns only repository-program progression. `docs/index.md` owns only routing. Neither may silently redefine Product/architecture.
