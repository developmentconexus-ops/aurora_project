---
id: REVIEW-AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY-2026-08-12
title: TA-01/TA-02 Module and Runtime Topology Adversarial Review
document_type: architecture_review
form: reference
authority: evidence
status: current
version: 0.1.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - adversarial review of the proposed TA-01 and TA-02 module/runtime topology
  - admitted TA-01 and TA-02 findings and current review verdict
related:
  - DESIGN-AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY
  - DESIGN-AURORA-TECHNICAL-ARCHITECTURE-BASELINE-MAP
  - PLAN-AURORA-TECHNICAL-ARCHITECTURE-BASELINE
  - DOC-AURORA-STATUS
review_target: 2192f182c24eed0c8406062957e035a0db6e88a5
last_reviewed: 2026-08-12
---

# TA-01/TA-02 Module and Runtime Topology — Adversarial Review

## 1. Review boundary

Target:

```text
2192f182c24eed0c8406062957e035a0db6e88a5
```

Review questions:

1. Does every representative global concept have one canonical owner?
2. Are any owners duplicated or overly broad?
3. Does the topology globalize an M0-scoped mechanism silently?
4. Can a provider/framework become sovereign through convenience?
5. Does the proposed Sovereign Host become a God process?
6. Does Stage A remain operationally small?
7. Does Stage B preserve identities and ownership?
8. Are process boundaries based on real runtime/security/fault evidence?
9. Are deferred mechanisms and later consumers explicit?
10. Is TA-03 blocked from inventing ownership?

## 2. Initial strengths

The target correctly:

- separates domain owners, application coordinators, deployables, enforcement boundaries, provider runtimes and adapters;
- prohibits direct provider/Harness/model/Presence writes to canonical state;
- distinguishes Project, Mission, Authority, Memory, Registry and Evidence ownership;
- keeps Effect Gateway and Credential Broker outside probabilistic model control;
- compares three complete topologies rather than isolated service choices;
- rejects early service decomposition;
- proposes one small sovereign deployable plus an independently replaceable cognitive/provider seam;
- preserves Stage A to Stage B identity continuity;
- avoids choosing repository, protocol, database, authentication and deployment products;
- provides explicit process split triggers and TA-03 inputs.

The Approach C direction remains credible, but the owner catalog is not yet complete enough for operator acceptance.

## 3. Admitted findings

### TA12-F01 — Missing canonical owner for Device and Environment inventory

**Severity:** MATERIAL / BLOCKING

The design lets Project reference devices, Presence own Presence context and a future provider control devices, but it does not name the canonical Aurora owner for:

- Device identity/inventory;
- Environment identity/topology;
- calibration/firmware/controller references;
- device registration/revocation relationships.

Live telemetry and target-system truth must remain external/source-owned, but the Aurora inventory/relationship owner must be explicit.

**Required remediation:** add a future/deferred `Environment and Device Registry` domain owner. It owns identity/inventory/metadata relationships, not live telemetry or actuation.

---

### TA12-F02 — Identity owner wording overlaps Provider and Presence identity lifecycle

**Severity:** MATERIAL / BLOCKING

C01 currently says it owns broad “stable actor identity references and identity lifecycle semantics,” while C05 owns Provider/Harness identity and C08 owns Presence identity/status. That wording can create duplicate ownership.

**Required remediation:** narrow C01 to root Aurora/Person/relationship identity plus cross-domain actor-reference vocabulary. Entity-specific identity lifecycle stays with the owning module:

```text
Provider/Harness/ProviderInstance → C05
Presence/InteractionSession       → C08
Device/Environment                → new Device/Environment owner
Mission/Delegation actor linkage  → C03 references the actor chain
```

Authentication proof mechanisms remain TA-06 adapters/decisions.

---

### TA12-F03 — Hypothesis, Experiment, Observation and Measurement ownership incomplete

**Severity:** MATERIAL / BLOCKING

The accepted world model includes Hypothesis, Experiment, Experiment Run, Observation and Measurement. The proposed matrix does not assign them completely.

**Required remediation:** establish:

- Project-owned Hypothesis/Experiment identity, lifecycle and relationship in C02 for general engineering work;
- provider-local execution steps/runs remain provider-local, with global references in C03/C02 as needed;
- immutable Observation/Measurement/Artifact/Evidence records and provenance in C07;
- evaluation-specific correlation and improvement lifecycle in C10.

This keeps raw provider execution separate from globally meaningful experimental state.

---

### TA12-F04 — Go wording can be read as silently extending ADR-0003 beyond M0

**Severity:** MATERIAL / BLOCKING

The topology reasonably uses the existing Go M0 Core as the Stage A sovereign seed, but the design says the Sovereign Host runtime is Go while also claiming not to globalize M0 decisions. Without clarification, acceptance could be misread as deciding every future Core module will use Go.

**Required remediation:** separate:

```text
current deployable fact/hypothesis
→ the M0 sovereign executable is Go and can seed the Stage A Host

cross-horizon language decision
→ placing new M1+ canonical modules in that Go process requires
   explicit consuming architecture/ADR revalidation

fixed topology decision
→ cognitive/provider execution remains behind a separate
   replaceable process/contract boundary when first consumed
```

Add a decision-register row for cross-horizon Go scope.

---

### TA12-F05 — Canonical Contract Model is visible but lacks an explicit non-runtime owner entry

**Severity:** MATERIAL / NON-BLOCKING

ADR-0001 owns the semantic policy, but the component model should explicitly classify the Contract Model as a governed architecture/specification asset rather than a runtime service.

**Required remediation:** add a non-runtime architecture asset entry:

```text
Canonical Contract Model
→ owned by Aurora Specs/accepted decisions
→ projected into schemas/types/bindings
→ not hosted/owned by Cognitive Runtime, AHDK or Capability Fabric
```

---

### TA12-F06 — Worklog and fixed review continuity not yet complete

**Severity:** TRACKING / BLOCKING FOR REVIEW CLOSEOUT

The proposal and STATUS exist, but `WORKLOG.md` has not yet been appended and no fixed remediated review target exists.

**Required remediation:** after semantic fixes:

- append the material work entry;
- update STATUS and PR metadata;
- run validation;
- re-review one fixed revision;
- record all findings as resolved or open.

## 4. Non-findings / rejected concerns

### One process per domain owner is not required

The design correctly distinguishes logical ownership from physical services. No change required.

### Separate Cognitive Runtime is not premature service fashion

The proposed seam has existing drivers: language/runtime diversity, framework cadence, resource variability, provider replaceability, fault containment and accepted provider-local ownership. Exact transport remains deferred. No runtime Spike is required merely to accept the conceptual seam; the first consuming capability still needs contract/recovery proof.

### Presence may remain co-packaged in Stage A

The accepted Stage A design permits logical separation with combined packaging. A second Presence, Stage B or OS/device lifecycle need is the correct split trigger.

### Artifact/Evidence versus Mission Outcome is coherent

C07 owns artifact/claim/receipt/evidence/verdict records. C03 owns final Mission/Delegation Outcome after consuming authorized verdicts. No duplicate owner found.

### Effect Request and execution are coherently separated

C04 owns Authority/Policy and canonical Effect Request/Policy Decision state. E01 executes and emits a receipt; C07 records the canonical receipt. C03 carries purpose/budget/delegation relationships. No change required beyond preserving these links.

## 5. Initial verdict

```text
VERDICT: BLOCKED PENDING REMEDIATION

blocking material findings: 4
material non-blocking findings: 1
tracking closeout findings: 1
```

The proposed Approach C remains the leading alternative, but the design must not be presented as acceptance-ready until TA12-F01 through TA12-F06 are resolved and a fixed revision passes validation/review.
