package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// upgradeCmd represents the upgrade command
var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the Osuny CLI",
	Long: `Osuny CLI is a Command-Line Interface made with Go.
It's code is hosted at github.com/osunyorg/osuny.
This command launches 'go install github.com/osunyorg/osuny' to get the latest version.`,
	Run: func(cmd *cobra.Command, args []string) {
		// FIXME does not seem to work
		command := exec.Command("go", "install", "github.com/osunyorg/osuny@latest")
		command.Stdout = os.Stdout
		if err := command.Run(); err != nil {
			fmt.Println("could not run command: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
}
