package golib

import (
	"os"

	log "github.com/rustyeddy/logrus"
)

// ResetDirectory will create a pristine test directory to begin our testing
func ResetDirectory(path string) {
	// If the file does not exist create it and be ready to go.
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.MkdirAll(path, 0755); err == nil {
			log.Fatal("failed to create ", path, err)
		}
	}

	// Check again for the basedir, if we don't have it now we never will
	if _, err := os.Stat(path); err != nil {
		log.Fatal("failed to create test directory ", path, err)
	}
}
