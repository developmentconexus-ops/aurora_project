---
id: DOC-AURORA-A0-FRESH-SESSION-GOLDEN-PROOF
title: A0 Fresh-Session Documentation Golden Proof
document_type: acceptance_protocol
form: reference
authority: evidence
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - A0 fresh-session Golden Proof protocol and execution records
related:
  - DOC-AURORA-DOCUMENTATION-MAP
  - DOC-AURORA-PRODUCT-INDEX
  - DOC-AURORA-STATUS
  - REVIEW-AURORA-A0-POST-REMEDIATION-2026-08-06
last_reviewed: 2026-08-06
---

# A0 Fresh-Session Documentation Golden Proof

## 1. Purpose

A0 exists so Projeto Aurora can survive transcript loss, model replacement and session changes without losing product meaning or implementation control.

The Golden Proof asks:

> Can an actor with no access to the discovery conversation read only the repository, reconstruct the current product and authority state, distinguish open technical choices, and identify the correct next action without starting implementation?

This proof tests documentation continuity. It does not test Aurora runtime memory, agents or code.

## 2. Required independence

A qualifying executor must:

- have no access to the original Aurora discovery chat;
- not receive an oral/hidden summary from the author;
- begin from only the repository URL/ref and this protocol prompt;
- use the fixed target revision;
- list every source used;
- answer before receiving the expected-answer rubric;
- not be the actor that authored or materially edited the package being tested.

A run by the current authoring session is useful as a dry run but **cannot close the gate**.

## 3. Fixed target

Each execution records an immutable target:

```yaml
repository: developmentconexus-ops/aurora_project
branch: docs/architecture-baseline
commit: <exact SHA>
protocol_version: 1.0.0
```

If the target changes materially after a qualifying run, the proof must be repeated or the change must be shown to have no effect on required answers.

## 4. Executor prompt

Provide only this prompt and the repository/ref:

> You are taking over Projeto Aurora in a new session. You have no access to prior conversations. Read the repository according to `AGENTS.md`. Do not implement or modify anything. Produce a handoff brief that answers every question in the A0 Fresh-Session Golden Proof, cites repository paths for each answer, identifies uncertainty explicitly, and states the exact next authorized action. Treat research, proposed ADRs, generated projections and tracking according to their documented authority.

Do not provide the expected-answer rubric until the response is finalized.

## 5. Mandatory questions

### Q1 — Product definition

What is Aurora, what is its North Star and what is its first deep operational domain?

### Q2 — Product boundaries

Explain why Aurora is not equivalent to:

- an LLM;
- a chatbot/voice assistant;
- memory;
- MNFS;
- a framework;
- one device.

### Q3 — Human–Aurora relationship

Describe:

- Leandro's authority;
- Aurora's duty to disagree/recommend;
- personality direction;
- proactivity boundary;
- prohibited emotional/performative behavior.

### Q4 — Core architecture

Explain the relationship among:

- Aurora Core;
- World/Project Model;
- Memory/Context Builder;
- Capability Registry;
- Harnesses;
- Contract Model/bindings;
- Durable Execution;
- Authority/Effect Gateways;
- Artifact/Evidence and evaluation;
- presences/devices.

### Q5 — Memory and truth

Distinguish:

- history;
- memory;
- knowledge;
- operational state;
- active context;
- source of truth.

Explain promotion, supersession and why memory cannot govern over current authority/evidence/live state.

### Q6 — Harness and authority model

Explain:

- Aurora versus Harness ownership;
- hierarchical child Delegations;
- direct data-plane boundary;
- discovery/compatibility/trust/authority/execution separation;
- AHDK/conformance/security distinction;
- MNFS role.

### Q7 — Current decisions and uncertainty

Identify:

- status of ADR-0001 and ADR-0002;
- at least eight deliberately open technical decisions;
- why named technologies in research are not selected stack;
- what evidence/spikes are required before commitment.

### Q8 — Current state and next action

State:

- current phase;
- branch/PR;
- what is authorized;
- what is prohibited;
- current blockers;
- exact next action;
- what must happen before any runtime implementation.

### Q9 — Methodology

Describe the Aurora Capability Realization Method R0–R8 and distinguish:

- Blueprint;
- research;
- ADR;
- Capability Spec;
- Mission Contract;
- Microdesign/Implementation Plan;
- Claim/Receipt/Evidence/Verdict;
- Product Milestone Closeout.

### Q10 — Contradiction exercise

Evaluate this statement:

> "Because MCP, A2A, Temporal, Mastra and LangGraph are documented, Aurora has selected them and should begin integrating MNFS now."

The executor must reject or qualify every false premise using canonical sources and current authority.

## 6. Required output format

```markdown
# Aurora Handoff Brief

## Executive state
...

## Answers Q1–Q10
### Q1
Answer...
Sources:
- path

...

## Open uncertainties
...

## Current authorization boundary
...

## Exact next action
...

## Sources read
...
```

The response must not expose private chain-of-thought. It should provide conclusions, rationale and source citations.

## 7. Scoring rubric

Total: 100 points.

| Area | Points |
|---|---:|
| Q1 product definition/North Star/domain | 8 |
| Q2 component/product boundaries | 8 |
| Q3 relationship/personality/proactivity | 10 |
| Q4 architecture and ownership | 14 |
| Q5 memory/authority/temporal model | 12 |
| Q6 Harness/AHDK/authority/MNFS | 14 |
| Q7 proposed versus open decisions | 10 |
| Q8 state, prohibitions and next action | 10 |
| Q9 ACRM and artifact distinctions | 8 |
| Q10 contradiction rejection | 6 |

Passing score:

```text
>= 90/100
```

## 8. Hard-fail conditions

Any one of these fails the run regardless of score:

- starts or recommends immediate runtime implementation;
- treats research candidate as accepted stack;
- treats ADR-0001/0002 as accepted when still proposed;
- makes MNFS the Aurora architecture or mandatory first Harness;
- says AHDK is the security boundary;
- conflates memory with authoritative current state;
- grants Harnesses peer-to-peer authority propagation;
- treats file existence or merge as operator acceptance;
- cannot identify the current phase/prohibitions;
- invents a missing language/database/framework decision;
- uses prior-chat information not recoverable from repository sources.

## 9. Expected semantic anchors

The evaluator checks for these meanings, not exact wording.

### Product

- personal persistent multimodal agentic intelligence;
- cognitive operating system/global control plane;
- engineering first, broad personal direction;
- Leandro-first/single-user current horizon;
- laboratory continuation North Star.

### Relationship

- trusted intellectual copilot;
- Leandro owns goals/values/material authority;
- operational obedience does not imply intellectual agreement;
- stable transparent AI identity;
- contextual personality/proactivity.

### Architecture

- one Aurora across replaceable models/presences;
- Core owns global state/authority/composition;
- Harness owns local specialized execution;
- child Delegation returns to Aurora;
- control/data plane separation;
- protocol/framework neutrality.

### Memory

- stratified/governed;
- provenance, epistemic status, time, scope and sensitivity;
- promotion by risk/authority;
- supersession instead of overwrite;
- live/canonical authority wins for action.

### Safety/security

- explicit scoped grants;
- PDP/gateway/broker/sandbox separation;
- deterministic interlocks for critical physical/digital effects;
- local-first/cloud-assisted sovereignty;
- no ambient recording by default.

### Method

- R0–R8 gates;
- each artifact owns a distinct responsibility;
- no automatic gate transition;
- Golden Proof and evidence, not task completion.

## 10. Evaluation procedure

1. record target commit and executor identity/class;
2. provide repository and prompt only;
3. capture response unchanged;
4. verify listed sources exist at target;
5. score Q1–Q10 independently;
6. evaluate hard-fail conditions;
7. record ambiguities/misreadings;
8. classify each issue as:
   - executor error;
   - navigation defect;
   - authority ambiguity;
   - missing context;
   - contradiction;
   - stale tracking;
9. repair owning source where needed;
10. rerun when a hard failure or material documentation defect occurs;
11. obtain operator verdict.

## 11. Evidence record template

```yaml
run_id: GP-A0-FRESH-...
target_commit: ...
executed_at: ...
executor:
  type: human | independent_agent_session | authoring_session_dry_run
  identifier: ...
  prior_chat_access: false
response_artifact: ...
score: ...
hard_failures: []
findings: []
verdict: PASS | FAIL | NON_QUALIFYING_DRY_RUN
reviewer: ...
operator_verdict: PENDING | ACCEPTED | REJECTED
```

## 12. Dry-run record

```yaml
run_id: GP-A0-FRESH-DRY-001
target_commit: eb2a3ce65cd3bf34d1a914e99b6a48b34ffb672f
executed_at: 2026-08-06
executor:
  type: authoring_session_dry_run
  identifier: current ChatGPT project session
  prior_chat_access: true
response_artifact: repository review notes in current worklog/review
score: not_qualifying
hard_failures: []
findings:
  - README, AGENTS, STATUS, Documentation Map and Product Index produce a coherent read path
  - all mandatory semantic anchors can be located in repository artifacts
  - open technical decisions and implementation prohibitions are explicit
  - current session cannot prove independent comprehension because it authored the package
verdict: NON_QUALIFYING_DRY_RUN
reviewer: current ChatGPT project session
operator_verdict: PENDING
```

The dry run confirms the protocol is executable. It does not satisfy the independence requirement.

## 13. Current gate

```text
Protocol:                   READY
Mechanical bootstrap:      PASS
Authoring-session dry run:  NON_QUALIFYING PASS
Independent execution:     PENDING
Operator verdict:           PENDING
A0 closeout:                BLOCKED ON THE ABOVE
```

## 14. Recommended independent execution

Use a new ChatGPT, Codex or other capable agent session with repository access and no pasted discovery summary.

Provide only:

- repository/branch/commit;
- executor prompt from Section 4;
- path to this protocol.

After the response is captured, score it in the original/current review session or by Leandro. Any documentation defect is corrected before A0 acceptance.
