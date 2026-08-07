#!/usr/bin/env python3
from pathlib import Path

EVIDENCE = "DOC-AURORA-M0-R4-SPK001-EVIDENCE-RECEIPT"
REVIEW = "REVIEW-AURORA-M0-R4-SPK001-SOVEREIGN-STORE-2026-08-07"
RUN = "31213792366"
EXEC = "4242342486f512320f12e0b603f052166264c4ea"

# STATUS
p = Path("docs/tracking/STATUS.md")
t = p.read_text(encoding="utf-8")
t = t.replace("version: 0.18.0", "version: 0.19.0", 1)
anchor = "  - DOC-AURORA-M0-R4-ADR-ACCEPTANCE-SPK001-AUTHORIZATION\n"
if anchor not in t:
    raise SystemExit("STATUS relation anchor missing")
t = t.replace(anchor, anchor + f"  - {EVIDENCE}\n  - {REVIEW}\n", 1)
repls = {
    "- **Mastra posture:** PROPOSED preferred/default substrate to evaluate first for first-party agentic Harness infrastructure; NOT Sovereign Core owner": "- **Mastra posture:** ACCEPTED preferred-first substrate to evaluate for first-party agentic Harness infrastructure; NOT Sovereign Core owner",
    "- **ADR-0007 / ADR-0008:** PROPOSED / spike-blocked": "- **ADR-0007:** PROPOSED / EVIDENCE-READY — SQLite + `modernc.org/sqlite` recommended by reviewed SPK-001; operator decision pending\n- **ADR-0008:** PROPOSED / SPK-002-BLOCKED",
    "- **SPK-AURORA-M0-SOVEREIGN-STORE-001:** AUTHORIZED FOR EXECUTION — exact canonical spec at `36f46956bc275d0aec32b7e3ea4d959010fa9dcb` / blob `6ad7397d46208a0a9c762073d2c5239ceff4e056`": f"- **SPK-AURORA-M0-SOVEREIGN-STORE-001:** PASS / EVIDENCE_COMPLETE / REVIEWED / DECISION_INFORMED / CLOSED — final run `{RUN}`, execution revision `{EXEC}`, 4/4 matrix PASS",
    "- **SPK-AURORA-M0-OWNER-TRUST-002:** PROPOSED / EXECUTION NOT AUTHORIZED / depends on reviewed SPK-001 result": "- **SPK-AURORA-M0-OWNER-TRUST-002:** PROPOSED / EXECUTION NOT AUTHORIZED — sequencing dependency on reviewed SPK-001 is now satisfied, but separate authorization is still required",
    "- **R4 verdict:** BLOCKED — existing executable spike evidence and required operator decision acceptance are not yet complete": "- **R4 verdict:** BLOCKED — SPK-001 is closed; ADR-0007 operator decision and SPK-002/ADR-0008 evidence/decision remain",
    "- **Accepted technical decisions:** ADR-0003, ADR-0004, ADR-0005, ADR-0006 and ADR-0009; operational store and owner-root mechanisms remain unresolved": "- **Accepted technical decisions:** ADR-0003, ADR-0004, ADR-0005, ADR-0006 and ADR-0009; ADR-0007 is evidence-ready but not accepted; owner-root mechanism remains unresolved",
    "R4 M0 decision coverage:         15/15 mapped; ADR-0003/4/5/6 accepted; ADR-0007/8 spike-blocked": "R4 M0 decision coverage:         15/15 mapped; ADR-0003/4/5/6 accepted; ADR-0007 evidence-ready; ADR-0008 SPK-002-blocked",
    "Architecture Spike execution:   SPK-001 AUTHORIZED ONLY; SPK-002 AND ALL OTHER SPIKES PROHIBITED": "Architecture Spike execution:   SPK-001 CLOSED; SPK-002 AND ALL OTHER SPIKES NOT AUTHORIZED",
}
for old, new in repls.items():
    if old not in t:
        raise SystemExit(f"STATUS replacement missing: {old}")
    t = t.replace(old, new, 1)

old_intro = "R4 has researched concrete dispositions for all M0 mechanism questions. ADR-0003, ADR-0004, ADR-0005 and ADR-0006 are accepted; ADR-0007 and ADR-0008 remain proposed because store/atomicity and owner-root/time/restore mechanisms still require their reviewed Architecture Spike evidence:"
new_intro = "R4 has researched concrete dispositions for all M0 mechanism questions. ADR-0003, ADR-0004, ADR-0005 and ADR-0006 are accepted. SPK-001 has now proved and closed the store/atomicity/backup/migration uncertainty, making ADR-0007 evidence-ready; ADR-0008 remains proposed because owner-root/time/restore-freshness behavior still requires SPK-002 evidence:"
if old_intro not in t:
    raise SystemExit("STATUS section 5 intro missing")
t = t.replace(old_intro, new_intro, 1)

t = t.replace("Aurora Sovereign Core\n→ proposed Go runtime", "Aurora Sovereign Core\n→ accepted Go runtime", 1)
t = t.replace("Mastra\n→ proposed preferred substrate to evaluate first for first-party agentic Harnesses", "Mastra\n→ accepted preferred-first substrate to evaluate for first-party agentic Harnesses", 1)

old_block = "The operator has now accepted ADR-0003, ADR-0004, ADR-0005, ADR-0006 and ADR-0009 and separately authorized execution of the exact current `SPK-AURORA-M0-SOVEREIGN-STORE-001` specification.\n\nThe remaining R4 blockers are now narrower:\n\n1. execute and independently review `SPK-AURORA-M0-SOVEREIGN-STORE-001`;\n2. use that evidence to accept/reject/revise ADR-0007;\n3. ADR-0008 remains blocked by `SPK-AURORA-M0-OWNER-TRUST-002`, which is still NOT AUTHORIZED and remains sequenced after reviewed SPK-001 evidence.\n\nADR-0009 acceptance remains cross-horizon and does not authorize Mastra implementation."
new_block = f"The operator accepted ADR-0003, ADR-0004, ADR-0005, ADR-0006 and ADR-0009 and authorized the exact SPK-001 specification. SPK-001 has now completed successfully: final workflow `{RUN}` passed all four Ubuntu/Windows × modernc/mattn correctness cases and the required evidence receipts; the independent review closed the spike as `PASS / REVIEWED / DECISION_INFORMED`.\n\nThe reviewed store recommendation is SQLite + `database/sql` + `modernc.org/sqlite` as the initial binding baseline. ADR-0007 revision `0.2.0` remains proposed and requires operator acceptance/rejection/revision before it becomes governing.\n\nThe remaining R4 blockers are now narrower:\n\n1. operator decides evidence-informed ADR-0007;\n2. `SPK-AURORA-M0-OWNER-TRUST-002` requires separate explicit execution authorization;\n3. after SPK-002 evidence is reviewed, ADR-0008 must be accepted/rejected/revised.\n\nADR-0009 acceptance remains cross-horizon and does not authorize Mastra implementation. SPK-001 completion does not authorize SPK-002."
if old_block not in t:
    raise SystemExit("STATUS blocker block missing")
t = t.replace(old_block, new_block, 1)

old_next = "ADR-0003 / ADR-0004 / ADR-0005 / ADR-0006 / ADR-0009 accepted\n→ SPK-AURORA-M0-SOVEREIGN-STORE-001 execution explicitly authorized\n→ execute the exact disposable persistence spike against a fixed branch/revision\n→ review evidence and decide ADR-0007\n→ STOP\n→ SPK-002 remains NOT AUTHORIZED until separately approved after SPK-001 review\n→ do not begin R5"
new_next = "SPK-AURORA-M0-SOVEREIGN-STORE-001 PASS / CLOSED\n→ operator reviews ADR-0007 v0.2.0 (SQLite + modernc.org/sqlite)\n→ accept / reject / revise ADR-0007\n→ STOP\n→ SPK-AURORA-M0-OWNER-TRUST-002 remains NOT AUTHORIZED until separately approved\n→ do not begin R5"
if old_next not in t:
    raise SystemExit("STATUS next-action block missing")
t = t.replace(old_next, new_next, 1)
p.write_text(t.rstrip() + "\n", encoding="utf-8")

# RESEARCH MAP
p = Path("docs/research/RESEARCH-MAP.md")
t = p.read_text(encoding="utf-8")
t = t.replace("version: 0.4.0", "version: 0.4.1", 1)
t = t.replace(
    "documentary support for Go/local-state shape; store/driver/crash commitment still requires SPK-AURORA-M0-SOVEREIGN-STORE-001",
    "documentary support for Go/local-state shape; SPK-AURORA-M0-SOVEREIGN-STORE-001 has now closed PASS and informs ADR-0007 v0.2.0",
    1,
)
t = t.replace(
    "ADR-0009 proposes Mastra as the default agentic Harness substrate to evaluate first, not as the specification.",
    "Accepted ADR-0009 establishes Mastra as the preferred-first agentic Harness substrate to evaluate, not as the specification.",
    1,
)
t = t.replace(
    "| Sovereign Core storage and recovery | local-first stores, event/state ownership, backup/restore | crash/restart/restore spike | documentary research complete; `SPK-AURORA-M0-SOVEREIGN-STORE-001` specified, execution not authorized |",
    "| Sovereign Core storage and recovery | local-first stores, event/state ownership, backup/restore | crash/restart/restore spike | `SPK-AURORA-M0-SOVEREIGN-STORE-001` PASS/CLOSED; SQLite + modernc evidence informs ADR-0007 v0.2.0 |",
    1,
)
p.write_text(t.rstrip() + "\n", encoding="utf-8")

# WORKLOG
p = Path("docs/tracking/WORKLOG.md")
t = p.read_text(encoding="utf-8").rstrip()
entry = f"""

## 2026-08-07 — SPK-001 Sovereign Store PASS / Evidence Closeout

The authorized disposable `SPK-AURORA-M0-SOVEREIGN-STORE-001` experiment executed from canonical authorization baseline `a3192afad3dba9c6e699588c07ca2bcaac1161fd` against the fixed spike specification at `36f46956bc275d0aec32b7e3ea4d959010fa9dcb` / blob `6ad7397d46208a0a9c762073d2c5239ceff4e056`.

Final executable evidence came from branch revision `{EXEC}` and GitHub Actions run `{RUN}`. The final matrix passed all four Ubuntu/Windows × `modernc.org/sqlite`/`mattn/go-sqlite3` correctness cases and the aggregate gate reported complete required evidence receipts. Process-kill, stale/invalid transitions, WAL/checkpoint recovery, interrupted backup, fresh restore after deleting the original work directory, identity collision, corruption, unsupported schema and migration fixtures all passed.

An intermediate Windows/modernc failure was traced to the test harness using `select {{}}` as an artificial deadlock before parent kill; replacing only that blocking mechanism with timer-backed sleep eliminated the false negative, and the final expanded run passed. The finding is preserved as a resolved test-harness issue, not a database defect.

Both bindings passed correctness. The reviewed tie-break recommends SQLite + `database/sql` + `modernc.org/sqlite` because it removes CGO/C-toolchain build dependence while preserving cross-platform correctness; `mattn/go-sqlite3` remains a tested fallback. PostgreSQL expansion was not triggered.

SPK-001 is `PASS / EVIDENCE_COMPLETE / REVIEWED / DECISION_INFORMED / CLOSED`. ADR-0007 revision `0.2.0` remains proposed/evidence-ready pending operator decision. `SPK-AURORA-M0-OWNER-TRUST-002`, ADR-0008, R5, Mission Contract, Microdesign and production implementation remain unauthorized/unresolved as applicable.
"""
if "## 2026-08-07 — SPK-001 Sovereign Store PASS / Evidence Closeout" not in t:
    t += entry
p.write_text(t.rstrip() + "\n", encoding="utf-8")
