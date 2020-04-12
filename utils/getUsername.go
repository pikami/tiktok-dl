package utils

import (
	"fmt"
	"regexp"
	"strings"

	config "github.com/pikami/tiktok-dl/models/config"
	res "github.com/pikami/tiktok-dl/resources"
)

// GetUsername - Get's username from passed URL param
func GetUsername() string {
	return GetUsernameFromString(config.Config.URL)
}

// GetUsernameFromString - Get's username from passed param
func GetUsernameFromString(str string) string {
	if match := strings.Contains(str, "/"); !match { // Not url
		return strings.Replace(str, "@", "", -1)
	}

	if match, _ := regexp.MatchString(".+tiktok\\.com/@.+", str); match { // URL
		stripedSuffix := strings.Split(str, "@")[1]
		return strings.Split(stripedSuffix, "/")[0]
	}

	panic(fmt.Sprintf(res.ErrorCouldNotRecogniseURL, str))
}
