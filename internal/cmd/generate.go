package cmd

import (
	"fmt"
	"os"
	"os/exec"
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
	Run: func(_ *cobra.Command, _ []string) {
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

	err = generatePrimitives(ast)
	if err != nil {
		return err
	}

	err = generateEnums(ast)
	if err != nil {
		return err
	}

	err = generateModels(ast)
	if err != nil {
		return err
	}

	cmd := exec.Command("go", "mod", "tidy")
	err = cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("gofmt", "-w", "./..")
	err = cmd.Run()
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

		err = gen.Generate(ast, item.Name.Identifier, f)
		if err != nil {
			return err
		}
	}

	return nil
}

func generatePrimitives(ast *ast.Ast) error {
	gen, err := createLibraryGenerator(ast)
	if err != nil {
		return nil
	}

	f, err := os.Create(path.Join(cfg.workingDir, "db", "DateTime.go"))
	if err != nil {
		return err
	}
	defer f.Close()

	err = gen.Generate(ast, "DateTime", f)
	if err != nil {
		return err
	}

	return nil
}

func generateEnums(ast *ast.Ast) error {
	gen, err := createLibraryGenerator(ast)
	if err != nil {
		return err
	}

	for _, item := range ast.Enums {
		f, err := os.Create(path.Join(cfg.workingDir, "db", fmt.Sprintf("%s.go", item.Name.Identifier)))
		if err != nil {
			return err
		}
		defer f.Close()

		f.Write([]byte("package db\n\n"))

		err = gen.Generate(ast, item.Name.Identifier, f)
		if err != nil {
			return err
		}
	}

	return nil
}

func generateModels(ast *ast.Ast) error {
	gen, err := createLibraryGenerator(ast)
	if err != nil {
		return err
	}

	for _, item := range ast.Models {
		f, err := os.Create(path.Join(cfg.workingDir, "db", fmt.Sprintf("%s.go", item.Name.Identifier)))
		if err != nil {
			return err
		}
		defer f.Close()

		f.Write([]byte("package db\n\n"))

		err = gen.Generate(ast, item.Name.Identifier, f)
		if err != nil {
			return err
		}
	}

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

func createLibraryGenerator(ast *ast.Ast) (generator.Generator, error) {
	for _, config := range ast.Config {
		if config.Type == "db" {
			for _, item := range config.Items {
				if item.Identifier.Identifier == "lib" {
					switch item.Value.Value {
					case "sqlx":
						return generator.NewSqlxGenerator(), nil
					}
				}
			}
		}
	}

	return nil, fmt.Errorf("unable to find db library")
}
