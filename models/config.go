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
	URL        string
	OutputPath string
	Debug      bool
	MetaData   bool
}

// GetConfig - Returns Config object
func GetConfig() {
	outputPath := flag.String("output", "./downloads", "Output path")
	debug := flag.Bool("debug", false, "Enables debug mode")
	metadata := flag.Bool("metadata", false, "Write video metadata to a .json file")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: tiktok-dl [OPTIONS] TIKTOK_USERNAME|TIKTOK_URL")
		os.Exit(2)
	}

	Config.URL = flag.Args()[len(args)-1]
	Config.OutputPath = *outputPath
	Config.Debug = *debug
	Config.MetaData = *metadata
}

// GetUsername - Get's username from passed URL param
func GetUsername() string {
	if match := strings.Contains(Config.URL, "/"); !match { // Not url
		return strings.Replace(Config.URL, "@", "", -1)
	}

	if match, _ := regexp.MatchString(".+tiktok\\.com/@.+", Config.URL); match { // URL
		stripedSuffix := strings.Split(Config.URL, "@")[1]
		return strings.Split(stripedSuffix, "/")[0]
	}

	panic("Could not recognise URL format")
}
