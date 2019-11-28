package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "gorew",
}

func init() {
	RootCmd.AddCommand(installCmd, addCmd, removeCmd, listCmd)
}
