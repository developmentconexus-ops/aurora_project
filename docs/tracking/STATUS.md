---
id: DOC-AURORA-STATUS
title: Aurora Project Status
document_type: project_status
form: reference
authority: tracking
status: current
version: 0.1.1
owners:
  - developmentconexus-ops
related:
  - DOC-AURORA-DOCUMENTATION-MAP
  - DOC-AURORA-ROADMAP
  - DESIGN-AURORA-CAPABILITY-HARNESS-ARCHITECTURE
last_reviewed: 2026-08-05
---

# Project Status

## Summary

- **Project:** Aurora
- **Repository:** `developmentconexus-ops/aurora_project`
- **Phase:** A0 — Product and Architecture Baseline
- **Working branch:** `docs/architecture-baseline`
- **Draft pull request:** #1 — product and harness architecture baseline
- **Implementation:** not started
- **Current proposal:** Capability and Harness Architecture
- **Research baseline:** Harness Architecture Research v1
- **Stack decisions:** none
- **MNFS integration:** deferred; MNFS is a future provider, not foundation

## Authorization boundary

```text
Documentation and research: AUTHORIZED
Design proposal: AUTHORIZED
Architecture spikes: PROPOSED, NOT AUTHORIZED
Implementation plan: NOT STARTED
Aurora Core implementation: PROHIBITED
Harness SDK implementation: PROHIBITED
MNFS integration: PROHIBITED
```

## Approved conversational direction captured

- Leandro-first and single-user now.
- Personal intelligence with engineering as first deep domain.
- Trusted intellectual copilot.
- Stable, expressive and transparent identity.
- Contextual proactivity.
- Progressive authority and delegated autonomy.
- Causal, supervised self-improvement.
- Governed multiscoped memory.
- Local-first, cloud-assisted sovereignty.
- Persistent Core and multi-presence continuity.
- Aurora as global control plane.
- Harnesses as specialized providers.
- Hierarchical orchestration.
- Governed manifests and trust.
- First-party AHDK + universal conformance.

These become canonical only after operator review and merge.

## Proposed ADRs

- ADR-0001 — Aurora-owned Contract Model and Replaceable Bindings.
- ADR-0002 — First-party Harness Development Kit and Universal Conformance.

## Open decisions

- Core language;
- first AHDK language;
- schema split;
- local RPC;
- durable engine;
- policy engine;
- workload identity;
- storage;
- event transport;
- artifact/evidence stores;
- reference harness runtime.

## Blockers

No technical blocker. The current gate is operator review of draft PR #1.

## Immediate next action

1. Review PR #1, especially the design spec and proposed ADRs.
2. Request corrections or explicitly approve the baseline.
3. After approval, promote statuses and write a detailed plan for A0 closeout / architecture spikes.
4. Do not implement before a separate plan approval.
