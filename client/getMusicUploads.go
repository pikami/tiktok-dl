package client

import (
	models "../models"
	config "../models/config"
	"fmt"
)

// GetMusicUploads - Get all uploads by given music
func GetMusicUploads(url string) []models.Upload {
	jsMethod := fmt.Sprintf("bootstrapIteratingVideos(%d)", config.Config.Limit)
	actionOutput := executeClientAction(url, jsMethod)
	return models.ParseUploads(actionOutput)
}

func GetMusicUploadsJson(url string) string {
	jsMethod := fmt.Sprintf("bootstrapIteratingVideos(%d)", config.Config.Limit)
	return executeClientAction(url, jsMethod)
}
