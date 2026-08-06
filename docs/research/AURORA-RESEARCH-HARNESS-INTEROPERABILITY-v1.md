---
id: RESEARCH-AURORA-HARNESS-INTEROPERABILITY-V1
title: Aurora Research — Harness Interoperability and Protocol Bindings
document_type: research_report
form: explanation
authority: research
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - protocol interoperability research through 2026-08-05
related:
  - DOC-AURORA-BLUEPRINT-05
  - DOC-AURORA-BLUEPRINT-07
  - ADR-AURORA-0001
source_manifest: AURORA-RESEARCH-HARNESS-INTEROPERABILITY-v1.sources.json
review_triggers:
  - MCP or A2A specification release
  - local RPC binding design
  - SPK-002 or SPK-003 result
last_reviewed: 2026-08-05
---

# Aurora Research — Harness Interoperability and Protocol Bindings

## 1. Research question

How should Aurora connect specialized harnesses while preserving:

- Aurora-owned domain semantics;
- long-running tasks;
- tools and resources;
- context minimization;
- authority and budgets;
- artifacts and evidence;
- cancellation/recovery;
- framework/runtime neutrality?

The decision surface includes MCP, A2A, native SDK/RPC and generic HTTP/gRPC/events.

---

## 2. Executive finding

MCP and A2A are complementary rather than mutually exclusive.

```text
MCP
→ strong fit for tools, resources and bounded operations

A2A
→ stronger fit for opaque remote agent applications and stateful tasks

Native AHDK / local RPC
→ strongest first-party ergonomics and control

Aurora Contract Model
→ remains canonical above all bindings
```

Neither MCP nor A2A fully models Aurora's Authority Grant, budget, evidence, global mission composition, memory policy or physical safety. Aurora should map its contracts to these standards where they fit and extend/wrap only proven gaps.

---

## 3. MCP architecture and fit

MCP defines host/client/server relationships and core primitives such as tools, resources and prompts. Current specification work includes structured authorization and a Tasks extension for asynchronous operations [S01][S02]. The project provides official conformance tooling [S03].

### Strong Aurora use cases

- query a resource;
- invoke a bounded tool;
- discover tool/resource schemas;
- expose repository, documentation, database or instrument adapters;
- invoke a narrow operation of a Harness;
- retrieve artifacts/context by reference;
- integrate broad ecosystem servers.

### Gaps relative to Aurora

MCP does not by itself own:

- global Mission/Delegation hierarchy;
- Authority Grant/effect policies;
- per-mission budget;
- project/source authority;
- provider trust/build approval;
- evidence-to-criterion semantics;
- child Delegation;
- direct data-channel governance;
- cross-provider recovery;
- physical interlocks.

### Tasks extension

Tasks can support async tool operations, polling and result retrieval. The Aurora adapter must still distinguish:

```text
MCP Task
≠ Aurora Delegation
```

A Task may implement a bounded Delegation or one step, but Aurora maintains global identity, authority and acceptance.

---

## 4. A2A architecture and fit

A2A 1.0 describes interoperable agent applications through Agent Cards, Tasks, Messages, Parts, Artifacts, streaming/polling and push notifications [S04][S05]. The project also provides a TCK/conformance direction [S06].

### Strong Aurora use cases

- remote opaque Harness;
- long-running stateful task;
- multi-turn input request;
- artifact-producing specialized agent;
- provider that should not expose internal agents/tools;
- asynchronous collaboration across process/network boundaries.

### Gaps relative to Aurora

A2A does not automatically define:

- Aurora Context Pack semantics;
- data classification/minimization;
- Authority Grant or effect enforcement;
- budgets and physical resource limits;
- global parent/child Mission authority;
- evidence profile;
- provider build trust/provenance;
- project memory or source authority;
- promotion/rollback.

These travel through Aurora-defined contracts/metadata or remain in the Core.

---

## 5. MCP and A2A complementarity

A possible composition:

```text
Aurora Core
→ delegates a research task to Research Harness using A2A

Research Harness
→ uses MCP tools for web search, repository and documents internally

Aurora
→ receives A2A artifacts and Aurora evidence metadata
```

Another:

```text
Aurora
→ calls an instrument read operation through an MCP server

No separate opaque long-running Harness is needed.
```

Selection depends on boundary, statefulness, opacity and lifecycle—not popularity.

---

## 6. Native AHDK and local RPC

For first-party providers, a native AHDK path can expose exact Aurora types and lifecycle with lower mapping overhead.

### In-process

Benefits:

- low latency;
- direct types;
- simple development;
- rich cancellation/telemetry.

Risks:

- shared permissions/process failure;
- language coupling;
- weaker isolation;
- provider bug can affect Core.

### Local RPC

Benefits:

- process boundary;
- independent restart;
- cross-language;
- environment isolation;
- explicit compatibility.

Costs:

- protocol/schema work;
- deployment/service management;
- serialization and troubleshooting.

Aurora should not invent local RPC semantics separately from the Contract Model. It may use generated HTTP/gRPC/JSON-RPC or another mechanism after a spike.

---

## 7. Mapping Aurora concepts

| Aurora | MCP | A2A | Native/RPC |
|---|---|---|---|
| Capability definition | tool/resource profile, server metadata | skill/capability in Agent Card | generated contract |
| Provider manifest | server config + Aurora registry | Agent Card + Aurora registry | native manifest |
| Delegation | tool call/Task for bounded cases | Task | typed delegation |
| Context Pack | args/resources/refs | Message Parts/artifacts/metadata | typed payload/refs |
| Progress | progress/task events | Task status/stream | domain event stream |
| Decision request | elicitation/input pattern | input-required/multi-turn message | typed signal |
| Artifact | content/resource ref | Artifact | Artifact API/ref |
| Evidence | Aurora schema over result/resource | Aurora profile over artifact | native evidence API |
| Authority | MCP auth + Aurora grant | A2A auth + Aurora grant | short-lived grant/token |
| Budget | Aurora metadata/enforcement | Aurora metadata/enforcement | native field |
| Cancel | cancellation/task cancel | task cancel | typed cancel |
| Recovery | task handle + Aurora state | task state + Aurora state | snapshot/checkpoint |

The table is a research mapping, not final adapter design.

---

## 8. Discovery and identity

MCP server availability and A2A Agent Cards are discovery inputs. Aurora Registry adds:

- exact provider/build identity;
- provenance;
- conformance results;
- trust dimensions;
- approval scopes;
- incidents;
- environment;
- current availability.

Remote discovery should not auto-install or auto-authorize providers.

Endpoint identity must be authenticated independently of self-described metadata.

---

## 9. Authorization and data policy

Protocol authentication proves access/session identity but does not replace Aurora authorization.

Aurora adapter must propagate or bind:

- actor/delegation;
- data classification;
- allowed actions/effects;
- resource scope;
- validity;
- budget;
- environment;
- trace/audit.

A remote provider should receive short-lived scoped access, not global memory or permanent credentials.

---

## 10. Long-running and recovery semantics

For each binding, spikes must test:

- process restart;
- network partition;
- duplicate/out-of-order events;
- lost event;
- polling versus streaming;
- cancellation;
- task expiration;
- reconnect;
- provider upgrade;
- artifact availability;
- ambiguous external effects.

A protocol Task status is not enough if the provider process disappears with local state. Provider contracts and Durable Execution remain separate concerns.

---

## 11. Direct data plane

High-rate telemetry or large artifacts may use a separate channel:

```text
A2A/MCP/native control message
→ Aurora authorizes channel
→ direct object store/stream transfer
→ event carries references and integrity
```

The protocol adapter must not hide:

- data destination;
- retention;
- rate/cost;
- classification;
- revocation.

---

## 12. Versioning

Versions to negotiate:

- protocol spec;
- transport binding;
- Aurora Contract Model;
- capability contract;
- provider manifest;
- event/artifact schemas;
- extension profiles.

A provider may support A2A 1.x but not the Aurora evidence profile. Compatibility is multidimensional.

Extension design should prefer:

- namespaced metadata/profiles;
- explicit version;
- fail-closed required extensions;
- graceful optional fields;
- no semantic overloading of generic text.

---

## 13. Security considerations

- remote endpoint spoofing;
- malicious manifest/Agent Card;
- broad OAuth scope;
- context exfiltration;
- task/result replay;
- artifact URL substitution;
- webhook spoofing;
- SSRF/network exposure;
- protocol parser vulnerabilities;
- provider using undeclared tools/services;
- SDK/client supply-chain risk.

Protocol conformance is not behavioral trust or sandboxing.

---

## 14. Operational burden

### MCP

Potentially broad ecosystem and simpler bounded calls. Remote auth and evolving extensions need careful current-spec implementation.

### A2A

Closer task semantics, but SDK/TCK maturity and operational patterns must be verified per language/version.

### Native/RPC

Best control but creates maintenance responsibility and risk of unnecessary proprietary protocol.

Aurora should prefer:

```text
native AHDK for first-party simplicity
+ standards adapters for interoperability
+ no custom distributed protocol until a spike proves the need
```

---

## 15. Required spikes

### SPK-002 — MCP

Prove:

- tool and resource;
- authorization;
- structured errors;
- progress;
- cancellation;
- Tasks for async bounded operation;
- conformance suite;
- Aurora Authority/Evidence profile mapping.

### SPK-003 — A2A

Prove:

- Agent Card discovery;
- Task lifecycle;
- streaming and polling;
- input-required/decision request;
- Artifact;
- cancel;
- restart/reconnect;
- TCK;
- Aurora Context/Authority/Budget/Evidence mapping.

### SPK-008 — Neutrality

Same capability:

- native AHDK provider;
- A2A provider or second runtime;

Aurora Mission/Delegation semantics unchanged.

---

## 16. Decision implications

### Supported

- Aurora owns the canonical Contract Model;
- MCP and A2A should be evaluated as complementary bindings;
- MCP is the primary candidate for tools/resources;
- A2A is the primary candidate for remote opaque tasks;
- first-party providers benefit from native AHDK/RPC;
- protocol conformance must be tested;
- authority/evidence remain Aurora extensions/contracts.

### Not decided

- mandatory adoption of MCP or A2A;
- exact local RPC;
- which A2A/MCP SDK language;
- protocol extension format;
- event transport;
- remote deployment topology.

---

## 17. Limitations

- Protocols and SDKs are evolving rapidly.
- Published specification support may differ from SDK implementation.
- Conformance suites test protocol behavior, not domain correctness/security.
- Aurora's physical safety and personal-memory requirements extend beyond current agent interop protocols.
- Real operational burden requires spikes and failure injection.

---

## 18. Conclusion

Aurora should not choose one “agent protocol.” It needs a binding strategy:

```text
AHDK/native for first-party Golden Path
MCP for broad tools/resources
A2A for remote opaque task providers
specialized data channels for large/high-rate artifacts
Aurora contracts and authority above all of them
```

The next step is executable mapping and conformance, not a bespoke protocol design.
