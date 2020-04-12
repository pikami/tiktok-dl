package utils

import (
	"bufio"
	"io/ioutil"
	"os"

	checkErr "github.com/pikami/tiktok-dl/utils/checkErr"
)

type delegateString func(string)

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

// ReadFileLineByLine - Reads file line by line and calls delegate
func ReadFileLineByLine(path string, delegate delegateString) {
	file, err := os.Open(path)
	checkErr.CheckErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		delegate(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

// AppendToFile - Appends line to file
func AppendToFile(str string, filePath string) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkErr.CheckErr(err)

	defer f.Close()
	if _, err := f.WriteString(str + "\n"); err != nil {
		checkErr.CheckErr(err)
	}
}
