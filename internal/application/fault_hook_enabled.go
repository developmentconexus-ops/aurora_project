//go:build aurora_testhooks

package application

import (
	"fmt"
	"os"
	"time"
)

func pauseAfterStateCommitForTest() error {
	marker := os.Getenv("AURORA_TEST_PAUSE_AFTER_STATE_COMMIT")
	if marker == "" {
		return nil
	}
	if err := os.WriteFile(marker, []byte("sqlite-committed\n"), 0o600); err != nil {
		return fmt.Errorf("publish test crash marker: %w", err)
	}
	for {
		time.Sleep(time.Hour)
	}
}
