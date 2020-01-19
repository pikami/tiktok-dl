package main

import (
	client "./client"
	models "./models"
	utils "./utils"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	models.GetConfig()
	url := models.Config.URL

	// Single video
	match, _ := regexp.MatchString("\\/@.+\\/video\\/[0-9]+", url)
	if match {
		getUsernameFromVidURLRegex, _ := regexp.Compile("com\\/@.*")
		parts := strings.Split(getUsernameFromVidURLRegex.FindString(url), "/")
		username := parts[1][1:]
		upload := client.GetVideoDetails(url)
		downloadDir := fmt.Sprintf("%s/%s", models.Config.OutputPath, username)

		utils.InitOutputDirectory(downloadDir)
		downloadVideo(upload, downloadDir)
		return
	}

	// Tiktok user
	downloadUser()
}

func downloadVideo(upload models.Upload, downloadDir string) {
	uploadID := upload.GetUploadID()
	downloadPath := fmt.Sprintf("%s/%s.mp4", downloadDir, uploadID)

	if utils.CheckIfExists(downloadPath) {
		fmt.Println("Upload '" + uploadID + "' already downloaded, skipping")
		return
	}

	fmt.Println("Downloading upload item '" + uploadID + "' to " + downloadPath)
	utils.DownloadFile(downloadPath, upload.URL)
}

func downloadUser() {
	username := models.Config.URL
	downloadDir := fmt.Sprintf("%s/%s", models.Config.OutputPath, username)
	uploads := client.GetUserUploads(username)

	utils.InitOutputDirectory(downloadDir)

	for _, upload := range uploads {
		downloadVideo(upload, downloadDir)
	}
}
