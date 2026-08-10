from pathlib import Path

p = Path('docs/capabilities/CAP-SOVEREIGN-CORE/M0-R6-MICRODESIGN.md')
s = p.read_text(encoding='utf-8')

def once(old, new):
    global s
    c = s.count(old)
    if c != 1:
        raise SystemExit(f'expected one match, found {c}: {old[:120]!r}')
    s = s.replace(old, new, 1)

once(
'''Allowed:\n\n```text\ncmd/adapters → application → domain + ports\nadapters → ports/domain types where required\n```\n''',
'''Allowed:\n\n```text\ncmd/adapters → application → domain + ports\nadapters → ports/domain types where required\n```\n\n`application` may receive the accepted non-authoritative standard observability APIs (`*slog.Logger`, OTel `trace.Tracer`, OTel `metric.Meter`) from the composition root. Exporter/provider construction remains in `adapters/observability`; domain packages never depend on observability.\n''')

once(
'''Accepted governing mutations include required audit/evidence references in the same SQLite transaction. Rejected/non-governing attempts may be recorded without advancing governing state.\n''',
'''The mutation/result structs named above are defined in `ports/state.go` and contain domain values plus operation IDs/preconditions only; they contain no SQL rows, statements or driver-specific types. Accepted governing mutations include required audit/evidence references in the same SQLite transaction. Rejected/non-governing attempts may be recorded without advancing governing state.\n''')

once(
'''Required because ADR-0008 mandates a physically independent trust boundary.\n\n```go\ntype OwnerTrustStore interface {\n''',
'''Required because ADR-0008 mandates a physically independent trust boundary. `RootEnvelope` and `Anchor` are port contract structs declared in `ports/owner_trust.go`; neither exposes the plaintext ORK.\n\n```go\ntype OwnerTrustStore interface {\n''')

once(
'''--data-dir\n→ AURORA_DATA_DIR\n→ platform user config directory / Aurora\n''',
'''--data-dir\n→ AURORA_DATA_DIR\n→ ~/.aurora\n''')

once(
'''authority_revision INTEGER PRIMARY KEY\npredecessor_revision INTEGER NULL\nauthority_state_json TEXT NOT NULL\nrevalidation_required INTEGER NOT NULL CHECK(revalidation_required IN (0,1))\nchanged_by TEXT NOT NULL\nchanged_at TEXT NOT NULL\n''',
'''authority_revision INTEGER PRIMARY KEY\npredecessor_revision INTEGER NULL\nauthority_state_json TEXT NOT NULL\nchanged_by TEXT NOT NULL\nchanged_at TEXT NOT NULL\n''')

old = '''The descriptor contains only governing heads:\n\n```json\n{\n  "version": 1,\n  "aurora_id": "AUR-...",\n  "governing_generation": 14,\n  "projects": [\n    {"project_id":"PRJ-...","state_revision":3}\n  ],\n  "authority_revision": 5\n}\n```\n\nProject entries are sorted by `project_id`. The logical object is JCS-canonicalized and authenticated once with HMAC-SHA-256 using the ORK-derived governing-state key.\n\nM0 does **not** add per-table HMACs, Merkle trees, hash chains, event-sourced integrity or a custom transaction protocol.\n'''
new = '''The descriptor is one compact **current governing logical snapshot**, not a list of physical SQLite bytes and not merely revision pointers. It binds:\n\n- immutable Aurora/owner identity binding required by M0;\n- `governing_generation`;\n- every Project's current canonical metadata, current revision number and current `StateEnvelope`/proposed-next-action/acceptance attribution;\n- the complete current logical `authority.State`, including revalidation status.\n\nConceptually:\n\n```json\n{\n  "version": 1,\n  "aurora": {"aurora_id":"AUR-...","owner_operator_id":"OWNER-..."},\n  "governing_generation": 14,\n  "projects": [\n    {\n      "project_id":"PRJ-...",\n      "display_label":"...",\n      "current_state_revision":3,\n      "current_state": {"schema_version":"1","kind":"...","summary":"...","payload":{}},\n      "proposed_next_action": null,\n      "accepted_by_actor":"...",\n      "accepted_at":"..."\n    }\n  ],\n  "authority": {"revision":5,"revalidation_required":false,"grants":[]}\n}\n```\n\nProject entries and any set-like fields use stable documented ordering before JCS. The logical object is JCS-canonicalized and authenticated **once** with HMAC-SHA-256 using the ORK-derived governing-state key. Therefore changing the contents of the current Project or authority revision without the ORK invalidates governing integrity even if revision numbers are left unchanged.\n\nM0 does **not** add per-table/per-row HMACs, Merkle trees, hash chains, event-sourced integrity or a custom transaction protocol.\n'''
once(old, new)

once(
'''7. on Unix, open and sync the parent directory;\n8. on Windows, record that temp-file Sync + replace is the supported M0 compatibility behavior and do not claim stronger directory-flush semantics than R7 evidence demonstrates.\n''',
'''7. on Unix, open and sync the parent directory;\n8. on Windows, after temp-file `Sync` + close, replace the target with `MoveFileExW` using `MOVEFILE_REPLACE_EXISTING | MOVEFILE_WRITE_THROUGH` through `golang.org/x/sys/windows`; do not claim Unix-style parent-directory fsync semantics on Windows beyond the R7 evidence.\n''')

once(
'''Logical content is validated by `schemas/sovereign-export-v1.schema.json`, canonicalized with JCS for digest input and protected externally with age.\n''',
'''Logical content is validated by `schemas/sovereign-export-v1.schema.json`. `integrity.payload_sha256` is SHA-256 over the JCS-canonical JSON of the export document with the top-level `integrity` member omitted; verification reconstructs that same canonical input before accepting the package. The complete document is then protected externally with age.\n''')

once(
'''- `golang.org/x/crypto` where required by ADR-0008;\n''',
'''- `golang.org/x/crypto` where required by ADR-0008;\n- `golang.org/x/sys/windows` for the Windows trust-file replacement primitive;\n''')

anchor = '''## 24. R6 self-review checklist\n'''
coverage = '''## 24. Mission-criterion design coverage\n\n| Mission criterion | Primary design realization | Vertical proof slice |\n|---|---|---|\n| `CRIT-001` Stable identity/scope | §§7.1, 9, 11 | Slices 1–2 |\n| `CRIT-002` Canonical state ownership | §§7.2, 9–10 | Slices 2–3 |\n| `CRIT-003` Revision-bound transitions | §§7.2, 8.1, 13 | Slice 3 |\n| `CRIT-004` Authority/next-safe-action | §§7.3, 11, 13 | Slices 4–5 |\n| `CRIT-005` Fresh-process recovery | §§9–13 | Slice 6 |\n| `CRIT-006` Export/restore/migration | §§14–15 | Slices 7–8 |\n| `CRIT-007` Security/sovereignty/secrets | §§10–14, 17 | Slices 1, 4–5, 7, 10 |\n| `CRIT-008` Audit/evidence/telemetry | §§9.3, 17, 19 | Slice 9 |\n| `CRIT-009` Reliability/fault containment | §§12–13, 20, 23 | Slices 5–6, 10 |\n| `CRIT-010` Architecture guards | §§2–6, 21–22 | all slices / static review |\n| `CRIT-011` Documentation/traceability | §§1–2, 21–25 | R6/R7 reviews |\n| `CRIT-012` Complete M0 Golden Proof | §§18–20 | Slice 11 |\n\nThe task-by-task Implementation Plan created after written Microdesign approval must expand this to explicit requirement/test allocations; this table proves no Mission criterion is structurally orphaned at the design level.\n\n## 25. R6 self-review checklist\n'''
once(anchor, coverage)

once('''## 25. Written-review gate\n''', '''## 26. Written-review gate\n''')

p.write_text(s.rstrip() + '\n', encoding='utf-8')
print('patched R6 microdesign self-review findings')
