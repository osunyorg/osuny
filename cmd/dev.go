package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// devCmd represents the dev command
var devCmd = &cobra.Command{
	Use:     "dev",
	Aliases: []string{"d"},
	Short:   "Start the Hugo server (with pagefind search)",
	Long: `This command builds the website in the public directory, 
then launches pagefind to index it, 
stores the pagefind artefacts in the static directory, 
and launches the server.

It makes the search work locally.`,
	Run: func(cmd *cobra.Command, args []string) {
		exec.Command(`hugo`)
		command := exec.Command("npx", "pagefind",
			"--site", "public",
			"--output-subdir", "../static/pagefind")
		// TODO exclusions
		// npx pagefind --site 'public' --output-subdir '../static/pagefind'
		command.Stdout = os.Stdout
		if err := command.Run(); err != nil {
			fmt.Println("could not run command: ", err)
		}
		serveCmd.Run(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(devCmd)
}
