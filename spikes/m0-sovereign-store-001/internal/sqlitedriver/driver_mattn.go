//go:build mattn

// Package sqlitedriver selects the disposable SQLite binding used by SPK-001.
package sqlitedriver

import _ "github.com/mattn/go-sqlite3"

const (
	Name      = "sqlite3"
	Candidate = "mattn"
)
