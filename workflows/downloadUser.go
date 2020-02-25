package workflows

import (
	client "../client"
	config "../models/config"
	res "../resources"
	utils "../utils"
	"fmt"
	"regexp"
	"strings"
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
		utils.LogErr(res.ErrorCouldNotGetUserUploads, err.Error())
		return
	}
	uploadCount := len(uploads)
	downloadDir := fmt.Sprintf("%s/%s", config.Config.OutputPath, username)

	utils.InitOutputDirectory(downloadDir)

	for index, upload := range uploads {
		downloadVideo(upload, downloadDir)
		utils.Logf("\r[%d/%d] Downloaded", index+1, uploadCount)
	}
	utils.Log()
}

func GetUserVideosJson(username string) {
	uploads := client.GetUserUploadsJson(username)
	fmt.Printf("%s", uploads)
}
