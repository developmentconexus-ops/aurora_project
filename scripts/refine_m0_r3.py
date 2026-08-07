#!/usr/bin/env python3
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
base = ROOT / "docs/capabilities/CAP-SOVEREIGN-CORE"
spec = base / "SPEC.md"
test = base / "TEST-PLAN.md"
cov = base / "R3-COVERAGE.md"

s = spec.read_text(encoding="utf-8")
old = "The `accepted_state` envelope MUST be semantically versioned but R3 does not select a serialization language. Core treats domain-specific content as data; it validates the envelope, revision/identity invariants and M0 transition rules rather than inventing domain workflow semantics."
new = """The logical `accepted_state` envelope has these minimum semantic fields:\n\n- `state_schema_version` — logical version of the state envelope;\n- `state_kind` — stable symbolic type/category for the state value;\n- `state_summary` — bounded operator-readable description of the accepted state;\n- `state_payload` — optional structured project data whose concrete serialization and size limits remain R4/R5 concerns.\n\n`state_payload` is intentionally opaque to M0 Core domain logic. Core validates envelope/version/revision/identity rules but MUST treat payload content as project data: it cannot redefine Aurora/Project identity, authority, policy, canonical ownership or the transition protocol merely because it is persisted. This is the complete M0 payload boundary, not an unresolved request for R4 to invent state semantics."""
if old not in s:
    raise SystemExit("SPEC accepted_state anchor not found")
s = s.replace(old, new)

old = "Only owner-authorized M0 authority administration may create/revoke/supersede authority. Stale revision or invalid scope fails closed."
new = """Only owner-authorized M0 authority administration may create/revoke/supersede authority. Stale revision or invalid scope fails closed.\n\nThe owner root for bootstrap/recovery authority is the authenticated `OperatorIdentityRef` for Leandro, not an `AuthorityGrantRecord` recovered from state. This root is narrowly scoped to initialization, M0 authority administration, restore/recovery resolution and authority revalidation needed to recover safe Core control. It does not authorize external effects and cannot be delegated by M0."""
if old not in s:
    raise SystemExit("SPEC ChangeAuthority anchor not found")
s = s.replace(old, new)

old = "- only a new explicit owner revalidation/grant operation, or another R4-approved freshness proof satisfying this Spec, may return it to `VALID`.\n\nThis rule applies to restore, not ordinary restart from the current canonical store."
new = """- only a new explicit owner revalidation/grant operation, or another R4-approved freshness proof satisfying this Spec, may return it to `VALID`;\n- a restored grant MUST NOT authorize its own revalidation;\n- owner revalidation MUST create a new attributable authority-state revision rather than mutate historical restored authority in place.\n\nThe authenticated owner `OperatorIdentityRef` is therefore the recovery root that can perform this narrow revalidation even while restored delegated authority is blocked. A non-owner actor cannot use `REVALIDATION_REQUIRED` as an escalation path.\n\nThis rule applies to restore, not ordinary restart from the current canonical store."""
if old not in s:
    raise SystemExit("SPEC restore authority anchor not found")
s = s.replace(old, new)

old = "Future effect-plane capabilities must extend authority semantics through later gates without weakening M0's owner/revision/expiry/revocation rules."
new = """The owner bootstrap/recovery root described in §9.4 is a product-authority boundary, not an effect credential and not a bypass around restore/state validation.\n\nFuture effect-plane capabilities must extend authority semantics through later gates without weakening M0's owner/revision/expiry/revocation rules."""
if old not in s:
    raise SystemExit("SPEC authority boundary anchor not found")
s = s.replace(old, new)
spec.write_text(s, encoding="utf-8")

t = test.read_text(encoding="utf-8")
old = "| `T-PORT-012` | DOCUMENT_REVIEW, USER_JOURNEY | inspect export/backup governance | Leandro can inspect material result; package classification at least `SENSITIVE` |"
new = old + "\n| `T-PORT-013` | SECURITY_TEST, USER_JOURNEY | after F-07 restore, attempt authority revalidation first as non-owner and then as authenticated owner | non-owner denied; owner creates a new attributable authority-state revision and only then may current authority return to `VALID` if scope/time/conditions pass |"
if old not in t:
    raise SystemExit("TEST portability anchor not found")
t = t.replace(old, new)

old = "→ old authority does NOT silently become valid\n→ revalidation required\n```"
new = "→ old authority does NOT silently become valid\n→ revalidation required\n→ non-owner revalidation denied\n→ authenticated owner performs explicit revalidation\n→ new authority-state revision becomes governing only if current scope/time/conditions pass\n```"
if old not in t:
    raise SystemExit("TEST J-002 anchor not found")
t = t.replace(old, new)
test.write_text(t, encoding="utf-8")

c = cov.read_text(encoding="utf-8")
repls = {
"| `CAP-SOVEREIGN-CORE-REQ-032` | SPEC §7.2, §7.7, §12 | T-AUTH-001, T-AUTH-011 | `ALLOCATED` |":
"| `CAP-SOVEREIGN-CORE-REQ-032` | SPEC §7.2, §7.7, §9.4, §12 | T-AUTH-001, T-AUTH-011, T-PORT-013 | `ALLOCATED` |",
"| `CAP-SOVEREIGN-CORE-REQ-042` | SPEC §9.8, §13; THREAT §7 | T-PORT-005, T-PORT-006 | `ALLOCATED` |":
"| `CAP-SOVEREIGN-CORE-REQ-042` | SPEC §9.4, §9.8, §13; THREAT §7 | T-PORT-005, T-PORT-006, T-PORT-013 | `ALLOCATED` |",
}
for old, new in repls.items():
    if old not in c:
        raise SystemExit(f"coverage anchor not found: {old[:40]}")
    c = c.replace(old, new)
cov.write_text(c, encoding="utf-8")

print("refined M0 R3 Spec/Test/Coverage")
