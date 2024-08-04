package generator_test

import (
	"bytes"
	"testing"

	"github.com/gophoria/gophoria/pkg/generator"
	"github.com/gophoria/gophoria/pkg/lexer"
	"github.com/gophoria/gophoria/pkg/parser"
)

func TestSqlx(t *testing.T) {
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
  id        int     @id @default(autoincrement())
  title     string
  content   string  @nullable
  public    bool    @default(false)
  author    User    @relation(field: authorId, reference: id)
  authorId  int
}`

	expected := `type Role string

const (
  RoleAdmin Role = "admin"
  RoleUser = "user"
)

type User struct {
  Id string
  Name string
  Surname string
  Role Role
  Posts []*Post
}

type Post struct {
  Id int
  Title string
  Content string
  Public bool
  Author *User
  AuthorId int
}

`

	lexer := lexer.NewLexer(input)
	parser := parser.NewParser(lexer)

	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("parser error: %s", err.Error())
	}

	generator := generator.NewSqlxGenerator()

	var buffer bytes.Buffer
	err = generator.GenerateAll(ast, nil)
	if err != nil {
		t.Fatalf("generator error: %s", err.Error())
	}

	if buffer.String() != expected {
		t.Fatalf("Generator output is not correct")
	}
}
