package main

import (
	config "github.com/pikami/tiktok-dl/models/config"
	workflows "github.com/pikami/tiktok-dl/workflows"
)

func main() {
	config.GetConfig()
	url := config.Config.URL
	batchFilePath := config.Config.BatchFilePath
	scrapedDataFilePath := config.Config.ScrapedDataFilePath

	// Batch file
	if workflows.CanUseDownloadBatchFile(batchFilePath) {
		workflows.DownloadBatchFile(batchFilePath)
		return
	}

	// Scraped data file
	if workflows.CanUseDownloadScrapedData(scrapedDataFilePath) {
		workflows.DownloadScrapedData(scrapedDataFilePath)
		return
	}

	workflows.StartWorkflowByParameter(url)
}
