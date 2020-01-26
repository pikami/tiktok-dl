package workflows

import (
	client "../client"
	models "../models"
	utils "../utils"
	"fmt"
	"regexp"
)

// CanUseDownloadMusic - Check's if DownloadMusic can be used for parameter
func CanUseDownloadMusic(url string) bool {
	match, _ := regexp.MatchString(".com\\/music\\/.+", url)
	return match
}

// DownloadMusic - Download all videos by given music
func DownloadMusic(url string) {
	uploads := client.GetMusicUploads(url)

	for _, upload := range uploads {
		username := models.GetUsernameFromString(upload.Uploader)
		downloadDir := fmt.Sprintf("%s/%s", models.Config.OutputPath, username)

		utils.InitOutputDirectory(downloadDir)
		downloadVideo(upload, downloadDir)
	}
}
