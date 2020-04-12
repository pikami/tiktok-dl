package utils

import (
	"fmt"
	"os"

	config "github.com/pikami/tiktok-dl/models/config"
)

// Log - Write to std out
func Log(a ...interface{}) {
	if !config.Config.Quiet {
		fmt.Println(a...)
	}
}

// Logf - Write formated text
func Logf(format string, a ...interface{}) {
	if !config.Config.Quiet {
		fmt.Printf(format, a...)
	}
}

// LogFatal - Write error and panic
func LogFatal(format string, a ...interface{}) {
	panic(fmt.Sprintf(format, a...))
}

// LogErr - Write error
func LogErr(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}
