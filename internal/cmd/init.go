package cmd

import (
	"os"
	"path"

	"github.com/gophoria/gophoria/internal/utils"
	"github.com/spf13/cobra"
)

type InitConfig struct {
	db    string
	dbUrl string
}

var initCfg InitConfig

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init gophoria project",
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.Create(cfg.file)
		if err != nil {
			exitWithError(err)
		}
		defer f.Close()

		utils.GenerateProject(f, initCfg.db, initCfg.dbUrl)

		err = createDirectoryStruct()
		if err != nil {
			exitWithError(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVar(&initCfg.db, "db", "sqlite3", "Database provider")
	initCmd.Flags().StringVar(&initCfg.dbUrl, "dbUrl", ":memory:", "Database url")
}

func createDirectoryStruct() error {
	err := utils.CreateDirIfNotExists(path.Join(cfg.workingDir, "db"))
	if err != nil {
		return err
	}
	err = utils.CreateDirIfNotExists(path.Join(cfg.workingDir, "migrations"))
	if err != nil {
		return err
	}

	return nil
}
