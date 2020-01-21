package workflows

import (
	client "../client"
	models "../models"
	utils "../utils"
	"fmt"
	"strings"
)

// CanUseDownloadUser - Test's if this workflow can be used for parameter
func CanUseDownloadUser(url string) bool {
	match := strings.Contains(url, "/")
	return !match
}

// DownloadUser - Download all user's videos
func DownloadUser(username string) {
	downloadDir := fmt.Sprintf("%s/%s", models.Config.OutputPath, username)
	uploads := client.GetUserUploads(username)

	utils.InitOutputDirectory(downloadDir)

	for _, upload := range uploads {
		downloadVideo(upload, downloadDir)
	}
}
