package cmd

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates the Osuny theme",
	Long: `All Osuny websites are based on the Osuny theme.
It is loaded as a submodule in themes/osuny.`,
	Run: func(cmd *cobra.Command, args []string) {

		// branch := "main"

		repository, err := git.PlainOpen(".")
		if err != nil {
			fmt.Println(err)
			return
		}

		worktree, err := repository.Worktree()
		if err != nil {
			fmt.Println(err)
			return
		}
		err = worktree.Pull(&git.PullOptions{
			RemoteName:        "origin",
			ReferenceName:     "refs/heads/main",
			SingleBranch:      true,
			RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
			Force:             true,
		})
		if err != nil && err != git.NoErrAlreadyUpToDate {
			fmt.Println("Pulled: %s", err)
		}

		submodules, err := worktree.Submodules()
		if err != nil {
			log.Fatalf("Error getting submodules: %s", err)
		}

		for _, submodule := range submodules {
			fmt.Printf("Updating submodule %s\n", submodule.Config().Name)
			subRepo, err := submodule.Repository()
			if err != nil {
				log.Fatalf("Error while getting submodule: %s", err)
			}

			w, err := subRepo.Worktree()
			if err != nil {
				log.Fatalf("Error getting submodule worktree: %s", err)
			}

			w.Checkout(&git.CheckoutOptions{
				Branch: "refs/heads/main",
			})
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
