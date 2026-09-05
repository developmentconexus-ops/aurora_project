# TA-03 Working Self-Review

> **TEMPORARY / NON-AUTHORITATIVE / BRANCH ONLY.** This file records corrections to `ta-03-analysis.md` before operator design review. It must be absorbed/deleted before the final TA-03 candidate.

## 1. Review lenses

The working operation set was challenged against:

```text
operation without real consumer
consumer without admitted operation
operation without one semantic owner
owner duplication or coordinator-as-owner
private owner method globalized by convenience
proposal/callback capable of arbitrary mutation
generic operation hiding materially different failure/authority semantics
provider/framework semantics promoted to Aurora meaning
retry/cancellation ambiguity
C12/C07/telemetry conflation
future capability implemented instead of seam prepared
```

## 2. Corrections

### SR-01 — Split memory lifecycle commands

`ManageMemoryItem` is rejected as too broad.

The accepted Product distinguishes candidate promotion, correction/supersession and deletion as materially different lifecycle consequences. TA-03 should use:

```text
SubmitMemoryCandidate       PROPOSAL  → C06
PromoteMemoryCandidate      COMMAND   → C06
SupersedeMemoryItem         COMMAND   → C06
DeleteMemoryItem            COMMAND   → C06
```

Semantics:

- promotion creates governed durable memory from a candidate only under C06 policy and retains provenance;
- correction of semantic content is represented through supersession/replacement lineage rather than silent overwrite;
- deletion is separate because it creates retention/privacy/derived-index/non-retrieval obligations and may be constrained by authoritative source/evidence retention;
- none of these operations changes C02/C03/C04 authoritative state.

### SR-02 — Admit Decision Request lifecycle

Blueprint 07 makes decision request/resolution a material Delegation control-plane concern. Omitting it would force M3/TA-04 to invent a foundational operation.

Add:

```text
SubmitDecisionRequest       PROPOSAL  → C03 pending-decision lifecycle/routing
ResolveDecisionRequest      COMMAND   → C03 pending-decision lifecycle/routing
```

`SubmitDecisionRequest` records the exact question, why local authority is insufficient, alternatives/evidence/recommendation, blocked work, safe default and deadline as applicable.

`ResolveDecisionRequest` records an attributable resolution or authoritative decision reference and releases/redirects the blocked Delegation according to current state.

Critical non-owner rule:

> C03 owns the pending Decision Request lifecycle and orchestration consequence. It does **not** acquire the semantic authority of the decision being requested. A Product/architecture/policy/provider/effect decision remains with its existing owner; any resulting canonical domain change still invokes that owner's admitted operation.

Examples:

```text
provider asks for new architecture choice
→ C03 records/routes Decision Request
→ operator/architecture authority decides outside provider
→ C03 records resolution reference
→ no architecture mutation is created by ResolveDecisionRequest

provider asks for broader effect authority
→ C03 records/routes Decision Request
→ C04/operator authority issues or refuses a grant through C04 semantics
→ C03 records resolution and resumes/blocks work
```

### SR-03 — Admit budget/consumption reporting without trusting provider counters

C03 owns global budgets/reconciled consumption, while providers/effect boundaries may observe usage. Add:

```text
ReportConsumption           CALLBACK/OBSERVATION → C03
```

The report is attributable input, not canonical budget truth by itself. C03 reconciles accepted provider/receipt/accounting evidence, applies soft/hard thresholds and blocks new work when material consumption is ambiguous according to the current budget profile.

Do not create a universal billing system or physical storage model in TA-03.

### SR-04 — Clarify B01 operation ownership

Replace working `semantic owner = G01/B01 profile` wording.

Correct split:

```text
semantic contract owner      → G01
provider-boundary profile    → B01 accepted TA-02 semantics
Aurora global work state     → C03
provider-local execution     → P01/P02
transport/dispatch mechanics → A03
process lifecycle policy     → A05
```

B01 is an accepted boundary profile, not a second semantic owner beside G01.

For every B01 operation record, use:

```text
semantic owner: G01
state consequence: explicit C03 and/or provider-local state as applicable
coordinator/executor: A03 / provider runtime as applicable
```

### SR-05 — Separate semantic owner from executor/enforcement boundary

The Planning Readiness `semantic owner` field is insufficient by itself to communicate operations where meaning, state consequence and execution belong to different accepted components.

TA-03 records, as applicable:

```text
semantic owner
state consequence owner(s)
coordinator / executor / enforcement boundary
```

This is descriptive decomposition of already accepted ownership, not a new owner model.

Canonical example:

```text
ExecuteAuthorizedEffect
semantic owner: C04 (Effect Request / Policy Decision meaning)
enforcement/executor: E01
receipt producer: E01
canonical receipt owner: C07
exact-history owner: C12
external live state: target system
```

E01 never creates authority.

### SR-06 — Replace generic runtime-control command with bounded lifecycle operations

`ControlProviderRuntime` is rejected as an arbitrary verb envelope.

TA-02 already fixes A05's lifecycle policy. TA-03 should expose only the material cross-process seams needed by current consumers:

```text
EnsureProviderRuntimeReady   COMMAND → A05
DrainProviderRuntime         COMMAND → A05
StopProviderRuntime          COMMAND → A05
ReconcileProviderRuntime     COMMAND → A05
```

`EnsureProviderRuntimeReady` may start or attach the exact approved provider instance and always establishes a fresh concrete `runtime_incarnation_id` when a process is started. It cannot create a new C03 Attempt.

`Drain` and `Stop` affect process acceptance/lifecycle only; they do not prove active provider requests or Delegations ended semantically.

`ReconcileProviderRuntime` maps observed process state/incarnation back to A05 desired state after restart/failure. B01 snapshot/reconciliation remains separately required for request semantics.

A distinct `RestartProviderRuntime` operation is not admitted now: restart is an A05 policy consequence of current approved-instance state plus failure/recovery policy, and any new process start is reflected through a new incarnation. If a future operator/consumer needs externally commanded restart semantics, reopen this surface.

### SR-07 — Keep interaction close/delivery operations

`CloseInteraction` and `DeliverInteractionOutput` survive the private-method challenge because both cross the Presence boundary and protect user-observable/privacy/history semantics:

- a Presence disconnect/timeout/explicit close must not silently imply C03 Mission cancellation;
- output delivery is environment/privacy/capability constrained and must not let a UI bypass C08 by reading arbitrary owner state.

### SR-08 — Keep C07 Verdict and C12 exact-history operations

They are not distant ceremony:

- M3/provider completion requires evidence-oriented acceptance before C03 global Outcome;
- C12 append is already a foundational active owner and materially required by accepted mutation/provider/effect profiles.

TA-03 still does not define physical Evidence/history stores or atomicity mechanisms.

## 3. Corrected operation deltas

The following replace/add to the working inventory:

| Working ID | Operation | Class | Semantic owner | Note |
|---|---|---|---|---|
| TA03-MEM-02 | `PromoteMemoryCandidate` | COMMAND | C06 | replaces broad `ManageMemoryItem` |
| TA03-MEM-03 | `SupersedeMemoryItem` | COMMAND | C06 | correction through lineage, no silent overwrite |
| TA03-MEM-04 | `DeleteMemoryItem` | COMMAND | C06 | explicit deletion/non-retrieval/derived-index consequence |
| TA03-DEL-07 | `SubmitDecisionRequest` | PROPOSAL | C03 | C03 owns pending-request lifecycle, not requested decision semantics |
| TA03-DEL-08 | `ResolveDecisionRequest` | COMMAND | C03 | records resolution/ref; resulting owner changes remain separate operations |
| TA03-DEL-09 | `ReportConsumption` | CALLBACK/OBSERVATION | C03 | provider/receipt observation; canonical budget remains reconciled by C03 |
| TA03-RUN-02 | `DrainProviderRuntime` | COMMAND | A05 | replaces generic runtime-control command |
| TA03-RUN-03 | `StopProviderRuntime` | COMMAND | A05 | process stop ≠ provider request/Delegation terminal proof |
| TA03-RUN-04 | `ReconcileProviderRuntime` | COMMAND | A05 | process/incarnation reconciliation; B01 request reconciliation remains separate |

B01 working records are corrected to `semantic owner = G01`, with B01 as accepted boundary profile and explicit state/executor roles.

## 4. Re-run of coverage logic

### Consumer without operation — corrected

Material current/foundational consumers now have an explicit semantic path for:

- Presence activation/interaction/output/close;
- interpreted Intent proposal and Mission promotion;
- authoritative identity/project/authority reads;
- governed memory query/context build/candidate promotion/supersession/delete;
- provider manifest/approval/revocation/resolution;
- Delegation creation/child capability/cancellation/reconciliation/outcome/read;
- Decision Request escalation/resolution;
- budget/consumption observation;
- provider process readiness/drain/stop/reconciliation;
- B01 request/cancel/snapshot/output;
- Authority Grant/revocation/effect decision/effect/reconciliation;
- Artifact/Observation/Claim/Evidence/Receipt/Verdict;
- C12 append/read exact history.

Known later consumers remain explicitly DEFER SAFELY rather than omitted.

### Operation without consumer — remaining challenge

Before durable design, every candidate will carry a `first/real consumer` field. Any operation that cannot name one is deleted or deferred.

## 5. Section-1 decision recommendation

The self-review strengthens the original recommendation:

> **TA-03 should adopt named semantic operations grouped into non-callable families, with one semantic owner per operation and explicit state/execution roles where they differ. It should reject both storage-shaped CRUD surfaces and a universal OperationEnvelope/CommandBus. Common cross-cutting fields remain conditional semantic obligations for TA-04 to project, not one mandatory physical envelope.**

This is ready to be presented as the first operator-reviewed TA-03 design section. No durable architecture file should be created until the operator approves that section.
