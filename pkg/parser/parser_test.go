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

	expectedItems := map[string]string{
		"provider": "sqlite3",
		"url":      ":memory:",
	}

	lexer := lexer.NewLexer(input)
	parser := parser.NewParser(lexer)

	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("parser error: %s", err.Error())
		return
	}

	if len(ast.Config) != 1 {
		t.Fatalf("expected 1 config but found %d", len(ast.Config))
		return
	}

	cfg := ast.Config[0]

	if cfg.Type != "db" {
		t.Fatalf("expected db config but found %s", cfg.Type)
		return
	}

	if len(cfg.Items) != 2 {
		t.Fatalf("expected 2 items in db but found %d", len(cfg.Items))
		return
	}

	for _, item := range cfg.Items {
		val, ok := expectedItems[item.Identifier.Identifier]
		if !ok {
			t.Fatalf("unexpected identifier %s", item.Identifier.Identifier)
			return
		}

		if item.Value.Value != val {
			t.Fatalf("expected value %s but got %s", val, item.Value.Value)
			return
		}
	}
}
