package application

import (
	"context"
	"fmt"
	"log/slog"
	"runtime"
	"sync/atomic"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/evidence"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

var operationSequence atomic.Uint64

type operationScope struct {
	service *Service
	ctx context.Context
	span trace.Span
	name string
	correlation evidence.Correlation
}

func (s *Service) startOperation(ctx context.Context, name, operationID string, c evidence.Correlation) *operationScope {
	if operationID == "" {
		operationID = fmt.Sprintf("OP-%d-%06d", s.operationNow().UnixNano(), operationSequence.Add(1))
	}
	c.OperationID = operationID
	if c.ProofRunID == "" { c.ProofRunID = evidence.ProofRunIDFromContext(ctx) }
	op := &operationScope{service:s,ctx:ctx,name:name,correlation:c}
	if s != nil && s.Tracer != nil {
		ctx2, span := s.Tracer.Start(ctx, name, trace.WithAttributes(operationAttributes(c,name)...))
		op.ctx, op.span = ctx2, span
	}
	if s != nil && s.Logger != nil { s.Logger.LogAttrs(op.ctx,slog.LevelInfo,"Aurora operation started",slogAttrs(c,name,"","")...) }
	if s != nil && s.Meter != nil {
		if counter, err := s.Meter.Int64Counter("aurora.m0.operations.total", metric.WithDescription("material M0 operations observed by the local Core")); err == nil {
			counter.Add(op.ctx,1,metric.WithAttributes(attribute.String("operation_name",name)))
		}
	}
	return op
}

func (op *operationScope) updateCorrelation(c evidence.Correlation) {
	if op == nil { return }
	c.OperationID = op.correlation.OperationID
	if c.ProofRunID == "" { c.ProofRunID = op.correlation.ProofRunID }
	op.correlation = c
}

func (op *operationScope) finish(result, reason string) {
	if op == nil || op.service == nil { return }
	if op.service.Logger != nil { op.service.Logger.LogAttrs(op.ctx,slog.LevelInfo,"Aurora operation finished",slogAttrs(op.correlation,op.name,result,reason)...) }
	if op.span != nil {
		op.span.SetAttributes(operationAttributes(op.correlation,op.name)...)
		op.span.SetAttributes(attribute.String("result",result),attribute.String("reason",reason))
		if result == "FAILED" || result == "REJECTED" || result == "BLOCKED" { op.span.SetStatus(codes.Error,reason) }
		op.span.End()
	}
	if (result == "FAILED" || result == "REJECTED" || result == "BLOCKED") && op.service.Meter != nil {
		if counter, err := op.service.Meter.Int64Counter("aurora.m0.operations.failures", metric.WithDescription("failed material M0 operations observed by the local Core")); err == nil {
			counter.Add(op.ctx,1,metric.WithAttributes(attribute.String("operation_name",op.name),attribute.String("result",result)))
		}
	}
}

func (s *Service) operationEvidence(op *operationScope, criterion, method string, refs []string, integrity string) evidence.EvidenceRecord {
	return evidence.EvidenceRecord{EvidenceID:"EVID-"+op.correlation.OperationID,Criterion:criterion,Producer:"aurora-sovereign-core",Verifier:"aurora-core-runtime-check",Method:method,Environment:runtime.GOOS+"/"+runtime.GOARCH,Versions:map[string]string{"go":runtime.Version(),"capability":"CAP-SOVEREIGN-CORE@0.2.0"},ArtifactRefs:append([]string(nil),refs...),IntegrityRef:integrity,Limitations:[]string{"self-produced runtime evidence is not an independent R8 product verdict"},Correlation:op.correlation,CreatedAt:s.operationNow()}
}

func (s *Service) appendOperationalRecord(ctx context.Context, kind, outcome, reason string, record evidence.EvidenceRecord) error {
	if s == nil || s.State == nil { return nil }
	return s.State.AppendRecord(ctx,evidence.StoredRecord{RecordID:"REC-"+record.Correlation.OperationID+"-"+kind,Kind:kind,OperationID:record.Correlation.OperationID,ProjectID:record.Correlation.ProjectID,StateRevision:record.Correlation.StateRevision,AuthorityRevision:record.Correlation.AuthorityRevision,Outcome:outcome,Reason:reason,Details:record,CreatedAt:record.CreatedAt})
}

func (s *Service) operationNow() time.Time { if s != nil && s.Clock != nil { return s.Clock.Now().UTC() }; return time.Now().UTC() }

func operationAttributes(c evidence.Correlation,name string) []attribute.KeyValue {
	a:=[]attribute.KeyValue{attribute.String("operation_id",c.OperationID),attribute.String("operation_name",name)}
	if c.ProofRunID!=""{a=append(a,attribute.String("proof_run_id",c.ProofRunID))};if c.AuroraID!=""{a=append(a,attribute.String("aurora_id",string(c.AuroraID)))};if c.ProjectID!=""{a=append(a,attribute.String("project_id",string(c.ProjectID)))};if c.StateRevision!=nil{a=append(a,attribute.Int64("state_revision",int64(*c.StateRevision)))};if c.AuthorityRevision!=nil{a=append(a,attribute.Int64("authority_revision",int64(*c.AuthorityRevision)))};return a
}
func slogAttrs(c evidence.Correlation,name,result,reason string) []slog.Attr {
	a:=[]slog.Attr{slog.String("operation_id",c.OperationID),slog.String("operation_name",name)}
	if c.ProofRunID!=""{a=append(a,slog.String("proof_run_id",c.ProofRunID))};if c.AuroraID!=""{a=append(a,slog.String("aurora_id",string(c.AuroraID)))};if c.ProjectID!=""{a=append(a,slog.String("project_id",string(c.ProjectID)))};if c.StateRevision!=nil{a=append(a,slog.Uint64("state_revision",uint64(*c.StateRevision)))};if c.AuthorityRevision!=nil{a=append(a,slog.Uint64("authority_revision",*c.AuthorityRevision))};if result!=""{a=append(a,slog.String("result",result))};if reason!=""{a=append(a,slog.String("reason",reason))};return a
}
