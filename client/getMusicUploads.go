package client

import (
	"fmt"

	models "github.com/pikami/tiktok-dl/models"
	config "github.com/pikami/tiktok-dl/models/config"
)

// GetMusicUploads - Get all uploads by given music
func GetMusicUploads(url string) ([]models.Upload, error) {
	actionOutput, err := GetMusicUploadsJSON(url)
	if err != nil {
		return nil, err
	}
	return models.ParseUploads(actionOutput), nil
}

// GetMusicUploadsJSON - Get music uploads scrape
func GetMusicUploadsJSON(url string) (string, error) {
	jsMethod := fmt.Sprintf("bootstrapIteratingVideos(%d)", config.Config.Limit)
	actionOutput, err := executeClientAction(url, jsMethod)
	if err != nil {
		return "", err
	}
	return actionOutput, nil
}
