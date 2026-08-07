from pathlib import Path


def replace_once(path: str, old: str, new: str) -> None:
    p = Path(path)
    text = p.read_text(encoding="utf-8")
    n = text.count(old)
    if n != 1:
        raise SystemExit(f"{path}: expected exactly one occurrence, found {n}: {old!r}")
    p.write_text(text.replace(old, new, 1), encoding="utf-8")


# ADR index
idx = "docs/adr/README.md"
replace_once(idx, "version: 0.4.1", "version: 0.4.2")
replace_once(
    idx,
    "| [ADR-0008](0008-m0-owner-root-recovery-trust.md) | M0 Owner Root and Recovery Trust Boundary | proposed / spike-blocked | wrapped random Owner Root + external trust anchor; requires SPK-002 evidence |",
    "| [ADR-0008](0008-m0-owner-root-recovery-trust.md) | M0 Owner Root and Recovery Trust Boundary | proposed / evidence-ready | random ORK + authenticated external trust high-water; SPK-002 PASS/CLOSED; operator decision pending |",
)
replace_once(
    idx,
    "ADR-0008 remains proposed and MUST NOT be accepted before its separately required `SPK-AURORA-M0-OWNER-TRUST-002` evidence is complete and reviewed. SPK-002 execution is explicitly authorized by `DOC-AURORA-M0-R4-ADR0007-ACCEPTANCE-SPK002-AUTHORIZATION`; no other spike is authorized by implication.",
    "ADR-0008 v0.2.0 is evidence-ready after reviewed `SPK-AURORA-M0-OWNER-TRUST-002` PASS/CLOSED. It remains `proposed` pending explicit operator ACCEPT/REJECT/REVISE; no implementation or R5 authorization follows by implication.",
)

# Research Map
research = "docs/research/RESEARCH-MAP.md"
replace_once(research, "version: 0.4.1", "version: 0.4.2")
replace_once(
    research,
    "documentary research complete; `SPK-AURORA-M0-OWNER-TRUST-002` specified and explicitly authorized for execution; evidence not yet reviewed",
    "`SPK-AURORA-M0-OWNER-TRUST-002` PASS/REVIEWED/CLOSED; executable evidence informs ADR-0008 v0.2.0; operator decision pending",
)

# STATUS
status = "docs/tracking/STATUS.md"
replace_once(status, "version: 0.20.0", "version: 0.21.0")
replace_once(
    status,
    "  - DOC-AURORA-M0-R4-ADR0007-ACCEPTANCE-SPK002-AUTHORIZATION\n  - REVIEW-AURORA-M0-R4-SPK001-SOVEREIGN-STORE-2026-08-07\n",
    "  - DOC-AURORA-M0-R4-ADR0007-ACCEPTANCE-SPK002-AUTHORIZATION\n  - DOC-AURORA-M0-R4-SPK002-EVIDENCE-RECEIPT\n  - REVIEW-AURORA-M0-R4-SPK001-SOVEREIGN-STORE-2026-08-07\n  - REVIEW-AURORA-M0-R4-SPK002-OWNER-TRUST-2026-08-07\n",
)
replace_once(
    status,
    "- **ADR-0008:** PROPOSED / SPK-002-BLOCKED",
    "- **ADR-0008:** PROPOSED / EVIDENCE-READY — SPK-002 PASS/CLOSED; operator decision pending",
)
replace_once(
    status,
    "- **SPK-AURORA-M0-OWNER-TRUST-002:** AUTHORIZED FOR EXECUTION — exact canonical spec v0.1.0 / blob `0ffb6fa2b35014e34b4301365dc2d5a8d96f021d`; no other spike authorized",
    "- **SPK-AURORA-M0-OWNER-TRUST-002:** PASS / EVIDENCE_COMPLETE / REVIEWED / DECISION_INFORMED / CLOSED — final run `31219882882`, execution revision `c76b96fee36878f15c54028b4ba1896f84ebdeca`, Linux/Windows PASS",
)
replace_once(
    status,
    "- **R4 verdict:** BLOCKED — ADR-0007 is accepted; SPK-002 execution/review and ADR-0008 operator decision remain",
    "- **R4 verdict:** BLOCKED ONLY ON ADR-0008 OPERATOR DECISION — all required M0 architecture spikes are PASS/REVIEWED/CLOSED",
)
replace_once(
    status,
    "R4 has researched concrete dispositions for all M0 mechanism questions. ADR-0003, ADR-0004, ADR-0005 and ADR-0006 are accepted. SPK-001 has now proved and closed the store/atomicity/backup/migration uncertainty, making ADR-0007 evidence-ready; ADR-0008 remains proposed because owner-root/time/restore-freshness behavior still requires SPK-002 evidence:",
    "R4 has researched concrete dispositions for all M0 mechanism questions. ADR-0003, ADR-0004, ADR-0005, ADR-0006 and ADR-0007 are accepted. SPK-001 and SPK-002 have both completed with reviewed PASS evidence. ADR-0008 v0.2.0 is evidence-ready and remains proposed pending the operator decision:",
)
replace_once(
    status,
    "R4 M0 decision coverage:         15/15 mapped; ADR-0003/4/5/6/7 accepted; ADR-0008 SPK-002-blocked",
    "R4 M0 decision coverage:         15/15 mapped; all required spikes closed; ADR-0008 evidence-ready/operator-decision pending",
)
replace_once(
    status,
    "Architecture Spike execution:   SPK-001 CLOSED; SPK-002 AUTHORIZED ONLY; all other spikes NOT AUTHORIZED",
    "Architecture Spike execution:   SPK-001 CLOSED; SPK-002 CLOSED; all other spikes NOT AUTHORIZED",
)
old_gate = """The remaining R4 blockers are now narrower:

1. execute and independently review the explicitly authorized `SPK-AURORA-M0-OWNER-TRUST-002`;
2. after SPK-002 evidence is reviewed, ADR-0008 must be accepted/rejected/revised by the operator.

ADR-0009 acceptance remains cross-horizon and does not authorize Mastra implementation. SPK-002 is now explicitly authorized by the operator; no other spike is authorized by implication.

Therefore the independent R4 verdict remains:

```text
R4 BLOCKED
```

R5, Mission Contract, Microdesign and runtime implementation remain unauthorized.

## 8. Immediate next action

```text
ADR-0007 v0.2.0 ACCEPTED
→ SPK-AURORA-M0-OWNER-TRUST-002 AUTHORIZED
→ execute exact disposable owner-root/trust spike against fixed revision
→ independently review evidence
→ present evidence-informed ADR-0008 decision to operator
→ STOP
→ do not begin R5
```"""
new_gate = """SPK-002 has now completed successfully. Final hardened workflow `31219882882` passed the full S01–S12, recovery-classification, mutation-boundary and secret-hygiene evidence on Ubuntu and Windows. The independent review closed the spike as `PASS / REVIEWED / DECISION_INFORMED` and recommends ADR-0008 v0.2.0.

No M0 architecture spike remains to execute. ADR-0008 v0.2.0 remains `proposed`; the exact remaining R4 blocker is the operator decision to ACCEPT, REJECT or REVISE it.

ADR-0009 acceptance remains cross-horizon and does not authorize Mastra implementation. Completion of the architecture spikes does not authorize R5 by implication.

Therefore the current independent R4 verdict remains:

```text
R4 BLOCKED
```

with exactly one remaining blocker:

```text
ADR-0008 v0.2.0 operator decision
```

R5, Mission Contract, Microdesign and runtime implementation remain unauthorized.

## 8. Immediate next action

```text
SPK-AURORA-M0-OWNER-TRUST-002 PASS / CLOSED
→ operator reviews ADR-0008 v0.2.0
→ ACCEPT | REJECT | REVISE
→ STOP
→ do not begin R5
```"""
replace_once(status, old_gate, new_gate)

# WORKLOG append
worklog = Path("docs/tracking/WORKLOG.md")
w = worklog.read_text(encoding="utf-8").rstrip()
if "## 2026-08-07 — SPK-002 Owner Trust PASS / Evidence Closeout" in w:
    raise SystemExit("WORKLOG already contains SPK-002 closeout entry")
w += """

## 2026-08-07 — SPK-002 Owner Trust PASS / Evidence Closeout

The authorized disposable `SPK-AURORA-M0-OWNER-TRUST-002` executed from canonical authorization baseline `ba7a211b40e21bfcd9aa1b90e5ca3c21cd229d1b` against canonical specification blob `0ffb6fa2b35014e34b4301365dc2d5a8d96f021d` (v0.1.0).

Final hardened executable evidence came from branch revision `c76b96fee36878f15c54028b4ba1896f84ebdeca` and GitHub Actions run `31219882882`. Ubuntu and Windows both passed S01–S12, all required DB/anchor recovery classifications, the post-review mutation-boundary hardening tests, measurements and secret-hygiene checks. The aggregate gate passed.

Candidate A—random 256-bit ORK, Argon2id-derived KEK, AES-256-GCM wrapped ORK, HKDF-SHA-256 purpose keys, HMAC-SHA-256 governing/trust descriptors and an authenticated external generation/observed-time high-water—passed the intended M0 local threat boundary. Candidate B was not duplicated because Candidate A proved passphrase rotation can preserve stable ORK/domain-root lineage while the direct passphrase-root alternative would couple rotation to the long-lived integrity root/rekey continuity. Candidate C remains a future hardware/OS trust strengthening class if the threat model expands.

Two findings were corrected before final evidence: (1) a measurement executable was initially placed inside the evidence artifact and contained deterministic fixture strings, so final packaging moved the binary outside evidence and both hygiene gates passed; (2) adversarial review found ordinary mutation needed to enforce anomaly classification itself, so `STATE_ROLLBACK`, `ANCHOR_LAG`, `TIME_UNTRUSTED` and `REVALIDATION_REQUIRED` now fail before normal mutation unless the explicit recovery path is used.

The reviewed lifecycle is `PASS / EVIDENCE_COMPLETE / REVIEWED / DECISION_INFORMED / CLOSED`. ADR-0008 was revised to v0.2.0 and remains `proposed / evidence-ready` pending explicit operator ACCEPT/REJECT/REVISE. R4 remains BLOCKED only on that operator decision. R5, Mission Contract, Microdesign and production implementation remain unauthorized.
"""
worklog.write_text(w + "\n", encoding="utf-8")
