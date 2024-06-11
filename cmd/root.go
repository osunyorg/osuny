package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.4"

var rootCmd = &cobra.Command{
	Use:     "osuny",
	Version: version,
	Short:   "Work seamlessly with Osuny from the command line",
	Long: `Osuny creates static websites with Hugo.
This command line interface (CLI) helps the developer
to work on the generated websites.

More help at https://developers.osuny.org`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true
}
