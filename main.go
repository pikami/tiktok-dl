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

	workflows.StartWorkflowByParameter(url)
}
