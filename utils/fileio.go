package utils

import (
	"os"
)

// CheckIfExists - Checks if file or directory exists
func CheckIfExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

// InitOutputDirectory - Creates output directory
func InitOutputDirectory(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}
}
