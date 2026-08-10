package cli

import (
	"io"
	"log/slog"

	"github.com/developmentconexus-ops/aurora_project/internal/adapters/observability"
	"github.com/developmentconexus-ops/aurora_project/internal/application"
)

func attachObservability(svc *application.Service, writer io.Writer) *observability.Runtime {
	if writer == nil {
		writer = io.Discard
	}
	runtime := observability.New(slog.New(slog.NewJSONHandler(writer, nil)))
	svc.Logger = runtime.Logger
	svc.Tracer = runtime.Tracer
	svc.Meter = runtime.Meter
	return runtime
}
