from pathlib import Path


def replace_once(path: str, old: str, new: str) -> None:
    p = Path(path)
    text = p.read_text(encoding="utf-8")
    n = text.count(old)
    if n != 1:
        raise SystemExit(f"{path}: expected exactly one occurrence, found {n}: {old!r}")
    p.write_text(text.replace(old, new, 1), encoding="utf-8")


# ADR-0008 acceptance
adr = "docs/adr/0008-m0-owner-root-recovery-trust.md"
replace_once(
    adr,
    "status: proposed\nversion: 0.2.0",
    "status: accepted\naccepted_at: 2026-08-07\nacceptance_evidence: DOC-AURORA-M0-R4-ADR0008-ACCEPTANCE\nversion: 0.2.0",
)
replace_once(
    adr,
    "  - proposed M0 local owner authentication, integrity-anchor, time rollback and restore revalidation mechanism",
    "  - M0 local owner authentication, integrity-anchor, time rollback and restore revalidation mechanism",
)
replace_once(
    adr,
    "  - REVIEW-AURORA-M0-R4-SPK002-OWNER-TRUST-2026-08-07\n",
    "  - REVIEW-AURORA-M0-R4-SPK002-OWNER-TRUST-2026-08-07\n  - DOC-AURORA-M0-R4-ADR0008-ACCEPTANCE\n  - REVIEW-AURORA-M0-R4-ARCHITECTURE-DECISION-READINESS-RERUN-2026-08-07\n",
)
replace_once(
    adr,
    "**Proposed, evidence-ready for operator acceptance:** use a random **256-bit Owner Root Key (ORK)** as the stable local owner/domain integrity root.",
    "**Accepted by the operator on 2026-08-07:** use a random **256-bit Owner Root Key (ORK)** as the stable local owner/domain integrity root.",
)
replace_once(
    adr,
    "Before this ADR becomes governing:\n\n```text\noperator reviews ADR-0008 v0.2.0\n→ ACCEPTED | REJECTED | REVISED\n```\n\nThis ADR remains `proposed` until that explicit operator decision.",
    "ADR-0008 v0.2.0 was explicitly accepted by the operator through `DOC-AURORA-M0-R4-ADR0008-ACCEPTANCE`. It is now a governing R4 architecture decision. Acceptance does not authorize R5, Microdesign or production implementation.",
)
replace_once(
    adr,
    "- `REVIEW-AURORA-M0-R4-SPK002-OWNER-TRUST-2026-08-07`\n- `SPK-AURORA-M0-OWNER-TRUST-002`",
    "- `REVIEW-AURORA-M0-R4-SPK002-OWNER-TRUST-2026-08-07`\n- `DOC-AURORA-M0-R4-ADR0008-ACCEPTANCE`\n- `REVIEW-AURORA-M0-R4-ARCHITECTURE-DECISION-READINESS-RERUN-2026-08-07`\n- `SPK-AURORA-M0-OWNER-TRUST-002`",
)

# ADR index
idx = "docs/adr/README.md"
replace_once(idx, "version: 0.4.2", "version: 0.4.3")
replace_once(
    idx,
    "| [ADR-0008](0008-m0-owner-root-recovery-trust.md) | M0 Owner Root and Recovery Trust Boundary | proposed / evidence-ready | random ORK + authenticated external trust high-water; SPK-002 PASS/CLOSED; operator decision pending |",
    "| [ADR-0008](0008-m0-owner-root-recovery-trust.md) | M0 Owner Root and Recovery Trust Boundary | accepted | random ORK + authenticated external trust high-water; historical restore requires owner revalidation |",
)
replace_once(
    idx,
    "ADR-0008 v0.2.0 is evidence-ready after reviewed `SPK-AURORA-M0-OWNER-TRUST-002` PASS/CLOSED. It remains `proposed` pending explicit operator ACCEPT/REJECT/REVISE; no implementation or R5 authorization follows by implication.",
    "ADR-0008 v0.2.0 was explicitly accepted by the operator on 2026-08-07 through `DOC-AURORA-M0-R4-ADR0008-ACCEPTANCE`. This acceptance does not authorize implementation or R5 by implication.",
)

# Research map
research = "docs/research/RESEARCH-MAP.md"
replace_once(research, "version: 0.4.2", "version: 0.4.3")
replace_once(
    research,
    "`SPK-AURORA-M0-OWNER-TRUST-002` PASS/REVIEWED/CLOSED; executable evidence informs ADR-0008 v0.2.0; operator decision pending",
    "`SPK-AURORA-M0-OWNER-TRUST-002` PASS/REVIEWED/CLOSED; ADR-0008 v0.2.0 accepted from reviewed executable evidence",
)

# R4 decision coverage
coverage = "docs/capabilities/CAP-SOVEREIGN-CORE/R4-DECISION-COVERAGE.md"
replace_once(coverage, "version: 0.4.0", "version: 0.5.0")
replace_once(
    coverage,
    "  - REVIEW-AURORA-M0-R4-SPK002-OWNER-TRUST-2026-08-07\n",
    "  - REVIEW-AURORA-M0-R4-SPK002-OWNER-TRUST-2026-08-07\n  - DOC-AURORA-M0-R4-ADR0008-ACCEPTANCE\n  - REVIEW-AURORA-M0-R4-ARCHITECTURE-DECISION-READINESS-RERUN-2026-08-07\n",
)
replace_once(
    coverage,
    "ADR-0003 / 0004 / 0005 / 0006 / 0007: ACCEPTED\nADR-0009: ACCEPTED cross-horizon\nSPK-001: PASS / REVIEWED / CLOSED\nSPK-002: PASS / REVIEWED / CLOSED\nADR-0008 v0.2.0: PROPOSED / EVIDENCE-READY\nR4: BLOCKED only on operator decision for ADR-0008",
    "ADR-0003 / 0004 / 0005 / 0006 / 0007 / 0008: ACCEPTED\nADR-0009: ACCEPTED cross-horizon\nSPK-001: PASS / REVIEWED / CLOSED\nSPK-002: PASS / REVIEWED / CLOSED\nR4: PASS",
)
replacements = {
    "| `R4-Q-INTEGRITY-001` integrity | HIGH_LOCK_IN_SECURITY | ADR-0005 accepted content digest boundary; ADR-0008 v0.2.0 proposes ORK-derived HMAC governing integrity + external trust anchor | SPK-002 PASS/CLOSED | `EVIDENCE_READY / ADR0008_OPERATOR_DECISION_REQUIRED` |":
    "| `R4-Q-INTEGRITY-001` integrity | HIGH_LOCK_IN_SECURITY | ADR-0005 accepted content digest boundary; ADR-0008 accepted ORK-derived HMAC governing integrity + external trust anchor | SPK-002 PASS/CLOSED | `DECIDED / ACCEPTED + EVIDENCED` |",
    "| `R4-Q-TIME-001` time/rollback | HIGH_LOCK_IN_SECURITY | ADR-0008 v0.2.0 proposes monotonic local duration + authenticated observed wall-time high-water / fail closed | SPK-002 S07 + final matrix PASS | `EVIDENCE_READY / ADR0008_OPERATOR_DECISION_REQUIRED` |":
    "| `R4-Q-TIME-001` time/rollback | HIGH_LOCK_IN_SECURITY | ADR-0008 accepted monotonic local duration + authenticated observed wall-time high-water / fail closed | SPK-002 S07 + final matrix PASS | `DECIDED / ACCEPTED + EVIDENCED` |",
    "| `R4-Q-AUTHN-001` owner auth/bootstrap | HIGH_LOCK_IN_SECURITY | ADR-0008 v0.2.0 proposes random ORK + Argon2id KEK + AES-256-GCM wrapped root | SPK-002 S01/S02/S11/S12 PASS | `EVIDENCE_READY / ADR0008_OPERATOR_DECISION_REQUIRED` |":
    "| `R4-Q-AUTHN-001` owner auth/bootstrap | HIGH_LOCK_IN_SECURITY | ADR-0008 accepted random ORK + Argon2id KEK + AES-256-GCM wrapped root | SPK-002 S01/S02/S11/S12 PASS | `DECIDED / ACCEPTED + EVIDENCED` |",
    "| `R4-Q-EXPORT-001` export/backup | REVERSIBLE_SECURITY | ADR-0005 accepted logical export; ADR-0007 accepted operational SQLite backup; ADR-0008 v0.2.0 defines encrypted root recovery without historical current high-water | SPK-001 + SPK-002 PASS/CLOSED | `EVIDENCE_READY / ADR0008_OPERATOR_DECISION_REQUIRED` |":
    "| `R4-Q-EXPORT-001` export/backup | REVERSIBLE_SECURITY | ADR-0005 accepted logical export; ADR-0007 accepted operational SQLite backup; ADR-0008 accepted encrypted root recovery without historical current high-water | SPK-001 + SPK-002 PASS/CLOSED | `DECIDED / ACCEPTED + EVIDENCED` |",
    "| `R4-Q-RESTORE-001` authority freshness after restore | HIGH_LOCK_IN_SECURITY | ADR-0008 v0.2.0 proposes `REVALIDATION_REQUIRED` + authenticated-owner-only new authority revision | SPK-002 S08–S12 PASS | `EVIDENCE_READY / ADR0008_OPERATOR_DECISION_REQUIRED` |":
    "| `R4-Q-RESTORE-001` authority freshness after restore | HIGH_LOCK_IN_SECURITY | ADR-0008 accepted `REVALIDATION_REQUIRED` + authenticated-owner-only new authority revision | SPK-002 S08–S12 PASS | `DECIDED / ACCEPTED + EVIDENCED` |",
}
for old, new in replacements.items():
    replace_once(coverage, old, new)
replace_once(
    coverage,
    "ADR-0007 — SQLite + database/sql + modernc.org/sqlite operational-state baseline\n```",
    "ADR-0007 — SQLite + database/sql + modernc.org/sqlite operational-state baseline\nADR-0008 — random ORK + authenticated external trust high-water + fail-closed restore revalidation\n```",
)
old_tail = """## 6. Remaining R4 decision

There is no remaining M0 architecture spike to execute.

The only remaining material R4 decision is:

```text
ADR-0008 v0.2.0
→ operator ACCEPT | REJECT | REVISE
```

ADR-0008 remains non-governing until that explicit operator decision.

## 7. R4 gate implication

Current state:

```text
15/15 M0 architecture questions accounted for
all required M0 architecture spikes PASS / REVIEWED / CLOSED
ADR-0003 / 0004 / 0005 / 0006 / 0007 accepted
ADR-0009 accepted cross-horizon
ADR-0008 evidence-ready / proposed
```

Therefore the correct current verdict is still:

```text
R4 BLOCKED
```

but the blocker is now exactly one operator decision:

```text
ADR-0008 v0.2.0
```

If ADR-0008 is accepted and no new material finding appears, R4 may then be reviewed for PASS. R5 still requires separate explicit operator authorization; it does not begin automatically from ADR acceptance or R4 PASS."""
new_tail = """## 6. R4 closure

There is no remaining M0 architecture spike or material R4 architecture decision.

ADR-0008 v0.2.0 was explicitly accepted through `DOC-AURORA-M0-R4-ADR0008-ACCEPTANCE`. The final R4 readiness rerun found no new material architecture blocker.

## 7. R4 gate implication

Current state:

```text
15/15 M0 architecture questions accounted for
all required M0 architecture spikes PASS / REVIEWED / CLOSED
ADR-0003 / 0004 / 0005 / 0006 / 0007 / 0008 accepted
ADR-0009 accepted cross-horizon
```

Therefore:

```text
R4 PASS
```

R4 PASS makes the architecture eligible to support a scoped Mission Contract, but R5 remains separately gated and is **NOT AUTHORIZED** until explicit operator approval."""
replace_once(coverage, old_tail, new_tail)

# Preserve the first R4 review as its checkpoint rather than current verdict owner.
old_review = "docs/reviews/2026-08-07-m0-r4-architecture-decision-readiness-review.md"
replace_once(old_review, "version: 1.0.0", "version: 1.0.1")
replace_once(
    old_review,
    "  - M0 R4 documentary architecture review observations and current verdict",
    "  - M0 R4 documentary architecture review observations and verdict at the pre-spike checkpoint",
)
replace_once(
    old_review,
    "  - DESIGN-AURORA-M0-OWNER-TRUST-SPIKE-002\nsource_revision:",
    "  - DESIGN-AURORA-M0-OWNER-TRUST-SPIKE-002\n  - REVIEW-AURORA-M0-R4-ARCHITECTURE-DECISION-READINESS-RERUN-2026-08-07\nsource_revision:",
)

# STATUS
status = "docs/tracking/STATUS.md"
replace_once(status, "version: 0.21.0", "version: 0.22.0")
replace_once(
    status,
    "  - REVIEW-AURORA-M0-R4-SPK002-OWNER-TRUST-2026-08-07\n",
    "  - REVIEW-AURORA-M0-R4-SPK002-OWNER-TRUST-2026-08-07\n  - DOC-AURORA-M0-R4-ADR0008-ACCEPTANCE\n  - REVIEW-AURORA-M0-R4-ARCHITECTURE-DECISION-READINESS-RERUN-2026-08-07\n",
)
replace_once(
    status,
    "- **Current readiness gate:** ACRM R4 — Architecture/Decision Readiness — BLOCKED",
    "- **Current readiness gate:** ACRM R4 — Architecture/Decision Readiness — PASS; R5 NOT AUTHORIZED",
)
replace_once(
    status,
    "- **ADR-0008:** PROPOSED / EVIDENCE-READY — SPK-002 PASS/CLOSED; operator decision pending",
    "- **ADR-0008:** ACCEPTED — random ORK + authenticated external trust high-water + owner-only restore revalidation",
)
replace_once(
    status,
    "- **R4 verdict:** BLOCKED ONLY ON ADR-0008 OPERATOR DECISION — all required M0 architecture spikes are PASS/REVIEWED/CLOSED",
    "- **R4 verdict:** PASS — all 15 M0 architecture questions decided; all required spikes PASS/REVIEWED/CLOSED",
)
replace_once(
    status,
    "- **Accepted technical decisions:** ADR-0003, ADR-0004, ADR-0005, ADR-0006, ADR-0007 and ADR-0009; owner-root mechanism remains unresolved",
    "- **Accepted technical decisions:** ADR-0003, ADR-0004, ADR-0005, ADR-0006, ADR-0007, ADR-0008 and ADR-0009",
)
replace_once(
    status,
    "R4 has researched concrete dispositions for all M0 mechanism questions. ADR-0003, ADR-0004, ADR-0005, ADR-0006 and ADR-0007 are accepted. SPK-001 and SPK-002 have both completed with reviewed PASS evidence. ADR-0008 v0.2.0 is evidence-ready and remains proposed pending the operator decision:",
    "R4 has researched concrete dispositions for all M0 mechanism questions. ADR-0003 through ADR-0008 are accepted for their M0 scopes, ADR-0009 is accepted cross-horizon, and SPK-001/SPK-002 have both completed with reviewed PASS evidence:",
)
replace_once(
    status,
    "M0 ACRM R4 — Architecture/Decision Readiness: BLOCKED — documentary baseline d00cc1abfc2a41ac7e81e1f3478e188b3c5e9e52",
    "M0 ACRM R4 — Architecture/Decision Readiness: PASS — final rerun after ADR-0008 v0.2.0 acceptance",
)
replace_once(
    status,
    "R4 M0 decision coverage:         15/15 mapped; all required spikes closed; ADR-0008 evidence-ready/operator-decision pending",
    "R4 M0 decision coverage:         15/15 decided; all required spikes closed; no unresolved material M0 architecture choice",
)
replace_once(
    status,
    "Accepted technical decisions:   ADR-0003 / ADR-0004 / ADR-0005 / ADR-0006 / ADR-0007 / ADR-0009",
    "Accepted technical decisions:   ADR-0003 / ADR-0004 / ADR-0005 / ADR-0006 / ADR-0007 / ADR-0008 / ADR-0009",
)
old_gate = """SPK-002 has now completed successfully. Final hardened workflow `31219882882` passed the full S01–S12, recovery-classification, mutation-boundary and secret-hygiene evidence on Ubuntu and Windows. The independent review closed the spike as `PASS / REVIEWED / DECISION_INFORMED` and recommends ADR-0008 v0.2.0.

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
new_gate = """SPK-002 completed successfully and ADR-0008 v0.2.0 was explicitly accepted by the operator through `DOC-AURORA-M0-R4-ADR0008-ACCEPTANCE`.

The final R4 readiness rerun reviewed all 15 M0 architecture questions, all accepted ADRs, both required closed spikes, Blueprint compatibility and migration/rollback posture. No unresolved material M0 architecture choice remains.

Final verdict:

```text
R4 PASS
```

Carry-forward R6/R7 constraints remain explicit, including bounded Argon2id envelope parameters, target filesystem publication/directory-sync validation, mutation-boundary anomaly enforcement and no overclaim beyond the tested local fault/threat model. These are implementation/evidence obligations, not unresolved R4 architecture choices.

R5, Mission Contract, Microdesign and runtime implementation remain unauthorized.

## 8. Immediate next action

```text
M0 ACRM R4 PASS
→ STOP
→ await explicit operator authorization for M0 ACRM R5 — Contract Readiness
```"""
replace_once(status, old_gate, new_gate)

# WORKLOG
worklog = Path("docs/tracking/WORKLOG.md")
w = worklog.read_text(encoding="utf-8").rstrip()
heading = "## 2026-08-07 — ADR-0008 Acceptance and M0 R4 PASS"
if heading in w:
    raise SystemExit("WORKLOG already contains ADR-0008/R4 PASS entry")
w += f"""

{heading}

The operator explicitly accepted ADR-0008 v0.2.0, bound to blob `2a1497f8311ba9d04cd61f5025d7eae2af2fc57f` at canonical pre-acceptance main `35ce98fb2ddde16133c01a4da2f3545c8ae5e308`.

The final M0 R4 readiness rerun resolved the last open gate blocker. All 15 M0 architecture questions now have governing dispositions; ADR-0003 through ADR-0008 are accepted for their scopes, ADR-0009 is accepted cross-horizon, and SPK-001/SPK-002 are `PASS / EVIDENCE_COMPLETE / REVIEWED / DECISION_INFORMED / CLOSED`.

The final verdict is `M0 ACRM R4 PASS`. This does not authorize R5. Mission Contract, Microdesign, production implementation, promotion of spike code and all other later-gate work remain prohibited until separately authorized by the operator.
"""
worklog.write_text(w + "\n", encoding="utf-8")
