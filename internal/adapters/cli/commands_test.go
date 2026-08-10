package cli

import (
	"bytes"
	"encoding/json"
	"testing"
)

type fixedSecretReader struct{ secret []byte }

func (f fixedSecretReader) ReadSecret(string) ([]byte, error) {
	return append([]byte(nil), f.secret...), nil
}

func TestInitStatusContinuityAndReinitializeRejection(t *testing.T) {
	dataDir := t.TempDir()
	secrets := fixedSecretReader{secret: []byte("fixture-owner-passphrase")}

	var initOut, initErr bytes.Buffer
	if code := runWithSecretReader([]string{"--json", "--data-dir", dataDir, "init"}, &initOut, &initErr, secrets); code != 0 {
		t.Fatalf("init code=%d stderr=%q", code, initErr.String())
	}
	var initialized map[string]any
	if err := json.Unmarshal(initOut.Bytes(), &initialized); err != nil {
		t.Fatal(err)
	}
	auroraID, _ := initialized["aurora_id"].(string)
	if auroraID == "" {
		t.Fatalf("init output missing aurora_id: %s", initOut.String())
	}

	var statusOut, statusErr bytes.Buffer
	if code := runWithSecretReader([]string{"--json", "--data-dir", dataDir, "status"}, &statusOut, &statusErr, secrets); code != 0 {
		t.Fatalf("status code=%d stderr=%q", code, statusErr.String())
	}
	var status map[string]any
	if err := json.Unmarshal(statusOut.Bytes(), &status); err != nil {
		t.Fatal(err)
	}
	if got, _ := status["aurora_id"].(string); got != auroraID {
		t.Fatalf("status aurora_id=%q want %q", got, auroraID)
	}

	var againOut, againErr bytes.Buffer
	if code := runWithSecretReader([]string{"--json", "--data-dir", dataDir, "init"}, &againOut, &againErr, secrets); code == 0 {
		t.Fatalf("second init unexpectedly succeeded: %s", againOut.String())
	}

	statusOut.Reset()
	statusErr.Reset()
	if code := runWithSecretReader([]string{"--json", "--data-dir", dataDir, "status"}, &statusOut, &statusErr, secrets); code != 0 {
		t.Fatalf("status after rejected init code=%d stderr=%q", code, statusErr.String())
	}
	status = nil
	if err := json.Unmarshal(statusOut.Bytes(), &status); err != nil {
		t.Fatal(err)
	}
	if got, _ := status["aurora_id"].(string); got != auroraID {
		t.Fatalf("identity changed after rejected init: got=%q want=%q", got, auroraID)
	}
}
