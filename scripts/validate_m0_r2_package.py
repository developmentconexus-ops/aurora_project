#!/usr/bin/env python3
from __future__ import annotations

import re
from collections import Counter
from pathlib import Path

REQ_PATH = Path('docs/capabilities/CAP-SOVEREIGN-CORE/REQUIREMENTS.md')
APP_PATH = Path('docs/capabilities/CAP-SOVEREIGN-CORE/APPLICABILITY.md')
COV_PATH = Path('docs/capabilities/CAP-SOVEREIGN-CORE/R2-COVERAGE.md')

req_text = REQ_PATH.read_text(encoding='utf-8')
app_text = APP_PATH.read_text(encoding='utf-8')
cov_text = COV_PATH.read_text(encoding='utf-8')

errors: list[str] = []

req_re = re.compile(
    r'^\| `(?P<id>CAP-SOVEREIGN-CORE-REQ-\d{3})` \| '
    r'(?P<statement>.*?) \| (?P<risk>critical|high|medium) \| '
    r'(?P<verification>.*?) \| (?P<sources>.*?) \|$', re.MULTILINE
)
reqs = list(req_re.finditer(req_text))
ids = [m.group('id') for m in reqs]
expected_ids = [f'CAP-SOVEREIGN-CORE-REQ-{i:03d}' for i in range(1, 123)]
if ids != expected_ids:
    errors.append(f'requirement IDs are not exactly sequential 001-122: got {len(ids)} rows')
if len(set(ids)) != len(ids):
    errors.append('duplicate derived requirement IDs')

allowed_verification = {
    'STATIC_ANALYSIS', 'SCHEMA_VALIDATION', 'UNIT_TEST', 'CONTRACT_TEST',
    'CONFORMANCE', 'INTEGRATION', 'FAULT_INJECTION', 'SECURITY_TEST',
    'SIMULATION', 'HIL', 'PHYSICAL_DRILL', 'BENCHMARK/EVAL',
    'USER_JOURNEY', 'DOCUMENT_REVIEW', 'OPERATOR_VERDICT'
}
modal_re = re.compile(r'\b(?:MUST|MUST NOT|SHOULD|MAY)\b')
source_re = re.compile(r'`(AUR-REQ-[A-Z0-9-]+-\d{3})`')

all_req_sources: list[str] = []
for m in reqs:
    rid = m.group('id')
    statement = m.group('statement')
    if not modal_re.search(statement):
        errors.append(f'{rid}: no normative modal')
    verifications = {v.strip() for v in m.group('verification').split(',') if v.strip()}
    if not verifications:
        errors.append(f'{rid}: no verification direction')
    bad = sorted(verifications - allowed_verification)
    if bad:
        errors.append(f'{rid}: unsupported verification methods {bad}')
    sources = source_re.findall(m.group('sources'))
    if not sources:
        errors.append(f'{rid}: no constitutional source')
    all_req_sources.extend(sources)

active_re = re.compile(
    r'^\| `(AUR-REQ-[A-Z0-9-]+-\d{3})` \| `(APPLIES|PARTIALLY_APPLIES)` \|',
    re.MULTILINE,
)
active_rows = active_re.findall(app_text)
active_ids = [sid for sid, _ in active_rows]
active_set = set(active_ids)
if len(active_ids) != 127 or len(active_set) != 127:
    errors.append(f'R1 active set expected 127 unique rows, got {len(active_ids)} rows/{len(active_set)} unique')

inactive_refs = sorted(set(all_req_sources) - active_set)
if inactive_refs:
    errors.append(f'derived requirements cite inactive R1 sources: {inactive_refs}')
uncovered = sorted(active_set - set(all_req_sources))
if uncovered:
    errors.append(f'R1 active sources not cited by requirements: {uncovered}')

cov_re = re.compile(
    r'^\| `(AUR-REQ-[A-Z0-9-]+-\d{3})` \| `(APPLIES|PARTIALLY_APPLIES)` \| (?P<reqs>.*?) \|$',
    re.MULTILINE,
)
cov_rows = list(cov_re.finditer(cov_text))
cov_sources = [m.group(1) for m in cov_rows]
if len(cov_sources) != 127 or len(set(cov_sources)) != 127:
    errors.append(f'coverage matrix expected 127 unique rows, got {len(cov_sources)} rows/{len(set(cov_sources))} unique')
if set(cov_sources) != active_set:
    errors.append('coverage matrix source set differs from R1 active set')

coverage_links = 0
for m in cov_rows:
    sid = m.group(1)
    mapped = re.findall(r'`(CAP-SOVEREIGN-CORE-REQ-\d{3})`', m.group('reqs'))
    coverage_links += len(mapped)
    if not mapped:
        errors.append(f'{sid}: coverage row has no derived requirement')
    unknown = sorted(set(mapped) - set(ids))
    if unknown:
        errors.append(f'{sid}: coverage maps unknown requirements {unknown}')
if coverage_links != 409:
    errors.append(f'expected 409 source->requirement coverage links, got {coverage_links}')

risk_counts = Counter(m.group('risk') for m in reqs)
expected_risk = Counter({'critical': 51, 'high': 53, 'medium': 18})
if risk_counts != expected_risk:
    errors.append(f'risk distribution changed unexpectedly: {dict(risk_counts)}')

for token in ('TBD', 'TODO', 'FIXME', 'XXX'):
    if re.search(rf'\b{token}\b', req_text):
        errors.append(f'placeholder present in requirements: {token}')

# Candidate implementation technologies must not be selected by R2. Generic
# architectural terms and explicitly negative patterns such as event sourcing are allowed.
forbidden = re.compile(
    r'\b(?:PostgreSQL|SQLite|Redis|Kafka|NATS|gRPC|Kubernetes|Docker|Mastra|LangGraph|OpenHands|TypeScript|Rust)\b',
    re.IGNORECASE,
)
for m in reqs:
    hit = forbidden.search(m.group('statement'))
    if hit:
        errors.append(f"{m.group('id')}: candidate technology appears in normative statement: {hit.group(0)}")

if 'status: proposed' not in req_text:
    errors.append('requirements package must remain proposed at R2')
if 'source_revision: 495b712142d7c3d722da2298f7a0b060707f9f5e' not in req_text:
    errors.append('requirements fixed source revision drifted')
if 'source_revision: 495b712142d7c3d722da2298f7a0b060707f9f5e' not in cov_text:
    errors.append('coverage fixed source revision drifted')

if errors:
    print('M0 R2 package validation FAILED')
    for error in errors:
        print(f'- {error}')
    raise SystemExit(1)

print('M0 R2 package validation PASS')
print(f'derived requirements: {len(reqs)}')
print(f'risk distribution: critical={risk_counts["critical"]}, high={risk_counts["high"]}, medium={risk_counts["medium"]}')
print(f'R1 active sources: {len(active_set)}')
print(f'coverage rows: {len(cov_rows)}')
print(f'source->requirement links: {coverage_links}')
print('uncovered active sources: 0')
print('inactive source references: 0')
print('unsupported verification methods: 0')
print('candidate technology selections: 0')
