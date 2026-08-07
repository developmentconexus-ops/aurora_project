package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/developmentconexus-ops/aurora_project/spikes/m0-owner-trust-002/internal/trust"
)

type evidence struct {
	GoVersion                  string                          `json:"go_version"`
	GOOS                       string                          `json:"goos"`
	GOARCH                     string                          `json:"goarch"`
	CGOEnabled                 string                          `json:"cgo_enabled"`
	KDF                        trust.KDFMeasurement            `json:"kdf"`
	BootstrapNS                int64                           `json:"bootstrap_ns"`
	PassphraseRotationNS       int64                           `json:"passphrase_rotation_ns"`
	DBCommitNS                 int64                           `json:"db_commit_ns"`
	AnchorWriteNS              int64                           `json:"anchor_write_ns"`
	RecoveryBundleNS           int64                           `json:"recovery_bundle_ns"`
	RestoreBundleNS            int64                           `json:"restore_bundle_ns"`
	OwnerRevalidationNS        int64                           `json:"owner_revalidation_ns"`
	RootArtifactBytes          int64                           `json:"root_artifact_bytes"`
	AnchorArtifactBytes        int64                           `json:"anchor_artifact_bytes"`
	StateDBBytes               int64                           `json:"state_db_bytes"`
	RecoveryBundleBytes        int64                           `json:"recovery_bundle_bytes"`
	RootFingerprintStable      bool                            `json:"root_fingerprint_stable_after_rotation"`
	RestoredClassification     trust.Classification            `json:"restored_classification"`
	PostRevalidationClass      trust.Classification            `json:"post_revalidation_classification"`
	PostRevalidationPermitting bool                            `json:"post_revalidation_permitting"`
	RecoveryStepModel          map[string]int                  `json:"recovery_step_model"`
	SecretPromptModel          map[string]int                  `json:"secret_prompt_model"`
	DiagnosticSamples          map[trust.Classification]string `json:"diagnostic_samples"`
	Limitations                []string                        `json:"limitations"`
}

func main() {
	out := flag.String("out", "", "output JSON path")
	flag.Parse()
	if *out == "" {
		fmt.Fprintln(os.Stderr, "-out is required")
		os.Exit(2)
	}
	if err := run(*out); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(out string) error {
	pass1 := []byte("spk002-evidence-owner-passphrase")
	pass2 := []byte("spk002-evidence-owner-passphrase-rotated")
	now := time.Date(2026, 8, 7, 20, 0, 0, 0, time.UTC)
	expiry := now.Add(3 * time.Hour)
	dir, err := os.MkdirTemp("", "aurora-spk002-evidence-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)
	layout := trust.NewLayout(filepath.Join(dir, "live"))

	kdf := trust.MeasureArgon2id(pass1)
	start := time.Now()
	session, err := trust.Bootstrap(layout, "OWNER-EVIDENCE", pass1, now, expiry)
	if err != nil {
		return fmt.Errorf("bootstrap: %w", err)
	}
	bootstrapNS := time.Since(start).Nanoseconds()

	beforeSession := session
	start = time.Now()
	if err := trust.RotatePassphrase(layout.RootPath, pass1, pass2); err != nil {
		return fmt.Errorf("rotate: %w", err)
	}
	rotationNS := time.Since(start).Nanoseconds()
	rotated, err := trust.Unlock(layout.RootPath, pass2)
	if err != nil {
		return fmt.Errorf("unlock rotated root: %w", err)
	}
	rootStable := trust.SameOwnerRoot(beforeSession, rotated)

	advance, err := trust.Advance(layout, rotated, trust.Mutation{
		ExpectedGeneration: 1,
		AuthorityRevision:  "AUTH-EVIDENCE-2",
		AuthorityStatus:    trust.StatusActive,
		ExpiresAt:          expiry,
		StateSummary:       "evidence generation two",
	}, now.Add(time.Minute), nil)
	if err != nil {
		return fmt.Errorf("advance: %w", err)
	}

	bundle := filepath.Join(dir, "bundle")
	start = time.Now()
	if err := trust.CreateRecoveryBundle(layout, bundle); err != nil {
		return fmt.Errorf("bundle: %w", err)
	}
	bundleNS := time.Since(start).Nanoseconds()

	restored := trust.NewLayout(filepath.Join(dir, "restored"))
	start = time.Now()
	if err := trust.RestoreRecoveryBundle(bundle, restored); err != nil {
		return fmt.Errorf("restore: %w", err)
	}
	restoreNS := time.Since(start).Nanoseconds()
	restoredSession, err := trust.Unlock(restored.RootPath, pass2)
	if err != nil {
		return fmt.Errorf("unlock restored root: %w", err)
	}
	restoredEval, err := trust.Evaluate(restored, restoredSession, now.Add(2*time.Minute))
	if err != nil {
		return fmt.Errorf("evaluate restored: %w", err)
	}
	start = time.Now()
	if err := trust.Revalidate(restored, restoredSession, trust.ActorOwner, trust.Revalidation{
		AuthorityRevision: "AUTH-EVIDENCE-REVALIDATED-3",
		ExpiresAt:         expiry.Add(time.Hour),
	}, now.Add(3*time.Minute)); err != nil {
		return fmt.Errorf("revalidate: %w", err)
	}
	revalidationNS := time.Since(start).Nanoseconds()
	postEval, err := trust.Evaluate(restored, restoredSession, now.Add(4*time.Minute))
	if err != nil {
		return fmt.Errorf("evaluate post revalidation: %w", err)
	}

	rootSize, err := fileSize(layout.RootPath)
	if err != nil {
		return err
	}
	anchorSize, err := fileSize(layout.AnchorPath)
	if err != nil {
		return err
	}
	dbSize, err := fileSize(layout.DBPath)
	if err != nil {
		return err
	}
	bundleBytes, err := treeSize(bundle)
	if err != nil {
		return err
	}

	classes := []trust.Classification{
		trust.ClassNormal,
		trust.ClassStateRollback,
		trust.ClassAnchorLag,
		trust.ClassInvalidDBMAC,
		trust.ClassInvalidAnchorMAC,
		trust.ClassMissingAnchor,
		trust.ClassMissingRoot,
		trust.ClassUnlockFailed,
		trust.ClassTimeUntrusted,
		trust.ClassRevalidationRequired,
	}
	diagnostics := make(map[trust.Classification]string, len(classes))
	for _, class := range classes {
		diagnostics[class] = trust.DiagnosticFor(class)
	}

	e := evidence{
		GoVersion:                  runtime.Version(),
		GOOS:                       runtime.GOOS,
		GOARCH:                     runtime.GOARCH,
		CGOEnabled:                 os.Getenv("CGO_ENABLED"),
		KDF:                        kdf,
		BootstrapNS:                bootstrapNS,
		PassphraseRotationNS:       rotationNS,
		DBCommitNS:                 advance.DBCommitNS,
		AnchorWriteNS:              advance.AnchorWriteNS,
		RecoveryBundleNS:           bundleNS,
		RestoreBundleNS:            restoreNS,
		OwnerRevalidationNS:        revalidationNS,
		RootArtifactBytes:          rootSize,
		AnchorArtifactBytes:        anchorSize,
		StateDBBytes:               dbSize,
		RecoveryBundleBytes:        bundleBytes,
		RootFingerprintStable:      rootStable,
		RestoredClassification:     restoredEval.Classification,
		PostRevalidationClass:      postEval.Classification,
		PostRevalidationPermitting: postEval.Permitting,
		RecoveryStepModel: map[string]int{
			"anchor_lag_owner_reconcile_steps":            3,
			"historical_restore_owner_revalidation_steps": 4,
			"fresh_machine_root_recovery_steps":           4,
		},
		SecretPromptModel: map[string]int{
			"bootstrap_passphrase_prompts":              1,
			"normal_unlock_prompts":                    1,
			"passphrase_rotation_authenticated_prompts": 2,
			"post_restore_owner_revalidation_prompts":   1,
		},
		DiagnosticSamples: diagnostics,
		Limitations: []string{
			"process-kill tests do not emulate all physical power-loss or storage-controller failures",
			"purely local trust cannot defeat replay of all trust artifacts plus compromise of owner secrets/runtime",
			"Argon2 observed heap peak uses 1ms Go HeapInuse polling and is not operating-system RSS",
		},
	}
	payload, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return err
	}
	payload = append(payload, '\n')
	if err := os.MkdirAll(filepath.Dir(out), 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(out, payload, 0o600); err != nil {
		return err
	}
	fmt.Print(string(payload))
	return nil
}

func fileSize(path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

func treeSize(root string) (int64, error) {
	var total int64
	err := filepath.Walk(root, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			total += info.Size()
		}
		return nil
	})
	return total, err
}
