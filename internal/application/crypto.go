package application

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/developmentconexus-ops/aurora_project/internal/domain/identity"
	"github.com/developmentconexus-ops/aurora_project/internal/ports"
	"github.com/gowebpki/jcs"
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/hkdf"
)

var ErrUnsupportedKDFProfile = errors.New("unsupported owner-root KDF profile")

const (
	rootEnvelopeVersion = 1
	argonMemoryKiB      = uint32(64 * 1024)
	argonIterations     = uint32(3)
	argonParallelism    = uint8(4)
	kekBytes            = uint32(32)
)

func newRootEnvelope(passphrase []byte) (ports.RootEnvelope, []byte, error) {
	ork := make([]byte, 32)
	if _, err := rand.Read(ork); err != nil {
		return ports.RootEnvelope{}, nil, fmt.Errorf("generate ORK: %w", err)
	}
	var rawID [16]byte
	if _, err := rand.Read(rawID[:]); err != nil {
		return ports.RootEnvelope{}, nil, fmt.Errorf("generate root ID: %w", err)
	}
	env, err := wrapORK(passphrase, "ROOT-"+hex.EncodeToString(rawID[:]), ork)
	if err != nil {
		return ports.RootEnvelope{}, nil, err
	}
	return env, ork, nil
}

func rewrapRootEnvelope(oldPassphrase, newPassphrase []byte, current ports.RootEnvelope) (ports.RootEnvelope, error) {
	ork, err := unlockORK(oldPassphrase, current)
	if err != nil {
		return ports.RootEnvelope{}, err
	}
	defer zero(ork)
	return wrapORK(newPassphrase, current.RootID, ork)
}

func wrapORK(passphrase []byte, rootID string, ork []byte) (ports.RootEnvelope, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return ports.RootEnvelope{}, fmt.Errorf("generate KDF salt: %w", err)
	}
	kek := argon2.IDKey(passphrase, salt, argonIterations, argonMemoryKiB, argonParallelism, kekBytes)
	defer zero(kek)
	block, err := aes.NewCipher(kek)
	if err != nil { return ports.RootEnvelope{}, err }
	gcm, err := cipher.NewGCM(block)
	if err != nil { return ports.RootEnvelope{}, err }
	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil { return ports.RootEnvelope{}, err }
	ciphertext := gcm.Seal(nil, nonce, ork, []byte(rootID))
	enc := base64.RawURLEncoding
	return ports.RootEnvelope{
		Version: rootEnvelopeVersion, RootID: rootID, KDF: "argon2id",
		MemoryKiB: argonMemoryKiB, Iterations: argonIterations, Parallelism: argonParallelism,
		Salt: enc.EncodeToString(salt), Nonce: enc.EncodeToString(nonce), WrappedORK: enc.EncodeToString(ciphertext),
	}, nil
}

func unlockORK(passphrase []byte, env ports.RootEnvelope) ([]byte, error) {
	if env.Version != rootEnvelopeVersion || env.KDF != "argon2id" || env.MemoryKiB != argonMemoryKiB || env.Iterations != argonIterations || env.Parallelism != argonParallelism {
		return nil, ErrUnsupportedKDFProfile
	}
	enc := base64.RawURLEncoding
	salt, err := enc.DecodeString(env.Salt)
	if err != nil { return nil, fmt.Errorf("decode KDF salt: %w", err) }
	nonce, err := enc.DecodeString(env.Nonce)
	if err != nil { return nil, fmt.Errorf("decode root nonce: %w", err) }
	wrapped, err := enc.DecodeString(env.WrappedORK)
	if err != nil { return nil, fmt.Errorf("decode wrapped ORK: %w", err) }
	if len(salt) != 16 || len(nonce) != 12 || len(wrapped) < 32+16 {
		return nil, errors.New("invalid owner-root envelope lengths")
	}
	kek := argon2.IDKey(passphrase, salt, env.Iterations, env.MemoryKiB, env.Parallelism, kekBytes)
	defer zero(kek)
	block, err := aes.NewCipher(kek)
	if err != nil { return nil, err }
	gcm, err := cipher.NewGCM(block)
	if err != nil { return nil, err }
	ork, err := gcm.Open(nil, nonce, wrapped, []byte(env.RootID))
	if err != nil { return nil, errors.New("owner-root unlock failed") }
	if len(ork) != 32 { zero(ork); return nil, errors.New("invalid ORK length") }
	return ork, nil
}

func derivePurposeKey(ork []byte, purpose string) ([]byte, error) {
	key := make([]byte, 32)
	if _, err := io.ReadFull(hkdf.New(sha256.New, ork, nil, []byte(purpose)), key); err != nil {
		return nil, err
	}
	return key, nil
}

type governingDescriptor struct {
	Version             int                     `json:"version"`
	Aurora              identity.AuroraIdentity `json:"aurora"`
	GoverningGeneration uint64                  `json:"governing_generation"`
	Authority           json.RawMessage         `json:"authority"`
}

func governingMAC(ork []byte, id identity.AuroraIdentity, generation uint64, authorityJSON []byte) ([]byte, error) {
	key, err := derivePurposeKey(ork, "aurora/m0/governing-state/v1")
	if err != nil { return nil, err }
	defer zero(key)
	raw, err := json.Marshal(governingDescriptor{Version:1, Aurora:id, GoverningGeneration:generation, Authority:json.RawMessage(authorityJSON)})
	if err != nil { return nil, err }
	canonical, err := jcs.Transform(raw)
	if err != nil { return nil, err }
	mac := hmac.New(sha256.New, key)
	_, _ = mac.Write(canonical)
	return mac.Sum(nil), nil
}

type anchorDescriptor struct {
	Version int `json:"version"`
	RootID string `json:"root_id"`
	AuroraID identity.AuroraID `json:"aurora_id"`
	GoverningGeneration uint64 `json:"governing_generation"`
	ObservedWallTimeHighWater string `json:"observed_wall_time_high_water"`
}

func anchorMAC(ork []byte, a ports.Anchor) ([]byte, error) {
	key, err := derivePurposeKey(ork, "aurora/m0/trust-anchor/v1")
	if err != nil { return nil, err }
	defer zero(key)
	raw, err := json.Marshal(anchorDescriptor{Version:a.Version, RootID:a.RootID, AuroraID:a.AuroraID, GoverningGeneration:a.GoverningGeneration, ObservedWallTimeHighWater:a.ObservedWallTimeHighWater.UTC().Format("2006-01-02T15:04:05.999999999Z07:00")})
	if err != nil { return nil, err }
	canonical, err := jcs.Transform(raw)
	if err != nil { return nil, err }
	mac := hmac.New(sha256.New, key)
	_, _ = mac.Write(canonical)
	return mac.Sum(nil), nil
}

func encodeMAC(v []byte) string { return base64.RawURLEncoding.EncodeToString(v) }
func decodeMAC(v string) ([]byte, error) { return base64.RawURLEncoding.DecodeString(v) }
func zero(v []byte) { for i := range v { v[i] = 0 } }
