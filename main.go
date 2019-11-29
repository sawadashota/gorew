package main

import (
	"fmt"
	"os"

	"github.com/sawadashota/gorew/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stdout, err)
	}
}
