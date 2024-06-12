package cmd

import (
	"fmt"

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
		if withSearch {
			fmt.Println("Preparing Pagefind index with exclusions:", pagefindExclude)

			Shell("hugo")

			Shell("npx", "pagefind",
				"--verbose",
				"--site", "public",
				"--exclude-selectors", pagefindExclude,
				"--output-subdir", "../static/pagefind")
		}

		Shell("hugo", "server",
			"--port", fmt.Sprint(serverPort))
	},
}

var withSearch bool
var serverPort int

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().BoolVarP(&withSearch, "with-search", "", false, "with Pagefind search")
	serveCmd.Flags().IntVarP(&serverPort, "port", "p", 1313, "port on which the server will listen")

}
