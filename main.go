package main

import (
	"github.com/sawadashota/gocmd/cmd"
	"github.com/sawadashota/gocmd/completion"
)

func main() {
	_ = cmd.RootCmd.Execute()
	completion.Run()
}
