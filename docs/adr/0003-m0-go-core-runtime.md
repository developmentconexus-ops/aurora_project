---
id: ADR-AURORA-0003
title: Go as the Initial Aurora Core Runtime
document_type: adr
form: explanation
authority: decision
status: proposed
version: 0.1.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - proposed M0 Aurora Core implementation language/runtime
related:
  - DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC
  - DESIGN-AURORA-M0-R4-DECISION-LANDSCAPE
  - RESEARCH-AURORA-M0-RUNTIME-PERSISTENCE-R4-V1
  - DOC-AURORA-M0-R4-OPERATOR-AUTHORIZATION
supersedes: []
superseded_by: null
last_reviewed: 2026-08-07
---

# ADR-0003 — Go as the Initial Aurora Core Runtime

## Context

R3 intentionally left `R4-Q-CORE-001` open. The Core needs a stable local runtime for domain/state/authority/recovery logic but does not need hard-real-time execution, memory-unsafe FFI-heavy systems programming or one language shared by every future Capability.

ADR-0001 already requires Aurora semantics to remain language/framework neutral across bindings.

## Decision drivers

- correctness and maintainability of state/authority code;
- small local operational footprint;
- stable compatibility policy;
- mature relational database integration;
- cross-platform build/test capability;
- observability/security standard-library ecosystem;
- low accidental coupling between Core language and future Harness/device languages;
- migration cost if the runtime is reconsidered later.

Affected requirements include `CAP-SOVEREIGN-CORE-REQ-005`, `006`, `017`, `095`, `096`, `097`, `098` and `107`.

## Evidence

`RESEARCH-AURORA-M0-RUNTIME-PERSISTENCE-R4-V1` compares current Go and Rust using primary/current sources.

The research finds both capable. Rust's ownership model gives stronger compile-time memory-safety guarantees without GC, but M0's dominant risks are transactional state, authority freshness, rollback, restore and evidence correctness. Go provides a stable compatibility promise, standard relational `database/sql`/transaction APIs, broad cross-platform support and lower language/tooling burden for this control/state Core.

## Options

### A — Go

Pros: stable compatibility model; simple local runtime; standard relational API; mature ecosystem; easy independent service/CLI packaging; future Rust components remain possible through contracts.

Cons: garbage collection; less compile-time memory discipline than Rust.

### B — Rust

Pros: ownership/borrowing, memory safety without GC, strong low-level control.

Cons: higher implementation/tooling complexity for a Core whose current main hazards are not memory unsafety; SQLite integration still needs native/binding proof.

### C — defer runtime selection

Would move a material implementation decision into R5/R6 and violate R4's purpose.

## Decision

**Proposed:** implement the initial Aurora Core in **Go**.

Mission/implementation baselines MUST pin a current supported Go patch release rather than treating one patch version as permanent product meaning.

Go does **not** become the language of all Aurora Capabilities. AHDK, Harnesses, device controllers and future specialized components may use other languages behind accepted contracts.

No agent framework is selected as Aurora Core runtime ownership.

## Explicit non-decisions

This ADR does not choose:

- SQLite versus PostgreSQL;
- SQLite binding;
- schema library/ORM;
- process supervision/deployment packaging;
- future AHDK first language;
- firmware/device language.

## Consequences

### Positive

- small, conventional Core implementation surface;
- relational/store choices stay accessible through `database/sql` where appropriate;
- strong compatibility/cross-platform story;
- low incentive to embed agent-framework semantics in Core.

### Negative

- GC exists in Core;
- memory-safety bugs are prevented primarily by Go's safe language/runtime model and review/tests rather than Rust's ownership system;
- some native dependencies may reintroduce CGO depending on selected SQLite binding.

### Risks

The main risk is language convenience encouraging direct SQL/framework types to leak into domain modules. R3 ports/architecture tests remain mandatory.

## Compatibility / migration / rollback

Aurora logical contracts and portable exports remain language neutral. If a future accepted ADR replaces Go, the migration target is the Core port/domain contract rather than Go-internal persistence objects.

## Validation

Documentary R4 evidence is sufficient to review this runtime choice; no standalone language spike is required. Storage bindings still require `SPK-AURORA-M0-SOVEREIGN-STORE-001`.

Acceptance still requires operator review. Acceptance does not authorize implementation.

## Reconsideration triggers

- a current Capability requirement cannot reasonably be satisfied in Go;
- measured Core performance/resource behavior becomes material and unacceptable;
- chosen critical library/platform loses support;
- a later architecture proves a different runtime materially reduces risk/operations without domain rewrite.

## References

- `RESEARCH-AURORA-M0-RUNTIME-PERSISTENCE-R4-V1`
- ADR-0001
- `CAP-SOVEREIGN-CORE/SPEC.md`
