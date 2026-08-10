//go:build linux

package cli

import (
	"errors"
	"os"

	"golang.org/x/sys/unix"
)

func readSecretLine(in *os.File) ([]byte, error) {
	fd := int(in.Fd())
	state, err := unix.IoctlGetTermios(fd, unix.TCGETS)
	if err != nil {
		if errors.Is(err, unix.ENOTTY) {
			return readLine(in)
		}
		return nil, err
	}
	next := *state
	next.Lflag &^= unix.ECHO
	if err := unix.IoctlSetTermios(fd, unix.TCSETS, &next); err != nil {
		return nil, err
	}
	defer func() { _ = unix.IoctlSetTermios(fd, unix.TCSETS, state) }()
	return readLine(in)
}
