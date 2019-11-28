package cmd

import (
	"github.com/sawadashota/gorew/pkg"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install Go CLI from .gorew",
	RunE: func(cmd *cobra.Command, args []string) error {
		h, err := pkg.NewHandler()
		if err != nil {
			return err
		}

		return h.InstallAll()
	},
}
