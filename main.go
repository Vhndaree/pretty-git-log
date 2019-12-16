package main

import (
	"github.com/Vhndaree/task-monitor/file"
	"github.com/Vhndaree/task-monitor/util"
)

func main() {
	util.ParseAndSetEnv("")
	file.Write()
}
