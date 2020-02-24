package client

import (
	models "../models"
	config "../models/config"
    "fmt"
)

// GetUserUploads - Get all uploads by user
func GetUserUploads(username string) []models.Upload {
    jsMethod := fmt.Sprintf("bootstrapIteratingVideos(%d)", config.Config.Limit)
	actionOutput := executeClientAction(`https://www.tiktok.com/@`+username, jsMethod)
	return models.ParseUploads(actionOutput)
}

func GetUserUploadsJson(username string) string {
    jsMethod := fmt.Sprintf("bootstrapIteratingVideos(%d)", config.Config.Limit)
	return executeClientAction(`https://www.tiktok.com/@`+username, jsMethod)
}