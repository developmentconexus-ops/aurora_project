package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/developmentconexus-ops/aurora_project/spikes/m0-sovereign-store-001/internal/store"
)

type metrics struct {
	Candidate            string         `json:"candidate"`
	GoVersion            string         `json:"go_version"`
	GOOS                 string         `json:"goos"`
	GOARCH               string         `json:"goarch"`
	CGOEnabled           string         `json:"cgo_enabled"`
	SQLiteVersion        string         `json:"sqlite_version"`
	IntegrityCheck       string         `json:"integrity_check"`
	BaselineBytes        int            `json:"baseline_state_summary_bytes"`
	InspectNS            int64          `json:"inspect_ns"`
	TransitionCount      int            `json:"transition_count"`
	TransitionP50NS      int64          `json:"transition_p50_ns"`
	TransitionP95NS      int64          `json:"transition_p95_ns"`
	TransitionMaxNS      int64          `json:"transition_max_ns"`
	BackupNS             int64          `json:"backup_ns"`
	RestoreNS            int64          `json:"restore_ns"`
	DatabaseBytes        int64          `json:"database_bytes"`
	BackupBytes          int64          `json:"backup_bytes"`
	BackupSHA256         string         `json:"backup_sha256"`
	RestoredSHA256       string         `json:"restored_sha256"`
	RecoveredSnapshot    store.Snapshot `json:"recovered_snapshot"`
}

func main() {
	out := flag.String("out", "", "write JSON evidence to this path")
	transitions := flag.Int("transitions", 100, "number of tiny accepted transitions")
	baselineBytes := flag.Int("baseline-bytes", 0, "deterministic initial state-summary payload size")
	flag.Parse()
	if *out == "" {
		fmt.Fprintln(os.Stderr, "-out is required")
		os.Exit(2)
	}
	if *transitions < 1 || *baselineBytes < 0 {
		fmt.Fprintln(os.Stderr, "-transitions must be positive and -baseline-bytes non-negative")
		os.Exit(2)
	}
	if err := run(*out, *transitions, *baselineBytes); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(out string, n, baselineBytes int) error {
	dir, err := os.MkdirTemp("", "aurora-spk001-metrics-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)
	dbPath := filepath.Join(dir, "aurora.db")
	backupPath := filepath.Join(dir, "backup.db")
	restorePath := filepath.Join(dir, "restore", "aurora.db")
	baselineSummary := "metrics baseline"
	if baselineBytes > 0 {
		baselineSummary = strings.Repeat("M", baselineBytes)
	}
	initial := store.Snapshot{
		SchemaVersion:     1,
		AuroraID:          "AURORA-METRICS",
		ProjectID:         "PROJECT-METRICS",
		CurrentRevision:   1,
		AuthorityRevision: "AUTH-1",
		StateKind:         "ACTIVE",
		StateSummary:      baselineSummary,
	}
	if err := store.Bootstrap(dbPath, initial); err != nil {
		return fmt.Errorf("bootstrap: %w", err)
	}

	start := time.Now()
	if _, err := store.Inspect(dbPath); err != nil {
		return fmt.Errorf("inspect: %w", err)
	}
	inspectNS := time.Since(start).Nanoseconds()

	durations := make([]int64, 0, n)
	for i := 0; i < n; i++ {
		expected := int64(i + 1)
		newRevision := expected + 1
		start = time.Now()
		err := store.ApplyTransition(dbPath, store.TransitionInput{
			AttemptID:          fmt.Sprintf("ATTEMPT-%06d", newRevision),
			ProjectID:          initial.ProjectID,
			ExpectedRevision:   expected,
			NewRevision:        newRevision,
			AuthorityRevision:  "AUTH-1",
			StateKind:          "ACTIVE",
			StateSummary:       fmt.Sprintf("metrics revision %d", newRevision),
			AuditID:            fmt.Sprintf("AUDIT-%06d", newRevision),
			EvidenceID:         fmt.Sprintf("EVIDENCE-%06d", newRevision),
			EvidenceRef:        fmt.Sprintf("sha256:metrics-%06d", newRevision),
		}, nil)
		if err != nil {
			return fmt.Errorf("transition %d: %w", newRevision, err)
		}
		durations = append(durations, time.Since(start).Nanoseconds())
	}
	sort.Slice(durations, func(i, j int) bool { return durations[i] < durations[j] })

	start = time.Now()
	if err := store.SupportedBackup(dbPath, backupPath, nil); err != nil {
		return fmt.Errorf("backup: %w", err)
	}
	backupNS := time.Since(start).Nanoseconds()
	backupHash, err := sha256File(backupPath)
	if err != nil {
		return fmt.Errorf("hash backup: %w", err)
	}
	start = time.Now()
	if err := store.RestoreBackup(backupPath, restorePath); err != nil {
		return fmt.Errorf("restore: %w", err)
	}
	restoreNS := time.Since(start).Nanoseconds()
	restoredHash, err := sha256File(restorePath)
	if err != nil {
		return fmt.Errorf("hash restored database: %w", err)
	}
	recovered, err := store.Inspect(restorePath)
	if err != nil {
		return fmt.Errorf("inspect restored state: %w", err)
	}

	sqliteVersion, err := store.SQLiteVersion(dbPath)
	if err != nil {
		return err
	}
	integrity, err := store.IntegrityStatus(dbPath)
	if err != nil {
		return err
	}
	dbInfo, err := os.Stat(dbPath)
	if err != nil {
		return err
	}
	backupInfo, err := os.Stat(backupPath)
	if err != nil {
		return err
	}

	m := metrics{
		Candidate:         store.Candidate(),
		GoVersion:         runtime.Version(),
		GOOS:              runtime.GOOS,
		GOARCH:            runtime.GOARCH,
		CGOEnabled:        os.Getenv("CGO_ENABLED"),
		SQLiteVersion:     sqliteVersion,
		IntegrityCheck:    integrity,
		BaselineBytes:     baselineBytes,
		InspectNS:         inspectNS,
		TransitionCount:   n,
		TransitionP50NS:   percentile(durations, 0.50),
		TransitionP95NS:   percentile(durations, 0.95),
		TransitionMaxNS:   durations[len(durations)-1],
		BackupNS:          backupNS,
		RestoreNS:         restoreNS,
		DatabaseBytes:     dbInfo.Size(),
		BackupBytes:       backupInfo.Size(),
		BackupSHA256:      backupHash,
		RestoredSHA256:    restoredHash,
		RecoveredSnapshot: recovered,
	}
	if err := os.MkdirAll(filepath.Dir(out), 0o755); err != nil {
		return err
	}
	payload, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	payload = append(payload, '\n')
	if err := os.WriteFile(out, payload, 0o600); err != nil {
		return err
	}
	fmt.Print(string(payload))
	return nil
}

func percentile(sorted []int64, p float64) int64 {
	if len(sorted) == 0 {
		return 0
	}
	idx := int(float64(len(sorted)-1) * p)
	return sorted[idx]
}

func sha256File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
