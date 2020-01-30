package client

import (
	models "../models"
)

// GetMusicUploads - Get all uploads by given music
func GetMusicUploads(url string) []models.Upload {
	actionOutput := executeClientAction(url, "bootstrapIteratingVideos()")
	return models.ParseUploads(actionOutput)
}
