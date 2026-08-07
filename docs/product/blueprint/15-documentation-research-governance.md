---
id: DOC-AURORA-BLUEPRINT-15
title: Governança Documental, Pesquisa e Evolução
document_type: product_blueprint_section
form: explanation
authority: constitutional
status: proposed
version: 0.2.1
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
last_reviewed: 2026-08-06
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

A chat approval or insight must be promoted to the correct document when it is meant to endure.

### P3 — Research is evidence, not authority

A primary source can support a decision without becoming the decision.

### P4 — Tracking coordinates current work, not product doctrine

`STATUS.md` can say which ADR is current; it cannot create the architecture itself.

### P5 — Historical information remains discoverable

Rejected, superseded and failed approaches are preserved with status.

### P6 — Accepted normative content cannot contain hidden placeholders

Open questions are explicit and owned by research/spike, not vague `TBD` text.

### P7 — Generated projections derive authority from sources

They are never edited directly.

### P8 — Documentation depth is mechanism-driven

Length is not a goal. Purpose, boundaries, flows, examples, failures, evaluation and non-goals are required where material.

### P9 — New sessions load the smallest correct authority set

`AGENTS.md` is an index and hard-rule bootstrap, not the complete doctrine.

### P10 — Implementation is gated by accepted intent and evidence

A repository containing detailed architecture does not authorize code.

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
- Documentation Map for authority/read paths.

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

### A8 — Tracking

Coordinates current work:

- issue;
- project board;
- status;
- worklog;
- checklist;
- coverage matrix.

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
| Specific technical choice | ADR |
| Reusable capability behavior | Capability Spec |
| Scoped implementation commitment | Approved Contract |
| Current project coordination | STATUS/tracking |
| Observed proof | Evidence/Acceptance |
| Research finding | Research Report |
| Exact code behavior | Code + generated/current Reference |
| Conversation origin | Discovery History |

---

## 15.7 Document identity and metadata

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

Target layout grows when a real document exists:

```text
README.md
AGENTS.md
CONTRIBUTING.md

.github/
├── CODEOWNERS
├── pull_request_template.md
└── workflows/
    └── docs.yml

docs/
├── DOCUMENTATION-MAP.md
├── roadmap.md
│
├── product/
│   ├── README.md
│   ├── PRODUCT-BLUEPRINT.md
│   ├── CAPABILITY-REALIZATION-METHOD.md
│   ├── REQUIREMENTS-TRACEABILITY.md
│   └── blueprint/01..15
│
├── adr/
├── capabilities/
├── standards/
├── golden-paths/
├── reference/
├── how-to/
├── tutorials/
├── research/
├── design/
├── reviews/
├── acceptance/
├── evidence/
├── history/
└── tracking/
```

Do not create empty directories solely to match the diagram.

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

The canonical method connects:

```text
Blueprint intent
→ applicability
→ requirements
→ research/ADR/spikes
→ Capability Spec
→ Mission Contract
→ implementation plan
→ code/config
→ evidence
→ Product Milestone closeout
```

It defines readiness gates and orphan detection:

- requirement without implementation/evidence;
- code without requirement;
- decision without research/impact;
- milestone without Golden Proof;
- evidence without criterion.

The separate method document owns details.

---

## 15.19 Tracking documents

### STATUS.md

Required fields:

- program/product phase;
- current branch/PR/issue;
- accepted/proposed artifacts;
- current gate/readiness;
- authorizations and prohibitions;
- blockers;
- verification evidence;
- exact next action.

### WORKLOG.md

Chronological material work:

- what changed;
- why;
- files/commits/PR;
- validation;
- unresolved items.

It does not own architecture.

### DECISIONS.md

Concise index linking to canonical decisions/Blueprint approvals.

### BACKLOG.md

Ideas, research questions and future capabilities without commitment.

### DOCUMENTATION-COVERAGE.md

Maps discovery/requirements to canonical owners and current coverage.

Tracking is updated at every material handoff.

---

## 15.20 Status and authorization vocabulary

A document can exist while implementation remains prohibited.

Status should explicitly separate:

```text
DISCOVERY AUTHORIZED
RESEARCH AUTHORIZED
DESIGN AUTHORIZED
SPIKE AUTHORIZED
PLAN AUTHORIZED
IMPLEMENTATION AUTHORIZED
EXTERNAL EFFECT AUTHORIZED
MERGE AUTHORIZED
```

Absence of prohibition is not authorization.

A fresh session reads current authorization before taking action.

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

Do not rewrite rejected rationale out of history.

For machine-readable contracts, accepted revisions should be immutable/content-addressed where appropriate.

---

## 15.22 Generated projections

Possible projections:

| Projection | Source |
|---|---|
| `PRODUCT-BLUEPRINT.md` | 15 modular Blueprint sections |
| `roadmap.md` | Blueprint 14 |
| static documentation site | canonical Markdown |
| API/reference | schemas/code |
| diagrams | structured source/Mermaid |
| review UI | plan/contract revisions |

Generated file includes:

- warning;
- source list/order;
- generation time/tool version;
- source hashes when possible;
- no independent edits.

CI verifies freshness.

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

## 15.24 Documentation checks

Required checks evolve from A0 to implementation.

### Structure

- Markdown/frontmatter parse;
- unique IDs;
- valid statuses/classes;
- owners;
- related links;
- file naming;
- section coverage.

### Authority

- one canonical owner;
- no accepted conflict;
- supersession consistency;
- generated source declaration;
- tracking/research not treated as normative.

### Research

- source manifest parse;
- cited/defined source matching;
- access dates;
- freshness triggers;
- limitation section;
- claim-source coverage.

### Quality

- no unresolved `TBD/TODO/FIXME` in accepted normative docs;
- no empty sections;
- no list-as-design for material concept;
- examples and failure modes where needed;
- open decisions explicit;
- non-goals;
- review triggers.

### Traceability

- requirements mapped;
- ADR/Spec/Contract links;
- evidence-to-criteria;
- code/documentation impact;
- roadmap milestone coverage.

### Projection

- Blueprint aggregate freshness;
- roadmap freshness;
- generated warning;
- source hashes/order.

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

The A0 adversarial review is preserved under `docs/reviews/`.

---

## 15.26 Human read paths

### First-time product reader

```text
README
→ Blueprint 01
→ Blueprint 02
→ Blueprint 14
```

### Architecture reader

```text
Documentation Map
→ Product index
→ Blueprint 03–13
→ ADR index
→ research map
```

### Current contributor/session

```text
AGENTS
→ STATUS
→ Documentation Map
→ current milestone/capability documents
→ relevant ADR/research/design
```

### Research reviewer

```text
Research Map
→ focused report
→ source manifest
→ decision/spike it informs
```

### Implementation worker

```text
AGENTS
→ STATUS and authorization
→ accepted Capability Spec/Contract
→ implementation plan
→ exact interfaces/tests
```

Do not load full Blueprint into every worker by default.

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
/docs/adr/
/docs/capabilities/
/docs/standards/
/docs/golden-paths/
/docs/research/
/docs/acceptance/
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

## 15.31 A0 acceptance rule and post-A0 current-state ownership

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
