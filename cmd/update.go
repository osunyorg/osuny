package cmd

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates the Osuny theme",
	Long: `All Osuny websites are based on the Osuny theme.
It is loaded as a submodule in themes/osuny.`,
	Run: func(cmd *cobra.Command, args []string) {

		branch := "main"

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
			RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		})
		if err != nil && err != git.NoErrAlreadyUpToDate {
			fmt.Println(err)
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

			err = subRepo.Fetch(&git.FetchOptions{
				RemoteName: "origin",
			})
			if err != nil && err != git.NoErrAlreadyUpToDate {
				log.Fatalf("Error while fetching submodule: %s", err)
			}

			w, err := subRepo.Worktree()
			if err != nil {
				log.Fatalf("Error getting submodule worktree: %s", err)
			}

			branchRefName := plumbing.NewBranchReferenceName(branch)
			branchCoOpts := git.CheckoutOptions{
				Branch: plumbing.ReferenceName(branchRefName),
				Force:  true,
			}

			if err := w.Checkout(&branchCoOpts); err != nil {
				// Warning("local checkout of branch '%s' failed, will attempt to fetch remote branch of same name.", branch)
				// Warning("like `git checkout <branch>` defaulting to `git checkout -b <branch> --track <remote>/<branch>`")

				mirrorRemoteBranchRefSpec := fmt.Sprintf("refs/heads/%s:refs/remotes/origin/%s", branch, branch)
				err = fetchOrigin(subRepo, mirrorRemoteBranchRefSpec)
				if err != nil {
					log.Fatalf("Error fetching origin: %s", err)
				}

				err = w.Checkout(&branchCoOpts)
				if err != nil {
					log.Fatalf("Error fetching origin: %s", err)
				}
			}
		}
	},
}

func fetchOrigin(repo *git.Repository, refSpecStr string) error {
	remote, err := repo.Remote("origin")
	if err != nil {
		log.Fatalf("Error getting remote origin: %s", err)
	}

	var refSpecs []config.RefSpec
	if refSpecStr != "" {
		refSpecs = []config.RefSpec{config.RefSpec(refSpecStr)}
	}

	if err = remote.Fetch(&git.FetchOptions{
		RefSpecs: refSpecs,
	}); err != nil {
		if err == git.NoErrAlreadyUpToDate {
			fmt.Print("refs already up to date")
		} else {
			return fmt.Errorf("fetch origin failed: %v", err)
		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
