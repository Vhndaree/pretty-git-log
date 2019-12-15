package service

import (
	"strings"
	"time"
)

// GetTodaysDate - returns todays date in YYYY-MM-DD format
func GetTodaysDate() string {
	return time.Now().Format("2006-01-02")
}

// Contains - checks if item contains in an array
func Contains(arr []string, item string) bool {
	for _, v := range arr {
		if strings.ToLower(v) == strings.ToLower(item) {
			return true
		}
	}

	return false
}

// Before - return string before pattern
func Before(value string, pattern string) string {
	// Get substring before a string.
	pos := strings.Index(value, pattern)
	if pos == -1 {
		return ""
	}
	return value[0:pos]
}
