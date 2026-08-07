// Package store contains disposable SPK-001 persistence experiment code.
//
// DISCARD: this package is executable evidence, not Aurora Core production code.
package store

// Snapshot is the minimum logical state the spike must recover across process death.
type Snapshot struct {
	SchemaVersion     int    `json:"schema_version"`
	AuroraID          string `json:"aurora_id"`
	ProjectID         string `json:"project_id"`
	CurrentRevision   int64  `json:"current_revision"`
	AuthorityRevision string `json:"authority_revision"`
	StateKind         string `json:"state_kind"`
	StateSummary      string `json:"state_summary"`
}
