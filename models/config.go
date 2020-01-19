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
	Debug      bool
}

// GetConfig - Returns Config object
func GetConfig() {
	outputPath := flag.String("output", "./downloads", "Output path")
	debug := flag.Bool("debug", false, "enables debug mode")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: tiktok-dl [OPTIONS] TIKTOK_USERNAME")
		os.Exit(2)
	}

	Config.UserName = flag.Args()[len(args)-1]
	Config.OutputPath = *outputPath
	Config.Debug = *debug
}
