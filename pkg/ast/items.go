package ast

import (
	"fmt"

	"github.com/gophoria/gophoria/pkg/lexer"
)

type ValueType int

const (
	ValueTypeInt ValueType = iota
	ValueTypeString
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
	Type  ValueType
	Value string
}

func NewValue(token *lexer.Token) *Value {
	valType := ValueTypeString
	if token.Type == lexer.TokenTypeInt {
		valType = ValueTypeInt
	}

	v := Value{
		Token: token,
		Value: token.Literal,
		Type:  valType,
	}

	return &v
}

func (v *Value) String() string {
	if v.Type == ValueTypeString {
		return "\"" + v.Value + "\""
	}

	return v.Value
}

type AssignItem struct {
	Token      *lexer.Token
	Identifier *Identifier
	Value      *Value
}

func NewAssignItem(token *lexer.Token, identifier *Identifier, value *Value) *AssignItem {
	a := AssignItem{
		Token:      token,
		Identifier: identifier,
		Value:      value,
	}

	return &a
}

func (a *AssignItem) String() string {
	return fmt.Sprintf("%s = %s", a.Identifier, a.Value)
}
