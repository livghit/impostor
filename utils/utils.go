package utils

import (
	"log"
	"os"
	"path/filepath"

)

// Remove hardcoded database name and get from env
func DatabasePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	path := filepath.Join(homeDir, "Development", "OpenSource", "impostor", "test.db")

	return path
}
