---
id: DOC-AURORA-BLUEPRINT-09
title: Tools, Dispositivos e Laboratório
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
  - digital tool model
  - device and instrument principles
  - laboratory observation and actuation boundaries
  - physical-digital engineering integration
related:
  - DOC-AURORA-BLUEPRINT-03
  - DOC-AURORA-BLUEPRINT-04
  - DOC-AURORA-BLUEPRINT-08
  - DOC-AURORA-BLUEPRINT-10
  - DOC-AURORA-BLUEPRINT-11
  - DOC-AURORA-BLUEPRINT-12
review_triggers:
  - device identity or control model changes
  - laboratory safety boundary changes
  - new physical effect class
last_reviewed: 2026-08-06
---

# 9. Tools, Dispositivos e Laboratório

## 9.1 Propósito

Aurora was conceived for an engineering environment in which software, firmware, electronics, instruments and physical experiments belong to the same project history.

The laboratory is not an optional integration added after chat. It is one of the product's defining future domains.

Aurora must eventually understand and coordinate:

- software tools;
- repositories and build systems;
- firmware compilers and programmers;
- development boards and custom PCBs;
- sensors and actuators;
- oscilloscopes, multimeters, sources, loads and analyzers;
- device telemetry;
- test fixtures;
- simulations;
- hardware-in-the-loop environments;
- safety interlocks;
- calibration and measurement uncertainty;
- experiment protocols and evidence.

This section defines constitutional boundaries. It does not authorize physical control or choose specific instruments and protocols.

---

## 9.2 Principles

### P1 — Tool access is not device authority

Being able to call an API or open a serial port does not authorize the action.

### P2 — Observe before actuate

Aurora's progression is:

```text
inventory
→ identity
→ read-only observation
→ verified telemetry
→ guided manual procedure
→ controlled actuation
→ autonomous campaign inside deterministic limits
```

### P3 — Physical state requires live verification

Memory and last-known state cannot replace:

- device identity;
- connection status;
- current configuration;
- calibration;
- interlock health;
- actual telemetry.

### P4 — Model reasoning is not a safety interlock

Current, voltage, temperature, mechanical movement and emergency shutdown require deterministic mechanisms appropriate to risk.

### P5 — Measurement carries conditions and uncertainty

A numeric value without unit, instrument, time and conditions is not sufficient engineering evidence.

### P6 — Every material command produces an observable receipt

Aurora must know whether an effect was requested, accepted, executed, denied, failed or remains ambiguous.

### P7 — Device and firmware identity are versioned separately

A PCB, its current firmware, bootloader, configuration and calibration are distinct entities.

---

## 9.3 Taxonomy

### Digital Tool

Bounded operation in software or service.

Examples:

- read file;
- compile firmware;
- run simulator;
- query database;
- create branch;
- inspect log.

### Device

Identifiable physical or virtual equipment with state and capabilities.

### Sensor

Produces observations.

Examples:

- temperature sensor;
- current monitor;
- camera;
- encoder;
- oscilloscope channel.

### Actuator

Changes physical state.

Examples:

- relay;
- motor;
- programmable power source output;
- electronic load setting;
- GPIO;
- firmware flash operation.

### Instrument

Measurement or stimulus device with calibration and operating constraints.

### Controller

Deterministic service/firmware that exposes a governed device boundary.

### Fixture

Physical arrangement that constrains or connects a test.

### Laboratory Harness

Specialized system coordinating instruments, protocols, observations, safety and experiment execution.

---

## 9.4 Device identity

A device requires stable identity independent of connection path.

Possible identifiers:

- Aurora device ID;
- manufacturer serial;
- hardware revision;
- board ID/EEPROM identity;
- USB/PCI/network identity;
- cryptographic identity;
- physical label/QR;
- firmware-reported identity.

No single identifier is universally sufficient. Identity confidence can combine evidence.

Example:

```yaml
device:
  id: DEV-PCB-POWER-REV-B-01
  class: custom_pcb
  project: PRJ-POWER-SUPPLY
  hardware:
    design_revision: B
    serial: PSB-0001
  firmware:
    artifact: FW-POWER-014
    digest: sha256:...
    observed_at: ...
  connection:
    controller: CTRL-LAB-USB-01
    path: usb://...
  trust:
    identity_confidence: high
    last_verified_at: ...
```

Aurora should not execute a protocol for revision B if only a generic serial path is known.

---

## 9.5 Device Manifest

A Device Manifest can declare:

### Identity

- device class;
- manufacturer/project;
- hardware revision;
- firmware/driver compatibility;
- instance identity.

### Capabilities

- observations;
- measurements;
- commands;
- supported rates;
- precision/resolution;
- channels;
- timing behavior.

### Effects

- possible physical changes;
- reversible/irreversible classification;
- safe defaults;
- command preconditions.

### Safety

- absolute limits;
- recommended operating area;
- interlocks;
- emergency behavior;
- required supervision;
- environmental restrictions.

### Data

- telemetry schemas;
- units;
- timestamps;
- buffering;
- retention;
- privacy classification.

### Operation

- connect/disconnect;
- health check;
- calibration;
- reset;
- cancellation;
- idempotency;
- expected acknowledgments.

A manifest describes intended capability. Verification and calibration determine actual trust.

---

## 9.6 Instrument Registry

The future laboratory needs a registry of:

- devices;
- controllers;
- firmware/drivers;
- calibration records;
- capabilities;
- availability;
- location;
- leases;
- safety class;
- health;
- current configuration;
- incidents;
- supported protocols.

This may integrate with the Capability Registry while preserving device-specific concerns.

Example:

```text
Oscilloscope 01
├── 4 channels
├── bandwidth 100 MHz
├── current probe available
├── calibration valid until 2027-03
├── controller online
├── approved capability: waveform_capture
└── prohibited effect: firmware_update_without_manual_mode
```

---

## 9.7 Observation versus command

### Observation path

```text
device/sensor
→ controller
→ schema/unit validation
→ timestamp and quality
→ telemetry/event channel
→ Artifact/Evidence Store
→ Aurora world model/context
```

### Command path

```text
Aurora/harness intent
→ Effect Request
→ policy decision
→ device gateway
→ command validation
→ deterministic limits/interlock
→ device/controller
→ acknowledgment/live observation
→ Effect Receipt
```

These paths must remain distinguishable.

---

## 9.8 Telemetry model

Telemetry needs:

- source identity;
- signal identity;
- value;
- unit;
- timestamp at source;
- receive time;
- sequence number;
- quality/status;
- sampling rate;
- calibration reference;
- experiment/run;
- configuration;
- uncertainty where applicable.

Example:

```yaml
telemetry:
  source: DEV-ELOAD-01
  signal: current
  value: 2.700
  unit: A
  source_time: ...
  sequence: 188211
  quality: VALID
  calibration: CAL-ELOAD-2026-01
  experiment_run: RUN-THERMAL-024
```

A graph rendered from telemetry is a projection. Raw or appropriately processed evidence remains referenced.

---

## 9.9 Time synchronization

Correlating firmware logs, oscilloscope traces, temperature and control commands requires a time model.

Future architecture must address:

- clock source;
- monotonic versus wall time;
- drift;
- synchronization uncertainty;
- timestamp at source versus arrival;
- sequence gaps;
- event correlation;
- offline buffering.

Aurora must not infer causality from timestamps when synchronization quality is insufficient.

---

## 9.10 Calibration

Measurement evidence requires calibration context.

A Calibration Record may contain:

- instrument identity;
- calibration procedure/source;
- date;
- valid interval;
- ranges/channels;
- uncertainty;
- environmental conditions;
- result;
- artifacts;
- approver.

Expired calibration does not always block exploratory work, but Aurora must:

- classify evidence quality;
- warn when criteria require calibrated measurement;
- prevent a low-confidence measurement from closing a critical criterion.

---

## 9.11 Laboratory Protocol

A protocol is a versioned procedure, not free-form chat instruction.

It may define:

```text
purpose
scope
applicable device revisions
prerequisites
fixture and connections
instrument configuration
safe limits
manual steps
automated steps
observation windows
expected signals
abort conditions
cleanup
artifacts/evidence
roles and authority
```

### Example: first power-up

```text
1. Verify board revision and visual inspection.
2. Confirm no short between power rails.
3. Set source to 24 V and 500 mA current limit.
4. Keep output disabled.
5. Connect ground and probe configuration.
6. Enable output under confirmation.
7. Observe inrush, rail voltages and temperature.
8. Abort if current > threshold or rail invalid.
9. Disable output.
10. Preserve waveform and readings.
```

Aurora may guide, observe or execute steps according to authority and safety maturity.

---

## 9.12 Deterministic interlocks

Interlocks should operate outside the probabilistic reasoning path.

Examples:

- physical emergency stop;
- current limiter;
- over-temperature cutoff;
- watchdog;
- mechanical limit switch;
- independent relay;
- gateway-enforced command range;
- timeout;
- mutually exclusive resource lock.

### Interlock properties

- independent of LLM availability;
- testable;
- fail-safe where practical;
- state observable;
- bypass controlled and explicit;
- incident-producing when triggered;
- version/configuration recorded.

Aurora can decide experiment strategy. She cannot silently disable an interlock because it prevents completion.

---

## 9.13 Levels of laboratory authority

### L0 — Inventory

Know the device exists and its documentation.

### L1 — Observe

Read telemetry and status.

### L2 — Prepare

Configure a proposed protocol without applying physical output.

### L3 — Guided manual

Leandro performs material actions while Aurora verifies steps and observations.

### L4 — Controlled actuation

Aurora commands within fixed limits with confirmation and interlocks.

### L5 — Autonomous experiment campaign

Aurora adapts parameters within an approved envelope and deterministic safety system.

Authority is capability-, device-, environment- and mission-specific.

---

## 9.14 Simulation and digital twins

Before physical actuation, Aurora may use:

- circuit simulation;
- firmware emulator;
- software model;
- recorded telemetry replay;
- test double;
- hardware-in-the-loop;
- controller simulator.

Simulation can reduce risk and cost but must declare model limitations.

A digital twin is not automatically current or accurate. It requires:

- model version;
- parameters;
- validation against measurements;
- scope;
- uncertainty;
- update path.

---

## 9.15 Hardware-in-the-loop

HIL can bridge deterministic control and physical validation.

Potential flow:

```text
firmware candidate
→ simulated/controlled plant
→ injected conditions
→ telemetry
→ failure and recovery tests
→ only then physical board campaign
```

HIL should be considered a capability/provider with manifests, evidence and authority boundaries.

---

## 9.16 Firmware lifecycle

Firmware artifacts require:

- source revision;
- toolchain version;
- build configuration;
- binary digest;
- target compatibility;
- signing/attestation where applicable;
- flash procedure;
- rollback/recovery image;
- observed device version;
- test evidence.

A filename such as `firmware-final.bin` is not identity.

Aurora must distinguish:

```text
source commit
build artifact
flashed artifact
reported running version
validated running version
```

---

## 9.17 Flashing as a physical/digital effect

Firmware flashing can:

- brick a device;
- change safety behavior;
- consume limited cycles;
- invalidate calibration;
- alter communication;
- require recovery hardware.

Effect Request should include:

- device identity;
- artifact digest;
- compatibility;
- preconditions;
- backup/recovery;
- cycle budget;
- authority;
- expected acknowledgment;
- post-flash verification.

A successful programmer exit code does not prove the target is running the intended firmware.

---

## 9.18 Resource leasing

Some resources cannot be shared safely:

- serial port;
- programmer;
- fixture;
- electronic load;
- oscilloscope channel;
- physical board;
- isolated bench;
- high-power source.

A future lease model should include:

- resource;
- holder;
- mission/delegation;
- start/expiry;
- mode;
- exclusivity;
- preemption;
- cleanup;
- orphan recovery.

Aurora should detect cross-mission collisions before execution.

---

## 9.19 High-volume data plane

Oscilloscope streams, images and high-rate telemetry should not pass through language-model prompts or every control-plane message.

Authorized direct channel:

```text
source controller
→ bounded data channel
→ processing/evaluation provider
→ artifact store
```

Aurora governs:

- channel identity;
- schema;
- producer/consumer;
- data class;
- rate/volume;
- retention;
- encryption;
- budget;
- start/stop;
- audit;
- revocation.

Aurora receives summaries and evidence references rather than every sample.

---

## 9.20 Example campaign — thermal stability

Objective:

> Compare five firmware control strategies and select the best thermal-stability candidate without exceeding bench limits.

### Baseline

- board revision B;
- current firmware baseline;
- ambient range;
- load profile;
- measurement setup;
- accepted metrics.

### Allowed variables

- control parameters;
- algorithm variant;
- test order;
- repetition count within budget.

### Fixed boundaries

- maximum voltage/current;
- maximum component temperature;
- ramp rates;
- minimum cool-down;
- flash cycle budget;
- emergency cutoff;
- no production promotion.

### Flow

```text
Firmware Harness builds variants
→ Artifact identity and compatibility verified
→ Laboratory Harness acquires board/instruments
→ preflight validates wiring, calibration and interlocks
→ variant flashed
→ post-flash identity confirmed
→ load profile executed
→ telemetry streamed to Evaluation Harness
→ thresholds monitored deterministically
→ failed/unsafe variant quarantined
→ repeat if inconclusive
→ cool-down and cleanup
→ evidence bundle
```

### Outcome

Aurora reports:

- full experiment matrix;
- conditions;
- raw artifact references;
- derived metrics;
- unsafe runs;
- cost/cycles;
- candidate recommendation;
- confidence;
- unresolved uncertainty;
- no automatic firmware promotion.

---

## 9.21 Example journey — visual component query through glasses

Leandro points to a component and asks:

> “Aurora, isso parece o conector da revisão anterior?”

Flow:

```text
camera activated on request
→ selected frame captured
→ component region and visual confidence
→ project memory retrieves prior BOM/photo
→ datasheet/part reference consulted
→ Aurora answers with confidence and distinguishing features
→ no inventory or design change without explicit decision
```

Aurora should say:

> “É compatível visualmente com o JST-XH usado na revisão A, mas não consigo confirmar pitch pela imagem. A foto sugere 2,5 mm; precisamos medir ou ler a marcação antes de atualizar o BOM.”

This is better than confidently naming the component.

---

## 9.22 Manual work and human observations

Not every laboratory action will be automated.

Aurora can:

- guide a checklist;
- request a photo or reading;
- time a step;
- record Leandro's observation;
- compare with expected state;
- mark uncertainty;
- stop when evidence is insufficient.

Human observation is attributed:

```text
observed_by: LEANDRO
method: manual_visual_inspection
statement: "no visible overheating"
confidence: operator_report
```

It does not become a calibrated temperature measurement.

---

## 9.23 Physical incident handling

Examples:

- overcurrent;
- overtemperature;
- unexpected movement;
- smoke/odor report;
- electrical isolation failure;
- device communication loss during actuation;
- firmware watchdog reset;
- fixture disconnect;
- interlock bypass;
- wrong board detected.

Response order:

```text
contain
→ enter safe state
→ stop/cancel commands
→ preserve evidence
→ assess people/environment
→ notify Leandro
→ identify ambiguous effects
→ quarantine resource
→ investigate before retry
```

Aurora must never prioritize experiment completion over containment.

---

## 9.24 Security concerns

Device and laboratory threats include:

- malicious firmware;
- spoofed identity;
- compromised controller;
- command injection;
- unsafe network bridge;
- stale manifest;
- falsified telemetry;
- hidden calibration drift;
- unauthorized remote control;
- supply-chain compromise;
- credential leakage;
- physical tampering.

Security architecture is owned with Blueprint 11, but laboratory design must expose the required enforcement points.

---

## 9.25 Evidence quality

Physical evidence levels may include:

```text
INFORMAL_OBSERVATION
INSTRUMENT_READING_UNCALIBRATED
CALIBRATED_MEASUREMENT
REPEATED_MEASUREMENT
CONTROLLED_EXPERIMENT
INDEPENDENT_REPLICATION
```

The level required depends on the criterion.

A prototype exploration may use informal readings. A safety limit or final specification needs stronger evidence.

---

## 9.26 Evaluation requirements

Future implementation must prove:

1. device identity is not inferred only from port/path;
2. running firmware identity is verified after flash;
3. stale calibration lowers evidence quality or blocks required criteria;
4. read-only authority cannot actuate;
5. out-of-range command is denied outside the model;
6. interlock trips under forced fault;
7. telemetry preserves units/time/source;
8. direct data channel can be revoked;
9. ambiguous command/effect is reconciled without blind retry;
10. resource leases prevent conflicting campaigns;
11. offline controller enters safe behavior;
12. raw artifacts remain linked to derived metrics;
13. emergency containment produces trace and receipt;
14. Aurora can guide a manual protocol without claiming automated observation.

---

## 9.27 Non-goals

This section does not choose:

- instrument brands/models;
- MQTT, OPC UA, ROS, SCPI or another device protocol;
- laboratory network topology;
- microcontroller/RTOS;
- digital-twin platform;
- HIL vendor;
- calibration laboratory;
- exact safety standard;
- autonomous high-voltage operation;
- unattended physical campaigns before independent safety validation;
- a universal device abstraction that hides all domain-specific behavior.
