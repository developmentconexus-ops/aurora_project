---
id: DESIGN-AURORA-SYSTEM-ARCHITECTURE-DECISION-LANDSCAPE
title: Aurora System Architecture Decision Landscape
document_type: architecture_decision_landscape
form: reference
authority: design
status: proposed
version: 0.1.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - proposed global Aurora architecture question and dependency map
  - proposed DECIDE RESEARCH SPIKE DEFER treatment and earliest-consumer map
related:
  - DOC-AURORA-STATUS
  - DOC-AURORA-CAPABILITY-REALIZATION-METHOD
  - DESIGN-AURORA-SYSTEM-ARCHITECTURE-REBASELINE
  - DOC-AURORA-SYSTEM-ARCHITECTURE-REBASELINE-OPERATOR-DIRECTION
  - DOC-AURORA-BLUEPRINT-03
  - DOC-AURORA-BLUEPRINT-05
  - DOC-AURORA-BLUEPRINT-06
  - DOC-AURORA-BLUEPRINT-07
  - DOC-AURORA-BLUEPRINT-08
  - DOC-AURORA-BLUEPRINT-09
  - DOC-AURORA-BLUEPRINT-10
  - DOC-AURORA-BLUEPRINT-11
  - DOC-AURORA-BLUEPRINT-12
  - DOC-AURORA-BLUEPRINT-13
  - DOC-AURORA-BLUEPRINT-14
  - DOC-AURORA-ADR-INDEX
  - DOC-AURORA-RESEARCH-MAP
  - DOC-AURORA-DECISIONS
source_revision: e7ca5ffb652fbbd68b35d4434506c58d26daf0e1
observed_r7_candidate_revision: 7ec999b093205a9d82eef2802eca60330d96e14d
last_reviewed: 2026-08-12
---

# Aurora System Architecture Decision Landscape

## 1. Purpose

Aurora already has an accepted product constitution and logical architecture. This landscape maps the remaining cross-system technical questions so later Product Milestones do not create isolated implementations with incompatible ownership, identity, contracts, security or data semantics.

The landscape answers:

```text
What is already constrained?
What has actually been decided?
What remains open?
Which questions depend on one another?
Which Product Milestone first consumes the answer?
What evidence is sufficient?
What should deliberately not be decided now?
```

It does not select a universal stack and does not authorize implementation.

## 2. Baseline interpretation

### Accepted product architecture

The following are already constitutional or accepted decisions:

- Aurora is the sovereign cognitive control plane and one persistent identity across models, processes and presences;
- Aurora owns global Project/Mission meaning, authority, budgets, governed context, provider trust and global Outcome;
- Harnesses own specialized local methodology, attempts, workers and provider-local state;
- Aurora owns language-, framework- and transport-neutral Contract Model semantics;
- protocols and runtimes are bindings rather than product ontology;
- first-party Harnesses use AHDK by policy unless waived, while conformance remains independent;
- memory, historical narrative, knowledge sources, operational state and active context are distinct;
- current operational truth is not inferred from model or Harness narrative;
- security/effect enforcement for material actions remains outside probabilistic model control;
- logical modularity precedes physical distribution;
- large data may move through governed data channels while control/authority remains Aurora-owned;
- local-first/cloud-assisted sovereignty is required;
- evidence, receipts, claims and verdicts remain distinct.

### Scoped technical decisions

Accepted ADRs currently establish:

- Go for the **M0 Sovereign Core**, not every Aurora subsystem;
- one local modular Core and explicit current state/revisions for **M0**;
- JSON Schema/JSON/JCS logical portability and age-protected export for **M0**;
- OTel/slog observability boundary for **M0**, with backend optional;
- SQLite + `modernc.org/sqlite` for **M0 operational state**;
- M0 owner-root and restore-revalidation trust semantics;
- Mastra as preferred-first substrate to evaluate for future first-party agentic Harnesses, never as global state/authority owner.

No M0-scoped mechanism is a universal platform decision by implication.

## 3. Decision treatment

```text
DECIDE
→ sufficient evidence exists and the answer is required before the next architecture commitment

RESEARCH
→ an upcoming decision is material but current evidence is insufficient or stale

SPIKE
→ documentary analysis cannot establish the required runtime/operational property

DEFER
→ the answer does not change the next evidence-supported build decision; record consumer and trigger
```

A row can preserve a structural `DECIDE` while deferring its concrete mechanism. Example: decide who owns authentication context now; defer the authentication product until a real topology and threat model exist.

## 4. Dependency clusters

### 4.1 Identity, authority and effects

```text
canonical identities
→ authentication proof
→ actor/delegation chain
→ authorization/policy decision
→ credential brokering
→ Effect Gateway enforcement
→ receipt/audit/reconciliation
```

Selecting Keycloak, SPIFFE, OPA, Cedar, Vault or another product before this chain is specified would invert product semantics and mechanism.

### 4.2 Data, storage, memory and context

```text
data categories and canonical owners
→ consistency/freshness/retention requirements
→ operational and knowledge access patterns
→ physical stores and derived indexes
→ backup/export/migration/deletion
→ Context Builder and model exposure
```

A vector index, graph or document store may serve retrieval without becoming the owner of Project truth or governed memory authority.

### 4.3 Contracts, APIs, events and Harnesses

```text
Aurora Contract Model
→ semantic versions and compatibility
→ schema representations
→ synchronous API profiles
→ event/message profiles
→ protocol bindings
→ AHDK generated/runtime APIs
→ independent conformance
```

REST, gRPC, MCP, A2A, ACP, CloudEvents and AsyncAPI remain candidates for specific boundaries, not competing definitions of Aurora.

### 4.4 Execution, durability and isolation

```text
Delegation/effect risk profile
→ execution environment and sandbox
→ process/service boundary
→ checkpoint/idempotency/reconciliation
→ durable engine applicability
→ scheduling and campaigns
```

The Development Harness sandbox decision does not automatically select Aurora's production execution plane.

### 4.5 Artifacts, evidence, telemetry and evaluation

```text
Artifact identity/content
→ Claim and Evidence relationships
→ audit/effect receipts
→ traces/logs/metrics
→ evaluation datasets/methods
→ Verdict and Outcome
```

Telemetry is diagnostic/operational evidence; it cannot silently become product state or a Verdict.

### 4.6 Presence, multimodality and physical systems

```text
Presence/device identity
→ environment/privacy classification
→ voice/vision channel activation
→ minimal handoff context
→ observation trust
→ Device Gateway and deterministic interlocks
```

Voice convenience cannot weaken confirmation, unit read-back, privacy or physical-effect boundaries.

## 5. Global landscape matrix

| # | Architecture area | Accepted constraint / existing scoped decision | Concrete open question | Earliest consumer | Treatment now | Evidence sufficient for promotion |
|---:|---|---|---|---|---|---|
| 01 | System context and external boundaries | Aurora is sovereign Core/control plane; external models, Harnesses, services, presences and devices remain bounded providers/adapters. | What are the authoritative system-of-interest boundary, trust zones and external actor/system map for Stage A and Stage B topology? | Rebaseline itself; M1/M2 | `DECIDE` | Context diagram plus ownership/trust-boundary review against Blueprints 03, 07, 11 and 12. |
| 02 | Logical modules and state ownership | Blueprint 12 names Core modules; one concept has one canonical owner; logical modularity precedes distribution. | Which module owns each global entity/projection, which dependencies are allowed, and which cross-module writes are forbidden? | Rebaseline; M1 | `DECIDE` | Module/ownership matrix, dependency rules and adversarial duplicate-owner review. |
| 03 | Contract Model and compatibility | ADR-0001: Aurora-owned semantics with replaceable bindings; ADR-0002: AHDK is not the specification. | Which contract families must be stabilized before M1/M2, and how are semantic, schema, binding, SDK and provider versions separated? | M1 context/model boundary; M2 registry/AHDK | `DECIDE` structural scope; `RESEARCH` representation | Contract inventory and compatibility model; current standards research before schema/binding ADRs. |
| 04 | API architecture | Protocols transport; contracts define. No universal API style is selected. | Which boundaries need in-process calls, local RPC, human-readable HTTP, typed RPC, streaming or asynchronous status queries, and what common error/idempotency semantics apply? | First multi-process capability, likely M1/M2 | `RESEARCH` | Boundary inventory, latency/failure/security requirements and current official protocol documentation; spike only for consequential streaming/recovery claims. |
| 05 | Events, messaging and synchronization | Domain Event, Transport Message, Telemetry Event and Audit Record are distinct; stream delivery is not durable truth. | Which events are durable, which are notifications, how are order/duplicates/versioning handled, and when is a broker justified? | M3 event ingestion; M4 durability | `DEFER` mechanism; `DECIDE` event taxonomy | Event taxonomy now; broker/transport research only when M3/M4 enters R4. |
| 06 | Canonical identity classes | Person, Aurora, Presence, Device, Provider, Provider Instance, Harness, actor and run are distinct. | What stable identifiers, lifecycle owners and relationships are required in the first technical domain model without collapsing roles into one `user_id`? | M1 and every later milestone | `DECIDE` | Identity/relationship matrix and persistence/rotation invariants; no authentication product required. |
| 07 | Authentication | M0 local owner root is scoped; Presence, services, providers and devices need different proof mechanisms. | How do Leandro, Aurora Core, local services, Harness instances, presences and devices authenticate in Stage A/B, including recovery and step-up? | M1 single-user interface; M2 provider identity; M8 Presence | `RESEARCH` | Threat/topology-specific authentication requirements and official mechanism comparison; mechanism decisions split by actor class. |
| 08 | Authorization, policy and Effect enforcement | Authority is explicit, scoped, expiring, revocable and non-transitive; access is not authority; PDP and Effect Gateway are separate. | What policy input/output contract and enforcement topology supports human, repository, network, credential and later device effects without framework-owned permission? | M3 first delegated effect; M10 physical effect | `DEFER` engine; `RESEARCH` before M3 | Authority/Effect Capability Spec plus allowed/denied/revoked/ambiguous effect drill; engine selected only after contract. |
| 09 | Secrets and credential brokering | Secret values must not enter prompts/manifests/general logs; minimum short-lived credentials are preferred. | Which secret classes exist, where is root custody, how are references resolved at gateways, and which actors may receive tokens rather than values? | First external/model/provider credential use; M2/M3 | `DEFER` product; `DECIDE` handling invariants | Secret-flow/threat map now; broker/product research at first consumer. |
| 10 | Data categories, ownership and lifecycle | Blueprint 11 defines PUBLIC, INTERNAL, CONFIDENTIAL, SENSITIVE, SECRET and DEVICE_RESTRICTED; one owner per concept. | Which canonical/derived/ephemeral datasets exist across Core, memory, history, artifacts, telemetry and providers, and who owns retention/deletion? | Rebaseline; M1 | `DECIDE` | Data inventory, owner, classification, retention, derivation and deletion-impact matrix. |
| 11 | Operational storage | M0 uses SQLite for its local state only; stores are adapters; current state is distinct from memory/history/logs. | Which logical stores are required by M1–M4, which can initially share one physical store, and what evidence would justify separation or PostgreSQL/another mechanism? | M1 governed memory; M2 registry | `DEFER` universal store; `RESEARCH` per consumer | Data/access/consistency/volume profile plus benchmark/fault proof only where mechanism choice is material. |
| 12 | Backup, export, migration and portability | M0 logical export/migration decisions are scoped but establish sovereignty and explicit migration principles. | What must be portable across Core, governed memory, provider-local history, artifacts and derived indexes, and what may be rebuildable/non-portable? | M1 memory deletion/export; M2 provider lifecycle | `DECIDE` portability classes; `DEFER` mechanisms | Portability/restore/deletion classes and migration ownership map; store-specific proof at consuming R4. |
| 13 | Memory, knowledge, retrieval and Context Builder | Memory guides; authority/evidence/live state govern. L0–L5 strata and promotion/supersession are accepted. | What is the M1 implementable write/manage/read architecture, evaluation corpus, derived-index model and deletion/supersession behavior? | M1 | `RESEARCH`, then bounded `SPIKE` | Current memory/RAG/temporal research, representative corpus, adversarial retrieval/deletion evaluations and mechanism comparison. |
| 14 | Model and inference architecture | Models are replaceable capabilities/runtimes, not Aurora identity or authority. | What role taxonomy, provider policy, routing/fallback contract, data-boundary checks and model-attribution records are needed for the first conversational capability? | M1 | `RESEARCH` | Current provider/runtime capabilities, data-policy comparison and eval on real M1 task classes; no global model winner. |
| 15 | Brain and cognitive runtime boundaries | Global cognitive lifecycle is Aurora-owned; Mastra is preferred-first only for agentic Harness/cognitive execution where fit is proven. | Which responsibilities remain deterministic/Core-owned, which are model-mediated, and which can be provider-local without duplicating global Mission, state or memory authority? | M1 and first Mastra-backed capability | `DECIDE` ownership; `RESEARCH` runtime fit | Cognitive responsibility matrix, failure/recovery rules and first-consumer conformance proof. |
| 16 | Harness, AHDK and conformance | ADR-0001/0002 accepted; M2 is first AHDK/Registry milestone; exact language/codegen/binding remain open. | Which minimum contract subset, schema source, first AHDK language and black-box profiles prove SDK/specification independence? | M2 | `DEFER` until M2 R1–R4 | SPK-001-style same-capability SDK/direct proof, build-bound manifest and conformance evidence. |
| 17 | Sandbox and execution environments | Policy and SDK are not security boundaries; environment isolation is required by risk. | What execution profiles exist for research, software, model, untrusted provider and later device work, and what isolation/egress/filesystem/credential guarantees are required? | M2 sandbox validation; M3 effects; Development Harness separately | `RESEARCH`, likely `SPIKE` | Threat-derived profile matrix plus process/container/microVM/provider proof for required properties; do not reuse Development Harness choice automatically. |
| 18 | Durable work, timers and scheduling | No durable engine in M0; provider-local history is not global truth; M4 owns durable Delegation risk. | What domain state remains Aurora-owned while a durable engine/provider owns timers, workflow history and checkpoints, and how is ambiguous-effect reconciliation performed? | M4 | `DEFER` | M4 Capability Spec and comparative restart/idempotency/operations spike across only current candidates. |
| 19 | Artifacts, evidence and provenance | Artifact, Claim, Receipt, Evidence, Verdict and Outcome are distinct; large content moves by reference. | What metadata/content split, content addressing, integrity, retention and source/build provenance are required for M2/M3? | M2 manifests/builds; M3 artifact/evidence flow | `RESEARCH` | Artifact/evidence access and lifecycle model, representative payloads and integrity/provenance proof. |
| 20 | Observability, audit and evaluation | OTel/slog is M0-scoped baseline; telemetry, domain events, audit and evidence remain separate. | What Aurora semantic conventions and correlation/redaction contracts must be stable across Core, models, Harnesses and effects before choosing backends? | M1 model/context evaluation; M2/M3 cross-provider traces | `DECIDE` semantics; `DEFER` backend | Semantic convention and redaction profile plus trace completeness tests; backend chosen only for actual operational need. |
| 21 | Voice and real-time interaction | Voice is a Presence channel with transcription uncertainty, speaker/environment risk, unit read-back and interruption needs. | What streaming, turn detection, barge-in, latency, confirmation and privacy contracts are required, independent of STT/TTS provider? | M8 or an explicitly promoted earlier voice milestone | `DEFER` | Reconsider when Voice enters executable horizon; then current provider research and end-to-end latency/interruption spike. |
| 22 | Vision and multimodality | Images/video are observations with confidence, not verified measurement by default; sensor activation is governed. | What capture, retention, confidence, calibration and cross-modal context contracts apply to workstation/wearable/lab vision? | M8 Presence; M9 observation | `DEFER` | Reconsider at M8/M9 R1; privacy threat model and representative perception/evidence evaluation. |
| 23 | Presence and device trust | One Aurora, multiple presences; Presence is not identity; environment changes disclosure and authority. | How are presences registered, authenticated, attested/revoked and handed off with minimal context in Stage B? | M8 | `DEFER` | M8 spec, device/workload identity research and multi-presence revoke/offline/handoff prototype. |
| 24 | Deployment and topology | Initial direction favors smallest local modular Core; distribution needs evidence; local-first is governance, not one disk forever. | What Stage A/B process/node topology, supervision, update and failure domains preserve sovereignty while allowing external models and isolated providers? | M1/M2 | `DECIDE` architecture stages; `DEFER` deployment product | Stage A/B topology and failure-domain diagram now; packaging/supervision choice at first multi-process implementation. |
| 25 | Configuration and environment management | Versions/configuration must be attributable; secrets are references; environment affects policy/trust. | Which configuration classes are immutable build metadata, deployment config, policy, secret reference or operator preference, and how are changes audited/migrated? | M1/M2 | `RESEARCH` | Configuration taxonomy and precedence/override/secret rules before selecting libraries/products. |
| 26 | Networking and local/cloud boundary | Intelligence may be distributed; sovereignty cannot. Data minimization and provider policy govern external calls. | Which network zones, egress classes, service discovery/channel identities and offline/degraded behaviors exist for Stage A/B? | M1 external models; M2 providers | `DECIDE` zones/egress classes; `DEFER` mesh/broker | Network/trust/data-flow map and provider invocation policy; no service mesh without evidence. |
| 27 | Supply chain and build provenance | Provider trust binds to exact build/environment; manifests are claims; generated contracts must be reproducible. | What source revision, dependency lock, build digest, SBOM/signature/attestation and promotion evidence are required by provider risk? | M2 | `RESEARCH` | Current SLSA/SBOM/signing ecosystem research and one reproducible reference-provider build proof. |
| 28 | Laboratory and physical-device boundary | Observe before actuate; model judgment is not an interlock; device identity, calibration, units and receipts are mandatory. | What controller/device manifest, telemetry time/quality model, leases, gateways and independent interlocks are required for read-only M9 and controlled M10? | M9/M10 | `DEFER` | Reconsider at M9 R1; then simulator/HIL, calibration and safety drill evidence. |

## 6. Current architecture work program

The matrix does not make all 28 areas current work.

### Phase A — Rebaseline completeness

Resolve/document only the program map:

1. system context and trust boundaries;
2. logical modules and ownership;
3. canonical identity classes;
4. data categories and ownership;
5. architecture dependency rules;
6. Stage A/B topology hypotheses;
7. scope labels for every accepted ADR;
8. exact near-horizon consumers.

Outputs remain design/decision proposals. No runtime implementation follows.

### Phase B — First executable-horizon architecture program

After Phase A review, the likely first substantive program is M1 readiness, not a universal platform build:

```text
M1 governed conversation/project context/memory
→ memory/context research and eval corpus
→ model/inference boundary research
→ authentication and single-user interface requirements
→ data/storage/portability decisions only for the M1 slice
→ accepted Capability Spec and R4 decisions
```

M1 order is still subject to operator confirmation after the rebaseline. This landscape does not authorize M1.

### Phase C — Capability fabric horizon

M2/M3 later consume:

- Contract Model representation/versioning;
- AHDK language/codegen;
- Registry/provider identity;
- sandbox profiles;
- build provenance;
- API/event bindings;
- artifacts/evidence;
- authority/effect minimum.

Do not prebuild these during M1 unless a current accepted consumer requires them.

## 7. Decision promotion flow

```text
landscape entry
→ exact question and earliest consumer
→ current primary-source research when needed
→ disposable Architecture Spike only for observed properties
→ proposed ADR / Spec / Standard
→ adversarial review
→ operator acceptance where material
→ consuming Capability R4 revalidation
→ exact Mission Contract and Microdesign
```

The landscape may be detailed without becoming the decision owner.

## 8. Stop rules

Stop and open a Finding when:

- two modules appear to own the same canonical concept;
- a proposed product is selected before its required contract/threat/data model exists;
- a local M0 mechanism is treated as universal without evidence;
- a framework/provider type leaks into Aurora canonical semantics;
- a deferred area receives implementation code without a promoted consumer;
- the Development Harness is made a runtime dependency of Aurora;
- a research report is treated as an accepted decision;
- an Architecture Spike begins without separate authorization;
- the landscape grows without eliminating uncertainty or naming consumers.

## 9. Explicit non-decisions

This revision does not select:

- a universal database or persistence stack;
- Keycloak, Zitadel, Authentik, Ory or another authentication platform;
- SPIFFE/SPIRE, OAuth Token Exchange or another workload identity mechanism;
- OPA, Cedar or another policy engine;
- Vault or another credential broker;
- REST, gRPC, MCP, A2A, ACP, CloudEvents, AsyncAPI or a broker as a universal binding;
- Kubernetes, a service mesh, high availability or multi-tenancy;
- a vector database, graph database or memory framework;
- an LLM/model provider or universal router;
- STT, TTS, voice transport or wearable hardware;
- a durable workflow engine;
- an artifact, telemetry or observability backend;
- a laboratory protocol/device stack;
- the first AHDK language;
- the next Product Milestone execution authorization.

## 10. Exit condition for this proposed landscape revision

This initial landscape is ready for operator review only when:

1. all 28 areas have constraints, questions, consumers, treatment and evidence direction;
2. scoped accepted decisions are not globalized;
3. current versus deferred work is explicit;
4. every `DEFER` has a reconsideration trigger;
5. no hidden stack selection or implementation appears;
6. ACRM integration and current STATUS agree;
7. mechanical documentation validation passes;
8. an adversarial review records Findings and limitations.
