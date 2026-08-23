---
id: work-current-methodology-repository-rebaseline-proposal
kind: temporary-work
owner: architecture
status: candidate
summary: Candidate complete methodology, implementation-readiness and repository-operating rebaseline for Projeto Aurora.
---

# Aurora Methodology, Planning-Readiness & Repository Rebaseline — Candidate

> **NON-AUTHORITATIVE / BRANCH-ONLY.** This proposal is the current Lead candidate for operator review. It must be absorbed into durable owners or deleted before merge.

## 1. Decision question

> How should Projeto Aurora adopt the engineering/repository discipline that matured across MetalDocs, Marketplace Central and Conexus OS so future implementation becomes constrained realization of accepted Product/architecture/contracts rather than a second system-design phase performed while coding?

The rebaseline has two equally binding concerns:

```text
A. ENGINEERING / PLANNING METHOD
   what must be decided, in what dependency order, with what evidence/proof,
   before Product implementation can begin

B. REPOSITORY OPERATING MODEL
   where authority/status/decisions/research/evidence live,
   how a fresh actor navigates them,
   how branches/reviews/CI/temporary work behave,
   and how context survives sessions without repository archaeology
```

This gate changes **method and routing**, not Aurora Product meaning by convenience.

---

## 2. Revalidated starting point

```text
repository: developmentconexus-ops/aurora_project
main: 35614c581cea32e04305c1ad63522fee151eb283
current Product implementation: PAUSED / BLOCKED
M0 R7 candidate: FROZEN / PRESERVED / NON-CANONICAL
TA-01: ACCEPTED / MERGED / CANONICAL
TA-02: ACCEPTED / MERGED / CANONICAL
TA-03 old definition: NOT AUTHORIZED
```

Accepted TA-01/TA-02 remains the current technical ownership/runtime baseline:

```text
G01 Contract Model Governance
C01–C12 canonical domain owners
A01–A05 application/runtime coordination
B01 transport-neutral provider-runtime boundary
E01 Effect Gateway
E02 Credential Broker

Stage A:
one small persistent Evolutionary Sovereign Host
+ one on-demand provider-runtime seam at first consumer
+ specialized Harnesses as independent providers
```

No evidence found in this rebaseline so far requires reopening that semantic ownership/topology direction.

---

## 3. New organizational authorities discovered after Aurora's last baseline

Aurora's current operating/documentation model predates two cross-repository authorities now used by MetalDocs, Marketplace Central and Conexus OS:

### 3.1 DevelopmentConexus Engineering Method v1.0.0

Canonical external owner:

```text
developmentconexus-ops/conexus-methodology/METHOD.md
```

It governs **how engineering decisions are reasoned about**:

```text
Evidence
→ Known / Inferred / Unknown / Deferred
→ Root Cause
→ Target Invariant
→ Constraints
→ Credible Alternatives
→ Local Maximum vs Global Maximum
→ Essential vs Accidental Complexity
→ YAGNI / Overengineering / Future Cost
→ Authority / Boundary
→ Enforcement
→ Proof Strategy
→ Adversarial Challenge
→ Decision
→ Reopen Triggers
```

Key laws consumed by Aurora:

- mechanism != authority;
- evidence strength is claim-relative;
- unknown remains unknown;
- current implementation is Evidence, not target authority merely because it exists;
- prepare the seam, not the entire future capability;
- proof strategy exists before implementation;
- material decisions receive adversarial challenge;
- Global Maximum does not mean maximum infrastructure/generalization;
- reopen accepted decisions only on material Evidence, not preference.

### 3.2 DevelopmentConexus Repository Standard v1.0.0

Canonical external owner:

```text
developmentconexus-ops/conexus-methodology/REPOSITORY-STANDARD.md
```

It governs **how repository knowledge and delivery are organized**.

Required fresh-actor route:

```text
AGENTS.md
→ docs/index.md
→ docs/roadmap.md
→ 1–2 task-specific owning documents
```

Default task context is at most five files unless a named material reason requires more.

Required repository roles:

```text
README.md       = landing only
AGENTS.md       = compact bootstrap/router + local hard stops
docs/index.md   = task/intention router
docs/roadmap.md = sole mutable current stage/status/allowed-work/next-action authority
```

Current Aurora diverges materially:

```text
no docs/index.md
mutable program status owned by docs/tracking/STATUS.md
docs/roadmap.md is a generated Product capability-roadmap projection
permanent docs/superpowers/** lives on main
large parallel tracking/status/worklog surfaces remain live
Documentation Map owns repository routing that is now partly organizational standard
main is currently not protected by GitHub branch protection
```

Those are methodology/repository drift findings, not Product defects.

---

# 4. Lessons imported as Evidence, not authority

## 4.1 MetalDocs

MetalDocs demonstrates that architecture can be closed down to executable details before Product implementation:

```text
Product / ownership
→ technical baseline
→ backend topology
→ internal owner contracts
→ persistence
→ executable wire
→ frontend realization
→ runtime/process/deployment
→ whole-system coherence
→ golden flows / validation
→ transition/cutover
→ implementation graph
→ adversarial implementation-readiness
→ explicit execution authorization
```

Important transferable properties:

- existing implementation is Evidence, not target authority;
- reuse is proof-backed and consumer-driven;
- one executable wire authority exists before frontend/backend coding;
- frontend is bidirectionally traced to Product operations;
- runtime mechanisms are selected only for named consumers;
- implementation remains blocked even after major architecture stages close.

## 4.2 Marketplace Central

Marketplace Central strengthens:

- Product/domain/identity/communication/integration/API separation before runtime mechanics;
- one machine-readable Product wire authority;
- decision-generation reconciliation so retired ideas do not return through implementation convenience;
- architecture Golden Flows selected by defect class rather than exhaustive operation count;
- controlled real external probes only where documentary evidence cannot close a material claim;
- frontend planning that cannot invent business/API authority.

## 4.3 Conexus OS

Conexus OS provides the strongest current implementation-readiness model:

```text
accepted Product / architecture authority
→ Product Surface & Authority Contract
→ Executable Wire Contract
→ Frontend Interaction & Authority Realization
→ Paved Road + Runtime/Persistence Realization
→ Whole-System Coherence & Golden Flows
→ Implementation Program & Execution Graph
→ Adversarial Implementation Readiness
→ explicit Product execution grant
→ implementation
```

It also formalizes the desired coding boundary:

```text
GENERATED
PLATFORM-CONTRACT
APP-OWNED
```

The purpose is to precompile repeated foundational decisions so coding agents legitimately choose mainly APP-OWNED local logic instead of repeatedly inventing auth, transport, persistence, contracts, security, repository structure and verification conventions.

## 4.4 Frontend Product Experience Planning Method v2.1

Current reusable reference lives in Marketplace Central:

```text
developmentconexus-ops/marketplace-central/
docs/development/frontend-product-experience-planning-method.md
```

It establishes P0–P14 from authority recovery and user needs through rendered structural wireframes, Screen Contracts, interactive low-fidelity prototype, adversarial walkthrough and implementation-readiness closure.

Disposition for Aurora now:

```text
REFERENCE / FUTURE SPECIALIZED METHOD
not copied locally
not current frontend work
not Product authority
```

Before Aurora's first material visual frontend/Presence surface enters implementation readiness, use the then-current organizationally accepted frontend method. If v2.1 has not yet been promoted into `conexus-methodology`, that cross-repository promotion should be handled separately rather than creating an Aurora fork.

---

# 5. Approaches considered

## Approach A — Minimal local patch

```text
keep current Aurora documentation tree
+ cite METHOD.md
+ cite REPOSITORY-STANDARD.md
+ continue old TA-03→TA-08 sequence
```

### Benefit

Smallest immediate change.

### Failure

Leaves known structural drift:

- parallel mutable status owner;
- no canonical task router;
- permanent temporary-plan tree;
- repository/source layout selected before operation/wire/consumer closure;
- no explicit implementation-readiness stage comparable to current best practice;
- coding can still inherit material unknowns.

**Disposition:** REJECT.

## Approach B — Repository cleanup only

```text
align AGENTS/index/roadmap/docs tree
+ preserve old TA-03→TA-08 sequence
```

### Benefit

Fixes session/repository behavior without changing technical architecture program.

### Failure

Does not address the deeper planning-order defect: old TA-03 asks for repositories/packages/build before exact control surface, executable contracts and consumer realization are closed.

**Disposition:** REJECT.

## Approach C — Full methodology + repository + implementation-readiness rebaseline

```text
adopt organizational Method + Repository Standard
+ preserve Aurora Product/TA-01/TA-02 authority
+ reframe ACRM as capability/slice realization lifecycle
+ replace TA-03+ planning order
+ introduce operation surface, executable contracts, Paved Road,
  whole-system proof, implementation graph and final readiness gate
+ migrate repository routing/status/tree to Repository Standard
```

### Benefit

Matches the real target:

> downstream coding receives exact authority, contracts, allowed degrees of freedom and proof obligations.

### Cost

Requires a deliberate documentary migration before new technical-stage work resumes.

**Recommendation:** ADOPT Approach C.

---

# 6. Method hierarchy after rebaseline

The rebaseline prevents several methods from competing for the same authority.

```text
┌──────────────────────────────────────────────────────────┐
│ DevelopmentConexus Engineering Method v1               │
│ HOW ENGINEERING IS REASONED ABOUT                       │
└──────────────────────────┬───────────────────────────────┘
                           ↓
┌──────────────────────────────────────────────────────────┐
│ DevelopmentConexus Repository Standard v1              │
│ HOW REPOSITORY AUTHORITY/CONTEXT/DELIVERY IS OPERATED   │
└──────────────────────────┬───────────────────────────────┘
                           ↓
┌──────────────────────────────────────────────────────────┐
│ Aurora Product + Architecture Authorities               │
│ WHAT AURORA MEANS / OWNS / MUST PRESERVE                │
└──────────────────────────┬───────────────────────────────┘
                           ↓
┌──────────────────────────────────────────────────────────┐
│ Aurora Planning & Implementation-Readiness Program      │
│ WHICH SYSTEM-WIDE DECISIONS MUST CLOSE BEFORE CODING    │
└──────────────────────────┬───────────────────────────────┘
                           ↓
┌──────────────────────────────────────────────────────────┐
│ Specialized methods                                     │
│ frontend / research / qualification / transition        │
│ activated only by a real consumer                       │
└──────────────────────────┬───────────────────────────────┘
                           ↓
┌──────────────────────────────────────────────────────────┐
│ ACRM R0–R8                                               │
│ CAPABILITY / SLICE REALIZATION + EVIDENCE LIFECYCLE     │
└──────────────────────────┬───────────────────────────────┘
                           ↓
┌──────────────────────────────────────────────────────────┐
│ Implementation Graph + constrained coding agents        │
└──────────────────────────────────────────────────────────┘
```

## 6.1 ACRM disposition

`docs/product/CAPABILITY-REALIZATION-METHOD.md` is not rejected.

Proposed disposition:

```text
PRESERVE core R0–R8 semantics
REFINE scope/placement
```

ACRM remains responsible for:

- applicability;
- capability requirements;
- capability readiness;
- capability-specific technical decisions/spikes;
- scoped contracts;
- microdesign;
- execution/evidence;
- milestone closeout.

It no longer needs to own repository routing or act as the only system-wide planning sequence.

System-wide readiness authority is provided by the planning program below. A capability R0–R6 consumes current global architecture; it cannot silently override an upstream global stage.

---

# 7. Target engineering invariant

> **Implementation must not be the stage where material Product, architecture, contract, security, data, runtime or UX decisions are discovered by convenience.**

Portuguese operational statement:

> Quando a implementação começar, toda decisão material previsível deve possuir autoridade explícita; o implementador recebe somente os graus de liberdade locais e reversíveis necessários para realizar os contratos aceitos.

Required stop law:

```text
material missing decision
→ STOP
→ identify smallest owning authority/stage
→ reopen only that owner
→ recompile affected downstream artifacts
→ resume only from reconciled authority
```

Forbidden behavior:

```text
missing decision
→ choose plausible default in code
→ let implementation convenience become architecture
```

---

# 8. Proposed Aurora planning / implementation-readiness graph

TA-01 and TA-02 are preserved. The accepted old ordering of TA-03→TA-08 is proposed for bounded reopen/supersession because newer cross-project Evidence shows repository/package topology was placed too early.

## MR-01 — Methodology, Planning-Readiness & Repository Rebaseline

**Current gate.**

Owns:

- method hierarchy;
- Repository Standard alignment target;
- reconciliation of old Aurora planning stages;
- new implementation-readiness graph;
- repository migration plan;
- implementation stop boundary.

Exit:

- operator-ratified rebaseline;
- repository migration target accepted;
- durable stage graph accepted;
- no Product implementation authorized.

---

## TA-01 — Logical Modules & Canonical Ownership

```text
STATUS: PRESERVE / CLOSED / CANONICAL
```

Owns:

- G01;
- C01–C12;
- A01–A05;
- E01/E02 responsibility boundaries;
- canonical ownership and forbidden writes.

Reopen only on material owner contradiction.

---

## TA-02 — Process, Runtime & Evolutionary Topology

```text
STATUS: PRESERVE / CLOSED / CANONICAL
```

Owns:

- Evolutionary Sovereign Host;
- initial provider-runtime seam;
- Stage A→B evolution principles;
- process split triggers;
- A05/B01 lifecycle boundary.

Concrete supervisors/transports/packages remain downstream.

---

## TA-03 — Control & Capability Operation Surface

### Purpose

Define the complete set of **material Aurora operations/interactions** needed by the current architecture horizon before choosing APIs, repositories or frameworks.

This is not only a public HTTP API census. It covers semantic operations crossing or protecting meaningful boundaries.

Candidate operation families to derive, not pre-admit:

```text
interaction / Presence
identity / Project state
Intent / Mission / Delegation
Authority / Policy / Effect request
Capability / Provider registry
Memory / Context
Artifact / Evidence / Verdict
Audit / exact history
runtime lifecycle
provider dispatch / reconciliation
```

Each admitted operation records where applicable:

```text
operation identity
caller/actor class
semantic owner
consumer
input/output meaning
read vs command vs proposal vs effect
scope / authority condition
state transition or non-transition
knowledge/outcome class
semantic idempotency/concurrency need
failure/ambiguity meaning
audit/evidence obligation
what it explicitly does NOT own
```

### Exit law

```text
material current consumer without operation = 0
operation without current/future-named consumer = 0 unless explicitly DEFERRED
ownerless operation = 0
mechanism/protocol operation presented as Product meaning = 0
```

---

## TA-04 — Executable Contract Model & Wire

### Purpose

Compile TA-03 semantics into Aurora-owned machine-checkable contract families before implementation frameworks create DTO/protocol authority.

Owns:

- G01 contract-family catalog;
- semantic version and compatibility law;
- schemas/types for boundary objects;
- request/result/event/proposal/effect envelope meanings;
- error/problem taxonomy;
- deadlines/cancellation;
- idempotency/correlation/causation;
- restart/reconciliation contract identities;
- streaming semantics when required;
- generated projection law;
- schema/binding/SDK/provider version separation.

Concrete bindings are selected per boundary only after requirements exist:

```text
in-process
local IPC
HTTP/OpenAPI
Connect/gRPC
SSE/WebSocket
message/event transport
MCP/A2A/other external bindings
```

There is no universal-protocol requirement.

### Exit law

Every material TA-02 process/provider crossing and every TA-03 externally consumed operation has an executable contract owner, compatibility/failure semantics and a deliberate binding disposition.

---

## TA-05 — Human Interaction & Presence Realization

### Purpose

Prove that accepted Product/control operations compose into coherent Leandro-facing interaction before infrastructure/package layout freezes around an incomplete user journey.

Aurora-specific interaction includes more than pages:

```text
activation
text interaction
voice-capable interaction when in current horizon
session/context continuation
clarification / decision request
Mission/progress inspection
Artifact/Evidence inspection
failure/recovery interaction
Presence handoff when current
```

Outputs:

- human goals and end-to-end flows;
- operation-to-interaction coverage;
- interaction/state authority map;
- material failure-message intent;
- initial Presence surface/shell semantics;
- exact UI/Presence prohibition against inventing authority;
- frontend visual-method trigger/disposition.

### Frontend specialized method

When the current Stage A implementation includes a material visual desktop/web surface, activate the then-current DevelopmentConexus frontend planning method (current reference: Frontend Product Experience Planning Method v2.1) proportionally.

The frontend method may produce IA, rendered structural wireframes, Screen Contracts and a low-fidelity functional prototype, but it cannot invent TA-03 operations or TA-04 wire.

### Exit law

Every current human goal is either representable through admitted operations/contracts or creates a named upstream finding. No visual/screen convenience creates Product authority.

---

## TA-06 — State, Data, Memory, Knowledge, Artifact & Evidence Architecture

### Purpose

Classify every durable/derived/ephemeral data family and its lifecycle before selecting physical stores.

Required families:

```text
canonical operational state
governed memory
knowledge/source documents
interaction/exact history
artifacts
observations/measurements
evidence/receipts/verdicts
audit
telemetry
derived retrieval indexes
ephemeral active context
provider-local state
secret/credential references
```

For each:

```text
owner/writer
canonical vs derived vs ephemeral
consistency/freshness
query/access pattern
confidentiality/integrity
retention/deletion
rebuildability
backup/export/restore
migration ownership
expected scale where deciding
co-location/separation criteria
```

Physical mechanisms may then be evaluated by real role:

```text
SQLite
PostgreSQL
filesystem/object store
full-text/vector/graph mechanisms
telemetry store
other proven mechanism
```

### Exit law

No store is selected because a framework offers it; each store has a named role, consumer, protected property and migration/replacement boundary.

---

## TA-07 — Identity, Authentication, Authorization, Policy, Secrets & Effect Enforcement

### Purpose

Close actor identity proof and deterministic authority/effect paths before coding.

Owns:

- actor/identity classes;
- stable identity vs runtime incarnation;
- human/session authentication requirements;
- service/Presence/Harness/provider/device authentication by class;
- Authority Grant/delegation semantics;
- authorization/policy-decision contract;
- Effect Gateway enforcement;
- secret classes/references/delivery;
- credential broker boundary;
- step-up/revocation/recovery/audit;
- Stage A→B changes.

Technology research starts from property, not product name. Candidate classes may include:

```text
OIDC/OAuth
Keycloak / Zitadel / Authentik / Ory
SPIFFE/SPIRE
Cedar / OPA
OS credential stores / Vault-like brokers
minimal Aurora-owned mechanisms where smaller and sufficient
```

No IAM/policy product may become Aurora domain identity or authority by convenience.

### Exit law

Every material actor/effect has an identity-proof, authority, credential and enforcement route or an explicit safely deferred consumer.

---

## TA-08 — Cognitive Runtime, Models, Context, Memory Use & Harness Realization

### Purpose

Close the deterministic Core ↔ cognitive/provider execution boundary with current consumer evidence.

Owns:

- deterministic vs model-mediated responsibility matrix;
- Context Builder inputs/authority;
- model-provider role taxonomy;
- model routing/fallback semantics;
- governed memory read/write/promotion path;
- tool/capability request path;
- Harness Delegation lifecycle;
- provider-local/global state reconciliation;
- Mastra fit for exact first consumers;
- restart/cancellation/ambiguous completion;
- artifact/evidence return path.

Mastra remains preferred-first to evaluate, not preselected as sovereign owner.

### Exit law

Replacing/loss of model, Mastra runtime or Harness cannot redefine Aurora identity, canonical state, Authority or governed memory semantics.

---

## TA-09 — Paved Road, Repository, Source, Build & Dependency Realization

### Purpose

Only after consumers/contracts/data/security/cognition are exact, define the environment future coding agents receive.

This stage supersedes the old early TA-03 repository stage.

### Ownership classes

Adopt the implementation-readiness distinction where useful:

```text
GENERATED
→ reproducible from accepted contract/source; never hand-owned divergent semantics

AURORA-CONTRACT
→ protected Aurora-controlled seam/mechanism consumed by implementation;
  implementer cannot silently weaken its invariant

APP-OWNED / MODULE-OWNED
→ module-specific logic future coding agents legitimately evolve
```

Exact naming may be refined, but the ownership distinction is required.

### Outputs

- monorepo/polyrepo/staged strategy;
- source/package/module layout;
- Go/TypeScript/other language placement;
- contract/schema source location;
- generated projection/codegen law;
- dependency directions and static guards;
- build/workspace tooling;
- dependency/version admission policy;
- test layers;
- CI aggregate gate;
- release/versioning relationships;
- scaffold/profile/Paved Road;
- Conexus OS / software-development Harness relationship;
- escape-hatch rule for a real unsupported property.

### Technology disposition

Every material dependency/framework decision follows:

```text
protected property
→ standards/current official evidence
→ credible alternatives
→ ADOPT | ADAPT | BUILD | DEFER | STOP
→ exact boundary
→ falsifier/proof
```

### Exit law

A future coding agent can place code/contracts/tests/dependencies without inventing foundational structure, and can clearly distinguish generated/protected/module-owned surfaces.

---

## TA-10 — Runtime, Configuration, Deployment, Observability & Recovery Realization

### Purpose

Select concrete operational mechanisms that realize TA-02 topology and TA-06/07/08 requirements.

Owns:

- runtime shells/process artifacts;
- service supervision;
- configuration source/precedence;
- secrets injection/reference realization;
- startup/readiness/liveness/shutdown;
- migrations;
- packaging/install/update/rollback;
- Windows/Linux Stage A posture;
- Stage B deployment posture;
- logs/traces/metrics vs Audit/Evidence boundaries;
- OTel exporter/backend realization where current;
- backup/restore/drills;
- supply chain and artifact provenance;
- operational resource/latency budgets when deciding.

### Exit law

The architecture can explain how a clean supported machine becomes a recoverable Aurora installation and how failure is diagnosed without telemetry becoming Product truth.

---

## TA-11 — Whole-System Coherence & Golden Flows

### Purpose

Prove that accepted TA-01→TA-10 authorities form one coherent system before implementation work is decomposed.

Golden Flows are selected by **defect class**, not exhaustive operation count.

Candidate proof families are derived later, but must represent at least the current highest-risk cross-boundary classes such as:

```text
restart-preserved sovereign state
human interaction → Intent → Mission / no-Mission disposition
Context Builder → cognition → proposal → canonical commit path
provider/Harness dispatch → ambiguous failure → reconciliation
memory promotion/supersession/deletion
forbidden effect denial / allowed effect receipt
Stage A process crash/restart
```

Real external/runtime probes execute only under separate authorization when documentary proof cannot establish a material claim.

### Exit law

Every material cross-system invariant has at least one representative positive/negative/recovery falsifier, and no unresolved global coherence contradiction remains.

---

## TA-12 — Implementation Program & Execution Graph

### Purpose

Compile accepted architecture into bounded implementation slices.

Each slice receives at minimum:

```text
operator-visible outcome
exact Product/control operations
owners/modules
contracts/generated projections
allowed data/persistence surfaces
allowed source/packages
exact dependency pins/profiles when material
GENERATED/AURORA-CONTRACT/MODULE-OWNED mutation permissions
configuration/migration/process surfaces allowed to change
proof/falsifier obligations
prerequisites
stop/reopen triggers
completion criteria
```

No slice chooses its own foundational stack.

### Exit law

The implementation graph is dependency-ordered, each slice is independently falsifiable/reviewable and all material architecture decisions are owned upstream.

---

## TA-13 — Adversarial Implementation Readiness

### Purpose

Fresh independent attack before Product execution becomes eligible.

Challenge at least:

```text
consumer without operation
operation without owner
owner duplication
missing executable contract
parallel DTO/wire authority
screen/Presence-shaped API
provider/framework semantic leakage
persistence class without owner/invariant
auth/effect path without deterministic enforcement
dependency without protected property
repository/package ambiguity
unowned generated code
Paved Road silent bypass
missing failure/recovery semantics
Golden Flow without real falsifiability
implementation slice containing architecture decisions
```

### Exit law

```text
material readiness findings = 0
independent challenge converged
operator ratification complete
```

Even then, Product implementation remains blocked until a separate explicit execution grant.

---

# 9. Relationship to Product Milestones

The existing long-horizon Product capability roadmap remains valuable and should not become mutable repository status.

Proposed separation:

```text
docs/product/blueprint/14-capability-roadmap.md
= canonical long-horizon Product Milestone sequence / Golden Proof direction

docs/roadmap.md
= sole mutable current repository/program stage/status/allowed-work/next-action authority
```

The current generated `docs/roadmap.md` projection therefore has a naming/authority collision and must be reworked during repository migration.

No Product Milestone is silently accepted/rejected by this rebaseline.

---

# 10. Target repository operating model

## 10.1 Fresh actor route

```text
AGENTS.md
→ docs/index.md
→ docs/roadmap.md
→ 1–2 task-specific owning documents
```

Default pack <= 5 files. Research/Evidence/history/Git are opt-in only when the task has a concrete reason.

## 10.2 Root responsibilities

### `README.md`

Landing only:

- concise Aurora description/North Star;
- links to `AGENTS.md` and `docs/index.md`;
- stable verification entrypoint if useful.

It must not contain mutable current stage/status/next action.

### `AGENTS.md`

Bootstrap only:

- fresh-actor route;
- reference to external Method and Repository Standard;
- Aurora-local hard stops;
- local verification command;
- Git/review rules that differ materially or need emphasis.

It must not duplicate Product prose, full ACRM, roadmap or full methodology.

### `docs/index.md`

Canonical task/intention router:

- task/question;
- smallest starting owner;
- optional add-on owner;
- what not to read by default.

It owns no mutable program status.

### `docs/roadmap.md`

Only mutable current program authority:

- current rebaseline/stage;
- closed stages relevant to progression;
- exact next action;
- implementation allowed/blocked;
- entry/exit/reopen conditions where material.

No other document may act as a second mutable status page.

---

## 10.3 Target documentation tree

Directories exist only with a real consumer. Target semantic model:

```text
docs/
├── index.md
├── roadmap.md
│
├── product/
│   ├── README.md
│   ├── blueprint/
│   ├── PRODUCT-BLUEPRINT.md        # generated aggregate if retained
│   └── REQUIREMENTS-TRACEABILITY.md
│
├── architecture/
│   ├── index.md
│   ├── ownership-and-runtime.md    # accepted TA-01/TA-02 authority or routed successor
│   ├── ... future TA authorities
│   └── ... current technical references only when they own current meaning
│
├── decisions/
│   ├── index.md                    # current disposition/reopen register
│   └── adr/
│       └── ... accepted/proposed ADRs with current disposition
│
├── phases/
│   ├── a0-product-baseline.md      # compact durable closure when useful
│   ├── m0-sovereign-core.md        # current/frozen M0 phase state and preserved Evidence route
│   ├── methodology-rebaseline.md   # MR-01 durable result after ratification
│   └── ... later stage contracts/closure summaries
│
├── development/
│   ├── engineering-rules.md        # Aurora-local specialization only
│   ├── capability-realization.md   # ACRM, after scope/refinement
│   └── documentation.md            # Aurora-specific authority/document rules not already external
│
├── research/
│   └── ... current research with named consumers
│
├── evidence/
│   └── ... durable proof summaries only when a current/future consumer exists
│
├── reference/
│   ├── origin-and-discovery.md     # optional durable rationale reference
│   └── ... detailed current technical reference only when needed
│
├── diagrams/                       # source-first diagrams only when a current owner uses them
└── work/                           # branch/review temporary only; NEVER in merge candidate/main
```

This is a semantic target, not an instruction to create every empty directory.

---

# 11. Current Aurora tree reconciliation

## 11.1 `docs/tracking/STATUS.md`

```text
CURRENT ROLE: mutable status authority
TARGET: SUPERSEDED by docs/roadmap.md
```

Surviving current stage/status semantics move into `docs/roadmap.md`. The old file is removed from the live tree after migration; Git retains history.

## 11.2 `docs/tracking/DECISIONS.md`

```text
TARGET: REHOME → docs/decisions/index.md
```

Refine it from a tracking pointer table into the durable current decision/disposition + reopen register where needed.

Controlled disposition vocabulary should remain small, e.g.:

```text
CURRENT
PRESERVE
REFINED
REOPEN
DEFERRED
SUPERSEDED
REJECTED
```

## 11.3 `docs/tracking/BACKLOG.md`

Loose ideas must not become hidden future commitments.

Target:

- material deferred obligations/reopen triggers → `docs/decisions/index.md` or a narrowly justified forward-obligations page;
- ordinary uncommitted ideas → GitHub issue/backlog mechanism when a real consumer exists;
- no second planning roadmap.

## 11.4 `docs/tracking/WORKLOG.md`

Current 50 KiB chronological log duplicates Git/PR history and is not required by Repository Standard.

Target:

```text
current semantic outcomes consolidated into phase/decision owners
→ remove WORKLOG from live tree
→ Git history remains archive
```

Do not delete its last unique semantic obligation until a durable owner is identified.

## 11.5 `docs/tracking/DOCUMENTATION-COVERAGE.md`

A0 coverage Evidence may remain useful but is not mutable program status.

Target:

```text
REHOME only if a named current/future proof consumer remains
otherwise remove from live tree after A0 closure summary references Git provenance
```

## 11.6 `docs/DOCUMENTATION-MAP.md`

Its current responsibilities split:

```text
cross-repo repository structure/routing law
→ external Repository Standard

task routing
→ docs/index.md

mutable current status route
→ docs/roadmap.md

Aurora-specific authority/artifact relationships
→ docs/development/documentation.md + routed Product/architecture owners
```

After migration there should be no second giant Documentation Map required by default.

## 11.7 `docs/superpowers/**`

Repository Standard v1 prohibits permanent `docs/superpowers/` in merge candidates/main.

Disposition:

- plans → historical Git provenance after their surviving outcomes are consolidated;
- specs containing current unique architecture meaning → absorb into current durable `docs/architecture/`, `docs/development/`, Product, phase or decision owner;
- remove entire live `docs/superpowers/**` tree before final standard-aligned merge.

## 11.8 `docs/design/**`

Current mixed category includes current global architecture, M0 designs, spike specifications and implementation plan.

Target disposition by meaning:

```text
current cross-system architecture → docs/architecture/
phase-specific M0 design/proof → docs/phases/ or docs/evidence/ as appropriate
spike definition/result → phase/decision/evidence owner depending lifecycle
old implementation plan → Git history after surviving constraints are consolidated
```

Do not rename everything mechanically; each file receives a consumer/authority disposition.

## 11.9 `docs/acceptance/**` and `docs/reviews/**`

Current live tree contains many dated gate/round artifacts.

Repository target:

- current durable acceptance/closeout facts → compact `docs/phases/` or decision/evidence records;
- independent review Evidence with a continuing consumer → `docs/evidence/`;
- superseded round-by-round detail with no current consumer → Git/PR history;
- reviewer output never becomes Product authority.

The migration must prove no accepted decision loses its evidence/provenance route before deleting live files.

## 11.10 `docs/history/**`

The original Aurora origin/discovery narrative is valuable rationale but not current authority.

Proposed disposition:

```text
REHOME → docs/reference/origin-and-discovery.md
```

It remains opt-in and excluded from default fresh-session context.

## 11.11 `docs/adr/**`

```text
REHOME → docs/decisions/adr/**
```

`docs/decisions/index.md` becomes the current disposition/reopen authority. Individual ADRs continue to own their exact technical decision meaning.

## 11.12 `docs/capabilities/**`

Current CAP-SOVEREIGN-CORE artifacts have real current provenance consumers because M0 R7 is frozen and may be requalified later.

Proposed disposition:

```text
PRESERVE initially
```

Do not churn them during repository cleanup unless a later phase-owner migration clearly improves routing without losing evidence.

---

# 12. Generated documentation collision

Current `scripts/generate_docs.py` generates:

```text
docs/product/PRODUCT-BLUEPRINT.md
docs/roadmap.md
```

Under the target Repository Standard this is no longer valid because `docs/roadmap.md` must own mutable current stage/status/next action, while the generated file is the Product capability-roadmap projection from Blueprint 14.

Required migration decision:

```text
Blueprint 14 remains canonical Product capability roadmap.
Generated Product Blueprint aggregate may remain.
Generated standalone docs/roadmap.md projection must stop owning that path.
```

Preferred target:

```text
docs/product/blueprint/14-capability-roadmap.md
= canonical Product roadmap source

docs/product/PRODUCT-BLUEPRINT.md
= generated aggregate containing it

docs/roadmap.md
= hand-owned sole mutable repository/program status authority
```

No second generated Product-roadmap file is required unless a real consumer later proves value.

`scripts/generate_docs.py`, `scripts/validate_docs.py` and `.github/workflows/docs.yml` must be changed during the repository migration to enforce this split.

---

# 13. Verification / repository conformance target

After repository migration the required aggregate check should prove at least:

## Bootstrap/authority

```text
AGENTS.md + docs/index.md + docs/roadmap.md <= 20 KiB
docs/roadmap.md is sole mutable current stage/status/next-action authority
README.md is landing only
default task pack <= 5 files
durable authorities reachable from docs/index.md or routed child index
relative router links resolve
```

## Temporary/bloat

Merge candidate and main contain no:

```text
docs/work/**
docs/superpowers/**
AI review-dialog artifacts
parallel status/roadmap tree
permanent session handoffs
active old/archive tree
```

## Decision/history

- current decision dispositions valid;
- removed live Evidence/history remains reachable in Git when required;
- no current authority references deleted paths without a replacement;
- exact unique unmerged provenance receives a durable ref before branch deletion when still needed.

## Guard quality

- material guards have deterministic negative control/equivalent falsifier;
- diff checks compare the intended base→candidate;
- implementation-blocked mode uses an allowlist of permitted top-level/source surfaces rather than a denylist of legacy names when appropriate.

## Aggregate gate

Target new/re-aligned check name:

```text
required
```

The exact CI implementation is a repository migration decision; current `Documentation/validate` remains Evidence until replaced.

---

# 14. Git / branch / review lifecycle target

Normal Aurora repository lifecycle after migration:

```text
main
→ one branch + one Draft PR for one coherent stage/gate
→ analysis / temporary docs/work if needed
→ consolidate candidate
→ exact verification
→ independent Fable challenge when Method requires it
   using isolated review/<stage>-fable branch
→ Lead adjudication
→ bounded corrections
→ operator ratification
→ explicit merge authorization
→ squash merge
→ delete head branch
→ next stage starts fresh from updated main
```

Rules:

- no direct commits to `main`;
- no force-push/shared-history rewrite;
- no later stage stacked on unmerged earlier stage by default;
- temporary review/work material never merges;
- reviewer output = Evidence, not authority;
- merge/ratification does not imply Product implementation;
- branch protection should require PR and the aggregate gate once repository settings are migrated.

---

# 15. Decision reconciliation of current Aurora architecture program

Proposed dispositions:

| Current item | Proposed disposition | Reason |
|---|---|---|
| A0 Product constitution | **PRESERVE** | still owns Product meaning |
| Product Blueprint 01–15 | **PRESERVE**, route refactor only | no Product contradiction found |
| ACRM R0–R8 | **REFINE** scope/placement; preserve semantics | useful capability/slice lifecycle; not repository operating model |
| ADR-0001/0002 | **PRESERVE** | Contract Model/AHDK direction remains coherent |
| ADR-0003..0008 | **PRESERVE WITH M0 SCOPE** | no globalization |
| ADR-0009 Mastra preferred-first | **PRESERVE WITH EXACT SCOPE / REQUALIFY AT CONSUMER** | current technology Evidence is time/version sensitive |
| frozen M0 R7 candidate | **PRESERVE AS EVIDENCE / NON-CANONICAL** | no Verdict/R8; useful implementation evidence only |
| System Architecture Rebaseline | **PRESERVE RESULT / ABSORB ROUTING** | core cross-system direction still valid |
| old Technical Architecture Map TA-01/02 | **PRESERVE** | accepted ownership/topology remains coherent |
| old TA-03 repository/source/build | **REOPEN / MOVE TO NEW TA-09** | newer Evidence proves repository topology should follow operations/contracts/consumers |
| old TA-04 contracts/APIs | **REFINE / becomes new TA-04 after new TA-03 operation surface** | semantic surface must precede wire |
| old TA-05 data | **REFINE / becomes new TA-06** | physical stores follow interaction/contract requirements |
| old TA-06 identity/security | **REFINE / becomes new TA-07** | property-first security realization preserved |
| old TA-07 Brain/models/memory/Harness | **REFINE / becomes new TA-08** | remains a distinct realization stage |
| old TA-08 operations | **SPLIT** between TA-09 Paved Road/build and TA-10 runtime/ops | source/build and operational mechanisms have different prerequisites |
| Presence activation/locked-workstation decisions | **PRESERVE downstream constraints** | do not expand micro-policy now |
| frontend P0–P14 | **DEFER as specialized method trigger** | no frontend planning work now; do not fork current reusable method |

This table is a candidate reconciliation record; ratification will promote exact dispositions into the durable decision register.

---

# 16. Rebaseline migration sequence after operator ratification

Repository migration itself should be one coherent documentary/operating gate, not mixed with TA-03 technical design.

Proposed sequence:

```text
M1  create docs/index.md + new docs/roadmap.md + compact AGENTS/README
M2  create docs/development local rules/method owners
M3  promote new planning/readiness graph into durable phase/architecture owners
M4  rehome decision register + ADR routing
M5  reconcile current architecture/design files to target owners
M6  consolidate acceptance/review/tracking/history; prove no unique obligation lost
M7  remove docs/superpowers/** and obsolete parallel status/worklog surfaces
M8  modify generator/validator/CI for Repository Standard conformance
M9  run Global Coherence Review + fresh-actor routing proof
M10 independent challenge + operator ratification + explicit merge authorization
```

No Product source/runtime implementation is included.

The migration plan must use Git rename/content-preserving operations where possible and prove link/ID/reference integrity. Bulk path changes do not gain semantic authority by movement.

---

# 17. Global coherence checks required before acceptance

The final rebaseline must attack:

```text
Does Method overlap ACRM?
Does Repository Standard overlap local Documentation Method?
Is there exactly one mutable program status owner?
Is Product capability roadmap confused with repository execution roadmap?
Can a fresh actor route without reading >5 files?
Does any accepted authority disappear because its dated acceptance file is removed?
Does any current decision have two owners after migration?
Does TA-03 operation-surface work duplicate Capability Specs?
Does TA-04 wire accidentally make one protocol authoritative?
Does TA-05 force visual frontend work before a real consumer?
Can TA-09 Paved Road become a speculative internal framework?
Are M0-scoped technology decisions protected from globalization?
Can implementation still encounter a material unowned choice?
Can a later capability reopen only the smallest owner instead of restarting architecture?
```

Material findings return to this proposal before promotion.

---

# 18. Exact current decision requested from operator after review

This proposal recommends:

```text
ACCEPT APPROACH C

→ adopt METHOD.md v1.0.0 as Aurora engineering-reasoning authority
→ adopt REPOSITORY-STANDARD.md v1.0.0 as Aurora repository-operating authority
→ preserve TA-01/TA-02
→ reopen/supersede old TA-03+ ordering with new TA-03→TA-13 readiness graph
→ preserve/refine ACRM as capability/slice realization lifecycle
→ align Aurora repository to README / AGENTS / docs/index / docs/roadmap model
→ make docs/roadmap sole mutable program status authority
→ remove permanent docs/superpowers and parallel tracking/worklog surfaces after safe consolidation
→ keep frontend methodology deferred until a real visual Presence consumer
→ keep Product implementation, M0 R7/R8 and Architecture Spikes blocked
```

An `ACCEPT` of this proposal authorizes writing the durable rebaseline design + exact repository migration plan only. It does **not** authorize migration merge, TA-03 work or Product implementation by implication.