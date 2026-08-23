---
id: DOC-AURORA-MR-01-MIGRATION-AUTHORIZATION
title: MR-01 Repository Migration Authorization
document_type: operator_authorization_evidence
form: reference
authority: evidence
status: accepted
version: 1.0.0
owners:
  - developmentconexus-ops
related:
  - DOC-AURORA-MR-01-OPERATOR-RATIFICATION
last_reviewed: 2026-08-23
---

# MR-01 Repository Migration Authorization

On 2026-08-23 the operator explicitly authorized execution of the MR-01 repository migration after MR-01 semantic ratification.

```text
base main: 35614c581cea32e04305c1ad63522fee151eb283
migration branch: docs/mr-01-repository-migration-20260823
migration PR: #8
integration model: one migration branch → one final merge candidate
```

This authorization covers the repository/documentation/control-plane migration only. It does not authorize TA-03+, Product/runtime implementation, M0 R7/R8, Architecture Spike execution, stack selection or merge.
