package utils

import (
	"log"
)

// CheckErr - Checks if error and log
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
