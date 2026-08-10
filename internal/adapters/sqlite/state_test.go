package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

func TestLoadCurrentFailsWhenProjectPointerDoesNotResolve(t *testing.T) {
	store, err := Open(t.TempDir()); if err != nil { t.Fatal(err) }; defer store.Close()
	seedStoreForProjectTest(t, store)
	p := project.Project{ProjectID:"PRJ-broken", DisplayLabel:"Broken", ObjectiveSummary:"Pointer test", CreatedAt:time.Unix(2,0).UTC(), UpdatedAt:time.Unix(2,0).UTC()}
	if _, err := store.CreateProject(context.Background(), ports.CreateProjectMutation{Project:p, ExpectedGeneration:1, NewGeneration:2, GoverningMAC:[]byte("m2"), OperationID:"OP-BROKEN"}); err != nil { t.Fatal(err) }
	if _, err := store.db.Exec(`INSERT INTO project_state_revisions(project_id,state_revision,state_schema_version,state_kind,state_summary,state_payload_json,accepted_by_actor,accepted_at,transition_attempt_id) VALUES(?,1,'1','note','R1','{}','OWNER','2026-08-10T00:00:00Z','ATT-1')`, p.ProjectID); err != nil { t.Fatal(err) }
	if _, err := store.db.Exec(`UPDATE projects SET current_state_revision=99 WHERE project_id=?`, p.ProjectID); err != nil { t.Fatal(err) }
	_, err = store.LoadCurrent(context.Background())
	if !errors.Is(err, ErrBrokenCurrentStatePointer) { t.Fatalf("err=%v want ErrBrokenCurrentStatePointer", err) }
}

func TestGetCurrentProjectStateReturnsExactlyPointedRevision(t *testing.T) {
	store, err := Open(t.TempDir()); if err != nil { t.Fatal(err) }; defer store.Close()
	seedStoreForProjectTest(t, store)
	p := project.Project{ProjectID:"PRJ-current", DisplayLabel:"Current", ObjectiveSummary:"State test", CreatedAt:time.Unix(2,0).UTC(), UpdatedAt:time.Unix(2,0).UTC()}
	if _, err := store.CreateProject(context.Background(), ports.CreateProjectMutation{Project:p, ExpectedGeneration:1, NewGeneration:2, GoverningMAC:[]byte("m2"), OperationID:"OP-CURRENT"}); err != nil { t.Fatal(err) }
	for rev, summary := range map[int]string{1:"R1",2:"R2"} {
		if _, err := store.db.Exec(`INSERT INTO project_state_revisions(project_id,state_revision,predecessor_revision,state_schema_version,state_kind,state_summary,state_payload_json,accepted_by_actor,accepted_at,transition_attempt_id) VALUES(?,?,?,?,?,?,?,?,?,?)`, p.ProjectID,rev,nullablePredecessor(rev),"1","note",summary,"{}","OWNER","2026-08-10T00:00:00Z",summary); err != nil { t.Fatal(err) }
	}
	if _, err := store.db.Exec(`UPDATE projects SET current_state_revision=2 WHERE project_id=?`, p.ProjectID); err != nil { t.Fatal(err) }
	state, err := store.GetCurrentProjectState(context.Background(), p.ProjectID); if err != nil { t.Fatal(err) }
	if state == nil || state.Revision != 2 || state.State.Summary != "R2" { t.Fatalf("state=%+v", state) }
	var count int; if err := store.db.QueryRow(`SELECT COUNT(*) FROM project_state_revisions WHERE project_id=?`,p.ProjectID).Scan(&count); err != nil { t.Fatal(err) }
	if count != 2 { t.Fatalf("history count=%d want 2",count) }
}

func nullablePredecessor(rev int) any { if rev == 1 { return sql.NullInt64{} }; return rev-1 }
