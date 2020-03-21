package workflows

import (
	"fmt"
	"regexp"

	client "../client"
	config "../models/config"
	res "../resources"
	utils "../utils"
	fileio "../utils/fileio"
	log "../utils/log"
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
		log.LogErr(res.ErrorCouldNotGetUserUploads, err.Error())
		return
	}

	uploads = utils.RemoveArchivedItems(uploads)
	uploadCount := len(uploads)

	for index, upload := range uploads {
		username := utils.GetUsernameFromString(upload.Uploader)
		downloadDir := fmt.Sprintf("%s/%s", config.Config.OutputPath, username)

		fileio.InitOutputDirectory(downloadDir)
		downloadVideo(upload, downloadDir)
		log.Logf("\r[%d/%d] Downloaded", index+1, uploadCount)
	}
	log.Log()
}

// GetMusicJSON - Prints scraped info from music
func GetMusicJSON(url string) {
	uploads, err := client.GetMusicUploadsJSON(url)
	if err != nil {
		log.LogErr(res.ErrorCouldNotGetUserUploads, err.Error())
		return
	}
	fmt.Printf("%s", uploads)
}
