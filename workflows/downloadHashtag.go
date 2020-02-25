package workflows

import (
	client "../client"
	config "../models/config"
	res "../resources"
	utils "../utils"
	"fmt"
	"strings"
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
		utils.LogErr(res.ErrorCouldNotGetUserUploads, err.Error())
		return
	}
	uploadCount := len(uploads)
	hashtag := utils.GetHashtagFromURL(url)
	downloadDir := fmt.Sprintf("%s/%s", config.Config.OutputPath, hashtag)

	utils.InitOutputDirectory(downloadDir)

	for index, upload := range uploads {
		downloadVideo(upload, downloadDir)
		utils.Logf("\r[%d/%d] Downloaded", index+1, uploadCount)
	}
	utils.Log()
}

func GetHashtagJson(url string) {
	uploads, err := client.GetHashtagUploads(url)
	if err != nil {
		utils.LogErr(res.ErrorCouldNotGetUserUploads, err.Error())
		return
	}
	fmt.Printf("%s", uploads)
}
