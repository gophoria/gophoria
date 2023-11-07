package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/gophoria/gophoria/internal/utils"
	"github.com/gophoria/gophoria/pkg/ast"
	"github.com/gophoria/gophoria/pkg/generator"
	"github.com/spf13/cobra"
)

type GenerateConfig struct {
	override bool
}

var generateCfg GenerateConfig

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate code",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var generateDbCommand = &cobra.Command{
	Use:   "db",
	Short: "Generate db",
	Run: func(cmd *cobra.Command, args []string) {
		err := generateDb()
		if err != nil {
			exitWithError(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.PersistentFlags().BoolVar(&generateCfg.override, "override", false, "Override files if exists")
	generateCmd.AddCommand(generateDbCommand)
}

func generateDb() error {
	err := createDirectoryStruct()
	if err != nil {
		return err
	}

	ast, err := utils.ParseFile(cfg.file)
	if err != nil {
		return err
	}

	err = generateMigrations(ast)
	if err != nil {
		return err
	}

	return nil
}

func generateMigrations(ast *ast.Ast) error {
	gen, err := createMigrationGenerator(ast)
	if err != nil {
		return err
	}

	for idx, item := range ast.Models {
		f, err := os.Create(path.Join(cfg.workingDir, "migrations", fmt.Sprintf("%d_%s.sql", idx+1, item.Name.Identifier)))
		if err != nil {
			return err
		}
		defer f.Close()

		gen.Generate(ast, item.Name.Identifier, f)
	}

	return nil
}

func generateEnums(ast *ast.Ast) error {
	return nil
}

func generateModels(ast *ast.Ast) error {
	return nil
}

func createMigrationGenerator(ast *ast.Ast) (generator.Generator, error) {
	for _, config := range ast.Config {
		if config.Type == "db" {
			for _, item := range config.Items {
				if item.Identifier.Identifier == "provider" {
					switch item.Value.Value {
					case "sqlite3":
						return generator.NewSqlite3Generator(), nil
					}
				}
			}
		}
	}

	return nil, fmt.Errorf("unable to find db provider")
}