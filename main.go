package main

import (
	models "./models"
	workflows "./workflows"
)

func main() {
	models.GetConfig()
	url := models.Config.URL
	batchFilePath := models.Config.BatchFilePath

	// Batch file
	if workflows.CanUseDownloadBatchFile(batchFilePath) {
		workflows.DownloadBatchFile(batchFilePath)
		return
	}

	// Single video
	if workflows.CanUseDownloadSingleVideo(url) {
		workflows.DownloadSingleVideo(url)
		return
	}

	// Tiktok user
	if workflows.CanUseDownloadUser(url) {
		workflows.DownloadUser(models.GetUsername())
		return
	}

	panic("Could not recognise URL format")
}
