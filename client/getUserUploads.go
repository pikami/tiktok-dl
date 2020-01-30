package client

import (
	models "../models"
)

// GetUserUploads - Get all uploads by user
func GetUserUploads(username string) []models.Upload {
	actionOutput := executeClientAction(`https://www.tiktok.com/@`+username, "bootstrapIteratingVideos()")
	return models.ParseUploads(actionOutput)
}
