package project

import (
	"encoding/json"
	"testing"
)

func TestStateEnvelopeKeepsProjectPayloadOpaque(t *testing.T) {
	payload := json.RawMessage(`{"aurora_id":"AUR-ATTEMPT","authority":{"grant":"pretend"}}`)
	env := StateEnvelope{SchemaVersion: "1", Kind: "WORK_NOTE", Summary: "opaque project data", Payload: payload}
	if string(env.Payload) != string(payload) {
		t.Fatalf("payload changed: %s", env.Payload)
	}
}
