package client

import (
	models "../models"
)

// GetUserUploads - Get all uploads by user
func GetUserUploads(username string) ([]models.Upload, error) {
	actionOutput, err := executeClientAction(`https://www.tiktok.com/@`+username, "bootstrapIteratingVideos()")
	if err != nil {
		return nil, err
	}
	return models.ParseUploads(actionOutput), nil
}
