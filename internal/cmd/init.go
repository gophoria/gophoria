package cmd

import (
	"os"

	"github.com/gophoria/gophoria/internal/utils"
	"github.com/spf13/cobra"
)

type InitConfig struct {
	db    string
	dbUrl string
}

var initCfg InitConfig

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Init gophoria project",
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.Create(cfg.file)
		if err != nil {
			exitWithError(err)
		}
		defer f.Close()

		utils.GenerateProject(f, initCfg.db, initCfg.dbUrl)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVar(&initCfg.db, "db", "sqlite3", "Database provider")
	initCmd.Flags().StringVar(&initCfg.dbUrl, "dbUrl", ":memory:", "Database url")
}
