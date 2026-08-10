package application

import (
	"context"
	"crypto/hmac"
	"errors"
	"fmt"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/identity"
)

type InspectResult struct {
	AuroraID                 identity.AuroraID `json:"aurora_id"`
	OwnerOperatorID          identity.OperatorID `json:"owner_operator_id"`
	CurrentAuthorityRevision uint64 `json:"current_authority_revision"`
	GoverningGeneration      uint64 `json:"governing_generation"`
	TrustStatus              string `json:"trust_status"`
}

func (s *Service) Inspect(ctx context.Context, passphrase []byte) (InspectResult, error) {
	if s.State == nil || s.Trust == nil || s.Clock == nil { return InspectResult{}, errors.New("service dependencies are incomplete") }
	env, err := s.Trust.LoadRootEnvelope(ctx)
	if err != nil { return InspectResult{}, fmt.Errorf("load owner root: %w", err) }
	ork, err := unlockORK(passphrase, env)
	if err != nil { return InspectResult{}, err }
	defer zero(ork)
	snap, err := s.State.LoadCurrent(ctx)
	if err != nil { return InspectResult{}, fmt.Errorf("load current state: %w", err) }
	wantDBMAC, err := governingMAC(ork, snap.Identity, snap.GoverningGeneration, snap.AuthorityStateJSON)
	if err != nil { return InspectResult{}, err }
	if !hmac.Equal(wantDBMAC, snap.GoverningMAC) { return InspectResult{}, errors.New("INVALID_DB_MAC") }
	anchor, err := s.Trust.LoadAnchor(ctx)
	if err != nil { return InspectResult{}, fmt.Errorf("load owner anchor: %w", err) }
	provided, err := decodeMAC(anchor.HMAC)
	if err != nil { return InspectResult{}, errors.New("INVALID_ANCHOR_MAC") }
	wantAnchor, err := anchorMAC(ork, anchor)
	if err != nil { return InspectResult{}, err }
	if !hmac.Equal(provided, wantAnchor) { return InspectResult{}, errors.New("INVALID_ANCHOR_MAC") }
	if anchor.AuroraID != snap.Identity.AuroraID || anchor.RootID != env.RootID { return InspectResult{}, errors.New("INVALID_ANCHOR_BINDING") }
	if snap.GoverningGeneration != anchor.GoverningGeneration { return InspectResult{}, errors.New("TRUST_GENERATION_MISMATCH") }
	if s.Clock.Now().UTC().Before(anchor.ObservedWallTimeHighWater) { return InspectResult{}, errors.New("TIME_UNTRUSTED") }
	return InspectResult{AuroraID:snap.Identity.AuroraID, OwnerOperatorID:snap.Identity.OwnerOperatorID, CurrentAuthorityRevision:snap.CurrentAuthorityRevision, GoverningGeneration:snap.GoverningGeneration, TrustStatus:"NORMAL"}, nil
}
