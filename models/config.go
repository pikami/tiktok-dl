package models

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Config - Runtime configuration
var Config struct {
	URL           string
	OutputPath    string
	BatchFilePath string
	Debug         bool
	MetaData      bool
	Deadline      int
}

// GetConfig - Returns Config object
func GetConfig() {
	outputPath := flag.String("output", "./downloads", "Output path")
	batchFilePath := flag.String("batch-file", "", "File containing URLs/Usernames to download, one value per line. Lines starting with '#', are considered as comments and ignored.")
	debug := flag.Bool("debug", false, "Enables debug mode")
	metadata := flag.Bool("metadata", false, "Write video metadata to a .json file")
	deadline := flag.Int("deadline", 1500, "Sets the timout for scraper logic in seconds (used as a workaround for 'context deadline exceeded' error)")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 && *batchFilePath == "" {
		fmt.Println("Usage: tiktok-dl [OPTIONS] TIKTOK_USERNAME|TIKTOK_URL")
		fmt.Println("  or:  tiktok-dl [OPTIONS] -batch-file path/to/users.txt")
		os.Exit(2)
	}

	if len(args) > 0 {
		Config.URL = flag.Args()[len(args)-1]
	} else {
		Config.URL = ""
	}
	Config.OutputPath = *outputPath
	Config.BatchFilePath = *batchFilePath
	Config.Debug = *debug
	Config.MetaData = *metadata
	Config.Deadline = *deadline
}

// GetUsername - Get's username from passed URL param
func GetUsername() string {
	return GetUsernameFromString(Config.URL)
}

// GetUsernameFromString - Get's username from passed param
func GetUsernameFromString(str string) string {
	if match := strings.Contains(str, "/"); !match { // Not url
		return strings.Replace(str, "@", "", -1)
	}

	if match, _ := regexp.MatchString(".+tiktok\\.com/@.+", str); match { // URL
		stripedSuffix := strings.Split(str, "@")[1]
		return strings.Split(stripedSuffix, "/")[0]
	}

	panic("Could not recognise URL format")
}
