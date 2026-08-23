# Projeto Aurora — Agent Bootstrap

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
