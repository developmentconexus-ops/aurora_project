//go:build modernc

// Package sqlitedriver selects the disposable SQLite binding used by SPK-001.
package sqlitedriver

import _ "modernc.org/sqlite"

const (
	Name      = "sqlite"
	Candidate = "modernc"
)
