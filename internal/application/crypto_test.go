package application

import (
	"bytes"
	"testing"
)

func TestOwnerRootEnvelopeRoundTripAndRewrap(t *testing.T) {
	oldPass := []byte("fixture-old-passphrase")
	newPass := []byte("fixture-new-passphrase")

	env, ork, err := newRootEnvelope(oldPass)
	if err != nil {
		t.Fatal(err)
	}
	if len(ork) != 32 {
		t.Fatalf("ORK length = %d, want 32", len(ork))
	}
	unlocked, err := unlockORK(env, oldPass)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(unlocked, ork) {
		t.Fatal("unlocked ORK differs from generated ORK")
	}
	if _, err := unlockORK(env, []byte("wrong-passphrase")); err == nil {
		t.Fatal("wrong passphrase unexpectedly unlocked ORK")
	}

	rewrapped, err := rewrapRootEnvelope(env, oldPass, newPass)
	if err != nil {
		t.Fatal(err)
	}
	if rewrapped.RootID != env.RootID {
		t.Fatalf("root ID changed on rewrap: %q -> %q", env.RootID, rewrapped.RootID)
	}
	unlockedNew, err := unlockORK(rewrapped, newPass)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(unlockedNew, ork) {
		t.Fatal("rewrap changed ORK lineage")
	}
}

func TestOwnerRootEnvelopeRejectsUnboundedKDFBeforeDerivation(t *testing.T) {
	env, _, err := newRootEnvelope([]byte("fixture-passphrase"))
	if err != nil {
		t.Fatal(err)
	}
	env.KDF.MemoryKiB = 1024 * 1024
	if _, err := unlockORK(env, []byte("fixture-passphrase")); err == nil {
		t.Fatal("extreme KDF parameters unexpectedly accepted")
	}
}
