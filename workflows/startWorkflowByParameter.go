package workflows

import (
	config "../models/config"
	res "../resources"
	utils "../utils"
)

// StartWorkflowByParameter - Start needed workflow by given parameter
func StartWorkflowByParameter(url string) {

	// Music
	if CanUseDownloadMusic(url) {
		if config.Config.JSONOnly {
			GetMusicJson(url)
		} else {
			DownloadMusic(url)
		}
		return
	}

	// Single video
	if CanUseDownloadSingleVideo(url) {
		DownloadSingleVideo(url)
		return
	}

	// Tiktok user
	if CanUseDownloadUser(url) {
		if config.Config.JSONOnly {
			GetUserVideosJson(utils.GetUsernameFromString(url))
		} else {
			DownloadUser(utils.GetUsernameFromString(url))
		}

		return
	}

	// Tiktok hashtag
	if CanUseDownloadHashtag(url) {
		if config.Config.JSONOnly {
			GetHashtagJson(url)
		} else {
			DownloadHashtag(url)
		}
		return
	}

	utils.LogFatal(res.ErrorCouldNotRecogniseURL, url)
}
