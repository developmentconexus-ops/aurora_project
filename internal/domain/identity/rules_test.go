package identity

import (
	"strings"
	"testing"
)

func TestNewAuroraIDIsRandomAndPrefixed(t *testing.T) {
	a, err := NewAuroraID()
	if err != nil {
		t.Fatal(err)
	}
	b, err := NewAuroraID()
	if err != nil {
		t.Fatal(err)
	}
	if a == b {
		t.Fatalf("two generated Aurora IDs are equal: %q", a)
	}
	if !strings.HasPrefix(string(a), "AUR-") || !strings.HasPrefix(string(b), "AUR-") {
		t.Fatalf("IDs %q and %q must use AUR- prefix", a, b)
	}
}
