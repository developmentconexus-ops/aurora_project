#!/usr/bin/env python3
from pathlib import Path

path = Path("docs/tracking/STATUS.md")
text = path.read_text(encoding="utf-8")
text = text.replace("version: 0.11.0", "version: 0.11.1", 1)
old = """start a fresh repository-only R0 review against canonical revision d0ddfb794296e599ac96bb73bf3772937d371bf9
→ read AGENTS.md and STATUS from that revision
→ execute M0 ACRM R0 only
→ re-run M0 ACRM R0 only
→ produce R0 PASS | FAIL | BLOCKED
→ stop before R1 unless R1 is separately authorized"""
new = """start a fresh repository-only R0 review from current canonical `main`
→ resolve and record the exact `main` HEAD as the fixed R0 target revision before reading scope sources
→ read AGENTS.md and STATUS from that exact revision
→ execute M0 ACRM R0 only
→ produce R0 PASS | FAIL | BLOCKED
→ stop before R1 unless R1 is separately authorized"""
if old not in text:
    raise SystemExit("expected stale R0 target sequence not found")
text = text.replace(old, new, 1)
path.write_text(text, encoding="utf-8")

Path(".github/workflows/fix-m0-r0-rerun-target.yml").unlink()
Path("scripts/fix_m0_r0_rerun_target.py").unlink()
