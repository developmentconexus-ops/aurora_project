package store

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func acceptedFixture(t *testing.T, path string) Snapshot {
	t.Helper()
	initial := Snapshot{
		SchemaVersion:     1,
		AuroraID:          "AURORA-SPIKE-001",
		ProjectID:         "PROJECT-SPIKE-001",
		CurrentRevision:   1,
		AuthorityRevision: "AUTH-1",
		StateKind:         "ACTIVE",
		StateSummary:      "revision one",
	}
	if err := Bootstrap(path, initial); err != nil {
		t.Fatalf("bootstrap: %v", err)
	}
	if err := ApplyTransition(path, TransitionInput{
		AttemptID:          "ATTEMPT-2",
		ProjectID:          initial.ProjectID,
		ExpectedRevision:   1,
		NewRevision:        2,
		AuthorityRevision:  "AUTH-1",
		StateKind:          "ACTIVE",
		StateSummary:       "revision two",
		AuditID:            "AUDIT-2",
		EvidenceID:         "EVIDENCE-2",
		EvidenceRef:        "sha256:fixture-2",
	}, nil); err != nil {
		t.Fatalf("transition: %v", err)
	}
	want := initial
	want.CurrentRevision = 2
	want.StateSummary = "revision two"
	return want
}

func TestCheckpointPreservesGoverningState(t *testing.T) {
	path := filepath.Join(t.TempDir(), "aurora.db")
	want := acceptedFixture(t, path)
	if err := Checkpoint(path, nil); err != nil {
		t.Fatalf("checkpoint: %v", err)
	}
	got, err := Inspect(path)
	if err != nil {
		t.Fatalf("inspect after checkpoint: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("checkpoint changed state\nwant: %#v\n got: %#v", want, got)
	}
}

func TestSupportedBackupRestoreAfterOriginalDestroyedAndIdentityCollision(t *testing.T) {
	dir := t.TempDir()
	workDir := filepath.Join(dir, "work")
	backupDir := filepath.Join(dir, "backup")
	restoreDir := filepath.Join(dir, "restored")
	source := filepath.Join(workDir, "aurora.db")
	backup := filepath.Join(backupDir, "backup.db")
	restored := filepath.Join(restoreDir, "aurora.db")
	want := acceptedFixture(t, source)

	keeper, err := openExistingDB(source)
	if err != nil {
		t.Fatalf("open active keeper: %v", err)
	}
	if err := SupportedBackup(source, backup, nil); err != nil {
		keeper.Close()
		t.Fatalf("online backup: %v", err)
	}
	if err := keeper.Close(); err != nil {
		t.Fatalf("close active keeper: %v", err)
	}
	if err := os.RemoveAll(workDir); err != nil {
		t.Fatalf("destroy original working directory: %v", err)
	}
	if _, err := os.Stat(source); !os.IsNotExist(err) {
		t.Fatalf("original store still exists after destroy: %v", err)
	}

	if err := RestoreBackup(backup, restored); err != nil {
		t.Fatalf("fresh restore: %v", err)
	}
	got, err := Inspect(restored)
	if err != nil {
		t.Fatalf("inspect restored: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("restored snapshot mismatch\nwant: %#v\n got: %#v", want, got)
	}

	collision := filepath.Join(dir, "collision.db")
	other := want
	other.AuroraID = "AURORA-OTHER"
	other.ProjectID = "PROJECT-OTHER"
	other.CurrentRevision = 1
	other.StateSummary = "other state"
	if err := Bootstrap(collision, other); err != nil {
		t.Fatalf("bootstrap collision target: %v", err)
	}
	before, _ := Inspect(collision)
	if err := RestoreBackup(backup, collision); !errors.Is(err, ErrIdentityCollision) {
		t.Fatalf("collision error = %v, want ErrIdentityCollision", err)
	}
	after, err := Inspect(collision)
	if err != nil {
		t.Fatalf("inspect collision target: %v", err)
	}
	if !reflect.DeepEqual(after, before) {
		t.Fatalf("collision restore mutated target\nbefore: %#v\n after: %#v", before, after)
	}
}

func TestCorruptAndUnsupportedSchemaFailExplicitly(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "aurora.db")
	_ = acceptedFixture(t, path)

	corrupt := filepath.Join(dir, "corrupt.db")
	copyFileForTest(t, path, corrupt)
	info, err := os.Stat(corrupt)
	if err != nil {
		t.Fatalf("stat corrupt fixture: %v", err)
	}
	if err := os.Truncate(corrupt, info.Size()/2); err != nil {
		t.Fatalf("truncate fixture: %v", err)
	}
	if _, err := Inspect(corrupt); err == nil {
		t.Fatal("corrupt database was accepted")
	}

	incompatible := filepath.Join(dir, "incompatible.db")
	copyFileForTest(t, path, incompatible)
	db, err := openExistingDB(incompatible)
	if err != nil {
		t.Fatalf("open incompatible fixture: %v", err)
	}
	if _, err := db.Exec(`UPDATE meta SET value = '999' WHERE key = 'schema_version'`); err != nil {
		db.Close()
		t.Fatalf("set incompatible schema: %v", err)
	}
	if err := db.Close(); err != nil {
		t.Fatalf("close incompatible fixture: %v", err)
	}
	if _, err := Inspect(incompatible); !errors.Is(err, ErrIncompatibleSchema) {
		t.Fatalf("incompatible schema error = %v, want ErrIncompatibleSchema", err)
	}
}

func TestMigrationV1ToV2PreservesProtectedSemantics(t *testing.T) {
	path := filepath.Join(t.TempDir(), "aurora.db")
	want := acceptedFixture(t, path)
	countsBefore, err := InspectCounts(path)
	if err != nil {
		t.Fatalf("counts before migration: %v", err)
	}
	if err := MigrateV1ToV2(path); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	got, err := Inspect(path)
	if err != nil {
		t.Fatalf("inspect migrated store: %v", err)
	}
	want.SchemaVersion = 2
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("migration changed protected semantics\nwant: %#v\n got: %#v", want, got)
	}
	countsAfter, err := InspectCounts(path)
	if err != nil {
		t.Fatalf("counts after migration: %v", err)
	}
	if !reflect.DeepEqual(countsAfter, countsBefore) {
		t.Fatalf("migration changed transition/audit evidence counts\nbefore: %#v\n after: %#v", countsBefore, countsAfter)
	}
}

func copyFileForTest(t *testing.T, src, dst string) {
	t.Helper()
	in, err := os.Open(src)
	if err != nil {
		t.Fatalf("open source copy: %v", err)
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		t.Fatalf("create copy: %v", err)
	}
	if _, err := io.Copy(out, in); err != nil {
		out.Close()
		t.Fatalf("copy fixture: %v", err)
	}
	if err := out.Close(); err != nil {
		t.Fatalf("close copy: %v", err)
	}
}
