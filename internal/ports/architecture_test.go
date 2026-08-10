package ports_test

import (
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestDomainDoesNotImportAdaptersOrSelectedMechanisms(t *testing.T) {
	root := filepath.Join("..", "domain")
	if _, err := os.Stat(root); err != nil {
		if os.IsNotExist(err) {
			return
		}
		t.Fatal(err)
	}

	forbidden := []string{
		"/adapters/",
		"modernc.org/sqlite",
		"go.opentelemetry.io",
		"filippo.io/age",
		"golang.org/x/sys/windows",
	}

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".go") || strings.HasSuffix(d.Name(), "_test.go") {
			return nil
		}
		file, err := parser.ParseFile(token.NewFileSet(), path, nil, parser.ImportsOnly)
		if err != nil {
			return err
		}
		for _, imp := range file.Imports {
			value, err := strconv.Unquote(imp.Path.Value)
			if err != nil {
				return err
			}
			for _, bad := range forbidden {
				if strings.Contains(value, bad) || value == bad {
					t.Errorf("domain file %s imports forbidden dependency %q", path, value)
				}
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
