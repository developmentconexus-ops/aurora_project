---
id: RESEARCH-AURORA-AUTHORITY-IDENTITY-EFFECTS-V1
title: Aurora Research — Identity, Delegated Authority and Effect Enforcement
document_type: research_report
form: explanation
authority: research
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - security architecture research on identity, authorization and effects through 2026-08-05
related:
  - DOC-AURORA-BLUEPRINT-05
  - DOC-AURORA-BLUEPRINT-10
  - DOC-AURORA-BLUEPRINT-11
  - DOC-AURORA-BLUEPRINT-12
source_manifest: AURORA-RESEARCH-AUTHORITY-IDENTITY-EFFECTS-v1.sources.json
review_triggers:
  - SPK-005 result
  - policy engine or workload identity selection
  - new effect gateway class
last_reviewed: 2026-08-05
---

# Aurora Research — Identity, Delegated Authority and Effect Enforcement

## 1. Research question

How can Aurora prove who is acting, preserve delegation from Leandro through Aurora and Harnesses, make policy decisions from context, issue minimum credentials and enforce effects at digital/physical boundaries?

Candidate patterns:

- Cedar;
- Open Policy Agent;
- OAuth 2.0 Token Exchange;
- SPIFFE/SPIRE;
- capability/scoped tokens;
- effect gateways;
- sandbox/environment restrictions;
- build provenance;
- agent-framework security limitations.

---

## 2. Executive finding

Security requires a chain, not one permission system:

```text
Identity
→ actor/workload is verifiable

Delegation
→ subject, actor and purpose are preserved

Policy Decision
→ request is evaluated against scope/context

Credential Broker
→ minimum short-lived access is issued

Effect Gateway / Sandbox / Device Controller
→ decision is enforced at the action

Receipt / Audit
→ outcome is attributable and reconcilable
```

Cedar and OPA are viable policy-decision candidates with different models and operational trade-offs. OAuth Token Exchange provides useful subject/actor delegation semantics. SPIFFE provides workload identity and short-lived credentials. None alone solves Aurora's complete authority model or physical safety. AHDK and agent guardrails are not security boundaries.

---

## 3. Identity categories

Aurora needs distinct identities for:

- Leandro/person;
- Aurora instance;
- Presence/device;
- Mission/Delegation;
- Provider/Provider Instance;
- Harness/worker/run;
- Effect Gateway;
- external service/device;
- artifact/build.

Identity claims must be bound to:

- issuer;
- audience;
- validity;
- environment;
- cryptographic or local trust root where appropriate;
- build/workload identity;
- revocation/rotation.

A string `agent_name=research` is not sufficient identity.

---

## 4. Subject, actor and executor

OAuth Token Exchange distinguishes the subject of a token from an actor acting on behalf of the subject [S03]. Aurora needs an expanded chain:

```text
Subject / authority origin: Leandro
Actor: Aurora
Executor: Harness/Worker
Originating Presence: device/channel
Purpose: Mission/Delegation
```

Benefits:

- audit can distinguish direct versus delegated action;
- policy can restrict which actor may act for a subject;
- child provider receives a new delegation, not copied parent authority;
- credentials can include audience/resource/action;
- incident scope is attributable.

Aurora may use OAuth-style token exchange for service APIs, but physical/local effects may require different tokens/capabilities.

---

## 5. Cedar

Cedar is an authorization policy language/engine built around:

```text
principal
Action
resource
context
```

with explicit allow/forbid policies and schema/entity models [S01]. Its documentation includes guidance for automated agents acting on behalf of users.

### Strengths

- explicit authorization domain;
- default-deny and forbid semantics;
- analyzable policy language;
- schema/entity hierarchy;
- separation from application logic;
- suitable for principal/action/resource/context model.

### Aurora fit

- actor/delegation policy;
- project/resource scope;
- data/effect decisions;
- device/provider relationships;
- explainable decision inputs.

### Questions/risks

- dynamic numeric budgets/physical constraints may need application checks;
- entity model synchronization;
- policy deployment/versioning;
- obligations/modifications workflow;
- integration language/runtime;
- operational decision logs.

Cedar is a strong candidate for Policy Decision Point, not a device interlock or credential system.

---

## 6. Open Policy Agent

OPA separates policy decision from enforcement and evaluates policies written in Rego over structured input/data [S02]. It supports decision logs and broad infrastructure/application integration.

### Strengths

- general-purpose policy decisions;
- flexible structured input;
- mature ecosystem;
- bundle/distribution and decision logs;
- can express context-rich rules;
- sidecar/library/service deployment options.

### Aurora fit

- provider/effect/data policy;
- environment and request context;
- externalized decision service;
- audit and policy versioning.

### Questions/risks

- Rego complexity and maintainability;
- policy performance/deployment;
- entity relationships modeled in data;
- explanation/user-facing reason codes;
- avoiding policy as a second domain store;
- physical constraints still need controller enforcement.

OPA is also a candidate, not an accepted choice.

---

## 7. Policy model requirements

Aurora's policy input may include:

```yaml
request:
  subject: LEANDRO
  actor: AURORA
  executor: LAB-HARNESS@sha256:...
  presence: PRS-LAB-DESKTOP
  mission: MIS-...
  delegation: DEL-...
  action: source.set_output
  resource: DEV-SOURCE-01
  data_class: DEVICE_RESTRICTED
  environment: LAB-BENCH-01
  authority_grant: GRANT-...
  budget_remaining: ...
  requested_parameters:
    voltage_v: 24
    current_limit_a: 0.5
  risk: PHYSICAL_CONTROLLED
  time: ...
```

Decision may return:

```text
ALLOW
DENY
REQUIRE_CONFIRMATION
ALLOW_WITH_CONSTRAINTS
DEFER
EMERGENCY_ALLOW
```

Constraints/obligations can include:

- parameter max;
- required interlock;
- private presence;
- receipt/audit;
- one-time credential;
- confirmation;
- evidence requirement.

The gateway must enforce them.

---

## 8. SPIFFE/SPIRE

SPIFFE defines workload identity using SPIFFE IDs and SVIDs, with short-lived X.509/JWT credentials delivered to workloads [S04]. SPIRE is a reference implementation.

### Strengths

- identity independent from IP/hostname;
- workload attestation;
- short-lived credentials;
- service-to-service trust;
- rotation;
- multi-environment identity.

### Aurora fit

- Core, Harness, gateways and device-node workload identity;
- remote/local service authentication;
- provider instance identity;
- reduced static API keys.

### Questions/risks

- operational overhead for single-user initial topology;
- attestation mechanisms on WSL2/home lab/devices;
- mapping workload identity to Aurora provider/build trust;
- embedded device support;
- not a user/person authentication system;
- does not define authorization.

Likely later-stage or selective adoption; architecture should not preclude it.

---

## 9. Credential Broker

A Broker mediates secrets and credentials:

```text
Effect Request
→ policy allows
→ Broker resolves secret reference or issues scoped token
→ Gateway/provider uses it
→ credential expires
→ receipt stores reference
```

Required properties:

- least privilege;
- short lifetime;
- audience/resource binding;
- actor/delegation binding;
- rotation/revocation;
- no secret in prompt/general logs;
- audit of use;
- provider cannot request arbitrary secret name;
- child delegation gets separate access.

Potential implementation may use OS keychain, Vault-like service, cloud secret manager, OAuth exchange or device-specific credentials. Selection remains open.

---

## 10. Effect Gateways

A Gateway owns an effect family and is the enforcement point.

Candidate families:

```text
Filesystem Gateway
Network/Egress Gateway
Repository Gateway
External Communication Gateway
Deployment Gateway
Database/Migration Gateway
Credential Gateway/Broker
Financial Gateway
Device/Laboratory Gateway
```

Gateway responsibilities:

- authenticate actor/request;
- validate contract/schema;
- call PDP;
- enforce constraints;
- obtain minimum credential;
- idempotency;
- execute or deny;
- capture external reference;
- emit receipt/audit;
- reconcile ambiguous state;
- support revoke/containment.

Avoid a universal shell gateway with unrestricted host permissions as the default boundary.

---

## 11. Sandbox and environment

Pi documentation explicitly warns that the coding agent process uses the permissions of the launching process and does not provide a complete sandbox by itself [S05]. OpenHands similarly documents security analysis/policy mechanisms but recommends sandboxing and notes bypass limitations [S06].

This validates:

> agent policy/confirmation is not sufficient if the process can directly use OS/network/credentials.

Sandbox/environment profiles should constrain:

- filesystem mounts;
- process/commands;
- network/egress;
- environment variables/secrets;
- CPU/memory/time;
- device access;
- kernel/host boundary;
- user identity;
- artifact output path.

First-party providers may use AHDK, but environment enforcement remains necessary for high-risk scopes.

---

## 12. Prompt injection

Untrusted content can instruct an agent to:

- reveal secrets;
- ignore project policy;
- call a tool;
- upload files;
- modify memory;
- contact another service.

Controls:

- treat content as data, never authority;
- minimize tools/context;
- policy/effect enforcement outside model;
- separate system/contract sources;
- no raw credential access;
- validate tool arguments;
- high-risk confirmation;
- provider sandbox/egress;
- memory write provenance;
- adversarial eval.

No prompt can guarantee prompt-injection immunity.

---

## 13. Build/provenance and provider trust

SLSA provenance describes how/where/from which source an artifact was built [S07]. Aurora can use provenance as one trust input:

- source repository/revision;
- builder identity;
- build parameters;
- artifact digest;
- dependencies/materials;
- attestation/signature.

Provenance proves lineage claims when verified. It does not prove functional correctness or absence of malicious source.

Registry approval should bind to exact build/provenance plus conformance/security evidence.

---

## 14. Revocation

Revocation needs layers:

- policy denies new requests;
- gateway blocks effects;
- token/credential invalidated;
- provider notified/canceled;
- channel closed;
- environment contained;
- device enters safe state;
- active effects reconciled;
- Registry/trust updated;
- audit preserved.

Offline/partition behavior requires short expiry and local enforcement. A remote Core cannot revoke a credential that remains valid indefinitely on a disconnected controller.

---

## 15. Emergency effects

Emergency containment may be preauthorized:

- power off;
- stop motion;
- revoke provider;
- close data channel;
- cancel cost runaway;
- isolate environment.

Properties:

- narrow target/action;
- independent trigger where possible;
- local availability;
- cannot be repurposed to continue work;
- receipt and notification;
- drill-tested;
- physical/manual alternative.

---

## 16. Decision logging and explanation

Leandro should see:

```text
Request: merge PR #42
Decision: REQUIRE_CONFIRMATION
Reason: production branch; irreversible external visibility; current grant permits branch write only
Policy version: ...
Alternatives: create draft PR / keep branch
```

Machine audit and user explanation may have different formats but derive from the same decision facts.

Sensitive details and policy internals should not be overexposed.

---

## 17. SPK-005

Scenario:

```text
1. Leandro grants Aurora a narrow project scope.
2. Aurora delegates to Provider A.
3. Provider receives short-lived actor/delegation identity.
4. Allowed artifact write succeeds.
5. Network and credential attempts outside scope fail.
6. Child Capability request creates a separate grant.
7. Revoke parent grant during execution.
8. New effects fail immediately.
9. Provider is partitioned; local credential expiry/gateway prevents continued effects.
10. Audit reconstructs subject, actor, executor, decision and receipt.
```

Compare:

- Cedar;
- OPA;
- minimal domain policy baseline;
- identity/token options;
- gateway enforcement.

Measure:

- expressiveness;
- explainability;
- policy lifecycle;
- performance;
- integration;
- operational burden;
- testability;
- failure behavior.

---

## 18. Decision implications

### Supported

- separate Identity, Authority, Credential and Effect layers;
- preserve subject/actor/executor;
- use external PDP candidate rather than hard-code all rules in agents;
- enforce at gateways/environment;
- short-lived scoped credentials;
- build-bound provider trust;
- independent physical interlocks;
- explicit revocation/receipts.

### Not decided

- Cedar versus OPA;
- SPIFFE timing;
- OAuth/token format;
- secret store/broker;
- sandbox technology;
- gateway process topology;
- policy data model;
- attestation/signature stack.

---

## 19. Limitations

- Enterprise policy/workload systems may be too heavy for early single-user Core.
- Policy correctness depends on the domain model and tests.
- Workload identity does not prove source/build correctness alone.
- Sandboxing varies by OS/WSL2/device.
- Physical safety needs domain standards and hardware engineering beyond general IAM.
- External-service APIs may limit fine-grained delegation/revocation.

---

## 20. Conclusion

Aurora's authority should not be implemented as a model asking “may I?” or a list of tool names.

The target chain is:

```text
Leandro authority
→ Aurora Delegation/Grant
→ verifiable actor/workload
→ contextual policy decision
→ minimum credential
→ effect-specific gateway/environment
→ target
→ receipt, observation and audit
```

SPK-005 must prove the chain under denial, revocation and partition before material autonomy is accepted.
