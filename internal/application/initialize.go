package application

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/identity"
	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

var ErrAlreadyInitialized = errors.New("Aurora is already initialized")

type InitializeResult struct { AuroraID identity.AuroraID `json:"aurora_id"`; OwnerOperatorID identity.OperatorID `json:"owner_operator_id"`; GoverningGeneration uint64 `json:"governing_generation"` }

func (s *Service) Initialize(ctx context.Context, passphrase []byte) (InitializeResult, error) {
	if s.State == nil || s.Trust == nil || s.Clock == nil { return InitializeResult{}, errors.New("service dependencies are incomplete") }
	if _, err := s.Trust.LoadRootEnvelope(ctx); err == nil { return InitializeResult{}, ErrAlreadyInitialized } else if !errors.Is(err,fs.ErrNotExist) { return InitializeResult{}, fmt.Errorf("check existing owner root: %w",err) }
	env, ork, err := newRootEnvelope(passphrase); if err != nil { return InitializeResult{}, err }; defer zero(ork)
	if err := s.Trust.StoreRootEnvelope(ctx,env); err != nil { return InitializeResult{}, fmt.Errorf("publish owner root: %w",err) }
	auroraID, err := identity.NewAuroraID(); if err != nil { return InitializeResult{}, err }
	now := s.Clock.Now().UTC()
	id := identity.AuroraIdentity{AuroraID:auroraID, OwnerOperatorID:"OWNER-LEANDRO", CreatedAt:now, IdentityRevision:1, CapabilityContractVersion:"0.2.0"}
	authority := map[string]any{"revision":uint64(1),"predecessor_revision":nil,"revalidation_required":false,"grants":[]any{map[string]any{"authority_id":"AUTH-OWNER-INITIAL","authority_revision":uint64(1),"subject_operator_id":"OWNER-LEANDRO","actor_id":"AURORA_CORE","project_scope":[]string{"*"},"permitted_action_classes":[]string{"project.create","project.state.transition","authority.admin","export","restore","migration"},"lifecycle_status":"ACTIVE","issued_at":now,"provenance":"owner-bootstrap"}}}
	authorityJSON, err := json.Marshal(authority); if err != nil { return InitializeResult{}, err }
	snap := ports.CurrentSnapshot{Identity:id, AuthorityStateJSON:authorityJSON, CurrentAuthorityRevision:1, GoverningGeneration:1}
	mac, err := governingMAC(ork,snap); if err != nil { return InitializeResult{}, err }
	_, err = s.State.Bootstrap(ctx,ports.BootstrapMutation{Identity:id,AuthorityStateJSON:authorityJSON,CurrentAuthorityRevision:1,GoverningGeneration:1,GoverningMAC:mac}); if err != nil { return InitializeResult{}, fmt.Errorf("bootstrap state: %w",err) }
	anchor := ports.Anchor{Version:1,RootID:env.RootID,AuroraID:auroraID,GoverningGeneration:1,ObservedWallTimeHighWater:now}
	am, err := anchorMAC(ork,anchor); if err != nil { return InitializeResult{},err }; anchor.HMAC=encodeMAC(am)
	if err := s.Trust.PublishAnchor(ctx,anchor); err != nil { return InitializeResult{},fmt.Errorf("publish owner anchor: %w",err) }
	return InitializeResult{AuroraID:auroraID,OwnerOperatorID:id.OwnerOperatorID,GoverningGeneration:1},nil
}
