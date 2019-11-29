package cmd

import (
	"github.com/sawadashota/gorew/pkg"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update package",
	Long: `Update given package or all package
if given no args, update all package`,
	Example: `gorew update github.com/some/package
gorew update`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		h, err := pkg.NewHandler()
		if err != nil {
			return err
		}

		if len(args) == 1 {
			src := args[0]
			p := pkg.NewGoPackage(src)
			return h.Update(p)
		}

		return h.UpdateAll()
	},
}
