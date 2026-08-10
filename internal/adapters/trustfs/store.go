package trustfs

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/developmentconexus-ops/aurora_project/internal/ports"
)

type Store struct { trustDir string }

func New(dataDir string) *Store { return &Store{trustDir:filepath.Join(dataDir,"trust")} }

func (s *Store) LoadRootEnvelope(context.Context) (ports.RootEnvelope, error) {
	var out ports.RootEnvelope
	if err := readJSON(filepath.Join(s.trustDir,"owner-root.json"), &out); err != nil { return out, err }
	return out, nil
}
func (s *Store) StoreRootEnvelope(_ context.Context, v ports.RootEnvelope) error { return publishJSON(filepath.Join(s.trustDir,"owner-root.json"), v) }
func (s *Store) LoadAnchor(context.Context) (ports.Anchor, error) {
	var out ports.Anchor
	if err := readJSON(filepath.Join(s.trustDir,"owner-anchor.json"), &out); err != nil { return out, err }
	return out, nil
}
func (s *Store) PublishAnchor(_ context.Context, v ports.Anchor) error { return publishJSON(filepath.Join(s.trustDir,"owner-anchor.json"), v) }

func publishJSON(path string, v any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil { return err }
	data, err := json.MarshalIndent(v,"","  ")
	if err != nil { return err }
	data = append(data,'\n')
	return atomicPublish(path,data,0o600)
}
func readJSON(path string, v any) error {
	data, err := os.ReadFile(path)
	if err != nil { return err }
	if err := json.Unmarshal(data,v); err != nil { return fmt.Errorf("decode %s: %w", filepath.Base(path), err) }
	return nil
}

var _ ports.OwnerTrustStore = (*Store)(nil)
