from __future__ import annotations

from pathlib import Path
import re
import subprocess

ROOT = Path(__file__).resolve().parents[1]
DATE = "2026-08-09"
BASELINE = "a6769fe8e28dc2dd693f12ad8b9f2460e95b8bc5"
MICRO_PROPOSAL_REV = "0f596602988a90205ff412fdb860e968700dbcb2"
MICRO_PROPOSAL_BLOB = "d76cf237211b7fe35c33d1a32f14905e769702a7"
ACCEPTANCE = "DOC-AURORA-M0-R6-MICRODESIGN-OPERATOR-ACCEPTANCE"
REVIEW = "REVIEW-AURORA-M0-R6-IMPLEMENTATION-DESIGN-READINESS-2026-08-09"


def read(path: str) -> str:
    return (ROOT / path).read_text(encoding="utf-8")


def write(path: str, text: str) -> None:
    p = ROOT / path
    p.parent.mkdir(parents=True, exist_ok=True)
    p.write_text(text.rstrip() + "\n", encoding="utf-8")


def replace_once(text: str, old: str, new: str, path: str) -> str:
    count = text.count(old)
    if count != 1:
        raise SystemExit(f"{path}: expected one match, found {count}: {old!r}")
    return text.replace(old, new, 1)


def git_blob(path: str) -> str:
    return subprocess.check_output(["git", "hash-object", path], cwd=ROOT, text=True).strip()

# 1) Promote the exactly accepted Microdesign without changing design substance.
micro_path = "docs/design/M0-R6-SOVEREIGN-CORE-MICRODESIGN.md"
micro = read(micro_path)
if git_blob(micro_path) != MICRO_PROPOSAL_BLOB:
    raise SystemExit("Microdesign proposal blob no longer matches operator acceptance")
micro = replace_once(micro, "status: proposed\nversion: 0.1.0", f"status: accepted\naccepted_at: {DATE}\nacceptance_evidence: {ACCEPTANCE}\naccepted_from_blob: {MICRO_PROPOSAL_BLOB}\nversion: 0.1.0", micro_path)
micro = replace_once(micro, "  - proposed M0 R6 implementation design for MIS-M0-SOVEREIGN-CORE-001 v0.1.0", "  - accepted M0 R6 implementation design for MIS-M0-SOVEREIGN-CORE-001 v0.1.0", micro_path)
micro = replace_once(micro, "This document remains `proposed` until the operator reviews the written synthesis. It does **not** authorize R7 or production/source implementation.", "This Microdesign was explicitly accepted by the operator through `DOC-AURORA-M0-R6-MICRODESIGN-OPERATOR-ACCEPTANCE`. Acceptance fixes the R6 design but does **not** authorize R7 or production/source implementation.", micro_path)
micro = replace_once(micro, "The document itself remains proposed until the operator reviews this written form.", "The written form was explicitly accepted by the operator. The remaining R6 work is exact task allocation plus adversarial plan review.", micro_path)
write(micro_path, micro)

# 2) Fix plan self-review finding: introduce dependencies only at first consumer.
plan_path = "docs/design/M0-R6-SOVEREIGN-CORE-IMPLEMENTATION-PLAN.md"
plan = read(plan_path)
plan = replace_once(
    plan,
    "- Every task ends with a green test/build result and an independently reviewable commit. No unrelated refactor.\n- Any implementation finding",
    "- Every task ends with a green test/build result and an independently reviewable commit. No unrelated refactor.\n- External modules are introduced just-in-time in the first TASK that imports them; `go mod tidy` must not be used to pre-stage unused framework/dependency surface.\n- Any implementation finding",
    plan_path,
)
plan = replace_once(
    plan,
    "- [ ] **Step 3 — pin baseline:** initialize `go.mod` with Go `1.26.5`; pin current R6-validated dependencies. Run `go mod tidy`. Verify `go list -m all` contains `modernc.org/sqlite v1.54.0` and `modernc.org/libc v1.74.1`; fail if modernc resolves a different libc pin.\n- [ ] **Step 4 — minimal GREEN:**",
    "- [ ] **Step 3 — module baseline:** initialize `go.mod` with module `github.com/developmentconexus-ops/aurora_project` and Go `1.26.5`. Do not pre-stage unused external modules in TASK-00.\n- [ ] **Step 4 — minimal GREEN:**",
    plan_path,
)
plan = replace_once(
    plan,
    "- [ ] **Step 1 — RED domain/crypto:** tests assert IDs are random stable prefixed IDs; root creation produces a 32-byte ORK only in memory; wrong passphrase fails; unsupported/extreme KDF params fail before Argon2 call; rewrapping preserves ORK bytes.\n- [ ] **Step 2 — RED store:**",
    "- [ ] **Step 1 — dependency lock for this slice:** add `modernc.org/sqlite v1.54.0`, exact compatible `modernc.org/libc v1.74.1`, `golang.org/x/crypto v0.54.0`, `github.com/gowebpki/jcs v1.0.1`, and the then-current supported exact `golang.org/x/sys` version required by the Windows publication adapter. Run `go mod tidy` only after imports exist; verify modernc/libc exact compatibility.\n- [ ] **Step 2 — RED domain/crypto:** tests assert IDs are random stable prefixed IDs; root creation produces a 32-byte ORK only in memory; wrong passphrase fails; unsupported/extreme KDF params fail before Argon2 call; rewrapping preserves ORK bytes.\n- [ ] **Step 3 — RED store:**",
    plan_path,
)
# Renumber remaining TASK-01 textual steps only for clarity.
for old, new in [
    ("- [ ] **Step 3 — RED trustfs:**", "- [ ] **Step 4 — RED trustfs:**"),
    ("- [ ] **Step 4 — implement crypto:**", "- [ ] **Step 5 — implement crypto:**"),
    ("- [ ] **Step 5 — implement bootstrap ordering:**", "- [ ] **Step 6 — implement bootstrap ordering:**"),
    ("- [ ] **Step 6 — CLI:**", "- [ ] **Step 7 — CLI:**"),
    ("- [ ] **Step 7 — verify journey:**", "- [ ] **Step 8 — verify journey:**"),
    ("- [ ] **Step 8 — verify hygiene:**", "- [ ] **Step 9 — verify hygiene:**"),
    ("- [ ] **Step 9 — commit:**", "- [ ] **Step 10 — commit:**"),
]:
    plan = replace_once(plan, old, new, plan_path)
plan = replace_once(
    plan,
    "- [ ] RED: schema/unit tests define `StateEnvelope{SchemaVersion,Kind,Summary,Payload}` and prove payload content resembling identity/authority remains opaque data.\n- [ ] RED: integration test",
    "- [ ] Add and pin `github.com/santhosh-tekuri/jsonschema/v6 v6.0.3`, then run `go mod tidy`; this is its first production consumer.\n- [ ] RED: schema/unit tests define `StateEnvelope{SchemaVersion,Kind,Summary,Payload}` and prove payload content resembling identity/authority remains opaque data.\n- [ ] RED: integration test",
    plan_path,
)
plan = replace_once(
    plan,
    "- [ ] RED schema/digest tests: valid export validates; a byte/semantic alteration fails digest/schema before apply; unsupported version fails explicitly.\n- [ ] Implement `ExportProtection`",
    "- [ ] Add and pin `filippo.io/age v1.3.1`, then run `go mod tidy`; this is its first production consumer.\n- [ ] RED schema/digest tests: valid export validates; a byte/semantic alteration fails digest/schema before apply; unsupported version fails explicitly.\n- [ ] Implement `ExportProtection`",
    plan_path,
)
plan = replace_once(
    plan,
    "- [ ] RED: accepted/rejected operations produce attributable records with stable `operation_id`, Aurora/Project/revision refs, outcome/reason; EvidenceRecord includes criterion/test method/environment/revision/limitations fields.\n- [ ] RED: telemetry sink",
    "- [ ] Add and pin OpenTelemetry Go API/SDK `v1.44.0`, then run `go mod tidy`; no exporter/backend dependency is added unless a current test needs one.\n- [ ] RED: accepted/rejected operations produce attributable records with stable `operation_id`, Aurora/Project/revision refs, outcome/reason; EvidenceRecord includes criterion/test method/environment/revision/limitations fields.\n- [ ] RED: telemetry sink",
    plan_path,
)
write(plan_path, plan)
plan_blob = git_blob(plan_path)

# 3) Generate exact 122/122 requirement -> primary task allocation.
def task_for_req(n: int) -> str:
    if 1 <= n <= 3: return "TASK-01"
    if 4 <= n <= 9: return "TASK-02"
    if 10 <= n <= 20: return "TASK-03"
    if 21 <= n <= 30: return "TASK-04"
    if n == 31: return "TASK-12"
    if 32 <= n <= 45: return "TASK-05"
    if 46 <= n <= 55: return "TASK-07"
    if 56 <= n <= 63: return "TASK-08"
    if 64 <= n <= 66: return "TASK-09"
    if 67 <= n <= 76: return "TASK-06"
    if 77 <= n <= 88: return "TASK-10"
    if 89 <= n <= 95: return "TASK-11"
    if 96 <= n <= 107: return "TASK-00"
    if 108 <= n <= 122: return "TASK-13"
    raise AssertionError(n)
rows = "\n".join(f"| `CAP-SOVEREIGN-CORE-REQ-{i:03d}` | `{task_for_req(i)}` |" for i in range(1, 123))
coverage = f'''---
id: DOC-AURORA-CAP-SOVEREIGN-CORE-R6-IMPLEMENTATION-COVERAGE
title: CAP-SOVEREIGN-CORE R6 Requirement-to-Implementation Allocation
document_type: implementation_coverage
form: reference
authority: design
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - M0 R6 primary requirement-to-implementation-task allocation
related:
  - DESIGN-AURORA-M0-R6-SOVEREIGN-CORE-MICRODESIGN
  - DESIGN-AURORA-M0-R6-SOVEREIGN-CORE-IMPLEMENTATION-PLAN
  - DOC-AURORA-CAP-SOVEREIGN-CORE-REQUIREMENTS
  - DOC-AURORA-CAP-SOVEREIGN-CORE-R5-COVERAGE
source_revision: {BASELINE}
last_reviewed: {DATE}
---

# CAP-SOVEREIGN-CORE — R6 Implementation Coverage

Every accepted `CAP-SOVEREIGN-CORE-REQ-001..122` has exactly one primary R7 implementation task. Secondary tasks/tests remain governed by the accepted Test Plan and the task text; this table prevents requirements from disappearing between Contract and code.

| Requirement | Primary implementation task |
|---|---|
{rows}

## Mission criterion task coverage

| Criterion | Primary tasks |
|---|---|
| `CRIT-001` | `TASK-01`, `TASK-02` |
| `CRIT-002` | `TASK-03` |
| `CRIT-003` | `TASK-04` |
| `CRIT-004` | `TASK-05`, `TASK-06` |
| `CRIT-005` | `TASK-07` |
| `CRIT-006` | `TASK-08`, `TASK-09` |
| `CRIT-007` | `TASK-01`, `TASK-06`, `TASK-08`, `TASK-11` |
| `CRIT-008` | `TASK-10` |
| `CRIT-009` | `TASK-06`, `TASK-07`, `TASK-11` |
| `CRIT-010` | `TASK-00` + architecture guard on every task |
| `CRIT-011` | `TASK-13` |
| `CRIT-012` | `TASK-12` |

`TASK-12` composes all prior product behavior into the complete real-binary Golden Proof. `TASK-13` packages traceability/evidence but cannot convert an implementation Claim into a Verdict.
'''
write("docs/capabilities/CAP-SOVEREIGN-CORE/R6-IMPLEMENTATION-COVERAGE.md", coverage)

# 4) Adversarial R6 review.
review = f'''---
id: {REVIEW}
title: M0 ACRM R6 Implementation Design Readiness Review
document_type: adversarial_review
form: explanation
authority: evidence
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - final M0 R6 Implementation Design Readiness observations and verdict
related:
  - DOC-AURORA-STATUS
  - DOC-AURORA-M0-R6-OPERATOR-AUTHORIZATION
  - {ACCEPTANCE}
  - DESIGN-AURORA-M0-R6-SOVEREIGN-CORE-MICRODESIGN
  - DESIGN-AURORA-M0-R6-SOVEREIGN-CORE-IMPLEMENTATION-PLAN
  - DOC-AURORA-CAP-SOVEREIGN-CORE-R6-IMPLEMENTATION-COVERAGE
  - DOC-AURORA-MIS-M0-SOVEREIGN-CORE-001
source_revision: {BASELINE}
reviewed_at: {DATE}
last_reviewed: {DATE}
---

# M0 ACRM R6 — Implementation Design Readiness Review

## 1. Executive verdict

```text
R6 PASS
```

The accepted Microdesign plus reviewed Implementation Plan are concrete enough for an R7 implementer to execute the approved Mission without inventing material product behavior or architecture. This verdict does **not** authorize R7 or source implementation.

## 2. Exact reviewed package

```text
canonical R5/R6 source baseline: {BASELINE}
accepted Microdesign proposal revision: {MICRO_PROPOSAL_REV}
accepted Microdesign proposal blob: {MICRO_PROPOSAL_BLOB}
Microdesign: v0.1.0 / accepted
Implementation Plan: v0.1.0 / blob {plan_blob}
Requirement-to-task allocation: 122/122
Mission criteria with task coverage: 12/12
```

## 3. Adversarial findings

### R6-F01 — heads-only governing HMAC would not authenticate current governing content — RESOLVED

The conversational draft initially proposed authenticating only governing revision heads. Self-review correctly found that an attacker could alter the current revision contents without changing its revision number. The written Microdesign was corrected before operator acceptance to HMAC one JCS-canonical **complete current governing logical snapshot** while still avoiding per-row HMACs, Merkle trees or a custom transaction protocol.

### R6-F02 — pre-staging all Go dependencies in TASK-00 conflicts with `go mod tidy` and vertical-slice YAGNI — RESOLVED

Initial task text attempted to pin every future module in TASK-00. Review corrected this to just-in-time introduction at first production consumer: SQLite/x-crypto/JCS in TASK-01, JSON Schema in TASK-03, age in TASK-08 and OTel in TASK-10. This preserves reproducibility without foundation-only dependency surface.

## 4. Dependency revalidation

Current primary/repository sources checked during R6 support the starting pins used by the Plan:

```text
Go                         1.26.5
modernc.org/sqlite         v1.54.0
modernc.org/libc           v1.74.1 exact compatible pin
golang.org/x/crypto        v0.54.0
filippo.io/age             v1.3.1
santhosh-tekuri/jsonschema v6.0.3
gowebpki/jcs               v1.0.1
OpenTelemetry Go           v1.44.0
```

`golang.org/x/sys` remains deliberately task-local: TASK-01 must pin the then-current supported exact version when the Windows atomic-replace code is introduced. That choice is reversible adapter plumbing and does not own product/domain semantics.

## 5. R6 gate checklist

| Condition | Result |
|---|---|
| accepted R5 Mission and A2/ADRs govern the plan | PASS |
| exact production source tree/modules defined | PASS |
| interfaces/types/schema/transitions concrete | PASS |
| SQLite/trust physical design implementable without inventing protocol | PASS |
| bounded Argon2 decode and OS publication semantics explicit | PASS |
| export/restore/migration exact enough to implement | PASS |
| test-first vertical sequence produces usable Aurora early | PASS |
| fault/security hooks and real-process proof explicit | PASS |
| observability optional/non-authoritative | PASS |
| all 122 requirements have one primary implementation task | PASS — 122/122 |
| all 12 Mission criteria have task coverage | PASS — 12/12 |
| no material design decision delegated to implementer | PASS |
| no hidden M1+/Mastra/AHDK/MNFS scope | PASS |
| no source/runtime implementation performed in R6 | PASS |
| R7 authority absent | PASS |

## 6. Complexity review

The design deliberately retains low-cost, high-value structural seams (`domain/application/ports/adapters`) while refusing speculative subsystems. Physical persistence remains six SQLite tables, one DB, two tiny trust files and one governing HMAC. No generic repository, policy engine, migration framework, Presence framework or agent framework is built for M0.

## 7. Carry-forward R7 obligations

R7 must:

- follow TASK-00..13 order unless a documented non-material parallelism rule applies;
- use RED→GREEN per task and preserve negative tests;
- never copy spike code into production;
- revalidate exact module locks when each dependency is first introduced;
- preserve the narrower local threat/fault claim and avoid power-loss overclaim;
- stop/replan on any finding that changes accepted behavior, mechanism class or topology;
- produce fixed-revision evidence before success claims.

## 8. Stop boundary

```text
M0 ACRM R6 — PASS
→ STOP
→ R7 NOT AUTHORIZED
→ await explicit operator authorization for M0 ACRM R7 — Execution and Evidence
```
'''
write("docs/reviews/2026-08-09-m0-r6-implementation-design-readiness-review.md", review)

# 5) Update STATUS to truthful R6 PASS / R7 not authorized.
status_path = "docs/tracking/STATUS.md"
status = read(status_path)
status = replace_once(status, "version: 0.24.0", "version: 0.25.0", status_path)
status = replace_once(status, "  - REVIEW-AURORA-M0-R5-CONTRACT-READINESS-RERUN-2026-08-09\nlast_reviewed:", f"  - REVIEW-AURORA-M0-R5-CONTRACT-READINESS-RERUN-2026-08-09\n  - DOC-AURORA-M0-R6-OPERATOR-AUTHORIZATION\n  - {ACCEPTANCE}\n  - DESIGN-AURORA-M0-R6-SOVEREIGN-CORE-MICRODESIGN\n  - DESIGN-AURORA-M0-R6-SOVEREIGN-CORE-IMPLEMENTATION-PLAN\n  - DOC-AURORA-CAP-SOVEREIGN-CORE-R6-IMPLEMENTATION-COVERAGE\n  - {REVIEW}\nlast_reviewed:", status_path)
status = replace_once(status, "- **Current readiness gate:** ACRM R5 — Contract Readiness — PASS; R6 NOT AUTHORIZED", "- **Current readiness gate:** ACRM R6 — Implementation Design Readiness — PASS; R7 NOT AUTHORIZED", status_path)
status = replace_once(status, "- **R6 and later gates:** NOT AUTHORIZED BY IMPLICATION", "- **R6 — Implementation Design Readiness:** PASS — Microdesign v0.1.0 accepted; Implementation Plan v0.1.0 reviewed; 122/122 requirements allocated\n- **R7 and later gates:** NOT AUTHORIZED BY IMPLICATION", status_path)
status = replace_once(status, "- **Runtime implementation:** not started and not authorized", "- **R6 implementation allocation:** 122/122 accepted requirements → TASK-00..13; 12/12 Mission criteria covered\n- **Runtime implementation:** not started and not authorized", status_path)
status = replace_once(status, "ACRM R6+:                       NOT AUTHORIZED", "ACRM R6 — Implementation Design Readiness: PASS — accepted Microdesign + reviewed exact Implementation Plan\nACRM R7+:                       NOT AUTHORIZED", status_path)
status = replace_once(status, "Microdesign/Implementation Plan: NOT STARTED", "Microdesign:                    ACCEPTED — M0 R6 Sovereign Core v0.1.0\nImplementation Plan:            REVIEWED — M0 R6 Sovereign Core v0.1.0; TASK-00..13", status_path)
start = status.find("## 7. Current blocker/gate")
if start < 0:
    raise SystemExit("STATUS section 7 missing")
status = status[:start] + f'''## 7. Current blocker/gate

M0 ACRM R6 is complete. The operator accepted `DESIGN-AURORA-M0-R6-SOVEREIGN-CORE-MICRODESIGN` v0.1.0 bound to proposal blob `{MICRO_PROPOSAL_BLOB}`. The derived Implementation Plan v0.1.0 was self-reviewed and adversarially reviewed with 122/122 requirement-to-task allocation and 12/12 Mission criterion coverage.

Final verdict:

```text
M0 ACRM R6 — PASS
```

There is no current R6 blocker. No production Go source was created or executed during R6. R7 remains a separate explicit authorization boundary.

## 8. Immediate next action

```text
R6 PASS
→ STOP
→ await explicit operator authorization for M0 ACRM R7 — Execution and Evidence
```

Aurora Core implementation, production source creation/execution, Mastra/AHDK/MNFS implementation and all later gates remain NOT AUTHORIZED.
'''
write(status_path, status)

# 6) Append WORKLOG.
worklog_path = "docs/tracking/WORKLOG.md"
worklog = read(worklog_path).rstrip()
entry = f'''\n\n## {DATE} — M0 R6 Microdesign Acceptance, Implementation Plan and PASS\n\nThe operator explicitly accepted M0 R6 Sovereign Core Microdesign v0.1.0, bound to proposal blob `{MICRO_PROPOSAL_BLOB}` at branch revision `{MICRO_PROPOSAL_REV}`.\n\nR6 then produced `M0-R6-SOVEREIGN-CORE-IMPLEMENTATION-PLAN.md` v0.1.0 with TASK-00..13 and exact file/interface/test/verification/commit boundaries. `R6-IMPLEMENTATION-COVERAGE.md` assigns all 122 accepted requirements to one primary implementation task and covers all 12 Mission criteria.\n\nAdversarial review resolved two material planning findings before verdict: the governing HMAC was corrected to cover the complete current governing logical snapshot rather than only revision heads, and dependency pins were moved from foundation-only pre-staging to just-in-time first consumers.\n\nFinal verdict is `M0 ACRM R6 PASS`. R7 Execution and Evidence, production Go source, deployment, Mastra/AHDK/MNFS implementation and later gates remain unauthorized.\n'''
write(worklog_path, worklog + entry)

print(f"plan_blob={plan_blob}")
print("R6 documentary closeout prepared")
