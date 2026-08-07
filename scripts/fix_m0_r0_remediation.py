#!/usr/bin/env python3
from pathlib import Path


def replace_once(path: str, old: str, new: str) -> None:
    p = Path(path)
    text = p.read_text(encoding="utf-8")
    if old not in text:
        raise SystemExit(f"expected text not found in {path}: {old[:120]!r}")
    p.write_text(text.replace(old, new, 1), encoding="utf-8")


def append_once(path: str, marker: str, addition: str) -> None:
    p = Path(path)
    text = p.read_text(encoding="utf-8")
    if marker in text:
        raise SystemExit(f"marker already present in {path}: {marker}")
    p.write_text(text.rstrip() + "\n\n" + addition.rstrip() + "\n", encoding="utf-8")


# ---------------------------------------------------------------------------
# F03 — remove mutable-current-state ownership from durable guidance/constitution.
# ---------------------------------------------------------------------------
replace_once(
    "AGENTS.md",
    """## 3. Current phase and authorization

```text
Phase: A0 — ACCEPTED Product, Discovery and Architecture Baseline
Repository disposition: PR #1 merge AUTHORIZED; consult STATUS for completion
A0 constitutional content: ACCEPTED
ADR-0001 / ADR-0002: ACCEPTED
First Product Milestone: NOT YET SELECTED
Aurora Core implementation: PROHIBITED
AHDK implementation: PROHIBITED
Architecture Spikes: PROPOSED, NOT AUTHORIZED
Stack selection: NOT PERFORMED
MNFS integration: PROHIBITED
```

A0 is accepted. Product Milestone selection and ACRM R0 preparation are authorized next; implementation remains prohibited until the applicable gates and a separate execution authorization pass.
""",
    """## 3. Current phase and authorization

Mutable coordination state has exactly one owner:

```text
docs/tracking/STATUS.md
```

A fresh session MUST read `STATUS.md` for the selected Product Milestone, current ACRM gate, blockers, authorizations, prohibitions and exact next action. This file intentionally does not duplicate those mutable values.

Stable accepted baseline:

```text
A0 constitutional baseline: ACCEPTED / MERGED
ADR-0001 / ADR-0002: ACCEPTED
Stack selection from A0: NOT PERFORMED
Aurora Core implementation: NOT AUTHORIZED BY A0
AHDK implementation: NOT AUTHORIZED BY A0
Architecture Spike execution: NOT AUTHORIZED BY A0
MNFS integration: NOT AUTHORIZED BY A0
```

After A0, every readiness transition remains explicit. Selection of a Product Milestone authorizes only the readiness work recorded in `STATUS`; completing one ACRM gate does not authorize the next gate, and no implementation follows from file existence or absence of a prohibition.
""",
)

replace_once(
    "README.md",
    """## Estado atual

```text
A0 baseline: ACCEPTED — 2026-08-06
ADR-0001 / ADR-0002: ACCEPTED
Next product gate: select first Product Milestone → begin ACRM R0
Runtime implementation: PROHIBITED
Architecture Spikes: NOT AUTHORIZED
Stack selection: NOT PERFORMED
```

A0 foi explicitamente aceita pelo operador após validação mecânica, revisão adversarial e Golden Proof independente. Essa aceitação estabelece a constituição do produto; não autoriza implementação. O próximo passo é selecionar o primeiro Product Milestone e iniciar ACRM R0.
""",
    """## Estado atual

O estado mutável do programa não é duplicado neste README. A fonte canônica para Product Milestone selecionado, gate ACRM atual, blockers, autorizações, proibições e próxima ação é:

```text
docs/tracking/STATUS.md
```

A baseline durável já estabelecida é:

```text
A0 baseline: ACCEPTED / MERGED — 2026-08-06
ADR-0001 / ADR-0002: ACCEPTED
Stack selection by A0: NOT PERFORMED
Runtime implementation: NOT AUTHORIZED BY A0
Architecture Spike execution: NOT AUTHORIZED BY A0
```

A aceitação de A0 estabelece a constituição do produto; cada milestone, gate, decisão técnica, spike e execução posterior continua exigindo sua própria promoção e autorização explícita. Consulte `STATUS.md` em vez de inferir o próximo passo deste resumo.
""",
)

replace_once(
    "CONTRIBUTING.md",
    """## Estado desta baseline

A branch `docs/architecture-baseline` é uma proposta para revisão do operador. Nenhum arquivo nela se torna aceito apenas por existir.
""",
    """## Estado e branches

`main` é a branch canônica do repositório. Branches não canônicas são superfícies de proposta/revisão e não promovem conteúdo por existência, commit, CI verde ou permissão de escrita.

A branch `docs/architecture-baseline` pode ser reutilizada pelo workflow documental para regeneração e revisão de mudanças constitucionais, mas seu conteúdo só se torna canônico após a autoridade aplicável aceitar a revisão e a mudança ser integrada a `main`.

O gate, as autorizações, os blockers e a próxima ação correntes são sempre lidos de `docs/tracking/STATUS.md`.
""",
)

replace_once("docs/DOCUMENTATION-MAP.md", "version: 0.2.0", "version: 0.2.1")
replace_once(
    "docs/DOCUMENTATION-MAP.md",
    """## 1. Current phase

```text
A0 — Product, Discovery and Architecture Baseline
```

A0 exists to transform the initial Aurora discovery dialogue into a complete, research-backed and resumable product constitution before runtime implementation begins.

Current work includes:

- Product Blueprint definition;
- preservation of origin, examples and decision reasoning;
- focused research with primary-source manifests;
- proposed architecture decisions;
- constitutional requirement derivation;
- a Blueprint-to-build realization method;
- roadmap and Golden Proof definition;
- mechanical and adversarial documentation validation.

Current work does **not** authorize:

- Aurora Core implementation;
- AHDK implementation;
- architecture spike execution;
- stack selection;
- MNFS integration;
- device or laboratory control.

The authoritative current boundary is always recorded in [`docs/tracking/STATUS.md`](tracking/STATUS.md).
""",
    """## 1. Accepted baseline and current-state owner

A0 transformed the initial Aurora discovery dialogue into the accepted, research-backed and resumable product constitution. The A0 baseline was explicitly accepted and merged on 2026-08-06.

A0 established:

- the Product Blueprint and canonical ownership model;
- preservation of origin, examples and decision reasoning;
- focused research with primary-source manifests;
- accepted ADR-0001 and ADR-0002;
- constitutional requirements and traceability;
- the Blueprint-to-build Capability Realization Method;
- roadmap and Golden Proof definitions;
- mechanical, adversarial and fresh-session validation.

This Documentation Map does **not** own mutable coordination state after A0. The authoritative source for the selected Product Milestone, current readiness gate, blockers, authorizations, prohibitions and exact next action is always [`docs/tracking/STATUS.md`](tracking/STATUS.md).

A0 acceptance does not, by itself, authorize Aurora Core/AHDK/MNFS implementation, Architecture Spike execution, stack selection or later ACRM gates. Those permissions must be explicit in `STATUS.md` and their owning evidence.
""",
)

replace_once("docs/product/README.md", "version: 0.2.0", "version: 0.2.1")
replace_once(
    "docs/product/README.md",
    """## 8. Current authority state

All A0 constitutional sources were explicitly accepted by the operator on 2026-08-06. Merge is a separate repository action and was authorized by the same operator decision.

```text
A0 baseline: ACCEPTED
Constitutional sources: ACCEPTED
ADR-0001 / ADR-0002: ACCEPTED
PR #1 merge: AUTHORIZED
Architecture Spikes: NOT AUTHORIZED
Aurora runtime implementation: PROHIBITED
AHDK implementation: PROHIBITED
MNFS integration: PROHIBITED
```

A merged file is not automatically accepted unless its lifecycle and operator decision say so. A0 acceptance is a separate explicit gate.
""",
    """## 8. Accepted baseline and current-state handoff

All A0 constitutional sources, ADR-0001 and ADR-0002 were explicitly accepted by the operator on 2026-08-06, and the accepted A0 package was subsequently merged to `main`.

This product index does not own mutable readiness state. For the selected Product Milestone, current ACRM gate, blockers, authorizations, prohibitions and exact next action, read:

```text
docs/tracking/STATUS.md
```

Stable governance remains:

- merge does not create acceptance by itself;
- accepted A0 intent does not select a stack;
- accepted ADRs govern only their stated decision scope;
- a selected milestone does not authorize later gates by implication;
- Architecture Spike execution and implementation require their own explicit authority.
""",
)

# Blueprint 01: convert stale "current phase" snapshot into durable authorization rule.
replace_once("docs/product/blueprint/01-product-vision.md", "version: 0.2.0", "version: 0.2.1")
replace_once(
    "docs/product/blueprint/01-product-vision.md",
    """## 1.19 Current product phase

```text
A0 — Product, Discovery and Architecture Baseline
```

Current work is documentation, research, adversarial review and architectural clarification.

Implementation remains prohibited until:

- all fifteen Blueprint sections are reviewed;
- A0 requirements and traceability are accepted;
- ADR status is explicitly decided;
- next milestone is promoted to readiness;
- architecture spikes and implementation plan are separately authorized.
""",
    """## 1.19 Authorization boundary after A0

A0 was explicitly accepted and merged on 2026-08-06. That acceptance establishes Aurora's constitutional product direction; it does not create runtime authority or select implementation mechanisms.

Mutable coordination state after A0 belongs to `docs/tracking/STATUS.md`, including:

- selected Product Milestone;
- current ACRM gate;
- active blockers/findings;
- authorized versus prohibited work;
- exact next action.

Material implementation can advance only through the applicable Capability Realization gates and separate execution authorization. Research, examples, accepted constitutional direction or the existence of a candidate spike/framework never substitute for that authority.
""",
)

# ---------------------------------------------------------------------------
# F01 — resolve Blueprint 14 executable-horizon ambiguity and complete M0 anatomy.
# ---------------------------------------------------------------------------
replace_once("docs/product/blueprint/14-capability-roadmap.md", "version: 0.2.0", "version: 0.3.0")
replace_once(
    "docs/product/blueprint/14-capability-roadmap.md",
    """## 14.5 Required milestone anatomy

Every Product Milestone must define:

```text
Outcome
Operator-visible value
Risk retired
Entry criteria
Capabilities involved
Architecture spikes
Golden Proof
Evidence requirements
Exit criteria
Telemetry baseline
Non-goals
Dependencies
Replan triggers
Promotion/authority boundary
```

A milestone cannot close because files exist or a demo looks plausible.
""",
    """## 14.5 Required milestone anatomy

Every Product Milestone **promoted into the current executable horizon** must define:

```text
Outcome
Operator-visible value
Risk retired
Entry criteria
Capabilities involved
Architecture spikes
Golden Proof
Evidence requirements
Exit criteria
Telemetry baseline
Non-goals
Dependencies
Replan triggers
Promotion/authority boundary
```

Directional future milestones may intentionally leave executable fields less specific while preserving enough outcome, risk, proof direction and boundaries to protect sequencing. Before such a milestone can pass its own R0, it must be expanded to the complete anatomy above without silently importing technical commitment from a distant horizon.

For the selected/current milestone, a missing required field is a constitutional readiness defect, not something R1/R2 should invent on behalf of the roadmap owner.

A milestone cannot close because files exist or a demo looks plausible.
""",
)

old_m0 = """# 14.8 M0 — Sovereign Core Walking Skeleton

## Outcome

A minimal Aurora Core persists identity, project state, authority snapshot and one interaction lifecycle across restart without depending on an external model or harness as authority.

## Operator-visible value

Leandro initializes Aurora, records a project state, restarts the process and receives the exact current state and permitted next action.

## Risk retired

```text
Aurora is merely a running session; restart destroys identity and state.
```

## Entry criteria

- A0 accepted;
- Core boundaries approved;
- minimal domain/entity spec;
- storage and language spikes complete enough for one local implementation;
- backup/restore and migration strategy for the slice.

## Capabilities

- sovereign identity;
- project registry;
- operational state;
- authority snapshot;
- event/audit minimum;
- CLI or simple interface.

## Golden Proof

```text
initialize Aurora instance
→ create project
→ record accepted state and next action
→ terminate all Aurora processes
→ start fresh process
→ recover same identities/state
→ reject an invalid transition
→ export and restore state
```

## Evidence

- state hashes/IDs;
- restart receipt;
- invalid transition test;
- backup/restore result;
- no transcript dependency.

## Non-goals

- conversational memory;
- model routing;
- harness registry;
- voice;
- multi-device;
- cloud;
- physical devices.

## Replan triggers

- store cannot preserve required state simply;
- domain model proves too broad for slice;
- operational burden exceeds single-user baseline.
"""

new_m0 = """# 14.8 M0 — Sovereign Core Walking Skeleton

## Outcome

A minimal Aurora Core persists identity, project state, authority snapshot and one interaction lifecycle across restart without depending on an external model or harness as authority.

## Operator-visible value

Leandro initializes Aurora, records a project state, restarts the process and receives the exact current state and permitted next action.

## Risk retired

```text
Aurora is merely a running session; restart destroys identity and state.
```

## Entry criteria

- A0 accepted;
- Core boundaries approved;
- minimal domain/entity spec;
- storage and language spikes complete enough for one local implementation;
- backup/restore and migration strategy for the slice.

## Capabilities involved

- sovereign identity;
- project registry;
- operational state;
- authority snapshot;
- event/audit minimum;
- CLI or simple interface.

## Architecture spikes

M0 requires evidence for the two implementation-blocking uncertainty classes already named by its entry criteria:

- local operational-state persistence/recovery, including the export/restore and migration needs of this slice;
- Core implementation language/runtime fit for the smallest local implementation.

The roadmap does **not** select spike IDs, candidate technologies, procedures or winners. R1–R4 must determine exact applicability and whether an existing portfolio item is reusable or a narrower M0 spike is required. Every spike still requires its own authorization before execution.

## Golden Proof

```text
initialize Aurora instance
→ create project
→ record accepted state and next action
→ terminate all Aurora processes
→ start fresh process
→ recover same identities/state
→ reject an invalid transition
→ export and restore state
```

## Evidence requirements

- state hashes/IDs;
- restart receipt;
- invalid transition test;
- backup/restore result;
- no transcript dependency.

## Exit criteria

- the complete M0 Golden Proof passes end to end against one fixed accepted revision;
- Aurora identity, project identity, accepted project state, authority snapshot and next action survive the process restart represented by the proof;
- the invalid transition is rejected without being accepted as current state;
- export and restore reproduce the M0 state required by the proof;
- evidence demonstrates that transcript, external model and Harness state are not the authority required for recovery;
- limitations and residual risks are explicit;
- operator-visible value is demonstrated and M0 receives its R8 closeout verdict rather than being inferred from component completion.

## Telemetry baseline

M0 records only the structured signals needed to explain and verify its own walking skeleton:

- stable Aurora/project/proof-run correlation identities;
- attempted and accepted/rejected lifecycle transitions;
- restart/recovery boundary and result;
- export/restore attempt and result;
- integrity/hash references used by the M0 evidence;
- classified failure reason when a proof step cannot complete.

No telemetry backend, event transport, schema technology or observability framework is selected by this baseline.

## Non-goals

- conversational memory;
- model routing;
- harness registry;
- voice;
- multi-device;
- cloud;
- physical devices.

## Dependencies

M0 depends on the prerequisites expressed by its entry criteria and the accepted constitutional owners for identity, project/state, authority, sovereignty, architecture and evidence.

M0 does **not** depend on M1 conversational memory, M2 Capability Registry/AHDK, MNFS integration, cloud deployment or physical-device capability. Those remain later roadmap concerns.

## Replan triggers

- store cannot preserve required state simply;
- domain model proves too broad for slice;
- operational burden exceeds single-user baseline.

## Promotion/authority boundary

M0 selection authorizes readiness analysis only to the extent recorded in `docs/tracking/STATUS.md`.

- R0–R6 may refine applicability, verifiable requirements, Capability design, technical decisions and the exact implementation contract, but MUST NOT silently change this outcome, named risk, Golden Proof direction or non-goals;
- completing one ACRM gate does not authorize the next;
- language, storage, runtime, topology, schema, event/audit and backup mechanisms remain technical decisions for later applicable gates;
- Architecture Spike execution and Aurora Core implementation require separate explicit authorization;
- M0 becomes accepted only through R8 closeout with end-to-end evidence and the required operator verdict.
"""
replace_once("docs/product/blueprint/14-capability-roadmap.md", old_m0, new_m0)

# Blueprint 15: keep the A0 acceptance rule but stop claiming mutable current gate.
replace_once("docs/product/blueprint/15-documentation-research-governance.md", "version: 0.2.0", "version: 0.2.1")
replace_once(
    "docs/product/blueprint/15-documentation-research-governance.md",
    """## 15.31 Current A0 gate

A0 remains `IN_REVIEW` until:

- all fifteen sections are complete and reviewed;
- discovery coverage is full;
- focused research program is published enough to support proposed ADRs;
- Capability Realization Method and traceability exist;
- aggregate/index/read paths are current;
- adversarial and fresh-session review pass;
- Leandro explicitly accepts baseline and ADR status.

Implementation remains prohibited.
""",
    """## 15.31 A0 acceptance rule and post-A0 current-state ownership

A0 acceptance required:

- all fifteen sections complete and reviewed;
- full discovery coverage;
- focused research sufficient to support the A0 decision set;
- Capability Realization Method and traceability;
- current aggregate/index/read paths;
- adversarial and fresh-session review;
- explicit Leandro acceptance of the baseline and ADR status.

Those conditions were satisfied and A0 was explicitly accepted on 2026-08-06, then merged to `main`.

After A0, mutable coordination state is not owned by this constitutional section. `docs/tracking/STATUS.md` owns the selected Product Milestone, current ACRM gate, blockers, authorization boundary and exact next action.

A0 acceptance never authorizes later gates, Architecture Spike execution or implementation by implication. Each transition still requires the authority defined by the Capability Realization Method and current `STATUS.md`.
""",
)

# Requirements Traceability: align accepted lifecycle and executable-horizon anatomy.
replace_once("docs/product/REQUIREMENTS-TRACEABILITY.md", "version: 0.1.0", "version: 0.2.0")
replace_once(
    "docs/product/REQUIREMENTS-TRACEABILITY.md",
    "| AUR-REQ-RDM-004 | Every Product Milestone MUST define outcome, value, risk, entry, capabilities, spikes, Golden Proof, evidence, exit, telemetry, non-goals, dependencies and replan triggers. | §14.5 | milestone schema/check |",
    "| AUR-REQ-RDM-004 | Every Product Milestone promoted into the current executable horizon MUST define outcome, value, risk, entry, capabilities, spikes, Golden Proof, evidence, exit, telemetry, non-goals, dependencies, replan triggers and promotion/authority boundary. | §14.5 | milestone schema/check |",
)
replace_once(
    "docs/product/REQUIREMENTS-TRACEABILITY.md",
    """## 17. Initial totals and allocation status

Current proposed constitutional requirements:
""",
    """## 17. Accepted A0 totals and allocation baseline

Accepted A0 constitutional requirements:
""",
)
replace_once(
    "docs/product/REQUIREMENTS-TRACEABILITY.md",
    """Current allocation:

```text
Product constitution: PROPOSED
Capability Specs: NOT_STARTED
Mission Contracts: NOT_STARTED
Implementation: PROHIBITED
Evidence: A0 documentation/research only
```

The large number does not mean every Capability implements all 294 requirements. R1 applicability selects relevant requirements and records rationale.
""",
    """A0 allocation baseline:

```text
Product constitution: ACCEPTED
Capability Specs: NOT_STARTED at A0 closeout
Mission Contracts: NOT_STARTED at A0 closeout
Implementation: NOT AUTHORIZED BY A0
Evidence: accepted A0 documentation/research + independent fresh-session Golden Proof
```

The large number does not mean every Capability implements all 294 requirements. R1 applicability selects relevant requirements and records rationale. Current milestone/gate authorization is intentionally not owned by this specification; consult `docs/tracking/STATUS.md`.
""",
)
replace_once(
    "docs/product/REQUIREMENTS-TRACEABILITY.md",
    """## 19. Next derivation step

When the next Product Milestone is selected, create the first Capability applicability/requirements package.

Example for M0:
""",
    """## 19. Next derivation step

`M0 — Sovereign Core Walking Skeleton` is the selected first Product Milestone. When ACRM R1 is separately authorized, create the first Capability applicability/requirements package.

Target package for M0:
""",
)

# F02 — align the ADR discovery/status index with accepted decision owners.
replace_once("docs/adr/README.md", "version: 0.2.0", "version: 0.2.1")
replace_once(
    "docs/adr/README.md",
    """| ADR | Title | Status | Primary decision |
|---|---|---|---|
| [ADR-0001](0001-aurora-owned-contract-model.md) | Aurora-owned Contract Model and Replaceable Bindings | proposed | Aurora owns cross-Harness semantics; protocols remain bindings |
| [ADR-0002](0002-first-party-harness-development-kit.md) | First-party Harness Development Kit and Universal Conformance | proposed | first-party Harnesses use AHDK by policy; conformance remains universal |

No ADR is accepted in A0 yet.
""",
    """| ADR | Title | Status | Primary decision |
|---|---|---|---|
| [ADR-0001](0001-aurora-owned-contract-model.md) | Aurora-owned Contract Model and Replaceable Bindings | accepted | Aurora owns cross-Harness semantics; protocols remain bindings |
| [ADR-0002](0002-first-party-harness-development-kit.md) | First-party Harness Development Kit and Universal Conformance | accepted | first-party Harnesses use AHDK by policy; conformance remains universal |

ADR-0001 and ADR-0002 were explicitly accepted by the operator as part of the A0 decision gate on 2026-08-06. Future ADRs remain non-governing until their own lifecycle and required authority promote them.
""",
)

# Documentation coverage: remove stale A0-closeout snapshot and align requirement status.
replace_once("docs/tracking/DOCUMENTATION-COVERAGE.md", "version: 2.0.0", "version: 2.0.1")
replace_once(
    "docs/tracking/DOCUMENTATION-COVERAGE.md",
    "| Requirements traceability | FULL | Requirements Traceability | 294 proposed constitutional requirements |",
    "| Requirements traceability | FULL | Requirements Traceability | 294 accepted A0 constitutional requirements |",
)
replace_once(
    "docs/tracking/DOCUMENTATION-COVERAGE.md",
    "At the current remediation state:",
    "At A0 acceptance:",
)
replace_once(
    "docs/tracking/DOCUMENTATION-COVERAGE.md",
    "- first Product Milestone contract after A0.",
    "- exact first Mission Contract for the selected M0 milestone.",
)
replace_once(
    "docs/tracking/DOCUMENTATION-COVERAGE.md",
    """Current state:

```text
CONTENT REMEDIATION COMPLETE
MECHANICAL VALIDATION IN PROGRESS
OPERATOR ACCEPTANCE NOT YET GRANTED
```
""",
    """A0 exit state:

```text
CONTENT REMEDIATION COMPLETE
MECHANICAL + ADVERSARIAL VALIDATION COMPLETE
INDEPENDENT FRESH-SESSION GOLDEN PROOF: PASS
OPERATOR ACCEPTANCE: GRANTED — 2026-08-06
A0 MERGE: COMPLETE
```

This coverage matrix does not own post-A0 milestone/gate coordination. Current readiness state is recorded only in `docs/tracking/STATUS.md`.
""",
)

# ---------------------------------------------------------------------------
# Tracking: record the failed R0 and the remediation proposal without advancing R1.
# ---------------------------------------------------------------------------
status = """---
id: DOC-AURORA-STATUS
title: Aurora Project Status
document_type: project_status
form: reference
authority: tracking
status: current
version: 0.9.0
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
  - DOC-AURORA-M0-OPERATOR-SELECTION
  - REVIEW-AURORA-M0-R0-CONSTITUTIONAL-BASELINE-2026-08-06
  - DOC-AURORA-M0-R0-REMEDIATION-AUTHORIZATION
last_reviewed: 2026-08-06
---

# Aurora Project Status

## 1. Current summary

- **Project:** Projeto Aurora
- **Canonical branch:** `main`
- **Canonical baseline reviewed by initial M0 R0:** `1da990f368a1bc693c09191c41d30a3db454d11e`
- **A0:** ACCEPTED and MERGED
- **A0 merge commit:** `f22085d97e198d99e89d52221b7b26d59d49bc12`
- **ADR-0001:** ACCEPTED
- **ADR-0002:** ACCEPTED
- **First Product Milestone:** `M0 — Sovereign Core Walking Skeleton` — SELECTED by operator
- **Current readiness gate:** ACRM R0 — Constitutional Baseline
- **Initial R0 verdict:** FAIL
- **R0 remediation:** AUTHORIZED / PROPOSED on non-canonical documentation branch / OPERATOR ACCEPTANCE OF CORRECTED REVISION PENDING
- **R1 and later gates:** NOT AUTHORIZED BY IMPLICATION
- **Stack decisions:** none
- **Runtime implementation:** not started and not authorized

## 2. M0 selection

The operator explicitly selected `M0 — Sovereign Core Walking Skeleton` as the first Product Milestone on 2026-08-06 after comparing M0 against M1 and M2.

Accepted M0 outcome:

> A minimal Aurora Core persists identity, project state, authority snapshot and one interaction lifecycle across restart without depending on an external model or Harness as authority.

Directional Golden Proof:

```text
initialize Aurora instance
→ create project
→ record accepted state and next action
→ terminate all Aurora processes
→ start fresh process
→ recover same identities/state
→ reject an invalid transition
→ export and restore state
```

Evidence: `docs/acceptance/2026-08-06-m0-operator-selection.md`.

## 3. Initial M0 R0 result

A fresh repository-only R0 review was executed against:

```text
1da990f368a1bc693c09191c41d30a3db454d11e
```

Verdict:

```text
R0 FAIL
```

Gate-failing findings:

- **R0-F01 — Product Milestone anatomy divergence:** Blueprint 14 required a complete executable-horizon milestone anatomy, while selected M0 lacked Architecture Spikes, Exit Criteria, Telemetry Baseline, Dependencies and Promotion/Authority Boundary; §14.5 also failed to distinguish selected executable milestones from intentionally directional future milestones.
- **R0-F02 — ADR status divergence:** `docs/adr/README.md`, despite owning ADR status discovery, still reported ADR-0001/0002 as proposed while their accepted ADR files and operator evidence reported ACCEPTED.
- **R0-F03 — mutable-state duplication/drift:** bootstrap/index/constitutional documents retained pre-A0/pre-M0 coordination snapshots even though `STATUS.md` and operator evidence had advanced.

Review record: `docs/reviews/2026-08-06-m0-r0-constitutional-baseline-review.md`.

## 4. Remediation boundary

The operator authorized remediation after the R0 FAIL. The authorized work is documentary/constitutional only:

- repair M0 roadmap anatomy without choosing technical mechanisms;
- make §14.5 consistent with the constitutional/executable two-horizon model;
- align ADR status discovery with accepted ADR owners/evidence;
- remove mutable-current-state ownership from durable constitutional/index documents and point it to `STATUS.md`;
- regenerate generated projections;
- run documentation validation;
- present the corrected revision for explicit operator acceptance;
- re-run M0 R0 only after the corrected constitutional revision is accepted/canonical.

This remediation does **not** authorize R1, Architecture Spike execution, Capability implementation, Aurora Core implementation, AHDK, MNFS integration, stack selection, Mission Contract or Microdesign.

## 5. Deliberately open M0-relevant technical decisions

The following remain open and must not be decided during R0 remediation:

- Aurora Core implementation language;
- initial process/deployment topology;
- operational state storage mechanism;
- state-versus-event model;
- schema representation;
- migration strategy details;
- backup/restore mechanism details;
- authority snapshot representation;
- audit/event mechanism;
- any durable execution engine beyond what M0 actually requires.

## 6. Current authorization boundary

```text
A0 baseline:                    ACCEPTED / MERGED
ADR-0001 / ADR-0002:           ACCEPTED
First Product Milestone:        M0 SELECTED
Initial M0 R0:                 FAIL
R0 documentary remediation:    AUTHORIZED
Corrected constitutional rev:  OPERATOR ACCEPTANCE PENDING
R0 re-run:                      PENDING corrected accepted revision
ACRM R1+:                       NOT AUTHORIZED
Architecture Spike execution:   PROHIBITED
Capability implementation:      PROHIBITED
Aurora Core implementation:     PROHIBITED
AHDK implementation:            PROHIBITED
MNFS integration:               PROHIBITED
Stack selection:                NOT PERFORMED
Mission Contract:               NOT STARTED
Microdesign/Implementation Plan: NOT STARTED
```

## 7. Current blocker/gate

The intentional blocker is the failed R0 constitutional baseline. The corrected documentation revision must be independently reviewable, validated and explicitly accepted before it can become the baseline for an R0 re-run.

No implementation blocker exists because implementation is not authorized work.

## 8. Immediate next action

```text
complete and validate R0 documentary remediation
→ operator reviews corrected constitutional revision
→ if accepted, integrate/promote it to canonical main
→ start a fresh repository-only session against the accepted revision
→ re-run M0 ACRM R0 only
→ produce R0 PASS | FAIL | BLOCKED
→ stop before R1 unless R1 is separately authorized
```
"""
Path("docs/tracking/STATUS.md").write_text(status, encoding="utf-8")

replace_once("docs/tracking/WORKLOG.md", "version: 0.6.0", "version: 0.7.0")
append_once(
    "docs/tracking/WORKLOG.md",
    "## 2026-08-06 — M0 R0 constitutional baseline review",
    """## 2026-08-06 — M0 R0 constitutional baseline review

A fresh repository-only review executed ACRM R0 for selected `M0 — Sovereign Core Walking Skeleton` against fixed commit:

```text
1da990f368a1bc693c09191c41d30a3db454d11e
```

Verdict:

```text
R0 FAIL
```

Findings:

- `R0-F01`: selected M0 did not satisfy the complete milestone anatomy required by Blueprint 14; the requirement also needed clarification that full executable anatomy applies when a milestone enters the executable horizon, preserving directional future milestones from premature detail;
- `R0-F02`: the ADR index reported ADR-0001/0002 as proposed despite accepted ADR owners and explicit operator acceptance;
- `R0-F03`: mutable current-state snapshots had drifted across AGENTS, README/index/constitutional/coverage documents after A0 merge and M0 selection.

The review deliberately did not choose language, storage, framework, runtime, protocol, topology or spike implementation and did not advance to R1.

Review record:

```text
docs/reviews/2026-08-06-m0-r0-constitutional-baseline-review.md
```

## 2026-08-06 — M0 R0 documentary remediation authorized

After receiving the R0 FAIL and exact allowed next action, Leandro instructed the project to continue. This is recorded as authorization to remediate the R0 documentation/constitutional findings only.

The remediation:

- completes selected M0's executable-horizon roadmap anatomy using only already-accepted intent;
- aligns the roadmap requirement with the two-horizon model;
- aligns ADR status discovery;
- removes duplicated mutable gate ownership from durable documents;
- regenerates Product Blueprint/Roadmap projections;
- validates documentation before operator review.

R1, Architecture Spike execution, stack selection, Aurora Core/AHDK/MNFS implementation, Mission Contract and Microdesign remain unauthorized.
""",
)

review = """---
id: REVIEW-AURORA-M0-R0-CONSTITUTIONAL-BASELINE-2026-08-06
title: M0 R0 Constitutional Baseline Review
document_type: adversarial_review
form: explanation
authority: evidence
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - observed M0 ACRM R0 findings against fixed revision 1da990f368a1bc693c09191c41d30a3db454d11e
related:
  - DOC-AURORA-STATUS
  - DOC-AURORA-M0-OPERATOR-SELECTION
  - DOC-AURORA-BLUEPRINT-14
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DOC-AURORA-M0-R0-REMEDIATION-AUTHORIZATION
last_reviewed: 2026-08-06
---

# M0 R0 Constitutional Baseline Review

## 1. Fixed scope

```text
Repository: developmentconexus-ops/aurora_project
Canonical branch reviewed: main
Target commit: 1da990f368a1bc693c09191c41d30a3db454d11e
Product Milestone: M0 — Sovereign Core Walking Skeleton
Gate: ACRM R0 — Constitutional Baseline
```

Question:

> Is the accepted constitutional intent required for M0 coherent, discoverable, sufficiently owned and authorized to proceed to applicability analysis?

No R1 applicability classification, technical selection, spike execution or implementation was performed.

## 2. Executive verdict

```text
R0 FAIL
```

M0's central intent is coherent and well aligned across Product Vision, Domain/World Model, Authority/Safety, Security/Sovereignty, System Architecture and Reliability/Evaluation. M0 is explicitly selected and R0 was authorized.

The gate failed because the accepted/documented baseline contained constitutional/documentation defects that R1 must not be forced to invent around.

## 3. R0-F01 — selected M0 milestone anatomy divergence

Blueprint 14 §14.5 requires the milestone in the executable horizon to define outcome, operator-visible value, named risk, entry criteria, capabilities, architecture spikes, Golden Proof, evidence, exit criteria, telemetry baseline, non-goals, dependencies, replan triggers and promotion/authority boundary.

Selected M0 defined the core outcome/value/risk/entry/capabilities/proof/evidence/non-goals/replan triggers but omitted:

- Architecture Spikes;
- Exit Criteria;
- Telemetry Baseline;
- Dependencies;
- Promotion/Authority Boundary.

The review also found an internal precision issue: §14.2 intentionally keeps distant milestones directional while §14.5 said every Product Milestone must carry the complete anatomy. The remediation must preserve both truths by requiring complete anatomy when a milestone is promoted into the current executable horizon.

Classification: constitutional/documentation owner defect. It cannot be delegated to R1/R2.

## 4. R0-F02 — ADR index status divergence

`docs/adr/README.md` owns ADR discovery/current status but reported ADR-0001 and ADR-0002 as `proposed` and stated that no ADR had been accepted in A0.

This contradicted:

- accepted frontmatter/content of ADR-0001;
- accepted frontmatter/content of ADR-0002;
- A0 operator acceptance evidence;
- STATUS and DECISIONS.

Classification: `DOCUMENTATION_DIVERGENCE` affecting fresh-session discoverability.

## 5. R0-F03 — mutable coordination state duplicated into durable sources

Post-A0/M0 commits correctly advanced `STATUS.md`, `DECISIONS.md`, WORKLOG and operator evidence but several durable guidance/constitutional/index documents retained pre-acceptance/pre-selection snapshots.

Observed examples included:

- AGENTS reporting first Product Milestone not yet selected;
- root README still saying the next gate was milestone selection;
- Documentation Map describing A0 work as current;
- Product index still reporting PR #1 merge merely authorized;
- Product Vision §1.19 reporting A0 as current phase;
- Blueprint 15 §15.31 reporting A0 `IN_REVIEW`;
- Requirements Traceability reporting the constitution as `PROPOSED`;
- Documentation Coverage reporting operator acceptance not yet granted.

Root cause: mutable coordination state was duplicated outside its canonical tracking owner and post-A0 promotion updated only a subset of those copies. The durable fix is to keep `STATUS.md` as the unique current-state owner and make constitutional/index documents point to it instead of owning competing snapshots.

## 6. Correctly deferred work

The following were deliberately not treated as R0 defects:

- M0 applicability classification → R1;
- atomic/verifiable Capability requirements → R2;
- M0 Capability/System Spec, lifecycle, threat model and test plan → R3;
- language, runtime, storage, topology, state/event model, schema and concrete backup/migration mechanisms → R4;
- Mission Contract → R5;
- Microdesign/Implementation Plan → R6;
- implementation/evidence → R7.

ADR-0001 is relevant only as a guardrail against framework/protocol-owned Aurora semantics. ADR-0002 does not pull AHDK into M0.

## 7. Remediation acceptance boundary

This review is evidence of a failed R0, not authority to self-accept its repair.

Required sequence:

```text
R0 FAIL evidence
→ documentary/constitutional remediation proposal
→ generated projections + validation
→ operator review/acceptance of corrected revision
→ canonical integration
→ fresh-session M0 R0 re-run
→ stop before R1 unless separately authorized
```
"""
Path("docs/reviews/2026-08-06-m0-r0-constitutional-baseline-review.md").write_text(review, encoding="utf-8")

# Regression assertions: stale/conflicting markers that caused this R0 failure must be gone.
checks = {
    "AGENTS.md": ["First Product Milestone: NOT YET SELECTED", "Product Milestone selection and ACRM R0 preparation are authorized next"],
    "README.md": ["Next product gate: select first Product Milestone", "O próximo passo é selecionar o primeiro Product Milestone"],
    "docs/product/blueprint/01-product-vision.md": ["## 1.19 Current product phase"],
    "docs/product/blueprint/15-documentation-research-governance.md": ["## 15.31 Current A0 gate", "A0 remains `IN_REVIEW`"],
    "docs/adr/README.md": ["No ADR is accepted in A0 yet."],
    "docs/product/REQUIREMENTS-TRACEABILITY.md": ["Product constitution: PROPOSED", "Current proposed constitutional requirements:"],
    "docs/tracking/DOCUMENTATION-COVERAGE.md": ["OPERATOR ACCEPTANCE NOT YET GRANTED"],
}
for path, stale in checks.items():
    text = Path(path).read_text(encoding="utf-8")
    for marker in stale:
        if marker in text:
            raise SystemExit(f"stale R0 marker remains in {path}: {marker}")

roadmap = Path("docs/product/blueprint/14-capability-roadmap.md").read_text(encoding="utf-8")
m0 = roadmap.split("# 14.8 M0 — Sovereign Core Walking Skeleton", 1)[1].split("# 14.9 M1", 1)[0]
for heading in [
    "## Outcome",
    "## Operator-visible value",
    "## Risk retired",
    "## Entry criteria",
    "## Capabilities involved",
    "## Architecture spikes",
    "## Golden Proof",
    "## Evidence requirements",
    "## Exit criteria",
    "## Telemetry baseline",
    "## Non-goals",
    "## Dependencies",
    "## Replan triggers",
    "## Promotion/authority boundary",
]:
    if heading not in m0:
        raise SystemExit(f"M0 anatomy still missing: {heading}")

# Remove the one-time remediation machinery from the final proposed diff.
Path(".github/workflows/fix-m0-r0-remediation.yml").unlink()
Path("scripts/fix_m0_r0_remediation.py").unlink()
