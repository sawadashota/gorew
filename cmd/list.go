package cmd

import (
	"fmt"
	"os"

	"github.com/sawadashota/gorew/pkg"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List package",
	RunE: func(cmd *cobra.Command, args []string) error {
		h, err := pkg.NewHandler()
		if err != nil {
			return err
		}

		ps, err := h.List()
		if err != nil {
			return err
		}

		if len(ps) == 0 {
			_, _ = fmt.Fprintln(os.Stdout, "there are no packages yet")
			return nil
		}

		for _, p := range ps {
			_, _ = fmt.Fprintln(os.Stdout, p.Source())
		}

		return nil
	},
}
