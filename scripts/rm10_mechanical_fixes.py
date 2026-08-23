#!/usr/bin/env python3
from __future__ import annotations

import subprocess
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]


def replace_once(path: Path, old: str, new: str) -> None:
    text = path.read_text(encoding="utf-8")
    count = text.count(old)
    if count != 1:
        raise RuntimeError(f"{path.relative_to(ROOT)} expected exactly one match, found {count}: {old!r}")
    path.write_text(text.replace(old, new, 1), encoding="utf-8")


# RM-I02: the two highest-value current router/program owners must participate in link validation.
replace_once(
    ROOT / "scripts/validate_docs.py",
    '        "CONTRIBUTING.md",\n        "docs/product",',
    '        "CONTRIBUTING.md",\n        "docs/index.md",\n        "docs/roadmap.md",\n        "docs/product",',
)

# RM-I05: archived snapshots must not advertise present-tense source-of-truth scope.
replace_once(
    ROOT / "docs/phases/pre-mr01-status-snapshot.md",
    "source_of_truth_for:\n  - current Aurora project phase\n  - current authorization boundary\n  - current blockers and immediate next action",
    "source_of_truth_for:\n  - pre-MR-01 repository-program phase snapshot\n  - pre-MR-01 authorization boundary snapshot\n  - pre-MR-01 blocker and next-action snapshot",
)
replace_once(
    ROOT / "docs/evidence/project-worklog.md",
    "source_of_truth_for:\n  - chronological material work history",
    "source_of_truth_for:\n  - pre-MR-01 chronological material work history through the repository-model cutover",
)

# RM-I07: remove duplicate future CODEOWNERS entry without changing ownership semantics.
replace_once(
    ROOT / "docs/product/blueprint/15-documentation-research-governance.md",
    "/docs/development/\n/docs/development/\n",
    "/docs/development/\n",
)

# Regenerate with the already-corrected checkout-independent generator (RM-I01).
subprocess.run(["python", "scripts/generate_docs.py"], cwd=ROOT, check=True)

# One-shot migration helper must not survive the correction commit.
Path(__file__).unlink()
