package portability

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/identity"
)

var ErrMigrationSemanticInvariant = errors.New("sovereign migration semantic invariant violation")

type MigrationResult struct {
	SourceVersion int      `json:"source_version"`
	TargetVersion int      `json:"target_version"`
	Document      Document `json:"-"`
}

type documentV0 struct {
	Format        string             `json:"format"`
	Version       int                `json:"version"`
	CreatedAt     time.Time          `json:"created_at"`
	Generation    uint64             `json:"generation"`
	Aurora        identity.AuroraIdentity `json:"aurora"`
	Projects      []ProjectBundle    `json:"projects"`
	Authority     AuthorityBundle    `json:"authority"`
	Attempts      []TransitionAttempt `json:"transition_attempts"`
	Records       []Record           `json:"records"`
	OwnerRecovery RootEnvelope       `json:"owner_recovery"`
}

func Migrate(raw []byte) (MigrationResult, error) {
	var header struct {
		Format  string `json:"format"`
		Version *int   `json:"version"`
	}
	if err := json.Unmarshal(raw, &header); err != nil {
		return MigrationResult{}, fmt.Errorf("read sovereign export header: %w", err)
	}
	if header.Format != FormatV1 || header.Version == nil {
		return MigrationResult{}, fmt.Errorf("%w: invalid format or missing version", ErrMigrationSemanticInvariant)
	}
	if *header.Version != 0 {
		return MigrationResult{}, fmt.Errorf("%w: source version %d", ErrUnsupportedVersion, *header.Version)
	}
	return migrateV0ToV1(raw)
}

func migrateV0ToV1(raw []byte) (MigrationResult, error) {
	var source documentV0
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&source); err != nil {
		return MigrationResult{}, fmt.Errorf("decode v0 export: %w", err)
	}
	if err := ensureJSONEOF(decoder); err != nil {
		return MigrationResult{}, err
	}
	if source.Format != FormatV1 || source.Version != 0 {
		return MigrationResult{}, fmt.Errorf("%w: invalid v0 header", ErrMigrationSemanticInvariant)
	}
	if err := validateV0Semantics(source); err != nil {
		return MigrationResult{}, err
	}

	exportID, err := deterministicMigrationExportID(source)
	if err != nil {
		return MigrationResult{}, err
	}
	doc := Document{
		Format:              FormatV1,
		Version:             1,
		ExportID:            exportID,
		CreatedAt:           source.CreatedAt,
		GoverningGeneration: source.Generation,
		Aurora:              source.Aurora,
		Projects:            source.Projects,
		Authority:           source.Authority,
		Attempts:            source.Attempts,
		Records:             source.Records,
		OwnerRecovery:       source.OwnerRecovery,
	}
	if err := Finalize(&doc); err != nil {
		return MigrationResult{}, fmt.Errorf("finalize v1 export: %w", err)
	}
	return MigrationResult{SourceVersion: 0, TargetVersion: 1, Document: doc}, nil
}

func ensureJSONEOF(decoder *json.Decoder) error {
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		if err == nil {
			return fmt.Errorf("%w: trailing JSON value", ErrMigrationSemanticInvariant)
		}
		return fmt.Errorf("decode trailing JSON: %w", err)
	}
	return nil
}

func deterministicMigrationExportID(source documentV0) (string, error) {
	raw, err := json.Marshal(source)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(raw)
	return "EXP-MIG-" + hex.EncodeToString(sum[:16]), nil
}

func validateV0Semantics(source documentV0) error {
	fail := func(format string, args ...any) error {
		return fmt.Errorf("%w: %s", ErrMigrationSemanticInvariant, fmt.Sprintf(format, args...))
	}
	if source.Generation == 0 {
		return fail("governing generation must be positive")
	}
	if source.Aurora.AuroraID == "" || source.Aurora.OwnerOperatorID == "" {
		return fail("Aurora and owner identities must be present")
	}

	projectIDs := make(map[string]struct{}, len(source.Projects))
	projectRevisions := make(map[string]map[uint64]struct{}, len(source.Projects))
	for _, bundle := range source.Projects {
		pid := string(bundle.Project.ProjectID)
		if pid == "" {
			return fail("project identity is empty")
		}
		if _, exists := projectIDs[pid]; exists {
			return fail("duplicate project identity %s", pid)
		}
		projectIDs[pid] = struct{}{}
		revisions := make(map[uint64]struct{}, len(bundle.Revisions))
		for _, revision := range bundle.Revisions {
			if string(revision.ProjectID) != pid {
				return fail("project revision identity drift: bundle=%s revision=%s", pid, revision.ProjectID)
			}
			r := uint64(revision.Revision)
			if r == 0 {
				return fail("project %s contains revision zero", pid)
			}
			if _, exists := revisions[r]; exists {
				return fail("project %s contains duplicate revision %d", pid, r)
			}
			revisions[r] = struct{}{}
			if revision.ProposedNextAction != nil && revision.ProposedNextAction.ProjectID != "" && string(revision.ProposedNextAction.ProjectID) != pid {
				return fail("project %s next-action scope drifted to %s", pid, revision.ProposedNextAction.ProjectID)
			}
		}
		if bundle.Project.CurrentStateRevision != nil {
			current := uint64(*bundle.Project.CurrentStateRevision)
			if _, exists := revisions[current]; !exists {
				return fail("project %s current revision %d is missing", pid, current)
			}
		}
		projectRevisions[pid] = revisions
	}

	authorityRevisions := make(map[uint64]struct{}, len(source.Authority.Revisions))
	for _, state := range source.Authority.Revisions {
		revision := uint64(state.Revision)
		if revision == 0 {
			return fail("authority contains revision zero")
		}
		if _, exists := authorityRevisions[revision]; exists {
			return fail("duplicate authority revision %d", revision)
		}
		authorityRevisions[revision] = struct{}{}
		for _, grant := range state.Grants {
			if grant.AuthorityID == "" || grant.ActorID == "" || grant.Provenance == "" {
				return fail("authority revision %d contains incomplete grant identity/provenance", revision)
			}
			for _, scopedProject := range grant.ProjectScope {
				if _, exists := projectIDs[string(scopedProject)]; !exists {
					return fail("authority %s references unknown project %s", grant.AuthorityID, scopedProject)
				}
			}
		}
	}
	currentAuthority := uint64(source.Authority.CurrentRevision)
	if currentAuthority == 0 {
		return fail("current authority revision must be positive")
	}
	if _, exists := authorityRevisions[currentAuthority]; !exists {
		return fail("current authority revision %d is missing", currentAuthority)
	}

	for _, attempt := range source.Attempts {
		pid := string(attempt.ProjectID)
		if _, exists := projectIDs[pid]; !exists {
			return fail("transition attempt %s references unknown project %s", attempt.AttemptID, pid)
		}
		if attempt.AcceptedStateRevision != nil {
			if _, exists := projectRevisions[pid][uint64(*attempt.AcceptedStateRevision)]; !exists {
				return fail("transition attempt %s references unknown accepted state revision", attempt.AttemptID)
			}
		}
	}
	for _, record := range source.Records {
		pid := string(record.ProjectID)
		if pid != "" {
			if _, exists := projectIDs[pid]; !exists {
				return fail("record %s references unknown project %s", record.RecordID, pid)
			}
			if record.StateRevision != nil {
				if _, exists := projectRevisions[pid][uint64(*record.StateRevision)]; !exists {
					return fail("record %s references unknown state revision", record.RecordID)
				}
			}
		}
		if record.AuthorityRevision != nil {
			if _, exists := authorityRevisions[uint64(*record.AuthorityRevision)]; !exists {
				return fail("record %s references unknown authority revision", record.RecordID)
			}
		}
	}
	return nil
}
