package workflows

import (
	"fmt"
	"regexp"

	client "github.com/pikami/tiktok-dl/client"
	config "github.com/pikami/tiktok-dl/models/config"
	res "github.com/pikami/tiktok-dl/resources"
	utils "github.com/pikami/tiktok-dl/utils"
	fileio "github.com/pikami/tiktok-dl/utils/fileio"
	log "github.com/pikami/tiktok-dl/utils/log"
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
		OnWorkflowFail(err, url)
		return
	}

	uploads = utils.RemoveArchivedItems(uploads)
	uploadCount := len(uploads)

	for index, upload := range uploads {
		username := utils.GetUsernameFromString(upload.Uploader)
		downloadDir := fmt.Sprintf("%s/%s", config.Config.OutputPath, username)

		fileio.InitOutputDirectory(downloadDir)
		downloadVideo(upload, downloadDir)
		log.Logf(res.Downloaded, index+1, uploadCount)
	}
	log.Log()
}

// GetMusicJSON - Prints scraped info from music
func GetMusicJSON(url string) {
	uploads, err := client.GetMusicUploadsJSON(url)
	if err != nil {
		OnWorkflowFail(err, url)
		return
	}
	fmt.Printf("%s", uploads)
}
