package cli

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/developmentconexus-ops/aurora_project/internal/application"
	"github.com/developmentconexus-ops/aurora_project/internal/domain/project"
)

func runProject(ctx context.Context, s *application.Service, pass []byte, args []string, out, errOut io.Writer, asJSON bool) int {
	if len(args) == 0 {
		return usageError(errOut, "project requires create, show or set-state")
	}
	switch args[0] {
	case "create":
		vals, err := parseNamed(args[1:], map[string]bool{"--label": true, "--objective": true})
		if err != nil {
			return usageError(errOut, err.Error())
		}
		if vals["--label"] == "" || vals["--objective"] == "" {
			return usageError(errOut, "--label and --objective are required")
		}
		p, err := s.CreateProject(ctx, pass, application.CreateProjectInput{DisplayLabel: vals["--label"], ObjectiveSummary: vals["--objective"]})
		if err != nil {
			return runtimeError(errOut, err)
		}
		if err := renderProject(out, application.ProjectInspection{Project: p}, asJSON); err != nil {
			return runtimeError(errOut, err)
		}
		return 0
	case "show":
		vals, err := parseNamed(args[1:], map[string]bool{"--project": true})
		if err != nil {
			return usageError(errOut, err.Error())
		}
		id := vals["--project"]
		if id == "" {
			return usageError(errOut, "--project is required")
		}
		inspection, err := s.InspectProject(ctx, pass, project.ProjectID(id))
		if err != nil {
			return runtimeError(errOut, err)
		}
		if err := renderProject(out, inspection, asJSON); err != nil {
			return runtimeError(errOut, err)
		}
		return 0
	case "set-state":
		vals, err := parseNamed(args[1:], map[string]bool{"--project": true, "--expected": true, "--kind": true, "--summary": true, "--payload": true, "--next-action": true, "--next-project": true})
		if err != nil {
			return usageError(errOut, err.Error())
		}
		if vals["--project"] == "" || vals["--expected"] == "" || vals["--kind"] == "" || vals["--summary"] == "" {
			return usageError(errOut, "--project, --expected, --kind and --summary are required")
		}
		var expected *project.StateRevision
		if vals["--expected"] != "none" {
			n, err := strconv.ParseUint(vals["--expected"], 10, 64)
			if err != nil {
				return usageError(errOut, "--expected must be none or an integer revision")
			}
			v := project.StateRevision(n)
			expected = &v
		}
		payload := json.RawMessage(vals["--payload"])
		if len(payload) == 0 {
			payload = json.RawMessage(`{}`)
		}
		var action *project.ActionDescriptor
		if vals["--next-action"] != "" {
			target := project.ProjectID(vals["--project"])
			if vals["--next-project"] != "" {
				target = project.ProjectID(vals["--next-project"])
			}
			action = &project.ActionDescriptor{ActionClass: vals["--next-action"], Summary: vals["--next-action"], ProjectID: target, RequiredAuthorityAction: vals["--next-action"]}
		}
		r, err := s.TransitionProjectState(ctx, pass, application.TransitionProjectInput{ProjectID: project.ProjectID(vals["--project"]), ExpectedRevision: expected, State: project.StateEnvelope{SchemaVersion: "1", Kind: vals["--kind"], Summary: vals["--summary"], Payload: payload}, ProposedNextAction: action})
		if err != nil {
			return runtimeError(errOut, err)
		}
		if err := renderTransition(out, r, asJSON); err != nil {
			return runtimeError(errOut, err)
		}
		return 0
	default:
		return usageError(errOut, "unknown project command")
	}
}

func parseNamed(args []string, allowed map[string]bool) (map[string]string, error) {
	vals := map[string]string{}
	for i := 0; i < len(args); i++ {
		key := args[i]
		if !allowed[key] {
			return nil, errors.New("unknown option " + key)
		}
		if i+1 >= len(args) {
			return nil, errors.New(key + " requires a value")
		}
		i++
		vals[key] = args[i]
	}
	return vals, nil
}
