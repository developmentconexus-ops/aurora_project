package identity

import "time"

type AuroraID string
type OperatorID string

type AuroraIdentity struct {
	AuroraID                  AuroraID  `json:"aurora_id"`
	OwnerOperatorID           OperatorID `json:"owner_operator_id"`
	CreatedAt                 time.Time `json:"created_at"`
	IdentityRevision          uint64    `json:"identity_revision"`
	CapabilityContractVersion string    `json:"capability_contract_version"`
}
