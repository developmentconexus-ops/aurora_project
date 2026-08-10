package cli

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/developmentconexus-ops/aurora_project/internal/application"
)

func renderInitialize(out io.Writer, result application.InitializeResult, asJSON bool) error {
	if asJSON {
		return json.NewEncoder(out).Encode(result)
	}
	_, err := fmt.Fprintf(out, "Aurora initialized\nAurora ID: %s\nOwner: %s\n", result.AuroraID, result.OwnerOperatorID)
	return err
}
func renderInspect(out io.Writer, result application.InspectResult, asJSON bool) error {
	if asJSON {
		return json.NewEncoder(out).Encode(result)
	}
	_, err := fmt.Fprintf(out, "Aurora ID: %s\nOwner: %s\nGeneration: %d\nTrust: %s\n", result.AuroraID, result.OwnerOperatorID, result.GoverningGeneration, result.TrustState)
	return err
}
func renderProject(out io.Writer, result application.ProjectInspection, asJSON bool) error {
	if asJSON {
		return json.NewEncoder(out).Encode(result)
	}
	_, err := fmt.Fprintf(out, "Project ID: %s\nLabel: %s\nObjective: %s\n", result.ProjectID, result.DisplayLabel, result.ObjectiveSummary)
	if err != nil {
		return err
	}
	if result.CurrentState != nil {
		_, err = fmt.Fprintf(out, "State revision: %d\nState: %s\n", result.CurrentState.Revision, result.CurrentState.State.Summary)
	}
	return err
}
