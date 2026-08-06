#!/usr/bin/env python3
"""Validate the Aurora documentation system without external dependencies."""

from __future__ import annotations

import argparse
import json
import re
import sys
from dataclasses import dataclass
from pathlib import Path
from typing import Iterable

ROOT = Path(__file__).resolve().parents[1]
NORMATIVE_AUTHORITIES = {
    "constitutional",
    "decision",
    "specification",
    "contract",
    "standard",
    "policy",
}
PLACEHOLDER_RE = re.compile(r"\b(?:TBD|TODO|FIXME|XXX)\b")
MARKDOWN_LINK_RE = re.compile(r"(?<!!)\[[^\]]*\]\(([^)]+)\)")
SOURCE_REF_RE = re.compile(r"\[(S\d{2,})\]")
REQ_ID_RE = re.compile(r"\bAUR-REQ-[A-Z0-9-]+-\d{3}\b")
DOC_ID_RE = re.compile(
    r"^(?:DOC|ADR|RESEARCH|DESIGN|PLAN|REVIEW|HISTORY)-[A-Z0-9-]+$"
)


@dataclass
class Document:
    path: Path
    text: str
    frontmatter_text: str
    body: str
    fields: dict[str, object]

    @property
    def doc_id(self) -> str | None:
        value = self.fields.get("id")
        return value if isinstance(value, str) else None

    @property
    def authority(self) -> str:
        value = self.fields.get("authority")
        return value if isinstance(value, str) else ""


def split_frontmatter(text: str, path: Path) -> tuple[str, str]:
    if not text.startswith("---\n"):
        raise ValueError(f"{path}: missing frontmatter")
    marker = text.find("\n---\n", 4)
    if marker < 0:
        raise ValueError(f"{path}: unterminated frontmatter")
    return text[4:marker], text[marker + 5 :].lstrip("\n")


def parse_simple_yaml(frontmatter: str, path: Path) -> dict[str, object]:
    """Parse Aurora's constrained top-level scalar/list frontmatter."""
    result: dict[str, object] = {}
    current_list: str | None = None
    for line_number, raw in enumerate(frontmatter.splitlines(), start=1):
        if not raw.strip() or raw.lstrip().startswith("#"):
            continue
        if raw.startswith("  - "):
            if current_list is None:
                raise ValueError(f"{path}:{line_number}: list item without key")
            value = raw[4:].strip().strip('"').strip("'")
            target = result[current_list]
            if not isinstance(target, list):
                raise ValueError(f"{path}:{line_number}: invalid list state")
            target.append(value)
            continue
        if raw.startswith(" "):
            raise ValueError(f"{path}:{line_number}: nested frontmatter is unsupported")
        if ":" not in raw:
            raise ValueError(f"{path}:{line_number}: malformed frontmatter line")
        key, value = raw.split(":", 1)
        key = key.strip()
        value = value.strip()
        current_list = None
        if not value:
            result[key] = []
            current_list = key
        elif value == "[]":
            result[key] = []
        elif value == "{}":
            result[key] = {}
        elif value.lower() in {"null", "~"}:
            result[key] = None
        elif value.lower() in {"true", "false"}:
            result[key] = value.lower() == "true"
        else:
            result[key] = value.strip('"').strip("'")
    return result


def iter_markdown_files(root: Path) -> Iterable[Path]:
    for path in sorted(root.rglob("*.md")):
        if ".git" not in path.parts:
            yield path


def load_documents(root: Path, errors: list[str]) -> list[Document]:
    docs: list[Document] = []
    candidates = [
        root / "AGENTS.md",
        root / "CONTRIBUTING.md",
        root / "README.md",
        *iter_markdown_files(root / "docs"),
    ]
    for path in candidates:
        if not path.exists():
            continue
        text = path.read_text(encoding="utf-8")
        if not text.startswith("---\n"):
            if path.name not in {"README.md", "AGENTS.md", "CONTRIBUTING.md"}:
                errors.append(f"{path.relative_to(root)}: canonical Markdown missing frontmatter")
            continue
        try:
            frontmatter, body = split_frontmatter(text, path)
            fields = parse_simple_yaml(frontmatter, path)
        except ValueError as exc:
            errors.append(str(exc))
            continue
        docs.append(Document(path, text, frontmatter, body, fields))
    return docs


def validate_expected_structure(root: Path, errors: list[str]) -> None:
    slugs = [
        "product-vision",
        "human-aurora-relationship",
        "domain-world-model",
        "cognitive-lifecycle-journeys",
        "capability-system",
        "memory-knowledge-context",
        "harness-orchestration",
        "interaction-multimodality-presence",
        "tools-devices-laboratory",
        "autonomy-authority-safety",
        "security-privacy-sovereignty",
        "system-architecture",
        "reliability-observability-evaluation",
        "capability-roadmap",
        "documentation-research-governance",
    ]
    expected = [
        *[
            root / "docs/product/blueprint" / f"{index:02d}-{slug}.md"
            for index, slug in enumerate(slugs, start=1)
        ],
        root / "docs/product/CAPABILITY-REALIZATION-METHOD.md",
        root / "docs/product/REQUIREMENTS-TRACEABILITY.md",
        root / "docs/product/PRODUCT-BLUEPRINT.md",
        root / "docs/history/2026-08-05-aurora-origin-and-discovery-record.md",
        root / "docs/tracking/DOCUMENTATION-COVERAGE.md",
        root / "docs/reviews/2026-08-05-a0-adversarial-documentation-review.md",
        root / "docs/roadmap.md",
    ]
    for path in expected:
        if not path.exists():
            errors.append(f"missing expected file: {path.relative_to(root)}")


def validate_metadata(docs: list[Document], errors: list[str]) -> dict[str, Document]:
    by_id: dict[str, Document] = {}
    required = {"id", "title", "document_type", "authority", "status", "owners"}
    for doc in docs:
        rel = doc.path.relative_to(ROOT)
        missing = sorted(required - doc.fields.keys())
        if missing:
            errors.append(f"{rel}: missing metadata fields: {', '.join(missing)}")
        doc_id = doc.doc_id
        if not doc_id:
            errors.append(f"{rel}: missing document id")
        elif not DOC_ID_RE.match(doc_id):
            errors.append(f"{rel}: invalid document id format: {doc_id}")
        elif doc_id in by_id:
            errors.append(
                f"duplicate document id {doc_id}: {by_id[doc_id].path.relative_to(ROOT)} and {rel}"
            )
        else:
            by_id[doc_id] = doc
        owners = doc.fields.get("owners")
        if not isinstance(owners, list) or not owners:
            errors.append(f"{rel}: owners must be a non-empty list")
        if doc.authority in NORMATIVE_AUTHORITIES and "source_of_truth_for" not in doc.fields:
            errors.append(f"{rel}: normative document missing source_of_truth_for")
    return by_id


def load_manifest_ids(root: Path, errors: list[str]) -> set[str]:
    result: set[str] = set()
    for path in sorted((root / "docs/research").glob("*.sources.json")):
        try:
            data = json.loads(path.read_text(encoding="utf-8"))
        except json.JSONDecodeError:
            continue
        manifest_id = data.get("id")
        if not isinstance(manifest_id, str):
            research_id = data.get("research_id")
            if isinstance(research_id, str):
                manifest_id = f"{research_id}-SOURCES"
            else:
                errors.append(
                    f"{path.relative_to(root)}: manifest requires id or research_id"
                )
                continue
        if manifest_id in result:
            errors.append(f"{path.relative_to(root)}: duplicate manifest id {manifest_id}")
        else:
            result.add(manifest_id)
    return result


def validate_relations(docs: list[Document], valid_ids: set[str], errors: list[str]) -> None:
    for doc in docs:
        related = doc.fields.get("related", [])
        if related is None:
            continue
        if not isinstance(related, list):
            errors.append(f"{doc.path.relative_to(ROOT)}: related must be a list")
            continue
        for target in related:
            if not isinstance(target, str):
                errors.append(f"{doc.path.relative_to(ROOT)}: non-string related target")
            elif target not in valid_ids:
                errors.append(f"{doc.path.relative_to(ROOT)}: unresolved related id {target}")


def validate_links(root: Path, docs: list[Document], errors: list[str]) -> None:
    for doc in docs:
        for raw_target in MARKDOWN_LINK_RE.findall(doc.body):
            target = raw_target.strip().split()[0].strip("<>")
            if not target or target.startswith(("http://", "https://", "mailto:", "#")):
                continue
            file_part = target.split("#", 1)[0]
            if not file_part:
                continue
            resolved = (doc.path.parent / file_part).resolve()
            try:
                resolved.relative_to(root.resolve())
            except ValueError:
                errors.append(f"{doc.path.relative_to(root)}: link escapes repository: {target}")
                continue
            if not resolved.exists():
                errors.append(f"{doc.path.relative_to(root)}: broken local link: {target}")


def validate_placeholders(docs: list[Document], errors: list[str]) -> None:
    for doc in docs:
        if doc.authority not in NORMATIVE_AUTHORITIES:
            continue
        for line_number, line in enumerate(doc.body.splitlines(), start=1):
            match = PLACEHOLDER_RE.search(line)
            if not match:
                continue
            token = match.group(0)
            if f"`{token}`" in line:
                continue
            lower = line.lower()
            if any(term in lower for term in ("placeholder", "unresolved", "scan", "prohibited")):
                continue
            errors.append(
                f"{doc.path.relative_to(ROOT)} body line {line_number}: unresolved placeholder: {line.strip()}"
            )


def validate_source_manifests(root: Path, errors: list[str], stats: dict[str, int]) -> None:
    research_dir = root / "docs/research"
    manifest_count = 0
    source_count = 0
    for manifest in sorted(research_dir.glob("*.sources.json")):
        manifest_count += 1
        try:
            data = json.loads(manifest.read_text(encoding="utf-8"))
        except json.JSONDecodeError as exc:
            errors.append(f"{manifest.relative_to(root)}: invalid JSON: {exc}")
            continue
        sources = data.get("sources")
        if not isinstance(sources, list) or not sources:
            errors.append(f"{manifest.relative_to(root)}: sources must be non-empty")
            continue
        ids: list[str] = []
        for source in sources:
            if not isinstance(source, dict):
                errors.append(f"{manifest.relative_to(root)}: source must be object")
                continue
            source_id = source.get("id")
            if not isinstance(source_id, str):
                errors.append(f"{manifest.relative_to(root)}: source missing id")
                continue
            ids.append(source_id)
            for field in ("title", "url", "publisher", "type", "accessed_at"):
                if not source.get(field):
                    errors.append(f"{manifest.relative_to(root)}: source {source_id} missing {field}")
        if len(ids) != len(set(ids)):
            errors.append(f"{manifest.relative_to(root)}: duplicate source ids")
        source_count += len(ids)

        report = manifest.with_name(manifest.name.replace(".sources.json", ".md"))
        if not report.exists():
            errors.append(f"{manifest.relative_to(root)}: matching report missing: {report.name}")
            continue
        referenced = set(SOURCE_REF_RE.findall(report.read_text(encoding="utf-8")))
        defined = set(ids)
        missing = sorted(referenced - defined)
        unused = sorted(defined - referenced)
        if missing:
            errors.append(f"{report.relative_to(root)}: undefined source references: {', '.join(missing)}")
        if unused:
            errors.append(f"{report.relative_to(root)}: manifest sources not cited: {', '.join(unused)}")
    stats["source_manifests"] = manifest_count
    stats["research_sources"] = source_count


def validate_requirements(root: Path, errors: list[str], stats: dict[str, int]) -> None:
    path = root / "docs/product/REQUIREMENTS-TRACEABILITY.md"
    if not path.exists():
        return
    ids = REQ_ID_RE.findall(path.read_text(encoding="utf-8"))
    counts: dict[str, int] = {}
    for req_id in ids:
        counts[req_id] = counts.get(req_id, 0) + 1
    duplicates = sorted(req_id for req_id, count in counts.items() if count > 1)
    if duplicates:
        errors.append(f"{path.relative_to(root)}: duplicate requirement IDs: {', '.join(duplicates[:20])}")
    stats["requirements"] = len(counts)
    if len(counts) < 200:
        errors.append(f"{path.relative_to(root)}: expected comprehensive traceability, found {len(counts)} requirements")


def validate_coverage(root: Path, errors: list[str]) -> None:
    path = root / "docs/tracking/DOCUMENTATION-COVERAGE.md"
    if not path.exists():
        return
    text = path.read_text(encoding="utf-8")
    gaps = re.findall(r"\|[^\n]*\|\s*(?:MISSING|UNMAPPED|OPEN_GAP)\s*\|", text)
    if gaps:
        errors.append(f"{path.relative_to(root)}: unresolved coverage gaps: {len(gaps)}")


def validate_generated(root: Path, generated_root: Path, errors: list[str]) -> None:
    for relative in (Path("docs/product/PRODUCT-BLUEPRINT.md"), Path("docs/roadmap.md")):
        expected = generated_root / relative
        actual = root / relative
        if not expected.exists():
            errors.append(f"generated artifact missing from validation input: {relative}")
        elif not actual.exists():
            errors.append(f"committed generated projection missing: {relative}")
        elif expected.read_bytes() != actual.read_bytes():
            errors.append(f"stale generated projection: {relative}")


def write_report(path: Path, errors: list[str], warnings: list[str], stats: dict[str, int]) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    payload = {
        "status": "PASS" if not errors else "FAIL",
        "errors": errors,
        "warnings": warnings,
        "statistics": stats,
    }
    path.write_text(json.dumps(payload, indent=2, ensure_ascii=False) + "\n", encoding="utf-8")


def main() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument("--generated-root", type=Path, required=True)
    parser.add_argument("--report", type=Path, default=ROOT / "docs-validation-report.json")
    args = parser.parse_args()

    errors: list[str] = []
    warnings: list[str] = []
    stats: dict[str, int] = {}

    validate_expected_structure(ROOT, errors)
    docs = load_documents(ROOT, errors)
    stats["canonical_documents"] = len(docs)
    by_id = validate_metadata(docs, errors)
    manifest_ids = load_manifest_ids(ROOT, errors)
    stats["document_ids"] = len(by_id)
    stats["manifest_ids"] = len(manifest_ids)
    validate_relations(docs, set(by_id) | manifest_ids, errors)
    validate_links(ROOT, docs, errors)
    validate_placeholders(docs, errors)
    validate_source_manifests(ROOT, errors, stats)
    validate_requirements(ROOT, errors, stats)
    validate_coverage(ROOT, errors)
    validate_generated(ROOT, args.generated_root.resolve(), errors)

    write_report(args.report.resolve(), errors, warnings, stats)
    print(json.dumps({"status": "PASS" if not errors else "FAIL", **stats}, indent=2))
    if errors:
        print("\nValidation errors:", file=sys.stderr)
        for error in errors:
            print(f"- {error}", file=sys.stderr)
        return 1
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
