package cli

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func promptSecret(prompt string, promptOut io.Writer) ([]byte, error) {
	fmt.Fprint(promptOut, prompt)
	secret, terminal, err := readSecretLine(os.Stdin)
	if terminal { fmt.Fprintln(promptOut) }
	return secret, err
}

func readLine(r io.Reader) ([]byte, error) {
	line, err := bufio.NewReader(r).ReadBytes('\n')
	if err != nil && err != io.EOF { return nil, err }
	line = bytes.TrimRight(line, "\r\n")
	if len(line) == 0 { return nil, fmt.Errorf("empty secret") }
	return line, nil
}
