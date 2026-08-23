# RM-10 — Independent Fable Review

Repository: `developmentconexus-ops/aurora_project`

## Fixed review identity

```text
candidate branch: docs/mr-01-repository-migration-20260823
candidate HEAD: eb402577c3cda638102408543c412b70c72b18d4
candidate Documentation: 32663114183 — SUCCESS
review branch: review/mr-01-migration-fable
review delta allowed: docs/work/current/ai-dialog.md only
```

This is independent review Evidence only. Do not modify candidate/durable files on this branch. Do not merge this review PR.

## Mandatory external authorities

Read and apply:

```text
developmentconexus-ops/conexus-methodology/METHOD.md
DevelopmentConexus Engineering Method v1.0.0

developmentconexus-ops/conexus-methodology/REPOSITORY-STANDARD.md
Repository Standard v1.0.0
```

Reviewer Evidence is not Product, architecture or repository authority. Lead adjudicates findings.

## Bounded Aurora review pack

Start with:

```text
AGENTS.md
docs/index.md
docs/roadmap.md
docs/decisions/methodology-repository-rebaseline.md
docs/development/planning-readiness.md
docs/development/engineering-rules.md
docs/development/capability-realization.md
docs/product/blueprint/15-documentation-research-governance.md
docs/architecture/index.md
docs/decisions/index.md
docs/evidence/mr-01-platform-enforcement.md
docs/evidence/mr-01-migration-coherence.md
docs/evidence/mr-01-migration-census.md
scripts/generate_docs.py
scripts/validate_docs.py
.github/workflows/docs.yml
```

Expand only when a concrete claim requires the owning source/Evidence.

## Context to preserve

```text
A0 Product constitution: ACCEPTED / MERGED
System Architecture Rebaseline: ACCEPTED / MERGED
TA-01: ACCEPTED / CANONICAL
TA-02: ACCEPTED / CANONICAL
M0 R7 candidate: FROZEN / PRESERVED / NON-CANONICAL
M0 R7 Verdict: NOT ISSUED
M0 R8: NOT AUTHORIZED
MR-01 semantic target: OPERATOR-RATIFIED
RM-01..RM-09: PASS
TA-03+: NOT AUTHORIZED
Product/runtime implementation: BLOCKED
merge: NOT AUTHORIZED
```

MR-01 intentionally changed repository/methodology/readiness operating structure without reopening Product meaning or TA-01/TA-02.

## Adversarial questions

Challenge the final migrated candidate, not the pre-migration proposal.

1. Is there exactly one mutable repository-program status/permission/next-action authority, and is it really `docs/roadmap.md`?
2. Can a fresh actor reconstruct current program/gate/prohibitions/next action through `AGENTS → docs/index → docs/roadmap → 1–2 owners` without loading stale history?
3. Did the migration lose any current Product, architecture, decision, acceptance, frozen-M0 or provenance semantics while retiring old live surfaces?
4. Are Blueprint 15 changes truly bounded to repository/documentation governance rather than hidden Product change?
5. Are Planning Readiness and ACRM now non-overlapping: global cross-system prerequisites versus capability/milestone/mission R0–R8 realization?
6. Does the TA-03→TA-13 graph remain dependency-driven and proportionate, with TA-TX conditional, or did migration turn it into ceremony?
7. Do TA-01/TA-02 remain semantically preserved and canonical without being silently rewritten by the new stage graph?
8. Are current Conexus OS references correct while historical MNFS provenance remains honest rather than mechanically rewritten?
9. Are the moved ADRs, architecture docs, capability docs, phase snapshots and Evidence discoverable with stable identity and without duplicate current authority?
10. Is `docs/roadmap.md` demonstrably isolated from Product Blueprint generation, while Blueprint 14 remains Product capability-roadmap authority?
11. Are `scripts/generate_docs.py`, `scripts/validate_docs.py` and `.github/workflows/docs.yml` mutually coherent with the target tree?
12. Are material repository guards falsifiable with deterministic negative controls rather than only positive CI?
13. Is final-candidate cleanliness real: no `docs/work/**`, `docs/superpowers/**`, legacy tracking/design/acceptance/reviews/adr surfaces in the candidate being reviewed?
14. Does RM-08 Evidence justify platform-enforcement closure without pretending the connected API can enumerate Ruleset internals it cannot observe?
15. Do repository merge settings and protected-main state support the accepted Git lifecycle without introducing unnecessary human-approval deadlock?
16. Did any migration step accidentally select a framework, database, IAM, model/provider, source topology or Product/runtime mechanism that belongs to TA-03+?
17. Does the migration preserve the law that implementation remains blocked until TA-13 plus a separate explicit operator Product execution grant?
18. Is there any hidden dependency on branch-only review Evidence, PR comments, old status files or unmerged branches for current semantic reconstruction?
19. Are historical snapshots clearly non-current so their old TA numbering/status language cannot outrank current owners?
20. What is the strongest credible alternative repository/readiness model now, and is the accepted candidate still the Global Maximum under current constraints rather than overbuilt?

## Required finding format

For every finding:

```text
ID: RM-Ixx
severity: BLOCKING | MATERIAL | MODERATE | MINOR
claim challenged:
evidence / exact authority:
why it matters:
recommended disposition:
  CORRECT_CANDIDATE | REOPEN_OWNER | ACCEPT_TRADEOFF | DEFER_SAFELY | NO_CHANGE
scope correction:
review coverage impact:
```

Also provide:

- strongest counterargument to the migrated model;
- best credible alternative;
- YAGNI / essential-vs-accidental complexity assessment;
- Repository Standard conformance assessment;
- whether TA-01 or TA-02 must reopen;
- whether a second independent review round is required if corrections are made.

## Final verdict

End with exactly one:

```text
CONVERGED
NOT_CONVERGED
STOP_PREREQUISITE
```

And separately state:

```text
PASS
PASS_WITH_FINDINGS
FAIL
```

Do not authorize merge. Do not authorize TA-03+. Do not implement anything. Write reviewer Evidence only in this file.

---

# Reviewer output — RM-10 independent Fable review

**Independent reviewer:** Claude Fable 5 (same independent challenger as Round 1 semantic review on PR #7; no involvement in candidate authorship, migration execution, or Lead adjudication).
**Reviewed candidate HEAD:** `eb402577c3cda638102408543c412b70c72b18d4` (branch `docs/mr-01-repository-migration-20260823`).
**Review branch HEAD at review start:** `9e638c892c2f29535afe9115def6829e43b4f9dc`.
**Isolation verified twice:** `git diff --stat eb402577..HEAD` = `docs/work/current/ai-dialog.md` only; and mechanically via the candidate's own guard: `python scripts/validate_docs.py --review-base eb402577…` → `Review isolation PASS`.
**Authority reconstruction:** performed from the migrated repository plus METHOD.md v1.0.0 and REPOSITORY-STANDARD.md v1.0.0 only.

**Live verification executed by the reviewer on this exact revision (not taken on faith from prior runs):**

```text
python scripts/validate_docs.py --self-test                      → PASS
python scripts/validate_docs.py --review-base eb402577…          → Review isolation PASS
python scripts/generate_docs.py --output-root <temp> + full
  validate_docs.py run against fresh generation                  → FAIL: "stale generated Product Blueprint"
  (root-caused below as RM-I01 — an environment-dependence defect
   in the generated projection, not actual semantic staleness)
router link targets (research/RESEARCH-MAP.md, reference/
  architecture-spikes.md, history/…, phases/m0/, evidence/…)     → all exist
frozen M0 R7 ref: origin feat/m0-r7-sovereign-core-20260810
  = 7ec999b093… matches revision cited by current architecture   → reachable
bootstrap budget: AGENTS 2948 + index 3021 + roadmap 4429
  = 10 398 bytes ≤ 20 480                                        → PASS
MNFS scan over current authorities: only explicit
  "Conexus OS (historically MNFS)" disambiguations remain        → PASS
candidate tree: no docs/work, docs/superpowers, docs/tracking,
  docs/adr, docs/design, docs/acceptance, docs/reviews,
  DOCUMENTATION-MAP at eb402577                                  → PASS
```

## Finding register

```text
ID: RM-I01
severity: MATERIAL
claim challenged: The Product Blueprint aggregate is a reproducible GENERATED projection and its
  staleness guard proves freshness (decision §8 GENERATED class: "reproducible projection from
  accepted machine-readable authority"; Blueprint 15.22/15.24; engineering-rules §13 fresh
  verification on the exact target revision).
evidence / exact authority: The committed docs/product/PRODUCT-BLUEPRINT.md embeds the CI runner's
  ABSOLUTE filesystem paths 45 times: the "Canonical source" manifest column and every
  BEGIN/END SOURCE marker read "/home/runner/work/aurora_project/aurora_project/docs/product/
  blueprint/…". Cause: scripts/generate_docs.py canonical_sources() returns absolute paths and
  generate_blueprint() emits path.as_posix() without making them repo-relative. Additionally the
  manifest SHA-256 is computed over raw checkout bytes, so with core.autocrlf=true the digests
  differ from the committed ones (verified: source 01 committed hash b80f328a… = LF-bytes hash;
  CRLF-bytes hash 0829aa6f…). Reviewer reproduction on this exact revision: regeneration on a
  non-runner machine produces a projection differing in 45 lines (all path rows/markers) and the
  full validator fails with "stale generated Product Blueprint" despite zero semantic drift.
why it matters: (a) The GENERATED-class law is violated in the constitution's own flagship
  projection — reproduction depends on the generating machine's directory layout and line-ending
  configuration, which is exactly the "hand-owned divergent semantics"/environment-coupling class
  the ownership law forbids. (b) The staleness negative-control is only meaningful on a
  byte-identical runner layout; every local/operator regeneration on Windows or any non-runner
  path produces a false STALE failure, training actors to distrust or bypass the guard
  (engineering-rules §1 forbids weakening guards to go green — this defect invites it).
  (c) Runner-absolute paths in a durable constitutional publication are wrong provenance: the
  canonical source identity is the repo-relative path, not a CI workspace path.
recommended disposition: CORRECT_CANDIDATE
scope correction: In generate_docs.py, emit repo-relative paths (path.relative_to(root)
  .as_posix()) in both the manifest table and the BEGIN/END SOURCE markers, and compute SHA-256
  over newline-normalized content (e.g. raw.replace(CRLF, LF)) with the normalization stated in
  the manifest header; regenerate the aggregate once; optionally add a self-test asserting the
  generated output contains no absolute path. No semantic source changes.
review coverage impact: none — the fix is mechanical in one generator + one regenerated
  projection; all semantic review of the Blueprint sources remains valid. No second round needed.
```

```text
ID: RM-I02
severity: MODERATE
claim challenged: "durable current documents are reachable from docs/index.md" and "current
  relative links resolve" are mechanically checked (engineering-rules §11 bootstrap/authority
  controls; RM-07 PASS claim).
evidence / exact authority: scripts/validate_docs.py validate_links() current_roots list covers
  README/AGENTS/CONTRIBUTING and docs/{product,architecture,decisions,development,capabilities,
  reference} — but NOT docs/index.md and NOT docs/roadmap.md. The router's own links (roadmap.md,
  research/RESEARCH-MAP.md, history/2026-08-05-…, phases/m0/, evidence/…) are never resolved;
  validate_router() only checks that route substrings appear in the index text. The two most
  important current files in the repository are the only current files whose links are
  unvalidated. (Reviewer manually confirmed all current router targets exist today.)
why it matters: A future rename/move of any router target (e.g. RESEARCH-MAP or a phase snapshot)
  would break the fresh-actor route while CI stays green — precisely the property class the
  migration claims is now guarded.
recommended disposition: CORRECT_CANDIDATE
scope correction: Add "docs/index.md" and "docs/roadmap.md" to current_roots in validate_links()
  (two lines). Optionally validate roadmap/index links in a dedicated router check.
review coverage impact: none.
```

```text
ID: RM-I03
severity: MODERATE
claim challenged: "every material behavioral guard has a deterministic negative control or
  equivalent falsifier" (engineering-rules §11 guard quality; RM-07 "deterministic negative
  controls" PASS; roadmap RM-07 row).
evidence / exact authority: self_test() proves falsifiability for: bootstrap budget (±), review
  isolation (±), final-tree (−), blocked-root allowlist (−), generator-roadmap isolation. It does
  NOT include a negative fixture for: (a) guard_program_owner — the sole-mutable-status-authority
  uniqueness guard, i.e. THE core duplicate-status control of the whole migration (no fixture
  with a second file carrying program_status_authority: true, nor with the flag removed from
  roadmap); (b) validate_generated staleness (no perturbed-byte fixture); (c) validate_router
  route presence. These are material behavioral guards under the candidate's own definition.
why it matters: The guard the entire status-model cutover depends on has never been demonstrated
  to fire. Per METHOD.md ("a control that cannot be shown to fire is not proven") the RM-07
  negative-control claim is overstated for exactly the highest-value control.
recommended disposition: CORRECT_CANDIDATE
scope correction: Extend self_test() with: temp tree containing two program_status_authority
  files → guard must fail; roadmap without the flag → must fail; perturbed generated aggregate →
  staleness must fail; index missing one route → router check must fail. ~20 lines, no CI
  structure change.
review coverage impact: none.
```

```text
ID: RM-I04
severity: MINOR
claim challenged: The RM-10 review channel and CI operate the review-isolation control end to end
  (this file's "Bounded Aurora review pack"; engineering-rules §11 "review-branch isolation is
  mechanically checked against the exact candidate").
evidence / exact authority: (a) The review pack in this channel names
  docs/evidence/mr-01-migration-coherence.md, which does not exist; the actual file is
  docs/evidence/mr-01-fresh-actor-coherence.md (id DOC-AURORA-MR-01-RM09-COHERENCE). A fresh
  reviewer following the pack hits a dead reference. (b) .github/workflows/docs.yml never invokes
  validate_docs.py --review-base on review/** branches or review PRs, although the capability
  exists and the PR base ref is available in the pull_request event; isolation for this review
  was proven by manual invocation (reviewer ran it: PASS). The guard is "capable of firing"
  (Standard §9 satisfied) but not wired into the automated path.
why it matters: (a) is a channel-accuracy defect in temporary review material; (b) leaves the
  isolation property dependent on someone remembering the manual command each review round.
recommended disposition: CORRECT_CANDIDATE
scope correction: Fix the filename in this temporary channel (or note it in adjudication —
  ai-dialog dies before merge anyway). Optionally add a docs.yml step: on pull_request events for
  review/** head branches, run validate_docs.py --review-base "origin/$GITHUB_BASE_REF".
review coverage impact: none.
```

```text
ID: RM-I05
severity: MINOR
claim challenged: Historical snapshots are clearly non-current so their status language cannot
  outrank current owners (adversarial question 19; census "archived pre-MR-01 snapshot").
evidence / exact authority: docs/phases/pre-mr01-status-snapshot.md carries status: archived and
  authority: evidence (correct) but retains frontmatter source_of_truth_for: "current Aurora
  project phase / current authorization boundary / current blockers and immediate next action".
  docs/evidence/project-worklog.md similarly retains source_of_truth_for on an archived surface.
  The words "source of truth for current …" on a non-current snapshot contradict the status
  field; guard_program_owner does not catch it because only program_status_authority: true is
  scanned.
why it matters: A metadata-driven consumer (or future tooling keyed on source_of_truth_for)
  could resurface an archived snapshot as an owner; it is the only residual place where old
  status surfaces still claim "current" anything.
recommended disposition: CORRECT_CANDIDATE
scope correction: In the two snapshot frontmatters, rewrite source_of_truth_for entries to
  historical phrasing ("pre-MR-01 snapshot of …") or delete the block. No body changes.
review coverage impact: none.
```

```text
ID: RM-I06
severity: MINOR
claim challenged: "important unique unmerged provenance remains reachable while a current
  consumer requires it" is a maintained property (engineering-rules §11; census "M0 R7 exact-ref
  preservation: retained").
evidence / exact authority: The frozen M0 R7 candidate is preserved only as live branch
  feat/m0-r7-sovereign-core-20260810 @ 7ec999b093… (verified reachable; cited as exact revision
  by current docs/architecture/module-runtime-topology.md and acceptance Evidence). There is no
  annotated tag (git tag list is empty), no branch protection on feat/*, and no mechanical
  reachability check in the validator, while delete_branch_on_merge is now enabled repo-wide.
  Protection is purely conventional (engineering-rules §4 prohibition).
why it matters: One accidental branch deletion or mistaken merge would silently destroy the
  byte-level provenance that current architecture authority cites by exact SHA — the reachability
  law (Repository Standard §10) would be violated only at the moment it becomes unrecoverable.
recommended disposition: DEFER_SAFELY
scope correction: none required for RM-10; recommended cheap hardening at or after promotion:
  one durable annotated tag on 7ec999b093… (Standard §10 step 3 applied proactively), recorded
  next to the existing citations. A validator/CI reachability probe is optional.
review coverage impact: none.
```

```text
ID: RM-I07
severity: MINOR
claim challenged: Blueprint 15.29 amended CODEOWNERS candidate list is correct.
evidence / exact authority: The list contains "/docs/development/" twice (consecutive duplicate
  lines); given the migrated tree, one of the two was plausibly intended as /docs/architecture/,
  which is absent from the list.
why it matters: Cosmetic duplication in a constitutional section; if CODEOWNERS is ever
  implemented from this list, architecture authority would be unprotected while development is
  listed twice.
recommended disposition: CORRECT_CANDIDATE
scope correction: Replace one duplicate with /docs/architecture/ (one line).
review coverage impact: none.
```

## Adversarial-question coverage (items producing no finding)

1. **Single status authority** — YES. `program_status_authority: true` exists only in docs/roadmap.md (guard-scanned); no router, README, index, or snapshot routes current status elsewhere; snapshot metadata residue is RM-I05, not a live second authority.
2. **Fresh-actor reconstruction** — PROVEN BY THIS REVIEW: the reviewer reconstructed the entire program state, prohibitions, and next action through AGENTS → index → roadmap → routed owners without conversation archaeology; representative packs ≤ 5 files; bootstrap 10.4 KiB.
3. **Semantic/provenance loss** — none found: ADR IDs 0001–0009 intact under decisions/adr; acceptance/review Evidence rehomed under evidence/** with dates as genuine subject identity; STATUS/WORKLOG archived as labeled snapshots; DECISIONS.md forward obligations live in decisions/index.md (O-002…O-024 with owners); superpowers semantics absorbed (architecture/phases); census records a real enforcement exception (two plan IDs + old roadmap ID given snapshots rather than weakening ID validation) — evidence the rule fired.
4. **Blueprint 15 boundedness** — HELD. All Round-1 gaps corrected (15.4 example, 15.19, 15.26 read paths, 15.31 post-A0 ownership now routes to roadmap/index); Product-meaning sections untouched; A0 acceptance preserved as historical fact.
5. **Planning Readiness vs ACRM** — non-overlapping and mutually referencing correctly (capability-realization: "ACRM does not duplicate that program", "a Capability does not replay every global TA stage"; repository rules explicitly owned by Standard + engineering-rules, not ACRM).
6. **TA graph proportionality** — unchanged from the ratified target, TA-TX conditional, tailoring law present; the Round-1 pre-TA-09 qualification-home clause is now in planning-readiness (line ~190), closing the executable-proof circularity.
7. **TA-01/TA-02 preservation** — DESIGN-AURORA-TA-01-02-MODULE-RUNTIME-TOPOLOGY lives unmodified in identity under docs/architecture/module-runtime-topology.md; decisions D-069–D-072 index it; old ordering exists only as a superseded snapshot under phases/.
8. **Conexus OS / MNFS** — current authorities carry only correct disambiguations; dual-referent clarification (Round-1 MR-I06) present in both decision §10 and engineering-rules §5; historical MNFS wording preserved where the name was current.
9. **Moved artifacts discoverability** — stable IDs verified (architecture files, ACRM DOC-AURORA-CAPABILITY-REALIZATION-METHOD, index reusing DOC-AURORA-DOCUMENTATION-MAP identity); validate_ids enforces uniqueness + related-target resolution repo-wide.
10. **Roadmap/generation isolation** — generator cannot emit docs/roadmap.md (code + negative self-test, reviewer-verified); Blueprint 14 remains Product-sequence authority; D-050 records the split.
11. **Generator/validator/CI coherence** — coherent with the target tree, with RM-I01 (environment-dependent output) and RM-I02/RM-I03 (coverage) as the exceptions.
12. **Guard falsifiability** — strong core (five falsified guard families run on every CI pass), minus the RM-I03 gaps.
13. **Final-candidate cleanliness** — REAL at eb402577: forbidden-surface list enforced in --final mode, which auto-activates exactly when docs/work is absent (i.e. on the candidate but not on review branches) — a clean design; verified against the tree.
14. **RM-08 honesty** — Evidence explicitly separates operator-attested Ruleset configuration from machine-verifiable state (protected=true, squash-only, auto-delete), records the 403-blocked probes as history, and does not overclaim. Residual: the required-check binding ("validate") itself remains attested-only until first merge attempt exercises it — acceptable; the promotion merge is its natural falsifier.
15. **Merge-flow deadlock** — none: required human approvals 0, squash-only, PR-based, auto-delete; single-operator flow works.
16. **Hidden TA-03+ selections** — none found. Python/GitHub Actions continue as existing migration-scope mechanisms; Go/SQLite/Mastra remain M0-scoped ADRs with an explicit AGENTS hard stop against globalization; O-023 keeps source topology open for TA-09.
17. **Implementation-blocked law** — preserved in roadmap authorization boundary, AGENTS hard stops, planning-readiness §18 (TA-13 + separate explicit operator Product execution grant), and 15.31.
18. **Hidden dependency on ephemeral material** — no current authority depends on docs/work, PR comments, or unmerged branches; durable Evidence summaries exist for census/coherence/enforcement; PR #8 working history is referenced as history only. Note (no finding): workflow-run IDs cited in Evidence have platform log-retention limits — the durable summaries already carry the outcomes, which is why this stays a note.
19. **Snapshot non-currency** — labeled correctly except the RM-I05 frontmatter residue.
20. **Global Maximum** — see closing assessments.

## Closing assessments

**Strongest counterargument to the migrated model:** For a single-operator pre-implementation repository, the four-index model (roadmap + index + decisions/index + architecture/index) plus a 400-line guard suite is heavier than the one STATUS.md it replaced, and each future gate now has to update more surfaces. Rebuttal that sustains the candidate: the old model's cost was not file count but authority ambiguity (status/architecture/chronology braided into tracking files that AGENTS forced every session to read); the new model pays a bounded, mechanically-guarded routing cost to make authority lookup O(1) for any actor, and this review itself is the demonstration — full program reconstruction from a cold start inside the read budget. The guard suite is proportionate because every guard maps to a named failure class that actually occurred in this repository's history.

**Best credible alternative:** Keep docs/roadmap.md as sole authority but generate a read-only status view (old STATUS path) for continuity. Rejected — it recreates a second status-shaped surface whose only defense against becoming authority is a label, reintroducing the retired failure class for zero routing gain. The migrated model is the Global Maximum under current constraints; nothing indicates overbuild.

**YAGNI / essential-vs-accidental complexity:** PASS. Every durable surface has a present consumer; snapshots were created only where ID-validation proved a consumer existed; no speculative tooling was added (no npm, no sync framework, no premature CODEOWNERS). The one accidental-complexity instance found is RM-I01 (environment-coupled generator output) — accidental in the literal sense and mechanically removable.

**Repository Standard conformance:** CONFORMANT. Hard-standard checklist verified: landing-only README; bootstrap-only AGENTS; index router with task table and do-not-read column; roadmap sole mutable authority (flag-guarded); semantic paths with real consumers; branch-only work excluded from candidate/main mechanically; squash-only protected main with required aggregate check "validate"; ≤5-file packs; 20 KiB budget enforced; reachability law honored (with RM-I06 hardening recommended). Deviations (capabilities/, history/, evidence/acceptance dates-in-names) are justified local specializations with real consumers.

**TA-01 / TA-02 reopen:** NOT REQUIRED. Preserved in identity and content routing; no migration Evidence contradicts them.

**Second independent round after corrections:** NOT REQUIRED. RM-I01–RM-I05/RM-I07 are mechanical generator/validator/metadata corrections that change no reviewed semantics; RM-I06 is deferred hardening. A second round would be warranted only if adjudication rewrites semantic content instead.

## Final verdict

```text
CONVERGED
```

```text
PASS_WITH_FINDINGS
```

Blocking: 0. Material: 1 (RM-I01). Moderate: 2 (RM-I02, RM-I03). Minor: 4 (RM-I04–RM-I07). The migrated repository achieves the ratified MR-01 operating model with authority, provenance, and prohibition boundaries intact; the material finding is confined to generator output reproducibility and guard meaningfulness, correctable on the candidate without invalidating this review's coverage.

This output is Evidence, not authority. It does not authorize merge, TA-03+, Architecture Spike execution, M0 R7/R8, or Product/runtime implementation. Lead adjudication and the separate operator merge/promotion gate remain required.
