---
id: DOC-AURORA-REPOSITORY-ROADMAP
title: Aurora Repository Roadmap
document_type: repository_program_roadmap
form: reference
authority: tracking
status: current
program_status_authority: true
version: 1.1.1
owners:
  - developmentconexus-ops
source_of_truth_for:
  - current repository stage and gate
  - implementation permission and blocked work
  - exact next action
  - progression and reopen triggers
related:
  - DOC-AURORA-MR-01-OPERATOR-RATIFICATION
  - DOC-AURORA-MR-01-MIGRATION-AUTHORIZATION
  - DOC-AURORA-PLANNING-READINESS
  - DOC-AURORA-MR-01-RM08-PLATFORM-ENFORCEMENT
  - DOC-AURORA-MR-01-RM09-COHERENCE
last_reviewed: 2026-08-23
---

# Aurora Repository Roadmap

> This is the sole mutable repository-program status/permission/next-action authority. Product capability sequence remains in Blueprint 14.

## Current canonical baseline

```text
A0 Product constitution: ACCEPTED / MERGED
System Architecture Rebaseline: ACCEPTED / MERGED
TA-01 Logical Modules & Canonical Ownership: ACCEPTED / CANONICAL
TA-02 Process/Runtime/Evolutionary Topology: ACCEPTED / CANONICAL
M0 R0–R6: historical PASS within M0 scope
M0 R7 candidate: FROZEN / PRESERVED / NON-CANONICAL
M0 R7 Verdict: NOT ISSUED
M0 R8: NOT AUTHORIZED
MR-01 Methodology/Repository/Readiness target: OPERATOR-RATIFIED
```

## Current program

```text
program: MR-01 Repository Migration
migration PR: #8
RM-01: PASS — semantic/provenance census; durable Evidence recorded
RM-02: PASS — bounded constitutional/method reconciliation
RM-03: PASS — target authority preparation
RM-04: PASS — decision/architecture/capability reconciliation
RM-05: PASS — atomic repository control-plane cutover
RM-06: PASS — legacy live-surface retirement/rehome
RM-07: PASS — positive validation + deterministic negative controls
RM-08: BLOCKED_BY_PLATFORM_CREDENTIAL_SCOPE
RM-09: PASS — fresh-actor + Global Coherence proof
RM-10: BLOCKED ON RM-08 — final independent review / promotion candidate follows platform enforcement
```

Evidence:

```text
RM-07 Documentation: 32651194229 — SUCCESS
RM-09 Coherence Audit: 32651194244 — SUCCESS
RM-08 observed platform: main protected=false; required checks off
RM-08 API attempt 1: 32652270369 — branch protection HTTP 403
RM-08 API attempt 2: 32652317828 — branch protection HTTP 403; repository settings HTTP 403
RM-08 exact Evidence: docs/evidence/mr-01-rm08-platform-enforcement.md
```

The connected GitHub integration has repository admin visibility, but its complete exposed tool surface does not include branch-protection/ruleset or repository-settings mutation. A one-shot Actions fallback was executed with maximum available `GITHUB_TOKEN` permissions (`write-all`); GitHub rejected both administrative endpoints with `Resource not accessible by integration`. No CI-based pseudo-protection is accepted as a substitute for server-side enforcement.

## Authorization boundary

```text
repository migration execution: AUTHORIZED
TA-03+: NOT AUTHORIZED
Architecture Spike execution: NOT AUTHORIZED
Aurora Product/runtime implementation: BLOCKED
M0 R7 continuation/Verdict/R8: NOT AUTHORIZED
framework/database/IAM/model/provider selection: NOT AUTHORIZED BY MIGRATION
merge: NOT AUTHORIZED
```

## Exact next action

```text
supply/apply GitHub credential or UI action with repository Administration:write for RM-08:
  - require PR-based integration on main
  - require aggregate repository check (`validate`)
  - forbid force-push
  - forbid branch deletion
  - retain squash as normal merge method
  - disable merge commits/rebase where permitted
  - enable automatic head-branch deletion where permitted
→ re-query main protection/settings and record RM-08 PASS Evidence
→ run final Documentation validation on the exact clean candidate
→ prepare isolated RM-10 Fable review
→ adjudicate any review findings
→ STOP before merge authorization
```

The repository/documentation candidate is already free of branch-only `docs/work/**` and temporary RM-08 workflow probes; platform credential scope is the only current blocker.

## Reopen triggers

Reopen the MR-01 target only if migration Evidence shows lost current semantics/provenance, duplicate/missing authority, an unworkable fresh-actor route, or a downstream material decision forced before its owner exists.
