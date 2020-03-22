package workflows

import (
	config "../models/config"
	res "../resources"
	utils "../utils"
	log "../utils/log"
)

// StartWorkflowByParameter - Start needed workflow by given parameter
func StartWorkflowByParameter(url string) {

	// Music
	if CanUseDownloadMusic(url) {
		if config.Config.JSONOnly {
			GetMusicJSON(url)
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
			GetUserVideosJSON(utils.GetUsernameFromString(url))
		} else {
			DownloadUser(utils.GetUsernameFromString(url))
		}

		return
	}

	// Tiktok hashtag
	if CanUseDownloadHashtag(url) {
		if config.Config.JSONOnly {
			GetHashtagJSON(url)
		} else {
			DownloadHashtag(url)
		}
		return
	}

	// Share URL
	if CanUseDownloadShareLink(url) {
		DownloadShareLink(url)
		return
	}

	log.LogFatal(res.ErrorCouldNotRecogniseURL, url)
}
