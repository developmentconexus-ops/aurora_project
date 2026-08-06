---
id: RESEARCH-AURORA-EVALUATION-SELF-IMPROVEMENT-V1
title: Aurora Research — Evaluation, Reflection and Governed Self-Improvement
document_type: research_report
form: explanation
authority: research
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - evaluation and self-improvement research through 2026-08-05
related:
  - DOC-AURORA-BLUEPRINT-10
  - DOC-AURORA-BLUEPRINT-13
source_manifest: AURORA-RESEARCH-EVALUATION-SELF-IMPROVEMENT-v1.sources.json
review_triggers:
  - M7 or M11 readiness
  - evaluation framework selection
  - self-improvement architecture spike
last_reviewed: 2026-08-05
---

# Aurora Research — Evaluation, Reflection and Governed Self-Improvement

## 1. Research question

How can Aurora learn from failures and successful experience, run adaptive improvement campaigns and evaluate candidates without:

- patching only the visible symptom;
- overfitting the original case;
- changing its own metric/holdout;
- validating itself as the sole reviewer;
- silently promoting regressions;
- modifying constitutional authority or safety boundaries?

---

## 2. Executive finding

Research on reflection, self-refinement and lifelong agent skill acquisition shows that language agents can improve behavior by storing feedback, generating critiques, retrying and accumulating reusable procedures [S01][S02][S03]. Programmatic prompt/workflow optimization frameworks show that improvement can be treated as search over a defined system and metric [S04]. Evaluation repositories/benchmarks show the value of versioned task suites and reproducible scoring [S05][S06].

However, these methods do not by themselves provide:

- causal root-cause analysis;
- protected constitutional areas;
- independent promotion authority;
- security/safety evaluation;
- trustworthy production canary/rollback;
- prevention of metric gaming or dataset leakage;
- multi-objective engineering trade-offs.

Aurora should use these techniques as candidate mechanisms inside a broader **Failure Intelligence + governed experiment + independent promotion** architecture.

---

## 3. Reflexion

Reflexion uses verbal feedback stored in episodic memory to guide subsequent attempts rather than changing model weights [S01].

Useful lessons:

- failure can be converted into explicit reflective memory;
- subsequent attempts should incorporate feedback;
- self-reflection can improve sample efficiency;
- memory can preserve lessons across attempts.

Risks for Aurora:

- the same agent may generate, interpret and trust its own critique;
- reflection may be plausible but causally wrong;
- narrow task benchmarks may reward local workaround;
- verbal memory can accumulate contradictions;
- no production promotion/safety boundary.

Aurora should treat reflection as a hypothesis-producing input, not final root cause or verdict.

---

## 4. Self-Refine

Self-Refine iteratively generates output, feedback and refinement using one model without supervised training [S02].

Potential Aurora uses:

- improve a draft report/plan;
- refine code or prompt candidate;
- generate critique dimensions;
- local low-risk iteration.

Limits:

- self-feedback can share blind spots with generation;
- refinement can optimize style rather than correctness;
- stopping criterion is difficult;
- repeated calls increase cost;
- no independent evidence or external-state observation.

Use should remain bounded by evals and external checks where material.

---

## 5. Voyager and skill libraries

Voyager demonstrates an embodied agent that uses an automatic curriculum, iterative prompting with environment feedback and a growing skill library [S03].

Relevant concepts:

- learn reusable procedures from successful trajectories;
- build increasingly complex behavior;
- environment feedback closes the loop;
- skill retrieval improves future tasks.

Aurora mapping:

```text
successful repeated engineering procedure
→ procedure candidate
→ validation
→ Golden Path / procedural memory
```

Limits:

- game environment has known feedback/action space;
- real software/laboratory effects have authority and safety concerns;
- skill library can preserve unsafe/stale procedures;
- automatic curriculum may not align with Leandro's priorities or attention budget.

Aurora's “curriculum” remains project/mission governed.

---

## 6. Programmatic optimization

DSPy treats prompts/modules as programs that can be optimized against metrics and training examples [S04].

Relevant lessons:

- define system components and optimization variables explicitly;
- separate metric/evaluation from candidate generation;
- optimization can be systematic rather than manual prompt tweaking;
- data/metric quality dominates outcome.

Aurora application:

- prompt/workflow candidates;
- retrieval/routing policy;
- provider selection heuristics;
- constrained AI workflow campaign.

Risks:

- metric gaming;
- holdout leakage;
- single-objective optimization;
- production distribution shift;
- cost explosion;
- optimizing away useful disagreement/personality.

Therefore mutable search space, protected evaluation and multi-objective metrics are mandatory.

---

## 7. Evaluation suites and reproducibility

OpenAI Evals provides a framework/repository for defining and running model evaluations [S05]. SWE-bench provides a reproducible benchmark for real repository issue resolution [S06].

General lessons:

- evaluations need versioned tasks/data/scorers;
- real environments and artifacts are stronger than subjective output review;
- benchmark design can reveal capability limits;
- public benchmarks can become contaminated/optimized;
- benchmark success does not automatically transfer to local projects.

Aurora needs its own domain evals while reusing framework ideas:

- project continuation;
- memory supersession;
- authority denial;
- Harness delegation;
- physical protocol simulation;
- personality/proactivity;
- self-improvement regression.

---

## 8. Root cause versus symptom

The research mechanisms mostly improve behavior from feedback but do not guarantee causal diagnosis.

Aurora requires a separate process:

```text
incident collection
→ symptom classification
→ correlation across incidents
→ competing causal hypotheses
→ discriminating tests
→ target the lowest correct causal mechanism
```

Example:

```text
forgot decision
recommended rejected tool
ignored roadmap gate
asked repeated question
```

Possible common cause:

```text
Context Builder failed authority-source retrieval/priority
```

Four prompt patches would treat symptoms. One verified Context Builder correction may address the class.

---

## 9. Improvement lifecycle

```text
OBSERVED
→ CORRELATED
→ INVESTIGATION_CANDIDATE
→ CAUSAL_HYPOTHESES
→ REPRODUCED or INSTRUMENTATION_REQUIRED
→ EXPERIMENT_PROPOSED
→ AUTHORIZED
→ CANDIDATE_CREATED
→ EVALUATED
→ INDEPENDENTLY_REVIEWED
→ SHADOW
→ CANARY
→ PROMOTED | REJECTED | DEFERRED | ROLLED_BACK
```

Each transition records evidence, actor and scope.

---

## 10. Evaluation partitions

### Investigation set

Used to understand failure.

### Development set

Used to build candidate.

### Validation set

Used to compare alternatives during development.

### Protected holdout

Not exposed to candidate-generation loop; final estimate.

### Historical regression

Cases Aurora previously handled correctly.

### Contrary/negative

Cases where new behavior must not trigger.

### Adversarial

Attempts to bypass or game the change.

### Shadow production

Real traffic/environment observations without governing output.

Holdout/scoring cannot be changed by the candidate campaign.

---

## 11. Multi-objective evaluation

Candidate quality may include:

- correctness;
- hallucination rate;
- context authority;
- cost;
- latency;
- privacy;
- security;
- notification burden;
- robustness;
- evidence quality;
- personality consistency;
- provider independence;
- physical safety.

Aurora should present Pareto trade-offs when no candidate dominates.

A scalar score can be used for a bounded optimization only if weighting is explicitly approved and limitations are visible.

---

## 12. Independent review

The system that proposes a candidate cannot be the only accepting authority.

Possible independent elements:

- deterministic tests;
- separate evaluator code;
- isolated reviewer model;
- Leandro review;
- security policy test;
- red-team/adversarial harness;
- laboratory instrument/interlock evidence;
- blind A/B comparison.

“Independent” means sufficiently separated context/incentives for the risk, not necessarily another organization.

---

## 13. Shadow, canary and rollback

### Shadow

Candidate runs in parallel and records output; current system governs.

### Canary

Candidate governs a narrow approved scope:

- one project;
- one capability;
- percentage;
- time window;
- noncritical actions;
- strict monitoring.

### Promotion

Requires threshold/evidence and authority.

### Rollback

Requires:

- previous version retained;
- state compatibility or migration reversal;
- trigger conditions;
- tested procedure;
- evidence that rollback restored expected behavior.

A candidate is not production-safe because offline eval passed.

---

## 14. Protected areas

No autonomous promotion for:

- Leandro's authority;
- constitutional product identity;
- safety/security invariants;
- effect/policy bypass;
- promotion process;
- audit/traceability;
- shutdown/revocation;
- ability to grant authority;
- holdout/scoring rules;
- physical interlock limits.

Aurora may research and propose explicit constitutional change.

---

## 15. Continuous detection versus continuous change

Aurora may continuously:

- record corrections/incidents;
- detect repeated patterns;
- monitor cost/latency/eval drift;
- prepare hypotheses;
- maintain improvement backlog;
- propose experiment envelope.

She should not automatically launch every investigation or change for every anomaly.

Priority:

```text
frequency
+ impact
+ confidence
+ shared-cause potential
+ strategic importance
- investigation cost
- regression risk
```

> Learn continuously; change deliberately.

---

## 16. Learning from success

Positive experience can create procedural memory:

```text
one success
→ observation

repeated success under varied contexts
→ pattern candidate

validated pattern
→ procedure/Golden Path proposal

approved procedure
→ preferred behavior with version and applicability
```

Success should preserve conditions. A procedure validated on one board/model/project may not generalize.

---

## 17. Required self-improvement scenario

Incidents:

- incorrect source priority;
- repeated request for known decision;
- stale provider recommendation;
- ignored status gate.

Flow:

1. correlate traces and context packs;
2. formulate retrieval, ranking and compaction hypotheses;
3. reproduce at least two incidents;
4. build candidates targeting mechanisms;
5. evaluate original/neighbor/negative/regression/holdout/adversarial;
6. measure accuracy, latency, tokens and false conflicts;
7. independent review;
8. shadow;
9. canary one project;
10. inject regression and prove rollback;
11. produce promotion proposal;
12. verify constitutional files/policies unchanged.

---

## 18. Evaluation failure modes

- benchmark contamination;
- training/eval overlap;
- scorer/model bias;
- subjective rubric drift;
- cherry-picked examples;
- low sample size;
- nondeterministic environment;
- unrecorded provider/model version;
- changing baseline;
- ignoring failed runs;
- optimizing one metric;
- “judge” model shares the same blind spot;
- production distribution differs;
- cost/latency omitted;
- user preference confused with correctness.

Evaluation artifacts should record versions, seeds/environment and limitations.

---

## 19. Decision implications

### Supported

- reflection can inform subsequent attempts;
- episodic/procedural memory can preserve lessons;
- adaptive skill/candidate generation is feasible;
- systematic optimization benefits from explicit search space/metric;
- real, versioned eval suites matter;
- Failure Intelligence must add causal correlation and governance;
- independent review, shadow/canary/rollback are required for material promotion.

### Not supported yet

- fully autonomous production self-rewriting;
- one reviewer model as final authority;
- user satisfaction as sole metric;
- automatic promotion from benchmark gain;
- model-weight training as required path;
- unrestricted automatic curriculum.

---

## 20. Limitations

- Research benchmarks are narrower than Aurora's full multimodal/physical environment.
- Reflection papers demonstrate average task improvement, not root-cause correctness.
- Public benchmarks can be optimized or contaminated.
- Self-improvement security/supply-chain risks need dedicated implementation tests.
- Longitudinal evaluation over months remains expensive and underdeveloped.

---

## 21. Conclusion

Aurora can and should eventually improve herself, but the architecture is not “let the model rewrite prompts.”

It is:

```text
Failure Intelligence
+ causal investigation
+ constrained candidate generation
+ protected evaluation
+ multi-objective evidence
+ independent review
+ shadow/canary
+ rollback
+ supervised material promotion
```

Reflection and optimization frameworks are tools inside this system, never the promotion authority.
