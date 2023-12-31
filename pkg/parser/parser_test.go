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

func TestEnum(t *testing.T) {
	input := `
enum Role {
  admin = "admin"
  user = "user"
}`

	expectedItems := map[string]string{
		"admin": "admin",
		"user":  "user",
	}

	lexer := lexer.NewLexer(input)
	parser := parser.NewParser(lexer)

	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("parser error: %s", err.Error())
		return
	}

	if len(ast.Enums) != 1 {
		t.Fatalf("expected 1 enum but found %d", len(ast.Config))
		return
	}

	en := ast.Enums[0]

	if en.Name.Identifier != "Role" {
		t.Fatalf("expected Role but found %s", en.Name)
		return
	}

	if len(en.Items) != 2 {
		t.Fatalf("expected 2 items in Role but found %d", len(en.Items))
		return
	}

	for _, item := range en.Items {
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

func TestSimpleModel(t *testing.T) {
	input := `
model User {
  id      string
  name    string
  surname string
  role    Role
  posts   Post[]
}`

	expectedItems := map[string]string{
		"id":      "string",
		"name":    "string",
		"surname": "string",
		"role":    "Role",
		"posts":   "Post[]",
	}

	lexer := lexer.NewLexer(input)
	parser := parser.NewParser(lexer)

	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("parser error: %s", err.Error())
		return
	}

	if len(ast.Models) != 1 {
		t.Fatalf("expected 1 model but found %d", len(ast.Config))
		return
	}

	model := ast.Models[0]

	if model.Name.Identifier != "User" {
		t.Fatalf("expected User but found %s", model.Name)
		return
	}

	if len(model.Items) != 5 {
		t.Fatalf("expected 5 items in Role but found %d", len(model.Items))
		return
	}

	for _, item := range model.Items {
		val, ok := expectedItems[item.Identifier.Identifier]
		if !ok {
			t.Fatalf("unexpected identifier %s", item.Identifier.Identifier)
			return
		}

		if item.DeclarationType.String() != val {
			t.Fatalf("expected value %s but got %s", val, item.DeclarationType)
			return
		}
	}
}

func TestDecorators(t *testing.T) {
	input := `
model User {
  id      string    @id
  name    string
  surname string
  role    Role      @nullable
  posts   Post[]
}`

	expectedItems := map[string]string{
		"id":      "string",
		"name":    "string",
		"surname": "string",
		"role":    "Role",
		"posts":   "Post[]",
	}

	expectedDecorators := map[string]string{
		"id":   "@id",
		"role": "@nullable",
	}

	lexer := lexer.NewLexer(input)
	parser := parser.NewParser(lexer)

	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("parser error: %s", err.Error())
		return
	}

	if len(ast.Models) != 1 {
		t.Fatalf("expected 1 model but found %d", len(ast.Config))
		return
	}

	model := ast.Models[0]

	if model.Name.Identifier != "User" {
		t.Fatalf("expected User but found %s", model.Name)
		return
	}

	if len(model.Items) != 5 {
		t.Fatalf("expected 5 items in Role but found %d", len(model.Items))
		return
	}

	for _, item := range model.Items {
		val, ok := expectedItems[item.Identifier.Identifier]
		if !ok {
			t.Fatalf("unexpected identifier %s", item.Identifier.Identifier)
			return
		}

		if item.DeclarationType.String() != val {
			t.Fatalf("expected value %s but got %s", val, item.DeclarationType)
			return
		}

		for _, dec := range item.Decorators {
			if dec.String() != expectedDecorators[item.Identifier.Identifier] {
				t.Fatalf("expected decorator %s but got %s", expectedDecorators[item.Identifier.Identifier], dec)
				return
			}
		}
	}
}

func TestCallableDecorators(t *testing.T) {
	input := `
model Post {
  id        int     @id @default(autoincrement())
  title     string
  content   string  @nullable
  public    bool    @default(false)
  author    User    @relation(field: authorId, reference: id)
  authorId  int
}`

	expectedItems := map[string]string{
		"id":       "int",
		"title":    "string",
		"content":  "string",
		"public":   "bool",
		"author":   "User",
		"authorId": "int",
	}

	expectedDecorators := map[string][]string{
		"id":      {"@id", "@default(autoincrement())"},
		"content": {"@nullable"},
		"public":  {"@default(false)"},
		"author":  {"@relation(field: authorId, reference: id)"},
	}

	lexer := lexer.NewLexer(input)
	parser := parser.NewParser(lexer)

	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("parser error: %s", err.Error())
		return
	}

	if len(ast.Models) != 1 {
		t.Fatalf("expected 1 model but found %d", len(ast.Config))
		return
	}

	model := ast.Models[0]

	if model.Name.Identifier != "Post" {
		t.Fatalf("expected Post but found %s", model.Name)
		return
	}

	if len(model.Items) != 6 {
		t.Fatalf("expected 6 items in Role but found %d", len(model.Items))
		return
	}

	for _, item := range model.Items {
		val, ok := expectedItems[item.Identifier.Identifier]
		if !ok {
			t.Fatalf("unexpected identifier %s", item.Identifier.Identifier)
			return
		}

		if item.DeclarationType.String() != val {
			t.Fatalf("expected value %s but got %s", val, item.DeclarationType)
			return
		}

		for i := range item.Decorators {
			if item.Decorators[i].String() != expectedDecorators[item.Identifier.Identifier][i] {
				t.Fatalf("expected decorator %s but got %s", expectedDecorators[item.Identifier.Identifier][i], item.Decorators[i])
				return
			}
		}
	}
}
