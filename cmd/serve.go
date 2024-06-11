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
		if WithSearch {
			fmt.Println("Preparing Pagefind index with exclusions:", pagefindExclude)

			Shell("hugo")

			Shell("npx", "pagefind",
				"--verbose",
				"--site", "public",
				"--exclude-selectors", pagefindExclude,
				"--output-subdir", "../static/pagefind")
		}
		Shell("hugo", "serve")
	},
}

var WithSearch bool

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().BoolVarP(&WithSearch, "with-search", "", false, "with Pagefind search")

}
