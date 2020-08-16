package main

var (
	outputDir = "../resources"

	res = []resource{
		resource{
			Package:  "resources",
			FileName: "scraper.go",
			Values: map[string]string{
				"ScraperPath":   "scraper.js",
				"ScraperScript": fileContentsOrDefault("../scraper.min.js"),
			},
		},
		resource{
			Package:  "resources",
			FileName: "errorStrings.go",
			Values: map[string]string{
				"ErrorCouldNotSerializeJSON": "Could not serialize json for video: %s\n",
				"ErrorCouldNotRecogniseURL":  "Could not recognise URL format of string %s",
				"Error":                      "Error : %s\n",
				"ErrorPathNotFound":          "File path %s not found.",
				"FailedOnItem":               "Failed while scraping item: %s\n",
				"FailedToLoadScraper":        "Failed to load scraper",
			},
		},
		resource{
			Package:  "resources",
			FileName: "messages.go",
			Values: map[string]string{
				"PreloadingItemsFound": "\rPreloading... %s items have been found.",
				"Preloading":           "\rPreloading...",
				"Retrieving":           "\nRetrieving items...",
				"ItemsFoundInArchive":  "%d items, found in archive. Skipping...\n",
				"Downloaded":           "\r[%d/%d] Downloaded",
				"UsageLine": "Usage: tiktok-dl [OPTIONS] TIKTOK_USERNAME|TIKTOK_URL\n" +
					"  or:  tiktok-dl [OPTIONS] -batch-file path/to/users.txt\n" +
					"  or:  tiktok-dl [OPTIONS] -scraped-data path/to/data.json",
			},
		},
		resource{
			Package:  "resources",
			FileName: "flags.go",
			Values: map[string]string{
				// Output
				"OutputFlag":        "output",
				"OutputDefault":     "./downloads",
				"OutputDescription": "Output path",
				// Batch file
				"BatchFlag":        "batch-file",
				"BatchDefault":     "",
				"BatchDescription": "File containing URLs/Usernames to download, one value per line. Lines starting with '#', are considered as comments and ignored.",
				// ScrapedData
				"ScrapedDataFlag":        "scraped-data",
				"ScrapedDataDefault":     "",
				"ScrapedDataDescription": "Download videos from scrape file (json format)",
				// Archive
				"ArchiveFlag":        "archive",
				"ArchiveDefault":     "",
				"ArchiveDescription": "Download only videos not listed in the archive file. Record the IDs of all downloaded videos in it.",
				// Fail log
				"FailLogFlag":        "fail-log",
				"FailLogDefault":     "",
				"FailLogDescription": "Write failed items to log file",
				// Debug
				"DebugFlag":        "debug",
				"DebugDefault":     "false",
				"DebugDescription": "Enables debug mode",
				// Metadata
				"MetadataFlag":        "metadata",
				"MetadataDefault":     "false",
				"MetadataDescription": "Write video metadata to a .json file",
				// Quiet
				"QuietFlag":        "quiet",
				"QuietDefault":     "false",
				"QuietDescription": "Suppress output",
				// JSON only
				"JsonFlag":        "json",
				"JsonDefault":     "false",
				"JsonDescription": "Just get JSON data from scraper (without video downloading)",
				// Deadline
				"DeadlineFlag":        "deadline",
				"DeadlineDefault":     "1500",
				"DeadlineDescription": "Sets the timout for scraper logic in seconds (used as a workaround for 'context deadline exceeded' error)",
				// Limit
				"LimitFlag":        "limit",
				"LimitDefault":     "0",
				"LimitDescription": "Sets the videos count limit (useful when there too many videos from the user or by hashtag)",
			},
		},
	}
)
