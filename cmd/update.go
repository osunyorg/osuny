package cmd

import (
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates the Osuny theme",
	Long: `All Osuny websites are based on the Osuny theme.
It is loaded as a submodule in themes/osuny.`,
	Run: func(cmd *cobra.Command, args []string) {
		Shell("git",
			"pull",
			"--recurse-submodules",
			"--depth", "1")
		Shell("git",
			"submodule",
			"foreach",
			"git",
			"checkout",
			"main")
		Shell("git",
			"submodule",
			"foreach",
			"git",
			"pull")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
