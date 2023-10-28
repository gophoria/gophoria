package generator

import (
	"io"

	"github.com/gophoria/gophoria/pkg/ast"
)

type Generator interface {
	GenerateAll(ast *ast.Ast, writer io.Writer) error
	Generate(ast *ast.Ast, name string, writer io.Writer) error
}
