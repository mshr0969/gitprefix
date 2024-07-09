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
		"feat: featureの略。機能の追加や更新の際に使用する。",
		"fix: bug fixの略。バグの修正の際に使用する。",
		"docs: documentationの略。ドキュメントの変更の際に使用する。",
		"style: コードの意味に影響を与えない変更の際に使用する。(インデント、セミコロンの追加など)",
		"refactor: コードのリファクタリングの際に使用する。",
		"perf: performanceの略。パフォーマンスの向上の際に使用する。",
		"test: テストの追加や修正の際に使用する。",
		"chore: その他、雑務的な変更の際に使用する。",
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
