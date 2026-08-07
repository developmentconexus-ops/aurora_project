#!/usr/bin/env python3
from pathlib import Path
import re
import subprocess

ACCEPTANCE_ID = 'DOC-AURORA-A0-OPERATOR-ACCEPTANCE'
ACCEPTED_AT = '2026-08-06'


def promote_frontmatter(path: Path) -> None:
    text = path.read_text(encoding='utf-8')
    if not text.startswith('---\n'):
        raise SystemExit(f'{path}: missing frontmatter')
    marker = text.find('\n---\n', 4)
    if marker < 0:
        raise SystemExit(f'{path}: unterminated frontmatter')
    fm = text[:marker + 1]
    body = text[marker + 1:]
    if re.search(r'(?m)^status: accepted$', fm):
        return
    fm, count = re.subn(
        r'(?m)^status: proposed$',
        f'status: accepted\naccepted_at: {ACCEPTED_AT}\nacceptance_evidence: {ACCEPTANCE_ID}',
        fm,
        count=1,
    )
    if count != 1:
        raise SystemExit(f'{path}: expected status: proposed in frontmatter')
    fm = re.sub(r'(?m)^last_reviewed: .*$', f'last_reviewed: {ACCEPTED_AT}', fm, count=1)
    path.write_text(fm + body, encoding='utf-8')


blueprints = sorted(Path('docs/product/blueprint').glob('[0-9][0-9]-*.md'))
if len(blueprints) != 15:
    raise SystemExit(f'expected 15 Blueprint sections, found {len(blueprints)}')

promote_paths = [
    Path('docs/DOCUMENTATION-MAP.md'),
    Path('docs/product/README.md'),
    Path('docs/product/CAPABILITY-REALIZATION-METHOD.md'),
    Path('docs/product/REQUIREMENTS-TRACEABILITY.md'),
    Path('docs/adr/README.md'),
    Path('docs/adr/0001-aurora-owned-contract-model.md'),
    Path('docs/adr/0002-first-party-harness-development-kit.md'),
    *blueprints,
]

for path in promote_paths:
    promote_frontmatter(path)

# Accepted constitutional requirement set.
req = Path('docs/product/REQUIREMENTS-TRACEABILITY.md')
text = req.read_text(encoding='utf-8')
text = text.replace(
    'This document converts the proposed Product Blueprint into explicit constitutional requirements.',
    'This document converts the accepted A0 Product Blueprint into explicit constitutional requirements.'
)
text = text.replace(
    'PROPOSED\n→ derived from proposed A0 Blueprint; not governing implementation yet',
    'ACCEPTED\n→ accepted A0 constitutional requirement; governs future applicability analysis but does not itself authorize implementation'
)
req.write_text(text, encoding='utf-8')

# A0 checklist closes, without authorizing the next execution gate.
acrm = Path('docs/product/CAPABILITY-REALIZATION-METHOD.md')
text = acrm.read_text(encoding='utf-8')
if '## 25. A0 acceptance checklist' in text:
    before, after = text.split('## 25. A0 acceptance checklist', 1)
    parts = after.split('\n---\n', 1)
    checklist = parts[0].replace('- [ ]', '- [x]')
    text = before + '## 25. A0 acceptance checklist' + checklist
    if len(parts) == 2:
        text += '\n---\n' + parts[1]
acrm.write_text(text, encoding='utf-8')

# Product index authority state.
product_index = Path('docs/product/README.md')
text = product_index.read_text(encoding='utf-8')
text = text.replace(
    'All constitutional sources in PR #1 remain `proposed` until explicit operator acceptance and merge.',
    'All A0 constitutional sources were explicitly accepted by the operator on 2026-08-06. Merge is a separate repository action and was authorized by the same operator decision.'
)
old_block = '''A0 documentation/research work: AUTHORIZED
A0 content: IN_REVIEW
ADRs: PROPOSED
Architecture Spikes: NOT AUTHORIZED
Aurora runtime implementation: PROHIBITED
AHDK implementation: PROHIBITED
MNFS integration: PROHIBITED'''
new_block = '''A0 baseline: ACCEPTED
Constitutional sources: ACCEPTED
ADR-0001 / ADR-0002: ACCEPTED
PR #1 merge: AUTHORIZED
Architecture Spikes: NOT AUTHORIZED
Aurora runtime implementation: PROHIBITED
AHDK implementation: PROHIBITED
MNFS integration: PROHIBITED'''
if old_block not in text:
    raise SystemExit('product index authority block changed unexpectedly')
product_index.write_text(text.replace(old_block, new_block), encoding='utf-8')

# ADR index lifecycle states.
adr_index = Path('docs/adr/README.md')
text = adr_index.read_text(encoding='utf-8')
text = re.sub(r'(\| \[0001\][^\n]*\|) proposed (\|)', r'\1 accepted \2', text)
text = re.sub(r'(\| \[0002\][^\n]*\|) proposed (\|)', r'\1 accepted \2', text)
adr_index.write_text(text, encoding='utf-8')

# Tracking decision index.
decisions = Path('docs/tracking/DECISIONS.md')
text = decisions.read_text(encoding='utf-8')
lines = []
for line in text.splitlines():
    if line.startswith('| D-') and line.rstrip().endswith('| proposed |'):
        line = line[:-len('| proposed |')] + '| accepted |'
    lines.append(line)
decisions.write_text('\n'.join(lines) + '\n', encoding='utf-8')

# Golden Proof reviewer verdict plus operator verdict.
evaluation = Path('docs/acceptance/2026-08-06-gp-a0-fresh-001-evaluation.md')
text = evaluation.read_text(encoding='utf-8')
text = text.replace('operator_verdict: PENDING', 'operator_verdict: ACCEPTED', 1)
if '## 9. Operator verdict' not in text:
    text += '''\n\n## 9. Operator verdict\n\nOn 2026-08-06, after reviewing the A0 reading package, Leandro explicitly approved all four pending decisions and stated that the idea is well structured.\n\n```text\nA0 baseline: ACCEPTED\nADR-0001: ACCEPTED\nADR-0002: ACCEPTED\nPR #1 merge: AUTHORIZED\n```\n\nThis approval does not authorize runtime implementation, Architecture Spike execution, AHDK implementation, MNFS integration or a stack choice. The next product gate is first Product Milestone selection followed by ACRM R0.\n'''
evaluation.write_text(text, encoding='utf-8')

# Acceptance index.
acceptance_index = Path('docs/acceptance/README.md')
text = acceptance_index.read_text(encoding='utf-8')
text = text.replace(
    'Protocol ready\nAuthoring-session dry run non-qualifying\nIndependent execution pending\nOperator verdict pending',
    'Protocol ready\nAuthoring-session dry run non-qualifying\nIndependent execution PASS — GP-A0-FRESH-001, 100/100, zero hard failures\nOperator verdict ACCEPTED — 2026-08-06'
)
if 'A0 Operator Acceptance' not in text:
    text = text.replace(
        '- [A0 Fresh-Session Documentation Golden Proof](2026-08-06-a0-fresh-session-golden-proof.md)',
        '- [A0 Fresh-Session Documentation Golden Proof](2026-08-06-a0-fresh-session-golden-proof.md)\n- [GP-A0-FRESH-001 Evaluation](2026-08-06-gp-a0-fresh-001-evaluation.md)\n- [A0 Operator Acceptance](2026-08-06-a0-operator-acceptance.md)'
    )
acceptance_index.write_text(text, encoding='utf-8')

# Agent bootstrap.
agents = Path('AGENTS.md')
text = agents.read_text(encoding='utf-8')
old = '''Phase: A0 — Product, Discovery and Architecture Baseline
Working branch: docs/architecture-baseline
Draft PR: #1
A0 content: IN REVIEW / REMEDIATION
Aurora Core implementation: PROHIBITED
AHDK implementation: PROHIBITED
Architecture Spikes: PROPOSED, NOT AUTHORIZED
Stack selection: NOT PERFORMED
MNFS integration: PROHIBITED
Automatic merge: NOT AUTHORIZED'''
new = '''Phase: A0 — ACCEPTED Product, Discovery and Architecture Baseline
Repository disposition: PR #1 merge AUTHORIZED; consult STATUS for completion
A0 constitutional content: ACCEPTED
ADR-0001 / ADR-0002: ACCEPTED
First Product Milestone: NOT YET SELECTED
Aurora Core implementation: PROHIBITED
AHDK implementation: PROHIBITED
Architecture Spikes: PROPOSED, NOT AUTHORIZED
Stack selection: NOT PERFORMED
MNFS integration: PROHIBITED'''
if old not in text:
    raise SystemExit('AGENTS current-phase block changed unexpectedly')
text = text.replace(old, new)
text = text.replace(
    'Documentation, research, adversarial review and validation are authorized.',
    'A0 is accepted. Product Milestone selection and ACRM R0 preparation are authorized next; implementation remains prohibited until the applicable gates and a separate execution authorization pass.'
)
agents.write_text(text, encoding='utf-8')

# Root README.
readme = Path('README.md')
text = readme.read_text(encoding='utf-8')
old = '''Phase: A0 — Product, Discovery and Architecture Baseline
Branch: docs/architecture-baseline
Draft PR: #1
Runtime implementation: PROHIBITED
Architecture Spikes: NOT AUTHORIZED
Stack selection: NOT PERFORMED'''
new = '''A0 baseline: ACCEPTED — 2026-08-06
ADR-0001 / ADR-0002: ACCEPTED
Next product gate: select first Product Milestone → begin ACRM R0
Runtime implementation: PROHIBITED
Architecture Spikes: NOT AUTHORIZED
Stack selection: NOT PERFORMED'''
if old not in text:
    raise SystemExit('README state block changed unexpectedly')
text = text.replace(old, new)
text = text.replace(
    'A0 está reconstruindo e validando a base documental completa. A existência de arquivos não significa aceitação nem autorização para implementar.',
    'A0 foi explicitamente aceita pelo operador após validação mecânica, revisão adversarial e Golden Proof independente. Essa aceitação estabelece a constituição do produto; não autoriza implementação. O próximo passo é selecionar o primeiro Product Milestone e iniciar ACRM R0.'
)
readme.write_text(text, encoding='utf-8')

# Exact pre-merge status.
Path('docs/tracking/STATUS.md').write_text('''---
id: DOC-AURORA-STATUS
title: Aurora Project Status
document_type: project_status
form: reference
authority: tracking
status: current
version: 0.6.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - current Aurora project phase
  - current authorization boundary
  - current blockers and immediate next action
related:
  - DOC-AURORA-DOCUMENTATION-MAP
  - DOC-AURORA-PRODUCT-INDEX
  - DOC-AURORA-ROADMAP
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DOC-AURORA-A0-OPERATOR-ACCEPTANCE
  - DOC-AURORA-GP-A0-FRESH-001-EVALUATION
last_reviewed: 2026-08-06
---

# Aurora Project Status

## 1. Current summary

- **Project:** Projeto Aurora
- **Phase:** A0 — ACCEPTED Product, Discovery and Architecture Baseline
- **A0 accepted:** 2026-08-06 by operator
- **Product Blueprint:** 15 modular accepted constitutional sections plus generated aggregate
- **Constitutional requirements:** 294 accepted A0 requirements
- **ADR-0001:** ACCEPTED
- **ADR-0002:** ACCEPTED
- **Independent Fresh-Session Golden Proof:** PASS — 100/100, zero hard failures
- **PR #1:** merge explicitly AUTHORIZED; pending repository merge at this tracking revision
- **First Product Milestone:** not yet selected
- **Stack decisions:** none
- **Runtime implementation:** not started and not authorized

## 2. Operator decision

The operator reviewed the required A0 package and explicitly approved all pending decisions, stating that the idea is well structured.

```text
A0 baseline: ACCEPTED
ADR-0001: ACCEPTED
ADR-0002: ACCEPTED
PR #1 merge: AUTHORIZED
```

Evidence: `docs/acceptance/2026-08-06-a0-operator-acceptance.md`.

## 3. Current authorization boundary

```text
A0 baseline:                    ACCEPTED
ADR-0001 / ADR-0002:           ACCEPTED
PR #1 merge:                   AUTHORIZED / PENDING
First Product Milestone choice: AUTHORIZED / NOT YET MADE
ACRM R0:                        NOT STARTED; begins after milestone selection
Architecture Spike planning:    NOT STARTED
Architecture Spike execution:   PROHIBITED
Capability implementation:      PROHIBITED
Aurora Core implementation:     PROHIBITED
AHDK implementation:            PROHIBITED
MNFS integration:               PROHIBITED
Stack selection:                NOT PERFORMED
```

A0 acceptance does not skip any ACRM gate. No runtime work is authorized by implication.

## 4. Deliberately open technical decisions

- Aurora Core language and deployment shape;
- first AHDK language/toolchain;
- schema representation per boundary;
- local RPC binding;
- exact MCP/A2A mapping;
- durable execution engine;
- policy engine;
- workload/device identity mechanism;
- operational state/event storage;
- Artifact/Evidence Store;
- event transport and telemetry backend;
- memory mechanism mix;
- first reference Harness runtime;
- first real engineering Harness;
- first Product Milestone and Mission Contract.

These remain future ACRM R3/R4 decisions unless the selected milestone requires them earlier through an explicit gate.

## 5. Immediate next action

```text
complete authorized merge of PR #1
→ select the first Product Milestone
→ begin ACRM R0 — Constitutional Baseline for that milestone
```

The next action is not implementation.
''', encoding='utf-8')

# Worklog decision record.
worklog = Path('docs/tracking/WORKLOG.md')
text = worklog.read_text(encoding='utf-8').rstrip()
if '## 2026-08-06 — A0 operator acceptance' not in text:
    text += '''\n\n## 2026-08-06 — A0 operator acceptance\n\nAfter the independent Fresh-Session Golden Proof passed 100/100 with zero hard failures, Leandro reviewed the A0 authorization reading package and explicitly approved all pending decisions, stating that the idea is well structured.\n\nDecisions:\n\n- A0 Product/Discovery/Architecture baseline: ACCEPTED;\n- ADR-0001: ACCEPTED;\n- ADR-0002: ACCEPTED;\n- PR #1 merge: AUTHORIZED.\n\nThe decision was recorded in `docs/acceptance/2026-08-06-a0-operator-acceptance.md`. Architecture Spikes and all runtime/AHDK/MNFS implementation remain prohibited. The next product step after repository merge is to select the first Product Milestone and begin ACRM R0.\n'''
worklog.write_text(text + '\n', encoding='utf-8')

# Regenerate read-only projections from promoted sources.
subprocess.run(['python', 'scripts/generate_docs.py'], check=True)

# One-time promotion machinery must not remain in the accepted repository.
Path('.github/workflows/promote-a0.yml').unlink()
Path('scripts/promote_a0.py').unlink()
