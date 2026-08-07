#!/usr/bin/env python3
from pathlib import Path
import re

ROOT = Path(__file__).resolve().parents[1]
BASE = ROOT / "docs/capabilities/CAP-SOVEREIGN-CORE"
REQ = BASE / "REQUIREMENTS.md"
SPEC = BASE / "SPEC.md"
THREAT = BASE / "THREAT-MODEL.md"
TEST = BASE / "TEST-PLAN.md"
COV = BASE / "R3-COVERAGE.md"
FRESH = ROOT / "docs/reviews/2026-08-07-m0-r3-research-freshness-review.md"

errors = []
for p in [REQ, SPEC, THREAT, TEST, COV, FRESH]:
    if not p.exists():
        errors.append(f"missing required R3 artifact: {p.relative_to(ROOT)}")

if errors:
    print("\n".join(errors))
    raise SystemExit(1)

req = REQ.read_text(encoding="utf-8")
spec = SPEC.read_text(encoding="utf-8")
threat = THREAT.read_text(encoding="utf-8")
test = TEST.read_text(encoding="utf-8")
cov = COV.read_text(encoding="utf-8")
fresh = FRESH.read_text(encoding="utf-8")

expected_req = {f"CAP-SOVEREIGN-CORE-REQ-{i:03d}" for i in range(1,123)}
req_ids = set(re.findall(r"CAP-SOVEREIGN-CORE-REQ-\d{3}", req))
if req_ids != expected_req:
    errors.append(f"requirements set mismatch missing={sorted(expected_req-req_ids)} extra={sorted(req_ids-expected_req)}")

cov_rows = re.findall(r"^\| `(CAP-SOVEREIGN-CORE-REQ-\d{3})` \| (.+?) \| (.+?) \| `ALLOCATED` \|$", cov, re.M)
if len(cov_rows) != 122:
    errors.append(f"expected 122 R3 coverage rows, found {len(cov_rows)}")
coverage_ids = [r[0] for r in cov_rows]
if set(coverage_ids) != expected_req or len(set(coverage_ids)) != 122:
    errors.append("R3 coverage requirement IDs are missing/duplicated/extra")

plan_test_ids = set(re.findall(r"`(T-[A-Z]+-\d{3})`", test))
referenced_tests = set()
for rid, mechanism, test_alloc in cov_rows:
    if not mechanism.strip():
        errors.append(f"{rid}: empty Spec/mechanism allocation")
    if not any(token in mechanism for token in ["SPEC", "THREAT"]):
        errors.append(f"{rid}: mechanism allocation does not point to SPEC/THREAT: {mechanism}")
    ids = set(re.findall(r"T-[A-Z]+-\d{3}", test_alloc))
    if not ids:
        errors.append(f"{rid}: no planned test allocation")
    referenced_tests |= ids
missing_tests = referenced_tests - plan_test_ids
if missing_tests:
    errors.append(f"coverage references undefined tests: {sorted(missing_tests)}")

required_spec_headings = [
    "## 2. Use cases",
    "## 3. Goals",
    "## 4. Non-goals",
    "## 5. Applicability and requirement baseline",
    "## 7. Domain model",
    "## 8. Logical architecture and ownership boundaries",
    "## 9. Semantic command contracts",
    "## 10. Lifecycle and state model",
    "## 11. Context and memory boundary",
    "## 12. Authority and effects boundary",
    "## 13. Security, privacy and data classification",
    "## 14. Failure and recovery model",
    "## 15. Observability, audit and evidence",
    "## 16. Evaluation and evidence model",
    "## 17. Compatibility and migration",
    "## 18. Rollout and graduation",
    "## 19. Open R4 decisions and uncertainty classes",
    "## 20. Requirement coverage",
    "## 21. Owner and reviewer",
]
for heading in required_spec_headings:
    if heading not in spec:
        errors.append(f"SPEC missing required section: {heading}")

required_spec_semantics = [
    "state_schema_version",
    "state_kind",
    "state_summary",
    "state_payload",
    "AuthorityGrantRecord",
    "AuthorityState",
    "AuthoritySnapshot",
    "NextSafeActionProjection",
    "REVALIDATION_REQUIRED",
    "authenticated `OperatorIdentityRef`",
    "restored grant MUST NOT authorize its own revalidation",
    "Time Source Port",
    "Integrity Port",
    "INVALID_TRANSITION",
    "AUTHORITY_INVALID",
    "STATE_CORRUPT",
    "VERSION_INCOMPATIBLE",
    "RESTORE_UNSAFE",
    "OPERATION_AMBIGUOUS",
]
for phrase in required_spec_semantics:
    if phrase not in spec:
        errors.append(f"SPEC missing required semantic closure: {phrase}")

owners = [
    "Identity Module",
    "Project State Module",
    "Authority Module",
    "Portability / Recovery Module",
    "Audit / Evidence Module",
]
for owner in owners:
    if owner not in spec:
        errors.append(f"SPEC missing logical owner: {owner}")

required_threats = [
    "canonical Project state is corrupted or tampered",
    "old backup predating revocation is restored",
    "identity collision",
    "clock rolls back or is stale",
    "crash occurs around state mutation",
    "ordinary Project content contains policy-like instructions",
    "migration silently changes domain meaning",
]
for phrase in required_threats:
    if phrase not in threat:
        errors.append(f"Threat model missing required class: {phrase}")

required_tests = [
    "T-TRANS-002",
    "T-AUTH-004",
    "T-AUTH-005",
    "T-AUTH-012",
    "T-PORT-002",
    "T-PORT-003",
    "T-PORT-005",
    "T-PORT-007",
    "T-PORT-010",
    "T-PORT-011",
    "T-PORT-013",
    "T-REC-004",
    "T-SEC-001",
    "T-REL-001",
]
for tid in required_tests:
    if tid not in plan_test_ids:
        errors.append(f"Test plan missing critical negative case: {tid}")

for name, text in [("SPEC", spec), ("THREAT", threat), ("TEST", test), ("COVERAGE", cov), ("FRESHNESS", fresh)]:
    if re.search(r"\b(?:TBD|TODO|FIXME|XXX)\b", text):
        errors.append(f"{name}: placeholder token present")

# Guard against R3 selecting named implementation candidates. These names may exist in research,
# but the normative R3 package must not make them implementation choices.
normative = "\n".join([spec, threat, test])
selected_patterns = [
    r"(?:selected|chosen|use|adopt)\s+(?:PostgreSQL|SQLite|Temporal|DBOS|Restate|Cedar|OPA|OpenTelemetry|CloudEvents|gRPC|JSON Schema)\b",
    r"(?:database|storage engine|runtime|telemetry backend|event transport)\s*:\s*(?:PostgreSQL|SQLite|Temporal|DBOS|Restate|OpenTelemetry|CloudEvents|gRPC)",
]
for pat in selected_patterns:
    if re.search(pat, normative, re.I):
        errors.append(f"R3 package appears to select a technology: pattern={pat}")

if "R3_RESEARCH_FRESHNESS: SUFFICIENT" not in fresh:
    errors.append("research freshness review missing sufficient disposition")
if "REQUIRES_R4_REVALIDATION" not in fresh:
    errors.append("research freshness review missing explicit R4 revalidation boundary")

if errors:
    print("M0 R3 package validation FAIL")
    for e in errors:
        print(f"- {e}")
    raise SystemExit(1)

print("M0 R3 package validation PASS")
print(f"R2 requirements: {len(expected_req)}")
print(f"R3 coverage rows: {len(cov_rows)}")
print(f"planned test IDs defined: {len(plan_test_ids)}")
print(f"test IDs referenced by coverage: {len(referenced_tests)}")
print(f"undefined referenced tests: {len(missing_tests)}")
print("required ACRM Spec sections: present")
print("required logical owners: present")
print("required semantic closure: present")
print("required threat classes: present")
print("critical negative test cases: present")
print("candidate technology selections: 0")
print("research freshness disposition: sufficient for R3; R4 revalidation required")
