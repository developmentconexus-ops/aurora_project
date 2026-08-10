---
id: DOC-AURORA-CAP-SOVEREIGN-CORE-TEST-PLAN
title: CAP-SOVEREIGN-CORE Capability Test Plan
document_type: capability_test_plan
form: reference
authority: specification
status: accepted
accepted_at: 2026-08-09
acceptance_evidence: DOC-AURORA-M0-R5-A2-CONTRACT-OPERATOR-ACCEPTANCE
accepted_from_blob: 8b42cc451439038e63e8b567702877b8951c5edb
version: 0.2.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - verification and evidence plan for CAP-SOVEREIGN-CORE
related:
  - DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC
  - DOC-AURORA-CAP-SOVEREIGN-CORE-THREAT-MODEL
  - DOC-AURORA-CAP-SOVEREIGN-CORE-REQUIREMENTS
  - DOC-AURORA-CAP-SOVEREIGN-CORE-R3-COVERAGE
  - DOC-AURORA-CAP-SOVEREIGN-CORE-R4-DECISION-COVERAGE
  - DOC-AURORA-CAP-SOVEREIGN-CORE-R5-COVERAGE
  - DOC-AURORA-MIS-M0-SOVEREIGN-CORE-001
  - DOC-AURORA-M0-R5-A2-CONTRACT-OPERATOR-ACCEPTANCE
source_revision: 9ea8adf5c115f54071d7e36e312695d19420d8b0
r4_alignment_revision: 74167bd1404d9076423ffdbae20f97958283527c
review_triggers:
  - R2 requirement change
  - R3 Spec or threat-model change
  - R4 mechanism decision changes verification feasibility
  - evidence or failure-model finding
last_reviewed: 2026-08-07
---

# CAP-SOVEREIGN-CORE — Capability Test Plan

## 1. Purpose

This plan defines how CAP-SOVEREIGN-CORE requirements and reusable design are verified in later authorized gates. Version 0.2.0 keeps the R3 behavioral tests while binding them to the accepted R4 architecture class.

R3 originally planned tests without choosing mechanisms. R4 has now selected the Go/SQLite/owner-trust architecture. This plan still does not select source files, Go test libraries, CLI syntax or the R6 fault-injection implementation.

The plan must make it possible for R5/R6 to allocate exact criteria/tasks without discovering missing product semantics.

---

## 2. Verification principles

1. **Evidence before success claims.** Dispatch/exit status alone is not success.
2. **Negative cases are first-class.** Invalid transition, authority failure, corruption and unsafe restore must be proven rejected/contained.
3. **Canonical state is inspected independently from UI narrative.** Tests verify stored/recovered semantics, not only rendered output.
4. **Fresh-process tests remove in-memory help.** Restart cases terminate all Aurora processes before recovery.
5. **Faults are injected at material boundaries.** Crash, stale revision, corruption, version mismatch and unavailable non-canonical sinks must be explicit.
6. **Security cases fail closed.** An inability to establish authority/integrity cannot become permission.
7. **Behavioral criteria stay stable across bindings.** Accepted R4 mechanisms constrain implementation but cannot weaken the expected product behavior.
8. **Fixed revision evidence.** Every material proof identifies exact code/spec/schema/runtime revisions applicable to the run.


### 2.1 R4-aligned execution baseline

R5/R6/R7 planning uses these accepted/evidence-qualified bindings:

```text
Core:                    Go, one local modular process
operational store:       SQLite + database/sql + modernc.org/sqlite
persistence posture:     WAL + synchronous=FULL
portable state:          JSON Schema 2020-12 + JSON/JCS
observability:           OTel traces/metrics + slog; exporter optional
owner root:              random ORK; Argon2id KEK; AES-256-GCM wrapped
integrity:               HKDF-SHA-256 purpose keys + HMAC-SHA-256
restore freshness:       REVALIDATION_REQUIRED + owner-only new authority revision
durable workflow engine: none in M0
Mastra/AHDK/model:       not required by M0
```

Evidence-qualified starting versions are Go 1.26.5, `modernc.org/sqlite` v1.54.0, compatible `modernc.org/libc` v1.74.1 and `golang.org/x/crypto` v0.54.0 under CGO=0. R6 must revalidate exact implementation pins; semantic mechanism changes require replan.

R6 must also design bounded Argon2 envelope parsing and target filesystem publication/fsync semantics before R7 can claim the corresponding guarantees.

---

## 3. Test levels

| Level | Purpose |
|---|---|
| `DOCUMENT_REVIEW` | verify ownership, scope, no hidden technology/roadmap commitment and artifact completeness |
| `STATIC_ANALYSIS` | verify forbidden dependency/type ownership after implementation exists |
| `SCHEMA_VALIDATION` | verify semantic contracts/required fields/version identities |
| `UNIT_TEST` | verify pure decision/state functions where later implementation permits isolation |
| `CONTRACT_TEST` | verify command/result/domain invariants across replaceable implementations |
| `INTEGRATION` | verify Core + durable/adapters as one vertical slice |
| `FAULT_INJECTION` | verify crash/corruption/ambiguity/recovery containment |
| `SECURITY_TEST` | verify authority/injection/confidentiality/integrity boundaries |
| `USER_JOURNEY` | verify operator-visible M0 flows |
| `BENCHMARK/EVAL` | verify later operational burden/efficiency only after R4 defines measurable mechanisms |
| `OPERATOR_VERDICT` | evaluate M0 Golden Proof/limitations at later closeout |

---

## 4. Deterministic fixture corpus

R3 defines logical fixtures, not serialization files.

### F-01 — empty sovereign context

No Aurora identity and no Project.

### F-02 — initialized Aurora

One stable Aurora identity plus owner operator identity and initial owner authority state.

### F-03 — Project with first accepted state

One Project, accepted revision `R1`, proposed next action and valid matching authority.

### F-04 — Project with revision history

Current revision `R3`, historical `R1/R2`, plus a stale transition request expecting `R1`.

### F-05 — expired authority

Authority record structurally active but current evaluation time is after `valid_until`.

### F-06 — revoked authority

Authority record revoked before process restart.

### F-07 — pre-revocation export

Export captured while authority was active; later live canonical state contains a revocation not present in the export.

### F-08 — corrupt state/export

Material bytes/semantic fields or integrity reference altered.

### F-09 — incompatible logical version

Persisted/exported version has no declared direct compatibility with current reader.

### F-10 — identity collision

Restore package contains Aurora/Project identity conflicting with an existing different target identity.

### F-11 — untrusted Project content

Project state contains strings/structures resembling policy, grants or system instructions.

### F-12 — ambiguous operation

Crash boundary leaves operation result unknown until canonical state is reconciled.

### F-13 — non-canonical sink unavailable

Telemetry/log backend unavailable while canonical state remains accessible.

### F-14 — time rollback

Time source moves backwards enough to affect an authority expiry decision.

### F-15 — migration semantic mutation

Candidate migration changes stable identity or authority meaning and must be rejected.

---

## 5. Test catalog — identity and scope

| Test ID | Method | Case | Expected result |
|---|---|---|---|
| `T-ID-001` | INTEGRATION, USER_JOURNEY | initialize from F-01 | one stable Aurora identity and owner authority created |
| `T-ID-002` | FAULT_INJECTION, INTEGRATION | restart F-02 | same Aurora identity recovered; no process/session identity substitution |
| `T-ID-003` | CONTRACT_TEST | initialize already initialized store | reject/no silent Aurora replacement |
| `T-ID-004` | INTEGRATION | create Project then restart | same Project identity recovered |
| `T-ID-005` | DOCUMENT_REVIEW, STATIC_ANALYSIS | inspect dependency/type boundaries | no model/Harness/framework/database/UI becomes identity/domain owner |
| `T-SCOPE-001` | DOCUMENT_REVIEW | inspect M0 package | M1/M2/MNFS/cloud/device behavior not required by M0 |
| `T-SCOPE-002` | DOCUMENT_REVIEW | inspect operator boundary | one single-user/Leandro-first semantic scope; no hidden multi-tenancy |

---

## 6. Test catalog — canonical ownership and state

| Test ID | Method | Case | Expected result |
|---|---|---|---|
| `T-STATE-001` | CONTRACT_TEST | inspect current Project | exactly one governing accepted revision referenced |
| `T-STATE-002` | FAULT_INJECTION | remove/break current pointer while history exists | fail/degraded; history/logs not promoted into current truth |
| `T-STATE-003` | CONTRACT_TEST | stale narrative conflicts with current state | current accepted revision governs |
| `T-STATE-004` | SCHEMA_VALIDATION | validate Project/state relationships | stable IDs, revision, predecessor, attribution and required provenance present |
| `T-STATE-005` | DOCUMENT_REVIEW, STATIC_ANALYSIS | inspect persistence/domain types | storage/protocol/framework-specific semantics do not own domain meaning |
| `T-STATE-006` | CONTRACT_TEST | inspect cached projection after source revision changes | projection stale/recomputed; cannot govern independently |

---

## 7. Test catalog — state transitions

| Test ID | Method | Case | Expected result |
|---|---|---|---|
| `T-TRANS-001` | CONTRACT_TEST, INTEGRATION | valid transition against current revision and authority | exactly one new accepted revision; current pointer advances |
| `T-TRANS-002` | CONTRACT_TEST, USER_JOURNEY | stale expected revision from F-04 | reject; current state unchanged |
| `T-TRANS-003` | CONTRACT_TEST | malformed state envelope | reject; no accepted revision |
| `T-TRANS-004` | SECURITY_TEST | unauthorized transition | reject; state/authority/next action unchanged |
| `T-TRANS-005` | CONTRACT_TEST | attempt to mutate stable Aurora/Project identity through state payload | reject |
| `T-TRANS-006` | FAULT_INJECTION | crash around transition commit | one coherent accepted revision or explicit ambiguous/failure; never duplicate success |
| `T-TRANS-007` | CONTRACT_TEST | corrective new revision | prior revision remains historical; new revision attributable |
| `T-TRANS-008` | INTEGRATION | accepted transition then inspect | new state, audit/evidence references and next-action sources mutually consistent |

---

## 8. Test catalog — authority and next safe action

| Test ID | Method | Case | Expected result |
|---|---|---|---|
| `T-AUTH-001` | SECURITY_TEST | technical access without owner authority | no authority mutation/state mutation permission |
| `T-AUTH-002` | CONTRACT_TEST | matching active grant + state candidate | authority snapshot `VALID`; matching next action may be `PERMITTED` |
| `T-AUTH-003` | SECURITY_TEST | wrong Project/action scope | blocked/not permitted |
| `T-AUTH-004` | SECURITY_TEST, FAULT_INJECTION | F-05 expired authority | `EXPIRED`; no permission |
| `T-AUTH-005` | SECURITY_TEST, INTEGRATION | revoke then restart F-06 | revocation recovered and governing |
| `T-AUTH-006` | CONTRACT_TEST | supersede authority | older record non-governing; new revision traceable |
| `T-AUTH-007` | CONTRACT_TEST | next action candidate without matching authority | projection `BLOCKED`/`NONE`; candidate not grant |
| `T-AUTH-008` | FAULT_INJECTION | missing/corrupt authority state | fail closed; diagnosable blocked state |
| `T-AUTH-009` | SECURITY_TEST | Project content attempts to grant authority | no authority change |
| `T-AUTH-010` | CONTRACT_TEST | change state/authority source revision | cached snapshot/next action invalidated |
| `T-AUTH-011` | SECURITY_TEST | authority admin by non-owner actor | reject |
| `T-AUTH-012` | FAULT_INJECTION, SECURITY_TEST | F-14 time rollback | expired/non-provable authority does not regain permission |

---

## 9. Test catalog — restart and recovery

| Test ID | Method | Case | Expected result |
|---|---|---|---|
| `T-REC-001` | USER_JOURNEY, INTEGRATION | terminate all Aurora processes then start fresh | same identities/current state/authority validation/next action recovered |
| `T-REC-002` | FAULT_INJECTION | remove in-memory/session/model/Harness context | recovery unchanged because those are not authority |
| `T-REC-003` | FAULT_INJECTION | canonical state unavailable | explicit degraded/failed result; no fabricated state |
| `T-REC-004` | FAULT_INJECTION, SECURITY_TEST | canonical state corrupt F-08 | blocked governing use; evidence/failure classification retained |
| `T-REC-005` | CONTRACT_TEST | inspect structured recovery result | state revision, authority outcome, classification/limitations present |
| `T-REC-006` | FAULT_INJECTION | stale cache after restart | cache ignored/recomputed from canonical sources |
| `T-REC-007` | CONTRACT_TEST | recovery error taxonomy | required M0 failure classes distinguishable |

---

## 10. Test catalog — export, restore and migration

| Test ID | Method | Case | Expected result |
|---|---|---|---|
| `T-PORT-001` | CONTRACT_TEST, INTEGRATION | export F-03 | package manifest contains required logical state/version/integrity metadata |
| `T-PORT-002` | SECURITY_TEST, FAULT_INJECTION | corrupt export F-08 | integrity validation fails before apply |
| `T-PORT-003` | CONTRACT_TEST | incompatible version F-09 | explicit migration required or clear failure; no silent coercion |
| `T-PORT-004` | INTEGRATION, USER_JOURNEY | restore valid export into fresh context | identities/state restored; authority safety evaluation applied |
| `T-PORT-005` | SECURITY_TEST, FAULT_INJECTION | restore F-07 pre-revocation export | apparently active authority becomes `REVALIDATION_REQUIRED` absent freshness proof |
| `T-PORT-006` | SECURITY_TEST | restore explicitly revoked/expired authority | remains non-permitting |
| `T-PORT-007` | SECURITY_TEST, FAULT_INJECTION | identity collision F-10 | no silent overwrite/merge |
| `T-PORT-008` | CONTRACT_TEST | inspect restore result | versions, identities, integrity/authority outcome, limitations/evidence refs present |
| `T-PORT-009` | INTEGRATION | valid explicit migration | identity/state/authority/provenance/evidence semantics preserved |
| `T-PORT-010` | FAULT_INJECTION | migration semantic mutation F-15 | reject/no silent governing replacement |
| `T-PORT-011` | FAULT_INJECTION | ambiguous crash during restore/migration | reconcile before retry; no duplicate/unknown success |
| `T-PORT-012` | DOCUMENT_REVIEW, USER_JOURNEY | inspect export/backup governance | Leandro can inspect material result; package classification at least `SENSITIVE` |
| `T-PORT-013` | SECURITY_TEST, USER_JOURNEY | after F-07 restore, attempt authority revalidation first as non-owner and then as authenticated owner | non-owner denied; owner creates a new attributable authority-state revision and only then may current authority return to `VALID` if scope/time/conditions pass |

---

## 11. Test catalog — security and sovereignty

| Test ID | Method | Case | Expected result |
|---|---|---|---|
| `T-SEC-001` | SECURITY_TEST | untrusted-content injection F-11 | content remains data; no policy/authority reinterpretation |
| `T-SEC-002` | DOCUMENT_REVIEW, SCHEMA_VALIDATION | inspect data-family classifications | every material M0 family has explicit accepted-class vocabulary assignment |
| `T-SEC-003` | STATIC_ANALYSIS, SECURITY_TEST | inspect normal telemetry/log path | no secret/sensitive payload included merely for correlation |
| `T-SEC-004` | SECURITY_TEST | spoof actor/operator context | mutation denied unless authentication boundary establishes owner identity |
| `T-SEC-005` | DOCUMENT_REVIEW, INTEGRATION | remove model/provider availability | canonical state/recovery remains available subject to local store |
| `T-SEC-006` | DOCUMENT_REVIEW | inspect threat-model coverage | all R2 minimum threats + rollback/time/atomicity covered |

---

## 12. Test catalog — audit, evidence and telemetry

| Test ID | Method | Case | Expected result |
|---|---|---|---|
| `T-EVID-001` | CONTRACT_TEST | accepted and rejected transition | audit differentiates attempt/result/reason and state revision refs |
| `T-EVID-002` | INTEGRATION | restart/recovery | recovery boundary/result correlated to Aurora/Project/state IDs |
| `T-EVID-003` | INTEGRATION | export/restore/migration | each operation produces attributable result/evidence metadata |
| `T-EVID-004` | CONTRACT_TEST | inspect EvidenceRecord | criterion, producer, verifier, method, environment, versions, refs, uncertainty/limitations present |
| `T-EVID-005` | DOCUMENT_REVIEW, CONTRACT_TEST | claim vs evidence vs verdict | implementation claim/exit code cannot establish verdict |
| `T-EVID-006` | FAULT_INJECTION | telemetry sink unavailable F-13 | canonical state remains owned by Core; missing required proof is explicit rather than fabricated |
| `T-EVID-007` | DOCUMENT_REVIEW | inspect retained metrics/signals | each states proof/decision/threshold purpose and limitation where applicable |
| `T-EVID-008` | SCHEMA_VALIDATION | inspect proof-run correlation | stable proof-run/Aurora/Project/operation/revision links present |

---

## 13. Test catalog — reliability and architecture guards

| Test ID | Method | Case | Expected result |
|---|---|---|---|
| `T-REL-001` | FAULT_INJECTION | retry unsafe/ambiguous operation | blind retry prohibited |
| `T-REL-002` | FAULT_INJECTION | retry classified safe/idempotent case | retry allowed without concealing systematic failure |
| `T-REL-003` | DOCUMENT_REVIEW | material verification result | exact target revision, limitations/residual risks explicit |
| `T-REL-004` | DOCUMENT_REVIEW | local component tests all green | M0 not closed without end-to-end Golden Proof/R8 verdict |
| `T-ARCH-001` | DOCUMENT_REVIEW | inspect logical boundaries | one owner per durable concept; adapters/projections distinct |
| `T-ARCH-002` | DOCUMENT_REVIEW | inspect topology proposal | logical modularity first; distribution requires evidence |
| `T-ARCH-003` | DOCUMENT_REVIEW | inspect implementation/design proposal | implementation conforms to accepted ADR-0003..0008 and no selected mechanism becomes domain authority |
| `T-ARCH-004` | DOCUMENT_REVIEW | inspect durable-engine proposal | engine considered only if M0 need is demonstrated |
| `T-ARCH-005` | DOCUMENT_REVIEW | inspect memory integration | M1 memory/vector/Context Builder not used as operational-state owner |
| `T-ARCH-006` | DOCUMENT_REVIEW | inspect future adapter evolution | domain meaning remains stable across binding/topology change |

---

## 14. Test catalog — documentation/readiness

| Test ID | Method | Case | Expected result |
|---|---|---|---|
| `T-DOC-001` | DOCUMENT_REVIEW | inspect R3 package | IDs/authority/status/owners/source baseline/relations declared |
| `T-DOC-002` | DOCUMENT_REVIEW | inspect 122 requirements | each allocated to Spec mechanism and test(s) in R3 coverage |
| `T-DOC-003` | DOCUMENT_REVIEW | inspect mechanism/readiness boundaries | accepted R4 bindings are reflected; remaining implementation details are explicitly R6-owned or trigger replan |
| `T-DOC-004` | DOCUMENT_REVIEW | inspect gate boundary | accepted A2/ADR/Contract artifacts do not auto-authorize R6/R7/R8 |
| `T-DOC-005` | DOCUMENT_REVIEW | inspect repository changes | no constitutional/milestone meaning silently changed |
| `T-DOC-006` | DOCUMENT_REVIEW | adversarial review | false inclusion, hidden stack, missing threat/recovery/test allocation checked |
| `T-DOC-007` | DOCUMENT_REVIEW | fresh read path | STATUS → applicability → requirements → Spec/threat/test/coverage is discoverable |

---

## 15. Representative journeys

### J-001 — M0 Golden Proof

```text
initialize Aurora
→ create Project
→ accept first state + proposed next action
→ inspect permitted/blocked next safe action from authority
→ terminate every Aurora process
→ start fresh process
→ recover same Aurora/Project/current state
→ reject stale/invalid transition
→ export
→ restore into fresh context
→ verify restored identity/state and safe authority handling
```

Evidence must support every major step and later R8 operator-visible value.

### J-002 — Authority continuity versus restore freshness

```text
active authority
→ ordinary restart
→ authority still valid if scope/time/status remain valid
→ create export
→ revoke live authority
→ restore old export
→ old authority does NOT silently become valid
→ revalidation required
→ non-owner revalidation denied
→ authenticated owner performs explicit revalidation
→ new authority-state revision becomes governing only if current scope/time/conditions pass
```

This journey distinguishes normal restart from potentially stale backup restore.

### J-003 — Corruption and containment

```text
valid current state
→ corrupt persisted/export state
→ fresh recovery/restore
→ validation failure
→ no fabricated current state
→ blocked/degraded result + evidence
```

### J-004 — Ambiguous mutation recovery

```text
request valid transition
→ crash at material commit boundary
→ fresh process
→ reconcile canonical state
→ exactly one accepted revision OR explicit failure/ambiguous result
→ no blind duplicate retry
```

---

## 16. Adversarial/fault coverage baseline

The future implementation test suite must include at least:

- stale revision;
- malformed state envelope;
- immutable identity mutation;
- wrong authority scope;
- expired grant;
- revoked grant;
- missing authority;
- corrupt authority;
- pre-revocation backup restore;
- identity collision restore;
- corrupt state/export;
- incompatible version;
- migration semantic drift;
- time rollback;
- ambiguous crash around mutation;
- unavailable canonical store;
- unavailable telemetry sink;
- untrusted content authority injection;
- adapter/framework state substitution attempt;
- unsafe retry.

---

## 17. Evaluation corpus and separation

M0 is deterministic and does not require model-quality datasets. Its evaluation corpus is the versioned fixture/adversarial set in this plan.

Later execution must preserve at least these logical roles:

- **development fixtures** — ordinary valid states and commands;
- **validation fixtures** — fixed expected restart/restore/migration outcomes;
- **adversarial fixtures** — corruption, stale authority, injection, collision and rollback cases;
- **Golden Proof fixture** — fixed operator-visible M0 journey.

R3 does not require a production-shadow or ML holdout corpus because those are outside this deterministic M0 capability.

---

## 18. Success thresholds

For future R7/R8 evidence unless a later approved contract adds stricter criteria:

1. **all contractually required deterministic cases:** 100% pass for the fixed accepted target revision;
2. **critical safety invariants:** zero unauthorized accepted state mutations;
3. **authority:** zero expired/revoked/unproven-restored authority treated as valid permission;
4. **identity:** zero silent Aurora/Project identity replacement or restore collision acceptance;
5. **recovery:** no fabricated current state when canonical state is missing/corrupt;
6. **coverage:** 122/122 R2 requirements mapped to tests/evidence criteria;
7. **evidence completeness:** every material proof record contains the required R3 evidence fields;
8. **Golden Proof:** complete M0 journey passes end to end before R8 can consider acceptance.

Performance/resource thresholds are not fixed at R3 because no mechanism exists to measure them fairly. R4/R5 may add operational-burden limits without weakening the behavioral thresholds above.

---

## 19. Evidence format

A material future test/proof evidence record must identify:

```text
evidence_id
criterion / requirement IDs
test_case_id
producer
verifier
fixed target revision
Capability/spec/schema/runtime versions as applicable
environment reference
fixture/input reference
observed result
artifact/integrity references
uncertainty / limitations
verdict reference when one exists
```

Raw logs are supporting artifacts, not the verdict.

---

## 20. Graduation levels

- **R3/G0:** COMPLETE — reusable behavioral plan and 122/122 allocation.
- **R4/G1:** COMPLETE — technical decisions/spikes make the plan implementable for one local slice.
- **R5/R6/G2:** CURRENT/NEXT — exact Mission criteria are proposed in R5; implementation tests/tasks require later R6.
- **R7/G3:** FUTURE — execute tests/evidence against implementation.
- **R8/G4:** FUTURE — complete M0 Golden Proof and operator verdict.

No level advances from CI/document existence alone.

---

## 21. Governance boundary

This Test Plan is proposed A2 specification authority. It does not authorize test code or Core implementation.

```text
R5 may approve exact Mission criteria
→ R6 separately owns test/implementation decomposition
→ R7 separately owns execution/evidence
```

The current R5 package must stop before R6 unless separately authorized.
