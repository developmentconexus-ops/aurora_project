package project

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

const ProjectStateSchemaV1 = `{
  "$schema":"https://json-schema.org/draft/2020-12/schema",
  "$id":"https://aurora.local/schemas/project-state-v1.schema.json",
  "type":"object",
  "additionalProperties":false,
  "required":["state_schema_version","state_kind","state_summary"],
  "properties":{
    "state_schema_version":{"const":"1"},
    "state_kind":{"type":"string","minLength":1,"maxLength":100},
    "state_summary":{"type":"string","minLength":1,"maxLength":2000},
    "state_payload":{}
  }
}`

var (
	stateSchemaOnce sync.Once
	stateSchema     *jsonschema.Schema
	stateSchemaErr  error
)

func compiledStateSchema() (*jsonschema.Schema, error) {
	stateSchemaOnce.Do(func() {
		resource, err := jsonschema.UnmarshalJSON(strings.NewReader(ProjectStateSchemaV1))
		if err != nil { stateSchemaErr = err; return }
		compiler := jsonschema.NewCompiler()
		if err := compiler.AddResource("project-state-v1.schema.json", resource); err != nil { stateSchemaErr = err; return }
		stateSchema, stateSchemaErr = compiler.Compile("project-state-v1.schema.json")
	})
	return stateSchema, stateSchemaErr
}

func ValidateStateEnvelope(env StateEnvelope) error {
	if env.Payload != nil && !json.Valid(env.Payload) { return errors.New("state payload is not valid JSON") }
	raw, err := json.Marshal(env)
	if err != nil { return fmt.Errorf("marshal state envelope: %w", err) }
	instance, err := jsonschema.UnmarshalJSON(bytes.NewReader(raw))
	if err != nil { return fmt.Errorf("decode state envelope for schema validation: %w", err) }
	schema, err := compiledStateSchema()
	if err != nil { return fmt.Errorf("compile project state schema: %w", err) }
	if err := schema.Validate(instance); err != nil { return fmt.Errorf("state envelope schema: %w", err) }
	return nil
}
