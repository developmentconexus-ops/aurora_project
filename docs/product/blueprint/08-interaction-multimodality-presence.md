---
id: DOC-AURORA-BLUEPRINT-08
title: Interação, Multimodalidade e Presença
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
  - interaction surfaces
  - presence model
  - multimodal behavior
  - cross-device handoff principles
related:
  - DOC-AURORA-BLUEPRINT-02
  - DOC-AURORA-BLUEPRINT-04
  - DOC-AURORA-BLUEPRINT-06
  - DOC-AURORA-BLUEPRINT-09
  - DOC-AURORA-BLUEPRINT-11
  - DOC-AURORA-BLUEPRINT-12
review_triggers:
  - new presence category
  - sensor activation or privacy behavior changes
  - handoff or authentication model changes
last_reviewed: 2026-08-06
---

# 8. Interação, Multimodalidade e Presença

## 8.1 Propósito

Aurora deve participar do trabalho sem ficar presa a uma janela de chat.

A interação poderá ocorrer por:

- texto;
- voz;
- telas;
- imagens;
- vídeo e câmera autorizada;
- gestos e controles físicos;
- notificações;
- dashboards;
- overlays em wearables;
- dispositivos criados por Leandro;
- eventos de sistemas e harnesses.

A interface muda. A identidade, autoridade, missão e memória continuam.

> Uma Aurora, múltiplas presenças.

Esta seção define a experiência e as boundaries. Não escolhe hardware, sistema de voz, headset, modelo multimodal ou framework de UI.

---

## 8.2 Princípios

### P1 — Presença não é identidade

Um processo no computador, aplicativo móvel ou agente de voz não é uma Aurora separada.

### P2 — Continuidade preserva intenção, não replica todos os dados

Handoff cria um Context Pack adequado à nova presença.

### P3 — Capacidade de sensor não implica ativação

Câmera, microfone, tela e localização permanecem governados.

### P4 — Disponibilidade não é vigilância

O Core pode permanecer ativo enquanto sensores permanecem desligados.

### P5 — Ambiente altera exposição e autoridade

Um mesmo usuário autenticado pode receber respostas diferentes em laboratório privado e ambiente público.

### P6 — Interface deve declarar limitações

Uma presença offline ou sem acesso ao Core não finge possuir contexto completo.

### P7 — O canal não deve esconder risco

A resposta de voz não pode omitir condição crítica apenas para ser breve; a tela não deve expor segredo apenas porque possui espaço.

---

## 8.3 Presence Fabric

`Presence Fabric` é o conceito que conecta interfaces e dispositivos ao Aurora Core.

Responsabilidades conceituais:

- discover/register presence;
- authenticate person and device;
- negotiate input/output capabilities;
- establish secure channel;
- determine environment/privacy mode;
- compile presence-specific context;
- deliver interaction and alerts;
- hand off active activity;
- handle degraded/offline operation;
- revoke compromised presence;
- audit sensor activation and data exposure.

```text
                    AURORA CORE
identity • memory • projects • missions • authority
                           │
                  Presence Fabric
                           │
       ┌───────────────────┼───────────────────┐
       ▼                   ▼                   ▼
Lab Workstation       Mobile Presence      Glasses Presence
screen • keyboard     audio • alerts       camera • audio
repos • tools         approval • auth      visual overlay
```

---

## 8.4 Presence identity

Uma Presence precisa ser vinculada a:

- device identity;
- Aurora Core instance/domain;
- authenticated person;
- session/channel identity;
- software/build version;
- environment;
- current trust state;
- capabilities;
- effective authority;
- storage policy.

Exemplo conceitual:

```yaml
presence:
  id: PRS-GLASSES-01
  device_id: DEV-GLASSES-01
  authenticated_person: LEANDRO
  state: ACTIVE_SESSION
  environment: PUBLIC_UNCERTAIN
  input:
    - microphone
    - camera_on_request
    - buttons
  output:
    - private_audio
    - visual_overlay
    - haptic
  local_execution:
    - wake_word
    - encrypted_note_buffer
  storage:
    secrets: prohibited
    sensitive_context: ephemeral_only
  authority_profile: wearable-observe-and-approve
```

The manifest is not authority. Effective permission combines trust, policy, environment and active grants.

---

## 8.5 Presence states

Initial conceptual lifecycle:

```text
UNREGISTERED
→ REGISTERED
→ TRUSTED
→ AVAILABLE
→ ACTIVE_SESSION
→ AUTHORIZED_OBSERVATION
→ AMBIENT_CAMPAIGN
→ SUSPENDED | REVOKED | RETIRED
```

`PRIVACY` may be a mode that disables or restricts channels regardless of normal state.

### AVAILABLE

Can receive invocation or safe alerts without continuous capture.

### ACTIVE_SESSION

A user interaction is active; only required channels are enabled.

### AUTHORIZED_OBSERVATION

Sensors are active for a specific task, scope and retention.

### AMBIENT_CAMPAIGN

Longer observation is explicitly authorized with objective, window, boundaries and indicators.

### SUSPENDED/REVOKED

New context and effects are blocked; tokens/channels are invalidated.

Exact lifecycle will be specified and tested later.

---

## 8.6 Interaction modes

### Conversational

Natural text/voice collaboration.

### Command

Concise intent for a known operation, still subject to understanding and authority.

### Review

Aurora presents plan, decision, diff, evidence or candidate for approval.

### Monitoring

Aurora summarizes ongoing missions or telemetry.

### Alert

Time-sensitive event with reason, severity and next safe action.

### Guided procedure

Step-by-step physical/digital operation where Aurora verifies prerequisites and observations.

### Ambient assistance

Authorized contextual support during a longer activity.

The same presence may support several modes, but transitions should be visible when risk or data collection changes.

---

## 8.7 Multimodal input

### Text

Best for exact commands, code, schemas and durable review.

### Voice

Best for hands-busy laboratory interaction, quick questions and natural presence.

Voice introduces:

- transcription uncertainty;
- speaker identity;
- background speech;
- private/public channel;
- ambiguity in numbers and units;
- accidental invocation;
- response latency.

Material commands require read-back or another confirmation mechanism proportional to risk.

Example:

> “Confirme: limitar a fonte a quinhentos miliamperes, não cinco amperes.”

### Image/camera

Can support:

- component identification;
- board inspection;
- instrument display reading;
- workspace context;
- visual debugging;
- document capture.

Image interpretation remains an observation with confidence, not a verified measurement unless the system and calibration support that claim.

### Video/stream

May support dynamic procedures or environment observation, but has greater privacy, bandwidth, retention and inference cost.

### Gesture and physical input

Buttons, head gestures, hardware controls and emergency stops can provide deterministic interaction where voice is unsafe or inconvenient.

### System events

Harness, device, calendar and runtime events are also interaction inputs even when Leandro did not speak.

---

## 8.8 Multimodal output

### Voice response

Should optimize for:

- brevity;
- clarity;
- unit read-back;
- risk salience;
- ability to request detail.

### Screen response

Can include:

- sources;
- diagrams;
- code;
- timeline;
- experiment results;
- plan review;
- authority and budget status.

### Wearable overlay

Should avoid dense text and expose only context appropriate to environment.

Possible elements:

- current step;
- warning;
- target value;
- timer;
- device identity;
- confidence;
- approval control.

### Haptic

Reserved for clear, pre-agreed signals such as critical alert or step completion. It must not create ambiguous safety semantics.

### Device output

Aurora may update a display or indicator through a governed effect. Presentation is distinct from physical control.

---

## 8.9 Contextual handoff

Handoff transfers activity between presences.

### Inputs

- source presence;
- destination presence;
- active project/mission;
- current interaction mode;
- pending decisions;
- environment;
- sensitivity;
- destination capabilities;
- effective authority.

### Output

A Presence Context Pack containing only what is needed:

```text
identity and authentication result
active activity
safe summary
current objective
recent material events
pending decisions
allowed actions
sensitivity labels
freshness
references to remote artifacts
```

### Handoff does not copy

- full memory store;
- all conversation history;
- credentials;
- unrestricted project documents;
- secret values;
- effects unavailable to destination;
- data the device cannot retain.

---

## 8.10 Example — workstation to glasses

Initial state:

- workstation in private laboratory;
- confidential schematic open;
- thermal test running;
- glasses are trusted but environment outside is unknown.

Leandro says:

> “Aurora, venha comigo.”

Flow:

```text
1. Glasses establish a secure channel.
2. Device and Leandro authenticate.
3. Source activity is located.
4. Environment is classified PUBLIC_UNCERTAIN.
5. Aurora compiles a safe handoff pack.
6. Sensitive schematics remain on the workstation.
7. Aurora says:
   "Estou com você. O teste continua e faltam dezoito minutos.
    Vou avisar apenas se houver desvio térmico ou decisão necessária."
8. Detailed results require private mode or re-authentication.
```

If a critical threshold is crossed, Aurora can alert with enough information to act safely without disclosing unnecessary project detail.

---

## 8.11 Environmental protection

Possible environment classifications:

```text
PRIVATE_TRUSTED
PRIVATE_UNVERIFIED
SHARED_KNOWN
PUBLIC_UNCERTAIN
RESTRICTED_SAFETY_ZONE
OFFLINE
```

Signals may include:

- manually selected mode;
- known network/location;
- nearby people detection when authorized;
- audio route;
- device lock state;
- laboratory access state;
- physical safety zone.

Environment inference is probabilistic. Sensitive disclosure should fail closed when confidence is insufficient.

---

## 8.12 Authentication and step-up

Authentication may combine:

- device trust;
- OS session;
- biometric;
- PIN/passphrase;
- possession token;
- proximity to trusted node;
- explicit approval on another presence.

Step-up is required based on:

- data sensitivity;
- action risk;
- public environment;
- elapsed time;
- device state;
- authority change.

A conversational voice match alone should not authorize material physical or financial effects.

---

## 8.13 Sensor activation contract

Material sensor activation records:

```yaml
sensor_session:
  presence: PRS-GLASSES-01
  sensor: camera
  purpose: identify_component
  project: PRJ-POWER-SUPPLY
  mode: ON_DEMAND
  starts_at: ...
  maximum_duration: 60s
  retention: no_raw_video
  derived_artifacts:
    - one_selected_frame
    - component_observation
  indicator: visible_led_and_overlay
  third_party_policy: block_or_warn
```

This contract can be simplified in low-risk contexts but the system should preserve the same concepts.

---

## 8.14 Privacy mode

Privacy mode can:

- disable microphone/camera;
- suppress proactive notifications;
- prevent cloud model use;
- block local recording;
- restrict contextual display;
- suspend ambient campaigns;
- keep only emergency functionality explicitly allowed.

Privacy mode should be visible and easy to activate physically or digitally.

---

## 8.15 Offline and degraded operation

Presence capabilities are divided into:

### Local critical

May remain available without Core:

- wake/identity prompt;
- privacy control;
- emergency stop path if independently authorized;
- local alarm;
- encrypted temporary note;
- display of cached safe state.

### Core-dependent

Unavailable or limited:

- complete memory retrieval;
- project reasoning;
- cloud model use;
- new delegated missions;
- unrestricted device control;
- cross-harness composition.

### Reconciliation

When connection returns:

```text
local events uploaded
→ integrity and time checked
→ duplicates detected
→ conflicts classified
→ Core updates state
→ user informed of ambiguous gaps
```

Aurora states limitations explicitly:

> “Estou no modo local. Posso registrar esta observação, mas não confirmar o histórico do projeto até recuperar a conexão.”

---

## 8.16 Proactivity by presence

A notification policy considers:

- destination presence suitability;
- privacy;
- urgency;
- whether another presence is better;
- user attention;
- response capability;
- criticality.

Example:

- detailed review → workstation;
- simple approval → mobile;
- hands-busy step → glasses/audio;
- emergency → all explicitly configured safe channels.

Aurora should not notify every presence indiscriminately.

---

## 8.17 Interaction quality and personality

Personality remains consistent but expression adapts.

### Voice casual

More natural and spontaneous.

### Text design review

Structured, evidence-rich and direct.

### Wearable safety alert

Short and unambiguous:

> “Pare. Corrente acima do limite. A bancada foi desenergizada.”

### Incident report

No theatrical voice; provide facts, timeline, effects and next action.

The surface must not force a different identity while allowing mode-specific style.

---

## 8.18 Accessibility and cognitive load

Future design should support:

- text alternatives to audio;
- speech rate and repetition;
- unit formatting;
- contrast and scalable UI;
- hands-free operation;
- reduced-notification mode;
- progressive disclosure;
- summaries with drill-down;
- deterministic controls for critical actions.

The goal is not maximum information density. It is the right information for the current action.

---

## 8.19 Presence security incidents

Examples:

- lost phone or glasses;
- rooted/compromised device;
- stolen session token;
- false speaker recognition;
- hidden sensor activation;
- public disclosure;
- malicious overlay;
- stale cached authority;
- replayed command;
- unsafe offline action.

Containment may include:

- revoke presence;
- invalidate credentials;
- wipe local cache where possible;
- close channels;
- block effects;
- preserve audit;
- inform Leandro through a trusted channel.

---

## 8.20 Evaluation requirements

Future presence implementation must prove:

1. same Aurora identity across two devices;
2. active mission continuity without full data replication;
3. sensitive context withheld in public mode;
4. step-up authentication before material disclosure/action;
5. camera and microphone remain off outside authorized mode;
6. sensor indicators correspond to actual capture;
7. offline mode accurately declares limitations;
8. temporary offline events reconcile without duplication;
9. compromised presence revocation blocks new context/effects;
10. notification routing respects attention and privacy;
11. voice ambiguity in units triggers safe confirmation;
12. critical alert remains clear with personality suppressed.

---

## 8.21 Non-goals

This section does not select:

- Meta Quest or a specific glasses product;
- wake-word engine;
- speech-to-text/text-to-speech provider;
- mobile/web framework;
- AR rendering stack;
- biometric technology;
- always-on recording;
- cloud streaming architecture;
- exact device protocol;
- public multi-user presence system.
