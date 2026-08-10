package cli

import (
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

// readLine intentionally does not buffer past the newline. Multiple secret
// prompts may share a piped stdin stream, so read-ahead by a short-lived
// buffered reader would consume bytes belonging to the next prompt.
func readLine(r io.Reader) ([]byte, error) {
	line := make([]byte, 0, 64)
	var one [1]byte
	for {
		n, err := r.Read(one[:])
		if n == 1 {
			if one[0] == '\n' { break }
			line = append(line, one[0])
		}
		if err != nil {
			if err == io.EOF { break }
			return nil, err
		}
	}
	line = bytes.TrimRight(line, "\r")
	if len(line) == 0 { return nil, fmt.Errorf("empty secret") }
	return line, nil
}
