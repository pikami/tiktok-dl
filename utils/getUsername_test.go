package utils

import (
	"testing"

	config "github.com/pikami/tiktok-dl/models/config"
	testUtil "github.com/pikami/tiktok-dl/unitTestUtil"
)

func TestGetUsername(t *testing.T) {
	testCaseDelegate := func(t *testing.T, url string, username string) {
		tu := testUtil.TestUtil{T: t}
		config.Config.URL = url
		actual := GetUsername()
		tu.AssertString(actual, username, "Username")
	}

	testVideoURL := func(t *testing.T) {
		testCaseDelegate(t, "https://www.tiktok.com/@some_username/video/0000000000000000000", "some_username")
	}

	testProfileURL := func(t *testing.T) {
		testCaseDelegate(t, "https://www.tiktok.com/@some_username", "some_username")
	}

	testPlainUsername := func(t *testing.T) {
		testCaseDelegate(t, "some_username", "some_username")
	}

	testAtUsername := func(t *testing.T) {
		testCaseDelegate(t, "@some_username", "some_username")
	}

	t.Run("Video URL", testVideoURL)
	t.Run("Username URL", testProfileURL)
	t.Run("Plain username", testPlainUsername)
	t.Run("Username with @ suffix", testAtUsername)
}
