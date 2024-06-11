package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.5"
var pagefindExclude = `"` +
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

var rootCmd = &cobra.Command{
	Use:     "osuny",
	Version: version,
	Short:   "Work seamlessly with Osuny from the command line",
	Long: `Osuny creates static websites with Hugo.
This command line interface (CLI) helps the developer
to work on the generated websites.

More help at https://developers.osuny.org`,
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
