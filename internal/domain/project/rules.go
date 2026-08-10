package project

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func NewProjectID() (ProjectID, error) {
	var raw [16]byte
	if _, err := rand.Read(raw[:]); err != nil {
		return "", fmt.Errorf("generate Project ID: %w", err)
	}
	return ProjectID("PRJ-" + hex.EncodeToString(raw[:])), nil
}
