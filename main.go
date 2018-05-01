package main

import (
	"github.com/sawadashota/gocmd/cmd"
	"github.com/sawadashota/gocmd/completion"
)

func main() {
	cmd.Commands().Execute()
	completion.Run()
}
