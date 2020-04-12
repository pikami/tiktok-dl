package client

import (
	models "github.com/pikami/tiktok-dl/models"
)

// GetVideoDetails - returns details of video
func GetVideoDetails(videoURL string) (models.Upload, error) {
	actionOutput, err := executeClientAction(videoURL, "bootstrapGetCurrentVideo()")
	if err != nil {
		return models.Upload{}, err
	}
	return models.ParseUpload(actionOutput), nil
}
