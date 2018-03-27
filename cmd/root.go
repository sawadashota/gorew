package cmd

import (
	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"os"
	"path"
)

var (
	srcPath     string
	binPath     string
	dotGoCmdDir string

	RootCmd = &cobra.Command{
		Use: "cmd",
	}

	InitCmd = &cobra.Command{
		Use:   "init",
		Short: "Save all golang cli",
		PreRun: func(cmd *cobra.Command, args []string) {
			if srcPath == "" {
				srcPath = getSrcPathFromEnv()
			}

			if binPath == "" {
				binPath = getBinPathFromEnv()
			}

			if dotGoCmdDir == "" {
				var err error
				dotGoCmdDir, err = homedir.Dir()

				if err != nil {
					panic(err)
				}
			}

			if errors := existsDir(srcPath, binPath, dotGoCmdDir); len(errors) > 0 {
				for _, err := range errors {
					color.Red(err.Error())
				}
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			gitDirs := goCommands(srcPath, binPath)
			recordFile(gitDirs, dotGoCmdDir)
			color.Green("created \"%v\" successfully!", path.Join(dotGoCmdDir, RecordFileName))
		},
	}

	InstallCmd = &cobra.Command{
		Use:   "install",
		Short: "Install golang cli from .gocmd",
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

func Commands() *cobra.Command {
	commands := RootCmd

	// Flags for init
	InitCmd.Flags().StringVarP(&srcPath, "srcPath", "s", "", "GOPATH (default: $GOPATH/src)")
	InitCmd.Flags().StringVarP(&binPath, "binPath", "b", "", "Bin path for golang commandDir (default: $GOPATH/bin)")
	InitCmd.Flags().StringVarP(&dotGoCmdDir, "file", "f", "", "Path for save .gocmd (default: $HOME)")

	// Flags for install
	InstallCmd.Flags().StringVarP(&dotGoCmdDir, "file", "f", "", "Directory to .gocmd (default: $HOME)")

	commands.AddCommand(InitCmd, InstallCmd)

	return commands
}

func getSrcPathFromEnv() string {
	return path.Join(os.Getenv("GOPATH"), "src")
}

func getBinPathFromEnv() string {
	return path.Join(os.Getenv("GOPATH"), "bin")
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
