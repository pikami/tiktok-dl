package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	res "github.com/pikami/tiktok-dl/resources"
)

// Config - Runtime configuration
var Config struct {
	URL                 string
	OutputPath          string
	BatchFilePath       string
	ScrapedDataFilePath string
	ArchiveFilePath     string
	FailLogFilePath     string
	Debug               bool
	MetaData            bool
	Quiet               bool
	JSONOnly            bool
	Deadline            int
	Limit               int
}

// GetConfig - Returns Config object
func GetConfig() {
	outputPath := flag.String(res.OutputFlag, res.OutputDefault, res.OutputDescription)
	batchFilePath := flag.String(res.BatchFlag, res.BatchDefault, res.BatchDescription)
	scrapedDataFilePath := flag.String(res.ScrapedDataFlag, res.ScrapedDataDefault, res.ScrapedDataDescription)
	archive := flag.String(res.ArchiveFlag, res.ArchiveDefault, res.ArchiveDescription)
	failLogPath := flag.String(res.FailLogFlag, res.FailLogDefault, res.FailLogDescription)
	debug := flag.Bool(res.DebugFlag, parseBool(res.DebugDefault), res.DebugDescription)
	metadata := flag.Bool(res.MetadataFlag, parseBool(res.MetadataDefault), res.MetadataDescription)
	quiet := flag.Bool(res.QuietFlag, parseBool(res.QuietDefault), res.QuietDescription)
	jsonOnly := flag.Bool(res.JsonFlag, parseBool(res.JsonDefault), res.JsonDescription)
	deadline := flag.Int(res.DeadlineFlag, parseInt(res.DeadlineDefault), res.DeadlineDescription)
	limit := flag.Int(res.LimitFlag, parseInt(res.LimitDefault), res.LimitDescription)
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 && *batchFilePath == "" && *scrapedDataFilePath == "" {
		fmt.Println(res.UsageLine)
		os.Exit(2)
	}

	if len(args) > 0 {
		Config.URL = flag.Args()[len(args)-1]
	} else {
		Config.URL = ""
	}
	Config.OutputPath = *outputPath
	Config.BatchFilePath = *batchFilePath
	Config.ScrapedDataFilePath = *scrapedDataFilePath
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

func parseBool(str string) bool {
	val, err := strconv.ParseBool(str)
	if err != nil {
		panic(err)
	}
	return val
}

func parseInt(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return val
}
