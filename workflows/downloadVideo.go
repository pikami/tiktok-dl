package workflows

import (
	client "../client"
	models "../models"
	utils "../utils"
	"fmt"
	"regexp"
)

// CanUseDownloadSingleVideo - Check's if DownloadSingleVideo can be used for parameter
func CanUseDownloadSingleVideo(url string) bool {
	match, _ := regexp.MatchString("\\/@.+\\/video\\/[0-9]+", url)
	return match
}

// DownloadSingleVideo - Downloads single video
func DownloadSingleVideo(url string) {
	username := models.GetUsername()
	upload := client.GetVideoDetails(url)
	downloadDir := fmt.Sprintf("%s/%s", models.Config.OutputPath, username)

	utils.InitOutputDirectory(downloadDir)
	downloadVideo(upload, downloadDir)
}

// DownloadVideo - Downloads one video
func downloadVideo(upload models.Upload, downloadDir string) {
	uploadID := upload.GetUploadID()
	downloadPath := fmt.Sprintf("%s/%s.mp4", downloadDir, uploadID)

	if utils.CheckIfExists(downloadPath) {
		fmt.Println("Upload '" + uploadID + "' already downloaded, skipping")
		return
	}

	fmt.Println("Downloading upload item '" + uploadID + "' to " + downloadPath)
	utils.DownloadFile(downloadPath, upload.URL)

	if models.Config.MetaData {
		metadataPath := fmt.Sprintf("%s/%s.json", downloadDir, uploadID)
		upload.WriteToFile(metadataPath)
	}
}
