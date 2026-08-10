package identity

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func NewAuroraID() (AuroraID, error) {
	var raw [16]byte
	if _, err := rand.Read(raw[:]); err != nil {
		return "", fmt.Errorf("generate Aurora identity: %w", err)
	}
	return AuroraID("AUR-" + hex.EncodeToString(raw[:])), nil
}
