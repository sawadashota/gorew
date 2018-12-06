package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	version string
	gitHash string

	TimeNow time.Time
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display this binary's version, build time and git hash of this build",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version:	%s\n", version)
		fmt.Printf("Git Hash:	%s\n", gitHash)
		fmt.Printf("Build Time:	%s\n", timeNow().UTC().Format("2006-01-02 15:04:05 MST"))
	},
}

func init() {
	version = getStringEnv("CIRCLE_TAG", "dev-master")
	gitHash = getStringEnv("CIRCLE_SHA1", "undefined")
	RootCmd.AddCommand(versionCmd)
}

// getStringEnv returns environment variable
// if the key doesn't exist, return default value
func getStringEnv(key string, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	return v
}

func timeNow() time.Time {
	if &TimeNow != nil {
		return TimeNow
	}

	return time.Now()
}
