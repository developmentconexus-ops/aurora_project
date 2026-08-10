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
