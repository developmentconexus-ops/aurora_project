#!/usr/bin/env python3
from pathlib import Path


def replace_once(path: str, old: str, new: str) -> None:
    p = Path(path)
    text = p.read_text(encoding="utf-8")
    if old not in text:
        raise SystemExit(f"expected text not found in {path}: {old[:120]!r}")
    p.write_text(text.replace(old, new, 1), encoding="utf-8")


normative = [
    "docs/DOCUMENTATION-MAP.md",
    "docs/product/README.md",
    "docs/product/REQUIREMENTS-TRACEABILITY.md",
    "docs/product/blueprint/01-product-vision.md",
    "docs/product/blueprint/14-capability-roadmap.md",
    "docs/product/blueprint/15-documentation-research-governance.md",
    "docs/adr/README.md",
]

accepted_block = "status: accepted\naccepted_at: 2026-08-06\nacceptance_evidence: DOC-AURORA-A0-OPERATOR-ACCEPTANCE"
for path in normative:
    replace_once(path, accepted_block, "status: proposed")

replace_once(
    "docs/adr/README.md",
    """## 5. Acceptance gate

Before acceptance:

- relevant ACRM R0–R3 gates are current;
- material uncertainty has evidence or an authorized spike;
- the decision does not silently expand current Product Milestone scope;
- affected requirements and documents are identified;
- operational burden and exit conditions are explicit;
- Leandro reviews material architecture decisions.

Acceptance of an ADR does not automatically authorize a Capability Spec, Mission Contract, Architecture Spike or implementation.
""",
    """## 5. Acceptance gate

For post-A0 Capability/Product-Milestone-scoped ADRs, before acceptance:

- relevant ACRM R0–R3 gates are current;
- material uncertainty has evidence or an authorized spike;
- the decision does not silently expand current Product Milestone scope;
- affected requirements and documents are identified;
- operational burden and exit conditions are explicit;
- Leandro reviews material architecture decisions.

ADR-0001 and ADR-0002 are A0 baseline decisions that were explicitly accepted through the A0 operator gate before the first Product Milestone readiness cycle. This post-A0 rule does not retroactively alter their accepted decision status.

Acceptance of an ADR does not automatically authorize a Capability Spec, Mission Contract, Architecture Spike or implementation.
""",
)

replace_once(
    "docs/tracking/WORKLOG.md",
    "R1, Architecture Spike execution, stack selection, Aurora Core/AHDK/MNFS implementation, Mission Contract and Microdesign remain unauthorized.",
    "R1, Architecture Spike execution, stack selection, Aurora Core/AHDK/MNFS implementation, Mission Contract and Microdesign remain unauthorized. The revised normative documents remain `PROPOSED` on the non-canonical branch until explicit operator acceptance; the last accepted versions remain canonical on `main`.",
)

# Assert no revised normative file falsely claims the new revision is already accepted.
for path in normative:
    text = Path(path).read_text(encoding="utf-8")
    if "status: accepted" in text.split("---", 2)[1]:
        raise SystemExit(f"candidate revision still marked accepted: {path}")
    if "status: proposed" not in text.split("---", 2)[1]:
        raise SystemExit(f"candidate revision not marked proposed: {path}")

Path(".github/workflows/align-m0-r0-candidate-lifecycle.yml").unlink()
Path("scripts/align_m0_r0_candidate_lifecycle.py").unlink()
