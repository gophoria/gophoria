package lexer_test

import (
	"testing"

	"github.com/gophoria/gophoria/pkg/lexer"
)

func TestLexer(t *testing.T) {
	input := `
db {
  provider = "sqlite3"
  url = ":memory:"
}

enum Role {
  admin = "admin"
  user = "user"
}

model User {
  id      string  @id @default(uuid())
  name    string
  surname string
  role    Role
  posts   Post[]
}

model Post {
  id        int       @id @default(autoincrement())
  title     string
  content   string  @nullable
  public    bool    @default(false)
  author    User    @relation(field: authorId, reference: id)
  authorId  int
}`

	expected := []*lexer.Token{
		lexer.NewToken(lexer.TokenTypeDb, "db", 0, 0),
		lexer.NewToken(lexer.TokenTypeLBrace, "{", 0, 0),
		lexer.NewToken(lexer.TokenTypeIdent, "provider", 0, 0),
		lexer.NewToken(lexer.TokenTypeAssign, "=", 0, 0),
		lexer.NewToken(lexer.TokenTypeString, "sqlite3", 0, 0),
		lexer.NewToken(lexer.TokenTypeIdent, "url", 0, 0),
		lexer.NewToken(lexer.TokenTypeAssign, "=", 0, 0),
		lexer.NewToken(lexer.TokenTypeString, ":memory:", 0, 0),
		lexer.NewToken(lexer.TokenTypeRBrace, "}", 0, 0),

		lexer.NewToken(lexer.TokenTypeEnum, "enum", 0, 0),
		lexer.NewToken(lexer.TokenTypeIdent, "Role", 0, 0),
		lexer.NewToken(lexer.TokenTypeLBrace, "{", 0, 0),
		lexer.NewToken(lexer.TokenTypeIdent, "admin", 0, 0),
		lexer.NewToken(lexer.TokenTypeAssign, "=", 0, 0),
		lexer.NewToken(lexer.TokenTypeString, "admin", 0, 0),
		lexer.NewToken(lexer.TokenTypeIdent, "user", 0, 0),
		lexer.NewToken(lexer.TokenTypeAssign, "=", 0, 0),
		lexer.NewToken(lexer.TokenTypeString, "user", 0, 0),
		lexer.NewToken(lexer.TokenTypeRBrace, "}", 0, 0),

		lexer.NewToken(lexer.TokenTypeModel, "model", 0, 0),
		lexer.NewToken(lexer.TokenTypeIdent, "User", 0, 0),
		lexer.NewToken(lexer.TokenTypeLBrace, "{", 0, 0),
		lexer.NewToken(lexer.TokenTypeIdent, "id", 0, 0),
		lexer.NewToken(lexer.TokenTypeString, "string", 0, 0),
		lexer.NewToken(lexer.TokenTypeDecId, "@id", 0, 0),
		lexer.NewToken(lexer.TokenTypeDecDefault, "@default", 0, 0),
		lexer.NewToken(lexer.TokenTypeLParen, "(", 0, 0),
		lexer.NewToken(lexer.TokenTypeIdent, "uuid", 0, 0),
		lexer.NewToken(lexer.TokenTypeLParen, "(", 0, 0),
		lexer.NewToken(lexer.TokenTypeRParen, ")", 0, 0),
		lexer.NewToken(lexer.TokenTypeRParen, ")", 0, 0),
		lexer.NewToken(lexer.TokenTypeIdent, "name", 0, 0),
		lexer.NewToken(lexer.TokenTypeString, "string", 0, 0),
		lexer.NewToken(lexer.TokenTypeIdent, "surname", 0, 0),
		lexer.NewToken(lexer.TokenTypeString, "string", 0, 0),
		lexer.NewToken(lexer.TokenTypeIdent, "role", 0, 0),
		lexer.NewToken(lexer.TokenTypeIdent, "Role", 0, 0),
		lexer.NewToken(lexer.TokenTypeIdent, "posts", 0, 0),
		lexer.NewToken(lexer.TokenTypeIdent, "Post", 0, 0),
		lexer.NewToken(lexer.TokenTypeLSquareBrace, "[", 0, 0),
		lexer.NewToken(lexer.TokenTypeRSquareBrace, "]", 0, 0),
		lexer.NewToken(lexer.TokenTypeRBrace, "}", 0, 0),

		lexer.NewToken(lexer.TokenTypeModel, "model", 0, 0),
		lexer.NewToken(lexer.TokenTypeIdent, "Post", 0, 0),
		lexer.NewToken(lexer.TokenTypeLBrace, "{", 0, 0),
		lexer.NewToken(lexer.TokenTypeIdent, "id", 0, 0),
		lexer.NewToken(lexer.TokenTypeString, "int", 0, 0),
		lexer.NewToken(lexer.TokenTypeDecId, "@id", 0, 0),
		lexer.NewToken(lexer.TokenTypeDecDefault, "@default", 0, 0),
		lexer.NewToken(lexer.TokenTypeLParen, "(", 0, 0),
		lexer.NewToken(lexer.TokenTypeIdent, "autoincrement", 0, 0),
		lexer.NewToken(lexer.TokenTypeLParen, "(", 0, 0),
		lexer.NewToken(lexer.TokenTypeRParen, ")", 0, 0),
		lexer.NewToken(lexer.TokenTypeRParen, ")", 0, 0),
		lexer.NewToken(lexer.TokenTypeIdent, "title", 0, 0),
		lexer.NewToken(lexer.TokenTypeString, "string", 0, 0),
		lexer.NewToken(lexer.TokenTypeIdent, "content", 0, 0),
		lexer.NewToken(lexer.TokenTypeString, "string", 0, 0),
		lexer.NewToken(lexer.TokenTypeDecNullable, "@nullable", 0, 0),
		lexer.NewToken(lexer.TokenTypeIdent, "public", 0, 0),
		lexer.NewToken(lexer.TokenTypeTBool, "bool", 0, 0),
		lexer.NewToken(lexer.TokenTypeDecDefault, "@default", 0, 0),
		lexer.NewToken(lexer.TokenTypeLParen, "(", 0, 0),
		lexer.NewToken(lexer.TokenTypeFalse, "false", 0, 0),
		lexer.NewToken(lexer.TokenTypeRParen, ")", 0, 0),
		lexer.NewToken(lexer.TokenTypeIdent, "author", 0, 0),
		lexer.NewToken(lexer.TokenTypeIdent, "User", 0, 0),
		lexer.NewToken(lexer.TokenTypeDecRelation, "@relation", 0, 0),
		lexer.NewToken(lexer.TokenTypeLParen, "(", 0, 0),
		lexer.NewToken(lexer.TokenTypeIdent, "field", 0, 0),
		lexer.NewToken(lexer.TokenTypeColon, ":", 0, 0),
		lexer.NewToken(lexer.TokenTypeIdent, "authorId", 0, 0),
		lexer.NewToken(lexer.TokenTypeComma, ",", 0, 0),
		lexer.NewToken(lexer.TokenTypeIdent, "reference", 0, 0),
		lexer.NewToken(lexer.TokenTypeColon, ":", 0, 0),
		lexer.NewToken(lexer.TokenTypeIdent, "id", 0, 0),
		lexer.NewToken(lexer.TokenTypeRParen, ")", 0, 0),
		lexer.NewToken(lexer.TokenTypeIdent, "authorId", 0, 0),
		lexer.NewToken(lexer.TokenTypeTInt, "int", 0, 0),
		lexer.NewToken(lexer.TokenTypeRBrace, "}", 0, 0),
	}

	lexer := lexer.NewLexer(input)

	for _, exp := range expected {
		tok := lexer.Next()

		if tok.Type != exp.Type {
			t.Fatalf("expected token type %v but got %v", exp.Type, tok.Type)
		}

		if tok.Literal != exp.Literal {
			t.Fatalf("expected token literal %v but got %v", exp.Literal, tok.Literal)
		}
	}
}
