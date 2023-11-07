package utils

import (
	"os"

	"github.com/gophoria/gophoria/pkg/ast"
	"github.com/gophoria/gophoria/pkg/lexer"
	"github.com/gophoria/gophoria/pkg/parser"
)

func ParseFile(fileName string) (*ast.Ast, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	lexer := lexer.NewLexer(string(data))
	parser := parser.NewParser(lexer)

	return parser.Parse()
}
