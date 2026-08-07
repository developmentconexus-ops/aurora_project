package store

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
)

func openExistingDB(path string) (*sql.DB, error) {
	if _, err := os.Stat(path); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, ErrMissingStore
		}
		return nil, fmt.Errorf("stat store: %w", err)
	}
	return openDB(path)
}
