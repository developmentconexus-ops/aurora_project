#!/usr/bin/env python3
from __future__ import annotations

import re
import shutil
import subprocess
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
RATIFIED = "f4007eca5406ff71a944f5f980cdf3007bfc7504"
RATIFICATION = "1de294cb1aa1223e1490aee96cce79fdbcf802f9"
TODAY = "2026-08-23"


def p(rel: str) -> Path:
    return ROOT / rel


def run(*args: str) -> str:
    return subprocess.check_output(args, cwd=ROOT, text=True).strip()


def git_show(ref: str, path: str) -> str:
    return subprocess.check_output(["git", "show", f"{ref}:{path}"], cwd=ROOT, text=True)


def write(rel: str, text: str) -> None:
    path = p(rel)
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(text.rstrip() + "\n", encoding="utf-8")


def move(src: str, dst: str) -> None:
    source, target = p(src), p(dst)
    if not source.exists():
        return
    target.parent.mkdir(parents=True, exist_ok=True)
    shutil.move(str(source), str(target))


def rm(rel: str) -> None:
    path = p(rel)
    if path.is_dir():
        shutil.rmtree(path)
    elif path.exists():
        path.unlink()


def set_frontmatter_scalar(text: str, key: str, value: str) -> str:
    if not text.startswith("---\n"):
        return text
    end = text.find("\n---\n", 4)
    if end < 0:
        return text
    fm = text[4:end]
    pat = re.compile(rf"(?m)^{re.escape(key)}:\s*.*$")
    line = f"{key}: {value}"
    if pat.search(fm):
        fm = pat.sub(line, fm, count=1)
    else:
        fm = fm.rstrip() + "\n" + line
    return "---\n" + fm + text[end:]


def add_related(text: str, doc_id: str) -> str:
    if f"  - {doc_id}" in text:
        return text
    m = re.search(r"(?m)^related:\s*\n(?P<body>(?:  - .*\n)*)", text)
    if m:
        body = m.group("body") + f"  - {doc_id}\n"
        return text[: m.start("body")] + body + text[m.end("body") :]
    return text


def promote_ratified(text: str, version: str) -> str:
    text = set_frontmatter_scalar(text, "status", "accepted")
    text = set_frontmatter_scalar(text, "version", version)
    text = set_frontmatter_scalar(text, "accepted_at", TODAY)
    text = set_frontmatter_scalar(text, "acceptance_evidence", "DOC-AURORA-MR-01-OPERATOR-RATIFICATION")
    text = set_frontmatter_scalar(text, "last_reviewed", TODAY)
    text = text.replace("proposed MR-01", "MR-01")
    text = text.replace("proposed Aurora", "Aurora")
    text = text.replace("> **PROPOSED / NOT YET RATIFIED.**", "> **ACCEPTED / OPERATOR-RATIFIED.**")
    text = text.replace("> **PROPOSED.**", "> **ACCEPTED / OPERATOR-RATIFIED.**")
    return text


def replace_section(text: str, start: str, end: str, replacement: str) -> str:
    s = text.find(start)
    if s < 0:
        raise RuntimeError(f"missing section start: {start}")
    e = text.find(end, s)
    if e < 0:
        raise RuntimeError(f"missing section end: {end}")
    return text[:s] + replacement.rstrip() + "\n\n---\n\n" + text[e:]


def archive_doc(text: str, note: str, version: str = "1.0.0") -> str:
    text = set_frontmatter_scalar(text, "authority", "evidence")
    text = set_frontmatter_scalar(text, "status", "archived")
    text = set_frontmatter_scalar(text, "version", version)
    text = set_frontmatter_scalar(text, "last_reviewed", TODAY)
    marker = text.find("\n---\n", 4)
    if marker >= 0:
        body_start = marker + 5
        text = text[:body_start] + f"\n> **ARCHIVED SNAPSHOT.** {note}\n\n" + text[body_start:].lstrip("\n")
    return text


def current_markdown_paths() -> list[Path]:
    roots = [p("README.md"), p("AGENTS.md"), p("CONTRIBUTING.md")]
    roots += list(p("docs/product").rglob("*.md")) if p("docs/product").exists() else []
    roots += list(p("docs/capabilities").rglob("*.md")) if p("docs/capabilities").exists() else []
    roots += list(p("docs/architecture").rglob("*.md")) if p("docs/architecture").exists() else []
    roots += list(p("docs/decisions").rglob("*.md")) if p("docs/decisions").exists() else []
    roots += list(p("docs/development").rglob("*.md")) if p("docs/development").exists() else []
    roots += list(p("docs/reference").rglob("*.md")) if p("docs/reference").exists() else []
    return [x for x in roots if x.exists()]


# ---------------------------------------------------------------------------
# RM-02/RM-03/RM-04: prepare target semantic homes without losing blobs/IDs.
# ---------------------------------------------------------------------------

# Evidence/reviews keep exact bytes and stable IDs; their directory taxonomy changes only.
move("docs/acceptance", "docs/evidence/acceptance")
move("docs/reviews", "docs/evidence/reviews")

# ADRs keep exact files/IDs under current decision taxonomy.
p("docs/decisions/adr").mkdir(parents=True, exist_ok=True)
if p("docs/adr").exists():
    for item in list(p("docs/adr").iterdir()):
        move(str(item.relative_to(ROOT)), f"docs/decisions/adr/{item.name}")
    rm("docs/adr")

# Tracking classes with live consumers are rehomed before old tracking disappears.
move("docs/tracking/DECISIONS.md", "docs/decisions/index.md")
move("docs/tracking/WORKLOG.md", "docs/evidence/project-worklog.md")
move("docs/tracking/DOCUMENTATION-COVERAGE.md", "docs/evidence/a0-documentation-coverage.md")
move("docs/tracking/BACKLOG.md", "docs/evidence/a0-backlog-snapshot.md")
move("docs/tracking/STATUS.md", "docs/phases/pre-mr01-status-snapshot.md")
rm("docs/tracking")

# Current cross-system architecture gets semantic paths; M0-specific material stays phase Evidence.
move("docs/design/AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY.md", "docs/architecture/module-runtime-topology.md")
move("docs/design/AURORA-SYSTEM-ARCHITECTURE-DECISION-LANDSCAPE.md", "docs/architecture/system-decision-landscape.md")
move("docs/design/SAR-A1-STAGE-A-AVAILABILITY-ACTIVATION.md", "docs/architecture/stage-a-availability-activation.md")
move("docs/design/ARCHITECTURE-SPIKES.md", "docs/reference/architecture-spikes.md")
move("docs/design/AURORA-TECHNICAL-ARCHITECTURE-BASELINE-MAP.md", "docs/phases/technical-architecture-baseline-map.md")
for name in [
    "M0-R4-DECISION-LANDSCAPE.md",
    "M0-R4-MASTRA-FIT-MATRIX.md",
    "M0-R6-SOVEREIGN-CORE-IMPLEMENTATION-PLAN.md",
    "M0-R6-SOVEREIGN-CORE-MICRODESIGN.md",
    "SPK-AURORA-M0-OWNER-TRUST-002.md",
    "SPK-AURORA-M0-SOVEREIGN-STORE-001.md",
]:
    move(f"docs/design/{name}", f"docs/phases/m0/{name.lower()}")
rm("docs/design")

# Two accepted superpowers specs still have live semantic consumers; plans/other chronology are Git history.
move("docs/superpowers/specs/2026-08-12-aurora-system-architecture-rebaseline-design.md", "docs/architecture/system-architecture.md")
move("docs/superpowers/specs/2026-08-05-aurora-capability-harness-architecture-design.md", "docs/architecture/capability-harness-boundary.md")
rm("docs/superpowers")

# ACRM keeps its stable ID but moves out of Product semantics.
move("docs/product/CAPABILITY-REALIZATION-METHOD.md", "docs/development/capability-realization.md")

# Ratified MR-01 durable owners are copied from the exact ratified semantic candidate.
write(
    "docs/decisions/methodology-repository-rebaseline.md",
    promote_ratified(git_show(RATIFIED, "docs/decisions/methodology-repository-rebaseline.md"), "1.0.0"),
)
write(
    "docs/development/planning-readiness.md",
    promote_ratified(git_show(RATIFIED, "docs/development/planning-readiness.md"), "1.0.0"),
)
write(
    "docs/development/engineering-rules.md",
    promote_ratified(git_show(RATIFIED, "docs/development/engineering-rules.md"), "1.0.0"),
)

# Durable ratification Evidence is normalized from the branch-only ratification record.
rat = git_show(RATIFICATION, "docs/work/current/operator-ratification.md")
rat = set_frontmatter_scalar(rat, "document_type", "operator_ratification_evidence")
rat = set_frontmatter_scalar(rat, "status", "accepted")
rat = set_frontmatter_scalar(rat, "version", "1.0.0")
rat = rat.replace("> **BRANCH-ONLY / TEMPORARY EVIDENCE.** This records the operator decision for the pre-migration candidate. It must be absorbed into durable owners and removed with `docs/work/**` before any final merge candidate.\n\n", "")
write("docs/evidence/mr-01-operator-ratification.md", rat)

write(
    "docs/evidence/mr-01-independent-review.md",
    """---
id: DOC-AURORA-MR-01-INDEPENDENT-REVIEW
title: MR-01 Independent Review Evidence
document_type: independent_review_evidence
form: reference
authority: evidence
status: accepted
version: 1.0.0
owners:
  - developmentconexus-ops
related:
  - DOC-AURORA-METHODOLOGY-REPOSITORY-REBASELINE
last_reviewed: 2026-08-23
---

# MR-01 Independent Review Evidence

Independent Fable review was performed against candidate `d3d453401a9b4244b14e6e833d0aeb9da2caea41` on isolated PR #7. The review branch differed from the candidate only by `docs/work/current/ai-dialog.md`.

```text
verdict: CONVERGED / PASS_WITH_FINDINGS
blocking: 0
material: 2
moderate: 3
minor: 2
TA-01 reopen: NOT REQUIRED
TA-02 reopen: NOT REQUIRED
second round: NOT REQUIRED for the accepted bounded corrections
```

Findings MR-I01..MR-I06 were corrected within existing ratified intent. MR-I07 was `DEFER_SAFELY`: any future Frontend Product Experience Planning Method adoption must pin the exact canonical identity/version/location at its real consumer. The review created no Product, stack or runtime authority.

Original review Evidence remains reachable in closed unmerged PR #7 / review commit `5ab89f8cc21c120afa439a89f35983cf072843b1`.
""",
)

write(
    "docs/evidence/mr-01-migration-authorization.md",
    """---
id: DOC-AURORA-MR-01-MIGRATION-AUTHORIZATION
title: MR-01 Repository Migration Authorization
document_type: operator_authorization_evidence
form: reference
authority: evidence
status: accepted
version: 1.0.0
owners:
  - developmentconexus-ops
related:
  - DOC-AURORA-MR-01-OPERATOR-RATIFICATION
last_reviewed: 2026-08-23
---

# MR-01 Repository Migration Authorization

On 2026-08-23 the operator explicitly authorized execution of the MR-01 repository migration after MR-01 semantic ratification.

```text
base main: 35614c581cea32e04305c1ad63522fee151eb283
migration branch: docs/mr-01-repository-migration-20260823
migration PR: #8
integration model: one migration branch → one final merge candidate
```

This authorization covers the repository/documentation/control-plane migration only. It does not authorize TA-03+, Product/runtime implementation, M0 R7/R8, Architecture Spike execution, stack selection or merge.
""",
)

# ---------------------------------------------------------------------------
# Refine the exact documents whose semantics changed.
# ---------------------------------------------------------------------------

# ACRM: preserve R0-R8 while moving global cross-system readiness to the ratified standard.
acrm_path = p("docs/development/capability-realization.md")
acrm = acrm_path.read_text(encoding="utf-8")
acrm = set_frontmatter_scalar(acrm, "version", "0.3.0")
acrm = set_frontmatter_scalar(acrm, "last_reviewed", TODAY)
acrm = add_related(acrm, "DOC-AURORA-PLANNING-READINESS")
acrm = add_related(acrm, "DOC-AURORA-ENGINEERING-RULES")
acrm = acrm.replace("  - program-level System Architecture Rebaseline integration\n", "  - capability/milestone/mission realization lifecycle\n")
acrm = replace_section(
    acrm,
    "## 6A. Program-level System Architecture Rebaseline",
    "## 7. Readiness gates R0–R8",
    """## 6A. Relationship to cross-system Planning and Implementation Readiness

Aurora is a system of systems. Cross-system architecture/readiness is governed by `docs/development/planning-readiness.md`; ACRM does not duplicate that program.

```text
Aurora Product / accepted architecture
        ↓
Planning & Implementation-Readiness authority
        ↓
selected Product Milestone / Capability / Mission
        ↓
R0 → R1 → R2 → R3 → R4 → R5 → R6
        ↓
R7 Execution and Evidence
        ↓
R8 Product Milestone Closeout
```

ACRM consumes the current cross-system owners, operations, contracts, data/security/cognitive boundaries and Paved-Road constraints that apply to its selected scope. A Capability does **not** replay every global TA stage.

If a capability exposes a missing cross-system material decision, it opens a Finding against the smallest global owner and blocks the affected commitment. It must not invent the missing architecture locally.

The `DECIDE | RESEARCH | SPIKE | DEFER` vocabulary remains valid inside R4 for capability-scoped uncertainty. Architecture Spike execution still requires separate explicit authorization.

Repository navigation, Git/review behavior and temporary-work rules are governed by the DevelopmentConexus Repository Standard plus `docs/development/engineering-rules.md`, not by ACRM.
""",
)
acrm = acrm.replace("### Roadmap/Product Milestone", "### Blueprint 14 / Product Milestone")
acrm_path.write_text(acrm, encoding="utf-8")

# Blueprint 07 current name.
bp07 = p("docs/product/blueprint/07-harness-orchestration.md")
text = bp07.read_text(encoding="utf-8")
text = text.replace("MNFS", "Conexus OS")
text = text.replace("### Conexus OS\n", "### Conexus OS (historically MNFS)\n", 1)
text = set_frontmatter_scalar(text, "version", "0.2.1")
text = set_frontmatter_scalar(text, "last_reviewed", TODAY)
bp07.write_text(text, encoding="utf-8")

# Blueprint 15 bounded operating-model amendment.
bp15 = p("docs/product/blueprint/15-documentation-research-governance.md")
text = bp15.read_text(encoding="utf-8")
text = set_frontmatter_scalar(text, "version", "0.3.0")
text = set_frontmatter_scalar(text, "accepted_at", TODAY)
text = set_frontmatter_scalar(text, "acceptance_evidence", "DOC-AURORA-MR-01-OPERATOR-RATIFICATION")
text = set_frontmatter_scalar(text, "last_reviewed", TODAY)
text = replace_section(
    text,
    "## 15.2 Governing principles",
    "## 15.3 Two classification axes",
    """## 15.2 Governing principles

### P1 — One durable concept, one canonical owner
Other documents may summarize, explain or apply; they do not silently redefine.

### P2 — Conversation is discovery, repository is canonical project memory
Durable intent is promoted to the correct repository owner.

### P3 — Research is evidence, not authority
Research can support a decision without becoming the decision.

### P4 — One mutable repository-program owner
`docs/roadmap.md` is the sole mutable repository-program authority for current stage/gate, implementation permission, blockers and exact next action. It does not own Product or architecture meaning.

### P5 — Historical information remains discoverable
Rejected, superseded and failed approaches remain reachable through current Evidence when consumed or through Git/closed PR history when live retention is unnecessary.

### P6 — Accepted normative content cannot hide placeholders
Open questions are explicit and owned by research/spike/decision, not vague TODOs.

### P7 — Generated projections derive authority from sources
They are never independent edit targets.

### P8 — Documentation depth is mechanism-driven
Length is not a goal; material boundaries, failures, proof and non-goals are.

### P9 — Fresh actors load the smallest correct authority set
The target route is `AGENTS.md → docs/index.md → docs/roadmap.md → 1–2 task-specific owners` and normal work fits five files or fewer unless a named material reason exists.

### P10 — Implementation is separately gated
Detailed architecture, migration completion or green CI does not authorize Product implementation by implication.
""",
)
text = text.replace("Documentation Map for authority/read paths.", "`docs/index.md` for task/intention routing and authority entrypoints.")
text = re.sub(
    r"(?ms)### A8 — Tracking\n.*?(?=### A9 — Research / Historical)",
    """### A8 — Program Coordination / Tracking

Owns current repository-program stage/status/next-action only through `docs/roadmap.md`. Issues/boards may coordinate work but do not become parallel Product, architecture or status authority.

""",
    text,
)
text = text.replace("| Current project coordination | STATUS/tracking |", "| Current repository program/gate/implementation permission/next action | `docs/roadmap.md` |")
text = text.replace("| Specific technical choice | ADR |", "| Specific technical choice | ADR under `docs/decisions/adr/` |")
text = text.replace("### 15.7 Document identity and metadata", "### 15.7 Document identity and metadata")
meta_insert = "\nMetadata is required when it materially improves authority, machine routing, traceability or generation. Repository routers/guidance MAY use path/index routing without a uniform metadata envelope when that is sufficient; existing rich metadata may remain when useful.\n"
idx = text.find("## 15.7 Document identity and metadata")
if idx >= 0 and meta_insert.strip() not in text:
    endline = text.find("\n", idx) + 1
    text = text[:endline] + meta_insert + text[endline:]
text = replace_section(
    text,
    "## 15.10 Documentation layout",
    "## 15.11 Conversation-to-canonical promotion",
    """## 15.10 Documentation layout

The live tree follows the DevelopmentConexus Repository Standard and creates only paths with real consumers:

```text
README.md
AGENTS.md
CONTRIBUTING.md when useful

docs/
├── index.md
├── roadmap.md
├── product/
├── architecture/
├── decisions/
├── phases/
├── development/
├── capabilities/    # Aurora-specific active ACRM consumer
├── reference/
├── research/
├── evidence/
├── diagrams/        # only when a real diagram consumer exists
└── work/            # branch-only temporary; forbidden in final candidate/main
```

Final merge candidates and `main` contain no `docs/work/**`, `docs/superpowers/**`, permanent session-handoff/dialogue/review-round trees, parallel mutable status/roadmap surfaces or active archive/old trees used as current authority. Do not create empty directories for aesthetics.
""",
)
text = replace_section(
    text,
    "## 15.18 Capability Realization Method",
    "## 15.19 Tracking documents",
    """## 15.18 Capability Realization Method

ACRM owns realization of a selected Product Milestone / Capability / Mission:

```text
applicability
→ requirements
→ capability readiness
→ decisions/research/spikes
→ scoped contract
→ implementation-design readiness
→ execution Evidence
→ Product Milestone closeout
```

Cross-system planning/readiness is owned separately by `docs/development/planning-readiness.md` and is consumed by ACRM as current authority. Repository navigation/Git/review/context rules are owned by the DevelopmentConexus Repository Standard plus `docs/development/engineering-rules.md`.

R0–R8 identities and already-recorded M0 Evidence remain valid; this is a scope refinement, not a rewrite of their historical outcome.
""",
)
text = replace_section(
    text,
    "## 15.19 Tracking documents",
    "## 15.20 Status and authorization vocabulary",
    """## 15.19 Repository-program coordination and history

Permanent `STATUS`, `WORKLOG`, `BACKLOG` and documentation-coverage dashboards are not required live surfaces after MR-01.

```text
docs/roadmap.md
→ sole mutable repository-program stage/status/implementation permission/next action

docs/decisions/index.md
→ current decision disposition and forward obligation discovery

Git + merged/closed PR history
→ chronological change record/provenance

docs/evidence/** / docs/phases/**
→ durable proof/closure/snapshots only when a current consumer exists

docs/work/**
→ temporary branch-only planning/review; never main authority
```

Old tracking content retires only after semantic census/rehome and reachability proof.
""",
)
text = replace_section(
    text,
    "## 15.20 Status and authorization vocabulary",
    "## 15.21 Supersession",
    """## 15.20 Status and authorization vocabulary

Keep authorization distinctions explicit:

```text
DISCOVERY
RESEARCH
DESIGN
SPIKE
PLAN
IMPLEMENTATION
EXTERNAL EFFECT
MERGE
```

After repository migration, current permission and exact next action are read only from `docs/roadmap.md`. Absence of a prohibition is not authorization.
""",
)
text = text.replace(
    "Do not rewrite rejected rationale out of history.",
    "Do not rewrite rejected rationale out of history. Git is archive only when required provenance remains reachable; unique unmerged provenance with a current consumer receives a durable ref before the last branch/reference is deleted.",
)
text = replace_section(
    text,
    "## 15.22 Generated projections",
    "## 15.23 Documentation impact",
    """## 15.22 Generated projections

Keep the Product aggregate:

```text
docs/product/PRODUCT-BLUEPRINT.md
← generated from Blueprint 01–15
```

`docs/roadmap.md` is **not** generated from Blueprint 14 after MR-01. It is a hand-maintained repository-program authority. Blueprint 14 remains the Product capability-roadmap authority; the two intentionally answer different questions.

Generated files declare provenance and are validated for freshness. No generated projection becomes an independent semantic owner.
""",
)
text = replace_section(
    text,
    "## 15.24 Documentation checks",
    "## 15.25 Adversarial documentation review",
    """## 15.24 Documentation and repository checks

The migrated repository verification covers at least:

```text
AGENTS + docs/index + docs/roadmap <= 20 KiB
docs/roadmap is the sole mutable program-status/next-action authority
README is landing-only
default routed task pack <= 5 files unless a named reason exists
durable current owners are reachable from docs/index or a routed child index
current router links resolve
no durable authority depends on docs/work
no docs/work in final candidate/main
no docs/superpowers in final candidate/main
no permanent handoff/dialogue/review-round tree
no duplicate mutable roadmap/status surface
current decision dispositions valid/discoverable
unique required unmerged provenance remains reachable
review branch isolation can be proved mechanically
blocked-implementation top-level/source allowlist is enforced
material guards have deterministic negative controls
at least one required aggregate CI gate protects main
```

Aurora-specific checks may continue validating Product Blueprint generation, requirement identities, research manifests and other real consumers. A green structural check proves only the properties it actually tests.
""",
)
text = replace_section(
    text,
    "## 15.26 Human read paths",
    "## 15.27 Agent read budgets and Context Packs",
    """## 15.26 Human and agent read paths

### Default fresh actor

```text
AGENTS.md
→ docs/index.md
→ docs/roadmap.md
→ 1–2 task-specific owners
```

### Product reader

```text
README.md
→ docs/product/README.md
→ relevant Blueprint owner(s)
→ Blueprint 14 when Product capability sequence is needed
```

### Architecture / decision work

```text
docs/index.md
→ docs/architecture/index.md or docs/decisions/index.md
→ exact current owner
```

### Capability work

```text
docs/roadmap.md
→ docs/development/capability-realization.md
→ exact Capability owner / relevant cross-system authority
```

### Research reviewer

```text
docs/index.md
→ exact question-specific research
→ source manifest
→ decision/spike/capability it informs
```

Do not load the full Product Blueprint, phase history, Evidence, Git history or raw research by default.
""",
)
text = text.replace("/docs/adr/", "/docs/decisions/adr/")
text = text.replace("/docs/standards/\n", "/docs/development/\n")
text = text.replace("/docs/golden-paths/\n", "/docs/development/\n")
text = text.replace("/docs/acceptance/", "/docs/evidence/")
text = replace_section(
    text,
    "## 15.31 A0 acceptance rule and post-A0 current-state ownership",
    "## 15.32 Non-goals",
    """## 15.31 A0 acceptance and post-A0 current-state ownership

A0 acceptance remains historical fact: the fifteen-section Product constitution, discovery coverage, research, ADR baseline, ACRM, traceability and fresh-session/adversarial proof were explicitly accepted and merged.

After MR-01 migration, mutable repository-program state is not owned by this constitutional section. `docs/roadmap.md` owns the current stage/gate, blockers, implementation permission and exact next action. `docs/index.md` owns task/intention routing. Product meaning remains with the Product Blueprint and accepted specific owners.

A0 acceptance, MR-01 ratification or repository migration never authorizes later TA stages, Architecture Spike execution or Product implementation by implication.
""",
)
bp15.write_text(text, encoding="utf-8")

# Archive old mutable tracking identities without keeping them current.
for rel, note in [
    ("docs/evidence/project-worklog.md", "Chronological project history through the pre-MR-01 repository model. Current status is owned only by docs/roadmap.md."),
    ("docs/evidence/a0-documentation-coverage.md", "A0 discovery/documentation coverage Evidence. It is not current repository routing authority."),
    ("docs/evidence/a0-backlog-snapshot.md", "A0-era non-commitment backlog snapshot. Current commitments and decisions live in Product/decision owners."),
    ("docs/phases/pre-mr01-status-snapshot.md", "Last STATUS snapshot before the MR-01 repository control-plane migration. Current status is owned only by docs/roadmap.md."),
]:
    path = p(rel)
    if path.exists():
        path.write_text(archive_doc(path.read_text(encoding="utf-8"), note), encoding="utf-8")

# The accepted old Technical Architecture map remains a phase snapshot; the stage ordering is superseded.
tamap = p("docs/phases/technical-architecture-baseline-map.md")
if tamap.exists():
    t = tamap.read_text(encoding="utf-8")
    t = set_frontmatter_scalar(t, "status", "superseded")
    t = set_frontmatter_scalar(t, "last_reviewed", TODAY)
    marker = t.find("\n---\n", 4)
    if marker >= 0:
        pos = marker + 5
        t = t[:pos] + "\n> **SUPERSEDED ORDERING SNAPSHOT.** TA-01/TA-02 results remain canonical; the old TA-03+ sequence is superseded by `docs/development/planning-readiness.md`.\n\n" + t[pos:].lstrip("\n")
    tamap.write_text(t, encoding="utf-8")

# Decision register: preserve existing disposition history while making MR-01 current.
dec = p("docs/decisions/index.md")
text = dec.read_text(encoding="utf-8")
text = set_frontmatter_scalar(text, "authority", "reference")
text = set_frontmatter_scalar(text, "version", "1.0.0")
text = set_frontmatter_scalar(text, "last_reviewed", TODAY)
text = text.replace("MNFS is a future software-engineering provider and does not define Aurora architecture", "Conexus OS (historically MNFS) is the software-engineering Harness/provider candidate and does not define Aurora architecture")
text = text.replace("aggregate Product Blueprint and roadmap are generated from modular canonical sources", "the Product Blueprint aggregate is generated from modular canonical sources; repository `docs/roadmap.md` is hand-maintained after MR-01")
text = re.sub(r"\| D-067 \|.*?\| accepted \|", "| D-067 | the pre-MR-01 TA-03..TA-08 ordering is superseded/refined by MR-01; TA-01/TA-02 remain preserved | MR-01 + Planning Readiness | superseded/refined |", text)
anchor = "| D-072 | A05 owns Aurora-side runtime lifecycle policy and B01 owns transport-neutral provider identity/lifecycle/idempotency/cancellation/reconciliation semantics before TA-04 selects a binding | TA-01/TA-02 design v0.5.0 + operator acceptance | accepted |"
addition = """
| D-073 | MR-01 adopts DevelopmentConexus Engineering Method v1.0.0 and Repository Standard v1.0.0 by reference while preserving Aurora Product authority | MR-01 operator ratification | accepted |
| D-074 | the default fresh-actor route is `AGENTS → docs/index → docs/roadmap → 1–2 owners`; `docs/roadmap.md` is the sole mutable repository-program status/permission/next-action owner | MR-01 + Repository Standard | accepted |
| D-075 | remaining global planning follows TA-03 Cross-System Operations → TA-04 Executable Contracts → TA-05 Human/Presence → TA-06 Data → TA-07 Identity/Security/Effects → TA-08 Cognitive Runtime/Harnesses → TA-09 Paved Road/Source/Build → TA-10 Operations → TA-11 Golden Flows → TA-12 Execution Graph → TA-13 Readiness; TA-TX is conditional | Planning Readiness | accepted |
| D-076 | MR-01 repository migration uses one migration branch and one final merge candidate; intermediate RM states do not partially merge | MR-01 migration plan + operator authorization | accepted |
"""
if anchor in text and "| D-073 |" not in text:
    text = text.replace(anchor, anchor + addition)
# Update the highest-value open-decision stage pointers.
repls = {
    "TA-06 Effect/Authority Spec + spike + ADR": "TA-07 Effect/Authority Spec + spike + ADR",
    "TA-06 security Spec + actor-specific spike": "TA-07 security Spec + actor-specific spike",
    "TA-05/TA-07 evidence capability research/spike": "TA-06 evidence capability research/spike",
    "TA-04/TA-08 research/spike": "TA-04/TA-10 research/spike",
    "TA-05/TA-07 + CAP-MEMORY-CONTEXT eval spikes": "TA-06/TA-08 + CAP-MEMORY-CONTEXT eval spikes",
    "TA-07/M2 after current architecture review": "TA-08/M2 after current architecture review",
    "TA-06 research + capability-specific ADRs": "TA-07 research + capability-specific ADRs",
    "TA-08 Standard/ADR": "TA-10 Standard/ADR",
    "TA-07 + M1 research/Capability Spec": "TA-08 + M1 research/Capability Spec",
    "TA-05 consumer-specific research/spike": "TA-06 consumer-specific research/spike",
    "TA-03 after TA-01/TA-02 acceptance": "TA-09 after upstream consumers/contracts are closed",
    "TA-02/TA-04/TA-07, first consumer evidence": "TA-02/TA-04/TA-08/TA-09, first consumer evidence",
}
for old, new in repls.items():
    text = text.replace(old, new)
dec.write_text(text, encoding="utf-8")

# ---------------------------------------------------------------------------
# RM-05 atomic control-plane cutover.
# ---------------------------------------------------------------------------

README = """# Projeto Aurora

Aurora é uma **inteligência artificial pessoal, persistente, multimodal e agêntica**: o control plane cognitivo e operacional de Leandro para projetos, conhecimento, capacidades, ferramentas, dispositivos e ambientes.

> **North Star:** Leandro pode retomar qualquer projeto com objetivo, estado, decisões e Evidence preservados; Aurora reúne contexto, coordena capacidades digitais/físicas dentro de authority explícita e ajuda a decidir o próximo passo.

Aurora é Leandro-first, local-first/cloud-assisted, framework-neutral, Evidence-driven e governada por memória/authority explícitas. Harnesses especializadas — incluindo Conexus OS quando delegado — executam trabalho local; elas não possuem a identidade, estado soberano ou autoridade global da Aurora.

## Entrar no repositório

- **Agentes:** leia [`AGENTS.md`](AGENTS.md).
- **Humanos / documentação:** comece em [`docs/index.md`](docs/index.md).
- **Estado atual / trabalho permitido:** [`docs/roadmap.md`](docs/roadmap.md).

O Product completo vive em `docs/product/`; arquitetura em `docs/architecture/`; decisões em `docs/decisions/`; regras locais em `docs/development/`.

> Conversa é descoberta. O repositório é a memória canônica do projeto.
"""
write("README.md", README)

AGENTS = """# Projeto Aurora — Agent Bootstrap

Cross-repository authorities:

- `developmentconexus-ops/conexus-methodology/METHOD.md` — DevelopmentConexus Engineering Method v1.0.0;
- `developmentconexus-ops/conexus-methodology/REPOSITORY-STANDARD.md` — Repository Standard v1.0.0.

Aurora-specific rules are in `docs/development/engineering-rules.md`.

## Fresh-actor route

Always start with:

```text
AGENTS.md
→ docs/index.md
→ docs/roadmap.md
→ 1–2 task-specific owners named by docs/index.md
```

Normal work fits five files or fewer. Do not recursively read Product Blueprint, phase history, Evidence, research, Git history, frozen M0 implementation or qualification harnesses before a concrete task requires them.

## Authority

- Product Blueprint owns constitutional Product meaning.
- `docs/architecture/**` owns accepted structural architecture in its stated scope.
- accepted ADRs/decisions own specific choices.
- Capability Specs own reusable capability behavior.
- approved Contracts own exact scoped commitments.
- `docs/development/**` owns local methods/engineering rules.
- Evidence proves observations; it does not create Product authority.
- research informs decisions; it does not decide.
- `docs/roadmap.md` is the **only** mutable repository-program stage/status/implementation-permission/next-action authority.

Conflict is `DOCUMENTATION_DIVERGENCE`: stop affected work, identify the smallest owner and replan there.

## Hard stops

Do not, by convenience:

- change Product meaning or a canonical owner;
- create/duplicate a trust, authority, effect or credential boundary;
- invent cross-system operations/contracts/persistent data during implementation;
- globalize M0-scoped Go/SQLite/other choices;
- make a model/framework/provider/Conexus OS own sovereign Aurora state/authority;
- execute external/physical effects outside exact authority;
- weaken guards to make CI green;
- start TA-03+, M0 R7/R8, Architecture Spikes or Product implementation unless `docs/roadmap.md` explicitly authorizes that exact work.

## Capability realization

For a selected Product Milestone / Capability / Mission read `docs/development/capability-realization.md`. R0→R8 remains the capability/slice realization lifecycle. Cross-system prerequisites come from `docs/development/planning-readiness.md` and are consumed rather than re-invented locally.

## Verification / Git

- one coherent gate per branch/PR by default;
- no direct commits/force-push to `main`;
- independent review output is Evidence only;
- temporary work lives only in branch `docs/work/current/` and never enters a final candidate/main;
- no completion claim without fresh verification on the exact target revision.

Current local verification is `python scripts/validate_docs.py --self-test` plus the `Documentation` GitHub workflow. Exact current work is always read from `docs/roadmap.md`.
"""
write("AGENTS.md", AGENTS)

CONTRIB = """# Contributing to Projeto Aurora

Start with `AGENTS.md`, then `docs/index.md` and `docs/roadmap.md`. Work only inside the authorization recorded by the roadmap and the exact task-specific owner.

Material changes follow the DevelopmentConexus Engineering Method and Repository Standard. Use one branch/Draft PR per coherent gate by default, preserve authority/provenance, run the repository verification, and never merge without explicit merge authority when required.

Product/runtime implementation remains a separate authorization from architecture, migration, CI or document acceptance.
"""
write("CONTRIBUTING.md", CONTRIB)

DOC_INDEX = """---
id: DOC-AURORA-DOCUMENTATION-MAP
title: Aurora Documentation Index
document_type: documentation_index
form: reference
authority: standard
status: accepted
accepted_at: 2026-08-23
acceptance_evidence: DOC-AURORA-MR-01-OPERATOR-RATIFICATION
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - task and intention routing
  - documentation authority entrypoints
  - smallest default read packs
related:
  - DOC-AURORA-BLUEPRINT-15
  - DOC-AURORA-REPOSITORY-ROADMAP
  - DOC-AURORA-DECISIONS
  - DOC-AURORA-PLANNING-READINESS
last_reviewed: 2026-08-23
---

# Aurora Documentation Index

This file routes a task to the smallest correct authority set. It does not own mutable program status or Product/architecture semantics.

| Intent | Start here | Do not read by default |
|---|---|---|
| Current gate, blockers, allowed work, exact next action | [`roadmap.md`](roadmap.md) | Product history, Evidence, Git history |
| Product / North Star / constitutional meaning | [`product/README.md`](product/README.md) → exact Blueprint owner | full Blueprint unless cross-domain impact is material |
| Logical modules, owners, Stage-A/B topology | [`architecture/index.md`](architecture/index.md) | M0 microdesign/history |
| Current decisions / open decisions / reopen triggers | [`decisions/index.md`](decisions/index.md) | review chronology |
| Cross-system planning/readiness | [`development/planning-readiness.md`](development/planning-readiness.md) | implementation plans |
| Capability/milestone realization R0–R8 | [`development/capability-realization.md`](development/capability-realization.md) | unrelated global TA stages |
| Repository/Git/review/verification rules | [`development/engineering-rules.md`](development/engineering-rules.md) | organizational Method copies |
| Capability-specific work | exact owner under [`capabilities/`](capabilities/) | other capabilities |
| Research | exact question/report via [`research/RESEARCH-MAP.md`](research/RESEARCH-MAP.md) | broad research corpus |
| Architecture-spike portfolio | [`reference/architecture-spikes.md`](reference/architecture-spikes.md) | spike execution unless authorized |
| M0 frozen design/evidence | [`phases/m0/`](phases/m0/) + exact Evidence | current source architecture inference |
| Acceptance/review proof | exact item under [`evidence/`](evidence/) when a current claim requires it | evidence chronology generally |
| Aurora origin/discovery | [`history/2026-08-05-aurora-origin-and-discovery-record.md`](history/2026-08-05-aurora-origin-and-discovery-record.md) | default task context |

## Authority law

```text
Product Blueprint / accepted structural owner / accepted decision
> Specification / Contract / Standard in its scope
> current machinery Reference
> Evidence for its observation
> research/history
```

`docs/roadmap.md` owns only repository-program progression. `docs/index.md` owns only routing. Neither may silently redefine Product/architecture.
"""
write("docs/index.md", DOC_INDEX)

ROADMAP = """---
id: DOC-AURORA-REPOSITORY-ROADMAP
title: Aurora Repository Roadmap
document_type: repository_program_roadmap
form: reference
authority: tracking
status: current
program_status_authority: true
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - current repository stage and gate
  - implementation permission and blocked work
  - exact next action
  - progression and reopen triggers
related:
  - DOC-AURORA-MR-01-OPERATOR-RATIFICATION
  - DOC-AURORA-MR-01-MIGRATION-AUTHORIZATION
  - DOC-AURORA-PLANNING-READINESS
last_reviewed: 2026-08-23
---

# Aurora Repository Roadmap

> This is the sole mutable repository-program status/permission/next-action authority. Product capability sequence remains in Blueprint 14.

## Current canonical baseline

```text
A0 Product constitution: ACCEPTED / MERGED
System Architecture Rebaseline: ACCEPTED / MERGED
TA-01 Logical Modules & Canonical Ownership: ACCEPTED / CANONICAL
TA-02 Process/Runtime/Evolutionary Topology: ACCEPTED / CANONICAL
M0 R0–R6: historical PASS within M0 scope
M0 R7 candidate: FROZEN / PRESERVED / NON-CANONICAL
M0 R7 Verdict: NOT ISSUED
M0 R8: NOT AUTHORIZED
MR-01 Methodology/Repository/Readiness target: OPERATOR-RATIFIED
```

## Current program

```text
program: MR-01 Repository Migration
migration PR: #8
RM-01: PASS — semantic/provenance census
RM-02: APPLIED — bounded constitutional/method reconciliation
RM-03: APPLIED — target authority preparation
RM-04: APPLIED — decision/architecture/capability reconciliation
RM-05: APPLIED — atomic repository control-plane cutover
RM-06: APPLIED — legacy live-surface retirement/rehome
RM-07: NEXT — verification + deterministic negative controls
RM-08: PENDING — Git/branch-protection enforcement
RM-09: PENDING — fresh-actor + Global Coherence proof
RM-10: PENDING — independent final review / promotion candidate
```

## Authorization boundary

```text
repository migration execution: AUTHORIZED
TA-03+: NOT AUTHORIZED
Architecture Spike execution: NOT AUTHORIZED
Aurora Product/runtime implementation: BLOCKED
M0 R7 continuation/Verdict/R8: NOT AUTHORIZED
framework/database/IAM/model/provider selection: NOT AUTHORIZED BY MIGRATION
merge: NOT AUTHORIZED
```

## Exact next action

```text
run RM-07 positive validation + guard negative controls on the migrated tree
→ correct only migration defects
→ RM-08 platform enforcement
→ RM-09 fresh-actor/global coherence proof
→ prepare isolated independent RM-10 review
→ STOP before merge authorization
```

## Reopen triggers

Reopen the MR-01 target only if migration Evidence shows lost current semantics/provenance, duplicate/missing authority, an unworkable fresh-actor route, or a downstream material decision forced before its owner exists.
"""
write("docs/roadmap.md", ROADMAP)

ARCH_INDEX = """---
id: DOC-AURORA-ARCHITECTURE-INDEX
title: Aurora Architecture Index
document_type: architecture_index
form: reference
authority: reference
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
related:
  - DESIGN-AURORA-SYSTEM-ARCHITECTURE-REBASELINE
  - DESIGN-AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY
  - DOC-AURORA-PLANNING-READINESS
last_reviewed: 2026-08-23
---

# Aurora Architecture Index

| Question | Current owner |
|---|---|
| Product-level logical architecture/invariants | `docs/product/blueprint/12-system-architecture.md` |
| System-architecture rebaseline/boundaries | `system-architecture.md` |
| Canonical modules, data ownership, dependency direction, Stage-A/B runtime topology | `module-runtime-topology.md` |
| Stage-A availability / activation / locked-workstation constraints | `stage-a-availability-activation.md` |
| Open cross-system architecture decision landscape | `system-decision-landscape.md`, interpreted through current Planning Readiness |
| Capability/Harness responsibility boundary | `capability-harness-boundary.md` + Blueprint 07 |
| Cross-system remaining stage order | `../development/planning-readiness.md` |
| M0-specific old microdesign/spikes | `../phases/m0/` — Evidence/history, not target architecture by existence |
| Architecture spike portfolio | `../reference/architecture-spikes.md` |

Old Technical Architecture Baseline ordering is preserved as a superseded snapshot under `../phases/technical-architecture-baseline-map.md`; TA-01/TA-02 results remain current.
"""
write("docs/architecture/index.md", ARCH_INDEX)

# Product index becomes Product-only rather than repository router.
prod = p("docs/product/README.md")
if prod.exists():
    t = prod.read_text(encoding="utf-8")
    t = t.replace("docs/product/CAPABILITY-REALIZATION-METHOD.md", "docs/development/capability-realization.md")
    t = t.replace("CAPABILITY-REALIZATION-METHOD.md", "../development/capability-realization.md")
    t = t.replace("`docs/roadmap.md`", "Blueprint 14 (`blueprint/14-capability-roadmap.md`)")
    prod.write_text(t, encoding="utf-8")

# Current-document path/ID updates. Evidence/phases/history preserve historical wording by default.
path_replacements = {
    "docs/acceptance/": "docs/evidence/acceptance/",
    "../acceptance/": "../evidence/acceptance/",
    "../../acceptance/": "../../evidence/acceptance/",
    "docs/reviews/": "docs/evidence/reviews/",
    "../reviews/": "../evidence/reviews/",
    "../../reviews/": "../../evidence/reviews/",
    "docs/adr/": "docs/decisions/adr/",
    "../adr/": "../decisions/adr/",
    "../../adr/": "../../decisions/adr/",
    "docs/design/AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY.md": "docs/architecture/module-runtime-topology.md",
    "docs/design/SAR-A1-STAGE-A-AVAILABILITY-ACTIVATION.md": "docs/architecture/stage-a-availability-activation.md",
    "docs/design/AURORA-SYSTEM-ARCHITECTURE-DECISION-LANDSCAPE.md": "docs/architecture/system-decision-landscape.md",
    "docs/design/ARCHITECTURE-SPIKES.md": "docs/reference/architecture-spikes.md",
    "docs/design/AURORA-TECHNICAL-ARCHITECTURE-BASELINE-MAP.md": "docs/phases/technical-architecture-baseline-map.md",
    "docs/product/CAPABILITY-REALIZATION-METHOD.md": "docs/development/capability-realization.md",
    "docs/DOCUMENTATION-MAP.md": "docs/index.md",
    "docs/tracking/STATUS.md": "docs/roadmap.md",
    "docs/tracking/DECISIONS.md": "docs/decisions/index.md",
    "docs/tracking/DOCUMENTATION-COVERAGE.md": "docs/evidence/a0-documentation-coverage.md",
    "docs/tracking/WORKLOG.md": "docs/evidence/project-worklog.md",
}
for path in current_markdown_paths():
    text = path.read_text(encoding="utf-8")
    for old, new in path_replacements.items():
        text = text.replace(old, new)
    text = text.replace("DOC-AURORA-ROADMAP", "DOC-AURORA-BLUEPRINT-14")
    path.write_text(text, encoding="utf-8")

# Old Documentation Map is replaced by docs/index with the same stable document ID.
rm("docs/DOCUMENTATION-MAP.md")

# ---------------------------------------------------------------------------
# Generator / validator / CI switch on the same commit.
# ---------------------------------------------------------------------------

GENERATOR = r'''#!/usr/bin/env python3
from __future__ import annotations

import argparse
import hashlib
import re
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
BLUEPRINT_DIR = Path("docs/product/blueprint")
BLUEPRINT_OUTPUT = Path("docs/product/PRODUCT-BLUEPRINT.md")


def split_frontmatter(text: str, path: Path) -> tuple[str, str]:
    if not text.startswith("---\n"):
        raise ValueError(f"{path}: missing YAML frontmatter")
    marker = text.find("\n---\n", 4)
    if marker < 0:
        raise ValueError(f"{path}: unterminated YAML frontmatter")
    return text[4:marker], text[marker + 5:].lstrip("\n")


def scalar(frontmatter: str, key: str) -> str:
    m = re.search(rf"(?m)^{re.escape(key)}:\s*([^\n]+)\s*$", frontmatter)
    if not m:
        raise ValueError(f"frontmatter missing {key}")
    return m.group(1).strip().strip('"').strip("'")


def canonical_sources(root: Path) -> list[Path]:
    sources = sorted((root / BLUEPRINT_DIR).glob("[0-9][0-9]-*.md"))
    actual = [int(x.name[:2]) for x in sources]
    if actual != list(range(1, 16)):
        raise ValueError(f"expected Blueprint 01..15, found {actual}")
    return sources


def generate_blueprint(root: Path) -> str:
    sections = []
    for path in canonical_sources(root):
        raw = path.read_text(encoding="utf-8")
        fm, body = split_frontmatter(raw, path)
        sections.append((path, scalar(fm, "id"), hashlib.sha256(raw.encode()).hexdigest(), body.rstrip()))
    ids = [x[1] for x in sections]
    out = [
        "---",
        "id: DOC-AURORA-PRODUCT-BLUEPRINT",
        "title: Aurora Product Blueprint",
        "document_type: product_blueprint_aggregate",
        "form: explanation",
        "authority: generated_projection",
        "status: generated",
        "version: 0.3.0",
        "owners:",
        "  - developmentconexus-ops",
        "generated_from:",
        *[f"  - {x}" for x in ids],
        "---",
        "",
        "<!-- GENERATED — DO NOT EDIT DIRECTLY",
        "Canonical sources: docs/product/blueprint/01-*.md through 15-*.md",
        "Generator: scripts/generate_docs.py",
        "-->",
        "",
        "# Aurora Product Blueprint",
        "",
        "> Generated publication of the fifteen modular constitutional sources. Edit sources, regenerate, validate.",
        "",
        "## Source manifest",
        "",
        "| Section | Canonical source | SHA-256 |",
        "|---:|---|---|",
    ]
    for path, _, digest, _ in sections:
        out.append(f"| {path.name[:2]} | `{path.as_posix()}` | `{digest}` |")
    out += ["", "---", ""]
    for i, (path, _, _, body) in enumerate(sections):
        out += [f"<!-- BEGIN SOURCE: {path.as_posix()} -->", body, f"<!-- END SOURCE: {path.as_posix()} -->"]
        if i != len(sections)-1:
            out += ["", "---", ""]
    return "\n".join(out).rstrip() + "\n"


def main() -> int:
    ap = argparse.ArgumentParser()
    ap.add_argument("--output-root", type=Path, default=ROOT)
    ap.add_argument("--check", action="store_true")
    args = ap.parse_args()
    content = generate_blueprint(ROOT)
    if args.check:
        target = ROOT / BLUEPRINT_OUTPUT
        if not target.exists() or target.read_text(encoding="utf-8") != content:
            print(f"STALE GENERATED DOCUMENT: {BLUEPRINT_OUTPUT}")
            return 1
        print("Generated Product Blueprint is current; repository roadmap is hand-maintained.")
        return 0
    target = args.output_root.resolve() / BLUEPRINT_OUTPUT
    target.parent.mkdir(parents=True, exist_ok=True)
    target.write_text(content, encoding="utf-8")
    print(target)
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
'''
write("scripts/generate_docs.py", GENERATOR)

VALIDATOR = r'''#!/usr/bin/env python3
from __future__ import annotations

import argparse
import json
import re
import shutil
import subprocess
import sys
import tempfile
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
LINK_RE = re.compile(r"(?<!!)\[[^\]]*\]\(([^)]+)\)")
REQ_RE = re.compile(r"\bAUR-REQ-[A-Z0-9-]+-\d{3}\b")
SOURCE_RE = re.compile(r"\[(S\d{2,})\]")
PROGRAM_OWNER_RE = re.compile(r"(?m)^program_status_authority:\s*true\s*$")
ALLOWED_ROOT = {".github", "docs", "scripts", "README.md", "AGENTS.md", "CONTRIBUTING.md"}
REQUIRED = [
    "README.md", "AGENTS.md", "docs/index.md", "docs/roadmap.md",
    "docs/product/README.md", "docs/product/PRODUCT-BLUEPRINT.md",
    "docs/product/blueprint/15-documentation-research-governance.md",
    "docs/development/engineering-rules.md", "docs/development/planning-readiness.md",
    "docs/development/capability-realization.md", "docs/decisions/index.md",
    "docs/architecture/index.md",
]
FINAL_FORBIDDEN = ["docs/work", "docs/superpowers", "docs/tracking", "docs/acceptance", "docs/reviews", "docs/adr", "docs/design", "docs/DOCUMENTATION-MAP.md"]


def md_files(root: Path):
    return sorted(x for x in root.rglob("*.md") if ".git" not in x.parts)


def frontmatter(text: str) -> str | None:
    if not text.startswith("---\n"):
        return None
    end = text.find("\n---\n", 4)
    return None if end < 0 else text[4:end]


def scalar(fm: str, key: str) -> str | None:
    m = re.search(rf"(?m)^{re.escape(key)}:\s*([^\n]+)\s*$", fm)
    return None if not m else m.group(1).strip().strip('"').strip("'")


def list_field(fm: str, key: str) -> list[str]:
    m = re.search(rf"(?m)^{re.escape(key)}:\s*\n(?P<body>(?:  - .*\n)*)", fm)
    if not m:
        return []
    return [line[4:].strip() for line in m.group("body").splitlines() if line.startswith("  - ")]


def guard_bootstrap(root: Path, errors: list[str]):
    total = sum((root / x).stat().st_size for x in ["AGENTS.md", "docs/index.md", "docs/roadmap.md"] if (root / x).exists())
    if total > 20 * 1024:
        errors.append(f"bootstrap budget exceeded: {total} bytes > 20480")


def guard_program_owner(root: Path, errors: list[str]):
    owners = []
    for path in md_files(root):
        if PROGRAM_OWNER_RE.search(path.read_text(encoding="utf-8")):
            owners.append(path.relative_to(root).as_posix())
    if owners != ["docs/roadmap.md"]:
        errors.append(f"program status authority must be only docs/roadmap.md; found {owners}")


def guard_root_allowlist(root: Path, errors: list[str]):
    roadmap = root / "docs/roadmap.md"
    if not roadmap.exists() or "Product/runtime implementation: BLOCKED" not in roadmap.read_text(encoding="utf-8"):
        return
    unexpected = sorted(x.name for x in root.iterdir() if x.name not in ALLOWED_ROOT and x.name != ".git")
    if unexpected:
        errors.append("implementation-blocked root surface not allowlisted: " + ", ".join(unexpected))


def guard_final_tree(root: Path, errors: list[str]):
    for rel in FINAL_FORBIDDEN:
        if (root / rel).exists():
            errors.append(f"final candidate forbidden legacy/temporary surface exists: {rel}")


def guard_review_changed_files(files: list[str], errors: list[str]):
    if files != ["docs/work/current/ai-dialog.md"]:
        errors.append(f"review isolation violation: {files}")


def validate_links(root: Path, errors: list[str]):
    # Current authority/router surfaces only; Evidence/phase/history snapshots may preserve historical paths.
    current_roots = ["README.md", "AGENTS.md", "CONTRIBUTING.md", "docs/product", "docs/architecture", "docs/decisions", "docs/development", "docs/capabilities", "docs/reference"]
    candidates = []
    for rel in current_roots:
        path = root / rel
        if path.is_file(): candidates.append(path)
        elif path.is_dir(): candidates += list(path.rglob("*.md"))
    for path in sorted(set(candidates)):
        text = path.read_text(encoding="utf-8")
        for raw in LINK_RE.findall(text):
            target = raw.strip().split()[0].strip("<>")
            if not target or target.startswith(("http://", "https://", "mailto:", "#")):
                continue
            file_part = target.split("#", 1)[0]
            if not file_part:
                continue
            resolved = (path.parent / file_part).resolve()
            try:
                resolved.relative_to(root.resolve())
            except ValueError:
                errors.append(f"{path.relative_to(root)} link escapes repo: {target}")
                continue
            if not resolved.exists():
                errors.append(f"{path.relative_to(root)} broken link: {target}")


def validate_ids(root: Path, errors: list[str]):
    seen = {}
    for path in md_files(root / "docs"):
        fm = frontmatter(path.read_text(encoding="utf-8"))
        if fm is None:
            continue
        doc_id = scalar(fm, "id")
        if not doc_id:
            continue
        if doc_id in seen:
            errors.append(f"duplicate doc id {doc_id}: {seen[doc_id]} and {path.relative_to(root)}")
        else:
            seen[doc_id] = path.relative_to(root)
    valid = set(seen)
    for path in md_files(root / "docs"):
        fm = frontmatter(path.read_text(encoding="utf-8"))
        if fm is None: continue
        for target in list_field(fm, "related"):
            if target and target not in valid:
                errors.append(f"{path.relative_to(root)} unresolved related id {target}")


def validate_research(root: Path, errors: list[str], stats: dict):
    manifests = list((root / "docs/research").glob("*.sources.json")) if (root / "docs/research").exists() else []
    source_count = 0
    for manifest in manifests:
        try: data = json.loads(manifest.read_text(encoding="utf-8"))
        except Exception as exc:
            errors.append(f"{manifest.relative_to(root)} invalid JSON: {exc}")
            continue
        sources = data.get("sources", [])
        if not sources:
            errors.append(f"{manifest.relative_to(root)} empty sources")
            continue
        ids = []
        for item in sources:
            sid = item.get("id") if isinstance(item, dict) else None
            if not sid:
                errors.append(f"{manifest.relative_to(root)} source missing id")
                continue
            ids.append(sid)
            for field in ("title", "url", "publisher", "type", "accessed_at"):
                if not item.get(field): errors.append(f"{manifest.relative_to(root)} source {sid} missing {field}")
        if len(ids) != len(set(ids)): errors.append(f"{manifest.relative_to(root)} duplicate source ids")
        source_count += len(ids)
        report = manifest.with_name(manifest.name.replace(".sources.json", ".md"))
        if report.exists():
            used = set(SOURCE_RE.findall(report.read_text(encoding="utf-8")))
            missing = used - set(ids)
            if missing: errors.append(f"{report.relative_to(root)} undefined sources: {sorted(missing)}")
    stats["source_manifests"] = len(manifests)
    stats["research_sources"] = source_count


def validate_requirements(root: Path, errors: list[str], stats: dict):
    path = root / "docs/product/REQUIREMENTS-TRACEABILITY.md"
    if not path.exists(): return
    ids = set(REQ_RE.findall(path.read_text(encoding="utf-8")))
    stats["requirements"] = len(ids)
    if len(ids) < 200: errors.append(f"requirements traceability unexpectedly small: {len(ids)}")


def validate_generated(root: Path, generated_root: Path, errors: list[str]):
    expected = generated_root / "docs/product/PRODUCT-BLUEPRINT.md"
    actual = root / "docs/product/PRODUCT-BLUEPRINT.md"
    if not expected.exists() or not actual.exists() or expected.read_bytes() != actual.read_bytes():
        errors.append("stale generated Product Blueprint")
    if (generated_root / "docs/roadmap.md").exists():
        errors.append("generator must never emit docs/roadmap.md")


def validate_router(root: Path, errors: list[str]):
    text = (root / "docs/index.md").read_text(encoding="utf-8") if (root / "docs/index.md").exists() else ""
    for needle in ["roadmap.md", "product/README.md", "architecture/index.md", "decisions/index.md", "development/planning-readiness.md", "development/capability-realization.md", "development/engineering-rules.md"]:
        if needle not in text: errors.append(f"docs/index missing required route: {needle}")


def validate(root: Path, generated_root: Path, final: bool) -> tuple[list[str], dict]:
    errors = []
    stats = {}
    for rel in REQUIRED:
        if not (root / rel).exists(): errors.append(f"missing required path: {rel}")
    if (root / "docs/superpowers").exists(): errors.append("docs/superpowers must not exist after MR-01 cutover")
    if (root / "docs/tracking").exists(): errors.append("docs/tracking must not exist after MR-01 cutover")
    if (root / "docs/DOCUMENTATION-MAP.md").exists(): errors.append("legacy Documentation Map path must not exist")
    guard_bootstrap(root, errors)
    guard_program_owner(root, errors)
    guard_root_allowlist(root, errors)
    if final: guard_final_tree(root, errors)
    validate_router(root, errors)
    validate_links(root, errors)
    validate_ids(root, errors)
    validate_research(root, errors, stats)
    validate_requirements(root, errors, stats)
    validate_generated(root, generated_root, errors)
    stats["markdown_files"] = len(md_files(root))
    return errors, stats


def self_test() -> int:
    failures = []
    with tempfile.TemporaryDirectory() as td:
        r = Path(td)
        (r/"docs").mkdir()
        (r/"AGENTS.md").write_text("a")
        (r/"docs/index.md").write_text("i")
        (r/"docs/roadmap.md").write_text("program_status_authority: true\nProduct/runtime implementation: BLOCKED")
        errs=[]; guard_bootstrap(r, errs)
        if errs: failures.append("bootstrap positive")
        (r/"AGENTS.md").write_text("x"*(21*1024))
        errs=[]; guard_bootstrap(r, errs)
        if not errs: failures.append("bootstrap negative")
    errs=[]; guard_review_changed_files(["docs/work/current/ai-dialog.md"], errs)
    if errs: failures.append("review isolation positive")
    errs=[]; guard_review_changed_files(["docs/work/current/ai-dialog.md", "README.md"], errs)
    if not errs: failures.append("review isolation negative")
    with tempfile.TemporaryDirectory() as td:
        r=Path(td); (r/"docs/work").mkdir(parents=True)
        errs=[]; guard_final_tree(r, errs)
        if not errs: failures.append("final-tree negative")
    with tempfile.TemporaryDirectory() as td:
        r=Path(td); (r/"docs").mkdir(); (r/"docs/roadmap.md").write_text("program_status_authority: true\nProduct/runtime implementation: BLOCKED")
        (r/"src").mkdir(); errs=[]; guard_root_allowlist(r, errs)
        if not errs: failures.append("blocked-root allowlist negative")
    with tempfile.TemporaryDirectory() as td:
        out=Path(td)
        rc=subprocess.run([sys.executable, str(ROOT/"scripts/generate_docs.py"), "--output-root", str(out)], cwd=ROOT).returncode
        if rc or (out/"docs/roadmap.md").exists(): failures.append("generator roadmap isolation")
    if failures:
        print("SELF-TEST FAIL:", ", ".join(failures)); return 1
    print("Documentation guard self-tests PASS")
    return 0


def main() -> int:
    ap=argparse.ArgumentParser()
    ap.add_argument("--generated-root", type=Path)
    ap.add_argument("--report", type=Path, default=ROOT/"docs-validation-report.json")
    ap.add_argument("--final", action="store_true")
    ap.add_argument("--self-test", action="store_true")
    ap.add_argument("--review-base")
    args=ap.parse_args()
    if args.self_test: return self_test()
    if args.review_base:
        changed=subprocess.check_output(["git","diff","--name-only",f"{args.review_base}...HEAD"],cwd=ROOT,text=True).splitlines()
        errors=[]; guard_review_changed_files(changed, errors)
        if errors:
            print("\n".join(errors)); return 1
        print("Review isolation PASS"); return 0
    if not args.generated_root:
        print("--generated-root required", file=sys.stderr); return 2
    errors,stats=validate(ROOT,args.generated_root.resolve(),args.final)
    payload={"status":"PASS" if not errors else "FAIL","errors":errors,"statistics":stats,"final_mode":args.final}
    args.report.parent.mkdir(parents=True,exist_ok=True); args.report.write_text(json.dumps(payload,indent=2)+"\n")
    if errors:
        print("DOCUMENTATION VALIDATION FAIL")
        for e in errors: print("-",e)
        return 1
    print("DOCUMENTATION VALIDATION PASS", json.dumps(stats,sort_keys=True))
    return 0


if __name__ == "__main__": raise SystemExit(main())
'''
write("scripts/validate_docs.py", VALIDATOR)

WORKFLOW = r'''name: Documentation

on:
  push:
    branches:
      - main
      - "docs/**"
      - "review/**"
    paths:
      - "**/*.md"
      - "docs/research/*.json"
      - "scripts/generate_docs.py"
      - "scripts/validate_docs.py"
      - ".github/workflows/docs.yml"
  pull_request:
    paths:
      - "**/*.md"
      - "docs/research/*.json"
      - "scripts/generate_docs.py"
      - "scripts/validate_docs.py"
      - ".github/workflows/docs.yml"
  workflow_dispatch:

permissions:
  contents: read

jobs:
  validate:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v4
        with:
          fetch-depth: 0

      - name: Set up Python
        uses: actions/setup-python@v5
        with:
          python-version: "3.12"

      - name: Generate Product documentation projection
        run: python scripts/generate_docs.py --output-root "$RUNNER_TEMP/aurora-generated"

      - name: Run deterministic guard self-tests
        run: python scripts/validate_docs.py --self-test

      - name: Validate documentation system
        id: validate
        shell: bash
        run: |
          FINAL=""
          if [[ ! -d docs/work ]]; then FINAL="--final"; fi
          python scripts/validate_docs.py \
            --generated-root "$RUNNER_TEMP/aurora-generated" \
            --report "$RUNNER_TEMP/aurora-generated/docs-validation-report.json" \
            $FINAL

      - name: Check candidate whitespace/diff range
        shell: bash
        run: |
          if [[ "$GITHUB_EVENT_NAME" == "pull_request" ]]; then
            git fetch origin "${{ github.base_ref }}" --depth=1
            git diff --check "origin/${{ github.base_ref }}...HEAD"
          elif [[ "$GITHUB_REF_NAME" == "main" ]]; then
            git diff --check HEAD^...HEAD
          else
            git diff --check origin/main...HEAD
          fi

      - name: Upload documentation proof
        if: always()
        uses: actions/upload-artifact@v4
        with:
          name: aurora-documentation
          path: |
            ${{ runner.temp }}/aurora-generated/docs/product/PRODUCT-BLUEPRINT.md
            ${{ runner.temp }}/aurora-generated/docs-validation-report.json
          if-no-files-found: error
          retention-days: 14
'''
write(".github/workflows/docs.yml", WORKFLOW)

# Regenerate Product aggregate under new constitutional source.
subprocess.check_call([sys.executable, str(ROOT / "scripts/generate_docs.py")], cwd=ROOT)

# Update migration work route to RM-07. It remains temporary.
work = p("docs/work/current/index.md")
if work.exists():
    t = work.read_text(encoding="utf-8")
    t = set_frontmatter_scalar(t, "version", "0.2.0")
    t = t.replace("RM-01 — CURRENT SEMANTIC + PROVENANCE CENSUS\nstate: IN PROGRESS", "RM-01 — PASS\nRM-02 — APPLIED\nRM-03 — APPLIED\nRM-04 — APPLIED\nRM-05 — APPLIED\nRM-06 — APPLIED\nRM-07 — NEXT / VERIFICATION")
    work.write_text(t, encoding="utf-8")

# One-shot migrator/workflow must not survive its own commit.
rm("scripts/mr01_migrate.py")
rm(".github/workflows/mr01-migrate.yml")
