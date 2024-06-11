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
			fmt.Println(`Preparing Pagefind index with exclusions:`)
			exec.Command(`hugo`)

			pagefindExclude := `"` +
				// Categories: No list of categories
				`.categories__taxonomy, .categories__term, ` +
				`.posts_categories__taxonomy, .posts_categories__term, ` +
				`.events_categories__taxonomy, .events_categories__term, ` +
				// Diplomas: No list of diplomas or block diplomas
				`.diplomas__taxonomy, .block-diplomas, ` +
				// Agenda events: No list of events or block events
				`.events__section, .block-agenda, ` +
				// Organizations: No list of organizations or block organizations
				`.organizations__section, .block-organizations, ` +
				// Pages: No block pages (there's no difference between list and page)
				`.block-pages, ` +
				// Persons: no list or block
				`.persons__section, .block-persons, ` +
				// No list of people's facets
				`.administrators__term, .authors__term, .researchers__term, .teachers__term, ` +
				// Posts: no list, block posts, or post sidebar
				`.posts__section, .block-posts, .post-sidebar, ` +
				// Programs: no block
				`.block-programs` +
				`"`
			fmt.Println(pagefindExclude)
			command := exec.Command(`npx`, `pagefind`,
				`--verbose`,
				`--site`, `public`,
				`--exclude-selectors`, pagefindExclude,
				`--output-subdir`, `../static/pagefind`)
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
