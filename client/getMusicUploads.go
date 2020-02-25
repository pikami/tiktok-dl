package client

import (
	models "../models"
	config "../models/config"
	"fmt"
)

// GetMusicUploads - Get all uploads by given music
func GetMusicUploads(url string) ([]models.Upload, error) {
  jsMethod := fmt.Sprintf("bootstrapIteratingVideos(%d)", config.Config.Limit)
	actionOutput, err := executeClientAction(url, jsMethod)
	if err != nil {
		return nil, err
	}
	return models.ParseUploads(actionOutput), nil
}

func GetMusicUploadsJson(url string) string {
	jsMethod := fmt.Sprintf("bootstrapIteratingVideos(%d)", config.Config.Limit)
	return executeClientAction(url, jsMethod)
}
