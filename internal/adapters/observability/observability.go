package observability

import (
	"context"
	"errors"
	"io"
	"log/slog"

	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type SignalDefinition struct {
	Name       string `json:"name"`
	Kind       string `json:"kind"`
	Purpose    string `json:"purpose"`
	Limitation string `json:"limitation"`
}

var signalDefinitions = []SignalDefinition{
	{
		Name:       "aurora.m0.operations.total",
		Kind:       "counter",
		Purpose:    "count material M0 operations for proof-run completeness and operational review",
		Limitation: "does not prove an operation was semantically correct or accepted by R8",
	},
	{
		Name:       "aurora.m0.operations.failures",
		Kind:       "counter",
		Purpose:    "count material M0 operation failures for recovery and reliability review",
		Limitation: "does not establish canonical state, authority, root cause, or product verdict",
	},
}

func SignalDefinitions() []SignalDefinition {
	out := make([]SignalDefinition, len(signalDefinitions))
	copy(out, signalDefinitions)
	return out
}

type Runtime struct {
	Logger *slog.Logger
	Tracer trace.Tracer
	Meter  metric.Meter

	traceProvider *sdktrace.TracerProvider
	meterProvider *sdkmetric.MeterProvider
}

func New(logger *slog.Logger) *Runtime {
	if logger == nil {
		logger = slog.New(slog.NewTextHandler(io.Discard, nil))
	}
	traceProvider := sdktrace.NewTracerProvider()
	meterProvider := sdkmetric.NewMeterProvider()
	return &Runtime{
		Logger:        logger,
		Tracer:        traceProvider.Tracer("aurora.sovereign-core"),
		Meter:         meterProvider.Meter("aurora.sovereign-core"),
		traceProvider: traceProvider,
		meterProvider: meterProvider,
	}
}

func (r *Runtime) Shutdown(ctx context.Context) error {
	if r == nil {
		return nil
	}
	var errs []error
	if r.traceProvider != nil {
		errs = append(errs, r.traceProvider.Shutdown(ctx))
	}
	if r.meterProvider != nil {
		errs = append(errs, r.meterProvider.Shutdown(ctx))
	}
	return errors.Join(errs...)
}
