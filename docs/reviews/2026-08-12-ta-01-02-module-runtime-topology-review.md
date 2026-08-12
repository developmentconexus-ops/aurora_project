---
id: REVIEW-AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY-2026-08-12
title: TA-01/TA-02 Module and Runtime Topology Adversarial Review
document_type: architecture_review
form: reference
authority: evidence
status: current
version: 0.4.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - adversarial review of the proposed TA-01 and TA-02 module/runtime topology
  - admitted TA-01 and TA-02 findings and final semantic review verdict
related:
  - DESIGN-AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY
  - DESIGN-AURORA-TECHNICAL-ARCHITECTURE-BASELINE-MAP
  - PLAN-AURORA-TECHNICAL-ARCHITECTURE-BASELINE
  - DOC-AURORA-STATUS
  - DOC-AURORA-WORKLOG
initial_review_target: 2192f182c24eed0c8406062957e035a0db6e88a5
fixed_design_target: 965211ad421f13994285031bbcf04b7e943cf75e
continuity_target: 3f720eb1785579ba3cc7c5df3c76e1cfcff8f3c9
last_reviewed: 2026-08-12
---

# TA-01/TA-02 Module and Runtime Topology — Adversarial Review

## 1. Review boundary

```text
initial target:
2192f182c24eed0c8406062957e035a0db6e88a5

fixed semantic design target:
965211ad421f13994285031bbcf04b7e943cf75e

continuity target after Worklog append and temporary-workflow removal:
3f720eb1785579ba3cc7c5df3c76e1cfcff8f3c9
```

The review tested:

1. complete and non-overlapping canonical ownership;
2. interpreted Intent ownership before Mission creation;
3. provider/framework capture;
4. accidental globalization of M0 decisions;
5. God-process and service-sprawl risks;
6. Stage A operational proportionality;
7. Stage B identity and ownership continuity;
8. binding-independent provider lifecycle and reconciliation;
9. ownership of runtime startup, shutdown, restart and supervision policy;
10. separation of audit/exact history, memory, evidence and telemetry;
11. Contract semantic authority versus Git/document custody;
12. handoff to later technical tranches without hidden product selection.

## 2. Fixed proposal reviewed

The design proposes:

```text
G01 Contract Model Governance

C01 Identity and Relationship
C02 Project, World and Experiment State
C03 Intent, Mission and Delegation Control
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

It separately defines:

```text
A01 Core Application Coordination
A02 Cognitive Coordination
A03 Capability Fabric / Harness Integration
A04 Durable Execution Port
A05 Runtime Lifecycle Coordination

B01 Transport-neutral Provider Runtime Boundary Profile

E01 Effect Gateway family
E02 Credential Broker boundary

P01 Cognitive Runtime Provider
P02 Specialized Harness Provider
P03 Model / External Service Adapter
P04 Device / Laboratory Provider
```

The recommendation is:

```text
Approach C — Evolutionary Sovereign Host

one small persistent Aurora Sovereign Host
+
one separate on-demand Cognitive Runtime Provider at first consumer
+
independent specialized Harness providers
+
all remaining physical splits require a named current trigger
```

## 3. Finding disposition

### TA12-F01 — Environment/Device owner missing

**RESOLVED.** C11 owns Device/Environment identity, inventory and registration relationships; live telemetry, actuation and target truth remain source-owned.

### TA12-F02 — Root identity overlapped Provider/Presence/Device identity

**RESOLVED.** C01 owns root Aurora/Person/relationship identity and actor-reference vocabulary. C05, C08 and C11 own entity-specific identity lifecycles.

### TA12-F03 — Hypothesis/Experiment/Observation/Measurement allocation incomplete

**RESOLVED.** C02 owns general engineering Hypothesis/Experiment lifecycle. C07 owns immutable Observation/Measurement/Artifact/Evidence records. Provider execution remains local; C10 owns evaluation/failure correlation.

### TA12-F04 — M0 Go could become a silent universal decision

**RESOLVED.** The design distinguishes:

```text
accepted M0 Go runtime fact
→ Stage A Go Sovereign Host seed hypothesis
→ M1+ module language/placement requires consuming revalidation
```

### TA12-F05 — Contract Model authority implicit

**RESOLVED.** G01 owns semantic families/versions, canonical meaning, compatibility, deprecation/removal, projection/generation authority and conformance criteria. C05 owns provider compatibility/trust/approval; Git owns artifact custody and history only.

### EXT-TA12-F01 — Provider process seam lacked lifecycle/reconciliation semantics

**RESOLVED.** B01 defines provider/build/instance/incarnation/capability/contract identities, request/attempt/idempotency, authority attenuation, lifecycle, deadlines, acknowledged cancellation, retry, snapshot/restart reconciliation and required response/error categories without selecting transport.

### CODEX-TA12-F01 — AuditRecord and L4 exact history lacked an owner

**RESOLVED.** C12 owns canonical Audit Records and attributable exact interaction/tool/provider-boundary/material-domain-event history, separate from C06 memory, C07 evidence and telemetry.

### CODEX-TA12-F02 — Interpreted Intent lacked a canonical owner

**RESOLVED.** C03 now owns validated interpreted Intent and its lifecycle before explicit promotion to Mission. C08/C12 retain raw interaction/history; A02/P01 may propose candidates but cannot commit Intent or create Mission automatically.

### CODEX-TA12-F03 — Multi-process lifecycle ownership missing

**RESOLVED.** A05 owns desired runtime state and start, attach, readiness, drain, stop, restart and reconciliation decisions. TA-08 may select Windows Service, systemd or another supervisor adapter, but cannot take lifecycle-policy ownership from A05.

### EXT-TA12-F02 — Effect receipt could appear provider-produced

**RESOLVED.** Providers may carry only references to E01-controlled receipts. E01 produces the receipt; C07 records its canonical evidence identity; C12 appends required history.

### EXT-TA12-F03 — Stage B provider identities were over-preserved

**RESOLVED.** Stable Aurora domain identities remain. A migrated/reinstalled exact provider environment receives a new `provider_instance_id`, fresh C05 approval snapshot and new `runtime_incarnation_id`; provider-local state is reconciled rather than canonized.

### EXT-TA12-F04 — Contract semantics versus Git custody ambiguity

**RESOLVED.** The ownership matrix separates G01 semantic governance from Git/document artifact custody and commit history.

### TA12-F06 — Worklog/material continuity missing

**RESOLVED.** `WORKLOG.md` v0.15.0 recorded the full proposal and findings; v0.16.0 records the final Intent/A05/provider-migration remediation. Temporary workflows used only to preserve append-only history were deleted and are absent from the final intended diff.

### Minor clarity/lint findings

**RESOLVED:** `near-term` wording, heading hierarchy, UI wording and topology-diagram alignment were corrected.

## 4. Canonical owner audit

| Potential overlap | Result |
|---|---|
| G01 semantics vs Git custody | distinct |
| C01 root identity vs C05/C08/C11 entity identity | distinct |
| C02 Experiment vs C07 Observation/Evidence | distinct |
| C03 Intent/Mission vs C08 interaction vs A02 inference | distinct |
| C03 Outcome vs C07 Verdict | distinct |
| C04 Effect Request/Decision vs E01 execution vs C07 Receipt | distinct |
| C06 memory/context vs provider thread/index | distinct |
| C07 evidence vs C12 audit/exact history | distinct |
| C10 evaluation correlation vs C02 lifecycle/C07 proof/C12 history | distinct |
| A05 lifecycle policy vs TA-08 supervisor mechanism | distinct |

No representative concept has two canonical commit owners in the fixed design.

## 5. Mutation, audit and provider-boundary audit

The fixed path is:

```text
request/proposal
→ A01 coordination
→ canonical owner validation
→ C04 authority/policy when required
→ owner COMMIT
→ C12 APPEND when audit/history profile requires
→ C07 evidence/receipt record when proof is required
→ notification/projection
```

Provider/model/Harness output is a candidate, proposal, observation, artifact, claim, capability request, provider-local status or reference to an E01-controlled receipt. It is never a direct canonical write.

B01 prevents TA-04 from inventing lifecycle semantics around a chosen transport. Provider `SUCCEEDED` does not commit Aurora state, prove an external effect or close a Mission.

A05 prevents TA-08 or an operating-system supervisor from becoming the implicit owner of provider lifecycle policy.

## 6. Topology audit

### Approach A — Core-centric single application

Rejected beyond deterministic M0-like slices because cognition, framework dependencies, resources and privileges would contaminate the sovereign process.

### Approach B — Early service decomposition

Rejected for Stage A because it adds RPC, workload identity, discovery, supervision and distributed consistency before current consumers justify them.

### Approach C — Evolutionary Sovereign Host

Retained because it introduces only the seam already justified by polyglot runtime, provider lifecycle, resource variability, supply-chain cadence, fault containment and provider-local state. Other components remain logically separate and physically co-located until a named split trigger exists.

The Sovereign Host is a deployable containing multiple owners, not a God domain package. TA-03 receives source-dependency requirements to preserve that distinction.

## 7. Stage B audit

Stage B may move the Sovereign Host to a persistent personal node and turn workstation/mobile into Presences. It adds transport, service identity, Presence authentication, egress, supervision, minimization and recovery mechanisms.

It does not move G01/C01–C12 ownership. Environment-bound provider instances and runtime incarnations are recreated and re-approved; provider-local state is reconciled under B01.

## 8. Validation evidence

```text
initial design:                 31623255539 — SUCCESS
v0.2 ownership remediation:     31624034248 — SUCCESS
v0.3 G01/B01 remediation:       31624510891 — SUCCESS
v0.4 C12 remediation:           31624906154 — SUCCESS
Worklog append automation:      31625671777 — SUCCESS
v0.5 Intent/A05 remediation:    31626534727 — SUCCESS
final continuity automation:    31626750628 — SUCCESS
```

The temporary append/remediation workflows were branch-restricted, used write authority only to make exact documentary commits and were removed before the final package. They are not architecture decisions or intended PR contents.

Final normal Documentation workflow runs after review/STATUS/PR closeout remain the completion evidence for the package.

## 9. Limitations and non-decisions

- this is architecture/document evidence, not executable proof of a future B01 binding or A05 supervisor adapter;
- first M1/M2 consumers must prove exact binding, restart and compatibility behavior;
- M1+ canonical module language/placement remains consumer-specific;
- repository strategy, physical stores, authentication, policy/secrets products, supervisor implementation, provider versions and deployment products remain open;
- no Architecture Spike or Aurora runtime implementation occurred;
- automated/adversarial review does not replace operator acceptance.

## 10. Verdict

```text
SEMANTIC VERDICT: PASS FOR OPERATOR REVIEW

blocking material findings open: 0
material findings open: 0
tracking findings open: 0
```

Exact next action:

```text
update STATUS and PR metadata to final package revision
→ run final normal branch and PR validation
→ resolve all review threads with evidence
→ present Approach C + G01/C01–C12 + A05 + B01 to operator
→ ACCEPT | REVISE | REJECT
→ STOP before merge, later-tranche finalization,
  Architecture Spike execution or Aurora implementation
```
