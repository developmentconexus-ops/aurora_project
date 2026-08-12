---
id: REVIEW-AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY-2026-08-12
title: TA-01/TA-02 Module and Runtime Topology Adversarial Review
document_type: architecture_review
form: reference
authority: evidence
status: current
version: 0.3.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - adversarial review of the proposed TA-01 and TA-02 module/runtime topology
  - admitted TA-01 and TA-02 findings and final review verdict
related:
  - DESIGN-AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY
  - DESIGN-AURORA-TECHNICAL-ARCHITECTURE-BASELINE-MAP
  - PLAN-AURORA-TECHNICAL-ARCHITECTURE-BASELINE
  - DOC-AURORA-STATUS
  - DOC-AURORA-WORKLOG
initial_review_target: 2192f182c24eed0c8406062957e035a0db6e88a5
fixed_design_target: 7e9b1b9baefd03a45d8fd0b1e04f0d1981b6cdc9
continuity_target: 99e4c1966bfeb075fa130b9f5078af7a7ac6991c
last_reviewed: 2026-08-12
---

# TA-01/TA-02 Module and Runtime Topology — Adversarial Review

## 1. Review boundary

Initial target:

```text
2192f182c24eed0c8406062957e035a0db6e88a5
```

Fixed semantic design target:

```text
7e9b1b9baefd03a45d8fd0b1e04f0d1981b6cdc9
```

Continuity target after preserving the complete Worklog and removing temporary append machinery:

```text
99e4c1966bfeb075fa130b9f5078af7a7ac6991c
```

The review tested:

1. complete, non-overlapping canonical ownership;
2. framework/provider capture;
3. accidental globalization of M0 choices;
4. God-process and service-sprawl risks;
5. Stage A operational proportionality;
6. Stage B ownership continuity;
7. binding-independent lifecycle and failure semantics across the first process seam;
8. audit/exact-history separation from memory, evidence and telemetry;
9. later-tranche handoff without hidden product selection;
10. repository-only continuity for a future session.

## 2. Reviewed proposal

The fixed target proposes:

```text
G01 Contract Model Governance

C01 Identity and Relationship
C02 Project, World and Experiment State
C03 Mission and Delegation Control
C04 Authority and Policy
C05 Capability and Provider Registry
C06 Memory, Knowledge and Context
C07 Artifact, Observation and Evidence
C08 Presence and Interaction
C09 Attention and Proactivity             [future/deferred]
C10 Failure Intelligence and Evaluation   [future/deferred]
C11 Environment and Device Registry       [future/deferred]
C12 Audit and Exact History
```

It separately classifies:

- application coordinators;
- deployables/composition boundaries;
- deterministic Effect/Credential enforcement;
- cognitive and specialized provider runtimes;
- storage, protocol, UI, model, telemetry and operating-system adapters.

The recommended topology is **Approach C — Evolutionary Sovereign Host with one early provider-runtime seam**:

```text
one small persistent Aurora Sovereign Host
+
one separate on-demand Cognitive Runtime Provider at first consumer
+
independent specialized Harness providers
+
all other physical splits require a named current trigger
```

## 3. Finding disposition

### TA12-F01 — Missing Environment/Device owner

**RESOLVED.** C11 owns Device/Environment identity, inventory and registration relationships; live telemetry, actuation and target truth remain external/source-owned.

### TA12-F02 — Root Identity owner overlapped Provider/Presence/Device identity

**RESOLVED.** C01 owns root Aurora/Person/relationship identity and actor-reference vocabulary. C05, C08 and C11 own their entity-specific lifecycles.

### TA12-F03 — Hypothesis/Experiment/Observation/Measurement ownership incomplete

**RESOLVED.** C02 owns general engineering Hypothesis/Experiment lifecycle; C07 owns immutable Observation/Measurement/Artifact/Evidence records; provider execution stays local; C10 owns evaluation/failure correlation only.

### TA12-F04 — Go wording could globalize ADR-0003

**RESOLVED.** The design distinguishes:

```text
M0 Go runtime fact
→ Stage A Go Sovereign Host seed hypothesis
→ M1+ module language/placement requires consuming revalidation
```

The topology fixes a sovereign/provider seam, not universal Go.

### TA12-F05 — Contract Model authority implicit

**RESOLVED.** G01 now owns semantic contract identity/version, compatibility, deprecation/removal, canonical meaning, projection/generation authority, conformance criteria and provider-manifest mapping. It is a governed non-deployable owner, not a service.

### EXT-TA12-F01 — Cognitive/provider seam lacked binding-independent lifecycle profile

**RESOLVED.** B01 now fixes provider/build/capability/contract identities; request/attempt/correlation/idempotency; authority/context attenuation; lifecycle/terminal meaning; deadlines; acknowledged cancellation; retries; ambiguous-completion reconciliation; provider restart/snapshot recovery; response/error categories; and TA-04 handoff.

### CODEX-TA12-F01 — AuditRecord and L4 exact history lacked an owner

**RESOLVED.** C12 now owns canonical Audit Records and exact L4 interaction/tool/provider-boundary/material-domain-event history. It has an append path distinct from:

- C06 governed memory and synthesis;
- C07 artifact/receipt/evidence/verdict;
- diagnostic telemetry and provider traces;
- current domain state.

M0 may physically co-store Audit/Evidence, but the logical owners remain distinct because authority, retention, query and promotion semantics differ.

### TA12-F06 — Material Worklog continuity missing

**RESOLVED.** `docs/tracking/WORKLOG.md` is now v0.15.0 and preserves the complete prior history plus the full TA-01/TA-02 proposal, findings, evidence, limitations and exact next action. The temporary branch-only append workflow was removed and is absent from the final PR diff.

## 4. Duplicate-owner audit

| Concept pair | Result |
|---|---|
| G01 contract semantics vs C05 provider compatibility | distinct |
| C01 root identity vs C05/C08/C11 entity identity | distinct |
| C02 Experiment vs C07 Observation/Evidence | distinct |
| C03 Mission Outcome vs C07 Verdict | distinct |
| C04 Effect Request/Decision vs E01 execution vs C07 Receipt | distinct |
| C06 memory/context vs provider thread/index | distinct |
| C07 evidence vs C12 audit/exact history | distinct |
| C08 Presence session vs C01 Person/C04 Authority | distinct |
| C11 inventory vs live device truth | distinct |
| C10 evaluation correlation vs C02 lifecycle/C07 proof/C12 history | distinct |

No representative canonical concept has two commit owners in the fixed design.

## 5. Mutation, audit and provider-boundary audit

The fixed mutation path is:

```text
request/proposal
→ A01 coordination
→ canonical owner validation
→ C04 authority/policy where required
→ owner COMMIT
→ C12 APPEND when audit/history profile requires
→ C07 evidence/receipt record when proof is required
→ notification/projection
```

Provider/model/Harness output remains a proposal, observation, artifact, claim, capability request, provider-local status or controlled gateway receipt. It is never a direct canonical write.

B01 prevents process boundaries from inventing lifecycle semantics during TA-04. Provider success does not commit Aurora state, prove an external effect or close a Mission.

## 6. Topology audit

### Approach A — Core-centric single application

Rejected beyond deterministic M0-like slices because cognition/framework/resource/privilege failure can contaminate the sovereign process.

### Approach B — Early service decomposition

Rejected for Stage A because it adds RPC, service identity, discovery, supervision and distributed consistency before current consumers justify them.

### Approach C — Evolutionary Sovereign Host

Retained because it introduces only the seam already justified by polyglot runtime, provider lifecycle, resource variability, supply-chain cadence, fault containment and provider-local state. Other modules remain logically separate and physically co-located until a named split trigger exists.

The Sovereign Host remains a deployable containing multiple owners, not one domain package. TA-03 receives source/dependency constraints to prevent God-process collapse.

## 7. Stage B audit

Stage B may move the Sovereign Host to a persistent personal node and turn the workstation/mobile into Presences. It adds transport, service identity, Presence authentication, egress, supervision, minimization and recovery mechanisms.

It does not move canonical ownership for G01, identity, Project/Mission, authority, memory/context, provider trust, artifacts/evidence, devices, audit/history or Outcome.

## 8. Validation evidence

```text
initial design:             31623255539 — SUCCESS
v0.2 ownership remediation: 31624034248 — SUCCESS
v0.3 G01/B01 remediation:   31624510891 — SUCCESS
v0.4 C12 remediation:       31624906154 — SUCCESS
Worklog append automation:  31625671777 — SUCCESS
```

The temporary append workflow was restricted to the branch, used `contents: write` only for the append commit and was deleted before final package review. It is not part of the proposed architecture or final diff.

Final branch/PR validation after review, STATUS and PR metadata updates remains the completion evidence for the review package.

## 9. Limitations and non-decisions

- this is architecture/document evidence, not executable proof of a future B01 binding;
- first M1/M2 consumers must prove exact binding, restart and compatibility behavior;
- M1+ canonical module language/placement remains consumer-specific;
- repository strategy, physical stores, authentication, policy/secrets products, supervisor, provider versions and deployment products remain open;
- no Architecture Spike or Aurora runtime implementation occurred;
- authoring/external automated review does not replace operator acceptance.

## 10. Verdict

```text
SEMANTIC VERDICT: PASS FOR OPERATOR REVIEW

blocking material findings open: 0
material findings open: 0
tracking findings open: 0
```

Exact next action:

```text
update STATUS and PR metadata to the final fixed package revision
→ run final branch and PR validation
→ present Approach C + G01/C01–C12 + B01 to operator
→ ACCEPT | REVISE | REJECT
→ STOP before TA-03/TA-04/TA-05/TA-08 finalization,
  Architecture Spike execution or Aurora implementation
```
