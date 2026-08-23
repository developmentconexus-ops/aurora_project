---
id: DESIGN-AURORA-METHODOLOGY-REPOSITORY-REBASELINE-CANDIDATE
title: Aurora Methodology, Planning-Readiness and Repository Rebaseline Candidate
document_type: temporary_design_candidate
form: explanation
authority: tracking
status: proposed
version: 0.2.0
owners:
  - developmentconexus-ops
related:
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DESIGN-AURORA-TECHNICAL-ARCHITECTURE-BASELINE-MAP
  - DESIGN-AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY
  - DOC-AURORA-BLUEPRINT-15
last_reviewed: 2026-08-23
---

# Aurora Methodology, Planning-Readiness & Repository Rebaseline — Candidate v0.2

> **NON-AUTHORITATIVE / BRANCH-ONLY.** This is the Lead candidate for operator review. It must be absorbed into durable owners or deleted before any merge candidate is promoted.

## 1. Decision question

> How should Projeto Aurora adopt the engineering and repository discipline that matured across MetalDocs, Marketplace Central and Conexus OS so future implementation becomes constrained realization of accepted Product, architecture and contracts rather than a second system-design phase performed while coding?

The rebaseline has two co-equal scopes:

```text
A. ENGINEERING / PLANNING
   what must be known, decided, contracted and falsifiable
   before Product implementation may begin

B. REPOSITORY OPERATING MODEL
   how current authority, status, decisions, Evidence, temporary work,
   Git/PR/review and fresh-session context are organized
```

This rebaseline changes method, planning order and repository routing. It does not redefine Aurora Product meaning by convenience.

---

## 2. Revalidated starting point

```text
repository                    developmentconexus-ops/aurora_project
revalidated main              35614c581cea32e04305c1ad63522fee151eb283
Product implementation        PAUSED / BLOCKED
Architecture Spike execution  NOT AUTHORIZED
M0 R7 candidate               FROZEN / PRESERVED / NON-CANONICAL
M0 R7 Verdict                 NOT ISSUED
M0 R8                         NOT AUTHORIZED
TA-01                         ACCEPTED / MERGED / CANONICAL
TA-02                         ACCEPTED / MERGED / CANONICAL
old TA-03                     NOT AUTHORIZED
```

No current Evidence requires reopening the semantic result of TA-01/TA-02:

```text
G01 Contract Model Governance
C01–C12 canonical domain owners
A01–A05 coordination/lifecycle responsibilities
B01 provider-runtime lifecycle/reconciliation boundary
E01 Effect Gateway
E02 Credential Broker

Stage A:
one small persistent Evolutionary Sovereign Host
+ one on-demand provider-runtime seam at first consumer
+ specialized Harnesses as independent providers
```

---

# 3. Current organizational authorities

Aurora's August 13 baseline predates the current cross-repository standards.

## 3.1 DevelopmentConexus Engineering Method v1.0.0

Canonical external authority:

```text
developmentconexus-ops/conexus-methodology/METHOD.md
```

It owns **how material engineering reasoning is performed**:

```text
Evidence
→ Known / Inferred / Unknown / Deferred
→ Root Cause
→ Target Invariant
→ Constraints
→ Credible Alternatives
→ Local Maximum vs Global Maximum
→ Essential vs Accidental Complexity
→ YAGNI / Future Cost
→ Authority / Boundary
→ Enforcement
→ Proof Strategy
→ Adversarial Challenge
→ Decision
→ Reopen Triggers
```

Aurora consumes these laws rather than maintaining a local duplicate:

- mechanism != authority;
- current implementation is Evidence, not target authority by existence;
- unknown remains unknown;
- prepare the seam, not the entire future capability;
- proof strategy precedes material implementation;
- Global Maximum does not mean maximum infrastructure;
- material architecture is challenged adversarially;
- accepted decisions reopen only on material Evidence.

## 3.2 DevelopmentConexus Repository Standard v1.0.0

Canonical external authority:

```text
developmentconexus-ops/conexus-methodology/REPOSITORY-STANDARD.md
```

It owns the common repository operating envelope:

```text
AGENTS.md
→ docs/index.md
→ docs/roadmap.md
→ 1–2 task-specific owning documents
```

Default task pack: five files or fewer unless a named material reason requires more.

Required role separation:

```text
README.md       landing only
AGENTS.md       compact bootstrap/router + local hard stops
docs/index.md   task/intention router only
docs/roadmap.md sole mutable stage/status/allowed-work/next-action authority
```

Temporary work is branch-only. Git/closed PRs are history subject to provenance reachability. One coherent gate owns one branch/PR by default. Independent Fable review is isolated and reviewer output is Evidence, never authority.

---

# 4. Evidence from the three mature product repositories

These repositories are references/Evidence, not Aurora authority.

## 4.1 MetalDocs

Current target planning evolved to:

```text
Product / ownership
→ technical baseline
→ backend boundaries
→ internal owner contracts
→ persistence
→ executable wire
→ frontend realization
→ runtime/process/deployment
→ whole-system coherence
→ Golden Flows / validation
→ transition/cutover
→ implementation graph
→ adversarial implementation-readiness
→ explicit execution grant
```

Transferable properties:

- existing implementation receives no preservation right from sunk cost;
- reuse requires a named current consumer and proof-backed fit;
- executable wire is closed before backend/frontend implementation;
- frontend is bidirectionally traced to accepted Product operations;
- runtime topology is derived from real consumers/failure classes;
- implementation remains blocked after architecture closure until readiness closes.

## 4.2 Marketplace Central

Transferable properties:

- Product, ownership, identity, communication, integration, API, frontend and runtime are separate authorities;
- one machine-readable Product wire authority prevents DTO/API drift;
- a decision-reconciliation map prevents retired ideas from returning through implementation convenience;
- Golden Flows are selected by defect class, not exhaustive operation count;
- real external probes are used only where a material claim cannot be closed documentarily;
- frontend cannot create Product/API authority merely to fit a screen.

## 4.3 Conexus OS

The strongest current implementation-readiness compilation model is:

```text
accepted Product / architecture
→ Product Surface & Authority Contract
→ Executable Wire Contract
→ Frontend Interaction & Authority Realization
→ Paved Road + runtime/persistence realization
→ Whole-System Coherence & Golden Flows
→ Implementation Program & Execution Graph
→ Adversarial Implementation Readiness
→ explicit Product execution grant
→ implementation
```

The Paved Road reduces implementation degrees of freedom through ownership classes conceptually equivalent to:

```text
GENERATED
PROTECTED PLATFORM/PRODUCT CONTRACT
APP/MODULE-OWNED
```

Conexus OS also uses property-first technology realization:

```text
protected property
→ normative/current official Evidence
→ credible alternatives
→ ADOPT | ADAPT | BUILD | DEFER | STOP
→ exact boundary
→ proof/falsifier
```

## 4.4 Frontend planning reference

Current reusable reference:

```text
developmentconexus-ops/marketplace-central/
docs/development/frontend-product-experience-planning-method.md
version 2.1
```

It is currently **reference methodology**, not yet an authority published from `conexus-methodology`.

Aurora will not fork/copy it now. Before the first material visual Presence/frontend implementation-readiness stage, either:

1. the then-current frontend method is promoted cross-repository through its own organizational decision; or
2. Aurora explicitly adopts an exact external reference for that consuming gate.

No frontend work is opened by this rebaseline.

---

# 5. Approaches considered

## A — Minimal citation patch

Cite Method/Repository Standard but preserve the existing Aurora tree and old TA-03→TA-08 order.

**Rejected:** leaves parallel status authority, no task router, permanent temporary-plan surfaces and repository topology too early in planning.

## B — Repository cleanup only

Align `AGENTS/index/roadmap/docs` but retain old architecture-stage order.

**Rejected:** fixes session behavior but not the deeper implementation-readiness defect.

## C — Full methodology + repository + readiness rebaseline

```text
external Method + Repository Standard
+ preserve Aurora Product / TA-01 / TA-02
+ boundedly refine documentation constitution
+ reframe ACRM as capability/slice lifecycle
+ replace old TA-03+ planning order
+ add operation surface, executable contracts, Paved Road,
  Golden Flows, implementation graph and final readiness attack
+ migrate repository routing/status/CI deliberately
```

**Recommended: ADOPT.**

---

# 6. Method authority hierarchy

The target deliberately prevents four methods from owning the same thing.

```text
DevelopmentConexus Engineering Method v1
  HOW WE REASON
        ↓
DevelopmentConexus Repository Standard v1
  HOW REPOSITORY AUTHORITY / CONTEXT / DELIVERY OPERATES
        ↓
Aurora Product + architecture authorities
  WHAT AURORA MEANS AND OWNS
        ↓
Aurora System Planning & Implementation-Readiness Program
  WHICH CROSS-SYSTEM DECISIONS MUST CLOSE BEFORE CODING
        ↓
Specialized methods when triggered
  frontend / research / qualification / transition
        ↓
ACRM R0–R8
  CAPABILITY / SLICE REALIZATION + EVIDENCE LIFECYCLE
        ↓
Implementation Graph
        ↓
Constrained coding agents
```

## 6.1 ACRM disposition

`CAPABILITY-REALIZATION-METHOD` is **REFINED, not rejected**.

Preserve R0–R8 semantics for:

- applicability;
- capability requirements;
- capability readiness;
- capability-specific research/decision/spike applicability;
- scoped Mission/implementation contracts;
- microdesign;
- execution/evidence;
- milestone Golden Proof/closeout.

Refine its scope so it does not duplicate:

- cross-repository reasoning Method;
- repository operating model;
- system-wide readiness-stage ordering.

A capability may consume accepted global architecture and refine its local behavior. It may not silently override a global owner/contract/readiness decision.

---

# 7. Target implementation invariant

> **Implementation must not be the stage where material Product, architecture, contract, security, data, runtime or UX decisions are discovered by convenience.**

Operational law:

```text
material missing decision
→ STOP
→ identify smallest owning authority/stage
→ reopen only that owner
→ recompile affected downstream artifacts
→ resume from reconciled authority
```

Forbidden:

```text
material unknown
→ plausible default in code
→ implementation convenience becomes architecture
```

Implementation may still choose local algorithms and reversible internals inside exact accepted boundaries.

---

# 8. Revised Aurora system planning/readiness graph

The order is dependency-bearing, not a ritual waterfall.

## Tailoring law

- A stage exists only when its protected property has a real current-horizon consumer.
- No empty artifact is created merely to satisfy numbering.
- Depth scales with materiality, irreversibility, uncertainty and blast radius.
- Adjacent stages may share one coherent branch/review package **only when the operator explicitly accepts that no independently material authority gate would be lost**.
- A later Product Milestone consumes already-current stages rather than replaying inception.
- A bounded change reopens only the smallest implicated authority.

## MR-01 — Methodology, Planning-Readiness & Repository Rebaseline

**Current gate.** Owns this method/repository reconciliation and the migration target. Does not authorize technical realization.

## TA-01 — Logical Modules & Canonical Ownership

```text
PRESERVE / CLOSED / CANONICAL
```

G01, C01–C12, A01–A05, E01/E02 and canonical owner/write boundaries.

## TA-02 — Process, Runtime & Evolutionary Topology

```text
PRESERVE / CLOSED / CANONICAL
```

Evolutionary Sovereign Host, initial provider-runtime seam, Stage A→B principles, process split triggers and A05/B01 lifecycle ownership.

## TA-03 — Cross-System Operation Surface

Own the **smallest complete material semantic operation surface for the current implementation horizon**, not every internal method of every Capability and not the whole future Aurora universe.

An operation belongs here when it crosses/protects a global owner, actor, trust, provider, canonical-state or human-consumer boundary.

Capability-local internal behavior remains owned by Capability Specs.

For each admitted material operation record as applicable:

```text
operation identity
caller/actor class
semantic owner
named consumer
input/output meaning
read | command | proposal | effect | observation
scope / authority condition
state transition or explicit non-transition
knowledge/outcome class
semantic idempotency / concurrency need
failure / ambiguity meaning
audit / Evidence obligation
explicit non-ownership
```

Exit:

```text
current-horizon cross-system consumer without operation       0
ownerless material operation                                 0
unjustified operation with no named consumer                 0
protocol/mechanism operation masquerading as Product meaning 0
```

## TA-04 — Executable Contract Model & Wire

Compile TA-03 semantics into Aurora-owned machine-checkable contract families before implementation frameworks create parallel DTO/protocol authority.

Own:

- G01 contract families and versions;
- machine-checkable boundary schemas/types;
- errors/problems;
- deadlines/cancellation;
- idempotency/correlation/causation;
- provider restart/reconciliation identities;
- streaming semantics where required;
- compatibility/deprecation;
- generated projection law;
- semantic/schema/binding/SDK/provider version separation.

Bindings are chosen per real boundary only after requirements:

```text
in-process
local IPC
HTTP/OpenAPI
Connect/gRPC
SSE/WebSocket
message/event transport
MCP/A2A/other external bindings
```

No universal protocol is required.

## TA-05 — Human Interaction & Presence Realization

Prove current Leandro-facing flows against accepted operations/contracts before package/runtime choices freeze around an incomplete interaction.

Aurora interaction may include:

```text
activation
text interaction
session/context continuation
clarification / decision request
Mission/progress visibility
Artifact/Evidence inspection
failure/recovery interaction
voice or handoff only when current-horizon consumer exists
```

Own:

- human goals and end-to-end flows;
- operation↔interaction coverage;
- human/Presence state authority;
- failure-message intent;
- Stage A Presence-shell semantics needed by the current horizon;
- frontend specialized-method trigger.

No screen/UX convenience creates an upstream operation.

## TA-06 — State, Data, Memory, Knowledge, Artifact & Evidence Architecture

Classify before selecting physical stores:

```text
canonical operational state
governed memory
knowledge/source documents
interaction / exact history
artifacts
observations / measurements
evidence / receipts / verdicts
audit
telemetry
derived retrieval indexes
ephemeral active context
provider-local runtime state
secret / credential references
```

For each, close owner/writer, canonical-vs-derived, consistency/freshness, access pattern, confidentiality/integrity, retention/deletion, rebuildability, backup/export/restore, migration, scale where deciding and co-location/separation criteria.

Only then compare SQLite, PostgreSQL, object/file stores, full-text/vector/graph or other mechanisms per role.

## TA-07 — Identity, Authentication, Authorization, Policy, Secrets & Effects

Close actor identity proof and deterministic authority/effect paths:

- actor/identity classes;
- stable identity vs runtime incarnation;
- human/session authentication;
- service/Presence/Harness/provider/device authentication when current;
- Authority Grants/delegation;
- policy-decision contract;
- E01 enforcement;
- secret classes/references/delivery;
- E02 credential broker;
- revocation/recovery/audit;
- Stage A→B changes.

Evaluate OIDC/Keycloak/Zitadel/Authentik/Ory/SPIFFE/Cedar/OPA/Vault-like or minimal mechanisms only from protected properties and current Evidence.

## TA-08 — Cognitive Runtime, Models, Context, Memory Use & Harness Realization

Close:

- deterministic vs model-mediated responsibilities;
- Context Builder authority/inputs;
- model-provider roles;
- routing/fallback semantics;
- governed memory use/promotion path;
- tool/capability request path;
- Harness Delegation lifecycle;
- provider-local/global state reconciliation;
- exact Mastra fit for the first real consumers;
- restart/cancellation/ambiguous completion;
- Artifact/Evidence return.

Mastra remains preferred-first to evaluate within exact accepted scope; it never inherits sovereign ownership by convenience.

## TA-09 — Paved Road, Repository, Source, Build & Dependency Realization

This stage supersedes the **old early TA-03** repository/source/build placement.

Only after operations/contracts/consumers/data/security/cognition are sufficiently exact does Aurora decide the coding environment.

Required ownership classes, naming refineable:

```text
GENERATED
  reproducible projection; no hand-owned divergent semantics

AURORA-CONTRACT
  Aurora-controlled protected seam; implementer may consume but not silently weaken

MODULE-OWNED
  module-specific Product/implementation logic a coding agent may legitimately evolve
```

Own:

- monorepo/polyrepo/staged strategy;
- source/package/module layout;
- Go/TypeScript/other placement;
- contract/schema source and generated projections;
- dependency directions/static guards;
- build/workspace tooling;
- exact dependency/version admission rules;
- test layers and CI aggregate;
- release/versioning relationship;
- deterministic scaffold/Paved Road;
- relationship to Conexus OS as software-development Harness;
- bounded escape hatch for a real unsupported property.

Every material dependency follows `ADOPT | ADAPT | BUILD | DEFER | STOP` with proof.

## TA-10 — Runtime, Configuration, Deployment, Observability & Recovery Realization

Select concrete operational mechanisms for the accepted topology:

- runtime shells/process artifacts;
- service supervision;
- config precedence;
- secret-reference injection;
- startup/readiness/liveness/shutdown;
- migrations;
- packaging/install/update/rollback;
- Stage A Windows/Linux support posture;
- Stage B deployment posture;
- logs/traces/metrics vs Audit/Evidence;
- OTel export/backend only when current;
- backup/restore/drills;
- supply-chain/artifact provenance;
- resource/latency budgets when deciding.

## TA-TX — Transition / Cutover — conditional

```text
CURRENT DISPOSITION: DEFER SAFELY
```

Aurora currently has no accepted production runtime whose business/operational continuity must be migrated. The frozen M0 R7 branch is Evidence, not current deployed authority.

Open TA-TX only when a real current installation/state/compatibility consumer must survive movement to the accepted target.

It then owns cutover, point of no return, compatibility entitlement, rollback/recovery barriers and retirement of old realization.

## TA-11 — Whole-System Coherence & Golden Flows

Select the smallest representative positive/negative/recovery flows by **defect class**.

The exact set is derived later. Candidate risk families include sovereign restart, Intent/Mission, context→cognition→canonical commit, provider ambiguity/reconciliation, memory supersession/deletion and effect denial/receipt.

Real runtime/external probes remain separately authorized.

## TA-12 — Implementation Program & Execution Graph

Each implementation slice receives:

```text
operator-visible outcome
exact operations/consumers
owners/modules
contracts/generated projections
allowed state/persistence surfaces
allowed packages/source ownership class
exact dependencies/pins when material
GENERATED/AURORA-CONTRACT/MODULE-OWNED mutation rights
config/migration/process surfaces allowed to change
proof/falsifier obligations
prerequisites
stop/reopen triggers
completion criteria
```

No slice selects its own foundational stack.

## TA-13 — Adversarial Implementation Readiness

Fresh independent challenge attacks at least:

```text
consumer without operation
operation without owner
owner duplication
missing executable contract
parallel DTO/wire authority
screen/Presence-shaped API
provider/framework semantic leakage
store without owner/invariant
auth/effect path without deterministic enforcement
dependency without protected property
repository/package ambiguity
unowned generated code
Paved Road bypass
missing failure/recovery semantics
non-falsifiable Golden Flow
implementation slice still containing a material architecture decision
```

Exit requires zero unresolved material readiness findings, converged independent challenge and operator ratification. A separate explicit Product execution grant is still required afterward.

---

# 9. Product Milestone roadmap vs repository roadmap

Two different meanings must no longer share one path/name.

```text
docs/product/blueprint/14-capability-roadmap.md
= canonical long-horizon Product Milestone sequence and Golden-Proof direction

docs/product/PRODUCT-BLUEPRINT.md
= optional generated aggregate containing Blueprint 14

docs/roadmap.md
= sole mutable current repository/program stage/status/allowed-work/next-action authority
```

Current `scripts/generate_docs.py` must stop generating `docs/roadmap.md` during repository migration.

The existing Product milestone sequence is not changed by this naming/authority correction except where explicit current terminology/status references are refined.

---

# 10. Bounded constitutional reopen — Blueprint 15

The current accepted Blueprint 15 itself owns documentation lifecycle/layout and explicitly establishes the old operating model:

```text
STATUS as current coordination owner
WORKLOG / DECISIONS / BACKLOG / DOCUMENTATION-COVERAGE tracking model
large Documentation Map
old documentation tree
generated docs/roadmap.md projection
```

Therefore Repository Standard alignment cannot be implemented as a non-semantic file move.

Required disposition:

```text
Blueprint 15 principles: PRESERVE
Blueprint 15 repository-operating clauses: BOUNDED REOPEN / REFINE
```

Preserve:

- one durable concept, one canonical owner;
- conversation is discovery, repository is canonical project memory;
- research is Evidence, not authority;
- generated projection does not gain authority;
- accepted decisions preserve provenance/supersession;
- new sessions load the smallest correct authority set;
- implementation is gated explicitly;
- document identities/ownership/lifecycle remain machine-discoverable where useful.

Refine:

- fresh actor route;
- status owner;
- documentation layout;
- tracking model;
- generated roadmap collision;
- temporary work/review treatment;
- Git-as-archive reachability;
- external Method/Repository Standard relationship.

The final durable rebaseline package must include the exact Blueprint 15 revision and operator ratification. An ADR cannot override this constitutional documentation authority alone.

---

# 11. Other current decision refinements required

## 11.1 Generated roadmap decision

Existing D-050 currently says both aggregate Product Blueprint **and roadmap** are generated from modular sources.

Required refinement:

```text
PRESERVE generated Product Blueprint aggregate
SUPERSEDE generated standalone repository docs/roadmap.md
PRESERVE Blueprint 14 as Product capability-roadmap authority
ESTABLISH docs/roadmap.md as mutable program/status authority
```

## 11.2 MNFS → Conexus OS identity

Current Aurora constitutional/examples still name **MNFS** as the future software-engineering Harness. The current product is **Conexus OS**, formerly MNFS.

Required bounded terminology/identity refinement:

```text
current name     Conexus OS
historical name  MNFS
```

Preserve the semantic invariant:

> The software-development Harness may plan/build/test/review/package Evidence for Aurora, but it does not define Aurora's sovereign Product architecture or become a required Aurora runtime dependency.

Update only current routed authority/README/decision references. Historical Evidence may retain MNFS as provenance.

## 11.3 Old TA-03+ program decisions

Preserve TA-01/TA-02. Reopen old TA-03+ ordering only to the extent necessary to adopt the revised dependency graph above. No technology winner changes by renumbering.

---

# 12. Target repository operating model

## 12.1 Fresh actor route

```text
AGENTS.md
→ docs/index.md
→ docs/roadmap.md
→ 1–2 task-specific owning documents
```

Normal task pack <= 5 files. Research/Evidence/history/Git are opt-in only for a named question.

## 12.2 Root files

### `README.md`

Landing only: concise Product description/North Star, stable links to `AGENTS.md` and `docs/index.md`, stable verification entrypoint if useful. No mutable stage/status.

### `AGENTS.md`

Compact bootstrap only:

- fresh actor route;
- citations to Method/Repository Standard;
- Aurora-local hard stops;
- verification command;
- local Git/review rules requiring emphasis.

No duplicated full Product/method/roadmap.

### `CONTRIBUTING.md`

Keep only if a real human-contributor consumer remains. If kept, it is compact guidance pointing to `AGENTS`, `docs/index` and repository-local engineering rules. It owns no status or architecture.

### `docs/index.md`

Canonical task/intention router. It names the smallest starting owner and what not to read by default. It owns no mutable status.

### `docs/roadmap.md`

Sole mutable current stage/status/allowed-work/next-action authority.

---

# 13. Target documentation model

Create directories only when a real current/future named consumer exists.

```text
docs/
├── index.md
├── roadmap.md
├── product/
│   ├── README.md
│   ├── blueprint/
│   ├── PRODUCT-BLUEPRINT.md
│   └── REQUIREMENTS-TRACEABILITY.md
├── architecture/
│   ├── index.md
│   └── current structural authorities
├── decisions/
│   ├── index.md
│   └── adr/
├── phases/
│   └── compact durable stage/closure authorities when useful
├── development/
│   ├── engineering-rules.md
│   ├── capability-realization.md
│   └── documentation.md when Aurora-specific rules remain
├── capabilities/
│   └── current capability owners such as CAP-SOVEREIGN-CORE
├── research/
├── evidence/
├── reference/
├── diagrams/        # only if a routed owner consumes them
└── work/            # branch/review temporary only; never final merge/main
```

This is a semantic target, not an empty-directory checklist.

---

# 14. Current-tree reconciliation

| Current surface | Candidate disposition |
|---|---|
| `docs/tracking/STATUS.md` | **SUPERSEDE** after semantics move to `docs/roadmap.md` |
| `docs/tracking/DECISIONS.md` | **REHOME/REFINE** to `docs/decisions/index.md` |
| `docs/tracking/WORKLOG.md` | consolidate surviving obligations, then remove live chronology; Git/PR history is archive |
| `docs/tracking/BACKLOG.md` | move material deferred obligations to decision/forward-obligation owner; avoid second roadmap |
| `docs/tracking/DOCUMENTATION-COVERAGE.md` | retain/rehome only if a named proof consumer remains |
| `docs/DOCUMENTATION-MAP.md` | split responsibilities among external Standard, `docs/index`, `docs/roadmap`, and local documentation rules |
| `docs/superpowers/**` | absorb unique current semantics, then remove permanently from live merge candidate/main |
| `docs/design/**` | rehome by current meaning to architecture/phases/evidence; do not bulk-rename mechanically |
| `docs/acceptance/**` | consolidate durable acceptance/closeout facts; preserve only Evidence with a real consumer |
| `docs/reviews/**` | preserve deciding Evidence where useful; round chronology otherwise remains in Git/PR history |
| `docs/history/**` | origin/discovery becomes opt-in reference if still useful |
| `docs/adr/**` | rehome under `docs/decisions/adr/**` after decision routing migration |
| `docs/capabilities/CAP-SOVEREIGN-CORE/**` | **PRESERVE initially** due frozen M0 Evidence/requalification consumers |

Deletion is permitted only after every still-current semantic obligation and required byte-level provenance has a durable reachable home/ref.

---

# 15. Repository verification target

The eventual aggregate gate should prove at least:

## Bootstrap / authority

```text
AGENTS.md + docs/index.md + docs/roadmap.md <= 20 KiB
docs/roadmap.md is sole mutable program status owner
README.md is landing only
default task route <= 5 files
durable authorities are routed
router links resolve
```

## Temporary/bloat

Final merge candidate/main contains no:

```text
docs/work/**
docs/superpowers/**
AI review-dialog artifact
parallel mutable status/roadmap tree
permanent session handoff/dialogue tree
active old/archive tree
```

## Decision/provenance

- current decision dispositions are valid;
- no accepted semantic owner disappears during path migration;
- deleted live historical files retain required Git/ref provenance;
- no current authority links to removed paths without replacement.

## Guard quality

- material guards have a deterministic negative control or equivalent falsifier;
- base→candidate diff checks use the intended range;
- architecture-only/implementation-blocked mode is fail-closed against unauthorized Product source/runtime additions.

## Aggregate gate

Target for the realigned repository:

```text
required
```

Current `Documentation/validate` remains transition Evidence until the replacement guard is proved and branch protection is migrated.

---

# 16. Git / PR / independent review target

```text
main
→ one branch + Draft PR for one coherent gate
→ bounded analysis / docs/work when useful
→ consolidated candidate
→ exact verification
→ isolated independent Fable challenge when Method requires it
→ Lead adjudication
→ bounded correction
→ operator ratification
→ separate explicit merge authorization
→ squash merge
→ head-branch cleanup
→ next stage starts from updated main
```

Required repository settings after gate migration:

- PR-based changes to `main`;
- no force-push/delete of `main`;
- aggregate required status check;
- no later stage stacked on unmerged earlier stage by default.

Current GitHub evidence says Aurora `main` is presently unprotected; that is a repository-conformance gap to close during migration, not permission to bypass PR discipline now.

---

# 17. Repository rebaseline migration sequence after ratification

The migration is a separate documentary/repository gate, not TA-03 technical design and not Product implementation.

```text
RM-01  create compact bootstrap: README / AGENTS / docs/index / hand-owned docs/roadmap
RM-02  refine Blueprint 15 + ACRM scope + local engineering/documentation owners
RM-03  promote revised planning/readiness graph to durable authority
RM-04  rehome current decision register and ADR routing
RM-05  reconcile architecture/design/capability authorities
RM-06  consolidate acceptance/review/tracking/history with provenance proof
RM-07  remove docs/superpowers/** and superseded parallel status/worklog surfaces
RM-08  change generator so only true projections are generated
RM-09  replace validator/CI with Repository Standard conformance + negative controls
RM-10  configure/prove branch protection and aggregate `required` gate
RM-11  fresh-actor routing proof + Global Coherence Review
RM-12  independent challenge + operator ratification + explicit merge authorization
```

These are migration work units, not mandatory separate PRs. Their branch/PR partition is decided by coherence/materiality and repository law.

No Product source/runtime implementation is included.

---

# 18. Decision reconciliation summary

| Current authority/item | Candidate disposition |
|---|---|
| A0 Product constitution | **PRESERVE** |
| Blueprint 01–14 Product meaning | **PRESERVE**, with bounded current-name/status reference corrections only |
| Blueprint 15 principles | **PRESERVE** |
| Blueprint 15 repository-operating clauses | **BOUNDED REOPEN / REFINE** |
| ACRM R0–R8 core semantics | **PRESERVE / REFINE SCOPE** |
| ADR-0001/0002 | **PRESERVE** |
| ADR-0003..0008 | **PRESERVE WITH M0 SCOPE** |
| ADR-0009 | **PRESERVE EXACT SCOPE / REQUALIFY AT CONSUMER** |
| frozen M0 R7 | **PRESERVE AS NON-CANONICAL EVIDENCE** |
| System Architecture Rebaseline | **PRESERVE RESULT / UPDATE ROUTING** |
| TA-01 / TA-02 | **PRESERVE / CLOSED / CANONICAL** |
| old TA-03 early repository stage | **REOPEN / SUPERSEDE BY NEW TA-09** |
| old TA-04 contracts | **REFINE**, follows new TA-03 operation surface |
| old TA-05 data | **REFINE → new TA-06** |
| old TA-06 identity/security | **REFINE → new TA-07** |
| old TA-07 cognition/Harness | **REFINE → new TA-08** |
| old TA-08 operations | **SPLIT → new TA-09 + TA-10** |
| Presence activation/locked-workstation constraints | **PRESERVE DOWNSTREAM** |
| frontend method v2.1 | **REFERENCE / DEFER CURRENT ADOPTION** |
| MNFS current name | **SUPERSEDED BY CONEXUS OS; preserve historical provenance** |

---

# 19. Explicit non-decisions

This rebaseline does **not** select:

- monorepo/polyrepo;
- a universal language;
- concrete Mastra version/integration;
- HTTP/gRPC/Connect/MCP/A2A/event/IPC binding;
- schema/codegen technology;
- PostgreSQL/vector/graph/object store architecture;
- authentication/policy/secrets products;
- supervisor/container/orchestrator;
- model/Voice/sandbox/durable-engine/observability backend;
- first AHDK language;
- Product implementation graph;
- production deployment.

Those become decisions only in their revised owning stages with current Evidence.

---

# 20. Acceptance meaning

Recommended operator decision:

```text
ACCEPT CANDIDATE v0.2
```

Acceptance would authorize only:

1. write the durable methodology/repository rebaseline authority;
2. boundedly revise Blueprint 15 and current terminology/decision routing;
3. write the exact repository migration design/plan;
4. prepare the migration gate for independent review.

Acceptance would **not** authorize:

- merging the migration to `main`;
- opening TA-03;
- Aurora Product implementation;
- M0 R7/R8 continuation;
- Architecture Spike execution;
- framework/store/IAM selection.

A separate operator ratification is required for the durable package, and a separate explicit merge authorization remains required afterward.