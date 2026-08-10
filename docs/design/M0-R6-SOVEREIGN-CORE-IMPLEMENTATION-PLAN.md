---
id: DESIGN-AURORA-M0-R6-SOVEREIGN-CORE-IMPLEMENTATION-PLAN
title: M0 Sovereign Core Implementation Plan
document_type: implementation_plan
form: reference
authority: design
status: proposed
version: 0.1.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - exact task-by-task R6 execution plan for MIS-M0-SOVEREIGN-CORE-001 v0.1.0
related:
  - DOC-AURORA-M0-R6-MICRODESIGN-OPERATOR-ACCEPTANCE
  - DESIGN-AURORA-M0-R6-SOVEREIGN-CORE-MICRODESIGN
  - DOC-AURORA-MIS-M0-SOVEREIGN-CORE-001
  - DOC-AURORA-CAP-SOVEREIGN-CORE-TEST-PLAN
  - DOC-AURORA-CAP-SOVEREIGN-CORE-R5-COVERAGE
source_revision: 0f596602988a90205ff412fdb860e968700dbcb2
microdesign_version: 0.1.0
mission_contract_revision: 0.1.0
last_reviewed: 2026-08-09
---

# M0 Sovereign Core Implementation Plan

> **For agentic workers:** execute this plan only after explicit M0 ACRM R7 authorization. Use test-driven development task-by-task; no task or green CI run authorizes the next ACRM gate.

**Goal:** Build the smallest real Aurora Sovereign Core that preserves stable identity, Project state and current authority across process death, safely rejects invalid transitions, and exports/restores/migrates sovereign logical state with fixed-revision evidence.

**Architecture:** One Go module and one local modular process. `cmd/adapters → application → domain + ports`; SQLite and trust filesystem are adapters, not domain owners. Implementation proceeds as vertical slices so Aurora becomes executable early and is hardened progressively rather than after a long foundation phase.

**Tech stack / initial pins:** Go `1.26.5`; `modernc.org/sqlite v1.54.0` with exact compatible `modernc.org/libc v1.74.1`; `golang.org/x/crypto v0.54.0`; `filippo.io/age v1.3.1`; `github.com/santhosh-tekuri/jsonschema/v6 v6.0.3`; `github.com/gowebpki/jcs v1.0.1`; OpenTelemetry Go API/SDK `v1.44.0`; `golang.org/x/sys` pinned at the then-current supported exact version when the Windows adapter task begins. `CGO_ENABLED=0` is mandatory for the selected SQLite path.

## Global constraints

- R7 is separately authorized; this Plan itself grants no execution authority.
- One Go module, one binary, one local Core process, one embedded SQLite operational store.
- No model call, Mastra runtime, AHDK, MNFS, broker, durable workflow engine, microservice, ORM, DI container, event bus, plugin framework or generalized authorization platform in M0.
- No production code may be copied/merged from SPK-001 or SPK-002; spike evidence informs behavior only.
- SQLite uses `database/sql`, `journal_mode=WAL`, `synchronous=FULL`, `foreign_keys=ON`, `MaxOpenConns(1)`, `MaxIdleConns(1)`.
- Owner trust is a random 256-bit ORK wrapped by Argon2id-derived 32-byte KEK and AES-256-GCM; Argon2id allowed profile is 64 MiB, 3 iterations, parallelism 4 and must be bounded before allocation.
- HKDF-SHA-256 separates governing-state and trust-anchor keys; HMAC-SHA-256 authenticates one current governing logical snapshot and the tiny external trust anchor.
- Portable logical state is JSON Schema 2020-12 + JSON; JCS is used for deterministic digest/MAC input; SHA-256 is content digest; normal sensitive exports are age-protected.
- Trust artifacts are `<data-dir>/trust/owner-root.json` and `<data-dir>/trust/owner-anchor.json`; they are not a second operational database.
- CLI is a proof/control adapter only; domain/application code never prints UI text or owns terminal input.
- Tests are slice-local TDD. The accepted 84-test catalog is a final coverage obligation, not an upfront file count.
- Every task ends with a green test/build result and an independently reviewable commit. No unrelated refactor.
- Any implementation finding that requires changing accepted A2 behavior, ADR mechanism class, Mission criteria, topology or threat claim stops execution and triggers replan.

---

## Locked production file map

```text
go.mod
go.sum
cmd/aurora/main.go

internal/domain/identity/{types.go,rules.go}
internal/domain/project/{types.go,transition.go}
internal/domain/authority/{types.go,evaluate.go}
internal/domain/portability/{types.go,invariants.go}
internal/domain/evidence/types.go

internal/application/{service.go,crypto.go,initialize.go,create_project.go,transition_project.go,authority.go,inspect.go,recover.go,export.go,restore.go,migrate.go}

internal/ports/{state.go,owner_trust.go,clock.go,export_protection.go}

internal/adapters/sqlite/{store.go,schema.go,mutations.go,queries.go}
internal/adapters/trustfs/{store.go,publish.go,publish_unix.go,publish_windows.go}
internal/adapters/exportage/protection.go
internal/adapters/cli/{cli.go,commands.go,render.go,secrets.go}
internal/adapters/observability/observability.go

schemas/{project-state-v1.schema.json,sovereign-export-v1.schema.json}
migrations/sqlite/0001_initial.sql

testdata/{fixtures,adversarial,migration,golden}/
tests/process/
tests/golden/
```

Test files live beside the package under test unless they require external process orchestration, in which case they live under `tests/process` or `tests/golden`.

---

### TASK-00 — Toolchain lock, module skeleton and architecture guards

**Mission coverage:** CRIT-010.  
**Primary requirements:** REQ-096..107.  
**Accepted tests:** `T-ID-005`, `T-SCOPE-001`, `T-SCOPE-002`, `T-ARCH-001..006`, `T-DOC-003`.

**Files:** create `go.mod`, `cmd/aurora/main.go`, package directories above, `internal/adapters/cli/cli.go`, `internal/adapters/cli/cli_test.go`, `internal/ports/architecture_test.go`.

**Produces:** module path `github.com/developmentconexus-ops/aurora_project`; `cli.Run(args []string, out, errOut io.Writer) int`; package dependency direction.

- [ ] **Step 1 — RED:** create `cli_test.go` asserting `Run([]string{"--help"}, ...) == 0` and output contains `Aurora Sovereign Core`; create `architecture_test.go` that parses imports under `internal/domain/**` and fails on imports containing `/adapters/`, `modernc.org/sqlite`, `go.opentelemetry.io`, `filippo.io/age` or OS-specific trust code.
- [ ] **Step 2 — verify RED:** run `go test ./...`; expected failure because module/CLI do not yet exist.
- [ ] **Step 3 — pin baseline:** initialize `go.mod` with Go `1.26.5`; pin current R6-validated dependencies. Run `go mod tidy`. Verify `go list -m all` contains `modernc.org/sqlite v1.54.0` and `modernc.org/libc v1.74.1`; fail if modernc resolves a different libc pin.
- [ ] **Step 4 — minimal GREEN:** implement `cmd/aurora/main.go` as `os.Exit(cli.Run(os.Args[1:], os.Stdout, os.Stderr))`; implement only `--help` and unknown-command failure in `cli.Run`.
- [ ] **Step 5 — verify:** `CGO_ENABLED=0 go test ./... && CGO_ENABLED=0 go vet ./... && CGO_ENABLED=0 go build ./cmd/aurora`; run built binary `--help`.
- [ ] **Step 6 — commit:** `git commit -m "build(core): establish M0 Go production skeleton"`.

---

### TASK-01 — Sovereign bootstrap: real ORK, SQLite, trust files, init/status

**Mission coverage:** CRIT-001, CRIT-007, partial CRIT-010.  
**Primary requirements:** REQ-001..003 plus bootstrap/security portions of REQ-067..076.  
**Accepted tests:** `T-ID-001..003`, `T-SEC-003..005` where applicable.

**Files:** create identity types/rules, `ports/state.go`, `ports/owner_trust.go`, `ports/clock.go`, `application/{service.go,crypto.go,initialize.go,inspect.go}`, SQLite store/schema/migration, trustfs store/publication files, CLI secrets/commands/render, tests for each.

**Core signatures:**

```go
type Service struct {
    State ports.StateStore
    Trust ports.OwnerTrustStore
    Clock ports.Clock
}

func (s *Service) Initialize(ctx context.Context, passphrase []byte) (InitializeResult, error)
func (s *Service) Inspect(ctx context.Context, passphrase []byte) (InspectResult, error)
```

`crypto.go` exposes package-private `newRootEnvelope`, `unlockORK`, `derivePurposeKey`, `governingMAC`, `anchorMAC`; plaintext ORK never enters persistence structs.

- [ ] **Step 1 — RED domain/crypto:** tests assert IDs are random stable prefixed IDs; root creation produces a 32-byte ORK only in memory; wrong passphrase fails; unsupported/extreme KDF params fail before Argon2 call; rewrapping preserves ORK bytes.
- [ ] **Step 2 — RED store:** SQLite integration test opens temp data dir, applies `0001_initial.sql`, verifies WAL/FULL/FK and single-connection posture, and proves bootstrap creates one `core_state` row + authority revision + governing record transactionally.
- [ ] **Step 3 — RED trustfs:** Unix/Windows-targeted tests verify temp-write + sync + replace behavior; no partial JSON is accepted. Build Windows package with `GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go test ./internal/adapters/trustfs` at CI time.
- [ ] **Step 4 — implement crypto:** use `crypto/rand`, `argon2.IDKey`, AES-GCM, HKDF-SHA-256 and HMAC-SHA-256 exactly as accepted; root JSON encodes binary fields with base64url and validates version/KDF allowlist before allocation.
- [ ] **Step 5 — implement bootstrap ordering:** create/publish root envelope first; then SQLite bootstrap with initial owner authority and generation `1`; then publish authenticated anchor generation `1`; incomplete bootstrap is explicitly classified, never silently reinitialized.
- [ ] **Step 6 — CLI:** `aurora init [--data-dir]` reads owner passphrase through adapter-local `SecretReader`; `aurora status` unlocks and verifies governing state, returning structured result rendered as text or `--json`.
- [ ] **Step 7 — verify journey:** with a temp data dir, run real binary `init`; terminate; run a second binary invocation `status`; assert identical `aurora_id`. Running `init` again must fail without replacing identity.
- [ ] **Step 8 — verify hygiene:** grep test logs/artifacts for fixture passphrase and ORK bytes; expected zero matches.
- [ ] **Step 9 — commit:** `git commit -m "feat(core): initialize sovereign Aurora identity"`.

---

### TASK-02 — Project identity and restart continuity

**Mission coverage:** CRIT-001, early CRIT-002.  
**Primary requirements:** REQ-004..009.  
**Accepted tests:** `T-ID-004`, `T-STATE-004`, scope/static checks.

**Files:** `domain/project/types.go`, `application/create_project.go`, state-port mutation types, SQLite project mutation/query, CLI project create/show, tests.

**Signatures:**

```go
type CreateProjectInput struct { DisplayLabel, ObjectiveSummary string }
func (s *Service) CreateProject(ctx context.Context, ownerPassphrase []byte, in CreateProjectInput) (project.Project, error)
func (s *Service) ShowProject(ctx context.Context, ownerPassphrase []byte, id project.ProjectID) (ProjectView, error)
```

- [ ] Write failing tests proving generated `ProjectID` is stable and not derived from label/path; duplicate request cannot silently replace an existing Project.
- [ ] Run focused tests; verify RED.
- [ ] Implement project domain type/rules, SQLite insert/query and application orchestration through `StateStore` only.
- [ ] Add CLI `project create --label ... --objective ...` and `project show --project ...`; both support `--json`.
- [ ] Run integration journey: create Project → terminate process → show Project from fresh process; same ID/metadata.
- [ ] Run `go test ./...`, `go vet ./...`, architecture guard.
- [ ] Commit `feat(project): persist stable project identity`.

---

### TASK-03 — Canonical accepted-state ownership and inspection

**Mission coverage:** CRIT-002.  
**Primary requirements:** REQ-010..020.  
**Accepted tests:** `T-STATE-001..006`.

**Files:** `domain/project/types.go`, `domain/project/transition.go`, `application/transition_project.go`, `application/inspect.go`, SQLite revision queries/mutations, `schemas/project-state-v1.schema.json`, tests.

- [ ] RED: schema/unit tests define `StateEnvelope{SchemaVersion,Kind,Summary,Payload}` and prove payload content resembling identity/authority remains opaque data.
- [ ] RED: integration test with broken `projects.current_state_revision` must return degraded/error; history may not be auto-promoted.
- [ ] Implement state envelope validation with the pinned JSON Schema validator and immutable revision loading.
- [ ] Implement `InspectProject` projection from current Project + current revision; no cached projection is authoritative after source revision changes.
- [ ] Verify all `T-STATE-*` cases and that DB/CLI-specific types do not leak into domain types.
- [ ] Commit `feat(state): establish canonical project state ownership`.

---

### TASK-04 — Revision-bound transitions and invalid-transition containment

**Mission coverage:** CRIT-003.  
**Primary requirements:** REQ-021..030.  
**Accepted tests:** `T-TRANS-001..008`, `T-EVID-001` partially.

**Files:** `domain/project/transition.go`, `application/transition_project.go`, `ports/state.go`, SQLite mutations, CLI `project set-state`, tests.

**Signature:**

```go
type TransitionProjectInput struct {
    AttemptID string
    ProjectID project.ProjectID
    ExpectedRevision *project.StateRevision
    State project.StateEnvelope
    ProposedNextAction *project.ActionDescriptor
}
```

- [ ] RED: valid transition creates exactly one revision and advances pointer; stale expected revision rejects; malformed envelope rejects; state payload attempting to mutate stable IDs rejects; rejected/failed attempts create no accepted revision.
- [ ] Verify RED.
- [ ] Implement validation order: identity/project → schema → expected revision → authority call boundary → immutable ownership → commitability. Re-read expected revision inside SQLite transaction before commit.
- [ ] Governing commit writes new revision + current pointer + required audit/evidence metadata + generation/HMAC in one transaction; rejected attempt uses `RecordNonGoverning` and does not advance generation/anchor.
- [ ] CLI accepts `--project`, `--expected <n|none>`, `--kind`, `--summary`, optional `--payload-json`, optional next-action fields.
- [ ] Verify `T-TRANS-001..008`, then run real binary R1→R2→stale-R1 rejection and inspect that R2 remains current.
- [ ] Commit `feat(state): enforce revision-bound project transitions`.

---

### TASK-05 — Minimum M0 authority and next-safe-action semantics

**Mission coverage:** CRIT-004.  
**Primary requirements:** REQ-032..045.  
**Accepted tests:** `T-AUTH-001..011` except trust-specific time rollback in TASK-06.

**Files:** authority domain types/evaluator, `application/authority.go`, SQLite authority revision mutation/query, CLI authority commands, tests.

**Pure API:**

```go
func Evaluate(state State, projectID project.ProjectID, action *project.ActionDescriptor, now time.Time) Snapshot
func NextSafeAction(state State, rev project.ProjectStateRevision, now time.Time) NextSafeActionProjection
```

- [ ] RED pure tests for active matching grant → `VALID/PERMITTED`; wrong scope → blocked; expired → `EXPIRED`; revoked/superseded → non-permitting; Project content cannot create permission.
- [ ] RED persistence tests prove each authority change creates an immutable complete authority revision and changes `core_state.current_authority_revision`; restart preserves revocation.
- [ ] Implement minimal `Grant`, `State`, snapshot and projection; no generalized policy engine.
- [ ] Implement owner-only `Grant`, `Revoke`, `Show`; stale authority revision fails closed.
- [ ] CLI: `authority grant --project ... --action ... [--valid-until RFC3339]`, `authority revoke --authority ...`, `authority show`.
- [ ] Run `T-AUTH-001..011` plus real restart after revoke.
- [ ] Commit `feat(authority): add bounded M0 authority evaluation`.

---

### TASK-06 — Owner trust classifications, governing MAC and rollback/time containment

**Mission coverage:** CRIT-004, CRIT-007, CRIT-009.  
**Primary requirements:** REQ-067..076 plus trust-specific portions of REQ-036, 041..045.  
**Accepted tests:** SPK-002 S01–S12 equivalents, `T-AUTH-008`, `T-AUTH-012`, `T-SEC-003..006`.

**Files:** `application/crypto.go`, trust preflight/reconcile logic in `application/recover.go` and `authority.go`, SQLite current-snapshot builder, trustfs tests, adversarial fixtures.

- [ ] RED: construct current governing logical snapshot, JCS-canonicalize, HMAC once; altering current state content while retaining revision number must yield `INVALID_DB_MAC`.
- [ ] RED classification table: equal generation/valid MACs → `NORMAL`; DB ahead + valid DB MAC → `ANCHOR_LAG`; DB behind → `STATE_ROLLBACK`; bad DB MAC → `INVALID_DB_MAC`; bad anchor → `INVALID_ANCHOR_MAC`; absent anchor outside restore → `MISSING_ANCHOR`; clock below authenticated high-water → `TIME_UNTRUSTED`; restored state marker → `REVALIDATION_REQUIRED`.
- [ ] Implement preflight before every governing mutation. `ANCHOR_LAG` is non-permitting and only authenticated owner reconciliation may publish anchor forward; no speculative `+1` restriction is added.
- [ ] Implement only the accepted observed-wall-time high-water semantics needed for S07/T-AUTH-012; do not invent a stronger policy during implementation. If reproducing accepted test behavior requires changing the Microdesign, stop/replan.
- [ ] Verify passphrase rotation rewraps same ORK and all secret-hygiene checks remain clean.
- [ ] Run Linux tests and Windows trustfs/classification build/test matrix under `CGO_ENABLED=0`.
- [ ] Commit `feat(trust): enforce owner-root rollback and restore safeguards`.

---

### TASK-07 — Fresh-process recovery and real process-kill continuity

**Mission coverage:** CRIT-005.  
**Primary requirements:** REQ-046..055.  
**Accepted tests:** `T-REC-001..007`, `T-REL-001..003` partially.

**Files:** `application/recover.go`, SQLite recovery queries, `tests/process/restart_test.go`, CLI status rendering.

- [ ] RED application tests for missing/corrupt canonical state → explicit degraded/failed result; stale projections discarded; failure taxonomy distinguishes store, authority/trust and version classes.
- [ ] RED process test builds `./cmd/aurora`, initializes state, creates Project/state/authority, exits all processes, starts a new OS process and compares IDs/revisions/results.
- [ ] Implement `RecoverCurrentState` from SQLite + trust artifacts + current clock only; no transcript/model/session data.
- [ ] Add structured `RecoveryResult` with recovered IDs/revisions, trust/authority classification and limitations.
- [ ] Verify all `T-REC-*`; process test must fail if it can succeed using in-memory same-process state.
- [ ] Commit `feat(recovery): recover sovereign state in a fresh process`.

---

### TASK-08 — Portable export and staged restore

**Mission coverage:** CRIT-006, CRIT-007.  
**Primary requirements:** REQ-056..063.  
**Accepted tests:** `T-PORT-001..008`, `T-PORT-012..013`.

**Files:** portability domain types/invariants, `ports/export_protection.go`, `adapters/exportage/protection.go`, `application/export.go`, `application/restore.go`, export schema, CLI export/restore, fixtures/tests.

**Logical format:** one `aurora-sovereign-export` JSON document v1. Digest is SHA-256 over JCS-canonical document with top-level `integrity` omitted; complete JSON is age passphrase protected.

- [ ] RED schema/digest tests: valid export validates; a byte/semantic alteration fails digest/schema before apply; unsupported version fails explicitly.
- [ ] Implement `ExportProtection` with `filippo.io/age` scrypt/passphrase API; export secret is separate from owner passphrase and only adapter/test injection sees it.
- [ ] Implement export from canonical domain snapshot, including version metadata, records and wrapped root recovery envelope but excluding current trust high-water as freshness authority.
- [ ] RED restore tests: fresh valid target succeeds to `REVALIDATION_REQUIRED`; identity collision leaves target unchanged; pre-revocation historical export cannot yield `VALID`; non-owner revalidation denied.
- [ ] Implement staged restore under a fresh temporary target data directory. Validate/decrypt/schema/digest/version/collision before publishing. Build fresh SQLite; publish recovered root; do not import historical anchor; owner revalidation creates a new authority revision/generation/anchor.
- [ ] CLI `export --out ...` and `restore --in ... --target-data-dir ...`; secrets are non-echo prompts; `--json` returns structured results only.
- [ ] Verify `T-PORT-001..008`, `012`, `013` and run real binary export→fresh-directory restore journey.
- [ ] Commit `feat(portability): export and safely restore sovereign state`.

---

### TASK-09 — Explicit logical and physical migration

**Mission coverage:** CRIT-006.  
**Primary requirements:** REQ-064..066.  
**Accepted tests:** `T-PORT-003`, `T-PORT-009..011`, `T-PORT-010` with F-15.

**Files:** `application/migrate.go`, portability invariants, `testdata/migration/v0-valid.json`, `v0-semantic-mutation.json`, migration tests; future SQL migration runner only enough to recognize/apply numbered SQLite files.

- [ ] RED: deterministic compatibility fixture v0→v1 preserves Aurora ID, Project IDs/current state meaning, authority semantics, provenance/evidence references.
- [ ] RED: semantic-mutation fixture changing stable identity/authority meaning is rejected; unsupported logical versions fail without coercion.
- [ ] Implement one explicit `migrateV0ToV1` compatibility transformation used only because the accepted Test Plan requires migration proof; no registry/graph/plugin framework.
- [ ] Implement physical migration runner that reads current schema version and applies explicit numbered SQL files transactionally; only `0001_initial.sql` exists initially.
- [ ] CLI `migrate --in ... --out ...` operates on portable logical documents and never mutates current store implicitly.
- [ ] Verify valid, invalid and interrupted/ambiguous migration behavior; retry only after reconciliation.
- [ ] Commit `feat(migration): verify explicit sovereign state migration`.

---

### TASK-10 — Audit/evidence metadata and minimal observability

**Mission coverage:** CRIT-008.  
**Primary requirements:** REQ-077..088.  
**Accepted tests:** `T-EVID-001..008`, `T-SEC-003`, F-13.

**Files:** evidence domain types, record helpers, observability adapter, application operation correlation, tests.

- [ ] RED: accepted/rejected operations produce attributable records with stable `operation_id`, Aurora/Project/revision refs, outcome/reason; EvidenceRecord includes criterion/test method/environment/revision/limitations fields.
- [ ] RED: telemetry sink absent/failing cannot alter Core operation result; sensitive payload/passphrase/ORK/export secret never appears in logs/traces.
- [ ] Implement `slog` structured fields and OTel trace/metric instrumentation with no required exporter. Initial metrics only operation total/failure total unless a current test requires another signal.
- [ ] Preserve Claim/Receipt/Evidence/Verdict as distinct domain types even though `records` co-locates physical metadata.
- [ ] Verify all `T-EVID-*`, telemetry failure fixture and secret redaction scan.
- [ ] Commit `feat(evidence): add attributable M0 proof and observability metadata`.

---

### TASK-11 — Complete adversarial/fault/security matrix

**Mission coverage:** CRIT-007, CRIT-009.  
**Primary requirements:** REQ-089..095 and secondary security allocations.  
**Accepted tests:** all remaining `FAULT_INJECTION`/`SECURITY_TEST` cases and F-08..F-15.

**Files:** `testdata/adversarial/**`, purpose-specific test hooks, `tests/process/faults_test.go`, no production chaos framework.

- [ ] Add deterministic cases for stale revision, malformed envelope, immutable identity mutation, wrong scope, expired/revoked/missing/corrupt authority, DB/anchor corruption/rollback, time rollback, store unavailable, pre-revocation restore, corrupt export, unsupported version, identity collision, migration semantic mutation, untrusted-content injection and unsafe retry.
- [ ] Add external process-kill hooks around the material SQLite-commit / anchor-publication boundary; use OS process termination, not an in-memory fake, for the crash proof.
- [ ] For each ambiguous mutation/restore/migration case, assert reconciliation precedes retry and duplicate success cannot be reported.
- [ ] Run complete deterministic catalog; expected 100% pass and zero unauthorized accepted mutation.
- [ ] Run Ubuntu and Windows CI classes required by Contract. Document that hosted process-kill is not a physical power-loss/storage-controller claim.
- [ ] Commit `test(m0): complete sovereign core fault and security matrix`.

---

### TASK-12 — Real-binary M0 Golden Proof

**Mission coverage:** CRIT-012 plus composition of CRIT-001..011.  
**Primary requirement:** REQ-031.  
**Accepted tests:** `J-001..J-004`, `T-REL-004`, operator-visible Golden Proof.

**Files:** `tests/golden/m0_test.go`, `testdata/golden/**`, optional `scripts/run-m0-golden-proof.*` only if the process test cannot remain portable in Go.

- [ ] RED: end-to-end runner must fail against any missing slice. It builds the real binary and uses a fresh temp directory.
- [ ] Execute: initialize Aurora → create Project → record accepted state + proposed action → establish/evaluate authority → terminate every Aurora process → fresh process recovers same IDs/state → stale transition rejected with zero governing mutation → export → fresh restore → `REVALIDATION_REQUIRED` → non-owner revalidation denied → owner revalidates → new authority revision/current next action.
- [ ] Verify real SQLite files/trust files/process boundaries are used. No direct calls to application internals may substitute for the operator-visible journey.
- [ ] Record exact binary/source revision, Go/dependency lock, OS, fixture IDs and limitations in machine-readable test output.
- [ ] Run `CGO_ENABLED=0 go test ./...`, `go vet ./...`, Golden Proof, docs validator and cross-platform required jobs.
- [ ] Commit `test(m0): prove Sovereign Core walking skeleton end to end`.

---

### TASK-13 — Traceability, documentation and R7 evidence packaging

**Mission coverage:** CRIT-011.  
**Primary requirements:** REQ-108..122.  
**Accepted tests:** `T-DOC-001..007`, `T-REL-003..004`.

**Files:** update implementation references under `docs/`, create R7 evidence receipts/review artifacts only when R7 is authorized and executed; keep `STATUS.md` truthful.

- [ ] Verify every durable code path has an owning Requirement/ADR/Contract/task and every `CAP-SOVEREIGN-CORE-REQ-001..122` has implementation/evidence allocation in `R6-IMPLEMENTATION-COVERAGE.md`.
- [ ] Verify no `TODO`, fake success, hardcoded production secret, mock-as-proof, M1+ capability, Mastra/AHDK/MNFS code or spike production copy appears.
- [ ] Record dependency license/maintenance receipt and exact pins actually used.
- [ ] Generate fixed-revision R7 evidence index linking tests, environment, artifacts, uncertainty and limitations; Claim is not Verdict.
- [ ] Run documentation generation/validator and `git diff --check`.
- [ ] Commit `docs: record M0 Sovereign Core implementation evidence`.

---

## Task dependency graph

```text
TASK-00
  ↓
TASK-01 → TASK-02 → TASK-03 → TASK-04 → TASK-05 → TASK-06 → TASK-07
                                                      ↓
                                                   TASK-08 → TASK-09
                                                      ↓        ↓
                                                   TASK-10 → TASK-11
                                                               ↓
                                                           TASK-12
                                                               ↓
                                                           TASK-13
```

Parallelism is intentionally limited for the first implementation because the source tree and state/trust contracts are new. After TASK-03, pure authority-domain tests and portability schema fixtures may be prepared in parallel only if they do not modify shared ports/application contracts.

## Verification commands required throughout R7

```bash
CGO_ENABLED=0 go test ./...
CGO_ENABLED=0 go vet ./...
CGO_ENABLED=0 go build ./cmd/aurora
gofmt -w <changed-go-files>
git diff --check
python scripts/generate_docs.py --check
python scripts/validate_docs.py
```

Cross-platform CI additionally builds/tests the trust filesystem and integrated required journeys on Windows amd64. Production success claims require the full target command, not an adjacent green check.

## Commit/review policy

- Each TASK is one independently reviewable green commit or a very small ordered commit set when OS-specific publication requires separation.
- Review after every TASK checks only that task's listed requirements/tests plus global constraints.
- Do not merge a task that weakens a negative case to get GREEN.
- A material finding pauses the sequence and records `REPLAN_REQUIRED`; later tasks are not a workaround.
- R8 Product Milestone closeout remains separate even after TASK-12/13 are green.

## R6 completion condition for this Plan

R6 can pass only when adversarial review confirms:

1. every task has exact files/interfaces/tests/verification;
2. all 122 accepted requirements have a primary implementation task;
3. all 12 Mission criteria have task coverage;
4. no material product/architecture choice is deferred to R7 implementers;
5. dependency pins/mechanism classes are reproducible enough to start;
6. vertical order produces real Aurora behavior early;
7. negative/fault/security obligations are explicit;
8. no future Presence/Mastra/M1+ work is prebuilt;
9. R7 remains separately unauthorized until operator action.
