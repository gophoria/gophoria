package cmd

import (
	"os"

	"github.com/gophoria/gophoria/internal/utils"
	"github.com/spf13/cobra"
)

type InitConfig struct {
	dbProvider string
	dbUrl      string
	dbLib      string

	uiLib        string
	uiComponents string

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

		utils.GenerateProject(f, utils.GenerateConfig{
			DbProvider: initCfg.dbProvider,
			DbUrl:      initCfg.dbUrl,
			DbLib:      initCfg.dbLib,

			UiLib:        initCfg.uiLib,
			UiComponents: initCfg.uiComponents,

			WithExample: initCfg.withExample,
		})

		err = createDirectoryStruct()
		if err != nil {
			exitWithError(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVar(&initCfg.dbProvider, "db", "sqlite3", "Database provider")
	initCmd.Flags().StringVar(&initCfg.dbUrl, "dbUrl", ":memory:", "Database url")
	initCmd.Flags().StringVar(&initCfg.dbLib, "dbLib", "sqlx", "Database library")
	initCmd.Flags().StringVar(&initCfg.uiLib, "ui", "templ", "UI library")
	initCmd.Flags().StringVar(&initCfg.uiComponents, "components", "daisyui", "UI components")
	initCmd.Flags().BoolVar(&initCfg.withExample, "example", false, "Example project")
}
