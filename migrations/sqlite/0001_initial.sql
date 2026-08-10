-- M0 physical migration 0001. Runtime migration 1 is kept semantically equivalent to this file.
CREATE TABLE IF NOT EXISTS authority_revisions (
  authority_revision INTEGER PRIMARY KEY,
  predecessor_revision INTEGER NULL,
  authority_state_json TEXT NOT NULL,
  changed_by TEXT NOT NULL,
  changed_at TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS core_state (
  singleton_key TEXT PRIMARY KEY CHECK(singleton_key = 'core'),
  aurora_id TEXT NOT NULL UNIQUE,
  owner_operator_id TEXT NOT NULL,
  logical_schema_version INTEGER NOT NULL,
  current_authority_revision INTEGER NOT NULL,
  governing_generation INTEGER NOT NULL,
  governing_descriptor_hmac BLOB NOT NULL,
  created_at TEXT NOT NULL,
  updated_at TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS projects (
  project_id TEXT PRIMARY KEY,
  display_label TEXT NOT NULL,
  objective_summary TEXT NOT NULL,
  current_state_revision INTEGER NULL,
  created_at TEXT NOT NULL,
  updated_at TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS project_state_revisions (
  project_id TEXT NOT NULL,
  state_revision INTEGER NOT NULL,
  predecessor_revision INTEGER NULL,
  state_schema_version TEXT NOT NULL,
  state_kind TEXT NOT NULL,
  state_summary TEXT NOT NULL,
  state_payload_json TEXT NULL,
  accepted_intent_ref TEXT NULL,
  proposed_next_action_json TEXT NULL,
  accepted_by_actor TEXT NOT NULL,
  accepted_at TEXT NOT NULL,
  transition_attempt_id TEXT NOT NULL UNIQUE,
  PRIMARY KEY(project_id, state_revision),
  FOREIGN KEY(project_id) REFERENCES projects(project_id)
);
CREATE TABLE IF NOT EXISTS transition_attempts (
  attempt_id TEXT PRIMARY KEY,
  project_id TEXT NOT NULL,
  actor_id TEXT NOT NULL,
  requested_at TEXT NOT NULL,
  expected_state_revision INTEGER NULL,
  requested_state_json TEXT NOT NULL,
  proposed_next_action_json TEXT NULL,
  authority_evaluation_ref TEXT NULL,
  result TEXT NOT NULL,
  reason TEXT NOT NULL,
  accepted_state_revision INTEGER NULL,
  FOREIGN KEY(project_id) REFERENCES projects(project_id)
);
CREATE TABLE IF NOT EXISTS records (
  record_id TEXT PRIMARY KEY,
  kind TEXT NOT NULL,
  operation_id TEXT NOT NULL,
  project_id TEXT NULL,
  state_revision INTEGER NULL,
  authority_revision INTEGER NULL,
  outcome TEXT NOT NULL,
  reason TEXT NOT NULL,
  details_json TEXT NOT NULL,
  created_at TEXT NOT NULL
);
