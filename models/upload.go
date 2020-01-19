package models

import (
	"encoding/json"
	"strings"
)

// Upload - Upload object
type Upload struct {
	ShareLink string `json:"shareLink"`
	URL       string `json:"url"`
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
