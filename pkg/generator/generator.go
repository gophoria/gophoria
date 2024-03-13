package generator

import (
	"fmt"
	"io"

	"github.com/gophoria/gophoria/pkg/ast"
)

var generators = map[string]Generator{}

type Generator interface {
	GenerateAll(ast *ast.Ast, writer io.Writer) error
	Generate(ast *ast.Ast, name string, writer io.Writer) error
}

func GetGenerator(name string) (Generator, error) {
	gen, ok := generators[name]
	if !ok {
		return nil, fmt.Errorf("generator %s not found", name)
	}

	return gen, nil
}

func RegisterGenerator(name string, gen Generator) {
	_, ok := generators[name]
	if ok {
		panic(fmt.Sprintf("generator %s already exists", name))
	}

	generators[name] = gen
}
