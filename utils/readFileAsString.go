package utils

import (
	"io/ioutil"
	"log"
)

// ReadFileAsString - Returns contents of given file
func ReadFileAsString(fileName string) string {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
