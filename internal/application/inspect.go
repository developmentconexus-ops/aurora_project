package application

import (
	"context"
	"errors"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/identity"
)

type InspectResult struct { AuroraID identity.AuroraID `json:"aurora_id"`; OwnerOperatorID identity.OperatorID `json:"owner_operator_id"`; CurrentAuthorityRevision uint64 `json:"current_authority_revision"`; GoverningGeneration uint64 `json:"governing_generation"`; TrustStatus string `json:"trust_status"` }

func (s *Service) Inspect(ctx context.Context, passphrase []byte) (InspectResult,error) {
	if s.State==nil || s.Trust==nil || s.Clock==nil { return InspectResult{},errors.New("service dependencies are incomplete") }
	trusted,err:=s.loadTrustedCurrent(ctx,passphrase); if err!=nil{return InspectResult{},err}; defer zero(trusted.ORK)
	return InspectResult{AuroraID:trusted.Snapshot.Identity.AuroraID,OwnerOperatorID:trusted.Snapshot.Identity.OwnerOperatorID,CurrentAuthorityRevision:trusted.Snapshot.CurrentAuthorityRevision,GoverningGeneration:trusted.Snapshot.GoverningGeneration,TrustStatus:"NORMAL"},nil
}
