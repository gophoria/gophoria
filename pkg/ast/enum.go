package ast

import (
	"strings"

	"github.com/gophoria/gophoria/pkg/lexer"
)

type Enum struct {
	Token *lexer.Token
	Name  *Identifier
	Items []*AssignItem
}

func NewEnum(token *lexer.Token, ident *Identifier) *Enum {
	e := Enum{
		Token: token,
		Name:  ident,
		Items: []*AssignItem{},
	}

	return &e
}

func (e *Enum) String() string {
	var sb strings.Builder
	sb.WriteString("enum ")
	sb.WriteString(e.Name.Identifier)
	sb.WriteString(" {\n")
	for _, item := range e.Items {
		sb.WriteString(item.String())
		sb.WriteString("\n")
	}
	sb.WriteString(" }\n")

	return sb.String()
}

func (e *Enum) AddItem(item *AssignItem) {
	e.Items = append(e.Items, item)
}
