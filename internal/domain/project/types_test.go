package project

import (
	"strings"
	"testing"
)

func TestNewProjectIDIsRandomAndNotLabelDerived(t *testing.T) {
	a, err := NewProjectID()
	if err != nil {
		t.Fatal(err)
	}
	b, err := NewProjectID()
	if err != nil {
		t.Fatal(err)
	}
	if a == b {
		t.Fatalf("generated IDs are equal: %q", a)
	}
	if !strings.HasPrefix(string(a), "PRJ-") || !strings.HasPrefix(string(b), "PRJ-") {
		t.Fatalf("IDs must use PRJ- prefix: %q %q", a, b)
	}
	if strings.Contains(strings.ToLower(string(a)), "project fonte") {
		t.Fatalf("ID derived from display label: %q", a)
	}
}
