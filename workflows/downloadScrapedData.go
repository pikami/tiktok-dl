package workflows

import (
	"fmt"
	models "github.com/pikami/tiktok-dl/models"
	config "github.com/pikami/tiktok-dl/models/config"
	res "github.com/pikami/tiktok-dl/resources"
	utils "github.com/pikami/tiktok-dl/utils"
	fileio "github.com/pikami/tiktok-dl/utils/fileio"
	log "github.com/pikami/tiktok-dl/utils/log"
)

// CanUseDownloadScrapedData - Check's if DownloadScrapedData can be used
func CanUseDownloadScrapedData(scrapedDataFilePath string) bool {
	return scrapedDataFilePath != ""
}

// DownloadScrapedData - Download items from scraped data file
func DownloadScrapedData(scrapedDataFilePath string) {
	if !fileio.CheckIfExists(scrapedDataFilePath) {
		log.LogFatal(res.ErrorPathNotFound, scrapedDataFilePath)
	}

	dataFileContent := fileio.ReadFileToString(scrapedDataFilePath)
	uploads := models.ParseUploads(dataFileContent)
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
