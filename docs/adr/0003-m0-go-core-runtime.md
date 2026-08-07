---
id: ADR-AURORA-0003
title: Go as the Initial Aurora Sovereign Core Runtime
document_type: adr
form: explanation
authority: decision
status: accepted
accepted_at: 2026-08-07
acceptance_evidence: DOC-AURORA-M0-R4-ADR-ACCEPTANCE-SPK001-AUTHORIZATION
version: 0.2.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - proposed M0 Aurora Sovereign Core implementation language/runtime
related:
  - DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC
  - DESIGN-AURORA-M0-R4-DECISION-LANDSCAPE
  - RESEARCH-AURORA-M0-RUNTIME-PERSISTENCE-R4-V1
  - RESEARCH-AURORA-MASTRA-COGNITIVE-HARNESS-R4-V1
  - ADR-AURORA-0009
  - DOC-AURORA-M0-R4-OPERATOR-AUTHORIZATION
supersedes: []
superseded_by: null
last_reviewed: 2026-08-07
---

# ADR-0003 — Go as the Initial Aurora Sovereign Core Runtime

## Context

R3 intentionally left `R4-Q-CORE-001` open. The Sovereign Core needs a stable local runtime for canonical domain/state/authority/recovery logic but does not need hard-real-time execution, memory-unsafe FFI-heavy systems programming or one language shared by every future Capability.

ADR-0001 already requires Aurora semantics to remain language/framework neutral across bindings.

The focused Mastra assessment subsequently clarified an important scope boundary: choosing Go for the **Sovereign Core** must not be interpreted as choosing Go as the universal runtime for Aurora's future agentic Harnesses. Mastra is now a proposed preferred TypeScript/Node cognitive/Harness substrate under ADR-0009, behind Aurora-owned contracts.

## Decision drivers

- correctness and maintainability of state/authority code;
- small local operational footprint;
- stable compatibility policy;
- mature relational database integration;
- cross-platform build/test capability;
- observability/security standard-library ecosystem;
- low accidental coupling between Core language and future Harness/device languages;
- clean polyglot boundary to future agentic runtimes such as Mastra;
- migration cost if the Core runtime is reconsidered later.

Affected requirements include `CAP-SOVEREIGN-CORE-REQ-005`, `006`, `017`, `095`, `096`, `097`, `098` and `107`.

## Evidence

`RESEARCH-AURORA-M0-RUNTIME-PERSISTENCE-R4-V1` compares current Go and Rust using primary/current sources.

The research finds both capable. Rust's ownership model gives stronger compile-time memory-safety guarantees without GC, but M0's dominant risks are transactional state, authority freshness, rollback, restore and evidence correctness. Go provides a stable compatibility promise, standard relational `database/sql`/transaction APIs, broad cross-platform support and lower language/tooling burden for this control/state Core.

`RESEARCH-AURORA-MASTRA-COGNITIVE-HARNESS-R4-V1` independently reinforces the value of keeping the Sovereign Core small and framework-neutral while allowing a different runtime to own rich provider-local cognition/execution.

## Options

### A — Go

Pros: stable compatibility model; simple local runtime; standard relational API; mature ecosystem; easy independent service/CLI packaging; clean process/API boundary to TypeScript agent runtimes; future Rust components remain possible through contracts.

Cons: garbage collection; less compile-time memory discipline than Rust.

### B — Rust

Pros: ownership/borrowing, memory safety without GC, strong low-level control.

Cons: higher implementation/tooling complexity for a Core whose current main hazards are not memory unsafety; SQLite integration still needs native/binding proof.

### C — TypeScript/Node/Mastra as the Sovereign Core

Pros: one runtime could reuse Mastra agent/storage/workflow primitives directly.

Cons: would make it easier for framework-local memory/workflow/session semantics to leak into canonical Project/Authority state; weakens independent Core recovery; increases framework lock-in exactly where Aurora requires sovereignty.

Rejected for Sovereign Core ownership. This does not reject Mastra as a Harness runtime.

### D — defer runtime selection

Would move a material M0 implementation decision into R5/R6 and violate R4's purpose.

## Decision

**Proposed:** implement the initial Aurora **Sovereign Core** in **Go**.

Mission/implementation baselines MUST pin a current supported Go patch release rather than treating one patch version as permanent product meaning.

Go does **not** become the language of all Aurora Capabilities. AHDK, Harnesses, agent runtimes, device controllers and future specialized components may use other languages behind accepted contracts.

In particular, ADR-0009 proposes Mastra/TypeScript as the preferred default substrate to evaluate first for first-party agentic Harnesses. That is an intentional polyglot boundary, not architectural inconsistency.

No agent framework is selected as Aurora Sovereign Core runtime ownership.

## Explicit non-decisions

This ADR does not choose:

- SQLite versus PostgreSQL;
- SQLite binding;
- schema library/ORM;
- process supervision/deployment packaging;
- future AHDK first language;
- Mastra implementation/version for a consuming Capability;
- firmware/device language.

## Consequences

### Positive

- small, conventional Core implementation surface;
- relational/store choices stay accessible through `database/sql` where appropriate;
- strong compatibility/cross-platform story;
- low incentive to embed agent-framework semantics in Core;
- clean failure domain between sovereign state and agentic runtime;
- preserves aggressive Mastra reuse without surrendering Core ownership.

### Negative

- Aurora becomes intentionally polyglot when Mastra-backed capabilities begin;
- Go↔TypeScript adapter/contract compatibility becomes an explicit concern;
- GC exists in Core;
- some native dependencies may reintroduce CGO depending on selected SQLite binding.

### Risks

The main risk is language convenience encouraging direct SQL/framework types to leak into domain modules, or future Mastra APIs leaking through the provider adapter into Go domain contracts. R3 ports, ADR-0001 and contract tests remain mandatory.

## Compatibility / migration / rollback

Aurora logical contracts and portable exports remain language neutral. If a future accepted ADR replaces Go, the migration target is the Core port/domain contract rather than Go-internal persistence objects.

Mastra or another Harness runtime can be replaced independently because it is outside the Sovereign Core ownership boundary.

## Validation

Documentary R4 evidence is sufficient to review this Sovereign Core runtime choice; no standalone language spike is required. Storage bindings still require `SPK-AURORA-M0-SOVEREIGN-STORE-001`.

The Mastra assessment does not add a new M0 runtime spike because M0 must function with every Harness/runtime absent.

Acceptance still requires operator review. Acceptance does not authorize implementation.

## Reconsideration triggers

- a current Core requirement cannot reasonably be satisfied in Go;
- measured Core performance/resource behavior becomes material and unacceptable;
- chosen critical library/platform loses support;
- a later architecture proves a different Core runtime materially reduces risk/operations without domain rewrite;
- Go domain contracts begin depending on provider-runtime internals.

## References

- `RESEARCH-AURORA-M0-RUNTIME-PERSISTENCE-R4-V1`
- `RESEARCH-AURORA-MASTRA-COGNITIVE-HARNESS-R4-V1`
- ADR-0001
- ADR-0009
- `CAP-SOVEREIGN-CORE/SPEC.md`
