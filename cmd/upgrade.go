package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the Osuny CLI",
	Long: `Osuny CLI is a Command-Line Interface made with Go.
It's code is hosted at github.com/osunyorg/osuny.
This command launches 'go install github.com/osunyorg/osuny' to get the latest version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("There is a delay between the release and the ability to update")
		Shell("go", "install", "github.com/osunyorg/osuny@latest")
		versionCmd.Run(cmd, []string{})
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
}
