#!/usr/bin/env python3
from __future__ import annotations

import argparse
import hashlib
import re
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
BLUEPRINT_DIR = Path("docs/product/blueprint")
BLUEPRINT_OUTPUT = Path("docs/product/PRODUCT-BLUEPRINT.md")


def split_frontmatter(text: str, path: Path) -> tuple[str, str]:
    if not text.startswith("---\n"):
        raise ValueError(f"{path}: missing YAML frontmatter")
    marker = text.find("\n---\n", 4)
    if marker < 0:
        raise ValueError(f"{path}: unterminated YAML frontmatter")
    return text[4:marker], text[marker + 5:].lstrip("\n")


def scalar(frontmatter: str, key: str) -> str:
    m = re.search(rf"(?m)^{re.escape(key)}:\s*([^\n]+)\s*$", frontmatter)
    if not m:
        raise ValueError(f"frontmatter missing {key}")
    return m.group(1).strip().strip('"').strip("'")


def canonical_sources(root: Path) -> list[Path]:
    sources = sorted((root / BLUEPRINT_DIR).glob("[0-9][0-9]-*.md"))
    actual = [int(x.name[:2]) for x in sources]
    if actual != list(range(1, 16)):
        raise ValueError(f"expected Blueprint 01..15, found {actual}")
    return sources


def generate_blueprint(root: Path) -> str:
    sections = []
    for path in canonical_sources(root):
        raw = path.read_text(encoding="utf-8")
        fm, body = split_frontmatter(raw, path)
        relative_path = path.relative_to(root)
        sections.append(
            (
                relative_path,
                scalar(fm, "id"),
                hashlib.sha256(raw.encode()).hexdigest(),
                body.rstrip(),
            )
        )
    ids = [x[1] for x in sections]
    out = [
        "---",
        "id: DOC-AURORA-PRODUCT-BLUEPRINT",
        "title: Aurora Product Blueprint",
        "document_type: product_blueprint_aggregate",
        "form: explanation",
        "authority: generated_projection",
        "status: generated",
        "version: 0.3.0",
        "owners:",
        "  - developmentconexus-ops",
        "generated_from:",
        *[f"  - {x}" for x in ids],
        "---",
        "",
        "<!-- GENERATED — DO NOT EDIT DIRECTLY",
        "Canonical sources: docs/product/blueprint/01-*.md through 15-*.md",
        "Generator: scripts/generate_docs.py",
        "-->",
        "",
        "# Aurora Product Blueprint",
        "",
        "> Generated publication of the fifteen modular constitutional sources. Edit sources, regenerate, validate.",
        "",
        "## Source manifest",
        "",
        "| Section | Canonical source | SHA-256 |",
        "|---:|---|---|",
    ]
    for path, _, digest, _ in sections:
        out.append(f"| {path.name[:2]} | `{path.as_posix()}` | `{digest}` |")
    out += ["", "---", ""]
    for i, (path, _, _, body) in enumerate(sections):
        out += [f"<!-- BEGIN SOURCE: {path.as_posix()} -->", body, f"<!-- END SOURCE: {path.as_posix()} -->"]
        if i != len(sections) - 1:
            out += ["", "---", ""]
    return "\n".join(out).rstrip() + "\n"


def main() -> int:
    ap = argparse.ArgumentParser()
    ap.add_argument("--output-root", type=Path, default=ROOT)
    ap.add_argument("--check", action="store_true")
    args = ap.parse_args()
    content = generate_blueprint(ROOT)
    if args.check:
        target = ROOT / BLUEPRINT_OUTPUT
        if not target.exists() or target.read_text(encoding="utf-8") != content:
            print(f"STALE GENERATED DOCUMENT: {BLUEPRINT_OUTPUT}")
            return 1
        print("Generated Product Blueprint is current; repository roadmap is hand-maintained.")
        return 0
    target = args.output_root.resolve() / BLUEPRINT_OUTPUT
    target.parent.mkdir(parents=True, exist_ok=True)
    target.write_text(content, encoding="utf-8")
    print(target)
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
