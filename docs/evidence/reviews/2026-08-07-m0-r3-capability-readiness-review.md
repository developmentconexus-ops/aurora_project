---
id: REVIEW-AURORA-M0-R3-CAPABILITY-READINESS-2026-08-07
title: M0 ACRM R3 Capability Readiness Review
document_type: adversarial_review
form: explanation
authority: evidence
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - M0 ACRM R3 Capability Readiness review observations and verdict
related:
  - DOC-AURORA-STATUS
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DOC-AURORA-M0-R3-OPERATOR-AUTHORIZATION
  - DOC-AURORA-CAP-SOVEREIGN-CORE-REQUIREMENTS
  - DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC
  - DOC-AURORA-CAP-SOVEREIGN-CORE-THREAT-MODEL
  - DOC-AURORA-CAP-SOVEREIGN-CORE-TEST-PLAN
  - DOC-AURORA-CAP-SOVEREIGN-CORE-R3-COVERAGE
  - REVIEW-AURORA-M0-R3-RESEARCH-FRESHNESS-2026-08-07
source_revision: 9ea8adf5c115f54071d7e36e312695d19420d8b0
r3_package_revision: 4b8558b724f28310fd8fbc6884944f7f59f16ea6
reviewed_at: 2026-08-07
last_reviewed: 2026-08-07
---

# M0 ACRM R3 — Capability Readiness Review

## 1. Executive verdict

```text
R3 PASS
```

Capability:

```text
CAP-SOVEREIGN-CORE
```

Fixed canonical R3 source baseline:

```text
9ea8adf5c115f54071d7e36e312695d19420d8b0
```

Reviewed clean R3 package revision:

```text
4b8558b724f28310fd8fbc6884944f7f59f16ea6
```

R3 passes because the proposed reusable Capability design now defines the complete M0 logical domain, contracts, lifecycle, ownership, authority, security/threat, recovery, observability/evidence, compatibility/migration and verification behavior needed to support later architecture decisions and a scoped implementation contract. All 122 R2 requirements are allocated to current Spec mechanisms and planned tests. No unresolved current-scope product/semantic placeholder remains; the remaining open questions are implementation mechanisms explicitly owned by R4.

This verdict does **not** authorize R4, Architecture Spike execution, Mission Contract, Microdesign or implementation. The R3 specification artifacts remain `status: proposed`; Git/CI and a gate PASS are not operator acceptance of normative content.

---

## 2. Gate definition used

The accepted Aurora Capability Realization Method asks at R3:

> Is the reusable Capability design complete enough to support a scoped implementation contract?

Required Spec areas:

- purpose/use cases;
- goals/non-goals;
- applicability;
- domain model;
- contracts/schemas;
- lifecycle/state;
- architecture/boundaries;
- context/memory;
- authority/effects;
- security/privacy/threat model;
- failure/recovery;
- observability;
- evaluation/evidence;
- rollout/graduation;
- compatibility/migration;
- open questions;
- requirement coverage.

Required test-plan areas:

- requirement-to-test mapping;
- representative journeys;
- adversarial/fault cases;
- evaluation corpus;
- thresholds/rationale;
- evidence format;
- graduation levels.

Gate conditions:

- all R2 requirements allocated to Spec mechanisms and tests;
- no architecture placeholder for current scope;
- open questions explicitly outside the current contract or blocking;
- owner/reviewer defined.

---

## 3. R3 package reviewed

The package consists of:

```text
docs/acceptance/2026-08-07-m0-r3-operator-authorization.md
docs/capabilities/CAP-SOVEREIGN-CORE/SPEC.md
docs/capabilities/CAP-SOVEREIGN-CORE/THREAT-MODEL.md
docs/capabilities/CAP-SOVEREIGN-CORE/TEST-PLAN.md
docs/capabilities/CAP-SOVEREIGN-CORE/R3-COVERAGE.md
docs/reviews/2026-08-07-m0-r3-research-freshness-review.md
```

Existing R1/R2 owners remain inputs:

```text
APPLICABILITY.md
REQUIREMENTS.md
R2-COVERAGE.md
```

`SPEC.md` owns reusable behavior. The threat model, test plan and coverage matrix are specialized subordinate artifacts; they do not redefine the Spec.

---

## 4. Architecture/design result

### One Capability, not five premature platforms

R3 considered:

1. one monolithic Spec;
2. one Capability with modular Spec/threat/test/coverage artifacts;
3. separate identity/state/authority/recovery/evidence capabilities.

Option 2 was selected because it preserves one `CAP-SOVEREIGN-CORE` boundary while giving enough review isolation. Option 3 would generalize platform boundaries before independent consumers exist.

### Logical modules

The Spec defines one Sovereign Core with exactly one logical owner for each M0 responsibility:

```text
Identity Module
Project State Module
Authority Module
Portability / Recovery Module
Audit / Evidence Module
```

A coordinator composes operations but does not own duplicate durable state.

Replaceable ports remain mechanism-neutral:

```text
Operator Adapter
Durable State Port
Evidence/Artifact Port
Time Source Port
Integrity Port
```

### State model

The logical model uses immutable accepted `ProjectStateRevision` records plus one current governing revision. This follows R2's explicit revision/precondition/supersession requirements and does not select event sourcing, snapshot persistence, database transactions or another physical persistence strategy.

`accepted_state` is no longer an opaque design placeholder. R3 defines its minimum semantic envelope:

```text
state_schema_version
state_kind
state_summary
optional state_payload
```

`state_payload` remains project data and cannot redefine identity, authority, policy, ownership or transition protocol.

### Authority model

R3 defines:

- owner/operator identity distinct from technical access;
- M0 authority grant/state/snapshot semantics;
- scope/expiry/revocation/supersession;
- `NextSafeActionProjection` derived from current state + current authority + evaluation time;
- fail-closed missing/corrupt/invalid authority;
- no external effect plane.

A cached next-action projection is non-authoritative when state revision, authority revision or time-validity changes.

---

## 5. Restore versus restart safety

The review treated ordinary restart and backup restore as different trust situations.

### Ordinary restart

A fresh process loading the current canonical store may recover a previously issued authority as effective only after current status, scope, conditions and time-validity are re-evaluated.

### Restore

A backup/export may predate a later revocation. Therefore an apparently active authority in restored data is not automatically trusted.

R3 now requires:

```text
apparently active restored authority
→ freshness cannot be proven
→ REVALIDATION_REQUIRED
→ no authority-dependent next safe action
→ restored grant cannot authorize its own revalidation
→ non-owner revalidation denied
→ authenticated Leandro owner identity may explicitly revalidate
→ new attributable authority-state revision
→ VALID only if current scope/time/conditions pass
```

This closes the recovery path without allowing a suspect restored grant to bootstrap itself back into authority.

---

## 6. Threat-model review

The R3 threat model identifies protected assets, classifications, six trust boundaries and twenty M0 threat classes.

The minimum R2 threats are covered:

- canonical-state tampering/corruption;
- stale/rolled-back authority;
- unsafe restore;
- export/backup exposure;
- identity collision;
- untrusted-content authority injection.

R3 additionally makes these implementation-relevant risks explicit for R4:

- clock rollback extending expired authority;
- crash/partial commit ambiguity;
- storage rollback;
- migration semantic drift;
- evidence/audit spoofing;
- sensitive telemetry leakage;
- unsafe retry;
- full-host compromise as residual risk.

Security semantics are fixed without selecting authentication, cryptographic, storage, policy-engine or telemetry products.

---

## 7. Data classification review

R3 assigns accepted Aurora data classes to every material M0 data family:

| Family | R3 class |
|---|---|
| Aurora identity metadata | `INTERNAL` |
| operator identity/authentication provenance | `SENSITIVE` |
| Project operational state | `CONFIDENTIAL` |
| authority grants/state/snapshots | `SENSITIVE` |
| audit records | `CONFIDENTIAL` |
| evidence metadata/references | `CONFIDENTIAL` minimum/inherit higher |
| telemetry/correlation identifiers | `INTERNAL` |
| export/backup package | `SENSITIVE` minimum/inherit higher |
| integrity descriptor | `INTERNAL` minimum/inherit as mechanism requires |

No secret material is required as canonical M0 state.

---

## 8. Failure/recovery review

R3 defines explicit failure classes rather than generic retry:

```text
INVALID_TRANSITION
AUTHORITY_INVALID
STATE_UNAVAILABLE
STATE_CORRUPT
VERSION_INCOMPATIBLE
RESTORE_UNSAFE
MIGRATION_FAILED
OPERATION_AMBIGUOUS
INTERNAL_OPERATIONAL
```

The containment invariant is consistent across them:

> unverifiable state or authority is never presented as current/permitted merely to preserve progress.

Blind retry after ambiguous state mutation, restore or migration is prohibited.

---

## 9. Test-plan review

The proposed test plan defines:

- 15 deterministic logical fixtures;
- 84 explicit test IDs;
- representative Golden Proof, authority/restore, corruption and ambiguous-mutation journeys;
- adversarial/fault cases for every critical M0 safety area;
- future evidence format;
- graduation levels;
- deterministic success thresholds.

Important negative cases include:

- stale state revision;
- malformed envelope;
- identity mutation attempt;
- wrong authority scope;
- expired authority;
- revoked authority;
- time rollback;
- pre-revocation backup restore;
- non-owner restore revalidation;
- identity collision;
- corrupt state/export;
- incompatible version;
- migration semantic drift;
- ambiguous crash;
- untrusted-content authority injection;
- unsafe retry.

R3 deliberately does not select a test framework or fault-injection tool.

---

## 10. Requirement allocation result

`R3-COVERAGE.md` contains exactly one allocation row for every R2 requirement:

```text
R2 requirements:          122
R3 coverage rows:         122
Unallocated requirements:   0
```

Every row names:

1. one or more current Spec/threat mechanisms/sections; and
2. one or more planned test/document-review cases.

No requirement is allocated to a concrete file/package/database/runtime implementation.

---

## 11. Mechanical verification

A dedicated one-shot R3 validator executed in GitHub Actions run:

```text
31150888970
```

Verified output:

```text
M0 R3 package validation PASS
R2 requirements: 122
R3 coverage rows: 122
planned test IDs defined: 84
test IDs referenced by coverage: 80
undefined referenced tests: 0
required ACRM Spec sections: present
required logical owners: present
required semantic closure: present
required threat classes: present
critical negative test cases: present
candidate technology selections: 0
research freshness disposition: sufficient for R3; R4 revalidation required
```

The same run also completed:

- generated documentation projection generation;
- repository documentation validation;
- generated-projection freshness check;
- helper cleanup with `git diff --cached --check`.

The repository documentation validator returned PASS and the one-shot validator removed itself.

---

## 12. Adversarial findings found and corrected before verdict

### R3-F01 — restore fail-closed state had no explicit safe recovery root

Initial Spec correctly blocked stale restored authority with `REVALIDATION_REQUIRED`, but it did not explicitly establish who could revalidate without relying on the authority record being questioned.

Risk:

- deadlock: safe restore could permanently block owner recovery; or
- bypass: implementation might let the restored grant authorize its own reactivation.

Resolution:

- authenticated `OperatorIdentityRef` for Leandro is the narrow owner bootstrap/recovery root;
- it is not an external-effect credential;
- restored authority cannot authorize its own revalidation;
- non-owner revalidation is denied;
- owner revalidation creates a new attributable authority-state revision;
- `T-PORT-013` and Journey J-002 verify the full blocked→owner-revalidated path.

Status:

```text
RESOLVED
```

### R3-F02 — `accepted_state` remained too generic for the current scope

Initial Spec named an `accepted_state` envelope but did not close its minimum semantic content. That left R4 at risk of deciding product-state semantics rather than only representation/mechanism.

Resolution:

R3 now fixes:

```text
state_schema_version
state_kind
state_summary
optional state_payload
```

with an explicit rule that `state_payload` is opaque project data and cannot redefine Core identity, authority, policy, ownership or transition semantics.

Status:

```text
RESOLVED
```

No further material R3 finding remains open.

---

## 13. False-inclusion / premature-generalization review

R3 does **not** build or specify future product systems that M0 does not require:

- M1 memory/Context Builder;
- M2 Registry/AHDK/provider trust;
- Mission/Delegation orchestration;
- Effect Gateway/PDP/Credential Broker;
- multi-Presence;
- physical devices;
- campaigns/budgets/timers;
- self-improvement;
- generic workflow engine;
- generic event platform;
- generic authorization platform.

The five internal logical modules are responsibilities inside one Capability, not independently productized services/capabilities.

---

## 14. Research freshness disposition

The Research Map's R3/R4 freshness rule was addressed in a dedicated review.

Three directly relevant reports were re-read:

- Durable Execution;
- Authority, Identity and Effects;
- Events, Observability and Schemas.

All are `current`, last reviewed 2026-08-05, and R3 uses only boundary findings already compatible with accepted constitutional/R2 semantics. R3 does not rely on current implementation-version claims for any candidate technology.

Disposition:

```text
R3_RESEARCH_FRESHNESS: SUFFICIENT
R4_MECHANISM/VERSION_REVALIDATION: REQUIRED
```

R4 must refresh primary sources for each actual technical decision before accepting a mechanism.

---

## 15. Open R4 decisions — not R3 placeholders

The Spec enumerates fifteen R4 questions:

```text
Core language/runtime
operational-state store
state-vs-event persistence mechanism
schema/serialization representation
crash-consistent commit mechanism
integrity mechanism
time/rollback semantics
local owner authentication/bootstrap
export/backup format/topology
migration mechanism/tooling
audit physical mechanism
telemetry backend/transport
process/deployment topology
durable-engine applicability
restore freshness/revalidation implementation
```

These do not block R3 because the product behavior, safety invariants, boundary contracts and tests they must satisfy are already explicit. They are mechanism choices, exactly the subject of R4.

No candidate technology or winner was selected.

---

## 16. ADR applicability

### ADR-0001

R3 obeys Aurora-owned, implementation-neutral domain contracts and replaceable ports. No protocol binding is selected.

### ADR-0002

AHDK/conformance remains later M2/provider work and is not a dependency of CAP-SOVEREIGN-CORE M0 behavior. R3 does not authorize AHDK implementation.

---

## 17. Gate checklist

| R3 condition | Result |
|---|---|
| purpose/use cases/goals/non-goals complete | PASS |
| applicability baseline explicit | PASS |
| domain model complete for M0 | PASS |
| semantic contracts/schemas complete for current scope | PASS |
| lifecycle/state model complete | PASS |
| logical owners/boundaries explicit | PASS |
| context/memory boundary explicit | PASS |
| authority/effects boundary explicit | PASS |
| threat model/data classification complete | PASS |
| failure/recovery model complete | PASS |
| observability/audit/evidence semantics complete | PASS |
| compatibility/migration semantics complete | PASS |
| rollout/graduation defined | PASS |
| Capability Test Plan complete | PASS |
| 122/122 requirements allocated to Spec mechanisms | PASS |
| 122/122 requirements allocated to planned verification | PASS |
| critical negative paths represented | PASS |
| no current-scope architecture placeholder | PASS |
| open questions assigned to R4 | PASS |
| owner/reviewer defined | PASS |
| research freshness disposition explicit | PASS |
| technology/stack selection absent | PASS |
| Architecture Spike execution absent | PASS |
| R4/R5/R6/R7 work not performed | PASS |

---

## 18. Verdict and stop boundary

```text
R3 PASS
```

Exact next permitted sequence:

```text
record R3 PASS
→ STOP at R3 boundary
→ await explicit operator authorization for M0 ACRM R4 — Architecture/Decision Readiness
```

R3 PASS does not authorize:

- R4 analysis by implication;
- Architecture Spike execution;
- stack selection;
- Mission Contract;
- Microdesign;
- Aurora Core implementation;
- AHDK implementation;
- MNFS integration.
