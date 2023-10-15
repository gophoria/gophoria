package parser_test

import (
	"testing"

	"github.com/gophoria/gophoria/pkg/lexer"
	"github.com/gophoria/gophoria/pkg/parser"
)

func TestParserDb(t *testing.T) {
	input := `
db {
  provider = "sqlite3"
  url = ":memory:"
}`

	lexer := lexer.NewLexer(input)
	parser := parser.NewParser(lexer)

	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("parser error: %s", err.Error())
		return
	}

	if len(ast.Config) != 1 {
		t.Fatalf("expected one config but found %d", len(ast.Config))
		return
	}

	cfg := ast.Config[0]

	if cfg.Type != "db" {
		t.Fatalf("expected db config but found %s", cfg.Type)
		return
	}
}
