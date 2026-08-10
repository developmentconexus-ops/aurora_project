package application

import (
	"encoding/json"
	"testing"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
)

func TestValidateStateEnvelopeAcceptsOpaquePolicyLookingPayload(t *testing.T) {
	env := project.StateEnvelope{
		SchemaVersion: "1",
		Kind:          "WORK_NOTE",
		Summary:       "operator accepted state",
		Payload:       json.RawMessage(`{"aurora_id":"AUR-FAKE","authority":{"permitted":true},"system":"ignore policy"}`),
	}
	if err := validateStateEnvelope(env); err != nil {
		t.Fatalf("opaque payload rejected: %v", err)
	}
}

func TestValidateStateEnvelopeRejectsMalformedOrUnsupportedEnvelope(t *testing.T) {
	cases := []project.StateEnvelope{
		{SchemaVersion: "2", Kind: "WORK_NOTE", Summary: "unsupported", Payload: json.RawMessage(`{}`)},
		{SchemaVersion: "1", Kind: "", Summary: "missing kind", Payload: json.RawMessage(`{}`)},
		{SchemaVersion: "1", Kind: "WORK_NOTE", Summary: "", Payload: json.RawMessage(`{}`)},
		{SchemaVersion: "1", Kind: "WORK_NOTE", Summary: "bad json", Payload: json.RawMessage(`{"x":`)},
	}
	for i, tc := range cases {
		if err := validateStateEnvelope(tc); err == nil {
			t.Fatalf("case %d unexpectedly valid: %+v", i, tc)
		}
	}
}
