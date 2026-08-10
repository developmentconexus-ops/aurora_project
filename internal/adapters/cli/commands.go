package cli

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/developmentconexus-ops/aurora_project/internal/adapters/sqlite"
	"github.com/developmentconexus-ops/aurora_project/internal/adapters/trustfs"
	"github.com/developmentconexus-ops/aurora_project/internal/application"
	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

type globalOptions struct {
	dataDir string
	json    bool
}

func parseGlobal(args []string) (globalOptions, []string, error) {
	var opts globalOptions
	var rest []string
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--json":
			opts.json = true
		case "--data-dir":
			if i+1 >= len(args) {
				return opts, nil, errors.New("--data-dir requires a path")
			}
			i++
			opts.dataDir = args[i]
		default:
			rest = append(rest, args[i])
		}
	}
	if opts.dataDir == "" {
		opts.dataDir = os.Getenv("AURORA_DATA_DIR")
	}
	if opts.dataDir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return opts, nil, err
		}
		opts.dataDir = filepath.Join(home, ".aurora")
	}
	return opts, rest, nil
}

func newService(dataDir string) (*application.Service, *sqlite.Store, error) {
	state, err := sqlite.Open(dataDir)
	if err != nil {
		return nil, nil, err
	}
	return &application.Service{State: state, Trust: trustfs.New(dataDir), Clock: ports.SystemClock{}}, state, nil
}
