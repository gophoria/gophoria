package ast

import (
	"fmt"

	"github.com/gophoria/gophoria/pkg/lexer"
)

type Identifier struct {
	Token      *lexer.Token
	Identifier string
}

func NewIdentifier(token *lexer.Token) *Identifier {
	i := Identifier{
		Token:      token,
		Identifier: token.Literal,
	}

	return &i
}

func (i *Identifier) String() string {
	return i.Identifier
}

type Value struct {
	Token *lexer.Token
	Value string
}

func NewValue(token *lexer.Token) *Value {
	v := Value{
		Token: token,
		Value: token.Literal,
	}

	return &v
}

func (v *Value) String() string {
	return v.Value
}

type AssignItem struct {
	Token      *lexer.Token
	Identifier *Identifier
	Value      *Value
}

func NewAssignItem(token *lexer.Token, identifier Identifier, value Value) *AssignItem {
	a := AssignItem{}

	return &a
}

func (a *AssignItem) String() string {
	return fmt.Sprintf("%s = %s", a.Identifier, a.Value)
}
