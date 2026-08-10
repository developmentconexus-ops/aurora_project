package trustfs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

const (
	rootFileName   = "owner-root.json"
	anchorFileName = "owner-anchor.json"
)

type Store struct {
	dir string
}

func New(dataDir string) *Store { return &Store{dir: filepath.Join(dataDir, "trust")} }

func (s *Store) LoadRootEnvelope(_ context.Context) (ports.RootEnvelope, error) {
	var out ports.RootEnvelope
	if err := s.load(rootFileName, &out); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return ports.RootEnvelope{}, ports.ErrRootNotFound
		}
		return ports.RootEnvelope{}, err
	}
	return out, nil
}

func (s *Store) StoreRootEnvelope(_ context.Context, env ports.RootEnvelope) error {
	return s.store(rootFileName, env)
}

func (s *Store) LoadAnchor(_ context.Context) (ports.Anchor, error) {
	var out ports.Anchor
	if err := s.load(anchorFileName, &out); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return ports.Anchor{}, ports.ErrAnchorNotFound
		}
		return ports.Anchor{}, err
	}
	return out, nil
}

func (s *Store) PublishAnchor(_ context.Context, anchor ports.Anchor) error {
	return s.store(anchorFileName, anchor)
}

func (s *Store) load(name string, out any) error {
	raw, err := os.ReadFile(filepath.Join(s.dir, name))
	if err != nil {
		return err
	}
	if err := json.Unmarshal(raw, out); err != nil {
		return fmt.Errorf("decode %s: %w", name, err)
	}
	return nil
}

func (s *Store) store(name string, value any) error {
	if err := os.MkdirAll(s.dir, 0o700); err != nil {
		return err
	}
	raw, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return atomicPublish(filepath.Join(s.dir, name), raw)
}
