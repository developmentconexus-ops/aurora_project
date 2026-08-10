package application

import (
	"log/slog"

	"github.com/developmentconexus-ops/aurora_project/internal/ports"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

type Service struct {
	State            ports.StateStore
	Trust            ports.OwnerTrustStore
	Clock            ports.Clock
	ExportProtection ports.ExportProtection
	Logger           *slog.Logger
	Tracer           trace.Tracer
	Meter            metric.Meter
}
