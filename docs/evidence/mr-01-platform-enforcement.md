---
id: DOC-AURORA-MR-01-RM08-PLATFORM-ENFORCEMENT
title: MR-01 Platform Enforcement Evidence
document_type: repository_platform_enforcement_evidence
form: reference
authority: evidence
status: accepted
version: 1.1.0
owners:
  - developmentconexus-ops
related:
  - DOC-AURORA-ENGINEERING-RULES
  - DOC-AURORA-REPOSITORY-ROADMAP
  - DOC-AURORA-MR-01-MIGRATION-AUTHORIZATION
last_reviewed: 2026-08-23
---

# MR-01 — RM-08 Platform Enforcement Evidence

## Target

RM-08 requires server-side GitHub enforcement for canonical `main`:

```text
PR-based integration required
required aggregate repository check: validate
force-push forbidden
branch deletion forbidden
squash retained as normal merge method
merge-commit/rebase disabled
automatic head-branch deletion enabled
```

This is platform state, not a documentation/CI convention.

## Prior blocked state

Before operator platform configuration, GitHub reported:

```text
main: 35614c581cea32e04305c1ad63522fee151eb283
protected: false
required status checks: off / none
allow_squash_merge: true
allow_merge_commit: true
allow_rebase_merge: true
delete_branch_on_merge: false
```

Connected-tool and `GITHUB_TOKEN` administrative mutation attempts were exhausted and failed with HTTP 403 because the integration lacked the required repository Administration mutation scope. Those probe runs remain historical Evidence:

```text
32652270369
32652317828
```

## Operator-applied platform configuration

On 2026-08-23 the operator completed the exact GitHub Ruleset/repository-settings procedure supplied for RM-08. The configured target was:

```text
target branch: main
enforcement: Active
PR-based integration: required
required check: validate (Documentation workflow job)
branch must be up to date before merge: required where available
bypass list: empty
force pushes: blocked
deletions: restricted
required human approvals: 0
squash merge: enabled
merge commits: disabled
rebase merge: disabled
automatic deletion of merged head branches: enabled
```

The connected GitHub API does not expose the Ruleset object/rules list through its available read surface, so individual Ruleset fields cannot be independently enumerated here. The operator configuration above is therefore paired with independently observable effective repository state below rather than represented as connector-derived fields.

## Post-configuration machine verification

Fresh GitHub repository/branch metadata after the operator action reported:

```text
main SHA:
35614c581cea32e04305c1ad63522fee151eb283

main protected:
true

repository merge policy:
allow_squash_merge: true
allow_merge_commit: false
allow_rebase_merge: false
delete_branch_on_merge: true
```

The classic branch-protection summary embedded by the branch endpoint continues to report its legacy protection subobject as disabled/off. This does not negate the top-level `protected: true`; RM-08 was configured through GitHub Rulesets rather than classic branch protection, and the connected read surface does not enumerate Ruleset internals.

The migration PR remained isolated and unmerged:

```text
PR: #8
state: OPEN / DRAFT / NOT MERGED
head: fed6d962f24c745dee2167507b1cc061eae935f9
base main: 35614c581cea32e04305c1ad63522fee151eb283
```

Fresh exact-head repository validation:

```text
workflow: Documentation
job/check: validate
run: 32652516829
result: SUCCESS
```

No canonical `main` commit changed while platform enforcement was applied.

## Verdict

```text
RM-08: PASS
server-side main protection: PRESENT
operator-applied PR/status-check/force-push/deletion rules: ATTESTED AGAINST EXACT CONFIGURATION PROCEDURE
observable main protected flag: TRUE
exact candidate check: validate — SUCCESS
squash-only repository merge policy: VERIFIED
auto-delete merged head branches: VERIFIED
main drift: NONE
PR #8 merged: NO
```

RM-08 Evidence is sufficient to proceed to the isolated RM-10 independent review gate. It does not authorize merge, TA-03+, Product/runtime implementation, M0 R7/R8 or Architecture Spike execution.
