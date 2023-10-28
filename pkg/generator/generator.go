package generator

import (
	"io"

	"github.com/gophoria/gophoria/pkg/ast"
)

type Generator interface {
	Generate(ast *ast.Ast, writer *io.Writer) string
}
