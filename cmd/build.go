package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the website for production",
	Long: `The command is based on Hugo build, 
plus some configuration and the commands related to Pagefind.`,
	Run: func(cmd *cobra.Command, args []string) {
		command := exec.Command(`hugo`, `--minify`)
		command.Stdout = os.Stdout
		if err := command.Run(); err != nil {
			fmt.Println("could not run command: ", err)
		}
		command = exec.Command("npx",
			"pagefind",
			"--site", "public",
			"--exclude-selectors", pagefindExclude)
		command.Env = os.Environ()
		command.Env = append(command.Env, "npm_config_yes=true")
		command.Stdout = os.Stdout
		if err := command.Run(); err != nil {
			fmt.Println("could not run command: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
