package client

import (
	models "../models"
	config "../models/config"
	"fmt"
)

// GetUserUploads - Get all uploads marked with given hashtag
func GetHashtagUploads(hashtagURL string) ([]models.Upload, error) {
	jsMethod := fmt.Sprintf("bootstrapIteratingVideos(%d)", config.Config.Limit)
	actionOutput, err := executeClientAction(hashtagURL, jsMethod)
	if err != nil {
		return nil, err
	}
	return models.ParseUploads(actionOutput), nil
}

func GetHashtagUploadsJson(hashtagURL string) (string, error) {
	jsMethod := fmt.Sprintf("bootstrapIteratingVideos(%d)", config.Config.Limit)
	actionOutput, err := executeClientAction(hashtagURL, jsMethod)
	if err != nil {
		return "", err
	}
	return actionOutput, nil
}
