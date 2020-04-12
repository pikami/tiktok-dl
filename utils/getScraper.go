package utils

import (
	resources "github.com/pikami/tiktok-dl/resources"
	fileio "github.com/pikami/tiktok-dl/utils/fileio"
)

// GetScraper - Retrieve scraper
func GetScraper() string {
	if fileio.CheckIfExists(resources.ScraperPath) {
		return ReadFileAsString(resources.ScraperPath)
	}

	if resources.ScraperScript != "" {
		return resources.ScraperScript
	}

	panic(resources.FailedToLoadScraper)
}
