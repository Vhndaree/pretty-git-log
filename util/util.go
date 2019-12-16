package util

import "time"

// GetTodaysDate - returns todays date in YYYY-MM-DD format
func GetTodaysDate() string {
	return time.Now().Format("2006-01-02")
}
