from pathlib import Path
import subprocess

BASELINE = "74167bd1404d9076423ffdbae20f97958283527c"

def read(path):
    return Path(path).read_text(encoding="utf-8")

def write(path, text):
    Path(path).write_text(text.rstrip() + "\n", encoding="utf-8")

def replace_once(path, old, new):
    text = read(path)
    count = text.count(old)
    if count != 1:
        raise SystemExit(f"{path}: expected exactly one match ({count}) for {old[:120]!r}")
    write(path, text.replace(old, new, 1))

def replace_section(path, start_heading, end_heading, new_section):
    text = read(path)
    start = text.find(start_heading)
    if start < 0:
        raise SystemExit(f"{path}: start heading missing: {start_heading}")
    end = text.find(end_heading, start + len(start_heading))
    if end < 0:
        raise SystemExit(f"{path}: end heading missing: {end_heading}")
    write(path, text[:start] + new_section.rstrip() + "\n\n---\n\n" + text[end:])

req = "docs/capabilities/CAP-SOVEREIGN-CORE/REQUIREMENTS.md"
replace_once(req, "version: 0.1.0", "version: 0.1.1")
replace_once(req, "  - proposed M0 ACRM R2 atomic requirements for CAP-SOVEREIGN-CORE", "  - M0 atomic requirements for CAP-SOVEREIGN-CORE")
replace_once(req, "  - DOC-AURORA-M0-R2-OPERATOR-AUTHORIZATION\nsource_revision:", "  - DOC-AURORA-M0-R2-OPERATOR-AUTHORIZATION\n  - DOC-AURORA-CAP-SOVEREIGN-CORE-R5-COVERAGE\n  - DOC-AURORA-MIS-M0-SOVEREIGN-CORE-001\nsource_revision:")
replace_once(req, "All 122 requirements below have `status: proposed`, `allocation: []` and `evidence: []` at R2. A future gate may allocate them, but R2 does not.", "All 122 requirements were proposed at R2. This v0.1.1 preserves their normative statements unchanged while making the package review-ready for R5 operator acceptance. R3 owns mechanism/test allocation and `R5-COVERAGE.md` owns requirement-to-Mission allocation; this document remains `proposed` until explicitly accepted by the operator.")

spec = "docs/capabilities/CAP-SOVEREIGN-CORE/SPEC.md"
replace_once(spec, "version: 0.1.0", "version: 0.2.0")
replace_once(spec, "  - proposed R3 reusable behavior and logical architecture for CAP-SOVEREIGN-CORE", "  - reusable behavior and logical architecture for CAP-SOVEREIGN-CORE")
replace_once(spec, "  - DOC-AURORA-CAP-SOVEREIGN-CORE-R3-COVERAGE\nsource_revision:", "  - DOC-AURORA-CAP-SOVEREIGN-CORE-R3-COVERAGE\n  - DOC-AURORA-CAP-SOVEREIGN-CORE-R4-DECISION-COVERAGE\n  - ADR-AURORA-0003\n  - ADR-AURORA-0004\n  - ADR-AURORA-0005\n  - ADR-AURORA-0006\n  - ADR-AURORA-0007\n  - ADR-AURORA-0008\n  - DOC-AURORA-MIS-M0-SOVEREIGN-CORE-001\nsource_revision:")
replace_once(spec, "source_revision: 9ea8adf5c115f54071d7e36e312695d19420d8b0", "source_revision: 9ea8adf5c115f54071d7e36e312695d19420d8b0\nr4_alignment_revision: " + BASELINE)
replace_once(spec, "This R3 specification fixes the reusable product behavior, domain semantics, state/lifecycle rules, ownership boundaries, authority rules, security/recovery behavior and evidence model needed by M0 without selecting the implementation mechanisms that belong to R4.", "The original R3 revision fixed reusable product behavior and deliberately left implementation mechanisms to R4. R4 has since reached PASS. This v0.2.0 preserves the R3 semantics and binds them to accepted ADR-0003..0008 while leaving source layout, Go API shape, SQL DDL, filesystem wrappers and test implementation to R6.")
replace_once(spec, "R3 does not authorize Architecture Spike execution, Mission Contract creation, Microdesign or implementation.", "This Capability Spec never grants gate authority by itself. R5 is separately authorized for Contract Readiness; R6, implementation and later gates remain separately gated.")
replace_once(spec, "only a new explicit owner revalidation/grant operation, or another R4-approved freshness proof satisfying this Spec, may return it to `VALID`;", "only the accepted ADR-0008 authenticated-owner revalidation/trust path may return it to `VALID` within M0;")
replace_once(spec, "implementation/operational efficiency once R4/R7 provide measurable mechanisms.", "implementation/operational efficiency using accepted R4 mechanisms once R7 provides measured evidence.")
replace_once(spec, "R4 chooses concrete representation.", "Accepted ADR-0005 binds portable logical representation to JSON Schema 2020-12 + JSON/JCS boundaries; accepted ADR-0007 binds the M0 operational store to SQLite through `database/sql` + `modernc.org/sqlite`. R6 defines exact schemas/files/APIs without changing these semantics.")
replace_section(spec, "## 18. Rollout and graduation", "## 19. Open R4 decisions and uncertainty classes", '''## 18. Rollout and graduation

Graduation is gate-based and does not follow document existence automatically.

### G0 — Capability semantics ready

`R3 PASS` established 122/122 allocation, threat model and deterministic test plan.

### G1 — Architecture ready

`R4 PASS` is complete. Accepted ADR-0003..0008 and reviewed SPK-001/SPK-002 make the Capability implementable for one local M0 slice.

### G2 — Contract/design ready

R5 must approve an exact Mission Contract and R6 must separately approve Microdesign/Implementation Plan.

### G3 — Implemented/evidenced

Future R7 executes implementation tests/evidence against the approved Contract/Plan.

### G4 — Product milestone accepted

Future R8 requires the complete M0 Golden Proof and operator verdict.

Current promotion state is `G1`; the R5 package is proposed and cannot advance to G2 without the required operator decisions.''')
replace_section(spec, "## 19. Open R4 decisions and uncertainty classes", "## 20. Requirement coverage", '''## 19. Accepted R4 mechanism bindings and R6-owned details

The fifteen implementation-blocking R4 questions are closed. Their current governing dispositions are:

| ID | Accepted M0 disposition | R6-owned detail / reconsideration boundary |
|---|---|---|
| `R4-Q-CORE-001` | Go Sovereign Core runtime | packages/interfaces/build layout |
| `R4-Q-STORE-001` | SQLite + `database/sql` + `modernc.org/sqlite` | exact schema/connection wrapper/module pins |
| `R4-Q-STATE-001` | current governing state + immutable revisions; audit/events distinct; no full Event Sourcing | exact tables/queries |
| `R4-Q-SCHEMA-001` | JSON Schema 2020-12 + JSON logical form + RFC 8785 JCS where deterministic bytes are required | exact schemas/code bindings |
| `R4-Q-ATOMIC-001` | SQLite transactional state/current-pointer/audit/evidence consistency | transaction API and fault hooks |
| `R4-Q-INTEGRITY-001` | SHA-256 content boundary + ORK-derived HMAC for governing/trust descriptors | exact descriptor fields/APIs |
| `R4-Q-TIME-001` | process-local monotonic duration + authenticated observed wall-time high-water; backward step fails closed | Time Source API/threshold handling |
| `R4-Q-AUTHN-001` | random ORK + Argon2id KEK + AES-256-GCM wrapped root | envelope format/API and bounded KDF parsing |
| `R4-Q-EXPORT-001` | logical portable export distinct from DB backup; normal SENSITIVE export uses `age`; encrypted ORK recovery allowed without importing historical freshness | archive layout/commands/library binding |
| `R4-Q-MIGRATE-001` | explicit application-owned version-pair migration | migration file/runner shape |
| `R4-Q-AUDIT-001` | audit/events logically distinct; M0 may co-locate transactionally | exact tables/retention fields |
| `R4-Q-TELEM-001` | OpenTelemetry traces/metrics + Go `slog`; exporter/backend optional | exact instrumentation/export adapter |
| `R4-Q-TOPOLOGY-001` | one local modular Sovereign Core process | executable/package layout |
| `R4-Q-ENGINE-001` | no durable workflow engine in M0 | reconsider only on accepted current requirement |
| `R4-Q-RESTORE-001` | historical restore → `REVALIDATION_REQUIRED`; restored grant cannot self-revalidate; authenticated owner creates new authority revision | exact recovery command/API |

R6 may choose reversible implementation details inside these decisions. A change to the accepted mechanism class or authority semantics requires replan/ADR rather than an implementation shortcut.''')
text = read(spec)
marker = "## 22. Stop boundary"
pos = text.find(marker)
if pos < 0:
    raise SystemExit("SPEC: stop boundary missing")
write(spec, text[:pos] + '''## 22. Governance boundary

This Spec is the proposed R4-aligned A2 reusable-behavior authority for CAP-SOVEREIGN-CORE.

It does not itself authorize any gate or implementation. Current progression is governed by `STATUS.md` and the Capability Realization Method:

```text
R4 PASS
→ R5 package/operator decision
→ STOP before R6 unless separately authorized
```

R6 may define implementation detail only after the A2 package and Mission Contract are accepted and R5 passes.''')

threat = "docs/capabilities/CAP-SOVEREIGN-CORE/THREAT-MODEL.md"
replace_once(threat, "version: 0.1.0", "version: 0.2.0")
replace_once(threat, "  - proposed R3 security and threat analysis for CAP-SOVEREIGN-CORE", "  - security and threat analysis for CAP-SOVEREIGN-CORE")
replace_once(threat, "  - DOC-AURORA-CAP-SOVEREIGN-CORE-R3-COVERAGE\nsource_revision:", "  - DOC-AURORA-CAP-SOVEREIGN-CORE-R3-COVERAGE\n  - DOC-AURORA-CAP-SOVEREIGN-CORE-R4-DECISION-COVERAGE\n  - ADR-AURORA-0007\n  - ADR-AURORA-0008\n  - DOC-AURORA-MIS-M0-SOVEREIGN-CORE-001\nsource_revision:")
replace_once(threat, "source_revision: 9ea8adf5c115f54071d7e36e312695d19420d8b0", "source_revision: 9ea8adf5c115f54071d7e36e312695d19420d8b0\nr4_alignment_revision: " + BASELINE)
replace_once(threat, "This document is part of the proposed R3 Capability Spec package and MUST NOT select R4 technologies.", "The original R3 revision intentionally did not select mechanisms. This v0.2.0 preserves the threat semantics and records the accepted R4 mitigations from ADR-0007/ADR-0008; exact implementation controls remain R6 work.")
replace_once(threat, "- authenticated operator context is established by an R4-selected mechanism;", "- authenticated operator context follows the accepted ADR-0008 owner-root/bootstrap boundary;")
replace_once(threat, "- R4 must choose a time/rollback approach proportionate to M0;", "- accepted ADR-0008 uses an authenticated observed wall-time high-water plus fail-closed backward-time semantics;")
replace_once(threat, "- mechanism and trust assumptions must be explicit in R4;", "- accepted ADR-0008 defines ORK-derived authenticated governing/trust descriptors and their local threat assumptions;")
replace_once(threat, "- a future R4 mechanism can provide some persistent local storage, operator-authentication boundary, time source and integrity check satisfying the logical contracts;", "- accepted R4 mechanisms provide the M0 local SQLite store, owner-root authentication boundary, time-high-water behavior and authenticated integrity model needed by the logical contracts;")
replace_once(threat, "| ID | Threat | Impact | Required R3 mitigation/behavior | Residual/R4 dependency |", "| ID | Threat | Impact | Required mitigation/behavior | Accepted R4 binding / later obligation |")
replace_section(threat, "## 10. R4 security questions", "## 11. Residual risks accepted only as R3-known constraints", '''## 10. Accepted R4 security bindings

The R3 security questions are now resolved for M0 by accepted decisions and executable evidence:

| Concern | Accepted M0 binding | Later implementation/evidence obligation |
|---|---|---|
| local owner authentication/bootstrap | random 256-bit ORK; passphrase→Argon2id KEK; AES-256-GCM wrapped ORK | versioned envelope; bound/allowlisted KDF parameters |
| local state persistence/atomicity | SQLite + `database/sql` + `modernc.org/sqlite`, WAL, `synchronous=FULL` | exact schema/transaction wrapper and target-environment fault proof |
| storage rollback detection | external authenticated generation high-water + DB HMAC | enforce anomaly preflight at mutation boundary |
| authority time rollback | authenticated observed wall-time high-water; backward movement → `TIME_UNTRUSTED` | exact Time Source abstraction and diagnostics |
| governing integrity | HKDF-SHA-256 purpose keys + HMAC-SHA-256 | exact protected descriptor fields/canonicalization |
| portable export confidentiality | logical JSON/JCS package; normal SENSITIVE outer package protected with `age` | exact archive/library/recipient handling |
| restore freshness | restored active-looking authority → `REVALIDATION_REQUIRED`; owner-only new revision | exact recovery UX/API; never import historical high-water as current |
| migration | explicit version-pair application-owned migration | exact runner/rollback and invariant verification |
| audit/evidence | logically distinct; may be co-located transactionally for M0 | exact persistence/retention fields |
| telemetry | OTel traces/metrics + `slog`, exporter optional/non-authoritative | redaction and exporter-failure proof |

No new security mechanism choice is delegated silently to R6. R6 owns only the implementation details inside these accepted boundaries.''')
replace_section(threat, "## 11. Residual risks accepted only as R3-known constraints", "## 12. Threat-model gate conclusion", '''## 11. Residual risks and carry-forward constraints

Accepted R4 architecture does not eliminate every local threat.

M0 still does not claim to solve:

- total compromise of host/administrator plus replay of **all** local trust artifacts and owner secrets/runtime;
- physical destruction/theft of every state/recovery copy;
- objectively trusted global time — only backward movement below an authenticated observed high-water is detected;
- every literal storage-controller/write-cache/power-loss failure beyond the tested process-kill model;
- accidental owner decision to explicitly reauthorize unsafe historical intent after receiving the required revalidation state;
- future network/provider/device threats outside M0.

Accepted controls that must carry into R6/R7:

- `age`-protected normal SENSITIVE portability envelope and explicit export classification;
- ORK/HKDF/HMAC integrity boundary independent from the operational DB;
- mutation-boundary enforcement for rollback/anchor-lag/time/restore anomalies;
- bounded Argon2id parameter parsing before allocation;
- explicit filesystem publication/fsync/directory-sync design for claimed target durability;
- secret/redaction hygiene in logs/evidence.

These limitations MUST remain visible in the Mission Contract and evidence. They are not permission to weaken fail-closed semantics.''')
replace_once(threat, "The R3 threat model is complete enough to support architecture decision work because:", "The R4-aligned threat model is complete enough to support the proposed M0 Mission Contract because:")
replace_once(threat, "- mechanism-dependent residual risks are named for R4;\n- no security mechanism or vendor is selected.", "- accepted mechanism bindings and remaining implementation/evidence obligations are explicit;\n- no unresolved M0 security mechanism choice is silently delegated to R6.")
replace_once(threat, "This conclusion is subject to the independent R3 adversarial review and does not itself constitute the final R3 verdict.", "This conclusion remains subject to the accepted R3 review, accepted R4 ADRs/spike evidence and the R5 Contract Readiness review; the Threat Model does not authorize implementation.")

test = "docs/capabilities/CAP-SOVEREIGN-CORE/TEST-PLAN.md"
replace_once(test, "version: 0.1.0", "version: 0.2.0")
replace_once(test, "  - proposed R3 verification and evidence plan for CAP-SOVEREIGN-CORE", "  - verification and evidence plan for CAP-SOVEREIGN-CORE")
replace_once(test, "  - DOC-AURORA-CAP-SOVEREIGN-CORE-R3-COVERAGE\nsource_revision:", "  - DOC-AURORA-CAP-SOVEREIGN-CORE-R3-COVERAGE\n  - DOC-AURORA-CAP-SOVEREIGN-CORE-R4-DECISION-COVERAGE\n  - DOC-AURORA-CAP-SOVEREIGN-CORE-R5-COVERAGE\n  - DOC-AURORA-MIS-M0-SOVEREIGN-CORE-001\nsource_revision:")
replace_once(test, "source_revision: 9ea8adf5c115f54071d7e36e312695d19420d8b0", "source_revision: 9ea8adf5c115f54071d7e36e312695d19420d8b0\nr4_alignment_revision: " + BASELINE)
replace_once(test, "This plan defines how the proposed CAP-SOVEREIGN-CORE requirements and R3 design will be verified in later authorized gates.", "This plan defines how CAP-SOVEREIGN-CORE requirements and reusable design are verified in later authorized gates. Version 0.2.0 keeps the R3 behavioral tests while binding them to the accepted R4 architecture class.")
replace_once(test, "R3 plans tests; it does not implement or execute production tests. No test framework, language, database, CI runner, benchmark tool or fault-injection technology is selected here.", "R3 originally planned tests without choosing mechanisms. R4 has now selected the Go/SQLite/owner-trust architecture. This plan still does not select source files, Go test libraries, CLI syntax or the R6 fault-injection implementation.")
replace_once(test, "7. **Technology-neutral criteria stay stable.** R4 may choose mechanisms but cannot change expected product behavior.", "7. **Behavioral criteria stay stable across bindings.** Accepted R4 mechanisms constrain implementation but cannot weaken the expected product behavior.")
insert_after = "8. **Fixed revision evidence.** Every material proof identifies exact code/spec/schema/runtime revisions applicable to the run.\n"
text = read(test)
if text.count(insert_after) != 1:
    raise SystemExit("TEST-PLAN: insertion anchor mismatch")
r4_baseline = '''\n\n### 2.1 R4-aligned execution baseline\n\nR5/R6/R7 planning uses these accepted/evidence-qualified bindings:\n\n```text\nCore:                    Go, one local modular process\noperational store:       SQLite + database/sql + modernc.org/sqlite\npersistence posture:     WAL + synchronous=FULL\nportable state:          JSON Schema 2020-12 + JSON/JCS\nobservability:           OTel traces/metrics + slog; exporter optional\nowner root:              random ORK; Argon2id KEK; AES-256-GCM wrapped\nintegrity:               HKDF-SHA-256 purpose keys + HMAC-SHA-256\nrestore freshness:       REVALIDATION_REQUIRED + owner-only new authority revision\ndurable workflow engine: none in M0\nMastra/AHDK/model:       not required by M0\n```\n\nEvidence-qualified starting versions are Go 1.26.5, `modernc.org/sqlite` v1.54.0, compatible `modernc.org/libc` v1.74.1 and `golang.org/x/crypto` v0.54.0 under CGO=0. R6 must revalidate exact implementation pins; semantic mechanism changes require replan.\n\nR6 must also design bounded Argon2 envelope parsing and target filesystem publication/fsync semantics before R7 can claim the corresponding guarantees.\n'''
write(test, text.replace(insert_after, insert_after + r4_baseline, 1))
replace_once(test, "| `T-ARCH-003` | DOCUMENT_REVIEW | inspect R4 candidate proposal | language/storage/event/schema/telemetry/backup choices not preselected by R3 |", "| `T-ARCH-003` | DOCUMENT_REVIEW | inspect implementation/design proposal | implementation conforms to accepted ADR-0003..0008 and no selected mechanism becomes domain authority |")
replace_once(test, "| `T-DOC-003` | DOCUMENT_REVIEW | inspect open questions | every unresolved current mechanism explicitly assigned to R4 or blocks R3 |", "| `T-DOC-003` | DOCUMENT_REVIEW | inspect mechanism/readiness boundaries | accepted R4 bindings are reflected; remaining implementation details are explicitly R6-owned or trigger replan |")
replace_once(test, "| `T-DOC-004` | DOCUMENT_REVIEW | inspect gate boundary | R3 PASS does not authorize R4/spikes/R5/R6/R7 |", "| `T-DOC-004` | DOCUMENT_REVIEW | inspect gate boundary | accepted A2/ADR/Contract artifacts do not auto-authorize R6/R7/R8 |")
replace_section(test, "## 20. Graduation levels", "## 21. Stop boundary", '''## 20. Graduation levels\n\n- **R3/G0:** COMPLETE — reusable behavioral plan and 122/122 allocation.\n- **R4/G1:** COMPLETE — technical decisions/spikes make the plan implementable for one local slice.\n- **R5/R6/G2:** CURRENT/NEXT — exact Mission criteria are proposed in R5; implementation tests/tasks require later R6.\n- **R7/G3:** FUTURE — execute tests/evidence against implementation.\n- **R8/G4:** FUTURE — complete M0 Golden Proof and operator verdict.\n\nNo level advances from CI/document existence alone.''')
text = read(test)
marker = "## 21. Stop boundary"
pos = text.find(marker)
if pos < 0:
    raise SystemExit("TEST-PLAN: stop boundary missing")
write(test, text[:pos] + '''## 21. Governance boundary\n\nThis Test Plan is proposed A2 specification authority. It does not authorize test code or Core implementation.\n\n```text\nR5 may approve exact Mission criteria\n→ R6 separately owns test/implementation decomposition\n→ R7 separately owns execution/evidence\n```\n\nThe current R5 package must stop before R6 unless separately authorized.''')

dec = "docs/tracking/DECISIONS.md"
replace_once(dec, "version: 0.4.0", "version: 0.5.0")
replace_once(dec, "| D-056 | Mastra is the accepted preferred-first substrate to evaluate for first-party agentic Harnesses while sovereign truth/authority remain Aurora-owned | ADR-0009 | accepted |", "| D-056 | Mastra is the accepted preferred-first substrate to evaluate for first-party agentic Harnesses while sovereign truth/authority remain Aurora-owned | ADR-0009 | accepted |\n| D-057 | SQLite + `database/sql` + `modernc.org/sqlite` is the accepted M0 operational-state baseline | ADR-0007 | accepted |\n| D-058 | M0 owner trust uses a random wrapped ORK + authenticated external generation/time high-water with fail-closed restore/revalidation semantics | ADR-0008 | accepted |")
replace_once(dec, "| O-009 | operational state/event storage | Sovereign Core Capability Spec + spike |\n", "")
replace_once(dec, "| O-015 | exact first Mission Contract for selected M0 | M0 ACRM R5 Mission Contract |", "| O-015 | exact first Mission Contract for selected M0 | `MIS-M0-SOVEREIGN-CORE-001` v0.1.0 proposed in R5; operator approval pending |")
replace_once(dec, "`O-015` previously combined milestone selection and first Contract. The milestone portion is resolved by `D-051`; only the exact Mission Contract remains open and cannot be chosen before its applicable gates.", "`O-015` previously combined milestone selection and first Contract. The milestone portion is resolved by `D-051`. R5 now proposes `MIS-M0-SOVEREIGN-CORE-001` v0.1.0; the decision remains open until explicit operator approval.")

def blob(path):
    return subprocess.check_output(["git", "hash-object", path], text=True).strip()

proposal_blobs = {
    "requirements": blob(req),
    "spec": blob(spec),
    "threat": blob(threat),
    "test": blob(test),
    "contract": blob("docs/capabilities/CAP-SOVEREIGN-CORE/MIS-M0-SOVEREIGN-CORE-001.md"),
    "coverage": blob("docs/capabilities/CAP-SOVEREIGN-CORE/R5-COVERAGE.md"),
}

review_path = "docs/reviews/2026-08-07-m0-r5-contract-readiness-review.md"
review = f'''---
id: REVIEW-AURORA-M0-R5-CONTRACT-READINESS-2026-08-07
title: M0 ACRM R5 Contract Readiness Review
document_type: adversarial_review
form: explanation
authority: evidence
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - M0 R5 Contract Readiness review observations and current verdict
related:
  - DOC-AURORA-STATUS
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DOC-AURORA-M0-R5-OPERATOR-AUTHORIZATION
  - DOC-AURORA-MIS-M0-SOVEREIGN-CORE-001
  - DOC-AURORA-CAP-SOVEREIGN-CORE-R5-COVERAGE
  - DOC-AURORA-CAP-SOVEREIGN-CORE-REQUIREMENTS
  - DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC
  - DOC-AURORA-CAP-SOVEREIGN-CORE-THREAT-MODEL
  - DOC-AURORA-CAP-SOVEREIGN-CORE-TEST-PLAN
  - REVIEW-AURORA-M0-R4-ARCHITECTURE-DECISION-READINESS-RERUN-2026-08-07
source_revision: {BASELINE}
reviewed_at: 2026-08-07
last_reviewed: 2026-08-07
---

# M0 ACRM R5 — Contract Readiness Review

## 1. Executive verdict

```text
R5 BLOCKED
```

This is **not** a scope/design failure. The R5 package is exact and review-ready, but the accepted documentation hierarchy requires an accepted A2 Capability Spec before an approved A3 Contract can become the implementation commitment. The R2/R3 gates intentionally left the A2 normative package `proposed`, and `MIS-M0-SOVEREIGN-CORE-001` is also still `proposed`.

The remaining blocker is therefore an explicit operator decision on the exact normative package + Mission Contract. R6 remains NOT AUTHORIZED.

## 2. Fixed R5 subject

```text
Capability: CAP-SOVEREIGN-CORE
Mission:    MIS-M0-SOVEREIGN-CORE-001
Mission Contract: v0.1.0
R5 source baseline: {BASELINE}
R4: PASS
```

One Mission is appropriate because M0 is one vertical Sovereign Core walking skeleton. Splitting internal identity/state/authority/recovery/evidence responsibilities into independent Missions would add coordination/platform boundaries without independent product outcomes.

## 3. Proposed normative package

R4-aligned A2 candidates:

| Artifact | Version | Proposal blob |
|---|---:|---|
| Requirements | `0.1.1` | `{proposal_blobs["requirements"]}` |
| Capability Spec | `0.2.0` | `{proposal_blobs["spec"]}` |
| Threat Model | `0.2.0` | `{proposal_blobs["threat"]}` |
| Capability Test Plan | `0.2.0` | `{proposal_blobs["test"]}` |

A3 candidate:

| Artifact | Version | Proposal blob |
|---|---:|---|
| `MIS-M0-SOVEREIGN-CORE-001` Mission Contract | `0.1.0` | `{proposal_blobs["contract"]}` |

Reference allocation:

| Artifact | Version | Blob |
|---|---:|---|
| R5 Requirement Allocation | `1.0.0` | `{proposal_blobs["coverage"]}` |

The Requirements revision changes lifecycle/traceability wording only; all 122 requirement statements remain unchanged. Spec/Threat/Test v0.2.0 remove stale “R4 will decide” wording and bind the already-accepted ADR-0003..0008 outcomes without choosing R6 source/API/DDL details.

## 4. R5 gate checklist

| R5 condition | Result |
|---|---|
| R4 complete for current Mission scope | PASS |
| one exact Mission identity/revision/baseline | PASS |
| operator-visible outcome and Golden Proof contribution | PASS |
| explicit scope/non-goals/assumptions/dependencies | PASS |
| contract-level decomposition without Microdesign | PASS |
| all in-scope requirements allocated | PASS — 122/122 |
| authority/prohibitions/complexity budget explicit | PASS |
| evidence profile and thresholds explicit | PASS |
| replan/supersession triggers explicit | PASS |
| no hidden future M1+/Mastra/AHDK/MNFS scope | PASS |
| no R6 implementation detail masquerading as Contract | PASS |
| A2 normative package accepted | **BLOCKED — operator decision required** |
| Mission Contract approved | **BLOCKED — operator decision required** |
| implementation/R6 authorization absent | PASS |

Therefore `R5 PASS` cannot be declared yet.

## 5. Requirement allocation result

`R5-COVERAGE.md` contains exactly:

```text
Capability requirements: 122
allocation rows:          122
unallocated:                0
primary Mission criteria:  12
```

The primary allocation is category-coherent and `REQ-031` is owned by the complete Golden Proof criterion. Detailed tests remain referenced through `R3-COVERAGE.md` / `TEST-PLAN.md`; R5 does not fork their 84-test catalog.

## 6. Mission Contract quality review

The Contract is implementation-exact at the correct level:

- binds the full M0 vertical outcome;
- contains 12 measurable criteria;
- names accepted architecture bindings and evidence-qualified environment class;
- keeps runtime model/cloud/Harness dependencies at zero;
- defines explicit non-goals;
- defines prospective later implementation authority without granting it;
- carries R4 residual risks into R6/R7;
- defines evidence, acceptance thresholds and replan triggers;
- intentionally leaves files/packages/Go interfaces/SQL DDL/CLI syntax/test framework to R6.

No material hidden scope was found.

## 7. A2 lifecycle finding

### R5-F01 — proposed reusable semantics cannot be silently outranked by an approved Contract

Status:

```text
OPEN / GATE BLOCKER / OPERATOR DECISION
```

Blueprint 15 precedence is:

```text
accepted Constitution
→ accepted ADR
→ accepted Capability/System Spec
→ approved scoped Contract
```

R2 and R3 reviews explicitly stated that their normative artifacts remained `proposed` despite gate PASS.

Resolution prepared in this R5 package:

- Requirements v0.1.1 preserves all 122 statements;
- Spec v0.2.0 incorporates completed R4 bindings;
- Threat Model v0.2.0 incorporates accepted store/owner-trust mitigations and residuals;
- Test Plan v0.2.0 incorporates the accepted/evidence-qualified execution class;
- all remain `proposed` pending operator acceptance.

R5 must not promote them by CI or inference.

## 8. Contract lifecycle finding

### R5-F02 — exact first Mission Contract requires operator approval

Status:

```text
OPEN / GATE BLOCKER / OPERATOR DECISION
```

`MIS-M0-SOVEREIGN-CORE-001` v0.1.0 is complete enough for approval but remains non-governing while `status: proposed`.

Approval must bind the exact proposal revision/blob; later material change requires a new Contract revision/supersession.

## 9. Environment and dependency review

The Contract does not mistake R4 spike pins for timeless dependencies.

Evidence-qualified baseline:

```text
Go 1.26.5
modernc.org/sqlite v1.54.0
modernc.org/libc v1.74.1 compatible pin
SQLite 3.53.3 observed
golang.org/x/crypto v0.54.0
CGO=0
Ubuntu 24.04 amd64 primary reference
Windows amd64 storage/trust compatibility evidence
```

R6 must revalidate exact implementation pins. A material semantic/durability/security change triggers Contract replan or renewed evidence.

## 10. Authority boundary

The R5 authorization permits contract preparation/review only.

Still explicitly prohibited:

```text
R6
Microdesign / Implementation Plan
production/source implementation
promotion of spike code
Mastra implementation
AHDK implementation
MNFS integration
deployment
external effects
```

The proposed Contract describes the envelope a later authorized implementation must obey; it does not grant that authority now.

## 11. Exact decision needed

To remove both R5 blockers without broadening scope, the operator must accept the exact R4-aligned A2 package and approve the exact Mission Contract.

Recommended operator decision:

```text
accept Requirements v0.1.1
+ accept Spec v0.2.0
+ accept Threat Model v0.2.0
+ accept Test Plan v0.2.0
+ approve MIS-M0-SOVEREIGN-CORE-001 v0.1.0
```

After that decision:

```text
record exact acceptance blobs
→ rerun R5
→ PASS | FAIL | BLOCKED
→ STOP before R6
```

## 12. Current verdict

```text
R5 BLOCKED
```

Blockers are intentionally narrow and documentary:

1. A2 package explicit operator acceptance;
2. Mission Contract explicit operator approval.

No additional architecture research, spike or Mission decomposition is required by the current evidence.
'''
write(review_path, review)

status = "docs/tracking/STATUS.md"
replace_once(status, "version: 0.22.0", "version: 0.23.0")
replace_once(status, "  - REVIEW-AURORA-M0-R4-MASTRA-MATERIALITY-2026-08-07\nlast_reviewed:", "  - REVIEW-AURORA-M0-R4-MASTRA-MATERIALITY-2026-08-07\n  - DOC-AURORA-M0-R5-OPERATOR-AUTHORIZATION\n  - DOC-AURORA-MIS-M0-SOVEREIGN-CORE-001\n  - DOC-AURORA-CAP-SOVEREIGN-CORE-R5-COVERAGE\n  - REVIEW-AURORA-M0-R5-CONTRACT-READINESS-2026-08-07\nlast_reviewed:")
replace_once(status, "- **Current readiness gate:** ACRM R4 — Architecture/Decision Readiness — PASS; R5 NOT AUTHORIZED", "- **Current readiness gate:** ACRM R5 — Contract Readiness — BLOCKED / OPERATOR APPROVAL REQUIRED; R4 PASS")
replace_once(status, "- **R5 — Contract Readiness:** NOT AUTHORIZED", "- **R5 — Contract Readiness:** AUTHORIZED / PACKAGE PREPARED / BLOCKED ON OPERATOR APPROVAL")
replace_once(status, "- **Runtime implementation:** not started and not authorized", "- **R5 proposed A2 package:** Requirements v0.1.1 + Spec v0.2.0 + Threat Model v0.2.0 + Test Plan v0.2.0 — operator acceptance pending\n- **R5 proposed Mission Contract:** `MIS-M0-SOVEREIGN-CORE-001` v0.1.0 — operator approval pending; 122/122 requirements allocated\n- **R5 verdict:** BLOCKED only on the exact operator approvals above; no new research/spike blocker\n- **Runtime implementation:** not started and not authorized")
replace_once(status, "ACRM R5 — Contract Readiness:    NOT AUTHORIZED", "ACRM R5 — Contract Readiness:    AUTHORIZED / BLOCKED ON OPERATOR APPROVAL")
replace_once(status, "R4 M0 decision coverage:         15/15 decided; all required spikes closed; no unresolved material M0 architecture choice", "R4 M0 decision coverage:         15/15 decided; all required spikes closed; no unresolved material M0 architecture choice\nR5 Mission proposal:              MIS-M0-SOVEREIGN-CORE-001 v0.1.0; 122/122 requirement allocation")
replace_once(status, "Mission Contract:               NOT STARTED", "Mission Contract:               PROPOSED — MIS-M0-SOVEREIGN-CORE-001 v0.1.0; not yet approved")
old_tail = '''## 8. Immediate next action

```text
M0 ACRM R4 PASS
→ STOP
→ await explicit operator authorization for M0 ACRM R5 — Contract Readiness
```'''
new_tail = '''## 8. Immediate next action

```text
R5 package review complete
→ operator reviews exact proposed A2 package + MIS-M0-SOVEREIGN-CORE-001 v0.1.0
→ ACCEPT | REJECT | REVISE
→ if accepted, rerun R5
→ STOP before R6
```

R6, Microdesign and implementation remain NOT AUTHORIZED.'''
replace_once(status, old_tail, new_tail)

worklog = "docs/tracking/WORKLOG.md"
w = read(worklog).rstrip()
entry = '''## 2026-08-07 — M0 R5 Contract Readiness package

The operator authorized M0 ACRM R5 with `Seguir` after canonical R4 PASS at `74167bd1404d9076423ffdbae20f97958283527c`.

R5 prepared one vertical Mission Contract, `MIS-M0-SOVEREIGN-CORE-001` v0.1.0, plus a 122/122 requirement-to-criterion allocation. The R5 preflight also found that R2/R3 gate PASS had intentionally left the normative Requirements/Spec/Threat/Test artifacts `proposed`. Because Blueprint 15 gives accepted A2 specifications precedence over approved A3 contracts, R5 prepared minimal R4-aligned A2 revisions rather than silently allowing a Contract to outrank proposed reusable semantics.

The adversarial R5 review finds the package technically/traceably ready but returns `R5 BLOCKED` pending explicit operator acceptance of Requirements v0.1.1, Spec v0.2.0, Threat Model v0.2.0 and Test Plan v0.2.0 plus approval of `MIS-M0-SOVEREIGN-CORE-001` v0.1.0. No R6 or implementation authority is implied.'''
if "## 2026-08-07 — M0 R5 Contract Readiness package" in w:
    raise SystemExit("WORKLOG R5 entry already present")
write(worklog, w + "\n\n" + entry)

print("R5 alignment/package transformation complete")
print(proposal_blobs)
