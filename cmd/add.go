package cmd

import (
	"github.com/sawadashota/gorew/pkg"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add package",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		h, err := pkg.NewHandler()
		if err != nil {
			return err
		}

		src := args[0]
		p := pkg.NewGoPackage(src)
		return h.Add(p)
	},
}
