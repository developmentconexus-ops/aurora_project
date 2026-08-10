//go:build linux

package cli

import (
	"os"

	"golang.org/x/sys/unix"
)

func readSecretLine(f *os.File) ([]byte, bool, error) {
	fd := int(f.Fd())
	term, err := unix.IoctlGetTermios(fd, unix.TCGETS)
	if err != nil {
		secret, readErr := readLine(f)
		return secret, false, readErr
	}
	original := *term
	term.Lflag &^= unix.ECHO
	if err := unix.IoctlSetTermios(fd, unix.TCSETS, term); err != nil { return nil, true, err }
	defer unix.IoctlSetTermios(fd, unix.TCSETS, &original)
	secret, err := readLine(f)
	return secret, true, err
}
