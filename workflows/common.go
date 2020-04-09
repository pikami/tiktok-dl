package workflows

import (
	config "../models/config"
	res "../resources"
	fileio "../utils/fileio"
	log "../utils/log"
)

// OnWorkflowFail - Funtion called when workflow fails
func OnWorkflowFail(err error, workItem string) {
	failLogFilePath := config.Config.FailLogFilePath

	if failLogFilePath != "" {
		fileio.AppendToFile(workItem, failLogFilePath)
	}

	log.LogErr(res.Error, err.Error())
	log.LogErr(res.FailedOnItem, workItem)
}
