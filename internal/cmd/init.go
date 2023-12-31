package cmd

import (
	"os"
	"path"

	"github.com/gophoria/gophoria/internal/utils"
	"github.com/spf13/cobra"
)

type InitConfig struct {
	db          string
	dbUrl       string
	withExample bool
}

var initCfg InitConfig

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init gophoria project",
	Run: func(_ *cobra.Command, _ []string) {
		f, err := os.Create(cfg.file)
		if err != nil {
			exitWithError(err)
		}
		defer f.Close()

		utils.GenerateProject(f, initCfg.db, initCfg.dbUrl, initCfg.withExample)

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
	initCmd.Flags().BoolVar(&initCfg.withExample, "example", false, "Example project")
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
