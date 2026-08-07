#!/usr/bin/env python3
from pathlib import Path

ACCEPTANCE_ID = "DOC-AURORA-M0-R0-REMEDIATION-OPERATOR-ACCEPTANCE"
ACCEPTED_AT = "2026-08-07"
APPROVED_SEMANTIC_SHA = "b32cfe134f84eed3797d866e607c92c227514186"

NORMATIVE = [
    "docs/DOCUMENTATION-MAP.md",
    "docs/adr/README.md",
    "docs/product/README.md",
    "docs/product/REQUIREMENTS-TRACEABILITY.md",
    "docs/product/blueprint/01-product-vision.md",
    "docs/product/blueprint/14-capability-roadmap.md",
    "docs/product/blueprint/15-documentation-research-governance.md",
]

for filename in NORMATIVE:
    path = Path(filename)
    text = path.read_text(encoding="utf-8")
    if "status: proposed" not in text:
        raise SystemExit(f"{filename}: expected status: proposed")
    if "accepted_at:" in text or "acceptance_evidence:" in text:
        raise SystemExit(f"{filename}: unexpected acceptance metadata before promotion")
    text = text.replace(
        "status: proposed\n",
        f"status: accepted\naccepted_at: {ACCEPTED_AT}\nacceptance_evidence: {ACCEPTANCE_ID}\n",
        1,
    )
    if "last_reviewed: 2026-08-06" in text:
        text = text.replace("last_reviewed: 2026-08-06", f"last_reviewed: {ACCEPTED_AT}", 1)
    path.write_text(text, encoding="utf-8")

status_path = Path("docs/tracking/STATUS.md")
status = status_path.read_text(encoding="utf-8")
status = status.replace("version: 0.9.0", "version: 0.10.0", 1)
if f"  - {ACCEPTANCE_ID}\n" not in status:
    marker = "  - DOC-AURORA-M0-R0-REMEDIATION-AUTHORIZATION\n"
    if marker not in status:
        raise SystemExit("STATUS: related marker not found")
    status = status.replace(marker, marker + f"  - {ACCEPTANCE_ID}\n", 1)
status = status.replace("last_reviewed: 2026-08-06", f"last_reviewed: {ACCEPTED_AT}", 1)
old = "- **R0 remediation:** AUTHORIZED / PROPOSED on non-canonical documentation branch / OPERATOR ACCEPTANCE OF CORRECTED REVISION PENDING"
new = "- **R0 remediation:** CORRECTED REVISION ACCEPTED BY OPERATOR / CANONICAL INTEGRATION AUTHORIZED"
if old not in status:
    raise SystemExit("STATUS: remediation summary marker not found")
status = status.replace(old, new, 1)
status = status.replace(
    "Corrected constitutional rev:  OPERATOR ACCEPTANCE PENDING\nR0 re-run:                      PENDING corrected accepted revision",
    "Corrected constitutional rev:  OPERATOR ACCEPTED\nCanonical integration:         AUTHORIZED / PENDING MERGE\nR0 re-run:                      AUTHORIZED AFTER CANONICAL INTEGRATION",
    1,
)
status = status.replace(
    "The intentional blocker is the failed R0 constitutional baseline. The corrected documentation revision must be independently reviewable, validated and explicitly accepted before it can become the baseline for an R0 re-run.",
    f"The corrected R0 remediation revision `{APPROVED_SEMANTIC_SHA}` has been explicitly accepted by the operator. The remaining gate is canonical integration of the validated accepted revision; R1 remains unauthorized.",
    1,
)
status = status.replace(
    "complete and validate R0 documentary remediation\n→ operator reviews corrected constitutional revision\n→ if accepted, integrate/promote it to canonical main\n→ start a fresh repository-only session against the accepted revision",
    "integrate the accepted R0 remediation revision into canonical main\n→ record the canonical merge/closeout revision\n→ start a fresh repository-only R0 review against that accepted revision",
    1,
)
status_path.write_text(status, encoding="utf-8")

worklog_path = Path("docs/tracking/WORKLOG.md")
worklog = worklog_path.read_text(encoding="utf-8")
worklog = worklog.replace("version: 0.7.0", "version: 0.8.0", 1)
worklog = worklog.replace("last_reviewed: 2026-08-06", f"last_reviewed: {ACCEPTED_AT}", 1)
entry = f'''\n\n## 2026-08-07 — M0 R0 remediation revision accepted\n\nThe operator reviewed the corrected R0 remediation candidate at exact semantic revision:\n\n```text\n{APPROVED_SEMANTIC_SHA}\n```\n\nand responded:\n\n> “Aprovo A revisao”\n\nThe decision accepts the corrected documentary/constitutional meaning and authorizes its lifecycle promotion and canonical integration. Promotion metadata and acceptance evidence may change after the approved semantic revision; the approved constitutional content itself must not drift during promotion.\n\nThis acceptance does **not** authorize R1 or any later ACRM gate, Architecture Spike execution, stack selection, Aurora Core/AHDK/MNFS implementation, Mission Contract or Microdesign. After canonical integration, the next allowed action is a fresh repository-only re-run of M0 ACRM R0 only.\n'''
if "## 2026-08-07 — M0 R0 remediation revision accepted" not in worklog:
    worklog += entry
worklog_path.write_text(worklog, encoding="utf-8")

Path(".github/workflows/promote-m0-r0-remediation.yml").unlink()
Path("scripts/promote_m0_r0_remediation.py").unlink()
