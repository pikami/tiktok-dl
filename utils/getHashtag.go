package utils

import (
	"fmt"
	"strings"

	res "github.com/pikami/tiktok-dl/resources"
)

// GetHashtagFromURL - Get's tag name from passed url
func GetHashtagFromURL(str string) string {
	if match := strings.Contains(str, "/tag/"); match {
		return strings.Split(str, "/tag/")[1]
	}

	panic(fmt.Sprintf(res.ErrorCouldNotRecogniseURL, str))
}
