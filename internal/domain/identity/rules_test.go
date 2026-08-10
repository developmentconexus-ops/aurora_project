package identity

import (
	"strings"
	"testing"
)

func TestNewAuroraIDIsPrefixedAndUnique(t *testing.T) {
	a, err := NewAuroraID()
	if err != nil { t.Fatal(err) }
	b, err := NewAuroraID()
	if err != nil { t.Fatal(err) }
	if a == b { t.Fatalf("generated duplicate Aurora IDs: %q", a) }
	if !strings.HasPrefix(string(a), "AUR-") { t.Fatalf("Aurora ID %q missing AUR- prefix", a) }
}
