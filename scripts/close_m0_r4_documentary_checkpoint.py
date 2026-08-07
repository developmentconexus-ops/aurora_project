#!/usr/bin/env python3
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
status_path = ROOT / "docs/tracking/STATUS.md"
worklog_path = ROOT / "docs/tracking/WORKLOG.md"

s = status_path.read_text(encoding="utf-8")
required = [
    "version: 0.15.0",
    "- **Current readiness gate:** ACRM R3 — Capability Readiness — PASS",
    "- **R4 — Architecture/Decision Readiness:** NOT AUTHORIZED / awaiting explicit operator authorization",
    "ACRM R4 — Architecture/Decision Readiness: NOT AUTHORIZED / AWAITING OPERATOR",
    "R3 PASS recorded\n→ stop at the R3 boundary\n→ await explicit operator authorization for M0 ACRM R4 — Architecture/Decision Readiness",
]
for item in required:
    if item not in s:
        raise SystemExit(f"STATUS precondition missing: {item}")

s = s.replace("version: 0.15.0", "version: 0.16.0", 1)
s = s.replace(
    "  - REVIEW-AURORA-M0-R3-CAPABILITY-READINESS-2026-08-07\nlast_reviewed:",
    "  - REVIEW-AURORA-M0-R3-CAPABILITY-READINESS-2026-08-07\n"
    "  - DOC-AURORA-M0-R4-OPERATOR-AUTHORIZATION\n"
    "  - DESIGN-AURORA-M0-R4-DECISION-LANDSCAPE\n"
    "  - DOC-AURORA-CAP-SOVEREIGN-CORE-R4-DECISION-COVERAGE\n"
    "  - ADR-AURORA-0003\n"
    "  - ADR-AURORA-0004\n"
    "  - ADR-AURORA-0005\n"
    "  - ADR-AURORA-0006\n"
    "  - ADR-AURORA-0007\n"
    "  - ADR-AURORA-0008\n"
    "  - DESIGN-AURORA-M0-SOVEREIGN-STORE-SPIKE-001\n"
    "  - DESIGN-AURORA-M0-OWNER-TRUST-SPIKE-002\n"
    "  - REVIEW-AURORA-M0-R4-ARCHITECTURE-DECISION-READINESS-2026-08-07\n"
    "last_reviewed:",
    1,
)
s = s.replace(
    "- **Current readiness gate:** ACRM R3 — Capability Readiness — PASS",
    "- **Current readiness gate:** ACRM R4 — Architecture/Decision Readiness — BLOCKED",
    1,
)
s = s.replace(
    "- **R3 verdict:** PASS\n- **R4 — Architecture/Decision Readiness:** NOT AUTHORIZED / awaiting explicit operator authorization\n- **R5 and later gates:** NOT AUTHORIZED BY IMPLICATION\n- **Stack decisions:** none\n- **Runtime implementation:** not started and not authorized",
    "- **R3 verdict:** PASS\n"
    "- **R4 source baseline:** `d00cc1abfc2a41ac7e81e1f3478e188b3c5e9e52`\n"
    "- **R4 documentary package/review integration:** `71f64bab2a82c2a7781d28274224f60abc277b2c`\n"
    "- **R4 architecture questions:** 15/15 mapped to research, proposed ADRs and/or exact spike evidence\n"
    "- **R4 focused research:** 4 current reports + 4 source manifests\n"
    "- **R4 ADRs:** ADR-0003..0008 all PROPOSED; 0003..0006 documentary-ready for operator review; 0007..0008 spike-blocked\n"
    "- **SPK-AURORA-M0-SOVEREIGN-STORE-001:** PROPOSED / EXECUTION NOT AUTHORIZED\n"
    "- **SPK-AURORA-M0-OWNER-TRUST-002:** PROPOSED / EXECUTION NOT AUTHORIZED / depends on reviewed SPK-001 result\n"
    "- **R4 verdict:** BLOCKED — required executable spike evidence and operator decision acceptance are not yet complete\n"
    "- **R5 — Contract Readiness:** NOT AUTHORIZED\n"
    "- **R6 and later gates:** NOT AUTHORIZED BY IMPLICATION\n"
    "- **Accepted stack decisions:** none; R4 technical choices remain proposed pending required review/evidence\n"
    "- **Runtime implementation:** not started and not authorized",
    1,
)

s = s.replace(
    "The following remain open and were not decided by R0, R1, R2 or R3. R3 fixed the behavior/boundaries they must satisfy; R4 owns the implementation mechanisms and current-source revalidation:",
    "R4 has now researched and proposed concrete dispositions for the M0 mechanism questions below, but none of the new ADRs is accepted yet. Store/atomicity and owner-root/time/restore mechanisms additionally require reviewed Architecture Spike evidence before acceptance:",
    1,
)

s = s.replace(
    "M0 ACRM R3 — Capability Readiness: PASS — source baseline 9ea8adf5c115f54071d7e36e312695d19420d8b0\n"
    "R3 proposed Capability package: Spec + threat model + test plan + 122/122 allocation\n"
    "ACRM R4 — Architecture/Decision Readiness: NOT AUTHORIZED / AWAITING OPERATOR\n"
    "ACRM R5+:                       NOT AUTHORIZED\n"
    "Architecture Spike execution:   PROHIBITED",
    "M0 ACRM R3 — Capability Readiness: PASS — source baseline 9ea8adf5c115f54071d7e36e312695d19420d8b0\n"
    "R3 proposed Capability package: Spec + threat model + test plan + 122/122 allocation\n"
    "M0 ACRM R4 — Architecture/Decision Readiness: BLOCKED — documentary baseline d00cc1abfc2a41ac7e81e1f3478e188b3c5e9e52\n"
    "R4 decision coverage:             15/15 mapped; 6 proposed ADRs; 2 required spikes unexecuted\n"
    "ACRM R5 — Contract Readiness:    NOT AUTHORIZED\n"
    "ACRM R6+:                       NOT AUTHORIZED\n"
    "Architecture Spike execution:   PROHIBITED / AWAITING EXPLICIT OPERATOR AUTHORIZATION",
    1,
)

old_gate = """The M0 ACRM R3 Capability Readiness analysis against fixed source baseline `9ea8adf5c115f54071d7e36e312695d19420d8b0` produced a proposed reusable `CAP-SOVEREIGN-CORE` Spec, threat model, Capability Test Plan and requirement-allocation matrix. All 122 R2 requirements are allocated to current Spec/threat mechanisms and planned verification. The R3-specific validator reported 122/122 allocation rows, 84 defined test IDs, 0 undefined test references, all required ACRM sections/owners/threat classes present and 0 selected technology candidates. Two material adversarial findings—restore revalidation without an explicit owner recovery root and an under-specified `accepted_state` envelope—were resolved before the final review returned `R3 PASS`.\n\nThe Capability Spec/threat model/test plan remain `proposed`; canonical integration and green CI do not constitute operator acceptance. R3 also determined that current research is sufficient for boundary reasoning while exact candidate/mechanism/version evidence MUST be revalidated in R4. The intentional boundary is now authorization: R4 has **not** been authorized. No Architecture/Decision Readiness work, Architecture Spike execution, stack choice, Mission Contract, Microdesign or implementation may begin by implication."""
new_gate = """M0 ACRM R4 has completed the documentary architecture/decision investigation against fixed source baseline `d00cc1abfc2a41ac7e81e1f3478e188b3c5e9e52`. All 15 R3-open architecture questions now have explicit dispositions. Four current focused research reports support six proposed ADRs: ADR-0003 through ADR-0006 are documentary-ready for operator review, while ADR-0007 (SQLite operational store) and ADR-0008 (Owner Root/recovery trust) are explicitly blocked by executable evidence. Two minimal sequential Architecture Spikes have been specified: `SPK-AURORA-M0-SOVEREIGN-STORE-001` first, then `SPK-AURORA-M0-OWNER-TRUST-002` only after the first spike has a reviewed viable store result.\n\nThe independent R4 review returned `R4 BLOCKED`, not FAIL and not PASS. The blocker is intentional and concrete: the accepted Research Map requires crash/restart/restore evidence before Sovereign Core storage/recovery commitment; Architecture Spike execution remains prohibited because no separate operator authorization exists; and all R4 ADRs remain `proposed`. No new technical choice is governing yet. R5, Mission Contract, Microdesign and implementation remain unauthorized."""
if old_gate not in s:
    raise SystemExit("R3 gate paragraph precondition missing")
s = s.replace(old_gate, new_gate, 1)

s = s.replace(
    "R3 PASS recorded\n→ stop at the R3 boundary\n→ await explicit operator authorization for M0 ACRM R4 — Architecture/Decision Readiness",
    "R4 documentary investigation reviewed\n"
    "→ R4 BLOCKED recorded\n"
    "→ operator reviews ADR-0003..0006\n"
    "→ separately authorize exact SPK-AURORA-M0-SOVEREIGN-STORE-001 execution if approved\n"
    "→ SPK-002 remains blocked until SPK-001 evidence is reviewed\n"
    "→ do not begin R5",
    1,
)

status_path.write_text(s, encoding="utf-8")

w = worklog_path.read_text(encoding="utf-8")
entry = r'''

## 2026-08-07 — M0 R4 Documentary Architecture/Decision Checkpoint

After R3 PASS, the operator authorized `M0 ACRM R4 — Architecture/Decision Readiness` with `Vamos para próximo passo então` and subsequently approved the decision philosophy of long-horizon exploration with evidence-bounded commitment.

R4 fixed its source baseline at:

```text
d00cc1abfc2a41ac7e81e1f3478e188b3c5e9e52
```

A decision landscape mapped all 15 R3-open architecture questions by dependency and lock-in. Four focused current research reports plus source manifests were created for runtime/persistence, portability/integrity/export, owner-root/authority/recovery trust and observability.

Six ADRs were proposed but not accepted:

```text
ADR-0003 Go as initial Core runtime
ADR-0004 local modular Core + explicit current state/revisions + no full event sourcing/durable engine in M0
ADR-0005 JSON Schema/JSON/JCS portable state + age outer export + explicit migrations
ADR-0006 OpenTelemetry traces/metrics + slog; optional backend
ADR-0007 SQLite operational store — spike-blocked
ADR-0008 Owner Root/recovery trust boundary — spike-blocked
```

Two sequential Architecture Spikes were specified but not authorized to execute:

```text
SPK-AURORA-M0-SOVEREIGN-STORE-001
→ compare Go SQLite bindings and prove crash/restart/backup/restore/atomicity

SPK-AURORA-M0-OWNER-TRUST-002
→ only after reviewed SPK-001 result; prove owner-root/rollback/time/restore-freshness protocol
```

The documentary validator initially exposed a test/index identifier mismatch; after correction it passed. The repository documentation validator then exposed invalid conceptual `SPK-*` values being used as canonical document IDs. The package was corrected to use `DESIGN-*` document IDs while preserving explicit `spike_id: SPK-*` identities. GitHub Actions run `31200153496` then passed R4 package validation, docs generation, documentation validation and projection freshness. The formal R4 review also passed validation in run `31200329122`.

The clean R4 documentary package/review was integrated to canonical `main` at:

```text
71f64bab2a82c2a7781d28274224f60abc277b2c
```

Normal main Documentation workflow run `31200439197` completed successfully.

The formal R4 verdict is:

```text
R4 BLOCKED
```

This is not an architecture failure. All 15 questions are accounted for, but accepted governance requires executable crash/restart/restore evidence before storage/recovery commitment, Architecture Spike execution has not been authorized, and all six new ADRs remain proposed. ADR-0003..0006 are ready for operator review; ADR-0007..0008 must remain proposed until their exact spike evidence is complete and reviewed.

R5, Mission Contract, Microdesign and implementation remain unauthorized.
'''
if "## 2026-08-07 — M0 R4 Documentary Architecture/Decision Checkpoint" in w:
    raise SystemExit("R4 worklog entry already exists")
worklog_path.write_text(w.rstrip() + entry.rstrip() + "\n", encoding="utf-8")
print("prepared M0 R4 documentary checkpoint tracking")
