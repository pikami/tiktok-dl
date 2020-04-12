package models

import (
	"os"
	"testing"

	testUtil "github.com/pikami/tiktok-dl/unitTestUtil"
	fileio "github.com/pikami/tiktok-dl/utils/fileio"
)

func TestParseUploads(t *testing.T) {
	tu := testUtil.TestUtil{T: t}
	jsonStr := "[{\"url\":\"some_url\",\"shareLink\":\"some_share_link\",\"caption\":\"some_caption\", \"uploader\": \"some.uploader\",\"sound\":{\"title\":\"some_title\",\"link\":\"some_link\"}}]"
	actual := ParseUploads(jsonStr)

	tu.AssertInt(len(actual), 1, "Array len")

	tu.AssertString(actual[0].URL, "some_url", "URL")
	tu.AssertString(actual[0].Caption, "some_caption", "Caption")
	tu.AssertString(actual[0].ShareLink, "some_share_link", "ShareLink")
	tu.AssertString(actual[0].Uploader, "some.uploader", "Uploader")

	tu.AssertString(actual[0].Sound.Link, "some_link", "Sound.Link")
	tu.AssertString(actual[0].Sound.Title, "some_title", "Sound.Title")
}

func TestParseUpload(t *testing.T) {
	tu := testUtil.TestUtil{T: t}
	jsonStr := "{\"url\":\"some_url\",\"shareLink\":\"some_share_link\",\"caption\":\"some_caption\",\"sound\":{\"title\":\"some_title\",\"link\":\"some_link\"}}"
	actual := ParseUpload(jsonStr)

	tu.AssertString(actual.URL, "some_url", "URL")
	tu.AssertString(actual.Caption, "some_caption", "Caption")
	tu.AssertString(actual.ShareLink, "some_share_link", "ShareLink")

	tu.AssertString(actual.Sound.Link, "some_link", "Sound.Link")
	tu.AssertString(actual.Sound.Title, "some_title", "Sound.Title")
}

func TestGetUploadID(t *testing.T) {
	tu := testUtil.TestUtil{T: t}
	var upload Upload
	upload.ShareLink = "http://pikami.org/some_thing/some_upload_id"
	actual := upload.GetUploadID()

	tu.AssertString(actual, "some_upload_id", "Upload ID")
}

func TestWriteToFile(t *testing.T) {
	tu := testUtil.TestUtil{T: t}
	expected := "{\"url\":\"some_url\",\"shareLink\":\"some_share_link\",\"caption\":\"some_caption\",\"uploader\":\"some.uploader\",\"sound\":{\"title\":\"some_title\",\"link\":\"some_link\"}}"
	filePath := "test_file.txt"
	upload := Upload{
		URL:       "some_url",
		Caption:   "some_caption",
		ShareLink: "some_share_link",
		Uploader:  "some.uploader",
		Sound: Sound{
			Link:  "some_link",
			Title: "some_title",
		},
	}

	upload.WriteToFile(filePath)

	actual := fileio.ReadFileToString(filePath)
	tu.AssertString(actual, expected, "File content")

	os.Remove(filePath)
}
