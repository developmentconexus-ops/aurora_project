#!/usr/bin/env python3
from __future__ import annotations

import argparse
import json
import re
import shutil
import subprocess
import sys
import sys
import tempfile
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
LINK_RE = re.compile(r"(?<!!)\[[^\]]*\]\(([^)]+)\)")
REQ_RE = re.compile(r"\bAUR-REQ-[A-Z0-9-]+-\d{3}\b")
SOURCE_RE = re.compile(r"\[(S\d{2,})\]")
PROGRAM_OWNER_RE = re.compile(r"(?m)^program_status_authority:\s*true\s*$")
ALLOWED_ROOT = {".github", "docs", "scripts", "README.md", "AGENTS.md", "CONTRIBUTING.md"}
REQUIRED = [
    "README.md", "AGENTS.md", "docs/index.md", "docs/roadmap.md",
    "docs/product/README.md", "docs/product/PRODUCT-BLUEPRINT.md",
    "docs/product/blueprint/15-documentation-research-governance.md",
    "docs/development/engineering-rules.md", "docs/development/planning-readiness.md",
    "docs/development/capability-realization.md", "docs/decisions/index.md",
    "docs/architecture/index.md",
]
FINAL_FORBIDDEN = ["docs/work", "docs/superpowers", "docs/tracking", "docs/acceptance", "docs/reviews", "docs/adr", "docs/design", "docs/DOCUMENTATION-MAP.md"]


def md_files(root: Path):
    return sorted(x for x in root.rglob("*.md") if ".git" not in x.parts)


def frontmatter(text: str) -> str | None:
    if not text.startswith("---\n"):
        return None
    end = text.find("\n---\n", 4)
    return None if end < 0 else text[4:end]


def scalar(fm: str, key: str) -> str | None:
    m = re.search(rf"(?m)^{re.escape(key)}:\s*([^\n]+)\s*$", fm)
    return None if not m else m.group(1).strip().strip('"').strip("'")


def list_field(fm: str, key: str) -> list[str]:
    m = re.search(rf"(?m)^{re.escape(key)}:\s*\n(?P<body>(?:  - .*\n)*)", fm)
    if not m:
        return []
    return [line[4:].strip() for line in m.group("body").splitlines() if line.startswith("  - ")]


def guard_bootstrap(root: Path, errors: list[str]):
    total = sum((root / x).stat().st_size for x in ["AGENTS.md", "docs/index.md", "docs/roadmap.md"] if (root / x).exists())
    if total > 20 * 1024:
        errors.append(f"bootstrap budget exceeded: {total} bytes > 20480")


def guard_program_owner(root: Path, errors: list[str]):
    owners = []
    for path in md_files(root):
        if PROGRAM_OWNER_RE.search(path.read_text(encoding="utf-8")):
            owners.append(path.relative_to(root).as_posix())
    if owners != ["docs/roadmap.md"]:
        errors.append(f"program status authority must be only docs/roadmap.md; found {owners}")


def guard_root_allowlist(root: Path, errors: list[str]):
    roadmap = root / "docs/roadmap.md"
    if not roadmap.exists() or "Product/runtime implementation: BLOCKED" not in roadmap.read_text(encoding="utf-8"):
        return
    unexpected = sorted(x.name for x in root.iterdir() if x.name not in ALLOWED_ROOT and x.name != ".git")
    if unexpected:
        errors.append("implementation-blocked root surface not allowlisted: " + ", ".join(unexpected))


def guard_final_tree(root: Path, errors: list[str]):
    for rel in FINAL_FORBIDDEN:
        if (root / rel).exists():
            errors.append(f"final candidate forbidden legacy/temporary surface exists: {rel}")


def guard_review_changed_files(files: list[str], errors: list[str]):
    if files != ["docs/work/current/ai-dialog.md"]:
        errors.append(f"review isolation violation: {files}")


def validate_links(root: Path, errors: list[str]):
    # Current authority/router surfaces only; Evidence/phase/history snapshots may preserve historical paths.
    current_roots = ["README.md", "AGENTS.md", "CONTRIBUTING.md", "docs/product", "docs/architecture", "docs/decisions", "docs/development", "docs/capabilities", "docs/reference"]
    candidates = []
    for rel in current_roots:
        path = root / rel
        if path.is_file(): candidates.append(path)
        elif path.is_dir(): candidates += list(path.rglob("*.md"))
    for path in sorted(set(candidates)):
        text = path.read_text(encoding="utf-8")
        for raw in LINK_RE.findall(text):
            target = raw.strip().split()[0].strip("<>")
            if not target or target.startswith(("http://", "https://", "mailto:", "#")):
                continue
            file_part = target.split("#", 1)[0]
            if not file_part:
                continue
            resolved = (path.parent / file_part).resolve()
            try:
                resolved.relative_to(root.resolve())
            except ValueError:
                errors.append(f"{path.relative_to(root)} link escapes repo: {target}")
                continue
            if not resolved.exists():
                errors.append(f"{path.relative_to(root)} broken link: {target}")


def validate_ids(root: Path, errors: list[str]):
    seen = {}
    for path in md_files(root / "docs"):
        fm = frontmatter(path.read_text(encoding="utf-8"))
        if fm is None:
            continue
        doc_id = scalar(fm, "id")
        if not doc_id:
            continue
        if doc_id in seen:
            errors.append(f"duplicate doc id {doc_id}: {seen[doc_id]} and {path.relative_to(root)}")
        else:
            seen[doc_id] = path.relative_to(root)
    valid = set(seen)
    for path in md_files(root / "docs"):
        fm = frontmatter(path.read_text(encoding="utf-8"))
        if fm is None: continue
        for target in list_field(fm, "related"):
            if target and target not in valid:
                errors.append(f"{path.relative_to(root)} unresolved related id {target}")


def validate_research(root: Path, errors: list[str], stats: dict):
    manifests = list((root / "docs/research").glob("*.sources.json")) if (root / "docs/research").exists() else []
    source_count = 0
    for manifest in manifests:
        try: data = json.loads(manifest.read_text(encoding="utf-8"))
        except Exception as exc:
            errors.append(f"{manifest.relative_to(root)} invalid JSON: {exc}")
            continue
        sources = data.get("sources", [])
        if not sources:
            errors.append(f"{manifest.relative_to(root)} empty sources")
            continue
        ids = []
        for item in sources:
            sid = item.get("id") if isinstance(item, dict) else None
            if not sid:
                errors.append(f"{manifest.relative_to(root)} source missing id")
                continue
            ids.append(sid)
            for field in ("title", "url", "publisher", "type", "accessed_at"):
                if not item.get(field): errors.append(f"{manifest.relative_to(root)} source {sid} missing {field}")
        if len(ids) != len(set(ids)): errors.append(f"{manifest.relative_to(root)} duplicate source ids")
        source_count += len(ids)
        report = manifest.with_name(manifest.name.replace(".sources.json", ".md"))
        if report.exists():
            used = set(SOURCE_RE.findall(report.read_text(encoding="utf-8")))
            missing = used - set(ids)
            if missing: errors.append(f"{report.relative_to(root)} undefined sources: {sorted(missing)}")
    stats["source_manifests"] = len(manifests)
    stats["research_sources"] = source_count


def validate_requirements(root: Path, errors: list[str], stats: dict):
    path = root / "docs/product/REQUIREMENTS-TRACEABILITY.md"
    if not path.exists(): return
    ids = set(REQ_RE.findall(path.read_text(encoding="utf-8")))
    stats["requirements"] = len(ids)
    if len(ids) < 200: errors.append(f"requirements traceability unexpectedly small: {len(ids)}")


def validate_generated(root: Path, generated_root: Path, errors: list[str]):
    expected = generated_root / "docs/product/PRODUCT-BLUEPRINT.md"
    actual = root / "docs/product/PRODUCT-BLUEPRINT.md"
    if not expected.exists() or not actual.exists() or expected.read_bytes() != actual.read_bytes():
        errors.append("stale generated Product Blueprint")
    if (generated_root / "docs/roadmap.md").exists():
        errors.append("generator must never emit docs/roadmap.md")


def validate_router(root: Path, errors: list[str]):
    text = (root / "docs/index.md").read_text(encoding="utf-8") if (root / "docs/index.md").exists() else ""
    for needle in ["roadmap.md", "product/README.md", "architecture/index.md", "decisions/index.md", "development/planning-readiness.md", "development/capability-realization.md", "development/engineering-rules.md"]:
        if needle not in text: errors.append(f"docs/index missing required route: {needle}")


def validate(root: Path, generated_root: Path, final: bool) -> tuple[list[str], dict]:
    errors = []
    stats = {}
    for rel in REQUIRED:
        if not (root / rel).exists(): errors.append(f"missing required path: {rel}")
    if (root / "docs/superpowers").exists(): errors.append("docs/superpowers must not exist after MR-01 cutover")
    if (root / "docs/tracking").exists(): errors.append("docs/tracking must not exist after MR-01 cutover")
    if (root / "docs/DOCUMENTATION-MAP.md").exists(): errors.append("legacy Documentation Map path must not exist")
    guard_bootstrap(root, errors)
    guard_program_owner(root, errors)
    guard_root_allowlist(root, errors)
    if final: guard_final_tree(root, errors)
    validate_router(root, errors)
    validate_links(root, errors)
    validate_ids(root, errors)
    validate_research(root, errors, stats)
    validate_requirements(root, errors, stats)
    validate_generated(root, generated_root, errors)
    stats["markdown_files"] = len(md_files(root))
    return errors, stats


def self_test() -> int:
    failures = []
    with tempfile.TemporaryDirectory() as td:
        r = Path(td)
        (r/"docs").mkdir()
        (r/"AGENTS.md").write_text("a")
        (r/"docs/index.md").write_text("i")
        (r/"docs/roadmap.md").write_text("program_status_authority: true\nProduct/runtime implementation: BLOCKED")
        errs=[]; guard_bootstrap(r, errs)
        if errs: failures.append("bootstrap positive")
        (r/"AGENTS.md").write_text("x"*(21*1024))
        errs=[]; guard_bootstrap(r, errs)
        if not errs: failures.append("bootstrap negative")
    errs=[]; guard_review_changed_files(["docs/work/current/ai-dialog.md"], errs)
    if errs: failures.append("review isolation positive")
    errs=[]; guard_review_changed_files(["docs/work/current/ai-dialog.md", "README.md"], errs)
    if not errs: failures.append("review isolation negative")
    with tempfile.TemporaryDirectory() as td:
        r=Path(td); (r/"docs/work").mkdir(parents=True)
        errs=[]; guard_final_tree(r, errs)
        if not errs: failures.append("final-tree negative")
    with tempfile.TemporaryDirectory() as td:
        r=Path(td); (r/"docs").mkdir(); (r/"docs/roadmap.md").write_text("program_status_authority: true\nProduct/runtime implementation: BLOCKED")
        (r/"src").mkdir(); errs=[]; guard_root_allowlist(r, errs)
        if not errs: failures.append("blocked-root allowlist negative")
    with tempfile.TemporaryDirectory() as td:
        out=Path(td)
        rc=subprocess.run([sys.executable, str(ROOT/"scripts/generate_docs.py"), "--output-root", str(out)], cwd=ROOT).returncode
        if rc or (out/"docs/roadmap.md").exists(): failures.append("generator roadmap isolation")
    if failures:
        print("SELF-TEST FAIL:", ", ".join(failures)); return 1
    print("Documentation guard self-tests PASS")
    return 0


def main() -> int:
    ap=argparse.ArgumentParser()
    ap.add_argument("--generated-root", type=Path)
    ap.add_argument("--report", type=Path, default=ROOT/"docs-validation-report.json")
    ap.add_argument("--final", action="store_true")
    ap.add_argument("--self-test", action="store_true")
    ap.add_argument("--review-base")
    args=ap.parse_args()
    if args.self_test: return self_test()
    if args.review_base:
        changed=subprocess.check_output(["git","diff","--name-only",f"{args.review_base}...HEAD"],cwd=ROOT,text=True).splitlines()
        errors=[]; guard_review_changed_files(changed, errors)
        if errors:
            print("\n".join(errors)); return 1
        print("Review isolation PASS"); return 0
    if not args.generated_root:
        print("--generated-root required", file=sys.stderr); return 2
    errors,stats=validate(ROOT,args.generated_root.resolve(),args.final)
    payload={"status":"PASS" if not errors else "FAIL","errors":errors,"statistics":stats,"final_mode":args.final}
    args.report.parent.mkdir(parents=True,exist_ok=True); args.report.write_text(json.dumps(payload,indent=2)+"\n")
    if errors:
        print("DOCUMENTATION VALIDATION FAIL")
        for e in errors: print("-",e)
        return 1
    print("DOCUMENTATION VALIDATION PASS", json.dumps(stats,sort_keys=True))
    return 0


if __name__ == "__main__": raise SystemExit(main())
