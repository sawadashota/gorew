package cmd

import (
	"github.com/sawadashota/gorew/pkg"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "rm",
	Short:   "Remove package",
	Example: "gorew rm github.com/some/package",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		h, err := pkg.NewHandler()
		if err != nil {
			return err
		}

		src := args[0]
		p := pkg.NewGoPackage(src)
		return h.Remove(p)
	},
}
