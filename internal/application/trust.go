package application

import (
	"context"
	"crypto/hmac"
	"errors"
	"fmt"
	"io/fs"
	"time"

	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

type TrustStatus string

const (
	TrustNormal               TrustStatus = "NORMAL"
	TrustAnchorLag            TrustStatus = "ANCHOR_LAG"
	TrustStateRollback        TrustStatus = "STATE_ROLLBACK"
	TrustInvalidDBMAC         TrustStatus = "INVALID_DB_MAC"
	TrustInvalidAnchorMAC     TrustStatus = "INVALID_ANCHOR_MAC"
	TrustMissingAnchor        TrustStatus = "MISSING_ANCHOR"
	TrustTimeUntrusted        TrustStatus = "TIME_UNTRUSTED"
	TrustRevalidationRequired TrustStatus = "REVALIDATION_REQUIRED"
)

var ErrReconcileNotAllowed = errors.New("anchor reconciliation is allowed only for ANCHOR_LAG")

type TrustResult struct {
	Status           TrustStatus `json:"status"`
	DBGeneration     uint64      `json:"db_generation"`
	AnchorGeneration uint64      `json:"anchor_generation"`
	Reason           string      `json:"reason"`
}

type TrustStateError struct{ Result TrustResult }
func (e *TrustStateError) Error() string { return string(e.Result.Status) + ": " + e.Result.Reason }

type trustedCurrent struct {
	Root     ports.RootEnvelope
	ORK      []byte
	Snapshot ports.CurrentSnapshot
	Anchor   ports.Anchor
}

type trustAssessment struct {
	trustedCurrent
	Result TrustResult
}

func (s *Service) ClassifyTrust(ctx context.Context, passphrase []byte) (TrustResult, error) {
	a, err := s.assessTrust(ctx, passphrase)
	if a.ORK != nil { zero(a.ORK) }
	if err != nil { return TrustResult{}, err }
	return a.Result, nil
}

func (s *Service) assessTrust(ctx context.Context, passphrase []byte) (trustAssessment, error) {
	var out trustAssessment
	root, err := s.Trust.LoadRootEnvelope(ctx)
	if err != nil { return out, fmt.Errorf("load owner root: %w", err) }
	ork, err := unlockORK(passphrase, root)
	if err != nil { return out, err }
	out.Root, out.ORK = root, ork
	fail := func(err error) (trustAssessment,error) { zero(ork); return trustAssessment{}, err }

	snap, err := s.State.LoadCurrent(ctx)
	if err != nil { return fail(fmt.Errorf("load current state: %w", err)) }
	out.Snapshot = snap
	wantDB, err := governingMAC(ork, snap)
	if err != nil { return fail(err) }
	if !hmac.Equal(wantDB, snap.GoverningMAC) {
		out.Result = TrustResult{Status:TrustInvalidDBMAC, DBGeneration:snap.GoverningGeneration, Reason:"governing descriptor HMAC does not verify"}
		return out, nil
	}

	anchor, err := s.Trust.LoadAnchor(ctx)
	if errors.Is(err, fs.ErrNotExist) {
		out.Result = TrustResult{Status:TrustMissingAnchor, DBGeneration:snap.GoverningGeneration, Reason:"current owner anchor is missing"}
		return out, nil
	}
	if err != nil { return fail(fmt.Errorf("load owner anchor: %w", err)) }
	out.Anchor = anchor
	out.Result.DBGeneration = snap.GoverningGeneration
	out.Result.AnchorGeneration = anchor.GoverningGeneration

	provided, err := decodeMAC(anchor.HMAC)
	if err != nil {
		out.Result.Status = TrustInvalidAnchorMAC
		out.Result.Reason = "owner anchor HMAC encoding is invalid"
		return out, nil
	}
	wantAnchor, err := anchorMAC(ork, anchor)
	if err != nil { return fail(err) }
	if !hmac.Equal(provided,wantAnchor) || anchor.AuroraID != snap.Identity.AuroraID || anchor.RootID != root.RootID {
		out.Result.Status = TrustInvalidAnchorMAC
		out.Result.Reason = "owner anchor authentication or identity binding does not verify"
		return out, nil
	}

	switch {
	case snap.GoverningGeneration > anchor.GoverningGeneration:
		out.Result.Status = TrustAnchorLag
		out.Result.Reason = "operational DB committed a newer authenticated generation than the anchor"
		return out, nil
	case snap.GoverningGeneration < anchor.GoverningGeneration:
		out.Result.Status = TrustStateRollback
		out.Result.Reason = "operational DB generation is behind authenticated anchor high-water"
		return out, nil
	}

	if s.Clock.Now().UTC().Before(anchor.ObservedWallTimeHighWater) {
		out.Result.Status = TrustTimeUntrusted
		out.Result.Reason = "current wall time is behind authenticated observed high-water"
		return out, nil
	}
	state, err := decodeAuthorityState(snap.AuthorityStateJSON)
	if err != nil { return fail(err) }
	if state.RevalidationRequired {
		out.Result.Status = TrustRevalidationRequired
		out.Result.Reason = "restored authority requires explicit owner revalidation"
		return out, nil
	}
	out.Result.Status = TrustNormal
	out.Result.Reason = "governing state and owner trust anchor are coherent"
	return out, nil
}

func (s *Service) loadTrustedCurrent(ctx context.Context, passphrase []byte) (trustedCurrent, error) {
	a, err := s.assessTrust(ctx, passphrase)
	if err != nil { return trustedCurrent{}, err }
	if a.Result.Status != TrustNormal {
		if a.ORK != nil { zero(a.ORK) }
		return trustedCurrent{}, &TrustStateError{Result:a.Result}
	}
	return a.trustedCurrent, nil
}

func (s *Service) ReconcileAnchor(ctx context.Context, passphrase []byte) error {
	a, err := s.assessTrust(ctx, passphrase)
	if err != nil { return err }
	defer zero(a.ORK)
	if a.Result.Status != TrustAnchorLag { return ErrReconcileNotAllowed }
	now := s.Clock.Now().UTC()
	anchor := a.Anchor
	anchor.GoverningGeneration = a.Snapshot.GoverningGeneration
	if now.After(anchor.ObservedWallTimeHighWater) { anchor.ObservedWallTimeHighWater = now }
	mac, err := anchorMAC(a.ORK, anchor)
	if err != nil { return err }
	anchor.HMAC = encodeMAC(mac)
	return s.Trust.PublishAnchor(ctx, anchor)
}

func (s *Service) RotateOwnerPassphrase(ctx context.Context, oldPassphrase, newPassphrase []byte) error {
	if len(newPassphrase) == 0 { return errors.New("new owner passphrase cannot be empty") }
	trusted, err := s.loadTrustedCurrent(ctx, oldPassphrase)
	if err != nil { return err }
	defer zero(trusted.ORK)
	next, err := rewrapRootEnvelope(oldPassphrase, newPassphrase, trusted.Root)
	if err != nil { return err }
	return s.Trust.StoreRootEnvelope(ctx, next)
}

var _ = time.Time{}
