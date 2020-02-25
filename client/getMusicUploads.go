package client

import (
	models "../models"
)

// GetMusicUploads - Get all uploads by given music
func GetMusicUploads(url string) ([]models.Upload, error) {
	actionOutput, err := executeClientAction(url, "bootstrapIteratingVideos()")
	if err != nil {
		return nil, err
	}
	return models.ParseUploads(actionOutput), nil
}
