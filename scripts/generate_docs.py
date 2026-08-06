#!/usr/bin/env python3
"""Generate Aurora documentation projections from canonical modular sources.

This script is documentation tooling only. It does not select or constrain the
Aurora product runtime.
"""

from __future__ import annotations

import argparse
import hashlib
import re
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
BLUEPRINT_DIR = Path("docs/product/blueprint")
BLUEPRINT_OUTPUT = Path("docs/product/PRODUCT-BLUEPRINT.md")
ROADMAP_SOURCE = BLUEPRINT_DIR / "14-capability-roadmap.md"
ROADMAP_OUTPUT = Path("docs/roadmap.md")


def split_frontmatter(text: str, path: Path) -> tuple[str, str]:
    if not text.startswith("---\n"):
        raise ValueError(f"{path}: missing YAML frontmatter")
    marker = text.find("\n---\n", 4)
    if marker < 0:
        raise ValueError(f"{path}: unterminated YAML frontmatter")
    return text[4:marker], text[marker + 5 :].lstrip("\n")


def scalar(frontmatter: str, key: str) -> str:
    match = re.search(rf"(?m)^{re.escape(key)}:\s*([^\n]+)\s*$", frontmatter)
    if not match:
        raise ValueError(f"frontmatter missing scalar field {key!r}")
    return match.group(1).strip().strip('"').strip("'")


def sha256(text: str) -> str:
    return hashlib.sha256(text.encode("utf-8")).hexdigest()


def canonical_sources(root: Path) -> list[Path]:
    source_dir = root / BLUEPRINT_DIR
    sources = sorted(source_dir.glob("[0-9][0-9]-*.md"))
    expected = list(range(1, 16))
    actual = [int(path.name[:2]) for path in sources]
    if actual != expected:
        raise ValueError(
            f"expected Blueprint sections 01..15, found {actual or 'none'}"
        )
    return sources


def generated_frontmatter(
    *,
    doc_id: str,
    title: str,
    document_type: str,
    version: str,
    generated_from: list[str],
) -> str:
    lines = [
        "---",
        f"id: {doc_id}",
        f"title: {title}",
        f"document_type: {document_type}",
        "form: explanation",
        "authority: generated_projection",
        "status: generated",
        f"version: {version}",
        "owners:",
        "  - developmentconexus-ops",
        "generated_from:",
    ]
    lines.extend(f"  - {item}" for item in generated_from)
    lines.extend(["---", ""])
    return "\n".join(lines)


def generate_blueprint(root: Path) -> str:
    sections: list[tuple[Path, str, str, str]] = []
    for path in canonical_sources(root):
        raw = path.read_text(encoding="utf-8")
        frontmatter, body = split_frontmatter(raw, path)
        sections.append((path, scalar(frontmatter, "id"), sha256(raw), body.rstrip()))

    header = generated_frontmatter(
        doc_id="DOC-AURORA-PRODUCT-BLUEPRINT",
        title="Aurora Product Blueprint",
        document_type="product_blueprint_aggregate",
        version="0.2.0",
        generated_from=[item[1] for item in sections],
    )

    publication = [
        header,
        "<!-- GENERATED — DO NOT EDIT DIRECTLY",
        "Canonical sources: docs/product/blueprint/01-*.md through 15-*.md",
        "Generator: scripts/generate_docs.py",
        "-->",
        "",
        "# Aurora Product Blueprint",
        "",
        "> This publication concatenates the fifteen modular constitutional sources. ",
        "> Edit the source section, regenerate, and validate; never edit this aggregate directly.",
        "",
        "## Source manifest",
        "",
        "| Section | Canonical source | SHA-256 |",
        "|---:|---|---|",
    ]
    for path, _doc_id, digest, _body in sections:
        publication.append(
            f"| {path.name[:2]} | `{path.as_posix()}` | `{digest}` |"
        )

    publication.extend(["", "---", ""])
    for index, (path, _doc_id, _digest, body) in enumerate(sections):
        publication.extend(
            [
                f"<!-- BEGIN SOURCE: {path.as_posix()} -->",
                body,
                f"<!-- END SOURCE: {path.as_posix()} -->",
            ]
        )
        if index != len(sections) - 1:
            publication.extend(["", "---", ""])

    return "\n".join(publication).rstrip() + "\n"


def generate_roadmap(root: Path) -> str:
    path = root / ROADMAP_SOURCE
    raw = path.read_text(encoding="utf-8")
    frontmatter, body = split_frontmatter(raw, path)
    source_id = scalar(frontmatter, "id")
    source_hash = sha256(raw)

    header = generated_frontmatter(
        doc_id="DOC-AURORA-ROADMAP",
        title="Aurora Capability Roadmap",
        document_type="product_roadmap",
        version="0.2.0",
        generated_from=[source_id],
    )
    return (
        header
        + "<!-- GENERATED — DO NOT EDIT DIRECTLY\n"
        + f"Canonical source: {ROADMAP_SOURCE.as_posix()}\n"
        + f"Source SHA-256: {source_hash}\n"
        + "Generator: scripts/generate_docs.py\n"
        + "-->\n\n"
        + body.rstrip()
        + "\n"
    )


def write_output(output_root: Path, relative: Path, content: str) -> None:
    destination = output_root / relative
    destination.parent.mkdir(parents=True, exist_ok=True)
    destination.write_text(content, encoding="utf-8")


def compare_output(root: Path, relative: Path, expected: str) -> bool:
    path = root / relative
    return path.exists() and path.read_text(encoding="utf-8") == expected


def main() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument(
        "--output-root",
        type=Path,
        default=ROOT,
        help="Root under which generated relative paths are written.",
    )
    parser.add_argument(
        "--check",
        action="store_true",
        help="Do not write; fail when committed projections differ.",
    )
    args = parser.parse_args()

    blueprint = generate_blueprint(ROOT)
    roadmap = generate_roadmap(ROOT)

    if args.check:
        failures = []
        if not compare_output(ROOT, BLUEPRINT_OUTPUT, blueprint):
            failures.append(BLUEPRINT_OUTPUT.as_posix())
        if not compare_output(ROOT, ROADMAP_OUTPUT, roadmap):
            failures.append(ROADMAP_OUTPUT.as_posix())
        if failures:
            print("STALE GENERATED DOCUMENTS:")
            for failure in failures:
                print(f"- {failure}")
            return 1
        print("Generated projections are current.")
        return 0

    output_root = args.output_root.resolve()
    write_output(output_root, BLUEPRINT_OUTPUT, blueprint)
    write_output(output_root, ROADMAP_OUTPUT, roadmap)
    print(output_root / BLUEPRINT_OUTPUT)
    print(output_root / ROADMAP_OUTPUT)
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
