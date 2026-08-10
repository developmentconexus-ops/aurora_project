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

	"github.com/developmentconexus-ops/aurora_project/internal/ports"
	"github.com/gowebpki/jcs"
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/hkdf"
)

const (
	rootEnvelopeVersion = 1
	allowedMemoryKiB    = 64 * 1024
	allowedIterations   = 3
	allowedParallelism  = 4
)

var errUnsupportedKDF = errors.New("unsupported owner-root KDF parameters")

func newRootEnvelope(passphrase []byte) (ports.RootEnvelope, []byte, error) {
	ork := make([]byte, 32)
	if _, err := rand.Read(ork); err != nil {
		return ports.RootEnvelope{}, nil, fmt.Errorf("generate ORK: %w", err)
	}
	rootID, err := randomIdentifier("ROOT-")
	if err != nil {
		clear(ork)
		return ports.RootEnvelope{}, nil, err
	}
	env, err := wrapORK(rootID, ork, passphrase)
	if err != nil {
		clear(ork)
		return ports.RootEnvelope{}, nil, err
	}
	return env, ork, nil
}

func rewrapRootEnvelope(env ports.RootEnvelope, oldPassphrase, newPassphrase []byte) (ports.RootEnvelope, error) {
	ork, err := unlockORK(env, oldPassphrase)
	if err != nil {
		return ports.RootEnvelope{}, err
	}
	defer clear(ork)
	return wrapORK(env.RootID, ork, newPassphrase)
}

func wrapORK(rootID string, ork, passphrase []byte) (ports.RootEnvelope, error) {
	if len(passphrase) == 0 {
		return ports.RootEnvelope{}, errors.New("owner passphrase must not be empty")
	}
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return ports.RootEnvelope{}, fmt.Errorf("generate KDF salt: %w", err)
	}
	params := ports.KDFParams{Algorithm: "argon2id", MemoryKiB: allowedMemoryKiB, Iterations: allowedIterations, Parallelism: allowedParallelism}
	kek := argon2.IDKey(passphrase, salt, params.Iterations, params.MemoryKiB, params.Parallelism, 32)
	defer clear(kek)
	block, err := aes.NewCipher(kek)
	if err != nil {
		return ports.RootEnvelope{}, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return ports.RootEnvelope{}, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return ports.RootEnvelope{}, fmt.Errorf("generate ORK nonce: %w", err)
	}
	wrapped := gcm.Seal(nil, nonce, ork, []byte(rootID))
	enc := base64.RawURLEncoding
	return ports.RootEnvelope{
		Version:    rootEnvelopeVersion,
		RootID:     rootID,
		KDF:        params,
		Salt:       enc.EncodeToString(salt),
		Nonce:      enc.EncodeToString(nonce),
		WrappedORK: enc.EncodeToString(wrapped),
	}, nil
}

func unlockORK(env ports.RootEnvelope, passphrase []byte) ([]byte, error) {
	if err := validateRootEnvelope(env); err != nil {
		return nil, err
	}
	enc := base64.RawURLEncoding
	salt, err := enc.DecodeString(env.Salt)
	if err != nil || len(salt) != 16 {
		return nil, errors.New("invalid owner-root salt")
	}
	nonce, err := enc.DecodeString(env.Nonce)
	if err != nil {
		return nil, errors.New("invalid owner-root nonce")
	}
	wrapped, err := enc.DecodeString(env.WrappedORK)
	if err != nil {
		return nil, errors.New("invalid wrapped ORK")
	}
	kek := argon2.IDKey(passphrase, salt, env.KDF.Iterations, env.KDF.MemoryKiB, env.KDF.Parallelism, 32)
	defer clear(kek)
	block, err := aes.NewCipher(kek)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	if len(nonce) != gcm.NonceSize() {
		return nil, errors.New("invalid owner-root nonce length")
	}
	ork, err := gcm.Open(nil, nonce, wrapped, []byte(env.RootID))
	if err != nil {
		return nil, errors.New("owner-root authentication failed")
	}
	if len(ork) != 32 {
		clear(ork)
		return nil, errors.New("invalid ORK length")
	}
	return ork, nil
}

func validateRootEnvelope(env ports.RootEnvelope) error {
	if env.Version != rootEnvelopeVersion || env.RootID == "" {
		return errUnsupportedKDF
	}
	if env.KDF.Algorithm != "argon2id" || env.KDF.MemoryKiB != allowedMemoryKiB || env.KDF.Iterations != allowedIterations || env.KDF.Parallelism != allowedParallelism {
		return errUnsupportedKDF
	}
	return nil
}

func derivePurposeKey(ork []byte, purpose string) ([]byte, error) {
	r := hkdf.New(sha256.New, ork, nil, []byte("aurora/m0/"+purpose))
	key := make([]byte, 32)
	if _, err := io.ReadFull(r, key); err != nil {
		return nil, err
	}
	return key, nil
}

func governingMAC(ork []byte, snapshot any) ([]byte, error) {
	key, err := derivePurposeKey(ork, "governing-state")
	if err != nil {
		return nil, err
	}
	defer clear(key)
	canonical, err := canonicalJSON(snapshot)
	if err != nil {
		return nil, err
	}
	mac := hmac.New(sha256.New, key)
	_, _ = mac.Write(canonical)
	return mac.Sum(nil), nil
}

func anchorMAC(ork []byte, anchor ports.Anchor) ([]byte, error) {
	key, err := derivePurposeKey(ork, "trust-anchor")
	if err != nil {
		return nil, err
	}
	defer clear(key)
	input := struct {
		Version                   int    `json:"version"`
		RootID                    string `json:"root_id"`
		AuroraID                  string `json:"aurora_id"`
		GoverningGeneration       uint64 `json:"governing_generation"`
		ObservedWallTimeHighWater string `json:"observed_wall_time_high_water"`
	}{anchor.Version, anchor.RootID, anchor.AuroraID, anchor.GoverningGeneration, anchor.ObservedWallTimeHighWater.UTC().Format("2006-01-02T15:04:05.999999999Z07:00")}
	canonical, err := canonicalJSON(input)
	if err != nil {
		return nil, err
	}
	mac := hmac.New(sha256.New, key)
	_, _ = mac.Write(canonical)
	return mac.Sum(nil), nil
}

func canonicalJSON(v any) ([]byte, error) {
	raw, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return jcs.Transform(raw)
}

func randomIdentifier(prefix string) (string, error) {
	var raw [16]byte
	if _, err := rand.Read(raw[:]); err != nil {
		return "", err
	}
	return prefix + hex.EncodeToString(raw[:]), nil
}
