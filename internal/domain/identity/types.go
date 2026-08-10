package identity

import "time"

type AuroraID string
type OperatorID string

type AuroraIdentity struct {
	AuroraID                  AuroraID
	OwnerOperatorID           OperatorID
	CreatedAt                 time.Time
	IdentityRevision          uint64
	CapabilityContractVersion string
}
