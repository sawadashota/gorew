package cmd

import (
	"os"
	"path"

	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var (
	dotGoCmdDir string

	RootCmd = &cobra.Command{
		Use: "gorew",
	}

	InstallCmd = &cobra.Command{
		Use:   "install",
		Short: "Install golang cli from .gorew",
		PreRun: func(cmd *cobra.Command, args []string) {
			if dotGoCmdDir == "" {
				var err error
				dotGoCmdDir, err = homedir.Dir()

				if err != nil {
					panic(err)
				}
			}

			if errors := existsDir(dotGoCmdDir); len(errors) > 0 {
				for _, err := range errors {
					color.Red(err.Error())
				}
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			install(dotGoCmdDir)
		},
	}
)

func init() {
	// Flags for install
	InstallCmd.Flags().StringVarP(&dotGoCmdDir, "file", "f", "", "Directory to .gorew (default: $HOME)")

	RootCmd.AddCommand(InstallCmd)
}

func getSrcPathFromEnv() string {
	return path.Join(os.Getenv("GOPATH"), "src")
}

func existsDir(dirs ...string) []error {
	var errors []error

	for _, dir := range dirs {
		if _, err := os.Stat(dir); err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}
