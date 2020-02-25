package client

import (
	models "../models"
	config "../models/config"
	"fmt"
)

// GetUserUploads - Get all uploads by user
func GetUserUploads(username string) ([]models.Upload, error) {
	jsMethod := fmt.Sprintf("bootstrapIteratingVideos(%d)", config.Config.Limit)
	actionOutput, err := executeClientAction(`https://www.tiktok.com/@`+username, jsMethod)
	if err != nil {
		return nil, err
	}
	return models.ParseUploads(actionOutput), nil
}

func GetUserUploadsJson(username string) string {
	jsMethod := fmt.Sprintf("bootstrapIteratingVideos(%d)", config.Config.Limit)
	return executeClientAction(`https://www.tiktok.com/@`+username, jsMethod)
}
