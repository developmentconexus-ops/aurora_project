package cli

import (
	"context"
	"io"
	"testing"

	"github.com/developmentconexus-ops/aurora_project/internal/application"
)

func TestAttachObservabilityWiresRealCLIService(t *testing.T) {
	svc := &application.Service{}
	runtime := attachObservability(svc, io.Discard)
	if runtime == nil {
		t.Fatal("observability runtime is nil")
	}
	if svc.Logger == nil || svc.Tracer == nil || svc.Meter == nil {
		t.Fatalf("service observability not wired: %+v", svc)
	}
	if err := runtime.Shutdown(context.Background()); err != nil {
		t.Fatalf("shutdown observability: %v", err)
	}
}
