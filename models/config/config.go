package config

import (
	"flag"
	"fmt"
	"os"
)

// Config - Runtime configuration
var Config struct {
	URL             string
	OutputPath      string
	BatchFilePath   string
	ArchiveFilePath string
	FailLogFilePath string
	Debug           bool
	MetaData        bool
	Quiet           bool
	JSONOnly        bool
	Deadline        int
	Limit           int
}

// GetConfig - Returns Config object
func GetConfig() {
	outputPath := flag.String("output", "./downloads", "Output path")
	batchFilePath := flag.String("batch-file", "", "File containing URLs/Usernames to download, one value per line. Lines starting with '#', are considered as comments and ignored.")
	archive := flag.String("archive", "", "Download only videos not listed in the archive file. Record the IDs of all downloaded videos in it.")
	failLogPath := flag.String("fail-log", "", "Write failed items to log file")
	debug := flag.Bool("debug", false, "Enables debug mode")
	metadata := flag.Bool("metadata", false, "Write video metadata to a .json file")
	quiet := flag.Bool("quiet", false, "Suppress output")
	jsonOnly := flag.Bool("json", false, "Just get JSON data from scraper (without video downloading)")
	deadline := flag.Int("deadline", 1500, "Sets the timout for scraper logic in seconds (used as a workaround for 'context deadline exceeded' error)")
	limit := flag.Int("limit", 0, "Sets the videos count limit (useful when there too many videos from the user or by hashtag)")
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
	Config.ArchiveFilePath = *archive
	Config.FailLogFilePath = *failLogPath
	Config.Debug = *debug
	Config.MetaData = *metadata
	Config.Quiet = *quiet
	if *jsonOnly {
		Config.Quiet = true
	}
	Config.JSONOnly = *jsonOnly
	Config.Deadline = *deadline
	Config.Limit = *limit
}
