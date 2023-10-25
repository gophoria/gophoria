package ast

import "github.com/gophoria/gophoria/pkg/lexer"

type Model struct {
	Token *lexer.Token
	Name  *Identifier
	Items []*Declaration
}

func NewModel(token *lexer.Token, name *Identifier) *Model {
	m := Model{
		Token: token,
		Name:  name,
	}

	return &m
}

func (m *Model) AddItem(item *Declaration) {
	m.Items = append(m.Items, item)
}

func (m *Model) String() string {
	return ""
}
