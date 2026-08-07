#!/usr/bin/env python3
from pathlib import Path

REQ = Path('docs/capabilities/CAP-SOVEREIGN-CORE/REQUIREMENTS.md')
text = REQ.read_text(encoding='utf-8')
replacements = {
"When M0 state refers to artifacts or evidence, the Project record SHOULD preserve stable references and provenance rather than duplicating all referenced content inline.":
"When M0 state refers to artifacts or evidence, the Project record MUST preserve stable references and provenance and MUST NOT require duplicating the full referenced content inline.",
"Secret values MUST NOT enter prompts, manifests, general telemetry or general logs through the normal M0 path; M0 MUST NOT require broad secret exposure merely to persist or recover state.":
"M0 MUST NOT require secret values to enter prompts or manifests, and secret values MUST NOT appear in general telemetry or general logs through the normal M0 path merely to persist, inspect or recover state.",
"M0 durable records MUST support explicit current-versus-superseded semantics and retention/archive handling sufficient to avoid ambiguous current state; broader M1 privacy/deletion behavior remains deferred.":
"M0 durable records MUST support explicit current-versus-superseded semantics and retention handling sufficient to avoid ambiguous current state; archive/deletion and the broader M1 privacy lifecycle remain deferred.",
"For security/integrity incidents relevant to M0, Capability MUST support requirements for detection, containment, evidence preservation, scope assessment, recovery and review; detailed incident design belongs to R3.":
"For security/integrity incidents relevant to M0, Capability MUST support detection of materially invalid or corrupt state, containment that prevents unsafe current-state or authority use, evidence preservation, and recovery/review hooks; the broader incident-response program remains deferred.",
"M0 telemetry and cross-component correlation MUST avoid sensitive payloads unless a specific proof need and policy justify them.":
"M0 cross-component observability MUST propagate stable identifiers without sensitive payloads. Sensitive content needed as evidence MUST be governed as separate evidence or artifact data rather than carried in correlation or general telemetry payloads.",
"Domain Events, audit records, logs and telemetry MUST NOT be the sole canonical source of M0 current state unless a later explicit accepted architecture decision defines a reconstruction model that still satisfies canonical ownership and recovery requirements.":
"Domain Events, audit records, logs and telemetry MUST NOT be the sole canonical source of M0 current state or authority.",
"M0 MUST NOT introduce a durable workflow/execution engine merely to satisfy restartable state; such engine MAY be considered only if later accepted requirements demonstrate a need beyond M0 operational-state lifecycle.":
"M0 MUST NOT introduce a durable workflow/execution engine merely because restartable state exists; R4 MAY consider one only if accepted M0 requirements and evidence demonstrate that it is necessary and proportionate to this slice.",
"R3/R4 SHOULD preserve adapter boundaries allowing later evolution from initial local Core without requiring Aurora/Project/authority domain meaning to be rewritten for a new process topology or binding.":
"R3/R4 MUST preserve replaceable adapter boundaries sufficient for later evolution from the initial local Core without requiring Aurora/Project/authority domain meaning to be rewritten for a new process topology or binding; later distributed stages need not be implemented in M0."
}
for old, new in replacements.items():
    count = text.count(old)
    if count != 1:
        raise SystemExit(f'expected one occurrence, got {count}: {old}')
    text = text.replace(old, new)
REQ.write_text(text, encoding='utf-8')

# Guard the exact R2 package shape after refinement.
if text.count('`CAP-SOVEREIGN-CORE-REQ-') < 122:
    raise SystemExit('requirements unexpectedly missing after refinement')
print('Refined 8 R2 requirement statements without changing IDs or coverage.')
