package cli

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/developmentconexus-ops/aurora_project/internal/application"
)

func renderResult(out io.Writer, asJSON bool, v any) error {
	if asJSON { return json.NewEncoder(out).Encode(v) }
	switch r := v.(type) {
	case application.InitializeResult:
		_, err := fmt.Fprintf(out, "Aurora initialized\nAurora ID: %s\nOwner: %s\nGeneration: %d\n", r.AuroraID, r.OwnerOperatorID, r.GoverningGeneration)
		return err
	case application.InspectResult:
		_, err := fmt.Fprintf(out, "Aurora ID: %s\nOwner: %s\nAuthority revision: %d\nGeneration: %d\nTrust: %s\n", r.AuroraID, r.OwnerOperatorID, r.CurrentAuthorityRevision, r.GoverningGeneration, r.TrustStatus)
		return err
	default:
		return fmt.Errorf("unsupported CLI result type %T", v)
	}
}
