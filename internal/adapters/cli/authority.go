package cli

import (
	"context"
	"io"
	"strconv"

	"github.com/developmentconexus-ops/aurora_project/internal/application"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/authority"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
)

func runAuthority(ctx context.Context, s *application.Service, pass []byte, args []string, out, errOut io.Writer, asJSON bool) int {
	if len(args) == 0 {
		return usageError(errOut, "authority requires show, grant or revoke")
	}
	switch args[0] {
	case "show":
		vals, err := parseNamed(args[1:], map[string]bool{"--project": true, "--action": true})
		if err != nil {
			return usageError(errOut, err.Error())
		}
		if vals["--project"] == "" || vals["--action"] == "" {
			return usageError(errOut, "--project and --action are required")
		}
		r, err := s.ShowAuthority(ctx, pass, project.ProjectID(vals["--project"]), vals["--action"])
		if err != nil {
			return runtimeError(errOut, err)
		}
		if err := renderJSONOrAuthority(out, r, asJSON); err != nil {
			return runtimeError(errOut, err)
		}
		return 0
	case "grant":
		vals, err := parseNamed(args[1:], map[string]bool{"--project": true, "--action": true, "--expected": true})
		if err != nil {
			return usageError(errOut, err.Error())
		}
		if vals["--project"] == "" || vals["--action"] == "" {
			return usageError(errOut, "--project and --action are required")
		}
		expected, err := parseAuthorityRevision(vals["--expected"])
		if err != nil {
			return usageError(errOut, err.Error())
		}
		r, err := s.GrantAuthority(ctx, pass, application.GrantAuthorityInput{ProjectID: project.ProjectID(vals["--project"]), Action: vals["--action"], ExpectedRevision: expected})
		if err != nil {
			return runtimeError(errOut, err)
		}
		if err := renderJSON(out, r, asJSON); err != nil {
			return runtimeError(errOut, err)
		}
		return 0
	case "revoke":
		vals, err := parseNamed(args[1:], map[string]bool{"--authority": true, "--expected": true})
		if err != nil {
			return usageError(errOut, err.Error())
		}
		if vals["--authority"] == "" {
			return usageError(errOut, "--authority is required")
		}
		expected, err := parseAuthorityRevision(vals["--expected"])
		if err != nil {
			return usageError(errOut, err.Error())
		}
		r, err := s.RevokeAuthority(ctx, pass, application.RevokeAuthorityInput{AuthorityID: vals["--authority"], ExpectedRevision: expected})
		if err != nil {
			return runtimeError(errOut, err)
		}
		if err := renderJSON(out, r, asJSON); err != nil {
			return runtimeError(errOut, err)
		}
		return 0
	default:
		return usageError(errOut, "unknown authority command")
	}
}

func parseAuthorityRevision(value string) (*authority.Revision, error) {
	if value == "" {
		return nil, nil
	}
	n, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return nil, err
	}
	r := authority.Revision(n)
	return &r, nil
}
