#!/usr/bin/env python3
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
status_path = ROOT / "docs/tracking/STATUS.md"
worklog_path = ROOT / "docs/tracking/WORKLOG.md"

s = status_path.read_text(encoding="utf-8")
required = [
    "version: 0.14.0",
    "- **Current readiness gate:** ACRM R2 — Requirements — PASS",
    "- **R3 — Capability Readiness:** NOT AUTHORIZED / awaiting explicit operator authorization",
    "ACRM R3 — Capability Readiness: NOT AUTHORIZED / AWAITING OPERATOR",
    "R2 PASS recorded\n→ stop at the R2 boundary\n→ await explicit operator authorization for M0 ACRM R3 — Capability Readiness",
]
for item in required:
    if item not in s:
        raise SystemExit(f"STATUS precondition missing: {item}")

s = s.replace("version: 0.14.0", "version: 0.15.0", 1)
s = s.replace(
    "  - REVIEW-AURORA-M0-R2-REQUIREMENTS-2026-08-07\nlast_reviewed:",
    "  - REVIEW-AURORA-M0-R2-REQUIREMENTS-2026-08-07\n"
    "  - DOC-AURORA-M0-R3-OPERATOR-AUTHORIZATION\n"
    "  - DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC\n"
    "  - DOC-AURORA-CAP-SOVEREIGN-CORE-THREAT-MODEL\n"
    "  - DOC-AURORA-CAP-SOVEREIGN-CORE-TEST-PLAN\n"
    "  - DOC-AURORA-CAP-SOVEREIGN-CORE-R3-COVERAGE\n"
    "  - REVIEW-AURORA-M0-R3-RESEARCH-FRESHNESS-2026-08-07\n"
    "  - REVIEW-AURORA-M0-R3-CAPABILITY-READINESS-2026-08-07\n"
    "last_reviewed:",
    1,
)
s = s.replace("- **Current readiness gate:** ACRM R2 — Requirements — PASS", "- **Current readiness gate:** ACRM R3 — Capability Readiness — PASS", 1)
s = s.replace(
    "- **R2 verdict:** PASS\n- **R3 — Capability Readiness:** NOT AUTHORIZED / awaiting explicit operator authorization\n- **R4 and later gates:** NOT AUTHORIZED BY IMPLICATION",
    "- **R2 verdict:** PASS\n"
    "- **R3 source baseline:** `9ea8adf5c115f54071d7e36e312695d19420d8b0`\n"
    "- **R3 reviewed clean package:** `4b8558b724f28310fd8fbc6884944f7f59f16ea6`\n"
    "- **R3 canonical integration:** `58a7946b62f27d8b8784169e7e3741eec24ecc95`\n"
    "- **R3 requirement allocation:** 122/122 R2 requirements allocated to Spec mechanisms and planned verification\n"
    "- **R3 test-plan baseline:** 84 planned test IDs; 80 referenced directly by requirement coverage\n"
    "- **R3 research freshness:** SUFFICIENT for boundary reasoning; R4 mechanism/version revalidation required\n"
    "- **R3 verdict:** PASS\n"
    "- **R4 — Architecture/Decision Readiness:** NOT AUTHORIZED / awaiting explicit operator authorization\n"
    "- **R5 and later gates:** NOT AUTHORIZED BY IMPLICATION",
    1,
)

start = "The following remain open and were not decided by R0, R1 or R2; they belong to later applicable readiness gates:\n\n"
end = "\n## 6. Current authorization boundary"
if start not in s:
    raise SystemExit("open decisions section precondition missing")
a, rest = s.split(start, 1)
old_body, b = rest.split(end, 1)
new_body = """The following remain open and were not decided by R0, R1, R2 or R3. R3 fixed the behavior/boundaries they must satisfy; R4 owns the implementation mechanisms and current-source revalidation:\n\n- Aurora Core implementation language/runtime;\n- operational-state storage mechanism;\n- state-versus-event persistence pattern;\n- schema/serialization representation;\n- crash-consistent commit/atomicity mechanism;\n- integrity mechanism;\n- time/rollback semantics used for authority expiry;\n- local owner authentication/bootstrap mechanism;\n- export/backup format and topology;\n- migration mechanism/tooling;\n- audit/event physical mechanism;\n- telemetry backend/transport;\n- initial process/deployment topology;\n- durable execution engine applicability, only if M0 need proves it proportionate;\n- authority freshness/revalidation mechanism after restore.\n"""
s = a + new_body + end + b

s = s.replace(
    "M0 ACRM R2 — Requirements:     PASS — source baseline 495b712142d7c3d722da2298f7a0b060707f9f5e\n"
    "R2 requirements baseline:       122 proposed atomic requirements; coverage 127/127\n"
    "ACRM R3 — Capability Readiness: NOT AUTHORIZED / AWAITING OPERATOR\n"
    "ACRM R4+:                       NOT AUTHORIZED",
    "M0 ACRM R2 — Requirements:     PASS — source baseline 495b712142d7c3d722da2298f7a0b060707f9f5e\n"
    "R2 requirements baseline:       122 proposed atomic requirements; coverage 127/127\n"
    "M0 ACRM R3 — Capability Readiness: PASS — source baseline 9ea8adf5c115f54071d7e36e312695d19420d8b0\n"
    "R3 proposed Capability package: Spec + threat model + test plan + 122/122 allocation\n"
    "ACRM R4 — Architecture/Decision Readiness: NOT AUTHORIZED / AWAITING OPERATOR\n"
    "ACRM R5+:                       NOT AUTHORIZED",
    1,
)

old_gate = """The M0 ACRM R2 requirements derivation against fixed source baseline `495b712142d7c3d722da2298f7a0b060707f9f5e` transformed all 127 R1-active constitutional source rows into a reviewed package of 122 proposed atomic requirements with 127/127 source coverage and returned `R2 PASS`. The R2-specific validator reported 0 uncovered active sources, 0 inactive source references, 0 unsupported verification methods and 0 selected technology candidates. Eight adversarial wording findings were corrected before the final verdict.\n\nThe requirements remain `proposed`; canonical integration and green CI do not constitute operator acceptance. The intentional boundary is now authorization: R3 has **not** been authorized. No Capability/System Spec, threat-model execution, Architecture Spike, stack choice, Mission Contract, Microdesign or implementation may begin by implication."""
new_gate = """The M0 ACRM R3 Capability Readiness analysis against fixed source baseline `9ea8adf5c115f54071d7e36e312695d19420d8b0` produced a proposed reusable `CAP-SOVEREIGN-CORE` Spec, threat model, Capability Test Plan and requirement-allocation matrix. All 122 R2 requirements are allocated to current Spec/threat mechanisms and planned verification. The R3-specific validator reported 122/122 allocation rows, 84 defined test IDs, 0 undefined test references, all required ACRM sections/owners/threat classes present and 0 selected technology candidates. Two material adversarial findings—restore revalidation without an explicit owner recovery root and an under-specified `accepted_state` envelope—were resolved before the final review returned `R3 PASS`.\n\nThe Capability Spec/threat model/test plan remain `proposed`; canonical integration and green CI do not constitute operator acceptance. R3 also determined that current research is sufficient for boundary reasoning while exact candidate/mechanism/version evidence MUST be revalidated in R4. The intentional boundary is now authorization: R4 has **not** been authorized. No Architecture/Decision Readiness work, Architecture Spike execution, stack choice, Mission Contract, Microdesign or implementation may begin by implication."""
if old_gate not in s:
    raise SystemExit("current gate paragraph precondition missing")
s = s.replace(old_gate, new_gate, 1)

s = s.replace(
    "R2 PASS recorded\n→ stop at the R2 boundary\n→ await explicit operator authorization for M0 ACRM R3 — Capability Readiness",
    "R3 PASS recorded\n→ stop at the R3 boundary\n→ await explicit operator authorization for M0 ACRM R4 — Architecture/Decision Readiness",
    1,
)

status_path.write_text(s, encoding="utf-8")

w = worklog_path.read_text(encoding="utf-8")
entry = """

## 2026-08-07 — M0 R3 Capability Readiness

The operator explicitly authorized `M0 ACRM R3 — Capability Readiness` after R2 PASS. The fixed R3 source baseline was `9ea8adf5c115f54071d7e36e312695d19420d8b0`.

R3 produced a proposed modular Capability package for `CAP-SOVEREIGN-CORE`:

```text
docs/capabilities/CAP-SOVEREIGN-CORE/SPEC.md
docs/capabilities/CAP-SOVEREIGN-CORE/THREAT-MODEL.md
docs/capabilities/CAP-SOVEREIGN-CORE/TEST-PLAN.md
docs/capabilities/CAP-SOVEREIGN-CORE/R3-COVERAGE.md
```

The logical design keeps one Sovereign Core Capability and assigns distinct owners for Identity, Project State, Authority, Portability/Recovery and Audit/Evidence. Replaceable ports keep operator interaction, durable storage, evidence/artifact payloads, time and integrity mechanisms outside domain ownership.

R3 fixed M0 domain/lifecycle semantics without selecting implementation technology. Accepted state uses explicit revision semantics and a minimum semantic envelope; authority snapshots and next-safe-action are projections over current state/authority; ordinary restart re-evaluates current authority while stale backup restore fails closed through `REVALIDATION_REQUIRED`. A restored grant cannot authorize its own revalidation; only the authenticated Leandro owner recovery root can create a new attributable authority-state revision.

The threat model covers state corruption/rollback, authority resurrection, unsafe restore, identity collision, export exposure/tampering, untrusted-content authority injection, clock rollback, crash ambiguity, migration drift, evidence spoofing and sensitive telemetry. R3 assigns the accepted Aurora data classes to all material M0 data families.

The Capability Test Plan defines 15 deterministic logical fixtures and 84 test IDs. `R3-COVERAGE.md` allocates all 122 R2 requirements to Spec mechanisms and planned verification. A dedicated validator run `31150888970` reported:

```text
R2 requirements: 122
R3 coverage rows: 122
planned test IDs defined: 84
test IDs referenced by coverage: 80
undefined referenced tests: 0
required ACRM Spec sections: present
required logical owners: present
required semantic closure: present
required threat classes: present
critical negative test cases: present
candidate technology selections: 0
```

Two material findings were corrected before verdict:

- `R3-F01`: restore fail-closed semantics lacked an explicit safe owner recovery/revalidation root;
- `R3-F02`: `accepted_state` was under-specified and could have left product semantics to R4.

The Research Map freshness boundary was reviewed. The current durability, authority/identity/effects and events/observability/schema reports remain sufficient for R3 boundary reasoning; exact mechanism/version primary-source revalidation remains mandatory for R4.

The clean R3 package reviewed at `4b8558b724f28310fd8fbc6884944f7f59f16ea6` received:

```text
R3 PASS
```

The package was integrated canonically at `58a7946b62f27d8b8784169e7e3741eec24ecc95` and passed the normal Documentation workflow. Spec/threat/test artifacts remain `proposed`; merge/CI do not constitute operator acceptance. No stack, Architecture Spike, Mission Contract, Microdesign or implementation was selected or executed. R4 remains separately gated and unauthorized.
"""
if "## 2026-08-07 — M0 R3 Capability Readiness" in w:
    raise SystemExit("R3 worklog entry already exists")
worklog_path.write_text(w.rstrip() + entry + "\n", encoding="utf-8")
print("closed M0 R3 tracking state")
