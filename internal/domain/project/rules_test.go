package project

import (
	"strings"
	"testing"
)

func TestNewProjectIDIsStableClassAndUnique(t *testing.T) {
	a, err := NewProjectID()
	if err != nil { t.Fatal(err) }
	b, err := NewProjectID()
	if err != nil { t.Fatal(err) }
	if a == b { t.Fatalf("duplicate ProjectID %q", a) }
	if !strings.HasPrefix(string(a), "PRJ-") { t.Fatalf("ProjectID %q missing PRJ- prefix", a) }
}
