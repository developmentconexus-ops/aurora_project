#!/usr/bin/env python3
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]

replacements = {
    "docs/design/SPK-AURORA-M0-SOVEREIGN-STORE-001.md": [
        (
            "id: SPK-AURORA-M0-SOVEREIGN-STORE-001\n",
            "id: DESIGN-AURORA-M0-SOVEREIGN-STORE-SPIKE-001\nspike_id: SPK-AURORA-M0-SOVEREIGN-STORE-001\n",
        ),
    ],
    "docs/design/SPK-AURORA-M0-OWNER-TRUST-002.md": [
        (
            "id: SPK-AURORA-M0-OWNER-TRUST-002\n",
            "id: DESIGN-AURORA-M0-OWNER-TRUST-SPIKE-002\nspike_id: SPK-AURORA-M0-OWNER-TRUST-002\n",
        ),
        (
            "  - SPK-AURORA-M0-SOVEREIGN-STORE-001\n",
            "  - DESIGN-AURORA-M0-SOVEREIGN-STORE-SPIKE-001\n",
        ),
    ],
    "docs/adr/0007-m0-sqlite-operational-store.md": [
        (
            "  - SPK-AURORA-M0-SOVEREIGN-STORE-001\n",
            "  - DESIGN-AURORA-M0-SOVEREIGN-STORE-SPIKE-001\n",
        ),
    ],
    "docs/adr/0008-m0-owner-root-recovery-trust.md": [
        (
            "  - SPK-AURORA-M0-OWNER-TRUST-002\n",
            "  - DESIGN-AURORA-M0-OWNER-TRUST-SPIKE-002\n",
        ),
    ],
}

for rel, reps in replacements.items():
    path = ROOT / rel
    text = path.read_text(encoding="utf-8")
    for old, new in reps:
        count = text.count(old)
        if count != 1:
            raise SystemExit(f"expected one occurrence in {rel}, found {count}: {old.strip()}")
        text = text.replace(old, new, 1)
    path.write_text(text, encoding="utf-8")

validator_path = ROOT / "scripts/validate_m0_r4_documentary_package.py"
validator = validator_path.read_text(encoding="utf-8")
old = '''spikes = {
    "SPK-AURORA-M0-SOVEREIGN-STORE-001": "docs/design/SPK-AURORA-M0-SOVEREIGN-STORE-001.md",
    "SPK-AURORA-M0-OWNER-TRUST-002": "docs/design/SPK-AURORA-M0-OWNER-TRUST-002.md",
}
for sid, path in spikes.items():
    txt = read(path)
    if f"id: {sid}" not in txt:
        errors.append(f"wrong/missing spike id: {path}")
'''
new = '''spikes = {
    "SPK-AURORA-M0-SOVEREIGN-STORE-001": ("DESIGN-AURORA-M0-SOVEREIGN-STORE-SPIKE-001", "docs/design/SPK-AURORA-M0-SOVEREIGN-STORE-001.md"),
    "SPK-AURORA-M0-OWNER-TRUST-002": ("DESIGN-AURORA-M0-OWNER-TRUST-SPIKE-002", "docs/design/SPK-AURORA-M0-OWNER-TRUST-002.md"),
}
for sid, (doc_id, path) in spikes.items():
    txt = read(path)
    if f"id: {doc_id}" not in txt:
        errors.append(f"wrong/missing spike document id: {path}")
    if f"spike_id: {sid}" not in txt:
        errors.append(f"wrong/missing conceptual spike id: {path}")
'''
if validator.count(old) != 1:
    raise SystemExit("validator spike block precondition not found exactly once")
validator = validator.replace(old, new, 1)
validator = validator.replace(
    'spk2 = read(spikes["SPK-AURORA-M0-OWNER-TRUST-002"])',
    'spk2 = read(spikes["SPK-AURORA-M0-OWNER-TRUST-002"][1])',
    1,
)
validator = validator.replace(
    'for path in [p for _,p in adrs.items()] + [p for _,p in spikes.items()] + [r[1] for r in reports]:',
    'for path in [p for _,p in adrs.items()] + [entry[1] for entry in spikes.values()] + [r[1] for r in reports]:',
    1,
)
validator_path.write_text(validator, encoding="utf-8")

print("fixed M0 R4 spike document IDs")
