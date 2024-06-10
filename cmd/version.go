/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get the installed version of Osuny CLI",
	Long: `Get the version of the Osuny command line interface (CLI).
This is neither the current theme version nore the current admin version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Osuny CLI v0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
