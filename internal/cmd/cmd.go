package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Config struct {
	file       string
	workingDir string
}

var cfg Config

var rootCmd = &cobra.Command{
	Use:   "gophoria",
	Short: "Gophoria is full-stack Go framework",
	Long: `Gophoria helps you by generateing code for dealing with database.
Additionally it can generate UI for Web, TUI and GUI`,
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	cwd, _ := os.Getwd()

	rootCmd.PersistentFlags().StringVarP(&cfg.file, "file", "f", "project.gophoria", "Gophoria main file")
	rootCmd.PersistentFlags().StringVarP(&cfg.workingDir, "workingDir", "w", cwd, "Gophoria main file")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		exitWithError(err)
	}
}

func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}
