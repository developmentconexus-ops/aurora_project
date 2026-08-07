#!/usr/bin/env python3
from pathlib import Path
import re

ROOT = Path(__file__).resolve().parents[1]
status_path = ROOT / "docs/tracking/STATUS.md"
worklog_path = ROOT / "docs/tracking/WORKLOG.md"

status = status_path.read_text(encoding="utf-8")

required = [
    "version: 0.12.0",
    "- **Current readiness gate:** ACRM R0 — Constitutional Baseline — PASS",
    "- **R1 — Applicability:** NOT AUTHORIZED / awaiting explicit operator authorization",
    "ACRM R1 — Applicability:       NOT AUTHORIZED / AWAITING OPERATOR",
    "R0 PASS recorded\n→ stop at the R0 boundary\n→ await explicit operator authorization for M0 ACRM R1 — Applicability",
]
for item in required:
    if item not in status:
        raise SystemExit(f"STATUS precondition missing: {item}")

status = status.replace("version: 0.12.0", "version: 0.13.0", 1)

anchor = "  - REVIEW-AURORA-M0-R0-CONSTITUTIONAL-BASELINE-RERUN-2026-08-07\n"
related = (
    anchor
    + "  - DOC-AURORA-M0-R1-OPERATOR-AUTHORIZATION\n"
    + "  - DOC-AURORA-CAP-SOVEREIGN-CORE-APPLICABILITY\n"
    + "  - REVIEW-AURORA-M0-R1-APPLICABILITY-2026-08-07\n"
)
status = status.replace(anchor, related, 1)

status = status.replace(
    "- **Current readiness gate:** ACRM R0 — Constitutional Baseline — PASS",
    "- **Current readiness gate:** ACRM R1 — Applicability — PASS",
    1,
)
status = status.replace(
    "- **R1 — Applicability:** NOT AUTHORIZED / awaiting explicit operator authorization\n"
    "- **R2 and later gates:** NOT AUTHORIZED BY IMPLICATION",
    "- **R1 source baseline:** `735f269025e2cc317424e4931f3a5cd414cd6f2a`\n"
    "- **R1 applicability artifact:** `7f10734ba6018154f196557de6c5735719046253` — 294/294 classified\n"
    "- **R1 review:** `fbbae69d529a53532e5858693394747081e11d0f` — PASS\n"
    "- **R1 active constitutional sources for future R2:** 127 (`78 APPLIES + 49 PARTIALLY_APPLIES`)\n"
    "- **R2 — Requirements:** NOT AUTHORIZED / awaiting explicit operator authorization\n"
    "- **R3 and later gates:** NOT AUTHORIZED BY IMPLICATION",
    1,
)
status = status.replace(
    "The following remain open and must not be decided during R0 remediation:",
    "The following remain open and were not decided by R0 or R1; they belong to later applicable readiness gates:",
    1,
)

status = re.sub(
    r"## 6\. Current authorization boundary\n\n```text\n.*?\n```\n\n## 7\.",
    """## 6. Current authorization boundary

```text
A0 baseline:                    ACCEPTED / MERGED
ADR-0001 / ADR-0002:           ACCEPTED
First Product Milestone:        M0 SELECTED
Initial M0 R0:                 FAIL — historical
R0 documentary remediation:    ACCEPTED / MERGED
M0 ACRM R0 re-run:             PASS — target 6054f84d007347c0aa9eef9e71317134b1047d3c
M0 ACRM R1 — Applicability:    PASS — source baseline 735f269025e2cc317424e4931f3a5cd414cd6f2a
R1 applicability coverage:     294/294 classified; 127 active sources
ACRM R2 — Requirements:        NOT AUTHORIZED / AWAITING OPERATOR
ACRM R3+:                       NOT AUTHORIZED
Architecture Spike execution:   PROHIBITED
Capability implementation:      PROHIBITED
Aurora Core implementation:     PROHIBITED
AHDK implementation:            PROHIBITED
MNFS integration:               PROHIBITED
Stack selection:                NOT PERFORMED
Mission Contract:               NOT STARTED
Microdesign/Implementation Plan: NOT STARTED
```

## 7.""",
    status,
    count=1,
    flags=re.S,
)

status = re.sub(
    r"## 7\. Current blocker/gate\n\n.*?\n\n## 8\. Immediate next action",
    """## 7. Current blocker/gate

The M0 ACRM R1 applicability analysis against fixed source baseline `735f269025e2cc317424e4931f3a5cd414cd6f2a` classified all 294 accepted constitutional requirements and returned `R1 PASS`. There is no unresolved `CONFLICT_REQUIRES_DECISION` and no unjustified exclusion identified by the R1 review.

The intentional boundary is now authorization: R2 has **not** been authorized. No atomic Capability requirement derivation, Capability/System Spec, threat-model execution, Architecture Spike, stack choice, Mission Contract, Microdesign or implementation may begin by implication.

## 8. Immediate next action""",
    status,
    count=1,
    flags=re.S,
)

status = re.sub(
    r"## 8\. Immediate next action\n\n```text\n.*?\n```\n?$",
    """## 8. Immediate next action

```text
R1 PASS recorded
→ stop at the R1 boundary
→ await explicit operator authorization for M0 ACRM R2 — Requirements
```
""",
    status,
    count=1,
    flags=re.S,
)

status_path.write_text(status, encoding="utf-8")

worklog = worklog_path.read_text(encoding="utf-8")
if "version: 0.10.0" not in worklog:
    raise SystemExit("WORKLOG version precondition missing")
if "## 2026-08-07 — M0 R1 Applicability" in worklog:
    raise SystemExit("R1 worklog entry already exists")
worklog = worklog.replace("version: 0.10.0", "version: 0.11.0", 1)
worklog += """

## 2026-08-07 — M0 R1 Applicability

After the M0 R0 re-run passed, the operator explicitly authorized proceeding to `ACRM R1 — Applicability` by responding “Vamos seguir” to the R1 authorization boundary. The authorization was recorded in `docs/acceptance/2026-08-07-m0-r1-operator-authorization.md` and did not authorize R2 or later work.

R1 fixed the canonical source baseline at:

```text
735f269025e2cc317424e4931f3a5cd414cd6f2a
```

The complete applicability artifact was created at:

```text
docs/capabilities/CAP-SOVEREIGN-CORE/APPLICABILITY.md
Applicability commit: 7f10734ba6018154f196557de6c5735719046253
```

All 294 accepted constitutional requirements were classified:

```text
APPLIES:                    78
PARTIALLY_APPLIES:          49
DEFERRED_BY_ROADMAP:       161
NOT_APPLICABLE:              6
CONFLICT_REQUIRES_DECISION:  0
TOTAL:                      294
ACTIVE SOURCES FOR R2:      127
```

The review preserved the high-risk M0 dependencies for sovereign identity, operational-state ownership, authority/restore safety, migration, backup/restore integrity, event/audit versus telemetry, security of the first canonical durable store and the CLI/interface boundary. It also deliberately deferred M1 memory, M2 Registry/AHDK, cross-Harness Delegation/effects, Presence, laboratory/physical control, adaptive campaigns and self-improvement.

Normal Documentation validation passed for the applicability artifact in workflow run `31146692949`. A separate R1 review was then recorded at `docs/reviews/2026-08-07-m0-r1-applicability-review.md` (commit `fbbae69d529a53532e5858693394747081e11d0f`), and its documentation validation job also completed successfully in workflow run `31146839010`.

R1 verdict:

```text
R1 PASS
```

No language, framework, database, storage, runtime, topology, schema, protocol, telemetry backend, Architecture Spike winner, Mission Contract or implementation was selected. R2 remains separately gated and unauthorized until an explicit operator decision.
"""
worklog_path.write_text(worklog, encoding="utf-8")

for rel in [
    "scripts/close_m0_r1.py",
    ".github/workflows/close-m0-r1.yml",
    ".github/m0-r1-closeout.trigger",
]:
    path = ROOT / rel
    if path.exists():
        path.unlink()
