---
id: REVIEW-AURORA-M0-R1-APPLICABILITY-2026-08-07
title: M0 ACRM R1 Applicability Review
document_type: adversarial_review
form: explanation
authority: evidence
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - M0 ACRM R1 applicability review observations and verdict
related:
  - DOC-AURORA-STATUS
  - DOC-AURORA-CAP-SOVEREIGN-CORE-APPLICABILITY
  - DOC-AURORA-M0-R1-OPERATOR-AUTHORIZATION
  - REVIEW-AURORA-M0-R0-CONSTITUTIONAL-BASELINE-RERUN-2026-08-07
source_revision: 735f269025e2cc317424e4931f3a5cd414cd6f2a
applicability_revision: 7f10734ba6018154f196557de6c5735719046253
reviewed_at: 2026-08-07
last_reviewed: 2026-08-07
---

# M0 ACRM R1 Applicability Review

## 1. Executive verdict

```text
R1 PASS
```

Capability:

```text
CAP-SOVEREIGN-CORE
```

Fixed constitutional/source baseline:

```text
735f269025e2cc317424e4931f3a5cd414cd6f2a
```

Applicability artifact revision reviewed:

```text
7f10734ba6018154f196557de6c5735719046253
```

R1 is satisfied because every accepted constitutional requirement was explicitly considered, the M0-active subset is traceable, exclusions/deferments are justified, high-risk cross-cutting dependencies are identified, and no unresolved applicability conflict requires a new constitutional/ADR decision before R2.

This verdict does **not** authorize R2 or any technical/implementation work.

## 2. Gate definition used

The accepted Aurora Capability Realization Method defines R1 as the applicability gate answering which constitutional, security, memory, authority, reliability and documentation requirements apply to the Capability.

Required process:

1. list source sections/principles;
2. classify applicability;
3. record rationale;
4. identify cross-capability dependencies;
5. identify conflicts/open research;
6. assign owner.

Gate condition:

- every known cross-cutting source considered;
- no unjustified exclusions;
- high-risk dependencies identified.

R2 atomic requirement derivation is explicitly a later gate and was not executed here.

## 3. Coverage result

The applicability artifact classifies all 294 accepted A0 constitutional requirements.

| Classification | Count |
|---|---:|
| `APPLIES` | 78 |
| `PARTIALLY_APPLIES` | 49 |
| `DEFERRED_BY_ROADMAP` | 161 |
| `NOT_APPLICABLE` | 6 |
| `CONFLICT_REQUIRES_DECISION` | 0 |
| **Total** | **294** |

The active source set for a future authorized R2 is therefore:

```text
78 APPLIES
+ 49 PARTIALLY_APPLIES
= 127 active constitutional source rows
```

This is a source-coverage count, not a prediction that R2 will contain exactly 127 atomic requirements. R2 may split or consolidate statements while retaining traceability.

## 4. Positive applicability review

The review checked that M0 did not omit the constitutional concerns most likely to be lost in a narrow persistence implementation.

### Sovereign identity — included

Stable Aurora and Project identity across process restart is active. Model, transcript, interface, Harness, runtime and database identity cannot substitute for Aurora's domain identity.

### Project/current operational state — included

Project/current state and next-action continuity are active. Operational state is explicitly distinguished from conversation history, model memory, Harness local state, workflow-engine history, Git and logs.

### Authority snapshot — included

Operator/final authority, access-versus-authority, scope/expiry/revocation and restore safety are active or partially active. The applicability artifact explicitly rejects treating a persisted text value such as “allowed next action” as an authority source by itself.

### Restart and recovery — included

Fresh-process reconstruction, valid/invalid transition behavior, classified recovery and state-outliving-process requirements are active.

### Sovereignty/export/restore/migration — included

Local sovereign ownership, export, backup/restore, migration and restore integrity/safety are active or partially active.

### Event/audit/evidence/telemetry separation — included

The M0 event/audit minimum and telemetry baseline are active without allowing events/logs/telemetry to become the source of current state. Claim, Evidence and Verdict distinctions remain active.

### Security — included

The first canonical durable state store triggers an R3 threat-model dependency. Data classification, integrity, authority corruption, backup sensitivity and restore safety were not excluded merely because M0 has no cloud/provider/device capability.

### Framework/stack neutrality — included

ADR-0001 and architecture invariants constrain M0 against framework/database/runtime ownership without selecting a mechanism.

### Documentation/authorization governance — included

R1 artifacts, future gate boundaries, explicit authorization and documentation findings remain governed and traceable.

## 5. Adversarial false-inclusion review

The matrix was also challenged for importing future product complexity into M0.

Correctly deferred:

- M1 conversational/project memory lifecycle and Context Builder;
- M2 Capability Registry, Provider trust lifecycle, AHDK and conformance;
- M3+ full Delegation/Effect/PDP/Gateway/Credential-Broker behavior;
- M7 adaptive campaigns;
- M8 Presence Fabric, handoff, sensors and multi-device behavior;
- M9/M10 laboratory/device/physical actuation;
- M11 self-improvement and promotion machinery.

A negative invariant may still apply even when the future subsystem is deferred. For example, Harness/framework state must not become M0's global authority/state owner.

## 6. `NOT_APPLICABLE` audit

Exactly six requirements are `NOT_APPLICABLE`, all with explicit rationale:

- `AUR-REQ-VIS-012` — historical A0 implementation gate;
- `AUR-REQ-RDM-006` — historical A0 closeout requirement;
- `AUR-REQ-DOC-005` — repository-global Blueprint publication requirement already satisfied by A0, not CAP behavior;
- `AUR-REQ-DOC-014` — R1 applicability is a canonical source artifact, not a generated projection;
- `AUR-REQ-DOC-019` — historical A0 fresh-session proof;
- `AUR-REQ-DOC-022` — historical A0 implementation-block requirement; current gate prohibition is owned by ACRM/STATUS.

No product/security/runtime obligation was excluded merely for implementation convenience.

## 7. High-risk dependencies

The applicability artifact identifies and assigns seven material dependencies:

1. **Authority snapshot correctness and restore safety** — R2/R3 must define minimum authority/freshness/validity semantics before R4 mechanism choice.
2. **Operational-state ownership** — current state cannot collapse into transcript, memory, Harness, engine history, Git or telemetry.
3. **Stable identity and migration** — identity/version invariants precede schema/store/migration choice.
4. **Backup/restore integrity and portability** — R2/R3 define behavioral contract before R4 mechanism.
5. **Event/audit versus telemetry** — semantics/evidence remain separate from backend/transport choice.
6. **Security of the first canonical durable store** — R3 threat model is required for state/authority/backup risk.
7. **Simple proof surface is not Presence Fabric** — CLI/simple interface cannot own state/identity or force UI architecture.

These are downstream readiness dependencies, not permission to execute R2/R3/R4.

## 8. Conflicts and open mechanisms

```text
CONFLICT_REQUIRES_DECISION = 0
```

No applicability conflict requires changing the accepted Product Blueprint or ADR set before requirement derivation.

The following remain open mechanisms rather than applicability conflicts:

- Core language/runtime;
- process/deployment topology;
- operational state store;
- state-versus-event implementation;
- schema representation;
- event/audit mechanism;
- telemetry backend;
- backup/restore mechanism/topology;
- migration mechanism;
- durable engine only if later requirements prove need;
- exact Architecture Spike scopes/IDs/candidates.

They remain R4 territory after R2/R3 and require their applicable separate authorizations.

## 9. ADR applicability

### ADR-0001

Applies as a boundary/ownership guard: M0 domain/state semantics remain language-, framework- and transport-independent. It does not cause selection of MCP, A2A, RPC, database, framework or runtime.

### ADR-0002

Accepted but deferred as M0 behavior because Registry/AHDK/Harness integration is explicitly not an M0 dependency. Its acceptance does not authorize AHDK implementation.

## 10. Mechanical verification

After `APPLICABILITY.md` was committed, the normal repository Documentation workflow executed against exact applicability revision:

```text
7f10734ba6018154f196557de6c5735719046253
```

Workflow run:

```text
31146692949
```

Result:

```text
SUCCESS
```

Mechanical validation is supporting evidence only. The R1 verdict additionally depends on the adversarial inclusion/exclusion and dependency review above.

## 11. Gate checklist

| R1 condition | Result |
|---|---|
| known cross-cutting sources considered | PASS |
| all 294 accepted constitutional requirements explicitly classified | PASS |
| `NOT_APPLICABLE` exclusions justified | PASS |
| future-roadmap scope not silently imported | PASS |
| negative invariants preserved where needed | PASS |
| high-risk dependencies identified/assigned | PASS |
| applicability conflicts identified | PASS — none unresolved |
| open technical mechanisms kept out of R1 | PASS |
| stack/Architecture Spike execution not selected | PASS |
| R2 atomic requirement derivation not executed | PASS |

## 12. Verdict and stop boundary

```text
R1 PASS
```

Exact next permitted sequence:

```text
record R1 PASS
→ stop at the R1 boundary
→ await explicit operator authorization for M0 ACRM R2 — Requirements
```

R2 is **not** authorized by this verdict. No atomic Capability requirements, Capability/System Spec, threat model execution, Architecture Spike, technical decision, Mission Contract, Microdesign or implementation may begin by implication.
