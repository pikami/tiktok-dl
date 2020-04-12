package client

import (
	"fmt"

	models "github.com/pikami/tiktok-dl/models"
	config "github.com/pikami/tiktok-dl/models/config"
)

// GetHashtagUploads - Get all uploads marked with given hashtag
func GetHashtagUploads(hashtagURL string) ([]models.Upload, error) {
	actionOutput, err := GetHashtagUploadsJSON(hashtagURL)
	if err != nil {
		return nil, err
	}

	return models.ParseUploads(actionOutput), nil
}

// GetHashtagUploadsJSON - Get hashtag uploads scrape
func GetHashtagUploadsJSON(hashtagURL string) (string, error) {
	jsMethod := fmt.Sprintf("bootstrapIteratingVideos(%d)", config.Config.Limit)
	actionOutput, err := executeClientAction(hashtagURL, jsMethod)
	if err != nil {
		return "", err
	}
	return actionOutput, nil
}
