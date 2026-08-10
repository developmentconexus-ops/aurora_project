package cli

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/developmentconexus-ops/aurora_project/internal/application"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
)

func renderResult(out io.Writer, asJSON bool, value any) error {
	if asJSON {
		return json.NewEncoder(out).Encode(value)
	}
	switch result := value.(type) {
	case application.InitializeResult:
		_, err := fmt.Fprintf(out, "Aurora initialized\nAurora ID: %s\nOwner: %s\nGeneration: %d\n", result.AuroraID, result.OwnerOperatorID, result.GoverningGeneration)
		return err
	case application.InspectResult:
		_, err := fmt.Fprintf(out, "Aurora ID: %s\nOwner: %s\nAuthority revision: %d\nGeneration: %d\nTrust: %s\n", result.AuroraID, result.OwnerOperatorID, result.CurrentAuthorityRevision, result.GoverningGeneration, result.TrustStatus)
		return err
	case project.Project:
		_, err := fmt.Fprintf(out, "Project created\nProject ID: %s\nLabel: %s\nObjective: %s\n", result.ProjectID, result.DisplayLabel, result.ObjectiveSummary)
		return err
	case application.ProjectView:
		_, err := fmt.Fprintf(out, "Project ID: %s\nLabel: %s\nObjective: %s\n", result.Project.ProjectID, result.Project.DisplayLabel, result.Project.ObjectiveSummary)
		if err != nil {
			return err
		}
		if result.CurrentState == nil {
			_, err = fmt.Fprintln(out, "Current state: NONE")
			return err
		}
		_, err = fmt.Fprintf(out, "Current state revision: %d\nState kind: %s\nState summary: %s\n", result.CurrentState.Revision, result.CurrentState.State.Kind, result.CurrentState.State.Summary)
		if err != nil {
			return err
		}
		if result.NextSafeAction != nil {
			_, err = fmt.Fprintf(out, "Next safe action: %s\n", result.NextSafeAction.Decision)
		}
		return err
	case project.ProjectStateRevision:
		_, err := fmt.Fprintf(out, "State accepted\nProject ID: %s\nState revision: %d\nState kind: %s\nState summary: %s\n", result.ProjectID, result.Revision, result.State.Kind, result.State.Summary)
		return err
	case application.AuthorityView:
		_, err := fmt.Fprintf(out, "Authority state revision: %d\nGrants: %d\nRevalidation required: %t\n", result.State.Revision, len(result.State.Grants), result.State.RevalidationRequired)
		return err
	case application.ExportResult:
		_, err := fmt.Fprintf(out, "Export created\nExport ID: %s\nSHA-256: %s\nCreated at: %s\n", result.ExportID, result.PayloadSHA256, result.CreatedAt.Format("2006-01-02T15:04:05Z07:00"))
		return err
	case application.RestoreResult:
		_, err := fmt.Fprintf(out, "Restore complete\nExport ID: %s\nAurora ID: %s\nAuthority revision: %d\nGeneration: %d\nTrust: %s\n", result.ExportID, result.AuroraID, result.AuthorityStateRevision, result.GoverningGeneration, result.TrustStatus)
		return err
	case application.MigratePackageResult:
		_, err := fmt.Fprintf(out, "Migration complete\nSource version: %d\nTarget version: %d\nSHA-256: %s\n", result.SourceVersion, result.TargetVersion, result.PayloadSHA256)
		return err
	default:
		return fmt.Errorf("unsupported CLI result type %T", value)
	}
}
