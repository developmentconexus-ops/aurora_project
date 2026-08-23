# TA-03 Working Analysis — Cross-System Operation Surface

> **TEMPORARY / NON-AUTHORITATIVE / BRANCH ONLY.** This file is analysis support for TA-03. It must be absorbed into durable architecture or deleted before any merge candidate is promoted.

## 1. Fixed authority baseline

```text
canonical main at start:
d151dd2e2bd069af2f1dbffd9608d3b7ce878022

current branch:
docs/ta-03-cross-system-operation-surface-20260823

TA-01 Logical Modules & Canonical Ownership:
ACCEPTED / CANONICAL

TA-02 Process/Runtime/Evolutionary Topology:
ACCEPTED / CANONICAL

TA-03 execution:
OPERATOR-AUTHORIZED / IN PROGRESS

TA-04+:
NOT AUTHORIZED

Architecture Spikes:
NOT AUTHORIZED

Aurora Product/runtime implementation:
BLOCKED

M0 R7 continuation/Verdict/R8:
NOT AUTHORIZED
```

Governing read set was deliberately expanded beyond the normal five-file fresh-actor pack because TA-03 is the accepted stage whose subject is the relationship between multiple constitutional/architecture owners. The expansion is bounded to Product/architecture owners needed to decide Presence/interaction, memory/context, provider/Harness, authority/effect and accepted Stage-A/runtime seams.

## 2. TA-03 target property

A future TA-04 or Capability actor must not need to invent:

- whether a cross-boundary interaction is a read, command, proposal, effect or callback/observation;
- who owns the operation meaning;
- who may request versus commit;
- which canonical state may change;
- what a provider/runtime result does **not** mean;
- what failure/ambiguity requires reconciliation;
- where idempotency/concurrency/cancellation/deadline are semantically required;
- whether C12 exact history or C07 Evidence/receipt is required;
- which shortcut is forbidden.

TA-03 does **not** decide encoding, schema technology, RPC/event protocol, HTTP shape, generated SDK surface, database, IAM product, model/provider product or implementation.

## 3. Evidence classification

### Known

- one durable Aurora concept has one canonical owner;
- G01 owns cross-system semantic contract governance, not runtime business state;
- C01–C12 own accepted domain meanings;
- A01/A02/A03 coordinate but do not gain canonical truth;
- A05 owns Aurora-side separately-running provider lifecycle policy;
- B01 already fixes transport-neutral provider identity/lifecycle/idempotency/cancellation/reconciliation semantics;
- P01/P02 own provider-local execution state only;
- C03 owns interpreted Intent, Mission/Delegation/global Attempt, budgets and global Outcome;
- C04 owns Authority Grant, Policy Decision and Effect Request semantics;
- C06 owns governed memory and Context Pack provenance;
- C07 owns Artifact/Observation/Evidence/Receipt/Verdict records;
- C08 owns Presence/Interaction Session semantics;
- C12 owns Audit/Exact History append semantics;
- E01 executes/denies effects and produces Effect Receipts but does not create authority;
- process death/provider completion/transport delivery/telemetry never imply canonical success;
- authority is non-transitive and effect ambiguity blocks blind retry;
- Stage A activation is a Presence request, not authentication or authority.

### Inferred, subject to TA-03 decision

- the sustainable operation surface should be organized around named semantic boundary operations rather than storage-shaped CRUD or one universal command envelope;
- common metadata obligations should be semantic rules applied as needed, not one mandatory physical envelope;
- owner-public read/proposal/command operations should preserve one semantic owner even when A01 coordinates them;
- provider output should be routable only to a known admitted target operation; a generic `OUTPUT_PROPOSAL` must not become an arbitrary mutation escape hatch;
- current TA-03 should admit the foundational Presence/context/provider/delegation/effect/evidence/history seams and defer capability-specific future seams whose accepted owner already exists and can be added later without dismantling authority.

### Unknown — intentionally left to later owning stages

- exact schema representation and code generation;
- exact wire/protocol/binding per boundary;
- exact channel authentication;
- physical stores and atomicity/outbox mechanisms;
- authentication/IAM/secret products;
- concrete policy/effect/credential products;
- concrete cognitive runtime/model/provider integration;
- runtime supervisor product;
- remote-node topology mechanisms.

### Deferred candidates

See §10. Deferral is safe only when an accepted owner/seam exists and later addition does not require semantic-owner movement.

## 4. Operation-surface granularity alternatives

### Alternative A — exhaustive owner CRUD

```text
C01 CRUD
C02 CRUD
...
C12 CRUD
```

Advantages:
- superficially complete;
- easy to turn into conventional APIs.

Rejected direction because:
- storage/entity shape leaks into architecture;
- encourages screen/database-shaped APIs;
- invents operations with no real consumer;
- obscures proposal versus canonical commit;
- creates a large speculative contract surface before Capability semantics exist.

### Alternative B — one universal OperationEnvelope / CommandBus

```text
operation(type, target, payload, metadata)
```

Advantages:
- small transport surface;
- uniform middleware seems attractive.

Rejected direction because:
- semantic differences collapse into strings/payloads;
- one generic route becomes an authority bypass risk;
- provider proposals can accidentally become arbitrary mutation;
- domain failure, concurrency, cancellation and Evidence semantics become hidden conventions;
- framework/protocol machinery can become de facto Aurora ontology.

### Alternative C — named semantic operations grouped into families

```text
family = navigation/taxonomy only
named operation = one semantic owner + explicit consequences
common requirements = invariants/profile, not one executable super-envelope
TA-04 = exact machine-readable projections/bindings
```

Working recommendation: **Alternative C**.

This preserves differentiated semantics while avoiding an endpoint per private method. A family is never itself a callable universal API.

## 5. Cross-system operation classes

Use only the Planning Readiness classes:

### READ

Requests owner-governed truth/projection. No canonical mutation. Material reads expose sufficient revision/freshness/provenance for the consuming decision.

### COMMAND

Requests the semantic owner to validate and, when valid, commit its own state transition. The caller does not receive owner authority merely because it may issue the command.

### PROPOSAL

A non-owner submits a candidate/request. The target owner may accept, record, transform, reject or require confirmation according to its own semantics. The caller cannot assume the target canonical state changed.

### EFFECT

Requests a consequential external/physical/digital action through the accepted C04 → E01 enforcement path. External success/failure can remain ambiguous and never implies an Aurora owner transition by itself.

### CALLBACK / OBSERVATION

Reports an externally/provider-produced fact, status, output, receipt or observation. It is attributable input. It never becomes another owner's truth merely because it arrived.

## 6. Global operation invariants — working candidate

### INV-OP-01 — One semantic owner

Every admitted operation has one semantic owner. Coordinators, transports, providers, stores, UI and generators cannot become parallel owners.

### INV-OP-02 — Caller permission is not commit authority

```text
READ     → consume only
PROPOSE  → candidate only
CALLBACK → reported input only
COMMAND  → owner validates and commits its own state only
EFFECT   → E01 executes only after C04 decision
```

### INV-OP-03 — No generic mutation escape hatch

Natural-language intent, provider `OUTPUT_PROPOSAL`, manifest, telemetry event, transport message or generic payload cannot target arbitrary canonical state. It must map to an admitted operation owned by the target owner, otherwise it is rejected/escalated or retained only as a non-authoritative artifact/claim/observation where applicable.

### INV-OP-04 — Logical operation identity precedes retry

Material commands/effects/provider requests carry a stable logical operation/request identity when duplication could change semantics. A new execution attempt is distinct from retry delivery. Unknown completion blocks blind replay.

### INV-OP-05 — Concurrency is explicit where stale state is material

A state-changing operation whose correctness depends on current state must carry a semantic precondition/revision/freshness constraint sufficient for the owner to reject stale/conflicting application. TA-04/TA-06 decide representation and physical enforcement.

### INV-OP-06 — Cancellation is intent, not proof

Cancellation/deadline expiry does not prove that provider/external work stopped or effects rolled back. Acknowledgement/snapshot/reconciliation is required where completion can be ambiguous.

Authority revocation is separate and must have operational effect at gateways even while cooperative cancellation is unresolved.

### INV-OP-07 — Provider terminal state is not global Outcome

P01/P02 completion is provider execution evidence only. C03 owns global Attempt/Delegation/Mission state and Outcome; C07 evidence/verdict may be required before C03 closure.

### INV-OP-08 — Effect outcome is receipt/reconciliation driven

No external effect is repeated because a timeout looks like failure. E01 returns/recovers a receipt or explicit ambiguity; C07 records the receipt and C12 records attributable history according to profile.

### INV-OP-09 — Exact history is not Evidence and telemetry is neither

C12 exact history, C07 Evidence/Receipt, domain state and telemetry remain separate. One may reference another only through explicit semantics.

### INV-OP-10 — Common metadata is a semantic profile, not a universal wire envelope

As applicable an operation may require:

```text
logical operation/request identity
actor / subject / executor / Presence refs
Interaction / Mission / Delegation / parent refs
semantic contract/version ref
owner/resource/action refs
source/target revision or freshness precondition
authority / policy refs
deadline / cancellation semantics
idempotency identity
correlation / causation refs
data classification
C12 profile
C07 Evidence/receipt profile
```

TA-03 does not require every field on every operation and does not create one physical `OperationEnvelope` type.

## 7. Candidate admitted operation families

The IDs below are working IDs only. They are not canonical until operator-approved TA-03 design is promoted.

### 7.1 Presence and interaction

| Working ID | Operation | Class | Semantic owner | Real consumer / protected property |
|---|---|---|---|---|
| TA03-PRS-01 | `RequestActivation` | PROPOSAL | C08 | Stage A/M1; one activation meaning across button/hotkey/UI/wake-word; trigger is not auth/authority |
| TA03-PRS-02 | `SubmitInteractionContribution` | CALLBACK/OBSERVATION | C08 | M1; attributable human/system input without treating raw input as interpreted Intent |
| TA03-PRS-03 | `DeliverInteractionOutput` | COMMAND | C08 | TA-05/M1; Presence/environment/privacy-aware delivery without giving UI business authority |
| TA03-PRS-04 | `CloseInteraction` | COMMAND | C08 | M1; explicit session termination/timeout boundary and exact-history closure |

Working consequence law:
- `RequestActivation` can open, restrict, require step-up, defer or deny interaction according to accepted policy; locked-state restrictions remain fixed.
- no Presence operation may mutate Project/Mission/Authority/Memory directly.

### 7.2 Intent and Mission entry

| Working ID | Operation | Class | Semantic owner | Real consumer / protected property |
|---|---|---|---|---|
| TA03-INT-01 | `ProposeIntentCandidate` | PROPOSAL | C03 | M1; A02/P01/C08 interpretation cannot become committed Intent by model assertion |
| TA03-INT-02 | `PromoteIntentToMission` | COMMAND | C03 | M3; Mission creation is an explicit C03 transition, not automatic classification |

Working consequence law:
- C03 may validate/reject/expire/commit an Intent proposal according to its owner rules;
- promotion requires current validated Intent plus purpose/authority/preconditions;
- raw transcript remains C08/C12, not C03 truth.

### 7.3 Owner-authoritative context and governed memory

Owner reads are named semantic projections, not generic database queries.

| Working ID | Operation | Class | Semantic owner | Real consumer / protected property |
|---|---|---|---|---|
| TA03-CTX-01 | `ReadIdentityContext` | READ | C01 | C04/C06/C08; actor/relationship refs without authenticating by profile data |
| TA03-CTX-02 | `ReadProjectContext` | READ | C02 | M1/C06; authoritative current Project context with revision/freshness |
| TA03-CTX-03 | `ReadAuthorityContext` | READ | C04 | C06/A01/A03; current grant/policy projection without copying authority into memory |
| TA03-CTX-04 | `QueryGovernedMemory` | READ | C06 | M1; scoped/provenance-aware memory retrieval below current authority |
| TA03-CTX-05 | `BuildContextPack` | COMMAND | C06 | M1/M2/M3; owner-governed minimized immutable/provenance-bearing projection for one consumer/purpose |
| TA03-MEM-01 | `SubmitMemoryCandidate` | PROPOSAL | C06 | M1; external/model/session synthesis cannot self-promote into durable memory |
| TA03-MEM-02 | `ManageMemoryItem` | COMMAND | C06 | M1; promote/supersede/correct/delete according to memory governance, with explicit downstream deletion/supersession consequences |

`ManageMemoryItem` is one owner lifecycle operation family only if the final design can specify its materially different actions without hiding them behind arbitrary verbs. Otherwise TA-03 should split promotion/supersession/deletion before acceptance.

### 7.4 Capability/provider governance

| Working ID | Operation | Class | Semantic owner | Real consumer / protected property |
|---|---|---|---|---|
| TA03-REG-01 | `SubmitProviderManifest` | PROPOSAL | C05 | M2/TA-08; manifest is claim, never approval or authority |
| TA03-REG-02 | `DecideProviderApproval` | COMMAND | C05 | M2; approval binds exact provider/build/environment/scope and current evidence/compatibility |
| TA03-REG-03 | `SuspendOrRevokeProvider` | COMMAND | C05 | M2+; trust/approval change blocks new use and triggers active-work containment/reconciliation |
| TA03-REG-04 | `ResolveProvider` | READ | C05 | A02/A03; resolve eligible approved provider refs without creating Delegation or Authority Grant |

Verification/Evidence remains C07-owned. C05 may consume evidence references and own compatibility/trust/approval assessment; it must not duplicate C07 evidence records or G01 semantics.

### 7.5 Mission/Delegation control

| Working ID | Operation | Class | Semantic owner | Real consumer / protected property |
|---|---|---|---|---|
| TA03-DEL-01 | `CreateDelegation` | COMMAND | C03 | M3; exact purpose/capability/provider/context/authority/budget/deadline binding before execution |
| TA03-DEL-02 | `RequestChildCapability` | PROPOSAL | C03 | M3+; provider cannot transitively delegate authority/context/budget |
| TA03-DEL-03 | `CancelDelegation` | COMMAND | C03 | M3/M4; global cancellation intent separated from provider stop acknowledgement and authority revocation |
| TA03-DEL-04 | `ReconcileAttemptState` | COMMAND | C03 | B01 restart/failure; missing provider state never inferred as success |
| TA03-DEL-05 | `CommitMissionOutcome` | COMMAND | C03 | M3+; provider completion cannot close global Outcome without required evidence/verdict composition |
| TA03-DEL-06 | `ReadDelegationState` | READ | C03 | Presence/A01/recovery; current global state distinct from provider-local state |

Budget consumption observations should be carried from provider/effect outputs but cannot become C03 budget truth automatically. If TA-03 finds that reconciliation needs a distinct operation, add it rather than overloading `ReconcileAttemptState`.

### 7.6 Runtime lifecycle and B01 provider execution

| Working ID | Operation | Class | Semantic owner | Real consumer / protected property |
|---|---|---|---|---|
| TA03-RUN-01 | `EnsureProviderRuntimeReady` | COMMAND | A05 | first separate provider; exact approved provider instance → new runtime incarnation → readiness before dispatch |
| TA03-RUN-02 | `ControlProviderRuntime` | COMMAND | A05 | first separate provider; start/attach/drain/stop/restart policy without owning Delegation terminal semantics |
| TA03-PRV-01 | `SubmitProviderRequest` | COMMAND | G01/B01 profile | A03→P01/P02; transport-neutral request/attempt/idempotency/authority/context/deadline semantics |
| TA03-PRV-02 | `CancelProviderRequest` | COMMAND | G01/B01 profile | cooperative provider cancellation with idempotent acknowledgement and known/unknown state |
| TA03-PRV-03 | `QueryProviderSnapshot` | READ | G01/B01 profile | restart/reconciliation; current provider-local state without inferred canonical success |
| TA03-PRV-04 | `ReportProviderOutput` | CALLBACK/OBSERVATION | G01/B01 profile | progress/proposal/artifact/claim/capability/effect/terminal outputs remain typed reported material routed to exact Aurora owner |

For B01 operations G01 owns cross-system semantic contract meaning; C03 owns global Delegation/Attempt state and P01/P02 own provider-local execution state. The final TA-03 design must state this split explicitly so `semantic owner` is not mistaken for business-state ownership.

`ReportProviderOutput` must not be a universal mutation endpoint. Every output category either maps to another admitted operation (`RequestChildCapability`, effect flow, artifact/claim flow, etc.) or remains a non-authoritative report/artifact.

### 7.7 Authority and effects

| Working ID | Operation | Class | Semantic owner | Real consumer / protected property |
|---|---|---|---|---|
| TA03-AUT-01 | `IssueAuthorityGrant` | COMMAND | C04 | M3/effects; explicit scoped/expiring/non-transitive authority |
| TA03-AUT-02 | `RevokeAuthorityGrant` | COMMAND | C04 | all effects; revocation must block future execution operationally even if cooperative work is still reconciling |
| TA03-AUT-03 | `EvaluateEffectRequest` | COMMAND | C04 | first effect; produce canonical Policy Decision from actor/purpose/resource/action/context/grant state without executing |
| TA03-EFF-01 | `ExecuteAuthorizedEffect` | EFFECT | C04 semantics; E01 enforcement/executor | first material effect; execute only exact current decision and preserve external ambiguity |
| TA03-EFF-02 | `ReconcileEffectOutcome` | EFFECT | C04 semantics; E01 enforcement/executor | unknown timeout/receipt; observe target/receipt without blindly repeating the effect |

Credential resolution remains behind E01/E02. No public provider-facing credential operation is admitted yet. TA-07 may require TA-03 reopen if a concrete credential/runtime topology introduces a new material boundary. Until then the invariant is: provider/model never receives raw secret by convenience and E01 cannot use a credential outside the exact C04 decision.

### 7.8 Artifact, Observation, Claim, Evidence, Receipt and Verdict

| Working ID | Operation | Class | Semantic owner | Real consumer / protected property |
|---|---|---|---|---|
| TA03-RES-01 | `RegisterArtifact` | CALLBACK/OBSERVATION | C07 | M2/M3; producer may publish content/integrity ref but cannot self-promote evidence |
| TA03-RES-02 | `RecordObservation` | CALLBACK/OBSERVATION | C07 | M3/M9; immutable provenance/quality/units without producer rewrite |
| TA03-RES-03 | `SubmitClaim` | PROPOSAL | C07 | M2/M3; provider assertion is not verified result |
| TA03-RES-04 | `RecordEvidence` | COMMAND | C07 | M2/M3; authorized verifier links method/result/limitations to claim/criterion/artifacts/observations |
| TA03-RES-05 | `RecordEffectReceipt` | CALLBACK/OBSERVATION | C07 | effects; only E01 produces the effect receipt, C07 owns canonical record/custody |
| TA03-RES-06 | `IssueVerdict` | COMMAND | C07 | M3+; only authorized verifier may create Verdict record according to profile |
| TA03-RES-07 | `ReadEvidencePackage` | READ | C07 | C03/operator/review; read proof without conflating it with Outcome/current state |

### 7.9 Audit and exact history

| Working ID | Operation | Class | Semantic owner | Real consumer / protected property |
|---|---|---|---|---|
| TA03-HIS-01 | `AppendExactHistory` | COMMAND | C12 | every governed material mutation/provider/effect/interaction profile requiring exact attribution/reconstruction |
| TA03-HIS-02 | `ReadExactHistory` | READ | C12 | recovery/audit/C06/operator; exact record without becoming memory/evidence/current state |

Where a future operation profile requires C12 append for correctness, append failure cannot be silently ignored. Physical transaction/outbox/atomicity is TA-06/TA-10 work; TA-03 owns only the semantic obligation and fail-closed/explicit-integrity consequence.

## 8. Candidate common semantics — not wire decisions

### 8.1 Identity and retry

Material operations distinguish:

```text
logical operation/request identity
execution attempt identity when execution occurs
idempotency identity when duplicate delivery can repeat consequence
correlation/causation identity for reconstruction
```

B01's accepted `request_id / attempt_id / correlation_id / idempotency_key` remains the provider-boundary specialization. TA-03 should not force the same field names onto every owner operation.

### 8.2 Preconditions and concurrency

Examples of semantic preconditions:

```text
expected owner revision / current lifecycle state
valid Intent or Delegation state
provider approval/build/environment still current
runtime incarnation READY
Authority Grant still valid
Policy Decision still current
context/source freshness within requirement
deadline not expired
required evidence/history path available
```

A failed precondition is not transport failure and must not be retried blindly.

### 8.3 Failure and ambiguity

TA-03 records the meaningful failure cases per operation. It does not create the exact TA-04 wire error taxonomy.

At minimum the design must be able to express materially distinct conditions such as:

```text
invalid semantic request
stale/conflicting precondition
unauthorized or revoked context
incompatible provider/contract
unavailable dependency/runtime
canceled/deadline exceeded
provider-local failure/state lost
external effect outcome unknown
reconciliation required
history/evidence integrity path unavailable
```

### 8.4 History/Evidence profile

Every admitted operation must state:

```text
C12 exact-history requirement:
NONE | MATERIAL-TRANSITION | ALWAYS/PROFILED

C07 proof implication:
NONE | ARTIFACT | OBSERVATION | CLAIM | RECEIPT | EVIDENCE | VERDICT
```

These are analysis categories, not proposed persisted enums yet.

## 9. Flow coverage used to falsify the operation set

TA-03 should prove completeness against accepted architecture flows without claiming TA-11 Golden Flow status.

### Flow A — Stage A activation to governed response

```text
Presence trigger
→ RequestActivation
→ C08 opens/restricts interaction
→ SubmitInteractionContribution
→ C12 exact history
→ A02 interpretation
→ ProposeIntentCandidate
→ C03 validates/commits or rejects
→ BuildContextPack if cognition needed
→ ResolveProvider / EnsureProviderRuntimeReady
→ SubmitProviderRequest
→ ReportProviderOutput
→ DeliverInteractionOutput
```

Failure probes:
- locked workstation;
- revoked Presence;
- heavy runtime unavailable;
- model output claims a Project mutation;
- exact history unavailable where required.

### Flow B — Provider admission

```text
SubmitProviderManifest
→ G01 compatibility consumed by C05
→ C07 verification Evidence refs
→ DecideProviderApproval
→ ResolveProvider
→ EnsureProviderRuntimeReady
```

Failure probes:
- build digest/environment changes;
- manifest declares more effects;
- stale verification;
- provider self-declares trust;
- runtime identity differs from approved instance.

### Flow C — Delegation and child capability

```text
PromoteIntentToMission
→ CreateDelegation
→ BuildContextPack
→ IssueAuthorityGrant
→ ResolveProvider
→ EnsureProviderRuntimeReady
→ SubmitProviderRequest
→ provider emits capability request
→ RequestChildCapability
→ C03 creates narrower child only if accepted
```

Failure probes:
- child tries to inherit parent token/context/budget;
- provider starts child directly;
- parent canceled while child runs;
- provider terminal result arrives before required evidence.

### Flow D — Material effect

```text
provider/Aurora proposes Effect Request
→ EvaluateEffectRequest
→ ExecuteAuthorizedEffect
→ RecordEffectReceipt
→ AppendExactHistory
→ owner reconciliation/transition
```

Unknown path:

```text
ExecuteAuthorizedEffect timeout / unknown
→ NO BLIND RETRY
→ ReconcileEffectOutcome
→ receipt | confirmed-no-effect | still-ambiguous
→ owner proceeds/blocks accordingly
```

Failure probes:
- grant revoked after decision but before execution;
- credential unavailable;
- duplicate request;
- gateway returns success without target/receipt support;
- provider fabricates receipt.

### Flow E — Memory/context

```text
C08/C12 interaction material
→ SubmitMemoryCandidate
→ C06 governance
→ ManageMemoryItem
→ later QueryGovernedMemory
→ ReadProjectContext / ReadAuthorityContext / sources
→ BuildContextPack
→ provider receives minimized Context Pack
```

Failure probes:
- memory conflicts with accepted source;
- superseded item retrieved as current;
- cross-project memory leakage;
- delete does not propagate to derived indexes;
- provider requests full memory store.

### Flow F — provider crash/restart

```text
A05 observes runtime failure
→ provider dispatch blocked
→ new runtime incarnation only under current approval
→ QueryProviderSnapshot
→ ReconcileAttemptState
→ C12 preserves reconstruction path
→ resume/new Attempt/block by C03
```

Failure probes:
- process restart creates duplicate attempt;
- state lost reported as success;
- stale authority/deadline used on resume;
- provider-local checkpoint copied into canonical state.

## 10. Working DEFER SAFELY register

These operations/seams have accepted logical owners but no current need to stabilize full cross-system semantics during TA-03 unless later analysis shows they are prerequisite for TA-04–TA-10.

| Subject | Current owner/seam | Earliest known consumer/trigger | Preserve now |
|---|---|---|---|
| Durable timers/waits/checkpoint-engine operations | C03 semantics + A04 port | M4 durable Delegation | engine cannot own Mission/Delegation truth; no inferred success; explicit reconciliation |
| Attention/proactivity nomination/delivery | C09 | M5 or earlier promoted proactivity | provider cannot self-interrupt; attention budget remains owner-governed |
| Failure-intelligence/evaluation/improvement operations | C10 | M6/M7 | evaluation cannot self-promote; Evidence/Outcome owners unchanged |
| Environment/device registry lifecycle | C11 | M8/M9 | device identity distinct from Presence/Provider; live telemetry not registry truth |
| Presence handoff / sensor session / ambient campaign operations | C08 | M8 or explicit earlier second-Presence/sensor consumer | Presence is not identity; context minimized; sensor capability is not activation authority |
| Direct high-volume data-channel establishment | Aurora control-plane owners + adapters | first real rate/size/latency consumer where reference transfer is insufficient | control/authority never bypassed; channel scoped/classified/revocable/auditable |
| Credential Broker as separately callable/runtime operation | E02 boundary; C04 permission semantics | TA-07 only if concrete topology requires a distinct material boundary | raw secrets never provider/model context; minimum scoped credential only |
| Device-specific effect verbs | C04 + E01 + C11/device source | M9/M10 | generic effect authorization/receipt/reconciliation law applies; deterministic interlocks independent |
| Remote-node migration/cutover operations | current domain owners + conditional TA-TX | Stage B/cutover trigger | domain identities/owners preserved; provider instance/incarnation identities re-established |

A deferred subject must reopen TA-03 if adding it later would otherwise require inventing a new cross-system operation during implementation or moving an accepted owner.

## 11. Strongest current adversarial objections

### Objection 1 — This is still too many operations

Response to test:
- remove any operation that is merely a private owner method;
- retain only entries whose removal forces a downstream actor to invent a material cross-boundary semantic decision;
- combine only when failure/authority/state consequences truly match.

### Objection 2 — `ManageMemoryItem` hides materially different operations

Likely valid. Promotion, supersession/correction and deletion have different authority/deletion-impact semantics. The final design may need to split them despite sharing C06 ownership.

### Objection 3 — B01 semantic owner is ambiguous

Must resolve explicitly:
- G01 owns B01 contract semantics/versioning;
- C03 owns global Delegation/Attempt state;
- P01/P02 own provider-local execution;
- A03 owns dispatch/reconciliation mechanics;
- A05 owns runtime process lifecycle policy.

If one `semantic owner` field cannot express this without ambiguity, the operation record needs `semantic owner` plus explicit `state consequence owner/executor` rather than changing accepted ownership.

### Objection 4 — `ExecuteAuthorizedEffect` appears C04-owned but E01 executes

This is deliberate if described correctly:
- C04 owns Effect Request/Policy Decision meaning and permission;
- E01 is the deterministic enforcement/execution boundary;
- C07 owns the canonical Effect Receipt record;
- C12 owns attributable exact history;
- external target owns its live external state.

TA-03 must not give E01 authority-generation semantics.

### Objection 5 — current M1 does not need M3/M4 operations

Planning Readiness is cross-system and must prepare downstream stages, but YAGNI still applies. Keep only foundational M2/M3 seams whose omission would make TA-04/TA-07/TA-08 invent foundational architecture. A04/M4 and distant C09–C11 operations are currently safe to defer because accepted seams/owners already exist.

## 12. Current material decision set before durable design

1. **Granularity:** accept Alternative C — named semantic operations grouped into non-callable families; no CRUD matrix and no universal operation envelope.
2. **Commit law:** accept that non-owner outputs are READ/PROPOSAL/CALLBACK only; owner COMMAND is the only canonical mutation path; EFFECT is exclusively C04→E01 mediated.
3. **Common metadata:** define conditional semantic obligations, not one physical common envelope.
4. **Provider routing:** require every provider output that could change Aurora to map to an admitted target operation; otherwise it cannot mutate canonical state.
5. **Concurrency/retry:** require semantic preconditions and stable logical identity only where the failure class exists; do not universalize M0 revision fields or B01 field names.
6. **Scope:** admit foundational Stage-A/M1 + provider/delegation/effect/evidence/history seams; DEFER SAFELY A04/M4, C09–C11, multi-Presence/sensors, direct data channels and topology-specific credential operations until named triggers.
7. **Memory lifecycle:** decide whether to split `ManageMemoryItem` into separate promote/supersede/delete operations before candidate design.

## 13. Proof path before candidate

```text
current authority reconstruction
→ operation admission test per entry
→ flow coverage A–F
→ operation-without-consumer / consumer-without-operation sweep
→ owner duplication sweep
→ forbidden-shortcut counterexamples
→ self-review against Planning Readiness TA-03 exit law
→ operator design review
→ durable candidate architecture document
→ exact-revision documentation validation
→ isolated Fable review
→ Lead adjudication
→ final verification
→ separate operator ratification
→ separate merge authorization
```

No TA-04 work or runtime implementation is authorized by this analysis.
