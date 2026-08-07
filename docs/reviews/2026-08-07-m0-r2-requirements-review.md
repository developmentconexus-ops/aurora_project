---
id: REVIEW-AURORA-M0-R2-REQUIREMENTS-2026-08-07
title: M0 ACRM R2 Requirements Review
document_type: adversarial_review
form: explanation
authority: evidence
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - M0 ACRM R2 requirements review observations and verdict
related:
  - DOC-AURORA-STATUS
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DOC-AURORA-CAP-SOVEREIGN-CORE-APPLICABILITY
  - DOC-AURORA-CAP-SOVEREIGN-CORE-REQUIREMENTS
  - DOC-AURORA-CAP-SOVEREIGN-CORE-R2-COVERAGE
  - DOC-AURORA-M0-R2-OPERATOR-AUTHORIZATION
  - REVIEW-AURORA-M0-R1-APPLICABILITY-2026-08-07
source_revision: 495b712142d7c3d722da2298f7a0b060707f9f5e
requirements_package_revision: a8ffbbe22995b8e683d9d49ad06f487c745709f9
reviewed_at: 2026-08-07
last_reviewed: 2026-08-07
---

# M0 ACRM R2 — Requirements Review

## 1. Executive verdict

```text
R2 PASS
```

Capability:

```text
CAP-SOVEREIGN-CORE
```

Fixed canonical R2 source baseline:

```text
495b712142d7c3d722da2298f7a0b060707f9f5e
```

Reviewed clean requirements-package revision:

```text
a8ffbbe22995b8e683d9d49ad06f487c745709f9
```

R2 passes because all 127 R1-active constitutional source rows are represented by traceable, normative and verifiable requirements; high-risk statements have verification direction; the package does not reactivate deferred roadmap scope; discovered semantic weakenings/overreach were corrected before verdict; no unresolved contradiction remains; and no implementation technology or Architecture Spike winner was selected.

This verdict does **not** authorize R3 or later work.

The requirements document remains `status: proposed`. `R2 PASS` means the derivation package satisfies the R2 readiness gate; it is not a claim that Git/CI promoted normative content to operator acceptance.

## 2. Gate used

The accepted Capability Realization Method defines R2 by the question:

> Are all applicable product statements transformed into verifiable requirements?

Required work:

- derive requirements;
- assign stable IDs;
- link sources/rationale;
- classify risk;
- define verification method;
- identify open decisions/spikes;
- detect duplicate/conflicting requirements;
- create a coverage matrix.

Gate conditions:

- no applicable source without requirement or rationale;
- no vague/high-risk requirement without verification direction;
- contradictions resolved or block.

R3 owns reusable Capability/System design, lifecycle/state model, threat model and test-plan allocation. R4 owns material technical choices/spikes. Neither was executed here.

## 3. Package reviewed

R2 produced:

```text
docs/acceptance/2026-08-07-m0-r2-operator-authorization.md
docs/capabilities/CAP-SOVEREIGN-CORE/REQUIREMENTS.md
docs/capabilities/CAP-SOVEREIGN-CORE/R2-COVERAGE.md
```

The package derives:

```text
122 atomic Capability requirements
```

from:

```text
127 R1-active constitutional source rows
= 78 APPLIES
+ 49 PARTIALLY_APPLIES
```

The number of derived requirements intentionally does not equal the number of source rows. Several constitutional rows express the same M0 invariant and were consolidated; several broad rows were decomposed into separately verifiable obligations.

## 4. Mechanical verification

A one-shot R2-specific validator executed against the reviewed package in GitHub Actions run:

```text
31149256807
```

Result:

```text
M0 R2 package validation PASS
```

Verified output:

```text
derived requirements: 122
risk distribution: critical=51, high=53, medium=18
R1 active sources: 127
coverage rows: 127
source->requirement links: 409
uncovered active sources: 0
inactive source references: 0
unsupported verification methods: 0
candidate technology selections: 0
```

The same run also completed:

- generated documentation projection generation;
- repository documentation validation;
- generated-projection freshness check;
- `git diff --cached --check` during helper cleanup.

The general documentation validator returned `PASS`; the one-shot validator/workflow removed itself after success.

## 5. Atomicity and verifiability review

The 122 requirements were reviewed for being atomic enough to assign concrete evidence later without freezing implementation design.

Accepted requirement shapes include:

- one stable-identity invariant even when it names several forbidden identity owners;
- one schema/semantic completeness invariant where a logically single record needs multiple required fields;
- one end-to-end Golden Proof requirement (`CAP-SOVEREIGN-CORE-REQ-031`) because it intentionally represents the Product Milestone journey rather than an implementation unit.

No requirement requires an implementation file/package/component allocation. `CAP-SOVEREIGN-CORE-REQ-122` explicitly preserves that later-gate boundary.

Every requirement uses a normative modal and one or more accepted ACRM verification-method classes. No `critical` or `high` requirement lacks verification direction.

## 6. Coverage review

The source-to-requirement matrix was checked as a set rather than relying on its printed summary.

Result:

```text
expected active sources: 127
covered active sources:  127
missing:                   0
inactive sources cited:    0
```

All 49 `PARTIALLY_APPLIES` sources remain bounded to the M0 slice defined in R1. R2 does not use their deferred remainder as current commitment.

No `DEFERRED_BY_ROADMAP` or `NOT_APPLICABLE` R1 row was silently reactivated as a source for a derived requirement.

## 7. High-risk dependency review

All seven high-risk dependencies carried from R1 have concrete R2 obligations.

### D1 — Authority snapshot correctness and restore safety

Covered principally by `REQ-032` through `REQ-045`, including scope, expiry, revocation, provenance, next-safe-action derivation, fail-closed behavior and restore safety.

### D2 — Operational-state ownership

Covered principally by `REQ-010` through `REQ-020`: canonical owners, separation from transcript/memory/Harness/engine/Git/logs, governing-state precedence and projection boundaries.

### D3 — Stable identity and migration

Covered by `REQ-001` through `REQ-003` plus export/version/migration obligations `REQ-057`, `REQ-064` and `REQ-065`.

### D4 — Backup/restore integrity and portability

Covered by `REQ-056` through `REQ-066`, including completeness, compatibility, integrity, authority safety, collision behavior, sovereignty and migration evidence.

### D5 — Event/audit versus telemetry

Covered by `REQ-077` through `REQ-088`. The final formulation explicitly prohibits event/audit/log/telemetry from becoming the sole canonical M0 current-state/authority source.

### D6 — Security of the first canonical durable store

Covered by `REQ-067` through `REQ-076`, with `REQ-075` making the R3 threat model a hard readiness obligation rather than performing that threat model during R2.

### D7 — Simple proof surface is not Presence Fabric

Covered by `REQ-008`, `REQ-030` and `REQ-072`: the interface is a proof/control surface and cannot own identity/state or force a Presence/UI architecture.

## 8. Adversarial findings found and corrected before verdict

The first R2 draft did **not** receive an automatic PASS. Eight wording defects were corrected while retaining requirement IDs and coverage.

### R2-F01 — event/audit authority exception weakened constitution

Initial `REQ-081` allowed a future architecture decision to make events/logs/telemetry the sole canonical state source. That weakened the active constitutional invariant that events are not the sole source of critical state.

Resolution:

```text
Domain Events, audit records, logs and telemetry MUST NOT be the sole
canonical source of M0 current state or authority.
```

Status: `RESOLVED`.

### R2-F02 — sensitive-observability exception weakened constitution

Initial `REQ-076` permitted sensitive payloads when a proof need/policy justified them, weakening the active observability requirement to propagate stable identifiers without sensitive payloads.

Resolution: correlation/general telemetry remains payload-minimized; sensitive evidence, when legitimately required, is governed as separate Evidence/Artifact data rather than correlation payload.

Status: `RESOLVED`.

### R2-F03 — archive scope exceeded the R1 partial boundary

Initial `REQ-073` pulled archive handling into M0 even though the R1 partial rationale activated retention/supersession while broader privacy/deletion lifecycle remained later.

Resolution: M0 requirement now covers current-versus-superseded semantics and retention; archive/deletion and broader M1 privacy lifecycle remain deferred.

Status: `RESOLVED`.

### R2-F04 — incident requirement imported too much future program

Initial `REQ-074` reproduced nearly the full future security-incident program despite R1 only partially applying it to M0 state/security failures.

Resolution: requirement is bounded to detection of material invalid/corrupt state, unsafe-use containment, evidence preservation and recovery/review hooks; broader incident-response program is deferred.

Status: `RESOLVED`.

### R2-F05 — durable-engine guard overconstrained later legitimate choice

Initial `REQ-105` said an engine could be considered only for a need beyond the M0 operational-state lifecycle. That could prohibit R4 from selecting one if accepted M0 requirements/evidence themselves prove it necessary.

Resolution: R2 prohibits introducing a durable engine merely because restartable state exists, while allowing R4 to consider one if accepted M0 requirements and evidence demonstrate necessity and proportionality.

Status: `RESOLVED`.

### R2-F06 — artifact-reference modal was weaker than source

Initial `REQ-014` used `SHOULD` where the active Project requirement establishes a stronger boundary against storing all referenced content inline.

Resolution: M0 Project state `MUST` preserve stable references/provenance when referring to artifacts/evidence and `MUST NOT` require duplicating full referenced content inline.

Status: `RESOLVED`.

### R2-F07 — evolutionary adapter modal was weaker than source

Initial `REQ-107` used `SHOULD` even though the active architecture source requires evolution through adapters rather than domain rewrite.

Resolution: the requirement now uses `MUST` while still saying later distributed stages need not be implemented in M0.

Status: `RESOLVED`.

### R2-F08 — secret hygiene needed a tighter M0 boundary

Initial `REQ-069` could be read as implying prompts/manifests are M0 runtime surfaces. They are not M0 dependencies.

Resolution: M0 is prohibited from requiring secrets to enter those future surfaces, while current general telemetry/logs remain prohibited from carrying secret values through the normal persistence/inspection/recovery path.

Status: `RESOLVED`.

## 9. Inference review

Several requirements are more precise than their constitutional source by design; that is the purpose of R2. The review checked that precision did not become unowned new product meaning.

### Restore identity collision — `REQ-062`

Requiring collision to fail or undergo explicit governed resolution is a direct safety consequence of stable identity, no silent overwrite, integrity and current-state ownership. It does not select merge strategy or storage mechanism.

Result: acceptable R2 derivation.

### Accepted-state/evidence consistency — `REQ-026`

The requirement says a partial/ambiguous update cannot be reported as successful and resulting state/evidence must be consistent. It does not prescribe database transactions, consensus or another mechanism.

Result: implementation-neutral.

### R3 threat-model minimum — `REQ-075`

This requirement identifies threat classes that R3 cannot ignore; it does not perform or decide the threat model in R2.

Result: correct downstream readiness dependency.

## 10. No hidden technical commitment

The final package selects none of the following:

- Core programming language/runtime;
- database or operational-state store;
- state-versus-event implementation pattern;
- schema/serialization/code-generation technology;
- event broker/transport/audit backend;
- telemetry backend/collector/framework;
- backup/archive topology or concrete format;
- migration tool;
- process/deployment topology;
- UI technology;
- durable execution engine;
- Architecture Spike IDs, candidates, procedures or winners.

`REQ-096` through `REQ-107` explicitly guard these open decisions rather than decide them.

No technology candidate name appears as a selected mechanism in a normative R2 requirement.

## 11. Requirement risk distribution

```text
critical  51
high      53
medium    18
--------------
total    122
```

The high proportion of critical/high requirements is expected for M0 because it creates Aurora's first sovereign identity/state/authority durability boundary. Risk labels direct later verification rigor; they do not imply implementation complexity or a specific architecture.

## 12. Remaining open work belongs to later gates

### R3 — if separately authorized

R3 must allocate all 122 requirements to a reusable Capability/System Spec and test plan, including:

- minimum domain model and exact logical ownership;
- state/lifecycle model;
- authority-state/snapshot semantics;
- security/privacy classification and threat model;
- failure/recovery semantics;
- observability/evidence model;
- compatibility/migration behavior;
- requirement-to-test coverage.

### R4 — later and separately authorized

Only after R3 may the project close implementation-blocking technical uncertainty and make technical decisions/spikes about language/runtime, storage, topology, schema representation, audit/event mechanism, backup/restore/migration mechanism and any durable engine need.

### R5/R6/R7

Mission Contract, Microdesign and implementation remain later independent gates.

## 13. Gate checklist

| R2 condition | Result |
|---|---|
| 127 R1-active sources accounted for | PASS — 127/127 |
| derived requirements exist | PASS — 122 |
| stable unique requirement IDs | PASS — `001`–`122` |
| normative statements | PASS |
| source traceability | PASS — 409 links |
| no active source orphan | PASS — 0 |
| no inactive/deferred source reactivated | PASS — 0 |
| risk classification present | PASS |
| high-risk verification direction present | PASS |
| verification methods use accepted vocabulary | PASS |
| duplicates/conflicts reviewed | PASS — none unresolved |
| R1 partial boundaries preserved | PASS after corrections |
| technology/stack remains open | PASS |
| Architecture Spikes not executed | PASS |
| R3 design/threat-model execution not performed | PASS |
| Mission Contract/Microdesign/implementation not performed | PASS |
| documentation validation | PASS |

## 14. Documentation impact

```yaml
documentation_impact:
  status: UPDATED
  affected:
    - docs/acceptance/2026-08-07-m0-r2-operator-authorization.md
    - docs/capabilities/CAP-SOVEREIGN-CORE/REQUIREMENTS.md
    - docs/capabilities/CAP-SOVEREIGN-CORE/R2-COVERAGE.md
    - docs/reviews/2026-08-07-m0-r2-requirements-review.md
  rationale: R2 creates the first atomic Capability requirements package for M0.
  follow_up: tracking closeout after canonical integration
```

## 15. Verdict and stop boundary

```text
R2 PASS
```

Exact boundary after canonical closeout:

```text
R2 PASS recorded
→ STOP
→ await explicit operator authorization for M0 ACRM R3 — Capability Readiness
```

R2 PASS does not authorize R3, R4, Architecture Spikes, stack selection, Mission Contract, Microdesign or implementation.
