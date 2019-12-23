package main

import (
	"github.com/Vhndaree/pretty-git-log/file"
	"github.com/Vhndaree/pretty-git-log/util"
)

func main() {
	util.ParseAndSetEnv("")
	file.Write()
}
