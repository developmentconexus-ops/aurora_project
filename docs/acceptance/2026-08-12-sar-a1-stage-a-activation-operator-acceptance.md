---
id: DOC-AURORA-SAR-A1-STAGE-A-ACTIVATION-OPERATOR-ACCEPTANCE
title: SAR-A1 Stage A Availability and Activation Operator Acceptance
document_type: operator_acceptance
form: reference
authority: evidence
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - operator acceptance of the Stage A workstation-sovereign topology
  - operator acceptance of the persistent-minimum and on-demand-cognition availability boundary
  - operator acceptance of Presence-owned activation and optional local wake-word semantics
related:
  - DESIGN-AURORA-SAR-A1-STAGE-A-AVAILABILITY-ACTIVATION
  - DOC-AURORA-STATUS
  - DOC-AURORA-DECISIONS
  - DESIGN-AURORA-SYSTEM-ARCHITECTURE-DECISION-LANDSCAPE
accepted_design_commit: 8902b0a03424a3580aaf2deadccd22dd6c632a7d
accepted_design_blob: ec491e9bd930697007a45163b3f6c033b138f420
recorded_at: 2026-08-12
last_reviewed: 2026-08-12
---

# SAR-A1 — Stage A Availability and Activation Operator Acceptance

## 1. Accepted decision sequence

During SAR-A1 discovery, the operator selected:

```text
Stage A topology option A
→ one Leandro-controlled workstation is the initial sovereign Aurora node
→ the workstation is also the first Presence
```

The availability discussion then clarified that the recommended operating model is:

```text
hybrid availability
→ lightweight Sovereign Core minimum remains active
→ local Presence / Activation Agent remains available
→ cognitive runtime, STT, models, tools/Harness connections and TTS start on demand
```

The operator further clarified the intended user experience as Alexa-like availability through an explicit trigger such as:

- a button;
- a keyboard shortcut;
- an interface action;
- or saying Aurora’s name.

The architecture discussion separated the responsibilities:

```text
Presence
→ detects the trigger

Sovereign Core
→ validates Presence/session/policy and opens or restricts the interaction

on-demand capabilities
→ perform speech, cognition, model and tool work
```

After reviewing the complete boundary—including optional local wake word, explicit indicators, sensor minimization and the rule that wake word is neither authentication nor authority—the operator responded:

```text
Aprovado
```

## 2. Accepted scope

The operator accepts `DESIGN-AURORA-SAR-A1-STAGE-A-AVAILABILITY-ACTIVATION` v0.1.0 with these decisions:

1. Stage A uses one workstation controlled by Leandro as the sovereign Aurora host and first Presence.
2. The design preserves later migration to a persistent Stage B home/laboratory node without changing Aurora domain identity or authority meaning.
3. Stage A uses a hybrid availability model.
4. A lightweight Sovereign Core minimum remains available for identity, state, recovery, minimal policy, authorized timers/alerts and capability startup.
5. A lightweight local Presence/Activation Agent receives or detects triggers.
6. Heavy cognition, full STT, models, retrieval, tools/Harnesses and TTS start only when required.
7. Button/UI/hotkey form the explicit initial trigger baseline.
8. A local wake word such as “Aurora” is an accepted optional activation adapter.
9. Wake word/button/hotkey trigger interaction activation only; they do not authenticate the actor or grant authority.
10. Full audio capture/transcription and provider transfer do not follow from Core availability alone and require explicit policy/indicator boundaries.
11. Logical Core and Presence responsibilities remain separate even if the first deployable packages them together.

## 3. Explicitly not accepted or authorized

This decision does not select or authorize:

- a wake-word engine or model;
- STT or TTS providers;
- speaker recognition;
- Windows Service, systemd or another supervisor;
- one-process versus two-process packaging;
- IPC, HTTP, gRPC or another local binding;
- behavior while the workstation is locked;
- disclosure policy for unknown speakers;
- continuous retained audio;
- cloud wake-word detection;
- Voice implementation;
- Presence implementation;
- Aurora runtime implementation;
- Architecture Spike execution;
- continuation or promotion of the frozen M0 R7 candidate;
- M0 R8 or M1+ implementation.

## 4. Next SAR-A1 boundary

The next discovery/design question is limited to the session-opening policy when trigger, workstation state and actor confidence differ:

```text
unlocked owner workstation session
locked workstation
unknown or non-owner speaker
uncertain actor identity
```

The next decision must preserve the separation between:

```text
availability
activation
authentication
disclosure
authority
```

No next technical mechanism or implementation is authorized by this acceptance.
