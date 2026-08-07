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
print("patched R5 helper expectations for current Spec")
