package utils

import (
	"io"
	"net/http"
	"os"
)

// DownloadFile - Downloads content from `url` and stores it in `outputPath`
func DownloadFile(outputPath string, url string) {
	// Get the data
	resp, err := http.Get(url)
	CheckErr(err)
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(outputPath)
	CheckErr(err)
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	CheckErr(err)
}
