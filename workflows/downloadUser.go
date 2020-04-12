package workflows

import (
	"fmt"
	"regexp"
	"strings"

	client "github.com/pikami/tiktok-dl/client"
	config "github.com/pikami/tiktok-dl/models/config"
	res "github.com/pikami/tiktok-dl/resources"
	utils "github.com/pikami/tiktok-dl/utils"
	fileio "github.com/pikami/tiktok-dl/utils/fileio"
	log "github.com/pikami/tiktok-dl/utils/log"
)

// CanUseDownloadUser - Test's if this workflow can be used for parameter
func CanUseDownloadUser(url string) bool {
	isURL := strings.Contains(url, "/")
	match, _ := regexp.MatchString(".+com\\/@[^\\/]+", url)
	return !isURL || match
}

// DownloadUser - Download all user's videos
func DownloadUser(username string) {
	uploads, err := client.GetUserUploads(username)
	if err != nil {
		OnWorkflowFail(err, username)
		return
	}

	uploads = utils.RemoveArchivedItems(uploads)
	uploadCount := len(uploads)

	downloadDir := fmt.Sprintf("%s/%s", config.Config.OutputPath, username)

	fileio.InitOutputDirectory(downloadDir)

	for index, upload := range uploads {
		downloadVideo(upload, downloadDir)
		log.Logf(res.Downloaded, index+1, uploadCount)
	}
	log.Log()
}

// GetUserVideosJSON - Prints scraped info from user
func GetUserVideosJSON(username string) {
	uploads, err := client.GetUserUploadsJSON(username)
	if err != nil {
		OnWorkflowFail(err, username)
		return
	}
	fmt.Printf("%s", uploads)
}
