package util

import (
	"os"
	"strings"
	"time"

	"github.com/Vhndaree/pretty-git-log/interfaces"
)

// GetTodaysDate - returns todays date in YYYY-MM-DD format
func GetTodaysDate() string {
	return time.Now().Format("2006-01-02")
}

// GetSpacesOfLength - gives string with space n
func GetSpacesOfLength(length int) string {
	spaces := ""
	for i := 0; i < length; i++ {
		spaces += " "
	}

	return spaces + " \t"
}

// FilterCommitsByUser - filters commits by username
func FilterCommitsByUser(c interfaces.Commits, userName string) (cs interfaces.Commits) {
	for _, v := range c {
		if strings.ToLower(v.UserName.Login) == strings.ToLower(userName) {
			cs = append(cs, v)
		}
	}

	return
}

// Exists returns whether the given file or directory exists
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return true, err
}
