package sqlite

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

func TestLoadProjectCurrentUsesExactPointerAndNeverPromotesHistory(t *testing.T) {
	ctx := context.Background()
	store, err := Open(t.TempDir())
	if err != nil { t.Fatal(err) }
	defer store.Close()
	now := time.Date(2026,8,9,22,0,0,0,time.UTC)
	_, err = store.Bootstrap(ctx, ports.BootstrapMutation{OperationID:"OP-B", AuroraID:"AUR-F", OwnerOperatorID:"OWNER-LOCAL", CreatedAt:now, AuthorityRevision:1, AuthorityJSON:[]byte(`{"revision":1,"grants":[]}`), GoverningGeneration:1, GoverningDescriptorHMAC:[]byte("mac1")})
	if err != nil { t.Fatal(err) }
	p := ports.ProjectRecord{ProjectID:"PRJ-F", DisplayLabel:"Fonte", ObjectiveSummary:"Fonte", CreatedAt:now, UpdatedAt:now}
	_, err = store.CreateProject(ctx, ports.CreateProjectMutation{OperationID:"OP-P", Project:p, ExpectedGeneration:1, NewGeneration:2, GoverningHMAC:[]byte("mac2")})
	if err != nil { t.Fatal(err) }
	_, err = store.db.ExecContext(ctx, `INSERT INTO project_state_revisions(project_id,state_revision,predecessor_revision,state_schema_version,state_kind,state_summary,state_payload_json,accepted_intent_ref,proposed_next_action_json,accepted_by_actor,accepted_at,transition_attempt_id) VALUES(?,1,NULL,'1','WORK_NOTE','R1','{}',NULL,NULL,'OWNER-LOCAL',?,'ATT-1')`, p.ProjectID, now.Format(time.RFC3339Nano))
	if err != nil { t.Fatal(err) }
	_, err = store.db.ExecContext(ctx, `UPDATE projects SET current_state_revision=1 WHERE project_id=?`, p.ProjectID)
	if err != nil { t.Fatal(err) }
	projectRow, state, err := store.LoadProjectCurrent(ctx, p.ProjectID)
	if err != nil { t.Fatal(err) }
	if projectRow.ProjectID != p.ProjectID || state == nil || state.Revision != 1 || state.State.Kind != "WORK_NOTE" {
		t.Fatalf("current = project=%+v state=%+v", projectRow, state)
	}
	_, err = store.db.ExecContext(ctx, `UPDATE projects SET current_state_revision=99 WHERE project_id=?`, p.ProjectID)
	if err != nil { t.Fatal(err) }
	_, state, err = store.LoadProjectCurrent(ctx, p.ProjectID)
	if !errors.Is(err, ports.ErrCurrentStateMissing) {
		t.Fatalf("broken pointer err=%v state=%+v, want ErrCurrentStateMissing", err, state)
	}
}
