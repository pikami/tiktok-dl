package client

import (
	"fmt"

	models "github.com/pikami/tiktok-dl/models"
	config "github.com/pikami/tiktok-dl/models/config"
)

// GetUserUploads - Get all uploads by user
func GetUserUploads(username string) ([]models.Upload, error) {
	actionOutput, err := GetUserUploadsJSON(username)
	if err != nil {
		return nil, err
	}
	return models.ParseUploads(actionOutput), nil
}

// GetUserUploadsJSON - Get user uploads scrape
func GetUserUploadsJSON(username string) (string, error) {
	jsMethod := fmt.Sprintf("bootstrapIteratingVideos(%d)", config.Config.Limit)
	actionOutput, err := executeClientAction(`https://www.tiktok.com/@`+username, jsMethod)
	if err != nil {
		return "", err
	}
	return actionOutput, nil
}
