package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get the installed version of Osuny CLI",
	Long: `Get the version of the Osuny command line interface (CLI).
This is neither the current theme version nore the current admin version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Osuny CLI", rootCmd.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
