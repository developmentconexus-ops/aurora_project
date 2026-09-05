---
id: DOC-AURORA-TA-03-DESIGN-REVIEW
title: TA-03 Design Review Evidence
document_type: architecture_design_review_evidence
form: reference
authority: evidence
status: complete
version: 0.1.0
owners:
  - developmentconexus-ops
related:
  - DOC-AURORA-REPOSITORY-ROADMAP
  - DOC-AURORA-TA-03-EXECUTION-AUTHORIZATION
  - DESIGN-AURORA-TA-03-CROSS-SYSTEM-OPERATION-SURFACE
  - DESIGN-AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY
  - DOC-AURORA-PLANNING-READINESS
  - DOC-AURORA-BLUEPRINT-14
last_reviewed: 2026-09-05
---

# TA-03 — Design Review Evidence

This is a frozen record of the operator's Section-1 review and the subsequent bounded admission finding. `complete` describes this evidence record, not completion or acceptance of TA-03. Current program status and next action belong only to [the repository roadmap](../roadmap.md).

## 1. Operator Section-1 review

Recorded on 2026-09-05 from the current operator conversation.

The preceding question was: “Aprova a Seção 1 nesse formato?” The operator answered: “Aprovo”. The subject was **TA-03 Design, Section 1 — admission model and cross-system mutation law**, not the complete operation catalogue or a merge candidate.

Reviewed basis: branch input `8c321e5ada41b6a8c10a0675766ad2fa11197f96` and the Section-1 presentation to the operator. No immutable conversation permalink was supplied; this is the Lead's attributable record of that direct instruction, not an independent review.

The approval covers:

- named semantic operations grouped into non-callable families;
- rejection of storage-shaped CRUD and a universal operation payload as semantic authority;
- one semantic owner per operation, with explicit state/executor/enforcement roles where they differ;
- READ / PROPOSAL / COMMAND / EFFECT / CALLBACK-OBSERVATION distinctions and owner-controlled canonical mutation;
- no arbitrary mutation through provider output, natural language, a manifest, telemetry or a generic payload;
- conditional semantic metadata rather than a mandatory physical envelope;
- proportional idempotency, concurrency and cancellation obligations;
- the presented corrections separating memory lifecycle commands, runtime lifecycle seams, Decision Request lifecycle and the G01/B01 roles.

The approval does not accept Section 2, ratify or close TA-03, authorize an upstream architecture amendment, authorize merge, advance TA-04, execute a Spike, select technology, resume M0 R7/R8 or permit Product/runtime implementation.

The faithful Section-1 design is recorded in [Cross-System Operation Surface](../architecture/cross-system-operation-surface.md). Its document lifecycle remains proposed pending the complete stage review and promotion.

## 2. Fixed inputs and method-version check

Repository inputs revalidated before this record:

```text
main: d151dd2e2bd069af2f1dbffd9608d3b7ce878022
PR: #11 — OPEN / DRAFT / NOT MERGED
candidate input: 8c321e5ada41b6a8c10a0675766ad2fa11197f96
candidate branch: docs/ta-03-cross-system-operation-surface-20260823
```

Aurora's accepted local bindings name Engineering Method v1.0.0 and Repository Standard v1.0.0. The external methodology main now exposes v1.1.0; that is not an automatic Aurora adoption. The accepted v1.0.0 texts were recovered and checked at methodology commit `ac832790b04491418d663960bd095a8ec8801693`:

```text
METHOD.md blob: 18c801c72bf4ec5273f8473b7ed0cf5e629d56ec
REPOSITORY-STANDARD.md blob: 1ea42dc8f57649882178c10ed639e804f42c48ed
```

This is source-resolution Evidence for the existing version binding, not a methodology upgrade or a new organizational pin policy.

Reading exceeded the default five-file pack only to resolve a concrete cross-owner question: whether the proposed first M1 provider invocation can satisfy the accepted C05/G01/A05 prerequisites. The read set included Planning Readiness, TA-01/TA-02, Blueprint 14's M1/M2 sections, the current Decision Index and the existing TA-03 working analysis/self-review. No frozen M0 implementation or broad research corpus was used as target authority.

## 3. TA03-F01 — First cognitive-provider admission prerequisite

### Observation and sources

All Aurora source observations below are bound to the fixed candidate input above; their accepted owner content is unchanged from canonical main.

| Source | Observed constraint |
|---|---|
| [TA-01/TA-02](../architecture/module-runtime-topology.md), §4.2 C05 | C05 owns provider identity/compatibility/trust/approval, but its Stage-A implementation is expressly deferred to M2. |
| Same owner, §4.3 A02 | The accepted cognitive coordination path resolves a provider through C05/G01, obtains A05 readiness, then invokes A03/B01. |
| Same owner, §4.3 A05 and §10 D02 | Dispatch/readiness/start require an approved compatible exact provider instance/build/environment; process readiness is not provider approval. |
| [Blueprint 14](../product/blueprint/14-capability-roadmap.md), §14.9 M1 | M1 includes governed conversation/context/memory and a minimum model adapter. |
| Same Product owner, §14.10 M2 | The full Registry/AHDK/reference-provider capability is a later milestone. |
| TA-03 working inventory at the fixed input | `ResolveProvider`, `EnsureProviderRuntimeReady` and `SubmitProviderRequest` are proposed for the M1/first-cognition flow. These first-consumer assignments are not yet accepted. |

The [Decision Index](../decisions/index.md) preserves the accepted topology and the remaining open runtime/provider choices. No explicit pre-M2 C05 admission exception was found in the consulted current owners. An empty connector search was not treated as proof of repository-wide absence.

### Finding, not a runtime claim

The proposed **M1 use of that provider path** is not yet justified by a coherent accepted prerequisite story:

```text
proposed M1 cognitive invocation
→ C05-approved provider required
→ C05 implementation deferred to M2
→ pre-M2 satisfaction of that requirement is not explicit
```

This is a material planning/admission gap. It does not prove that all possible M1 implementations are impossible, that a particular framework fails, or that every TA-01/TA-02 decision is wrong. A different M1 realization could be valid, but must be supported by accepted authority rather than invented here.

The gap cannot be hidden by treating a configured endpoint, model API key, local process, first-party origin or successful readiness handshake as provider approval.

### Scope of STOP

Disposition at observation: **STOP / SPLIT PREREQUISITE**, limited to admitting the disputed first-consumer provider path and claiming the complete TA-03 catalogue is ready.

Unaffected admission-law work and independent catalogue analysis remain inside the existing TA-03 execution grant. Existing TA-01/TA-02 ownership and topology remain accepted. MR-01 is not reopened. No upstream normative text was changed by this record.

### Smallest owning authority

First return to `docs/architecture/module-runtime-topology.md`, specifically C05's Stage-A disposition in §4.2 read with A02/A05 and D02. A Product Blueprint 14 amendment is necessary only if the chosen resolution actually changes milestone outcomes, sequence or non-goals.

### Credible alternatives

| Alternative | Consequence | Assessment |
|---|---|---|
| Explicit minimum C05 responsibility before the first consumer; full Registry/AHDK/reference-provider realization remains M2 | Preserves the existing owner and approval boundary while making prerequisite availability explicit; may require a bounded upstream clarification/amendment. | Recommended candidate, not accepted. |
| Keep C05 deferred and establish an accepted M1 path not consuming that provider seam | Can preserve current staging only if the complete model/data/authority path is justified without a hidden provider-approval owner or bypass. | Must be demonstrated in the smallest owning architecture decision; not assumed here. |
| Move the full M2 capability before M1 | Broadens Product sequencing and prerequisite work. | Disproportionate unless the smaller alternatives fail; would involve Blueprint 14. |
| Let P01, A05, configuration or a credential substitute for C05 approval | Creates implicit approval or a parallel trust authority. | Rejected against accepted ownership; not a permitted shortcut. |

### Recommended bounded resolution to submit

Clarify that the **minimum C05-owned identity, compatibility and scoped approval responsibility must exist before any first consumer requiring it**. Keep the complete Registry/AHDK/reference-provider milestone at M2. This does not select storage, a service, a framework, an IAM product, a model, a provider or an executable contract.

The exact minimum, admission/proof prerequisites and failure behavior must be stated by the upstream owner; this Evidence record does not prescribe a new bootstrap approval procedure or waive conformance.

### Proof and falsifiers before resumption

A bounded owner review must establish:

1. The first conversational/model-provider flow has a named owner path for exact identity, current compatibility/approval and readiness before dispatch.
2. No configuration, API key, process state or first-party label creates approval by itself.
3. A changed/revoked/unavailable approval blocks new dispatch without requiring the entire M2 product capability to be declared complete.
4. A normal interaction need not be silently promoted into a Mission just to make the provider path executable.
5. M2's operator-visible outcome and its conformance/approval proof obligations remain intact, or any real Product change returns to Blueprint 14.

Static counterexample analysis can support the architecture decision. Runtime identity, revocation, restart and enforcement claims require later separately authorized executable Evidence; none was produced here.

Falsifier of this finding: an already accepted, scoped authority shows how the proposed pre-M2 invocation satisfies C05/A05 without a new owner, an approval bypass or a changed Product commitment. If demonstrated, record NO CHANGE REQUIRED and correct only the TA-03 interpretation.

After the owning resolution is accepted, rebaseline only the affected first-consumer assignments/preconditions in TA-03 and resume catalogue review. Final independent review, stage ratification and merge remain separate gates.
