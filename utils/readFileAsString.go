package utils

import (
	"io/ioutil"

	checkErr "github.com/pikami/tiktok-dl/utils/checkErr"
)

// ReadFileAsString - Returns contents of given file
func ReadFileAsString(fileName string) string {
	content, err := ioutil.ReadFile(fileName)
	checkErr.CheckErr(err)
	return string(content)
}
