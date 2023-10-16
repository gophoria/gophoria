package ast

import (
	"strings"

	"github.com/gophoria/gophoria/pkg/lexer"
)

type Enum struct {
	Token *lexer.Token
	Name  string
	Items []*AssignItem
}

func NewEnum(token *lexer.Token) *Enum {
	e := Enum{
		Token: token,
		Name:  token.Literal,
		Items: []*AssignItem{},
	}

	return &e
}

func (e *Enum) String() string {
	var sb strings.Builder
	sb.WriteString("enum ")
	sb.WriteString(e.Name)
	sb.WriteString(" {\n")
	for _, item := range e.Items {
		sb.WriteString(item.String())
		sb.WriteString(",\n")
	}
	sb.WriteString(" }\n")

	return sb.String()
}

func (e *Enum) AddItem(item *AssignItem) {
	e.Items = append(e.Items, item)
}
