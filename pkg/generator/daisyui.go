package generator

import (
	"fmt"
	"os"
	"path"

	"github.com/gophoria/gophoria/pkg/ast"
)

type DaisyUiGenerator struct {
	ast *ast.Ast
	cfg *GeneratorConfig
}

func init() {
	RegisterGenerator("daisyui", NewDaisyUiGenerator())
}

func NewDaisyUiGenerator() *DaisyUiGenerator {
	g := DaisyUiGenerator{}

	return &g
}

func (d *DaisyUiGenerator) GenerateAll(ast *ast.Ast, cfg *GeneratorConfig) error {
	d.ast = ast
	d.cfg = cfg

	return nil
}

func (d *DaisyUiGenerator) Generate(ast *ast.Ast, cfg *GeneratorConfig, name string) error {
	d.ast = ast
	d.cfg = cfg

	return nil
}

func (d *DaisyUiGenerator) generateUi(item *ast.Model) error {
	f, err := os.Create(path.Join(d.cfg.WorkingDir, "view", fmt.Sprintf("%s.templ", item.Name.Identifier)))
	if err != nil {
		return err
	}
	defer f.Close()

	f.Write([]byte("package view\n\n"))

	return nil
}
