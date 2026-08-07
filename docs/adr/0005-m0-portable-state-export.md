---
id: ADR-AURORA-0005
title: M0 Portable Logical State and Export Envelope
document_type: adr
form: explanation
authority: decision
status: accepted
accepted_at: 2026-08-07
acceptance_evidence: DOC-AURORA-M0-R4-ADR-ACCEPTANCE-SPK001-AUTHORIZATION
version: 0.1.0
owners:
  - developmentconexus-ops
approvers:
  - operator
source_of_truth_for:
  - proposed M0 portable logical schema, canonicalization, export protection and migration boundary
related:
  - DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC
  - RESEARCH-AURORA-M0-PORTABILITY-INTEGRITY-R4-V1
  - DESIGN-AURORA-M0-R4-DECISION-LANDSCAPE
supersedes: []
superseded_by: null
last_reviewed: 2026-08-07
---

# ADR-0005 — M0 Portable Logical State and Export Envelope

## Context

M0 must export, restore and migrate sovereign state without making a database file or one transport serialization the permanent product format.

## Decision drivers

- human inspectability;
- storage independence;
- schema/version validation;
- deterministic digest/MAC input;
- SENSITIVE export confidentiality;
- explicit migration and exit path;
- compatibility with future languages/bindings.

Affected requirements include `REQ-056..066`, `068`, `076`, `083`, `087`, `101`, `104`, `107`.

## Options

- database-native backup as sovereignty format;
- JSON + JSON Schema + canonicalization;
- Protobuf as sole logical/export format;
- deterministic CBOR;
- multiple simultaneous canonical formats.

## Decision

**Proposed:** separate the physical operational store from the portable logical sovereignty format.

For M0:

```text
logical schema        = JSON Schema Draft 2020-12
logical interchange   = JSON
canonical hash/MAC input = RFC 8785 JCS canonical JSON
member/content digest = SHA-256
normal SENSITIVE outer export encryption = age
```

Operational DB backup/snapshot is a separate same-mechanism recovery artifact and does not replace logical export.

Protobuf remains available for future typed bindings but is not the M0 canonical fingerprint/export representation. CBOR is deferred until compact binary representation becomes a demonstrated need.

Migration is application-owned, explicit source-version → target-version transformation with validation before and after, invariant checks and evidence. Physical DDL migrations cannot silently redefine logical state meaning.

The exact archive/container layout inside the encrypted outer file is deferred to R6 because it is reversible and does not own semantics.

## Consequences

### Positive

- SQLite/PostgreSQL migration does not redefine sovereignty format;
- exports are inspectable before encryption;
- stable digest input is explicit rather than serializer-dependent;
- encryption can be replaced without changing logical state.

### Negative

- requires schema + canonicalization discipline in addition to DB schema;
- JCS/I-JSON numeric constraints must be respected;
- conversion code must be tested for semantic compatibility.

### Risks

Treating ordinary non-canonical JSON bytes as digest identity would invalidate stable hashes. Schema version and canonicalization version must therefore be explicit.

## Security / integrity boundary

SHA-256 detects content change/corruption but is not authenticated against an attacker who can rewrite both data and digest. Authenticated trust anchors/HMAC key custody are governed by ADR-0008.

`age` protects the outer export; export keys are not owner authority credentials.

## Compatibility / migration / rollback

Logical export versioning isolates changes. Replacing age only re-envelopes the logical package. Replacing JCS or JSON requires an explicit new logical/export version and migration.

## Validation

Documentary evidence is sufficient for this representation boundary. End-to-end export/restore still participates in SPK-001 and owner recovery semantics in SPK-002.

## Reconsideration triggers

- M0 payloads require exact values incompatible with I-JSON representation;
- embedded/binary channels make JSON size/material parsing cost unacceptable;
- JCS implementation interoperability fails conformance tests;
- age loses maintained interoperable support or a stronger product requirement appears.

## References

- `RESEARCH-AURORA-M0-PORTABILITY-INTEGRITY-R4-V1`
- RFC 8785
- JSON Schema Draft 2020-12
- ADR-0001
