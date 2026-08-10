//go:build windows

package trustfs

import (
	"os"
	"path/filepath"

	"golang.org/x/sys/windows"
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
	from, err := windows.UTF16PtrFromString(tmpName)
	if err != nil { return err }
	to, err := windows.UTF16PtrFromString(path)
	if err != nil { return err }
	return windows.MoveFileEx(from,to,windows.MOVEFILE_REPLACE_EXISTING|windows.MOVEFILE_WRITE_THROUGH)
}
