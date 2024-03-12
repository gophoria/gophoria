package generator

import (
	"io"

	"github.com/gophoria/gophoria/pkg/ast"
)

type DaisyUiGenerator struct{}

func init() {
	RegisterGenerator("daisyui", NewDaisyUiGenerator())
}

func NewDaisyUiGenerator() *DaisyUiGenerator {
	g := DaisyUiGenerator{}

	return &g
}

func (d *DaisyUiGenerator) GenerateAll(ast *ast.Ast, writer io.Writer) error {
	return nil
}

func (d *DaisyUiGenerator) Generate(ast *ast.Ast, name string, writer io.Writer) error {
	return nil
}
