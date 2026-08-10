//go:build windows

package cli

import (
	"os"

	"golang.org/x/sys/windows"
)

func readSecretLine(in *os.File) ([]byte, error) {
	handle := windows.Handle(in.Fd())
	var mode uint32
	if err := windows.GetConsoleMode(handle, &mode); err != nil {
		return readLine(in)
	}
	if err := windows.SetConsoleMode(handle, mode&^windows.ENABLE_ECHO_INPUT); err != nil {
		return nil, err
	}
	defer func() { _ = windows.SetConsoleMode(handle, mode) }()
	return readLine(in)
}
