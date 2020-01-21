package utils

import (
	"io/ioutil"
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

// ReadFileToString - Reads file and returns content
func ReadFileToString(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(content)
}
