#!/usr/bin/env python3
from pathlib import Path

MERGE_SHA = "d0ddfb794296e599ac96bb73bf3772937d371bf9"
CLOSEOUT_ID = "DOC-AURORA-M0-R0-REMEDIATION-MERGE-CLOSEOUT"
DATE = "2026-08-07"

status_path = Path("docs/tracking/STATUS.md")
status = status_path.read_text(encoding="utf-8")
status = status.replace("version: 0.10.0", "version: 0.11.0", 1)
if f"  - {CLOSEOUT_ID}\n" not in status:
    marker = "  - DOC-AURORA-M0-R0-REMEDIATION-OPERATOR-ACCEPTANCE\n"
    if marker not in status:
        raise SystemExit("STATUS: acceptance relation marker not found")
    status = status.replace(marker, marker + f"  - {CLOSEOUT_ID}\n", 1)
status = status.replace("last_reviewed: 2026-08-07", f"last_reviewed: {DATE}", 1)
status = status.replace(
    "- **R0 remediation:** CORRECTED REVISION ACCEPTED BY OPERATOR / CANONICAL INTEGRATION AUTHORIZED",
    f"- **R0 remediation:** ACCEPTED AND MERGED — `{MERGE_SHA}`",
    1,
)
summary_marker = "- **Initial R0 verdict:** FAIL\n"
if summary_marker in status and "- **Canonical R0 remediation merge:**" not in status:
    status = status.replace(summary_marker, summary_marker + f"- **Canonical R0 remediation merge:** `{MERGE_SHA}`\n", 1)
status = status.replace(
    "Canonical integration:         AUTHORIZED / PENDING MERGE\nR0 re-run:                      AUTHORIZED AFTER CANONICAL INTEGRATION",
    f"Canonical integration:         COMPLETE — {MERGE_SHA}\nR0 re-run:                      AUTHORIZED / CURRENT NEXT ACTION",
    1,
)
status = status.replace(
    "The corrected R0 remediation revision `b32cfe134f84eed3797d866e607c92c227514186` has been explicitly accepted by the operator. The remaining gate is canonical integration of the validated accepted revision; R1 remains unauthorized.",
    f"The corrected R0 remediation has been accepted and canonically integrated at `{MERGE_SHA}`. There is no remaining remediation blocker. The current gate is the fresh M0 ACRM R0 re-run; R1 remains unauthorized.",
    1,
)
status = status.replace(
    "integrate the accepted R0 remediation revision into canonical main\n→ record the canonical merge/closeout revision\n→ start a fresh repository-only R0 review against that accepted revision",
    f"start a fresh repository-only R0 review against canonical revision {MERGE_SHA}\n→ read AGENTS.md and STATUS from that revision\n→ execute M0 ACRM R0 only",
    1,
)
status_path.write_text(status, encoding="utf-8")

worklog_path = Path("docs/tracking/WORKLOG.md")
worklog = worklog_path.read_text(encoding="utf-8")
worklog = worklog.replace("version: 0.8.0", "version: 0.9.0", 1)
worklog = worklog.replace("last_reviewed: 2026-08-07", f"last_reviewed: {DATE}", 1)
entry = f'''\n\n## 2026-08-07 — M0 R0 remediation canonically integrated\n\nThe operator-accepted R0 remediation was merged through PR #2 into canonical `main`:\n\n```text\nMerge commit: {MERGE_SHA}\n```\n\nPre-merge evidence included semantic-drift verification from approved revision `b32cfe134f84eed3797d866e607c92c227514186`, accepted-lifecycle promotion, generated projection refresh, promotion validation run `31144371490` and normal Documentation workflow run `31144424887`, both successful.\n\nThe remediation merge does not authorize R1, Architecture Spike execution, stack selection, Aurora Core/AHDK/MNFS implementation, Mission Contract or Microdesign. The next authorized action is a fresh repository-only M0 ACRM R0 re-run.\n'''
if "## 2026-08-07 — M0 R0 remediation canonically integrated" not in worklog:
    worklog += entry
worklog_path.write_text(worklog, encoding="utf-8")

Path(".github/workflows/close-m0-r0-remediation-merge.yml").unlink()
Path("scripts/close_m0_r0_remediation_merge.py").unlink()
