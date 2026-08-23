---
id: DOC-AURORA-MR-01-RM08-PLATFORM-ENFORCEMENT
title: MR-01 RM-08 Platform Enforcement Evidence
document_type: platform_enforcement_evidence
form: reference
authority: evidence
status: blocked
version: 1.0.0
owners:
  - developmentconexus-ops
related:
  - DOC-AURORA-REPOSITORY-ROADMAP
  - DOC-AURORA-MR-01-MIGRATION-AUTHORIZATION
last_reviewed: 2026-08-23
---

# MR-01 RM-08 — Platform Enforcement Evidence

## Target

RM-08 requires server-side GitHub enforcement for the canonical `main` branch:

```text
PR-based integration required
required aggregate repository check
force-push forbidden
branch deletion forbidden
squash retained as normal merge method
merge-commit/rebase disabled where repository settings permit
automatic head-branch deletion enabled where repository settings permit
```

This is platform state, not a documentation/CI convention.

## Observed canonical platform state

Before enforcement attempts, GitHub reported:

```text
main protected: false
protection enabled: false
required status checks: off / none
allow_squash_merge: true
allow_merge_commit: true
allow_rebase_merge: true
delete_branch_on_merge: false
```

## Connected GitHub capability check

The connected GitHub integration has repository admin visibility, but its exposed mutation surface does not include branch-protection/ruleset creation or repository-settings mutation.

A complementary plugin search returned no installable GitHub integration exposing those missing operations.

## One-shot API attempts

To exhaust the remaining safe route, the connected GitHub integration published a temporary GitHub Actions workflow that called the official GitHub REST endpoints using `github.token`.

### Attempt 1

```text
workflow: RM-08 Platform Enforcement
run: 32652270369
branch-protection endpoint:
  PUT /repos/developmentconexus-ops/aurora_project/branches/main/protection
result: HTTP 403
message: Resource not accessible by integration
```

The repository-settings step did not execute because the branch-protection step failed first.

### Attempt 2 — maximum available GITHUB_TOKEN permissions

The workflow was retried with `permissions: write-all` and both administrative calls executed independently.

```text
workflow: RM-08 Platform Enforcement
run: 32652317828

branch-protection endpoint:
  PUT /repos/developmentconexus-ops/aurora_project/branches/main/protection
  result: HTTP 403
  message: Resource not accessible by integration

repository-settings endpoint:
  PATCH /repos/developmentconexus-ops/aurora_project
  requested:
    allow_squash_merge: true
    allow_merge_commit: false
    allow_rebase_merge: false
    delete_branch_on_merge: true
  result: HTTP 403
  message: Resource not accessible by integration
```

The job log confirms `write-all` granted every repository permission available to `GITHUB_TOKEN`, but no `Administration` permission exists in that permission set.

GitHub's branch-protection API requires repository `Administration: write`. GitHub's Actions documentation states that operations requiring permissions unavailable to `GITHUB_TOKEN` need another GitHub App installation token or a personal access token with the required permission.

Official references:

- https://docs.github.com/en/rest/branches/branch-protection
- https://docs.github.com/en/actions/tutorials/authenticate-with-github_token
- https://docs.github.com/en/actions/reference/workflows-and-actions/workflow-syntax

## Verdict

```text
RM-08: BLOCKED_BY_PLATFORM_CREDENTIAL_SCOPE
repository/documentation defect: NO
operator authorization missing: NO
connector repository visibility/admin role: PRESENT
required administrative mutation credential: NOT AVAILABLE TO CURRENT CONNECTED/ACTIONS TOKEN
```

No pseudo-enforcement was substituted. CI cannot replace server-side branch protection because a direct push would already mutate `main` before a post-push workflow could reject it.

RM-10 and canonical promotion remain blocked until the required GitHub platform state is applied and reverified.
