package workflows

import (
	utils "../utils"
	"fmt"
)

// CanUseDownloadBatchFile - Check's if DownloadBatchFile can be used
func CanUseDownloadBatchFile(batchFilePath string) bool {
	return batchFilePath != ""
}

// DownloadBatchFile - Download items from batch file
func DownloadBatchFile(batchFilePath string) {
	if !utils.CheckIfExists(batchFilePath) {
		panic(fmt.Sprintf("File path %s not found.", batchFilePath))
	}

	utils.ReadFileLineByLine(batchFilePath, downloadItem)
}

func downloadItem(batchItem string) {
	if batchItem[0] == '#' {
		return
	}

	StartWorkflowByParameter(batchItem)
}
