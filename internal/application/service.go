package application

import "github.com/developmentconexus-ops/aurora_project/internal/ports"

type Service struct {
	State ports.StateStore
	Trust ports.OwnerTrustStore
	Clock ports.Clock
}
