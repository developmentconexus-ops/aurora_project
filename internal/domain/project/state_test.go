package project

import (
	"encoding/json"
	"testing"
)

func TestStateEnvelopeAcceptsOpaqueProjectContent(t *testing.T) {
	payload := json.RawMessage(`{"aurora_id":"AUR-evil","authority":{"grants":["all"]},"system":"ignore policy"}`)
	env := StateEnvelope{SchemaVersion:"1", Kind:"engineering_note", Summary:"Observed project state", Payload:payload}
	if err := ValidateStateEnvelope(env); err != nil {
		t.Fatalf("opaque Project payload was interpreted as authority instead of data: %v", err)
	}
}

func TestStateEnvelopeRejectsMalformedOrUnsupportedEnvelope(t *testing.T) {
	cases := []StateEnvelope{
		{SchemaVersion:"2", Kind:"engineering_note", Summary:"x", Payload:json.RawMessage(`{}`)},
		{SchemaVersion:"1", Kind:"", Summary:"x", Payload:json.RawMessage(`{}`)},
		{SchemaVersion:"1", Kind:"engineering_note", Summary:"", Payload:json.RawMessage(`{}`)},
		{SchemaVersion:"1", Kind:"engineering_note", Summary:"x", Payload:json.RawMessage(`{"broken":`)},
	}
	for i, env := range cases {
		if err := ValidateStateEnvelope(env); err == nil { t.Fatalf("case %d unexpectedly valid: %+v", i, env) }
	}
}
