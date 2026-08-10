//go:build !linux && !windows

package cli

import "os"

func readSecretLine(in *os.File) ([]byte, error) { return readLine(in) }
