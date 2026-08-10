package ports

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"strings"
	"testing"
)

func TestDomainDoesNotImportAdaptersOrMechanisms(t *testing.T) {
	root := filepath.Clean("../domain")
	forbidden := []string{"/adapters/", "modernc.org/sqlite", "go.opentelemetry.io", "filippo.io/age", "golang.org/x/sys/windows"}
	_ = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || filepath.Ext(path) != ".go" {
			return err
		}
		file, parseErr := parser.ParseFile(token.NewFileSet(), path, nil, parser.ImportsOnly)
		if parseErr != nil {
			t.Fatalf("parse %s: %v", path, parseErr)
		}
		for _, imp := range file.Imports {
			value := strings.Trim(imp.Path.Value, `"`)
			for _, bad := range forbidden {
				if strings.Contains(value, bad) {
					t.Errorf("domain file %s imports forbidden dependency %q", path, value)
				}
			}
		}
		_ = ast.File{}
		return nil
	})
}
