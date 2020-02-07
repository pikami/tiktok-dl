package main

import (
	config "./models/config"
	workflows "./workflows"
)

func main() {
	config.GetConfig()
	url := config.Config.URL
	batchFilePath := config.Config.BatchFilePath

	// Batch file
	if workflows.CanUseDownloadBatchFile(batchFilePath) {
		workflows.DownloadBatchFile(batchFilePath)
		return
	}

	workflows.StartWorkflowByParameter(url)
}
