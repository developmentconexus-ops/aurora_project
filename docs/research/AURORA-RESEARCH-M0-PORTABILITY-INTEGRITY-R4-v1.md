---
id: RESEARCH-AURORA-M0-PORTABILITY-INTEGRITY-R4-V1
title: Aurora M0 R4 Research — Portability, Schema, Integrity and Export Protection
document_type: research_report
form: explanation
authority: research
status: current
version: 1.0.0
owners:
  - developmentconexus-ops
source_of_truth_for:
  - M0 R4 portability, serialization, integrity and encrypted-export research through 2026-08-07
related:
  - DOC-AURORA-M0-R4-OPERATOR-AUTHORIZATION
  - DESIGN-AURORA-M0-R4-DECISION-LANDSCAPE
  - DOC-AURORA-CAP-SOVEREIGN-CORE-SPEC
  - DOC-AURORA-CAP-SOVEREIGN-CORE-THREAT-MODEL
source_manifest: AURORA-RESEARCH-M0-PORTABILITY-INTEGRITY-R4-v1.sources.json
review_triggers:
  - export package version change
  - selected canonicalization or encryption format change
  - migration/restore finding
  - selected integrity/key-custody mechanism change
last_reviewed: 2026-08-07
---

# Aurora M0 R4 — Portability, Schema, Integrity and Export Protection

## 1. Research question

Which representation and integrity architecture can make M0 state:

- human-inspectable;
- portable across future storage engines;
- schema-validatable;
- deterministic when hashed;
- detectable when corrupted;
- protected against unauthorized modification where a separate key is available;
- safely encrypted as `SENSITIVE` export/backup material;
- migratable without making one transport or database format the product constitution?

## 2. Key architectural separation

R4 should preserve four distinct layers:

```text
Database physical schema
    ↓ implements local operational persistence

Aurora logical state schema
    ↓ owns portable meaning

Canonical byte representation
    ↓ used only where stable digest/MAC input is needed

Encrypted outer package
    ↓ protects the SENSITIVE export at rest/in transit
```

Using one mechanism for all four would increase lock-in and blur semantics.

## 3. Logical schema: JSON Schema Draft 2020-12

JSON Schema Draft 2020-12 provides a mature, human-readable schema language for JSON validation [S01]. It fits M0 because the logical envelope is modest and operator inspection/debugging matters more than compact binary size.

Proposed M0 use:

```text
export-manifest.schema.json
project-state.schema.json
authority-state.schema.json
audit-record.schema.json
evidence-record.schema.json
```

Each logical schema carries an Aurora-owned schema ID/version. Database migrations may implement those semantics differently, but export/migration compatibility is evaluated against the logical schema and protected invariants.

### Why not make database DDL the portable schema?

Because a future SQLite → PostgreSQL or other store migration would then change the sovereignty format merely because physical indexing/types changed.

### Why not Protobuf as the sovereignty representation?

Protobuf remains a strong future typed binding/transport candidate, but its official documentation explicitly states that protobuf serialization is not canonical and that even deterministic serialization is not a stable canonical fingerprint [S04]. That makes raw protobuf bytes a poor foundation for long-lived integrity identities/hashes.

## 4. Canonical bytes: JCS for hash/MAC input

Ordinary JSON serialization does not guarantee one byte representation for semantically equivalent objects.

RFC 8785 JCS defines a deterministic JSON representation by constraining data to I-JSON-compatible values and recursively sorting object properties [S02]. It is therefore a good fit for:

- manifest digest input;
- protected trust-anchor payloads;
- stable integrity descriptors;
- repeatable migration/evidence hashes.

JCS is Informational rather than Standards Track, so Aurora should treat it as a selected reversible canonicalization mechanism, not constitutional product meaning.

M0 logical schemas should avoid values that do not fit JCS/I-JSON safely. Where exact large integer/decimal semantics are needed in future, the schema should represent them explicitly (for example as decimal strings) instead of relying on ambiguous JSON number precision.

## 5. CBOR alternative

RFC 8949 CBOR is Standards Track and defines deterministic encoding rules [S03]. It is technically stronger as a compact deterministic binary format.

For M0 it provides little current value:

- payload size is not a problem;
- direct human inspection is valuable;
- we already need a readable schema/evidence surface;
- using both JSON and CBOR would create conversion/compatibility obligations without a named need.

Therefore:

```text
CBOR = DEFERRED ALTERNATIVE
reconsider if binary/size/embedded-channel requirements become material
```

## 6. Digest versus authenticated integrity

### 6.1 SHA-256 — corruption/change fingerprint

FIPS 180-4 specifies SHA-2 algorithms including SHA-256 [S05]. A SHA-256 digest is appropriate for:

- artifact/export member content hashes;
- repeatable evidence references;
- accidental corruption detection;
- stable fingerprinting of JCS-canonical logical documents.

But an unkeyed digest is **not** sufficient against an attacker who can modify both the content and the stored digest.

### 6.2 HMAC-SHA-256 — authenticated integrity

FIPS 198-1 defines HMAC as a keyed message authentication mechanism [S06]. NIST is transitioning that guidance toward SP 800-224, but HMAC remains the relevant standard primitive; the currently published SP 800-224 is still an initial public draft [S07].

HMAC-SHA-256 is appropriate where Aurora needs to know that a protected descriptor/anchor was created by a holder of the owner-root-derived integrity key.

Critical rule:

```text
HMAC key MUST NOT live only inside the same rollback/tamper domain as the data it authenticates.
```

Otherwise a store compromise or rollback can defeat the intended trust separation.

### 6.3 HKDF subkey separation

RFC 5869 HKDF derives independent keys from one strong master key using labeled `info` contexts [S08]. If the owner-root design selects one random master root key, Aurora should derive separate subkeys such as:

```text
"aurora/m0/state-integrity/v1"
"aurora/m0/trust-anchor/v1"
"aurora/m0/recovery-envelope/v1"
```

This prevents one key from being reused directly across different cryptographic purposes.

## 7. Logical export package

Proposed M0 logical export layout:

```text
aurora-export/
├── manifest.json
├── aurora.json
├── projects/
│   └── <project-id>/
│       ├── project.json
│       ├── state-revisions.json
│       ├── authority.json
│       ├── audit.json
│       └── evidence-index.json
└── schemas/
    └── logical schema/version descriptors or stable references
```

The exact file split remains an R6 implementation detail. R4 fixes the properties:

- manifest identifies Aurora export-format version and Capability semantic version;
- every required logical member is schema-validatable;
- every member has SHA-256 digest over the defined canonical bytes;
- manifest itself has a deterministic canonical representation;
- unknown incompatible major versions fail clearly;
- protected identity/authority invariants are checked after parse and before restored state becomes governing;
- physical DB internals, WAL files and telemetry are not part of the sovereignty contract.

## 8. Operational backup is not logical export

For a selected store, Aurora may have a fast operational snapshot/backup mechanism. For SQLite that may use the supported Online Backup API; for PostgreSQL it would use PostgreSQL backup mechanisms.

That artifact is useful for same-mechanism recovery but is **not** the product-portable export.

Therefore:

```text
operational backup
!=
logical export
```

Both can coexist.

## 9. Export confidentiality: age

The R3 data classification says export/backup material is `SENSITIVE` minimum. Plain portable JSON should therefore not be the normal at-rest export artifact.

`age` is a current, versioned, multi-platform encrypted-file format and Go library with explicit keys/recipients and non-malleable encrypted files [S09][S10]. It fits the M0 reversible-mechanism profile because:

- it is independent of the operational DB;
- it has implementations beyond Go;
- it can be invoked through library or tool interfaces;
- recipient/key policy can evolve separately from the logical package;
- encryption failure does not change domain semantics.

Recommended architecture:

```text
logical export package
→ deterministic member digests / manifest
→ archive/container representation selected in R6
→ age encrypted outer file
```

Exact archive/container choice (`zip`, `tar`, stream framing, etc.) is not an R4 product decision as long as the logical manifest/schema remain authoritative.

## 10. Recipient/key policy

R4 should **not** use export encryption keys as canonical owner authority.

The export layer needs a recovery policy, but its key can be:

- a dedicated owner export recipient;
- a passphrase recipient for a recovery artifact;
- a future hardware recipient/plugin.

For M0, the smallest portable starting point is a dedicated recovery/export identity protected by the owner-root recovery process. Exact UI/key-storage packaging should be tested with the owner-root security spike rather than decided from age documentation alone.

## 11. Migration semantics

Migration belongs to Aurora, not the database engine.

R4 architectural rule:

```text
version-pair migration
source logical version → target logical version
→ validate source
→ transform
→ validate target
→ verify protected invariants
→ record evidence
→ commit/promote explicitly
```

Store-specific DDL migrations implement physical representation but cannot silently redefine logical meaning.

Rollback rule:

- prior governing state remains available until target validation/promotion succeeds;
- failure produces explicit `MIGRATION_FAILED` evidence;
- no best-effort partial coercion of identity/authority semantics.

Exact migration library/tooling follows the runtime/store decision and is reversible.

## 12. Proposed R4 decisions

Documentary evidence is sufficient to recommend:

```text
R4-Q-SCHEMA-001
→ JSON Schema Draft 2020-12 for portable M0 logical schema
→ JSON as operator-readable logical interchange
→ JCS (RFC 8785) for canonical hash/MAC input
→ Protobuf NOT the M0 sovereignty/fingerprint format
→ CBOR deferred until a binary-size/embedded need exists

R4-Q-EXPORT-001
→ separate operational backup from portable logical export
→ normal portable export protected by age outer encryption

R4-Q-MIGRATE-001
→ application-owned explicit version-pair migrations with source/target validation and protected-invariant checks
```

The following remains coupled to the security spike:

```text
R4-Q-INTEGRITY-001
→ SHA-256 for content digest/fingerprints
→ HMAC-SHA-256 for authenticated trust anchors
→ exact owner-root key custody / rollback-anchor mechanism requires executable proof
```

## 13. Migration/exit path

If JCS or JSON Schema later becomes inadequate:

- the logical schema ID/version explicitly changes;
- migration produces the new logical representation;
- old export readers remain version-bound;
- database state model remains unaffected by the wire/canonicalization change;
- integrity algorithm/version is carried in descriptors rather than assumed globally.

If age is replaced:

- decrypt the old outer container;
- retain the same logical package;
- re-encrypt under a new envelope version.

This is precisely why encryption is outside the logical sovereignty format.

## 14. Limitations

This report does not prove:

- owner-root key custody;
- HMAC rollback protection across crashes;
- safe recovery-key UX;
- actual cross-version migration implementation;
- age key-loss/recovery behavior in Aurora;
- end-to-end restore after malicious/accidental corruption.

Those mechanism-dependent claims belong to Architecture Spike evidence or later implementation tests.
