package store

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var ErrIdentityCollision = errors.New("restore identity collision")

const (
	bootstrapSchemaVersion = 1
	currentSchemaVersion   = 2
)

// Checkpoint performs an explicit WAL checkpoint around observable hook points.
func Checkpoint(path string, hook FaultHook) error {
	db, err := openExistingDB(path)
	if err != nil {
		return err
	}
	defer db.Close()
	callHook(hook, "before_checkpoint")
	var busy, logFrames, checkpointed int
	if err := db.QueryRow(`PRAGMA wal_checkpoint(TRUNCATE)`).Scan(&busy, &logFrames, &checkpointed); err != nil {
		return fmt.Errorf("wal checkpoint: %w", err)
	}
	if busy != 0 {
		return fmt.Errorf("wal checkpoint remained busy: log=%d checkpointed=%d", logFrames, checkpointed)
	}
	callHook(hook, "after_checkpoint")
	return nil
}

// SupportedBackup creates a standalone SQLite-consistent backup using VACUUM INTO.
// The final destination is published only after the temporary database validates.
func SupportedBackup(source, destination string, hook FaultHook) error {
	if source == destination {
		return errors.New("backup destination must differ from source")
	}
	if err := os.MkdirAll(filepath.Dir(destination), 0o755); err != nil {
		return fmt.Errorf("mkdir backup dir: %w", err)
	}
	tmp := destination + ".partial"
	_ = os.Remove(tmp)
	_ = os.Remove(destination)

	db, err := openExistingDB(source)
	if err != nil {
		return err
	}
	callHook(hook, "before_backup")
	query := "VACUUM INTO '" + strings.ReplaceAll(filepath.ToSlash(tmp), "'", "''") + "'"
	if _, err := db.Exec(query); err != nil {
		db.Close()
		return fmt.Errorf("vacuum into backup: %w", err)
	}
	if err := db.Close(); err != nil {
		return fmt.Errorf("close backup source: %w", err)
	}
	callHook(hook, "after_backup_sql")
	if _, err := Inspect(tmp); err != nil {
		return fmt.Errorf("validate backup temp: %w", err)
	}
	if err := os.Rename(tmp, destination); err != nil {
		return fmt.Errorf("publish backup: %w", err)
	}
	callHook(hook, "after_backup_publish")
	return nil
}

// RestoreBackup validates a standalone backup before publishing it as the destination.
// Existing stores are never overwritten; a different Aurora identity is an explicit collision.
func RestoreBackup(backup, destination string) error {
	backupSnapshot, err := Inspect(backup)
	if err != nil {
		return fmt.Errorf("inspect backup: %w", err)
	}
	if _, err := os.Stat(destination); err == nil {
		existing, inspectErr := Inspect(destination)
		if inspectErr != nil {
			return fmt.Errorf("inspect restore target: %w", inspectErr)
		}
		if existing.AuroraID != backupSnapshot.AuroraID {
			return fmt.Errorf("%w: target=%s backup=%s", ErrIdentityCollision, existing.AuroraID, backupSnapshot.AuroraID)
		}
		return ErrAlreadyExists
	} else if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("stat restore target: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(destination), 0o755); err != nil {
		return fmt.Errorf("mkdir restore dir: %w", err)
	}
	tmp := destination + ".partial"
	_ = os.Remove(tmp)
	if err := copyStandaloneDB(backup, tmp); err != nil {
		return err
	}
	if _, err := Inspect(tmp); err != nil {
		return fmt.Errorf("validate restore temp: %w", err)
	}
	if err := os.Rename(tmp, destination); err != nil {
		return fmt.Errorf("publish restore: %w", err)
	}
	return nil
}

// MigrateV1ToV2 is a deliberately tiny version-pair fixture proving that migration
// can change logical schema version without changing protected identity/state semantics.
func MigrateV1ToV2(path string) error {
	db, err := openExistingDB(path)
	if err != nil {
		return err
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("begin migration: %w", err)
	}
	defer tx.Rollback()
	var raw string
	if err := tx.QueryRow(`SELECT value FROM meta WHERE key = 'schema_version'`).Scan(&raw); err != nil {
		return fmt.Errorf("read migration schema version: %w", err)
	}
	if raw != "1" {
		return fmt.Errorf("migration requires v1 source, got %s", raw)
	}
	if _, err := tx.Exec(`CREATE TABLE migration_metadata (key TEXT PRIMARY KEY, value TEXT NOT NULL)`); err != nil {
		return fmt.Errorf("create v2 migration metadata: %w", err)
	}
	if _, err := tx.Exec(`INSERT INTO migration_metadata(key, value) VALUES ('migration', 'v1-to-v2')`); err != nil {
		return fmt.Errorf("record migration: %w", err)
	}
	if _, err := tx.Exec(`UPDATE meta SET value = '2' WHERE key = 'schema_version'`); err != nil {
		return fmt.Errorf("update schema version: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit migration: %w", err)
	}
	return nil
}

func copyStandaloneDB(source, destination string) error {
	in, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("open backup source: %w", err)
	}
	defer in.Close()
	out, err := os.OpenFile(destination, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o600)
	if err != nil {
		return fmt.Errorf("create restore temp: %w", err)
	}
	ok := false
	defer func() {
		_ = out.Close()
		if !ok {
			_ = os.Remove(destination)
		}
	}()
	if _, err := io.Copy(out, in); err != nil {
		return fmt.Errorf("copy backup: %w", err)
	}
	if err := out.Sync(); err != nil {
		return fmt.Errorf("sync restore temp: %w", err)
	}
	if err := out.Close(); err != nil {
		return fmt.Errorf("close restore temp: %w", err)
	}
	ok = true
	return nil
}
