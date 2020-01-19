package main

import (
	client "./client"
	models "./models"
	utils "./utils"
	"fmt"
)

func main() {
	models.GetConfig()

	username := models.Config.UserName
	downloadDir := fmt.Sprintf("%s/%s", models.Config.OutputPath, username)
	uploads := client.GetUserUploads(username)

	utils.InitOutputDirectory(downloadDir)

	for _, upload := range uploads {
		uploadID := upload.GetUploadID()
		downloadPath := fmt.Sprintf("%s/%s.mp4", downloadDir, uploadID)

		if utils.CheckIfExists(downloadPath) {
			fmt.Println("Upload '" + uploadID + "' already downloaded, skipping")
			continue
		}

		utils.DownloadFile(downloadPath, upload.URL)
	}
}
