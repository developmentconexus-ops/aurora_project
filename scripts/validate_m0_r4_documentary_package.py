#!/usr/bin/env python3
from pathlib import Path
import re

ROOT = Path(__file__).resolve().parents[1]
errors = []

def read(path):
    p = ROOT / path
    if not p.exists():
        errors.append(f"missing: {path}")
        return ""
    return p.read_text(encoding="utf-8")

baseline = "d00cc1abfc2a41ac7e81e1f3478e188b3c5e9e52"

auth = read("docs/acceptance/2026-08-07-m0-r4-operator-authorization.md")
landscape = read("docs/design/M0-R4-DECISION-LANDSCAPE.md")
coverage = read("docs/capabilities/CAP-SOVEREIGN-CORE/R4-DECISION-COVERAGE.md")
adr_index = read("docs/adr/README.md")
research_map = read("docs/research/RESEARCH-MAP.md")

if baseline not in auth or baseline not in landscape or baseline not in coverage:
    errors.append("fixed R4 source baseline missing from one or more owners")

expected_questions = {
    "R4-Q-CORE-001", "R4-Q-STORE-001", "R4-Q-STATE-001", "R4-Q-SCHEMA-001",
    "R4-Q-ATOMIC-001", "R4-Q-INTEGRITY-001", "R4-Q-TIME-001", "R4-Q-AUTHN-001",
    "R4-Q-EXPORT-001", "R4-Q-MIGRATE-001", "R4-Q-AUDIT-001", "R4-Q-TELEM-001",
    "R4-Q-TOPOLOGY-001", "R4-Q-ENGINE-001", "R4-Q-RESTORE-001",
}
land_ids = set(re.findall(r"R4-Q-[A-Z]+-\d{3}", landscape))
coverage_ids = set(re.findall(r"R4-Q-[A-Z]+-\d{3}", coverage))
if not expected_questions.issubset(land_ids):
    errors.append(f"decision landscape missing questions: {sorted(expected_questions-land_ids)}")
if not expected_questions.issubset(coverage_ids):
    errors.append(f"coverage missing questions: {sorted(expected_questions-coverage_ids)}")

adrs = {
    "ADR-AURORA-0003": "docs/adr/0003-m0-go-core-runtime.md",
    "ADR-AURORA-0004": "docs/adr/0004-m0-local-state-execution-shape.md",
    "ADR-AURORA-0005": "docs/adr/0005-m0-portable-state-export.md",
    "ADR-AURORA-0006": "docs/adr/0006-m0-observability-boundary.md",
    "ADR-AURORA-0007": "docs/adr/0007-m0-sqlite-operational-store.md",
    "ADR-AURORA-0008": "docs/adr/0008-m0-owner-root-recovery-trust.md",
}
for aid, path in adrs.items():
    txt = read(path)
    if f"id: {aid}" not in txt:
        errors.append(f"wrong/missing ADR id: {path}")
    if "status: proposed" not in txt:
        errors.append(f"ADR is not proposed: {path}")
    if "status: accepted" in txt:
        errors.append(f"R4 ADR accidentally accepted: {path}")
    if aid not in adr_index:
        errors.append(f"ADR missing from index: {aid}")

adr7 = read(adrs["ADR-AURORA-0007"])
adr8 = read(adrs["ADR-AURORA-0008"])
if "SPK-AURORA-M0-SOVEREIGN-STORE-001" not in adr7 or "MUST NOT move to `accepted`" not in adr7:
    errors.append("ADR-0007 spike blocker not explicit")
if "SPK-AURORA-M0-OWNER-TRUST-002" not in adr8 or "MUST NOT move to `accepted`" not in adr8:
    errors.append("ADR-0008 spike blocker not explicit")

spikes = {
    "SPK-AURORA-M0-SOVEREIGN-STORE-001": "docs/design/SPK-AURORA-M0-SOVEREIGN-STORE-001.md",
    "SPK-AURORA-M0-OWNER-TRUST-002": "docs/design/SPK-AURORA-M0-OWNER-TRUST-002.md",
}
for sid, path in spikes.items():
    txt = read(path)
    if f"id: {sid}" not in txt:
        errors.append(f"wrong/missing spike id: {path}")
    if "status: proposed" not in txt:
        errors.append(f"spike is not proposed: {path}")
    if "Execution: NOT AUTHORIZED" not in txt:
        errors.append(f"spike execution boundary missing: {path}")
    if "DISCARD" not in txt:
        errors.append(f"spike disposal rule missing: {path}")

spk2 = read(spikes["SPK-AURORA-M0-OWNER-TRUST-002"])
if "Dependency: SPK-AURORA-M0-SOVEREIGN-STORE-001 reviewed result" not in spk2:
    errors.append("SPK-002 sequencing dependency missing")

reports = [
    ("RESEARCH-AURORA-M0-RUNTIME-PERSISTENCE-R4-V1", "docs/research/AURORA-RESEARCH-M0-RUNTIME-PERSISTENCE-R4-v1.md", "docs/research/AURORA-RESEARCH-M0-RUNTIME-PERSISTENCE-R4-v1.sources.json"),
    ("RESEARCH-AURORA-M0-PORTABILITY-INTEGRITY-R4-V1", "docs/research/AURORA-RESEARCH-M0-PORTABILITY-INTEGRITY-R4-v1.md", "docs/research/AURORA-RESEARCH-M0-PORTABILITY-INTEGRITY-R4-v1.sources.json"),
    ("RESEARCH-AURORA-M0-OWNER-AUTHORITY-RECOVERY-R4-V1", "docs/research/AURORA-RESEARCH-M0-OWNER-AUTHORITY-RECOVERY-R4-v1.md", "docs/research/AURORA-RESEARCH-M0-OWNER-AUTHORITY-RECOVERY-R4-v1.sources.json"),
    ("RESEARCH-AURORA-M0-OBSERVABILITY-R4-V1", "docs/research/AURORA-RESEARCH-M0-OBSERVABILITY-R4-v1.md", "docs/research/AURORA-RESEARCH-M0-OBSERVABILITY-R4-v1.sources.json"),
]
for rid, report, manifest in reports:
    rt = read(report); mt = read(manifest)
    if f"id: {rid}" not in rt:
        errors.append(f"research report id missing: {report}")
    if f'"research_id": "{rid}"' not in mt:
        errors.append(f"research manifest id mismatch: {manifest}")
    if rid not in research_map:
        errors.append(f"research report missing from map: {rid}")
    if '"selection_policy"' not in mt or '"limitations"' not in mt:
        errors.append(f"research manifest incomplete: {manifest}")

required_coverage_phrases = [
    "R4 BLOCKED", "SPIKE_REQUIRED / EXECUTION_NOT_AUTHORIZED",
    "PROPOSED_ADR_READY / OPERATOR_REVIEW_REQUIRED", "INTENTIONAL_NON_SELECTION",
]
for phrase in required_coverage_phrases:
    if phrase not in coverage:
        errors.append(f"coverage missing disposition phrase: {phrase}")

# Nothing in the documentary package may claim implementation or spike execution authorization.
for path in [p for _,p in adrs.items()] + [p for _,p in spikes.items()] + [r[1] for r in reports]:
    txt = read(path)
    if "Implementation: AUTHORIZED" in txt or "Execution: AUTHORIZED" in txt:
        errors.append(f"unauthorized execution/implementation claim: {path}")

if errors:
    print("M0 R4 documentary package validation FAIL")
    for e in errors:
        print("-", e)
    raise SystemExit(1)

print("M0 R4 documentary package validation PASS")
print("R4 questions expected: 15")
print("R4 questions covered: 15")
print("focused R4 reports: 4")
print("source manifests: 4")
print("proposed ADRs: 6")
print("accepted new ADRs: 0")
print("proposed spikes: 2")
print("authorized spike executions: 0")
print("documentary-ready ADRs: 4")
print("spike-blocked ADRs: 2")
print("current readiness implication: R4 BLOCKED")
