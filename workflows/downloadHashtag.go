package workflows

import (
	"fmt"
	"strings"

	client "../client"
	config "../models/config"
	res "../resources"
	utils "../utils"
	fileio "../utils/fileio"
	log "../utils/log"
)

// CanUseDownloadHashtag - Test's if this workflow can be used for parameter
func CanUseDownloadHashtag(url string) bool {
	match := strings.Contains(url, "/tag/")
	return match
}

// DownloadHashtag - Download videos marked with given hashtag
func DownloadHashtag(url string) {
	uploads, err := client.GetHashtagUploads(url)
	if err != nil {
		log.LogErr(res.ErrorCouldNotGetUserUploads, err.Error())
		return
	}

	uploads = utils.RemoveArchivedItems(uploads)
	uploadCount := len(uploads)

	hashtag := utils.GetHashtagFromURL(url)
	downloadDir := fmt.Sprintf("%s/%s", config.Config.OutputPath, hashtag)

	fileio.InitOutputDirectory(downloadDir)

	for index, upload := range uploads {
		downloadVideo(upload, downloadDir)
		log.Logf("\r[%d/%d] Downloaded", index+1, uploadCount)
	}
	log.Log()
}

// GetHashtagJSON - Prints scraped info from hashtag
func GetHashtagJSON(url string) {
	uploads, err := client.GetHashtagUploadsJSON(url)
	if err != nil {
		log.LogErr(res.ErrorCouldNotGetUserUploads, err.Error())
		return
	}
	fmt.Printf("%s", uploads)
}
