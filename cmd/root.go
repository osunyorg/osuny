package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "osuny",
	Version: "0.0.2",
	Short:   "A command line interface to work with Osuny",
	Long: `Osuny lets you create static websites with Hugo.
This command line interface helps you interact simply as a developer.`,
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
