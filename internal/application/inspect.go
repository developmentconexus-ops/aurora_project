package application

import (
	"context"
	"crypto/hmac"
	"encoding/base64"
	"encoding/json"
	"errors"
)

var ErrTrustValidation = errors.New("owner trust validation failed")

type InspectResult struct {
	AuroraID            string `json:"aurora_id"`
	OwnerOperatorID     string `json:"owner_operator_id"`
	GoverningGeneration uint64 `json:"governing_generation"`
	TrustState          string `json:"trust_state"`
}

func (s *Service) Inspect(ctx context.Context, passphrase []byte) (InspectResult, error) {
	current, err := s.State.LoadCurrent(ctx)
	if err != nil {
		return InspectResult{}, err
	}
	root, err := s.Trust.LoadRootEnvelope(ctx)
	if err != nil {
		return InspectResult{}, err
	}
	ork, err := unlockORK(root, passphrase)
	if err != nil {
		return InspectResult{}, err
	}
	defer clear(ork)
	snapshot := bootstrapGoverningSnapshot{
		Version:             1,
		Aurora:              bootstrapAurora{AuroraID: current.AuroraID, OwnerOperatorID: current.OwnerOperatorID},
		GoverningGeneration: current.GoverningGeneration,
		Projects:            []any{},
		Authority:           json.RawMessage(current.AuthorityJSON),
	}
	wantDBMAC, err := governingMAC(ork, snapshot)
	if err != nil {
		return InspectResult{}, err
	}
	if !hmac.Equal(wantDBMAC, current.GoverningHMAC) {
		return InspectResult{}, ErrTrustValidation
	}
	anchor, err := s.Trust.LoadAnchor(ctx)
	if err != nil {
		return InspectResult{}, err
	}
	anchorBytes, err := base64.RawURLEncoding.DecodeString(anchor.HMAC)
	if err != nil {
		return InspectResult{}, ErrTrustValidation
	}
	wantAnchorMAC, err := anchorMAC(ork, anchor)
	if err != nil {
		return InspectResult{}, err
	}
	if !hmac.Equal(wantAnchorMAC, anchorBytes) || anchor.RootID != root.RootID || anchor.AuroraID != current.AuroraID || anchor.GoverningGeneration != current.GoverningGeneration {
		return InspectResult{}, ErrTrustValidation
	}
	return InspectResult{AuroraID: current.AuroraID, OwnerOperatorID: current.OwnerOperatorID, GoverningGeneration: current.GoverningGeneration, TrustState: "NORMAL"}, nil
}
