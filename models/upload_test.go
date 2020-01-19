package models

import "testing"

// TestParseUploads - Test parsing
func TestParseUploads(t *testing.T) {
	jsonStr := "[{\"shareLink\":\"some_share_link\", \"url\": \"some_url\"}]"
	actual := ParseUploads(jsonStr)

	expectedLen := 1
	if len(actual) != expectedLen {
		t.Errorf("Array len incorrect: Expected %d, but got %d", expectedLen, len(actual))
	}

	expectedShareLink := "some_share_link"
	if actual[0].ShareLink != expectedShareLink {
		t.Errorf("ShareLink is incorrect: Expected %s, but got %s", expectedShareLink, actual[0].ShareLink)
	}

	expectedURL := "some_url"
	if actual[0].URL != expectedURL {
		t.Errorf("URL is incorrect: Expected %s, but got %s", expectedURL, actual[0].URL)
	}
}

func TestGetUploadID(t *testing.T) {
	var upload Upload
	upload.ShareLink = "http://pikami.org/some_thing/some_upload_id"
	expected := "some_upload_id"

	actual := upload.GetUploadID()

	if actual != expected {
		t.Errorf("UploadId is incorrect: Expected %s, but got %s", expected, actual)
	}
}
