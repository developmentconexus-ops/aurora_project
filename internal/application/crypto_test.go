package application

import (
	"bytes"
	"errors"
	"testing"

	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

func TestRootEnvelopeRoundTripAndRotationPreserveORK(t *testing.T) {
	env, ork, err := newRootEnvelope([]byte("correct horse battery staple"))
	if err != nil { t.Fatal(err) }
	if len(ork) != 32 { t.Fatalf("ORK length=%d want 32", len(ork)) }
	got, err := unlockORK([]byte("correct horse battery staple"), env)
	if err != nil { t.Fatal(err) }
	if !bytes.Equal(got, ork) { t.Fatal("unwrapped ORK changed") }
	if _, err := unlockORK([]byte("wrong"), env); err == nil { t.Fatal("wrong passphrase unexpectedly unlocked ORK") }
	rotated, err := rewrapRootEnvelope([]byte("correct horse battery staple"), []byte("new passphrase"), env)
	if err != nil { t.Fatal(err) }
	got2, err := unlockORK([]byte("new passphrase"), rotated)
	if err != nil { t.Fatal(err) }
	if !bytes.Equal(got2, ork) { t.Fatal("rotation changed ORK lineage") }
}

func TestUnlockRejectsExtremeKDFBeforeAllocation(t *testing.T) {
	env := ports.RootEnvelope{Version: 1, KDF: "argon2id", MemoryKiB: 1<<31, Iterations: 3, Parallelism: 4}
	_, err := unlockORK([]byte("x"), env)
	if !errors.Is(err, ErrUnsupportedKDFProfile) { t.Fatalf("err=%v want ErrUnsupportedKDFProfile", err) }
}
