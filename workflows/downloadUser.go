package workflows

import (
	"fmt"
	"regexp"
	"strings"

	client "../client"
	config "../models/config"
	res "../resources"
	utils "../utils"
	log "../utils/log"
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
		log.LogErr(res.ErrorCouldNotGetUserUploads, err.Error())
		return
	}

	uploads = utils.RemoveArchivedItems(uploads)
	uploadCount := len(uploads)

	downloadDir := fmt.Sprintf("%s/%s", config.Config.OutputPath, username)

	utils.InitOutputDirectory(downloadDir)

	for index, upload := range uploads {
		downloadVideo(upload, downloadDir)
		log.Logf("\r[%d/%d] Downloaded", index+1, uploadCount)
	}
	log.Log()
}

// GetUserVideosJSON - Prints scraped info from user
func GetUserVideosJSON(username string) {
	uploads, err := client.GetUserUploadsJSON(username)
	if err != nil {
		log.LogErr(res.ErrorCouldNotGetUserUploads, err.Error())
		return
	}
	fmt.Printf("%s", uploads)
}
