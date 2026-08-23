---
id: DOC-AURORA-MR-01-RATIFICATION-NOTE
title: MR-01 Ratification Routing Note
document_type: temporary_ratification_note
form: reference
authority: tracking
status: current
version: 0.1.0
owners:
  - developmentconexus-ops
related:
  - DOC-AURORA-MR-01-OPERATOR-RATIFICATION
last_reviewed: 2026-08-23
---

# MR-01 Ratification Routing Note

MR-01 is operator-ratified as a **semantic target**. The durable target documents remain `PROPOSED` on this non-canonical branch until a separately authorized repository migration applies the ratified target, removes temporary `docs/work/**`, and integrates a coherent final candidate to `main`.

This is deliberate lifecycle separation:

```text
operator-ratified target
≠ current main authority
≠ repository migration authorization
≠ TA-03 authorization
≠ Product implementation authorization
≠ merge authorization
```

Exact ratification Evidence: `docs/work/current/operator-ratification.md` and PR #6 discussion.
