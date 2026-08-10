package sqlite

import (
	"context"
	"encoding/json"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/evidence"
)

func (s *Store) AppendRecord(ctx context.Context, record evidence.StoredRecord) error {
	details := []byte(`{}`)
	if record.Details != nil {
		encoded, err := json.Marshal(record.Details)
		if err != nil {
			return err
		}
		details = encoded
	}
	var projectID, stateRevision, authorityRevision any
	if record.ProjectID != "" {
		projectID = record.ProjectID
	}
	if record.StateRevision != nil {
		stateRevision = uint64(*record.StateRevision)
	}
	if record.AuthorityRevision != nil {
		authorityRevision = *record.AuthorityRevision
	}
	_, err := s.db.ExecContext(ctx,
		`INSERT INTO records(record_id,kind,operation_id,project_id,state_revision,authority_revision,outcome,reason,details_json,created_at) VALUES(?,?,?,?,?,?,?,?,?,?)`,
		record.RecordID, record.Kind, record.OperationID, projectID, stateRevision, authorityRevision, record.Outcome, record.Reason, string(details), record.CreatedAt.UTC().Format(timeLayout),
	)
	return err
}

const timeLayout = "2006-01-02T15:04:05.999999999Z07:00"
