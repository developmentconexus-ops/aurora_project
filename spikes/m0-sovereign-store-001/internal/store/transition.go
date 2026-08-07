package store

import (
	"errors"
	"fmt"
)

var ErrStaleRevision = errors.New("stale expected revision")

// TransitionInput is the minimal accepted mutation used by SPK-001.
type TransitionInput struct {
	AttemptID         string
	ProjectID         string
	ExpectedRevision  int64
	NewRevision       int64
	AuthorityRevision string
	StateKind         string
	StateSummary      string
	AuditID           string
	EvidenceID        string
	EvidenceRef       string
}

// Counts exposes only row counts needed to prove atomicity and stale no-change behavior.
type Counts struct {
	StateRevisions      int
	AcceptedTransitions int
	AuditRecords        int
	EvidenceRefs        int
}

// FaultHook marks a deterministic boundary. Tests may use it to block a child process
// so the parent can terminate it with the OS process-kill primitive.
type FaultHook func(point string)

func callHook(h FaultHook, point string) {
	if h != nil {
		h(point)
	}
}

// ApplyTransition executes the exact one-transaction mutation boundary under test.
func ApplyTransition(path string, in TransitionInput, hook FaultHook) error {
	db, err := openExistingDB(path)
	if err != nil {
		return err
	}
	defer db.Close()

	callHook(hook, "before_tx")
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("begin transition: %w", err)
	}
	defer tx.Rollback()

	var current int64
	if err := tx.QueryRow(`SELECT current_revision FROM projects WHERE project_id = ?`, in.ProjectID).Scan(&current); err != nil {
		return fmt.Errorf("read current revision: %w", err)
	}
	if current != in.ExpectedRevision {
		return fmt.Errorf("%w: got %d want %d", ErrStaleRevision, in.ExpectedRevision, current)
	}
	if in.NewRevision != current+1 {
		return fmt.Errorf("new revision must be predecessor+1: current=%d new=%d", current, in.NewRevision)
	}
	callHook(hook, "after_validation")

	if _, err := tx.Exec(`
		INSERT INTO state_revisions(project_id, revision, state_kind, state_summary)
		VALUES (?, ?, ?, ?)
	`, in.ProjectID, in.NewRevision, in.StateKind, in.StateSummary); err != nil {
		return fmt.Errorf("insert new revision: %w", err)
	}
	callHook(hook, "after_revision_insert")

	result, err := tx.Exec(`
		UPDATE projects
		SET current_revision = ?, authority_revision = ?
		WHERE project_id = ? AND current_revision = ?
	`, in.NewRevision, in.AuthorityRevision, in.ProjectID, in.ExpectedRevision)
	if err != nil {
		return fmt.Errorf("update current pointer: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("current pointer rows affected: %w", err)
	}
	if rows != 1 {
		return fmt.Errorf("%w: pointer compare-and-set affected %d rows", ErrStaleRevision, rows)
	}
	callHook(hook, "after_pointer_update")

	if _, err := tx.Exec(`
		INSERT INTO transition_attempts(attempt_id, project_id, expected_revision, proposed_revision, result, reason)
		VALUES (?, ?, ?, ?, 'ACCEPTED', 'spike accepted transition')
	`, in.AttemptID, in.ProjectID, in.ExpectedRevision, in.NewRevision); err != nil {
		return fmt.Errorf("insert transition attempt: %w", err)
	}
	if _, err := tx.Exec(`
		INSERT INTO audit_records(audit_id, attempt_id, project_id, revision, event_type)
		VALUES (?, ?, ?, ?, 'STATE_TRANSITION_ACCEPTED')
	`, in.AuditID, in.AttemptID, in.ProjectID, in.NewRevision); err != nil {
		return fmt.Errorf("insert audit record: %w", err)
	}
	if _, err := tx.Exec(`
		INSERT INTO evidence_refs(evidence_id, attempt_id, ref)
		VALUES (?, ?, ?)
	`, in.EvidenceID, in.AttemptID, in.EvidenceRef); err != nil {
		return fmt.Errorf("insert evidence ref: %w", err)
	}
	callHook(hook, "before_commit")

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit transition: %w", err)
	}
	callHook(hook, "after_commit")
	return nil
}

func InspectCounts(path string) (Counts, error) {
	db, err := openExistingDB(path)
	if err != nil {
		return Counts{}, err
	}
	defer db.Close()

	var out Counts
	queries := []struct {
		q   string
		dst *int
	}{
		{`SELECT COUNT(*) FROM state_revisions`, &out.StateRevisions},
		{`SELECT COUNT(*) FROM transition_attempts WHERE result = 'ACCEPTED'`, &out.AcceptedTransitions},
		{`SELECT COUNT(*) FROM audit_records`, &out.AuditRecords},
		{`SELECT COUNT(*) FROM evidence_refs`, &out.EvidenceRefs},
	}
	for _, item := range queries {
		if err := db.QueryRow(item.q).Scan(item.dst); err != nil {
			return Counts{}, fmt.Errorf("inspect count: %w", err)
		}
	}
	return out, nil
}
