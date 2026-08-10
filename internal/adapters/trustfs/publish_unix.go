//go:build !windows

package trustfs

import (
	"os"
	"path/filepath"
)

func atomicPublish(path string, data []byte, mode os.FileMode) error {
	dir := filepath.Dir(path)
	tmp, err := os.CreateTemp(dir, "."+filepath.Base(path)+".tmp-*")
	if err != nil { return err }
	tmpName := tmp.Name()
	defer os.Remove(tmpName)
	if err := tmp.Chmod(mode); err != nil { tmp.Close(); return err }
	if _, err := tmp.Write(data); err != nil { tmp.Close(); return err }
	if err := tmp.Sync(); err != nil { tmp.Close(); return err }
	if err := tmp.Close(); err != nil { return err }
	if err := os.Rename(tmpName,path); err != nil { return err }
	d, err := os.Open(dir)
	if err != nil { return err }
	defer d.Close()
	return d.Sync()
}
