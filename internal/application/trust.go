package application

import (
	"context"
	"crypto/hmac"
	"errors"
	"fmt"

	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

type trustedCurrent struct {
	Root ports.RootEnvelope
	ORK []byte
	Snapshot ports.CurrentSnapshot
	Anchor ports.Anchor
}

func (s *Service) loadTrustedCurrent(ctx context.Context, passphrase []byte) (trustedCurrent, error) {
	var out trustedCurrent
	root, err := s.Trust.LoadRootEnvelope(ctx); if err != nil { return out, fmt.Errorf("load owner root: %w", err) }
	ork, err := unlockORK(passphrase, root); if err != nil { return out, err }
	fail := func(err error) (trustedCurrent,error) { zero(ork); return trustedCurrent{}, err }
	snap, err := s.State.LoadCurrent(ctx); if err != nil { return fail(fmt.Errorf("load current state: %w", err)) }
	wantDB, err := governingMAC(ork, snap); if err != nil { return fail(err) }
	if !hmac.Equal(wantDB, snap.GoverningMAC) { return fail(errors.New("INVALID_DB_MAC")) }
	anchor, err := s.Trust.LoadAnchor(ctx); if err != nil { return fail(fmt.Errorf("load owner anchor: %w", err)) }
	provided, err := decodeMAC(anchor.HMAC); if err != nil { return fail(errors.New("INVALID_ANCHOR_MAC")) }
	wantAnchor, err := anchorMAC(ork, anchor); if err != nil { return fail(err) }
	if !hmac.Equal(provided,wantAnchor) { return fail(errors.New("INVALID_ANCHOR_MAC")) }
	if anchor.AuroraID != snap.Identity.AuroraID || anchor.RootID != root.RootID { return fail(errors.New("INVALID_ANCHOR_BINDING")) }
	if snap.GoverningGeneration != anchor.GoverningGeneration { return fail(errors.New("TRUST_GENERATION_MISMATCH")) }
	if s.Clock.Now().UTC().Before(anchor.ObservedWallTimeHighWater) { return fail(errors.New("TIME_UNTRUSTED")) }
	return trustedCurrent{Root:root, ORK:ork, Snapshot:snap, Anchor:anchor}, nil
}
