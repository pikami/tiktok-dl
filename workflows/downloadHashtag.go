package workflows

import (
	"fmt"
	"strings"

	client "github.com/pikami/tiktok-dl/client"
	config "github.com/pikami/tiktok-dl/models/config"
	res "github.com/pikami/tiktok-dl/resources"
	utils "github.com/pikami/tiktok-dl/utils"
	fileio "github.com/pikami/tiktok-dl/utils/fileio"
	log "github.com/pikami/tiktok-dl/utils/log"
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
		OnWorkflowFail(err, url)
		return
	}

	uploads = utils.RemoveArchivedItems(uploads)
	uploadCount := len(uploads)

	hashtag := utils.GetHashtagFromURL(url)
	downloadDir := fmt.Sprintf("%s/%s", config.Config.OutputPath, hashtag)

	fileio.InitOutputDirectory(downloadDir)

	for index, upload := range uploads {
		downloadVideo(upload, downloadDir)
		log.Logf(res.Downloaded, index+1, uploadCount)
	}
	log.Log()
}

// GetHashtagJSON - Prints scraped info from hashtag
func GetHashtagJSON(url string) {
	uploads, err := client.GetHashtagUploadsJSON(url)
	if err != nil {
		OnWorkflowFail(err, url)
		return
	}
	fmt.Printf("%s", uploads)
}
