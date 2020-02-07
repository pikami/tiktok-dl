package utils

import (
	"io/ioutil"
)

// ReadFileAsString - Returns contents of given file
func ReadFileAsString(fileName string) string {
	content, err := ioutil.ReadFile(fileName)
	CheckErr(err)
	return string(content)
}
