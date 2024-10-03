package cmd

import (
	"fmt"
	"os/exec"

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

var generateUiCommand = &cobra.Command{
	Use:   "ui",
	Short: "Generate ui",
	Run: func(_ *cobra.Command, _ []string) {
		err := generateUi()
		if err != nil {
			exitWithError(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.PersistentFlags().BoolVar(&generateCfg.override, "override", false, "Override files if exists")
	generateCmd.AddCommand(generateDbCommand)
	generateCmd.AddCommand(generateUiCommand)
}

func generateDb() error {
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

	err = formatProject()
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

	cfg := createGeneratorCfg()

	for _, item := range ast.Models {
		err = gen.Generate(ast, cfg, item.Name.Identifier)
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

	cfg := createGeneratorCfg()

	err = gen.Generate(ast, cfg, "DateTime")
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

	cfg := createGeneratorCfg()

	for _, item := range ast.Enums {
		err = gen.Generate(ast, cfg, item.Name.Identifier)
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

	cfg := createGeneratorCfg()

	for _, item := range ast.Models {
		err = gen.Generate(ast, cfg, item.Name.Identifier)
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
					gen, err := generator.GetGenerator(item.Value.Value)
					if err != nil {
						return nil, err
					}

					return gen, nil
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
					gen, err := generator.GetGenerator(item.Value.Value)
					if err != nil {
						return nil, err
					}

					return gen, nil
				}
			}
		}
	}

	return nil, fmt.Errorf("unable to find db library")
}

func generateUi() error {
	ast, err := utils.ParseFile(cfg.file)
	if err != nil {
		return err
	}

	err = generatePages(ast)
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

func generatePages(ast *ast.Ast) error {
	gen, err := createPageGenerator(ast)
	if err != nil {
		return err
	}

	cfg := createGeneratorCfg()

	for _, item := range ast.Models {
		err = gen.Generate(ast, cfg, item.Name.Identifier)
		if err != nil {
			return err
		}
	}

	return nil
}

func createPageGenerator(ast *ast.Ast) (generator.Generator, error) {
	for _, config := range ast.Config {
		if config.Type == "ui" {
			for _, item := range config.Items {
				if item.Identifier.Identifier == "components" {
					gen, err := generator.GetGenerator(item.Value.Value)
					if err != nil {
						return nil, err
					}

					return gen, nil
				}
			}
		}
	}

	return nil, fmt.Errorf("unable to find ui components")
}

func formatProject() error {
	cmd := exec.Command("go", "mod", "tidy")
	err := cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("gofmt", "-s", "-w", ".")
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func createGeneratorCfg() *generator.GeneratorConfig {
	return &generator.GeneratorConfig{
		Override:   generateCfg.override,
		WorkingDir: cfg.workingDir,
	}
}
