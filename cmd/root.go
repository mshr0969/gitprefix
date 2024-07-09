package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gitprefix",
	Short: "A CLI tool to show git commit message prefixes.",
	Run: func(cmd *cobra.Command, args []string) {
		showPrefixes()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func showPrefixes() {
	green := "\033[32m"
	reset := "\033[0m"

	prefixes := []string{
		"feat: A new feature",
		"fix: A bug fix",
		"docs: Documentation only changes",
		"style: Changes that do not affect the meaning of the code (white-space, formatting, etc)",
		"refactor: A code change that neither fixes a bug nor adds a feature",
		"perf: A code change that improves performance",
		"test: Adding missing or correcting existing tests",
		"chore: Changes to the build process or auxiliary tools and libraries such as documentation generation",
	}

	fmt.Println("Commit Message Prefixes:")
	for _, prefix := range prefixes {
		parts := strings.SplitN(prefix, ": ", 2)
		if len(parts) == 2 {
			fmt.Println(green + parts[0] + ":" + reset + " " + parts[1])
		} else {
			fmt.Println(prefix)
		}
	}
}
