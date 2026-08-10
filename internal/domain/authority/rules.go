package authority

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func NewAuthorityID() (AuthorityID, error) {
	var raw [16]byte
	if _, err := rand.Read(raw[:]); err != nil { return "", fmt.Errorf("generate Authority ID: %w", err) }
	return AuthorityID("AUTH-"+hex.EncodeToString(raw[:])), nil
}
