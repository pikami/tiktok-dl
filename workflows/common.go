package workflows

import (
	config "github.com/pikami/tiktok-dl/models/config"
	res "github.com/pikami/tiktok-dl/resources"
	fileio "github.com/pikami/tiktok-dl/utils/fileio"
	log "github.com/pikami/tiktok-dl/utils/log"
)

// OnWorkflowFail - Function called when workflow fails
func OnWorkflowFail(err error, workItem string) {
	failLogFilePath := config.Config.FailLogFilePath

	if failLogFilePath != "" {
		fileio.AppendToFile(workItem, failLogFilePath)
	}

	log.LogErr(res.Error, err.Error())
	log.LogErr(res.FailedOnItem, workItem)
}
