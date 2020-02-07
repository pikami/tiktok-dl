package workflows

import (
	res "../resources"
	utils "../utils"
)

// StartWorkflowByParameter - Start needed workflow by given parameter
func StartWorkflowByParameter(url string) {

	// Music
	if CanUseDownloadMusic(url) {
		DownloadMusic(url)
		return
	}

	// Single video
	if CanUseDownloadSingleVideo(url) {
		DownloadSingleVideo(url)
		return
	}

	// Tiktok user
	if CanUseDownloadUser(url) {
		DownloadUser(utils.GetUsernameFromString(url))
		return
	}

	utils.LogFatal(res.ErrorCouldNotRecogniseURL, url)
}
