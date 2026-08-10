package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

func (s *Store) CreateProject(ctx context.Context, in ports.CreateProjectMutation) (ports.ProjectRecord, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil { return ports.ProjectRecord{}, err }
	defer func(){ _ = tx.Rollback() }()
	var generation uint64
	if err := tx.QueryRowContext(ctx, `SELECT governing_generation FROM core_state WHERE singleton_key='core'`).Scan(&generation); err != nil {
		if errors.Is(err, sql.ErrNoRows) { return ports.ProjectRecord{}, ports.ErrNotInitialized }
		return ports.ProjectRecord{}, err
	}
	if generation != in.ExpectedGeneration { return ports.ProjectRecord{}, ports.ErrGenerationConflict }
	var exists int
	if err := tx.QueryRowContext(ctx, `SELECT COUNT(*) FROM projects WHERE project_id=?`, in.Project.ProjectID).Scan(&exists); err != nil { return ports.ProjectRecord{}, err }
	if exists != 0 { return ports.ProjectRecord{}, ports.ErrProjectExists }
	created := in.Project.CreatedAt.UTC().Format(time.RFC3339Nano)
	updated := in.Project.UpdatedAt.UTC().Format(time.RFC3339Nano)
	if _, err := tx.ExecContext(ctx, `INSERT INTO projects(project_id,display_label,objective_summary,current_state_revision,created_at,updated_at) VALUES(?,?,?,NULL,?,?)`, in.Project.ProjectID, in.Project.DisplayLabel, in.Project.ObjectiveSummary, created, updated); err != nil { return ports.ProjectRecord{}, err }
	if _, err := tx.ExecContext(ctx, `UPDATE core_state SET governing_generation=?, governing_descriptor_hmac=?, updated_at=? WHERE singleton_key='core' AND governing_generation=?`, in.NewGeneration, in.GoverningHMAC, updated, in.ExpectedGeneration); err != nil { return ports.ProjectRecord{}, err }
	if _, err := tx.ExecContext(ctx, `INSERT INTO records(record_id,kind,operation_id,project_id,state_revision,authority_revision,outcome,reason,details_json,created_at) VALUES(?, 'AUDIT', ?, ?, NULL, NULL, 'ACCEPTED', 'PROJECT_CREATED', '{}', ?)`, in.OperationID, in.OperationID, in.Project.ProjectID, updated); err != nil { return ports.ProjectRecord{}, err }
	if err := tx.Commit(); err != nil { return ports.ProjectRecord{}, err }
	return in.Project, nil
}

func (s *Store) LoadProject(ctx context.Context, id string) (ports.ProjectRecord, error) {
	return s.loadProjectBase(ctx, id)
}

func (s *Store) loadProjectBase(ctx context.Context, id string) (ports.ProjectRecord, error) {
	var r ports.ProjectRecord
	var rev sql.NullInt64
	var created, updated string
	if err := s.db.QueryRowContext(ctx, `SELECT project_id,display_label,objective_summary,current_state_revision,created_at,updated_at FROM projects WHERE project_id=?`, id).Scan(&r.ProjectID,&r.DisplayLabel,&r.ObjectiveSummary,&rev,&created,&updated); err != nil {
		if errors.Is(err, sql.ErrNoRows) { return ports.ProjectRecord{}, ports.ErrProjectNotFound }
		return ports.ProjectRecord{}, err
	}
	if rev.Valid { v := uint64(rev.Int64); r.CurrentStateRevision = &v }
	var err error
	r.CreatedAt, err = time.Parse(time.RFC3339Nano, created); if err != nil { return ports.ProjectRecord{}, err }
	r.UpdatedAt, err = time.Parse(time.RFC3339Nano, updated); if err != nil { return ports.ProjectRecord{}, err }
	return r, nil
}

func (s *Store) loadProjects(ctx context.Context) ([]ports.ProjectRecord, error) {
	rows, err := s.db.QueryContext(ctx, `SELECT project_id FROM projects ORDER BY project_id`)
	if err != nil { return nil, err }
	defer rows.Close()
	var ids []string
	for rows.Next() { var id string; if err := rows.Scan(&id); err != nil { return nil, err }; ids=append(ids,id) }
	if err := rows.Err(); err != nil { return nil, err }
	out:=make([]ports.ProjectRecord,0,len(ids))
	for _,id:=range ids { p,state,err:=s.LoadProjectCurrent(ctx,id); if err!=nil{return nil,err}; p.CurrentState=state; out=append(out,p) }
	return out,nil
}
