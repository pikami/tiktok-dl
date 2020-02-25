package workflows

import (
	client "../client"
	config "../models/config"
	res "../resources"
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
	uploads, err := client.GetMusicUploads(url)
	if err != nil {
		utils.LogErr(res.ErrorCouldNotGetUserUploads, err.Error())
		return
	}
	uploadCount := len(uploads)

	for index, upload := range uploads {
		username := utils.GetUsernameFromString(upload.Uploader)
		downloadDir := fmt.Sprintf("%s/%s", config.Config.OutputPath, username)

		utils.InitOutputDirectory(downloadDir)
		downloadVideo(upload, downloadDir)
		utils.Logf("\r[%d/%d] Downloaded", index+1, uploadCount)
	}
	utils.Log()
}
