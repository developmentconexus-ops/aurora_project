---
id: DESIGN-AURORA-TA-03-CROSS-SYSTEM-OPERATION-SURFACE
title: Aurora Cross-System Operation Surface
document_type: system_architecture_design
form: reference
authority: design
status: proposed
version: 0.1.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - proposed TA-03 cross-system operation admission and mutation law
related:
  - DOC-AURORA-REPOSITORY-ROADMAP
  - DOC-AURORA-PLANNING-READINESS
  - DESIGN-AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY
  - DOC-AURORA-TA-03-DESIGN-REVIEW
  - DOC-AURORA-BLUEPRINT-14
source_revision: d151dd2e2bd069af2f1dbffd9608d3b7ce878022
review_triggers:
  - a consumer requires an unadmitted material cross-system operation
  - operation ownership or a trust or effect boundary changes
  - a generic payload bypasses owner validation
  - a deferred consumer becomes a prerequisite
  - a material operation lacks a failure or proof path
last_reviewed: 2026-09-05
---

# TA-03 — Cross-System Operation Surface

## 1. Scope and lifecycle

This revision records the operator-reviewed **Section 1: admission model and cross-system mutation law**. The [design-review Evidence](../evidence/ta-03-design-review.md) records the exact approval scope.

The document remains proposed: partial design review is not complete TA-03 ratification or canonical promotion. [The repository roadmap](../roadmap.md) alone owns current permission, blockers and next action.

This revision does not admit a complete operation catalogue. It cannot be consumed as TA-04-ready operation contracts. It selects no schema, protocol, transport, SDK, framework, store, IAM mechanism, model or provider, and authorizes no Product implementation.

Accepted Product meaning, [TA-01/TA-02 ownership and topology](module-runtime-topology.md), and [Planning Readiness](../development/planning-readiness.md) remain upstream constraints.

## 2. TA03-D01 — Admission unit and granularity

Use named semantic operations grouped into non-callable families:

```text
real consumer / protected property
→ material boundary crossing
→ named operation
→ one semantic owner
→ explicit consequences and failure law
→ separately authorized TA-04 projection
```

Admit an operation only when it crosses a material owner, runtime/provider, actor/Presence, effect/credential, external-system or durable cross-owner boundary. An ordinary private owner method remains private.

A proposed operation must name its real consumer/protected property and explain why its semantics must be stable before the consuming work. A future possibility alone is not sufficient. Safe deferral requires an owner, preserved invariant and revisit trigger.

### Alternatives and outcome

| Approach | Disposition | Basis |
|---|---|---|
| Exhaustive CRUD per owner | Rejected as the semantic architecture | Storage/entity shape does not establish a real consumer or distinguish proposal, authority and commitment. |
| Universal operation type plus arbitrary payload | Rejected as Aurora semantic authority | An unbounded semantic dispatcher can conceal missing operations, owners and failure rules. |
| Named semantic operations; families for navigation only | Operator-approved Section-1 direction | Preserves differentiated meaning without globalizing private methods. |

This does not decide how future transport or application machinery dispatches requests. A shared mechanism may implement accepted operations; it cannot introduce unadmitted semantics or become a parallel authority. No physical CommandBus/broker/framework selection is made here.

## 3. Operation classes and canonical mutation

| Class | Meaning | Authority consequence |
|---|---|---|
| READ | Consume an owner-governed truth or projection. | No requested domain-state mutation; a required access/audit record remains C12-owned. |
| PROPOSAL | A non-owner submits a candidate or request. | The caller cannot commit the target meaning. The receiving owner may validate and record it under its own rules. |
| COMMAND | Request the semantic owner to perform its own transition. | Only that owner may validate and commit the authorized transition. |
| EFFECT | Request a governed consequential action. | C04 owns Effect Request/Policy Decision meaning; E01 enforces execution. Target state is not Aurora state. |
| CALLBACK / OBSERVATION | Supply attributable externally produced material. | Arrival is not canonical truth; the receiver validates and records only its own meaning. |

A read's audit obligation is not permission to mutate Project, Mission or Authority. Similarly, storing a received proposal/observation in its owning domain is not permission for its producer to commit another owner's state.

The accepted mutation path remains:

```text
request / proposal / reported input
→ application coordination
→ exact owner validation
→ C04 policy when required
→ owner commits its own state
→ C12 append when required
→ C07 Evidence/Receipt when required
→ post-commit projection / notification
```

The path expresses ownership and semantic obligations, not a selected transaction protocol or an assumption that external effects share one transaction.

## 4. No arbitrary mutation route

Natural-language output, a model classification, provider OUTPUT_PROPOSAL, a manifest, telemetry, a transport message, a UI action or a generic payload cannot target arbitrary canonical state.

Material input must map to an admitted named operation and reach its exact owner. Without that mapping it is rejected, escalated or retained only as a non-authoritative artifact/claim/observation through the appropriate owner.

Permission to call does not transfer authority to commit. A provider's claim that a Project changed is not a Project transition.

## 5. Semantic owner, state owner and executor

Each operation identifies one semantic owner. Where relevant it separately names the state consequence owner, coordinator, executor/enforcement boundary, receipt producer and record custodian. Those roles do not create multiple owners for the same meaning.

### Effect boundary

```text
Effect Request / Policy Decision meaning → C04
execution / enforcement / receipt production → E01
canonical Receipt record → C07
attributable exact history → C12
external live state → target system
```

E01 cannot create a grant or authorize itself.

### Provider boundary

```text
cross-system contract semantics → G01
accepted provider-boundary profile → B01
global Delegation / Attempt / Outcome → C03
provider-local execution state → P01 / P02
transport / dispatch mechanics → A03
provider-process lifecycle policy → A05
```

B01 is a profile, not a second owner beside G01. G01's contract governance does not acquire runtime business-state ownership.

### Decision Request

C03 owns pending Decision Request lifecycle and routing, not the authority of the requested decision. Resolving the request records the appropriate decision/reference and orchestration consequence. Any resulting Project, policy, approval or other canonical change still follows its own owner.

## 6. Conditional metadata, not a universal envelope

Material operations specify applicable logical request identity, actor/subject/executor/Presence references, Interaction/Mission/Delegation purpose references, semantic contract revision, resource/action references, state/freshness preconditions, authority/policy references, deadline/cancellation, idempotency, correlation/causation, data classification and C12/C07 obligations.

These are conditional semantic requirements, not one mandatory physical record. A Mission or Delegation must not be fabricated merely to fill a shared envelope. Exact profiles and projections remain subject to the appropriate admission and later contract gates.

## 7. Temporal and proof obligations

Require stable logical identity when duplicate delivery could repeat a material consequence. Preserve B01's distinction between a logical request, its execution Attempt and redelivery; do not universalize its field names across unrelated owner operations.

Require a semantic revision/current-state precondition when stale application can violate correctness. Do not impose one global optimistic-lock mechanism.

Cancellation is an intent, not proof of stopped execution or undone effects. Acknowledgement, snapshot and reconciliation establish the applicable conclusion. Authority revocation is separate from cooperative cancellation and must have operational effect at the enforcement boundary.

Process death, missing delivery, provider success and elapsed deadline do not independently establish an Aurora Outcome. Unknown external completion blocks blind retry. Required audit/Evidence failure cannot be hidden as success; the owning operation must specify fail-closed or explicit integrity/reconciliation consequences.

Architectural counterexamples support semantic design. Executable enforcement, external behavior, restart and integration claims require later authorized proof matching the claim.

## 8. Granularity corrections preserved from Section 1

Memory candidate submission, promotion, supersession and deletion have different authority/lifecycle/deletion consequences and must not be hidden in an arbitrary ManageMemoryItem verb.

Provider readiness, drain, stop and process reconciliation preserve A05's existing boundaries; a generic ControlProviderRuntime command must not silently create a global work transition. Process restart cannot create a new C03 Attempt by itself.

These examples preserve the reviewed granularity decision. They do not substitute for complete per-operation admission records.

## 9. Next inventory boundary

The operation inventory must supply, for each admitted entry, all applicable Planning Readiness §10 fields: identity, consumer/caller, semantic owner, class, semantic inputs/outputs, authority/preconditions, state consequence, failure/ambiguity, concurrency/idempotency, cancellation/deadline, history, Evidence/receipt and forbidden shortcuts.

No collection of names or family table is treated as those contracts. First-consumer assignments must be demonstrated against accepted prerequisites, not inferred from an intended milestone name.

The design-review Evidence records TA03-F01, which challenges the proposed M1 provider-path admission. Its owning resolution must precede admission of that disputed path; current scheduling and permission remain in the repository roadmap.

## 10. Reopen and acceptance boundary

Reopen the smallest implicated authority when a consumer lacks a material operation, an owner is missing/duplicated, a required prerequisite is absent, or an operation changes Product/authority/effect meaning outside its scope.

Retain valid Section-1 decisions while that finding is resolved. Neither a section approval nor a green documentation workflow is complete TA-03 acceptance. Complete operation review, proportional independent challenge, finding adjudication, operator ratification and separate merge authorization remain required before canonical promotion.
