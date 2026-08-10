package observability

import (
	"context"
	"io"
	"log/slog"
	"testing"
)

func TestSignalDefinitionsStatePurposeAndLimitation(t *testing.T) {
	definitions := SignalDefinitions()
	if len(definitions) != 2 {
		t.Fatalf("signal definitions=%d want 2", len(definitions))
	}
	for _, definition := range definitions {
		if definition.Name == "" || definition.Kind == "" || definition.Purpose == "" || definition.Limitation == "" {
			t.Fatalf("incomplete signal definition: %+v", definition)
		}
	}
}

func TestRuntimeRequiresNoExporter(t *testing.T) {
	runtime := New(slog.New(slog.NewTextHandler(io.Discard, nil)))
	if runtime.Logger == nil || runtime.Tracer == nil || runtime.Meter == nil {
		t.Fatalf("runtime incomplete: %+v", runtime)
	}
	if err := runtime.Shutdown(context.Background()); err != nil {
		t.Fatalf("shutdown: %v", err)
	}
}
