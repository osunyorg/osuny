package cmd

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone [repository name]",
	Short: "Properly clone a repository",
	Args:  cobra.ExactArgs(1),
	Long: `The Osuny admin generates repositories on github.com/osunyorg.
As there are submodules, they need to be cloned with the --recurse-submodule option.

The repository name should be like "epv-site", not "osunyorg/epv-site".
This is intended to maintain the centralized maintenance of the theme.`,
	Run: func(cmd *cobra.Command, args []string) {
		var repository = args[0]
		var url = fmt.Sprintf("git@github.com:osunyorg/%s.git", repository)
		git.PlainClone(repository, false, &git.CloneOptions{
			URL:               url,
			RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		})
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}
