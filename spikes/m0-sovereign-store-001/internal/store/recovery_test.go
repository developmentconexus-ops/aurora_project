package store

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestProcessKillAroundCheckpointPreservesCommittedState(t *testing.T) {
	for _, point := range []string{"before_checkpoint", "after_checkpoint"} {
		t.Run(point, func(t *testing.T) {
			dir := t.TempDir()
			path := filepath.Join(dir, "aurora.db")
			want := acceptedFixture(t, path)
			cmd, _ := startHelperAtPoint(t, "checkpoint", path, "", point)
			killHelper(t, cmd)
			got, err := Inspect(path)
			if err != nil {
				t.Fatalf("inspect after checkpoint kill at %s: %v", point, err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("checkpoint kill changed governing state at %s\nwant: %#v\n got: %#v", point, want, got)
			}
		})
	}
}

func TestInterruptedBackupNeverPublishesPartialAsFinal(t *testing.T) {
	for _, point := range []string{"before_backup", "after_backup_sql", "after_backup_publish"} {
		t.Run(point, func(t *testing.T) {
			dir := t.TempDir()
			source := filepath.Join(dir, "source.db")
			destination := filepath.Join(dir, "backup.db")
			want := acceptedFixture(t, source)
			cmd, _ := startHelperAtPoint(t, "backup", source, destination, point)
			killHelper(t, cmd)

			_, statErr := os.Stat(destination)
			if point != "after_backup_publish" {
				if statErr == nil {
					t.Fatalf("backup was published before safe publication point %s", point)
				}
				if !os.IsNotExist(statErr) {
					t.Fatalf("stat interrupted backup: %v", statErr)
				}
				return
			}
			if statErr != nil {
				t.Fatalf("published backup missing after publish point: %v", statErr)
			}
			got, err := Inspect(destination)
			if err != nil {
				t.Fatalf("published backup invalid: %v", err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("published backup changed state\nwant: %#v\n got: %#v", want, got)
			}
		})
	}
}

func TestNaiveMainFileCopyWhileCommittedRevisionLivesInWALIsNotCurrentBackup(t *testing.T) {
	dir := t.TempDir()
	source := filepath.Join(dir, "source.db")
	naive := filepath.Join(dir, "naive-main-only.db")
	initial := Snapshot{
		SchemaVersion:     1,
		AuroraID:          "AURORA-SPIKE-001",
		ProjectID:         "PROJECT-SPIKE-001",
		CurrentRevision:   1,
		AuthorityRevision: "AUTH-1",
		StateKind:         "ACTIVE",
		StateSummary:      "revision one",
	}
	if err := Bootstrap(source, initial); err != nil {
		t.Fatalf("bootstrap: %v", err)
	}
	// Put the bootstrap baseline in the main database before creating a new WAL commit.
	if err := Checkpoint(source, nil); err != nil {
		t.Fatalf("baseline checkpoint: %v", err)
	}

	cmd, _ := startHelperAtPoint(t, "transition", source, "", "after_commit")
	copyFileForTest(t, source, naive)

	naiveSnapshot, naiveErr := Inspect(naive)
	if naiveErr == nil && naiveSnapshot.CurrentRevision == 2 {
		killHelper(t, cmd)
		t.Fatalf("naive main-file copy unexpectedly contained current WAL revision: %#v", naiveSnapshot)
	}
	killHelper(t, cmd)

	got, err := Inspect(source)
	if err != nil {
		t.Fatalf("inspect original after kill: %v", err)
	}
	if got.CurrentRevision != 2 || got.StateSummary != "revision two after crash boundary" {
		t.Fatalf("original did not recover committed WAL state: %#v", got)
	}
}
