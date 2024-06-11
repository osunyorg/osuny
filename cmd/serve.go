package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:     "serve",
	Aliases: []string{"watch", "server", "w"},
	Short:   "Start the Hugo server (without pagefind search)",
	Long: `Osuny's websites are based on Hugo.
When you code a website, you work with a local version on a local server.
This command launches the server.
It's just a proxy for "hugo server".`,
	Run: func(cmd *cobra.Command, args []string) {
		command := exec.Command("hugo", "serve")
		command.Stdout = os.Stdout
		if err := command.Run(); err != nil {
			fmt.Println("could not run command: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
