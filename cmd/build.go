package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the website for production",
	Long: `The command is based on Hugo build, 
plus some configuration and the commands related to Pagefind.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Errors should crash build, so we don't use function Shell
		command := exec.Command(`hugo`, `--minify`)
		command.Stdout = os.Stdout
		if err := command.Run(); err != nil {
			panic(err)
		}
		command = exec.Command("npx",
			"pagefind",
			"--site", "public",
			"--exclude-selectors", pagefindExclude)
		command.Env = os.Environ()
		command.Env = append(command.Env, "npm_config_yes=true")
		command.Stdout = os.Stdout
		if err := command.Run(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
