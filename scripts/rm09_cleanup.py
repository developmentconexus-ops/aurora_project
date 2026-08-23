#!/usr/bin/env python3
from __future__ import annotations

import re
import subprocess
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]


def path(rel: str) -> Path:
    return ROOT / rel


def update(rel: str, replacements: list[tuple[str, str]], regexes: list[tuple[str, str]] | None = None) -> None:
    p = path(rel)
    text = p.read_text(encoding="utf-8")
    original = text
    for old, new in replacements:
        text = text.replace(old, new)
    for pattern, repl in regexes or []:
        text = re.sub(pattern, repl, text)
    if text != original:
        p.write_text(text, encoding="utf-8")


# Current architecture: current program owner + current Harness name.
for rel in [
    "docs/architecture/capability-harness-boundary.md",
    "docs/architecture/module-runtime-topology.md",
    "docs/architecture/stage-a-availability-activation.md",
    "docs/architecture/system-architecture.md",
    "docs/architecture/system-decision-landscape.md",
]:
    update(
        rel,
        [
            ("DOC-AURORA-STATUS", "DOC-AURORA-REPOSITORY-ROADMAP"),
            ("repair STATUS / DECISIONS / WORKLOG continuity", "repair docs/roadmap.md / decision-register / Git-Evidence continuity"),
            ("ACRM integration and current STATUS agree", "ACRM integration and current docs/roadmap.md agree"),
            ("MNFS", "Conexus OS"),
        ],
    )

# M0 Capability package remains accepted M0 semantics but consumes the current repository control plane.
cap_root = path("docs/capabilities/CAP-SOVEREIGN-CORE")
for p in sorted(cap_root.glob("*.md")):
    text = p.read_text(encoding="utf-8")
    text = text.replace("DOC-AURORA-STATUS", "DOC-AURORA-REPOSITORY-ROADMAP")
    text = text.replace("MNFS", "Conexus OS")
    text = text.replace("STATUS.md", "docs/roadmap.md")
    text = re.sub(r"\bSTATUS\b", "docs/roadmap.md", text)
    p.write_text(text, encoding="utf-8")

# Current ADR terminology.
update("docs/decisions/adr/0001-aurora-owned-contract-model.md", [("MNFS", "Conexus OS")])

# Decision register: remove stale 'current' program language while preserving history/dispositions.
update(
    "docs/decisions/index.md",
    [
        (
            "Current implementation remains paused while the accepted Technical Architecture Baseline begins with TA-01/TA-02.",
            "Product implementation remains blocked. TA-01/TA-02 are accepted/canonical; remaining cross-system planning order is owned by the current Planning and Implementation-Readiness Standard.",
        ),
        ("Blueprint 15, Documentation Map", "Blueprint 15, Documentation Index"),
        (
            "| D-064 | the M0 R7 implementation candidate remains frozen, preserved and non-canonical; code/CI existence is not acceptance | operator direction + STATUS | accepted current coordination |",
            "| D-064 | the M0 R7 implementation candidate remains frozen, preserved and non-canonical; code/CI existence is not acceptance | operator direction + current roadmap/phase Evidence | accepted |",
        ),
        (
            "| D-065 | current program priority is the cross-system Technical Architecture Baseline, not additional broad product-definition dialogue or Presence micro-policy decomposition | Technical Architecture Map + operator acceptance | accepted |",
            "| D-065 | before MR-01, program priority moved to the cross-system Technical Architecture Baseline rather than more broad Product discovery or Presence micro-policy | Technical Architecture Map + operator acceptance | superseded by MR-01 program sequencing |",
        ),
        (
            "| D-066 | a technical question is current only when it changes ownership, structural/runtime/contract/security/data boundaries or the next implementation decision; otherwise it is deferred | Technical Architecture Map | accepted |",
            "| D-066 | a technical question is current only when it changes ownership, structural/runtime/contract/security/data boundaries or the next implementation decision; otherwise it is deferred | MR-01 Planning Readiness | accepted |",
        ),
        (
            "| D-068 | accepted Stage A Presence/activation/locked-workstation rules remain downstream constraints, while further session-policy detail is deferred until a consuming Capability | Stage A design + Technical Architecture Map | accepted |",
            "| D-068 | accepted Stage A Presence/activation/locked-workstation rules remain downstream constraints, while further session-policy detail is deferred until a consuming Capability | Stage A design + MR-01 Planning Readiness | accepted |",
        ),
        (
            "The complete dependency map is `DESIGN-AURORA-TECHNICAL-ARCHITECTURE-BASELINE-MAP` and `DESIGN-AURORA-SYSTEM-ARCHITECTURE-DECISION-LANDSCAPE`.",
            "The remaining cross-system dependency order is owned by `DOC-AURORA-PLANNING-READINESS`; structural open questions remain indexed by `DESIGN-AURORA-SYSTEM-ARCHITECTURE-DECISION-LANDSCAPE` and this decision register.",
        ),
        ("roadmap readiness; MNFS remains one candidate", "roadmap readiness; Conexus OS remains one candidate"),
    ],
)

# Planning-readiness language is accepted, not proposed.
update(
    "docs/development/planning-readiness.md",
    [
        ("proposed relationship between global readiness and ACRM", "relationship between global readiness and ACRM"),
        ("proposed technology-research and Paved-Road decision law", "technology-research and Paved-Road decision law"),
        (
            "TA-01 and TA-02 are already accepted/canonical at the time of this proposal and are not reopened by naming this graph.",
            "TA-01 and TA-02 are accepted/canonical and are not reopened by this graph.",
        ),
    ],
)

# ACRM consumes the new router explicitly.
update("docs/development/capability-realization.md", [("Documentation Map;", "`docs/index.md`;"), ("Documentation Map", "Documentation Index")])

# Current Product references: preserve explicitly historical origin/A0/M0 wording, update present/future references.
update(
    "docs/product/blueprint/01-product-vision.md",
    [("MNFS\n→ future software-engineering harness", "Conexus OS (historically MNFS)\n→ future software-engineering harness")],
)
update("docs/product/blueprint/03-domain-world-model.md", [("PRJ-MNFS", "PRJ-CONEXUS-OS")])
update("docs/product/blueprint/05-capability-system.md", [("integrating MNFS", "integrating Conexus OS")])
update("docs/product/blueprint/06-memory-knowledge-context.md", [("Aurora USES MNFS as a future provider", "Aurora USES Conexus OS as a future provider")])
update("docs/product/blueprint/12-system-architecture.md", [("│ MNFS       │", "│ Conexus OS  │")])
update(
    "docs/product/blueprint/14-capability-roadmap.md",
    [
        ("real MNFS/lab integration", "real Conexus OS/lab integration"),
        ("- MNFS;\n- Firmware Harness;", "- Conexus OS;\n- Firmware Harness;"),
        ("MNFS is a strong future candidate", "Conexus OS is a strong future candidate"),
        ("require MNFS as M6 provider", "require Conexus OS as M6 provider"),
    ],
)
update(
    "docs/product/blueprint/15-documentation-research-governance.md",
    [
        ("| STATUS | Tracking | Reference |", "| `docs/roadmap.md` | Program coordination | Reference |"),
        ("- generated roadmap;", "- generated Product Blueprint publication;"),
        ("| Product sequence | Blueprint 14 / generated roadmap |", "| Product sequence | Blueprint 14 |"),
    ],
)

# Traceability follows the constitutional owner refinement; requirement IDs remain stable.
update(
    "docs/product/REQUIREMENTS-TRACEABILITY.md",
    [
        ("STATUS/readiness check", "docs/roadmap.md/readiness check"),
        ("STATUS MUST state current gate, authorizations/prohibitions, blockers, verification and exact next action.", "`docs/roadmap.md` MUST state current gate, authorizations/prohibitions, blockers, verification and exact next action."),
        ("STATUS/gate enforcement", "docs/roadmap.md/gate enforcement"),
        ("MNFS MUST be integrated as a future provider", "Conexus OS MUST be integrated as a future provider"),
        ("not assume MNFS automatically", "not assume Conexus OS automatically"),
    ],
)

update("docs/reference/architecture-spikes.md", [("STATUS records whether any next step is authorized.", "`docs/roadmap.md` records whether any next step is authorized.")])

# Regenerate only the Product aggregate after source changes.
subprocess.check_call([sys.executable, str(ROOT / "scripts/generate_docs.py")], cwd=ROOT)

# Remove this one-shot helper from the resulting commit.
path("scripts/rm09_cleanup.py").unlink()
