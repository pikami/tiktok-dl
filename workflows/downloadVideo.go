package workflows

import (
	"fmt"
	"regexp"

	client "github.com/pikami/tiktok-dl/client"
	models "github.com/pikami/tiktok-dl/models"
	config "github.com/pikami/tiktok-dl/models/config"
	res "github.com/pikami/tiktok-dl/resources"
	utils "github.com/pikami/tiktok-dl/utils"
	fileio "github.com/pikami/tiktok-dl/utils/fileio"
	log "github.com/pikami/tiktok-dl/utils/log"
)

// CanUseDownloadSingleVideo - Check's if DownloadSingleVideo can be used for parameter
func CanUseDownloadSingleVideo(url string) bool {
	match, _ := regexp.MatchString("\\/@.+\\/video\\/[0-9]+", url)
	return match
}

// DownloadSingleVideo - Downloads single video
func DownloadSingleVideo(url string) {
	username := utils.GetUsernameFromString(url)
	upload, err := client.GetVideoDetails(url)
	if err != nil {
		OnWorkflowFail(err, url)
		return
	}

	if utils.IsItemInArchive(upload) {
		return
	}
	downloadDir := fmt.Sprintf("%s/%s", config.Config.OutputPath, username)

	fileio.InitOutputDirectory(downloadDir)
	downloadVideo(upload, downloadDir)
	log.Logf(res.Downloaded, 1, 1)
}

// DownloadVideo - Downloads one video
func downloadVideo(upload models.Upload, downloadDir string) {
	uploadID := upload.GetUploadID()
	downloadPath := fmt.Sprintf("%s/%s.mp4", downloadDir, uploadID)

	if fileio.CheckIfExists(downloadPath) {
		return
	}

	utils.DownloadFile(downloadPath, upload.URL)

	if config.Config.MetaData {
		metadataPath := fmt.Sprintf("%s/%s.json", downloadDir, uploadID)
		upload.WriteToFile(metadataPath)
	}

	utils.AddItemToArchive(upload.GetUploadID())
}
