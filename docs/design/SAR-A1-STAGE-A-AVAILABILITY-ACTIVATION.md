---
id: DESIGN-AURORA-SAR-A1-STAGE-A-AVAILABILITY-ACTIVATION
title: SAR-A1 Stage A Availability and Activation Boundary
document_type: system_architecture_design
form: explanation
authority: design
status: accepted
accepted_at: 2026-08-12
acceptance_evidence: DOC-AURORA-SAR-A1-STAGE-A-ACTIVATION-OPERATOR-ACCEPTANCE
version: 0.1.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - Stage A workstation-sovereign topology
  - Stage A persistent-minimum and on-demand-cognition availability boundary
  - Presence-owned activation trigger semantics
  - baseline relationship between button, hotkey and optional local wake-word activation
related:
  - DOC-AURORA-BLUEPRINT-08
  - DOC-AURORA-BLUEPRINT-10
  - DOC-AURORA-BLUEPRINT-11
  - DOC-AURORA-BLUEPRINT-12
  - DESIGN-AURORA-SYSTEM-ARCHITECTURE-DECISION-LANDSCAPE
  - DOC-AURORA-STATUS
  - DOC-AURORA-DECISIONS
  - DOC-AURORA-SAR-A1-STAGE-A-ACTIVATION-OPERATOR-ACCEPTANCE
source_revision: 564d677daee4f7b27ec7203d75317976076e7205
review_triggers:
  - Stage A moves from one workstation to a persistent home or laboratory node
  - activation is required while the workstation is locked or no owner session is active
  - wake-word processing requires cloud transfer or continuous retained audio
  - Presence and Sovereign Core must become separate operating-system services
  - a second Presence or remote device enters the executable horizon
last_reviewed: 2026-08-12
---

# SAR-A1 — Stage A Availability and Activation Boundary

## 1. Purpose

This design fixes the first bounded Stage A interpretation for Aurora availability and activation without selecting a Voice provider, wake-word engine, authentication product, operating-system service manager or process topology.

It answers four questions:

1. Where does the first sovereign Aurora installation live?
2. What remains active when no conversation is occurring?
3. Which component detects that Leandro wants to interact?
4. What does an activation trigger prove—and what does it not prove?

This document is architecture only. It does not authorize implementation.

---

## 2. Accepted Stage A topology

Stage A uses a **single workstation controlled by Leandro as the sovereign node**.

```text
Leandro
   │
   ▼
Stage A workstation
├── Aurora Sovereign Core
├── canonical operational state
├── governed-memory capability when separately authorized
├── artifact/evidence storage
├── first local Presence
└── adapters to approved external models, services and future Harnesses
```

The workstation is both:

- the initial sovereign Aurora host; and
- the first Presence through which Leandro interacts.

Stage A does not require:

- a dedicated home server;
- a laboratory node;
- a cluster;
- Kubernetes;
- service mesh;
- remote-device authentication;
- network discovery between Aurora-controlled nodes.

The design must preserve migration to Stage B:

```text
Stage A
workstation = sovereign node + Presence

        ↓ explicit migration

Stage B
persistent personal server = sovereign node
workstation = Presence
```

Aurora domain identities, project meaning, authority semantics and portable state must not require redesign merely because the sovereign host moves.

---

## 3. Availability model

Stage A uses a **hybrid availability model**.

### 3.1 Always-active logical responsibilities

Two lightweight logical responsibilities remain available:

```text
Sovereign Core minimum
+
local Presence / Activation Agent
```

They may initially share one deployable process or package if that is the smallest safe topology. Logical ownership must remain distinct even if physical packaging is combined.

### 3.2 Sovereign Core minimum

The minimum Core is responsible for:

- Aurora instance identity and bootstrap;
- canonical operational state access;
- restart recovery and integrity checks;
- current Presence registration/status;
- minimal session-opening policy;
- timers and alerts that belong to an authorized current capability;
- a local controlled endpoint for the Presence;
- starting approved on-demand capabilities.

The minimum Core does not imply that an LLM, STT, TTS, model router, full retrieval pipeline or Harness runtime remains active continuously.

### 3.3 Local Presence / Activation Agent

The local Presence side is responsible for detecting or receiving interaction triggers such as:

- application or tray-button activation;
- keyboard shortcut;
- physical button through a future adapter;
- optional local wake word;
- a governed operating-system event;
- a future Presence-specific trigger.

The Presence side owns channel, device, environment and activation context. It does not become the owner of Aurora identity, canonical state, authority or global Mission meaning.

---

## 4. On-demand responsibilities

The following start only when required by an authorized interaction or capability:

- full speech-to-text;
- cognitive runtime;
- external or local model connection;
- model routing;
- memory retrieval and Context Builder work beyond the minimum state lookup;
- tools and Harness connections;
- text-to-speech;
- heavy evaluation, indexing or analysis workloads.

Conceptual flow:

```text
trigger detected
→ Presence emits Activation Request
→ Core validates Presence and current policy
→ interaction session opens or is restricted/denied
→ required on-demand capabilities start
→ interaction completes or times out
→ heavy capabilities stop or return to idle
```

---

## 5. Unified activation contract

Button, hotkey, UI and wake word are adapters to one logical activation boundary.

Conceptual request:

```text
ActivationRequest
├── presence_ref
├── trigger_type
├── observed_at
├── requested_channel
├── environment_context_ref
├── detector_confidence when applicable
├── device/session context
└── correlation identity
```

Exact schema and wire representation remain undecided.

The Core may respond by:

- opening a text or Voice interaction;
- showing a notification;
- requesting authentication or step-up;
- limiting disclosure;
- refusing because the Presence is unavailable/revoked;
- deferring until an approved capability is available.

Separate trigger adapters must not create separate product meanings for “activate Aurora.”

---

## 6. Trigger baseline

### 6.1 Initial required trigger class

The initial baseline must support at least one explicit local trigger:

- application/UI button; or
- keyboard shortcut.

This lets Aurora prove session opening, state recovery, policy, indicators, timeout and on-demand capability startup without making wake-word quality an initial architecture dependency.

### 6.2 Optional local wake word

A local wake word such as “Aurora” is an accepted optional Stage A activation adapter.

The intended behavior is:

```text
restricted local detector observes only for wake-word pattern
→ wake word detected
→ visible/audible activation indicator
→ full interaction capture/transcription begins under policy
→ session completes or times out
→ full capture/transcription stops
```

The exact detector, model, acoustic pipeline, latency target and packaging remain future research/decision work.

Cloud wake-word detection or continuous cloud audio transfer is not selected and would require explicit privacy/security review.

---

## 7. Security meaning of activation

A trigger means only:

> A Presence reports that an interaction with Aurora was requested.

It does **not** prove:

- that the speaker is Leandro;
- that the operating-system session is trusted;
- that a requested action is authorized;
- that private context may be disclosed;
- that a material digital or physical effect may execute.

The separation is mandatory:

```text
wake word / button / hotkey
→ activation

authentication / Presence trust
→ who or what is interacting

authority / policy
→ what may be disclosed or performed
```

Speaker recognition, if later adopted, is an authentication signal with uncertainty—not a universal authority grant.

---

## 8. Privacy and sensor boundary

Core availability does not imply continuous sensing.

Outside an explicit policy:

- camera remains inactive;
- full STT remains inactive;
- full audio recording/retention remains inactive;
- model/provider audio transfer remains inactive;
- wake-word observation, when enabled, is restricted to the minimum local detection purpose;
- activation must produce an understandable indicator before broader capture begins;
- retained audio and derived data require a separately defined lifecycle.

The wake-word adapter must fail closed or degrade visibly when its privacy/indicator requirements cannot be met.

---

## 9. Failure behavior

### Core unavailable

The Presence reports that Aurora is unavailable. It must not simulate a successful session.

### Wake-word detector unavailable

Button/hotkey/UI activation may remain available. The system reports wake-word degradation rather than silently transferring detection to an unapproved provider.

### False activation

The session can be canceled by timeout, explicit dismissal or no-follow-up detection. False activation does not create authority or durable memory automatically.

### Ambiguous speaker or environment

Aurora restricts disclosure or requests authentication according to the future session-opening policy.

### Heavy runtime fails to start

The Core preserves the activation/session receipt and reports capability unavailability. It does not claim that Aurora understood or completed the request.

---

## 10. Stage A to Stage B invariants

When the sovereign node later moves to a persistent server:

- the workstation Presence keeps the same conceptual activation contract;
- wake word/button/hotkey remain Presence adapters;
- the Core remains the session/policy/state authority;
- remote transport and Presence authentication become new mechanisms, not new product semantics;
- Presence-local audio may remain local until minimized/authorized data crosses to the sovereign node;
- a lost or revoked Presence cannot become a second Aurora identity.

---

## 11. Explicit non-decisions

This accepted design does not choose:

- Windows Service, systemd, launchd or another supervisor;
- one process versus two processes;
- local IPC, HTTP, gRPC or another binding;
- wake-word engine or model;
- audio capture library;
- STT/TTS provider;
- Voice streaming protocol;
- speaker-recognition mechanism;
- behavior while the workstation is locked;
- response policy for non-owner speakers;
- audio retention duration;
- Voice latency/SLO;
- microphone hardware;
- a permanent background model runtime.

These choices remain governed by SAR-A1 or later consuming capability work.

---

## 12. Next SAR-A1 question

The next boundary to decide is the **session-opening policy for actor and workstation state**, especially:

```text
What happens when a trigger occurs while:
- Leandro’s workstation session is unlocked;
- the workstation is locked;
- another person says the wake word;
- actor identity is uncertain?
```

That decision must separate availability, disclosure, authentication and authority without requiring speaker recognition to become the sole security boundary.
