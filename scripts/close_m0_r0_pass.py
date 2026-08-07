#!/usr/bin/env python3
from pathlib import Path

TARGET = "6054f84d007347c0aa9eef9e71317134b1047d3c"
REVIEW_ID = "REVIEW-AURORA-M0-R0-CONSTITUTIONAL-BASELINE-RERUN-2026-08-07"
TARGET_FINDING_ID = "DOC-AURORA-M0-R0-RERUN-TARGET-FINDING"
DATE = "2026-08-07"

status_path = Path("docs/tracking/STATUS.md")
status = status_path.read_text(encoding="utf-8")
status = status.replace("version: 0.11.1", "version: 0.12.0", 1)
marker = "  - DOC-AURORA-M0-R0-REMEDIATION-MERGE-CLOSEOUT\n"
if marker not in status:
    raise SystemExit("STATUS related marker not found")
addition = ""
if f"  - {TARGET_FINDING_ID}\n" not in status:
    addition += f"  - {TARGET_FINDING_ID}\n"
if f"  - {REVIEW_ID}\n" not in status:
    addition += f"  - {REVIEW_ID}\n"
status = status.replace(marker, marker + addition, 1)
status = status.replace(
    "- **Current readiness gate:** ACRM R0 — Constitutional Baseline\n- **Initial R0 verdict:** FAIL",
    f"- **Current readiness gate:** ACRM R0 — Constitutional Baseline — PASS\n- **Initial R0 verdict:** FAIL\n- **R0 re-run target:** `{TARGET}`\n- **R0 re-run verdict:** PASS",
    1,
)
status = status.replace(
    "- **R1 and later gates:** NOT AUTHORIZED BY IMPLICATION",
    "- **R1 — Applicability:** NOT AUTHORIZED / awaiting explicit operator authorization\n- **R2 and later gates:** NOT AUTHORIZED BY IMPLICATION",
    1,
)
old_boundary = """First Product Milestone:        M0 SELECTED
Initial M0 R0:                 FAIL
R0 documentary remediation:    AUTHORIZED
Corrected constitutional rev:  OPERATOR ACCEPTED
Canonical integration:         COMPLETE — d0ddfb794296e599ac96bb73bf3772937d371bf9
R0 re-run:                      AUTHORIZED / CURRENT NEXT ACTION
ACRM R1+:                       NOT AUTHORIZED"""
new_boundary = f"""First Product Milestone:        M0 SELECTED
Initial M0 R0:                 FAIL — historical
R0 documentary remediation:    ACCEPTED / MERGED
Corrected constitutional rev:  OPERATOR ACCEPTED
Canonical integration:         COMPLETE — d0ddfb794296e599ac96bb73bf3772937d371bf9
M0 ACRM R0 re-run:             PASS — target {TARGET}
ACRM R1 — Applicability:       NOT AUTHORIZED / AWAITING OPERATOR
ACRM R2+:                       NOT AUTHORIZED"""
if old_boundary not in status:
    raise SystemExit("STATUS authorization boundary block not found")
status = status.replace(old_boundary, new_boundary, 1)
old_gate = """The corrected R0 remediation has been accepted and canonically integrated at `d0ddfb794296e599ac96bb73bf3772937d371bf9`. There is no remaining remediation blocker. The current gate is the fresh M0 ACRM R0 re-run; R1 remains unauthorized.

No implementation blocker exists because implementation is not authorized work."""
new_gate = f"""The M0 ACRM R0 re-run against fixed canonical target `{TARGET}` returned `R0 PASS`. There is no unresolved material constitutional baseline blocker for M0 applicability analysis.

The intentional boundary is now authorization: R1 has **not** been authorized. No applicability package, requirement derivation, Architecture Spike execution, stack choice, Mission Contract, Microdesign or implementation may begin by implication."""
if old_gate not in status:
    raise SystemExit("STATUS gate paragraph not found")
status = status.replace(old_gate, new_gate, 1)
old_next = """start a fresh repository-only R0 review from current canonical `main`
→ resolve and record the exact `main` HEAD as the fixed R0 target revision before reading scope sources
→ read AGENTS.md and STATUS from that exact revision
→ execute M0 ACRM R0 only
→ produce R0 PASS | FAIL | BLOCKED
→ stop before R1 unless R1 is separately authorized"""
new_next = """R0 PASS recorded
→ stop at the R0 boundary
→ await explicit operator authorization for M0 ACRM R1 — Applicability"""
if old_next not in status:
    raise SystemExit("STATUS next-action block not found")
status = status.replace(old_next, new_next, 1)
status_path.write_text(status, encoding="utf-8")

worklog_path = Path("docs/tracking/WORKLOG.md")
worklog = worklog_path.read_text(encoding="utf-8")
worklog = worklog.replace("version: 0.9.0", "version: 0.10.0", 1)
entry = f'''\n\n## 2026-08-07 — M0 ACRM R0 re-run PASS\n\nAfter the operator-accepted R0 remediation was canonically integrated and the fresh-review target procedure was repaired, the re-run fixed the immutable repository target at:\n\n```text\n{TARGET}\n```\n\nThe review re-read the mandatory bootstrap/current-state path and M0-relevant constitutional owners from that exact revision, including Product Vision, Domain/World Model, Cognitive Lifecycle, operational-state/memory separation, Harness boundary, Authority/Safety, Security/Sovereignty, System Architecture, Reliability/Evaluation, Blueprint 14, Blueprint 15, Requirements Traceability, ADR-0001/0002, decision/coverage/research indexes and operator evidence.\n\nFormal R0 result:\n\n```text\nR0 PASS\n```\n\nAll initial gate-failing findings were resolved:\n\n- R0-F01 — executable-horizon M0 milestone anatomy: RESOLVED;\n- R0-F02 — ADR status divergence: RESOLVED;\n- R0-F03 — mutable-state duplication/drift: RESOLVED;\n- R0-F04 — re-run target continuity: RESOLVED as tracking-only continuity repair before the fixed review.\n\nThe review confirmed that remaining open work belongs to later gates rather than missing constitutional intent: R1 applicability classification, R2 verifiable requirements, R3 `CAP-SOVEREIGN-CORE` Capability/System Spec and threat/test model, R4 technical decisions/spikes, R5 Mission Contract and R6 Microdesign.\n\nOne non-blocking documentation-hygiene note remains: the Documentation Map entrypoint table still describes `REQUIREMENTS-TRACEABILITY.md` with the historical word “proposed” even though the canonical requirement document is accepted. This does not compete with requirement lifecycle authority or current authorization and was not silently edited during the R0 verdict.\n\nReview record:\n\n```text\ndocs/reviews/2026-08-07-m0-r0-constitutional-baseline-rerun.md\n```\n\nR1 is **not authorized** by the PASS verdict. Architecture Spike execution, stack selection, Aurora Core/AHDK/MNFS implementation, Mission Contract and Microdesign remain unauthorized. The exact next action is to stop and await explicit operator authorization for `M0 ACRM R1 — Applicability`.\n'''
if "## 2026-08-07 — M0 ACRM R0 re-run PASS" not in worklog:
    worklog += entry
worklog_path.write_text(worklog, encoding="utf-8")

Path(".github/workflows/close-m0-r0-pass.yml").unlink()
Path("scripts/close_m0_r0_pass.py").unlink()
