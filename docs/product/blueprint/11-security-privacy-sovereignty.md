---
id: DOC-AURORA-BLUEPRINT-11
title: Segurança, Privacidade e Soberania
document_type: product_blueprint_section
form: explanation
authority: constitutional
status: accepted
accepted_at: 2026-08-06
acceptance_evidence: DOC-AURORA-A0-OPERATOR-ACCEPTANCE
version: 0.2.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - security principles
  - privacy and data governance
  - local-first cloud-assisted sovereignty
  - threat and trust boundaries
related:
  - DOC-AURORA-BLUEPRINT-02
  - DOC-AURORA-BLUEPRINT-06
  - DOC-AURORA-BLUEPRINT-08
  - DOC-AURORA-BLUEPRINT-09
  - DOC-AURORA-BLUEPRINT-10
  - DOC-AURORA-BLUEPRINT-12
  - RESEARCH-AURORA-HARNESS-ARCHITECTURE-V1
review_triggers:
  - data classification changes
  - provider or credential model changes
  - threat boundary changes
  - physical control or ambient sensing changes
last_reviewed: 2026-08-06
---

# 11. Segurança, Privacidade e Soberania

## 11.1 Propósito

Aurora may accumulate a uniquely sensitive map of Leandro's life and work:

- conversations;
- projects;
- designs and source code;
- personal preferences;
- devices and laboratory topology;
- credentials and authorized actions;
- images, voice and environmental context;
- failures, habits and inferred patterns;
- company and third-party information;
- physical actuation paths.

Security and privacy cannot be delegated to a model prompt or added after the product works.

Aurora must be useful precisely because she knows context and can act. The same properties create the greatest risk if:

- context leaks;
- authority is confused;
- a provider is compromised;
- a device is stolen;
- memory becomes incorrect or impossible to delete;
- a harness bypasses an SDK;
- physical control is exposed;
- third parties are observed without governance.

The constitutional model is:

> local-first, cloud-assisted and governed.

And the sovereignty rule is:

> Intelligence can be distributed; sovereignty cannot.

---

## 11.2 Security objectives

Aurora must protect:

### Confidentiality

Only authorized actors/providers receive the minimum necessary context.

### Integrity

Memory, decisions, manifests, builds, commands and evidence cannot be silently altered or substituted.

### Availability

Critical local capabilities and recovery remain possible when a provider or network fails.

### Authority integrity

No component can grant itself or another component permissions it does not own.

### Safety

Digital compromise cannot easily become unsafe physical actuation.

### Accountability

Material data disclosure and effects are attributable and reviewable.

### User sovereignty

Leandro can inspect, correct, export, delete, revoke and migrate the system.

---

## 11.3 Security planes

Aurora separates several planes.

```text
Identity Plane
→ who is the person, Aurora, harness, worker, device or provider?

Authority Plane
→ what may this actor do, to which resource, under which delegation?

Context/Data Plane
→ what information may be read or transferred?

Effect Plane
→ where is an action enforced?

Execution Plane
→ in which process/container/device does work run?

Presence Plane
→ which interface/sensors are active and what may they reveal?

Supply-Chain Plane
→ which source/build/version is actually running?

Audit/Evidence Plane
→ what happened, why and with which result?
```

A control in one plane does not replace controls in another.

---

## 11.4 Threat model

### Threat actors and failure sources

- compromised external provider;
- malicious or vulnerable harness;
- dependency/supply-chain compromise;
- stolen device or token;
- malware on local host;
- prompt injection through documents/web/tools;
- incorrect model reasoning;
- accidental user authorization;
- misconfigured policy;
- spoofed instrument/device;
- network attacker;
- insider/third-party data exposure;
- stale memory or state;
- Aurora self-improvement candidate with harmful regression.

### Assets

- constitutional identity;
- personal memory;
- project intellectual property;
- secrets;
- credentials;
- code repositories;
- device controls;
- provider trust records;
- audit/evidence;
- backup/recovery material;
- physical safety systems.

### Consequences

- exfiltration;
- unauthorized modification;
- false context;
- financial or communication effect;
- unsafe laboratory action;
- loss of trust and continuity;
- irrecoverable memory corruption;
- hidden surveillance;
- provider lock-in;
- deletion failure.

The threat model evolves with capabilities. A text-only local prototype and an autonomous laboratory system require different controls.

---

## 11.5 Trust zones

Initial conceptual zones:

```text
Z0 — Constitutional/Sovereign Core
Z1 — Trusted local services
Z2 — Sandboxed first-party harness
Z3 — Verified external provider
Z4 — Unverified/discovered provider
Z5 — Public internet and untrusted content
Z6 — Physical device/laboratory control zone
```

Trust zone influences:

- accessible data classes;
- credential form;
- network policy;
- effect gateway;
- monitoring;
- retention;
- verification;
- authority ceiling.

Trust is not permanent. Incident, version change or environment change can downgrade a provider.

---

## 11.6 Data classification

Initial classes:

### PUBLIC

Intended for public disclosure.

Examples:

- public documentation;
- open-source repository;
- published paper.

### INTERNAL

Project or personal operational information with limited harm if disclosed.

### CONFIDENTIAL

Intellectual property, private project details or company information.

### SENSITIVE

Personal, financial, health, security, location, private audio/video or high-impact business data.

### SECRET

Credentials, private keys, recovery codes and values that should not enter model context.

### DEVICE_RESTRICTED

Data or commands constrained to a physical/device zone.

Classification can include caveats:

```text
third_party
retention_limited
no_cloud
no_model_training
project_only
presence_private_only
```

A classification label is not enough. Policy and enforcement must consume it.

---

## 11.7 Data minimization

Before external or cross-boundary transfer:

```text
requested capability
→ required information fields
→ authority and purpose
→ source selection
→ redaction/pseudonymization/reference
→ provider policy
→ transfer record
```

Aurora should prefer:

- artifact references over copying complete content;
- summaries over full history;
- scoped project context over global memory;
- ephemeral credentials over secrets;
- exact fields over broad database access;
- local processing for sensitive derivation.

### Example

A cloud model needs help choosing a circuit topology.

It may receive:

- electrical requirements;
- public component constraints;
- sanitized project context.

It may not need:

- Leandro's global personal profile;
- company credentials;
- unrelated project documents;
- complete conversation history.

---

## 11.8 Local-first boundary

The sovereign Core should keep under Leandro's control:

- Aurora identity;
- constitutional documents;
- canonical personal/project memory;
- provider registry and trust;
- authority grants;
- policy configuration;
- secrets and credential references;
- audit and effect receipts;
- deletion/export controls;
- active mission state.

“Local” may evolve from a workstation to a server controlled by Leandro. The principle is governance and ownership, not one physical disk forever.

### Cloud-assisted processing

External services may provide:

- advanced models;
- web research;
- specialized inference;
- durable/cloud compute;
- synchronization;
- selected storage.

Use is conditional on:

- provider identity;
- approved purpose;
- data classes;
- retention/training policy;
- geographic/legal constraints when applicable;
- minimization;
- cost;
- audit;
- fallback.

---

## 11.9 Provider policy

A provider profile may include:

- organization/service;
- model/endpoint;
- authentication method;
- data retention;
- training use;
- region;
- encryption;
- supported data classes;
- incident status;
- contract version;
- cost limits;
- allowed capabilities;
- availability and degraded behavior.

Provider configured does not mean provider approved for every context.

### Provider invocation record

```yaml
provider_call:
  provider: MODEL-PROVIDER-X
  model: MODEL-Y
  mission: MIS-...
  purpose: compare_control_topologies
  data_classes:
    - INTERNAL
  context_manifest: CTX-...
  redactions: 4
  retention_policy: provider-approved-profile-02
  cost: ...
  result_artifact: ART-...
```

Sensitive prompts themselves may require protected audit storage rather than general logs.

---

## 11.10 Secrets and Credential Broker

Secrets should not be copied into prompts, manifests or environment variables broadly.

A Credential Broker can issue:

- short-lived token;
- scoped credential;
- one-operation capability;
- secret reference resolved at gateway;
- device-bound credential;
- delegated actor token.

Flow:

```text
Delegation needs effect
→ Effect Gateway asks policy
→ Broker issues minimum credential
→ action occurs
→ credential expires/revokes
→ receipt preserves reference, not secret
```

Harness logs and model context must not reveal secret values.

---

## 11.11 Identity and delegation chain

Aurora must preserve:

```text
subject: Leandro
actor: Aurora
executor: Harness/Worker
presence/device: where request originated
mission/delegation: why it exists
action/resource: what is attempted
```

This avoids ambiguity between:

- Leandro acting directly;
- Aurora acting under delegated authority;
- harness acting for Aurora;
- compromised worker acting outside scope.

Identity credentials should be short-lived, verifiable and bound to context when practical.

Specific adoption of OAuth Token Exchange, SPIFFE or another mechanism remains a research/spike question.

---

## 11.12 Policy Decision and Effect Enforcement

```text
Request
→ Policy Decision Point
→ allow/deny/require_confirmation/modify
→ Enforcement at Effect Gateway
→ target system
→ receipt and audit
```

The PDP considers:

- principal/actor/executor;
- action;
- resource;
- delegation;
- project;
- environment;
- device;
- data class;
- risk;
- budget;
- time;
- prior approvals;
- emergency state.

The gateway enforces actual access.

A model response saying “not allowed” is not enforcement.

---

## 11.13 Sandboxing and execution environments

Sandboxing limits:

- filesystem;
- process creation;
- network;
- devices;
- credentials;
- resource use;
- host access.

Possible environment classes:

```text
READ_ONLY_ANALYSIS
LOCAL_RESTRICTED
PROJECT_WRITE
NETWORK_RESTRICTED
DEVICE_OBSERVE
DEVICE_CONTROLLED
PRODUCTION_DEPLOY
```

An environment profile must declare what is technically enforced versus merely requested.

A first-party SDK cannot prevent direct OS access if the process already possesses it. Security requires the environment to withhold unnecessary capabilities.

---

## 11.14 Prompt injection and untrusted content

Aurora will process hostile or misleading input from:

- web pages;
- repositories;
- issues;
- documents;
- emails;
- images;
- tool results;
- device messages.

Rules:

- external content is data, not authority;
- content cannot redefine system policy;
- tool instructions from data require independent validation;
- credentials are not exposed because a document asks;
- cross-domain instructions are rejected or escalated;
- sources and transformations remain attributable;
- high-risk effects require stronger deterministic policy.

Prompt-injection resistance cannot rely solely on prompting.

---

## 11.15 Memory privacy

Memory introduces special risks:

- inferred sensitive traits;
- outdated personal beliefs;
- accidental third-party retention;
- cross-project retrieval;
- deleted source retained in summary;
- embeddings/vector copies surviving deletion;
- model-generated memory presented as fact;
- hidden use outside intended purpose.

Memory governance requires:

- source and scope;
- sensitivity;
- promotion policy;
- access policy;
- review/edit/delete;
- derived-data tracking;
- cascade deletion or tombstone strategy;
- project isolation;
- export;
- retention and expiry;
- audit of use where material.

A memory deletion request should identify derived copies and indexes, not only remove one row.

---

## 11.16 Conversation and sensor privacy

### Conversation

Raw transcript, summaries, extracted memories and tool outputs are separate data products with potentially different retention.

### Audio/video

Policies must distinguish:

- live processing;
- temporary buffer;
- selected frame/clip;
- full recording;
- derived observations;
- biometric/speaker data;
- third-party content.

Default should avoid continuous raw retention.

### Third parties

Aurora should provide visible/appropriate behavior when other people may be captured. Exact legal requirements depend on jurisdiction and use and require future research; the constitutional principle is minimization and explicit governance.

---

## 11.17 Supply-chain security

Provider/harness trust should bind to:

- source repository/revision;
- build digest;
- dependencies;
- build provenance;
- signature/attestation;
- contract version;
- AHDK version;
- conformance result;
- environment;
- approval scope;
- known vulnerabilities/incidents.

An updated package does not inherit full trust automatically.

Potential controls:

- lockfiles;
- reproducible build where practical;
- signed artifacts;
- SLSA-style provenance;
- dependency scanning;
- isolated build;
- promotion gates;
- rollback.

Exact tooling remains open.

---

## 11.18 Presence and device security

Threats:

- lost wearable;
- rooted phone;
- malicious firmware;
- fake device identity;
- stale cached grant;
- unauthorized camera activation;
- public audio disclosure;
- replayed command;
- compromised lab controller.

Controls:

- device registration and attestation where appropriate;
- short-lived sessions;
- step-up authentication;
- local cache minimization;
- remote revoke;
- sensor indicators;
- command nonce/idempotency;
- independent interlocks;
- environment-specific authority;
- safe offline behavior.

---

## 11.19 Physical safety boundary

Physical effects require defense in depth:

```text
product policy
→ mission envelope
→ authority grant
→ device gateway validation
→ deterministic interlock
→ controller/firmware safety
→ physical protection
→ observation and emergency path
```

No single layer is sufficient.

Remote/cloud model failure must not remove local emergency behavior.

---

## 11.20 Audit and transparency

Aurora should allow Leandro to answer:

- what data was used?
- which provider received it?
- why was the provider selected?
- what redaction occurred?
- which authority allowed the action?
- which actor executed?
- what effect happened?
- what was denied?
- which credential reference was used?
- how much did it cost?
- what evidence exists?
- can the action/data be revoked or deleted?

Audit must balance detail with secret/privacy protection.

---

## 11.21 Data lifecycle

Lifecycle concepts:

```text
created/observed
→ classified
→ stored/indexed
→ used/transferred
→ updated/superseded
→ archived
→ expired/deleted
→ deletion verified
```

Retention policy depends on data type.

Examples:

- ephemeral voice buffer: seconds/minutes;
- selected experiment waveform: project evidence retention;
- active grant: until expiry plus audit record;
- secret: broker-managed, no prompt/log copy;
- superseded memory: historical but excluded from normal context;
- third-party video: minimal or no retention.

---

## 11.22 Export, deletion and portability

Sovereignty requires:

- export of canonical memory and project data;
- human-readable and machine-readable formats;
- provenance preservation;
- deletion requests;
- provider revocation;
- credential rotation;
- backup/restore;
- migration to new storage/model/provider;
- verification of restored authority and integrity.

A system that cannot be migrated or deleted is not sovereign merely because it runs on a local machine.

---

## 11.23 Backup and recovery security

Backups can be more sensitive than active storage.

Requirements:

- encryption;
- access control;
- version and integrity;
- retention;
- tested restore;
- secret separation;
- deletion implications;
- offline copy strategy where appropriate;
- incident response.

Restore must not reactivate expired grants or compromised provider trust silently.

---

## 11.24 Degraded and disconnected operation

When cloud or network is unavailable:

- local canonical state remains accessible according to policy;
- critical local safety remains independent;
- provider-dependent capabilities declare unavailable;
- queued effects do not execute later without freshness/authority recheck;
- temporary observations reconcile explicitly;
- stale authorization is not assumed current.

Aurora should favor honest limitation over simulated capability.

---

## 11.25 Security incident lifecycle

```text
DETECTED
→ CONTAINED
→ EVIDENCE_PRESERVED
→ SCOPE_ASSESSED
→ RECOVERY
→ ROOT_CAUSE
→ CORRECTIVE_ACTION
→ REVIEW
→ CLOSED
```

Possible actions:

- revoke provider/device/grant;
- isolate environment;
- disable gateway;
- rotate credentials;
- preserve logs and artifacts;
- block promotion;
- notify Leandro;
- verify physical safe state;
- update trust and memory.

Incident response itself requires authority and tested runbooks.

---

## 11.26 Security versus usability

Security that requires confirmation for every low-risk action destroys autonomy. Autonomy without enforceable boundaries destroys trust.

Aurora should use:

- scoped pre-authorization;
- risk-based step-up;
- short-lived grants;
- safe defaults;
- transparent reason codes;
- progressive trust;
- reversible actions;
- local emergency controls.

The goal is not maximum friction. It is predictable authority.

---

## 11.27 Evaluation and drills

Future system must test:

1. external provider receives only allowed data classes;
2. secret value cannot enter prompt/log through normal path;
3. expired/revoked grant blocks an effect;
4. SDK bypass cannot bypass Effect Gateway/environment;
5. prompt injection cannot grant authority;
6. cross-project memory is denied;
7. deleted memory is removed from derived indexes according to policy;
8. lost device revocation blocks context and actions;
9. compromised provider build loses inherited approval;
10. restore does not reactivate expired credentials;
11. network partition does not replay stale physical commands;
12. emergency stop works without cloud/model;
13. audit can reconstruct a provider call and effect without revealing secrets;
14. public presence withholds sensitive context;
15. third-party sensor policy is applied.

Drills produce evidence, not merely policy documents.

---

## 11.28 Open research and decisions

Still unresolved:

- policy engine;
- workload identity mechanism;
- token/delegation format;
- secret store/broker;
- sandbox/container/VM strategy;
- encryption and key architecture;
- provider policy automation;
- attestation;
- backup topology;
- legal/privacy requirements by domain/jurisdiction;
- security standards applicable to physical laboratory operation.

These require focused research and spikes.

---

## 11.29 Non-goals

This section does not claim:

- absolute security;
- a prompt can enforce policy;
- local-only execution is automatically safe;
- signed manifest is automatically trustworthy;
- one trust score can summarize all risk;
- all data must remain local;
- all external providers are unsafe;
- all laboratory control can become autonomous;
- current implementation complies with a named regulatory standard.
