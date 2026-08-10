package application

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
	"github.com/developmentconexus-ops/aurora_project/schemas"
	"github.com/santhosh-tekuri/jsonschema/v6"
)

var (
	projectStateSchemaOnce sync.Once
	projectStateSchema     *jsonschema.Schema
	projectStateSchemaErr  error
)

func compileProjectStateSchema() (*jsonschema.Schema, error) {
	projectStateSchemaOnce.Do(func() {
		var doc any
		if err := json.Unmarshal(schemas.ProjectStateV1, &doc); err != nil {
			projectStateSchemaErr = fmt.Errorf("decode embedded project state schema: %w", err)
			return
		}
		compiler := jsonschema.NewCompiler()
		compiler.DefaultDraft(jsonschema.Draft2020)
		const location = "urn:aurora:project-state:v1"
		if err := compiler.AddResource(location, doc); err != nil {
			projectStateSchemaErr = fmt.Errorf("add project state schema: %w", err)
			return
		}
		projectStateSchema, projectStateSchemaErr = compiler.Compile(location)
	})
	return projectStateSchema, projectStateSchemaErr
}

func validateStateEnvelope(env project.StateEnvelope) error {
	encoded, err := json.Marshal(env)
	if err != nil {
		return fmt.Errorf("encode state envelope: %w", err)
	}
	var value any
	if err := json.Unmarshal(encoded, &value); err != nil {
		return fmt.Errorf("decode state envelope: %w", err)
	}
	schema, err := compileProjectStateSchema()
	if err != nil {
		return err
	}
	if err := schema.Validate(value); err != nil {
		return fmt.Errorf("validate state envelope: %w", err)
	}
	return nil
}
