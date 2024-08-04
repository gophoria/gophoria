package generator

import (
	"fmt"

	"github.com/gophoria/gophoria/pkg/ast"
)

var generators = map[string]Generator{}

type GeneratorConfig struct {
	Override   bool
	WorkingDir string
}

type Generator interface {
	GenerateAll(ast *ast.Ast, cfg *GeneratorConfig) error
	Generate(ast *ast.Ast, cfg *GeneratorConfig, name string) error
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
