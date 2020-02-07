package utils

import (
	config "../models/config"
	"fmt"
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
