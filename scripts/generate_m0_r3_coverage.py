#!/usr/bin/env python3
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
OUT = ROOT / "docs/capabilities/CAP-SOVEREIGN-CORE/R3-COVERAGE.md"

spec = {}
tests = {}

def a(nums, s, t):
    for n in nums:
        if n in spec:
            raise SystemExit(f"duplicate allocation: {n}")
        spec[n] = s
        tests[n] = t

# Identity/scope
a([1], "SPEC §7.1, §9.1, §10.1", "T-ID-001, T-ID-002, T-REC-001")
a([2], "SPEC §7.3, §10.2", "T-ID-004, T-PORT-004")
a([3], "SPEC §7.1, §7.3, §9.6, §9.8", "T-ID-002, T-PORT-004, T-ID-003")
a([4], "SPEC §3–4, §7.2", "T-SCOPE-002")
a([5], "SPEC §1, §7.5, §17.1", "T-ID-005, T-ARCH-003")
a([6], "SPEC §7.1–7.3, §8, §11", "T-ID-005")
a([7], "SPEC §4, §11", "T-SCOPE-001")
a([8], "SPEC §7.10, §8.7, §9.5", "T-SCOPE-002, T-ID-005")
a([9], "SPEC §4–5, §19", "T-SCOPE-001, T-DOC-005")
# Ownership/state
a([10], "SPEC §8.1–8.4", "T-ARCH-001, T-STATE-001")
a([11], "SPEC §7–8, §11", "T-STATE-002, T-ID-005")
a([12], "SPEC §7.4, §7.8–7.10", "T-STATE-003")
a([13], "SPEC §7.3", "T-STATE-001, T-STATE-004")
a([14], "SPEC §7.3, §7.12", "T-STATE-004")
a([15], "SPEC §7, §15.1", "T-STATE-004, T-EVID-008")
a([16], "SPEC §7.9–7.10", "T-STATE-006")
a([17], "SPEC §7, §8.8, §17", "T-STATE-005, T-ARCH-003")
a([18], "SPEC §8", "T-ARCH-001")
a([19], "SPEC §7.11–7.12, §15", "T-EVID-004, T-EVID-005, T-ARCH-001")
a([20], "SPEC §8, §11", "T-STATE-002, T-ID-005")
# Transitions
a([21], "SPEC §7.5, §9.3", "T-EVID-001")
a([22], "SPEC §7.5, §9.3", "T-TRANS-002")
a([23], "SPEC §7.4, §9.3", "T-TRANS-001")
a([24], "SPEC §9.3, §10.4", "T-TRANS-001, T-TRANS-004")
a([25], "SPEC §9.3, §10.4", "T-TRANS-002, T-TRANS-003, T-TRANS-004")
a([26], "SPEC §8.1, §9.3, §14.1", "T-TRANS-006, T-TRANS-008")
a([27], "SPEC §10.3", "T-TRANS-002, T-TRANS-007")
a([28], "SPEC §9, §15", "T-TRANS-008, T-REC-001, T-PORT-004")
a([29], "SPEC §8, §15", "T-EVID-001, T-STATE-002")
a([30], "SPEC §9.5", "T-STATE-001, T-AUTH-002, T-REC-005")
a([31], "SPEC §2, §9–10, §18", "T-ID-001, T-REC-001, T-TRANS-002, T-PORT-004")
# Authority
a([32], "SPEC §7.2, §7.7, §12", "T-AUTH-001, T-AUTH-011")
a([33], "SPEC §7.7, §12", "T-AUTH-001")
a([34], "SPEC §7.7–7.8", "T-AUTH-002, T-AUTH-006")
a([35], "SPEC §7.7, §12", "T-AUTH-003")
a([36], "SPEC §7.9, §10.5", "T-AUTH-004, T-AUTH-012")
a([37], "SPEC §7.7, §10.5", "T-AUTH-005, T-PORT-006")
a([38], "SPEC §7.9", "T-AUTH-002, T-AUTH-010")
a([39], "SPEC §7.10", "T-AUTH-007, T-AUTH-002")
a([40], "SPEC §12", "T-SCOPE-001, T-AUTH-001")
a([41], "SPEC §9.6, §10.6", "T-REC-001, T-AUTH-004, T-AUTH-005")
a([42], "SPEC §9.8, §13; THREAT §7", "T-PORT-005, T-PORT-006")
a([43], "SPEC §9.3–9.4, §12; THREAT §7", "T-SEC-001")
a([44], "SPEC §7.9, §14", "T-AUTH-008")
a([45], "SPEC §14; THREAT §7", "T-AUTH-008, T-REC-004")
# Recovery
a([46], "SPEC §8.8, §9.6, §10.6", "T-REC-001")
a([47], "SPEC §9.6, §11", "T-REC-002")
a([48], "SPEC §9.6, §10.6", "T-REC-001")
a([49], "SPEC §9.6, §14", "T-REC-003, T-REC-004")
a([50], "SPEC §14.1", "T-REC-007")
a([51], "SPEC §14", "T-REC-003, T-REC-004")
a([52], "SPEC §14.2", "T-REL-001, T-REL-002")
a([53], "SPEC §14.2", "T-REL-001, T-REL-002")
a([54], "SPEC §7.14, §9.6", "T-REC-005")
a([55], "SPEC §14.3", "T-REC-006")
# Portability
a([56], "SPEC §7.13, §9.7, §17", "T-PORT-001, T-PORT-004")
a([57], "SPEC §7.13, §17.1", "T-PORT-001, T-PORT-003")
a([58], "SPEC §7.13, §9.7; THREAT §7", "T-PORT-002")
a([59], "SPEC §9.7–9.8, §17", "T-PORT-001, T-PORT-004")
a([60], "SPEC §9.8, §17.2; THREAT §7", "T-PORT-002, T-PORT-003, T-PORT-005")
a([61], "SPEC §9.8, §17", "T-PORT-004")
a([62], "SPEC §9.8; THREAT §7", "T-PORT-007")
a([63], "SPEC §9.7, §13", "T-PORT-012")
a([64], "SPEC §9.9, §17.3", "T-PORT-009")
a([65], "SPEC §9.9, §17.2–17.3", "T-PORT-003, T-PORT-010")
a([66], "SPEC §7.14, §9.9, §15.3", "T-PORT-008, T-PORT-009")
# Security
a([67], "SPEC §8.8, §11, §13.2", "T-SEC-005, T-PORT-012")
a([68], "SPEC §13.1; THREAT §3", "T-SEC-002")
a([69], "SPEC §13.1, §15.1; THREAT §7", "T-SEC-003")
a([70], "SPEC §7.5, §7.11, §15.2", "T-SEC-004, T-EVID-001")
a([71], "SPEC §11–12; THREAT §7", "T-SEC-001")
a([72], "SPEC §9.5, §13.2, §15.2", "T-STATE-001, T-AUTH-002, T-PORT-012")
a([73], "SPEC §10.3, §13.1", "T-STATE-003, T-TRANS-007")
a([74], "SPEC §14; THREAT §7", "T-REC-004, T-AUTH-008")
a([75], "THREAT §§1–12", "T-SEC-006")
a([76], "SPEC §13.1, §15.1; THREAT §7", "T-SEC-003, T-EVID-008")
# Evidence/telemetry
a([77], "SPEC §7.11, §15.1–15.2", "T-EVID-001, T-EVID-002, T-EVID-003")
a([78], "SPEC §15.2", "T-EVID-001")
a([79], "SPEC §15.2", "T-EVID-002")
a([80], "SPEC §15.2", "T-EVID-003")
a([81], "SPEC §7.11, §15.4", "T-EVID-006, T-STATE-002")
a([82], "SPEC §7.12, §15.3", "T-EVID-005")
a([83], "SPEC §7.12, §15.3", "T-EVID-004")
a([84], "SPEC §15.1–15.3", "T-EVID-008")
a([85], "SPEC §15.4", "T-EVID-007")
a([86], "SPEC §15.4", "T-EVID-006")
a([87], "SPEC §7.12–7.13, §15.3", "T-PORT-002, T-EVID-004")
a([88], "SPEC §16, §18", "T-REL-004")
# Reliability
a([89], "SPEC §16, §18", "T-REL-003, T-REC-001, T-SEC-006")
a([90], "SPEC §14.1", "T-REC-007")
a([91], "SPEC §14, §16; TEST §§7–13", "T-TRANS-002, T-AUTH-004, T-AUTH-005, T-PORT-002, T-PORT-003, T-REL-001")
a([92], "SPEC §15.3, §18", "T-REL-003")
a([93], "SPEC §15.3, §18", "T-REL-003")
a([94], "SPEC §14, §16, §21", "T-DOC-006")
a([95], "SPEC §8, §19", "T-ARCH-001, T-ARCH-003")
# Open decision guards
a([96], "SPEC §8, §19", "T-ARCH-001, T-ARCH-002")
a([97], "SPEC §19", "T-ARCH-002")
a(range(98, 105), "SPEC §19", "T-ARCH-003")
a([105], "SPEC §19", "T-ARCH-004")
a([106], "SPEC §11, §19", "T-ARCH-005")
a([107], "SPEC §8, §17.4, §19", "T-ARCH-006")
# Documentation/readiness
a([108], "SPEC §5, §8, §21", "T-DOC-001, T-ARCH-001")
a([109], "SPEC §1, §5, §21", "T-DOC-001")
a([110], "SPEC §5, §20", "T-DOC-002, T-DOC-005")
a([111], "SPEC §20", "T-DOC-002")
a([112], "SPEC §19, §21", "T-ARCH-003, T-DOC-005")
a([113], "SPEC §1, §5, §21", "T-DOC-002, T-DOC-005")
a([114], "SPEC §22", "T-DOC-004, T-DOC-007")
a([115], "SPEC §21", "T-DOC-005")
a([116], "SPEC §5, §20–21", "T-DOC-001")
a([117], "SPEC §20–21", "T-DOC-006")
a([118], "SPEC §5, §20", "T-DOC-007")
a([119], "SPEC §21–22", "T-DOC-004")
a([120], "SPEC §1, §4, §22", "T-DOC-005")
a([121], "SPEC §1, §22", "T-DOC-004")
a([122], "SPEC §8, §20, §22", "T-DOC-002, T-ARCH-003")

expected = set(range(1, 123))
actual = set(spec)
if actual != expected:
    raise SystemExit(f"allocation mismatch missing={sorted(expected-actual)} extra={sorted(actual-expected)}")

rows = []
for n in range(1, 123):
    rid = f"CAP-SOVEREIGN-CORE-REQ-{n:03d}"
    rows.append(f"| `{rid}` | {spec[n]} | {tests[n]} | `ALLOCATED` |")

text = f'''---
id: DOC-AURORA-CAP-SOVEREIGN-CORE-R3-COVERAGE
title: CAP-SOVEREIGN-CORE R3 Requirement Allocation
document_type: requirement_coverage
form: reference
authority: reference
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - current R3 allocation of CAP-SOVEREIGN-CORE R2 requirements to Spec mechanisms and tests
related:
  - DOC-AURORA-CAP-SOVEREIGN-CORE-REQUIREMENTS
  - DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC
  - DOC-AURORA-CAP-SOVEREIGN-CORE-THREAT-MODEL
  - DOC-AURORA-CAP-SOVEREIGN-CORE-TEST-PLAN
  - DOC-AURORA-M0-R3-OPERATOR-AUTHORIZATION
source_revision: 9ea8adf5c115f54071d7e36e312695d19420d8b0
last_reviewed: 2026-08-07
---

# CAP-SOVEREIGN-CORE — R3 Requirement Allocation

## 1. Purpose

R3 must allocate every R2 requirement to a concrete current Spec mechanism/section and to one or more planned verification cases.

This matrix contains exactly one row for every `CAP-SOVEREIGN-CORE-REQ-001..122`.

```text
R2 requirements expected: 122
R3 allocation rows:        122
Unallocated:                 0
```

The matrix is allocation evidence; `SPEC.md` remains the reusable behavior owner and `TEST-PLAN.md` remains the test-plan owner.

## 2. Allocation matrix

| Requirement | Spec / mechanism allocation | Planned test allocation | Status |
|---|---|---|---|
''' + "\n".join(rows) + '''

## 3. R3 gate assertions

- all 122 R2 requirement IDs are represented exactly once;
- every row names at least one current Spec/threat/test mechanism or section;
- every row names at least one planned test/document-review case;
- no row allocates a requirement to a concrete implementation file/package/database/runtime;
- R4 open mechanism questions remain in `SPEC.md` §19 rather than masquerading as completed architecture decisions;
- the threat model and test plan are specialized subordinate artifacts and do not redefine the Spec.

## 4. Stop boundary

This coverage matrix does not authorize R4, Architecture Spike execution, Mission Contract, Microdesign or implementation.
'''

OUT.write_text(text, encoding="utf-8")
print(f"wrote {OUT.relative_to(ROOT)} with {len(rows)} allocations")
