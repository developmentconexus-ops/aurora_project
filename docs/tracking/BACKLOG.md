---
id: DOC-AURORA-BACKLOG
title: Aurora Backlog
document_type: backlog
form: reference
authority: tracking
status: current
version: 0.2.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - non-committed Aurora ideas and future investigations
last_reviewed: 2026-08-06
---

# Aurora Backlog

## 1. Authority notice

Items here are **not commitments, requirements, roadmap guarantees or implementation authorization**.

A backlog item enters execution only through:

```text
Product Milestone selection
→ ACRM applicability and requirements
→ research/spikes/ADRs
→ Capability Spec
→ approved Mission Contract
→ approved implementation design
```

## 2. A0 closeout work

- final documentation CI PASS;
- refreshed adversarial review against repaired fixed head;
- fresh-session documentation Golden Proof;
- PR summary and review guide update;
- operator review and explicit A0 acceptance;
- lifecycle promotion/merge procedure after acceptance.

## 3. Candidate next-milestone discovery

These are candidates to evaluate after A0, not a preselected sequence beyond the constitutional roadmap:

- Sovereign Core walking skeleton;
- Project/World identity and state model;
- governed conversation and project-context slice;
- reference Capability Registry/provider slice;
- Contract Model/AHDK/conformance spike slice;
- durable Delegation/recovery slice.

The first Product Milestone should be selected by risk reduction and Golden Proof value, not by framework availability.

## 4. Focused research programs

Current A0 reports exist, but implementation-adjacent revalidation/spikes remain for:

### Sovereign Core

- local-first state and event ownership;
- backup/restore and corruption recovery;
- local/cloud topology;
- process/service lifecycle;
- cryptographic identity and key storage.

### Memory and Context

- observational memory experiments;
- exact transcript/compaction boundaries;
- structured/project/episodic/relational stores;
- temporal and authority-aware retrieval;
- supersession and deletion propagation;
- long-horizon memory eval datasets;
- scale/cost/latency degradation.

### Capability and AHDK

- source schema and code generation;
- first SDK language;
- conformance profile design;
- simulator/fault injection;
- scaffolder/Golden Path UX;
- provider compatibility and migration.

### Interoperability and durability

- MCP Tasks mapping;
- A2A task/artifact mapping and SDK maturity;
- local RPC alternatives;
- CloudEvents/AsyncAPI/Protobuf boundary selection;
- Temporal/DBOS/Restate comparison;
- idempotency and ambiguous-effect recovery.

### Authority and security

- policy language/engine;
- delegated identity/token exchange;
- workload/device identity;
- Effect Gateway profiles;
- Credential Broker;
- sandbox/OS enforcement;
- supply-chain provenance and attestation.

### Presence and multimodality

- voice pipeline and interruption latency;
- local wake/activation;
- visual/spatial interfaces;
- environment classification;
- device pairing/trust/revocation;
- offline/degraded behavior;
- third-party privacy in shared environments.

### Laboratory and physical safety

- instrument discovery and identity;
- calibration/measurement uncertainty;
- telemetry transport;
- simulator/HIL architecture;
- interlocks and fail-safe controller;
- emergency stop and watchdogs;
- physical campaign evidence.

### Evaluation and self-improvement

- trace-based evals;
- root-cause correlation;
- holdout governance;
- candidate/canary comparison;
- user trust and personality evals;
- cost/latency/reliability multi-objective scoring;
- rollback drills.

## 5. Candidate capabilities/Harnesses

- `CAP-SOVEREIGN-CORE`;
- `CAP-PROJECT-WORLD-MODEL`;
- `CAP-MEMORY-CONTEXT`;
- `CAP-CAPABILITY-REGISTRY`;
- `CAP-CONTRACT-MODEL`;
- `CAP-AHDK-CONFORMANCE`;
- `CAP-DELEGATION-DURABILITY`;
- `CAP-AUTHORITY-EFFECTS`;
- `CAP-ARTIFACT-EVIDENCE`;
- `CAP-PRESENCE-FABRIC`;
- `CAP-DEVICE-LABORATORY`;
- `CAP-EVALUATION-FAILURE-INTELLIGENCE`;
- Research Harness;
- Evaluation Harness;
- Hardware Engineering Harness;
- Firmware Harness;
- Laboratory Harness;
- MNFS adapter/provider;
- mobile and wearable presences;
- engineering journal/project graph.

Names are candidates until Capability identity/specification is accepted.

## 6. Platform enablers

- AHDK;
- Conformance Kit;
- scaffolder and templates;
- Capability Registry;
- Durable Execution Port;
- Policy Decision Point;
- Effect Gateways;
- Credential Broker;
- Artifact/Evidence Stores;
- OpenTelemetry semantic conventions;
- provider provenance/attestation;
- documentation site/projections;
- automated traceability/orphan checks.

Enablers must have a named consumer and Golden Proof. The project must not become a generic platform before Aurora journeys work.

## 7. Future personal domains

Potential domains that require separate capability, privacy and authority analysis:

- Metal Nobre operational intelligence;
- personal calendar and planning;
- learning/study assistance;
- home/laboratory environment;
- finance;
- health;
- communications.

They are not included in the current executable horizon.

## 8. Explicit non-commitments

- public marketplace;
- public SDK ecosystem;
- SaaS/multi-tenancy;
- generic enterprise agent platform;
- cloud-first canonical memory;
- unrestricted peer-to-peer agent federation;
- autonomous production/self-promotion;
- ambient recording by default;
- physical actuation before deterministic safety readiness;
- implementation of every roadmap milestone;
- simultaneous TypeScript/Python/Go SDK parity;
- selection of a framework because it is currently popular.

## 9. Promotion rule

Before promoting a backlog item, record:

- which current problem/risk it solves;
- operator-visible outcome;
- canonical Blueprint sources;
- Product Milestone contribution;
- dependencies/non-goals;
- required research/spikes;
- authority and data implications;
- Golden Proof;
- removal condition if experimental.
