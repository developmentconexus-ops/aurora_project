package ports

import (
	"context"
	"errors"
	"time"
)

var (
	ErrRootNotFound   = errors.New("owner root envelope not found")
	ErrAnchorNotFound = errors.New("owner trust anchor not found")
)

type KDFParams struct {
	Algorithm   string `json:"algorithm"`
	MemoryKiB   uint32 `json:"memory_kib"`
	Iterations  uint32 `json:"iterations"`
	Parallelism uint8  `json:"parallelism"`
}

type RootEnvelope struct {
	Version    int       `json:"version"`
	RootID     string    `json:"root_id"`
	KDF        KDFParams `json:"kdf"`
	Salt       string    `json:"salt"`
	Nonce      string    `json:"nonce"`
	WrappedORK string    `json:"wrapped_ork"`
}

type Anchor struct {
	Version                   int       `json:"version"`
	RootID                    string    `json:"root_id"`
	AuroraID                  string    `json:"aurora_id"`
	GoverningGeneration       uint64    `json:"governing_generation"`
	ObservedWallTimeHighWater time.Time `json:"observed_wall_time_high_water"`
	HMAC                      string    `json:"hmac"`
}

type OwnerTrustStore interface {
	LoadRootEnvelope(context.Context) (RootEnvelope, error)
	StoreRootEnvelope(context.Context, RootEnvelope) error
	LoadAnchor(context.Context) (Anchor, error)
	PublishAnchor(context.Context, Anchor) error
}
