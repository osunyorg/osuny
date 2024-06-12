package cmd

import (
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup the local project",
	Long: `When the project is cloned, it is not completely ready.
It's necessary to install the javascript dependencies, 
and to set the theme to the main branch.`,
	Run: func(cmd *cobra.Command, args []string) {
		Shell("yarn", "install")
		updateCmd.Run(cmd, []string{})
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
