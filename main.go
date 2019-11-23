package main

import (
	"github.com/sawadashota/gorew/cmd"
	"github.com/sawadashota/gorew/completion"
)

func main() {
	_ = cmd.RootCmd.Execute()
	completion.Run()
}
