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

replacements = {
    'dec = replace_once(dec, "version: 0.4.0", "version: 0.5.0", dec_path)':
        'dec = replace_once(dec, "version: 0.5.0", "version: 0.6.0", dec_path)',
    'anchor = "| D-056 | Mastra is the accepted preferred-first substrate to evaluate for first-party agentic Harnesses while sovereign truth/authority remain Aurora-owned | ADR-0009 | accepted |"\ninsert = anchor + "\\n| D-057 | CAP-SOVEREIGN-CORE Requirements v0.1.1, Spec v0.2.0, Threat Model v0.2.0 and Test Plan v0.2.0 are the accepted R4-aligned M0 A2 package | CAP-SOVEREIGN-CORE A2 documents + R5 operator acceptance | accepted |\\n| D-058 | MIS-M0-SOVEREIGN-CORE-001 v0.1.0 is the approved first scoped M0 Mission Contract | MIS-M0-SOVEREIGN-CORE-001 + R5 operator acceptance | approved |"':
        'anchor = "| D-058 | M0 owner trust uses a random wrapped ORK + authenticated external generation/time high-water with fail-closed restore/revalidation semantics | ADR-0008 | accepted |"\ninsert = anchor + "\\n| D-059 | CAP-SOVEREIGN-CORE Requirements v0.1.1, Spec v0.2.0, Threat Model v0.2.0 and Test Plan v0.2.0 are the accepted R4-aligned M0 A2 package | CAP-SOVEREIGN-CORE A2 documents + R5 operator acceptance | accepted |\\n| D-060 | MIS-M0-SOVEREIGN-CORE-001 v0.1.0 is the approved first scoped M0 Mission Contract | MIS-M0-SOVEREIGN-CORE-001 + R5 operator acceptance | approved |"',
    'o15 = "| O-015 | exact first Mission Contract for selected M0 | M0 ACRM R5 Mission Contract |\\n"':
        'o15 = "| O-015 | exact first Mission Contract for selected M0 | `MIS-M0-SOVEREIGN-CORE-001` v0.1.0 proposed in R5; operator approval pending |\\n"',
    'old_para = "`O-015` previously combined milestone selection and first Contract. The milestone portion is resolved by `D-051`; only the exact Mission Contract remains open and cannot be chosen before its applicable gates.\\n\\n"':
        'old_para = "`O-015` previously combined milestone selection and first Contract. The milestone portion is resolved by `D-051`. R5 now proposes `MIS-M0-SOVEREIGN-CORE-001` v0.1.0; the decision remains open until explicit operator approval.\\n\\n"',
}
for old_text, new_text in replacements.items():
    count = text.count(old_text)
    if count != 1:
        raise SystemExit(f"decision-index patch target count={count}: {old_text[:140]}")
    text = text.replace(old_text, new_text, 1)

path.write_text(text.rstrip() + "\n", encoding="utf-8")
print("patched R5 closeout helper for current contract and decision-index lifecycle wording")
