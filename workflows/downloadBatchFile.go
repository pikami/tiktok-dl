package workflows

import (
	res "../resources"
	utils "../utils"
)

// CanUseDownloadBatchFile - Check's if DownloadBatchFile can be used
func CanUseDownloadBatchFile(batchFilePath string) bool {
	return batchFilePath != ""
}

// DownloadBatchFile - Download items from batch file
func DownloadBatchFile(batchFilePath string) {
	if !utils.CheckIfExists(batchFilePath) {
		utils.LogFatal(res.ErrorPathNotFound, batchFilePath)
	}

	utils.ReadFileLineByLine(batchFilePath, downloadItem)
}

func downloadItem(batchItem string) {
	if batchItem[0] == '#' {
		return
	}

	StartWorkflowByParameter(batchItem)
}
