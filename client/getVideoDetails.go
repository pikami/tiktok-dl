package client

import (
	models "../models"
)

// GetVideoDetails - returns details of video
func GetVideoDetails(videoURL string) models.Upload {
	actionOutput := executeClientAction(videoURL, "bootstrapGetCurrentVideo()")
	return models.ParseUpload(actionOutput)
}
