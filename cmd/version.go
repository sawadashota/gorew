package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display this binary's version, build time and git hash of this build",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version:	%s\n", version)
		fmt.Printf("Git Hash:	%s\n", commit)
		fmt.Printf("Build Time:	%s\n", date)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
