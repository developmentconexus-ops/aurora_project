package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type SecretReader interface {
	ReadSecret(prompt string) ([]byte, error)
}

type terminalSecretReader struct {
	in  *os.File
	out io.Writer
}

func newTerminalSecretReader(out io.Writer) SecretReader {
	return terminalSecretReader{in: os.Stdin, out: out}
}

func (r terminalSecretReader) ReadSecret(prompt string) ([]byte, error) {
	_, _ = fmt.Fprint(r.out, prompt)
	secret, err := readSecretLine(r.in)
	_, _ = fmt.Fprintln(r.out)
	return secret, err
}

func readLine(in *os.File) ([]byte, error) {
	line, err := bufio.NewReader(in).ReadString('\n')
	if err != nil && len(line) == 0 {
		return nil, err
	}
	return []byte(strings.TrimRight(line, "\r\n")), nil
}
