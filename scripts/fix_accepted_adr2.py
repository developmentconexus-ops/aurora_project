#!/usr/bin/env python3
from pathlib import Path

path = Path('docs/adr/0002-first-party-harness-development-kit.md')
text = path.read_text(encoding='utf-8')
replacements = {
    '  - proposed first-party AHDK policy\n  - proposed universal conformance policy\n  - proposed waiver boundary':
        '  - first-party AHDK policy\n  - universal conformance policy\n  - waiver boundary',
    '## 1. Status\n\n```text\nPROPOSED\n```\n\nThis ADR proposes an organizational and technical policy. It does not choose the first SDK language or authorize implementation.':
        '## 1. Status\n\n```text\nACCEPTED\n```\n\nThis ADR establishes an organizational and technical policy accepted in A0. It does not choose the first SDK language or authorize implementation.',
    '#### Assessment\n\nRecommended, subject to SPK-001.':
        '#### Assessment\n\nSelected as the A0 policy direction. Implementation readiness remains subject to SPK-001 and the applicable ACRM gates.',
    '## 7. Proposed decision': '## 7. Decision',
    '## 17. Implementation constraints if accepted': '## 17. Implementation constraints under the accepted policy',
    '## 19. Acceptance evidence\n\nBefore accepting this policy for implementation planning:':
        '## 19. Implementation-readiness evidence\n\nBefore advancing this accepted A0 policy into implementation planning:',
    'A0 may accept the policy direction without selecting language or implementing the kit.':
        'A0 accepts the policy direction without selecting a language or implementing the kit; implementation readiness remains a later gate.'
}
for old, new in replacements.items():
    if old not in text:
        raise SystemExit(f'expected text not found: {old[:80]!r}')
    text = text.replace(old, new, 1)
path.write_text(text, encoding='utf-8')

Path('.github/workflows/fix-accepted-adr2.yml').unlink()
Path('scripts/fix_accepted_adr2.py').unlink()
