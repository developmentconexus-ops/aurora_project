from pathlib import Path

path = Path("scripts/tmp_prepare_m0_r5.py")
text = path.read_text(encoding="utf-8")

repls = {
    'replace_once(spec, "  - proposed R3 reusable behavior and logical architecture for CAP-SOVEREIGN-CORE", "  - reusable behavior and logical architecture for CAP-SOVEREIGN-CORE")':
    'replace_once(spec, "  - proposed reusable behavior and logical design of CAP-SOVEREIGN-CORE", "  - reusable behavior and logical design of CAP-SOVEREIGN-CORE")\nreplace_once(spec, "  - proposed M0 sovereign Core domain and lifecycle semantics", "  - M0 sovereign Core domain and lifecycle semantics")',

    'replace_once(spec, "  - DOC-AURORA-CAP-SOVEREIGN-CORE-R3-COVERAGE\\nsource_revision:", "  - DOC-AURORA-CAP-SOVEREIGN-CORE-R3-COVERAGE\\n  - DOC-AURORA-CAP-SOVEREIGN-CORE-R4-DECISION-COVERAGE\\n  - ADR-AURORA-0003\\n  - ADR-AURORA-0004\\n  - ADR-AURORA-0005\\n  - ADR-AURORA-0006\\n  - ADR-AURORA-0007\\n  - ADR-AURORA-0008\\n  - DOC-AURORA-MIS-M0-SOVEREIGN-CORE-001\\nsource_revision:")':
    'replace_once(spec, "  - ADR-AURORA-0002\\nsource_revision:", "  - ADR-AURORA-0002\\n  - DOC-AURORA-CAP-SOVEREIGN-CORE-R4-DECISION-COVERAGE\\n  - ADR-AURORA-0003\\n  - ADR-AURORA-0004\\n  - ADR-AURORA-0005\\n  - ADR-AURORA-0006\\n  - ADR-AURORA-0007\\n  - ADR-AURORA-0008\\n  - DOC-AURORA-MIS-M0-SOVEREIGN-CORE-001\\nsource_revision:")',

    'replace_once(spec, "This R3 specification fixes the reusable product behavior, domain semantics, state/lifecycle rules, ownership boundaries, authority rules, security/recovery behavior and evidence model needed by M0 without selecting the implementation mechanisms that belong to R4.", "The original R3 revision fixed reusable product behavior and deliberately left implementation mechanisms to R4. R4 has since reached PASS. This v0.2.0 preserves the R3 semantics and binds them to accepted ADR-0003..0008 while leaving source layout, Go API shape, SQL DDL, filesystem wrappers and test implementation to R6.")':
    'replace_once(spec, "This R3 specification fixes **logical semantics and boundaries**. It does not select implementation language, process topology, database/storage engine, state-versus-event persistence pattern, serialization format, event transport, telemetry backend, backup technology, migration tooling, durable workflow engine or UI technology. Those remain R4 decisions where applicable.", "The original R3 revision fixed reusable product behavior and deliberately left implementation mechanisms to R4. R4 has since reached PASS. This v0.2.0 preserves the R3 semantics and binds them to accepted ADR-0003..0008 while leaving source layout, Go API shape, SQL DDL, filesystem wrappers and test implementation to R6.")',
}

for old, new in repls.items():
    count = text.count(old)
    if count != 1:
        raise SystemExit(f"patch target count={count}: {old[:120]}")
    text = text.replace(old, new, 1)

path.write_text(text.rstrip() + "\n", encoding="utf-8")

coverage_path = Path("docs/capabilities/CAP-SOVEREIGN-CORE/R5-COVERAGE.md")
coverage = coverage_path.read_text(encoding="utf-8")
placeholder = '"+table+"'
if coverage.count(placeholder) != 1:
    raise SystemExit(f"R5 coverage placeholder count={coverage.count(placeholder)}")

def criterion_for(req: int) -> int:
    if 1 <= req <= 9:
        return 1
    if 10 <= req <= 20:
        return 2
    if 21 <= req <= 30:
        return 3
    if req == 31:
        return 12
    if 32 <= req <= 45:
        return 4
    if 46 <= req <= 55:
        return 5
    if 56 <= req <= 66:
        return 6
    if 67 <= req <= 76:
        return 7
    if 77 <= req <= 88:
        return 8
    if 89 <= req <= 95:
        return 9
    if 96 <= req <= 107:
        return 10
    if 108 <= req <= 122:
        return 11
    raise ValueError(req)

rows = []
for req in range(1, 123):
    crit = criterion_for(req)
    rows.append(
        f"| `CAP-SOVEREIGN-CORE-REQ-{req:03d}` | "
        f"`MIS-M0-SOVEREIGN-CORE-001-CRIT-{crit:03d}` | "
        "`R3-COVERAGE.md` + `TEST-PLAN.md` | `ALLOCATED` |"
    )
coverage = coverage.replace(placeholder, "\n".join(rows), 1)
coverage_path.write_text(coverage.rstrip() + "\n", encoding="utf-8")

print("patched R5 helper expectations and materialized 122/122 coverage rows")
