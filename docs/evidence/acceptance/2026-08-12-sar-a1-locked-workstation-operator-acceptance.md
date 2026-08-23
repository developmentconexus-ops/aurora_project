---
id: DOC-AURORA-SAR-A1-LOCKED-WORKSTATION-OPERATOR-ACCEPTANCE
title: SAR-A1 Locked Workstation Activation Boundary Operator Acceptance
document_type: operator_acceptance
form: reference
authority: evidence
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - operator acceptance of the Stage A locked-workstation activation boundary
  - operator acceptance of unlock-required disclosure and interaction behavior
related:
  - DESIGN-AURORA-SAR-A1-STAGE-A-AVAILABILITY-ACTIVATION
  - DOC-AURORA-SAR-A1-STAGE-A-ACTIVATION-OPERATOR-ACCEPTANCE
  - DOC-AURORA-STATUS
  - DOC-AURORA-DECISIONS
accepted_design_commit: aaae0f5e88a655479e287058e6a857d3c3995cae
accepted_design_blob: 3dd8e2b0503da10b7baebcfaedc211b2b84b3d47
recorded_at: 2026-08-12
last_reviewed: 2026-08-12
---

# SAR-A1 — Locked Workstation Activation Boundary Operator Acceptance

## 1. Decision context

After accepting the Stage A workstation-sovereign, hybrid-availability and Presence-owned activation design, SAR-A1 compared three behaviors for a trigger received while the workstation is locked:

```text
A — recognize activation and require workstation unlock before continuing
B — ignore or disable activation while locked
C — provide a limited public assistant mode while locked
```

The recommended option was A because it preserves visible Aurora availability without allowing the wake word, button or hotkey to become an authentication or disclosure boundary.

The operator selected:

```text
A
```

## 2. Accepted behavior

The operator accepts `DESIGN-AURORA-SAR-A1-STAGE-A-AVAILABILITY-ACTIVATION` v0.2.0 with this locked-workstation behavior:

1. An enabled local Presence may detect an explicit trigger or optional local wake word while the workstation is locked.
2. Detection opens only a restricted activation path, not a normal Aurora interaction session.
3. Aurora may acknowledge local availability and request workstation unlock/authentication.
4. The acknowledgement is narrow and should not require a general cognitive runtime or model call.
5. No private Project, Mission, memory, conversation or personal context may be loaded or disclosed while locked.
6. No personal context may be sent to a model for general reasoning while locked.
7. No tool, Harness, digital effect or physical effect may execute from the locked activation path.
8. No notification or alert content may be disclosed while locked.
9. A previously authorized critical alert may signal that attention is required without revealing sensitive content.
10. Stage A does not provide general public question answering while locked.
11. Wake word, button, hotkey and speaker similarity remain activation signals only; they do not authenticate Leandro or grant authority.
12. Failure to determine workstation lock state fails closed.

Representative allowed response:

> “Estou disponível. Desbloqueie o computador para continuar.”

Representative allowed critical-alert signal:

> “Existe um alerta importante. Desbloqueie para consultar.”

## 3. Explicit non-decisions

This acceptance does not decide or authorize:

- authentication policy after the workstation is unlocked;
- speaker recognition;
- a future public assistant mode;
- critical-alert classification rules;
- wake-word engine or model;
- STT/TTS provider;
- Voice or Presence implementation;
- operating-system service packaging;
- Architecture Spike execution;
- Aurora runtime implementation;
- continuation or promotion of the frozen M0 R7 candidate;
- M0 R8 or M1+ implementation.

## 4. Next SAR-A1 question

The next discovery question is the real usage assumption for the unlocked Stage A workstation:

```text
exclusive Leandro operating-system account
separate accounts for other users
or other people may use Leandro’s already-unlocked session
```

That assumption must be established before deciding whether an unlocked operating-system session is sufficient baseline authentication for low-risk interaction and when Aurora-specific step-up is required.