package client

import (
	models "../models"
	config "../models/config"
	"fmt"
)

// GetUserUploads - Get all uploads marked with given hashtag
func GetHashtagUploads(hashtagURL string) []models.Upload {
	jsMethod := fmt.Sprintf("bootstrapIteratingVideos(%d)", config.Config.Limit)
	actionOutput := executeClientAction(hashtagURL, jsMethod)
	return models.ParseUploads(actionOutput)
}

func GetHashtagUploadsJson(hashtagURL string) string {
	jsMethod := fmt.Sprintf("bootstrapIteratingVideos(%d)", config.Config.Limit)
	return executeClientAction(hashtagURL, jsMethod)
}
