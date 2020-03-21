package workflows

import (
	res "../resources"
	fileio "../utils/fileio"
	log "../utils/log"
)

// CanUseDownloadBatchFile - Check's if DownloadBatchFile can be used
func CanUseDownloadBatchFile(batchFilePath string) bool {
	return batchFilePath != ""
}

// DownloadBatchFile - Download items from batch file
func DownloadBatchFile(batchFilePath string) {
	if !fileio.CheckIfExists(batchFilePath) {
		log.LogFatal(res.ErrorPathNotFound, batchFilePath)
	}

	fileio.ReadFileLineByLine(batchFilePath, downloadItem)
}

func downloadItem(batchItem string) {
	if batchItem[0] == '#' {
		return
	}

	StartWorkflowByParameter(batchItem)
}
