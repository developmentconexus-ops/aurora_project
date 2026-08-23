---
id: DOC-AURORA-BLUEPRINT-15
title: Governança Documental, Pesquisa e Evolução
document_type: product_blueprint_section
form: explanation
authority: constitutional
status: accepted
accepted_at: 2026-08-23
acceptance_evidence: DOC-AURORA-MR-01-OPERATOR-RATIFICATION
version: 0.3.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - documentation authority and lifecycle
  - research lifecycle and evidence rules
  - source-of-truth protocol
  - conversation-to-canonical promotion
  - documentation validation and continuity
related:
  - DOC-AURORA-DOCUMENTATION-MAP
  - DOC-AURORA-RESEARCH-MAP
  - DOC-AURORA-PRODUCT-INDEX
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DOC-AURORA-DOCUMENTATION-COVERAGE
  - HISTORY-AURORA-ORIGIN-DISCOVERY-2026-08-05
review_triggers:
  - authority hierarchy changes
  - documentation layout or lifecycle changes
  - research source policy changes
  - canonical read path changes
  - implementation/documentation gate changes
last_reviewed: 2026-08-23
---

# 15. Governança Documental, Pesquisa e Evolução

## 15.1 Propósito

Aurora is designed to preserve context, identity and continuity over years. The project documentation must demonstrate those properties before the runtime exists.

Documentation governance prevents:

- a deep conversation remaining the only place where product intent exists;
- research becoming an implicit decision;
- a status file redefining architecture;
- one framework's documentation becoming Aurora doctrine;
- multiple files owning the same concept;
- stale alternatives being read as current direction;
- accepted decisions being silently rewritten;
- agents loading everything and still missing the relevant authority;
- code changing behavior without updating contracts and evaluations;
- failed approaches disappearing and being rediscovered;
- source URLs existing without claim traceability;
- a generated aggregate becoming the edited source;
- a new session reconstructing the project from memory or guesswork.

> **Documentation is part of Aurora's project control plane. It requires identity, authority, lifecycle, ownership, source provenance, validation, supersession and operational gates.**

The first A0 proposal failed this standard by preserving conclusions while compressing most mechanisms, examples and reasoning. The adversarial review and remediation package explicitly correct that failure.

---

## 15.2 Governing principles

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

---

## 15.3 Two classification axes

Every material document has:

### Authority class

What it is allowed to govern.

### Reader form

How the reader should use it.

Reader forms may follow:

```text
Tutorial
How-to
Reference
Explanation
```

Examples:

| Document | Authority | Reader form |
|---|---|---|
| Product Blueprint | Constitutional | Explanation |
| ADR | Decision | Explanation/Reference |
| Capability Spec | Specification | Explanation/Reference |
| Approved Mission Contract | Contract | Reference |
| Security Standard | Standard/Policy | Reference |
| CLI/API reference | Reference | Reference |
| Runbook | Guidance | How-to |
| Research Report | Research | Explanation |
| Acceptance Report | Evidence | Reference/Explanation |
| STATUS | Tracking | Reference |

Form never elevates authority. A clear tutorial does not override an ADR.

---

## 15.4 Authority classes

### A0 — Constitutional

Owns:

- product identity;
- North Star;
- domain principles;
- authority hierarchy;
- constitutional invariants;
- long-horizon scope and non-goals;
- documentation governance.

Examples:

- Product Blueprint sections;
- `docs/index.md` for task/intention routing and authority entrypoints.

### A1 — Decision

Owns one material choice:

- alternatives;
- rationale;
- consequences;
- supersession;
- reconsideration triggers.

Example: ADR selecting one durable engine after spike.

An ordinary ADR cannot silently violate a constitutional invariant.

### A2 — Specification

Owns complete reusable behavior:

- Capability Spec;
- protocol/schema spec;
- subsystem design;
- test/evaluation plan;
- rollout/graduation.

### A3 — Contract

Owns a bounded commitment:

- Mission/Delegation Contract;
- API contract;
- approved environment;
- exact authority/budget;
- closeout contract.

### A4 — Standard / Policy / Golden Path

Owns recurring rules and preferred execution paths.

Examples:

- first-party AHDK policy;
- data classification policy;
- laboratory first-power Golden Path;
- documentation standard.

### A5 — Reference

Describes exact current machinery:

- CLI/API;
- schema;
- state-machine reference;
- provider matrix;
- device registry reference.

### A6 — Guidance

Explains how to use/operate:

- tutorials;
- how-to;
- runbooks;
- contributor guide.

### A7 — Evidence

Records observation:

- test report;
- benchmark;
- architecture-spike result;
- safety drill;
- acceptance/closeout;
- provenance.

Evidence owns what was observed under which conditions, not the general product rule.

### A8 — Program Coordination / Tracking

Owns current repository-program stage/status/next-action only through `docs/roadmap.md`. Issues/boards may coordinate work but do not become parallel Product, architecture or status authority.

### A9 — Research / Historical

Owns:

- investigation;
- source analysis;
- landscape comparison;
- discovery history;
- rejected/superseded proposals;
- legacy maps.

### A10 — Generated Projection

Examples:

- complete Blueprint aggregate;
- generated roadmap;
- static site;
- rendered diagram;
- API docs.

Authority follows the source.

---

## 15.5 Precedence and conflict

Current precedence:

```text
1. Constitutional invariant and current Product Blueprint
2. Accepted specific ADR compatible with constitution
3. Accepted Capability/System Spec
4. Approved scoped Contract
5. Standard/Policy/Golden Path
6. Current implementation-derived Reference
7. Guidance
8. Tracking
9. Research/Historical
10. Generated Projection follows source
```

This ordering does not permit silent contradiction.

A material conflict creates:

```text
DOCUMENTATION_DIVERGENCE
```

Possible responses:

- update stale document;
- propose constitutional change;
- create/supersede ADR;
- replan contract;
- update Capability Spec;
- open Finding;
- block implementation/delegation;
- explicitly accept temporary documented risk.

### Example

Research report recommends A2A as remote binding. An accepted ADR later selects native RPC for the current slice. Research remains valid evidence; implementation follows the ADR/contract.

---

## 15.6 Canonical ownership table

| Durable concept | Canonical owner |
|---|---|
| Product promise/North Star | Blueprint 01 |
| Human–Aurora relationship/personality | Blueprint 02 |
| Domain entities/relationships | Blueprint 03 |
| Cognitive lifecycle/global journeys | Blueprint 04 |
| Capability/Registry/AHDK principles | Blueprint 05 |
| Memory/context principles | Blueprint 06 |
| Harness/delegation boundary | Blueprint 07 |
| Presence/multimodality | Blueprint 08 |
| Device/laboratory principles | Blueprint 09 |
| Autonomy/authority/safety | Blueprint 10 |
| Security/privacy/sovereignty | Blueprint 11 |
| Logical system architecture | Blueprint 12 |
| Reliability/evaluation/self-improvement | Blueprint 13 |
| Product sequence | Blueprint 14 / generated roadmap |
| Documentation/research governance | Blueprint 15 |
| Specific technical choice | ADR under `docs/decisions/adr/` |
| Reusable capability behavior | Capability Spec |
| Scoped implementation commitment | Approved Contract |
| Current repository program/gate/implementation permission/next action | `docs/roadmap.md` |
| Observed proof | Evidence/Acceptance |
| Research finding | Research Report |
| Exact code behavior | Code + generated/current Reference |
| Conversation origin | Discovery History |

---

## 15.7 Document identity and metadata

Metadata is required when it materially improves authority, machine routing, traceability or generation. Repository routers/guidance MAY use path/index routing without a uniform metadata envelope when that is sufficient; existing rich metadata may remain when useful.

Canonical Markdown uses structured frontmatter where applicable.

Example:

```yaml
---
id: DOC-AURORA-BLUEPRINT-06
title: Memory, Knowledge and Context

document_type: product_blueprint_section
form: explanation
authority: constitutional
status: proposed
version: 0.2.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - memory principles
related:
  - DOC-AURORA-BLUEPRINT-03
supersedes: []
superseded_by: null
review_triggers:
  - memory architecture changes
last_reviewed: 2026-08-05
tracking_issue: null
---
```

### Required fields for normative documents

- `id`;
- `title`;
- `document_type`;
- `authority`;
- `status`;
- `version` or revision;
- `owners`;
- `source_of_truth_for`;
- `related`;
- `last_reviewed`.

### Optional fields

- `form`;
- `approvers`;
- `supersedes`;
- `superseded_by`;
- `review_triggers`;
- `generated_from`;
- `tracking_issue`;
- `canonical_environment`;
- `implementation_status`;
- `freshness`.

### Rules

- IDs unique and stable;
- relationships resolve;
- owner exists;
- status allowed for class;
- generated documents declare source;
- accepted normative documents cannot be ownerless;
- Research declares non-normative scope;
- historical document does not masquerade as current.

---

## 15.8 Document lifecycle

General normative lifecycle:

```text
DRAFT
→ PROPOSED
→ ACCEPTED
→ SUPERSEDED | REJECTED | WITHDRAWN
```

### DRAFT

Incomplete working material; not relied upon.

### PROPOSED

Complete enough for review; not yet governing.

### ACCEPTED

Approved by required authority and active.

### SUPERSEDED

Replaced but preserved.

### REJECTED

Reviewed and not accepted; rationale preserved.

### WITHDRAWN

Author/owner removed before decision or because no longer applicable.

### Tracking lifecycle

```text
CURRENT
BLOCKED
ARCHIVED
```

### Research lifecycle

```text
CURRENT
STALE
HISTORICAL
SUPERSEDED
```

A stale report can still be historically useful but cannot support a current technical claim without verification.

---

## 15.9 Product Blueprint architecture

The Product Blueprint uses fifteen modular editable sections.

```text
docs/product/blueprint/
├── 01-product-vision.md
├── 02-human-aurora-relationship.md
├── 03-domain-world-model.md
├── 04-cognitive-lifecycle-journeys.md
├── 05-capability-system.md
├── 06-memory-knowledge-context.md
├── 07-harness-orchestration.md
├── 08-interaction-multimodality-presence.md
├── 09-tools-devices-laboratory.md
├── 10-autonomy-authority-safety.md
├── 11-security-privacy-sovereignty.md
├── 12-system-architecture.md
├── 13-reliability-observability-evaluation.md
├── 14-capability-roadmap.md
└── 15-documentation-research-governance.md
```

### Aggregate publication

```text
docs/product/PRODUCT-BLUEPRINT.md
```

Properties:

- generated/read-only projection;
- source order and versions;
- warning header;
- source hashes when generator exists;
- freshness check;
- convenient for full reading/export;
- never independent authority.

A manual aggregate may be created during documentation-only A0 only if clearly generated from current sources and validated. Long-term, a generator/check becomes required.

---

## 15.10 Documentation layout

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

---

## 15.11 Conversation-to-canonical promotion

Conversation is the discovery workspace.

Potential durable content:

- product decision;
- architecture alternative;
- user preference;
- example scenario;
- correction;
- requirement;
- research question;
- implementation authorization.

Promotion flow:

```text
conversation statement
→ classify content
→ determine scope and authority
→ identify canonical owner
→ create/update proposed document
→ preserve source/history reference
→ review
→ accept or reject
→ update status/coverage
```

### Examples

**“Aurora should be Leandro-first.”**

→ Blueprint 01 constitutional product decision.

**“MCP may be useful for tools.”**

→ Research finding; later ADR if adopted.

**“Use this command now.”**

→ ephemeral interaction, not durable architecture.

**“I prefer detailed technical research before framework choice.”**

→ potential global relationship memory; confirmed by repeated explicit behavior, not Product Blueprint unless made a product rule.

### Discovery History

A detailed historical record preserves examples and reasoning without forcing every conversation sentence into normative sections.

---

## 15.12 Research lifecycle

### Step 1 — Research question

Specific and decision-linked.

Bad:

> “Research AI frameworks.”

Better:

> “Which current agent/workflow runtimes can execute a first-party research harness while preserving Aurora-owned contracts, durable recovery and local-first operation?”

### Step 2 — Scope and exclusions

Define:

- technologies/concepts;
- required currentness;
- environments;
- risks;
- non-goals.

### Step 3 — Source strategy

Priority:

1. specifications/standards;
2. official documentation;
3. primary research papers;
4. official repositories/releases;
5. reproducible benchmark/code;
6. secondary sources only for discovery/context.

### Step 4 — Evidence capture

Record:

- source ID;
- publisher;
- title;
- URL/reference;
- date/version;
- access date;
- claims supported;
- limitations;
- relevant excerpts/sections without excessive copying.

### Step 5 — Analysis

Compare:

- problem fit;
- architecture;
- maturity;
- security;
- durability;
- operational burden;
- language/runtime;
- lock-in;
- gaps;
- contradictions.

### Step 6 — Implications

State what research supports, does not prove and which decisions/spikes it informs.

### Step 7 — Promotion

```text
Research
→ recommendation/alternatives
→ architecture spike if uncertainty is executable
→ ADR/Spec proposed
→ review
→ accepted/rejected
```

Research does not authorize implementation by itself.

---

## 15.13 Focused research reports

Research should be split by decision surface so each report can be refreshed independently.

Initial program:

```text
AURORA-RESEARCH-MEMORY-CONTEXT
AURORA-RESEARCH-HARNESS-INTEROPERABILITY
AURORA-RESEARCH-AHDK-CONFORMANCE-GOLDEN-PATHS
AURORA-RESEARCH-DURABLE-EXECUTION
AURORA-RESEARCH-AUTHORITY-IDENTITY-EFFECTS
AURORA-RESEARCH-EVENTS-OBSERVABILITY-SCHEMAS
AURORA-RESEARCH-AGENT-FRAMEWORKS-RUNTIMES
AURORA-RESEARCH-PRESENCE-MULTIMODALITY
AURORA-RESEARCH-LABORATORY-DEVICE-SAFETY
AURORA-RESEARCH-EVALUATION-SELF-IMPROVEMENT
```

One aggregate synthesis may index conclusions, but cannot replace focused ownership.

---

## 15.14 Source manifests

Every material report has a `.sources.json` or equivalent manifest.

Example:

```json
{
  "research_id": "RESEARCH-AURORA-MEMORY-CONTEXT-V1",
  "accessed_at": "2026-08-05",
  "selection_policy": "Primary sources for normative claims",
  "sources": [
    {
      "id": "S01",
      "title": "...",
      "publisher": "...",
      "url": "...",
      "type": "official_documentation",
      "version_or_date": "...",
      "accessed_at": "2026-08-05",
      "supports": ["claim A"],
      "limitations": ["does not prove project memory quality"]
    }
  ]
}
```

Validation should detect:

- cited source missing manifest;
- manifest source unused;
- unsupported claim category;
- missing access date/version;
- duplicate source IDs;
- stale time-sensitive source.

---

## 15.15 Research freshness

Research includes review triggers such as:

- protocol/spec release;
- framework major version;
- SDK maturity change;
- security advisory;
- benchmark correction;
- new requirement;
- spike result contradicting documentation;
- provider product/policy change.

A current date in Git does not prove current content. Review uses source version and claim applicability.

---

## 15.16 Architecture Decision Records

ADR contains:

- context/problem;
- requirements;
- research/evidence;
- alternatives;
- decision;
- rationale;
- consequences;
- risks/mitigations;
- compatibility/migration;
- review triggers;
- supersession.

ADR does not repeat an entire Blueprint section. It owns a specific choice.

### ADR status

```text
PROPOSED
ACCEPTED
REJECTED
SUPERSEDED
WITHDRAWN
```

### Constitutional change

If a decision changes a constitutional principle, the Blueprint is updated through explicit constitutional review. An ADR alone cannot override it.

---

## 15.17 Capability Specifications

Future path:

```text
docs/capabilities/CAP-*/
├── SPEC.md
├── REQUIREMENTS.md
├── TEST-PLAN.md
├── THREAT-MODEL.md when applicable
├── ROLLOUT.md
└── IMPLEMENTATION-HISTORY.md
```

A Capability Spec owns reusable design and evaluation, independent of one implementation Mission.

It should include:

- purpose/use cases;
- applicability;
- goals/non-goals;
- domain model;
- contracts/lifecycle;
- architecture;
- authority/security;
- failure/recovery;
- observability;
- evaluation;
- rollout/graduation;
- open questions;
- requirements traceability.

Create a Capability Spec when its Product Milestone approaches readiness, not for every idea in the roadmap.

---

## 15.18 Capability Realization Method

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

---

## 15.19 Repository-program coordination and history

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

---

## 15.20 Status and authorization vocabulary

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

---

## 15.21 Supersession

Material change preserves history.

Steps:

1. create proposed revision/new document;
2. state what changes and why;
3. analyze dependents;
4. review/approve;
5. mark old source superseded;
6. update indexes/relations/projections;
7. preserve Git history;
8. re-evaluate contracts/implementations where required.

Do not rewrite rejected rationale out of history. Git is archive only when required provenance remains reachable; unique unmerged provenance with a current consumer receives a durable ref before the last branch/reference is deleted.

For machine-readable contracts, accepted revisions should be immutable/content-addressed where appropriate.

---

## 15.22 Generated projections

Keep the Product aggregate:

```text
docs/product/PRODUCT-BLUEPRINT.md
← generated from Blueprint 01–15
```

`docs/roadmap.md` is **not** generated from Blueprint 14 after MR-01. It is a hand-maintained repository-program authority. Blueprint 14 remains the Product capability-roadmap authority; the two intentionally answer different questions.

Generated files declare provenance and are validated for freshness. No generated projection becomes an independent semantic owner.

---

## 15.23 Documentation impact

Every material change declares:

```yaml
documentation_impact:
  status: NONE | UPDATED | FOLLOW_UP_REQUIRED
  affected: []
  rationale: ""
  follow_up: null
```

`NONE` requires a specific explanation.

Examples of material change:

- domain behavior;
- API/contract/schema;
- authority/security;
- memory behavior;
- provider capability;
- roadmap/gate;
- operation/recovery;
- user-visible interaction;
- evaluation.

A PR cannot claim no impact merely because documentation is inconvenient.

---

## 15.24 Documentation and repository checks

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

---

## 15.25 Adversarial documentation review

Before accepting a baseline or major capability, reviewer asks:

- What would a new session misunderstand?
- Which concepts exist only in conversation?
- Are conclusions present without rationale/examples?
- Does a list hide missing behavior?
- Can implementation satisfy wording while violating intent?
- Which source owns each term?
- What is explicitly open?
- Are failure/recovery/security addressed?
- Is research current and primary?
- Does the roadmap have executable proof?
- Is implementation accidentally authorized?

The A0 adversarial review is preserved under `docs/evidence/reviews/`.

---

## 15.26 Human and agent read paths

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

---

## 15.27 Agent read budgets and Context Packs

Agents receive role/scoped packs.

### Lead/architect

Current status, authority, relevant Blueprint/ADRs/research.

### Harness implementer

Capability Spec, contract, AHDK/standard, exact task.

### Reviewer

Fixed diff/version, criteria, relevant rules and threat/evaluation context.

### QA/evaluation

Journey, environment, expected observations and evidence profile.

Full product context remains discoverable but not automatically injected.

---

## 15.28 Golden Proof — fresh-session continuity

A0 closes only after a fresh session using repository only can:

1. state what Aurora is and is not;
2. reproduce the North Star and major scenarios;
3. identify Leandro-first scope;
4. explain relationship/personality/proactivity;
5. distinguish memory classes and source authority;
6. explain Capability/Provider/Harness/AHDK;
7. explain hierarchy and child Delegation;
8. describe autonomy envelope and self-improvement boundary;
9. identify local-first/cloud-assisted policy;
10. identify laboratory progression and physical interlock principle;
11. find current roadmap/gate;
12. list open technical decisions;
13. refuse implementation because it is not authorized;
14. point to the exact next review action.

The test result is Evidence and should record misunderstandings/corrections.

---

## 15.29 Ownership

Initial owner:

```text
developmentconexus-ops
```

Operator approval remains a domain authority distinct from Git write permission.

Future CODEOWNERS may protect:

```text
/docs/product/
/docs/decisions/adr/
/docs/capabilities/
/docs/development/
/docs/development/
/docs/research/
/docs/evidence/
AGENTS.md
contract/schema directories
security/policy paths
```

A bot or model with repository write permission cannot self-approve constitutional change.

---

## 15.30 Documentation incident

Examples:

- accepted documents contradict;
- implementation changed without spec;
- source manifest contains dead/incorrect source;
- generated aggregate stale;
- agent implemented from research instead of ADR;
- conversation approval never promoted;
- sensitive data committed;
- old provider docs caused unsafe command.

Response:

```text
contain/block relevant implementation
→ open documentation Finding/Incident
→ identify canonical owner
→ correct/supersede
→ update dependent context/evals
→ preserve evidence
→ review process failure
```

Documentation defects can be product/security defects.

---

## 15.31 A0 acceptance and post-A0 current-state ownership

A0 acceptance remains historical fact: the fifteen-section Product constitution, discovery coverage, research, ADR baseline, ACRM, traceability and fresh-session/adversarial proof were explicitly accepted and merged.

After MR-01 migration, mutable repository-program state is not owned by this constitutional section. `docs/roadmap.md` owns the current stage/gate, blockers, implementation permission and exact next action. `docs/index.md` owns task/intention routing. Product meaning remains with the Product Blueprint and accepted specific owners.

A0 acceptance, MR-01 ratification or repository migration never authorizes later TA stages, Architecture Spike execution or Product implementation by implication.

---

## 15.32 Non-goals

- document every casual message;
- maximize Markdown volume;
- force full Blueprint into every model context;
- use Git issues as architecture;
- treat research recommendation as accepted choice;
- erase rejected/superseded history;
- create directories/documents without consumers;
- freeze technology forever;
- let documentation delay every trivial implementation detail once proper authority exists;
- accept generated content without source validation;
- allow a writing agent to approve its own constitutional change.
