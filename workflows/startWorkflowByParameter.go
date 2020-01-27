package workflows

import (
	models "../models"
	"fmt"
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
		DownloadUser(models.GetUsernameFromString(url))
		return
	}

	panic(fmt.Sprintf("Could not recognise URL format of string %s", url))
}
