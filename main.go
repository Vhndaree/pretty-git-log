package main

import (
	"github.com/vhndaree/pretty-git-log/file"
	"github.com/vhndaree/pretty-git-log/util"
)

// GO auto calls init before main
func init() {
	util.ParseAndSetEnv("")
}

func main() {
	file.Write()
}
