package cli

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/developmentconexus-ops/aurora_project/internal/application"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/authority"
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
func renderTransition(out io.Writer, result application.TransitionProjectResult, asJSON bool) error {
	if asJSON {
		return json.NewEncoder(out).Encode(result)
	}
	_, err := fmt.Fprintf(out, "Project ID: %s\nState revision: %d\nGeneration: %d\n", result.ProjectID, result.StateRevision, result.GoverningGeneration)
	return err
}
func renderJSON(out io.Writer, value any, asJSON bool) error {
	if asJSON {
		return json.NewEncoder(out).Encode(value)
	}
	_, err := fmt.Fprintf(out, "%+v\n", value)
	return err
}
func renderJSONOrAuthority(out io.Writer, value authority.Snapshot, asJSON bool) error {
	if asJSON {
		return json.NewEncoder(out).Encode(value)
	}
	_, err := fmt.Fprintf(out, "Authority revision: %d\nDecision: %s\nStatus: %s\n", value.SourceRevision, value.Decision, value.Status)
	return err
}
