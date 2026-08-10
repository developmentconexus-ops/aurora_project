from pathlib import Path

path = Path("scripts/tmp_close_m0_r5.py")
text = path.read_text(encoding="utf-8")

old = '''    append_related(fm, ACCEPTANCE)\n    ensure_approver(fm)\n    for old, new in body_replacements:\n'''
new = '''    append_related(fm, ACCEPTANCE)\n    ensure_approver(fm)\n    if path.endswith("MIS-M0-SOVEREIGN-CORE-001.md"):\n        expected = "  - proposed exact implementation commitment for M0 Sovereign Core Walking Skeleton"\n        replacement = "  - exact scoped implementation commitment for M0 Sovereign Core Walking Skeleton"\n        if expected not in fm:\n            raise SystemExit(f"{path}: contract source_of_truth proposal wording missing")\n        fm[fm.index(expected)] = replacement\n    for old, new in body_replacements:\n'''
if text.count(old) != 1:
    raise SystemExit("promote patch anchor mismatch")
text = text.replace(old, new, 1)

old_tuple = '''        (\n            "  - proposed exact implementation commitment for M0 Sovereign Core Walking Skeleton",\n            "  - exact scoped implementation commitment for M0 Sovereign Core Walking Skeleton",\n        ),\n'''
if text.count(old_tuple) != 1:
    raise SystemExit("contract frontmatter body replacement tuple mismatch")
text = text.replace(old_tuple, "", 1)

path.write_text(text.rstrip() + "\n", encoding="utf-8")
print("patched R5 closeout helper for contract frontmatter lifecycle wording")
