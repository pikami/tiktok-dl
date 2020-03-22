package workflows

import (
	"fmt"
	"regexp"

	client "../client"
	models "../models"
	config "../models/config"
	res "../resources"
	utils "../utils"
	log "../utils/log"
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
		log.LogErr(res.ErrorCouldNotGetUserUploads, err.Error())
		return
	}

	if utils.IsItemInArchive(upload) {
		return
	}
	downloadDir := fmt.Sprintf("%s/%s", config.Config.OutputPath, username)

	utils.InitOutputDirectory(downloadDir)
	downloadVideo(upload, downloadDir)
	log.Log("[1/1] Downloaded\n")
}

// DownloadVideo - Downloads one video
func downloadVideo(upload models.Upload, downloadDir string) {
	uploadID := upload.GetUploadID()
	downloadPath := fmt.Sprintf("%s/%s.mp4", downloadDir, uploadID)

	if utils.CheckIfExists(downloadPath) {
		return
	}

	utils.DownloadFile(downloadPath, upload.URL)

	if config.Config.MetaData {
		metadataPath := fmt.Sprintf("%s/%s.json", downloadDir, uploadID)
		upload.WriteToFile(metadataPath)
	}

	utils.AddItemToArchive(upload.GetUploadID())
}
