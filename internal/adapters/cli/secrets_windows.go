//go:build windows

package cli

import (
	"os"

	"golang.org/x/sys/windows"
)

func readSecretLine(f *os.File) ([]byte, bool, error) {
	h := windows.Handle(f.Fd())
	var mode uint32
	if err := windows.GetConsoleMode(h, &mode); err != nil {
		secret, readErr := readLine(f)
		return secret, false, readErr
	}
	if err := windows.SetConsoleMode(h, mode &^ windows.ENABLE_ECHO_INPUT); err != nil { return nil, true, err }
	defer windows.SetConsoleMode(h, mode)
	secret, err := readLine(f)
	return secret, true, err
}
