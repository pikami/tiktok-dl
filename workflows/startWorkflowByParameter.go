package workflows

import (
	config "github.com/pikami/tiktok-dl/models/config"
	res "github.com/pikami/tiktok-dl/resources"
	utils "github.com/pikami/tiktok-dl/utils"
	log "github.com/pikami/tiktok-dl/utils/log"
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
