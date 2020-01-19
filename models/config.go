package models

import (
	"flag"
	"fmt"
	"os"
)

// Config - Runtime configuration
var Config struct {
	UserName   string
	OutputPath string
}

// GetConfig - Returns Config object
func GetConfig() {
	outputPath := flag.String("output", "./downloads", "Output path")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: tiktok-dl [OPTIONS] TIKTOK_USERNAME")
		os.Exit(2)
	}

	Config.UserName = flag.Args()[len(args)-1]
	Config.OutputPath = *outputPath
}
