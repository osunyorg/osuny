package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:     "serve",
	Aliases: []string{"watch", "server", "dev", "s"},
	Short:   "Start the Hugo server",
	Long: `Osuny's websites are based on Hugo.
When you code a website, you work with a local version on a local server.
This command launches the server.
It's just a proxy for "hugo server".`,
	Run: func(cmd *cobra.Command, args []string) {
		if Search {
			fmt.Println("Preparing Pagefind index with exclusions:")
			fmt.Println(pagefindExclude)

			exec.Command("hugo")

			command := exec.Command("npx", "pagefind",
				"--verbose",
				"--site", "public",
				"--exclude-selectors", pagefindExclude,
				"--output-subdir", "../static/pagefind")
			command.Stdout = os.Stdout
			if err := command.Run(); err != nil {
				fmt.Println("could not run command: ", err)
			}
		}
		command := exec.Command("hugo", "serve")
		command.Stdout = os.Stdout
		if err := command.Run(); err != nil {
			fmt.Println("could not run command: ", err)
		}
	},
}

var Search bool

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().BoolVarP(&Search, "search", "s", false, "with Pagefind search")

}
