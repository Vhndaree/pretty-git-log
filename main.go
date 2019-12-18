package main

import (
	"github.com/Vhndaree/pretty-git-log/file"
	"github.com/Vhndaree/pretty-git-log/util"
)

// GO auto calls init before main
func init() {
	util.ParseAndSetEnv("")
}

func main() {
	file.Write()
}
