package sqlite

import (
	"context"
	"encoding/json"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/authority"
	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

func(s *Store)CommitAuthorityRevision(ctx context.Context,in ports.AuthorityMutation)(authority.State,error){tx,err:=s.db.BeginTx(ctx,nil);if err!=nil{return authority.State{},err};defer tx.Rollback();var current uint64;var generation uint64;if err:=tx.QueryRowContext(ctx,`SELECT current_authority_revision,governing_generation FROM core_state WHERE singleton_key='core'`).Scan(&current,&generation);err!=nil{return authority.State{},err};if authority.Revision(current)!=in.ExpectedAuthorityRevision{return authority.State{},ports.ErrStaleAuthorityRevision};if generation!=in.ExpectedGeneration{return authority.State{},ErrConcurrentMutation};raw,err:=json.Marshal(in.State);if err!=nil{return authority.State{},err};changed:=in.State.ChangedAt.UTC().Format(time.RFC3339Nano);var pred any;if in.State.PredecessorRevision!=nil{pred=uint64(*in.State.PredecessorRevision)};if _,err:=tx.ExecContext(ctx,`INSERT INTO authority_revisions(authority_revision,predecessor_revision,authority_state_json,changed_by,changed_at) VALUES(?,?,?,?,?)`,in.State.Revision,pred,string(raw),in.State.ChangedBy,changed);err!=nil{return authority.State{},err};res,err:=tx.ExecContext(ctx,`UPDATE core_state SET current_authority_revision=?,governing_generation=?,governing_descriptor_hmac=?,updated_at=? WHERE singleton_key='core' AND current_authority_revision=? AND governing_generation=?`,in.State.Revision,in.NewGeneration,in.GoverningMAC,changed,in.ExpectedAuthorityRevision,in.ExpectedGeneration);if err!=nil{return authority.State{},err};n,err:=res.RowsAffected();if err!=nil{return authority.State{},err};if n!=1{return authority.State{},ports.ErrStaleAuthorityRevision};if _,err:=tx.ExecContext(ctx,`INSERT INTO records(record_id,kind,operation_id,authority_revision,outcome,reason,details_json,created_at) VALUES(?,?,?,?,?,?,?,?)`,`REC-`+in.OperationID,"AUDIT",in.OperationID,in.State.Revision,"ACCEPTED","AUTHORITY_CHANGED",`{}`,changed);err!=nil{return authority.State{},err};if err:=tx.Commit();err!=nil{return authority.State{},err};return in.State,nil}
