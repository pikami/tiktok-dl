package models

import (
	"encoding/json"
	"os"
	"strings"

	res "github.com/pikami/tiktok-dl/resources"
	checkErr "github.com/pikami/tiktok-dl/utils/checkErr"
	log "github.com/pikami/tiktok-dl/utils/log"
)

// Upload - Upload object
type Upload struct {
	URL       string `json:"url"`
	ShareLink string `json:"shareLink"`
	Caption   string `json:"caption"`
	Uploader  string `json:"uploader"`
	Sound     Sound  `json:"sound"`
}

// Sound - Sound object
type Sound struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

// ParseUploads - Parses json uploads array
func ParseUploads(str string) []Upload {
	var uploads []Upload
	json.Unmarshal([]byte(str), &uploads)
	return uploads
}

// ParseUpload - Parses json uploads array
func ParseUpload(str string) Upload {
	var upload Upload
	json.Unmarshal([]byte(str), &upload)
	return upload
}

// GetUploadID - Returns upload id
func (u Upload) GetUploadID() string {
	parts := strings.Split(u.ShareLink, "/")
	return parts[len(parts)-1]
}

// WriteToFile - Writes object to file
func (u Upload) WriteToFile(outputPath string) {
	bytes, err := json.Marshal(u)
	if err != nil {
		log.Logf(res.ErrorCouldNotSerializeJSON, u.GetUploadID())
		panic(err)
	}

	// Create the file
	out, err := os.Create(outputPath)
	checkErr.CheckErr(err)
	defer out.Close()

	// Write to file
	_, err = out.Write(bytes)
	checkErr.CheckErr(err)
}
