from __future__ import annotations

from pathlib import Path
import hashlib
import re
import subprocess

BASELINE = "abbcb063c90c834ad45f6b04ca5abe308f9dacb2"
DATE = "2026-08-09"
ACCEPTANCE = "DOC-AURORA-M0-R5-A2-CONTRACT-OPERATOR-ACCEPTANCE"
RERUN_REVIEW = "REVIEW-AURORA-M0-R5-CONTRACT-READINESS-RERUN-2026-08-09"

TARGETS = {
    "docs/capabilities/CAP-SOVEREIGN-CORE/REQUIREMENTS.md": ("0.1.1", "de234e4a57c04d1d0b68cd017597e06a618fd68b"),
    "docs/capabilities/CAP-SOVEREIGN-CORE/SPEC.md": ("0.2.0", "dd6f66c23c08fc635d780aac5e70533a82e72a75"),
    "docs/capabilities/CAP-SOVEREIGN-CORE/THREAT-MODEL.md": ("0.2.0", "7e97f816d0c4966ba6b12cf0447c7a2210fbea34"),
    "docs/capabilities/CAP-SOVEREIGN-CORE/TEST-PLAN.md": ("0.2.0", "8b42cc451439038e63e8b567702877b8951c5edb"),
    "docs/capabilities/CAP-SOVEREIGN-CORE/MIS-M0-SOVEREIGN-CORE-001.md": ("0.1.0", "1db39012874828f54f293bf76571259494ba5a79"),
}


def read(path: str) -> str:
    return Path(path).read_text(encoding="utf-8")


def write(path: str, text: str) -> None:
    Path(path).write_text(text.rstrip() + "\n", encoding="utf-8")


def git_blob(path: str) -> str:
    return subprocess.check_output(["git", "hash-object", path], text=True).strip()


def replace_once(text: str, old: str, new: str, label: str) -> str:
    count = text.count(old)
    if count != 1:
        raise SystemExit(f"{label}: expected one match, found {count}: {old[:100]!r}")
    return text.replace(old, new, 1)


def split_frontmatter(text: str) -> tuple[list[str], str]:
    if not text.startswith("---\n"):
        raise SystemExit("missing frontmatter")
    end = text.find("\n---\n", 4)
    if end < 0:
        raise SystemExit("unterminated frontmatter")
    return text[4:end].splitlines(), text[end + 5:].lstrip("\n")


def join_frontmatter(lines: list[str], body: str) -> str:
    return "---\n" + "\n".join(lines) + "\n---\n\n" + body.rstrip() + "\n"


def set_scalar(lines: list[str], key: str, value: str) -> None:
    prefix = key + ":"
    hits = [i for i, line in enumerate(lines) if line.startswith(prefix)]
    if len(hits) != 1:
        raise SystemExit(f"frontmatter key {key} count={len(hits)}")
    lines[hits[0]] = f"{key}: {value}"


def insert_after_key(lines: list[str], key: str, additions: list[str]) -> None:
    prefix = key + ":"
    hits = [i for i, line in enumerate(lines) if line.startswith(prefix)]
    if len(hits) != 1:
        raise SystemExit(f"frontmatter key {key} count={len(hits)}")
    idx = hits[0] + 1
    for line in reversed(additions):
        lines.insert(idx, line)


def append_related(lines: list[str], doc_id: str) -> None:
    if f"  - {doc_id}" in lines:
        return
    try:
        idx = lines.index("related:") + 1
    except ValueError as exc:
        raise SystemExit("related list missing") from exc
    while idx < len(lines) and lines[idx].startswith("  - "):
        idx += 1
    lines.insert(idx, f"  - {doc_id}")


def ensure_approver(lines: list[str]) -> None:
    if "approvers:" in lines:
        return
    try:
        idx = lines.index("owners:") + 1
    except ValueError as exc:
        raise SystemExit("owners list missing") from exc
    while idx < len(lines) and lines[idx].startswith("  - "):
        idx += 1
    lines[idx:idx] = ["approvers:", "  - operator"]


def promote(path: str, version: str, expected_blob: str, body_replacements: list[tuple[str, str]]) -> None:
    actual = git_blob(path)
    if actual != expected_blob:
        raise SystemExit(f"{path}: accepted blob mismatch {actual} != {expected_blob}")
    text = read(path)
    fm, body = split_frontmatter(text)
    if f"version: {version}" not in fm:
        raise SystemExit(f"{path}: version mismatch")
    if "status: proposed" not in fm:
        raise SystemExit(f"{path}: expected status proposed")
    set_scalar(fm, "status", "accepted")
    status_idx = fm.index("status: accepted")
    fm[status_idx + 1:status_idx + 1] = [
        f"accepted_at: {DATE}",
        f"acceptance_evidence: {ACCEPTANCE}",
        f"accepted_from_blob: {expected_blob}",
    ]
    append_related(fm, ACCEPTANCE)
    ensure_approver(fm)
    for old, new in body_replacements:
        body = replace_once(body, old, new, path)
    write(path, join_frontmatter(fm, body))


# Snapshot semantic content that lifecycle promotion is forbidden to change.
req_before = read("docs/capabilities/CAP-SOVEREIGN-CORE/REQUIREMENTS.md")
req_rows_before = re.findall(r"^\| `CAP-SOVEREIGN-CORE-REQ-\d{3}` \|.*$", req_before, re.M)
if len(req_rows_before) != 122:
    raise SystemExit(f"expected 122 requirement rows before promotion, got {len(req_rows_before)}")
contract_before = read("docs/capabilities/CAP-SOVEREIGN-CORE/MIS-M0-SOVEREIGN-CORE-001.md")
criteria_before = re.findall(
    r"^### `MIS-M0-SOVEREIGN-CORE-001-CRIT-\d{3}`.*?(?=^### `MIS-M0-SOVEREIGN-CORE-001-CRIT-|^## )",
    contract_before,
    re.M | re.S,
)
if len(criteria_before) != 12:
    raise SystemExit(f"expected 12 Mission criteria before promotion, got {len(criteria_before)}")

promote(
    "docs/capabilities/CAP-SOVEREIGN-CORE/REQUIREMENTS.md",
    *TARGETS["docs/capabilities/CAP-SOVEREIGN-CORE/REQUIREMENTS.md"],
    body_replacements=[
        (
            "this document remains `proposed` until explicitly accepted by the operator.",
            "this document was explicitly accepted by the operator through `DOC-AURORA-M0-R5-A2-CONTRACT-OPERATOR-ACCEPTANCE`.",
        ),
    ],
)

promote(
    "docs/capabilities/CAP-SOVEREIGN-CORE/SPEC.md",
    *TARGETS["docs/capabilities/CAP-SOVEREIGN-CORE/SPEC.md"],
    body_replacements=[
        (
            "This document is the proposed reusable Capability/System Specification for:",
            "This document is the accepted reusable Capability/System Specification for:",
        ),
        (
            "Current promotion state is `G1`; the R5 package is proposed and cannot advance to G2 without the required operator decisions.",
            "The R4-aligned A2 package and `MIS-M0-SOVEREIGN-CORE-001` v0.1.0 were explicitly accepted/approved by the operator. The R5 rerun determines Contract Readiness PASS; R6 remains separately gated.",
        ),
        (
            "This Spec is the proposed R4-aligned A2 reusable-behavior authority for CAP-SOVEREIGN-CORE.",
            "This Spec is the accepted R4-aligned A2 reusable-behavior authority for CAP-SOVEREIGN-CORE.",
        ),
        (
            "R4 PASS\n→ R5 package/operator decision\n→ STOP before R6 unless separately authorized",
            "R4 PASS\n→ A2 package + Mission Contract accepted/approved\n→ R5 rerun\n→ STOP before R6 unless separately authorized",
        ),
        (
            "R6 may define implementation detail only after the A2 package and Mission Contract are accepted and R5 passes.",
            "R6 may define implementation detail only after R5 passes and R6 is separately authorized.",
        ),
    ],
)

promote(
    "docs/capabilities/CAP-SOVEREIGN-CORE/THREAT-MODEL.md",
    *TARGETS["docs/capabilities/CAP-SOVEREIGN-CORE/THREAT-MODEL.md"],
    body_replacements=[],
)

promote(
    "docs/capabilities/CAP-SOVEREIGN-CORE/TEST-PLAN.md",
    *TARGETS["docs/capabilities/CAP-SOVEREIGN-CORE/TEST-PLAN.md"],
    body_replacements=[],
)

promote(
    "docs/capabilities/CAP-SOVEREIGN-CORE/MIS-M0-SOVEREIGN-CORE-001.md",
    *TARGETS["docs/capabilities/CAP-SOVEREIGN-CORE/MIS-M0-SOVEREIGN-CORE-001.md"],
    body_replacements=[
        (
            "  - proposed exact implementation commitment for M0 Sovereign Core Walking Skeleton",
            "  - exact scoped implementation commitment for M0 Sovereign Core Walking Skeleton",
        ),
        ("Status:           PROPOSED", "Status:           APPROVED"),
        (
            "The Contract becomes governing only after explicit operator approval. Approval of this Contract does not itself authorize R6, production implementation or merge/promotion.",
            "The Contract was explicitly approved by the operator through `DOC-AURORA-M0-R5-A2-CONTRACT-OPERATOR-ACCEPTANCE` and is the governing scoped M0 implementation commitment. Approval does not itself authorize R6, production implementation or merge/promotion.",
        ),
    ],
)

# Verify normative/criterion semantics survived the lifecycle-only promotion.
req_after = read("docs/capabilities/CAP-SOVEREIGN-CORE/REQUIREMENTS.md")
req_rows_after = re.findall(r"^\| `CAP-SOVEREIGN-CORE-REQ-\d{3}` \|.*$", req_after, re.M)
if req_rows_after != req_rows_before:
    raise SystemExit("requirement statements changed during lifecycle promotion")
contract_after = read("docs/capabilities/CAP-SOVEREIGN-CORE/MIS-M0-SOVEREIGN-CORE-001.md")
criteria_after = re.findall(
    r"^### `MIS-M0-SOVEREIGN-CORE-001-CRIT-\d{3}`.*?(?=^### `MIS-M0-SOVEREIGN-CORE-001-CRIT-|^## )",
    contract_after,
    re.M | re.S,
)
if criteria_after != criteria_before:
    raise SystemExit("Mission criteria changed during lifecycle promotion")

# Create final R5 rerun review.
review_path = Path("docs/reviews/2026-08-09-m0-r5-contract-readiness-rerun.md")
review_path.write_text(f'''---
id: {RERUN_REVIEW}
title: M0 ACRM R5 Contract Readiness Final Rerun
document_type: adversarial_review
form: explanation
authority: evidence
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - final M0 R5 Contract Readiness observations and verdict after exact operator acceptance
related:
  - DOC-AURORA-STATUS
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - {ACCEPTANCE}
  - DOC-AURORA-CAP-SOVEREIGN-CORE-REQUIREMENTS
  - DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC
  - DOC-AURORA-CAP-SOVEREIGN-CORE-THREAT-MODEL
  - DOC-AURORA-CAP-SOVEREIGN-CORE-TEST-PLAN
  - DOC-AURORA-CAP-SOVEREIGN-CORE-R5-COVERAGE
  - DOC-AURORA-MIS-M0-SOVEREIGN-CORE-001
  - REVIEW-AURORA-M0-R5-CONTRACT-READINESS-2026-08-07
source_revision: {BASELINE}
reviewed_at: {DATE}
last_reviewed: {DATE}
---

# M0 ACRM R5 — Contract Readiness Final Rerun

## 1. Executive verdict

```text
R5 PASS
```

The two blockers from the initial R5 review are resolved by explicit operator authority bound to the exact reviewed proposal blobs. No new architecture, scope, security or traceability blocker was introduced during lifecycle promotion.

This verdict does **not** authorize R6, Microdesign or implementation.

## 2. Exact authority binding

Accepted proposal baseline:

```text
{BASELINE}
```

Accepted/approved proposal blobs:

```text
Requirements v0.1.1  de234e4a57c04d1d0b68cd017597e06a618fd68b
Spec v0.2.0          dd6f66c23c08fc635d780aac5e70533a82e72a75
Threat Model v0.2.0  7e97f816d0c4966ba6b12cf0447c7a2210fbea34
Test Plan v0.2.0     8b42cc451439038e63e8b567702877b8951c5edb
Mission v0.1.0       1db39012874828f54f293bf76571259494ba5a79
```

Authority evidence: `{ACCEPTANCE}`.

Lifecycle promotion changes only acceptance metadata and stale lifecycle wording. The 122 requirement statements and 12 Mission criteria were mechanically compared before/after promotion and remained byte-identical as extracted normative rows/criterion sections.

## 3. Gate checklist

| R5 condition | Result |
|---|---|
| R4 complete for Mission scope | PASS |
| exact Mission identity/revision/baseline | PASS |
| operator-visible outcome + Golden Proof contribution | PASS |
| scope/non-goals/assumptions/dependencies explicit | PASS |
| contract-level decomposition without R6 Microdesign | PASS |
| in-scope requirement allocation | PASS — 122/122 |
| authority/prohibitions/budget explicit | PASS |
| evidence profile/thresholds explicit | PASS |
| change/replan/supersession path explicit | PASS |
| hidden M1+/Mastra/AHDK/MNFS scope absent | PASS |
| accepted A2 normative package | PASS — exact operator acceptance |
| approved Mission Contract | PASS — exact operator approval |
| implementation authorization absent | PASS |

## 4. Contract readiness conclusion

`MIS-M0-SOVEREIGN-CORE-001` v0.1.0 is now an approved A3 scoped commitment over an accepted A2 reusable specification package. The Contract carries all 122 requirements through 12 primary criteria and preserves the accepted R4 architecture boundaries.

R6 may later define files, Go interfaces/types, SQL DDL, filesystem publication, CLI/proof adapter, test files and implementation sequence, but it may not alter accepted behavior/architecture by implementation convenience.

## 5. Residual carry-forward obligations

R6/R7 must preserve, among other accepted constraints:

- bounded Argon2id envelope parsing before memory allocation;
- target-filesystem publication/fsync/directory-sync design and verification for claimed durability;
- mutation-boundary enforcement for rollback/anchor-lag/time/restore anomalies;
- no overclaim beyond the local fault/threat model evidenced by SPK-001/SPK-002;
- exact dependency/version revalidation before implementation;
- full deterministic negative/fault/security paths from the accepted Test Plan;
- no promotion of disposable spike code as production code.

These are R6/R7 design/evidence obligations, not R5 blockers.

## 6. Stop boundary

```text
M0 ACRM R5 — PASS
→ STOP
→ R6 NOT AUTHORIZED
→ await explicit operator authorization for M0 ACRM R6 — Implementation Design Readiness
```

No source/runtime implementation is authorized by this verdict.
''', encoding="utf-8")

# Update STATUS.
status_path = "docs/tracking/STATUS.md"
status = read(status_path)
status = replace_once(status, "version: 0.23.0", "version: 0.24.0", status_path)
status = replace_once(status, "last_reviewed: 2026-08-07", f"last_reviewed: {DATE}", status_path)
status = replace_once(status, "  - REVIEW-AURORA-M0-R5-CONTRACT-READINESS-2026-08-07\nlast_reviewed:", f"  - REVIEW-AURORA-M0-R5-CONTRACT-READINESS-2026-08-07\n  - {ACCEPTANCE}\n  - {RERUN_REVIEW}\nlast_reviewed:", status_path)
status = replace_once(status, "- **Current readiness gate:** ACRM R5 — Contract Readiness — BLOCKED / OPERATOR APPROVAL REQUIRED; R4 PASS", "- **Current readiness gate:** ACRM R5 — Contract Readiness — PASS; R6 NOT AUTHORIZED", status_path)
status = replace_once(status, "- **R2 derived requirements:** 122 proposed atomic requirements; coverage 127/127 active sources", "- **R2 derived requirements:** 122 accepted atomic requirements; coverage 127/127 active sources", status_path)
status = replace_once(status, "- **R5 — Contract Readiness:** AUTHORIZED / PACKAGE PREPARED / BLOCKED ON OPERATOR APPROVAL", "- **R5 — Contract Readiness:** PASS — exact A2 package accepted and MIS-M0-SOVEREIGN-CORE-001 v0.1.0 approved", status_path)
status = replace_once(status, "- **R5 proposed A2 package:** Requirements v0.1.1 + Spec v0.2.0 + Threat Model v0.2.0 + Test Plan v0.2.0 — operator acceptance pending", "- **R5 accepted A2 package:** Requirements v0.1.1 + Spec v0.2.0 + Threat Model v0.2.0 + Test Plan v0.2.0 — accepted from exact proposal blobs", status_path)
status = replace_once(status, "- **R5 proposed Mission Contract:** `MIS-M0-SOVEREIGN-CORE-001` v0.1.0 — operator approval pending; 122/122 requirements allocated", "- **R5 approved Mission Contract:** `MIS-M0-SOVEREIGN-CORE-001` v0.1.0 — approved; 122/122 requirements allocated", status_path)
status = replace_once(status, "- **R5 verdict:** BLOCKED only on the exact operator approvals above; no new research/spike blocker", "- **R5 verdict:** PASS — final rerun after exact operator acceptance; no unresolved Contract Readiness blocker", status_path)
status = replace_once(status, "R2 requirements baseline:       122 proposed atomic requirements; coverage 127/127", "R2 requirements baseline:       122 accepted atomic requirements; coverage 127/127", status_path)
status = replace_once(status, "R3 proposed Capability package: Spec + threat model + test plan + 122/122 allocation", "R3 Capability package:          ACCEPTED — R4-aligned Requirements/Spec/Threat/Test + 122/122 allocation", status_path)
status = replace_once(status, "R5 Mission proposal:              MIS-M0-SOVEREIGN-CORE-001 v0.1.0; 122/122 requirement allocation", "R5 Mission Contract:             APPROVED — MIS-M0-SOVEREIGN-CORE-001 v0.1.0; 122/122 requirement allocation", status_path)
status = replace_once(status, "ACRM R5 — Contract Readiness:    AUTHORIZED / BLOCKED ON OPERATOR APPROVAL", "ACRM R5 — Contract Readiness:    PASS — exact A2 package + Mission Contract accepted/approved", status_path)
status = replace_once(status, "Mission Contract:               PROPOSED — MIS-M0-SOVEREIGN-CORE-001 v0.1.0; not yet approved", "Mission Contract:               APPROVED — MIS-M0-SOVEREIGN-CORE-001 v0.1.0", status_path)

start = status.find("## 7. Current blocker/gate")
end = status.find("## 8. Immediate next action", start)
if start < 0 or end < 0:
    raise SystemExit("STATUS sections 7/8 not found")
new7 = f'''## 7. Current blocker/gate

The operator explicitly accepted the R4-aligned A2 package and approved `MIS-M0-SOVEREIGN-CORE-001` v0.1.0 through `{ACCEPTANCE}`, bound to the exact proposal blobs at baseline `{BASELINE}`.

The final R5 rerun verified:

```text
accepted Requirements/Spec/Threat/Test package
+ approved Mission Contract
+ 122/122 requirements allocated to 12 Mission criteria
+ accepted R4 architecture carried forward
+ no hidden M1+/Mastra/AHDK/MNFS scope
+ no R6 Microdesign performed
+ implementation remains prohibited
```

Final verdict:

```text
M0 ACRM R5 — PASS
```

There is no current R5 blocker. R6 remains a separate authorization boundary.

'''
status = status[:start] + new7 + status[end:]
start8 = status.find("## 8. Immediate next action")
if start8 < 0:
    raise SystemExit("STATUS section 8 missing")
status = status[:start8] + '''## 8. Immediate next action

```text
R5 PASS
→ STOP
→ await explicit operator authorization for M0 ACRM R6 — Implementation Design Readiness
```

R6, Microdesign and Aurora Core implementation remain NOT AUTHORIZED.
'''
write(status_path, status)

# Update decision index and close O-015.
dec_path = "docs/tracking/DECISIONS.md"
dec = read(dec_path)
dec = replace_once(dec, "version: 0.4.0", "version: 0.5.0", dec_path)
dec = replace_once(dec, "last_reviewed: 2026-08-06", f"last_reviewed: {DATE}", dec_path)
dec = replace_once(
    dec,
    "A0 and ADR-0001/0002 were explicitly accepted by the operator and merged to `main`. The first Product Milestone was subsequently selected as M0. Technical mechanisms and later ACRM gates remain open unless separately accepted/authorized.",
    "A0, the current M0 ADR set and the R4-aligned CAP-SOVEREIGN-CORE A2 package are operator-accepted; `MIS-M0-SOVEREIGN-CORE-001` v0.1.0 is the approved first M0 Mission Contract. Later ACRM gates remain explicit and separately authorized through `STATUS.md`.",
    dec_path,
)
anchor = "| D-056 | Mastra is the accepted preferred-first substrate to evaluate for first-party agentic Harnesses while sovereign truth/authority remain Aurora-owned | ADR-0009 | accepted |"
insert = anchor + "\n| D-057 | CAP-SOVEREIGN-CORE Requirements v0.1.1, Spec v0.2.0, Threat Model v0.2.0 and Test Plan v0.2.0 are the accepted R4-aligned M0 A2 package | CAP-SOVEREIGN-CORE A2 documents + R5 operator acceptance | accepted |\n| D-058 | MIS-M0-SOVEREIGN-CORE-001 v0.1.0 is the approved first scoped M0 Mission Contract | MIS-M0-SOVEREIGN-CORE-001 + R5 operator acceptance | approved |"
dec = replace_once(dec, anchor, insert, dec_path)
o15 = "| O-015 | exact first Mission Contract for selected M0 | M0 ACRM R5 Mission Contract |\n"
if o15 not in dec:
    raise SystemExit("DECISIONS O-015 row missing")
dec = dec.replace(o15, "", 1)
old_para = "`O-015` previously combined milestone selection and first Contract. The milestone portion is resolved by `D-051`; only the exact Mission Contract remains open and cannot be chosen before its applicable gates.\n\n"
if old_para not in dec:
    raise SystemExit("DECISIONS O-015 explanation missing")
dec = dec.replace(old_para, "", 1)
write(dec_path, dec)

# Append worklog.
worklog_path = "docs/tracking/WORKLOG.md"
worklog = read(worklog_path).rstrip()
entry = f'''\n\n## {DATE} — M0 R5 A2/Contract Acceptance and PASS\n\nThe operator explicitly accepted the exact R4-aligned `CAP-SOVEREIGN-CORE` A2 package and approved `MIS-M0-SOVEREIGN-CORE-001` v0.1.0. Acceptance is bound to proposal baseline `{BASELINE}` and proposal blobs recorded in `{ACCEPTANCE}`.\n\nLifecycle promotion changed acceptance metadata and stale lifecycle wording only. The workflow mechanically verified that all 122 atomic requirement table rows and all 12 Mission criterion sections remained identical to the approved proposal semantics.\n\nThe final R5 rerun verified 122/122 requirement allocation, exact Mission identity/baseline, scope/non-goals, authority/prohibitions, evidence/thresholds, replan path, accepted R4 bindings and absence of hidden R6/M1+/Mastra/AHDK/MNFS scope.\n\nFinal verdict:\n\n```text\nM0 ACRM R5 — PASS\n```\n\nR6, Microdesign and Aurora Core implementation remain NOT AUTHORIZED. Exact next action is to stop and await separate operator authorization for M0 ACRM R6 — Implementation Design Readiness.\n'''
write(worklog_path, worklog + entry)

print("R5 acceptance closeout transformation prepared")
