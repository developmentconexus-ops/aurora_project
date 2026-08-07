package store

import "fmt"

func SQLiteVersion(path string) (string, error) {
	db, err := openExistingDB(path)
	if err != nil {
		return "", err
	}
	defer db.Close()
	var version string
	if err := db.QueryRow(`SELECT sqlite_version()`).Scan(&version); err != nil {
		return "", fmt.Errorf("sqlite runtime version: %w", err)
	}
	return version, nil
}
